// Package bing provides a simple way to access bing API.
package bing

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/1995parham/gosimac/internal/source"
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
func (b *Bing) Fetch(ctx context.Context) error {
	r, err := b.gather(ctx)
	if err != nil {
		return fmt.Errorf("gathering information from bing failed %w", err)
	}

	images := make([]source.Image, 0, len(r.Images))
	for _, img := range r.Images {
		images = append(images, source.Image{
			Name: img.StartDate,
			URL:  "https://www.bing.com" + img.URL,
		})
	}

	if err := source.Download(ctx, b.Client, b.Path, b.Prefix, images); err != nil {
		return fmt.Errorf("bing download failed: %w", err)
	}

	return nil
}

// gather fetches image metadata from bing.
func (b *Bing) gather(ctx context.Context) (*Response, error) {
	r := new(Response)

	resp, err := b.Client.
		R().
		SetContext(ctx).
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
