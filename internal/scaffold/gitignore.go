package scaffold

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/JadenB9/projectmaker/internal/config"
)

// writeGitignore creates a .gitignore file tailored to the project config.
func writeGitignore(cfg *config.ProjectConfig) error {
	var lines []string

	// Common ignores
	lines = append(lines,
		"# Dependencies",
		"node_modules/",
		"",
		"# Environment",
		".env",
		".env.local",
		".env.*.local",
		"",
		"# Build output",
		"dist/",
		"build/",
		".next/",
		".nuxt/",
		".svelte-kit/",
		".output/",
		"",
		"# IDE",
		".idea/",
		".vscode/",
		"",
		"# OS",
		".DS_Store",
		"",
		"# Logs",
		"npm-debug.log*",
	)

	// Python
	if cfg.Language == "python" {
		lines = append(lines,
			"",
			"# Python",
			"venv/",
			"__pycache__/",
			"*.pyc",
			".pytest_cache/",
		)
	}

	// Go
	if cfg.Language == "go" {
		lines = append(lines,
			"",
			"# Go",
			"*.exe",
			"*.dll",
			"*.so",
			"*.dylib",
		)
	}

	// Rust
	if cfg.Language == "rust" {
		lines = append(lines,
			"",
			"# Rust",
			"target/",
			"Cargo.lock",
		)
	}

	// Docker
	if containsExtra(cfg.Extras, "dockerfile") {
		lines = append(lines,
			"",
			"# Docker",
			"docker-compose.override.yml",
		)
	}

	lines = append(lines, "")

	content := strings.Join(lines, "\n")
	return os.WriteFile(filepath.Join(cfg.ProjectDir, ".gitignore"), []byte(content), 0644)
}

func containsExtra(extras []string, target string) bool {
	for _, e := range extras {
		if e == target {
			return true
		}
	}
	return false
}
