// Package pexels provides a simple way to access pexels API.
package pexels

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

func New(count int, query string, orientation string, apiKey string, path string, size string) *Pexels {
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
func (p *Pexels) Fetch(ctx context.Context) error {
	photos, err := p.gather(ctx)
	if err != nil {
		return fmt.Errorf("gathering information from pexels failed %w", err)
	}

	images := make([]source.Image, 0, len(photos))

	for _, photo := range photos {
		url, err := p.photoURL(photo)
		if err != nil {
			return err
		}

		images = append(images, source.Image{
			Name: strconv.Itoa(photo.ID),
			URL:  url,
		})
	}

	return source.Download(ctx, p.Client, p.Path, p.Prefix, images)
}

func (p *Pexels) photoURL(photo Photo) (string, error) {
	switch p.Size {
	case OriginalSize:
		return photo.Src.Original, nil
	case Large2xSize:
		return photo.Src.Large2x, nil
	case LargeSize:
		return photo.Src.Large, nil
	case MediumSize:
		return photo.Src.Medium, nil
	case SmallSize:
		return photo.Src.Small, nil
	case PortraitSize:
		return photo.Src.Portrait, nil
	case LandscapeSize:
		return photo.Src.Landscape, nil
	case TinySize:
		return photo.Src.Tiny, nil
	default:
		return "", ErrInvalidSize
	}
}

// gather fetches image metadata from pexels.
func (p *Pexels) gather(ctx context.Context) ([]Photo, error) {
	r := new(Response)

	req := p.Client.R().
		SetContext(ctx).
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
