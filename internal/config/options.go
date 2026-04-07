package config

// Option represents a selectable choice in the TUI.
type Option struct {
	Label string
	Value string
}

// Languages lists all supported programming languages.
var Languages = []Option{
	{Label: "TypeScript", Value: "typescript"},
	{Label: "JavaScript", Value: "javascript"},
	{Label: "Python", Value: "python"},
	{Label: "Go", Value: "go"},
	{Label: "Rust", Value: "rust"},
	{Label: "Ruby", Value: "ruby"},
	{Label: "PHP", Value: "php"},
}

// Frameworks lists all supported frontend and backend frameworks.
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

// Styling lists all supported styling/CSS options.
var Styling = []Option{
	{Label: "Tailwind CSS", Value: "tailwind"},
	{Label: "Tailwind + shadcn/ui", Value: "tailwind-shadcn"},
	{Label: "CSS Modules", Value: "css-modules"},
	{Label: "Styled Components", Value: "styled-components"},
	{Label: "None", Value: "none"},
}

// Databases lists all supported database/ORM options.
var Databases = []Option{
	{Label: "Drizzle ORM", Value: "drizzle"},
	{Label: "Prisma", Value: "prisma"},
	{Label: "Supabase", Value: "supabase"},
	{Label: "MongoDB (Mongoose)", Value: "mongodb-mongoose"},
	{Label: "SQLite", Value: "sqlite"},
	{Label: "None", Value: "none"},
}

// AuthProviders lists all supported authentication providers.
var AuthProviders = []Option{
	{Label: "Clerk", Value: "clerk"},
	{Label: "NextAuth / Auth.js", Value: "nextauth"},
	{Label: "Supabase Auth", Value: "supabase-auth"},
	{Label: "Lucia", Value: "lucia"},
	{Label: "None", Value: "none"},
}

// PaymentProviders lists all supported payment providers.
var PaymentProviders = []Option{
	{Label: "Stripe", Value: "stripe"},
	{Label: "None", Value: "none"},
}

// EmailProviders lists all supported email providers.
var EmailProviders = []Option{
	{Label: "Resend", Value: "resend"},
	{Label: "None", Value: "none"},
}

// PackageManagers lists all supported package managers.
var PackageManagers = []Option{
	{Label: "npm", Value: "npm"},
	{Label: "pnpm", Value: "pnpm"},
	{Label: "bun", Value: "bun"},
	{Label: "yarn", Value: "yarn"},
}

// DeploymentTargets lists all supported deployment targets.
var DeploymentTargets = []Option{
	{Label: "Vercel", Value: "vercel"},
	{Label: "Railway", Value: "railway"},
	{Label: "Cloudflare", Value: "cloudflare"},
	{Label: "Docker (self-hosted)", Value: "docker"},
	{Label: "None", Value: "none"},
}

// ExtraOptions lists all optional extras that can be added to a project.
var ExtraOptions = []Option{
	{Label: "GitHub Actions CI", Value: "github-actions"},
	{Label: "Dockerfile + docker-compose", Value: "dockerfile"},
	{Label: "n8n (workflow automation)", Value: "n8n"},
	{Label: "Vercel Analytics", Value: "vercel-analytics"},
	{Label: "PostHog Analytics", Value: "posthog"},
	{Label: "ESLint + Prettier", Value: "eslint-prettier"},
}
