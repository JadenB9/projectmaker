package tui

import (
	"fmt"

	figure "github.com/common-nighthawk/go-figure"
	"github.com/charmbracelet/lipgloss"
)

const Version = "0.5.4"

func welcomeBanner() string {
	artStyle := lipgloss.NewStyle().Bold(true).Foreground(primary)

	fig := figure.NewFigure("PROJECT", "small", true)
	art := fig.String()

	line := "──────────────────────────────"

	return fmt.Sprintf("\n%s%s\n%s\n",
		artStyle.Render(art),
		dimStyle().Render(line),
		dimStyle().Render(fmt.Sprintf("  v%s — Scaffold your next project in seconds", Version)),
	)
}
