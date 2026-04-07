package config

// Option represents a selectable choice in the TUI.
type Option struct {
	Label string
	Value string
}

// Languages lists all supported programming languages.
var Languages = []Option{
	{Label: "TypeScript — typed JavaScript, most popular for web dev", Value: "typescript"},
	{Label: "JavaScript — runs everywhere, huge ecosystem", Value: "javascript"},
	{Label: "Python — great for AI/ML, scripting, backends", Value: "python"},
	{Label: "Go — fast compiled language, great for APIs and CLIs", Value: "go"},
	{Label: "Rust — memory-safe systems language, very fast", Value: "rust"},
	{Label: "Ruby — elegant syntax, famous for Rails web framework", Value: "ruby"},
	{Label: "PHP — powers WordPress, Laravel, most of the web", Value: "php"},
	{Label: "Java — enterprise standard, Android, Spring Boot", Value: "java"},
	{Label: "Kotlin — modern Java alternative, Android and server", Value: "kotlin"},
	{Label: "C# — Microsoft ecosystem, Unity, .NET APIs", Value: "csharp"},
	{Label: "C++ — game engines, systems, high-performance code", Value: "cpp"},
	{Label: "C — OS kernels, embedded systems, bare metal", Value: "c"},
	{Label: "Swift — Apple platforms, iOS/macOS apps", Value: "swift"},
	{Label: "Dart — Flutter cross-platform mobile and web apps", Value: "dart"},
	{Label: "Elixir — real-time apps, fault-tolerant, Phoenix framework", Value: "elixir"},
	{Label: "Scala — JVM language, big data, functional programming", Value: "scala"},
	{Label: "Zig — low-level control, C alternative, no hidden allocations", Value: "zig"},
	{Label: "Lua — lightweight scripting, game modding, embedded", Value: "lua"},
	{Label: "R — statistics, data analysis, research", Value: "r"},
	{Label: "Perl — text processing, sysadmin, legacy systems", Value: "perl"},
	{Label: "Haskell — purely functional, academic, type theory", Value: "haskell"},
	{Label: "OCaml — functional with imperative, compiler writing", Value: "ocaml"},
	{Label: "Clojure — Lisp on JVM, data-oriented, concurrent", Value: "clojure"},
}

// Frameworks lists all supported frontend and backend frameworks.
var Frameworks = []Option{
	{Label: "Next.js — React framework with SSR, API routes, full-stack", Value: "nextjs"},
	{Label: "React (Vite) — fast React SPA with Vite bundler", Value: "react-vite"},
	{Label: "SvelteKit — lightweight full-stack, compiles away the framework", Value: "sveltekit"},
	{Label: "Nuxt — Vue.js framework with SSR and auto-imports", Value: "nuxt"},
	{Label: "Astro — content-focused sites, ships zero JS by default", Value: "astro"},
	{Label: "Angular — enterprise frontend, batteries-included", Value: "angular"},
	{Label: "Express — minimal Node.js web server, most popular", Value: "express"},
	{Label: "Hono — ultrafast web framework, runs on edge and Node", Value: "hono"},
	{Label: "Fastify — fast Node.js server with schema validation", Value: "fastify"},
	{Label: "FastAPI — modern Python API framework, async, auto-docs", Value: "fastapi"},
	{Label: "Django — Python batteries-included, admin panel, ORM", Value: "django"},
	{Label: "Flask — lightweight Python web framework, flexible", Value: "flask"},
	{Label: "Rails — Ruby full-stack, convention over configuration", Value: "rails"},
	{Label: "Laravel — PHP full-stack, elegant syntax, huge ecosystem", Value: "laravel"},
	{Label: "Gin — fast Go HTTP framework, minimal and performant", Value: "gin"},
	{Label: "Fiber — Express-inspired Go framework, very fast", Value: "fiber"},
	{Label: "None — no framework, start from scratch", Value: "none"},
}

// Styling lists all supported styling/CSS options.
var Styling = []Option{
	{Label: "Tailwind CSS — utility-first CSS, rapid styling", Value: "tailwind"},
	{Label: "Tailwind + shadcn/ui — Tailwind with pre-built accessible components", Value: "tailwind-shadcn"},
	{Label: "CSS Modules — scoped CSS per component, no conflicts", Value: "css-modules"},
	{Label: "Styled Components — CSS-in-JS, dynamic styles in React", Value: "styled-components"},
	{Label: "None — plain CSS or bring your own", Value: "none"},
}

// Databases lists all supported database/ORM options.
var Databases = []Option{
	{Label: "Drizzle ORM — lightweight TypeScript ORM, SQL-like syntax", Value: "drizzle"},
	{Label: "Prisma — popular TypeScript ORM, auto-generated types", Value: "prisma"},
	{Label: "Supabase — hosted Postgres with auth, storage, real-time", Value: "supabase"},
	{Label: "MongoDB — NoSQL document database, flexible schemas", Value: "mongodb-mongoose"},
	{Label: "SQLite — embedded file-based database, zero config", Value: "sqlite"},
	{Label: "None — no database", Value: "none"},
}

// AuthProviders lists all supported authentication providers.
var AuthProviders = []Option{
	{Label: "Clerk — drop-in auth UI, user management, social login", Value: "clerk"},
	{Label: "NextAuth / Auth.js — open-source auth for Next.js, many providers", Value: "nextauth"},
	{Label: "Supabase Auth — built into Supabase, email/social/magic links", Value: "supabase-auth"},
	{Label: "Lucia — lightweight auth library, session-based, flexible", Value: "lucia"},
	{Label: "None — no authentication", Value: "none"},
}

// PaymentProviders lists all supported payment providers.
var PaymentProviders = []Option{
	{Label: "Stripe — payment processing, subscriptions, invoices", Value: "stripe"},
	{Label: "None — no payments", Value: "none"},
}

// EmailProviders lists all supported email providers.
var EmailProviders = []Option{
	{Label: "Resend — modern email API, React email templates", Value: "resend"},
	{Label: "None — no email service", Value: "none"},
}

// PackageManagers lists all supported package managers.
var PackageManagers = []Option{
	{Label: "npm — Node default, largest registry, most compatible", Value: "npm"},
	{Label: "pnpm — fast, disk-efficient, strict dependency resolution", Value: "pnpm"},
	{Label: "bun — all-in-one JS runtime, fastest installs", Value: "bun"},
	{Label: "yarn — Facebook's alternative, workspaces, plug'n'play", Value: "yarn"},
}

// DeploymentTargets lists all supported deployment targets.
var DeploymentTargets = []Option{
	{Label: "Vercel — frontend hosting, serverless functions, auto-deploy", Value: "vercel"},
	{Label: "Railway — backend services, databases, containers", Value: "railway"},
	{Label: "Cloudflare — DNS, CDN, edge workers, Pages", Value: "cloudflare"},
	{Label: "Docker — self-hosted with Dockerfile + compose", Value: "docker"},
	{Label: "AWS — EC2, Lambda, S3, or ECS", Value: "aws"},
	{Label: "Fly.io — global edge deployment, containers", Value: "flyio"},
}

// ExtraOptions lists all optional extras that can be added to a project.
var ExtraOptions = []Option{
	{Label: "GitHub Actions CI — auto-run tests and linting on push", Value: "github-actions"},
	{Label: "Dockerfile + docker-compose — containerize your app for any host", Value: "dockerfile"},
	{Label: "n8n — open-source workflow automation, like Zapier but self-hosted", Value: "n8n"},
	{Label: "Vercel Analytics — web analytics and Core Web Vitals tracking", Value: "vercel-analytics"},
	{Label: "PostHog — open-source product analytics, feature flags, session replay", Value: "posthog"},
	{Label: "ESLint + Prettier — auto-format and catch code issues on save", Value: "eslint-prettier"},
}
