package eval

import (
	"testing"

	"github.com/ronniegeraghty/hyoka/internal/config"
)

func TestCollectRequiredModels_Basic(t *testing.T) {
	configs := []config.ToolConfig{
		{
			Generator: &config.GeneratorConfig{Model: "claude-opus-4.6"},
			Reviewer:  &config.ReviewerConfig{Models: []string{"claude-opus-4.6", "gpt-4.1"}},
		},
	}
	models := collectRequiredModels(configs)
	if len(models) != 2 {
		t.Errorf("expected 2 unique models, got %d: %v", len(models), models)
	}
	// claude-opus-4.6 appears in both generator and reviewer but should only appear once
	found := map[string]bool{}
	for _, m := range models {
		found[m] = true
	}
	if !found["claude-opus-4.6"] || !found["gpt-4.1"] {
		t.Errorf("missing models: %v", models)
	}
}

func TestCollectRequiredModels_MultipleConfigs(t *testing.T) {
	configs := []config.ToolConfig{
		{
			Name:      "baseline",
			Generator: &config.GeneratorConfig{Model: "claude-sonnet-4.5"},
			Reviewer:  &config.ReviewerConfig{Models: []string{"claude-opus-4.6", "gemini-3-pro-preview", "gpt-4.1"}},
		},
		{
			Name:      "azure-mcp",
			Generator: &config.GeneratorConfig{Model: "claude-opus-4.6"},
			Reviewer:  &config.ReviewerConfig{Models: []string{"claude-opus-4.6", "gemini-3-pro-preview", "gpt-4.1"}},
		},
	}
	models := collectRequiredModels(configs)
	// Unique models: claude-sonnet-4.5, claude-opus-4.6, gemini-3-pro-preview, gpt-4.1
	if len(models) != 4 {
		t.Errorf("expected 4 unique models, got %d: %v", len(models), models)
	}
}

func TestCollectRequiredModels_EmptyConfig(t *testing.T) {
	configs := []config.ToolConfig{{Name: "empty"}}
	models := collectRequiredModels(configs)
	if len(models) != 0 {
		t.Errorf("expected 0 models for empty config, got %d", len(models))
	}
}

func TestCollectRequiredModels_NilGenerator(t *testing.T) {
	configs := []config.ToolConfig{
		{
			Reviewer: &config.ReviewerConfig{Models: []string{"gpt-4.1"}},
		},
	}
	models := collectRequiredModels(configs)
	if len(models) != 1 || models[0] != "gpt-4.1" {
		t.Errorf("expected [gpt-4.1], got %v", models)
	}
}

func TestModelCheckResult_OK(t *testing.T) {
	r := &ModelCheckResult{
		Available: []string{"claude-opus-4.6", "gpt-4.1"},
	}
	if !r.OK() {
		t.Error("expected OK with no unavailable models")
	}
	if r.Error() != "" {
		t.Errorf("expected empty error, got %q", r.Error())
	}
}

func TestModelCheckResult_NotOK(t *testing.T) {
	r := &ModelCheckResult{
		Available:   []string{"claude-opus-4.6"},
		Unavailable: []string{"gemini-3-pro-preview"},
	}
	if r.OK() {
		t.Error("expected NOT OK with unavailable models")
	}
	errMsg := r.Error()
	if errMsg == "" {
		t.Error("expected non-empty error message")
	}
	if !stringContains(errMsg, "gemini-3-pro-preview") {
		t.Errorf("error should mention unavailable model: %q", errMsg)
	}
}

func TestModelCheckResult_MultipleUnavailable(t *testing.T) {
	r := &ModelCheckResult{
		Available:   []string{"claude-opus-4.6"},
		Unavailable: []string{"gemini-3-pro-preview", "nonexistent-model"},
	}
	errMsg := r.Error()
	if !stringContains(errMsg, "gemini-3-pro-preview") || !stringContains(errMsg, "nonexistent-model") {
		t.Errorf("error should mention all unavailable models: %q", errMsg)
	}
}

func stringContains(s, substr string) bool {
	return len(s) >= len(substr) && stringContainsHelper(s, substr)
}

func stringContainsHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
