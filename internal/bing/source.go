// Package bing provides a simple way to access bing API.
package bing

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/pterm/pterm"
)

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
		Client: resty.New().SetBaseURL("https://www.bing.com"),
	}
}

// Fetch images from bing daily wallpaper.
func (b *Bing) Fetch() error {
	r, err := b.gather()
	if err != nil {
		return fmt.Errorf("gatering information from bing failed %w", err)
	}

	for _, image := range r.Images {
		pterm.Info.Printf("Getting %s\n", image.StartDate)

		resp, err := resty.New().R().SetDoNotParseResponse(true).Get("http://www.bing.com/" + image.URL)
		if err != nil {
			return fmt.Errorf("network failure: %w", err)
		}

		if resp.IsError() {
			pterm.Error.Printf("bing response code is %d: %s", resp.StatusCode(), resp.String())

			return ErrRequestFailed
		}

		pterm.Success.Printf("%s was gotten\n", image.StartDate)

		go b.Store(image.StartDate, resp.RawBody())
	}

	return nil
}

func (b *Bing) Store(name string, content io.ReadCloser) {
	path := path.Join(
		b.Path,
		fmt.Sprintf("%s-%s.jpg", b.Prefix, name),
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

	bytes, err := io.Copy(file, content)
	if err != nil {
		pterm.Error.Printf("io.Copy (%d bytes): %v\n", bytes, err)
	}

	if err := file.Close(); err != nil {
		pterm.Error.Printf("(*os.File).Close: %v", err)
	}

	if err := content.Close(); err != nil {
		pterm.Error.Printf("(*io.ReadCloser).Close: %v", err)
	}
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
