package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const Version = "0.5.3"

func welcomeBanner() string {
	// Use a plain style with just color — no margins or padding that break alignment
	artStyle := lipgloss.NewStyle().Bold(true).Foreground(primary)

	art := "" +
		" ___  ___  ___    _ ___ ___ _____\n" +
		"| _ \\| _ \\/ _ \\  | | __/ __|_   _|\n" +
		"|  _/|   / (_) | | | _| (__  | |\n" +
		"|_|  |_|_\\___/\\_/ |___|\\___| |_|"

	line := "─────────────────────────────────"

	return fmt.Sprintf("\n%s\n%s\n%s\n",
		artStyle.Render(art),
		dimStyle().Render(line),
		dimStyle().Render(fmt.Sprintf("  v%s — Scaffold your next project in seconds", Version)),
	)
}
