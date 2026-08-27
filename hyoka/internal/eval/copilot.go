package eval

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	copilot "github.com/github/copilot-sdk/go"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/config/tool"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/copilotevent"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/copilotperm"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/logging"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/pidfile"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/process"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/progress"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/prompt"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/report"
)

// CopilotPromptRunner uses the Copilot SDK to run real evaluations.
type CopilotPromptRunner struct {
	clientOpts        *copilot.ClientOptions
	allowCloud        bool
	maxSessionActions int
	maxTurns          int
	maxFiles          int
	sessionTimeout    time.Duration
	progressFn        progress.ProgressFunc

	// Per-eval resolved limits (set by SetLimitsForEval before each Run call).
	// Protected by evalLimitsMu for concurrent eval safety (#bugfix-maxturns).
	evalLimitsMu          sync.RWMutex
	evalMaxTurns          int
	evalMaxFiles          int
	evalMaxSessionActions int
}

// SetProgressFunc registers a callback for live progress updates.
func (e *CopilotPromptRunner) SetProgressFunc(fn progress.ProgressFunc) {
	e.progressFn = fn
}

// SetSessionTimeout configures the maximum duration for a single generation
// SendAndWait call. Zero means use the default (10 minutes). Per-prompt
// Timeout frontmatter still overrides this value.
func (e *CopilotPromptRunner) SetSessionTimeout(d time.Duration) {
	e.sessionTimeout = d
}

// SetLimitsForEval updates the per-eval resolved limits before each Run call.
// These values override the CLI-level defaults during real-time enforcement.
// Safe for concurrent calls — the engine may run multiple evals in parallel
// against the same runner instance (#bugfix-maxturns).
func (e *CopilotPromptRunner) SetLimitsForEval(maxTurns, maxFiles, maxSessionActions int) {
	e.evalLimitsMu.Lock()
	e.evalMaxTurns = maxTurns
	e.evalMaxFiles = maxFiles
	e.evalMaxSessionActions = maxSessionActions
	e.evalLimitsMu.Unlock()
}

// PromptRunnerOptions configures the CopilotPromptRunner.
type PromptRunnerOptions struct {
	// GitHubToken for SDK authentication (optional; falls back to logged-in user).
	GitHubToken string
	// CLIPath overrides the Copilot CLI executable path.
	CLIPath string
	// AllowCloud permits agent output to provision real cloud resources (#36).
	AllowCloud bool
	// MaxSessionActions limits the total number of actions (reasoning, message,
	// tool execution start) during generation. When reached, the session context
	// is cancelled to stop the run immediately.
	MaxSessionActions int
	// MaxTurns limits the number of assistant turns during generation (#347).
	MaxTurns int
	// MaxFiles limits the number of files created during generation (#347).
	MaxFiles int
}

// NewCopilotPromptRunner creates a new evaluator backed by the Copilot SDK.
func NewCopilotPromptRunner(opts PromptRunnerOptions) *CopilotPromptRunner {
	clientOpts := BuildBaseClientOpts()
	if opts.GitHubToken != "" {
		clientOpts.GitHubToken = opts.GitHubToken
	}
	if opts.CLIPath != "" {
		clientOpts.Connection = copilot.StdioConnection{Path: opts.CLIPath}
	}
	return &CopilotPromptRunner{
		clientOpts:        clientOpts,
		allowCloud:        opts.AllowCloud,
		maxSessionActions: opts.MaxSessionActions,
		maxTurns:          opts.MaxTurns,
		maxFiles:          opts.MaxFiles,
	}
}

// Evaluate runs a prompt through a real Copilot session and returns generated files and events.
func (e *CopilotPromptRunner) Run(ctx context.Context, p *prompt.Prompt, cfg *config.ToolConfig, workDir string) (*EvalResult, error) {
	// Starter files are copied by the engine before Evaluate is called (#127).

	// Build session config from tool config
	// Create isolated config directory to prevent user-level skills from
	// leaking into the eval session (#21). Only skills explicitly listed
	// in the eval config's SkillDirectories are loaded.
	configDir, err := NewIsolatedConfigDir()
	if err != nil {
		return nil, fmt.Errorf("creating isolated config dir: %w", err)
	}

	// Pre-session tool validation (WU-1): resolve every declared plugin,
	// skill directory, and MCP server BEFORE the copilot client is started.
	// A missing plugin, missing skill path, or empty skill_dir aborts the
	// eval with error_category=tool_load_failure. Progress events are emitted
	// here so the interactive renderer can show the Tools block before the
	// Agent Attempt header — buildSessionConfigForEval consumes the
	// resulting report without re-resolving. Validation runs BEFORE
	// client.Start so a missing copilot binary or auth failure cannot
	// mask a tool-load failure (the hard-fail contract).
	taggedEmit := e.buildTaggedEmit(cfg, p.ID+"/"+cfg.Name, p.ID)
	toolReport, toolErr := tool.ValidateAndExpand(ctx, tool.ValidationInput{
		GeneratorTools: cfgGeneratorTools(cfg),
		// Reviewer tools are validated separately in cmd/run.go per-config
		// (WU-2) so missing reviewer skills fail fast there; including them
		// here would double-validate and couple engine runs to CLI flags.
		ReviewerTools: nil,
		ConfigDir:     configDir,
		PluginsDir:    config.ResolvePluginsDir(),
		Emit:          taggedEmit,
	})
	if toolErr != nil {
		_ = os.RemoveAll(configDir)
		return &EvalResult{
			Success: false,
			// toolErr.Error() is the multi-line summary produced by
			// tool.SummarizeToolLoadErrors — every failed tool, not just
			// the first. Surface it verbatim so operators can fix all
			// broken tools in one pass.
			Error:         "tool_load_failure:\n" + toolErr.Error(),
			ErrorDetails:  toolErr.Error(),
			ErrorCategory: "tool_load_failure",
		}, fmt.Errorf("tool load failure: %w", toolErr)
	}
	defer os.RemoveAll(configDir)

	// Create Copilot client
	opts := *e.clientOpts
	opts.WorkingDirectory = workDir
	// Enrich env with prompt/config metadata for this specific eval (#70).
	opts.Env = process.HyokaEvalEnv(p.ID, cfg.Name)
	client := copilot.NewClient(&opts)

	if err := client.Start(ctx); err != nil {
		return &EvalResult{
			Error:        fmt.Sprintf("copilot client start failed: %v", err),
			ErrorDetails: err.Error(),
		}, fmt.Errorf("starting copilot client: %w", err)
	}
	// The SDK does not expose the child PID, so we discover it by
	// scanning for direct child processes whose name contains "copilot".
	// The PID is written to a file so the clean command can find orphaned
	// processes even if hyoka crashes.
	var trackedPIDs []int
	for _, cpid := range process.FindChildCopilotPIDs() {
		if err := pidfile.Write(pidfile.Info{PID: cpid, PromptID: p.ID, Config: cfg.Name}); err != nil {
			slog.Debug("failed to write PID file", "pid", cpid, "error", err)
		} else {
			trackedPIDs = append(trackedPIDs, cpid)
		}
	}

	// Track session ID for cleanup — set after CreateSession.
	var sessionID string

	// buildCleanupFn returns a function that deletes session state and stops
	// the client. It is stored in EvalResult.CleanupFn so the engine can call
	// it AFTER copying generated files out of the workspace (#261). Previously,
	// DeleteSession was deferred here, which removed workspace artifacts before
	// the engine could copy them — causing baseline configs (which lack MCP
	// server processes that keep files alive) to report zero generated files.
	buildCleanupFn := func() func() {
		return func() {
			if sessionID != "" {
				deleteCtx, deleteCancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer deleteCancel()
				if err := client.DeleteSession(deleteCtx, sessionID); err != nil {
					slog.Debug("session delete failed, session-state may remain",
						"sessionID", sessionID, "error", err)
				}
			}
			done := make(chan struct{})
			go func() { client.Stop(); close(done) }()
			select {
			case <-done:
			case <-time.After(10 * time.Second):
				client.ForceStop()
			}
			// Remove PID files for processes we tracked.
			for _, cpid := range trackedPIDs {
				pidfile.Remove(cpid)
			}
		}
	}

	// Safety net: if we return early (error paths that don't attach CleanupFn
	// to a result), ensure the client is stopped and PIDs cleaned up.
	var cleanupCalled bool
	defer func() {
		if !cleanupCalled {
			buildCleanupFn()()
		}
	}()

	// Build session config from tool config. configDir, taggedEmit and
	// toolReport were resolved before client.Start above (pre-session
	// hard-fail contract); reuse them here.
	sessionCfg := e.buildSessionConfigForEval(ctx, cfg, workDir, configDir, mergePromptProperties(p), p.ID+"/"+cfg.Name, p.ID, toolReport)

	// Subscribe to events with detailed capture and debug logging.
	// This MUST be set before CreateSession — the SDK reads OnEvent during
	// session creation and won't pick up a callback assigned afterwards.
	var events []copilot.SessionEvent
	var sessionRecords []report.SessionEventRecord
	var mu sync.Mutex
	var toolTracker copilotevent.ToolTracker
	debugPrefix := p.ID + "/" + cfg.Name
	// Structured logger for this eval session (#42)
	lg := logging.EvalLogger(p.ID, cfg.Name, "generation", 0)

	// Capture turn counter for expanded events
	var turnCounter int
	var actionCounter int
	var fileCounter int // tracks files created during session (#347)

	// Mid-generation action limit. Create a cancellable child context
	// so the OnEvent callback can stop runaway sessions in real time.
	genCtx, genCancel := context.WithCancel(ctx)
	defer genCancel()
	var actionLimitHit bool
	var turnLimitHit bool
	var fileLimitHit bool

	// Resolve effective limits for real-time enforcement.
	// Prefer per-eval resolved values (set by engine via SetLimitsForEval),
	// fall back to CLI defaults from runner construction, then hardcoded defaults.
	e.evalLimitsMu.RLock()
	maxTurnsLimit := e.evalMaxTurns
	if maxTurnsLimit <= 0 {
		maxTurnsLimit = e.maxTurns
	}
	if maxTurnsLimit <= 0 {
		maxTurnsLimit = 25
	}

	maxFilesLimit := e.evalMaxFiles
	if maxFilesLimit <= 0 {
		maxFilesLimit = e.maxFiles
	}
	if maxFilesLimit <= 0 {
		maxFilesLimit = 50
	}

	maxSessionActionsLimit := e.evalMaxSessionActions
	if maxSessionActionsLimit <= 0 {
		maxSessionActionsLimit = e.maxSessionActions
	}
	e.evalLimitsMu.RUnlock()

	// Build expected tool sets for verification after session creation (#347)
	expectedMCPServers := make(map[string]bool)
	for _, entry := range cfg.Generator.Tools {
		if entry.ResolvedType() == "mcp" {
			expectedMCPServers[entry.Name] = true
		}
	}
	// verifier accumulates SDK-reported tool loads and produces exactly one
	// bulk EventToolsVerified per eval once every configured kind has fired.
	// Emission is gated on a non-nil progressFn; the slice is built under the
	// OnEvent mutex below but dispatched after unlock.
	verifier := newToolVerifier(sessionCfg.SkillDirectories, expectedMCPServers)

	sessionCfg.OnEvent = func(event copilot.SessionEvent) {
		mu.Lock()
		events = append(events, event)
		details := copilotevent.Extract(event)
		toolTracker.Enrich(event, &details)
		var verifiedTools []progress.ToolStatus

		// Build serializable event record
		rec := report.SessionEventRecord{
			Type: string(event.Type()),
		}
		if details.ToolName != nil {
			rec.ToolName = *details.ToolName
		}
		if details.Content != nil {
			rec.Content = *details.Content
		}
		if details.Arguments != nil {
			if argsBytes, err := json.Marshal(details.Arguments); err == nil {
				rec.ToolArgs = string(argsBytes)
			}
		}
		if details.Result != nil {
			if details.Result.DetailedContent != nil {
				rec.ToolResult = *details.Result.DetailedContent
			} else if details.Result.Content != nil {
				rec.ToolResult = *details.Result.Content
			}
		}
		if details.Error != nil {
			if details.Error.ErrorClass != nil {
				rec.Error = details.Error.ErrorClass.Message
			} else if details.Error.String != nil {
				rec.Error = *details.Error.String
			}
		}
		if details.Success != nil {
			rec.ToolSuccess = details.Success
		}
		if details.Duration != nil {
			rec.Duration = *details.Duration
		}
		if details.MCPServerName != nil {
			rec.MCPServerName = *details.MCPServerName
		}
		if details.MCPToolName != nil {
			rec.MCPToolName = *details.MCPToolName
		}
		if details.Path != nil {
			rec.FilePath = *details.Path
		}

		// Expanded event fields
		switch event.Type() {
		case copilot.SessionEventTypeAssistantTurnStart:
			turnCounter++
			rec.TurnNumber = turnCounter
			lg.Info("Turn started", "turn", turnCounter)
			// Tool loading MUST be complete by first turn start — the SDK won't
			// begin generation until tools are loaded or definitively failed.
			// Signal the verifier so postSessionToolVerification doesn't wait
			// forever for events that will never arrive (#347 / Option A).
			verifier.onSessionReady()
			if e.progressFn != nil {
				if t := verifier.emitIfReady(); t != nil {
					verifiedTools = t
				}
			}
			// Real-time turn limit enforcement (#347)
			if maxTurnsLimit > 0 && turnCounter > maxTurnsLimit && !turnLimitHit {
				turnLimitHit = true
				lg.Warn("Turn limit reached, cancelling session",
					"turns", turnCounter, "max_turns", maxTurnsLimit)
				genCancel()
			}
		case copilot.SessionEventTypeAssistantTurnEnd:
			rec.TurnNumber = turnCounter
			lg.Info("Turn ended", "turn", turnCounter)
		case copilot.SessionEventTypeAssistantReasoning:
			actionCounter++
			if maxSessionActionsLimit > 0 && actionCounter > maxSessionActionsLimit && !actionLimitHit {
				actionLimitHit = true
				lg.Warn("Action limit reached, cancelling session", "actions", actionCounter, "max_session_actions", maxSessionActionsLimit)
				genCancel()
			}
			// Content already captured above
		case copilot.SessionEventTypeAssistantIntent:
			if details.Intent != nil {
				rec.Intent = *details.Intent
			}
		case copilot.SessionEventTypeAssistantUsage:
			if details.InputTokens != nil {
				rec.InputTokens = int(*details.InputTokens)
			}
			if details.OutputTokens != nil {
				rec.OutputTokens = int(*details.OutputTokens)
			}
		case copilot.SessionEventTypeSessionWorkspaceFileChanged:
			if details.Operation != nil {
				rec.FileOperation = *details.Operation
				// Real-time file count enforcement (#347)
				if *details.Operation == "create" {
					fileCounter++
					if maxFilesLimit > 0 && fileCounter > maxFilesLimit && !fileLimitHit {
						fileLimitHit = true
						lg.Warn("File limit reached, cancelling session",
							"files", fileCounter, "max_files", maxFilesLimit)
						genCancel()
					}
				}
			}
		case copilot.SessionEventTypeCommandExecute:
			if details.Command != nil {
				rec.CommandText = *details.Command
			}
		case copilot.SessionEventTypeSkillInvoked:
			if details.SkillName != nil {
				rec.SkillName = *details.SkillName
			}
		case copilot.SessionEventTypeExternalToolRequested, copilot.SessionEventTypeExternalToolCompleted:
			if details.ToolName != nil {
				rec.ToolName = *details.ToolName
			}
		case copilot.SessionEventTypeSessionTruncation:
			rec.IsTruncation = true
			lg.Warn("Context truncated")
		case copilot.SessionEventTypeSessionCompactionStart:
			lg.Info("Context compaction started")
		case copilot.SessionEventTypeSessionCompactionComplete:
			lg.Info("Context compaction complete")
		case copilot.SessionEventTypeSessionWarning:
			if details.Message != nil {
				rec.WarningText = *details.Message
				lg.Warn("Session warning", "message", *details.Message)
			}
		case copilot.SessionEventTypeAbort:
			lg.Error("Session aborted")
		case copilot.SessionEventTypePermissionRequested:
			tn, tc := "", ""
			if details.ToolName != nil {
				tn = *details.ToolName
			}
			if details.ToolCallID != nil {
				tc = *details.ToolCallID
			}
			lg.Debug("Permission requested", "toolName", tn, "toolCallID", tc)
		case copilot.SessionEventTypePermissionCompleted:
			tn, tc, rsn, ern, msg := "", "", "", "", ""
			if details.ToolName != nil {
				tn = *details.ToolName
			}
			if details.ToolCallID != nil {
				tc = *details.ToolCallID
			}
			if details.Reason != nil {
				rsn = *details.Reason
			}
			if details.ErrorReason != nil {
				ern = *details.ErrorReason
			}
			if details.Message != nil {
				msg = *details.Message
			}
			lg.Debug("Permission completed", "toolName", tn, "toolCallID", tc, "reason", rsn, "errorReason", ern, "message", msg)
		case copilot.SessionEventTypeSessionSkillsLoaded:
			names := make([]string, 0, len(details.Skills))
			for _, s := range details.Skills {
				names = append(names, s.Name)
			}
			if len(names) > 0 {
				rec.Content = strings.Join(names, ", ")
				lg.Info("Skills loaded", "skills", rec.Content)
			} else if tool.CountSkills(sessionCfg.SkillDirectories) > 0 {
				lg.Warn("No skills loaded despite configured skill directories",
					"expected_dirs", len(sessionCfg.SkillDirectories))
			}
			verifier.onSkillsLoaded(names)
			if e.progressFn != nil {
				if t := verifier.emitIfReady(); t != nil {
					verifiedTools = t
				}
			}
		case copilot.SessionEventTypeSessionMCPServersLoaded:
			names := make([]string, 0, len(details.Servers))
			loadedNames := make(map[string]bool, len(details.Servers))
			for _, s := range details.Servers {
				names = append(names, s.Name)
				loadedNames[s.Name] = true
			}
			if len(names) > 0 {
				rec.Content = strings.Join(names, ", ")
				lg.Info("MCP servers loaded", "servers", rec.Content)
				// Verify all expected MCP servers loaded (#347)
				for expected := range expectedMCPServers {
					if !loadedNames[expected] {
						lg.Warn("Expected MCP server not loaded",
							"server", expected, "loaded", names)
					}
				}
			} else if len(expectedMCPServers) > 0 {
				lg.Warn("No MCP servers loaded despite configuration",
					"expected", len(expectedMCPServers))
			}
			verifier.onMCPLoaded(names)
			if e.progressFn != nil {
				if t := verifier.emitIfReady(); t != nil {
					verifiedTools = t
				}
			}
		case copilot.SessionEventTypeSessionToolsUpdated:
			lg.Info("Tools updated")
		case copilot.SessionEventTypeSubagentCompleted:
			if details.ToolCallID != nil {
				rec.SubagentID = *details.ToolCallID
			}
		case copilot.SessionEventTypeSubagentFailed:
			if details.ToolCallID != nil {
				rec.SubagentID = *details.ToolCallID
			}
		}

		sessionRecords = append(sessionRecords, rec)
		mu.Unlock()

		// Emit the bulk tool-verification event outside the lock (matches the
		// rest of the progress-forwarding in this handler; avoids any risk of
		// deadlock with renderers that take their own locks in progressFn).
		if verifiedTools != nil && e.progressFn != nil {
			e.progressFn(progress.ProgressEvent{
				EvalID:     debugPrefix,
				PromptID:   p.ID,
				ConfigName: cfg.Name,
				Type:       progress.EventToolsVerified,
				Tools:      verifiedTools,
			})
		}

		// Forward progress events to display
		if e.progressFn != nil {
			evalID := debugPrefix
			switch event.Type() {
			case copilot.SessionEventTypeToolExecutionStart:
				toolName := ""
				if details.ToolName != nil {
					toolName = *details.ToolName
				}
				if isFileWriteTool(toolName) {
					arg := toolArgSummary(event)
					e.progressFn(progress.ProgressEvent{
						EvalID: evalID, PromptID: p.ID, ConfigName: cfg.Name,
						Type:    progress.EventWritingFile,
						Message: toolName + " → " + arg,
					})
				} else {
					arg := toolArgSummary(event)
					msg := toolName
					if arg != "" {
						msg = toolName + " → " + arg
					}
					e.progressFn(progress.ProgressEvent{
						EvalID: evalID, PromptID: p.ID, ConfigName: cfg.Name,
						Type:    progress.EventToolStart,
						Message: msg,
					})
				}
			case copilot.SessionEventTypeToolExecutionComplete:
				toolName := ""
				if details.ToolName != nil {
					toolName = *details.ToolName
				}
				result := ""
				if details.Result != nil && details.Result.Content != nil {
					result = truncateStr(*details.Result.Content, 60)
				}
				msg := toolName
				if result != "" {
					msg = toolName + " → " + result
				}
				e.progressFn(progress.ProgressEvent{
					EvalID: evalID, PromptID: p.ID, ConfigName: cfg.Name,
					Type:    progress.EventToolComplete,
					Message: msg,
				})
			case copilot.SessionEventTypeAssistantMessage:
				content := ""
				if details.Content != nil {
					content = *details.Content
				}
				if content != "" {
					summary := truncateStr(content, 80)
					e.progressFn(progress.ProgressEvent{
						EvalID: evalID, PromptID: p.ID, ConfigName: cfg.Name,
						Type:    progress.EventReasoning,
						Message: summary,
					})
				}
			case copilot.SessionEventTypeAssistantTurnStart:
				e.progressFn(progress.ProgressEvent{
					EvalID: evalID, PromptID: p.ID, ConfigName: cfg.Name,
					Type:    progress.EventReasoning,
					Message: fmt.Sprintf("Turn %d started", turnCounter),
				})
			case copilot.SessionEventTypeSessionTruncation:
				e.progressFn(progress.ProgressEvent{
					EvalID: evalID, PromptID: p.ID, ConfigName: cfg.Name,
					Type:    progress.EventReasoning,
					Message: "⚠ Context truncated",
				})
			}
		}

		// Debug logging — all event types (slog.Debug is a no-op at higher levels)
		switch event.Type() {
		case copilot.SessionEventTypeToolExecutionStart:
			actionCounter++
			if maxSessionActionsLimit > 0 && actionCounter > maxSessionActionsLimit && !actionLimitHit {
				actionLimitHit = true
				lg.Warn("Action limit reached, cancelling session", "actions", actionCounter, "max_session_actions", maxSessionActionsLimit)
				genCancel()
			}
			toolName := ""
			if details.ToolName != nil {
				toolName = *details.ToolName
			}
			lg.Debug("Tool start", "tool", toolName)
		case copilot.SessionEventTypeToolExecutionComplete:
			toolName := ""
			if details.ToolName != nil {
				toolName = *details.ToolName
			}
			content := ""
			if details.Content != nil {
				content = truncateStr(*details.Content, 200)
			}
			lg.Debug("Tool done", "tool", toolName, "result", content)
		case copilot.SessionEventTypeAssistantMessage:
			actionCounter++
			if maxSessionActionsLimit > 0 && actionCounter > maxSessionActionsLimit && !actionLimitHit {
				actionLimitHit = true
				lg.Warn("Action limit reached, cancelling session", "actions", actionCounter, "max_session_actions", maxSessionActionsLimit)
				genCancel()
			}
			content := ""
			if details.Content != nil {
				content = *details.Content
			}
			if content != "" {
				if summary := detectFileCreation(content); summary != "" {
					lg.Debug("Assistant creating file", "summary", summary)
				} else {
					lg.Debug("Assistant message", "content", truncateStr(content, 200))
				}
			}
		case copilot.SessionEventTypeSessionError:
			content := ""
			if details.Content != nil {
				content = *details.Content
			}
			lg.Debug("Session error", "content", content)
		case copilot.SessionEventTypeAssistantTurnStart:
			lg.Debug("Turn started", "turn", turnCounter)
		case copilot.SessionEventTypeAssistantTurnEnd:
			lg.Debug("Turn ended", "turn", turnCounter)
		case copilot.SessionEventTypeAssistantUsage:
			in, out := 0, 0
			if details.InputTokens != nil {
				in = int(*details.InputTokens)
			}
			if details.OutputTokens != nil {
				out = int(*details.OutputTokens)
			}
			lg.Debug("Token usage", "input_tokens", in, "output_tokens", out)
		case copilot.SessionEventTypeSessionTruncation:
			lg.Debug("Context truncated")
		case copilot.SessionEventTypeSkillInvoked:
			name := ""
			if details.SkillName != nil {
				name = *details.SkillName
			}
			lg.Debug("Skill invoked", "skill", name)
		case copilot.SessionEventTypeSubagentCompleted, copilot.SessionEventTypeSubagentFailed:
			lg.Debug("Subagent event", "type", string(event.Type()))
		default:
			content := ""
			if details.Content != nil {
				content = truncateStr(*details.Content, 100)
			}
			lg.Debug("SDK event", "type", string(event.Type()), "content", content)
		}
	}

	slog.Info("Creating Copilot session",
		"model", cfg.Generator.Model,
		"skill_dirs", len(sessionCfg.SkillDirectories),
		"skills", tool.CountSkills(sessionCfg.SkillDirectories),
		"mcp_servers", len(sessionCfg.MCPServers),
		"work_dir", workDir,
	)
	session, err := client.CreateSession(genCtx, sessionCfg)
	if err != nil {
		return &EvalResult{
			Error:        fmt.Sprintf("session creation failed: %v", err),
			ErrorDetails: err.Error(),
		}, fmt.Errorf("creating session: %w", err)
	}
	sessionID = session.SessionID

	// Send the prompt
	if e.progressFn != nil {
		e.progressFn(progress.ProgressEvent{
			EvalID: debugPrefix, PromptID: p.ID, ConfigName: cfg.Name,
			Type:    progress.EventSendingPrompt,
			Message: fmt.Sprintf("Sending prompt (%d chars)...", len(p.PromptText)),
		})
	}
	// Apply an explicit deadline so the SDK does not fall back to its
	// hard-coded 60-second default (see copilot-sdk session.go).
	// Honour per-prompt Timeout (seconds) if set; otherwise use configured
	// session timeout (default 10 min).
	promptTimeout := 10 * time.Minute
	if e.sessionTimeout > 0 {
		promptTimeout = e.sessionTimeout
	}
	if p.Timeout > 0 {
		promptTimeout = time.Duration(p.Timeout) * time.Second
	}
	sendCtx, sendCancel := context.WithTimeout(genCtx, promptTimeout)
	defer sendCancel()

	lg.Debug("Sending prompt", "chars", len(p.PromptText), "timeout", promptTimeout)
	_, err = session.SendAndWait(sendCtx, copilot.MessageOptions{
		Prompt: p.PromptText,
	})
	if err != nil {
		mu.Lock()
		captured := make([]report.SessionEventRecord, len(sessionRecords))
		copy(captured, sessionRecords)
		capturedEvts := make([]copilot.SessionEvent, len(events))
		copy(capturedEvts, events)
		mu.Unlock()

		// Mid-generation action limit: return partial results so the
		// post-generation guardrail in engine.go can mark the eval as failed
		// with a proper reason instead of treating it as an SDK error.
		if actionLimitHit {
			generatedFiles, listErr := listFiles(workDir)
			if listErr != nil {
				lg.Warn("Failed to list generated files after action-limit", "dir", workDir, "error", listErr)
			}
			lg.Warn("Returning partial results after action-limit cancellation",
				"actions", actionCounter, "files", len(generatedFiles))
			cleanupCalled = true
			return &EvalResult{
				GeneratedFiles: generatedFiles,
				EventCount:     len(captured),
				ToolCalls:      extractToolCalls(capturedEvts),
				SessionEvents:  captured,
				ActionTimeline: BuildActionTimeline(captured),
				Success:        true, // Let engine.go guardrail set the proper failure
				FinalResponse:  extractLastAssistantMessage(captured),
				ToolReport:     toolReport,
				CleanupFn:      buildCleanupFn(),
			}, nil
		}

		return &EvalResult{
			SessionEvents:  captured,
			ActionTimeline: BuildActionTimeline(captured),
			EventCount:     len(captured),
			ToolCalls:      extractToolCalls(capturedEvts),
			Error:          fmt.Sprintf("prompt send failed: %v", err),
			ErrorDetails:   err.Error(),
			FinalResponse:  extractLastAssistantMessage(captured),
			ToolReport:     toolReport,
			CleanupFn:      buildCleanupFn(),
		}, fmt.Errorf("sending prompt: %w", err)
	}

	// Collect results
	mu.Lock()
	capturedEvents := make([]copilot.SessionEvent, len(events))
	copy(capturedEvents, events)
	capturedRecords := make([]report.SessionEventRecord, len(sessionRecords))
	copy(capturedRecords, sessionRecords)
	mu.Unlock()

	// Post-session tool verification gate (#347 / Item E / Option A).
	// The SDK emits SessionSkillsLoaded / SessionMcpServersLoaded only after
	// the first message round-trip, so this gate runs AFTER SendAndWait
	// returned — by which point the verifier's readyChan has typically already
	// closed from inside the OnEvent callback (either from normal tool events
	// OR from onSessionReady when AssistantTurnStart fired).
	//
	// waitForToolVerification now uses a 5-minute absolute ceiling as a
	// fail-safe in case the session never reached first turn (auth hang,
	// network failure, SDK bug). This is NOT the primary gate — the real
	// signal is AssistantTurnStart, which marks tool registration as
	// definitively complete. The ceiling is ONLY for broken sessions.
	//
	// Failure here is fatal to the eval: grading code that ran without the
	// configured tools produces false-positive scores. Match the
	// pre-session error format (Item D) by using
	// tool.SummarizeToolLoadErrors so operators see consistent messaging
	// regardless of which validation layer caught the breakage.
	if summary := postSessionToolVerification(ctx, verifier, 5*time.Minute); summary != "" {
		lg.Warn("Post-session tool verification failed; aborting before grading",
			"summary", summary)
		generatedFiles, listErr := listFiles(workDir)
		if listErr != nil {
			lg.Warn("Failed to list generated files after tool verification failure",
				"dir", workDir, "error", listErr)
		}
		cleanupCalled = true
		return &EvalResult{
			GeneratedFiles: generatedFiles,
			EventCount:     len(capturedEvents),
			ToolCalls:      extractToolCalls(capturedEvents),
			SessionEvents:  capturedRecords,
			ActionTimeline: BuildActionTimeline(capturedRecords),
			Success:        false,
			Error:          "tool_load_failure:\n" + summary,
			ErrorDetails:   summary,
			ErrorCategory:  "tool_load_failure",
			FinalResponse:  extractLastAssistantMessage(capturedRecords),
			ToolReport:     toolReport,
			CleanupFn:      buildCleanupFn(),
		}, fmt.Errorf("post-session tool verification: %s", summary)
	}

	generatedFiles, listErr := listFiles(workDir)
	if listErr != nil {
		lg.Warn("Failed to list generated files", "dir", workDir, "error", listErr)
	}
	toolCalls := extractToolCalls(capturedEvents)
	hasError := hasSessionError(capturedEvents)

	lg.Debug("Session results",
		"events", len(capturedEvents),
		"tool_calls", len(toolCalls),
		"files", len(generatedFiles))

	cleanupCalled = true
	return &EvalResult{
		GeneratedFiles: generatedFiles,
		EventCount:     len(capturedEvents),
		ToolCalls:      toolCalls,
		SessionEvents:  capturedRecords,
		ActionTimeline: BuildActionTimeline(capturedRecords),
		Success:        !hasError,
		Error:          "",
		FinalResponse:  extractLastAssistantMessage(capturedRecords),
		ToolReport:     toolReport,
		CleanupFn:      buildCleanupFn(),
	}, nil
}

// truncateStr truncates a string to maxLen characters, appending "..." if truncated.
func truncateStr(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

// detectFileCreation checks if assistant content looks like file creation
// and returns a summary like "key_vault_crud.py (89 lines)".
func detectFileCreation(content string) string {
	lines := strings.Split(strings.ReplaceAll(content, "\r\n", "\n"), "\n")
	// Look for patterns like "```python", file path references, or create_file tool patterns
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		// Detect "Creating file: filename" or "Writing filename" patterns
		for _, prefix := range []string{"creating file:", "writing file:", "creating ", "writing "} {
			lower := strings.ToLower(trimmed)
			if strings.HasPrefix(lower, prefix) {
				filename := strings.TrimSpace(trimmed[len(prefix):])
				if filename != "" && !strings.Contains(filename, " ") {
					lineCount := len(lines)
					return fmt.Sprintf("%s (%d lines)", filename, lineCount)
				}
			}
		}
	}
	// If content is very long (likely code), summarize by line count
	if len(lines) > 20 {
		// Try to find a filename from a markdown code fence
		for _, line := range lines[:min(5, len(lines))] {
			trimmed := strings.TrimSpace(line)
			if strings.HasPrefix(trimmed, "```") && len(trimmed) > 3 {
				return fmt.Sprintf("code block (%d lines)", len(lines))
			}
		}
	}
	return ""
}

// Client returns a new Copilot client for the given working directory.
// Exported for use by the review package.
func (e *CopilotPromptRunner) Client(ctx context.Context, workDir string) (*copilot.Client, error) {
	opts := *e.clientOpts
	opts.WorkingDirectory = workDir
	client := copilot.NewClient(&opts)
	if err := client.Start(ctx); err != nil {
		return nil, err
	}
	return client, nil
}

// mergePromptProperties builds a property map from a prompt's metadata fields.
// This map is used by ResolveTools to evaluate conditional tool entries.
func mergePromptProperties(p *prompt.Prompt) map[string]string {
	if p.Properties != nil {
		return p.Properties
	}
	return make(map[string]string)
}

func (e *CopilotPromptRunner) buildSessionConfig(ctx context.Context, cfg *config.ToolConfig, workDir string, configDir string, promptProps map[string]string) *copilot.SessionConfig {
	return e.buildSessionConfigForEval(ctx, cfg, workDir, configDir, promptProps, "", "", nil)
}

// buildTaggedEmit constructs the ProgressEmitter used by both the
// pre-session validator (WU-1) and buildSessionConfigForEval for any
// legacy emission path. Each event is tagged with evalID/promptID/
// configName so the interactive renderer routes it to the right block.
// Returns nil when progressFn is unset (e.g. tests).
func (e *CopilotPromptRunner) buildTaggedEmit(cfg *config.ToolConfig, evalID, promptID string) tool.ProgressEmitter {
	if e.progressFn == nil {
		return nil
	}
	fn := e.progressFn
	configName := ""
	if cfg != nil {
		configName = cfg.Name
	}
	return func(evt progress.ProgressEvent) {
		if evt.EvalID == "" {
			evt.EvalID = evalID
			evt.PromptID = promptID
			evt.ConfigName = configName
		}
		fn(evt)
	}
}

// cfgGeneratorTools returns the generator tool entries, or nil if absent.
func cfgGeneratorTools(cfg *config.ToolConfig) []tool.Entry {
	if cfg == nil || cfg.Generator == nil {
		return nil
	}
	return cfg.Generator.Tools
}

// buildSessionConfigForEval is the same as buildSessionConfig but tags
// tool-resolution progress events with the given evalID/promptID so the
// interactive renderer can route them to the active eval block.
//
// When toolReport is non-nil (the WU-1 path from Run), the session config
// is populated from the pre-validated report — no re-resolution, no
// duplicate progress emission. When toolReport is nil (the legacy
// buildSessionConfig path used by tests), the old emit + resolve code
// path runs as a fallback.
func (e *CopilotPromptRunner) buildSessionConfigForEval(ctx context.Context, cfg *config.ToolConfig, workDir string, configDir string, promptProps map[string]string, evalID, promptID string, toolReport *tool.ToolLoadReport) *copilot.SessionConfig {
	var skillDirs []string
	if toolReport != nil {
		skillDirs = toolReport.GeneratorSkillDirs()
	} else {
		// Legacy path: replicate the pre-WU-1 behavior (emit + resolve)
		// so existing tests that call buildSessionConfig directly keep
		// producing the same SessionConfig.
		taggedEmit := e.buildTaggedEmit(cfg, evalID, promptID)
		if taggedEmit != nil {
			if cfg.Generator != nil {
				tool.EmitMCPResolutions(cfg.Generator.Tools, taggedEmit)
			}
		}
		if cfg.Generator != nil {
			resolved, err := tool.ResolveSkillsWithReporter(ctx, cfg.Generator.Tools, configDir, taggedEmit)
			if err != nil {
				slog.Warn("Failed to resolve generator skill directories", "error", err)
			} else {
				skillDirs = resolved
			}
		}
	}
	// Use the config-driven system prompt (#115, #116). The default is zero
	// system prompt — all behavioral instructions belong in the config YAML.
	systemMsg := ""
	if cfg.Generator != nil && cfg.Generator.SystemPrompt != "" {
		systemMsg = cfg.Generator.SystemPrompt
	}

	// Safety boundaries (#36): when --allow-cloud is false (default), instruct
	// the generator to avoid provisioning real Azure resources. The agent should
	// use mock data, local emulators, environment variable placeholders, and IaC
	// templates instead of live CLI commands.
	if !e.allowCloud {
		systemMsg += "\n\nSAFETY BOUNDARIES:\n" +
			"Do NOT provision, create, modify, or delete real Azure resources. " +
			"Do NOT run `az`, `azd`, or ARM/Bicep deployment commands that target live Azure subscriptions. " +
			"Instead, use mock/fake connection strings, environment variable placeholders (e.g., " +
			"os.environ[\"AZURE_STORAGE_CONNECTION_STRING\"]), local emulators (Azurite, CosmosDB emulator), " +
			"and Infrastructure-as-Code templates (Bicep/Terraform) that define resources declaratively " +
			"without deploying them. All code must be runnable in a local-only, offline environment."
	}

	// Instruct the agent to use available skills before producing output.
	// Without this hint, models tend to go straight to producing output
	// and never invoke the skill tool, even when skills are loaded.
	// Only add if there are actual skills, not just empty directories (#291).
	if tool.CountSkills(skillDirs) > 0 {
		systemMsg += "\n\nSKILLS:\n" +
			"You have Azure SDK skills available. BEFORE writing any code, invoke the relevant skill " +
			"using the skill tool to get SDK-specific patterns, API examples, and acceptance criteria. " +
			"Also read the skill's reference files (acceptance-criteria.md, examples.md) for detailed guidance. " +
			"Then use that information to produce correct, modern Azure SDK output."
	}

	sc := &copilot.SessionConfig{
		Model:               cfg.Generator.Model,
		ConfigDirectory:     configDir,
		WorkingDirectory:    workDir,
		OnPermissionRequest: copilotperm.ApproveAll,
		Hooks: &copilot.SessionHooks{
			OnPreToolUse: func(input copilot.PreToolUseHookInput, invocation copilot.HookInvocation) (*copilot.PreToolUseHookOutput, error) {
				toolName := input.ToolName
				// Validate file paths for file-write tools
				if isFileWriteTool(toolName) {
					if args, ok := input.ToolArgs.(map[string]interface{}); ok {
						if p, ok := args["path"].(string); ok {
							resolved := p
							if !filepath.IsAbs(resolved) {
								resolved = filepath.Join(workDir, resolved)
							}
							resolved = filepath.Clean(resolved)
							absWork := filepath.Clean(workDir)
							rel, err := filepath.Rel(absWork, resolved)
							if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
								slog.Warn("File path outside workspace",
									"tool", toolName, "path", p, "resolved", resolved, "workspace", absWork)
								return &copilot.PreToolUseHookOutput{
									PermissionDecision:       "deny",
									PermissionDecisionReason: fmt.Sprintf("path %q is outside workspace %q", p, absWork),
								}, nil
							}
						}
					}
				}
				// Log bash/command tools
				if toolName == "bash" || toolName == "shell" || toolName == "run_command" {
					if args, ok := input.ToolArgs.(map[string]interface{}); ok {
						if cmd, ok := args["command"].(string); ok {
							slog.Debug("Command execution", "tool", toolName, "command", truncateStr(cmd, 120))
						}
					}
				}
				return &copilot.PreToolUseHookOutput{PermissionDecision: "allow"}, nil
			},
			OnPostToolUse: func(input copilot.PostToolUseHookInput, invocation copilot.HookInvocation) (*copilot.PostToolUseHookOutput, error) {
				slog.Debug("Tool complete", "tool", input.ToolName)
				// Check file sizes for file operations
				if isFileWriteTool(input.ToolName) {
					if args, ok := input.ToolArgs.(map[string]interface{}); ok {
						if p, ok := args["path"].(string); ok {
							if info, err := os.Stat(p); err == nil && info.Size() > 100*1024 {
								slog.Warn("Large file created", "path", p, "bytes", info.Size())
							}
						}
					}
				}
				return &copilot.PostToolUseHookOutput{}, nil
			},
		},
		SkillDirectories: skillDirs,
	}

	// Resolve tools: conditional tool entries are filtered by prompt properties.
	// An empty slice serializes as JSON [] which tells the CLI "zero tools" —
	// nil serializes as null which means "all default tools available."
	var toolEntries []config.ToolEntry
	for _, entry := range cfg.Generator.Tools {
		if entry.ResolvedType() == "tool" {
			toolEntries = append(toolEntries, entry)
		}
	}
	availableTools := config.ResolveTools(toolEntries, promptProps)
	if len(toolEntries) > 0 {
		slog.Debug("Resolved conditional tools",
			"entries", len(toolEntries),
			"matched", len(availableTools),
			"tools", availableTools,
			"properties", promptProps)
	}
	excludedTools := cfg.Generator.ExcludedTools
	if len(availableTools) > 0 {
		sc.AvailableTools = availableTools
	}
	if len(excludedTools) > 0 {
		sc.ExcludedTools = excludedTools
	}

	// Map MCP servers
	var mcpEntries []config.ToolEntry
	for _, entry := range cfg.Generator.Tools {
		if entry.ResolvedType() == "mcp" {
			mcpEntries = append(mcpEntries, entry)
		}
	}
	if len(mcpEntries) > 0 {
		sc.MCPServers = make(map[string]copilot.MCPServerConfig, len(mcpEntries))
		for _, entry := range mcpEntries {
			mcpType := entry.ResolvedMCPType()
			var mcpCfg copilot.MCPServerConfig
			if mcpType == "remote" {
				mcpCfg = copilot.MCPHTTPServerConfig{
					Tools: entry.MCPTools,
					URL:   entry.URL,
				}
			} else {
				mcpCfg = copilot.MCPStdioServerConfig{
					Tools:   entry.MCPTools,
					Command: entry.Command,
					Args:    entry.Args,
				}
			}
			sc.MCPServers[entry.Name] = mcpCfg
			slog.Info("MCP server configured",
				"name", entry.Name,
				"type", mcpType,
				"command", entry.Command,
				"url", entry.URL,
				"args", entry.Args,
				"tools", entry.MCPTools,
			)
		}
	} else {
		slog.Debug("No MCP servers configured")
	}

	// Set the system message: start with config-driven system prompt,
	// then append any accumulated hints (skills, MCP, etc.)
	if systemMsg != "" {
		sc.SystemMessage = &copilot.SystemMessageConfig{
			Mode:    "append",
			Content: systemMsg,
		}
		slog.Info("Generator system prompt configured", "length", len(systemMsg))
	}

	return sc
}

// extractToolCalls returns unique tool names from session events.
func extractToolCalls(events []copilot.SessionEvent) []string {
	seen := make(map[string]bool)
	var tools []string
	for _, e := range events {
		if e.Type() == copilot.SessionEventTypeToolExecutionStart ||
			e.Type() == copilot.SessionEventTypeToolExecutionComplete {
			name := ""
			if details := copilotevent.Extract(e); details.ToolName != nil {
				name = *details.ToolName
			}
			if name != "" && !seen[name] {
				seen[name] = true
				tools = append(tools, name)
			}
		}
	}
	return tools
}

// hasSessionError checks for error events.
func hasSessionError(events []copilot.SessionEvent) bool {
	for _, e := range events {
		if e.Type() == copilot.SessionEventTypeSessionError {
			return true
		}
	}
	return false
}

// extractLastAssistantMessage returns the content of the last assistant message
// from session event records. Returns empty string if no assistant messages found.
func extractLastAssistantMessage(records []report.SessionEventRecord) string {
	for i := len(records) - 1; i >= 0; i-- {
		if records[i].Type == "assistant.message" && records[i].Content != "" {
			return records[i].Content
		}
	}
	return ""
}

// isFileWriteTool returns true for tools that create or modify files.
func isFileWriteTool(name string) bool {
	switch name {
	case "create", "edit", "write_file", "create_file",
		"insert_edit_into_file", "write_to_file":
		return true
	}
	return false
}

// extractAbsPathsFromCommand extracts absolute paths from a shell command string.
// Used for containment checking of bash/shell tool invocations.
func extractAbsPathsFromCommand(cmd string) []string {
	var paths []string
	for _, part := range strings.Fields(cmd) {
		if strings.HasPrefix(part, "/") && len(part) > 1 {
			abs, err := filepath.Abs(part)
			if err == nil {
				paths = append(paths, abs)
			} else {
				paths = append(paths, part)
			}
		}
	}
	return paths
}

// toolArgSummary extracts a short summary of the tool's primary argument.
func toolArgSummary(event copilot.SessionEvent) string {
	details := copilotevent.Extract(event)
	if details.Path != nil && *details.Path != "" {
		return filepath.Base(*details.Path)
	}
	if details.Arguments != nil {
		if args, ok := details.Arguments.(map[string]interface{}); ok {
			for _, key := range []string{"path", "file", "command"} {
				if v, ok := args[key]; ok {
					if s, ok := v.(string); ok && s != "" {
						if key == "path" || key == "file" {
							return filepath.Base(s)
						}
						return truncateStr(s, 40)
					}
				}
			}
		}
	}
	return ""
}

// isolateSkills copies each resolved skill directory into the per-session
// configDir so sessions don't share mutable skill state. Returns the new
// isolated paths. If a copy fails, the original path is kept.
func isolateSkills(resolved []string, configDir string) []string {
	if len(resolved) == 0 {
		return nil
	}
	skillsBase := filepath.Join(configDir, "skills")
	if err := os.MkdirAll(skillsBase, 0755); err != nil {
		slog.Warn("Failed to create per-session skills dir, using originals", "error", err)
		return resolved
	}
	isolated := make([]string, 0, len(resolved))
	for _, src := range resolved {
		name := filepath.Base(src)
		dst := filepath.Join(skillsBase, name)
		if err := copyDir(src, dst); err != nil {
			slog.Warn("Failed to isolate skill, using original", "skill", name, "error", err)
			isolated = append(isolated, src)
			continue
		}
		isolated = append(isolated, dst)
	}
	return isolated
}
