package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/JadenB9/projectmaker/internal/scaffold"
)

var (
	doneIcon = successStyle.Render("  [done] ")
	todoIcon = accentStyle.Render("  [todo] ")
	skipIcon = lipgloss.NewStyle().Foreground(dimColor).Render("  [skip] ")
)

// PrintStepResult displays a single step result as it completes (live progress).
func PrintStepResult(step scaffold.StepResult) {
	var icon string
	switch step.Status {
	case "done":
		icon = doneIcon
	case "manual":
		icon = todoIcon
	default:
		icon = skipIcon
	}

	fmt.Printf("%s%s\n", icon, step.Name)

	if step.Message != "" && step.Status == "manual" {
		fmt.Println(lipgloss.NewStyle().Foreground(dimColor).PaddingLeft(12).Render(step.Message))
	}

	// Small pause between steps for visual effect
	time.Sleep(80 * time.Millisecond)
}

// PrintComplete displays the final summary after all steps.
func PrintComplete(result *scaffold.Result, projectName string) {
	fmt.Println()

	var manualCount int
	for _, s := range result.Steps {
		if s.Status == "manual" {
			manualCount++
		}
	}

	if manualCount > 0 {
		fmt.Println(accentStyle.Render(fmt.Sprintf("  %d step(s) need manual action — see PROJECT_SPEC.md", manualCount)))
		fmt.Println()
	}

	fmt.Println(successStyle.Render("  PROJECT_SPEC.md generated"))
	fmt.Println()
}

// PrintProjectReady displays a welcome message after clearing the screen.
func PrintProjectReady(projectName string) {
	fmt.Println()
	fmt.Println(titleStyle.Render("  " + projectName))
	fmt.Println(successStyle.Render("  Project ready. You're now in the project directory."))
	fmt.Println()
	fmt.Println(dimStyle().Render("  Run `cat PROJECT_SPEC.md` for setup details"))
	fmt.Println()
}

// DimText returns styled dim text.
func DimText(s string) string {
	return dimStyle().Render(s)
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
