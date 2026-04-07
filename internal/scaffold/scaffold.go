package scaffold

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/JadenB9/projectmaker/internal/config"
	"github.com/JadenB9/projectmaker/internal/services"
)

// Result holds the outcome of the full scaffold operation.
type Result struct {
	Steps    []StepResult
	SpecPath string
}

// StepResult represents the outcome of a single scaffold step.
type StepResult struct {
	Name    string
	Status  string // "done", "skipped", "manual"
	Message string
}

// StepCallback is called after each step completes, for live progress display.
type StepCallback func(step StepResult)

// Run orchestrates the full project scaffold process.
// If onStep is non-nil, it's called after each step completes.
func Run(cfg *config.ProjectConfig, clis services.CLIStatus, onStep StepCallback) (*Result, error) {
	var steps []StepResult

	addStep := func(s StepResult) {
		steps = append(steps, s)
		if onStep != nil {
			onStep(s)
		}
	}

	// 1. Create project directory
	if err := os.MkdirAll(cfg.ProjectDir, 0755); err != nil {
		return nil, fmt.Errorf("creating project directory: %w", err)
	}
	addStep(StepResult{
		Name:    "Create directory",
		Status:  "done",
		Message: cfg.ProjectDir,
	})

	// 2. git init
	if clis.Git {
		if _, err := services.RunCmdInDir(cfg.ProjectDir, "git", "init"); err != nil {
			addStep(StepResult{
				Name:    "Git init",
				Status:  "manual",
				Message: "Run: git init",
			})
		} else {
			addStep(StepResult{
				Name:   "Git init",
				Status: "done",
			})
		}
	} else {
		addStep(StepResult{
			Name:    "Git init",
			Status:  "skipped",
			Message: "git not found",
		})
	}

	// 3. Write .gitignore
	if err := writeGitignore(cfg); err != nil {
		addStep(StepResult{
			Name:    "Write .gitignore",
			Status:  "manual",
			Message: fmt.Sprintf("Failed to write .gitignore: %v", err),
		})
	} else {
		addStep(StepResult{
			Name:   "Write .gitignore",
			Status: "done",
		})
	}

	// 4. Write .env and .env.example
	envVars := collectEnvVars(cfg)
	if len(envVars) > 0 {
		if err := writeEnvFile(cfg, envVars); err != nil {
			addStep(StepResult{
				Name:    "Write .env files",
				Status:  "manual",
				Message: fmt.Sprintf("Failed: %v", err),
			})
		} else {
			addStep(StepResult{
				Name:    "Write .env files",
				Status:  "done",
				Message: fmt.Sprintf("%d variables configured", len(envVars)),
			})
		}
	} else {
		addStep(StepResult{
			Name:    "Write .env files",
			Status:  "skipped",
			Message: "No env vars needed",
		})
	}

	// 5. Scaffold framework
	frameworkSteps := scaffoldFramework(cfg, clis)
	for _, s := range frameworkSteps {
		addStep(s)
	}

	// 6. Install additional deps
	depSteps := installDeps(cfg, clis)
	for _, s := range depSteps {
		addStep(s)
	}

	// 7. Create GitHub repo
	if clis.GitHub {
		_, err := services.RunCmdInDir(cfg.ProjectDir, "gh", "repo", "create", cfg.Name, "--private", "--source=.", "--remote=origin")
		if err != nil {
			addStep(StepResult{
				Name:    "Create GitHub repo",
				Status:  "manual",
				Message: fmt.Sprintf("Run: gh repo create %s --private --source=. --remote=origin", cfg.Name),
			})
		} else {
			addStep(StepResult{
				Name:   "Create GitHub repo",
				Status: "done",
			})
		}
	} else {
		addStep(StepResult{
			Name:    "Create GitHub repo",
			Status:  "skipped",
			Message: "gh CLI not found",
		})
	}

	// 8. Link deployment targets
	for _, dep := range cfg.Deployment {
		switch dep {
		case "vercel":
			if clis.Vercel {
				_, err := services.RunCmdInDir(cfg.ProjectDir, "vercel", "link", "--yes")
				if err != nil {
					addStep(StepResult{
						Name:    "Link Vercel",
						Status:  "manual",
						Message: "Run: vercel link --yes",
					})
				} else {
					addStep(StepResult{
						Name:   "Link Vercel",
						Status: "done",
					})
				}
			} else {
				addStep(StepResult{
					Name:    "Link Vercel",
					Status:  "skipped",
					Message: "vercel CLI not found",
				})
			}
		case "railway":
			addStep(StepResult{
				Name:    "Railway setup",
				Status:  "manual",
				Message: "Go to railway.app > New Project > Deploy from GitHub",
			})
		case "cloudflare":
			addStep(StepResult{
				Name:    "Cloudflare setup",
				Status:  "manual",
				Message: "Go to dash.cloudflare.com > Add site or Pages project",
			})
		case "docker":
			addStep(StepResult{
				Name:    "Docker setup",
				Status:  "manual",
				Message: "Add Dockerfile and docker-compose.yml to project root",
			})
		case "aws":
			addStep(StepResult{
				Name:    "AWS setup",
				Status:  "manual",
				Message: "Configure AWS CLI and deploy via your preferred service (Lambda, ECS, EC2)",
			})
		case "flyio":
			addStep(StepResult{
				Name:    "Fly.io setup",
				Status:  "manual",
				Message: "Run: fly launch (install flyctl first: curl -L https://fly.io/install.sh | sh)",
			})
		}
	}

	return &Result{Steps: steps}, nil
}

// scaffoldFramework dispatches to the appropriate framework scaffolder.
func scaffoldFramework(cfg *config.ProjectConfig, clis services.CLIStatus) []StepResult {
	switch cfg.Framework {
	case "nextjs":
		return scaffoldNextJS(cfg, clis)
	case "react-vite":
		return scaffoldViteReact(cfg, clis)
	case "sveltekit":
		return scaffoldSvelteKit(cfg, clis)
	case "nuxt":
		return scaffoldNuxt(cfg, clis)
	case "express":
		return scaffoldExpress(cfg, clis)
	case "rails":
		return scaffoldRails(cfg, clis)
	case "laravel":
		return scaffoldLaravel(cfg, clis)
	default:
		// Handle language-based scaffolding when no framework specified
		switch cfg.Language {
		case "python":
			return scaffoldPython(cfg, clis)
		case "go":
			return scaffoldGoBackend(cfg, clis)
		default:
			return []StepResult{{
				Name:    "Scaffold framework",
				Status:  "skipped",
				Message: fmt.Sprintf("No scaffolder for framework=%q language=%q", cfg.Framework, cfg.Language),
			}}
		}
	}
}

// scaffoldNextJS creates a Next.js project with create-next-app.
func scaffoldNextJS(cfg *config.ProjectConfig, clis services.CLIStatus) []StepResult {
	runner := pmRunner(cfg.PackageManager)
	args := []string{"create-next-app@latest", cfg.ProjectDir,
		"--ts", "--eslint", "--app", "--src-dir", "--import-alias", "@/*",
	}

	if cfg.Styling == "tailwind" || cfg.Styling == "tailwind-shadcn" {
		args = append(args, "--tailwind")
	} else {
		args = append(args, "--no-tailwind")
	}

	switch cfg.PackageManager {
	case "bun":
		args = append(args, "--use-bun")
	case "pnpm":
		args = append(args, "--use-pnpm")
	case "yarn":
		args = append(args, "--use-yarn")
	}

	var steps []StepResult

	// Use interactive mode so create-next-app can prompt the user
	err := services.RunInteractive("", runner, args...)
	if err != nil {
		steps = append(steps, StepResult{
			Name:    "Scaffold Next.js",
			Status:  "manual",
			Message: fmt.Sprintf("Run: %s %s", runner, joinArgs(args)),
		})
		return steps
	}
	steps = append(steps, StepResult{
		Name:   "Scaffold Next.js",
		Status: "done",
	})

	// shadcn init if needed
	if cfg.Styling == "tailwind-shadcn" {
		err := services.RunInteractive(cfg.ProjectDir, runner, "shadcn@latest", "init", "--yes")
		if err != nil {
			steps = append(steps, StepResult{
				Name:    "Init shadcn/ui",
				Status:  "manual",
				Message: fmt.Sprintf("Run: %s shadcn@latest init --yes", runner),
			})
		} else {
			steps = append(steps, StepResult{
				Name:   "Init shadcn/ui",
				Status: "done",
			})
		}
	}

	return steps
}

// scaffoldViteReact creates a Vite + React + TypeScript project.
func scaffoldViteReact(cfg *config.ProjectConfig, clis services.CLIStatus) []StepResult {
	runner := pmRunner(cfg.PackageManager)
	_, err := services.RunCmd(runner, "create-vite@latest", cfg.ProjectDir, "--template", "react-ts")
	if err != nil {
		return []StepResult{{
			Name:    "Scaffold React (Vite)",
			Status:  "manual",
			Message: fmt.Sprintf("Run: %s create-vite@latest %s --template react-ts", runner, cfg.ProjectDir),
		}}
	}
	return []StepResult{{
		Name:   "Scaffold React (Vite)",
		Status: "done",
	}}
}

// scaffoldSvelteKit creates a SvelteKit project.
func scaffoldSvelteKit(cfg *config.ProjectConfig, clis services.CLIStatus) []StepResult {
	runner := pmRunner(cfg.PackageManager)
	_, err := services.RunCmd(runner, "sv", "create", cfg.ProjectDir, "--template", "minimal", "--types", "ts")
	if err != nil {
		return []StepResult{{
			Name:    "Scaffold SvelteKit",
			Status:  "manual",
			Message: fmt.Sprintf("Run: %s sv create %s --template minimal --types ts", runner, cfg.ProjectDir),
		}}
	}
	return []StepResult{{
		Name:   "Scaffold SvelteKit",
		Status: "done",
	}}
}

// scaffoldNuxt creates a Nuxt project.
func scaffoldNuxt(cfg *config.ProjectConfig, clis services.CLIStatus) []StepResult {
	runner := pmRunner(cfg.PackageManager)
	_, err := services.RunCmd(runner, "nuxi@latest", "init", cfg.ProjectDir)
	if err != nil {
		return []StepResult{{
			Name:    "Scaffold Nuxt",
			Status:  "manual",
			Message: fmt.Sprintf("Run: %s nuxi@latest init %s", runner, cfg.ProjectDir),
		}}
	}
	return []StepResult{{
		Name:   "Scaffold Nuxt",
		Status: "done",
	}}
}

// scaffoldExpress creates an Express + TypeScript project.
func scaffoldExpress(cfg *config.ProjectConfig, clis services.CLIStatus) []StepResult {
	var steps []StepResult
	pm := cfg.PackageManager
	if pm == "" {
		pm = "npm"
	}

	// npm init -y
	_, err := services.RunCmdInDir(cfg.ProjectDir, pm, "init", "-y")
	if err != nil {
		return []StepResult{{
			Name:    "Scaffold Express",
			Status:  "manual",
			Message: fmt.Sprintf("Run: cd %s && %s init -y && %s install express typescript @types/express @types/node ts-node", cfg.ProjectDir, pm, pm),
		}}
	}
	steps = append(steps, StepResult{
		Name:   "Init package.json",
		Status: "done",
	})

	// Install express + ts deps
	installCmd := installVerb(pm)
	_, err = services.RunCmdInDir(cfg.ProjectDir, pm, installCmd, "express", "typescript", "@types/express", "@types/node", "ts-node")
	if err != nil {
		steps = append(steps, StepResult{
			Name:    "Install Express deps",
			Status:  "manual",
			Message: fmt.Sprintf("Run: %s %s express typescript @types/express @types/node ts-node", pm, installCmd),
		})
	} else {
		steps = append(steps, StepResult{
			Name:   "Install Express deps",
			Status: "done",
		})
	}

	return steps
}

// scaffoldPython creates a Python project with a virtual environment.
func scaffoldPython(cfg *config.ProjectConfig, clis services.CLIStatus) []StepResult {
	if !clis.Python {
		return []StepResult{{
			Name:    "Scaffold Python",
			Status:  "manual",
			Message: "Run: python3 -m venv venv",
		}}
	}

	_, err := services.RunCmdInDir(cfg.ProjectDir, "python3", "-m", "venv", "venv")
	if err != nil {
		return []StepResult{{
			Name:    "Scaffold Python",
			Status:  "manual",
			Message: fmt.Sprintf("Run: cd %s && python3 -m venv venv", cfg.ProjectDir),
		}}
	}
	return []StepResult{{
		Name:   "Create Python venv",
		Status: "done",
	}}
}

// scaffoldRails creates a Rails API project.
func scaffoldRails(cfg *config.ProjectConfig, clis services.CLIStatus) []StepResult {
	if !clis.Ruby {
		return []StepResult{{
			Name:    "Scaffold Rails",
			Status:  "manual",
			Message: "Run: rails new " + cfg.ProjectDir + " --api --database=postgresql",
		}}
	}

	_, err := services.RunCmd("rails", "new", cfg.ProjectDir, "--api", "--database=postgresql")
	if err != nil {
		return []StepResult{{
			Name:    "Scaffold Rails",
			Status:  "manual",
			Message: "Run: rails new " + cfg.ProjectDir + " --api --database=postgresql",
		}}
	}
	return []StepResult{{
		Name:   "Scaffold Rails",
		Status: "done",
	}}
}

// scaffoldLaravel creates a Laravel project.
func scaffoldLaravel(cfg *config.ProjectConfig, clis services.CLIStatus) []StepResult {
	if !clis.PHP {
		return []StepResult{{
			Name:    "Scaffold Laravel",
			Status:  "manual",
			Message: "Run: composer create-project laravel/laravel " + cfg.ProjectDir,
		}}
	}

	_, err := services.RunCmd("composer", "create-project", "laravel/laravel", cfg.ProjectDir)
	if err != nil {
		return []StepResult{{
			Name:    "Scaffold Laravel",
			Status:  "manual",
			Message: "Run: composer create-project laravel/laravel " + cfg.ProjectDir,
		}}
	}
	return []StepResult{{
		Name:   "Scaffold Laravel",
		Status: "done",
	}}
}

// scaffoldGoBackend creates a Go module project.
func scaffoldGoBackend(cfg *config.ProjectConfig, clis services.CLIStatus) []StepResult {
	if !clis.Go {
		return []StepResult{{
			Name:    "Scaffold Go backend",
			Status:  "manual",
			Message: fmt.Sprintf("Run: cd %s && go mod init %s", cfg.ProjectDir, cfg.Name),
		}}
	}

	_, err := services.RunCmdInDir(cfg.ProjectDir, "go", "mod", "init", cfg.Name)
	if err != nil {
		return []StepResult{{
			Name:    "Scaffold Go backend",
			Status:  "manual",
			Message: fmt.Sprintf("Run: cd %s && go mod init %s", cfg.ProjectDir, cfg.Name),
		}}
	}
	return []StepResult{{
		Name:   "Init Go module",
		Status: "done",
	}}
}

// installDeps installs additional npm packages based on selected services.
func installDeps(cfg *config.ProjectConfig, clis services.CLIStatus) []StepResult {
	var packages []string

	// Database packages
	switch cfg.Database {
	case "prisma":
		packages = append(packages, "prisma", "@prisma/client")
	case "drizzle":
		packages = append(packages, "drizzle-orm", "drizzle-kit")
	case "mongodb-mongoose":
		packages = append(packages, "mongoose")
	case "supabase":
		packages = append(packages, "@supabase/supabase-js")
	}

	// Auth packages
	switch cfg.Auth {
	case "clerk":
		packages = append(packages, "@clerk/nextjs")
	case "nextauth":
		packages = append(packages, "next-auth")
	case "lucia":
		packages = append(packages, "lucia")
	}

	// Payment packages
	if cfg.Payments == "stripe" {
		packages = append(packages, "stripe", "@stripe/stripe-js")
	}

	// Email packages
	if cfg.Email == "resend" {
		packages = append(packages, "resend")
	}

	if len(packages) == 0 {
		return nil
	}

	// Only install npm packages if we're in a JS/TS project
	if !isJSProject(cfg) {
		return []StepResult{{
			Name:    "Install dependencies",
			Status:  "skipped",
			Message: "Non-JS project; install packages manually",
		}}
	}

	pm := cfg.PackageManager
	if pm == "" {
		pm = "npm"
	}

	verb := installVerb(pm)
	args := append([]string{verb}, packages...)

	_, err := services.RunCmdInDir(cfg.ProjectDir, pm, args...)
	if err != nil {
		return []StepResult{{
			Name:    "Install dependencies",
			Status:  "manual",
			Message: fmt.Sprintf("Run: %s %s %s", pm, verb, joinArgs(packages)),
		}}
	}

	return []StepResult{{
		Name:    "Install dependencies",
		Status:  "done",
		Message: fmt.Sprintf("%d packages via %s", len(packages), pm),
	}}
}

// pmRunner returns the package runner command for a given package manager.
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

// installVerb returns the install subcommand for a given package manager.
func installVerb(pm string) string {
	switch pm {
	case "bun":
		return "add"
	case "pnpm":
		return "add"
	case "yarn":
		return "add"
	default:
		return "install"
	}
}

// isJSProject returns true if the project is JavaScript/TypeScript based.
func isJSProject(cfg *config.ProjectConfig) bool {
	switch cfg.Language {
	case "typescript", "javascript":
		return true
	}
	switch cfg.Framework {
	case "nextjs", "react-vite", "sveltekit", "nuxt", "express", "hono", "fastify", "astro", "angular":
		return true
	}
	return false
}

// joinArgs joins string slices into a space-separated string for display.
func joinArgs(args []string) string {
	result := ""
	for i, a := range args {
		if i > 0 {
			result += " "
		}
		result += a
	}
	return result
}

// EnsureProjectDir makes sure the project directory path is set on the config.
func EnsureProjectDir(cfg *config.ProjectConfig, baseDir string) {
	if cfg.ProjectDir == "" {
		cfg.ProjectDir = filepath.Join(baseDir, cfg.Name)
	}
}
