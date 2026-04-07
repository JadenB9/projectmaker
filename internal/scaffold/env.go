package scaffold

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/JadenB9/projectmaker/internal/config"
)

// EnvVar represents an environment variable entry for .env files.
type EnvVar struct {
	Key         string
	Placeholder string
	Comment     string
}

// collectEnvVars returns environment variables based on the project config.
func collectEnvVars(cfg *config.ProjectConfig) []EnvVar {
	var vars []EnvVar

	// Database vars
	switch cfg.Database {
	case "prisma", "drizzle":
		vars = append(vars, EnvVar{
			Key:         "DATABASE_URL",
			Placeholder: "postgresql://user:password@localhost:5432/mydb",
			Comment:     "Database connection string",
		})
	case "supabase":
		vars = append(vars, EnvVar{
			Key:         "NEXT_PUBLIC_SUPABASE_URL",
			Placeholder: "https://your-project.supabase.co",
			Comment:     "Supabase project URL",
		})
		vars = append(vars, EnvVar{
			Key:         "NEXT_PUBLIC_SUPABASE_ANON_KEY",
			Placeholder: "your-anon-key",
			Comment:     "Supabase anonymous key",
		})
		vars = append(vars, EnvVar{
			Key:         "SUPABASE_SERVICE_ROLE_KEY",
			Placeholder: "your-service-role-key",
			Comment:     "Supabase service role key (server-side only)",
		})
	case "mongodb-mongoose":
		vars = append(vars, EnvVar{
			Key:         "MONGODB_URI",
			Placeholder: "mongodb://localhost:27017/mydb",
			Comment:     "MongoDB connection string",
		})
	}

	// Auth vars
	switch cfg.Auth {
	case "clerk":
		vars = append(vars, EnvVar{
			Key:         "NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY",
			Placeholder: "pk_test_...",
			Comment:     "Clerk publishable key",
		})
		vars = append(vars, EnvVar{
			Key:         "CLERK_SECRET_KEY",
			Placeholder: "sk_test_...",
			Comment:     "Clerk secret key",
		})
	case "nextauth":
		vars = append(vars, EnvVar{
			Key:         "NEXTAUTH_URL",
			Placeholder: "http://localhost:3000",
			Comment:     "NextAuth base URL",
		})
		vars = append(vars, EnvVar{
			Key:         "NEXTAUTH_SECRET",
			Placeholder: "your-secret-here",
			Comment:     "NextAuth secret (generate with: openssl rand -base64 32)",
		})
	}

	// Payment vars
	if cfg.Payments == "stripe" {
		vars = append(vars, EnvVar{
			Key:         "NEXT_PUBLIC_STRIPE_PUBLISHABLE_KEY",
			Placeholder: "pk_test_...",
			Comment:     "Stripe publishable key",
		})
		vars = append(vars, EnvVar{
			Key:         "STRIPE_SECRET_KEY",
			Placeholder: "sk_test_...",
			Comment:     "Stripe secret key",
		})
		vars = append(vars, EnvVar{
			Key:         "STRIPE_WEBHOOK_SECRET",
			Placeholder: "whsec_...",
			Comment:     "Stripe webhook signing secret",
		})
	}

	// Email vars
	if cfg.Email == "resend" {
		vars = append(vars, EnvVar{
			Key:         "RESEND_API_KEY",
			Placeholder: "re_...",
			Comment:     "Resend API key",
		})
	}

	return vars
}

// writeEnvFile writes both .env (with placeholders) and .env.example (empty values).
func writeEnvFile(cfg *config.ProjectConfig, vars []EnvVar) error {
	if len(vars) == 0 {
		return nil
	}

	var envLines, exampleLines []string

	for _, v := range vars {
		if v.Comment != "" {
			envLines = append(envLines, fmt.Sprintf("# %s", v.Comment))
			exampleLines = append(exampleLines, fmt.Sprintf("# %s", v.Comment))
		}
		envLines = append(envLines, fmt.Sprintf("%s=%s", v.Key, v.Placeholder))
		exampleLines = append(exampleLines, fmt.Sprintf("%s=", v.Key))
		envLines = append(envLines, "")
		exampleLines = append(exampleLines, "")
	}

	envPath := filepath.Join(cfg.ProjectDir, ".env")
	if err := os.WriteFile(envPath, []byte(strings.Join(envLines, "\n")+"\n"), 0644); err != nil {
		return fmt.Errorf("writing .env: %w", err)
	}

	examplePath := filepath.Join(cfg.ProjectDir, ".env.example")
	if err := os.WriteFile(examplePath, []byte(strings.Join(exampleLines, "\n")+"\n"), 0644); err != nil {
		return fmt.Errorf("writing .env.example: %w", err)
	}

	return nil
}
