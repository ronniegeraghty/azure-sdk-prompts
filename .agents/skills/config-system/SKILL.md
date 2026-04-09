---
name: "config-system"
description: "Generator/Reviewer sub-structs, property-based tool filters"
domain: "architecture"
confidence: "high"
source: "hyoka/internal/config/config.go"
---

## Context

Hyoka's config system is YAML-driven and highly composable. Each evaluation config defines separate generator and reviewer specifications, with tool filtering via properties (allow/deny lists and skill injection).

## Config Structure

### Top-Level ToolConfig
```yaml
name: baseline/claude-opus-4.6
description: "Baseline Opus eval"
generator:
  model: claude-opus-4.6
  system_prompt: "You are..."
  skills: [...]
  mcp_servers: {...}
  tools: [...]
  available_tools: [...]    # Allowlist
  excluded_tools: [...]     # Denylist
reviewer:
  models: [claude-opus-4.6, gpt-5.4]
  system_prompt: "Review this code..."
plugins: [...]
limits:
  max_turns: 25
  max_files: 50
  max_output_size: 1048576
  max_session_actions: 50
```

### Generator vs. Reviewer Separation

**GeneratorConfig** (code generation):
- Single `model` field (one generator model)
- Optional `system_prompt` override
- Skills injection for generator context
- MCP servers for external tools
- Tool filters (available_tools / excluded_tools)

**ReviewerConfig** (code review):
- `models` array (multiple reviewers in parallel)
- Optional `system_prompt` for review instructions
- Skills injection for review guidance
- No MCP servers (reviewers don't call tools)

## Tool Filtering

Tools are filtered via **ToolEntry** specs:
```go
type ToolEntry struct {
    Name       string            `yaml:"name"`
    Include    bool              `yaml:"include"`           // whitelist
    Exclude    bool              `yaml:"exclude"`           // blacklist
    Properties map[string]string `yaml:"properties"`        // filters by context
}
```

### Property-Based Tool Filtering

Properties are **snake_case** metadata on tools (e.g., `language:python`, `service:keyvault`). Generator model sees only tools matching the current prompt's properties.

**Example:**
```yaml
# Config
tools:
  - name: python_test_runner
    include: true
    properties:
      language: python

# Prompt metadata
language: python  # Matches → tool included

language: dotnet  # No match → tool hidden
```

## Composition Patterns

**Skill Injection:**
```yaml
generator:
  skills:
    - type: local
      path: ./skills/generator/**/*.md
    - type: remote
      name: azure-prompt-patterns
      repo: github.com/Azure/hyoka-skills
```

**MCP Server Configuration:**
```yaml
generator:
  mcp_servers:
    azure-cli:
      type: command
      command: az
      args: ["--version"]
      tools: ["azure_resource_list", "azure_resource_get"]
```

## Key Patterns

- **Single model for generation, multiple for review:** Enables comparison across reviewers
- **Property-based filtering:** Tools scope automatically to prompt language/service
- **SessionLimits as guardrails:** Per-config turn/file/action limits override engine defaults
- **Plugins for extensibility:** Config-level plugin registration (for future custom scoring)

## Related Code Locations

- **Config loading:** `hyoka/internal/config/config.go`
- **Tool filtering logic:** `hyoka/internal/config/tools.go`
- **Example configs:** `configs/*.yaml`

## Anti-Patterns

- Hardcoding tool availability (use properties instead)
- Using different skill sets for generator and reviewer without documentation
- Conflicting available_tools and excluded_tools (excluded wins)
- Ignoring SessionLimits in custom configs (defaults should be sufficient)
