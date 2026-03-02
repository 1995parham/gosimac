package store

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/pterm/pterm"
)

var (
	// ErrAlreadyExists indicates the file already exists.
	ErrAlreadyExists = errors.New("file already exists")
	// ErrFileCreate indicates failure to create the destination file.
	ErrFileCreate = errors.New("file creation failed")
	// ErrFileCopy indicates failure to copy content into the file.
	ErrFileCopy = errors.New("file copy failed")
)

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
		return fmt.Errorf("%w: %w", ErrFileCreate, err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			pterm.Error.Printf("(*os.File).Close: %v", err)
		}
	}()

	if _, err := io.Copy(file, content); err != nil {
		return fmt.Errorf("%w: %w", ErrFileCopy, err)
	}

	return nil
}
