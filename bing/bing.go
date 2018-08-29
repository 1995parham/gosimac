// Package bing provides a simple way to access bing API.
package bing

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Source is source implmentation for bing everyday image
type Source struct {
	response Response
	Idx      int
	N        int
}

// Init initiates source and return number of avaiable images
func (s *Source) Init() (int, error) {
	// Create HTTP GET request
	resp, err := http.Get(
		fmt.Sprintf("http://www.bing.com/HPImageArchive.aspx?format=js&idx=%d&n=%d&mkt=en-US",
			s.Idx, s.N))
	if err != nil {
		return 0, fmt.Errorf("network failure on %s: %v", "http://www.bing.com/hpimagearchive.aspx", err)
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Invalid response: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&s.response); err != nil {
		return 0, fmt.Errorf("decoding json: %v", err)
	}

	if err := resp.Body.Close(); err != nil {
		return 0, fmt.Errorf("(io.Closer).Close: %v", err)
	}

	return len(s.response.Images), nil
}

// Name returns source name
func (s *Source) Name() string {
	return "bing"
}

// Fetch fetches given index from source
func (s *Source) Fetch(index int) (string, io.ReadCloser, error) {
	image := s.response.Images[index]

	logrus.Infof("Getting %s\n", image.StartDate)

	resp, err := http.Get(fmt.Sprintf("http://www.bing.com/%s", image.URL))
	if err != nil {
		return "", nil, err
	}

	logrus.Infof("%s was gotten\n", image.StartDate)

	return fmt.Sprintf("%s.jpg", image.FullStartDate), resp.Body, nil
}
