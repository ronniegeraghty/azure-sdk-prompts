package eval

import (
	"context"
	"fmt"
	"io"
	"math"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/comparison"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/config/tool"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria/graders"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/pairwise"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/process"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/progress"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/prompt"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/report"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/review"
)

// EvalResult holds the raw output from a Copilot evaluation.
type EvalResult struct {
	GeneratedFiles []string
	EventCount     int
	ToolCalls      []string
	SessionEvents  []report.SessionEventRecord
	ActionTimeline *ActionTimeline
	Success        bool
	Error          string
	ErrorDetails   string
	IsStub         bool
	StarterFiles   []string
	// CleanupFn deletes session state after the caller has consumed the
	// workspace files. Must be called after copying generated files out
	// of the workspace directory (#261). Nil for stub evaluators.
	CleanupFn func()
}

// PromptRunner defines the interface for running evaluations.
type PromptRunner interface {
	Run(ctx context.Context, prompt *prompt.Prompt, config *config.ToolConfig, workDir string) (*EvalResult, error)
}

// ReviewerFactory creates a reviewer for a specific config.
// Returns nil if no reviewer should be created (e.g., stub mode or review disabled).
type ReviewerFactory func(cfg *config.ToolConfig) (review.Reviewer, *review.PanelReviewer, error)

// StubRunner returns placeholder results for testing.
type StubRunner struct{}

// Evaluate returns a stub result and creates a stub output file in the workspace.
func (s *StubRunner) Run(ctx context.Context, p *prompt.Prompt, cfg *config.ToolConfig, workDir string) (*EvalResult, error) {
	// Write a stub file so file graders can find it on disk.
	if workDir != "" {
		if err := os.WriteFile(filepath.Join(workDir, "stub_output.txt"), []byte("stub"), 0644); err != nil {
			slog.Warn("stub file write failed", "error", err)
		}
	}
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
	Workers   int
	OutputDir string
	SkipReview   bool
	DryRun       bool
	ProgressMode string // "auto", "live", "log", "off"

	// Fan-out visibility (#34)
	ConfirmLargeRuns bool
	AutoConfirm      bool
	// Generator guardrails (#35). Phase 3.5 (#566) dropped the byte-size
	// guardrail entirely — review no longer inlines file contents, and the
	// review-side caps in internal/utils prevent runaway memory.
	MaxTurns          int
	MaxSessionActions int
	MaxFiles          int
	// Tiered limits (#347): separate generator and reviewer guardrails.
	// Reviewer limits default to half of generator limits when zero.
	ReviewerMaxTurns   int
	ReviewerMaxActions int
	// Process lifecycle (#46)
	StrictCleanup bool // Fail run if orphaned processes detected after cleanup.
	// Session timeout — maximum duration for a single SendAndWait call
	// (generation or review). Defaults to 10 minutes. Per-prompt Timeout
	// frontmatter overrides this for the generation phase.
	SessionTimeout time.Duration
	// Resource monitoring (#45)
	MonitorResources bool
	// Tiered criteria (#30)
	CriteriaDir string // Directory containing attribute-matched criteria YAML files.
	// Review session splitting (#580). When "isolated", graders/groups marked
	// with isolate: true are reviewed in their own Copilot session per panel
	// model. Empty string or "combined" preserves legacy single-session
	// behavior. Validated at the CLI layer.
	ReviewMode string
	// Generator safety (#36)
	AllowCloud bool // Allow agent output to provision real Azure resources.
	// Directory exclusion (#63)
	ExcludeDirs []string // Directories to exclude from generated_files output.
	// Pre-flight model availability check (#264).
	// When true, the engine queries the Copilot backend for available models
	// before starting evaluations and fails fast if any configured model
	// (generator or reviewer) is unavailable.
	CheckModels bool
	// Output writer for user-facing messages (defaults to os.Stdout).
	Stdout io.Writer
	// Tracker overrides the default process tracker (used in tests to avoid
	// killing real Copilot CLI processes during orphan scans).
	Tracker *process.ProcessTracker
}

// Engine orchestrates evaluation runs.
type Engine struct {
	evaluator       PromptRunner
	reviewerFactory ReviewerFactory
	opts            EngineOptions
	tracker         *process.ProcessTracker
	// graderBundle holds the unified grader configuration loaded from
	// CriteriaDir. It replaces the pre-#625 dual storage of
	// criteria.GraderConfig (prompt-review criteria) and
	// graders.GraderConfig (typed graders). Every matched entry — prompt
	// or typed — flows through the same Bundle.
	graderBundle *criteria.Bundle
}

// NewEngine creates a new Engine with the given evaluator and options.
func NewEngine(evaluator PromptRunner, opts EngineOptions) *Engine {
	return NewEngineWithReviewerFactory(evaluator, nil, opts)
}

// NewEngineWithReviewerFactory creates a new Engine with an evaluator and reviewer factory.
func NewEngineWithReviewerFactory(evaluator PromptRunner, factory ReviewerFactory, opts EngineOptions) *Engine {
	if opts.Workers <= 0 {
		opts.Workers = 1
	} else if opts.Workers > 8 {
		opts.Workers = 8
	}
	if opts.OutputDir == "" {
		opts.OutputDir = "./reports"
	}
	// Generator guardrail defaults (#35)
	if opts.MaxTurns <= 0 {
		opts.MaxTurns = 25
	}
	if opts.MaxSessionActions <= 0 {
		opts.MaxSessionActions = 50
	}
	if opts.MaxFiles <= 0 {
		opts.MaxFiles = 50
	}
	// Tiered reviewer limits: default to half of generator limits (#347)
	if opts.ReviewerMaxTurns <= 0 {
		opts.ReviewerMaxTurns = opts.MaxTurns / 2
		if opts.ReviewerMaxTurns < 5 {
			opts.ReviewerMaxTurns = 5
		}
	}
	if opts.ReviewerMaxActions <= 0 {
		opts.ReviewerMaxActions = opts.MaxSessionActions / 2
		if opts.ReviewerMaxActions < 10 {
			opts.ReviewerMaxActions = 10
		}
	}
	if opts.SessionTimeout <= 0 {
		opts.SessionTimeout = 10 * time.Minute
	}
	// Resolve to absolute path so workspace directories passed to the Copilot CLI
	// are always absolute. Without this, the agent constructs wrong paths like
	// /home/user/reports/... instead of /home/user/projects/repo/reports/...
	if abs, err := filepath.Abs(opts.OutputDir); err == nil {
		opts.OutputDir = abs
	}
	if opts.Stdout == nil {
		opts.Stdout = os.Stdout
	}
	tracker := opts.Tracker
	if tracker == nil {
		tracker = process.DefaultTracker
	}
	return &Engine{
		evaluator:       evaluator,
		reviewerFactory: factory,
		opts:            opts,
		tracker:         tracker,
	}
}

// printf writes user-facing output to the configured writer.
func (e *Engine) printf(format string, args ...any) {
	fmt.Fprintf(e.opts.Stdout, format, args...)
}

// loadBundle loads the unified grader bundle from CriteriaDir (#625).
// Per-file parse/validation failures are captured in Bundle.FileErrors and
// only surface when a specific eval's properties would have selected the
// malformed file (Q4 deferred-error semantics).
func (e *Engine) loadBundle() {
	if e.opts.CriteriaDir == "" {
		slog.Debug("No criteria directory configured, skipping bundle load")
		return
	}
	if _, err := os.Stat(e.opts.CriteriaDir); os.IsNotExist(err) {
		slog.Debug("Criteria directory does not exist, skipping", "dir", e.opts.CriteriaDir)
		return
	}
	bundle, err := criteria.LoadUnifiedDir(e.opts.CriteriaDir)
	if err != nil {
		slog.Warn("Failed to walk criteria directory", "dir", e.opts.CriteriaDir, "error", err)
		return
	}
	e.graderBundle = bundle
	slog.Info("Loaded unified grader bundle",
		"dir", e.opts.CriteriaDir,
		"files", len(bundle.Configs),
		"file_errors", len(bundle.FileErrors),
	)
}

// matchedForEval returns the grader entries whose file/group/grader `when`
// blocks match the eval's prompt properties. Thin wrapper kept for call-site
// symmetry with reviewBuckets.
func (e *Engine) matchedForEval(props map[string]string) []criteria.MatchedUnifiedEntry {
	return criteria.MatchingUnifiedEntries(e.graderBundle, props)
}

// reviewBuckets builds the set of review buckets for a prompt under the
// configured ReviewMode. In combined mode (default) it returns one bucket
// containing all matched prompt-type graders + prompt criteria. In isolated
// mode it returns one bucket per isolated grader/group plus a shared
// "combined" bucket for the rest. When isolated mode is requested but
// nothing is marked isolate, it logs a warning so the flag is observably
// no-op rather than silently dead.
func (e *Engine) reviewBuckets(p *prompt.Prompt, props map[string]string) []graders.ReviewBucket {
	mode := e.opts.ReviewMode
	if mode == "" {
		mode = criteria.ReviewModeCombined
	}
	matched := e.matchedForEval(props)
	promptMatched, _ := criteria.PartitionMatched(matched)
	if mode == criteria.ReviewModeIsolated && !criteria.HasUnifiedIsolation(promptMatched) {
		slog.Warn("review-mode=isolated requested but no graders or groups are marked isolate; falling back to combined",
			"prompt_id", p.ID)
	}
	return criteria.BuildUnifiedReviewBuckets(promptMatched, p.EvaluationCriteria, mode)
}

// mergedCriteria returns the combined attribute-matched + prompt-specific
// evaluation criteria text, using only the prompt-type entries from the
// Bundle. Kept on Engine for back-compat with the GraderInput.EvalCriteria
// field consumed by review-aware graders in the single-bucket path.
func (e *Engine) mergedCriteria(p *prompt.Prompt, props map[string]string) string {
	matched := e.matchedForEval(props)
	promptMatched, _ := criteria.PartitionMatched(matched)
	entries := make([]criteria.UnifiedGraderEntry, 0, len(promptMatched))
	for _, m := range promptMatched {
		entries = append(entries, m.Entry)
	}
	merged := criteria.MergeUnifiedCriteria(entries, p.EvaluationCriteria)
	if merged == "" {
		return p.EvaluationCriteria
	}
	return merged
}

// EvalTask represents a single prompt+config evaluation to run.
type EvalTask struct {
	Prompt *prompt.Prompt
	Config config.ToolConfig
}

// resolvedLimits holds the effective guardrail limits for a single eval,
// resolved from config-level overrides and engine-level defaults.
type resolvedLimits struct {
	maxTurns          int
	maxFiles          int
	maxSessionActions int
}

func expandGeneratorModels(configs []config.ToolConfig) ([]config.ToolConfig, error) {
	var expanded []config.ToolConfig
	for _, cfg := range configs {
		if cfg.Generator == nil {
			return nil, fmt.Errorf("config %q: generator.model or generator.models is required", cfg.Name)
		}
		models := cfg.Generator.ResolveModels()
		if len(models) == 0 {
			return nil, fmt.Errorf("config %q: generator.model or generator.models is required", cfg.Name)
		}
		for _, model := range models {
			clone := cloneToolConfigForModel(cfg, model)
			if len(models) > 1 {
				clone.Name = fmt.Sprintf("%s/%s", cfg.Name, model)
			}
			expanded = append(expanded, clone)
		}
	}
	return expanded, nil
}

func cloneToolConfigForModel(src config.ToolConfig, model string) config.ToolConfig {
	dst := src
	if src.Generator != nil {
		gen := *src.Generator
		gen.Model = model
		gen.Models = nil
		if len(src.Generator.Tools) > 0 {
			gen.Tools = cloneToolEntries(src.Generator.Tools)
		}
		if len(src.Generator.ExcludedTools) > 0 {
			gen.ExcludedTools = append([]string(nil), src.Generator.ExcludedTools...)
		}
		dst.Generator = &gen
	}
	if src.Reviewer != nil {
		rev := *src.Reviewer
		if len(src.Reviewer.Tools) > 0 {
			rev.Tools = cloneToolEntries(src.Reviewer.Tools)
		}
		if len(src.Reviewer.Models) > 0 {
			rev.Models = append([]string(nil), src.Reviewer.Models...)
		}
		dst.Reviewer = &rev
	}
	if len(src.Plugins) > 0 {
		dst.Plugins = append([]string(nil), src.Plugins...)
	}
	return dst
}

func cloneToolEntries(entries []config.ToolEntry) []config.ToolEntry {
	clone := make([]config.ToolEntry, len(entries))
	for i, te := range entries {
		clone[i] = te
		if te.When != nil {
			m := make(map[string]string, len(te.When))
			for k, v := range te.When {
				m[k] = v
			}
			clone[i].When = m
		}
		if te.Args != nil {
			args := make([]string, len(te.Args))
			copy(args, te.Args)
			clone[i].Args = args
		}
		if te.MCPTools != nil {
			tools := make([]string, len(te.MCPTools))
			copy(tools, te.MCPTools)
			clone[i].MCPTools = tools
		}
	}
	return clone
}

// resolveLimits merges per-prompt and per-config session limits with engine defaults.
// Resolution order: prompt frontmatter > config YAML > CLI flag/engine default.
// Values > 0 take precedence; zero values fall back to the next layer.
func (e *Engine) resolveLimits(cfg config.ToolConfig, p *prompt.Prompt) resolvedLimits {
	rl := resolvedLimits{
		maxTurns:          e.opts.MaxTurns,
		maxFiles:          e.opts.MaxFiles,
		maxSessionActions: e.opts.MaxSessionActions,
	}
	if cfg.Limits != nil {
		if cfg.Limits.MaxTurns > 0 {
			rl.maxTurns = cfg.Limits.MaxTurns
		}
		if cfg.Limits.MaxFiles > 0 {
			rl.maxFiles = cfg.Limits.MaxFiles
		}
		if cfg.Limits.MaxSessionActions > 0 {
			rl.maxSessionActions = cfg.Limits.MaxSessionActions
		}
	}
	// Prompt-level overrides take highest priority (#284)
	if p != nil {
		if p.MaxTurns > 0 {
			rl.maxTurns = p.MaxTurns
		}
		if p.MaxSessionActions > 0 {
			rl.maxSessionActions = p.MaxSessionActions
		}
	}
	return rl
}

// Run executes evaluations for the given prompts crossed with configs.
func (e *Engine) Run(ctx context.Context, prompts []*prompt.Prompt, configs []config.ToolConfig) (*report.RunSummary, error) {
	// Load the unified grader bundle (#625) from CriteriaDir.
	e.loadBundle()

	expandedConfigs, err := expandGeneratorModels(configs)
	if err != nil {
		return nil, err
	}
	configs = expandedConfigs

	// Build task list (cross product: prompts × configs)
	var tasks []EvalTask
	for _, p := range prompts {
		for _, c := range configs {
			tasks = append(tasks, EvalTask{Prompt: p, Config: c})
		}
	}

	// Pre-run summary (#34: fan-out visibility)
	evalCount := len(tasks)
	sessionsPerEval := 2 // generate + review
	sessionLabel := fmt.Sprintf("%d × 2 for generate/review", evalCount)
	if e.opts.SkipReview {
		sessionsPerEval = 1
		sessionLabel = fmt.Sprintf("%d × 1 for generate only", evalCount)
	}
	estimatedSessions := evalCount * sessionsPerEval
	maxSessions := e.opts.Workers * 3
	slog.Info("Evaluation plan",
		"evaluations", evalCount,
		"prompts", len(prompts),
		"configs", len(configs),
		"estimated_sessions", estimatedSessions,
		"workers", e.opts.Workers,
		"max_sessions", maxSessions)

	if e.opts.DryRun {
		e.printf("\n🔍 DRY RUN MODE — No evaluations will be executed\n")
	}
	e.printf("\n📊 Evaluation plan: %d evaluations (%d prompts × %d configs)\n", evalCount, len(prompts), len(configs))
	e.printf("   Estimated Copilot sessions: %d (%s)\n", estimatedSessions, sessionLabel)
	e.printf("   Workers: %d | Max sessions: %d\n\n", e.opts.Workers, maxSessions)

	// Pre-flight model availability check (#264). Query the backend for
	// available models and fail fast if any configured model is unavailable.
	// This prevents mid-eval failures from unavailable reviewer models
	// (e.g. gemini-3-pro-preview).
	if e.opts.CheckModels {
		if checker, ok := e.evaluator.(*CopilotPromptRunner); ok {
			e.printf("🔍 Checking model availability...\n")
			if err := checker.ValidateModelAvailability(ctx, configs); err != nil {
				return nil, fmt.Errorf("pre-flight check failed: %w", err)
			}
			e.printf("✅ All models available\n\n")
		}
	}

	// Confirmation prompt for large runs (#34)
	if evalCount > 10 && e.opts.ConfirmLargeRuns && !e.opts.AutoConfirm {
		e.printf("⚠️  Large run detected (%d evaluations). Continue? [y/N] ", evalCount)
		var answer string
		if _, err := fmt.Scanln(&answer); err != nil {
			// On piped input or EOF, default to "no" for safety.
			slog.Info("No TTY input detected, defaulting to abort")
			return nil, fmt.Errorf("no interactive input available for confirmation")
		}
		answer = strings.TrimSpace(strings.ToLower(answer))
		if answer != "y" && answer != "yes" {
			return nil, fmt.Errorf("run aborted by user (use -y to skip confirmation)")
		}
	}

	if e.opts.DryRun {
		return e.dryRun(ctx, tasks)
	}

	// Resource monitor (#45) — opt-in via --monitor-resources.
	var resMonitor *process.ResourceMonitor
	if e.opts.MonitorResources {
		resMonitor = process.NewResourceMonitor(e.tracker, 5*time.Second)
		resMonitor.Start()
		defer resMonitor.Stop()
	}

	// Wrap context with cancel so signal handler can trigger graceful shutdown (#67).
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Ensure all tracked Copilot processes are cleaned up when Run exits.
	defer func() {
		if errs := e.tracker.TerminateAll(5 * time.Second); len(errs) > 0 {
			for _, err := range errs {
				slog.Warn("Process cleanup error", "error", err)
			}
		}
	}()

	// Set up signal handler so SIGINT/SIGTERM terminates spawned processes.
	sigCh := make(chan os.Signal, 1)
	process.NotifyShutdownSignals(sigCh)
	// Unregister signal handler before closing the channel to prevent
	// a send-on-closed-channel panic (defers execute LIFO).
	defer close(sigCh)
	defer signal.Stop(sigCh)
	go func() {
		first := true
		for sig := range sigCh {
			if first {
				slog.Warn("Received signal — terminating tracked Copilot processes", "signal", sig.String())
				cancel() // Cancel context to unwind in-flight goroutines
				if errs := e.tracker.TerminateAll(5 * time.Second); len(errs) > 0 {
					for _, err := range errs {
						slog.Warn("Process cleanup error", "error", err)
					}
				}
				first = false
			} else {
				// Second signal: cancel context, allow brief grace period for
				// defers to run, then force exit (#67).
				slog.Warn("Received second signal — forcing exit", "signal", sig.String())
				cancel()
				time.Sleep(2 * time.Second)
				os.Exit(1)
			}
		}
	}()

	runID := time.Now().Format("20060102-150405")
	summary := &report.RunSummary{
		RunID:        runID,
		Timestamp:    time.Now().UTC().Format(time.RFC3339),
		TotalPrompts: len(prompts),
		TotalConfigs: len(configs),
		TotalEvals:   len(tasks),
	}

	slog.Info("Starting run", "workers", e.opts.Workers)

	start := time.Now()

	runDir := filepath.Join(e.opts.OutputDir, runID)

	// Progress display — mode is controlled by --progress flag.
	// When --log-level debug/info, main.go sets ProgressMode to "log" automatically.
	display := progress.NewDisplay(progress.DisplayConfig{
		Total:     len(tasks),
		Workers:   e.opts.Workers,
		ReportDir: runDir + "/",
		Mode:      progress.ProgressMode(e.opts.ProgressMode),
	})

	// Wire progress reporting if evaluator supports it
	if pr, ok := e.evaluator.(progress.Reporter); ok {
		pr.SetProgressFunc(display.HandleEvent)
	}

	sem := make(chan struct{}, e.opts.Workers)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, task := range tasks {
		wg.Add(1)
		go func(t EvalTask) {
			defer wg.Done()

			// Acquire worker semaphore. Use select so context cancellation
			// unblocks waiting goroutines instead of leaking them (#129).
			select {
			case sem <- struct{}{}:
			case <-ctx.Done():
				return
			}
			defer func() { <-sem }()

			taskName := t.Prompt.ID + "/" + t.Config.Name

			// Register eval with resource monitor if active (#45)
			if resMonitor != nil {
				resMonitor.RegisterEval(taskName)
			}

			display.HandleEvent(progress.ProgressEvent{
				EvalID:     taskName,
				PromptID:   t.Prompt.ID,
				ConfigName: t.Config.Name,
				Type:       progress.EventStarting,
				Message:    "Waiting for session...",
			})

			// Progress callbacks for runSingleEval
			sendPhase := func(phase progress.Phase) {
				display.HandleEvent(progress.ProgressEvent{
					EvalID: taskName, Type: progress.EventPhaseChange, Phase: phase,
				})
			}
			sendEvent := func(evtType progress.EventType, msg string) {
				display.HandleEvent(progress.ProgressEvent{
					EvalID: taskName, PromptID: t.Prompt.ID, ConfigName: t.Config.Name,
					Type: evtType, Message: msg,
				})
			}
			// sendRawEvent lets callers emit ProgressEvents with richer fields
			// (e.g. GraderID/GraderKind/Result/Score) while still auto-filling
			// identity fields so downstream renderers can match events to evals.
			sendRawEvent := func(evt progress.ProgressEvent) {
				if evt.EvalID == "" {
					evt.EvalID = taskName
				}
				if evt.PromptID == "" {
					evt.PromptID = t.Prompt.ID
				}
				if evt.ConfigName == "" {
					evt.ConfigName = t.Config.Name
				}
				display.HandleEvent(evt)
			}

			evalReport := e.runSingleEval(ctx, t, runID, sendPhase, sendEvent, sendRawEvent)

			// Attach per-eval resource stats (#45)
			if resMonitor != nil {
				if es := resMonitor.EvalStats(taskName); es != nil {
					evalReport.ResourceUsage = &report.ResourceStats{
						PeakCPUPercent: es.PeakCPUPercent,
						PeakMemoryMB:   es.PeakMemoryMB,
						SampleCount:    es.SampleCount,
					}
				}
			}

			evtType := progress.EventPassed
			msg := ""
			reviewScore := 0
			if evalReport.Review != nil {
				reviewScore = evalReport.Review.OverallScore
			}
			if evalReport.Error != "" {
				evtType = progress.EventError
				msg = "ERROR"
			} else if !evalReport.Success {
				evtType = progress.EventFailed
				if evalReport.Review != nil {
					msg = fmt.Sprintf("%d/%d criteria", evalReport.Review.OverallScore, evalReport.Review.MaxScore)
				}
			}
			display.HandleEvent(progress.ProgressEvent{
				EvalID:      taskName,
				PromptID:    t.Prompt.ID,
				ConfigName:  t.Config.Name,
				Type:        evtType,
				Message:     msg,
				FileCount:   len(evalReport.GeneratedFiles),
				ReviewScore: reviewScore,
			})

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
	display.Done()

	// Post-run orphan scan — terminate any leaked copilot processes (#46)
	// Only scan when using the DefaultTracker (production). Test-injected
	// trackers have no registrations, so every real copilot process looks
	// like an orphan — which would kill the user's Copilot CLI.
	if e.tracker == process.DefaultTracker {
		if orphans := e.tracker.TerminateOrphans(); orphans > 0 {
			slog.Warn("Terminated orphaned copilot processes", "count", orphans)
			if e.opts.StrictCleanup {
				return summary, fmt.Errorf("strict-cleanup: %d orphaned copilot processes detected and terminated", orphans)
			}
		}
	}

	summary.Duration = time.Since(start).Seconds()

	// Calculate per-phase average durations across all reports (#44)
	var genSum, reviewSum float64
	var genCount, reviewCount int
	for _, r := range summary.Results {
		if r.GenerationDuration > 0 {
			genSum += r.GenerationDuration
			genCount++
		}
		if r.ReviewDuration > 0 {
			reviewSum += r.ReviewDuration
			reviewCount++
		}
	}
	if genCount > 0 {
		summary.AvgGenerationDuration = genSum / float64(genCount)
	}
	if reviewCount > 0 {
		summary.AvgReviewDuration = reviewSum / float64(reviewCount)
	}

	// Attach aggregate resource stats and print summary (#45)
	if resMonitor != nil {
		rs := resMonitor.RunStats()
		summary.ResourceUsage = &report.RunResourceStats{
			PeakCPUPercent: rs.PeakCPUPercent,
			PeakMemoryMB:   rs.PeakMemoryMB,
			SessionCount:   rs.SessionCount,
		}
		e.printf("\n🔍 Resource usage: %s\n", resMonitor.SummaryLine())
	}

	pairwiseReports, pairwiseImpacts := collectPairwiseReports(summary.Results)
	if len(pairwiseReports) > 0 {
		summary.PairwiseResults = pairwiseReports
		if _, err := report.WritePairwiseReport(runID, pairwiseReports, pairwiseImpacts, e.opts.OutputDir); err != nil {
			slog.Warn("Failed to write pairwise report", "error", err)
		}
	}

	// Auto-generate pairwise comparisons for multi-config runs (#357). Written
	// to {runDir}/comparisons.json so the site comparison surface can render
	// without recomputing and is guaranteed to match what `hyoka compare`
	// prints for the same inputs.
	if len(summary.Results) > 0 {
		runDir := filepath.Join(e.opts.OutputDir, runID)
		derefReports := make([]report.EvalReport, 0, len(summary.Results))
		for _, r := range summary.Results {
			if r != nil {
				derefReports = append(derefReports, *r)
			}
		}
		if _, err := comparison.WriteForRun(runDir, derefReports); err != nil {
			slog.Warn("Failed to write auto-generated comparisons", "error", err)
		}
	}

	// Write JSON summary
	if _, err := report.WriteSummary(summary, e.opts.OutputDir); err != nil {
		slog.Error("Failed to write run summary", "error", err)
	}
	// Write Markdown summary
	if _, err := report.WriteSummaryMarkdown(summary, e.opts.OutputDir); err != nil {
		slog.Error("Failed to write Markdown summary", "error", err)
	}

	return summary, nil
}

func (e *Engine) dryRun(ctx context.Context, tasks []EvalTask) (*report.RunSummary, error) {
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

	// Validate skill directories for each unique config (#291).
	// This surfaces warnings about empty/missing skill dirs during dry run.
	validatedConfigs := make(map[string]bool)
	for _, t := range tasks {
		if validatedConfigs[t.Config.Name] {
			continue
		}
		validatedConfigs[t.Config.Name] = true
		if t.Config.Generator != nil {
			entries := countSkillEntries(t.Config.Generator.Tools)
			resolved, err := tool.ResolveSkills(ctx, t.Config.Generator.Tools, "")
			if err != nil {
				slog.Warn("Failed to resolve generator skills", "config", t.Config.Name, "error", err)
				continue
			}
			count := tool.CountSkills(resolved)
			if entries > 0 {
				e.printf("   Config %q: %d generator dir(s) searched, %d skill(s) found\n", t.Config.Name, entries, count)
			}
		}
		if t.Config.Reviewer != nil {
			entries := countSkillEntries(t.Config.Reviewer.Tools)
			resolved, err := tool.ResolveSkills(ctx, t.Config.Reviewer.Tools, "")
			if err != nil {
				slog.Warn("Failed to resolve reviewer skills", "config", t.Config.Name, "error", err)
				continue
			}
			count := tool.CountSkills(resolved)
			if entries > 0 {
				e.printf("   Config %q: %d reviewer dir(s) searched, %d skill(s) found\n", t.Config.Name, entries, count)
			}
		}
	}

	return summary, nil
}
// countSkillEntries counts how many ToolEntries in the list are skills.
func countSkillEntries(entries []config.ToolEntry) int {
	n := 0
	for _, e := range entries {
		if e.ResolvedType() == "skill" {
			n++
		}
	}
	return n
}
func collectPairwiseReports(results []*report.EvalReport) ([]*pairwise.PairwiseReport, []pairwise.ToolImpact) {
	type key struct {
		promptID string
		baseName string
	}

	byKey := make(map[key][]pairwise.VariantResult)

	for _, r := range results {
		model := ""
		if r.ConfigUsed != nil {
			if m, ok := r.ConfigUsed["model"].(string); ok {
				model = m
			}
		}
		baseName, removedTool, ok := parsePairwiseConfigName(r.ConfigName, model)
		if !ok {
			continue
		}
		score, maxScore := pairwiseScore(r)
		k := key{promptID: r.PromptID, baseName: baseName}
		byKey[k] = append(byKey[k], pairwise.VariantResult{
			ConfigName:  r.ConfigName,
			RemovedTool: removedTool,
			Score:       score,
			MaxScore:    maxScore,
			Success:     r.Success,
		})
	}

	if len(byKey) == 0 {
		return nil, nil
	}

	var reports []*pairwise.PairwiseReport
	for k, variants := range byKey {
		report, err := pairwise.ComputeImpacts(k.promptID, variants)
		if err != nil {
			slog.Warn("Pairwise impact computation failed", "prompt", k.promptID, "config", k.baseName, "error", err)
			continue
		}
		reports = append(reports, report)
	}

	if len(reports) == 0 {
		return nil, nil
	}

	sort.Slice(reports, func(i, j int) bool {
		if reports[i].PromptID != reports[j].PromptID {
			return reports[i].PromptID < reports[j].PromptID
		}
		return reports[i].Baseline.ConfigName < reports[j].Baseline.ConfigName
	})

	impacts := pairwise.AggregateImpacts(reports)
	return reports, impacts
}
func parsePairwiseConfigName(configName, model string) (string, string, bool) {
	suffix := ""
	if model != "" && strings.HasSuffix(configName, "/"+model) {
		suffix = "/" + model
		configName = strings.TrimSuffix(configName, suffix)
	}

	if strings.HasSuffix(configName, "/baseline") {
		base := strings.TrimSuffix(configName, "/baseline")
		return base + suffix, "", true
	}

	if idx := strings.LastIndex(configName, "/without-"); idx != -1 {
		base := configName[:idx]
		removed := configName[idx+len("/without-"):]
		return base + suffix, removed, true
	}

	return "", "", false
}
func pairwiseScore(r *report.EvalReport) (int, int) {
	if r.Review != nil {
		return r.Review.OverallScore, r.Review.MaxScore
	}
	if r.ScoreBreakdown != nil {
		score := int(math.Round(r.ScoreBreakdown.FinalScorePct))
		return score, 100
	}
	return 0, 0
}
