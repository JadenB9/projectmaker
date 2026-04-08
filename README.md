# projectmaker

A TUI tool for scaffolding full-stack projects. Type `project` to create a new app with everything connected.

## Install

```bash
brew install JadenB9/tap/project
```

Or with Go:

```bash
go install github.com/JadenB9/projectmaker@latest
```

## Usage

```bash
project              # Create a new project
project remove       # Delete a project (local + GitHub)
project help         # Show all commands
project version      # Show version
```

## What it does

- **12 preset stacks** — Next.js, T3, MERN, Supabase, Convex, SvelteKit, Nuxt, Rails, Django, FastAPI, Go+HTMX, Laravel
- **Custom stack builder** — pick every layer: language, framework, styling, DB, auth, payments, email, package manager, deployment
- **23 languages** — TypeScript, Python, Go, Rust, Java, Kotlin, Swift, and more
- **Auto-scaffolding** — runs `create-next-app`, `create-vite`, etc. interactively
- **GitHub integration** — creates repo, sets remote, validates name availability
- **Vercel integration** — links project after scaffolding
- **Auth checks** — verifies GitHub and Vercel are authenticated before building
- **Multi-deploy** — select multiple targets (Vercel + Railway + Cloudflare)
- **Environment setup** — generates `.env` with placeholders and `.env.example`
- **PROJECT_SPEC.md** — always generated as the source of truth for AI or human setup
- **Project removal** — `project remove` deletes local dir + GitHub repo with confirmation

## How it works

1. Name your project (auto-lowercase, checks GitHub for conflicts)
2. Pick a preset stack or build custom
3. Select deployment targets and extras
4. Confirm your choices
5. Auth check — verifies connections, offers to log in
6. Scaffold — steps animate as they complete
7. Press Enter — clears screen, drops you into the project directory

## Requirements

- **Required:** Go 1.21+ (only for `go install` method)
- **Optional:** `gh` (GitHub CLI), `vercel`, `bun`/`pnpm`/`npm`/`yarn`

The tool auto-detects what's installed and adjusts accordingly.
