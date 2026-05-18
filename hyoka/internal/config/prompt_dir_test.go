package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// writeYAML is a small helper for the tests below.
func writeYAML(t *testing.T, dir, name, body string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(body), 0644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
	return path
}

const minimalCfgBody = `configs:
  - name: example/baseline
    generator:
      model: claude-sonnet-4.5
`

const cfgWithPromptDirBody = `prompt_directory: ./my-prompts
configs:
  - name: example/baseline
    generator:
      model: claude-sonnet-4.5
`

func TestLoad_PromptDirectoryRelativeResolved(t *testing.T) {
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "my-prompts"), 0755); err != nil {
		t.Fatal(err)
	}
	path := writeYAML(t, dir, "config.yaml", cfgWithPromptDirBody)

	cf, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	want := filepath.Join(dir, "my-prompts")
	if filepath.Clean(cf.PromptDirectory) != want {
		t.Errorf("PromptDirectory = %q, want %q", cf.PromptDirectory, want)
	}
}

func TestLoad_PromptDirectoryAbsolutePreserved(t *testing.T) {
	dir := t.TempDir()
	abs := filepath.Join(t.TempDir(), "elsewhere")
	if err := os.MkdirAll(abs, 0755); err != nil {
		t.Fatal(err)
	}
	body := "prompt_directory: " + abs + "\n" + minimalCfgBody
	path := writeYAML(t, dir, "config.yaml", body)

	cf, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cf.PromptDirectory != abs {
		t.Errorf("PromptDirectory = %q, want %q", cf.PromptDirectory, abs)
	}
}

func TestLoad_NoPromptDirectoryDefaultsEmpty(t *testing.T) {
	dir := t.TempDir()
	path := writeYAML(t, dir, "config.yaml", minimalCfgBody)
	cf, err := Load(path)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cf.PromptDirectory != "" {
		t.Errorf("PromptDirectory = %q, want empty", cf.PromptDirectory)
	}
}

func TestLoadDir_PromptDirectoryPropagated(t *testing.T) {
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "p"), 0755); err != nil {
		t.Fatal(err)
	}
	writeYAML(t, dir, "a.yaml", "prompt_directory: ./p\n"+minimalCfgBody)
	writeYAML(t, dir, "b.yaml", `configs:
  - name: example/other
    generator:
      model: claude-sonnet-4.5
`)
	cf, err := LoadDir(dir)
	if err != nil {
		t.Fatalf("LoadDir: %v", err)
	}
	want := filepath.Join(dir, "p")
	if filepath.Clean(cf.PromptDirectory) != want {
		t.Errorf("PromptDirectory = %q, want %q", cf.PromptDirectory, want)
	}
}

func TestLoadDir_PromptDirectoryConflict(t *testing.T) {
	dir := t.TempDir()
	os.MkdirAll(filepath.Join(dir, "p1"), 0755)
	os.MkdirAll(filepath.Join(dir, "p2"), 0755)
	writeYAML(t, dir, "a.yaml", "prompt_directory: ./p1\n"+minimalCfgBody)
	writeYAML(t, dir, "b.yaml", `prompt_directory: ./p2
configs:
  - name: example/other
    generator:
      model: claude-sonnet-4.5
`)
	_, err := LoadDir(dir)
	if err == nil {
		t.Fatal("expected conflict error, got nil")
	}
	if !strings.Contains(err.Error(), "conflicting prompt_directory") {
		t.Errorf("error %q does not mention conflicting prompt_directory", err)
	}
}

func TestResolvePromptDirCandidates_ConfigDirWinsWhenSet(t *testing.T) {
	root := t.TempDir()
	hyokaPrompts := filepath.Join(root, ".hyoka", "prompts")
	customPrompts := filepath.Join(root, "custom")
	os.MkdirAll(hyokaPrompts, 0755)
	os.MkdirAll(customPrompts, 0755)
	proj := &ProjectDir{Root: filepath.Join(root, ".hyoka")}

	got := ResolvePromptDirCandidates(proj, customPrompts)
	if len(got) == 0 || got[0] != customPrompts {
		t.Fatalf("expected customPrompts first, got %v", got)
	}
	found := false
	for _, c := range got {
		if c == hyokaPrompts {
			found = true
			break
		}
	}
	if !found {
		t.Errorf(".hyoka/prompts missing from fallback chain: %v", got)
	}
}

func TestResolvePromptDirCandidates_NoOverrideDelegates(t *testing.T) {
	root := t.TempDir()
	hyokaPrompts := filepath.Join(root, ".hyoka", "prompts")
	os.MkdirAll(hyokaPrompts, 0755)
	proj := &ProjectDir{Root: filepath.Join(root, ".hyoka")}

	got := ResolvePromptDirCandidates(proj, "")
	if len(got) == 0 || got[0] != hyokaPrompts {
		t.Fatalf("expected .hyoka/prompts first when no override, got %v", got)
	}
}

func TestResolvePromptDirCandidates_NonexistentOverrideFallsBack(t *testing.T) {
	root := t.TempDir()
	hyokaPrompts := filepath.Join(root, ".hyoka", "prompts")
	os.MkdirAll(hyokaPrompts, 0755)
	proj := &ProjectDir{Root: filepath.Join(root, ".hyoka")}

	got := ResolvePromptDirCandidates(proj, filepath.Join(root, "does-not-exist"))
	if len(got) == 0 || got[0] != hyokaPrompts {
		t.Fatalf("expected fallback to .hyoka/prompts, got %v", got)
	}
}

func TestPeekPromptDirectory_FindsValueAcrossFiles(t *testing.T) {
	dir := t.TempDir()
	os.MkdirAll(filepath.Join(dir, "p"), 0755)
	writeYAML(t, dir, "a.yaml", "configs: []\n")
	writeYAML(t, dir, "b.yaml", "prompt_directory: ./p\nconfigs: []\n")

	got := PeekPromptDirectory(dir)
	want := filepath.Join(dir, "p")
	if filepath.Clean(got) != want {
		t.Errorf("PeekPromptDirectory = %q, want %q", got, want)
	}
}

func TestPeekPromptDirectory_EmptyWhenAbsent(t *testing.T) {
	dir := t.TempDir()
	writeYAML(t, dir, "a.yaml", "configs: []\n")
	if got := PeekPromptDirectory(dir); got != "" {
		t.Errorf("PeekPromptDirectory = %q, want empty", got)
	}
}

func TestPeekPromptDirectory_NonexistentDir(t *testing.T) {
	if got := PeekPromptDirectory(filepath.Join(t.TempDir(), "missing")); got != "" {
		t.Errorf("PeekPromptDirectory on missing dir = %q, want empty", got)
	}
}
