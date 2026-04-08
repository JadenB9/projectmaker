package remove

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

var (
	errColor  = lipgloss.Color("#C27171")
	warnColor = lipgloss.Color("#D4A574")
	dimColor  = lipgloss.Color("#888888")
	doneColor = lipgloss.Color("#8FAE8B")
	primary   = lipgloss.Color("#7C9CBF")
)

// Run launches the project removal flow.
func Run() {
	fmt.Println()
	fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(errColor).Render("  Remove Project"))
	fmt.Println(lipgloss.NewStyle().Foreground(dimColor).Render("  This will permanently delete a project locally and from GitHub.\n"))

	// Get project name
	var projectName string
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Project name to remove").
				Placeholder("my-project").
				Value(&projectName),
		),
	).Run()
	if err != nil || projectName == "" {
		fmt.Println(lipgloss.NewStyle().Foreground(dimColor).Render("  Cancelled."))
		return
	}

	projectName = strings.TrimSpace(projectName)
	cwd, _ := os.Getwd()
	localDir := filepath.Join(cwd, projectName)

	// Check what exists
	localExists := dirExists(localDir)
	ghExists := ghRepoExists(projectName)

	if !localExists && !ghExists {
		fmt.Println()
		fmt.Println(lipgloss.NewStyle().Foreground(warnColor).Render(
			fmt.Sprintf("  No project %q found locally or on GitHub.", projectName)))
		fmt.Println()
		return
	}

	// Show what will be deleted
	fmt.Println()
	fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(warnColor).Render("  The following will be permanently deleted:"))
	fmt.Println()
	if localExists {
		fmt.Printf("  %s  Local directory: %s\n",
			lipgloss.NewStyle().Foreground(errColor).Render("[delete]"),
			localDir)
	}
	if ghExists {
		ghUser := getGHUser()
		repoName := projectName
		if ghUser != "" {
			repoName = ghUser + "/" + projectName
		}
		fmt.Printf("  %s  GitHub repo: %s\n",
			lipgloss.NewStyle().Foreground(errColor).Render("[delete]"),
			repoName)
	}
	fmt.Println()

	// Require typing the name to confirm
	var confirmation string
	err = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title(fmt.Sprintf("Type %q to confirm deletion", projectName)).
				Placeholder(projectName).
				Validate(func(s string) error {
					if s != projectName {
						return fmt.Errorf("name doesn't match")
					}
					return nil
				}).
				Value(&confirmation),
		),
	).Run()
	if err != nil {
		fmt.Println(lipgloss.NewStyle().Foreground(dimColor).Render("  Cancelled."))
		return
	}

	fmt.Println()

	// Delete GitHub repo first
	if ghExists {
		fmt.Print(lipgloss.NewStyle().Foreground(dimColor).Render("  Deleting GitHub repo..."))
		err := deleteGHRepo(projectName)
		if err != nil {
			fmt.Println()
			fmt.Printf("  %s GitHub repo deletion failed: %v\n",
				lipgloss.NewStyle().Foreground(errColor).Render("[error]"), err)
			fmt.Println(lipgloss.NewStyle().Foreground(dimColor).Render(
				"  Try manually: gh repo delete " + projectName + " --yes"))
		} else {
			fmt.Printf("\r  %s GitHub repo deleted\n",
				lipgloss.NewStyle().Foreground(doneColor).Render("[done] "))
		}
	}

	// Delete local directory
	if localExists {
		fmt.Print(lipgloss.NewStyle().Foreground(dimColor).Render("  Deleting local directory..."))
		err := os.RemoveAll(localDir)
		if err != nil {
			fmt.Println()
			fmt.Printf("  %s Local deletion failed: %v\n",
				lipgloss.NewStyle().Foreground(errColor).Render("[error]"), err)
		} else {
			fmt.Printf("\r  %s Local directory deleted\n",
				lipgloss.NewStyle().Foreground(doneColor).Render("[done] "))
		}
	}

	fmt.Println()
	fmt.Println(lipgloss.NewStyle().Foreground(doneColor).Render(
		fmt.Sprintf("  Project %q has been removed.", projectName)))
	fmt.Println()
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

func ghRepoExists(name string) bool {
	if _, err := exec.LookPath("gh"); err != nil {
		return false
	}
	out, err := exec.Command("gh", "repo", "view", name, "--json", "name").CombinedOutput()
	return err == nil && len(out) > 0
}

func deleteGHRepo(name string) error {
	cmd := exec.Command("gh", "repo", "delete", name, "--yes")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func getGHUser() string {
	out, err := exec.Command("gh", "api", "user", "--jq", ".login").CombinedOutput()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}
