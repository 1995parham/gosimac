// Package source defines the common interface and shared logic for wallpaper providers.
package source

import (
	"context"
	"errors"
	"fmt"

	"github.com/1995parham/gosimac/internal/store"
	"github.com/pterm/pterm"
	"golang.org/x/sync/errgroup"
	"resty.dev/v3"
)

// DownloadFailedError indicates that downloading an image failed with a non-2xx status.
type DownloadFailedError struct {
	Name       string
	StatusCode int
}

func (e *DownloadFailedError) Error() string {
	return fmt.Sprintf("download %s failed with status %d", e.Name, e.StatusCode)
}

// Image represents a downloadable wallpaper with a name and URL.
type Image struct {
	Name string
	URL  string
}

// Source represents a wallpaper provider that can fetch images.
type Source interface {
	Fetch(ctx context.Context) error
}

// Download fetches images concurrently and saves them to disk.
// It uses errgroup for goroutine management and error propagation.
func Download(ctx context.Context, client *resty.Client, path, prefix string, images []Image) error {
	g, ctx := errgroup.WithContext(ctx)

	for _, img := range images {
		pterm.Info.Printf("Getting %s\n", img.Name)

		resp, err := client.R().SetContext(ctx).SetDoNotParseResponse(true).Get(img.URL)
		if err != nil {
			return fmt.Errorf("download %s: %w", img.Name, err)
		}

		if resp.IsError() {
			return &DownloadFailedError{Name: img.Name, StatusCode: resp.StatusCode()}
		}

		pterm.Success.Printf("%s downloaded\n", img.Name)

		g.Go(func() error {
			err := store.Save(path, prefix, img.Name, resp.Body)
			if errors.Is(err, store.ErrAlreadyExists) {
				return nil
			}

			return fmt.Errorf("store save failed: %w", err)
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("saving images: %w", err)
	}

	return nil
}
