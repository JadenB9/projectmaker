package main

import (
	"fmt"
	"os"

	"github.com/JadenB9/projectmaker/internal/scaffold"
	"github.com/JadenB9/projectmaker/internal/services"
	"github.com/JadenB9/projectmaker/internal/spec"
	"github.com/JadenB9/projectmaker/internal/tui"
)

func main() {
	cfg, err := tui.Run()
	if err != nil {
		tui.PrintError(err.Error())
		os.Exit(1)
	}
	if cfg == nil {
		tui.PrintCancelled()
		os.Exit(0)
	}

	// Detect available CLIs
	clis := services.DetectCLIs()

	// Run scaffold
	tui.PrintDivider()
	tui.PrintStep("Scaffolding project...")
	fmt.Println()

	result, err := scaffold.Run(cfg, clis)
	if err != nil {
		tui.PrintError(err.Error())
		os.Exit(1)
	}

	// Generate PROJECT_SPEC.md
	if err := spec.Generate(cfg, result); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to generate PROJECT_SPEC.md: %v\n", err)
	}

	tui.PrintResults(result, cfg.Name)
}
