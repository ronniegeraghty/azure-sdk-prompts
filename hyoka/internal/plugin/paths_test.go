package plugin

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPluginCacheCandidates(t *testing.T) {
	tests := []struct {
		name     string
		repoDir  string
		plugin   string
		wantNil  bool
		wantSubs []string // expected dir suffixes in order
	}{
		{
			name:    "happy path returns three candidates in precedence order",
			repoDir: "/cache/repos/owner/repo/default",
			plugin:  "my-plugin",
			wantSubs: []string{
				filepath.Join(".github", "plugins", "my-plugin"),
				filepath.Join(".github", "skills", "my-plugin"),
				filepath.Join("skills", "my-plugin"),
			},
		},
		{
			name:    "empty repoDir returns nil",
			repoDir: "",
			plugin:  "my-plugin",
			wantNil: true,
		},
		{
			name:    "empty name returns nil",
			repoDir: "/cache/repos/owner/repo/default",
			plugin:  "",
			wantNil: true,
		},
		{
			name:    "weird repo dir with trailing slash still produces clean joins",
			repoDir: "/some/odd/path/",
			plugin:  "p",
			wantSubs: []string{
				filepath.Join(".github", "plugins", "p"),
				filepath.Join(".github", "skills", "p"),
				filepath.Join("skills", "p"),
			},
		},
		{
			name:    "plugin name with hyphens preserved",
			repoDir: "/r",
			plugin:  "azure-sdk-tools",
			wantSubs: []string{
				filepath.Join(".github", "plugins", "azure-sdk-tools"),
				filepath.Join(".github", "skills", "azure-sdk-tools"),
				filepath.Join("skills", "azure-sdk-tools"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PluginCacheCandidates(tt.repoDir, tt.plugin)
			if tt.wantNil {
				if got != nil {
					t.Fatalf("expected nil, got %v", got)
				}
				return
			}
			if len(got) != len(tt.wantSubs) {
				t.Fatalf("expected %d candidates, got %d: %v", len(tt.wantSubs), len(got), got)
			}
			for i, want := range tt.wantSubs {
				expected := filepath.Join(tt.repoDir, want)
				if got[i] != expected {
					t.Errorf("candidate %d:\n  got  %s\n  want %s", i, got[i], expected)
				}
			}
		})
	}
}

func TestFindPluginInRepo_Precedence(t *testing.T) {
	root := t.TempDir()

	// Seed all three candidate locations with valid plugin dirs (each a
	// single-skill plugin: top-level SKILL.md). The first one (.github/plugins)
	// must win.
	for _, sub := range []string{
		filepath.Join(".github", "plugins", "p"),
		filepath.Join(".github", "skills", "p"),
		filepath.Join("skills", "p"),
	} {
		dir := filepath.Join(root, sub)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", dir, err)
		}
		if err := os.WriteFile(filepath.Join(dir, "SKILL.md"), []byte("ok"), 0o644); err != nil {
			t.Fatalf("write SKILL.md: %v", err)
		}
	}

	got, err := FindPluginInRepo(root, "p")
	if err != nil {
		t.Fatalf("FindPluginInRepo: %v", err)
	}
	want := filepath.Join(root, ".github", "plugins", "p")
	if got != want {
		t.Errorf("expected precedence winner %s, got %s", want, got)
	}
}

func TestFindPluginInRepo_NotFound_EnumeratesAllChecked(t *testing.T) {
	root := t.TempDir()
	_, err := FindPluginInRepo(root, "missing")
	if err == nil {
		t.Fatal("expected error for missing plugin")
	}
	msg := err.Error()
	for _, want := range []string{
		filepath.Join(root, ".github", "plugins", "missing"),
		filepath.Join(root, ".github", "skills", "missing"),
		filepath.Join(root, "skills", "missing"),
	} {
		if !strings.Contains(msg, want) {
			t.Errorf("error message missing checked path %q\nfull:\n%s", want, msg)
		}
	}
}
