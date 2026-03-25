# Waza POC: KeyVault Data Plane Python Eval Suite

Proof of concept to evaluate whether [Waza](https://github.com/microsoft/waza) can replicate
[hyoka](../tool/)'s Azure SDK code generation evaluation model.

## What This Tests

Three Azure Key Vault Data Plane Python prompts from `prompts/key-vault/data-plane/python/`:

| Task | Hyoka ID | Difficulty |
|------|----------|------------|
| CRUD Secrets | `key-vault-dp-python-crud` | Basic |
| Error Handling | `key-vault-dp-python-error-handling` | Intermediate |
| Pagination | `key-vault-dp-python-pagination` | Intermediate |

## How Hyoka's Model Maps to Waza

### Criteria Structure

| Hyoka Concept | Waza Equivalent | Status |
|---------------|-----------------|--------|
| General rubric (rubric.md) | Spec-level `prompt` + `text` graders in eval.yaml | ✅ Mapped |
| Per-prompt criteria | Task-level `prompt` + `text` graders in task YAML | ✅ Mapped |
| Dual criteria merge | `RunAll()` merges spec + task graders | ✅ Native support |

### Grader Type Mapping

| Hyoka Criterion | Waza Grader | Notes |
|-----------------|-------------|-------|
| Code Builds | `text` (syntax keywords) + `prompt` (LLM judge) | No Python build verification in Waza |
| Latest Packages | `text` (contains azure-keyvault-secrets) | Deterministic check |
| Best Practices | `prompt` (LLM-as-judge) | Works well |
| Error Handling | `prompt` (LLM-as-judge) | Works well |
| Code Quality | `prompt` (LLM-as-judge) | Works well |
| Per-prompt criteria | `prompt` (per-criterion pass/fail calls) | Partial credit via tool call counting |

### What's NOT Replicated

| Hyoka Feature | Waza Status | Impact |
|---------------|-------------|--------|
| **Multi-model review panel** | ❌ Not supported | Hyoka uses 3 reviewers + opus consolidation + majority voting. Waza's `prompt` grader uses single LLM judge. This is hyoka's key differentiator. |
| **Build verification** | ❌ Not built-in | Hyoka runs language-specific compilers. Could use Waza's `program` grader with a Python syntax checker. |
| **MCP server configs** | ⚠️ Partial | `mcp_servers` field exists in Waza spec but isn't plumbed to execution engine. |
| **Config variants** | ⚠️ Manual | Hyoka has 6 configs (baseline × model). Waza uses `--model` flag or `--baseline` for A/B. Need separate eval.yaml files for different skill/MCP configs. |
| **Session transcripts** | ⚠️ Partial | Waza has `--session-log` and `--transcript-dir` but format differs from hyoka. |

## Running

```bash
# Dry run with mock executor (no API calls):
# Edit eval.yaml: change executor to "mock"
waza run waza-test/eval.yaml -v

# Real run (requires GITHUB_TOKEN):
waza run waza-test/eval.yaml --model claude-sonnet-4.5 -v -o results.json

# Multi-model comparison:
waza run waza-test/eval.yaml \
  --model claude-sonnet-4.5 \
  --model claude-opus-4.6 \
  --model gpt-4.1 \
  -o results.json

# With result caching:
waza run waza-test/eval.yaml --cache -v
```

## Directory Structure

```
waza-test/
├── eval.yaml                          # Eval spec (global graders + config)
├── tasks/
│   ├── crud-secrets.yaml              # CRUD operations task
│   ├── error-handling.yaml            # Error handling task
│   └── pagination-list-secrets.yaml   # Pagination task
└── README.md                          # This file
```

## Key Findings

1. **Waza CAN replicate the dual-criteria model** — spec-level graders for general rubric + task-level graders for per-prompt criteria. `RunAll()` merges both naturally.

2. **The `prompt` grader is powerful** — Multi-criteria rubrics with per-criterion pass/fail tool calls give partial credit scoring. This is a good analog for hyoka's criteria-based scoring.

3. **Multi-model review panel is the gap** — Hyoka's 3-reviewer panel with majority voting and opus consolidation is NOT replicable in Waza. You'd need to build a custom grader or run N prompt graders and aggregate externally.

4. **Build verification would need a `program` grader** — Waza's `program` grader could shell out to `python -m py_compile` for syntax verification, but hyoka's multi-language build pipeline is more sophisticated.

5. **`--baseline` flag is a nice Waza feature** — Automatically runs with/without skills, computing skill impact. More elegant than hyoka's manual baseline/MCP config pairs.
