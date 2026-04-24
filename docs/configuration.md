# Configuration Guide

hyoka uses YAML configuration files to define evaluation setups. Each config specifies a generator model, reviewer models, and tool entries (tools, skills, MCP servers).

## Config Directory

By default, configs are loaded from `./configs/`. Use `--config-dir` to specify a different location.

## Custom Prompt Directory

By default, hyoka looks for prompts in this order:

1. `./.hyoka/prompts/` (created by `hyoka init`)
2. `./prompts/` (legacy fallback)
3. `../prompts/` (legacy fallback)

You can override this by setting `prompt_directory:` at the **top level** of any config YAML file:

```yaml
prompt_directory: ../shared-prompt-library
configs:
  - name: baseline/claude-opus-4.6
    generator:
      model: claude-opus-4.6
```

Notes:

- The path is resolved **relative to the config file** that contains it (so `../shared-prompt-library` from `.hyoka/configs/foo.yaml` points at `.hyoka/../shared-prompt-library`). Absolute paths are honored as-is.
- When loading multiple config files via `--config-dir`, only one file may set `prompt_directory`; conflicting values across files are an error.
- The `--prompts` CLI flag still wins over the config-driven value, so a one-off `--prompts ./other` takes precedence.
- If you don't set `prompt_directory`, behavior is identical to previous releases — existing repos require no changes.

Resolution priority: `--prompts` flag → `prompt_directory:` in config YAML → `.hyoka/prompts/` → `./prompts/` → `../prompts/`.

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

### Plugins

Plugins bundle related skills, MCP servers, and hooks into a single reusable YAML definition. Declare plugins as tool entries with `type: plugin` under `generator.tools` or `reviewer.tools`. Each plugin can contain:

- **Skills** — Local or remote Copilot skills
- **MCP Servers** — MCP server configurations
- **Hooks** — Pre/post tool-use hooks for custom logic

#### Plugin Declaration

```yaml
generator:
  tools:
    - name: azure-sdk-python
      type: plugin
      source: remote
      repo: microsoft/skills
    - name: my-plugin
      type: plugin
      source: local
```

| Field | Required | Default | Description |
|-------|----------|---------|-------------|
| `name` | yes | — | Plugin identifier (filename without `.yaml` for local; the plugin folder name within the repo for remote) |
| `type` | yes | — | Must be `"plugin"` |
| `source` | yes | — | `local` or `remote` |
| `repo` | for `source: remote` | — | Source repository. Canonical form is `owner/repo` (e.g. `microsoft/skills`); GitHub is assumed, so the `github.com/` prefix is redundant but accepted for backward compatibility. hyoka has no implicit marketplace — declare it explicitly. |
| `version` | no | repo default | Git ref (branch, tag, or commit) to pin |

#### Local Plugins

Local plugins are YAML files placed in the `.hyoka/plugins/` directory. Each plugin file defines a bundle of tools (skills + MCPs).

```yaml
generator:
  tools:
    - name: my-plugin
      type: plugin
      source: local
```

hyoka resolves this by looking for `.hyoka/plugins/my-plugin.yaml`. If found, the plugin is parsed and its child tools (skills and MCP servers) are registered with the generator session.

#### Remote Plugins

Remote plugins live in a Git repository (commonly the GitHub Copilot CLI plugin marketplace at `microsoft/skills`, but any repo following the same layout works). The `repo:` field tells hyoka exactly where to fetch from.

```yaml
generator:
  tools:
    - name: azure-sdk-python
      type: plugin
      source: remote
      repo: microsoft/skills
```

hyoka resolves remote plugins from `~/.hyoka/cache/default/<owner>/<repo>/...` (populated by your prior `/plugin install` or by hyoka's fetch flow). To pin to a specific git ref, add `version: <branch-tag-or-sha>`.

> **No implicit marketplace.** Earlier versions accepted a bare `name: foo@skills` shorthand that resolved to `microsoft/skills`. That magic has been removed — every remote plugin entry must declare `repo:` explicitly.

#### Dual-Role Plugins

If you want a plugin available to both the generator and reviewer environments, declare it in both `generator.tools` and `reviewer.tools`:

```yaml
generator:
  tools:
    - name: my-plugin
      type: plugin
      source: local
reviewer:
  tools:
    - name: my-plugin
      type: plugin
      source: local
```

There is no automatic sharing between generator and reviewer tools — each environment receives only the tools explicitly declared for it.

#### Hard-Fail Semantics

- **Fetch errors** (remote plugins) fail before session creation, preventing partial or incorrect evaluations.
- **Missing tools** in a plugin are reported at pre-session validation time (`EventToolsVerified`), not during eval.
- The evaluation aborts immediately if any plugin fails to load.

### Tool Load Validation

All tools declared in `generator.tools` or `reviewer.tools` are **implicitly required**. Before attempting code generation, hyoka performs a **pre-session static validation** to resolve every declared plugin, skill directory, and MCP server. If any tool fails to resolve, the evaluation aborts immediately with a `tool_load_failure` error — the generator is never invoked.

#### Hard-Fail Contract

hyoka hard-fails (with `error_category: "tool_load_failure"`) when:

- **Plugin not found** — A configured plugin is not in the plugin registry and hasn't been installed locally
- **Skill path missing** — A local skill's `path` points to a non-existent directory
- **Missing `SKILL.md`** — A configured skill directory exists but doesn't contain a required `SKILL.md` file
- **Empty skill directory** — A skill directory glob pattern (`glob: "..."`) matches zero SKILL.md files
- **MCP server unavailable** — The `command` (for local) or `url` (for remote) specified for an MCP server is invalid or unreachable

#### Tools Progress Output

During `hyoka run`, each configured tool is resolved and reported in the **Tools** progress section. Plugins and skill directories are **expanded into their children** — each child (individual skill or tool from a plugin) reports its load status individually:

```
Tools:
  ✓ Loaded      azure-sdk-python (plugin)
  ✓ Loaded        ├── skill-azure-sdk-patterns
  ✓ Loaded        └── mcp-azure-resource-tools
  ✓ Loaded      generator-skills (skills dir)
  ✓ Loaded        ├── coding-standards
  ✗ Failed         └── sdk-version-check (missing SKILL.md)
  ✓ Loaded      azure-mcp
```

This grouped display lets you see at a glance which children succeeded and which failed, making diagnostics faster.

#### Error Reporting in EvalReport JSON

When a tool load fails, the evaluation report includes:

```json
{
  "error": "tool_load_failure: skill directory './skills/gen' missing SKILL.md in ./skills/gen/my-skill",
  "error_category": "tool_load_failure",
  "error_details": "skill directory './skills/gen' missing SKILL.md in ./skills/gen/my-skill"
}
```

#### Config Scoping

Reviewer tool validation is **scoped per-config**. When running multiple configs in a single `hyoka run` invocation (e.g., `--config "config-a,config-b"`), each config's `reviewer.tools` are validated independently. This prevents skills configured for one config from accidentally leaking into another.

#### Diagnosing Tool Load Failures

Re-run your evaluation with debug logging:

```bash
hyoka run --prompt-id <prompt-id> --config <config> \
  --log-level debug --log-file hyoka-debug.log
```

Then check the report and logs for details:

```bash
# View the error in the eval report
cat reports/<eval-id>/report.json | jq '.error, .error_category'

# Grep logs for tool-resolution details
grep -i "tool\|skill\|mcp" hyoka-debug.log
```

#### Future: Optional Tools

In a future release, you may be able to mark specific tools as optional using `optional: true`. For now, all configured tools are required.

## Limits

Guardrail limits can be set at multiple levels with the following resolution order:

**prompt frontmatter > config YAML > CLI flag > engine default**

This allows fine-grained control at the prompt level while maintaining sensible defaults.

| Field | Type | Default | CLI Flag | Description |
|-------|------|---------|----------|-------------|
| `max_turns` | int | 25 | `--max-turns` | Maximum assistant turns per generation |
| `max_files` | int | 50 | `--max-files` | Maximum generated files per evaluation |
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
      max_session_actions: 25
```

### Prompt-Level Limits

Override limits for a specific prompt via frontmatter (highest priority). This is useful for complex prompts that require more actions or turns:

```yaml
---
id: storage-dp-python-batch
properties:
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

## Tool Versioning & Custom Fetchers

Remote skills (and other future remote tools) are pinned and fetched through hyoka's pluggable **Fetcher** system. By default, hyoka uses an `npx skills add`-backed fetcher that follows the repo's default branch.

### Pinning versions

Use `tool_version_override` at the top of any config file to pin tools by `name`:

```yaml
tool_version_override:
  azure-sdk-java: "v1.4.2"      # git tag
  copilot-knowledge: "main"      # branch
  azsdk-samples: "abc123def"     # commit SHA

configs:
  - name: baseline/sonnet
    generator:
      model: claude-sonnet-4.5
      tools:
        - name: azure-sdk-java
          type: skill
          source: remote
          repo: Azure/azure-sdk-skills
        - name: pinned-by-entry
          type: skill
          source: remote
          repo: x/y
          version: "v2.0.0"      # per-entry version always wins
```

**Resolution order:** per-entry `version:` field > `tool_version_override` map > fetcher default (latest).

The version is forwarded to the fetcher; the default `npx` fetcher appends it as a git ref (`repo@version`) and caches each version under a separate directory (`.skills-cache/<version>/<repo>/<name>/`) so toggling pins doesn't poison the cache.

`tool_version_override` maps are scoped to the config file they live in. Conflicting values across files in a directory load are an error.

### Custom fetchers

The `tool.Fetcher` interface (in `hyoka/internal/config/tool/`) lets embedders register additional fetchers — useful for internal mirror caches, alternate package managers, or custom version-pinning rules:

```go
import "github.com/ronniegeraghty/hyoka/hyoka/internal/config/tool"

type artifactoryFetcher struct{ /* ... */ }

func (artifactoryFetcher) Name() string { return "artifactory" }
func (artifactoryFetcher) CanFetch(e tool.Entry) bool {
    return e.ResolvedType() == tool.TypeSkill && e.Source == "remote" &&
           strings.HasPrefix(e.Repo, "internal/")
}
func (artifactoryFetcher) Fetch(ctx context.Context, req tool.FetchRequest) (tool.FetchResult, error) {
    // download from internal Artifactory, return FetchResult{Dir, Version}
}

func init() { _ = tool.Register(artifactoryFetcher{}) }
```

**Lookup order:** custom fetchers are consulted before the built-in `npx` default; the first whose `CanFetch` returns true wins. This means custom fetchers can shadow the default for specific entries while leaving everything else on the default path.

Hyoka calls `tool.ValidateFetchers(...)` at the start of every run, so any remote skill without a matching registered fetcher fails fast — before a session is spawned.

## Multiple Config Files

Place multiple `.yaml` files in the config directory. All are loaded automatically. Use `--config <name>` to select specific configs, or `--all-configs` to run all.

## Evaluation Criteria

Use `--criteria-dir` to point to a directory of criteria YAML files. These define graders that evaluate generated outputs across multiple dimensions: file output, build success, LLM-based code review, tool usage constraints, and more.

**Grader types:** Each criteria file contains a list of graders. Graders are matched against prompt metadata (language, service, plane) and evaluated independently. See [graders/index.md](graders/index.md) for complete reference and type documentation.

**Criteria files are merged** across the criteria directory. Graders can be conditional (via `when:` properties), weighted for scoring, and configured per-type with specific checks.

For authoring criteria files, hierarchical organization, and advanced patterns, see [criteria-authoring.md](criteria-authoring.md).

## Plugins

**For declaring plugins in your config, see [Plugins](#plugins) under Configuration Fields above.**

Plugin YAML definitions are stored in the `.hyoka/plugins/` directory for local plugins. Each plugin file (e.g., `.hyoka/plugins/my-plugin.yaml`) declares:

- **Skills** — Local or remote Copilot skills
- **MCP Servers** — MCP server configurations
- **Hooks** — Pre/post tool-use hooks for custom logic

Use `hyoka tools` (or `hyoka plugins`) to list all discovered plugins, skills, and MCP servers.
