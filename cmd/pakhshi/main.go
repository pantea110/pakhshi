package main

import (
	"github.com/snapp-incubator/pakhshi/internal/cmd"
	"github.com/pterm/pterm"
)

func main() {
	if err := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("Pa", pterm.NewStyle(pterm.FgCyan)),
		pterm.NewLettersFromStringWithStyle("kh", pterm.NewStyle(pterm.FgLightMagenta)),
		pterm.NewLettersFromStringWithStyle("shi", pterm.NewStyle(pterm.FgLightRed)),
	).Render(); err != nil {
		_ = err
	}

	pterm.Description.Println("One Client to Rule Them All")

	cmd.Execute()
}
