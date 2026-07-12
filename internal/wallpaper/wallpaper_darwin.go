//go:build darwin

// Package wallpaper sets the desktop wallpaper on the host operating system.
package wallpaper

import (
	"context"
	"fmt"
	"os/exec"
)

// Set sets the macOS desktop wallpaper to the image at the given path.
// It drives System Events through osascript, so no extra dependencies are needed.
func Set(ctx context.Context, path string) error {
	script := fmt.Sprintf(
		`tell application "System Events" to set picture of every desktop to %q`,
		path,
	)

	// nolint: gosec
	out, err := exec.CommandContext(ctx, "osascript", "-e", script).CombinedOutput()
	if err != nil {
		return fmt.Errorf("osascript failed: %w: %s", err, out)
	}

	return nil
}
