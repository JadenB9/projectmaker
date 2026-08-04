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

- **13 preset stacks** — Blank Project, Next.js, T3, MERN, Supabase, Convex, SvelteKit, Nuxt, Rails, Django, FastAPI, Go+HTMX, Laravel
- **Blank Project mode** — just GitHub + .gitignore + .env + PROJECT_SPEC.md, no runtime setup
- **Custom stack builder** — pick every layer: language, framework, styling, DB, auth, payments, email, package manager, deployment
- **23 languages** — TypeScript, Python, Go, Rust, Java, Kotlin, Swift, and more
- **Auto-scaffolding** — runs `create-next-app`, `create-vite`, etc. interactively
- **GitHub integration** — creates repo, sets remote, validates name availability
- **Vercel integration** — links project interactively after scaffolding
- **Auth checks** — verifies GitHub and Vercel are authenticated, offers to log in
- **Multi-deploy** — select multiple targets (Vercel + Railway + Cloudflare)
- **Environment setup** — generates `.env` with placeholders and `.env.example`
- **PROJECT_SPEC.md** — source of truth with manual steps at top, AI guidelines included
- **Project removal** — `project remove` deletes local dir + GitHub repo with confirmation
- **Jump to a Projects folder** — press Up at the name box to skip creating and drop straight into one of your Projects folders
- **Descriptions everywhere** — every option explains what it does, beginner-friendly

## Jumping into a Projects folder

`project` still opens on the name box, so typing works exactly as before. Press
**Up** there and it swaps to a folder picker listing your Projects directories —
the iCloud ones (`Projects`, `Projects2026`, ...) plus the local `~/Projects` on
this Mac, each with a note and project count:

```
  Project Directory

  > Projects           iCloud — main project archive · 30 projects
    Projects2026       iCloud — 2026 projects · 1 project
    Projects (Local)   This Mac — local only, not synced · 0 projects
```

Arrow keys scroll, Enter drops you into that folder, Esc goes back to the name
box. This mirrors the Project Directory menu in Mode Terminal.

## How it works

1. Name your project (auto-lowercase, checks GitHub for conflicts) — or press Up to jump into an existing Projects folder
2. Pick a preset stack or build custom (Blank Project for GitHub-only setup)
3. Select deployment targets and extras
4. Confirm your choices
5. Auth check — verifies connections, offers to log in
6. Scaffold — steps animate as they complete
7. Press Enter — clears screen, drops you into the project directory

## PROJECT_SPEC.md

Every project gets a `PROJECT_SPEC.md` that includes:

- **Manual setup steps** at the top so you see them first
- Stack summary, environment variables, dev commands, deployment docs
- **AI guidelines** — instructions for any AI tool working in the project:
  - No co-authoring or AI attribution in commits
  - Code should look student-written, not AI-generated
  - Security-first development practices

## Requirements

- **Required:** Go 1.21+ (only for `go install` method)
- **Optional:** `gh` (GitHub CLI), `vercel`, `bun`/`pnpm`/`npm`/`yarn`

The tool auto-detects what's installed and adjusts accordingly.
