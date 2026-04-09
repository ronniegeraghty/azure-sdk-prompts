package plugin

import (
	"encoding/json"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"testing"
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
