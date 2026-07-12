//go:build !darwin

package wallpaper

import (
	"context"
	"fmt"
	"runtime"
)

// ErrUnsupported indicates that setting the wallpaper is not implemented for this OS.
var ErrUnsupported = fmt.Errorf("setting the wallpaper is not supported on %s", runtime.GOOS)

// Set is a stub on operating systems where wallpaper setting is not implemented yet.
func Set(_ context.Context, _ string) error {
	return ErrUnsupported
}
