---
name: "copilot-sdk-integration"
description: "How hyoka uses the Copilot SDK"
domain: "architecture"
confidence: "high"
source: "hyoka/internal/eval/copilot.go"
---

## Context

Hyoka integrates with the Copilot SDK to launch agent sessions (both for code generation and code review). The SDK provides session lifecycle management, tool invocation, and action history capture.

## Session Lifecycle

### Creation
```go
sess, err := SDK.NewSession(ctx, &SessionOptions{
    Model: "claude-opus-4.6",
    SystemPrompt: "You are an Azure SDK code generator...",
    Tools: toolSet,
    Skills: skillPaths,
    MCPServers: mcpConfigs,
    MaxTurns: 25,
    MaxFiles: 50,
    MaxOutputSize: 1MB,
})
```

### Agent Turn Loop
```go
// Send initial prompt
resp, err := sess.SendMessage(ctx, prompt)

// Agent responds with actions (tool calls, file operations, etc)
for action := range resp.Actions {
    // Hyoka captures action metadata for timeline
    timeline.Append(action)
    
    // Let SDK handle tool execution
    result, err := sess.ExecuteAction(ctx, action)
}

// Repeat until agent stops or max turns reached
```

### Cleanup
```go
sess.Close()  // Kills Copilot process, cleans workspace
```

## Action Capture

The SDK returns **SessionEventRecord** for each agent action:
- Type: `tool_call`, `tool_result`, `file_read`, `file_write`, `bash_command`
- Tool name, arguments, output
- Duration, success/failure status
- Turn number

Hyoka maps these into **ActionEvent** for graders and reports.

## Skills Integration

Skills are **SKILL.md files** injected into the SDK session. They provide domain knowledge (e.g., Azure SDK patterns, Python conventions) without needing direct prompt modification.

```yaml
generator:
  skills:
    - type: local
      path: ./skills/generator/**/*.md
```

The SDK concatenates skills into the system prompt at session creation.

## MCP Server Integration

MCP (Model Context Protocol) servers extend tool availability beyond built-in SDK tools.

```yaml
generator:
  mcp_servers:
    azure-cli:
      type: command
      command: /usr/bin/az
      args: ["mcp", "run"]
      tools: ["azure_resource_list"]
```

The SDK spawns the MCP server process and routes tool calls to it.

## Timeout and Resource Monitoring

- **Session timeout:** If agent doesn't respond within `--gen-timeout` (default 600s), SDK terminates
- **Resource monitoring:** Optional per-config (via `--monitor-resources`), tracks CPU/memory per session
- **Process tracking:** Hyoka tracks child PIDs for cleanup on interruption

## Tool Result Feedback Loop

After each tool execution:
1. SDK captures result (stdout, stderr, exit code)
2. Hyoka records in action timeline
3. Result fed back to agent in next turn
4. Agent continues or stops

## Error Handling

- **Tool not found:** Agent sees error, can adjust
- **Tool timeout:** Reported to agent, can retry
- **SDK session crash:** Entire eval marked as failed, workspace preserved for debugging

## Code Locations

- **Copilot SDK session management:** `hyoka/internal/eval/copilot.go`
- **Action capture and timeline:** `hyoka/internal/eval/action.go`
- **Engine integration:** `hyoka/internal/eval/engine.go` (orchestrates SDK calls)

## Anti-Patterns

- Modifying agent response mid-turn (let SDK handle tool execution)
- Killing process without calling `sess.Close()` (orphans may remain)
- Assuming all tools are available (check config tool filters)
- Not capturing action timeline (required for graders and debugging)
