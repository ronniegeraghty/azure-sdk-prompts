# Configuration Guide

This guide covers the hyoka configuration system — how to define evaluation environments with models, skills, and MCP servers.

## Overview

Configuration files live in the `configs/` directory as YAML files. Each file defines one or more evaluation configurations specifying a **generator** model (for code generation) and a **reviewer** panel (for code review).

hyoka auto-discovers all `.yaml` files in `configs/` at runtime. You can also point to a specific file with `--config-file`.

## Config File Structure

A config file contains one or more configurations under a top-level `configs:` key:

```yaml
configs:
  - name: baseline/claude-sonnet-4.5
    description: "Baseline — Claude Sonnet 4.5"
    generator:
      model: "claude-sonnet-4.5"
    reviewer:
      models:
        - "claude-opus-4.6"
        - "gemini-3-pro-preview"
        - "gpt-4.1"
```

## Generator Configuration

The `generator:` section defines the code generation agent:

```yaml
generator:
  model: "claude-sonnet-4.5"
  skills:
    - type: remote
      name: azure-keyvault-py
      repo: microsoft/skills
    - type: local
      path: "./skills/generator"
  mcp_servers:
    azure:
      type: local
      command: npx
      args: ["-y", "@azure/mcp@latest"]
      tools: ["*"]
  available_tools: ["create", "edit", "bash"]
  excluded_tools: ["web_fetch"]
```

### Generator Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `model` | string | ✓ | LLM model for code generation. Examples: `claude-sonnet-4.5`, `claude-opus-4.6`, `gpt-4.1` |
| `skills` | list | | Skills providing domain knowledge. See [Skills](#skills). |
| `mcp_servers` | map | | MCP server configurations. See [MCP Servers](#mcp-servers). |
| `available_tools` | list | | Restrict generator to these tools only. Empty = all tools. |
| `excluded_tools` | list | | Tools to exclude from the generator. |

## Reviewer Configuration

The `reviewer:` section defines the multi-model review panel:

```yaml
reviewer:
  models:
    - "claude-opus-4.6"
    - "gemini-3-pro-preview"
    - "gpt-4.1"
  skills:
    - type: local
      path: "./skills/reviewer"
```

### Reviewer Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `models` | list | | List of reviewer models. First model acts as consolidator. No duplicates. |
| `skills` | list | | Skills for reviewer agents. See [Skills](#skills). |

The review panel runs all models in parallel, then the first model consolidates results using majority voting per criterion.

## Skills

Skills give agents domain-specific knowledge (SDK patterns, API examples, acceptance criteria) that improve code generation and review quality.

Each skill has a `type` — either `local` or `remote`:

### Local Skills

Load from a filesystem directory containing a `SKILL.md` file:

```yaml
skills:
  - type: local
    path: "./skills/generator"
  - type: local
    path: "./skills/reviewer/*"   # glob patterns supported
```

| Field | Required | Description |
|-------|----------|-------------|
| `type` | ✓ | Must be `local` |
| `path` | ✓ | Directory path or glob pattern |

### Remote Skills

Fetched from a GitHub repository at runtime via `npx skills add`:

```yaml
skills:
  - type: remote
    name: azure-keyvault-py
    repo: microsoft/skills
```

| Field | Required | Description |
|-------|----------|-------------|
| `type` | ✓ | Must be `remote` |
| `repo` | ✓ | GitHub repository (`owner/repo` format) |
| `name` | | Skill name within the repository |

> **Tip:** The [microsoft/skills](https://github.com/microsoft/skills) repo contains 132+ skills across Azure SDK scenarios.

## MCP Servers

MCP (Model Context Protocol) servers provide tools to the generator agent:

```yaml
mcp_servers:
  azure:
    type: local
    command: npx
    args: ["-y", "@azure/mcp@latest"]
    tools: ["*"]
```

### MCP Server Fields

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `type` | string | | `local` | Server type (typically `local` for executable commands) |
| `command` | string | ✓ | | Command to execute (e.g., `npx`, `python3`) |
| `args` | list | | `[]` | Command-line arguments |
| `tools` | list | | `[]` | Tools exposed. Use `["*"]` for all tools. |

## Complete Example

A full config demonstrating all options:

```yaml
configs:
  - name: full-example/claude-sonnet-4.5
    description: "Full example — all config options"
    
    generator:
      model: "claude-sonnet-4.5"
      skills:
        - type: remote
          name: azure-keyvault-py
          repo: microsoft/skills
        - type: local
          path: "./skills/generator"
      mcp_servers:
        azure:
          type: local
          command: npx
          args: ["-y", "@azure/mcp@latest"]
          tools: ["*"]
    
    reviewer:
      models:
        - "claude-opus-4.6"
        - "gemini-3-pro-preview"
        - "gpt-4.1"
      skills:
        - type: local
          path: "./skills/reviewer"
```

## Included Configs

The `configs/` directory ships with these configurations:

| Config | Generator Model | MCP | Skills |
|--------|----------------|-----|--------|
| `baseline/claude-sonnet-4.5` | Claude Sonnet 4.5 | — | — |
| `baseline/claude-opus-4.6` | Claude Opus 4.6 | — | — |
| `baseline-skills/claude-opus-4.6` | Claude Opus 4.6 | — | Generator skills |
| `baseline/gpt-codex` | GPT Codex | — | — |
| `azure-mcp/claude-sonnet-4.5` | Claude Sonnet 4.5 | Azure MCP | — |
| `azure-mcp/claude-opus-4.6` | Claude Opus 4.6 | Azure MCP | — |
| `azure-mcp/gpt-codex` | GPT Codex | Azure MCP | — |

All configs use the same 3-model reviewer panel: Claude Opus 4.6 (consolidator), Gemini 3 Pro Preview, and GPT-4.1.

## Backward Compatibility (Legacy Format)

Hyoka supports a legacy flat format that is automatically migrated to the new `generator`/`reviewer` structure during parsing via the `Normalize()` function.

### Legacy Fields → New Location

| Legacy Field | Migrates To |
|-------------|-------------|
| `model` | `generator.model` |
| `reviewer_models` | `reviewer.models` |
| `reviewer_model` | `reviewer.models` (wrapped in list) |
| `mcp_servers` | `generator.mcp_servers` |
| `skill_directories` | `generator.skills` (as `type: local`) |
| `generator_skill_directories` | `generator.skills` (as `type: local`) |
| `reviewer_skill_directories` | `reviewer.skills` (as `type: local`) |
| `available_tools` | `generator.available_tools` |
| `excluded_tools` | `generator.excluded_tools` |

### Example: Legacy Format

```yaml
configs:
  - name: baseline/claude-opus-4.6
    description: "Baseline — raw Copilot"
    model: "claude-opus-4.6"
    reviewer_models:
      - "claude-opus-4.6"
      - "gemini-3-pro-preview"
      - "gpt-4.1"
```

This is equivalent to:

```yaml
configs:
  - name: baseline/claude-opus-4.6
    description: "Baseline — raw Copilot"
    generator:
      model: "claude-opus-4.6"
    reviewer:
      models:
        - "claude-opus-4.6"
        - "gemini-3-pro-preview"
        - "gpt-4.1"
```

Both formats work. The new format is recommended for new configs.

## Validation Rules

Configs are validated at parse time:

1. **At least one config** must be defined
2. **All configs must have a `name`** — used for CLI selection
3. **Skill types must be `local` or `remote`** — with required fields per type
4. **No duplicate reviewer models** — each model appears at most once
5. **Generator model required** — either via `generator.model` or legacy `model`

## Accessor Methods

When working with configs programmatically, use accessor methods which handle legacy-to-new format resolution:

| Method | Returns | Description |
|--------|---------|-------------|
| `EffectiveModel()` | `string` | Generator model |
| `EffectiveReviewerModels()` | `[]string` | Reviewer model list |
| `EffectiveGeneratorSkills()` | `[]Skill` | Generator skills |
| `EffectiveReviewerSkills()` | `[]Skill` | Reviewer skills |
| `EffectiveMCPServers()` | `map[string]*MCPServer` | MCP server configs |
| `EffectiveAvailableTools()` | `[]string` | Allowed tools |
| `EffectiveExcludedTools()` | `[]string` | Blocked tools |
