package bing

import (
	"fmt"

	"github.com/1995parham/gosimac/bing"
	"github.com/1995parham/gosimac/cmd/common"
	"github.com/spf13/cobra"
)

const flagIndex = "index"

// Register registers bing command.
func Register(root *cobra.Command) {
	// nolint: exhaustruct
	cmd := &cobra.Command{
		Use:     "bing",
		Aliases: []string{"b"},
		Short:   "fetches images from https://bing.com",

		RunE: func(cmd *cobra.Command, args []string) error {
			n, err := cmd.Flags().GetInt(common.FlagCount)
			if err != nil {
				return fmt.Errorf("count flag parse failed: %w", err)
			}

			i, err := cmd.Flags().GetInt(flagIndex)
			if err != nil {
				return fmt.Errorf("index flag parse failed: %w", err)
			}

			b := &bing.Source{
				N:     n,
				Index: i,
			}

			if err := common.Run(b, cmd); err != nil {
				return fmt.Errorf("bing engine failed: %w", err)
			}

			return nil
		},
	}

	cmd.Flags().IntP(flagIndex, "i", 0, "Index of the first image to fetch")
	root.AddCommand(cmd)
}
