package services

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// AuthCheck represents the authentication status of a service.
type AuthCheck struct {
	Service string
	Ready   bool
	User    string // authenticated user/account if available
	Message string // error or instruction
}

// CheckAuth verifies authentication for all selected services.
// Returns a list of auth checks with their status.
func CheckAuth(clis CLIStatus, deployment []string) []AuthCheck {
	var checks []AuthCheck

	// GitHub (always check if gh is available)
	if clis.GitHub {
		user, err := ghAuthUser()
		if err != nil {
			checks = append(checks, AuthCheck{
				Service: "GitHub",
				Ready:   false,
				Message: "Not authenticated. Run: gh auth login",
			})
		} else {
			checks = append(checks, AuthCheck{
				Service: "GitHub",
				Ready:   true,
				User:    user,
			})
		}
	} else {
		checks = append(checks, AuthCheck{
			Service: "GitHub",
			Ready:   false,
			Message: "gh CLI not installed. Install: brew install gh",
		})
	}

	// Vercel
	for _, dep := range deployment {
		if dep == "vercel" {
			if clis.Vercel {
				user, err := vercelAuthUser()
				if err != nil {
					checks = append(checks, AuthCheck{
						Service: "Vercel",
						Ready:   false,
						Message: "Not authenticated. Run: vercel login",
					})
				} else {
					checks = append(checks, AuthCheck{
						Service: "Vercel",
						Ready:   true,
						User:    user,
					})
				}
			} else {
				checks = append(checks, AuthCheck{
					Service: "Vercel",
					Ready:   false,
					Message: "vercel CLI not installed. Install: npm i -g vercel",
				})
			}
			break
		}
	}

	return checks
}

// LoginService attempts to interactively log in to a service.
func LoginService(service string) error {
	switch service {
	case "GitHub":
		cmd := exec.Command("gh", "auth", "login")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	case "Vercel":
		cmd := exec.Command("vercel", "login")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	default:
		return fmt.Errorf("unknown service: %s", service)
	}
}

func ghAuthUser() (string, error) {
	out, err := exec.Command("gh", "auth", "status").CombinedOutput()
	if err != nil {
		return "", err
	}
	// Parse "Logged in to github.com account USERNAME"
	s := string(out)
	if idx := strings.Index(s, "account "); idx >= 0 {
		rest := s[idx+8:]
		if end := strings.IndexAny(rest, " \n("); end > 0 {
			return rest[:end], nil
		}
		return strings.TrimSpace(rest), nil
	}
	return "authenticated", nil
}

func vercelAuthUser() (string, error) {
	out, err := exec.Command("vercel", "whoami").CombinedOutput()
	if err != nil {
		return "", err
	}
	user := strings.TrimSpace(string(out))
	if user == "" {
		return "", fmt.Errorf("not logged in")
	}
	return user, nil
}
