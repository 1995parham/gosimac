package unsplash

import (
	"fmt"

	"github.com/1995parham/gosimac/internal/unsplash"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

const (
	flagQuery       = "query"
	flagOrientation = "orientation"
	flagCount       = "number"
	flagToken       = "token"
	flagSize        = "size"

	// DefaultCount is a default number of fetching images from sources.
	defaultCount = 10

	// nolint: gosec
	gosimacToken = "4c483af1b27cf8d55fc29504bc48e3755e47eb7a3dd3a320e92b23fc4e5aa1b8"
)

// Register registers unsplash command.
// nolint: funlen
func Register(root *cobra.Command, path string) {
	// nolint: exhaustruct
	cmd := &cobra.Command{
		Use:     "unsplash",
		Aliases: []string{"u"},
		Short:   "fetches images from https://unsplash.org",

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

			t, err := cmd.Flags().GetString(flagToken)
			if err != nil {
				return fmt.Errorf("token flag parse failed: %w", err)
			}

			pterm.Info.Printf("token: %s\n", t)

			s, err := cmd.Flags().GetString(flagSize)
			if err != nil {
				return fmt.Errorf("size flag parse failed: %w", err)
			}

			pterm.Info.Printf("size: %s\n", s)

			u := unsplash.New(n, q, o, t, path, s)

			if err := u.Fetch(); err != nil {
				return fmt.Errorf("unsplash fetch failed %w", err)
			}

			return nil
		},
	}

	cmd.Flags().StringP(flagQuery, "q", "", "Limit selection to photos matching a search term.")
	cmd.Flags().StringP(
		flagOrientation,
		"o",
		"landscape",
		"Filter search results by photo orientation, Valid values are landscape, portrait, and squarish.",
	)
	cmd.Flags().StringP(flagSize, "s", "full", "Image size on unsplash: small, regular, full and raw")
	cmd.Flags().IntP(flagCount, "n", defaultCount, "The number of photos to return")
	cmd.Flags().StringP(flagToken, "t", gosimacToken, "The unplash api token")
	root.AddCommand(cmd)
}
