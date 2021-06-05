package bing

import (
	"github.com/1995parham/gosimac/bing"
	"github.com/1995parham/gosimac/cmd/common"
	"github.com/spf13/cobra"
)

const flagIndex = "index"

// Register registers bing command.
func Register(root *cobra.Command) {
	// nolint: exhaustivestruct
	cmd := &cobra.Command{
		Use:     "bing",
		Aliases: []string{"b"},
		Short:   "fetches images from https://bing.com",

		RunE: func(cmd *cobra.Command, args []string) error {
			n, err := cmd.Flags().GetInt(common.FlagCount)
			if err != nil {
				return err
			}

			i, err := cmd.Flags().GetInt(flagIndex)
			if err != nil {
				return err
			}

			b := &bing.Source{
				N:     n,
				Index: i,
			}

			return common.Run(b, cmd)
		},
	}

	cmd.Flags().IntP(flagIndex, "i", 0, "Index of the first image to fetch")
	root.AddCommand(cmd)
}
