package eval

import (
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/config/tool"
)

func TestInjectConfigProps_PopulatesGeneratorAndConfigKeys(t *testing.T) {
	cfg := config.ToolConfig{
		Name: "azure-mcp/claude-opus-4.6",
		Generator: &config.GeneratorConfig{
			Model: "claude-opus-4.6",
			Tools: []tool.Entry{
				{Name: "azure", Type: tool.TypeMCP},
				{Name: "reviewer-skills", Type: tool.TypeSkill},
				{Name: "azure-functions", Type: tool.TypePlugin},
				{Name: "list-files", Type: tool.TypeTool},          // not surfaced in Phase 1
				{Name: "ignored", Type: "bogus"},                   // unknown type, skipped
				{Name: "", Type: tool.TypeMCP},                     // empty name, skipped
				{Name: "default-skill"},                            // empty type defaults to "tool"
			},
		},
	}

	props := map[string]string{}
	injectConfigProps(props, cfg)

	want := map[string]string{
		"generator":              "claude-opus-4.6",
		"config":                 "azure-mcp/claude-opus-4.6",
		"mcp_server:azure":       "true",
		"skill:reviewer-skills":  "true",
		"plugin:azure-functions": "true",
	}
	for k, v := range want {
		if got := props[k]; got != v {
			t.Errorf("props[%q] = %q, want %q", k, got, v)
		}
	}

	// Negative: unknown / tool / empty / colliding-name entries shouldn't leak.
	for _, k := range []string{
		"mcp_server:", "tool:list-files", "skill:ignored",
		"mcp_server:bogus", "skill:", "skill:default-skill",
	} {
		if v, ok := props[k]; ok {
			t.Errorf("props[%q] should be absent, got %q", k, v)
		}
	}
}

func TestInjectConfigProps_ZeroToolsStillEmitsGeneratorAndConfig(t *testing.T) {
	cfg := config.ToolConfig{
		Name: "baseline/claude-opus-4.6",
		Generator: &config.GeneratorConfig{
			Model: "claude-opus-4.6",
		},
	}
	props := map[string]string{}
	injectConfigProps(props, cfg)

	if props["generator"] != "claude-opus-4.6" {
		t.Errorf("generator = %q, want claude-opus-4.6", props["generator"])
	}
	if props["config"] != "baseline/claude-opus-4.6" {
		t.Errorf("config = %q, want baseline/claude-opus-4.6", props["config"])
	}
	if len(props) != 2 {
		t.Errorf("expected only 2 keys, got %d: %v", len(props), props)
	}
}

func TestInjectConfigProps_NilGenerator(t *testing.T) {
	cfg := config.ToolConfig{Name: "weird"}
	props := map[string]string{}
	injectConfigProps(props, cfg)
	if props["config"] != "weird" {
		t.Errorf("config not set with nil generator")
	}
	if _, ok := props["generator"]; ok {
		t.Errorf("generator should not be set when Generator is nil")
	}
}

func TestInjectConfigProps_OverwritesExistingPromptKeys(t *testing.T) {
	// Engine-injected keys must win over prompt frontmatter clashes.
	cfg := config.ToolConfig{
		Name: "real-config",
		Generator: &config.GeneratorConfig{
			Model: "real-model",
			Tools: []tool.Entry{{Name: "azure", Type: tool.TypeMCP}},
		},
	}
	props := map[string]string{
		"generator":        "spoofed-by-prompt",
		"config":           "spoofed-by-prompt",
		"mcp_server:azure": "false-spoof",
	}
	injectConfigProps(props, cfg)

	if props["generator"] != "real-model" {
		t.Errorf("generator overwrite failed: got %q", props["generator"])
	}
	if props["config"] != "real-config" {
		t.Errorf("config overwrite failed: got %q", props["config"])
	}
	if props["mcp_server:azure"] != "true" {
		t.Errorf("mcp_server:azure overwrite failed: got %q", props["mcp_server:azure"])
	}
}
