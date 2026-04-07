package services

import (
	"os/exec"
	"strings"
)

// CLIStatus tracks which CLI tools are available on the system.
type CLIStatus struct {
	Git, GitHub, Vercel, Docker          bool
	Node, Bun, Pnpm, Yarn, Npm          bool
	Python, Go, Cargo, Ruby, PHP         bool
}

// DetectCLIs checks the system PATH for all supported CLI tools.
func DetectCLIs() CLIStatus {
	return CLIStatus{
		Git:    hasCmd("git"),
		GitHub: hasCmd("gh"),
		Vercel: hasCmd("vercel"),
		Docker: hasCmd("docker"),
		Node:   hasCmd("node"),
		Bun:    hasCmd("bun"),
		Pnpm:   hasCmd("pnpm"),
		Yarn:   hasCmd("yarn"),
		Npm:    hasCmd("npm"),
		Python: hasCmd("python3"),
		Go:     hasCmd("go"),
		Cargo:  hasCmd("cargo"),
		Ruby:   hasCmd("ruby"),
		PHP:    hasCmd("php"),
	}
}

func hasCmd(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

// RunCmd executes a command and returns its combined output.
func RunCmd(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(out)), err
}

// RunCmdInDir executes a command in a specific directory and returns its combined output.
func RunCmdInDir(dir, name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(out)), err
}
