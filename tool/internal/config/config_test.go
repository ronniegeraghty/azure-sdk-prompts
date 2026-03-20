package config

import (
"os"
"path/filepath"
"testing"
)

func TestParseValidConfig(t *testing.T) {
data := []byte(`
configs:
  - name: test-config
    description: "Test configuration"
    model: "gpt-4"
    mcp_servers: {}
    skill_directories: []
    available_tools: []
    excluded_tools: []
  - name: test-config-2
    description: "Second test"
    model: "claude-sonnet-4.5"
    mcp_servers:
      azure:
        type: local
        command: npx
        args: ["-y", "@azure/mcp@latest"]
        tools: ["*"]
    skill_directories: []
    available_tools: []
    excluded_tools: []
`)
cfg, err := Parse(data)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if len(cfg.Configs) != 2 {
t.Fatalf("expected 2 configs, got %d", len(cfg.Configs))
}
if cfg.Configs[0].Name != "test-config" {
t.Errorf("expected name 'test-config', got %q", cfg.Configs[0].Name)
}
if cfg.Configs[0].Model != "gpt-4" {
t.Errorf("expected model 'gpt-4', got %q", cfg.Configs[0].Model)
}
// Check MCP server on second config
if cfg.Configs[1].MCPServers == nil {
t.Fatal("expected MCP servers on second config")
}
azure, ok := cfg.Configs[1].MCPServers["azure"]
if !ok {
t.Fatal("expected 'azure' MCP server")
}
if azure.Command != "npx" {
t.Errorf("expected command 'npx', got %q", azure.Command)
}
}

func TestParseEmptyConfig(t *testing.T) {
data := []byte(`configs: []`)
_, err := Parse(data)
if err == nil {
t.Fatal("expected error for empty configs")
}
}

func TestParseConfigMissingName(t *testing.T) {
data := []byte(`
configs:
  - description: "No name"
    model: "gpt-4"
`)
_, err := Parse(data)
if err == nil {
t.Fatal("expected error for config missing name")
}
}

func TestParseInvalidYAML(t *testing.T) {
data := []byte(`not: valid: yaml: [`)
_, err := Parse(data)
if err == nil {
t.Fatal("expected error for invalid YAML")
}
}

func TestGetConfig(t *testing.T) {
data := []byte(`
configs:
  - name: alpha
    description: "Alpha"
    model: "gpt-4"
  - name: beta
    description: "Beta"
    model: "claude-sonnet-4.5"
`)
cfg, err := Parse(data)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}

tc, err := cfg.GetConfig("beta")
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if tc.Name != "beta" {
t.Errorf("expected 'beta', got %q", tc.Name)
}

_, err = cfg.GetConfig("nonexistent")
if err == nil {
t.Fatal("expected error for nonexistent config")
}
}

func TestGetConfigs(t *testing.T) {
data := []byte(`
configs:
  - name: alpha
    description: "Alpha"
    model: "gpt-4"
  - name: beta
    description: "Beta"
    model: "claude-sonnet-4.5"
  - name: gamma
    description: "Gamma"
    model: "gpt-4"
`)
cfg, err := Parse(data)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}

// Empty names returns all
all, err := cfg.GetConfigs(nil)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if len(all) != 3 {
t.Errorf("expected 3 configs, got %d", len(all))
}

// Specific names
subset, err := cfg.GetConfigs([]string{"alpha", "gamma"})
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if len(subset) != 2 {
t.Errorf("expected 2 configs, got %d", len(subset))
}

// Missing name
_, err = cfg.GetConfigs([]string{"alpha", "missing"})
if err == nil {
t.Fatal("expected error for missing config name")
}
}

func TestLoadFromFile(t *testing.T) {
dir := t.TempDir()
path := filepath.Join(dir, "config.yaml")
content := []byte(`
configs:
  - name: file-test
    description: "From file"
    model: "gpt-4"
`)
if err := os.WriteFile(path, content, 0644); err != nil {
t.Fatalf("failed to write test file: %v", err)
}

cfg, err := Load(path)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if cfg.Configs[0].Name != "file-test" {
t.Errorf("expected 'file-test', got %q", cfg.Configs[0].Name)
}

// Non-existent file
_, err = Load(filepath.Join(dir, "nonexistent.yaml"))
if err == nil {
t.Fatal("expected error for nonexistent file")
}
}
