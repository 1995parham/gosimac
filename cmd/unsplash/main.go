package unsplash

import (
	"fmt"

	"github.com/1995parham/gosimac/cmd/common"
	"github.com/1995parham/gosimac/unsplash"
	"github.com/spf13/cobra"
)

const (
	flagQuery       = "query"
	flagOrientation = "orientation"
)

// Register registers unsplash command.
func Register(root *cobra.Command) {
	// nolint: exhaustivestruct
	cmd := &cobra.Command{
		Use:     "unsplash",
		Aliases: []string{"u"},
		Short:   "fetches images from https://unsplash.org",

		RunE: func(cmd *cobra.Command, args []string) error {
			n, err := cmd.Flags().GetInt(common.FlagCount)
			if err != nil {
				return fmt.Errorf("count flag parse failed: %w", err)
			}

			q, err := cmd.Flags().GetString(flagQuery)
			if err != nil {
				return fmt.Errorf("query flag parse failed: %w", err)
			}

			o, err := cmd.Flags().GetString(flagOrientation)
			if err != nil {
				return fmt.Errorf("orientation flag parse failed: %w", err)
			}

			s := &unsplash.Source{
				N:           n,
				Query:       q,
				Orientation: o,
			}

			return common.Run(s, cmd)
		},
	}

	cmd.Flags().StringP(flagQuery, "q", "", "Limit selection to photos matching a search term.")
	cmd.Flags().StringP(
		flagOrientation,
		"o",
		"landscape",
		"Filter search results by photo orientation, Valid values are landscape, portrait, and squarish.",
	)
	root.AddCommand(cmd)
}
