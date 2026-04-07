package config

// ProjectConfig holds the full configuration for a project to be scaffolded.
type ProjectConfig struct {
	Name           string
	Stack          string // preset name or "custom"
	Language       string
	Framework      string
	Styling        string
	Backend        string
	Database       string
	Auth           string
	Payments       string
	Email          string
	PackageManager string
	Deployment     string
	Extras         []string
	ProjectDir     string
}
