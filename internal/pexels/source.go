// Package pexels provides a simple way to access pexels API.
package pexels

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

// nolint: gosec
const apiKey = "MyMeHXzRhKAPr4VovWihfhq0E28DOug0wqmZhPs7djy0ubRH52hZnwG0"

// ErrRequestFailed indicates a general error in service request.
var ErrRequestFailed = errors.New("request failed")

var ErrInvalidSize = errors.New("invalid size request")

const (
	OriginalSize  = "original"
	Large2xSize   = "large2x"
	LargeSize     = "large"
	MediumSize    = "medium"
	SmallSize     = "small"
	PortraitSize  = "portrait"
	LandscapeSize = "landscape"
	TinySize      = "tiny"
)

// Pexels image provider.
type Pexels struct {
	N           int
	Query       string
	Orientation string
	Path        string
	Prefix      string
	Size        string
	APIKey      string
	Client      *resty.Client
}

func New(count int, query string, orientation string, path string, size string) *Pexels {
	return &Pexels{
		N:           count,
		Query:       query,
		Orientation: orientation,
		Path:        path,
		Size:        size,
		APIKey:      apiKey,
		Prefix:      "pexels",
		Client: resty.New().
			SetBaseURL("https://api.pexels.com").
			SetHeader("Authorization", apiKey).
			SetTimeout(timeout),
	}
}

// Fetch images from pexels based on given criteria.
// nolint: cyclop
func (p *Pexels) Fetch() error {
	photos, err := p.gather()
	if err != nil {
		return fmt.Errorf("gathering information from pexels failed %w", err)
	}

	var wg sync.WaitGroup

	for _, photo := range photos {
		pterm.Info.Printf("Getting %d (%s)\n", photo.ID, photo.Alt)

		var url string

		switch p.Size {
		case OriginalSize:
			url = photo.Src.Original
		case Large2xSize:
			url = photo.Src.Large2x
		case LargeSize:
			url = photo.Src.Large
		case MediumSize:
			url = photo.Src.Medium
		case SmallSize:
			url = photo.Src.Small
		case PortraitSize:
			url = photo.Src.Portrait
		case LandscapeSize:
			url = photo.Src.Landscape
		case TinySize:
			url = photo.Src.Tiny
		}

		if url == "" {
			return ErrInvalidSize
		}

		resp, err := p.Client.R().SetDoNotParseResponse(true).Get(url)
		if err != nil {
			return fmt.Errorf("network failure: %w", err)
		}

		if resp.IsError() {
			pterm.Error.Printf("pexels response code is %d: %s", resp.StatusCode(), resp.String())

			return ErrRequestFailed
		}

		pterm.Success.Printf("%d was gotten\n", photo.ID)

		wg.Add(1)

		go func(name string, content io.ReadCloser) {
			defer wg.Done()

			store.Save(p.Path, p.Prefix, name, content)
		}(strconv.Itoa(photo.ID), resp.Body)
	}

	wg.Wait()

	return nil
}

// gather images urls from pexels based on given criteria.
func (p *Pexels) gather() ([]Photo, error) {
	r := new(Response)

	req := p.Client.R().
		SetResult(r).
		SetQueryParam("per_page", strconv.Itoa(p.N))

	if p.Orientation != "" {
		req = req.SetQueryParam("orientation", p.Orientation)
	}

	var (
		resp *resty.Response
		err  error
	)

	if p.Query != "" {
		resp, err = req.
			SetQueryParam("query", p.Query).
			Get("/v1/search")
	} else {
		resp, err = req.Get("/v1/curated")
	}

	if err != nil {
		return nil, fmt.Errorf("network failure: %w", err)
	}

	if resp.IsError() {
		pterm.Error.Printf("pexels response code is %d: %s", resp.StatusCode(), resp.String())

		return nil, ErrRequestFailed
	}

	return r.Photos, nil
}
