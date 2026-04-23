package eval

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/process"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/prompt"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/report"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/review"
)

func TestMain(m *testing.M) {
	// Suppress slog output during tests to keep output clean.
	slog.SetDefault(slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{Level: slog.LevelError + 1})))
	os.Exit(m.Run())
}

// quietOpts returns EngineOptions with stdout suppressed and an isolated
// ProcessTracker so tests never scan/kill real Copilot CLI processes.
func quietOpts(opts EngineOptions) EngineOptions {
	opts.Stdout = io.Discard
	opts.Tracker = &process.ProcessTracker{}
	return opts
}

// slowRunner blocks until context cancellation, simulating a timeout.
type slowRunner struct{}

func (s *slowRunner) Run(ctx context.Context, _ *prompt.Prompt, _ *config.ToolConfig, _ string) (*EvalResult, error) {
	<-ctx.Done()
	return nil, fmt.Errorf("prompt send failed: %w", ctx.Err())
}

func TestStubRunner(t *testing.T) {
stub := &StubRunner{}
p := &prompt.Prompt{ID: "test-prompt", Properties: map[string]string{"language": "go"}}
cfg := &config.ToolConfig{Name: "test-config", Generator: &config.GeneratorConfig{Model: "gpt-4"}}

result, err := stub.Run(context.Background(), p, cfg, t.TempDir())
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if !result.Success {
t.Error("expected stub to succeed")
}
if len(result.GeneratedFiles) == 0 {
t.Error("expected stub to return generated files")
}
if !result.IsStub {
t.Error("expected IsStub to be true for stub evaluator")
}
}

func TestEngineDryRun(t *testing.T) {
engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{
Workers: 2,
DryRun:  true,
}))

prompts := []*prompt.Prompt{
{ID: "p1", Properties: map[string]string{"service": "storage", "language": "dotnet"}},
{ID: "p2", Properties: map[string]string{"service": "keyvault", "language": "python"}},
}
configs := []config.ToolConfig{
{Name: "baseline", Generator: &config.GeneratorConfig{Model: "gpt-4"}},
{Name: "azure-mcp", Generator: &config.GeneratorConfig{Model: "claude-sonnet-4.5"}},
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

func TestDryRunSkillDir_Populated(t *testing.T) {
	dir := t.TempDir()
	// Create a skill_dir with 2 skill subdirectories
	skillsDir := filepath.Join(dir, "skills")
	for _, name := range []string{"skill-a", "skill-b"} {
		subDir := filepath.Join(skillsDir, name)
		if err := os.MkdirAll(subDir, 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(subDir, "SKILL.md"), []byte("# Skill"), 0644); err != nil {
			t.Fatal(err)
		}
	}

	var buf strings.Builder
	opts := quietOpts(EngineOptions{Workers: 1, DryRun: true})
	opts.Stdout = &buf
	engine := NewEngine(&StubRunner{}, opts)

	prompts := []*prompt.Prompt{
		{ID: "p1", Properties: map[string]string{"language": "python"}},
	}
	configs := []config.ToolConfig{
		{
			Name: "populated-skills",
			Generator: &config.GeneratorConfig{
				Model: "gpt-4",
				Tools: []config.ToolEntry{
					{Name: "gen-skills", Type: "skill", Source: "local", SkillDir: true, Path: skillsDir},
				},
			},
		},
	}

	_, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	output := buf.String()
	if !strings.Contains(output, "1 generator dir(s) searched, 2 skill(s) found") {
		t.Errorf("expected '1 generator dir(s) searched, 2 skill(s) found' in output, got:\n%s", output)
	}
}

func TestDryRunSkillDir_Empty(t *testing.T) {
	dir := t.TempDir()
	// Create an empty skill_dir (only .gitkeep)
	skillsDir := filepath.Join(dir, "skills")
	if err := os.MkdirAll(skillsDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(skillsDir, ".gitkeep"), nil, 0644); err != nil {
		t.Fatal(err)
	}

	var buf strings.Builder
	opts := quietOpts(EngineOptions{Workers: 1, DryRun: true})
	opts.Stdout = &buf
	engine := NewEngine(&StubRunner{}, opts)

	prompts := []*prompt.Prompt{
		{ID: "p1", Properties: map[string]string{"language": "python"}},
	}
	configs := []config.ToolConfig{
		{
			Name: "empty-skills",
			Generator: &config.GeneratorConfig{
				Model: "gpt-4",
				Tools: []config.ToolEntry{
					{Name: "gen-skills", Type: "skill", Source: "local", SkillDir: true, Path: skillsDir},
				},
			},
		},
	}

	_, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	output := buf.String()
	if !strings.Contains(output, "1 generator dir(s) searched, 0 skill(s) found") {
		t.Errorf("expected '1 generator dir(s) searched, 0 skill(s) found' in output, got:\n%s", output)
	}
}

func TestDryRunSkillDir_NonExistent(t *testing.T) {
	var buf strings.Builder
	opts := quietOpts(EngineOptions{Workers: 1, DryRun: true})
	opts.Stdout = &buf
	engine := NewEngine(&StubRunner{}, opts)

	prompts := []*prompt.Prompt{
		{ID: "p1", Properties: map[string]string{"language": "python"}},
	}
	configs := []config.ToolConfig{
		{
			Name: "missing-dir",
			Generator: &config.GeneratorConfig{
				Model: "gpt-4",
				Tools: []config.ToolEntry{
					{Name: "gen-skills", Type: "skill", Source: "local", SkillDir: true, Path: "/does/not/exist"},
				},
			},
		},
	}

	_, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	output := buf.String()
	if !strings.Contains(output, "1 generator dir(s) searched, 0 skill(s) found") {
		t.Errorf("expected '1 generator dir(s) searched, 0 skill(s) found' in output, got:\n%s", output)
	}
}

func TestDryRunSingleSkill_MissingSKILLMD(t *testing.T) {
	dir := t.TempDir()
	// Create a directory but no SKILL.md
	skillDir := filepath.Join(dir, "my-skill")
	if err := os.Mkdir(skillDir, 0755); err != nil {
		t.Fatal(err)
	}

	var buf strings.Builder
	opts := quietOpts(EngineOptions{Workers: 1, DryRun: true})
	opts.Stdout = &buf
	engine := NewEngine(&StubRunner{}, opts)

	prompts := []*prompt.Prompt{
		{ID: "p1", Properties: map[string]string{"language": "python"}},
	}
	configs := []config.ToolConfig{
		{
			Name: "bad-single-skill",
			Generator: &config.GeneratorConfig{
				Model: "gpt-4",
				Tools: []config.ToolEntry{
					{Name: "my-skill", Type: "skill", Source: "local", Path: skillDir},
				},
			},
		},
	}

	_, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	output := buf.String()
	if !strings.Contains(output, "1 generator dir(s) searched, 0 skill(s) found") {
		t.Errorf("expected '1 generator dir(s) searched, 0 skill(s) found' in output, got:\n%s", output)
	}
}

func TestEngineRun(t *testing.T) {
outputDir := t.TempDir()
engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{
Workers:   1,
OutputDir: outputDir,
}))

prompts := []*prompt.Prompt{
{ID: "test-prompt", Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "go", "category": "auth"}},
}
configs := []config.ToolConfig{
{Name: "test-config", Generator: &config.GeneratorConfig{Model: "gpt-4"}},
}

summary, err := engine.Run(context.Background(), prompts, configs)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if summary.TotalEvals != 1 {
t.Errorf("expected 1 evaluation, got %d", summary.TotalEvals)
}
}

func TestEngineRunCapturesGeneratedFiles(t *testing.T) {
	// The evaluator returns GeneratedFiles in its result, but may not leave
	// files on disk (e.g., SDK cleanup removes them). The engine must use the
	// evaluator's captured list rather than relying solely on ws.ListFiles().
	outputDir := t.TempDir()
	engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{
		Workers:   1,
		OutputDir: outputDir,
	}))

	prompts := []*prompt.Prompt{
		{ID: "filelist-test", Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "python", "category": "crud"}},
	}
	configs := []config.ToolConfig{
		{Name: "baseline", Generator: &config.GeneratorConfig{Model: "gpt-4"}},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(summary.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(summary.Results))
	}
	r := summary.Results[0]
	if len(r.GeneratedFiles) == 0 {
		t.Error("expected GeneratedFiles to be populated from evaluator result, got 0 files")
	}
}

func TestEngineRunTimeoutError(t *testing.T) {
	// An evaluator that blocks until the context is cancelled.
	slowEval := &slowRunner{}
	outputDir := t.TempDir()
	engine := NewEngine(slowEval, quietOpts(EngineOptions{
		Workers:   1,
		OutputDir: outputDir,
	}))

	prompts := []*prompt.Prompt{
		{ID: "timeout-test", Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "go", "category": "auth"}},
	}
	configs := []config.ToolConfig{
		{Name: "baseline", Generator: &config.GeneratorConfig{Model: "gpt-4"}},
	}

	// Use a short-lived context to simulate cancellation
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	summary, err := engine.Run(ctx, prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(summary.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(summary.Results))
	}
	r := summary.Results[0]
	if r.Error == "" {
		t.Fatal("expected error in report for cancelled eval")
	}
}

func TestNewWorkspace(t *testing.T) {
ws, err := NewWorkspace("test-prompt", "test-config")
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

// Verify it's in temp dir, not in reports
if !strings.HasPrefix(ws.Dir, os.TempDir()) {
t.Errorf("expected workspace in temp dir, got %s", ws.Dir)
}

// Test ListFiles on empty workspace
files, err := ws.ListFiles()
if err != nil {
t.Fatalf("ListFiles failed: %v", err)
}
if len(files) != 0 {
t.Errorf("expected 0 files in empty workspace, got %d", len(files))
}

// Create a test file and verify ListFiles
testFile := filepath.Join(ws.Dir, "test.py")
if err := os.WriteFile(testFile, []byte("print('hello')"), 0644); err != nil {
t.Fatalf("failed to write test file: %v", err)
}
files, err = ws.ListFiles()
if err != nil {
t.Fatalf("ListFiles failed: %v", err)
}
if len(files) != 1 || files[0] != "test.py" {
t.Errorf("expected [test.py], got %v", files)
}

// Test CopyFilesTo
destDir := t.TempDir()
copied, err := ws.CopyFilesTo(destDir)
if err != nil {
t.Fatalf("CopyFilesTo failed: %v", err)
}
if len(copied) != 1 || copied[0] != "test.py" {
t.Errorf("expected [test.py] copied, got %v", copied)
}
destFile := filepath.Join(destDir, "test.py")
data, err := os.ReadFile(destFile)
if err != nil {
t.Fatalf("failed to read copied file: %v", err)
}
if string(data) != "print('hello')" {
t.Errorf("unexpected file content: %s", data)
}

// Cleanup
if err := ws.Cleanup(); err != nil {
t.Fatalf("cleanup failed: %v", err)
}
if _, err := os.Stat(ws.Dir); !os.IsNotExist(err) {
t.Error("expected workspace to be removed after cleanup")
}
}

// manyFilesRunner generates N files to trigger the max-files guardrail.
type manyFilesRunner struct {
	fileCount int
}

func (m *manyFilesRunner) Run(ctx context.Context, p *prompt.Prompt, cfg *config.ToolConfig, workDir string) (*EvalResult, error) {
	var files []string
	for i := 0; i < m.fileCount; i++ {
		name := fmt.Sprintf("file_%d.txt", i)
		path := filepath.Join(workDir, name)
		os.WriteFile(path, []byte("content"), 0644)
		files = append(files, name)
	}
	return &EvalResult{
		GeneratedFiles: files,
		Success:        true,
		IsStub:         true,
	}, nil
}

// manyTurnsRunner produces session events to trigger the max-turns guardrail.
type manyTurnsRunner struct {
	turnCount int
}

func (m *manyTurnsRunner) Run(ctx context.Context, p *prompt.Prompt, cfg *config.ToolConfig, workDir string) (*EvalResult, error) {
	name := "output.txt"
	os.WriteFile(filepath.Join(workDir, name), []byte("hello"), 0644)
	var events []report.SessionEventRecord
	for i := 0; i < m.turnCount; i++ {
		events = append(events, report.SessionEventRecord{Type: "assistant.message"})
	}
	return &EvalResult{
		GeneratedFiles: []string{name},
		SessionEvents:  events,
		Success:        true,
		IsStub:         true,
	}, nil
}

func TestGuardrailMaxFiles(t *testing.T) {
	outputDir := t.TempDir()
	engine := NewEngine(&manyFilesRunner{fileCount: 10}, quietOpts(EngineOptions{
		Workers:   1,
		OutputDir: outputDir,
		SkipReview: true,
		MaxFiles:  5,
	}))

	prompts := []*prompt.Prompt{
		{ID: "guardrail-files", Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "go", "category": "auth"}},
	}
	configs := []config.ToolConfig{
		{Name: "test", Generator: &config.GeneratorConfig{Model: "gpt-4"}},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(summary.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(summary.Results))
	}
	r := summary.Results[0]
	// MaxFiles remains a HARD FAIL (the byte-size guardrail was dropped in #566).
	if r.Success {
		t.Error("expected guardrail to fail the eval")
	}
	if !strings.Contains(r.GuardrailAbortReason, "file count") {
		t.Errorf("expected guardrail abort reason about file count, got %q", r.GuardrailAbortReason)
	}
}

func TestGuardrailMaxTurns(t *testing.T) {
	outputDir := t.TempDir()
	// 30 assistant.message events exceeds the default MaxTurns=25.
	engine := NewEngine(&manyTurnsRunner{turnCount: 30}, quietOpts(EngineOptions{
		Workers:    1,
		OutputDir:  outputDir,
		SkipReview: true,
	}))

	prompts := []*prompt.Prompt{
		{ID: "guardrail-turns", Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "go", "category": "auth"}},
	}
	configs := []config.ToolConfig{
		{Name: "test", Generator: &config.GeneratorConfig{Model: "gpt-4"}},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(summary.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(summary.Results))
	}
	r := summary.Results[0]
	if r.Success {
		t.Error("expected guardrail to fail the eval")
	}
	if !strings.Contains(r.GuardrailAbortReason, "turn count") {
		t.Errorf("expected guardrail abort reason about turn count, got %q", r.GuardrailAbortReason)
	}
}

// TestActionLimitSoftCap verifies that exceeding the session action limit is
// treated as a soft cap: the eval is NOT counted as an error, and the review
// phase determines pass/fail. The report records action_limit_reached and
// action_count for transparency.
func TestActionLimitSoftCap(t *testing.T) {
	outputDir := t.TempDir()
	// 10 events with MaxSessionActions=5 triggers the soft cap.
	// MaxTurns=999 prevents the turn-count hard guardrail from firing.
	engine := NewEngine(&manyTurnsRunner{turnCount: 10}, quietOpts(EngineOptions{
		Workers:           1,
		OutputDir:         outputDir,
		SkipReview:        true,
		MaxTurns:          999,
		MaxSessionActions: 5,
	}))

	prompts := []*prompt.Prompt{
		{ID: "soft-cap-test", Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "go", "category": "auth"}},
	}
	configs := []config.ToolConfig{
		{Name: "test", Generator: &config.GeneratorConfig{Model: "gpt-4"}},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(summary.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(summary.Results))
	}
	r := summary.Results[0]

	// Soft cap: no error set, so the eval is NOT counted as "Errors".
	if r.Error != "" {
		t.Errorf("expected no error for action limit soft cap, got %q", r.Error)
	}
	// With SkipReview, Success stays as returned by evaluator (true).
	if !r.Success {
		t.Error("expected success=true when action limit is soft cap with no review")
	}
	// The report records that the action limit was reached.
	if !r.ActionLimitReached {
		t.Error("expected ActionLimitReached=true")
	}
	if r.ActionCount != 10 {
		t.Errorf("expected ActionCount=10, got %d", r.ActionCount)
	}
	// Summary: counted as Passed (not Error).
	if summary.Passed != 1 {
		t.Errorf("expected Passed=1, got %d", summary.Passed)
	}
	if summary.Errors != 0 {
		t.Errorf("expected Errors=0, got %d", summary.Errors)
	}
}

// TestGuardrailMaxNewFiles and the byte-size guardrail test were removed
// (#566 amendment): the byte-size guardrail was dropped entirely and the
// new-files soft cap was rolled back. MaxFiles=50 hard fail remains the
// agent-output backstop.

// TestWorkspaceDeltaCaptured verifies WorkspaceDelta is populated on every
// successful eval (#566) and reaches the report (grader coverage via #571 nil-safety tests).
func TestWorkspaceDeltaCaptured(t *testing.T) {
	outputDir := t.TempDir()
	engine := NewEngine(&manyFilesRunner{fileCount: 3}, quietOpts(EngineOptions{
		Workers:    1,
		OutputDir:  outputDir,
		SkipReview: true,
	}))

	prompts := []*prompt.Prompt{
		{ID: "delta-capture", Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "go", "category": "auth"}},
	}
	configs := []config.ToolConfig{
		{Name: "test", Generator: &config.GeneratorConfig{Model: "gpt-4"}},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	r := summary.Results[0]
	if r.WorkspaceDelta == nil {
		t.Fatal("expected WorkspaceDelta to be populated")
	}
	if got := r.WorkspaceDelta.NewFileCount; got != 3 {
		t.Errorf("NewFileCount: expected 3, got %d", got)
	}
	if r.WorkspaceDelta.BytesAdded == 0 {
		t.Error("expected BytesAdded > 0")
	}
}

func TestGuardrailDefaultValues(t *testing.T) {
	engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{}))
	if engine.opts.MaxTurns != 25 {
		t.Errorf("default MaxTurns: expected 25, got %d", engine.opts.MaxTurns)
	}
	if engine.opts.MaxSessionActions != 50 {
		t.Errorf("default MaxSessionActions: expected 50, got %d", engine.opts.MaxSessionActions)
	}
	if engine.opts.MaxFiles != 50 {
		t.Errorf("default MaxFiles: expected 50, got %d", engine.opts.MaxFiles)
	}
}

func TestResolveLimitsNilFallsBackToDefaults(t *testing.T) {
	engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{}))
	cfg := config.ToolConfig{Name: "no-limits", Generator: &config.GeneratorConfig{Model: "gpt-4"}}
	lim := engine.resolveLimits(cfg, nil)
	if lim.maxTurns != 25 { t.Errorf("expected maxTurns 25, got %d", lim.maxTurns) }
	if lim.maxSessionActions != 50 { t.Errorf("expected maxSessionActions 50, got %d", lim.maxSessionActions) }
	if lim.maxFiles != 50 { t.Errorf("expected maxFiles 50, got %d", lim.maxFiles) }
}

func TestResolveLimitsZeroFieldsFallBackToDefaults(t *testing.T) {
	engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{}))
	cfg := config.ToolConfig{Name: "zero", Generator: &config.GeneratorConfig{Model: "gpt-4"}, Limits: &config.SessionLimits{}}
	lim := engine.resolveLimits(cfg, nil)
	if lim.maxTurns != 25 { t.Errorf("expected 25, got %d", lim.maxTurns) }
	if lim.maxSessionActions != 50 { t.Errorf("expected 50, got %d", lim.maxSessionActions) }
	if lim.maxFiles != 50 { t.Errorf("expected 50, got %d", lim.maxFiles) }
}

func TestResolveLimitsConfigOverridesDefaults(t *testing.T) {
	engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{}))
	cfg := config.ToolConfig{Name: "custom", Generator: &config.GeneratorConfig{Model: "gpt-4"}, Limits: &config.SessionLimits{MaxTurns: 10, MaxFiles: 20, MaxSessionActions: 30}}
	lim := engine.resolveLimits(cfg, nil)
	if lim.maxTurns != 10 { t.Errorf("expected 10, got %d", lim.maxTurns) }
	if lim.maxSessionActions != 30 { t.Errorf("expected 30, got %d", lim.maxSessionActions) }
	if lim.maxFiles != 20 { t.Errorf("expected 20, got %d", lim.maxFiles) }
}

func TestResolveLimitsPartialOverride(t *testing.T) {
	engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{}))
	cfg := config.ToolConfig{Name: "partial", Generator: &config.GeneratorConfig{Model: "gpt-4"}, Limits: &config.SessionLimits{MaxTurns: 10}}
	lim := engine.resolveLimits(cfg, nil)
	if lim.maxTurns != 10 { t.Errorf("expected 10, got %d", lim.maxTurns) }
	if lim.maxSessionActions != 50 { t.Errorf("expected 50, got %d", lim.maxSessionActions) }
	if lim.maxFiles != 50 { t.Errorf("expected 50, got %d", lim.maxFiles) }
}

func TestResolveLimitsPromptOverridesConfig(t *testing.T) {
	engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{}))
	cfg := config.ToolConfig{Name: "cfg", Generator: &config.GeneratorConfig{Model: "gpt-4"}, Limits: &config.SessionLimits{MaxTurns: 10, MaxSessionActions: 30}}
	p := &prompt.Prompt{ID: "test", MaxSessionActions: 100, MaxTurns: 40}
	lim := engine.resolveLimits(cfg, p)
	if lim.maxTurns != 40 { t.Errorf("expected prompt maxTurns 40, got %d", lim.maxTurns) }
	if lim.maxSessionActions != 100 { t.Errorf("expected prompt maxSessionActions 100, got %d", lim.maxSessionActions) }
	if lim.maxFiles != 50 { t.Errorf("expected default maxFiles 50, got %d", lim.maxFiles) }
}

func TestResolveLimitsPromptOverridesDefaults(t *testing.T) {
	engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{}))
	cfg := config.ToolConfig{Name: "no-limits", Generator: &config.GeneratorConfig{Model: "gpt-4"}}
	p := &prompt.Prompt{ID: "test", MaxSessionActions: 75}
	lim := engine.resolveLimits(cfg, p)
	if lim.maxSessionActions != 75 { t.Errorf("expected prompt maxSessionActions 75, got %d", lim.maxSessionActions) }
	if lim.maxTurns != 25 { t.Errorf("expected default maxTurns 25, got %d", lim.maxTurns) }
}

func TestResolveLimitsPromptPartialOverride(t *testing.T) {
	engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{}))
	cfg := config.ToolConfig{Name: "cfg", Generator: &config.GeneratorConfig{Model: "gpt-4"}, Limits: &config.SessionLimits{MaxTurns: 10, MaxSessionActions: 30}}
	p := &prompt.Prompt{ID: "test", MaxSessionActions: 100}
	lim := engine.resolveLimits(cfg, p)
	if lim.maxTurns != 10 { t.Errorf("expected config maxTurns 10, got %d", lim.maxTurns) }
	if lim.maxSessionActions != 100 { t.Errorf("expected prompt maxSessionActions 100, got %d", lim.maxSessionActions) }
}

func TestResolveLimitsPromptZeroDoesNotOverride(t *testing.T) {
	engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{}))
	cfg := config.ToolConfig{Name: "cfg", Generator: &config.GeneratorConfig{Model: "gpt-4"}, Limits: &config.SessionLimits{MaxSessionActions: 30}}
	p := &prompt.Prompt{ID: "test", MaxSessionActions: 0, MaxTurns: 0}
	lim := engine.resolveLimits(cfg, p)
	if lim.maxSessionActions != 30 { t.Errorf("expected config maxSessionActions 30, got %d", lim.maxSessionActions) }
	if lim.maxTurns != 25 { t.Errorf("expected default maxTurns 25, got %d", lim.maxTurns) }
}

func TestConfigLimitsRespectedByGuardrail(t *testing.T) {
	outputDir := t.TempDir()
	engine := NewEngine(&manyTurnsRunner{turnCount: 15}, quietOpts(EngineOptions{Workers: 1, OutputDir: outputDir, SkipReview: true}))
	prompts := []*prompt.Prompt{{ID: "config-limit-test", Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "go", "category": "auth"}}}
	configs := []config.ToolConfig{{Name: "strict", Generator: &config.GeneratorConfig{Model: "gpt-4"}, Limits: &config.SessionLimits{MaxTurns: 5, MaxSessionActions: 99}}}
	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil { t.Fatalf("unexpected error: %v", err) }
	r := summary.Results[0]
	if r.Success { t.Error("expected guardrail to fail") }
	if !strings.Contains(r.GuardrailAbortReason, "turn count") { t.Errorf("expected turn count guardrail, got %q", r.GuardrailAbortReason) }
	if r.GuardrailMaxTurns != 5 { t.Errorf("expected GuardrailMaxTurns 5, got %d", r.GuardrailMaxTurns) }
	if r.GuardrailMaxSessionActions != 99 { t.Errorf("expected GuardrailMaxSessionActions 99, got %d", r.GuardrailMaxSessionActions) }
}

func TestConfigLimitsOverrideEngineDefaults(t *testing.T) {
	outputDir := t.TempDir()
	engine := NewEngine(&manyFilesRunner{fileCount: 10}, quietOpts(EngineOptions{Workers: 1, OutputDir: outputDir, SkipReview: true, MaxFiles: 100}))
	prompts := []*prompt.Prompt{{ID: "override-test", Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "go", "category": "auth"}}}
	configs := []config.ToolConfig{{Name: "restrictive", Generator: &config.GeneratorConfig{Model: "gpt-4"}, Limits: &config.SessionLimits{MaxFiles: 3}}}
	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil { t.Fatalf("unexpected error: %v", err) }
	r := summary.Results[0]
	if r.Success { t.Error("expected config-level MaxFiles=3 to fail") }
	if !strings.Contains(r.GuardrailAbortReason, "file count") { t.Errorf("expected file count guardrail, got %q", r.GuardrailAbortReason) }
}

// Integration-style: full stub eval lifecycle — verifies reports are generated and result is consistent.
func TestStubEvalLifecycle(t *testing.T) {
	outputDir := t.TempDir()
	engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{
		Workers:    1,
		OutputDir:  outputDir,
		SkipReview: true,
	}))

	prompts := []*prompt.Prompt{
		{ID: "lifecycle-test", Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "go", "category": "crud"}},
	}
	configs := []config.ToolConfig{
		{Name: "baseline", Generator: &config.GeneratorConfig{Model: "gpt-4"}},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify summary fields
	if summary.RunID == "" {
		t.Error("expected non-empty RunID")
	}
	if summary.TotalEvals != 1 {
		t.Errorf("expected 1 eval, got %d", summary.TotalEvals)
	}
	if summary.TotalPrompts != 1 {
		t.Errorf("expected 1 prompt, got %d", summary.TotalPrompts)
	}
	if summary.TotalConfigs != 1 {
		t.Errorf("expected 1 config, got %d", summary.TotalConfigs)
	}

	// Verify result has correct identifiers
	if len(summary.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(summary.Results))
	}
	r := summary.Results[0]
	if r.PromptID != "lifecycle-test" {
		t.Errorf("expected PromptID 'lifecycle-test', got %q", r.PromptID)
	}
	if r.ConfigName != "baseline" {
		t.Errorf("expected ConfigName 'baseline', got %q", r.ConfigName)
	}
	if r.Timestamp == "" {
		t.Error("expected non-empty Timestamp")
	}
	if r.Duration <= 0 {
		t.Errorf("expected positive Duration, got %f", r.Duration)
	}
	if !r.IsStub {
		t.Error("expected IsStub to be true")
	}

	// Verify guardrail limits are recorded
	if r.GuardrailMaxTurns != 25 {
		t.Errorf("expected GuardrailMaxTurns 25, got %d", r.GuardrailMaxTurns)
	}
	if r.GuardrailMaxFiles != 50 {
		t.Errorf("expected GuardrailMaxFiles 50, got %d", r.GuardrailMaxFiles)
	}
	if r.GuardrailMaxSessionActions != 50 {
		t.Errorf("expected GuardrailMaxSessionActions 50, got %d", r.GuardrailMaxSessionActions)
	}

	// Verify report files exist on disk
	reportDir := filepath.Join(outputDir, summary.RunID)
	if _, err := os.Stat(reportDir); os.IsNotExist(err) {
		t.Errorf("expected report directory %s to exist", reportDir)
	}
}

// Integration-style: verify multi-prompt multi-config fan-out
func TestMultiPromptMultiConfigFanOut(t *testing.T) {
	outputDir := t.TempDir()
	engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{
		Workers:    2,
		OutputDir:  outputDir,
		SkipReview: true,
	}))

	prompts := []*prompt.Prompt{
		{ID: "p1", Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "go", "category": "crud"}},
		{ID: "p2", Properties: map[string]string{"service": "keyvault", "plane": "data-plane", "language": "python", "category": "auth"}},
		{ID: "p3", Properties: map[string]string{"service": "cosmos-db", "plane": "data-plane", "language": "java", "category": "query"}},
	}
	configs := []config.ToolConfig{
		{Name: "config-a", Generator: &config.GeneratorConfig{Model: "gpt-4"}},
		{Name: "config-b", Generator: &config.GeneratorConfig{Model: "claude-sonnet-4.5"}},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if summary.TotalEvals != 6 {
		t.Errorf("expected 6 evaluations (3 prompts × 2 configs), got %d", summary.TotalEvals)
	}
	if summary.TotalPrompts != 3 {
		t.Errorf("expected 3 prompts, got %d", summary.TotalPrompts)
	}
	if summary.TotalConfigs != 2 {
		t.Errorf("expected 2 configs, got %d", summary.TotalConfigs)
	}
	if len(summary.Results) != 6 {
		t.Errorf("expected 6 results, got %d", len(summary.Results))
	}

	// Verify all prompt/config combinations are represented
	seen := make(map[string]bool)
	for _, r := range summary.Results {
		key := r.PromptID + "/" + r.ConfigName
		if seen[key] {
			t.Errorf("duplicate result for %s", key)
		}
		seen[key] = true
	}
	for _, p := range prompts {
		for _, c := range configs {
			key := p.ID + "/" + c.Name
			if !seen[key] {
				t.Errorf("missing result for %s", key)
			}
		}
	}
}

// Integration-style: per-phase duration tracking
func TestPhaseDurationTracking(t *testing.T) {
	outputDir := t.TempDir()
	engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{
		Workers:    1,
		OutputDir:  outputDir,
		SkipReview: true,
	}))

	prompts := []*prompt.Prompt{
		{ID: "timing-test", Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "go", "category": "crud"}},
	}
	configs := []config.ToolConfig{
		{Name: "test", Generator: &config.GeneratorConfig{Model: "gpt-4"}},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(summary.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(summary.Results))
	}
	r := summary.Results[0]
	if r.GenerationDuration <= 0 {
		t.Errorf("expected positive GenerationDuration, got %f", r.GenerationDuration)
	}
	if r.Duration <= 0 {
		t.Errorf("expected positive overall Duration, got %f", r.Duration)
	}
}

func TestLargeRunAutoConfirmBypass(t *testing.T) {
	// With AutoConfirm=true, a run of >10 evals should proceed without blocking on stdin.
	outputDir := t.TempDir()
	engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{
		Workers:          1,
		OutputDir:        outputDir,
		SkipReview:       true,
		ConfirmLargeRuns: true,
		AutoConfirm:      true,
	}))

	// Create 12 prompt×config combinations to exceed the 10-eval threshold.
	var prompts []*prompt.Prompt
	for i := 0; i < 12; i++ {
		prompts = append(prompts, &prompt.Prompt{
			ID:       fmt.Sprintf("auto-confirm-%d", i),

			Properties: map[string]string{

				"service":  "storage",

				"plane":    "data-plane",

				"language": "go",

				"category": "crud",

			},
		})
	}
	configs := []config.ToolConfig{
		{Name: "test", Generator: &config.GeneratorConfig{Model: "gpt-4"}},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("expected no error with AutoConfirm, got: %v", err)
	}
	if len(summary.Results) != 12 {
		t.Errorf("expected 12 results, got %d", len(summary.Results))
	}
}

func TestLargeRunConfirmAbort(t *testing.T) {
	// With ConfirmLargeRuns=true and stdin providing "n", the run should abort.
	outputDir := t.TempDir()
	engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{
		Workers:          1,
		OutputDir:        outputDir,
		SkipReview:       true,
		ConfirmLargeRuns: true,
		AutoConfirm:      false,
	}))

	var prompts []*prompt.Prompt
	for i := 0; i < 12; i++ {
		prompts = append(prompts, &prompt.Prompt{
			ID:       fmt.Sprintf("abort-%d", i),

			Properties: map[string]string{

				"service":  "storage",

				"plane":    "data-plane",

				"language": "go",

				"category": "crud",

			},
		})
	}

	// Redirect stdin to provide "n"
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r
	_, _ = w.Write([]byte("n\n"))
	w.Close()
	defer func() { os.Stdin = oldStdin }()

	_, err := engine.Run(context.Background(), prompts, []config.ToolConfig{{Name: "test", Generator: &config.GeneratorConfig{Model: "gpt-4"}}})
	if err == nil {
		t.Fatal("expected error for aborted run")
	}
	if !strings.Contains(err.Error(), "run aborted by user") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestStubRunnerWriteErrorLogsWarning(t *testing.T) {
	// When workDir is an invalid path, the stub evaluator should still
	// succeed (logging a warning) rather than returning an error.
	stub := &StubRunner{}
	p := &prompt.Prompt{ID: "test-prompt", Properties: map[string]string{"language": "go"}}
	cfg := &config.ToolConfig{Name: "test-config", Generator: &config.GeneratorConfig{Model: "gpt-4"}}

	result, err := stub.Run(context.Background(), p, cfg, filepath.Join(t.TempDir(), "nonexistent", "subdir"))
	if err != nil {
		t.Fatalf("stub evaluator should not return an error even when write fails: %v", err)
	}
	if !result.Success {
		t.Error("expected stub to report success even when file write fails")
	}
	if !result.IsStub {
		t.Error("expected IsStub to be true")
	}
}

func TestLargeRunConfirmNoTTY(t *testing.T) {
	// With ConfirmLargeRuns=true and stdin closed (simulating piped/no-TTY input),
	// the run should abort with an "no interactive input" error.
	outputDir := t.TempDir()
	engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{
		Workers:          1,
		OutputDir:        outputDir,
		SkipReview:       true,
		ConfirmLargeRuns: true,
		AutoConfirm:      false,
	}))

	var prompts []*prompt.Prompt
	for i := 0; i < 12; i++ {
		prompts = append(prompts, &prompt.Prompt{
			ID:         fmt.Sprintf("no-tty-%d", i),
			Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "go", "category": "crud"},
		})
	}

	// Redirect stdin to a closed pipe (EOF immediately).
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()
	w.Close()
	os.Stdin = r
	defer func() { os.Stdin = oldStdin }()

	_, err := engine.Run(context.Background(), prompts, []config.ToolConfig{
		{Name: "test", Generator: &config.GeneratorConfig{Model: "gpt-4"}},
	})
	if err == nil {
		t.Fatal("expected error when stdin is closed (no TTY)")
	}
	if !strings.Contains(err.Error(), "no interactive input available") {
		t.Errorf("expected 'no interactive input available' error, got: %v", err)
	}
}

// capturingReviewer records the evaluation criteria passed to Review.
type capturingReviewer struct {
	capturedCriteria string
}

func (c *capturingReviewer) Review(_ context.Context, _ string, _ string, _ string, evaluationCriteria string) (*review.ReviewResult, error) {
	c.capturedCriteria = evaluationCriteria
	return &review.ReviewResult{
		OverallScore: 5,
		MaxScore:     5,
	}, nil
}

func TestCriteriaMergedIntoReview(t *testing.T) {
	// Create criteria directory with a language-matched grader config
	criteriaDir := t.TempDir()
	os.MkdirAll(filepath.Join(criteriaDir, "language"), 0755)
	os.WriteFile(filepath.Join(criteriaDir, "language", "go.yaml"), []byte(`
when:
  language: go
graders:
  - name: Uses DefaultAzureCredential
    weight: 1.0
    prompt: Must use azidentity.DefaultAzureCredential
`), 0644)

	reviewer := &capturingReviewer{}
	reviewerFactory := func(cfg *config.ToolConfig) (review.Reviewer, *review.PanelReviewer, error) {
		return reviewer, nil, nil
	}
	engine := NewEngineWithReviewerFactory(&StubRunner{}, reviewerFactory, quietOpts(EngineOptions{
		Workers:     1,
		OutputDir:   t.TempDir(),
		CriteriaDir: criteriaDir,
	}))

	prompts := []*prompt.Prompt{
		{
			ID: "criteria-test", Properties: map[string]string{"service": "storage", "plane": "data-plane",
				"language": "go", "category": "crud"},
			EvaluationCriteria: "- Must handle errors properly",
		},
	}
	configs := []config.ToolConfig{{Name: "test", Generator: &config.GeneratorConfig{Model: "gpt-4"}}}

	_, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(reviewer.capturedCriteria, "DefaultAzureCredential") {
		t.Errorf("expected grader criteria in review, got: %s", reviewer.capturedCriteria)
	}
	if !strings.Contains(reviewer.capturedCriteria, "handle errors properly") {
		t.Errorf("expected prompt criteria in review, got: %s", reviewer.capturedCriteria)
	}
}

func TestCriteriaDirNotExist(t *testing.T) {
	// Non-existent criteria dir should not cause an error
	engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{
		Workers:     1,
		OutputDir:   t.TempDir(),
		CriteriaDir: "/nonexistent/path",
		SkipReview:  true,
	}))

	prompts := []*prompt.Prompt{
		{ID: "dir-test", Properties: map[string]string{"service": "storage", "language": "go", "plane": "data-plane", "category": "crud"}},
	}
	configs := []config.ToolConfig{{Name: "test", Generator: &config.GeneratorConfig{Model: "gpt-4"}}}

	_, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("non-existent criteria dir should not fail: %v", err)
	}
}

func TestCriteriaDirEmpty(t *testing.T) {
	// Empty criteria dir should work fine — no criteria matched
	reviewer := &capturingReviewer{}
	reviewerFactory := func(cfg *config.ToolConfig) (review.Reviewer, *review.PanelReviewer, error) {
		return reviewer, nil, nil
	}
	engine := NewEngineWithReviewerFactory(&StubRunner{}, reviewerFactory, quietOpts(EngineOptions{
		Workers:     1,
		OutputDir:   t.TempDir(),
		CriteriaDir: t.TempDir(), // empty dir
	}))

	prompts := []*prompt.Prompt{
		{
			ID: "empty-criteria", Properties: map[string]string{"service": "storage", "language": "go",
				"plane": "data-plane", "category": "crud"},
			EvaluationCriteria: "- Prompt specific criterion",
		},
	}
	configs := []config.ToolConfig{{Name: "test", Generator: &config.GeneratorConfig{Model: "gpt-4"}}}

	_, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should fall back to prompt-only criteria
	if !strings.Contains(reviewer.capturedCriteria, "Prompt specific criterion") {
		t.Errorf("expected prompt criteria as fallback, got: %s", reviewer.capturedCriteria)
	}
}

// TestCancelledContextNoGoroutineLeak verifies that goroutines spawned by
// Engine.Run terminate promptly when the parent context is already cancelled.
// This catches semaphore-acquisition leaks (#129).
func TestCancelledContextNoGoroutineLeak(t *testing.T) {
	outputDir := t.TempDir()
	// Use 1 worker but many tasks — excess goroutines would block on
	// semaphore acquisition if context cancellation is not respected.
	engine := NewEngine(&slowRunner{}, quietOpts(EngineOptions{
		Workers:    1,
		OutputDir:  outputDir,
		SkipReview: true,
	}))

	prompts := []*prompt.Prompt{
		{ID: "leak-1", Properties: map[string]string{"service": "s", "plane": "data-plane", "language": "go", "category": "c"}},
		{ID: "leak-2", Properties: map[string]string{"service": "s", "plane": "data-plane", "language": "go", "category": "c"}},
		{ID: "leak-3", Properties: map[string]string{"service": "s", "plane": "data-plane", "language": "go", "category": "c"}},
		{ID: "leak-4", Properties: map[string]string{"service": "s", "plane": "data-plane", "language": "go", "category": "c"}},
	}
	configs := []config.ToolConfig{
		{Name: "cfg", Generator: &config.GeneratorConfig{Model: "gpt-4"}},
	}

	before := runtime.NumGoroutine()

	// Context is already cancelled — all goroutines should bail immediately.
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := engine.Run(ctx, prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Allow goroutines a brief window to wind down.
	deadline := time.After(2 * time.Second)
	for {
		after := runtime.NumGoroutine()
		// Tolerate a small delta for unrelated runtime goroutines.
		if after <= before+2 {
			break
		}
		select {
		case <-deadline:
			t.Fatalf("goroutine leak: before=%d, after=%d (delta %d)", before, after, after-before)
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func TestStrictCleanupOptionWired(t *testing.T) {
	// Verify StrictCleanup option flows through to the engine.
	engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{
		Workers:       1,
		OutputDir:     t.TempDir(),
		SkipReview:    true,
		StrictCleanup: true,
	}))

	if !engine.opts.StrictCleanup {
		t.Error("expected StrictCleanup to be true")
	}

	// Verify it defaults to false
	engine2 := NewEngine(&StubRunner{}, quietOpts(EngineOptions{
		Workers:  1,
		OutputDir: t.TempDir(),
	}))

	if engine2.opts.StrictCleanup {
		t.Error("expected StrictCleanup to default to false")
	}
}

// ---------------------------------------------------------------------------
// Tiered limits tests (#347)
// ---------------------------------------------------------------------------

func TestReviewerLimitsDefaultToHalfGenerator(t *testing.T) {
	engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{
		Workers:           1,
		OutputDir:         t.TempDir(),
		MaxTurns:          20,
		MaxSessionActions: 40,
	}))

	if engine.opts.ReviewerMaxTurns != 10 {
		t.Errorf("ReviewerMaxTurns = %d, want 10 (half of 20)", engine.opts.ReviewerMaxTurns)
	}
	if engine.opts.ReviewerMaxActions != 20 {
		t.Errorf("ReviewerMaxActions = %d, want 20 (half of 40)", engine.opts.ReviewerMaxActions)
	}
}

func TestReviewerLimitsMinimumFloor(t *testing.T) {
	engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{
		Workers:           1,
		OutputDir:         t.TempDir(),
		MaxTurns:          6,
		MaxSessionActions: 10,
	}))

	if engine.opts.ReviewerMaxTurns < 5 {
		t.Errorf("ReviewerMaxTurns = %d, want >= 5 (minimum floor)", engine.opts.ReviewerMaxTurns)
	}
	if engine.opts.ReviewerMaxActions < 10 {
		t.Errorf("ReviewerMaxActions = %d, want >= 10 (minimum floor)", engine.opts.ReviewerMaxActions)
	}
}

func TestReviewerLimitsExplicitOverride(t *testing.T) {
	engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{
		Workers:            1,
		OutputDir:          t.TempDir(),
		MaxTurns:           20,
		MaxSessionActions:  40,
		ReviewerMaxTurns:   8,
		ReviewerMaxActions: 15,
	}))

	if engine.opts.ReviewerMaxTurns != 8 {
		t.Errorf("ReviewerMaxTurns = %d, want 8 (explicit override)", engine.opts.ReviewerMaxTurns)
	}
	if engine.opts.ReviewerMaxActions != 15 {
		t.Errorf("ReviewerMaxActions = %d, want 15 (explicit override)", engine.opts.ReviewerMaxActions)
	}
}

// ---------------------------------------------------------------------------
// Tool availability tracking tests (#348)
// ---------------------------------------------------------------------------

func TestBuildToolAvailability(t *testing.T) {
	env := &report.EnvironmentInfo{
		AvailableTools: []string{"bash", "create", "edit"},
		MCPServers:     []string{"azure-mcp"},
		SkillsLoaded:   []string{"sdk-skill"},
		SkillsInvoked:  []string{"sdk-skill"},
	}
	events := []report.SessionEventRecord{
		{Type: "tool.execution_complete", ToolName: "bash"},
		{Type: "tool.execution_complete", ToolName: "create"},
		{Type: "skill.invoked", SkillName: "sdk-skill"},
		{Type: "external_tool.completed", MCPServerName: "azure-mcp", ToolName: "mcp_tool"},
	}

	entries := buildToolAvailability(env, events)

	if len(entries) != 5 {
		t.Fatalf("expected 5 entries (3 tools + 1 mcp + 1 skill), got %d", len(entries))
	}

	// Verify bash: available=true, used=true
	found := false
	for _, e := range entries {
		if e.Name == "bash" {
			found = true
			if !e.Available || !e.Used {
				t.Errorf("bash: available=%v used=%v, want true/true", e.Available, e.Used)
			}
			if e.Type != "tool" {
				t.Errorf("bash type = %q, want tool", e.Type)
			}
		}
	}
	if !found {
		t.Error("bash entry not found")
	}

	// Verify edit: available=true, used=false
	for _, e := range entries {
		if e.Name == "edit" {
			if !e.Available || e.Used {
				t.Errorf("edit: available=%v used=%v, want true/false", e.Available, e.Used)
			}
		}
	}

	// Verify azure-mcp: available=true, used=true
	for _, e := range entries {
		if e.Name == "azure-mcp" {
			if !e.Available || !e.Used {
				t.Errorf("azure-mcp: available=%v used=%v, want true/true", e.Available, e.Used)
			}
			if e.Type != "mcp" {
				t.Errorf("azure-mcp type = %q, want mcp", e.Type)
			}
		}
	}

	// Verify skill: available=true, used=true
	for _, e := range entries {
		if e.Name == "sdk-skill" {
			if !e.Available || !e.Used {
				t.Errorf("sdk-skill: available=%v used=%v, want true/true", e.Available, e.Used)
			}
			if e.Type != "skill" {
				t.Errorf("sdk-skill type = %q, want skill", e.Type)
			}
		}
	}
}

func TestBuildToolAvailabilityNilEnv(t *testing.T) {
	entries := buildToolAvailability(nil, nil)
	if entries != nil {
		t.Errorf("expected nil for nil env, got %v", entries)
	}
}

func TestBuildToolAvailabilityNoEvents(t *testing.T) {
	env := &report.EnvironmentInfo{
		AvailableTools: []string{"bash"},
		MCPServers:     []string{"mcp-1"},
	}
	entries := buildToolAvailability(env, nil)
	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(entries))
	}
	for _, e := range entries {
		if e.Used {
			t.Errorf("%s should not be used (no events)", e.Name)
		}
	}
}

// ---------------------------------------------------------------------------
// Unified grading pipeline tests (WI-023)
// ---------------------------------------------------------------------------

// TestReviewResultsAppendedNotOverwritten verifies the fix for the results
// overwrite bug: typed grader results must coexist with review results in
// the same GraderResults slice. Cut over to the unified Bundle path (#625).
func TestReviewResultsAppendedNotOverwritten(t *testing.T) {
	// Set up a file grader that will pass (stub_output.txt exists).
	gradersDir := t.TempDir()
	graderYAML := `graders:
  - type: file
    name: "stub_exists"
    details:
      path: "stub_output.txt"
    weight: 1.0
`
	os.WriteFile(filepath.Join(gradersDir, "test.yaml"), []byte(graderYAML), 0644)

	reviewer := &review.StubReviewer{}
	reviewerFactory := func(cfg *config.ToolConfig) (review.Reviewer, *review.PanelReviewer, error) {
		return reviewer, nil, nil
	}

	engine := NewEngineWithReviewerFactory(&StubRunner{}, reviewerFactory, quietOpts(EngineOptions{
		Workers:     1,
		OutputDir:   t.TempDir(),
		CriteriaDir: gradersDir,
	}))

	prompts := []*prompt.Prompt{
		{
			ID:         "overwrite-test",
			Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "python", "category": "crud"},
			PromptText: "Create a test file",
		},
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

	// The bug previously caused review results to overwrite grader results.
	// After the fix, both must be present.
	if len(r.GraderResults) < 2 {
		t.Fatalf("expected at least 2 grader results (file grader + review), got %d", len(r.GraderResults))
	}

	hasFile := false
	hasReview := false
	for _, gr := range r.GraderResults {
		if gr.GraderType == "file" {
			hasFile = true
		}
		// v3 schema collapses panel-member expansion into a single
		// "ai_review" row carrying GraderType "prompt_review".
		if gr.GraderType == "prompt_review" || gr.GraderType == "review" {
			hasReview = true
		}
	}
	if !hasFile {
		t.Error("expected file grader result to be present")
	}
	if !hasReview {
		t.Error("expected review grader result to be present")
	}

	// Review backward-compat fields should still be populated.
	if r.Review == nil {
		t.Error("expected Review field to be populated for backward compat")
	}
}

// TestUnifiedGraderSuccessIncludesReview verifies that evalReport.Success
// is determined from ALL grader results, not just the review.
func TestUnifiedGraderSuccessIncludesReview(t *testing.T) {
	// File grader passes (stub_output.txt exists).
	// Review passes (StubReviewer always passes).
	// Overall should pass.
	gradersDir := t.TempDir()
	graderYAML := `graders:
  - type: file
    name: "stub_check"
    details:
      path: "stub_output.txt"
    weight: 1.0
`
	os.WriteFile(filepath.Join(gradersDir, "test.yaml"), []byte(graderYAML), 0644)

	reviewerFactory := func(cfg *config.ToolConfig) (review.Reviewer, *review.PanelReviewer, error) {
		return &review.StubReviewer{}, nil, nil
	}

	engine := NewEngineWithReviewerFactory(&StubRunner{}, reviewerFactory, quietOpts(EngineOptions{
		Workers:     1,
		OutputDir:   t.TempDir(),
		CriteriaDir: gradersDir,
	}))

	prompts := []*prompt.Prompt{
		{
			ID:         "unified-success-test",
			Properties: map[string]string{"service": "storage", "plane": "data-plane", "language": "python", "category": "crud"},
			PromptText: "test",
		},
	}
	configs := []config.ToolConfig{
		{Name: "cfg", Generator: &config.GeneratorConfig{Model: "gpt-4"}},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r := summary.Results[0]
	if !r.Success {
		t.Error("expected unified success when all graders pass")
	}
}

