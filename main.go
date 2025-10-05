package main

import (
	"github.com/1995parham/gosimac/internal/cmd"
	"github.com/carlmjohnson/versioninfo"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

func main() {
	if err := pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("Go", pterm.NewStyle(pterm.FgCyan)),
		putils.LettersFromStringWithStyle("Si", pterm.NewStyle(pterm.FgLightMagenta)),
		putils.LettersFromStringWithStyle("Mac", pterm.NewStyle(pterm.FgLightRed)),
	).Render(); err != nil {
		pterm.Error.Printf("failed to render banner: %v\n", err)
	}

	pterm.Description.Printf("gosimac %s\n", versioninfo.Short())

	cmd.Execute()
}
