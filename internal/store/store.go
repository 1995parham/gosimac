package store

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/pterm/pterm"
)

// ErrAlreadyExists indicates the file already exists.
var ErrAlreadyExists = errors.New("file already exists")

// Save stores image content to a file with the given prefix and name.
func Save(basePath, prefix, name string, content io.ReadCloser) error {
	defer func() {
		if err := content.Close(); err != nil {
			pterm.Error.Printf("(*io.ReadCloser).Close: %v", err)
		}
	}()

	filePath := path.Join(
		basePath,
		fmt.Sprintf("%s-%s.jpg", prefix, name),
	)

	if _, err := os.Stat(filePath); err == nil {
		pterm.Warning.Printf("%s already exists\n", filePath)

		return ErrAlreadyExists
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("os.Create: %w", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			pterm.Error.Printf("(*os.File).Close: %v", err)
		}
	}()

	bytes, err := io.Copy(file, content)
	if err != nil {
		return fmt.Errorf("io.Copy (%d bytes): %w", bytes, err)
	}

	return nil
}
