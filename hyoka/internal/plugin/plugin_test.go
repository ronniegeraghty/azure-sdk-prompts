package plugin

import (
	"encoding/json"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/toolload"
)

func TestMain(m *testing.M) {
	slog.SetDefault(slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{Level: slog.LevelError + 1})))
	os.Exit(m.Run())
}

func TestLoadPlugin(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "azure-sdk-tools.yaml"), []byte(`
name: azure-sdk-tools
description: Azure SDK development tools
skills:
  - type: local
    path: ../skills/generator
  - type: remote
    repo: github.com/Azure/ai-hub-sdk
    name: azure-sdk-tools
mcp_servers:
  azure:
    type: sse
    command: npx
    args: ["-y", "@azure/mcp@latest"]
hooks:
  pre_tool_use:
    - validate_workspace_paths
  post_tool_use:
    - validate_file_sizes
`), 0644)

	reg := NewRegistry()
	if err := reg.LoadDir(dir); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if reg.Count() != 1 {
		t.Fatalf("expected 1 plugin, got %d", reg.Count())
	}

	p, err := reg.Get("azure-sdk-tools")
	if err != nil {
		t.Fatalf("plugin not found: %v", err)
	}
	if p.Description != "Azure SDK development tools" {
		t.Errorf("wrong description: %q", p.Description)
	}
	if len(p.Skills) != 2 {
		t.Errorf("expected 2 skills, got %d", len(p.Skills))
	}
	if len(p.MCPServers) != 1 {
		t.Errorf("expected 1 MCP server, got %d", len(p.MCPServers))
	}
	if p.Hooks == nil {
		t.Fatal("expected hooks")
	}
	if len(p.Hooks.PreToolUse) != 1 {
		t.Errorf("expected 1 pre_tool_use hook, got %d", len(p.Hooks.PreToolUse))
	}
	if len(p.Hooks.PostToolUse) != 1 {
		t.Errorf("expected 1 post_tool_use hook, got %d", len(p.Hooks.PostToolUse))
	}
}

func TestLoadDirMultiplePlugins(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "plugin-a.yaml"), []byte(`
name: plugin-a
skills:
  - type: local
    path: ./a
`), 0644)
	os.WriteFile(filepath.Join(dir, "plugin-b.yaml"), []byte(`
name: plugin-b
skills:
  - type: local
    path: ./b
`), 0644)

	reg := NewRegistry()
	if err := reg.LoadDir(dir); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if reg.Count() != 2 {
		t.Errorf("expected 2 plugins, got %d", reg.Count())
	}

	names := reg.List()
	sort.Strings(names)
	if names[0] != "plugin-a" || names[1] != "plugin-b" {
		t.Errorf("unexpected names: %v", names)
	}
}

func TestLoadDirDuplicateNames(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "a.yaml"), []byte("name: same-name\nskills:\n  - type: local\n    path: ./a\n"), 0644)
	os.WriteFile(filepath.Join(dir, "b.yaml"), []byte("name: same-name\nskills:\n  - type: local\n    path: ./b\n"), 0644)

	reg := NewRegistry()
	err := reg.LoadDir(dir)
	if err == nil {
		t.Error("expected error for duplicate plugin names")
	}
}

func TestLoadDirSkipsInvalid(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "bad.yaml"), []byte("not: valid: yaml: ["), 0644)
	os.WriteFile(filepath.Join(dir, "noname.yaml"), []byte("skills:\n  - type: local\n    path: ./x\n"), 0644)
	os.WriteFile(filepath.Join(dir, "readme.txt"), []byte("not yaml"), 0644)

	reg := NewRegistry()
	if err := reg.LoadDir(dir); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if reg.Count() != 0 {
		t.Errorf("expected 0 valid plugins, got %d", reg.Count())
	}
}

func TestLoadDirNonexistent(t *testing.T) {
	reg := NewRegistry()
	// Non-existent directory should not error (just skip)
	if err := reg.LoadDir("/nonexistent/path"); err != nil {
		t.Fatalf("expected no error for missing dir, got %v", err)
	}
}

func TestGetNotFound(t *testing.T) {
	reg := NewRegistry()
	_, err := reg.Get("nonexistent")
	if err == nil {
		t.Error("expected error for missing plugin")
	}
}

func TestPluginJSON(t *testing.T) {
	p := Plugin{
		Name:        "test",
		Description: "Test plugin",
		Skills: []PluginSkill{
			{Type: "local", Path: "./skills"},
		},
	}
	data, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	var decoded Plugin
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if decoded.Name != "test" {
		t.Errorf("expected name 'test', got %q", decoded.Name)
	}
}

func TestAll(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "p.yaml"), []byte("name: p\nskills:\n  - type: local\n    path: ./x\n"), 0644)

	reg := NewRegistry()
	reg.LoadDir(dir)

	all := reg.All()
	if len(all) != 1 {
		t.Errorf("expected 1 plugin, got %d", len(all))
	}
}

func TestPluginToToolEntries(t *testing.T) {
	p := &Plugin{
		Name: "sdk-tools",
		Skills: []PluginSkill{
			{Type: "local", Path: "./skills/sdk"},
			{Type: "remote", Repo: "github.com/example/repo", Name: "sdk"},
		},
		MCPServers: map[string]*MCPServer{
			"azure-cli": {Type: "stdio", Command: "az", Args: []string{"mcp"}, Tools: []string{"*"}},
		},
	}

	entries := p.ToToolEntries()
	if len(entries) != 3 {
		t.Fatalf("expected 3 tool entries, got %d", len(entries))
	}
	if entries[0].Type != "skill" || entries[0].Name != "./skills/sdk" {
		t.Errorf("expected first entry to use path as name, got %+v", entries[0])
	}
	if entries[1].Type != "skill" || entries[1].Repo != "github.com/example/repo" {
		t.Errorf("expected second entry to be remote skill, got %+v", entries[1])
	}
	if entries[2].Type != "mcp" || entries[2].Name != "azure-cli" || entries[2].Command != "az" {
		t.Errorf("expected MCP entry for azure-cli, got %+v", entries[2])
	}
}

// TestResolveInstalled_ContainerPluginFanOut verifies the regression for
// the microsoft/skills `.github/plugins/<name>/skills/<child>/SKILL.md`
// layout. Before the fix, ResolveInstalled required a SKILL.md at the
// plugin directory root and returned "" for container plugins (which
// caused tool_load_failure even after the cache was populated). The
// fix accepts container directories and a sibling EnumerateChildSkills
// returns one path per child — the validator emits one report row per
// child so the SDK loads them and the verifier checks the right names.
func TestResolveInstalled_ContainerPluginFanOut(t *testing.T) {
home := t.TempDir()
restore := toolload.SetTestRoot(filepath.Join(home, ".hyoka", "cache"))
defer restore()

// Build a fake cache mirroring microsoft/skills' layout under the new
// canonical path:
//   <CacheRoot>/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/
//     ├── README.md            (no SKILL.md at the root)
//     └── skills/
//         ├── azure-keyvault-py/SKILL.md
//         └── azure-identity-py/SKILL.md
pluginDir := filepath.Join(toolload.RepoCacheDir("microsoft", "skills", "default"),
".github", "plugins", "azure-sdk-python")
if err := os.MkdirAll(pluginDir, 0o755); err != nil {
t.Fatal(err)
}
// README.md at the root, deliberately NOT a SKILL.md — this is what
// fooled the old isSkillDir check into returning false.
if err := os.WriteFile(filepath.Join(pluginDir, "README.md"), []byte("plugin readme"), 0o644); err != nil {
t.Fatal(err)
}
for _, child := range []string{"azure-keyvault-py", "azure-identity-py"} {
childDir := filepath.Join(pluginDir, "skills", child)
if err := os.MkdirAll(childDir, 0o755); err != nil {
t.Fatal(err)
}
if err := os.WriteFile(filepath.Join(childDir, "SKILL.md"), []byte("# "+child), 0o644); err != nil {
t.Fatal(err)
}
}

got := ResolveInstalled("microsoft/skills", "azure-sdk-python")
if got == "" {
t.Fatal("ResolveInstalled returned empty for a container plugin (regression: pre-fix behavior)")
}
if got != pluginDir {
t.Errorf("ResolveInstalled returned %q, want %q", got, pluginDir)
}

children := EnumerateChildSkills(got)
if len(children) != 2 {
t.Fatalf("EnumerateChildSkills returned %d children, want 2: %v", len(children), children)
}
// Sorted lexicographically: azure-identity-py < azure-keyvault-py
wantBases := []string{"azure-identity-py", "azure-keyvault-py"}
for i, child := range children {
if filepath.Base(child) != wantBases[i] {
t.Errorf("children[%d] base = %q, want %q", i, filepath.Base(child), wantBases[i])
}
if _, err := os.Stat(filepath.Join(child, "SKILL.md")); err != nil {
t.Errorf("child %q missing SKILL.md: %v", child, err)
}
}
}

// TestResolveInstalled_SingleSkillPluginStillWorks verifies the fix did
// not regress the single-skill plugin layout (top-level SKILL.md, no
// `skills/` subdirectory). EnumerateChildSkills must return nil so the
// validator falls back to recording one skill row named after the plugin.
func TestResolveInstalled_SingleSkillPluginStillWorks(t *testing.T) {
home := t.TempDir()
restore := toolload.SetTestRoot(filepath.Join(home, ".hyoka", "cache"))
defer restore()

pluginDir := filepath.Join(toolload.RepoCacheDir("acme", "skills", "default"),
".github", "plugins", "solo")
if err := os.MkdirAll(pluginDir, 0o755); err != nil {
t.Fatal(err)
}
if err := os.WriteFile(filepath.Join(pluginDir, "SKILL.md"), []byte("# solo"), 0o644); err != nil {
t.Fatal(err)
}

got := ResolveInstalled("acme/skills", "solo")
if got != pluginDir {
t.Errorf("ResolveInstalled returned %q, want %q", got, pluginDir)
}
if children := EnumerateChildSkills(got); children != nil {
t.Errorf("EnumerateChildSkills returned %v for a single-skill plugin, want nil", children)
}
}

// TestEnumerateChildSkills_IgnoresChildrenWithoutSkillMd ensures the
// fan-out helper does not report directories that are missing SKILL.md
// (e.g. a `__pycache__` dir or a child that's been partially deleted).
func TestEnumerateChildSkills_IgnoresChildrenWithoutSkillMd(t *testing.T) {
dir := t.TempDir()
mustMkdir := func(p string) {
if err := os.MkdirAll(p, 0o755); err != nil {
t.Fatal(err)
}
}
mustMkdir(filepath.Join(dir, "skills", "good"))
if err := os.WriteFile(filepath.Join(dir, "skills", "good", "SKILL.md"), []byte("# good"), 0o644); err != nil {
t.Fatal(err)
}
// Empty subdir — no SKILL.md.
mustMkdir(filepath.Join(dir, "skills", "empty"))
// File at the top of skills/ — should be skipped (not a dir).
if err := os.WriteFile(filepath.Join(dir, "skills", "loose.md"), []byte("loose"), 0o644); err != nil {
t.Fatal(err)
}

got := EnumerateChildSkills(dir)
if len(got) != 1 {
t.Fatalf("got %d children, want 1: %v", len(got), got)
}
if filepath.Base(got[0]) != "good" {
t.Errorf("got child %q, want basename %q", got[0], "good")
}
}
