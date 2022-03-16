package main

import (
	"runtime/debug"

	"github.com/1995parham/gosimac/cmd"
	"github.com/pterm/pterm"
)

func main() {
	if err := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("Go", pterm.NewStyle(pterm.FgCyan)),
		pterm.NewLettersFromStringWithStyle("Si", pterm.NewStyle(pterm.FgLightMagenta)),
		pterm.NewLettersFromStringWithStyle("Mac", pterm.NewStyle(pterm.FgLightRed)),
	).Render(); err != nil {
		_ = err
	}

	// nolint: varnamelen
	if bi, ok := debug.ReadBuildInfo(); ok {
		vcsReversion := ""
		vcsTime := ""

		for _, value := range bi.Settings {
			switch value.Key {
			case "vcs.revision":
				vcsReversion = value.Value
			case "vcs.time":
				vcsTime = value.Value
			}
		}

		pterm.Description.Printf("gosimac %s %s %s\n", bi.Main.Version, vcsReversion, vcsTime)
	}

	cmd.Execute()
}
