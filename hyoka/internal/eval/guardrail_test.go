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
	"github.com/ronniegeraghty/hyoka/internal/report"
)

// turnHeavyEvaluator returns many assistant.turn_end events to trigger the turn guardrail.
type turnHeavyEvaluator struct {
	turnCount int
}

func (e *turnHeavyEvaluator) Evaluate(ctx context.Context, _ *prompt.Prompt, _ *config.ToolConfig, workDir string) (*EvalResult, error) {
	events := make([]report.SessionEventRecord, e.turnCount)
	for i := 0; i < e.turnCount; i++ {
		events[i] = report.SessionEventRecord{Type: "assistant.turn_end"}
	}
	// Write one file so we have non-zero output
	os.WriteFile(filepath.Join(workDir, "main.py"), []byte("print('hello')"), 0644)
	return &EvalResult{
		GeneratedFiles: []string{"main.py"},
		EventCount:     e.turnCount,
		SessionEvents:  events,
		Success:        true,
	}, nil
}

// fileHeavyEvaluator creates many files to trigger the file count guardrail.
type fileHeavyEvaluator struct {
	fileCount int
}

func (e *fileHeavyEvaluator) Evaluate(ctx context.Context, _ *prompt.Prompt, _ *config.ToolConfig, workDir string) (*EvalResult, error) {
	var files []string
	for i := 0; i < e.fileCount; i++ {
		name := "file_" + strings.Repeat("x", 3) + "_" + string(rune('a'+i%26)) + ".py"
		if i >= 26 {
			name = "sub/file_" + string(rune('a'+i%26)) + string(rune('0'+i/26)) + ".py"
			os.MkdirAll(filepath.Join(workDir, "sub"), 0755)
		}
		os.WriteFile(filepath.Join(workDir, name), []byte("# generated"), 0644)
		files = append(files, name)
	}
	return &EvalResult{
		GeneratedFiles: files,
		EventCount:     1,
		SessionEvents:  []report.SessionEventRecord{{Type: "assistant.turn_end"}},
		Success:        true,
	}, nil
}

// sizeHeavyEvaluator creates large files to trigger the output size guardrail.
type sizeHeavyEvaluator struct {
	sizeBytes int64
}

func (e *sizeHeavyEvaluator) Evaluate(ctx context.Context, _ *prompt.Prompt, _ *config.ToolConfig, workDir string) (*EvalResult, error) {
	data := make([]byte, e.sizeBytes)
	for i := range data {
		data[i] = 'x'
	}
	os.WriteFile(filepath.Join(workDir, "bigfile.bin"), data, 0644)
	return &EvalResult{
		GeneratedFiles: []string{"bigfile.bin"},
		EventCount:     1,
		SessionEvents:  []report.SessionEventRecord{{Type: "assistant.turn_end"}},
		Success:        true,
	}, nil
}

func TestGuardrail_TurnCapTriggered(t *testing.T) {
	outputDir := t.TempDir()
	maxTurns := 5
	engine := NewEngine(&turnHeavyEvaluator{turnCount: maxTurns + 3}, EngineOptions{
		Workers:         1,
		OutputDir:       outputDir,
		GenerateTimeout: 30 * time.Second,
		MaxTurns:        maxTurns,
	})

	prompts := []*prompt.Prompt{
		{ID: "turn-test", Service: "storage", Plane: "data-plane", Language: "go", Category: "auth"},
	}
	configs := []config.ToolConfig{
		{Name: "test-config", Model: "gpt-4"},
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
		t.Error("expected Success=false when turn cap is exceeded")
	}
	if r.GuardrailAbortReason == "" {
		t.Error("expected GuardrailAbortReason to be set")
	}
	if !strings.Contains(r.GuardrailAbortReason, "turn count") {
		t.Errorf("expected turn count in abort reason, got %q", r.GuardrailAbortReason)
	}
	if !strings.Contains(r.Error, "turn count") {
		t.Errorf("expected turn count in error, got %q", r.Error)
	}
}

func TestGuardrail_TurnCapNotTriggered(t *testing.T) {
	outputDir := t.TempDir()
	maxTurns := 25
	engine := NewEngine(&turnHeavyEvaluator{turnCount: 3}, EngineOptions{
		Workers:         1,
		OutputDir:       outputDir,
		GenerateTimeout: 30 * time.Second,
		MaxTurns:        maxTurns,
	})

	prompts := []*prompt.Prompt{
		{ID: "turn-ok", Service: "storage", Plane: "data-plane", Language: "go", Category: "auth"},
	}
	configs := []config.ToolConfig{
		{Name: "test-config", Model: "gpt-4"},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r := summary.Results[0]
	if !r.Success {
		t.Errorf("expected Success=true when under turn cap, got error: %q", r.Error)
	}
	if r.GuardrailAbortReason != "" {
		t.Errorf("expected empty GuardrailAbortReason, got %q", r.GuardrailAbortReason)
	}
}

func TestGuardrail_FileCapTriggered(t *testing.T) {
	outputDir := t.TempDir()
	maxFiles := 5
	engine := NewEngine(&fileHeavyEvaluator{fileCount: maxFiles + 3}, EngineOptions{
		Workers:         1,
		OutputDir:       outputDir,
		GenerateTimeout: 30 * time.Second,
		MaxFiles:        maxFiles,
	})

	prompts := []*prompt.Prompt{
		{ID: "file-test", Service: "storage", Plane: "data-plane", Language: "python", Category: "crud"},
	}
	configs := []config.ToolConfig{
		{Name: "test-config", Model: "gpt-4"},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r := summary.Results[0]
	if r.Success {
		t.Error("expected Success=false when file cap is exceeded")
	}
	if r.GuardrailAbortReason == "" {
		t.Error("expected GuardrailAbortReason to be set")
	}
	if !strings.Contains(r.GuardrailAbortReason, "file count") {
		t.Errorf("expected 'file count' in abort reason, got %q", r.GuardrailAbortReason)
	}
}

func TestGuardrail_SizeCapTriggered(t *testing.T) {
	outputDir := t.TempDir()
	maxSize := int64(1024) // 1KB
	engine := NewEngine(&sizeHeavyEvaluator{sizeBytes: 2048}, EngineOptions{
		Workers:         1,
		OutputDir:       outputDir,
		GenerateTimeout: 30 * time.Second,
		MaxOutputSize:   maxSize,
	})

	prompts := []*prompt.Prompt{
		{ID: "size-test", Service: "storage", Plane: "data-plane", Language: "python", Category: "crud"},
	}
	configs := []config.ToolConfig{
		{Name: "test-config", Model: "gpt-4"},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r := summary.Results[0]
	if r.Success {
		t.Error("expected Success=false when size cap is exceeded")
	}
	if r.GuardrailAbortReason == "" {
		t.Error("expected GuardrailAbortReason to be set")
	}
	if !strings.Contains(r.GuardrailAbortReason, "output size") {
		t.Errorf("expected 'output size' in abort reason, got %q", r.GuardrailAbortReason)
	}
}

func TestGuardrail_ReportRecordsLimits(t *testing.T) {
	outputDir := t.TempDir()
	engine := NewEngine(&StubEvaluator{}, EngineOptions{
		Workers:         1,
		OutputDir:       outputDir,
		GenerateTimeout: 30 * time.Second,
		MaxTurns:        10,
		MaxFiles:        20,
		MaxOutputSize:   512000,
	})

	prompts := []*prompt.Prompt{
		{ID: "limits-test", Service: "storage", Plane: "data-plane", Language: "go", Category: "auth"},
	}
	configs := []config.ToolConfig{
		{Name: "test-config", Model: "gpt-4"},
	}

	summary, err := engine.Run(context.Background(), prompts, configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r := summary.Results[0]
	if r.GuardrailMaxTurns != 10 {
		t.Errorf("expected GuardrailMaxTurns=10, got %d", r.GuardrailMaxTurns)
	}
	if r.GuardrailMaxFiles != 20 {
		t.Errorf("expected GuardrailMaxFiles=20, got %d", r.GuardrailMaxFiles)
	}
	if r.GuardrailMaxOutputSize != 512000 {
		t.Errorf("expected GuardrailMaxOutputSize=512000, got %d", r.GuardrailMaxOutputSize)
	}
}

func TestNewEngine_DefaultGuardrails(t *testing.T) {
	engine := NewEngine(&StubEvaluator{}, EngineOptions{})

	if engine.opts.MaxTurns != 25 {
		t.Errorf("expected default MaxTurns=25, got %d", engine.opts.MaxTurns)
	}
	if engine.opts.MaxFiles != 50 {
		t.Errorf("expected default MaxFiles=50, got %d", engine.opts.MaxFiles)
	}
	if engine.opts.MaxOutputSize != 1048576 {
		t.Errorf("expected default MaxOutputSize=1MB, got %d", engine.opts.MaxOutputSize)
	}
}

func TestNewEngine_DefaultTimeouts(t *testing.T) {
	engine := NewEngine(&StubEvaluator{}, EngineOptions{})

	if engine.opts.GenerateTimeout != 10*time.Minute {
		t.Errorf("expected default GenerateTimeout=10m, got %v", engine.opts.GenerateTimeout)
	}
	if engine.opts.BuildTimeout != 5*time.Minute {
		t.Errorf("expected default BuildTimeout=5m, got %v", engine.opts.BuildTimeout)
	}
	if engine.opts.ReviewTimeout != 5*time.Minute {
		t.Errorf("expected default ReviewTimeout=5m, got %v", engine.opts.ReviewTimeout)
	}
}

func TestNewEngine_LegacyTimeoutBackwardCompat(t *testing.T) {
	engine := NewEngine(&StubEvaluator{}, EngineOptions{
		Timeout: 120 * time.Second,
	})

	if engine.opts.GenerateTimeout != 120*time.Second {
		t.Errorf("expected GenerateTimeout=120s from legacy Timeout, got %v", engine.opts.GenerateTimeout)
	}
}

func TestNewEngine_GenerateTimeoutOverridesLegacy(t *testing.T) {
	engine := NewEngine(&StubEvaluator{}, EngineOptions{
		Timeout:         120 * time.Second,
		GenerateTimeout: 180 * time.Second,
	})

	if engine.opts.GenerateTimeout != 180*time.Second {
		t.Errorf("expected GenerateTimeout=180s (explicit), got %v", engine.opts.GenerateTimeout)
	}
}

func TestNewEngine_WorkerDefaults(t *testing.T) {
	engine := NewEngine(&StubEvaluator{}, EngineOptions{})

	if engine.opts.Workers <= 0 || engine.opts.Workers > 8 {
		t.Errorf("expected workers between 1-8, got %d", engine.opts.Workers)
	}
	if engine.opts.MaxSessions != engine.opts.Workers*3 {
		t.Errorf("expected MaxSessions=Workers*3 (%d), got %d", engine.opts.Workers*3, engine.opts.MaxSessions)
	}
}
