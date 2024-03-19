package bing

import (
	"fmt"

	"github.com/1995parham/gosimac/internal/bing"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

const (
	flagIndex = "index"
	flagCount = "numer"

	// DefaultCount is a default number of fetching images from sources.
	defaultCount = 10
)

// Register registers bing command.
func Register(root *cobra.Command, path string) {
	// nolint: exhaustruct
	cmd := &cobra.Command{
		Use:     "bing",
		Aliases: []string{"b"},
		Short:   "fetches images from https://bing.com",

		RunE: func(cmd *cobra.Command, _ []string) error {
			n, err := cmd.Flags().GetInt(flagCount)
			if err != nil {
				return fmt.Errorf("count flag parse failed: %w", err)
			}
			pterm.Info.Printf("count: %d\n", n)

			i, err := cmd.Flags().GetInt(flagIndex)
			if err != nil {
				return fmt.Errorf("index flag parse failed: %w", err)
			}
			pterm.Info.Printf("index: %d\n", i)

			b := bing.New(n, i, path)

			if err := b.Fetch(); err != nil {
				return fmt.Errorf("bing fetch failed %w", err)
			}

			return nil
		},
	}

	cmd.Flags().IntP(flagIndex, "i", 0, "Index of the first image to fetch")
	cmd.Flags().IntP(flagCount, "n", defaultCount, "The number of photos to return")
	root.AddCommand(cmd)
}
