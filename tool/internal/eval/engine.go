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
)

// EvalResult holds the raw output from a Copilot evaluation.
type EvalResult struct {
GeneratedFiles []string
EventCount     int
ToolCalls      []string
Success        bool
Error          string
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
Success:        true,
Error:          "",
}, nil
}

// EngineOptions configures the evaluation engine.
type EngineOptions struct {
Workers    int
Timeout    time.Duration
OutputDir  string
SkipTests  bool
SkipReview bool
Debug      bool
DryRun     bool
}

// Engine orchestrates evaluation runs.
type Engine struct {
evaluator CopilotEvaluator
opts      EngineOptions
}

// NewEngine creates a new Engine with the given evaluator and options.
func NewEngine(evaluator CopilotEvaluator, opts EngineOptions) *Engine {
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

	// Write run summary
	if _, err := report.WriteSummary(summary, e.opts.OutputDir); err != nil && e.opts.Debug {
		log.Printf("failed to write run summary: %v", err)
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

// Setup workspace
ws, err := NewWorkspace(e.opts.OutputDir, task.Prompt.ID, task.Config.Name)
if err != nil {
evalReport.Error = fmt.Sprintf("workspace setup failed: %v", err)
evalReport.Duration = time.Since(start).Seconds()
return evalReport
}

// Run evaluation (stub)
result, err := e.evaluator.Evaluate(evalCtx, task.Prompt, &task.Config, ws.Dir)
if err != nil {
evalReport.Error = fmt.Sprintf("evaluation failed: %v", err)
evalReport.Duration = time.Since(start).Seconds()
return evalReport
}

evalReport.GeneratedFiles = result.GeneratedFiles
evalReport.EventCount = result.EventCount
evalReport.ToolCalls = result.ToolCalls

// Build verification
buildResult, err := build.Verify(evalCtx, task.Prompt.Language, ws.Dir)
if err != nil {
evalReport.Error = fmt.Sprintf("build verification failed: %v", err)
evalReport.Duration = time.Since(start).Seconds()
return evalReport
}
evalReport.Build = buildResult
evalReport.Success = result.Success && (buildResult == nil || buildResult.Success)

evalReport.Duration = time.Since(start).Seconds()

// Write report
reportPath, err := report.WriteReport(evalReport, e.opts.OutputDir, runID, task.Prompt)
if err != nil {
if e.opts.Debug {
log.Printf("failed to write report: %v", err)
}
} else if e.opts.Debug {
log.Printf("report written to %s", reportPath)
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
