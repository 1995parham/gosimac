//go:build darwin

// Package wallpaper sets the desktop wallpaper on the host operating system.
package wallpaper

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"

	"howett.net/plist"
)

// storeIndexPath is where WallpaperAgent keeps the wallpaper configuration on
// macOS Sonoma and later. Every space, and every display within a space, has
// its own entry.
const storeIndexPath = "Library/Application Support/com.apple.wallpaper/Store/Index.plist"

// imageChoiceProvider marks a desktop entry as showing a still image, as
// opposed to an aerial or a dynamic system wallpaper.
const imageChoiceProvider = "com.apple.wallpaper.choice.image"

// nullMarker is how the store spells an absent value.
const nullMarker = "$null"

// ErrNoDesktops indicates the wallpaper store held no desktop entries to update,
// which means the store exists but is not in the layout we know how to patch.
var ErrNoDesktops = errors.New("wallpaper store contains no desktop entries")

// imageConfiguration is the payload of a still-image choice. It is stored as a
// binary property list nested inside the outer one.
type imageConfiguration struct {
	Type string `plist:"type"`
	URL  struct {
		Relative string `plist:"relative"`
	} `plist:"url"`
}

// Set sets the macOS desktop wallpaper to the image at the given path.
//
// macOS keeps a separate wallpaper configuration for every space, so driving
// System Events through osascript only repaints the space that happens to be
// active; the others keep whatever they were last set to and drift apart. We
// rewrite every desktop entry in the store instead, then restart WallpaperAgent
// so it picks the change up. On macOS versions without the store we fall back
// to osascript, which is authoritative there.
func Set(ctx context.Context, path string) error {
	abs, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("resolving %s: %w", path, err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("locating home directory: %w", err)
	}

	index := filepath.Join(home, storeIndexPath)

	if _, err := os.Stat(index); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return setViaSystemEvents(ctx, abs)
		}

		return fmt.Errorf("reading wallpaper store: %w", err)
	}

	if err := setViaStore(index, abs); err != nil {
		return err
	}

	return restartWallpaperAgent(ctx)
}

// setViaStore points every desktop entry in the wallpaper store at abs.
func setViaStore(index, abs string) error {
	raw, err := os.ReadFile(index)
	if err != nil {
		return fmt.Errorf("reading wallpaper store: %w", err)
	}

	var store map[string]any
	if _, err := plist.Unmarshal(raw, &store); err != nil {
		return fmt.Errorf("decoding wallpaper store: %w", err)
	}

	choice, err := imageChoice(abs)
	if err != nil {
		return err
	}

	if n := patchDesktops(store, choice); n == 0 {
		return ErrNoDesktops
	}

	patched, err := plist.Marshal(store, plist.BinaryFormat)
	if err != nil {
		return fmt.Errorf("encoding wallpaper store: %w", err)
	}

	// Write through a temporary file next to the original so a failure part way
	// through cannot leave WallpaperAgent with a truncated store.
	tmp, err := os.CreateTemp(filepath.Dir(index), ".Index.plist.*")
	if err != nil {
		return fmt.Errorf("creating temporary wallpaper store: %w", err)
	}

	defer func() {
		_ = os.Remove(tmp.Name())
	}()

	if _, err := tmp.Write(patched); err != nil {
		_ = tmp.Close()

		return fmt.Errorf("writing temporary wallpaper store: %w", err)
	}

	if err := tmp.Close(); err != nil {
		return fmt.Errorf("closing temporary wallpaper store: %w", err)
	}

	if err := os.Rename(tmp.Name(), index); err != nil {
		return fmt.Errorf("replacing wallpaper store: %w", err)
	}

	return nil
}

// imageChoice builds the store entry that shows the image at abs.
func imageChoice(abs string) (map[string]any, error) {
	var cfg imageConfiguration

	cfg.Type = "imageFile"
	// nolint: exhaustruct
	cfg.URL.Relative = (&url.URL{Scheme: "file", Path: abs}).String()

	// The store nests this as an opaque binary property list.
	encoded, err := plist.Marshal(cfg, plist.BinaryFormat)
	if err != nil {
		return nil, fmt.Errorf("encoding image choice: %w", err)
	}

	return map[string]any{
		"Configuration": encoded,
		"Files":         []any{},
		"Provider":      imageChoiceProvider,
	}, nil
}

// patchDesktops walks the store and rewrites the content of every desktop entry
// it finds, returning how many it touched. The store nests desktops under
// per-space and per-display keys that we cannot know ahead of time, so we
// recurse rather than address them directly. Idle (screen saver) entries are
// left alone.
func patchDesktops(node any, choice map[string]any) int {
	patched := 0

	switch typed := node.(type) {
	case map[string]any:
		if desktop, ok := typed["Desktop"].(map[string]any); ok {
			if content, ok := desktop["Content"].(map[string]any); ok {
				content["Choices"] = []any{choice}
				content["Shuffle"] = nullMarker
				content["EncodedOptionValues"] = nullMarker
				patched++
			}
		}

		for _, child := range typed {
			patched += patchDesktops(child, choice)
		}
	case []any:
		for _, child := range typed {
			patched += patchDesktops(child, choice)
		}
	}

	return patched
}

// restartWallpaperAgent makes WallpaperAgent reread the store. launchd brings it
// back automatically, so a missing process is not an error.
func restartWallpaperAgent(ctx context.Context) error {
	// nolint: gosec
	if out, err := exec.CommandContext(ctx, "killall", "WallpaperAgent").CombinedOutput(); err != nil {
		var exit *exec.ExitError
		if !errors.As(err, &exit) {
			return fmt.Errorf("restarting WallpaperAgent: %w: %s", err, out)
		}
	}

	return nil
}

// setViaSystemEvents is the pre-Sonoma path, where System Events is the
// authoritative way to set the wallpaper.
func setViaSystemEvents(ctx context.Context, abs string) error {
	script := fmt.Sprintf(
		`tell application "System Events" to set picture of every desktop to %q`,
		abs,
	)

	// nolint: gosec
	out, err := exec.CommandContext(ctx, "osascript", "-e", script).CombinedOutput()
	if err != nil {
		return fmt.Errorf("osascript failed: %w: %s", err, out)
	}

	return nil
}
