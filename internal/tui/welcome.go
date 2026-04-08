package tui

import (
	"fmt"
	"strings"

	figure "github.com/common-nighthawk/go-figure"
	"github.com/charmbracelet/lipgloss"
)

const Version = "0.5.5"

func welcomeBanner() string {
	colorStyle := lipgloss.NewStyle().Bold(true).Foreground(primary)

	fig := figure.NewFigure("PROJECT", "doom", true)
	art := fig.String()

	// Color each line individually to avoid lipgloss padding issues
	var lines []string
	lines = append(lines, "") // leading blank line
	for _, l := range strings.Split(art, "\n") {
		if strings.TrimSpace(l) != "" {
			lines = append(lines, colorStyle.Render(l))
		}
	}

	// Measure the width of the first art line for the separator
	artLines := strings.Split(art, "\n")
	width := 0
	for _, l := range artLines {
		if len(l) > width {
			width = len(l)
		}
	}

	lines = append(lines, dimStyle().Render(strings.Repeat("─", width)))
	lines = append(lines, dimStyle().Render(fmt.Sprintf("  v%s — Scaffold your next project in seconds", Version)))
	lines = append(lines, "") // trailing blank line

	return strings.Join(lines, "\n")
}
