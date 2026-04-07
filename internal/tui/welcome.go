package tui

import "fmt"

func welcomeBanner() string {
	banner := `
                 _           _                 _
  _ __ _ _ ___  (_) ___  __ | |_  _ __   __ _ | | __ ___  _ _
 | '_ \ '_/ _ \ | |/ -_)/ _||  _|| '  \ / _' || |/ // -_)| '_|
 | .__/_| \___/_/ |\___|\__| \__||_|_|_|\__,_||_|\_\\___||_|
 |_|          |__/
`
	return fmt.Sprintf("%s\n%s",
		titleStyle.Render(banner),
		subtitleStyle.Render("  Scaffold your next project in seconds"),
	)
}
