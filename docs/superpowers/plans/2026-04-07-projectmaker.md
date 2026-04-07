# ProjectMaker Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build a Go CLI tool (`project`) that launches an elegant TUI for scaffolding full-stack project stacks with automated GitHub/Vercel integration.

**Architecture:** Huh forms drive the TUI wizard. Config structs define presets and user choices. A scaffold engine executes commands and generates files. PROJECT_SPEC.md is always generated as the source of truth.

**Tech Stack:** Go, Huh (forms), Lip Gloss (styling), Bubbles (spinner/progress)

---

### Task 1: Initialize Go Project

**Files:**
- Create: `go.mod`
- Create: `main.go`

- [ ] **Step 1: Initialize Go module**

```bash
cd "/Users/jaden/Library/Mobile Documents/com~apple~CloudDocs/Projects/projectmaker"
go mod init github.com/JadenB9/projectmaker
```

- [ ] **Step 2: Install dependencies**

```bash
go get github.com/charmbracelet/huh
go get github.com/charmbracelet/lipgloss
go get github.com/charmbracelet/bubbles
go get github.com/charmbracelet/bubbletea
```

- [ ] **Step 3: Create main.go entry point**

```go
package main

import (
	"fmt"
	"os"

	"github.com/JadenB9/projectmaker/internal/config"
	"github.com/JadenB9/projectmaker/internal/tui"
)

func main() {
	cfg, err := tui.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	if cfg == nil {
		// User cancelled
		os.Exit(0)
	}
	_ = cfg // Will be used by scaffold engine in Task 5
}
```

- [ ] **Step 4: Commit**

```bash
git add go.mod main.go
git commit -m "init: go module with charm dependencies"
```

---

### Task 2: Define Config and Stack Presets

**Files:**
- Create: `internal/config/config.go`
- Create: `internal/config/presets.go`
- Create: `internal/config/options.go`

- [ ] **Step 1: Create config struct**

Create `internal/config/config.go`:

```go
package config

// ProjectConfig holds all user selections from the TUI wizard.
type ProjectConfig struct {
	Name           string
	Stack          string // preset name or "custom"
	Language       string
	Framework      string
	Styling        string
	Backend        string
	Database       string
	Auth           string
	Payments       string
	Email          string
	PackageManager string
	Deployment     string
	Extras         []string
	ProjectDir     string // absolute path to the project directory
}
```

- [ ] **Step 2: Create options constants**

Create `internal/config/options.go`:

```go
package config

type Option struct {
	Label string
	Value string
}

var Languages = []Option{
	{Label: "TypeScript", Value: "typescript"},
	{Label: "JavaScript", Value: "javascript"},
	{Label: "Python", Value: "python"},
	{Label: "Go", Value: "go"},
	{Label: "Rust", Value: "rust"},
	{Label: "Ruby", Value: "ruby"},
	{Label: "PHP", Value: "php"},
}

var Frameworks = []Option{
	{Label: "Next.js", Value: "nextjs"},
	{Label: "React (Vite)", Value: "react-vite"},
	{Label: "SvelteKit", Value: "sveltekit"},
	{Label: "Nuxt", Value: "nuxt"},
	{Label: "Astro", Value: "astro"},
	{Label: "Angular", Value: "angular"},
	{Label: "Express", Value: "express"},
	{Label: "Hono", Value: "hono"},
	{Label: "Fastify", Value: "fastify"},
	{Label: "FastAPI", Value: "fastapi"},
	{Label: "Django", Value: "django"},
	{Label: "Flask", Value: "flask"},
	{Label: "Rails", Value: "rails"},
	{Label: "Laravel", Value: "laravel"},
	{Label: "Gin", Value: "gin"},
	{Label: "Fiber", Value: "fiber"},
	{Label: "None", Value: "none"},
}

var Styling = []Option{
	{Label: "Tailwind CSS", Value: "tailwind"},
	{Label: "Tailwind + shadcn/ui", Value: "tailwind-shadcn"},
	{Label: "CSS Modules", Value: "css-modules"},
	{Label: "Styled Components", Value: "styled-components"},
	{Label: "None", Value: "none"},
}

var Databases = []Option{
	{Label: "Drizzle ORM", Value: "drizzle"},
	{Label: "Prisma", Value: "prisma"},
	{Label: "Supabase", Value: "supabase"},
	{Label: "MongoDB (Mongoose)", Value: "mongodb"},
	{Label: "SQLite", Value: "sqlite"},
	{Label: "None", Value: "none"},
}

var AuthProviders = []Option{
	{Label: "Clerk", Value: "clerk"},
	{Label: "NextAuth / Auth.js", Value: "nextauth"},
	{Label: "Supabase Auth", Value: "supabase-auth"},
	{Label: "Lucia", Value: "lucia"},
	{Label: "None", Value: "none"},
}

var PaymentProviders = []Option{
	{Label: "Stripe", Value: "stripe"},
	{Label: "None", Value: "none"},
}

var EmailProviders = []Option{
	{Label: "Resend", Value: "resend"},
	{Label: "None", Value: "none"},
}

var PackageManagers = []Option{
	{Label: "npm", Value: "npm"},
	{Label: "pnpm", Value: "pnpm"},
	{Label: "bun", Value: "bun"},
	{Label: "yarn", Value: "yarn"},
}

var DeploymentTargets = []Option{
	{Label: "Vercel", Value: "vercel"},
	{Label: "Railway", Value: "railway"},
	{Label: "Cloudflare", Value: "cloudflare"},
	{Label: "Docker (self-hosted)", Value: "docker"},
	{Label: "None", Value: "none"},
}

var ExtraOptions = []Option{
	{Label: "GitHub Actions CI", Value: "github-actions"},
	{Label: "Dockerfile + docker-compose", Value: "docker"},
	{Label: "n8n (workflow automation)", Value: "n8n"},
	{Label: "Vercel Analytics", Value: "vercel-analytics"},
	{Label: "PostHog Analytics", Value: "posthog"},
	{Label: "ESLint + Prettier", Value: "eslint-prettier"},
}
```

- [ ] **Step 3: Create preset stacks**

Create `internal/config/presets.go`:

```go
package config

type Preset struct {
	Name        string
	Description string
	Config      ProjectConfig
}

var Presets = []Preset{
	{
		Name:        "Next.js Full-Stack",
		Description: "Next.js App Router + TypeScript + Tailwind + shadcn/ui",
		Config: ProjectConfig{
			Language:       "typescript",
			Framework:      "nextjs",
			Styling:        "tailwind-shadcn",
			Backend:        "api-routes",
			PackageManager: "auto",
			Deployment:     "vercel",
		},
	},
	{
		Name:        "T3 Stack",
		Description: "Next.js + tRPC + Prisma + NextAuth + Tailwind",
		Config: ProjectConfig{
			Language:       "typescript",
			Framework:      "nextjs",
			Styling:        "tailwind",
			Backend:        "trpc",
			Database:       "prisma",
			Auth:           "nextauth",
			PackageManager: "auto",
			Deployment:     "vercel",
		},
	},
	{
		Name:        "MERN Stack",
		Description: "MongoDB + Express + React + Node + TypeScript",
		Config: ProjectConfig{
			Language:       "typescript",
			Framework:      "react-vite",
			Styling:        "tailwind",
			Backend:        "express",
			Database:       "mongodb",
			PackageManager: "auto",
		},
	},
	{
		Name:        "Next.js + Supabase",
		Description: "Next.js App Router + Supabase Auth/DB + Tailwind + shadcn/ui",
		Config: ProjectConfig{
			Language:       "typescript",
			Framework:      "nextjs",
			Styling:        "tailwind-shadcn",
			Backend:        "api-routes",
			Database:       "supabase",
			Auth:           "supabase-auth",
			PackageManager: "auto",
			Deployment:     "vercel",
		},
	},
	{
		Name:        "Next.js + Convex",
		Description: "Next.js + Convex real-time backend + Tailwind",
		Config: ProjectConfig{
			Language:       "typescript",
			Framework:      "nextjs",
			Styling:        "tailwind",
			Backend:        "convex",
			PackageManager: "auto",
			Deployment:     "vercel",
		},
	},
	{
		Name:        "SvelteKit Full-Stack",
		Description: "SvelteKit + TypeScript + Tailwind + Drizzle",
		Config: ProjectConfig{
			Language:       "typescript",
			Framework:      "sveltekit",
			Styling:        "tailwind",
			Database:       "drizzle",
			PackageManager: "auto",
		},
	},
	{
		Name:        "Nuxt Full-Stack",
		Description: "Nuxt 4 + TypeScript + Tailwind + Nitro",
		Config: ProjectConfig{
			Language:       "typescript",
			Framework:      "nuxt",
			Styling:        "tailwind",
			PackageManager: "auto",
		},
	},
	{
		Name:        "Rails + React",
		Description: "Ruby on Rails API + React SPA + TypeScript",
		Config: ProjectConfig{
			Language:       "ruby",
			Framework:      "rails",
			Styling:        "tailwind",
			Backend:        "rails",
			PackageManager: "auto",
		},
	},
	{
		Name:        "Django + React",
		Description: "Django REST Framework + React SPA + TypeScript",
		Config: ProjectConfig{
			Language:       "python",
			Framework:      "django",
			Styling:        "tailwind",
			Backend:        "django",
			PackageManager: "auto",
		},
	},
	{
		Name:        "FastAPI + React",
		Description: "FastAPI + React SPA + TypeScript",
		Config: ProjectConfig{
			Language:       "python",
			Framework:      "fastapi",
			Styling:        "tailwind",
			Backend:        "fastapi",
			PackageManager: "auto",
		},
	},
	{
		Name:        "Go + HTMX",
		Description: "Go backend + HTMX + Tailwind",
		Config: ProjectConfig{
			Language:  "go",
			Framework: "gin",
			Styling:   "tailwind",
			Backend:   "go",
		},
	},
	{
		Name:        "Laravel + Vue",
		Description: "Laravel + Vue 3 + Inertia + Tailwind",
		Config: ProjectConfig{
			Language:  "php",
			Framework: "laravel",
			Styling:   "tailwind",
			Backend:   "laravel",
		},
	},
}
```

- [ ] **Step 4: Commit**

```bash
git add internal/config/
git commit -m "feat: add config structs and preset stack definitions"
```

---

### Task 3: Build TUI Styles

**Files:**
- Create: `internal/tui/styles.go`

- [ ] **Step 1: Create calm color palette and styles**

Create `internal/tui/styles.go`:

```go
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
```

- [ ] **Step 2: Commit**

```bash
git add internal/tui/styles.go
git commit -m "feat: add calm TUI color palette and styles"
```

---

### Task 4: Build TUI Wizard

**Files:**
- Create: `internal/tui/wizard.go`
- Create: `internal/tui/welcome.go`

- [ ] **Step 1: Create welcome banner**

Create `internal/tui/welcome.go`:

```go
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
```

- [ ] **Step 2: Create the main wizard flow**

Create `internal/tui/wizard.go`:

```go
package tui

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/JadenB9/projectmaker/internal/config"
)

// Run launches the TUI wizard and returns the completed config.
func Run() (*config.ProjectConfig, error) {
	fmt.Println(welcomeBanner())

	cfg := &config.ProjectConfig{}

	// Get current directory name as default project name
	cwd, _ := os.Getwd()
	defaultName := filepath.Base(cwd)

	// Step 1: Project name
	var projectName string
	err := huh.NewInput().
		Title("Project Name").
		Description("What should your project be called?").
		Placeholder(defaultName).
		Value(&projectName).
		Run()
	if err != nil {
		return nil, err
	}
	if projectName == "" {
		projectName = defaultName
	}
	cfg.Name = projectName
	cfg.ProjectDir = filepath.Join(cwd, projectName)

	// Step 2: Stack selection
	stackChoice, err := selectStack()
	if err != nil {
		return nil, err
	}

	if stackChoice == "custom" {
		cfg.Stack = "custom"
		if err := customStackWizard(cfg); err != nil {
			return nil, err
		}
	} else {
		// Apply preset
		for _, p := range config.Presets {
			if p.Name == stackChoice {
				cfg.Stack = p.Name
				cfg.Language = p.Config.Language
				cfg.Framework = p.Config.Framework
				cfg.Styling = p.Config.Styling
				cfg.Backend = p.Config.Backend
				cfg.Database = p.Config.Database
				cfg.Auth = p.Config.Auth
				cfg.Payments = p.Config.Payments
				cfg.Email = p.Config.Email
				cfg.Deployment = p.Config.Deployment
				cfg.PackageManager = p.Config.PackageManager
				break
			}
		}
	}

	// Resolve "auto" package manager
	if cfg.PackageManager == "auto" || cfg.PackageManager == "" {
		cfg.PackageManager = detectPackageManager()
	}

	// Step 3: Extras (always ask)
	extras, err := selectExtras()
	if err != nil {
		return nil, err
	}
	cfg.Extras = extras

	// Step 4: Confirmation
	confirmed, err := confirmConfig(cfg)
	if err != nil {
		return nil, err
	}
	if !confirmed {
		return nil, nil
	}

	return cfg, nil
}

func selectStack() (string, error) {
	options := make([]huh.Option[string], 0, len(config.Presets)+1)
	for _, p := range config.Presets {
		label := fmt.Sprintf("%s  %s", p.Name, subtitleStyle.Render(p.Description))
		options = append(options, huh.NewOption(label, p.Name))
	}
	options = append(options, huh.NewOption(
		fmt.Sprintf("%s  %s", "Custom Stack", subtitleStyle.Render("Pick every layer yourself")),
		"custom",
	))

	var choice string
	err := huh.NewSelect[string]().
		Title("Choose a Stack").
		Description("Select a preset or build your own").
		Options(options...).
		Value(&choice).
		Run()

	return choice, err
}

func customStackWizard(cfg *config.ProjectConfig) error {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Language / Runtime").
				Options(toHuhOptions(config.Languages)...).
				Value(&cfg.Language),

			huh.NewSelect[string]().
				Title("Framework").
				Options(toHuhOptions(config.Frameworks)...).
				Value(&cfg.Framework),

			huh.NewSelect[string]().
				Title("Frontend Styling").
				Options(toHuhOptions(config.Styling)...).
				Value(&cfg.Styling),
		),

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
		),

		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Email Service").
				Options(toHuhOptions(config.EmailProviders)...).
				Value(&cfg.Email),

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

func selectExtras() ([]string, error) {
	var extras []string
	err := huh.NewMultiSelect[string]().
		Title("Extras").
		Description("Select any additional features (space to toggle)").
		Options(toHuhOptions(config.ExtraOptions)...).
		Value(&extras).
		Run()
	return extras, err
}

func confirmConfig(cfg *config.ProjectConfig) (bool, error) {
	summary := buildSummary(cfg)
	fmt.Println(boxStyle.Render(summary))

	var confirmed bool
	err := huh.NewConfirm().
		Title("Create this project?").
		Affirmative("Yes, create it").
		Negative("No, cancel").
		Value(&confirmed).
		Run()

	return confirmed, err
}

func buildSummary(cfg *config.ProjectConfig) string {
	var b strings.Builder
	b.WriteString(titleStyle.Render("Project Summary") + "\n\n")

	row := func(label, value string) {
		if value != "" && value != "none" {
			b.WriteString(fmt.Sprintf("  %s  %s\n",
				lipgloss.NewStyle().Foreground(dimColor).Width(20).Render(label),
				lipgloss.NewStyle().Foreground(textColor).Render(value),
			))
		}
	}

	row("Name:", cfg.Name)
	row("Stack:", cfg.Stack)
	row("Language:", cfg.Language)
	row("Framework:", cfg.Framework)
	row("Styling:", cfg.Styling)
	row("Database:", cfg.Database)
	row("Auth:", cfg.Auth)
	row("Payments:", cfg.Payments)
	row("Email:", cfg.Email)
	row("Pkg Manager:", cfg.PackageManager)
	row("Deployment:", cfg.Deployment)

	if len(cfg.Extras) > 0 {
		row("Extras:", strings.Join(cfg.Extras, ", "))
	}

	return b.String()
}

func toHuhOptions(opts []config.Option) []huh.Option[string] {
	result := make([]huh.Option[string], len(opts))
	for i, o := range opts {
		result[i] = huh.NewOption(o.Label, o.Value)
	}
	return result
}

func detectPackageManager() string {
	for _, pm := range []string{"bun", "pnpm", "yarn", "npm"} {
		if _, err := exec.LookPath(pm); err == nil {
			return pm
		}
	}
	return "npm"
}
```

- [ ] **Step 3: Commit**

```bash
git add internal/tui/
git commit -m "feat: TUI wizard with stack selection and custom builder"
```

---

### Task 5: Build CLI Detection

**Files:**
- Create: `internal/services/detect.go`

- [ ] **Step 1: Create CLI detection utility**

Create `internal/services/detect.go`:

```go
package services

import (
	"os/exec"
	"strings"
)

type CLIStatus struct {
	Git     bool
	GitHub  bool
	Vercel  bool
	Docker  bool
	Node    bool
	Bun     bool
	Pnpm    bool
	Yarn    bool
	Npm     bool
	Python  bool
	Go      bool
	Cargo   bool
	Ruby    bool
	PHP     bool
}

func DetectCLIs() CLIStatus {
	return CLIStatus{
		Git:    hasCmd("git"),
		GitHub: hasCmd("gh"),
		Vercel: hasCmd("vercel"),
		Docker: hasCmd("docker"),
		Node:   hasCmd("node"),
		Bun:    hasCmd("bun"),
		Pnpm:   hasCmd("pnpm"),
		Yarn:   hasCmd("yarn"),
		Npm:    hasCmd("npm"),
		Python: hasCmd("python3"),
		Go:     hasCmd("go"),
		Cargo:  hasCmd("cargo"),
		Ruby:   hasCmd("ruby"),
		PHP:    hasCmd("php"),
	}
}

func hasCmd(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

func RunCmd(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(out)), err
}

func RunCmdInDir(dir, name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(out)), err
}
```

- [ ] **Step 2: Commit**

```bash
git add internal/services/detect.go
git commit -m "feat: CLI tool detection utility"
```

---

### Task 6: Build Scaffold Engine

**Files:**
- Create: `internal/scaffold/scaffold.go`
- Create: `internal/scaffold/gitignore.go`
- Create: `internal/scaffold/env.go`

- [ ] **Step 1: Create the main scaffold orchestrator**

Create `internal/scaffold/scaffold.go`:

```go
package scaffold

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/JadenB9/projectmaker/internal/config"
	"github.com/JadenB9/projectmaker/internal/services"
)

type Result struct {
	Steps    []StepResult
	SpecPath string
}

type StepResult struct {
	Name    string
	Status  string // "done", "skipped", "manual"
	Message string
}

func Run(cfg *config.ProjectConfig, clis services.CLIStatus) (*Result, error) {
	result := &Result{}

	// Create project directory
	if err := os.MkdirAll(cfg.ProjectDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create project directory: %w", err)
	}

	// Git init
	if clis.Git {
		_, err := services.RunCmdInDir(cfg.ProjectDir, "git", "init")
		if err != nil {
			result.Steps = append(result.Steps, StepResult{"Git init", "skipped", "failed to init git"})
		} else {
			result.Steps = append(result.Steps, StepResult{"Git init", "done", ""})
		}
	}

	// Create .gitignore
	if err := writeGitignore(cfg); err != nil {
		result.Steps = append(result.Steps, StepResult{".gitignore", "skipped", err.Error()})
	} else {
		result.Steps = append(result.Steps, StepResult{".gitignore", "done", ""})
	}

	// Create .env
	envVars := collectEnvVars(cfg)
	if len(envVars) > 0 {
		if err := writeEnvFile(cfg, envVars); err != nil {
			result.Steps = append(result.Steps, StepResult{".env", "skipped", err.Error()})
		} else {
			result.Steps = append(result.Steps, StepResult{".env", "done", ""})
		}
	}

	// Run framework scaffolding
	frameworkResult := scaffoldFramework(cfg, clis)
	result.Steps = append(result.Steps, frameworkResult...)

	// GitHub repo
	if clis.GitHub {
		_, err := services.RunCmdInDir(cfg.ProjectDir, "gh", "repo", "create", cfg.Name, "--private", "--source=.", "--remote=origin")
		if err != nil {
			result.Steps = append(result.Steps, StepResult{"GitHub repo", "manual", "Run: gh repo create " + cfg.Name + " --private --source=. --remote=origin"})
		} else {
			result.Steps = append(result.Steps, StepResult{"GitHub repo", "done", ""})
		}
	} else {
		result.Steps = append(result.Steps, StepResult{"GitHub repo", "manual", "Install gh CLI and run: gh repo create"})
	}

	// Vercel link
	if cfg.Deployment == "vercel" {
		if clis.Vercel {
			_, err := services.RunCmdInDir(cfg.ProjectDir, "vercel", "link", "--yes")
			if err != nil {
				result.Steps = append(result.Steps, StepResult{"Vercel link", "manual", "Run: vercel link"})
			} else {
				result.Steps = append(result.Steps, StepResult{"Vercel link", "done", ""})
			}
		} else {
			result.Steps = append(result.Steps, StepResult{"Vercel link", "manual", "Install vercel CLI and run: vercel link"})
		}
	}

	return result, nil
}

func scaffoldFramework(cfg *config.ProjectConfig, clis services.CLIStatus) []StepResult {
	var results []StepResult

	switch cfg.Framework {
	case "nextjs":
		results = append(results, scaffoldNextJS(cfg, clis)...)
	case "react-vite":
		results = append(results, scaffoldViteReact(cfg, clis)...)
	case "sveltekit":
		results = append(results, scaffoldSvelteKit(cfg, clis)...)
	case "nuxt":
		results = append(results, scaffoldNuxt(cfg, clis)...)
	case "express":
		results = append(results, scaffoldExpress(cfg, clis)...)
	case "django", "fastapi", "flask":
		results = append(results, scaffoldPython(cfg, clis)...)
	case "rails":
		results = append(results, scaffoldRails(cfg, clis)...)
	case "laravel":
		results = append(results, scaffoldLaravel(cfg, clis)...)
	case "gin", "fiber":
		results = append(results, scaffoldGoBackend(cfg, clis)...)
	default:
		results = append(results, StepResult{"Framework", "manual", "No auto-scaffold available for " + cfg.Framework})
	}

	return results
}

func scaffoldNextJS(cfg *config.ProjectConfig, clis services.CLIStatus) []StepResult {
	var results []StepResult
	pm := cfg.PackageManager

	// Use create-next-app
	args := []string{"create-next-app@latest", cfg.ProjectDir,
		"--ts", "--eslint", "--app", "--src-dir",
		"--import-alias", "@/*",
	}

	if cfg.Styling == "tailwind" || cfg.Styling == "tailwind-shadcn" {
		args = append(args, "--tailwind")
	} else {
		args = append(args, "--no-tailwind")
	}

	usePm := pm
	if pm == "bun" {
		args = append(args, "--use-bun")
		usePm = "bunx"
	} else if pm == "pnpm" {
		args = append(args, "--use-pnpm")
		usePm = "pnpx"
	} else if pm == "yarn" {
		args = append(args, "--use-yarn")
		usePm = "npx"
	} else {
		usePm = "npx"
	}

	_, err := services.RunCmd(usePm, args...)
	if err != nil {
		results = append(results, StepResult{"Next.js scaffold", "manual", fmt.Sprintf("Run: %s %s", usePm, joinArgs(args))})
		return results
	}
	results = append(results, StepResult{"Next.js scaffold", "done", ""})

	// shadcn/ui
	if cfg.Styling == "tailwind-shadcn" {
		_, err := services.RunCmdInDir(cfg.ProjectDir, usePm, "shadcn@latest", "init", "--yes")
		if err != nil {
			results = append(results, StepResult{"shadcn/ui", "manual", "Run: npx shadcn@latest init"})
		} else {
			results = append(results, StepResult{"shadcn/ui", "done", ""})
		}
	}

	// Install additional deps based on selections
	results = append(results, installDeps(cfg, clis)...)

	return results
}

func scaffoldViteReact(cfg *config.ProjectConfig, clis services.CLIStatus) []StepResult {
	var results []StepResult
	pm := cfg.PackageManager
	runner := pmRunner(pm)
	template := "react-ts"
	if cfg.Language == "javascript" {
		template = "react"
	}

	_, err := services.RunCmd(runner, "create-vite@latest", cfg.ProjectDir, "--template", template)
	if err != nil {
		results = append(results, StepResult{"Vite React scaffold", "manual", fmt.Sprintf("Run: %s create-vite@latest %s --template %s", runner, cfg.Name, template)})
	} else {
		results = append(results, StepResult{"Vite React scaffold", "done", ""})
	}
	results = append(results, installDeps(cfg, clis)...)
	return results
}

func scaffoldSvelteKit(cfg *config.ProjectConfig, clis services.CLIStatus) []StepResult {
	var results []StepResult
	runner := pmRunner(cfg.PackageManager)
	_, err := services.RunCmd(runner, "sv", "create", cfg.ProjectDir, "--template", "minimal", "--types", "ts")
	if err != nil {
		results = append(results, StepResult{"SvelteKit scaffold", "manual", "Run: npx sv create " + cfg.Name})
	} else {
		results = append(results, StepResult{"SvelteKit scaffold", "done", ""})
	}
	return results
}

func scaffoldNuxt(cfg *config.ProjectConfig, clis services.CLIStatus) []StepResult {
	var results []StepResult
	runner := pmRunner(cfg.PackageManager)
	_, err := services.RunCmd(runner, "nuxi@latest", "init", cfg.ProjectDir)
	if err != nil {
		results = append(results, StepResult{"Nuxt scaffold", "manual", "Run: npx nuxi@latest init " + cfg.Name})
	} else {
		results = append(results, StepResult{"Nuxt scaffold", "done", ""})
	}
	return results
}

func scaffoldExpress(cfg *config.ProjectConfig, clis services.CLIStatus) []StepResult {
	var results []StepResult
	if err := os.MkdirAll(cfg.ProjectDir, 0755); err != nil {
		results = append(results, StepResult{"Express scaffold", "manual", "Create project directory manually"})
		return results
	}
	_, err := services.RunCmdInDir(cfg.ProjectDir, cfg.PackageManager, "init", "-y")
	if err != nil {
		results = append(results, StepResult{"Express init", "manual", "Run: npm init -y"})
	} else {
		results = append(results, StepResult{"Express init", "done", ""})
	}
	results = append(results, installDeps(cfg, clis)...)
	return results
}

func scaffoldPython(cfg *config.ProjectConfig, clis services.CLIStatus) []StepResult {
	var results []StepResult
	if err := os.MkdirAll(cfg.ProjectDir, 0755); err != nil {
		return results
	}
	_, err := services.RunCmdInDir(cfg.ProjectDir, "python3", "-m", "venv", "venv")
	if err != nil {
		results = append(results, StepResult{"Python venv", "manual", "Run: python3 -m venv venv"})
	} else {
		results = append(results, StepResult{"Python venv", "done", ""})
	}
	return results
}

func scaffoldRails(cfg *config.ProjectConfig, clis services.CLIStatus) []StepResult {
	var results []StepResult
	if clis.Ruby {
		_, err := services.RunCmd("rails", "new", cfg.ProjectDir, "--api", "--database=postgresql")
		if err != nil {
			results = append(results, StepResult{"Rails scaffold", "manual", "Run: rails new " + cfg.Name + " --api"})
		} else {
			results = append(results, StepResult{"Rails scaffold", "done", ""})
		}
	} else {
		results = append(results, StepResult{"Rails scaffold", "manual", "Install Ruby and Rails, then run: rails new " + cfg.Name})
	}
	return results
}

func scaffoldLaravel(cfg *config.ProjectConfig, clis services.CLIStatus) []StepResult {
	var results []StepResult
	if clis.PHP {
		_, err := services.RunCmd("composer", "create-project", "laravel/laravel", cfg.ProjectDir)
		if err != nil {
			results = append(results, StepResult{"Laravel scaffold", "manual", "Run: composer create-project laravel/laravel " + cfg.Name})
		} else {
			results = append(results, StepResult{"Laravel scaffold", "done", ""})
		}
	} else {
		results = append(results, StepResult{"Laravel scaffold", "manual", "Install PHP and Composer, then: composer create-project laravel/laravel " + cfg.Name})
	}
	return results
}

func scaffoldGoBackend(cfg *config.ProjectConfig, clis services.CLIStatus) []StepResult {
	var results []StepResult
	if err := os.MkdirAll(cfg.ProjectDir, 0755); err != nil {
		return results
	}
	_, err := services.RunCmdInDir(cfg.ProjectDir, "go", "mod", "init", "github.com/"+cfg.Name)
	if err != nil {
		results = append(results, StepResult{"Go mod init", "manual", "Run: go mod init github.com/" + cfg.Name})
	} else {
		results = append(results, StepResult{"Go mod init", "done", ""})
	}
	return results
}

func installDeps(cfg *config.ProjectConfig, clis services.CLIStatus) []StepResult {
	var results []StepResult
	var deps []string

	switch cfg.Database {
	case "prisma":
		deps = append(deps, "prisma", "@prisma/client")
	case "drizzle":
		deps = append(deps, "drizzle-orm", "drizzle-kit")
	case "mongodb":
		deps = append(deps, "mongoose")
	case "supabase":
		deps = append(deps, "@supabase/supabase-js")
	}

	switch cfg.Auth {
	case "clerk":
		deps = append(deps, "@clerk/nextjs")
	case "nextauth":
		deps = append(deps, "next-auth")
	case "lucia":
		deps = append(deps, "lucia")
	}

	if cfg.Payments == "stripe" {
		deps = append(deps, "stripe", "@stripe/stripe-js")
	}
	if cfg.Email == "resend" {
		deps = append(deps, "resend")
	}

	if len(deps) == 0 {
		return results
	}

	pm := cfg.PackageManager
	installCmd := "install"
	if pm == "yarn" {
		installCmd = "add"
	}
	if pm == "bun" {
		installCmd = "add"
	}
	if pm == "pnpm" {
		installCmd = "add"
	}

	args := append([]string{installCmd}, deps...)
	_, err := services.RunCmdInDir(cfg.ProjectDir, pm, args...)
	if err != nil {
		results = append(results, StepResult{"Install deps", "manual", fmt.Sprintf("Run: %s %s %s", pm, installCmd, joinArgs(deps))})
	} else {
		results = append(results, StepResult{"Install deps", "done", ""})
	}

	return results
}

func pmRunner(pm string) string {
	switch pm {
	case "bun":
		return "bunx"
	case "pnpm":
		return "pnpx"
	default:
		return "npx"
	}
}

func joinArgs(args []string) string {
	quoted := make([]string, len(args))
	for i, a := range args {
		quoted[i] = a
	}
	return fmt.Sprintf("%s", quoted)
}
```

- [ ] **Step 2: Create .gitignore generator**

Create `internal/scaffold/gitignore.go`:

```go
package scaffold

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/JadenB9/projectmaker/internal/config"
)

func writeGitignore(cfg *config.ProjectConfig) error {
	var lines []string

	// Common
	lines = append(lines, "# Dependencies", "node_modules/", ".pnp/", ".pnp.js", "")
	lines = append(lines, "# Environment", ".env", ".env.local", ".env.*.local", "")
	lines = append(lines, "# Build", "dist/", "build/", ".next/", ".nuxt/", ".svelte-kit/", ".output/", "")
	lines = append(lines, "# IDE", ".idea/", ".vscode/", "*.swp", "*.swo", ".DS_Store", "")
	lines = append(lines, "# Debug", "npm-debug.log*", "yarn-debug.log*", "yarn-error.log*", "")

	// Python
	if cfg.Language == "python" {
		lines = append(lines, "# Python", "venv/", "__pycache__/", "*.pyc", ".pytest_cache/", "")
	}

	// Go
	if cfg.Language == "go" {
		lines = append(lines, "# Go", "*.exe", "*.exe~", "*.dll", "*.so", "*.dylib", "")
	}

	// Rust
	if cfg.Language == "rust" {
		lines = append(lines, "# Rust", "target/", "Cargo.lock", "")
	}

	// Docker
	for _, e := range cfg.Extras {
		if e == "docker" {
			lines = append(lines, "# Docker", "docker-compose.override.yml", "")
			break
		}
	}

	content := strings.Join(lines, "\n")
	return os.WriteFile(filepath.Join(cfg.ProjectDir, ".gitignore"), []byte(content), 0644)
}
```

- [ ] **Step 3: Create .env generator**

Create `internal/scaffold/env.go`:

```go
package scaffold

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/JadenB9/projectmaker/internal/config"
)

type EnvVar struct {
	Key         string
	Placeholder string
	Comment     string
}

func collectEnvVars(cfg *config.ProjectConfig) []EnvVar {
	var vars []EnvVar

	// Database
	switch cfg.Database {
	case "prisma", "drizzle":
		vars = append(vars, EnvVar{"DATABASE_URL", "postgresql://user:password@localhost:5432/dbname", "Database connection string"})
	case "supabase":
		vars = append(vars, EnvVar{"NEXT_PUBLIC_SUPABASE_URL", "https://your-project.supabase.co", "Supabase project URL (Settings > API)"})
		vars = append(vars, EnvVar{"NEXT_PUBLIC_SUPABASE_ANON_KEY", "your-anon-key", "Supabase anon key (Settings > API)"})
		vars = append(vars, EnvVar{"SUPABASE_SERVICE_ROLE_KEY", "your-service-role-key", "Supabase service role key (Settings > API)"})
	case "mongodb":
		vars = append(vars, EnvVar{"MONGODB_URI", "mongodb://localhost:27017/dbname", "MongoDB connection string"})
	}

	// Auth
	switch cfg.Auth {
	case "clerk":
		vars = append(vars, EnvVar{"NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY", "pk_test_...", "Clerk publishable key (Dashboard > API Keys)"})
		vars = append(vars, EnvVar{"CLERK_SECRET_KEY", "sk_test_...", "Clerk secret key (Dashboard > API Keys)"})
	case "nextauth":
		vars = append(vars, EnvVar{"NEXTAUTH_URL", "http://localhost:3000", "NextAuth base URL"})
		vars = append(vars, EnvVar{"NEXTAUTH_SECRET", "your-secret-here", "Run: openssl rand -base64 32"})
	}

	// Payments
	if cfg.Payments == "stripe" {
		vars = append(vars, EnvVar{"STRIPE_SECRET_KEY", "sk_test_...", "Stripe secret key (Dashboard > Developers > API Keys)"})
		vars = append(vars, EnvVar{"NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY", "pk_test_...", "Stripe publishable key (Dashboard > Developers > API Keys)"})
		vars = append(vars, EnvVar{"STRIPE_WEBHOOK_SECRET", "whsec_...", "Stripe webhook secret (Dashboard > Developers > Webhooks)"})
	}

	// Email
	if cfg.Email == "resend" {
		vars = append(vars, EnvVar{"RESEND_API_KEY", "re_...", "Resend API key (Dashboard > API Keys)"})
	}

	return vars
}

func writeEnvFile(cfg *config.ProjectConfig, vars []EnvVar) error {
	var lines []string
	lines = append(lines, "# Environment Variables")
	lines = append(lines, "# Generated by projectmaker")
	lines = append(lines, "# Fill in the values below", "")

	for _, v := range vars {
		lines = append(lines, "# "+v.Comment)
		lines = append(lines, v.Key+"="+v.Placeholder)
		lines = append(lines, "")
	}

	content := strings.Join(lines, "\n")

	// Write .env
	envPath := filepath.Join(cfg.ProjectDir, ".env")
	if err := os.WriteFile(envPath, []byte(content), 0644); err != nil {
		return err
	}

	// Write .env.example (safe to commit)
	exampleLines := []string{"# Environment Variables - Copy to .env and fill in values", ""}
	for _, v := range vars {
		exampleLines = append(exampleLines, "# "+v.Comment)
		exampleLines = append(exampleLines, v.Key+"=")
		exampleLines = append(exampleLines, "")
	}
	exampleContent := strings.Join(exampleLines, "\n")
	return os.WriteFile(filepath.Join(cfg.ProjectDir, ".env.example"), []byte(exampleContent), 0644)
}
```

- [ ] **Step 4: Commit**

```bash
git add internal/scaffold/ internal/services/
git commit -m "feat: scaffold engine with framework, env, and gitignore generation"
```

---

### Task 7: Build PROJECT_SPEC.md Generator

**Files:**
- Create: `internal/spec/generator.go`

- [ ] **Step 1: Create spec generator**

Create `internal/spec/generator.go`:

```go
package spec

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/JadenB9/projectmaker/internal/config"
	"github.com/JadenB9/projectmaker/internal/scaffold"
)

func Generate(cfg *config.ProjectConfig, result *scaffold.Result) error {
	var b strings.Builder

	b.WriteString("# Project Specification\n\n")
	b.WriteString(fmt.Sprintf("> Generated by [projectmaker](https://github.com/JadenB9/projectmaker) on %s\n\n", time.Now().Format("2006-01-02")))

	// Project Info
	b.WriteString("## Project Info\n\n")
	b.WriteString(fmt.Sprintf("- **Name:** %s\n", cfg.Name))
	b.WriteString(fmt.Sprintf("- **Directory:** `%s`\n", cfg.ProjectDir))
	b.WriteString(fmt.Sprintf("- **Stack:** %s\n", cfg.Stack))
	b.WriteString(fmt.Sprintf("- **Created:** %s\n\n", time.Now().Format("2006-01-02 15:04")))

	// Stack Summary
	b.WriteString("## Stack Summary\n\n")
	b.WriteString("| Layer | Choice |\n")
	b.WriteString("|-------|--------|\n")
	writeRow := func(layer, value string) {
		if value != "" && value != "none" {
			b.WriteString(fmt.Sprintf("| %s | %s |\n", layer, value))
		}
	}
	writeRow("Language", cfg.Language)
	writeRow("Framework", cfg.Framework)
	writeRow("Styling", cfg.Styling)
	writeRow("Database", cfg.Database)
	writeRow("Auth", cfg.Auth)
	writeRow("Payments", cfg.Payments)
	writeRow("Email", cfg.Email)
	writeRow("Package Manager", cfg.PackageManager)
	writeRow("Deployment", cfg.Deployment)
	if len(cfg.Extras) > 0 {
		writeRow("Extras", strings.Join(cfg.Extras, ", "))
	}
	b.WriteString("\n")

	// Environment Variables
	envVars := collectEnvDocs(cfg)
	if len(envVars) > 0 {
		b.WriteString("## Environment Variables\n\n")
		b.WriteString("Copy `.env.example` to `.env` and fill in the values:\n\n")
		b.WriteString("| Variable | Where to Get It |\n")
		b.WriteString("|----------|----------------|\n")
		for _, v := range envVars {
			b.WriteString(fmt.Sprintf("| `%s` | %s |\n", v.key, v.source))
		}
		b.WriteString("\n")
	}

	// Setup Steps (manual)
	manualSteps := collectManualSteps(result)
	if len(manualSteps) > 0 {
		b.WriteString("## Manual Setup Steps\n\n")
		b.WriteString("The following could not be automated and need manual action:\n\n")
		for i, step := range manualSteps {
			b.WriteString(fmt.Sprintf("%d. **%s:** %s\n", i+1, step.name, step.instruction))
		}
		b.WriteString("\n")
	}

	// Dev Commands
	b.WriteString("## Development Commands\n\n")
	b.WriteString(devCommands(cfg))

	// Deployment
	b.WriteString("## Deployment\n\n")
	b.WriteString(deploymentDocs(cfg))

	// Architecture Notes
	b.WriteString("## Architecture Notes\n\n")
	b.WriteString(architectureNotes(cfg))

	return os.WriteFile(filepath.Join(cfg.ProjectDir, "PROJECT_SPEC.md"), []byte(b.String()), 0644)
}

type envDoc struct {
	key    string
	source string
}

func collectEnvDocs(cfg *config.ProjectConfig) []envDoc {
	var docs []envDoc

	switch cfg.Database {
	case "prisma", "drizzle":
		docs = append(docs, envDoc{"DATABASE_URL", "Your database provider (e.g., Neon, PlanetScale, local PostgreSQL)"})
	case "supabase":
		docs = append(docs, envDoc{"NEXT_PUBLIC_SUPABASE_URL", "[Supabase Dashboard](https://supabase.com/dashboard) > Settings > API"})
		docs = append(docs, envDoc{"NEXT_PUBLIC_SUPABASE_ANON_KEY", "[Supabase Dashboard](https://supabase.com/dashboard) > Settings > API"})
		docs = append(docs, envDoc{"SUPABASE_SERVICE_ROLE_KEY", "[Supabase Dashboard](https://supabase.com/dashboard) > Settings > API"})
	case "mongodb":
		docs = append(docs, envDoc{"MONGODB_URI", "[MongoDB Atlas](https://cloud.mongodb.com) > Connect > Drivers"})
	}

	switch cfg.Auth {
	case "clerk":
		docs = append(docs, envDoc{"NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY", "[Clerk Dashboard](https://dashboard.clerk.com) > API Keys"})
		docs = append(docs, envDoc{"CLERK_SECRET_KEY", "[Clerk Dashboard](https://dashboard.clerk.com) > API Keys"})
	case "nextauth":
		docs = append(docs, envDoc{"NEXTAUTH_URL", "Your app URL (http://localhost:3000 for dev)"})
		docs = append(docs, envDoc{"NEXTAUTH_SECRET", "Run: `openssl rand -base64 32`"})
	}

	if cfg.Payments == "stripe" {
		docs = append(docs, envDoc{"STRIPE_SECRET_KEY", "[Stripe Dashboard](https://dashboard.stripe.com) > Developers > API Keys"})
		docs = append(docs, envDoc{"NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY", "[Stripe Dashboard](https://dashboard.stripe.com) > Developers > API Keys"})
		docs = append(docs, envDoc{"STRIPE_WEBHOOK_SECRET", "[Stripe Dashboard](https://dashboard.stripe.com) > Developers > Webhooks"})
	}

	if cfg.Email == "resend" {
		docs = append(docs, envDoc{"RESEND_API_KEY", "[Resend Dashboard](https://resend.com/api-keys) > API Keys"})
	}

	return docs
}

type manualStep struct {
	name        string
	instruction string
}

func collectManualSteps(result *scaffold.Result) []manualStep {
	var steps []manualStep
	for _, s := range result.Steps {
		if s.Status == "manual" {
			steps = append(steps, manualStep{s.Name, s.Message})
		}
	}
	return steps
}

func devCommands(cfg *config.ProjectConfig) string {
	var b strings.Builder
	pm := cfg.PackageManager
	run := pm + " run"
	if pm == "bun" {
		run = "bun"
	}

	switch cfg.Framework {
	case "nextjs":
		b.WriteString(fmt.Sprintf("```bash\n%s dev      # Start development server\n%s build    # Production build\n%s start    # Start production server\n%s lint     # Run linter\n```\n\n", run, run, run, run))
	case "react-vite":
		b.WriteString(fmt.Sprintf("```bash\n%s dev      # Start development server\n%s build    # Production build\n%s preview  # Preview production build\n```\n\n", run, run, run))
	case "sveltekit":
		b.WriteString(fmt.Sprintf("```bash\n%s dev      # Start development server\n%s build    # Production build\n%s preview  # Preview production build\n```\n\n", run, run, run))
	case "nuxt":
		b.WriteString(fmt.Sprintf("```bash\n%s dev      # Start development server\n%s build    # Production build\n%s preview  # Preview production build\n```\n\n", run, run, run))
	case "django":
		b.WriteString("```bash\npython manage.py runserver   # Start dev server\npython manage.py migrate     # Run migrations\npython manage.py test        # Run tests\n```\n\n")
	case "fastapi":
		b.WriteString("```bash\nuvicorn main:app --reload    # Start dev server\npytest                       # Run tests\n```\n\n")
	case "rails":
		b.WriteString("```bash\nrails server                 # Start dev server\nrails db:migrate             # Run migrations\nrails test                   # Run tests\n```\n\n")
	case "laravel":
		b.WriteString("```bash\nphp artisan serve            # Start dev server\nphp artisan migrate          # Run migrations\nphp artisan test             # Run tests\n```\n\n")
	default:
		b.WriteString(fmt.Sprintf("```bash\n%s dev      # Start development server\n%s build    # Production build\n```\n\n", run, run))
	}

	return b.String()
}

func deploymentDocs(cfg *config.ProjectConfig) string {
	switch cfg.Deployment {
	case "vercel":
		return "This project is configured for **Vercel** deployment.\n\n1. Push to GitHub (auto-deploys if linked)\n2. Or run: `vercel --prod`\n\n"
	case "railway":
		return "Deploy to **Railway**:\n\n1. Go to [railway.app](https://railway.app)\n2. New Project > Deploy from GitHub repo\n3. Add environment variables from `.env.example`\n4. Railway auto-detects the framework and builds\n\n"
	case "cloudflare":
		return "Deploy to **Cloudflare Pages**:\n\n1. Go to [Cloudflare Dashboard](https://dash.cloudflare.com) > Pages\n2. Create a project > Connect to Git\n3. Set build command and output directory\n4. Add environment variables from `.env.example`\n\n"
	case "docker":
		return "Deploy with **Docker**:\n\n1. Build: `docker build -t " + cfg.Name + " .`\n2. Run: `docker run -p 3000:3000 " + cfg.Name + "`\n3. Or use docker-compose: `docker-compose up`\n\n"
	default:
		return "No specific deployment target configured. Choose a hosting provider and follow their deployment guide.\n\n"
	}
}

func architectureNotes(cfg *config.ProjectConfig) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("This is a **%s** project", cfg.Framework))
	if cfg.Language != "" {
		b.WriteString(fmt.Sprintf(" using **%s**", cfg.Language))
	}
	b.WriteString(".\n\n")

	if cfg.Database != "" && cfg.Database != "none" {
		b.WriteString(fmt.Sprintf("- **Database:** %s\n", cfg.Database))
	}
	if cfg.Auth != "" && cfg.Auth != "none" {
		b.WriteString(fmt.Sprintf("- **Auth:** %s — handles user authentication and session management\n", cfg.Auth))
	}
	if cfg.Payments != "" && cfg.Payments != "none" {
		b.WriteString(fmt.Sprintf("- **Payments:** %s — handles payment processing and subscriptions\n", cfg.Payments))
	}
	if cfg.Email != "" && cfg.Email != "none" {
		b.WriteString(fmt.Sprintf("- **Email:** %s — handles transactional email\n", cfg.Email))
	}
	b.WriteString("\n")

	return b.String()
}
```

- [ ] **Step 2: Commit**

```bash
git add internal/spec/
git commit -m "feat: PROJECT_SPEC.md generator with env docs and manual steps"
```

---

### Task 8: Wire Everything Together in main.go

**Files:**
- Modify: `main.go`
- Create: `internal/tui/progress.go`

- [ ] **Step 1: Create progress display**

Create `internal/tui/progress.go`:

```go
package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/JadenB9/projectmaker/internal/scaffold"
)

func PrintResults(result *scaffold.Result, specPath string) {
	fmt.Println()
	fmt.Println(titleStyle.Render("Setup Complete"))
	fmt.Println()

	doneIcon := lipgloss.NewStyle().Foreground(secondary).Render("[done]")
	manualIcon := lipgloss.NewStyle().Foreground(accent).Render("[manual]")
	skipIcon := lipgloss.NewStyle().Foreground(errColor).Render("[skip]")

	for _, step := range result.Steps {
		var icon string
		switch step.Status {
		case "done":
			icon = doneIcon
		case "manual":
			icon = manualIcon
		case "skipped":
			icon = skipIcon
		}

		line := fmt.Sprintf("  %s %s", icon, step.Name)
		if step.Message != "" && step.Status != "done" {
			line += "\n" + lipgloss.NewStyle().Foreground(dimColor).PaddingLeft(10).Render(step.Message)
		}
		fmt.Println(line)
	}

	fmt.Println()

	// Count manual steps
	var manualCount int
	for _, s := range result.Steps {
		if s.Status == "manual" {
			manualCount++
		}
	}

	if manualCount > 0 {
		fmt.Println(accentStyle.Render(fmt.Sprintf("  %d manual step(s) remaining — see PROJECT_SPEC.md for details", manualCount)))
	}

	fmt.Println()
	fmt.Println(successStyle.Render("  PROJECT_SPEC.md generated at: " + specPath))
	fmt.Println()

	// Helpful next steps
	var nextSteps []string
	nextSteps = append(nextSteps, "cd "+result.Steps[0].Name) // project dir
	nextSteps = append(nextSteps, "Open PROJECT_SPEC.md for full setup details")

	fmt.Println(lipgloss.NewStyle().Foreground(dimColor).Render("  Next steps:"))
	for _, s := range nextSteps {
		fmt.Println(lipgloss.NewStyle().Foreground(textColor).PaddingLeft(4).Render("$ " + s))
	}
	fmt.Println()
}

func PrintError(msg string) {
	fmt.Println()
	fmt.Println(errorStyle.Render("  Error: " + msg))
	fmt.Println()
}

func PrintCancelled() {
	fmt.Println()
	fmt.Println(lipgloss.NewStyle().Foreground(dimColor).Render("  Project creation cancelled."))
	fmt.Println()
}
```

- [ ] **Step 2: Update main.go to wire everything**

Update `main.go`:

```go
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/JadenB9/projectmaker/internal/scaffold"
	"github.com/JadenB9/projectmaker/internal/services"
	"github.com/JadenB9/projectmaker/internal/spec"
	"github.com/JadenB9/projectmaker/internal/tui"
)

func main() {
	cfg, err := tui.Run()
	if err != nil {
		tui.PrintError(err.Error())
		os.Exit(1)
	}
	if cfg == nil {
		tui.PrintCancelled()
		os.Exit(0)
	}

	// Detect available CLIs
	clis := services.DetectCLIs()

	// Run scaffold
	result, err := scaffold.Run(cfg, clis)
	if err != nil {
		tui.PrintError(err.Error())
		os.Exit(1)
	}

	// Generate PROJECT_SPEC.md
	if err := spec.Generate(cfg, result); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to generate PROJECT_SPEC.md: %v\n", err)
	}

	specPath := filepath.Join(cfg.ProjectDir, "PROJECT_SPEC.md")
	tui.PrintResults(result, specPath)
}
```

- [ ] **Step 3: Commit**

```bash
git add main.go internal/tui/progress.go
git commit -m "feat: wire TUI, scaffold, and spec generator together"
```

---

### Task 9: Build, Test, and Install

- [ ] **Step 1: Build and fix any compilation errors**

```bash
cd "/Users/jaden/Library/Mobile Documents/com~apple~CloudDocs/Projects/projectmaker"
go build -o project .
```

- [ ] **Step 2: Test the binary locally**

```bash
./project
```

Walk through the TUI, select a preset, confirm, verify output.

- [ ] **Step 3: Install globally**

```bash
go install .
```

Verify `project` command works from any directory.

- [ ] **Step 4: Commit final state**

```bash
git add .
git commit -m "feat: projectmaker v1 — TUI project scaffolder"
```
