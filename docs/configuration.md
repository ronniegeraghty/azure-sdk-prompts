# Configuration Guide

hyoka uses YAML configuration files to define evaluation setups. Each config specifies a generator model, reviewer models, and tool entries (tools, skills, MCP servers).

## Config Directory

By default, configs are loaded from `./configs/`. Use `--config-dir` to specify a different location.

## Config Names vs Filenames

The `name` field in a config is what you pass to the `--config` CLI flag. It is **not** the filename. For example, a config file called `azure-mcp-opus.yaml` might define `name: azure-mcp/claude-opus-4.6`. You'd run it with: `--config azure-mcp/claude-opus-4.6`.

Current config names and their filenames:

| Filename | Config Name |
|----------|-------------|
| `baseline-sonnet.yaml` | `baseline/claude-sonnet-4.5` |
| `baseline-opus.yaml` | `baseline/claude-opus-4.6` |
| `baseline-opus-skills.yaml` | `baseline-skills/claude-opus-4.6` |
| `baseline-codex.yaml` | `baseline/gpt-5.3-codex` |
| `azure-mcp-sonnet.yaml` | `azure-mcp/claude-sonnet-4.5` |
| `azure-mcp-opus.yaml` | `azure-mcp/claude-opus-4.6` |
| `azure-mcp-codex.yaml` | `azure-mcp/gpt-5.3-codex` |

## Config File Format

A config file contains one or more named configurations:

```yaml
configs:
  - name: baseline/claude-opus-4.6
    generator:
      model: claude-opus-4.6
      tools:
        - name: generator-skills
          type: skill
          source: local
          path: ../skills/generator
    reviewer:
      models:
        - claude-opus-4.6
        - gpt-5.3-codex
        - claude-sonnet-4.5
      tools:
        - name: reviewer-skills
          type: skill
          source: local
          path: ../skills/reviewer
```

## Configuration Fields

### Top-Level

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | yes | Unique name for this configuration |
| `description` | string | no | Human-readable description |

### Generator Section

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `model` | string | required | Model to use for code generation |
| `system_prompt` | string | "" | Custom system instruction for the generator agent |
| `tools` | list | [] | Tool entries (type: tool, skill, or mcp) |
| `excluded_tools` | list | [] | Denylist of tools |

### Reviewer Section

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `model` | string | — | Single reviewer model |
| `models` | list | — | Multiple reviewer models (panel review) |
| `system_prompt` | string | "" | Custom system instruction for reviewer agents |
| `tools` | list | [] | Tool entries (type: skill) for review sessions |

When `models` lists multiple models, hyoka uses a **panel review** where all models review independently and the first model acts as consolidator to produce a consensus result.

### Skills

Skills are Copilot agent instructions packaged as directories containing a `SKILL.md` file. When attached to a generator or reviewer session, the agent receives the skill's content as additional context that guides its behavior.

Skills can be attached to either the **generator** (code generation agent) or the **reviewer** (grading panel agents), or both. They are configured as tool entries (`type: skill`) under the `generator.tools` and `reviewer.tools` fields respectively.

#### Skill Types

There are two ways to load skills: **local** (from the filesystem) and **remote** (fetched from a GitHub repository).

##### Local Skills

Local skills reference a directory on disk. Paths can be absolute or relative to the config file's directory.

```yaml
generator:
  model: claude-opus-4.6
  tools:
    - name: generator-skills
      type: skill
      source: local
      path: ./skills/generator
```

| Field | Required | Description |
|-------|----------|-------------|
| `type` | yes | Must be `"skill"` |
| `source` | yes | Must be `"local"` |
| `path` | yes | Path to a skill directory (absolute or relative) |

**Glob patterns** are supported, letting you load multiple skill directories at once. Only directories are included — files are filtered out.

```yaml
generator:
  tools:
    # Loads every subdirectory under ./skills/generator/
    - name: generator-skills
      type: skill
      source: local
      path: "./skills/generator/*"
```

For example, if `./skills/reviewer/` contains three subdirectories (`code-review-comments/`, `reviewer-build/`, `sdk-version-check/`), the pattern `./skills/reviewer/*` expands to all three.

##### Remote Skills

Remote skills are fetched from a GitHub repository using `npx skills add`. They are cached locally under `.skills-cache/` so subsequent runs don't re-download.

```yaml
generator:
  model: claude-sonnet-4.5
  tools:
    - name: azure-keyvault-py
      type: skill
      source: remote
      repo: microsoft/skills
```

| Field | Required | Description |
|-------|----------|-------------|
| `type` | yes | Must be `"skill"` |
| `source` | yes | Must be `"remote"` |
| `repo` | yes | GitHub repository in `owner/repo` format |
| `name` | no | Specific skill name within the repo |

Under the hood, hyoka runs:

```
npx skills add <repo> --directory .skills-cache/<repo>/<name> [--name <name>]
```

#### Generator vs Reviewer Skills

- **Generator skills** guide the code generation agent — for example, providing SDK usage patterns, coding conventions, or language-specific best practices.
- **Reviewer skills** guide the review panel agents — for example, instructing them to add inline review comments or verify builds.

```yaml
configs:
  - name: my-eval/claude-opus-4.6
    generator:
      model: claude-opus-4.6
      tools:
        - name: generator-skills
          type: skill
          source: local
          path: ./skills/generator
        - name: azure-keyvault-py
          type: skill
          source: remote
          repo: microsoft/skills
    reviewer:
      models:
        - claude-opus-4.6
        - gpt-4.1
      tools:
        - name: reviewer-skills
          type: skill
          source: local
          path: "./skills/reviewer/*"
```

#### Skill Directory Structure

Each skill is a directory containing at minimum a `SKILL.md` file:

```
skills/
├── generator/
│   └── my-sdk-skill/
│       └── SKILL.md        # Instructions for the generator agent
└── reviewer/
    ├── code-review-comments/
    │   └── SKILL.md        # Adds inline REVIEW: comments
    ├── reviewer-build/
    │   └── SKILL.md        # Verifies generated code builds
    └── sdk-version-check/
        └── SKILL.md        # Checks SDK package versions
```

The `SKILL.md` file contains markdown instructions that the Copilot agent receives as context during its session.

#### Legacy Skill Format

Legacy fields (`skill_directories`, `generator_skill_directories`, `reviewer_skill_directories`) are no longer supported. Migrate them to `generator.tools` / `reviewer.tools` entries with `type: skill`.

### MCP Servers

```yaml
generator:
  tools:
    - name: azure
      type: mcp
      command: npx
      args: ["-y", "@azure/mcp@latest", "server", "start"]
      mcp_tools: ["*"]
```

| Field | Required | Description |
|-------|----------|-------------|
| `name` | yes | MCP server identifier |
| `type` | yes | Must be `"mcp"` |
| `command` | yes | Command to launch the MCP server |
| `args` | no | Arguments passed to the command |
| `mcp_tools` | no | Tool filter — `["*"]` for all tools, or a list of specific tool names |

> **Important:** The `mcp_tools` field must be set (typically `["*"]`) for the MCP server's tools to be registered with the agent. Without it, the server starts but its tools won't be available.

## Limits

Guardrail limits can be set at multiple levels with the following resolution order:

**prompt frontmatter > config YAML > CLI flag > engine default**

This allows fine-grained control at the prompt level while maintaining sensible defaults.

| Field | Type | Default | CLI Flag | Description |
|-------|------|---------|----------|-------------|
| `max_turns` | int | 25 | `--max-turns` | Maximum assistant turns per generation |
| `max_files` | int | 50 | `--max-files` | Maximum generated files per evaluation |
| `max_output_size` | string | "1MB" | `--max-output-size` | Maximum total output size (supports KB, MB suffixes) |
| `max_session_actions` | int | 50 | `--max-session-actions` | Maximum actions per Copilot session |

### Config-Level Limits

Set limits in the config YAML file (overridden by CLI flags):

```yaml
configs:
  - name: strict-config
    generator:
      model: claude-opus-4.6
    limits:
      max_turns: 15
      max_files: 30
      max_output_size: "512KB"
      max_session_actions: 25
```

### Prompt-Level Limits

Override limits for a specific prompt via frontmatter (highest priority). This is useful for complex prompts that require more actions or turns:

```yaml
---
id: storage-dp-python-batch
service: storage
language: python
plane: data-plane
category: crud
difficulty: advanced
max_session_actions: 100  # This prompt needs more reasoning steps
max_turns: 40             # Allow more back-and-forth turns
---
```

Prompts with unset limit fields fall through to config > CLI > default, ensuring backward compatibility.

## Multiple Config Files

Place multiple `.yaml` files in the config directory. All are loaded automatically. Use `--config <name>` to select specific configs, or `--all-configs` to run all.

## Tiered Evaluation Criteria

Use `--criteria-dir` to point to a directory of attribute-matched criteria YAML files. These are matched against prompt metadata (language, service, plane) and merged with prompt-specific criteria at review time. See [prompt-authoring.md](prompt-authoring.md) for details.
