/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 30-08-2018
 * |
 * | File Name:     app.go
 * +===============================================
 */

package core

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"sync"

	"github.com/sirupsen/logrus"
)

// App is a GoSiMac application. It contains all gosimac
// functionality. it fetchs background from given source and store them in files.
type App struct {
	source Source
	path   string

	fetchStream chan int
	storeStream chan image

	wait sync.WaitGroup
}

type image struct {
	name string
	data io.ReadCloser
}

// NewApp creates new app from given source
func NewApp(path string, source Source) *App {
	return &App{
		source: source,
		path:   path,

		fetchStream: make(chan int),
		storeStream: make(chan image),
	}
}

// Run application that fetches images and store them
func (a *App) Run() error {
	n, err := a.source.Init()
	if err != nil {
		return err
	}
	a.wait.Add(n)

	go func() {
		for i := 0; i < n; i++ {
			a.fetchStream <- i
		}
	}()

	for i := 0; i < runtime.NumCPU(); i++ {
		go a.fetch()
		go a.store()
	}

	return nil
}

// Wait waits until all images are fetched
func (a *App) Wait() {
	a.wait.Wait()
}

func (a *App) fetch() {
	logrus.Infof("Fetch from %s", a.source.Name())
	for index := range a.fetchStream {
		name, data, err := a.source.Fetch(index)
		if err != nil {
			a.wait.Done()
			continue
		}
		a.storeStream <- image{name, data}
	}
}

func (a *App) store() {
	logrus.Infof("Store from %s", a.source.Name())
	for image := range a.storeStream {
		path := path.Join(
			fmt.Sprintf("%s", a.path),
			fmt.Sprintf("%s-%s", a.source.Name(), image.name),
		)

		if _, err := os.Stat(path); err == nil {
			logrus.Infof("%s is already exists\n", path)
			a.wait.Done()
			continue
		}

		file, err := os.Create(path)
		if err != nil {
			logrus.Errorf("os.Create: %v", err)
			a.wait.Done()
			continue
		}

		if _, err := io.Copy(file, image.data); err != nil {
			logrus.Errorf("io.Copy: %v", err)
		}

		if err := file.Close(); err != nil {
			logrus.Errorf("(*os.File).Close: %v", err)
		}

		if err := image.data.Close(); err != nil {
			logrus.Errorf("(*io.ReadCloser).Close: %v", err)
		}
		a.wait.Done()
	}
}
