package config

import (
	"testing"
)

func TestApplyVersionOverrides_PinsByName(t *testing.T) {
	cf := &ConfigFile{
		ToolVersionOverride: map[string]string{
			"my-mcp":   "v2.0.0",
			"my-skill": "v1.5.0",
		},
		Configs: []ToolConfig{{
			Name: "c1",
			Generator: &GeneratorConfig{
				Tools: []ToolEntry{
					{Name: "my-mcp", Type: "mcp", Command: "npx"},
					{Name: "untouched", Type: "skill", Source: "remote", Repo: "x/y"},
				},
			},
			Reviewer: &ReviewerConfig{
				Tools: []ToolEntry{
					{Name: "my-skill", Type: "skill", Source: "remote", Repo: "x/my-skill"},
				},
			},
		}},
	}

	cf.ApplyVersionOverrides()

	if got := cf.Configs[0].Generator.Tools[0].Version; got != "v2.0.0" {
		t.Errorf("generator[0] version: got %q, want v2.0.0", got)
	}
	if got := cf.Configs[0].Generator.Tools[1].Version; got != "" {
		t.Errorf("generator[1] should be untouched, got version %q", got)
	}
	if got := cf.Configs[0].Reviewer.Tools[0].Version; got != "v1.5.0" {
		t.Errorf("reviewer[0] version: got %q, want v1.5.0", got)
	}
}

func TestApplyVersionOverrides_PerEntryWinsOverMap(t *testing.T) {
	cf := &ConfigFile{
		ToolVersionOverride: map[string]string{"my-skill": "v1.0.0"},
		Configs: []ToolConfig{{
			Name: "c1",
			Generator: &GeneratorConfig{
				Tools: []ToolEntry{
					{Name: "my-skill", Type: "skill", Source: "remote", Repo: "x/y", Version: "v9.9.9"},
				},
			},
		}},
	}
	cf.ApplyVersionOverrides()
	if got := cf.Configs[0].Generator.Tools[0].Version; got != "v9.9.9" {
		t.Errorf("per-entry version should win: got %q, want v9.9.9", got)
	}
}

func TestApplyVersionOverrides_NoMapNoOp(t *testing.T) {
	cf := &ConfigFile{
		Configs: []ToolConfig{{
			Name: "c1",
			Generator: &GeneratorConfig{
				Tools: []ToolEntry{{Name: "x", Type: "skill", Source: "remote", Repo: "a/b"}},
			},
		}},
	}
	cf.ApplyVersionOverrides()
	if got := cf.Configs[0].Generator.Tools[0].Version; got != "" {
		t.Errorf("no override map → version should stay empty, got %q", got)
	}
}

func TestApplyVersionOverrides_Idempotent(t *testing.T) {
	cf := &ConfigFile{
		ToolVersionOverride: map[string]string{"x": "v1"},
		Configs: []ToolConfig{{
			Name: "c1",
			Generator: &GeneratorConfig{
				Tools: []ToolEntry{{Name: "x", Type: "skill", Source: "remote", Repo: "a/b"}},
			},
		}},
	}
	cf.ApplyVersionOverrides()
	cf.ApplyVersionOverrides()
	if got := cf.Configs[0].Generator.Tools[0].Version; got != "v1" {
		t.Errorf("idempotent apply should keep v1, got %q", got)
	}
}

func TestParseConfig_ToolVersionOverride(t *testing.T) {
	yaml := []byte(`
tool_version_override:
  my-skill: "v2.0.0"
configs:
  - name: c1
    generator:
      model: gpt-5.2
      tools:
        - name: my-skill
          type: skill
          source: remote
          repo: acme/my-skill
`)
	cf, err := Parse(yaml)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if cf.ToolVersionOverride["my-skill"] != "v2.0.0" {
		t.Errorf("override map not parsed: %v", cf.ToolVersionOverride)
	}
	// Parse does not call ApplyVersionOverrides (Load does); ensure the API works
	// when invoked explicitly.
	cf.ApplyVersionOverrides()
	if got := cf.Configs[0].Generator.Tools[0].Version; got != "v2.0.0" {
		t.Errorf("post-apply version: got %q, want v2.0.0", got)
	}
}
