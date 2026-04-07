package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/JadenB9/projectmaker/internal/scaffold"
)

// PrintResults displays the scaffold results with styled status icons.
func PrintResults(result *scaffold.Result, projectName string) {
	fmt.Println()
	fmt.Println(titleStyle.Render("  Setup Complete"))
	fmt.Println()

	doneIcon := successStyle.Render("[done]  ")
	manualIcon := accentStyle.Render("[todo]  ")
	skipIcon := lipgloss.NewStyle().Foreground(dimColor).Render("[skip]  ")

	var manualCount int
	for _, step := range result.Steps {
		var icon string
		switch step.Status {
		case "done":
			icon = doneIcon
		case "manual":
			icon = manualIcon
			manualCount++
		default:
			icon = skipIcon
		}

		line := fmt.Sprintf("  %s%s", icon, step.Name)
		fmt.Println(line)

		if step.Message != "" && step.Status == "manual" {
			fmt.Println(lipgloss.NewStyle().Foreground(dimColor).PaddingLeft(12).Render(step.Message))
		}
	}

	fmt.Println()

	if manualCount > 0 {
		fmt.Println(accentStyle.Render(fmt.Sprintf("  %d step(s) need manual action — see PROJECT_SPEC.md", manualCount)))
		fmt.Println()
	}

	fmt.Println(successStyle.Render("  PROJECT_SPEC.md generated successfully"))
	fmt.Println()

	nextSteps := []string{
		"cd " + projectName,
		"cat PROJECT_SPEC.md",
	}

	fmt.Println(lipgloss.NewStyle().Foreground(dimColor).Render("  Next:"))
	for _, s := range nextSteps {
		fmt.Println(lipgloss.NewStyle().Foreground(textColor).PaddingLeft(4).Render("$ " + s))
	}
	fmt.Println()
}

// PrintError displays a styled error message.
func PrintError(msg string) {
	fmt.Println()
	fmt.Println(errorStyle.Render("  Error: " + msg))
	fmt.Println()
}

// PrintCancelled displays a cancellation message.
func PrintCancelled() {
	fmt.Println()
	fmt.Println(lipgloss.NewStyle().Foreground(dimColor).Render("  Project creation cancelled."))
	fmt.Println()
}

// PrintStep displays a single step during execution.
func PrintStep(name string) {
	fmt.Println(lipgloss.NewStyle().Foreground(primary).Render("  -> " + name))
}

// PrintDivider prints a subtle divider line.
func PrintDivider() {
	fmt.Println(lipgloss.NewStyle().Foreground(dimColor).Render("  " + strings.Repeat("─", 50)))
}
