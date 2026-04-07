package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/JadenB9/projectmaker/internal/config"
	"github.com/JadenB9/projectmaker/internal/scaffold"
	"github.com/JadenB9/projectmaker/internal/services"
	"github.com/JadenB9/projectmaker/internal/spec"
)

func TestCLIDetection(t *testing.T) {
	clis := services.DetectCLIs()
	// On this machine we know git and node exist
	if !clis.Git {
		t.Error("expected git to be detected")
	}
	if !clis.Node {
		t.Error("expected node to be detected")
	}
	t.Logf("CLIs: git=%v gh=%v vercel=%v bun=%v npm=%v pnpm=%v docker=%v",
		clis.Git, clis.GitHub, clis.Vercel, clis.Bun, clis.Npm, clis.Pnpm, clis.Docker)
}

func TestPresets(t *testing.T) {
	if len(config.Presets) != 12 {
		t.Errorf("expected 12 presets, got %d", len(config.Presets))
	}
	for _, p := range config.Presets {
		if p.Name == "" {
			t.Error("preset has empty name")
		}
		if p.Description == "" {
			t.Errorf("preset %q has empty description", p.Name)
		}
		if p.Config.Language == "" {
			t.Errorf("preset %q has empty language", p.Name)
		}
		if p.Config.Framework == "" {
			t.Errorf("preset %q has empty framework", p.Name)
		}
	}
}

func TestOptions(t *testing.T) {
	checks := map[string][]config.Option{
		"Languages":       config.Languages,
		"Frameworks":      config.Frameworks,
		"Styling":         config.Styling,
		"Databases":       config.Databases,
		"AuthProviders":   config.AuthProviders,
		"PaymentProviders": config.PaymentProviders,
		"EmailProviders":  config.EmailProviders,
		"PackageManagers": config.PackageManagers,
		"DeploymentTargets": config.DeploymentTargets,
		"ExtraOptions":    config.ExtraOptions,
	}
	for name, opts := range checks {
		if len(opts) == 0 {
			t.Errorf("%s has no options", name)
		}
		for _, o := range opts {
			if o.Label == "" || o.Value == "" {
				t.Errorf("%s has option with empty label or value", name)
			}
		}
	}
}

func TestGitignoreGeneration(t *testing.T) {
	tmpDir := t.TempDir()
	cfg := &config.ProjectConfig{
		Language:   "typescript",
		ProjectDir: tmpDir,
	}

	// Run scaffold just for gitignore by creating the dir
	clis := services.CLIStatus{Git: false} // disable git to avoid init
	result, err := scaffold.Run(cfg, clis, nil)
	if err != nil {
		t.Fatalf("scaffold.Run failed: %v", err)
	}
	_ = result

	gitignorePath := filepath.Join(tmpDir, ".gitignore")
	data, err := os.ReadFile(gitignorePath)
	if err != nil {
		t.Fatalf("failed to read .gitignore: %v", err)
	}

	content := string(data)
	if !strings.Contains(content, "node_modules/") {
		t.Error(".gitignore missing node_modules/")
	}
	if !strings.Contains(content, ".env") {
		t.Error(".gitignore missing .env")
	}
	if !strings.Contains(content, ".DS_Store") {
		t.Error(".gitignore missing .DS_Store")
	}
}

func TestEnvGeneration(t *testing.T) {
	tmpDir := t.TempDir()
	cfg := &config.ProjectConfig{
		Language:   "typescript",
		Framework:  "nextjs",
		Database:   "drizzle",
		Auth:       "clerk",
		Payments:   "stripe",
		Email:      "resend",
		ProjectDir: tmpDir,
	}

	clis := services.CLIStatus{Git: false}
	_, err := scaffold.Run(cfg, clis, nil)
	if err != nil {
		t.Fatalf("scaffold.Run failed: %v", err)
	}

	// Check .env
	envPath := filepath.Join(tmpDir, ".env")
	data, err := os.ReadFile(envPath)
	if err != nil {
		t.Fatalf("failed to read .env: %v", err)
	}

	content := string(data)
	for _, key := range []string{"DATABASE_URL", "CLERK_SECRET_KEY", "STRIPE_SECRET_KEY", "RESEND_API_KEY"} {
		if !strings.Contains(content, key) {
			t.Errorf(".env missing %s", key)
		}
	}

	// Check .env.example
	examplePath := filepath.Join(tmpDir, ".env.example")
	if _, err := os.Stat(examplePath); os.IsNotExist(err) {
		t.Error(".env.example not created")
	}
}

func TestSpecGeneration(t *testing.T) {
	tmpDir := t.TempDir()
	cfg := &config.ProjectConfig{
		Name:           "test-project",
		Stack:          "Next.js Full-Stack",
		Language:       "typescript",
		Framework:      "nextjs",
		Styling:        "tailwind-shadcn",
		Database:       "drizzle",
		Auth:           "clerk",
		Payments:       "stripe",
		Email:          "resend",
		PackageManager: "bun",
		Deployment:     []string{"vercel"},
		Extras:         []string{"github-actions"},
		ProjectDir:     tmpDir,
	}

	result := &scaffold.Result{
		Steps: []scaffold.StepResult{
			{Name: "Create directory", Status: "done"},
			{Name: "Scaffold Next.js", Status: "manual", Message: "Run: bunx create-next-app@latest"},
			{Name: "Create GitHub repo", Status: "done"},
		},
	}

	err := spec.Generate(cfg, result)
	if err != nil {
		t.Fatalf("spec.Generate failed: %v", err)
	}

	specPath := filepath.Join(tmpDir, "PROJECT_SPEC.md")
	data, err := os.ReadFile(specPath)
	if err != nil {
		t.Fatalf("failed to read PROJECT_SPEC.md: %v", err)
	}

	content := string(data)

	// Check key sections exist
	for _, section := range []string{
		"# Project Specification",
		"## Project Info",
		"## Stack Summary",
		"## Environment Variables",
		"## Manual Setup Steps",
		"## Development Commands",
		"## Deployment",
		"## Architecture Notes",
	} {
		if !strings.Contains(content, section) {
			t.Errorf("PROJECT_SPEC.md missing section: %s", section)
		}
	}

	// Check specific content
	if !strings.Contains(content, "test-project") {
		t.Error("PROJECT_SPEC.md missing project name")
	}
	if !strings.Contains(content, "Clerk Dashboard") {
		t.Error("PROJECT_SPEC.md missing Clerk dashboard link")
	}
	if !strings.Contains(content, "Stripe Dashboard") {
		t.Error("PROJECT_SPEC.md missing Stripe dashboard link")
	}
	if !strings.Contains(content, "bunx create-next-app") {
		t.Error("PROJECT_SPEC.md missing manual step for Next.js scaffold")
	}
	if !strings.Contains(content, "bun dev") {
		t.Error("PROJECT_SPEC.md missing bun dev command")
	}

	t.Logf("PROJECT_SPEC.md: %d bytes, all sections present", len(data))
}
