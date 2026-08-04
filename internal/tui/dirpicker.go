package tui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

// iCloudDir is the container that holds the synced Projects folders.
const iCloudDir = "Library/Mobile Documents/com~apple~CloudDocs"

// projectsPrefix is the folder-name prefix shared by every Projects directory.
const projectsPrefix = "Projects"

// ProjectDir is one Projects folder offered by the jump picker.
type ProjectDir struct {
	Path  string
	Label string
	Note  string
	Count int // -1 when the folder could not be read
}

// description is the dim note shown beside the folder name.
func (d ProjectDir) description() string {
	if d.Count < 0 {
		return d.Note
	}
	unit := "projects"
	if d.Count == 1 {
		unit = "project"
	}
	return fmt.Sprintf("%s · %d %s", d.Note, d.Count, unit)
}

// icloudNote explains what an iCloud Projects folder holds. A folder ending in
// a year (Projects2026) is described by that year, anything else is the archive.
func icloudNote(name string) string {
	suffix := name[len(projectsPrefix):]
	if suffix != "" && strings.IndexFunc(suffix, func(r rune) bool { return r < '0' || r > '9' }) == -1 {
		return "iCloud — " + suffix + " projects"
	}
	return "iCloud — main project archive"
}

// countProjects counts the visible sub-directories, or -1 if unreadable.
func countProjects(path string) int {
	entries, err := os.ReadDir(path)
	if err != nil {
		return -1
	}
	n := 0
	for _, e := range entries {
		if e.IsDir() && !strings.HasPrefix(e.Name(), ".") {
			n++
		}
	}
	return n
}

// isDir reports whether path is a directory, following symlinks.
func isDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

// FindProjectDirs lists the Projects folders to jump into: the iCloud ones
// (Projects, Projects2026, ...) followed by the local ~/Projects on this Mac.
// This mirrors the Project Directory picker in Mode Terminal.
func FindProjectDirs() []ProjectDir {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil
	}

	var dirs []ProjectDir
	seen := map[string]bool{}

	icloud := filepath.Join(home, iCloudDir)
	entries, err := os.ReadDir(icloud) // already sorted by name
	if err == nil {
		for _, e := range entries {
			name := e.Name()
			if !strings.HasPrefix(strings.ToLower(name), strings.ToLower(projectsPrefix)) {
				continue
			}
			path := filepath.Join(icloud, name)
			if !isDir(path) {
				continue
			}
			dirs = append(dirs, ProjectDir{
				Path:  path,
				Label: name,
				Note:  icloudNote(name),
				Count: countProjects(path),
			})
			seen[path] = true
		}
	}

	// Local Mac projects — same folder name as iCloud, so label it apart
	local := filepath.Join(home, projectsPrefix)
	if isDir(local) && !seen[local] {
		dirs = append(dirs, ProjectDir{
			Path:  local,
			Label: projectsPrefix + " (Local)",
			Note:  "This Mac — local only, not synced",
			Count: countProjects(local),
		})
	}

	return dirs
}

// nameStep wraps the project-name form so that pressing Up swaps it for a
// picker listing the Projects folders. The form behaves exactly as before
// until Up is pressed.
type nameStep struct {
	form    *huh.Form
	dirs    []ProjectDir
	picking bool
	cursor  int
	jumpDir string
	aborted bool
}

func (m nameStep) Init() tea.Cmd {
	return m.form.Init()
}

func (m nameStep) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	key, isKey := msg.(tea.KeyMsg)

	if m.picking {
		if !isKey {
			return m, nil
		}
		switch key.String() {
		case "up", "k":
			m.cursor = (m.cursor - 1 + len(m.dirs)) % len(m.dirs)
		case "down", "j":
			m.cursor = (m.cursor + 1) % len(m.dirs)
		case "enter":
			m.jumpDir = m.dirs[m.cursor].Path
			return m, tea.Quit
		case "esc", "b", "left":
			m.picking = false
		case "ctrl+c":
			m.aborted = true
			return m, tea.Quit
		}
		return m, nil
	}

	// Up on the name box opens the folder picker instead of reaching the form
	if isKey && key.String() == "up" && len(m.dirs) > 0 {
		m.picking = true
		m.cursor = 0
		return m, nil
	}

	f, cmd := m.form.Update(msg)
	if form, ok := f.(*huh.Form); ok {
		m.form = form
	}
	if m.form.State == huh.StateCompleted || m.form.State == huh.StateAborted {
		return m, tea.Quit
	}
	return m, cmd
}

func (m nameStep) View() string {
	// Leave a clean screen behind once we're done
	if m.jumpDir != "" || m.aborted || m.form.State != huh.StateNormal {
		return ""
	}
	if m.picking {
		return m.pickerView()
	}
	hint := dimStyle().Render("  ↑  jump to a Projects folder instead")
	return m.form.View() + "\n" + hint + "\n"
}

// pickerView renders the folder list in the same shape as Mode Terminal's
// Project Directory menu: name on the left, dim note on the right.
func (m nameStep) pickerView() string {
	labelWidth := 0
	for _, d := range m.dirs {
		if w := lipgloss.Width(d.Label); w > labelWidth {
			labelWidth = w
		}
	}

	selected := lipgloss.NewStyle().Bold(true).Foreground(primary)

	var b strings.Builder
	b.WriteString("\n")
	b.WriteString(titleStyle.Render("  Project Directory"))
	b.WriteString("\n")

	for i, d := range m.dirs {
		marker, label := "    ", d.Label
		if i == m.cursor {
			marker = selected.Render("  > ")
			label = selected.Render(label)
		}
		gap := strings.Repeat(" ", labelWidth-lipgloss.Width(d.Label)+3)
		b.WriteString(marker + label + gap + dimStyle().Render(d.description()) + "\n")
	}

	b.WriteString("\n")
	b.WriteString(dimStyle().Render("  ↑/↓ navigate   Enter: jump into folder   Esc: back to name"))
	b.WriteString("\n")
	return b.String()
}

// runNameStep runs the name form with the folder picker attached. It returns
// the folder to jump into (empty when the user typed a name instead) and
// whether the user backed out entirely.
func runNameStep(form *huh.Form, dirs []ProjectDir) (jumpDir string, aborted bool, err error) {
	final, err := tea.NewProgram(nameStep{form: form, dirs: dirs}).Run()
	if err != nil {
		return "", false, err
	}

	m, ok := final.(nameStep)
	if !ok {
		return "", true, nil
	}
	if m.jumpDir != "" {
		return m.jumpDir, false, nil
	}
	return "", m.aborted || m.form.State == huh.StateAborted, nil
}
