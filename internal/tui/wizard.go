package tui

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/JadenB9/projectmaker/internal/config"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// toHuhOptions converts a slice of config.Option to huh select options.
func toHuhOptions(opts []config.Option) []huh.Option[string] {
	out := make([]huh.Option[string], len(opts))
	for i, o := range opts {
		out[i] = huh.NewOption(o.Label, o.Value)
	}
	return out
}

// detectPackageManager checks which package managers are installed and returns
// the first available one from the preferred order: bun, pnpm, yarn, npm.
func detectPackageManager() string {
	for _, pm := range []string{"bun", "pnpm", "yarn", "npm"} {
		if _, err := exec.LookPath(pm); err == nil {
			return pm
		}
	}
	return "npm"
}

// buildSummary formats the project config into a styled summary box.
func buildSummary(cfg *config.ProjectConfig) string {
	labelStyle := lipgloss.NewStyle().
		Foreground(primary).
		Bold(true).
		Width(18).
		Align(lipgloss.Right).
		PaddingRight(1)

	valueStyle := lipgloss.NewStyle().
		Foreground(textColor)

	line := func(label, value string) string {
		if value == "" || value == "none" {
			value = dimStyle().Render("--")
		} else {
			value = valueStyle.Render(value)
		}
		return labelStyle.Render(label) + value
	}

	extras := "--"
	if len(cfg.Extras) > 0 {
		extras = strings.Join(cfg.Extras, ", ")
	}

	lines := strings.Join([]string{
		line("Project", cfg.Name),
		line("Stack", cfg.Stack),
		line("Language", cfg.Language),
		line("Framework", cfg.Framework),
		line("Styling", cfg.Styling),
		line("Database", cfg.Database),
		line("Auth", cfg.Auth),
		line("Payments", cfg.Payments),
		line("Email", cfg.Email),
		line("Pkg Manager", cfg.PackageManager),
		line("Deployment", cfg.Deployment),
		line("Extras", extras),
	}, "\n")

	header := headerStyle.Render("Project Summary")
	return header + "\n" + boxStyle.Render(lines)
}

// dimStyle returns a style for dimmed text.
func dimStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(dimColor)
}

// isAbort returns true if the error indicates user pressed Escape/Ctrl+C.
func isAbort(err error) bool {
	return err != nil && errors.Is(err, huh.ErrUserAborted)
}

// Run launches the TUI wizard and returns the completed config.
// Press Escape at any step to go back to the previous step.
func Run() (*config.ProjectConfig, error) {
	fmt.Println(welcomeBanner())
	fmt.Println(dimStyle().Render("  Press Esc to go back at any step\n"))

	cfg := &config.ProjectConfig{}
	cwd, _ := os.Getwd()
	defaultName := "my-project"
	if cwd != "" {
		defaultName = filepath.Base(cwd)
	}
	cfg.Name = defaultName

	var stackChoice string

	// Steps: 0=name, 1=stack, 2=custom(if needed), 3=extras, 4=confirm
	step := 0
	for step >= 0 {
		switch step {
		case 0: // Project name
			err := huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Title("Project Name").
						Placeholder(defaultName).
						Validate(func(s string) error {
							name := s
							if name == "" {
								name = defaultName
							}
							dir := filepath.Join(cwd, name)
							if info, err := os.Stat(dir); err == nil && info.IsDir() {
								return fmt.Errorf("directory %q already exists", name)
							}
							return nil
						}).
						Value(&cfg.Name),
				),
			).Run()
			if isAbort(err) {
				return nil, nil // Can't go back from first step, cancel
			}
			if err != nil {
				return nil, err
			}
			if cfg.Name == "" {
				cfg.Name = defaultName
			}
			step++

		case 1: // Stack selection
			stackOptions := make([]huh.Option[string], 0, len(config.Presets)+1)
			for _, p := range config.Presets {
				label := fmt.Sprintf("%s  %s", p.Name, dimStyle().Render(p.Description))
				stackOptions = append(stackOptions, huh.NewOption(label, p.Name))
			}
			stackOptions = append(stackOptions, huh.NewOption(
				fmt.Sprintf("Custom Stack  %s", dimStyle().Render("pick each layer yourself")),
				"custom",
			))

			err := huh.NewForm(
				huh.NewGroup(
					huh.NewSelect[string]().
						Title("Choose a Stack").
						Options(stackOptions...).
						Value(&stackChoice),
				),
			).Run()
			if isAbort(err) {
				step--
				continue
			}
			if err != nil {
				return nil, err
			}

			if stackChoice == "custom" {
				cfg.Stack = "custom"
				step = 2 // Go to custom flow
			} else {
				applyPreset(cfg, stackChoice)
				step = 3 // Skip custom, go to extras
			}

		case 2: // Custom stack flow (3 groups with back support)
			backFromCustom := runCustomFlow(cfg)
			if backFromCustom {
				step = 1 // Back to stack selection
				continue
			}
			step = 3

		case 3: // Extras
			if cfg.PackageManager == "" {
				cfg.PackageManager = detectPackageManager()
			}

			err := huh.NewForm(
				huh.NewGroup(
					huh.NewMultiSelect[string]().
						Title("Extras").
						Description("Space to toggle, Enter to continue").
						Options(toHuhOptions(config.ExtraOptions)...).
						Value(&cfg.Extras),
				),
			).Run()
			if isAbort(err) {
				if stackChoice == "custom" {
					step = 2
				} else {
					step = 1
				}
				continue
			}
			if err != nil {
				return nil, err
			}
			step++

		case 4: // Confirm
			fmt.Println(buildSummary(cfg))

			var confirmed bool
			err := huh.NewForm(
				huh.NewGroup(
					huh.NewConfirm().
						Title("Create this project?").
						Affirmative("Yes, create it").
						Negative("No, go back").
						Value(&confirmed),
				),
			).Run()
			if isAbort(err) {
				step = 3
				continue
			}
			if err != nil {
				return nil, err
			}

			if !confirmed {
				step = 3 // Go back to extras
				continue
			}

			// Set project directory
			cfg.ProjectDir = filepath.Join(cwd, cfg.Name)
			fmt.Println(successStyle.Render("\n  Let's build it!"))
			return cfg, nil
		}
	}

	return nil, nil
}

// applyPreset copies a preset's config values into cfg.
func applyPreset(cfg *config.ProjectConfig, name string) {
	for _, p := range config.Presets {
		if p.Name == name {
			cfg.Stack = p.Name
			cfg.Language = p.Config.Language
			cfg.Framework = p.Config.Framework
			cfg.Styling = p.Config.Styling
			cfg.Backend = p.Config.Backend
			cfg.Database = p.Config.Database
			cfg.Auth = p.Config.Auth
			cfg.Payments = p.Config.Payments
			cfg.Email = p.Config.Email
			cfg.PackageManager = p.Config.PackageManager
			cfg.Deployment = p.Config.Deployment
			return
		}
	}
}

// runCustomFlow walks through each config layer with individual steps.
// Returns true if user wants to go back to stack selection.
func runCustomFlow(cfg *config.ProjectConfig) bool {
	type customStep struct {
		title string
		opts  []config.Option
		value *string
	}

	steps := []customStep{
		{"Language", config.Languages, &cfg.Language},
		{"Framework", config.Frameworks, &cfg.Framework},
		{"Styling", config.Styling, &cfg.Styling},
		{"Database / ORM", config.Databases, &cfg.Database},
		{"Authentication", config.AuthProviders, &cfg.Auth},
		{"Payments", config.PaymentProviders, &cfg.Payments},
		{"Email Provider", config.EmailProviders, &cfg.Email},
		{"Package Manager", config.PackageManagers, &cfg.PackageManager},
		{"Deployment Target", config.DeploymentTargets, &cfg.Deployment},
	}

	i := 0
	for i < len(steps) {
		s := steps[i]
		err := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title(s.title).
					Options(toHuhOptions(s.opts)...).
					Value(s.value),
			),
		).Run()

		if isAbort(err) {
			if i == 0 {
				return true // Back to stack selection
			}
			i--
			continue
		}
		if err != nil {
			return true
		}
		i++
	}

	return false
}
