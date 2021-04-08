package main

import (
	"github.com/1995parham/gosimac/cmd"
	"github.com/pterm/pterm"
)

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	if err := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("Go", pterm.NewStyle(pterm.FgCyan)),
		pterm.NewLettersFromStringWithStyle("Si", pterm.NewStyle(pterm.FgLightMagenta)),
		pterm.NewLettersFromStringWithStyle("Mac", pterm.NewStyle(pterm.FgLightRed)),
	).Render(); err != nil {
		_ = err
	}

	pterm.Description.Printf("gosimac %s, commit %s, built at %s\n", version, commit, date)

	cmd.Execute()
}
