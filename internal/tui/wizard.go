package tui

import (
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

// dimStyle returns a style for dimmed text (helper to avoid exporting the var).
func dimStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(dimColor)
}

// Run launches the TUI wizard and returns the completed config.
func Run() (*config.ProjectConfig, error) {
	fmt.Println(welcomeBanner())

	cfg := &config.ProjectConfig{}

	// Default project name to current directory name.
	defaultName := "my-project"
	if wd, err := os.Getwd(); err == nil {
		defaultName = filepath.Base(wd)
	}
	cfg.Name = defaultName

	// --- Step 1: Project name ---
	nameForm := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Project Name").
				Placeholder(defaultName).
				Value(&cfg.Name),
		),
	)
	if err := nameForm.Run(); err != nil {
		return nil, err
	}
	if cfg.Name == "" {
		cfg.Name = defaultName
	}

	// --- Step 2: Stack selection ---
	stackOptions := make([]huh.Option[string], 0, len(config.Presets)+1)
	for _, p := range config.Presets {
		label := fmt.Sprintf("%s  %s", p.Name, dimStyle().Render(p.Description))
		stackOptions = append(stackOptions, huh.NewOption(label, p.Name))
	}
	stackOptions = append(stackOptions, huh.NewOption(
		fmt.Sprintf("Custom Stack  %s", dimStyle().Render("pick each layer yourself")),
		"custom",
	))

	var stackChoice string
	stackForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose a Stack").
				Options(stackOptions...).
				Value(&stackChoice),
		),
	)
	if err := stackForm.Run(); err != nil {
		return nil, err
	}

	// --- Step 3: Apply preset or walk custom flow ---
	if stackChoice == "custom" {
		cfg.Stack = "custom"
		if err := runCustomFlow(cfg); err != nil {
			return nil, err
		}
	} else {
		applyPreset(cfg, stackChoice)
	}

	// Resolve package manager if empty.
	if cfg.PackageManager == "" {
		cfg.PackageManager = detectPackageManager()
	}

	// --- Step 4: Extras (always asked) ---
	extrasForm := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title("Extras").
				Options(toHuhOptions(config.ExtraOptions)...).
				Value(&cfg.Extras),
		),
	)
	if err := extrasForm.Run(); err != nil {
		return nil, err
	}

	// --- Step 5: Show summary and confirm ---
	fmt.Println(buildSummary(cfg))

	var confirmed bool
	confirmForm := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Create this project?").
				Affirmative("Yes").
				Negative("No").
				Value(&confirmed),
		),
	)
	if err := confirmForm.Run(); err != nil {
		return nil, err
	}

	if !confirmed {
		fmt.Println(errorStyle.Render("Cancelled."))
		return nil, nil
	}

	// Set project directory.
	if wd, err := os.Getwd(); err == nil {
		cfg.ProjectDir = filepath.Join(wd, cfg.Name)
	}

	fmt.Println(successStyle.Render("Let's build it!"))
	return cfg, nil
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

// runCustomFlow walks through each config layer with select fields,
// organized into 3 groups within a single form.
func runCustomFlow(cfg *config.ProjectConfig) error {
	form := huh.NewForm(
		// Group 1: Core stack
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Language").
				Options(toHuhOptions(config.Languages)...).
				Value(&cfg.Language),
			huh.NewSelect[string]().
				Title("Framework").
				Options(toHuhOptions(config.Frameworks)...).
				Value(&cfg.Framework),
			huh.NewSelect[string]().
				Title("Styling").
				Options(toHuhOptions(config.Styling)...).
				Value(&cfg.Styling),
		),
		// Group 2: Backend services
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Database / ORM").
				Options(toHuhOptions(config.Databases)...).
				Value(&cfg.Database),
			huh.NewSelect[string]().
				Title("Authentication").
				Options(toHuhOptions(config.AuthProviders)...).
				Value(&cfg.Auth),
			huh.NewSelect[string]().
				Title("Payments").
				Options(toHuhOptions(config.PaymentProviders)...).
				Value(&cfg.Payments),
			huh.NewSelect[string]().
				Title("Email Provider").
				Options(toHuhOptions(config.EmailProviders)...).
				Value(&cfg.Email),
		),
		// Group 3: Tooling
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Package Manager").
				Options(toHuhOptions(config.PackageManagers)...).
				Value(&cfg.PackageManager),
			huh.NewSelect[string]().
				Title("Deployment Target").
				Options(toHuhOptions(config.DeploymentTargets)...).
				Value(&cfg.Deployment),
		),
	)

	return form.Run()
}
