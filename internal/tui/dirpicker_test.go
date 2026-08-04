package tui

import (
	"strings"
	"testing"
)

func TestICloudNote(t *testing.T) {
	cases := map[string]string{
		"Projects":        "iCloud — main project archive",
		"Projects2026":    "iCloud — 2026 projects",
		"Projects2027":    "iCloud — 2027 projects",
		"ProjectsOld":     "iCloud — main project archive",
		"ProjectsArchive": "iCloud — main project archive",
	}
	for name, want := range cases {
		if got := icloudNote(name); got != want {
			t.Errorf("icloudNote(%q) = %q, want %q", name, got, want)
		}
	}
}

func TestProjectDirDescription(t *testing.T) {
	cases := []struct {
		dir  ProjectDir
		want string
	}{
		{ProjectDir{Note: "iCloud — 2026 projects", Count: 1}, "iCloud — 2026 projects · 1 project"},
		{ProjectDir{Note: "iCloud — main project archive", Count: 30}, "iCloud — main project archive · 30 projects"},
		{ProjectDir{Note: "This Mac — local only, not synced", Count: 0}, "This Mac — local only, not synced · 0 projects"},
		{ProjectDir{Note: "unreadable", Count: -1}, "unreadable"},
	}
	for _, c := range cases {
		if got := c.dir.description(); got != c.want {
			t.Errorf("description() = %q, want %q", got, c.want)
		}
	}
}

// FindProjectDirs must never return a folder that doesn't exist, and the local
// Mac folder must be labelled apart from the iCloud one of the same name.
func TestFindProjectDirs(t *testing.T) {
	dirs := FindProjectDirs()
	if len(dirs) == 0 {
		t.Skip("no Projects folders on this machine")
	}

	labels := map[string]bool{}
	for _, d := range dirs {
		if !isDir(d.Path) {
			t.Errorf("%s is not a directory", d.Path)
		}
		if d.Label == "" || d.Note == "" {
			t.Errorf("%s has an empty label or note", d.Path)
		}
		if labels[d.Label] {
			t.Errorf("duplicate label %q — folders would be indistinguishable", d.Label)
		}
		labels[d.Label] = true
		t.Logf("%-20s %s", d.Label, d.description())
	}

	// The local folder, when present, is tagged so it can't be confused
	// with the iCloud folder of the same name.
	for _, d := range dirs {
		if strings.HasSuffix(d.Path, "/Projects") && !strings.Contains(d.Path, "CloudDocs") {
			if !strings.Contains(d.Label, "Local") {
				t.Errorf("local folder %s labelled %q, want it marked Local", d.Path, d.Label)
			}
		}
	}
}
