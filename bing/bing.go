// Package bing provides a simple way to access bing API.
package bing

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
)

var w sync.WaitGroup

func getBingImage(path string, image Image) {
	log.Infof("Getting %s\n", image.StartDate)

	defer w.Done()

	if _, err := os.Stat(fmt.Sprintf("%s/%s.jpg", path, image.FullStartDate)); err == nil {
		log.Infof("%s is already exists\n", image.StartDate)
		return
	}

	resp, err := http.Get(fmt.Sprintf("http://www.bing.com/%s", image.URL))
	if err != nil {
		log.Errorf("http.Get: %v", err)
		return
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Errorf("(io.Closer).Close: %v", err)
		}
	}()

	destFile, err := os.Create(fmt.Sprintf("%s/%s.jpg", path, image.FullStartDate))
	if err != nil {
		log.Errorf("os.Create: %v", err)
		return
	}

	defer func() {
		if err := destFile.Close(); err != nil {
			log.Errorf("(*os.File).Close: %v", err)
		}
	}()

	if _, err := io.Copy(destFile, resp.Body); err != nil {
		log.Errorf("io.Copy: %v", err)
	}

	log.Infof("%s was gotten\n", image.StartDate)
}

// GetBingDesktop function gets `n` Bing Wallpaper since `idx` and stores them in `path`.
func GetBingDesktop(path string, idx int, n int) error {
	// Create HTTP GET request
	resp, err := http.Get(
		fmt.Sprintf("http://www.bing.com/hpimagearchive.aspx?format=js&index=%d&number=%d&mkt=en-US",
			idx, n))
	if err != nil {
		return fmt.Errorf("network failure on %s: %v", "http://www.bing.com/hpimagearchive.aspx", err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Errorf("(io.Closer).Close: %v", err)
		}
	}()

	var bingResp Response
	if err := json.NewDecoder(resp.Body).Decode(&bingResp); err != nil {
		return fmt.Errorf("decoding json: %v", err)
	}

	// Create spreate thread for each image
	for _, image := range bingResp.Images {
		w.Add(1)
		go getBingImage(path, image)
	}

	// Waiting for getting all the images
	w.Wait()

	return nil
}
