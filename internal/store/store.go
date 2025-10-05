package store

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/pterm/pterm"
)

// Save stores image content to a file with the given prefix and name.
func Save(basePath, prefix, name string, content io.ReadCloser) {
	filePath := path.Join(
		basePath,
		fmt.Sprintf("%s-%s.jpg", prefix, name),
	)

	if _, err := os.Stat(filePath); err == nil {
		pterm.Warning.Printf("%s is already exists\n", filePath)

		return
	}

	file, err := os.Create(filePath)
	if err != nil {
		pterm.Error.Printf("os.Create: %v\n", err)

		return
	}

	defer func() {
		if err := file.Close(); err != nil {
			pterm.Error.Printf("(*os.File).Close: %v", err)
		}
	}()

	defer func() {
		if err := content.Close(); err != nil {
			pterm.Error.Printf("(*io.ReadCloser).Close: %v", err)
		}
	}()

	bytes, err := io.Copy(file, content)
	if err != nil {
		pterm.Error.Printf("io.Copy (%d bytes): %v\n", bytes, err)
	}
}
