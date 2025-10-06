package pexels

import (
	"fmt"

	"github.com/1995parham/gosimac/internal/pexels"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

const (
	flagQuery       = "query"
	flagOrientation = "orientation"
	flagCount       = "number"
	flagSize        = "size"

	// DefaultCount is a default number of fetching images from sources.
	defaultCount = 10
)

// Register registers pexels command.
// nolint: funlen
func Register(root *cobra.Command, path string) {
	// nolint: exhaustruct
	cmd := &cobra.Command{
		Use:     "pexels",
		Aliases: []string{"p"},
		Short:   "fetches images from https://pexels.com",

		RunE: func(cmd *cobra.Command, _ []string) error {
			n, err := cmd.Flags().GetInt(flagCount)
			if err != nil {
				return fmt.Errorf("count flag parse failed: %w", err)
			}

			pterm.Info.Printf("count: %d\n", n)

			q, err := cmd.Flags().GetString(flagQuery)
			if err != nil {
				return fmt.Errorf("query flag parse failed: %w", err)
			}

			pterm.Info.Printf("query: %s\n", q)

			o, err := cmd.Flags().GetString(flagOrientation)
			if err != nil {
				return fmt.Errorf("orientation flag parse failed: %w", err)
			}

			pterm.Info.Printf("orientation: %s\n", o)

			s, err := cmd.Flags().GetString(flagSize)
			if err != nil {
				return fmt.Errorf("size flag parse failed: %w", err)
			}

			pterm.Info.Printf("size: %s\n", s)

			p := pexels.New(n, q, o, path, s)

			if err := p.Fetch(); err != nil {
				return fmt.Errorf("pexels fetch failed %w", err)
			}

			return nil
		},
	}

	cmd.Flags().StringP(flagQuery, "q", "", "Limit selection to photos matching a search term.")
	cmd.Flags().StringP(
		flagOrientation,
		"o",
		"landscape",
		"Filter search results by photo orientation, Valid values are landscape, portrait, and square.",
	)
	cmd.Flags().StringP(
		flagSize,
		"s",
		"large",
		"Image size on pexels: original, large2x, large, medium, small, portrait, landscape, tiny",
	)
	cmd.Flags().IntP(flagCount, "n", defaultCount, "The number of photos to return")
	root.AddCommand(cmd)
}
