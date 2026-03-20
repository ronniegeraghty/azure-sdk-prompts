package eval

import (
"context"
"os"
"testing"
"time"

"github.com/ronniegeraghty/azure-sdk-prompts/tool/internal/config"
"github.com/ronniegeraghty/azure-sdk-prompts/tool/internal/prompt"
)

func TestStubEvaluator(t *testing.T) {
stub := &StubEvaluator{}
p := &prompt.Prompt{ID: "test-prompt", Language: "go"}
cfg := &config.ToolConfig{Name: "test-config", Model: "gpt-4"}

result, err := stub.Evaluate(context.Background(), p, cfg, t.TempDir())
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if !result.Success {
t.Error("expected stub to succeed")
}
if len(result.GeneratedFiles) == 0 {
t.Error("expected stub to return generated files")
}
}

func TestEngineDryRun(t *testing.T) {
engine := NewEngine(&StubEvaluator{}, EngineOptions{
Workers: 2,
DryRun:  true,
})

prompts := []*prompt.Prompt{
{ID: "p1", Service: "storage", Language: "dotnet"},
{ID: "p2", Service: "keyvault", Language: "python"},
}
configs := []config.ToolConfig{
{Name: "baseline", Model: "gpt-4"},
{Name: "azure-mcp", Model: "claude-sonnet-4.5"},
}

summary, err := engine.Run(context.Background(), prompts, configs)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if summary.RunID != "dry-run" {
t.Errorf("expected run ID 'dry-run', got %q", summary.RunID)
}
if summary.TotalEvals != 4 {
t.Errorf("expected 4 evaluations (2 prompts x 2 configs), got %d", summary.TotalEvals)
}
if summary.TotalPrompts != 2 {
t.Errorf("expected 2 prompts, got %d", summary.TotalPrompts)
}
if summary.TotalConfigs != 2 {
t.Errorf("expected 2 configs, got %d", summary.TotalConfigs)
}
}

func TestEngineRun(t *testing.T) {
outputDir := t.TempDir()
engine := NewEngine(&StubEvaluator{}, EngineOptions{
Workers:   1,
Timeout:   30 * time.Second,
OutputDir: outputDir,
})

prompts := []*prompt.Prompt{
{ID: "test-prompt", Service: "storage", Plane: "data-plane", Language: "go", Category: "auth"},
}
configs := []config.ToolConfig{
{Name: "test-config", Model: "gpt-4"},
}

summary, err := engine.Run(context.Background(), prompts, configs)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if summary.TotalEvals != 1 {
t.Errorf("expected 1 evaluation, got %d", summary.TotalEvals)
}
}

func TestNewWorkspace(t *testing.T) {
baseDir := t.TempDir()
ws, err := NewWorkspace(baseDir, "test-prompt", "test-config")
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if ws.Dir == "" {
t.Error("expected non-empty workspace dir")
}
// Verify directory exists
info, err := os.Stat(ws.Dir)
if err != nil {
t.Fatalf("workspace dir does not exist: %v", err)
}
if !info.IsDir() {
t.Error("expected workspace to be a directory")
}

// Cleanup
if err := ws.Cleanup(); err != nil {
t.Fatalf("cleanup failed: %v", err)
}
if _, err := os.Stat(ws.Dir); !os.IsNotExist(err) {
t.Error("expected workspace to be removed after cleanup")
}
}
