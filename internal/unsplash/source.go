package unsplash

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
			SetTimeout(30 * time.Second),
	}
}

// Fetch images from unsplash based on given critarias.
// nolint: cyclop
func (u *Unsplash) Fetch() error {
	images, err := u.gather()
	if err != nil {
		return fmt.Errorf("gathering information from unsplash failed %w", err)
	}

	var wg sync.WaitGroup

	// unsplash rate limiter is sensitive we reduce the number of goroutines.
	for _, image := range images {
		pterm.Info.Printf("Getting %s (%s)\n", image.ID, image.Description)

		var url string

		switch u.Size {
		case RawSize:
			url = image.URLs.Raw
		case FullSize:
			url = image.URLs.Full
		case RegularSize:
			url = image.URLs.Regular
		case SmallSize:
			url = image.URLs.Small
		case ThumbSize:
			url = image.URLs.Thumb
		}

		if url == "" {
			return ErrInvalidSize
		}

		resp, err := u.Client.R().SetDoNotParseResponse(true).Get(url)
		if err != nil {
			return fmt.Errorf("network failure: %w", err)
		}

		if resp.IsError() {
			pterm.Error.Printf("unsplash response code is %d: %s", resp.StatusCode(), resp.String())

			return ErrRequestFailed
		}

		pterm.Success.Printf("%s was gotten\n", image.ID)

		wg.Add(1)

		go func(name string, content io.ReadCloser) {
			defer wg.Done()

			store.Save(u.Path, u.Prefix, name, content)
		}(image.ID, resp.Body)
	}

	wg.Wait()

	return nil
}

// gather images urls from unsplash based on given critarias.
func (u *Unsplash) gather() ([]Image, error) {
	var images []Image

	resp, err := u.Client.R().
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
