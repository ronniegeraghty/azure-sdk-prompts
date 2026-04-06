package tools

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

const validRegistryYAML = `
version: "1"
tools:
  - name: azure-mcp
    type: mcp
    description: "Azure MCP server"
    version: "0.1.0"
    mcp:
      command: npx
      args: ["-y", "@azure/mcp@latest", "server", "start"]
      env:
        AZURE_SUB: "test-sub"
      tools: ["*"]

  - name: keyvault-skill
    type: skill
    description: "Key Vault Python skill"
    version: "1.0.0"
    skill:
      source: remote
      repo: microsoft/skills
      name: azure-keyvault-py

  - name: local-gen
    type: skill
    description: "Local generator skill"
    skill:
      source: local
      path: "./skills/generator/*"
`

func TestParseRegistryValid(t *testing.T) {
	reg, err := ParseRegistry([]byte(validRegistryYAML))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if reg.Version != "1" {
		t.Errorf("version = %q, want %q", reg.Version, "1")
	}
	if len(reg.Tools) != 3 {
		t.Fatalf("got %d tools, want 3", len(reg.Tools))
	}

	// MCP entry
	mcp := reg.Tools[0]
	if mcp.Name != "azure-mcp" {
		t.Errorf("tool[0].name = %q, want %q", mcp.Name, "azure-mcp")
	}
	if mcp.Type != TypeMCP {
		t.Errorf("tool[0].type = %q, want %q", mcp.Type, TypeMCP)
	}
	if mcp.MCP == nil {
		t.Fatal("tool[0].mcp is nil")
	}
	if mcp.MCP.Command != "npx" {
		t.Errorf("mcp.command = %q, want %q", mcp.MCP.Command, "npx")
	}
	if len(mcp.MCP.Args) != 4 {
		t.Errorf("mcp.args len = %d, want 4", len(mcp.MCP.Args))
	}
	if mcp.MCP.Env["AZURE_SUB"] != "test-sub" {
		t.Errorf("mcp.env[AZURE_SUB] = %q, want %q", mcp.MCP.Env["AZURE_SUB"], "test-sub")
	}
	if len(mcp.MCP.Tools) != 1 || mcp.MCP.Tools[0] != "*" {
		t.Errorf("mcp.tools = %v, want [*]", mcp.MCP.Tools)
	}

	// Remote skill entry
	skill := reg.Tools[1]
	if skill.Type != TypeSkill {
		t.Errorf("tool[1].type = %q, want %q", skill.Type, TypeSkill)
	}
	if skill.Skill == nil {
		t.Fatal("tool[1].skill is nil")
	}
	if skill.Skill.Source != SourceRemote {
		t.Errorf("skill.source = %q, want %q", skill.Skill.Source, SourceRemote)
	}
	if skill.Skill.Repo != "microsoft/skills" {
		t.Errorf("skill.repo = %q, want %q", skill.Skill.Repo, "microsoft/skills")
	}

	// Local skill entry
	local := reg.Tools[2]
	if local.Skill == nil || local.Skill.Source != SourceLocal {
		t.Fatal("tool[2] should be a local skill")
	}
	if local.Skill.Path != "./skills/generator/*" {
		t.Errorf("skill.path = %q, want %q", local.Skill.Path, "./skills/generator/*")
	}
}

func TestParseRegistryInvalidVersion(t *testing.T) {
	data := []byte(`version: "2"
tools: []`)
	_, err := ParseRegistry(data)
	if err == nil {
		t.Fatal("expected error for unsupported version")
	}
}

func TestParseRegistryMissingName(t *testing.T) {
	data := []byte(`
version: "1"
tools:
  - type: mcp
    mcp:
      command: echo`)
	_, err := ParseRegistry(data)
	if err == nil {
		t.Fatal("expected error for missing tool name")
	}
}

func TestParseRegistryDuplicateName(t *testing.T) {
	data := []byte(`
version: "1"
tools:
  - name: dup
    type: mcp
    mcp:
      command: echo
  - name: dup
    type: mcp
    mcp:
      command: echo`)
	_, err := ParseRegistry(data)
	if err == nil {
		t.Fatal("expected error for duplicate tool name")
	}
}

func TestParseRegistryInvalidType(t *testing.T) {
	data := []byte(`
version: "1"
tools:
  - name: bad
    type: unknown`)
	_, err := ParseRegistry(data)
	if err == nil {
		t.Fatal("expected error for invalid tool type")
	}
}

func TestParseRegistryMCPMissingCommand(t *testing.T) {
	data := []byte(`
version: "1"
tools:
  - name: no-cmd
    type: mcp
    mcp:
      args: ["--help"]`)
	_, err := ParseRegistry(data)
	if err == nil {
		t.Fatal("expected error for missing mcp.command")
	}
}

func TestParseRegistryMCPMissingBlock(t *testing.T) {
	data := []byte(`
version: "1"
tools:
  - name: no-block
    type: mcp`)
	_, err := ParseRegistry(data)
	if err == nil {
		t.Fatal("expected error for missing mcp block")
	}
}

func TestParseRegistrySkillMissingBlock(t *testing.T) {
	data := []byte(`
version: "1"
tools:
  - name: no-block
    type: skill`)
	_, err := ParseRegistry(data)
	if err == nil {
		t.Fatal("expected error for missing skill block")
	}
}

func TestParseRegistrySkillInvalidSource(t *testing.T) {
	data := []byte(`
version: "1"
tools:
  - name: bad-source
    type: skill
    skill:
      source: cloud`)
	_, err := ParseRegistry(data)
	if err == nil {
		t.Fatal("expected error for invalid skill.source")
	}
}

func TestParseRegistryRemoteSkillMissingRepo(t *testing.T) {
	data := []byte(`
version: "1"
tools:
  - name: no-repo
    type: skill
    skill:
      source: remote
      name: my-skill`)
	_, err := ParseRegistry(data)
	if err == nil {
		t.Fatal("expected error for missing skill.repo")
	}
}

func TestParseRegistryRemoteSkillMissingName(t *testing.T) {
	data := []byte(`
version: "1"
tools:
  - name: no-skill-name
    type: skill
    skill:
      source: remote
      repo: org/repo`)
	_, err := ParseRegistry(data)
	if err == nil {
		t.Fatal("expected error for missing skill.name")
	}
}

func TestParseRegistryLocalSkillMissingPath(t *testing.T) {
	data := []byte(`
version: "1"
tools:
  - name: no-path
    type: skill
    skill:
      source: local`)
	_, err := ParseRegistry(data)
	if err == nil {
		t.Fatal("expected error for missing skill.path")
	}
}

func TestParseRegistryInvalidYAML(t *testing.T) {
	_, err := ParseRegistry([]byte(`{{{`))
	if err == nil {
		t.Fatal("expected error for invalid YAML")
	}
}

func TestGet(t *testing.T) {
	reg, err := ParseRegistry([]byte(validRegistryYAML))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	entry, ok := reg.Get("azure-mcp")
	if !ok {
		t.Fatal("expected to find azure-mcp")
	}
	if entry.Type != TypeMCP {
		t.Errorf("type = %q, want %q", entry.Type, TypeMCP)
	}

	_, ok = reg.Get("nonexistent")
	if ok {
		t.Fatal("expected Get to return false for nonexistent tool")
	}
}

func TestLoadRegistry(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "registry.yaml")
	if err := os.WriteFile(path, []byte(validRegistryYAML), 0o644); err != nil {
		t.Fatalf("writing test file: %v", err)
	}

	reg, err := LoadRegistry(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(reg.Tools) != 3 {
		t.Errorf("got %d tools, want 3", len(reg.Tools))
	}
}

func TestLoadRegistryFileNotFound(t *testing.T) {
	_, err := LoadRegistry("/nonexistent/registry.yaml")
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}

func TestLoadRemoteRegistry(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-yaml")
		w.Write([]byte(validRegistryYAML))
	}))
	defer srv.Close()

	reg, err := LoadRemoteRegistry(srv.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(reg.Tools) != 3 {
		t.Errorf("got %d tools, want 3", len(reg.Tools))
	}
}

func TestLoadRemoteRegistryBadStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()

	_, err := LoadRemoteRegistry(srv.URL)
	if err == nil {
		t.Fatal("expected error for non-200 status")
	}
}
