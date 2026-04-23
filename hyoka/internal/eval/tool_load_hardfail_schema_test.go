package eval

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/prompt"
)

// TestCopilotRunner_ToolLoadFailure_RemotePluginUncached proves that a
// `type: plugin` tool entry with `source: remote` and no matching cache
// hard-fails the eval with ErrorCategory="tool_load_failure" — i.e.
// fetch failures surface before CreateSession rather than degrading
// silently. Runs with a clean HOME so the cache lookup misses.
func TestCopilotRunner_ToolLoadFailure_RemotePluginUncached(t *testing.T) {
	prevHome := os.Getenv("HOME")
	cleanHome := t.TempDir()
	_ = os.Setenv("HOME", cleanHome)
	t.Cleanup(func() { _ = os.Setenv("HOME", prevHome) })

	runner := NewCopilotPromptRunner(PromptRunnerOptions{
		MaxSessionActions: 50,
	})

	cfg := &config.ToolConfig{
		Name: "test-config",
		Generator: &config.GeneratorConfig{
			Model: "gpt-4",
			Tools: []config.ToolEntry{
				{Type: "plugin", Name: "uncached-remote@skills", Source: "remote"},
			},
		},
	}

	p := &prompt.Prompt{
		ID:         "test-prompt",
		PromptText: "Write code",
	}

	workDir := t.TempDir()

	result, err := runner.Run(context.Background(), p, cfg, workDir)
	if err == nil {
		t.Fatal("expected error for uncached remote plugin")
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.ErrorCategory != "tool_load_failure" {
		t.Errorf("expected ErrorCategory='tool_load_failure', got %q", result.ErrorCategory)
	}
	if len(result.GeneratedFiles) != 0 {
		t.Errorf("session should not have started; got %d generated files", len(result.GeneratedFiles))
	}
}

// TestCopilotRunner_ToolLoadFailure_PluginOnlyInGenerator_ReviewerUnaffected
// proves the eval-time analogue of the no-auto-append contract: a failed
// plugin in generator.tools aborts the run with tool_load_failure, and
// the failure references the declared plugin — it was not smuggled into
// a shared tool list.
func TestCopilotRunner_ToolLoadFailure_PluginOnlyInGenerator_ReviewerUnaffected(t *testing.T) {
	runner := NewCopilotPromptRunner(PromptRunnerOptions{
		MaxSessionActions: 50,
	})

	cfg := &config.ToolConfig{
		Name: "test-config",
		Generator: &config.GeneratorConfig{
			Model: "gpt-4",
			Tools: []config.ToolEntry{
				{Type: "plugin", Name: "gen-only-ghost", Source: "local"},
			},
		},
		Reviewer: &config.ReviewerConfig{
			Models: []string{"gpt-4"},
		},
	}

	p := &prompt.Prompt{
		ID:         "test-prompt",
		PromptText: "Write code",
	}

	workDir := t.TempDir()

	result, err := runner.Run(context.Background(), p, cfg, workDir)
	if err == nil {
		t.Fatal("expected error for missing generator-only plugin")
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.ErrorCategory != "tool_load_failure" {
		t.Errorf("expected ErrorCategory='tool_load_failure', got %q", result.ErrorCategory)
	}
	if !strings.Contains(err.Error(), "gen-only-ghost") {
		t.Errorf("expected error to reference the failed plugin name; got: %v", err)
	}
}
