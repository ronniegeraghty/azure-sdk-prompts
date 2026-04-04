package trends

import (
"context"
"fmt"
"log/slog"
"os"
"strings"
"sync"
"time"

copilot "github.com/github/copilot-sdk/go"
)

// AnalyzeTrends uses a Copilot SDK session to perform AI-powered trend analysis.
// It returns the analysis text to be included in the report.
func AnalyzeTrends(ctx context.Context, tr *TrendReport) (string, error) {
lg := slog.With("role", "trend-analysis", "model", "gpt-4.1")
lg.Info("Starting AI trend analysis", "total_runs", tr.TotalRuns, "prompts", len(tr.PromptTrends))
prompt := formatTrendPrompt(tr)

lg.Debug("Starting Copilot client")
client := copilot.NewClient(nil)
if err := client.Start(ctx); err != nil {
return "", fmt.Errorf("starting copilot client: %w", err)
}
var trendSessionID string
defer func() {
// Delete session state before stopping client (#62)
if trendSessionID != "" {
deleteCtx, deleteCancel := context.WithTimeout(context.Background(), 5*time.Second)
defer deleteCancel()
if err := client.DeleteSession(deleteCtx, trendSessionID); err != nil {
lg.Debug("Session delete failed", "sessionID", trendSessionID, "error", err)
}
}
client.Stop()
}()

// Create isolated config directory to prevent user-level skills from
// leaking into the analysis session (#21).
configDir, err := os.MkdirTemp("", "hyoka-config-*")
if err != nil {
return "", fmt.Errorf("creating isolated config dir: %w", err)
}
defer os.RemoveAll(configDir)

lg.Debug("Creating session")
session, err := client.CreateSession(ctx, &copilot.SessionConfig{
Model: "gpt-4.1",
SystemMessage: &copilot.SystemMessageConfig{
Mode:    "append",
Content: `You are an Azure SDK code evaluation analyst. You analyze results from automated code generation evaluations where different AI agent configurations (baseline, with-skills, with-MCP-tools) are compared.

Your output must be structured markdown with these exact sections:

## Key Findings
A 3-5 bullet summary of the most important takeaways.

## Config Comparison
For each prompt evaluated, a table comparing configs:
| Criteria | Config A | Config B | Config C | Notes |
Use ✅/❌ for pass/fail. Add a Notes column explaining differences.

## What Tools/Skills Helped
- Which specific tools or skills were invoked and what impact they had
- What the agent got RIGHT because of tools/skills that baseline got wrong
- What the agent got WRONG despite having tools/skills

## What Needs Investigation
- Criteria that failed across ALL configs (systemic gaps)
- Unexpected regressions (config with more tools scored lower)
- Non-deterministic behavior (same config, different results)

Be specific — reference exact criteria names, tool names, and scores. Do NOT reference historical runs or prompts not in the current data.`,
},
ConfigDir:           configDir,
OnPermissionRequest: copilot.PermissionHandler.ApproveAll,
})
if err != nil {
return "", fmt.Errorf("creating analysis session: %w", err)
}
trendSessionID = session.SessionID
defer session.Disconnect()

// Capture assistant messages via event subscription
var assistantContent strings.Builder
var mu sync.Mutex
unsub := session.On(func(event copilot.SessionEvent) {
if event.Type == copilot.SessionEventTypeAssistantMessage && event.Data.Content != nil {
mu.Lock()
assistantContent.WriteString(*event.Data.Content)
mu.Unlock()
}
})
defer unsub()

lg.Debug("Sending trend analysis prompt", "prompt_chars", len(prompt))
_, err = session.SendAndWait(ctx, copilot.MessageOptions{
Prompt: prompt,
})
if err != nil {
return "", fmt.Errorf("analysis request failed: %w", err)
}

mu.Lock()
result := strings.TrimSpace(assistantContent.String())
mu.Unlock()

lg.Info("AI trend analysis complete", "result_chars", len(result))

if result == "" {
return "", fmt.Errorf("empty analysis response")
}

return result, nil
}

// formatTrendPrompt builds a structured prompt focused on tool usage and resource impact.
func formatTrendPrompt(tr *TrendReport) string {
var b strings.Builder

b.WriteString("Analyze the following evaluation results from a SINGLE RUN comparing different AI agent configurations.\n\n")
b.WriteString("Each config represents a different tool setup for the same AI model:\n")
b.WriteString("- 'baseline' = no tools, no skills (agent relies on built-in knowledge)\n")
b.WriteString("- 'baseline-skills' = SDK-specific skills loaded (patterns, examples, acceptance criteria)\n")
b.WriteString("- 'azure-mcp' = Azure MCP server tools (get_best_practices, documentation lookup)\n\n")
b.WriteString("Analyze ONLY the data below. Do NOT reference prompts or runs not listed here.\n\n")
b.WriteString("---\n\n")

// Overall stats
passed, failed := 0, 0
for _, e := range tr.Entries {
if e.Success {
passed++
} else {
failed++
}
}
fmt.Fprintf(&b, "## Overview\n- Total evaluations: %d\n- Passed: %d (%.0f%%)\n- Failed: %d\n- Unique prompts: %d\n- Run count: %d\n\n",
tr.TotalRuns, passed, pct(passed, tr.TotalRuns), failed, len(tr.PromptTrends), len(tr.RunIDs))

// Tool usage summary across configs
b.WriteString("## Tool Usage by Config\n\n")
configTools := map[string]map[string]int{} // config -> tool -> count
configRuns := map[string]int{}
for _, e := range tr.Entries {
configRuns[e.ConfigName]++
if configTools[e.ConfigName] == nil {
configTools[e.ConfigName] = map[string]int{}
}
for _, t := range e.ToolCalls {
configTools[e.ConfigName][t]++
}
}
for cfg, tools := range configTools {
fmt.Fprintf(&b, "### %s (%d runs)\n", cfg, configRuns[cfg])
if len(tools) == 0 {
b.WriteString("- No tool calls (agent relied on built-in knowledge)\n")
} else {
for tool, count := range tools {
fmt.Fprintf(&b, "- **%s**: used in %d runs\n", tool, count)
}
}
b.WriteString("\n")
}

// Per-prompt detail with tool call info
b.WriteString("## Per-Prompt Performance & Tool Usage\n\n")
for _, pt := range tr.PromptTrends {
fmt.Fprintf(&b, "### %s\n", pt.PromptID)
for cfg, runs := range pt.Configs {
p, f := 0, 0
totalDur := 0.0
toolSet := map[string]int{}
for _, r := range runs {
if r.Success {
p++
} else {
f++
}
totalDur += r.Duration
for _, t := range r.ToolCalls {
toolSet[t]++
}
}
avgDur := 0.0
if len(runs) > 0 {
avgDur = totalDur / float64(len(runs))
}
trend := pt.Trend[cfg]
fmt.Fprintf(&b, "- **%s**: %d/%d passed (%.0f%%), avg duration: %s, trend: %s\n",
cfg, p, p+f, pct(p, p+f), formatDuration(avgDur), trend)

// Include score details from most recent run
if len(runs) > 0 {
latest := runs[len(runs)-1]
if latest.HasReview && latest.MaxScore > 0 {
fmt.Fprintf(&b, "  Score: %d/%d (%.0f%%)\n", latest.Score, latest.MaxScore, pct(latest.Score, latest.MaxScore))
}
}

if len(toolSet) > 0 {
b.WriteString("  Tools used: ")
first := true
for tool, count := range toolSet {
if !first {
b.WriteString(", ")
}
fmt.Fprintf(&b, "%s(×%d)", tool, count)
first = false
}
b.WriteString("\n")
} else {
b.WriteString("  Tools used: none (built-in knowledge only)\n")
}

if f > 0 {
b.WriteString("  Run history: ")
for i, r := range runs {
if i > 0 {
b.WriteString(", ")
}
icon := "PASS"
if !r.Success {
icon = "FAIL"
}
fmt.Fprintf(&b, "%s(%s)", icon, r.RunID)
}
b.WriteString("\n")
}
}
b.WriteString("\n")
}

return b.String()
}
