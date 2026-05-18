# hyoka Examples

This directory contains example configurations, prompts, and criteria demonstrating hyoka's capabilities and recommended patterns.

## Directory Structure

```
examples/
├── configs/          # Example evaluation configurations
├── prompts/          # Example prompt files
├── criteria/         # Example grading criteria
└── README.md         # This file
```

## Examples Overview

### Configs (`examples/configs/`)

Configuration files demonstrate the unified tool system and multi-model review setup:

- **`example-full.yaml`** — Comprehensive example showing all available config options:
  - Generator section with local and remote skills
  - Reviewer section with multiple models
  - Session limits (max_turns, max_files, max_session_actions)
  - Unified `tools:` block (no separate `mcp_servers:` section)

- **`example-generator-skills.yaml`** — Demonstrates loading skills from a local directory:
  - `skill_dir: true` — loads all subdirectories with SKILL.md files
  - `skill_dir: false` (default) — loads a single skill directory

- **`example-remote-skill.yaml`** — Demonstrates remote skill fetching:
  - Whole-repo fetch via `npx skills add`
  - Subpath fetch via `git sparse-checkout`

### Prompts (`examples/prompts/`)

Prompt files demonstrate different prompt formats and features:

- **`prompt-template.prompt.md`** — Template for creating new prompts using Markdown format
- **`example.prompt.yaml`** — Pure YAML prompt format (alternative to Markdown)
- **`existing-files-example.prompt.md`** — Demonstrates `starter_project` feature:
  - `project_context: type: existing` — agent works on pre-existing code
  - `starter_project:` path — files copied to agent workspace before session starts

### Criteria (`examples/criteria/`)

Criteria files define evaluation graders with optional property-based filtering:

- **`language/` subdirectory** — Language-specific criteria:
  - Language-level `when:` condition (applies to all graders in the file)
  - Example: `language/python.yaml`, `language/rust.yaml`, etc.

- **`service/` subdirectory** — Service-specific criteria:
  - Service-level `when:` condition
  - Example: `service/storage.yaml`, `service/key-vault.yaml`

- **`hierarchical-when-example.yaml`** — Demonstrates hierarchical `when` syntax:
  - **File-level `when:`** — Condition applies to all graders in the file
  - **Group-level `when:`** — Condition applies to graders in a section (separated by `---`)
  - **Grader-level `when:`** — Condition applies to individual graders
  - When multiple levels are present, they are AND-ed together

## Architecture Patterns

### Unified Tool System

All example configs use the unified `tools:` block for both generator and reviewer:

```yaml
generator:
  tools:
    - name: azure
      type: mcp
      command: npx
      args: ["-y", "@azure/mcp@latest", "server", "start"]
    - name: reviewer-skills
      type: skill
      source: local
      skill_dir: true
      path: "./skills/reviewer"
```

No separate `mcp_servers:` or `skills:` sections — everything is configured in one place.

### Hierarchical `when` Conditions

Criteria can use `when:` at multiple levels for precise control:

```yaml
# File-level: applies to ALL graders in this file
when:
  language: python

graders:
  - name: "File-level grader"
    prompt: "Applies to all Python prompts"
  
  - name: "Specific grader"
    when:
      category: async-operations
    prompt: "Applies only to Python + async-operations"

---
# Group-level: new section with different file-level when
when:
  language: rust

graders:
  - name: "Rust-specific grader"
    prompt: "Applies to all Rust prompts"
```

## Running Examples

### Validate all examples

```bash
hyoka validate
```

Output:
```
✓ All 89 prompt(s) are valid
✓ All 12 config(s) are valid
✓ All 2 criteria file(s) valid (25 grader(s))
```

### Run an example evaluation

```bash
# Using example config with a specific prompt
go run ./hyoka run \
  --config-file examples/configs/example-full.yaml \
  --config example-full/claude-sonnet-4.5 \
  --language python \
  --dry-run
```

### List prompts matched by criteria

```bash
go run ./hyoka list --language python --service key-vault
```

## Adding New Examples

When adding new example files:

1. **Configs** — Place in `examples/configs/`, use `.yaml` extension
2. **Prompts** — Place in `examples/prompts/`, use `.prompt.md` (Markdown) or `.prompt.yaml` (pure YAML)
3. **Criteria** — Place in `examples/criteria/`, organized by subdirectory (language/, service/, etc.)
4. **Starter Files** — If a prompt uses `starter_project`, follow the naming convention:
   - Prompt file: `examples/prompts/my-prompt.prompt.md`
   - Starters directory: `examples/prompts/my-prompt.starters/`

All examples must validate with `hyoka validate` before committing.

## See Also

- [Configuration Guide](../docs/configuration.md) — Complete config schema
- [Prompt Authoring](../docs/prompt-authoring.md) — Prompt format details
- [Grader Config Schema](../docs/grader-config-schema.md) — Criteria and grader types
