package config

// Preset represents a pre-configured project stack.
type Preset struct {
	Name        string
	Description string
	UseCase     string
	Config      ProjectConfig
}

// Presets contains all available stack presets.
var Presets = []Preset{
	{
		Name:        "Next.js Full-Stack",
		Description: "Next.js App Router + TypeScript + Tailwind + shadcn/ui",
		UseCase:     "SaaS apps, dashboards, marketing sites, landing pages",
		Config: ProjectConfig{
			Stack:          "nextjs-fullstack",
			Language:       "typescript",
			Framework:      "nextjs",
			Styling:        "tailwind-shadcn",
			PackageManager: "pnpm",
			Deployment:     []string{"vercel"},
		},
	},
	{
		Name:        "T3 Stack",
		Description: "Next.js + tRPC + Prisma + NextAuth + Tailwind",
		UseCase:     "Type-safe full-stack apps with end-to-end typesafety",
		Config: ProjectConfig{
			Stack:          "t3",
			Language:       "typescript",
			Framework:      "nextjs",
			Styling:        "tailwind",
			Database:       "prisma",
			Auth:           "nextauth",
			PackageManager: "pnpm",
			Deployment:     []string{"vercel"},
		},
	},
	{
		Name:        "MERN Stack",
		Description: "MongoDB + Express + React + Node + TypeScript",
		UseCase:     "REST APIs, real-time apps, flexible NoSQL data models",
		Config: ProjectConfig{
			Stack:          "mern",
			Language:       "typescript",
			Framework:      "react-vite",
			Backend:        "express",
			Database:       "mongodb-mongoose",
			PackageManager: "npm",
		},
	},
	{
		Name:        "Next.js + Supabase",
		Description: "Next.js + Supabase Auth/DB + Tailwind + shadcn/ui",
		UseCase:     "Apps needing auth, Postgres, and real-time out of the box",
		Config: ProjectConfig{
			Stack:          "nextjs-supabase",
			Language:       "typescript",
			Framework:      "nextjs",
			Styling:        "tailwind-shadcn",
			Database:       "supabase",
			Auth:           "supabase-auth",
			PackageManager: "pnpm",
			Deployment:     []string{"vercel"},
		},
	},
	{
		Name:        "Next.js + Convex",
		Description: "Next.js + Convex real-time backend + Tailwind",
		UseCase:     "Real-time collaborative apps, live dashboards, chat apps",
		Config: ProjectConfig{
			Stack:          "nextjs-convex",
			Language:       "typescript",
			Framework:      "nextjs",
			Styling:        "tailwind",
			PackageManager: "pnpm",
			Deployment:     []string{"vercel"},
		},
	},
	{
		Name:        "SvelteKit Full-Stack",
		Description: "SvelteKit + TypeScript + Tailwind + Drizzle",
		UseCase:     "Fast, lightweight apps with minimal JS bundle size",
		Config: ProjectConfig{
			Stack:          "sveltekit-fullstack",
			Language:       "typescript",
			Framework:      "sveltekit",
			Styling:        "tailwind",
			Database:       "drizzle",
			PackageManager: "pnpm",
		},
	},
	{
		Name:        "Nuxt Full-Stack",
		Description: "Nuxt 4 + TypeScript + Tailwind + Nitro",
		UseCase:     "Vue-based apps, SSR content sites, hybrid rendering",
		Config: ProjectConfig{
			Stack:          "nuxt-fullstack",
			Language:       "typescript",
			Framework:      "nuxt",
			Styling:        "tailwind",
			PackageManager: "pnpm",
		},
	},
	{
		Name:        "Rails + React",
		Description: "Ruby on Rails API + React SPA + TypeScript",
		UseCase:     "Rapid prototyping, CRUD-heavy apps, startups",
		Config: ProjectConfig{
			Stack:     "rails-react",
			Language:  "ruby",
			Framework: "rails",
			Backend:   "rails",
		},
	},
	{
		Name:        "Django + React",
		Description: "Django REST Framework + React SPA + TypeScript",
		UseCase:     "Data-driven apps, admin panels, ML/AI backends",
		Config: ProjectConfig{
			Stack:     "django-react",
			Language:  "python",
			Framework: "django",
			Backend:   "django",
		},
	},
	{
		Name:        "FastAPI + React",
		Description: "FastAPI + React SPA + TypeScript",
		UseCase:     "High-performance APIs, async backends, ML model serving",
		Config: ProjectConfig{
			Stack:     "fastapi-react",
			Language:  "python",
			Framework: "fastapi",
			Backend:   "fastapi",
		},
	},
	{
		Name:        "Go + HTMX",
		Description: "Go backend + HTMX + Tailwind",
		UseCase:     "Server-rendered apps, minimal JS, high concurrency",
		Config: ProjectConfig{
			Stack:     "go-htmx",
			Language:  "go",
			Framework: "gin",
			Styling:   "tailwind",
			Backend:   "go",
		},
	},
	{
		Name:        "Laravel + Vue",
		Description: "Laravel + Vue 3 + Inertia + Tailwind",
		UseCase:     "PHP monoliths, CMS-backed sites, e-commerce",
		Config: ProjectConfig{
			Stack:     "laravel-vue",
			Language:  "php",
			Framework: "laravel",
			Backend:   "laravel",
			Styling:   "tailwind",
		},
	},
}
