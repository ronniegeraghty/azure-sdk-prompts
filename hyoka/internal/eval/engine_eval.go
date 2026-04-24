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
"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria"
"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria/graders"
"github.com/ronniegeraghty/hyoka/hyoka/internal/logging"
"github.com/ronniegeraghty/hyoka/hyoka/internal/progress"
"github.com/ronniegeraghty/hyoka/hyoka/internal/report"
"github.com/ronniegeraghty/hyoka/hyoka/internal/review"
"github.com/ronniegeraghty/hyoka/hyoka/internal/workspace"
)

func (e *Engine) runSingleEval(ctx context.Context, task EvalTask, runID string, sendPhase func(progress.Phase), sendEvent func(progress.EventType, string), sendRawEvent func(progress.ProgressEvent)) *report.EvalReport {
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
	evalReport.GuardrailMaxSessionActions = lim.maxSessionActions

	if len(task.Prompt.Tags) > 0 {
		evalReport.PromptMeta["tags"] = strings.Join(task.Prompt.Tags, ", ")
	}
	if task.Prompt.Group != "" {
		evalReport.PromptMeta["group"] = task.Prompt.Group
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

	// Snapshot starter-file sizes so guardrails can charge the agent only for
	// the bytes/files it actually contributed, not the starter project it was
	// handed (#565).
	starterSnapshot := snapshotStarterSizes(genDir, starterFiles)

	// Take a content-hashed snapshot of the workspace post-starter-copy.
	// Pairing this with a post-generation snapshot yields a WorkspaceDelta
	// describing exactly what the agent created/modified/deleted (#566).
	beforeSnap, snapErr := workspace.TakeSnapshot(genDir)
	if snapErr != nil {
		// Non-fatal: if we cannot snapshot, the eval still proceeds; delta
		// will simply be unavailable for graders and report consumers.
		lg.Warn("Workspace pre-snapshot failed; WorkspaceDelta will be unavailable",
			"error", snapErr)
		beforeSnap = nil
	}

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
		// If EvalResult returned an ErrorCategory, use it; otherwise infer
		if result != nil && result.ErrorCategory != "" {
			evalReport.ErrorCategory = result.ErrorCategory
			evalReport.Error = result.Error
			evalReport.ErrorDetails = result.ErrorDetails
			evalReport.FailureReason = result.Error // Use the error message as failure reason
		} else if genCtxErr == context.Canceled {
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

	// Compute the WorkspaceDelta now that the agent has finished writing to
	// genDir (#566). If the pre-snapshot failed we skip — graders nil-check.
	if beforeSnap != nil {
		afterSnap, snapErr := workspace.TakeSnapshot(genDir)
		if snapErr != nil {
			lg.Warn("Workspace post-snapshot failed; WorkspaceDelta unavailable",
				"error", snapErr)
		} else {
			evalReport.WorkspaceDelta = workspace.ComputeDelta(beforeSnap, afterSnap)
			lg.Debug("WorkspaceDelta captured",
				"new_files", evalReport.WorkspaceDelta.NewFileCount,
				"modified_files", evalReport.WorkspaceDelta.ModifiedFileCount,
				"deleted_files", evalReport.WorkspaceDelta.DeletedFileCount,
				"bytes_net", evalReport.WorkspaceDelta.BytesNet)
		}
	}

	// Graders now run on every eval, regardless of file count. The legacy
	// engine-level "no files generated" failure was removed in favor of the
	// configurable `output_check` grader (see hyoka/internal/graders/output_check_grader.go).
	// Evals that evaluate the agent's final response text rather than files
	// (e.g., planning, recommendations, explanations) benefit from this change.
	// The agent's final response is now threaded through GraderInput.AgentFinalResponse.

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
		resolved, err := tool.ResolveSkills(ctx, task.Config.Generator.Tools, "")
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
	// SkillGroups: enrich the flat SkillsLoaded list with parent linkage
	// from the post-validation tool topology (Plan 5.5 Option B). The
	// loaded set (SkillsLoaded) is authoritative for which skills the SDK
	// actually loaded; toolReport.Items provides the parent/kind for each
	// skill that was attempted at session start. Skills the SDK loaded
	// without a corresponding validator entry (rare, mostly stub paths)
	// land in SkillGroups with empty Parent/ParentKind.
	if result != nil && result.ToolReport != nil && len(env.SkillsLoaded) > 0 {
		topology := make(map[string]report.SkillLoadEntry, len(result.ToolReport.Items))
		for _, item := range result.ToolReport.Items {
			if item.Kind != progress.ToolKindSkill && item.Kind != progress.ToolKindMCP {
				continue
			}
			topology[item.Name] = report.SkillLoadEntry{
				Name:       item.Name,
				Kind:       item.Kind,
				Parent:     item.Parent,
				ParentKind: item.ParentKind,
			}
		}
		groups := make([]report.SkillLoadEntry, 0, len(env.SkillsLoaded))
		for _, s := range env.SkillsLoaded {
			if entry, ok := topology[s]; ok {
				groups = append(groups, entry)
			} else {
				groups = append(groups, report.SkillLoadEntry{Name: s, Kind: progress.ToolKindSkill})
			}
		}
		env.SkillGroups = groups
	}
	evalReport.Environment = env

	// Build tool availability summary: what was available vs actually used (#348).
	evalReport.ToolAvailability = buildToolAvailability(env, evalReport.SessionEvents)

	// Build SessionSetup from config and starter files (#219).
	setup := &report.SessionSetupEvent{
		Tools:        reportAvailableTools,
		StarterFiles: evalReport.StarterFiles,
	}
	// Prefer post-validation tool topology (carries parent linkage and
	// runtime status) when available — falls back to the raw config entries
	// for stub runs and any path that did not perform tool validation.
	if result != nil && result.ToolReport != nil {
		setup.MCPServers, setup.Skills = buildToolLoadResults(result.ToolReport, task.Config.Generator.Tools)
	} else {
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
				Kind:    "mcp",
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
				Kind:    "skill",
			})
		}
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

		// File count — HARD FAIL (#35). Counts agent-attributable files
		// (starter-aware via #565). Phase 3.5 (#566) dropped the byte-size
		// guardrail entirely; MaxFiles remains the only generator-side cap
		// with MaxTurns / MaxSessionActions covering session length.
		agentFileCount := computeAgentFileCount(generatedFiles, starterSnapshot)
		if agentFileCount > lim.maxFiles {
			reason := fmt.Sprintf("guardrail: agent file count %d exceeded limit of %d", agentFileCount, lim.maxFiles)
			evalReport.GuardrailAbortReason = reason
			evalReport.Error = reason
			evalReport.Success = false
			lg.Warn("Guardrail triggered", "reason", reason, "agent_files", agentFileCount, "max_files", lim.maxFiles)
		}
	}

	// Emit EventSessionDetails to signal that the generation phase is complete.
	// This lets the interactive renderer flip the Agent Attempt line to
	// "Completed" BEFORE graders take over the tail (#Bug-2-fix). Without this
	// event, agentComplete would only be triggered by the terminal event
	// (EventPassed/Failed) which arrives AFTER all graders have run, causing
	// the frozen agent row index to be stale when rewriteFrozenLine tries to
	// update it in place.
	turnCount := 0
	for _, ev := range evalReport.SessionEvents {
		if ev.Type == "assistant.message" {
			turnCount++
		}
	}
	// Compute cost from token usage
	cost := 0.0
	for _, ev := range evalReport.SessionEvents {
		if ev.Type == "assistant.usage" {
			cost += float64(ev.InputTokens+ev.OutputTokens) * 0.00001 // rough estimate
		}
	}
	sendRawEvent(progress.ProgressEvent{
		Type:      progress.EventSessionDetails,
		Files:     generatedFiles,
		Turns:     turnCount,
		ToolCalls: len(evalReport.ToolCalls),
		Cost:      cost,
	})

	// Unified grading pipeline (#625) — the single Bundle drives both
	// prompt-type entries (LLM review panel) and typed entries (file,
	// program, output_check, ...). A malformed grader file only fails THIS
	// eval if its file-level when: block would have matched these props
	// (Q4 deferred-error semantics).
	gradeStart := time.Now()
	glg := logging.WithPhase(lg, "grading")

	// Fail this eval (and only this eval) if a grader file relevant to
	// its properties failed to load.
	if bundleErr := e.graderBundle.MatchingErrors(props); bundleErr != nil {
		glg.Error("Grader bundle has errors matching this eval", "error", bundleErr)
		if !evalFailed {
			evalReport.Success = false
			evalReport.FailureReason = fmt.Sprintf("grader config error: %v", bundleErr)
			evalReport.Error = "grader config error"
			evalReport.ErrorDetails = bundleErr.Error()
			evalReport.ErrorCategory = "grader_config_error"
		}
		evalReport.ReviewDuration = time.Since(gradeStart).Seconds()
		// Skip grading entirely — we don't trust any partial result
		// when the bundle couldn't be fully parsed for this eval.
	} else {
		// Resolve applicable graders, then partition into prompt (review
		// panel) and typed (NewGrader) entries. promptMatched is
		// consumed indirectly via reviewBuckets/mergedCriteria on
		// graderInput; we only need typedMatched here.
		matched := e.matchedForEval(props)
		_, typedMatched := criteria.PartitionMatched(matched)

		// Collect all grader results across typed graders and AI review.
		var allGraderResults []graders.GraderResult

		// Build common grader input shared by all grader types.
		graderInput := graders.GraderInput{
			WorkspacePath:       genWs.Dir,
			OriginalPrompt:      task.Prompt.PromptText,
			EvalCriteria:        e.mergedCriteria(task.Prompt, props),
			EvalCriteriaBuckets: e.reviewBuckets(task.Prompt, props),
			WorkspaceDelta:      evalReport.WorkspaceDelta,
		}
		if task.Prompt.ReferenceAnswer != "" {
			graderInput.ReferenceDir = task.Prompt.ReferenceAnswer
		}
		if result != nil {
			if result.ActionTimeline != nil {
				graderInput.ActionLog = result.ActionTimeline.ToGraderActionLog()
			}
		graderInput.AgentFinalResponse = result.FinalResponse
	}

	// --- Phase 1: Typed graders (file, program, output_check, ...) ---
	if len(typedMatched) > 0 {
		typedConfigs := make([]graders.GraderConfig, 0, len(typedMatched))
		for _, m := range typedMatched {
			typedConfigs = append(typedConfigs, m.Entry.ToRuntimeConfig())
		}
		glg.Debug("Typed graders matched", "count", len(typedConfigs))
		instances, instErr := criteria.InstantiateGraders(typedConfigs)
		if instErr != nil {
			glg.Error("Failed to instantiate typed graders", "error", instErr)
		} else {
			hooks := buildGraderHooks(sendRawEvent)
			typedResults := criteria.RunGradersWithHooks(ctx, instances, typedConfigs, graderInput, hooks)
			allGraderResults = append(allGraderResults, typedResults...)
			glg.Debug("Typed graders complete", "count", len(typedResults))
		}
	}

	// --- Phase 2: AI review grader (runs alongside typed graders) ---
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

			// NOTE: Do NOT bracket this with sendEvent(EventToolStart)/
			// sendEvent(EventToolComplete). The grader Start/Complete
			// events emitted just below already convey this state for
			// the renderer; the dual emission disturbs the active tail
			// right around the grader handoff and produces duplicate
			// rendered rows for ai_review (see Tank Phase 1 issue (e)).
			if panelReviewer != nil {
				glg.Debug("Review panel models", "models", panelReviewer.Models())
			}

			emitGraderStart(sendRawEvent, reviewGrader)
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
			} else if reviewGrader.LastConsolidated != nil {
				glg.Debug("Review complete",
					"passed", reviewGrader.LastConsolidated.OverallScore,
					"max", reviewGrader.LastConsolidated.MaxScore)
			}
			// Apply default weight — review has weight 1.0, not a gate grader.
			if reviewResult.Weight == 0 {
				reviewResult.Weight = 1.0
			}
			emitGraderComplete(sendRawEvent, reviewGrader, reviewResult)
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
			// Pre-computed roll-ups (#schema_v3): site reads these
			// directly. Counts are taken off agg.Results, which
			// is the authoritative pre-conversion grader set —
			// one entry per grader, including prompt_review.
			// Total is the sum of all grader points across all graders;
			// graders with zero points count as 1 point for backward compat.
			evalReport.GradersTotal = countTotalPoints(agg.Results)
			evalReport.GradersPassed = countPassedPoints(agg.Results)

			if !agg.Pass && !evalFailed {
				evalReport.Success = false
				if evalReport.FailureReason == "" {
					evalReport.FailureReason = "one or more graders failed"
				}
			}

			glg.Info("Grader execution complete",
				"graders", len(allGraderResults),
				"score", fmt.Sprintf("%.2f", agg.Score),
				"passed", agg.Pass)
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

	// Populate file contents for generated files (Bug #2: site display).
	// Read each generated file from the workspace directory up to 1MB per file.
	evalReport.FileContents = readGeneratedFileContents(ws.Dir, evalReport.GeneratedFiles, lg)

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

// buildToolLoadResults converts a post-validation tool topology into the
// report shape used by SessionSetupEvent. The validator emits a flat list of
// leaves (skills, MCP servers, plugin children, skill_dir children) each
// carrying Parent/ParentKind back-pointers; we preserve that linkage on the
// report side so v3 JSON reports can render the same grouped view as the
// live progress display. Configured-but-unresolved entries from the raw
// config (e.g., a skill that failed to fetch and was filtered out before
// validation) are not duplicated — the toolReport is authoritative for what
// was actually attempted at session start.
func buildToolLoadResults(toolReport *tool.ToolLoadReport, configured []config.ToolEntry) (mcpServers, skills []report.ToolLoadResult) {
	if toolReport == nil {
		return nil, nil
	}
	// Map from (kind, name) -> details from the raw config so children we
	// emit can carry the same Details string the legacy "configured" rows
	// surfaced (e.g., MCP command lines, skill source). Only top-level
	// config entries are recorded here; plugin children inherit from the
	// plugin parent's details where appropriate.
	cfgDetails := make(map[string]string, len(configured))
	for _, e := range configured {
		key := e.ResolvedType() + ":" + e.Name
		switch e.ResolvedType() {
		case "mcp":
			d := e.Command
			if len(e.Args) > 0 {
				d += " " + strings.Join(e.Args, " ")
			}
			cfgDetails[key] = d
		case "skill":
			cfgDetails[key] = e.SkillSource()
		}
	}

	// Track which parents we've emitted so each container is recorded once.
	emittedParent := make(map[string]bool)

	for _, item := range toolReport.Items {
		if item.Parent != "" && !emittedParent[item.ParentKind+":"+item.Parent] {
			emittedParent[item.ParentKind+":"+item.Parent] = true
			parentRow := report.ToolLoadResult{
				Name: item.Parent,
				Kind: item.ParentKind, // "plugin" | "skill_dir"
				// Status omitted: parents have no runtime status, only
				// their children do (the plugin/skill_dir is a container).
			}
			// Plugin parents may have config-side Details (the package
			// reference), if recorded under a top-level plugin entry.
			if d, ok := cfgDetails["plugin:"+item.Parent]; ok {
				parentRow.Details = d
			}
			// Plugin parents land in the same group as their children.
			switch item.ParentKind {
			case progress.ToolParentKindPlugin:
				// Plugin children can be either skills or MCP servers, so
				// the plugin parent row is duplicated into both buckets
				// only when the plugin actually contributed children of
				// that kind. We resolve that lazily by inspecting child
				// kinds in a follow-up pass below.
				switch item.Kind {
				case progress.ToolKindMCP:
					mcpServers = append(mcpServers, parentRow)
				case progress.ToolKindSkill:
					skills = append(skills, parentRow)
				}
			case progress.ToolParentKindSkillDir:
				skills = append(skills, parentRow)
			}
		}

		row := report.ToolLoadResult{
			Name:       item.Name,
			Status:     item.Status,
			Error:      item.Reason,
			Kind:       item.Kind,
			Parent:     item.Parent,
			ParentKind: item.ParentKind,
		}
		// Carry config-side Details onto the leaf when the entry was
		// declared at the top level (no parent).
		if item.Parent == "" {
			if d, ok := cfgDetails[item.Kind+":"+item.Name]; ok {
				row.Details = d
			}
		}
		switch item.Kind {
		case progress.ToolKindMCP:
			mcpServers = append(mcpServers, row)
		case progress.ToolKindSkill:
			skills = append(skills, row)
		case progress.ToolKindPlugin:
			// Top-level plugin rows that have not yet been emitted as a
			// parent above (rare — usually they appear via children) are
			// captured into both buckets so consumers see the container.
			if !emittedParent[progress.ToolParentKindPlugin+":"+item.Name] {
				emittedParent[progress.ToolParentKindPlugin+":"+item.Name] = true
				skills = append(skills, row)
				mcpServers = append(mcpServers, row)
			}
		}
	}
	return mcpServers, skills
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

// countTotalPoints returns the total number of grader points across all graders.
// Graders with zero points are counted as having 1 point for backward compatibility
// with legacy graders that don't populate the Points slice.
func countTotalPoints(results []graders.GraderResult) int {
	total := 0
	for _, r := range results {
		if len(r.Points) > 0 {
			total += len(r.Points)
		} else {
			// Legacy grader with no Points: treat as 1 point
			total++
		}
	}
	return total
}

// countPassedPoints returns the number of passed grader points across all graders.
// For graders without Points, the grader's overall Pass field is used (1 if true, 0 if false).
func countPassedPoints(results []graders.GraderResult) int {
	passed := 0
	for _, r := range results {
		if len(r.Points) > 0 {
			for _, pt := range r.Points {
				if pt.Pass {
					passed++
				}
			}
		} else {
			// Legacy grader with no Points: use overall Pass
			if r.Pass {
				passed++
			}
		}
	}
	return passed
}

// readGeneratedFileContents reads the contents of generated files from the workspace.
// Files exceeding maxFileContentSize (1MB) are capped with a truncation marker.
// Binary files (detected via extension) are skipped with a marker message.
func readGeneratedFileContents(workspaceDir string, generatedFiles []string, lg *slog.Logger) map[string]string {
	const maxFileContentSize = 1024 * 1024 // 1MB cap per file
	binaryExtensions := map[string]bool{
		".png": true, ".jpg": true, ".jpeg": true, ".gif": true, ".bmp": true,
		".pdf": true, ".zip": true, ".tar": true, ".gz": true, ".7z": true,
		".exe": true, ".dll": true, ".so": true, ".dylib": true,
		".bin": true, ".dat": true, ".db": true, ".sqlite": true,
	}

	contents := make(map[string]string, len(generatedFiles))
	for _, relPath := range generatedFiles {
		fullPath := filepath.Join(workspaceDir, relPath)

		// Check if file is binary by extension
		ext := strings.ToLower(filepath.Ext(relPath))
		if binaryExtensions[ext] {
			contents[relPath] = "[Binary file — not displayed]"
			continue
		}

		// Read file with size check
		info, err := os.Stat(fullPath)
		if err != nil {
			lg.Debug("Failed to stat generated file for content read", "file", relPath, "error", err)
			contents[relPath] = fmt.Sprintf("[Error reading file: %v]", err)
			continue
		}

		if info.Size() > maxFileContentSize {
			contents[relPath] = fmt.Sprintf("[File too large to display (%d bytes) — view on disk at %s]",
				info.Size(), fullPath)
			continue
		}

		data, err := os.ReadFile(fullPath)
		if err != nil {
			lg.Debug("Failed to read generated file contents", "file", relPath, "error", err)
			contents[relPath] = fmt.Sprintf("[Error reading file: %v]", err)
			continue
		}

		contents[relPath] = string(data)
	}

	return contents
}

// convertGraderResults converts internal grader results to report format.
//
// Schema v3 (this code path) emits exactly one row per grader, including
// the prompt_review grader: per-criterion outcomes ride on the row's
// Points slice (populated by the grader itself in Phase 2), while the
// existing ReviewDetails struct still carries PanelResults / Criteria for
// backward-compat with the static Markdown/HTML report templates.
//
// The legacy "expand to one entry per panel member + one consensus row"
// behaviour was removed in Phase 5 — its expanded shape was the structural
// cause of the "all panel members passed but rows render red" bug, since
// site renderers had no way to distinguish a panel-member entry (no Points,
// no Pass) from a real grader failure. Old v2 reports on disk keep the
// expanded shape on read; only freshly written v3 reports use this path.
func convertGraderResults(results []graders.GraderResult) []report.GraderResult {
	var reportResults []report.GraderResult
	for _, r := range results {
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
		if len(r.Points) > 0 {
			rr.Points = make([]report.GraderPoint, len(r.Points))
			for j, p := range r.Points {
				rr.Points[j] = report.GraderPoint{Name: p.Name, Pass: p.Pass, Message: p.Message}
			}
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
		// prompt_review graders: re-shape the grader's name to a stable
		// "ai_review" identifier (the panel-member-expansion code used to
		// produce one row per model — site filters keyed off "ai_review"
		// or "review" type for years) and copy the existing ReviewDetails
		// struct verbatim so static Markdown/HTML templates keep working.
		if r.Kind == graders.KindPromptReview && r.ReviewDetails != nil {
			rd := r.ReviewDetails
			rr.GraderName = "ai_review"
			rr.GraderType = "prompt_review"
			rr.Model = rd.Model
			rr.OverallScore = rd.OverallScore
			rr.MaxScore = rd.MaxScore
			rr.Issues = rd.Issues
			rr.Strengths = rd.Strengths
			rr.IsConsensus = len(rd.PanelResults) > 0
			scores := review.ReviewScores{}
			for _, c := range rd.Criteria {
				scores.Criteria = append(scores.Criteria, review.CriterionResult{
					Name: c.Name, Passed: c.Passed, Reason: c.Reason,
				})
			}
			rr.Scores = scores
			// ReviewDetails struct is preserved for templates that walk
			// PanelResults / Criteria directly.
			panelDetails := make([]report.ReviewGraderPanelEntry, 0, len(rd.PanelResults))
			for _, p := range rd.PanelResults {
				crit := make([]report.ReviewGraderCriterion, 0, len(p.Criteria))
				for _, c := range p.Criteria {
					crit = append(crit, report.ReviewGraderCriterion{
						Name: c.Name, Passed: c.Passed, Reason: c.Reason,
					})
				}
				panelDetails = append(panelDetails, report.ReviewGraderPanelEntry{
					Model:        p.Model,
					OverallScore: p.OverallScore,
					MaxScore:     p.MaxScore,
					Summary:      p.Summary,
					Issues:       p.Issues,
					Strengths:    p.Strengths,
					Criteria:     crit,
				})
			}
			critDetails := make([]report.ReviewGraderCriterion, 0, len(rd.Criteria))
			for _, c := range rd.Criteria {
				critDetails = append(critDetails, report.ReviewGraderCriterion{
					Name: c.Name, Passed: c.Passed, Reason: c.Reason,
				})
			}
			rr.ReviewDetails = &report.ReviewGraderDetail{
				Model:        rd.Model,
				PanelResults: panelDetails,
				Criteria:     critDetails,
				OverallScore: rd.OverallScore,
				MaxScore:     rd.MaxScore,
				Issues:       rd.Issues,
				Strengths:    rd.Strengths,
				IsConsensus:  len(rd.PanelResults) > 0,
			}
		}
		reportResults = append(reportResults, rr)
	}
	return reportResults
}

// buildGraderHooks wires criteria.GraderHooks to a raw progress emitter so
// each grader run in RunGradersWithHooks produces GraderStart/GraderComplete
// events. A nil sendRawEvent means "no reporter wired" — in that case we
// return zero-value hooks and RunGradersWithHooks will skip emission.
func buildGraderHooks(sendRawEvent func(progress.ProgressEvent)) criteria.GraderHooks {
if sendRawEvent == nil {
return criteria.GraderHooks{}
}
return criteria.GraderHooks{
OnStart: func(g graders.Grader) {
emitGraderStart(sendRawEvent, g)
},
OnComplete: func(g graders.Grader, result graders.GraderResult) {
emitGraderComplete(sendRawEvent, g, result)
},
}
}

// emitGraderStart sends a GraderStart progress event. Safe to call with a nil
// sender (no-op) so callers in the review path don't need to guard.
func emitGraderStart(sendRawEvent func(progress.ProgressEvent), g graders.Grader) {
if sendRawEvent == nil || g == nil {
return
}
sendRawEvent(progress.ProgressEvent{
Type:       progress.EventGraderStart,
GraderID:   g.Name(),
GraderKind: g.Kind(),
})
}

// emitGraderComplete sends a GraderComplete progress event. Score is only
// populated for grader kinds that produce a meaningful numeric score
// (prompt_review and prompt LLM-as-judge). Output-check / file / program /
// behavior graders are binary pass/fail, so Score stays nil for them to
// avoid misleading renders like "pass (0/10)".
func emitGraderComplete(sendRawEvent func(progress.ProgressEvent), g graders.Grader, result graders.GraderResult) {
if sendRawEvent == nil || g == nil {
return
}
outcome := progress.GraderResultFail
if result.Pass {
outcome = progress.GraderResultPass
}
var scorePtr *float64
switch g.Kind() {
case graders.KindPromptReview, graders.KindPrompt:
s := result.Score
scorePtr = &s
}
var points []progress.GraderPoint
if len(result.Points) > 0 {
points = make([]progress.GraderPoint, len(result.Points))
for i, p := range result.Points {
points[i] = progress.GraderPoint{Name: p.Name, Pass: p.Pass, Message: p.Message}
}
}
sendRawEvent(progress.ProgressEvent{
Type:       progress.EventGraderComplete,
GraderID:   g.Name(),
GraderKind: g.Kind(),
Result:     outcome,
Score:      scorePtr,
Message:    result.Message,
Points:     points,
})
}
