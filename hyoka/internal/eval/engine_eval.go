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

	debugPrefix := task.Prompt.ID + "/" + task.Config.Name
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

	// Snapshot home directory and CWD before eval so we can recover misplaced files after
	homeDir, _ := os.UserHomeDir()
	var preEvalHomeFiles map[string]bool
	if homeDir != "" {
		preEvalHomeFiles = snapshotDir(homeDir)
	}
	// Also snapshot CWD — agents may write files relative to the process working directory
	cwdDir, _ := os.Getwd()
	var preEvalCwdFiles map[string]bool
	if cwdDir != "" && cwdDir != homeDir && cwdDir != ws.Dir {
		preEvalCwdFiles = snapshotDir(cwdDir)
	}

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
	// First, recover any files the agent wrote to the home directory instead of the workspace.
	// The Copilot CLI sometimes creates files in ~ when the agent omits the path parameter.
	if homeDir != "" && preEvalHomeFiles != nil {
		recovered := recoverMisplacedFiles(homeDir, preEvalHomeFiles, genDir, debugPrefix)
		if recovered > 0 {
			lg.Info("Recovered misplaced files from home dir", "count", recovered)
		}
		// Post-recovery validation: flag anything recovery couldn't handle (#26)
		if remaining := ValidateWorkspaceContainment(homeDir, preEvalHomeFiles); len(remaining) > 0 {
			lg.Warn("Items still outside workspace after recovery (home)", "count", len(remaining), "items", remaining)
		}
	}
	// Also recover from CWD
	if cwdDir != "" && preEvalCwdFiles != nil {
		recovered := recoverMisplacedFiles(cwdDir, preEvalCwdFiles, genDir, debugPrefix)
		if recovered > 0 {
			lg.Info("Recovered misplaced files from CWD", "count", recovered)
		}
		if remaining := ValidateWorkspaceContainment(cwdDir, preEvalCwdFiles); len(remaining) > 0 {
			lg.Warn("Items still outside workspace after recovery (CWD)", "count", len(remaining), "items", remaining)
		}
	}

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

	// Pluggable grader execution (#136) — runs after generation, before review.
	if len(e.pluginGraders) > 0 && len(generatedFiles) > 0 {
		glg := logging.WithPhase(lg, "grading")

		applicable := graders.ApplicableGraders(e.pluginGraders, props)
		glg.Debug("Applicable graders", "total", len(e.pluginGraders), "applicable", len(applicable))

		if len(applicable) > 0 {
			instances, err := graders.InstantiateGraders(applicable)
			if err != nil {
				glg.Error("Failed to instantiate graders", "error", err)
			} else {
				input := graders.GraderInput{
					WorkspacePath: genWs.Dir,
				}
				// Populate ActionLog from the structured action timeline (#139)
				if result != nil && result.ActionTimeline != nil {
					input.ActionLog = result.ActionTimeline.ToGraderActionLog()
				}

				results := graders.RunGraders(ctx, instances, applicable, input)
				agg, aggErr := graders.AggregateResults(results)
				if aggErr != nil {
					glg.Error("Failed to aggregate grader results", "error", aggErr)
				} else {
					reportResults := make([]report.GraderResult, len(agg.Results))
					for i, r := range agg.Results {
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
						reportResults[i] = rr
					}
					evalReport.GraderResults = reportResults
					evalReport.ScoreBreakdown = report.BuildScoreBreakdown(reportResults)

					if !agg.Pass && !evalFailed {
						evalReport.Success = false
						if agg.GateFailed {
							evalReport.FailureReason = "gate grader(s) failed"
						}
					}

					glg.Info("Grader execution complete",
						"graders", len(results),
						"score", fmt.Sprintf("%.2f", agg.Score),
						"passed", agg.Pass,
						"gate_failed", agg.GateFailed)
					sendEvent(progress.EventToolComplete, fmt.Sprintf("Graders: %.0f%% (%d/%d passed)",
						agg.Score*100, countPassed(results), len(results)))
				}
			}
		}
	}

	// Code review — use panel reviewer if available, otherwise single reviewer
	// Uses its own independent timeout context (fixes issue #3).
	if !e.opts.SkipReview && len(generatedFiles) > 0 {
		reviewStart := time.Now()
		sendPhase(progress.PhaseReviewing)
		rlg := logging.WithPhase(lg, "review")

		// Create an isolated reviewer workspace with a copy of the generated
		// files. Reviewers operate on this copy and cannot modify the original
		// output in the report directory (#26).
		reviewWorkDir, err := NewReviewerWorkspace(ws.Dir)
		if err != nil {
			rlg.Warn("Reviewer workspace creation failed, using original", "error", err)
			reviewWorkDir = ws.Dir
		} else {
			defer os.RemoveAll(reviewWorkDir)
		}

		referenceDir := ""
		if task.Prompt.ReferenceAnswer != "" {
			referenceDir = task.Prompt.ReferenceAnswer
		}

		// Merge evaluation criteria (#30)
		evalCriteria := e.mergedCriteria(task.Prompt, props)

		// Create reviewer for this specific config using the factory (#92)
		var reviewer review.Reviewer
		var panelReviewer *review.PanelReviewer
		if e.reviewerFactory != nil {
			reviewer, panelReviewer, err = e.reviewerFactory(&task.Config)
			if err != nil {
				rlg.Warn("Reviewer creation failed, skipping review", "error", err)
			}
		}

		if panelReviewer != nil {
			models := panelReviewer.Models()
			rlg.Debug("Starting review panel")
			sendEvent(progress.EventToolStart, fmt.Sprintf("Review panel: %v", models))
			panel, consolidated, err := panelReviewer.ReviewPanel(ctx, task.Prompt.PromptText, reviewWorkDir, referenceDir, evalCriteria)
			if err != nil {
				rlg.Error("Review panel failed", "error", err)
				sendEvent(progress.EventReasoning, fmt.Sprintf("Review panel failed: %v", err))
			} else {
				evalReport.ReviewPanel = panel
				evalReport.Review = consolidated
				evalReport.GraderResults = report.GraderResultsFromReview(consolidated, panel)
				// With criteria-based scoring, success = all criteria passed
				if !evalFailed {
					evalReport.Success = consolidated.Scores.AllPassed()
				}
				sendEvent(progress.EventToolComplete, fmt.Sprintf("Review complete: %d/%d criteria passed", consolidated.OverallScore, consolidated.MaxScore))
				rlg.Debug("Review panel complete",
					"reviewers", len(panel),
					"score", consolidated.OverallScore,
					"max_score", consolidated.MaxScore)
			}
		} else if reviewer != nil {
			rlg.Debug("Starting single review session")
			sendEvent(progress.EventToolStart, "Single model review")
			reviewResult, err := reviewer.Review(ctx, task.Prompt.PromptText, reviewWorkDir, referenceDir, evalCriteria)
			if err != nil {
				rlg.Error("Code review failed", "error", err)
				sendEvent(progress.EventReasoning, fmt.Sprintf("Review failed: %v", err))
			} else {
				evalReport.Review = reviewResult
				evalReport.GraderResults = report.GraderResultsFromReview(reviewResult, nil)
				// With criteria-based scoring, success = all criteria passed
				if !evalFailed {
					evalReport.Success = reviewResult.Scores.AllPassed()
				}
				sendEvent(progress.EventToolComplete, fmt.Sprintf("Review complete: %d/%d criteria passed", reviewResult.OverallScore, reviewResult.MaxScore))
				rlg.Debug("Review complete",
					"score", reviewResult.OverallScore,
					"max_score", reviewResult.MaxScore)
			}
		}

		// Capture reviewed (annotated) files from the reviewer workspace
		reviewedFiles, err := readReviewedFiles(reviewWorkDir)
		if err == nil && len(reviewedFiles) > 0 {
			evalReport.ReviewedFiles = reviewedFiles
			rlg.Debug("Captured reviewed files", "count", len(reviewedFiles))
		}
		evalReport.ReviewDuration = time.Since(reviewStart).Seconds()
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
