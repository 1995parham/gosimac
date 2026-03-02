package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Build-time variables set via ldflags.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func Version() string {
	return fmt.Sprintf("%s (%s) built on %s", version, commit, date)
}

func registerVersion(root *cobra.Command) {
	// nolint: exhaustruct
	root.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println(Version())
		},
	})
}
