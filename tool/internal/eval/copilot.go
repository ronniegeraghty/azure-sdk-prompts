package eval

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	copilot "github.com/github/copilot-sdk/go"
	"github.com/ronniegeraghty/azure-sdk-prompts/tool/internal/config"
	"github.com/ronniegeraghty/azure-sdk-prompts/tool/internal/prompt"
)

// CopilotSDKEvaluator uses the Copilot SDK to run real evaluations.
type CopilotSDKEvaluator struct {
	clientOpts *copilot.ClientOptions
	debug      bool
}

// CopilotEvalOptions configures the CopilotSDKEvaluator.
type CopilotEvalOptions struct {
	// GitHubToken for SDK authentication (optional; falls back to logged-in user).
	GitHubToken string
	// CLIPath overrides the Copilot CLI executable path.
	CLIPath string
	// Debug enables verbose logging.
	Debug bool
}

// NewCopilotSDKEvaluator creates a new evaluator backed by the Copilot SDK.
func NewCopilotSDKEvaluator(opts CopilotEvalOptions) *CopilotSDKEvaluator {
	clientOpts := &copilot.ClientOptions{}
	if opts.GitHubToken != "" {
		clientOpts.GitHubToken = opts.GitHubToken
	}
	if opts.CLIPath != "" {
		clientOpts.CLIPath = opts.CLIPath
	}
	if opts.Debug {
		clientOpts.LogLevel = "debug"
	}
	return &CopilotSDKEvaluator{
		clientOpts: clientOpts,
		debug:      opts.Debug,
	}
}

// Evaluate runs a prompt through a real Copilot session and returns generated files and events.
func (e *CopilotSDKEvaluator) Evaluate(ctx context.Context, p *prompt.Prompt, cfg *config.ToolConfig, workDir string) (*EvalResult, error) {
	// Copy starter project if configured
	if p.StarterProject != "" {
		starterDir := p.StarterProject
		if !filepath.IsAbs(starterDir) && p.FilePath != "" {
			starterDir = filepath.Join(filepath.Dir(p.FilePath), starterDir)
		}
		if err := copyDir(starterDir, workDir); err != nil {
			return nil, fmt.Errorf("copying starter project: %w", err)
		}
	}

	// Create Copilot client
	opts := *e.clientOpts
	opts.Cwd = workDir
	client := copilot.NewClient(&opts)

	if err := client.Start(ctx); err != nil {
		return nil, fmt.Errorf("starting copilot client: %w", err)
	}
	defer client.Stop()

	// Build session config from tool config
	sessionCfg := e.buildSessionConfig(cfg, workDir)

	session, err := client.CreateSession(ctx, sessionCfg)
	if err != nil {
		return nil, fmt.Errorf("creating session: %w", err)
	}
	defer session.Disconnect()

	// Subscribe to events
	var events []copilot.SessionEvent
	var mu sync.Mutex
	unsub := session.On(func(event copilot.SessionEvent) {
		mu.Lock()
		events = append(events, event)
		mu.Unlock()
	})
	defer unsub()

	// Send the prompt
	_, err = session.SendAndWait(ctx, copilot.MessageOptions{
		Prompt: p.PromptText,
	})
	if err != nil {
		return nil, fmt.Errorf("sending prompt: %w", err)
	}

	// Collect results
	mu.Lock()
	capturedEvents := make([]copilot.SessionEvent, len(events))
	copy(capturedEvents, events)
	mu.Unlock()

	generatedFiles := listWorkspaceFiles(workDir)
	toolCalls := extractToolCalls(capturedEvents)
	hasError := hasSessionError(capturedEvents)

	if e.debug {
		log.Printf("[copilot] %s/%s: %d events, %d tool calls, %d files",
			p.ID, cfg.Name, len(capturedEvents), len(toolCalls), len(generatedFiles))
	}

	return &EvalResult{
		GeneratedFiles: generatedFiles,
		EventCount:     len(capturedEvents),
		ToolCalls:      toolCalls,
		Success:        !hasError,
		Error:          "",
	}, nil
}

// Client returns a new Copilot client for the given working directory.
// Exported for use by the review package.
func (e *CopilotSDKEvaluator) Client(ctx context.Context, workDir string) (*copilot.Client, error) {
	opts := *e.clientOpts
	opts.Cwd = workDir
	client := copilot.NewClient(&opts)
	if err := client.Start(ctx); err != nil {
		return nil, err
	}
	return client, nil
}

func (e *CopilotSDKEvaluator) buildSessionConfig(cfg *config.ToolConfig, workDir string) *copilot.SessionConfig {
	sc := &copilot.SessionConfig{
		Model: cfg.Model,
		SystemMessage: &copilot.SystemMessageConfig{
			Mode: "append",
			Content: `You are being evaluated on code generation quality.
Write complete, working code. Use the specified SDK packages.
Do not ask clarifying questions — make reasonable assumptions.
Write all code to the current working directory.`,
		},
		WorkingDirectory:    workDir,
		OnPermissionRequest: copilot.PermissionHandler.ApproveAll,
		AvailableTools:      cfg.AvailableTools,
		ExcludedTools:       cfg.ExcludedTools,
		SkillDirectories:    cfg.SkillDirectories,
	}

	// Map MCP servers
	if len(cfg.MCPServers) > 0 {
		sc.MCPServers = make(map[string]copilot.MCPServerConfig, len(cfg.MCPServers))
		for name, srv := range cfg.MCPServers {
			sc.MCPServers[name] = copilot.MCPServerConfig{
				"type":    srv.Type,
				"command": srv.Command,
				"args":    srv.Args,
			}
		}
	}

	return sc
}

// extractToolCalls returns unique tool names from session events.
func extractToolCalls(events []copilot.SessionEvent) []string {
	seen := make(map[string]bool)
	var tools []string
	for _, e := range events {
		if e.Type == copilot.SessionEventTypeToolExecutionStart ||
			e.Type == copilot.SessionEventTypeToolExecutionComplete {
			name := ""
			if e.Data.ToolName != nil {
				name = *e.Data.ToolName
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
		if e.Type == copilot.SessionEventTypeSessionError {
			return true
		}
	}
	return false
}

// listWorkspaceFiles returns all file paths relative to the workspace dir.
func listWorkspaceFiles(dir string) []string {
	var files []string
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		if strings.HasPrefix(filepath.Base(path), ".") {
			return nil
		}
		rel, err := filepath.Rel(dir, path)
		if err != nil {
			return nil
		}
		files = append(files, rel)
		return nil
	})
	return files
}

// copyDir recursively copies src to dst.
func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if info.IsDir() {
			return os.MkdirAll(target, 0755)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, info.Mode())
	})
}
