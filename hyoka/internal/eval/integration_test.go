package eval

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/ronniegeraghty/hyoka/internal/config"
	"github.com/ronniegeraghty/hyoka/internal/prompt"
)

// --- Integration Tests ---

// TestStubEvalLifecycle_FullRun runs a complete stub evaluation end-to-end
// and verifies reports are generated, results are correct, and exit is clean.
func TestStubEvalLifecycle_FullRun(t *testing.T) {
	outputDir := t.TempDir()
	engine := NewEngine(&StubEvaluator{}, EngineOptions{
		Workers:         2,
		OutputDir:       outputDir,
		GenerateTimeout: 30 * time.Second,
	})

	prompts := []*prompt.Prompt{
		{ID: "storage-auth-go", Service: "storage", Plane: "data-plane", Language: "go", Category: "auth"},
		{ID: "keyvault-crud-py", Service: "keyvault", Plane: "data-plane", Language: "python", Category: "crud"},
	}
	configs := []config.ToolConfig{
		{Name: "baseline", Model: "gpt-4"},
		{Name: "azure-mcp", Model: "claude-sonnet-4.5"},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify cross product: 2 prompts × 2 configs = 4 evaluations
	if summary.TotalEvals != 4 {
		t.Errorf("expected 4 total evals, got %d", summary.TotalEvals)
	}
	if summary.TotalPrompts != 2 {
		t.Errorf("expected 2 total prompts, got %d", summary.TotalPrompts)
	}
	if summary.TotalConfigs != 2 {
		t.Errorf("expected 2 total configs, got %d", summary.TotalConfigs)
	}
	if len(summary.Results) != 4 {
		t.Fatalf("expected 4 results, got %d", len(summary.Results))
	}

	// All stub evals should succeed
	for _, r := range summary.Results {
		if !r.Success {
			t.Errorf("eval %s/%s failed unexpectedly: %s", r.PromptID, r.ConfigName, r.Error)
		}
		if !r.IsStub {
			t.Errorf("expected IsStub=true for %s/%s", r.PromptID, r.ConfigName)
		}
		if r.Duration <= 0 {
			t.Errorf("expected positive duration for %s/%s, got %f", r.PromptID, r.ConfigName, r.Duration)
		}
	}

	// Verify summary counts
	if summary.Passed != 4 {
		t.Errorf("expected 4 passed, got %d", summary.Passed)
	}
	if summary.Failed != 0 {
		t.Errorf("expected 0 failed, got %d", summary.Failed)
	}
	if summary.Errors != 0 {
		t.Errorf("expected 0 errors, got %d", summary.Errors)
	}

	// Verify run duration
	if summary.Duration <= 0 {
		t.Error("expected positive run duration")
	}

	// Verify RunID format (timestamp)
	if summary.RunID == "" {
		t.Error("expected non-empty RunID")
	}

	// Verify report files were written to output directory
	entries, err := os.ReadDir(outputDir)
	if err != nil {
		t.Fatalf("failed to read output dir: %v", err)
	}
	if len(entries) == 0 {
		t.Error("expected report files in output directory")
	}
}

// TestStubEvalLifecycle_SinglePromptSingleConfig verifies the simplest eval case.
func TestStubEvalLifecycle_SinglePromptSingleConfig(t *testing.T) {
	outputDir := t.TempDir()
	engine := NewEngine(&StubEvaluator{}, EngineOptions{
		Workers:         1,
		OutputDir:       outputDir,
		GenerateTimeout: 30 * time.Second,
	})

	prompts := []*prompt.Prompt{
		{ID: "simple-test", Service: "storage", Plane: "data-plane", Language: "go", Category: "auth"},
	}
	configs := []config.ToolConfig{
		{Name: "baseline", Model: "gpt-4"},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if summary.TotalEvals != 1 {
		t.Errorf("expected 1 eval, got %d", summary.TotalEvals)
	}

	r := summary.Results[0]
	if r.PromptID != "simple-test" {
		t.Errorf("expected PromptID 'simple-test', got %q", r.PromptID)
	}
	if r.ConfigName != "baseline" {
		t.Errorf("expected ConfigName 'baseline', got %q", r.ConfigName)
	}

	// Verify prompt metadata is populated
	if r.PromptMeta == nil {
		t.Fatal("expected PromptMeta to be set")
	}
	if r.PromptMeta["service"] != "storage" {
		t.Errorf("expected service='storage' in metadata, got %v", r.PromptMeta["service"])
	}
	if r.PromptMeta["language"] != "go" {
		t.Errorf("expected language='go' in metadata, got %v", r.PromptMeta["language"])
	}

	// Verify config used is populated
	if r.ConfigUsed == nil {
		t.Fatal("expected ConfigUsed to be set")
	}
	if r.ConfigUsed["model"] != "gpt-4" {
		t.Errorf("expected model='gpt-4' in config, got %v", r.ConfigUsed["model"])
	}
}

// TestEvalReport_EnvironmentInfo verifies environment info is populated.
func TestEvalReport_EnvironmentInfo(t *testing.T) {
	outputDir := t.TempDir()
	engine := NewEngine(&StubEvaluator{}, EngineOptions{
		Workers:         1,
		OutputDir:       outputDir,
		GenerateTimeout: 30 * time.Second,
	})

	prompts := []*prompt.Prompt{
		{ID: "env-test", Service: "storage", Plane: "data-plane", Language: "go", Category: "auth"},
	}
	configs := []config.ToolConfig{
		{Name: "test-config", Model: "gpt-4"},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r := summary.Results[0]
	if r.Environment == nil {
		t.Fatal("expected Environment to be set")
	}
	if r.Environment.Model != "gpt-4" {
		t.Errorf("expected model 'gpt-4', got %q", r.Environment.Model)
	}
}

// TestZeroPromptDetection verifies helpful error on empty prompt directory.
func TestZeroPromptDetection(t *testing.T) {
	dir := t.TempDir()

	_, err := prompt.LoadPrompts(dir)
	if err == nil {
		t.Fatal("expected error for empty prompt directory")
	}
	errMsg := err.Error()
	if !strings.Contains(errMsg, "no prompts found") {
		t.Errorf("expected 'no prompts found' in error, got %q", errMsg)
	}
}

// TestZeroPromptDetection_WithNearMiss verifies helpful suggestions in error.
func TestZeroPromptDetection_WithNearMiss(t *testing.T) {
	dir := t.TempDir()
	// Create a file that almost matches the prompt pattern
	os.WriteFile(filepath.Join(dir, "auth-prompt.md"), []byte("# auth"), 0644)

	_, err := prompt.LoadPrompts(dir)
	if err == nil {
		t.Fatal("expected error for directory with only near-miss files")
	}
	errMsg := err.Error()
	if !strings.Contains(errMsg, "Did you mean") {
		t.Errorf("expected 'Did you mean' in error, got %q", errMsg)
	}
	if !strings.Contains(errMsg, "auth.prompt.md") {
		t.Errorf("expected suggested fix 'auth.prompt.md' in error, got %q", errMsg)
	}
}

// TestEvalToolUsage verifies expected vs actual tool comparison logic.
func TestEvalToolUsage(t *testing.T) {
	tests := []struct {
		name         string
		expected     []string
		actual       []string
		wantMatched  int
		wantMissing  int
		wantExtra    int
		wantFullMatch bool
	}{
		{
			name:         "perfect match",
			expected:     []string{"create_file", "run_terminal_command"},
			actual:       []string{"create_file", "run_terminal_command"},
			wantMatched:  2,
			wantMissing:  0,
			wantExtra:    0,
			wantFullMatch: true,
		},
		{
			name:         "missing tools",
			expected:     []string{"create_file", "run_terminal_command"},
			actual:       []string{"create_file"},
			wantMatched:  1,
			wantMissing:  1,
			wantExtra:    0,
			wantFullMatch: false,
		},
		{
			name:         "extra tools",
			expected:     []string{"create_file"},
			actual:       []string{"create_file", "run_terminal_command", "read_file"},
			wantMatched:  1,
			wantMissing:  0,
			wantExtra:    2,
			wantFullMatch: true, // Match=true means all expected tools found (extra tools are OK)
		},
		{
			name:         "no overlap",
			expected:     []string{"create_file"},
			actual:       []string{"run_terminal_command"},
			wantMatched:  0,
			wantMissing:  1,
			wantExtra:    1,
			wantFullMatch: false,
		},
		{
			name:         "both empty",
			expected:     []string{},
			actual:       []string{},
			wantMatched:  0,
			wantMissing:  0,
			wantExtra:    0,
			wantFullMatch: true,
		},
		{
			name:         "nil expected",
			expected:     nil,
			actual:       []string{"create_file"},
			wantMatched:  0,
			wantMissing:  0,
			wantExtra:    1,
			wantFullMatch: true, // No expected tools → nothing missing → match
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := evaluateToolUsage(tt.expected, tt.actual)
			if len(result.MatchedTools) != tt.wantMatched {
				t.Errorf("matched: got %d, want %d", len(result.MatchedTools), tt.wantMatched)
			}
			if len(result.MissingTools) != tt.wantMissing {
				t.Errorf("missing: got %d, want %d", len(result.MissingTools), tt.wantMissing)
			}
			if len(result.ExtraTools) != tt.wantExtra {
				t.Errorf("extra: got %d, want %d", len(result.ExtraTools), tt.wantExtra)
			}
			if result.Match != tt.wantFullMatch {
				t.Errorf("match: got %v, want %v", result.Match, tt.wantFullMatch)
			}
		})
	}
}

// TestDryRun_PromptCounting verifies dry run tracks unique prompts/configs.
func TestDryRun_PromptCounting(t *testing.T) {
	engine := NewEngine(&StubEvaluator{}, EngineOptions{
		Workers: 2,
		DryRun:  true,
	})

	prompts := []*prompt.Prompt{
		{ID: "p1", Service: "storage", Language: "dotnet"},
		{ID: "p2", Service: "keyvault", Language: "python"},
		{ID: "p3", Service: "storage", Language: "java"},
	}
	configs := []config.ToolConfig{
		{Name: "baseline", Model: "gpt-4"},
		{Name: "azure-mcp", Model: "claude-sonnet-4.5"},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if summary.TotalEvals != 6 {
		t.Errorf("expected 6 evals (3×2), got %d", summary.TotalEvals)
	}
	if summary.TotalPrompts != 3 {
		t.Errorf("expected 3 unique prompts, got %d", summary.TotalPrompts)
	}
	if summary.TotalConfigs != 2 {
		t.Errorf("expected 2 unique configs, got %d", summary.TotalConfigs)
	}
	if summary.RunID != "dry-run" {
		t.Errorf("expected run ID 'dry-run', got %q", summary.RunID)
	}
}

// TestWorkspaceContainmentValidation verifies escaped file detection.
func TestWorkspaceContainmentValidation(t *testing.T) {
	dir := t.TempDir()

	// Pre-snapshot: only existing files
	preSnapshot := snapshotDir(dir)

	// Simulate a new file appearing (escaped workspace)
	os.WriteFile(filepath.Join(dir, "escaped_file.py"), []byte("print('oops')"), 0644)

	escaped := ValidateWorkspaceContainment(dir, preSnapshot)
	if len(escaped) != 1 {
		t.Fatalf("expected 1 escaped file, got %d: %v", len(escaped), escaped)
	}
	if escaped[0] != "escaped_file.py" {
		t.Errorf("expected 'escaped_file.py', got %q", escaped[0])
	}
}

func TestWorkspaceContainmentValidation_NoEscape(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "existing.py"), []byte("# exists"), 0644)

	preSnapshot := snapshotDir(dir)

	escaped := ValidateWorkspaceContainment(dir, preSnapshot)
	if len(escaped) != 0 {
		t.Errorf("expected 0 escaped files, got %d: %v", len(escaped), escaped)
	}
}

func TestWorkspaceContainmentValidation_NilSnapshot(t *testing.T) {
	dir := t.TempDir()
	escaped := ValidateWorkspaceContainment(dir, nil)
	if escaped != nil {
		t.Errorf("expected nil for nil snapshot, got %v", escaped)
	}
}

// TestRecoverMisplacedFiles verifies file recovery from non-workspace locations.
func TestRecoverMisplacedFiles(t *testing.T) {
	srcDir := t.TempDir()
	destDir := t.TempDir()

	// Create pre-snapshot
	os.WriteFile(filepath.Join(srcDir, "existing.txt"), []byte("keep"), 0644)
	preSnapshot := snapshotDir(srcDir)

	// Simulate new files appearing in srcDir
	os.WriteFile(filepath.Join(srcDir, "main.py"), []byte("print('recovered')"), 0644)
	os.WriteFile(filepath.Join(srcDir, "app.go"), []byte("package main"), 0644)

	recovered := recoverMisplacedFiles(srcDir, preSnapshot, destDir, "test")
	if recovered != 2 {
		t.Errorf("expected 2 recovered files, got %d", recovered)
	}

	// Verify files were moved to destDir
	for _, name := range []string{"main.py", "app.go"} {
		if _, err := os.Stat(filepath.Join(destDir, name)); err != nil {
			t.Errorf("expected %s in dest dir: %v", name, err)
		}
		// Verify removed from source
		if _, err := os.Stat(filepath.Join(srcDir, name)); !os.IsNotExist(err) {
			t.Errorf("expected %s to be removed from source dir", name)
		}
	}
}

// TestRecoverMisplacedFiles_JunkDirs verifies junk directory cleanup.
func TestRecoverMisplacedFiles_JunkDirs(t *testing.T) {
	srcDir := t.TempDir()
	destDir := t.TempDir()

	preSnapshot := snapshotDir(srcDir)

	// Simulate junk directories appearing
	os.MkdirAll(filepath.Join(srcDir, "__pycache__"), 0755)
	os.MkdirAll(filepath.Join(srcDir, "node_modules"), 0755)
	os.WriteFile(filepath.Join(srcDir, "__pycache__", "cache.pyc"), []byte("cached"), 0644)

	recovered := recoverMisplacedFiles(srcDir, preSnapshot, destDir, "test")
	if recovered != 2 {
		t.Errorf("expected 2 recovered (deleted junk dirs), got %d", recovered)
	}

	// Verify junk dirs are deleted
	if _, err := os.Stat(filepath.Join(srcDir, "__pycache__")); !os.IsNotExist(err) {
		t.Error("expected __pycache__ to be deleted")
	}
	if _, err := os.Stat(filepath.Join(srcDir, "node_modules")); !os.IsNotExist(err) {
		t.Error("expected node_modules to be deleted")
	}
}

// TestSnapshotDir verifies directory snapshot creation.
func TestSnapshotDir(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "file1.txt"), []byte("a"), 0644)
	os.WriteFile(filepath.Join(dir, "file2.txt"), []byte("b"), 0644)
	os.MkdirAll(filepath.Join(dir, "subdir"), 0755)
	// Hidden files should be excluded
	os.WriteFile(filepath.Join(dir, ".hidden"), []byte("h"), 0644)

	snap := snapshotDir(dir)
	if len(snap) != 3 {
		t.Errorf("expected 3 entries in snapshot, got %d: %v", len(snap), snap)
	}
	if !snap["file1.txt"] || !snap["file2.txt"] || !snap["subdir"] {
		t.Errorf("expected file1.txt, file2.txt, subdir in snapshot, got %v", snap)
	}
	if snap[".hidden"] {
		t.Error("hidden files should not be in snapshot")
	}
}

// TestNewWorkspaceAt verifies persistent workspace creation.
func TestNewWorkspaceAt_PersistentCleanup(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "persistent-ws")

	ws, err := NewWorkspaceAt(dir)
	if err != nil {
		t.Fatalf("NewWorkspaceAt failed: %v", err)
	}

	// Verify directory was created
	if _, err := os.Stat(ws.Dir); err != nil {
		t.Fatalf("workspace dir not created: %v", err)
	}

	// Cleanup should be a no-op for persistent workspaces
	if err := ws.Cleanup(); err != nil {
		t.Fatalf("cleanup error: %v", err)
	}
	// Directory should still exist
	if _, err := os.Stat(ws.Dir); err != nil {
		t.Error("persistent workspace should not be deleted by Cleanup")
	}
}

// TestTimedOutEval_ErrorDetails verifies timeout error details.
func TestTimedOutEval_ErrorDetails(t *testing.T) {
	outputDir := t.TempDir()
	engine := NewEngine(&slowEvaluator{}, EngineOptions{
		Workers:         1,
		OutputDir:       outputDir,
		GenerateTimeout: 100 * time.Millisecond,
	})

	prompts := []*prompt.Prompt{
		{ID: "timeout-detail", Service: "storage", Plane: "data-plane", Language: "go", Category: "auth"},
	}
	configs := []config.ToolConfig{
		{Name: "baseline", Model: "gpt-4"},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r := summary.Results[0]
	if r.ErrorCategory != "timeout" {
		t.Errorf("expected ErrorCategory='timeout', got %q", r.ErrorCategory)
	}
	if !strings.Contains(r.Error, "timed out") {
		t.Errorf("expected 'timed out' in error, got %q", r.Error)
	}
	if !strings.Contains(r.ErrorDetails, "generate-timeout") {
		t.Errorf("expected '--generate-timeout' suggestion in details, got %q", r.ErrorDetails)
	}
}

// TestEvalReport_GuardrailFieldsAlwaysRecorded verifies guardrail limits are
// always written to the report, even when not triggered.
func TestEvalReport_GuardrailFieldsAlwaysRecorded(t *testing.T) {
	outputDir := t.TempDir()
	engine := NewEngine(&StubEvaluator{}, EngineOptions{
		Workers:         1,
		OutputDir:       outputDir,
		GenerateTimeout: 30 * time.Second,
		MaxTurns:        42,
		MaxFiles:        99,
		MaxOutputSize:   2097152,
	})

	prompts := []*prompt.Prompt{
		{ID: "gr-fields", Service: "storage", Plane: "data-plane", Language: "go", Category: "auth"},
	}
	configs := []config.ToolConfig{{Name: "baseline", Model: "gpt-4"}}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r := summary.Results[0]
	if r.GuardrailMaxTurns != 42 {
		t.Errorf("GuardrailMaxTurns = %d, want 42", r.GuardrailMaxTurns)
	}
	if r.GuardrailMaxFiles != 99 {
		t.Errorf("GuardrailMaxFiles = %d, want 99", r.GuardrailMaxFiles)
	}
	if r.GuardrailMaxOutputSize != 2097152 {
		t.Errorf("GuardrailMaxOutputSize = %d, want 2097152", r.GuardrailMaxOutputSize)
	}
	if r.GuardrailAbortReason != "" {
		t.Errorf("GuardrailAbortReason = %q, want empty (not triggered)", r.GuardrailAbortReason)
	}
}
