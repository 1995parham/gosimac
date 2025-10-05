// Package bing provides a simple way to access bing API.
package bing

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"sync"
	"time"

	"github.com/1995parham/gosimac/internal/store"
	"github.com/pterm/pterm"
	"resty.dev/v3"
)

const timeout = 30 * time.Second

// ErrRequestFailed indicates a general error in service request.
var ErrRequestFailed = errors.New("request failed")

// Bing image provider based on daily wallpaper on bing.
type Bing struct {
	N      int
	Index  int
	Path   string
	Prefix string
	Client *resty.Client
}

func New(count int, index int, path string) *Bing {
	return &Bing{
		N:      count,
		Path:   path,
		Index:  index,
		Prefix: "bing",
		Client: resty.New().
			SetBaseURL("https://www.bing.com").
			SetTimeout(timeout),
	}
}

// Fetch images from bing daily wallpaper.
func (b *Bing) Fetch() error {
	r, err := b.gather()
	if err != nil {
		return fmt.Errorf("gathering information from bing failed %w", err)
	}

	var wg sync.WaitGroup

	for _, image := range r.Images {
		pterm.Info.Printf("Getting %s\n", image.StartDate)

		resp, err := b.Client.R().SetDoNotParseResponse(true).Get("http://www.bing.com/" + image.URL)
		if err != nil {
			return fmt.Errorf("network failure: %w", err)
		}

		if resp.IsError() {
			pterm.Error.Printf("bing response code is %d: %s", resp.StatusCode(), resp.String())

			return ErrRequestFailed
		}

		pterm.Success.Printf("%s was gotten\n", image.StartDate)

		wg.Add(1)

		go func(name string, content io.ReadCloser) {
			defer wg.Done()

			store.Save(b.Path, b.Prefix, name, content)
		}(image.StartDate, resp.Body)
	}

	wg.Wait()

	return nil
}

// Init initiates source and return number of available images.
func (b *Bing) gather() (*Response, error) {
	r := new(Response)

	resp, err := b.Client.
		R().
		SetResult(r).
		SetQueryParam("format", "js").
		SetQueryParam("mkt", "en-US").
		SetQueryParam("idx", strconv.Itoa(b.Index)).
		SetQueryParam("n", strconv.Itoa(b.N)).
		Get("/HPImageArchive.aspx")
	if err != nil {
		return nil, fmt.Errorf("network failure: %w", err)
	}

	if resp.IsError() {
		pterm.Error.Printf("bing response code is %d: %s", resp.StatusCode(), resp.String())

		return nil, ErrRequestFailed
	}

	return r, nil
}
