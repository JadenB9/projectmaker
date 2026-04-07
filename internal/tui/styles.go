package tui

import "github.com/charmbracelet/lipgloss"

var (
	primary   = lipgloss.Color("#7C9CBF") // soft blue
	secondary = lipgloss.Color("#8FAE8B") // muted sage
	accent    = lipgloss.Color("#D4A574") // warm amber
	textColor = lipgloss.Color("#E0E0E0") // light gray
	dimColor  = lipgloss.Color("#888888") // medium gray
	errColor  = lipgloss.Color("#C27171") // soft red

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(primary).
			MarginBottom(1)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(dimColor).
			Italic(true)

	accentStyle = lipgloss.NewStyle().
			Foreground(accent)

	successStyle = lipgloss.NewStyle().
			Foreground(secondary)

	errorStyle = lipgloss.NewStyle().
			Foreground(errColor)

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(primary).
			Padding(1, 2)

	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(primary).
			Padding(1, 2).
			MarginBottom(1)
)
