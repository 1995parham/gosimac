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

	"github.com/pterm/pterm"
)

// App is a GoSiMac application. It contains all gosimac functionality.
// it fetchs background from given source and store them in the given path.
type App struct {
	source Source
	path   string

	// these streams connect fetch and store stage
	fetchStream chan int
	storeStream chan image

	wait sync.WaitGroup
}

// image contains name and bytes of the fetched image.
type image struct {
	name string
	data io.ReadCloser
}

// NewApp creates new app from given source.
func NewApp(path string, source Source) *App {
	return &App{
		source: source,
		path:   path,

		fetchStream: make(chan int),
		storeStream: make(chan image),
	}
}

// Run application that fetches images and store them.
func (a *App) Run() error {
	// finds number of the images
	n, err := a.source.Init()
	if err != nil {
		return err
	}

	pterm.Warning.Printf("%d Images are available from %s\n", n, a.source.Name())
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

// Wait waits until all images are fetched.
func (a *App) Wait() {
	a.wait.Wait()
}

// fetch runs the fetch stage.
func (a *App) fetch() {
	for index := range a.fetchStream {
		name, data, err := a.source.Fetch(index)
		if err != nil {
			a.wait.Done()

			continue
		}
		a.storeStream <- image{name, data}
	}
}

// store runs the store stage.
func (a *App) store() {
	for image := range a.storeStream {
		func() {
			defer a.wait.Done()

			path := path.Join(
				a.path,
				fmt.Sprintf("%s-%s", a.source.Name(), image.name),
			)

			if _, err := os.Stat(path); err == nil {
				pterm.Warning.Printf("%s is already exists\n", path)

				return
			}

			file, err := os.Create(path)
			if err != nil {
				pterm.Error.Printf("os.Create: %v\n", err)

				return
			}

			bytes, err := io.Copy(file, image.data)
			if err != nil {
				pterm.Error.Printf("io.Copy (%d bytes): %v\n", bytes, err)
			}

			if err := file.Close(); err != nil {
				pterm.Error.Printf("(*os.File).Close: %v", err)
			}

			if err := image.data.Close(); err != nil {
				pterm.Error.Printf("(*io.ReadCloser).Close: %v", err)
			}
		}()
	}
}
