//go:build integration

package eval

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/ronniegeraghty/hyoka/internal/config"
	"github.com/ronniegeraghty/hyoka/internal/prompt"
)

func TestEngineRunWithGraders(t *testing.T) {
	// Create a temporary graders directory with a config file
	gradersDir := t.TempDir()
	graderYAML := `graders:
  - kind: file
    name: "main_exists"
    config:
      path: "stub_output.txt"
    weight: 1.0
    gate: true
`
	if err := os.WriteFile(filepath.Join(gradersDir, "test.yaml"), []byte(graderYAML), 0644); err != nil {
		t.Fatal(err)
	}

	outputDir := t.TempDir()
	engine := NewEngine(&StubEvaluator{}, quietOpts(EngineOptions{
		Workers:    1,
		OutputDir:  outputDir,
		GradersDir: gradersDir,
	}))

	prompts := []*prompt.Prompt{
		{ID: "grader-test", Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "python", "category": "crud"}},
	}
	configs := []config.ToolConfig{
		{Name: "test-config", Generator: &config.GeneratorConfig{Model: "gpt-4"}},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(summary.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(summary.Results))
	}

	r := summary.Results[0]
	if len(r.GraderResults) == 0 {
		t.Fatal("expected GraderResults to be populated")
	}
	if len(r.GraderResults) != 1 {
		t.Fatalf("expected 1 grader result, got %d", len(r.GraderResults))
	}
	gr := r.GraderResults[0]
	if gr.GraderName != "main_exists" {
		t.Errorf("expected grader name 'main_exists', got %q", gr.GraderName)
	}
	if gr.GraderType != "file" {
		t.Errorf("expected type 'file', got %q", gr.GraderType)
	}
	// StubEvaluator returns GeneratedFiles: ["stub_output.txt"]
	// The file grader checks path: "stub_output.txt" — should pass
	if !r.Success {
		t.Errorf("expected eval to succeed, got failure: %s", r.FailureReason)
	}
}

func TestEngineRunWithGraderGateFails(t *testing.T) {
	gradersDir := t.TempDir()
	graderYAML := `graders:
  - kind: file
    name: "missing_file"
    config:
      path: "does_not_exist.py"
    weight: 1.0
    gate: true
`
	if err := os.WriteFile(filepath.Join(gradersDir, "test.yaml"), []byte(graderYAML), 0644); err != nil {
		t.Fatal(err)
	}

	outputDir := t.TempDir()
	engine := NewEngine(&StubEvaluator{}, quietOpts(EngineOptions{
		Workers:    1,
		OutputDir:  outputDir,
		GradersDir: gradersDir,
	}))

	prompts := []*prompt.Prompt{
		{ID: "gate-fail-test", Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "python", "category": "crud"}},
	}
	configs := []config.ToolConfig{
		{Name: "test-config", Generator: &config.GeneratorConfig{Model: "gpt-4"}},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r := summary.Results[0]
	if len(r.GraderResults) == 0 {
		t.Fatal("expected GraderResults to be populated")
	}
	if r.Success {
		t.Error("expected eval to fail when gate grader fails")
	}
	if r.FailureReason == "" {
		t.Error("expected failure reason to be set for gate failure")
	}
	if r.Success {
		t.Error("expected eval to fail when gate grader fails")
	}
}

func TestEngineRunWithGraderWhenFilter(t *testing.T) {
	gradersDir := t.TempDir()
	graderYAML := `graders:
  - kind: file
    name: "python_only"
    config:
      path: "stub_output.txt"
    weight: 1.0
    when:
      language: python
  - kind: file
    name: "go_only"
    config:
      path: "go_output.go"
    weight: 1.0
    when:
      language: go
`
	if err := os.WriteFile(filepath.Join(gradersDir, "test.yaml"), []byte(graderYAML), 0644); err != nil {
		t.Fatal(err)
	}

	outputDir := t.TempDir()
	engine := NewEngine(&StubEvaluator{}, quietOpts(EngineOptions{
		Workers:    1,
		OutputDir:  outputDir,
		GradersDir: gradersDir,
	}))

	// Python prompt — only python_only grader should apply
	prompts := []*prompt.Prompt{
		{ID: "filter-test", Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "python", "category": "crud"}},
	}
	configs := []config.ToolConfig{
		{Name: "test-config", Generator: &config.GeneratorConfig{Model: "gpt-4"}},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r := summary.Results[0]
	if len(r.GraderResults) == 0 {
		t.Fatal("expected GraderResults to be populated")
	}
	if len(r.GraderResults) != 1 {
		t.Fatalf("expected 1 grader result (python_only), got %d", len(r.GraderResults))
	}
	if r.GraderResults[0].GraderName != "python_only" {
		t.Errorf("expected grader name 'python_only', got %q", r.GraderResults[0].GraderName)
	}
}

func TestEngineRunNoGradersConfigured(t *testing.T) {
	// No GradersDir set — should fall back to reviewer pipeline only
	outputDir := t.TempDir()
	engine := NewEngine(&StubEvaluator{}, quietOpts(EngineOptions{
		Workers:   1,
		OutputDir: outputDir,
	}))

	prompts := []*prompt.Prompt{
		{ID: "no-graders-test", Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "python", "category": "crud"}},
	}
	configs := []config.ToolConfig{
		{Name: "test-config", Generator: &config.GeneratorConfig{Model: "gpt-4"}},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r := summary.Results[0]
	if r.GraderResults != nil {
		t.Error("expected nil GraderResults when no graders configured")
	}
	// Success should still be set by the stub evaluator
	if !r.Success {
		t.Error("expected success when no graders and no review")
	}
}
