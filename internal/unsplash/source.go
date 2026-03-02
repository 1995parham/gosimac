package unsplash

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
	RawSize     = "raw"
	FullSize    = "full"
	RegularSize = "regular"
	SmallSize   = "small"
	ThumbSize   = "thumb"
)

// Unsplash image provider.
type Unsplash struct {
	N           int
	Query       string
	Orientation string
	Path        string
	Prefix      string
	Size        string
	Client      *resty.Client
}

func New(count int, query string, orientation string, token string, path string, size string) *Unsplash {
	return &Unsplash{
		N:           count,
		Query:       query,
		Orientation: orientation,
		Path:        path,
		Size:        size,
		Prefix:      "unsplash",
		Client: resty.New().
			SetBaseURL("https://api.unsplash.com").
			SetHeader("Accept-Version", "v1").
			SetHeader("Authorization", "Client-ID "+token).
			SetTimeout(timeout),
	}
}

// Fetch images from unsplash based on given criteria.
func (u *Unsplash) Fetch(ctx context.Context) error {
	apiImages, err := u.gather(ctx)
	if err != nil {
		return fmt.Errorf("gathering information from unsplash failed %w", err)
	}

	images := make([]source.Image, 0, len(apiImages))

	for _, img := range apiImages {
		url, err := u.imageURL(img)
		if err != nil {
			return err
		}

		images = append(images, source.Image{
			Name: img.ID,
			URL:  url,
		})
	}

	return source.Download(ctx, u.Client, u.Path, u.Prefix, images)
}

func (u *Unsplash) imageURL(img Image) (string, error) {
	switch u.Size {
	case RawSize:
		return img.URLs.Raw, nil
	case FullSize:
		return img.URLs.Full, nil
	case RegularSize:
		return img.URLs.Regular, nil
	case SmallSize:
		return img.URLs.Small, nil
	case ThumbSize:
		return img.URLs.Thumb, nil
	default:
		return "", ErrInvalidSize
	}
}

// gather fetches image metadata from unsplash.
func (u *Unsplash) gather(ctx context.Context) ([]Image, error) {
	var images []Image

	resp, err := u.Client.R().
		SetContext(ctx).
		SetResult(&images).
		SetQueryParam("count", strconv.Itoa(u.N)).
		SetQueryParam("orientation", u.Orientation).
		SetQueryParam("query", u.Query).
		Get("/photos/random")
	if err != nil {
		return nil, fmt.Errorf("network failure: %w", err)
	}

	if resp.IsError() {
		pterm.Error.Printf("unsplash response code is %d: %s", resp.StatusCode(), resp.String())

		return nil, ErrRequestFailed
	}

	return images, nil
}
