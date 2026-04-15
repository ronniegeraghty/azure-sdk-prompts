package eval

import (
"context"
"fmt"
"log/slog"
"os"
"path/filepath"
"strings"
"time"

"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
"github.com/ronniegeraghty/hyoka/hyoka/internal/config/tool"
"github.com/ronniegeraghty/hyoka/hyoka/internal/graders"
"github.com/ronniegeraghty/hyoka/hyoka/internal/logging"
"github.com/ronniegeraghty/hyoka/hyoka/internal/progress"
"github.com/ronniegeraghty/hyoka/hyoka/internal/report"
"github.com/ronniegeraghty/hyoka/hyoka/internal/review"
)

func (e *Engine) runSingleEval(ctx context.Context, task EvalTask, runID string, sendPhase func(progress.Phase), sendEvent func(progress.EventType, string)) *report.EvalReport {
	// Each phase gets its own independent timeout so a slow generation
	// doesn't starve build or review (fixes issue #3).
	genCtx, genCancel := context.WithCancel(ctx)
	defer genCancel()

	// Structured logger with eval context fields (#42)
	lg := logging.EvalLogger(task.Prompt.ID, task.Config.Name, "generation", 0)
	start := time.Now()

	// Compute prompt properties once — used for criteria matching, grader
	// matching, tool resolution, and report metadata.
	promptProps := mergePromptProperties(task.Prompt)
	props := map[string]string{
		"language": task.Prompt.Language(),
		"service":  task.Prompt.Service(),
		"plane":    task.Prompt.Plane(),
		"category": task.Prompt.Category(),
		"sdk":      task.Prompt.SDKPackage(),
	}
	// Merge prompt-level properties (from frontmatter) into props.
	for k, v := range promptProps {
		if _, exists := props[k]; !exists {
			props[k] = v
		}
	}

	evalReport := &report.EvalReport{
		SchemaVersion: report.CurrentSchemaVersion,
		PromptID:      task.Prompt.ID,
		ConfigName:    task.Config.Name,
		Timestamp:     time.Now().UTC().Format(time.RFC3339),
		PromptMeta: map[string]any{
			"service":     props["service"],
			"plane":       props["plane"],
			"language":    props["language"],
			"category":    props["category"],
			"description": task.Prompt.Description(),
			"difficulty":  task.Prompt.Difficulty(),
			"sdk_package": props["sdk"],
		},
		ConfigUsed: map[string]any{
			"name":  task.Config.Name,
			"model": task.Config.Generator.Model,
		},
	}

	// Resolve effective limits: prompt > per-config > engine defaults (#125, #284)
	lim := e.resolveLimits(task.Config, task.Prompt)
	evalReport.GuardrailMaxTurns = lim.maxTurns
	evalReport.GuardrailMaxFiles = lim.maxFiles
	evalReport.GuardrailMaxOutputSize = lim.maxOutputSize
	evalReport.GuardrailMaxSessionActions = lim.maxSessionActions

	if len(task.Prompt.Tags) > 0 {
		evalReport.PromptMeta["tags"] = strings.Join(task.Prompt.Tags, ", ")
	}

	lg.Info("Starting Copilot session")

	// Build the report directory path early — workspace lives in the report tree (Issue 2)
	reportDir := filepath.Join(report.ReportDir(e.opts.OutputDir, runID, task.Prompt), task.Config.Name)
	codeDir := filepath.Join(reportDir, "generated-code")

	ws, err := NewWorkspaceAt(codeDir)
	if err != nil {
		evalReport.Error = fmt.Sprintf("workspace setup failed: %v", err)
		evalReport.ErrorDetails = err.Error()
		evalReport.ErrorCategory = "generation_failure"
		evalReport.FailureReason = fmt.Sprintf("Could not create workspace directory: %v", err)
		evalReport.Duration = time.Since(start).Seconds()
		return evalReport
	}

	// Create an isolated temporary workspace for the generator (#26, #126).
	// The agent writes files here — not directly to the report tree.
	// After generation, files are copied to the persistent report directory.
	genWs, err := NewWorkspace(task.Prompt.ID, task.Config.Name)
	if err != nil {
		evalReport.Error = fmt.Sprintf("generator workspace setup failed: %v", err)
		evalReport.ErrorDetails = err.Error()
		evalReport.ErrorCategory = "generation_failure"
		evalReport.FailureReason = fmt.Sprintf("Could not create isolated eval workspace: %v", err)
		evalReport.Duration = time.Since(start).Seconds()
		return evalReport
	}
	defer genWs.Cleanup()
	genDir := genWs.Dir

	// Copy starter files into the generator workspace before evaluation (#127).
	starterFiles, starterErr := genWs.CopyStarterFiles(task.Prompt)
	if starterErr != nil {
		evalReport.Error = fmt.Sprintf("starter file copy failed: %v", starterErr)
		evalReport.ErrorDetails = starterErr.Error()
		evalReport.ErrorCategory = "generation_failure"
		evalReport.FailureReason = fmt.Sprintf("Could not copy starter project: %v", starterErr)
		evalReport.Duration = time.Since(start).Seconds()
		return evalReport
	}
	evalReport.StarterFiles = starterFiles

	lg.Debug("Workspace created", "workspace", ws.Dir, "gen_dir", genDir,
		"starter_files", len(starterFiles))

	// Run evaluation (generation phase — uses its own timeout)
	sendPhase(progress.PhaseGenerating)

	genStart := time.Now()
	result, err := e.evaluator.Run(genCtx, task.Prompt, &task.Config, genDir)
	genCtxErr := genCtx.Err() // capture before cancel to distinguish real cancellation from cleanup
	genCancel()               // release generation context immediately
	evalReport.GenerationDuration = time.Since(genStart).Seconds()
	evalFailed := err != nil
	if evalFailed {
		if genCtxErr == context.Canceled {
			evalReport.Error = "generation cancelled (action limit reached)"
			evalReport.ErrorDetails = "context cancelled — the session exceeded the --max-session-actions limit"
			evalReport.ErrorCategory = "timeout"
			evalReport.FailureReason = "Generation cancelled due to action limit"
		} else {
			evalReport.Error = fmt.Sprintf("evaluation failed: %v", err)
			evalReport.ErrorDetails = err.Error()
			evalReport.ErrorCategory = "sdk_error"
			evalReport.FailureReason = fmt.Sprintf("SDK evaluation error: %v", err)
		}
		// Capture whatever session events were collected before failure
		if result != nil {
			evalReport.SessionEvents = result.SessionEvents
			evalReport.EventCount = result.EventCount
			evalReport.ToolCalls = result.ToolCalls
			evalReport.IsStub = result.IsStub
			if result.ActionTimeline != nil {
				evalReport.ActionTimeline = result.ActionTimeline.ToReport()
			}
		}
		// Don't return early — continue to collect files and run review for diagnostics
	}

	if result != nil && !evalFailed {
		evalReport.EventCount = result.EventCount
		evalReport.ToolCalls = result.ToolCalls
		evalReport.SessionEvents = result.SessionEvents
		evalReport.IsStub = result.IsStub
		evalReport.Success = result.Success
		if result.ActionTimeline != nil {
			evalReport.ActionTimeline = result.ActionTimeline.ToReport()
		}
	}

	// Collect generated files — workspace listing is the primary source since
	// ForceStop preserves files on disk.
	// PreToolUse hooks now enforce containment at the tool level (#346),
	// so the snapshot-and-recovery approach is no longer needed.

	// Copy generated files from isolated workspace to persistent report directory (#26)
	if err := copyDir(genDir, ws.Dir); err != nil {
		lg.Warn("Failed to copy generated files to report dir", "error", err)
	}

	// Release session state AFTER workspace files are safely copied (#261).
	// Previously, Evaluate's defer called DeleteSession before returning,
	// which removed workspace artifacts before copyDir could read them.
	// Baseline configs (no MCP) were affected because there were no MCP
	// server processes keeping files alive during cleanup.
	if result != nil && result.CleanupFn != nil {
		result.CleanupFn()
	}

	generatedFiles, listErr := ws.ListFiles()
	if listErr != nil {
		lg.Warn("Failed to list workspace files", "error", listErr)
	}
	if len(generatedFiles) == 0 && result != nil && len(result.GeneratedFiles) > 0 {
		generatedFiles = result.GeneratedFiles
	}
	// Apply directory exclusion filter (#63)
	if len(e.opts.ExcludeDirs) > 0 {
		before := len(generatedFiles)
		generatedFiles = filterExcludedDirs(generatedFiles, e.opts.ExcludeDirs)
		if excluded := before - len(generatedFiles); excluded > 0 {
			lg.Debug("Excluded files by directory filter", "excluded", excluded, "remaining", len(generatedFiles))
		}
	}
	evalReport.GeneratedFiles = generatedFiles

	// Diagnostic: if 0 files generated, check if agent attempted file creation
	if len(generatedFiles) == 0 && !evalFailed {
		fileToolAttempts := 0
		for _, ev := range evalReport.SessionEvents {
			if ev.Type == "tool.execution_start" && isFileWriteTool(ev.ToolName) {
				fileToolAttempts++
			}
		}
		if fileToolAttempts > 0 {
			lg.Warn("0 files generated despite file-write tool attempts", "attempts", fileToolAttempts)
			if evalReport.Error == "" {
				evalReport.Error = fmt.Sprintf("0 files generated despite %d file-write tool attempts", fileToolAttempts)
				evalReport.ErrorCategory = "no_files"
				evalReport.FailureReason = fmt.Sprintf("Generator made %d file-write attempts but no files appeared in the workspace — files may have been written to the wrong location", fileToolAttempts)
				evalReport.Success = false
			}
		} else {
			lg.Warn("0 files generated — agent did not use any file-write tools")
			if evalReport.Error == "" {
				evalReport.Error = "0 files generated — agent did not create any files"
				evalReport.ErrorCategory = "no_files"
				evalReport.FailureReason = "Generator produced no files — the agent did not invoke any file-write tools"
				evalReport.Success = false
			}
		}
	}

	lg.Debug("Session complete",
		"tool_calls", len(evalReport.ToolCalls),
		"files_generated", len(generatedFiles),
		"elapsed", time.Since(start).Truncate(time.Millisecond).String())

	// Per-phase generation duration already captured above (evalReport.GenerationDuration).
	// Overall Duration is set at the end of the function after all phases complete.

	// Populate environment info from config and captured events.
	// Use ResolveSkillDirs for accurate directory resolution (#291).
	var skillDirectories []string
	if task.Config.Generator != nil {
		resolved, err := tool.ResolveSkills(task.Config.Generator.Tools, "")
		if err != nil {
			slog.Warn("Failed to resolve skill directories for report", "error", err)
		} else {
			skillDirectories = resolved
		}
	}
	// Resolve tools for reporting — mirrors the resolution in buildSessionConfig.
	var toolEntries []config.ToolEntry
	for _, entry := range task.Config.Generator.Tools {
		if entry.ResolvedType() == "tool" {
			toolEntries = append(toolEntries, entry)
		}
	}
	reportAvailableTools := config.ResolveTools(toolEntries, promptProps)
	env := &report.EnvironmentInfo{
		Model:            task.Config.Generator.Model,
		SkillDirectories: skillDirectories,
		AvailableTools:   reportAvailableTools,
		ExcludedTools:    task.Config.Generator.ExcludedTools,
		SafetyBoundaries: !e.opts.AllowCloud,
		AllowCloud:       e.opts.AllowCloud,
		WorkingDirectory: ws.Dir,
	}
	// Extract MCP server names
	for _, entry := range task.Config.Generator.Tools {
		if entry.ResolvedType() == "mcp" {
			env.MCPServers = append(env.MCPServers, entry.Name)
		}
	}
	// Derive token usage, turn count, truncation, skills from events
	for _, ev := range evalReport.SessionEvents {
		switch ev.Type {
		case "assistant.usage":
			env.TotalInputTokens += ev.InputTokens
			env.TotalOutputTokens += ev.OutputTokens
		case "assistant.turn_start":
			env.TurnCount++
		case "session.truncation":
			env.ContextTruncated = true
		case "skill.invoked":
			if ev.SkillName != "" {
				env.SkillsInvoked = append(env.SkillsInvoked, ev.SkillName)
			}
		case "session.skills_loaded":
			if ev.Content != "" {
				env.SkillsLoaded = strings.Split(ev.Content, ", ")
			}
		}
	}
	// Derive MCP tools invoked from action timeline events.
	if evalReport.ActionTimeline != nil {
		seen := map[string]bool{}
		for _, ev := range evalReport.ActionTimeline.Events {
			if ev.MCPServer != "" && ev.Tool != "" && ev.Action == "start" {
				if !seen[ev.Tool] {
					seen[ev.Tool] = true
					env.MCPToolsInvoked = append(env.MCPToolsInvoked, ev.Tool)
				}
			}
		}
	}
	evalReport.Environment = env

	// Build tool availability summary: what was available vs actually used (#348).
	evalReport.ToolAvailability = buildToolAvailability(env, evalReport.SessionEvents)

	// Build SessionSetup from config and starter files (#219).
	setup := &report.SessionSetupEvent{
		Tools:        reportAvailableTools,
		StarterFiles: evalReport.StarterFiles,
	}
	// Record configured MCP servers with details.
	for _, entry := range task.Config.Generator.Tools {
		if entry.ResolvedType() != "mcp" {
			continue
		}
		details := entry.Command
		if len(entry.Args) > 0 {
			details += " " + strings.Join(entry.Args, " ")
		}
		setup.MCPServers = append(setup.MCPServers, report.ToolLoadResult{
			Name:    entry.Name,
			Status:  "configured",
			Details: details,
		})
	}
	// Record configured skills with details.
	for _, entry := range task.Config.Generator.Tools {
		if entry.ResolvedType() != "skill" {
			continue
		}
		name := entry.Name
		if name == "" {
			if entry.Path != "" {
				name = entry.Path
			} else {
				name = entry.Repo
			}
		}
		details := entry.SkillSource()
		setup.Skills = append(setup.Skills, report.ToolLoadResult{
			Name:    name,
			Status:  "configured",
			Details: details,
		})
	}
	// Determine system prompt status.
	if task.Config.Generator.SystemPrompt != "" {
		setup.SystemPrompt = fmt.Sprintf("custom (%d chars)", len(task.Config.Generator.SystemPrompt))
	} else {
		setup.SystemPrompt = "none (default)"
	}
	evalReport.SessionSetup = setup

	// Generator guardrail checks (#35, #125)
	if !evalFailed {
		// Check turn count (assistant.message events = conversation turns)
		turnCount := 0
		for _, ev := range evalReport.SessionEvents {
			if ev.Type == "assistant.message" {
				turnCount++
			}
		}
		if turnCount > lim.maxTurns {
			reason := fmt.Sprintf("guardrail: turn count %d exceeded limit of %d", turnCount, lim.maxTurns)
			evalReport.GuardrailAbortReason = reason
			evalReport.Error = reason
			evalReport.Success = false
			lg.Warn("Guardrail triggered", "reason", reason, "turns", turnCount, "max_turns", lim.maxTurns)
		}

		// Check action count — soft cap: note but don't error, let review decide pass/fail
		actionCount := 0
		for _, ev := range evalReport.SessionEvents {
			if ev.Type == "assistant.reasoning" || ev.Type == "assistant.message" || ev.Type == "tool.execution_start" {
				actionCount++
			}
		}
		if actionCount > lim.maxSessionActions {
			evalReport.ActionLimitReached = true
			evalReport.ActionCount = actionCount
			lg.Warn("Action limit reached (soft cap — review will proceed with partial results)",
				"actions", actionCount, "max_session_actions", lim.maxSessionActions)
		}

		// Check file count
		if len(generatedFiles) > lim.maxFiles {
			reason := fmt.Sprintf("guardrail: file count %d exceeded limit of %d", len(generatedFiles), lim.maxFiles)
			evalReport.GuardrailAbortReason = reason
			evalReport.Error = reason
			evalReport.Success = false
			lg.Warn("Guardrail triggered", "reason", reason, "files", len(generatedFiles), "max_files", lim.maxFiles)
		}

		// Check total output size
		var totalSize int64
		for _, f := range generatedFiles {
			absPath := f
			if !filepath.IsAbs(f) {
				absPath = filepath.Join(ws.Dir, f)
			}
			if info, err := os.Stat(absPath); err == nil {
				totalSize += info.Size()
			}
		}
		if totalSize > lim.maxOutputSize {
			reason := fmt.Sprintf("guardrail: total output size %d bytes exceeded limit of %d bytes", totalSize, lim.maxOutputSize)
			evalReport.GuardrailAbortReason = reason
			evalReport.Error = reason
			evalReport.Success = false
			lg.Warn("Guardrail triggered", "reason", reason, "total_size", totalSize, "max_size", lim.maxOutputSize)
		}
	}

	// Unified grading pipeline (WI-023) — all graders (pluggable + AI review)
	// run in a single phase. The review is now a grader type ("prompt_review"),
	// not a separate phase. This fixes the results overwrite bug where
	// GraderResultsFromReview() clobbered pluggable grader results.
	if len(generatedFiles) > 0 {
		gradeStart := time.Now()
		glg := logging.WithPhase(lg, "grading")

		// Collect all grader results across pluggable graders and AI review.
		var allGraderResults []graders.GraderResult

		// Build common grader input shared by all grader types.
		graderInput := graders.GraderInput{
			WorkspacePath:  genWs.Dir,
			OriginalPrompt: task.Prompt.PromptText,
			EvalCriteria:   e.mergedCriteria(task.Prompt, props),
		}
		if task.Prompt.ReferenceAnswer != "" {
			graderInput.ReferenceDir = task.Prompt.ReferenceAnswer
		}
		if result != nil && result.ActionTimeline != nil {
			graderInput.ActionLog = result.ActionTimeline.ToGraderActionLog()
		}

		// --- Phase 1: Pluggable graders (file, program, behavior, etc.) ---
		if len(e.pluginGraders) > 0 {
			applicable := graders.ApplicableGraders(e.pluginGraders, props)
			glg.Debug("Applicable graders", "total", len(e.pluginGraders), "applicable", len(applicable))

			if len(applicable) > 0 {
				instances, err := graders.InstantiateGraders(applicable)
				if err != nil {
					glg.Error("Failed to instantiate graders", "error", err)
				} else {
					pluginResults := graders.RunGraders(ctx, instances, applicable, graderInput)
					allGraderResults = append(allGraderResults, pluginResults...)
					glg.Debug("Pluggable graders complete", "count", len(pluginResults))
				}
			}
		}

		// --- Phase 2: AI review grader (runs alongside pluggable graders) ---
		var reviewGrader *graders.PromptReviewGrader
		if !e.opts.SkipReview {
			sendPhase(progress.PhaseReviewing)

			var reviewer review.Reviewer
			var panelReviewer *review.PanelReviewer
			if e.reviewerFactory != nil {
				reviewer, panelReviewer, err = e.reviewerFactory(&task.Config)
				if err != nil {
					glg.Warn("Reviewer creation failed, skipping review", "error", err)
				}
			}

			if panelReviewer != nil || reviewer != nil {
				reviewGrader = graders.NewPromptReviewGrader("ai_review", reviewer, panelReviewer)

				if panelReviewer != nil {
					models := panelReviewer.Models()
					sendEvent(progress.EventToolStart, fmt.Sprintf("Review panel: %v", models))
				} else {
					sendEvent(progress.EventToolStart, "Single model review")
				}

				reviewResult, reviewErr := reviewGrader.Grade(ctx, graderInput)
				if reviewErr != nil {
					glg.Error("Review grader failed", "error", reviewErr)
					sendEvent(progress.EventReasoning, fmt.Sprintf("Review failed: %v", reviewErr))
					// Add a failing result so aggregation accounts for the review attempt.
					reviewResult = graders.GraderResult{
						Kind:    graders.KindPromptReview,
						Name:    "ai_review",
						Pass:    false,
						Score:   0,
						Message: fmt.Sprintf("review grader error: %v", reviewErr),
					}
				} else {
					if reviewGrader.LastConsolidated != nil {
						sendEvent(progress.EventToolComplete, fmt.Sprintf("Review complete: %d/%d criteria passed",
							reviewGrader.LastConsolidated.OverallScore, reviewGrader.LastConsolidated.MaxScore))
					}
				}
				// Apply default weight — review has weight 1.0, not a gate grader.
				if reviewResult.Weight == 0 {
					reviewResult.Weight = 1.0
				}
				allGraderResults = append(allGraderResults, reviewResult)

				// Populate backward-compat report fields.
				evalReport.ReviewPanel = reviewGrader.LastPanel
				evalReport.Review = reviewGrader.LastConsolidated
			}
		}

		// --- Aggregate all results and update report ---
		if len(allGraderResults) > 0 {
			agg, aggErr := graders.AggregateResults(allGraderResults)
			if aggErr != nil {
				glg.Error("Failed to aggregate grader results", "error", aggErr)
			} else {
				reportResults := convertGraderResults(agg.Results)
				evalReport.GraderResults = reportResults
				evalReport.ScoreBreakdown = report.BuildScoreBreakdown(reportResults)

				if !agg.Pass && !evalFailed {
					evalReport.Success = false
					if agg.GateFailed {
						evalReport.FailureReason = "gate grader(s) failed"
					}
				}

				glg.Info("Grader execution complete",
					"graders", len(allGraderResults),
					"score", fmt.Sprintf("%.2f", agg.Score),
					"passed", agg.Pass,
					"gate_failed", agg.GateFailed)
				sendEvent(progress.EventToolComplete, fmt.Sprintf("Graders: %.0f%% (%d/%d passed)",
					agg.Score*100, countPassed(allGraderResults), len(allGraderResults)))
			}
		}

		// Capture reviewed (annotated) files from the reviewer workspace.
		if reviewGrader != nil && reviewGrader.LastReviewWorkDir != "" {
			reviewedFiles, rfErr := readReviewedFiles(reviewGrader.LastReviewWorkDir)
			if rfErr == nil && len(reviewedFiles) > 0 {
				evalReport.ReviewedFiles = reviewedFiles
				glg.Debug("Captured reviewed files", "count", len(reviewedFiles))
			}
			reviewGrader.CleanupWorkspace()
		}

		evalReport.ReviewDuration = time.Since(gradeStart).Seconds()
	}

	// Tool usage evaluation (compare expected vs actual tools)
	if len(task.Prompt.ExpectedTools) > 0 {
		evalReport.ToolUsage = evaluateToolUsage(task.Prompt.ExpectedTools, evalReport.ToolCalls)
		lg.Debug("Tool usage evaluated",
			"match", evalReport.ToolUsage.Match,
			"matched", evalReport.ToolUsage.MatchedTools,
			"missing", evalReport.ToolUsage.MissingTools)
	}

	// Copy reviewed (annotated) files into report under reviewed-code/
	if len(evalReport.ReviewedFiles) > 0 {
		reviewedDir := filepath.Join(reportDir, "reviewed-code")
		if err := writeReviewedFiles(reviewedDir, evalReport.ReviewedFiles); err != nil {
			lg.Error("Failed to write reviewed files", "error", err)
		} else {
			lg.Debug("Wrote reviewed files", "count", len(evalReport.ReviewedFiles), "dir", reviewedDir)
		}
	}

	// Build re-run command so users can reproduce this evaluation
	evalReport.RerunCommand = buildRerunCommand(task.Prompt.ID, task.Config.Name, e.opts)

	// Capture overall duration after all phases (generation, build, review) complete.
	evalReport.Duration = time.Since(start).Seconds()

	// Write JSON report
	reportPath, err := report.WriteReport(evalReport, e.opts.OutputDir, runID, task.Prompt)
	if err != nil {
		lg.Error("Failed to write report", "error", err)
	} else {
		lg.Debug("Report written", "path", reportPath)
	}
	// Write Markdown report
	if _, err := report.WriteMarkdownReport(evalReport, e.opts.OutputDir, runID,
		props["service"], props["plane"], props["language"], props["category"]); err != nil {
		lg.Error("Failed to write Markdown report", "error", err)
	}

	lg.Info("Evaluation complete",
		"success", evalReport.Success,
		"files_generated", len(evalReport.GeneratedFiles),
		"elapsed", fmt.Sprintf("%.2fs", evalReport.Duration))

	return evalReport
}

// buildRerunCommand constructs the CLI command to reproduce a single evaluation.
func buildRerunCommand(promptID, configName string, opts EngineOptions) string {
	parts := []string{"hyoka run"}
	parts = append(parts, "--prompt-id", promptID)
	parts = append(parts, "--config", configName)

	if opts.SkipReview {
		parts = append(parts, "--skip-review")
	}
	if opts.MonitorResources {
		parts = append(parts, "--monitor-resources")
	}

	if opts.MaxSessionActions != 50 {
		parts = append(parts, fmt.Sprintf("--max-session-actions=%d", opts.MaxSessionActions))
	}

	return strings.Join(parts, " ")
}

// buildToolAvailability constructs a summary of tools available vs actually used
// during a generation session. It combines AvailableTools, MCPServers, and
// SkillsInvoked from EnvironmentInfo with events to determine usage (#348).
func buildToolAvailability(env *report.EnvironmentInfo, events []report.SessionEventRecord) []report.ToolAvailabilityEntry {
	if env == nil {
		return nil
	}

	// Collect tools actually used from session events
	toolsUsed := make(map[string]bool)
	skillsUsed := make(map[string]bool)
	mcpUsed := make(map[string]bool)

	for _, ev := range events {
		switch ev.Type {
		case "tool.execution_complete":
			if ev.ToolName != "" {
				toolsUsed[ev.ToolName] = true
			}
		case "skill.invoked":
			if ev.SkillName != "" {
				skillsUsed[ev.SkillName] = true
			}
		case "external_tool.completed":
			if ev.MCPServerName != "" {
				mcpUsed[ev.MCPServerName] = true
			}
			if ev.ToolName != "" {
				toolsUsed[ev.ToolName] = true
			}
		}
	}

	var entries []report.ToolAvailabilityEntry

	// Available tools (built-in tools like bash, create, etc.)
	for _, t := range env.AvailableTools {
		entries = append(entries, report.ToolAvailabilityEntry{
			Name:      t,
			Type:      "tool",
			Available: true,
			Used:      toolsUsed[t],
		})
	}

	// MCP servers
	for _, s := range env.MCPServers {
		entries = append(entries, report.ToolAvailabilityEntry{
			Name:      s,
			Type:      "mcp",
			Available: true,
			Used:      mcpUsed[s],
		})
	}

	// Skills — include loaded and invoked
	skillSet := make(map[string]bool)
	for _, s := range env.SkillsLoaded {
		if !skillSet[s] {
			skillSet[s] = true
			entries = append(entries, report.ToolAvailabilityEntry{
				Name:      s,
				Type:      "skill",
				Available: true,
				Used:      skillsUsed[s],
			})
		}
	}
	// Add any invoked skills that weren't in the loaded list
	for _, s := range env.SkillsInvoked {
		if !skillSet[s] {
			skillSet[s] = true
			entries = append(entries, report.ToolAvailabilityEntry{
				Name:      s,
				Type:      "skill",
				Available: true,
				Used:      true,
			})
		}
	}

	return entries
}

// evaluateToolUsage compares expected tools from prompt frontmatter with actual tool calls.
func evaluateToolUsage(expected, actual []string) *report.ToolUsageResult {
	actualSet := make(map[string]bool, len(actual))
	for _, t := range actual {
		actualSet[t] = true
	}

	var matched, missing []string
	expectedSet := make(map[string]bool, len(expected))
	for _, t := range expected {
		expectedSet[t] = true
		if actualSet[t] {
			matched = append(matched, t)
		} else {
			missing = append(missing, t)
		}
	}

	var extra []string
	for _, t := range actual {
		if !expectedSet[t] {
			extra = append(extra, t)
		}
	}

	return &report.ToolUsageResult{
		ExpectedTools: expected,
		ActualTools:   actual,
		MatchedTools:  matched,
		MissingTools:  missing,
		ExtraTools:    extra,
		Match:         len(missing) == 0,
	}
}

// readReviewedFiles reads all files from the workspace and returns them as ReviewedFile entries.
// Files that contain "REVIEW:" comments are considered annotated.
func readReviewedFiles(dir string) ([]report.ReviewedFile, error) {
	var reviewed []report.ReviewedFile
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			if info != nil && info.IsDir() && strings.HasPrefix(info.Name(), ".") && path != dir {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasPrefix(filepath.Base(path), ".") || info.Size() > 1<<20 {
			return nil
		}
		rel, err := filepath.Rel(dir, path)
		if err != nil {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		content := string(data)
		if strings.Contains(content, "REVIEW:") {
			reviewed = append(reviewed, report.ReviewedFile{
				Path:    rel,
				Content: content,
			})
		}
		return nil
	})
	return reviewed, err
}

// writeReviewedFiles writes annotated files to the reviewed-code directory.
func writeReviewedFiles(dir string, files []report.ReviewedFile) error {
	for _, f := range files {
		dst := filepath.Join(dir, f.Path)
		if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
			return fmt.Errorf("creating dir for %s: %w", f.Path, err)
		}
		if err := os.WriteFile(dst, []byte(f.Content), 0644); err != nil {
			return fmt.Errorf("writing %s: %w", f.Path, err)
		}
	}
	return nil
}

// countPassed returns the number of passed results in a grader result list.
func countPassed(results []graders.GraderResult) int {
	n := 0
	for _, r := range results {
		if r.Pass {
			n++
		}
	}
	return n
}

// convertGraderResults converts internal grader results to report format.
// Review grader results are expanded into individual panel + consensus entries
// for backward compatibility with the report schema.
func convertGraderResults(results []graders.GraderResult) []report.GraderResult {
	var reportResults []report.GraderResult
	for _, r := range results {
		if r.Kind == graders.KindPromptReview && r.ReviewDetails != nil {
			reportResults = append(reportResults, expandReviewGraderResult(r)...)
			continue
		}
		pass := r.Pass
		rr := report.GraderResult{
			GraderName: r.Name,
			GraderType: r.Kind,
			Summary:    r.Message,
			Score:      r.Score,
			Weight:     r.Weight,
			Pass:       &pass,
			Gate:       r.Gate,
		}
		if r.FileDetails != nil {
			checks := make([]report.FileCheckDetail, len(r.FileDetails.CheckedFiles))
			for j, c := range r.FileDetails.CheckedFiles {
				checks[j] = report.FileCheckDetail{
					Path: c.Path, Exists: c.Exists,
					PatternMatched: c.PatternMatched, Pattern: c.Pattern,
				}
			}
			rr.FileDetails = &report.FileGraderDetail{CheckedFiles: checks}
		}
		if r.ProgramDetails != nil {
			rr.ProgramDetails = &report.ProgramGraderDetail{
				Command: r.ProgramDetails.Command, ExitCode: r.ProgramDetails.ExitCode,
				Stdout: r.ProgramDetails.Stdout, Stderr: r.ProgramDetails.Stderr,
			}
		}
		if r.PromptDetails != nil {
			rr.PromptDetails = &report.PromptGraderDetail{
				Model: r.PromptDetails.Model, Rubric: r.PromptDetails.Rubric,
				Reasoning: r.PromptDetails.Reasoning,
				RawScore:  r.PromptDetails.RawScore, MaxScore: r.PromptDetails.MaxScore,
			}
		}
		if r.BehaviorDetails != nil {
			rr.BehaviorDetails = &report.BehaviorGraderDetail{
				ToolsUsed: r.BehaviorDetails.ToolsUsed, MissingTools: r.BehaviorDetails.MissingTools,
				ForbiddenUsed: r.BehaviorDetails.ForbiddenUsed,
				TurnCount:     r.BehaviorDetails.TurnCount, MaxTurns: r.BehaviorDetails.MaxTurns,
				ActualTurns: r.BehaviorDetails.ActualTurns, TotalActions: r.BehaviorDetails.TotalActions,
				TurnLimitHit: r.BehaviorDetails.TurnLimitHit, Violations: r.BehaviorDetails.Violations,
				SequenceMatch:    r.BehaviorDetails.SequenceMatch,
				ExpectedSequence: r.BehaviorDetails.ExpectedSequence,
				ActualSequence:   r.BehaviorDetails.ActualSequence,
				MatchedActions:   r.BehaviorDetails.MatchedActions,
				ConstraintsMet:   r.BehaviorDetails.ConstraintsMet,
				ToolCounts:       r.BehaviorDetails.ToolCounts,
			}
		}
		reportResults = append(reportResults, rr)
	}
	return reportResults
}

// expandReviewGraderResult converts a review grader result into multiple
// report entries — one per panel member plus the consensus. This replaces
// the removed GraderResultsFromReview function in the main code path.
func expandReviewGraderResult(r graders.GraderResult) []report.GraderResult {
	rd := r.ReviewDetails
	var results []report.GraderResult

	// Panel member entries.
	for _, p := range rd.PanelResults {
		scores := review.ReviewScores{}
		for _, c := range p.Criteria {
			scores.Criteria = append(scores.Criteria, review.CriterionResult{
				Name: c.Name, Passed: c.Passed, Reason: c.Reason,
			})
		}
		results = append(results, report.GraderResult{
			GraderName:   p.Model,
			GraderType:   "review",
			Model:        p.Model,
			Scores:       scores,
			OverallScore: p.OverallScore,
			MaxScore:     p.MaxScore,
			Summary:      p.Summary,
			Issues:       p.Issues,
			Strengths:    p.Strengths,
		})
	}

	// Consolidated entry.
	consensusName := rd.Model
	if consensusName == "" {
		consensusName = "consensus"
	}
	scores := review.ReviewScores{}
	for _, c := range rd.Criteria {
		scores.Criteria = append(scores.Criteria, review.CriterionResult{
			Name: c.Name, Passed: c.Passed, Reason: c.Reason,
		})
	}
	results = append(results, report.GraderResult{
		GraderName:   consensusName,
		GraderType:   "review",
		Model:        rd.Model,
		Scores:       scores,
		OverallScore: rd.OverallScore,
		MaxScore:     rd.MaxScore,
		Summary:      rd.Summary,
		Issues:       rd.Issues,
		Strengths:    rd.Strengths,
		IsConsensus:  len(rd.PanelResults) > 0,
	})

	return results
}
