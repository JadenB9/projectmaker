package tui

import "fmt"

const Version = "0.5.1"

func welcomeBanner() string {
	art := "" +
		" ___  ___  ___    _ ___ ___ _____\n" +
		"| _ \\| _ \\/ _ \\  | | __/ __|_   _|\n" +
		"|  _/|   / (_) | | | _| (__  | |  \n" +
		"|_|  |_|_\\\\___/\\_/ |___|\\___| |_|  \n"

	line := "───────────────────────────────"

	return fmt.Sprintf("%s%s\n%s\n",
		titleStyle.Render(art),
		dimStyle().Render(line),
		dimStyle().Render(fmt.Sprintf("  v%s — Scaffold your next project in seconds", Version)),
	)
}
