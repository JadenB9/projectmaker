package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/JadenB9/projectmaker/internal/remove"
	"github.com/JadenB9/projectmaker/internal/scaffold"
	"github.com/JadenB9/projectmaker/internal/services"
	"github.com/JadenB9/projectmaker/internal/spec"
	"github.com/JadenB9/projectmaker/internal/tui"
)

func main() {
	// Check for subcommands
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "remove", "rm", "delete":
			remove.Run()
			return
		case "help", "--help", "-h":
			printHelp()
			return
		case "version", "--version", "-v":
			fmt.Println("projectmaker v" + tui.Version)
			return
		}
	}

	// Default: run the create wizard
	runCreate()
}

func printHelp() {
	fmt.Println()
	fmt.Println(tui.DimText("  Usage:"))
	fmt.Println("    project              Create a new project")
	fmt.Println("    project remove       Delete a project (local + GitHub)")
	fmt.Println("    project help         Show this help")
	fmt.Println("    project version      Show version")
	fmt.Println()
}

func runCreate() {
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

	// Check authentication for selected services
	tui.PrintDivider()
	fmt.Println()
	checks := services.CheckAuth(clis, cfg.Deployment)
	allReady := tui.PrintAuthChecks(checks)

	// If any service isn't authenticated, offer to log in
	if !allReady {
		for _, check := range checks {
			if !check.Ready && canLogin(check.Service) {
				if tui.AskLogin(check.Service) {
					fmt.Println()
					if err := services.LoginService(check.Service); err != nil {
						tui.PrintError(fmt.Sprintf("Failed to log in to %s: %v", check.Service, err))
					} else {
						fmt.Println(tui.DimText(fmt.Sprintf("  Logged in to %s", check.Service)))
					}
					fmt.Println()
					clis = services.DetectCLIs()
				}
			}
		}
	}

	// Run scaffold with live progress
	tui.PrintDivider()
	fmt.Println()

	result, err := scaffold.Run(cfg, clis, func(step scaffold.StepResult) {
		tui.PrintStepResult(step)
	})
	if err != nil {
		tui.PrintError(err.Error())
		os.Exit(1)
	}

	// Generate PROJECT_SPEC.md
	if err := spec.Generate(cfg, result); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to generate PROJECT_SPEC.md: %v\n", err)
	}

	tui.PrintComplete(result, cfg.Name)

	// Wait for Enter, then clear screen and cd into project
	fmt.Print(tui.DimText("  Press Enter to jump into your project..."))
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	// Clear screen
	fmt.Print("\033[2J\033[H")

	// Print welcome in the new project
	tui.PrintProjectReady(cfg.Name)

	// Exec a new shell in the project directory
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/zsh"
	}
	shellPath, err := exec.LookPath(filepath.Base(shell))
	if err != nil {
		shellPath = shell
	}
	os.Chdir(cfg.ProjectDir)
	syscall.Exec(shellPath, []string{filepath.Base(shell)}, os.Environ())
}

func canLogin(service string) bool {
	return service == "GitHub" || service == "Vercel"
}
