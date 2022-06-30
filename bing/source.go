// Package bing provides a simple way to access bing API.
package bing

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/pterm/pterm"
)

// ErrRequestFailed indicates a general error in service request.
var ErrRequestFailed = errors.New("request failed")

// Source is source implmentation for bing everyday image.
type Source struct {
	response Response
	N        int
	Index    int
}

// Init initiates source and return number of available images.
func (s *Source) Init() (int, error) {
	resp, err := resty.New().
		SetBaseURL("https://www.bing.com").
		R().
		SetResult(&s.response).
		SetQueryParam("format", "js").
		SetQueryParam("mkt", "en-US").
		SetQueryParam("idx", strconv.Itoa(s.Index)).
		SetQueryParam("n", strconv.Itoa(s.N)).
		Get("/HPImageArchive.aspx")
	if err != nil {
		return 0, fmt.Errorf("network failure: %w", err)
	}

	if !resp.IsSuccess() {
		return 0, ErrRequestFailed
	}

	return len(s.response.Images), nil
}

// Name returns source name.
func (s *Source) Name() string {
	return "bing"
}

// Fetch fetches given index from source.
func (s *Source) Fetch(index int) (string, io.ReadCloser, error) {
	image := s.response.Images[index]

	pterm.Info.Printf("Getting %s", image.StartDate)

	resp, err := resty.New().R().SetDoNotParseResponse(true).Get(fmt.Sprintf("http://www.bing.com/%s", image.URL))
	if err != nil {
		return "", nil, fmt.Errorf("network failure: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return "", nil, ErrRequestFailed
	}

	pterm.Success.Printf("%s was gotten", image.StartDate)

	return fmt.Sprintf("%s.jpg", image.FullStartDate), resp.RawBody(), nil
}
