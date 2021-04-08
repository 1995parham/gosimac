package common

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path"

	"github.com/1995parham/gosimac/core"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

const (
	// FlagPath flag indicates where to store the wallpapers.
	FlagPath = "path"

	// FlagCount flag indicates number of fetching images from source.
	FlagCount = "number"

	// DefaultCount is a default number of fetching images from sources.
	DefaultCount = 10
)

// DefaultPath is a default path for storing the wallpapers.
func DefaultPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("user.Current: %v", err)
	}

	p := path.Join(usr.HomeDir, "Pictures", "GoSiMac")
	if _, err := os.Stat(p); err != nil {
		if err := os.Mkdir(p, 0755); err != nil {
			log.Fatalf("os.Mkdir: %v", err)
		}
	}

	return p
}

// Run runs given source on given path and waits for its results.
func Run(s core.Source, cmd *cobra.Command) error {
	p, err := cmd.Flags().GetString(FlagPath)
	if err != nil {
		return fmt.Errorf("flag path parse failed: %w", err)
	}

	pterm.Description.Printf("Source: %+v\n", s)
	pterm.Description.Printf("Path: %s\n", p)

	a := core.NewApp(p, s)
	if err := a.Run(); err != nil {
		return err
	}

	a.Wait()

	return nil
}
