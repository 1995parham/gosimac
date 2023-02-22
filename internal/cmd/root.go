package cmd

import (
	"log"
	"os"
	"path"

	"github.com/1995parham/gosimac/internal/cmd/bing"
	"github.com/1995parham/gosimac/internal/cmd/unsplash"
	"github.com/adrg/xdg"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// ExitFailure status code.
const (
	ExitFailure = 1
	// DirectoryPermission used for creating GoSiMac Directory.
	// nolint: gofumpt
	DirectoryPermission os.FileMode = 0755
)

// DefaultPath is a default path for storing the wallpapers.
func DefaultPath() string {

	p := path.Join(xdg.UserDirs.Pictures, "GoSiMac")
	if _, err := os.Stat("/proc/sys/fs/binfmt_misc/WSLInterop"); err == nil { //running on wsl (windows subsystem for linux)
		log.Println("Program is running in a WSL environment")
		p = path.Join(xdg.Home, "Pictures", "GoSiMac")
	}
	if _, err := os.Stat(p); err != nil {
		if err := os.MkdirAll(p, DirectoryPermission); err != nil {
			log.Fatalf("os.Mkdir: %v", err)
		}
	}

	return p
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// nolint: exhaustruct
	root := &cobra.Command{
		Use:   "GoSiMac",
		Short: "Fetch the wallpaper from Bings, Unsplash...",
	}

	var path string

	root.PersistentFlags().StringVarP(&path, "path", "p", DefaultPath(), "A path to where photos are stored")

	unsplash.Register(root, path)
	bing.Register(root, path)

	if err := root.Execute(); err != nil {
		pterm.Error.Println(err.Error())
		os.Exit(ExitFailure)
	}
}
