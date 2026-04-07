# ProjectMaker Design Spec

## Overview

A Go CLI tool that launches an elegant TUI for scaffolding full-stack project stacks. Invoked by typing `project` anywhere in the terminal. Automates GitHub repo creation, Vercel linking, dependency installation, and environment setup. Generates a `PROJECT_SPEC.md` as the single source of truth for every project.

## Tech Stack

- **Language:** Go
- **TUI Framework:** Bubble Tea + Lip Gloss (styling) + Huh (forms)
- **Templates:** Go `embed` for bundled file templates
- **Distribution:** `go install` / goreleaser for cross-platform binaries
- **Binary name:** `project`

## Design Principles

- Calm, muted color palette — no flashy neon, just clean and readable
- Single binary, zero runtime dependencies
- Auto-detect installed tools and smart defaults
- Always generate PROJECT_SPEC.md regardless of automation level
- Per-workspace package manager selection

## User Flow

1. Welcome screen with tool name
2. Project name input (defaults to current directory name)
3. Stack selection: preset OR custom
4. Layer-by-layer selection (custom mode or override preset defaults):
   - Language/Runtime
   - Framework
   - Frontend styling
   - Backend/API approach
   - Database/ORM
   - Auth provider
   - Payments
   - Email
   - Package manager (per workspace, auto-detected defaults)
   - Deployment target
   - Extras (CI, Docker, n8n, Analytics)
5. Confirmation summary with option to go back
6. Execution: scaffold, create repo, link services, generate spec

## Preset Stacks

| # | Name | Components |
|---|------|------------|
| 1 | Next.js Full-Stack | Next.js App Router + TypeScript + Tailwind + shadcn/ui |
| 2 | T3 Stack | Next.js + tRPC + Prisma + NextAuth + Tailwind |
| 3 | MERN | MongoDB + Express + React + Node + TypeScript |
| 4 | Next.js + Supabase | Next.js + Supabase Auth/DB + Tailwind + shadcn/ui |
| 5 | Next.js + Convex | Next.js + Convex real-time backend + Tailwind |
| 6 | SvelteKit Full-Stack | SvelteKit + TypeScript + Tailwind + Drizzle |
| 7 | Nuxt Full-Stack | Nuxt 4 + TypeScript + Tailwind + Nitro |
| 8 | Rails + React | Ruby on Rails API + React SPA + TypeScript |
| 9 | Django + React | Django REST Framework + React SPA + TypeScript |
| 10 | FastAPI + React | FastAPI + React SPA + TypeScript |
| 11 | Go + HTMX | Go backend + HTMX + Tailwind |
| 12 | Laravel + Vue | Laravel + Vue 3 + Inertia + Tailwind |
| 13 | Custom Stack | Pick every layer individually |

## Selectable Options Per Layer

### Language/Runtime
- TypeScript, JavaScript, Python, Go, Rust, Ruby, PHP

### Framework
- **Frontend:** Next.js, React (Vite), Svelte/SvelteKit, Vue/Nuxt, Astro, Angular
- **Backend:** Express, Hono, Fastify, FastAPI, Django, Flask, Rails, Laravel, Gin, Fiber, Actix

### Frontend Styling
- Tailwind CSS, shadcn/ui (with Tailwind), CSS Modules, Styled Components, None

### Database / ORM
- Drizzle, Prisma, Supabase, MongoDB (Mongoose), SQLite, None

### Auth
- Clerk, NextAuth / Auth.js, Supabase Auth, Lucia, None

### Payments
- Stripe, None

### Email
- Resend, None

### Package Manager (per workspace)
- npm, pnpm, bun, yarn
- Auto-detects what's installed, defaults to first found

### Deployment Target
- Vercel, Railway, Cloudflare Pages/Workers, Docker (self-hosted), None

### Extras
- GitHub Actions CI template
- Dockerfile + docker-compose
- n8n (workflow automation, Docker service)
- Vercel Analytics
- PostHog Analytics
- ESLint + Prettier config

## Automation Tiers

### Fully Automated (if CLI detected)
- `git init` + `.gitignore`
- `gh repo create` — creates GitHub repo, sets remote
- `vercel link` — links to Vercel project (if Vercel selected)
- Package installation via chosen manager
- `.env` file with placeholder values
- Base project files (layout, config, etc.) from templates

### Semi-Automated (needs manual auth tokens)
- Clerk: creates `.env` with `NEXT_PUBLIC_CLERK_PUBLISHABLE_KEY` and `CLERK_SECRET_KEY` placeholders, documents where to get them
- Stripe: creates `.env` with `STRIPE_SECRET_KEY` and `STRIPE_WEBHOOK_SECRET` placeholders
- Supabase: `SUPABASE_URL` and `SUPABASE_ANON_KEY` placeholders
- Resend: `RESEND_API_KEY` placeholder
- Database URLs: connection string placeholders

### Spec-Only (documented in PROJECT_SPEC.md)
- Railway deployment setup
- Cloudflare Pages/Workers config
- n8n Docker service setup
- Any service where CLI automation isn't reliable

## PROJECT_SPEC.md Format

Always generated. Contains:

```markdown
# Project Specification

## Project Info
- Name, created date, directory

## Stack Summary
- Table of all chosen technologies

## Environment Variables
- Every required env var with description and where to get the value

## Setup Instructions (Manual Steps)
- Numbered list of anything that couldn't be automated
- Links to dashboards/docs for each service

## Architecture Notes
- How the pieces connect
- File structure overview

## Development Commands
- Dev server, build, test, lint commands

## Deployment
- How to deploy based on chosen target
```

## Project Structure (of the tool itself)

```
projectmaker/
├── main.go
├── cmd/
│   └── root.go            # Cobra CLI setup
├── internal/
│   ├── tui/
│   │   ├── app.go          # Main Bubble Tea app model
│   │   ├── welcome.go      # Welcome screen
│   │   ├── name.go         # Project name input
│   │   ├── stacks.go       # Stack selection list
│   │   ├── options.go      # Layer-by-layer selection
│   │   ├── confirm.go      # Confirmation summary
│   │   ├── progress.go     # Execution progress screen
│   │   └── styles.go       # Lip Gloss styles, color palette
│   ├── config/
│   │   └── config.go       # Project config struct
│   ├── stacks/
│   │   └── presets.go      # Preset stack definitions
│   ├── scaffold/
│   │   ├── scaffold.go     # Orchestrates project creation
│   │   ├── files.go        # File generation from templates
│   │   └── gitignore.go    # .gitignore generation
│   ├── services/
│   │   ├── github.go       # gh CLI integration
│   │   ├── vercel.go       # vercel CLI integration
│   │   ├── detect.go       # CLI availability detection
│   │   └── packages.go     # Package manager operations
│   ├── spec/
│   │   └── generator.go    # PROJECT_SPEC.md generator
│   └── templates/          # Embedded template files
│       ├── env/             # .env templates per service
│       ├── gitignore/       # .gitignore templates
│       ├── docker/          # Dockerfile, docker-compose templates
│       └── ci/              # GitHub Actions templates
└── go.mod
```

## Color Palette

Calm, muted tones:
- **Primary:** Soft blue (#7C9CBF)
- **Secondary:** Muted sage (#8FAE8B)
- **Accent:** Warm amber (#D4A574)
- **Text:** Light gray (#E0E0E0)
- **Dim text:** Medium gray (#888888)
- **Error:** Soft red (#C27171)
- **Background:** Terminal default (transparent)

## CLI Detection

On startup, detect which tools are available:
- `gh` — GitHub CLI
- `vercel` — Vercel CLI
- `npm`, `pnpm`, `bun`, `yarn` — package managers
- `docker` — Docker
- `python3`, `go`, `cargo`, `ruby`, `php` — language runtimes

This informs which options to show and what can be automated vs spec-only.

## Installation

```bash
go install github.com/JadenB9/projectmaker@latest
```

The binary is named `project`, so it's available globally as `project` in the terminal.
