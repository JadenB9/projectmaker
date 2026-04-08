package tui

import "fmt"

const Version = "0.5.2"

func welcomeBanner() string {
	art := ` ___  ___  ___    _ ___ ___ _____
| _ \| _ \/ _ \  | | __/ __|_   _|
|  _/|   / (_) | | | _| (__  | |
|_|  |_|_\___/\_/ |___|\___| |_|`

	line := "─────────────────────────────────"

	return fmt.Sprintf("%s\n%s\n%s\n",
		titleStyle.Render(art),
		dimStyle().Render(line),
		dimStyle().Render(fmt.Sprintf("  v%s — Scaffold your next project in seconds", Version)),
	)
}
