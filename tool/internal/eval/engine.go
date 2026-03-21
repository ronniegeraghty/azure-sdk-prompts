package eval

import (
"context"
"fmt"
"log"
"sync"
"time"

"github.com/ronniegeraghty/azure-sdk-prompts/tool/internal/build"
"github.com/ronniegeraghty/azure-sdk-prompts/tool/internal/config"
"github.com/ronniegeraghty/azure-sdk-prompts/tool/internal/prompt"
"github.com/ronniegeraghty/azure-sdk-prompts/tool/internal/report"
"github.com/ronniegeraghty/azure-sdk-prompts/tool/internal/review"
)

// EvalResult holds the raw output from a Copilot evaluation.
type EvalResult struct {
GeneratedFiles []string
EventCount     int
ToolCalls      []string
SessionEvents  []report.SessionEventRecord
Success        bool
Error          string
ErrorDetails   string
IsStub         bool
}

// CopilotEvaluator defines the interface for running evaluations.
type CopilotEvaluator interface {
Evaluate(ctx context.Context, prompt *prompt.Prompt, config *config.ToolConfig, workDir string) (*EvalResult, error)
}

// StubEvaluator returns placeholder results for testing.
type StubEvaluator struct{}

// Evaluate returns a stub result.
func (s *StubEvaluator) Evaluate(ctx context.Context, p *prompt.Prompt, cfg *config.ToolConfig, workDir string) (*EvalResult, error) {
return &EvalResult{
GeneratedFiles: []string{"stub_output.txt"},
EventCount:     0,
ToolCalls:      []string{},
SessionEvents:  nil,
Success:        true,
Error:          "",
IsStub:         true,
}, nil
}

// EngineOptions configures the evaluation engine.
type EngineOptions struct {
Workers     int
Timeout     time.Duration
OutputDir   string
SkipTests   bool
SkipReview  bool
VerifyBuild bool
Debug       bool
DryRun      bool
}

// Verifier evaluates generated code against prompt requirements.
type Verifier interface {
Verify(ctx context.Context, originalPrompt string, workDir string, expectedCoverage string) (*report.VerifyResult, error)
}

// StubVerifier returns a placeholder pass result.
type StubVerifier struct{}

// Verify returns a stub verification pass.
func (s *StubVerifier) Verify(_ context.Context, _ string, _ string, _ string) (*report.VerifyResult, error) {
return &report.VerifyResult{
Pass:      true,
Reasoning: "Verification skipped (stub mode)",
Summary:   "Stub mode — no Copilot verification performed",
}, nil
}

// Engine orchestrates evaluation runs.
type Engine struct {
evaluator CopilotEvaluator
reviewer  review.Reviewer
verifier  Verifier
opts      EngineOptions
}

// NewEngine creates a new Engine with the given evaluator and options.
func NewEngine(evaluator CopilotEvaluator, opts EngineOptions) *Engine {
return NewEngineWithReviewer(evaluator, nil, nil, opts)
}

// NewEngineWithReviewer creates a new Engine with an evaluator, verifier, and reviewer.
func NewEngineWithReviewer(evaluator CopilotEvaluator, verifier Verifier, reviewer review.Reviewer, opts EngineOptions) *Engine {
if opts.Workers <= 0 {
opts.Workers = 4
}
if opts.Timeout <= 0 {
opts.Timeout = 5 * time.Minute
}
if opts.OutputDir == "" {
opts.OutputDir = "./reports"
}
return &Engine{
evaluator: evaluator,
reviewer:  reviewer,
verifier:  verifier,
opts:      opts,
}
}

// EvalTask represents a single prompt+config evaluation to run.
type EvalTask struct {
Prompt *prompt.Prompt
Config config.ToolConfig
}

// Run executes evaluations for the given prompts crossed with configs.
func (e *Engine) Run(ctx context.Context, prompts []*prompt.Prompt, configs []config.ToolConfig) (*report.RunSummary, error) {
// Build task list (cross product)
var tasks []EvalTask
for _, p := range prompts {
for _, c := range configs {
tasks = append(tasks, EvalTask{Prompt: p, Config: c})
}
}

if e.opts.DryRun {
return e.dryRun(tasks)
}

runID := time.Now().Format("20060102-150405")
summary := &report.RunSummary{
RunID:        runID,
Timestamp:    time.Now().UTC().Format(time.RFC3339),
TotalPrompts: len(prompts),
TotalConfigs: len(configs),
TotalEvals:   len(tasks),
}

start := time.Now()

sem := make(chan struct{}, e.opts.Workers)
var mu sync.Mutex
var wg sync.WaitGroup

for _, task := range tasks {
wg.Add(1)
go func(t EvalTask) {
defer wg.Done()

sem <- struct{}{}
defer func() { <-sem }()

evalReport := e.runSingleEval(ctx, t, runID)

mu.Lock()
defer mu.Unlock()

summary.Results = append(summary.Results, evalReport)

if evalReport.Success {
summary.Passed++
} else if evalReport.Error != "" {
summary.Errors++
} else {
summary.Failed++
}
}(task)
}

wg.Wait()

summary.Duration = time.Since(start).Seconds()

// Write JSON summary
if _, err := report.WriteSummary(summary, e.opts.OutputDir); err != nil && e.opts.Debug {
log.Printf("failed to write run summary: %v", err)
}

// Write HTML summary
if _, err := report.WriteSummaryHTML(summary, e.opts.OutputDir); err != nil && e.opts.Debug {
log.Printf("failed to write HTML summary: %v", err)
}

return summary, nil
}

func (e *Engine) runSingleEval(ctx context.Context, task EvalTask, runID string) *report.EvalReport {
evalCtx, cancel := context.WithTimeout(ctx, e.opts.Timeout)
defer cancel()

start := time.Now()

evalReport := &report.EvalReport{
PromptID:   task.Prompt.ID,
ConfigName: task.Config.Name,
Timestamp:  time.Now().UTC().Format(time.RFC3339),
PromptMeta: map[string]any{
"service":  task.Prompt.Service,
"plane":    task.Prompt.Plane,
"language": task.Prompt.Language,
"category": task.Prompt.Category,
},
ConfigUsed: map[string]any{
"name":  task.Config.Name,
"model": task.Config.Model,
},
}

if e.opts.Debug {
log.Printf("[DEBUG] Starting Copilot session for %s with config %s...", task.Prompt.ID, task.Config.Name)
}

// Setup workspace
ws, err := NewWorkspace(e.opts.OutputDir, task.Prompt.ID, task.Config.Name)
if err != nil {
evalReport.Error = fmt.Sprintf("workspace setup failed: %v", err)
evalReport.ErrorDetails = err.Error()
evalReport.Duration = time.Since(start).Seconds()
return evalReport
}

// Run evaluation
result, err := e.evaluator.Evaluate(evalCtx, task.Prompt, &task.Config, ws.Dir)
if err != nil {
evalReport.Error = fmt.Sprintf("evaluation failed: %v", err)
evalReport.ErrorDetails = err.Error()
evalReport.Duration = time.Since(start).Seconds()
// Capture whatever session events were collected before failure
if result != nil {
evalReport.SessionEvents = result.SessionEvents
evalReport.EventCount = result.EventCount
evalReport.ToolCalls = result.ToolCalls
evalReport.IsStub = result.IsStub
}
return evalReport
}

evalReport.GeneratedFiles = result.GeneratedFiles
evalReport.EventCount = result.EventCount
evalReport.ToolCalls = result.ToolCalls
evalReport.SessionEvents = result.SessionEvents
evalReport.IsStub = result.IsStub
evalReport.Success = result.Success

if e.opts.Debug {
log.Printf("[DEBUG] Session complete: %d tool calls, %d files generated, %s",
len(result.ToolCalls), len(result.GeneratedFiles), time.Since(start).Truncate(time.Millisecond))
}

// Copilot-based verification (default, unless stub mode has its own stub verifier)
if e.verifier != nil {
if e.opts.Debug {
log.Printf("[DEBUG] Starting verification session...")
}
verifyResult, err := e.verifier.Verify(evalCtx, task.Prompt.PromptText, ws.Dir, task.Prompt.ExpectedCoverage)
if err != nil {
if e.opts.Debug {
log.Printf("[DEBUG] ERROR: verification failed: %v", err)
}
} else {
evalReport.Verification = verifyResult
evalReport.Success = verifyResult.Pass
if e.opts.Debug {
passStr := "FAIL"
if verifyResult.Pass {
passStr = "PASS"
}
log.Printf("[DEBUG] Verification: %s — %s", passStr, verifyResult.Summary)
}
}
}

// Optional build verification (--verify-build flag)
if e.opts.VerifyBuild {
buildResult, err := build.Verify(evalCtx, task.Prompt.Language, ws.Dir)
if err != nil {
if e.opts.Debug {
log.Printf("[DEBUG] ERROR: build verification failed: %v", err)
}
} else {
evalReport.Build = buildResult
if !buildResult.Success {
evalReport.Success = false
}
}
}

// Code review (unless skipped)
if !e.opts.SkipReview && e.reviewer != nil {
if e.opts.Debug {
log.Printf("[DEBUG] Starting review session...")
}
referenceDir := ""
if task.Prompt.ReferenceAnswer != "" {
referenceDir = task.Prompt.ReferenceAnswer
}
reviewResult, err := e.reviewer.Review(evalCtx, task.Prompt.PromptText, ws.Dir, referenceDir)
if err != nil {
if e.opts.Debug {
log.Printf("[DEBUG] ERROR: code review failed for %s/%s: %v", task.Prompt.ID, task.Config.Name, err)
}
} else {
evalReport.Review = reviewResult
if e.opts.Debug {
log.Printf("[DEBUG] Review score: %d/10", reviewResult.OverallScore)
}
}
}

evalReport.Duration = time.Since(start).Seconds()

// Write JSON report
reportPath, err := report.WriteReport(evalReport, e.opts.OutputDir, runID, task.Prompt)
if err != nil {
if e.opts.Debug {
log.Printf("[DEBUG] ERROR: failed to write report: %v", err)
}
} else if e.opts.Debug {
log.Printf("[DEBUG] report written to %s", reportPath)
}

// Write HTML report
if _, err := report.WriteHTMLReport(evalReport, e.opts.OutputDir, runID,
task.Prompt.Service, task.Prompt.Plane, task.Prompt.Language, task.Prompt.Category); err != nil {
if e.opts.Debug {
log.Printf("[DEBUG] ERROR: failed to write HTML report: %v", err)
}
}

return evalReport
}

func (e *Engine) dryRun(tasks []EvalTask) (*report.RunSummary, error) {
summary := &report.RunSummary{
RunID:        "dry-run",
Timestamp:    time.Now().UTC().Format(time.RFC3339),
TotalPrompts: 0,
TotalConfigs: 0,
TotalEvals:   len(tasks),
}

promptIDs := make(map[string]bool)
configNames := make(map[string]bool)

for _, t := range tasks {
promptIDs[t.Prompt.ID] = true
configNames[t.Config.Name] = true
}

summary.TotalPrompts = len(promptIDs)
summary.TotalConfigs = len(configNames)

return summary, nil
}
