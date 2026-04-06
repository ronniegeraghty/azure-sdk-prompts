# Tool Registry Design

## Overview

The **tool registry** is a YAML catalog of curated tool configurations—MCP servers
and Copilot skills—that teams can reference by name instead of inlining full
connection details in every evaluation config.

### Goals

| Goal | Detail |
|------|--------|
| **DRY configs** | Declare a tool once, reference it everywhere by name. |
| **Versioned catalog** | Pin tool versions so evaluations are reproducible. |
| **Local + remote** | Load registries from a local file or a URL (for shared team catalogs). |
| **Type safety** | Each tool type (`mcp`, `skill`) carries only the fields it needs. |

## YAML Schema

```yaml
# version: Schema version (currently "1").
# tools:   Ordered list of tool entries.

version: "1"
tools:
  - name: <string>              # unique identifier (kebab-case)
    type: <"mcp" | "skill">     # tool type
    description: <string>       # human-readable summary
    version: <string>           # semver of the tool itself

    # --- MCP-specific (required when type == "mcp") ---
    mcp:
      command: <string>         # executable to spawn
      args: [<string>, ...]     # command arguments
      env:                      # optional environment variables
        KEY: VALUE
      tools: [<string>, ...]    # tool allow-list ("*" = all)

    # --- Skill-specific (required when type == "skill") ---
    skill:
      source: <"local" | "remote">
      repo: <string>           # GitHub owner/repo (remote only)
      name: <string>           # skill name within repo (remote only)
      path: <string>           # local path or glob (local only)
```

### Field Reference

| Field | Required | Description |
|-------|----------|-------------|
| `version` | yes | Schema version. Must be `"1"`. |
| `tools` | yes | List of `ToolEntry` objects. |
| `tools[].name` | yes | Unique kebab-case identifier (e.g., `azure-mcp`). |
| `tools[].type` | yes | `"mcp"` or `"skill"`. |
| `tools[].description` | no | Human-readable summary. |
| `tools[].version` | no | Semver string for the tool version. |
| `tools[].mcp` | when type=mcp | MCP server configuration block. |
| `tools[].mcp.command` | yes | Executable command to spawn the server. |
| `tools[].mcp.args` | no | Command-line arguments. |
| `tools[].mcp.env` | no | Environment variables as key-value pairs. |
| `tools[].mcp.tools` | no | Tool allow-list. `["*"]` permits all. |
| `tools[].skill` | when type=skill | Skill configuration block. |
| `tools[].skill.source` | yes | `"local"` or `"remote"`. |
| `tools[].skill.repo` | remote | GitHub `owner/repo`. |
| `tools[].skill.name` | remote | Skill name within the repository. |
| `tools[].skill.path` | local | File system path or glob pattern. |

## Example

```yaml
version: "1"
tools:
  - name: azure-mcp
    type: mcp
    description: "Azure MCP server for cloud resource management"
    version: "0.1.0"
    mcp:
      command: npx
      args: ["-y", "@azure/mcp@latest", "server", "start"]
      tools: ["*"]

  - name: azure-keyvault-skill
    type: skill
    description: "Azure Key Vault Python skill"
    version: "1.0.0"
    skill:
      source: remote
      repo: microsoft/skills
      name: azure-keyvault-py

  - name: local-generator-skill
    type: skill
    description: "Custom generator skill from local path"
    version: "0.0.1"
    skill:
      source: local
      path: "./skills/generator/*"
```

## Go API

```go
import "github.com/ronniegeraghty/hyoka/internal/tools"

// Load from a local YAML file.
reg, err := tools.LoadRegistry("tools.yaml")

// Load from a URL (shared team catalog).
reg, err := tools.LoadRemoteRegistry("https://example.com/registry.yaml")

// Look up a tool by name.
entry, ok := reg.Get("azure-mcp")
```

## Integration with Evaluation Configs

Evaluation configs (`configs/*.yaml`) can reference registry entries by name
using the `tools:` field on a generator or reviewer:

```yaml
configs:
  - name: my-eval/claude-opus-4.6
    generator:
      model: claude-opus-4.6
      tools:
        - name: azure-mcp      # resolved from the active registry
          always_on: true
```

The evaluation engine resolves tool names against the loaded registry and
expands them into full MCP server or skill configurations before starting a
Copilot session.

## Validation Rules

1. `version` must equal `"1"`.
2. Every tool must have a non-empty `name` and a valid `type`.
3. Tool names must be unique within a registry.
4. MCP entries must have a non-empty `mcp.command`.
5. Skill entries must have a valid `skill.source` (`"local"` or `"remote"`).
6. Remote skills require `skill.repo` and `skill.name`.
7. Local skills require `skill.path`.
