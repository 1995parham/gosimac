// Package bing provides a simple way to access bing API.
package bing

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
	resty "gopkg.in/resty.v1"
)

// Source is source implmentation for bing everyday image
type Source struct {
	response Response
	N        int
	Index    int
}

// Init initiates source and return number of available images
func (s *Source) Init() (int, error) {
	resp, err := resty.New().
		SetHostURL("https://www.bing.com").
		R().
		SetResult(&s.response).
		SetQueryParam("format", "js").
		SetQueryParam("mkt", "en-US").
		SetQueryParam("idx", strconv.Itoa(s.Index)).
		SetQueryParam("n", strconv.Itoa(s.N)).
		Get("/HPImageArchive.aspx")
	if err != nil {
		return 0, err
	}

	if resp.StatusCode() != http.StatusOK {
		return 0, fmt.Errorf("invalid response: %s", resp.Status())
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

	logrus.Infof("Getting %s", image.StartDate)

	// nolint: bodyclose
	resp, err := http.Get(fmt.Sprintf("http://www.bing.com/%s", image.URL))
	if err != nil {
		return "", nil, err
	}

	logrus.Infof("%s was gotten", image.StartDate)

	return fmt.Sprintf("%s.jpg", image.FullStartDate), resp.Body, nil
}
