// Package source defines the common interface and shared logic for wallpaper providers.
package source

import (
	"context"
	"errors"
	"fmt"
	"math/rand/v2"

	"github.com/1995parham/gosimac/internal/store"
	"github.com/1995parham/gosimac/internal/wallpaper"
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
// When set is true, one of the downloaded images (chosen at random) is applied
// as the desktop wallpaper once all downloads have finished.
func Download(ctx context.Context, client *resty.Client, path, prefix string, images []Image, set bool) error {
	group, groupCtx := errgroup.WithContext(ctx)

	// saved is written by index so each goroutine touches a distinct slot,
	// avoiding a mutex. Empty slots are skipped when picking a wallpaper.
	saved := make([]string, len(images))

	for i, img := range images {
		pterm.Info.Printf("Getting %s\n", img.Name)

		resp, err := client.R().SetContext(groupCtx).SetResponseDoNotParse(true).Get(img.URL)
		if err != nil {
			return fmt.Errorf("download %s: %w", img.Name, err)
		}

		if resp.IsStatusFailure() {
			return &DownloadFailedError{Name: img.Name, StatusCode: resp.StatusCode()}
		}

		pterm.Success.Printf("%s downloaded\n", img.Name)

		group.Go(func() error {
			filePath, err := store.Save(path, prefix, img.Name, resp.Body)
			if err != nil && !errors.Is(err, store.ErrAlreadyExists) {
				return fmt.Errorf("store save failed: %w", err)
			}

			saved[i] = filePath

			return nil
		})
	}

	if err := group.Wait(); err != nil {
		return fmt.Errorf("saving images: %w", err)
	}

	if set {
		return setWallpaper(ctx, saved)
	}

	return nil
}

// setWallpaper applies a random downloaded image as the desktop wallpaper.
func setWallpaper(ctx context.Context, paths []string) error {
	available := make([]string, 0, len(paths))

	for _, p := range paths {
		if p != "" {
			available = append(available, p)
		}
	}

	if len(available) == 0 {
		pterm.Warning.Println("no images available to set as wallpaper")

		return nil
	}

	choice := available[rand.IntN(len(available))]

	pterm.Info.Printf("Setting wallpaper to %s\n", choice)

	if err := wallpaper.Set(ctx, choice); err != nil {
		return fmt.Errorf("set wallpaper failed: %w", err)
	}

	pterm.Success.Println("wallpaper set")

	return nil
}
