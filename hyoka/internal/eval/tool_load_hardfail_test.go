package eval

import (
	"context"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/prompt"
)

func TestCopilotRunner_ToolLoadFailure_HardFail(t *testing.T) {
	// This test verifies that when ValidateAndExpand fails (e.g., missing plugin),
	// the eval aborts with ErrorCategory="tool_load_failure" BEFORE CreateSession.
	
	runner := NewCopilotPromptRunner(PromptRunnerOptions{
		MaxSessionActions: 50,
	})
	
	// Config with a non-existent plugin
	cfg := &config.ToolConfig{
		Name: "test-config",
		Generator: &config.GeneratorConfig{
			Model: "gpt-4",
		},
		Plugins: []string{"nonexistent-plugin"},
	}
	
	p := &prompt.Prompt{
		ID:         "test-prompt",
		PromptText: "Write a hello world program",
	}
	
	// Create temp work dir
	workDir := t.TempDir()
	
	result, err := runner.Run(context.Background(), p, cfg, workDir)
	
	// Expect error
	if err == nil {
		t.Fatal("expected error for missing plugin")
	}
	
	// Expect EvalResult with tool_load_failure category
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	
	if result.ErrorCategory != "tool_load_failure" {
		t.Errorf("expected ErrorCategory='tool_load_failure', got %q", result.ErrorCategory)
	}
	
	if result.Error == "" {
		t.Error("expected non-empty Error field")
	}
	
	if result.ErrorDetails == "" {
		t.Error("expected non-empty ErrorDetails field")
	}
	
	if result.Success {
		t.Error("expected Success=false")
	}
	
	// Key contract: CreateSession should never have been called
	// We verify this indirectly: if the session had been created,
	// we'd have session events or generated files. With tool_load_failure,
	// there should be neither.
	if len(result.GeneratedFiles) > 0 {
		t.Errorf("expected no generated files on tool_load_failure, got %d", len(result.GeneratedFiles))
	}
}

func TestCopilotRunner_ToolLoadFailure_MissingSkill(t *testing.T) {
	runner := NewCopilotPromptRunner(PromptRunnerOptions{
		MaxSessionActions: 50,
	})
	
	cfg := &config.ToolConfig{
		Name: "test-config",
		Generator: &config.GeneratorConfig{
			Model: "gpt-4",
			Tools: []config.ToolEntry{
				{Type: "skill", Source: "local", Path: "./nonexistent-skill", Name: "missing"},
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
		t.Fatal("expected error for missing skill")
	}
	
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	
	if result.ErrorCategory != "tool_load_failure" {
		t.Errorf("expected ErrorCategory='tool_load_failure', got %q", result.ErrorCategory)
	}
}

func TestCopilotRunner_ToolLoadFailure_EmptySkillDir(t *testing.T) {
	runner := NewCopilotPromptRunner(PromptRunnerOptions{
		MaxSessionActions: 50,
	})
	
	// Create empty skills dir
	emptyDir := t.TempDir()
	
	cfg := &config.ToolConfig{
		Name: "test-config",
		Generator: &config.GeneratorConfig{
			Model: "gpt-4",
			Tools: []config.ToolEntry{
				{Type: "skill", Source: "local", Path: emptyDir, Name: "empty", SkillDir: true},
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
		t.Fatal("expected error for empty skill_dir")
	}
	
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	
	if result.ErrorCategory != "tool_load_failure" {
		t.Errorf("expected ErrorCategory='tool_load_failure', got %q", result.ErrorCategory)
	}
}

func TestCopilotRunner_ToolLoadFailure_MCPMissingCommand(t *testing.T) {
	runner := NewCopilotPromptRunner(PromptRunnerOptions{
		MaxSessionActions: 50,
	})
	
	cfg := &config.ToolConfig{
		Name: "test-config",
		Generator: &config.GeneratorConfig{
			Model: "gpt-4",
			Tools: []config.ToolEntry{
				{Type: "mcp", Name: "broken-mcp", Command: ""},
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
		t.Fatal("expected error for MCP missing command")
	}
	
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	
	if result.ErrorCategory != "tool_load_failure" {
		t.Errorf("expected ErrorCategory='tool_load_failure', got %q", result.ErrorCategory)
	}
}
