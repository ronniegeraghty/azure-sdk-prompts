# Hyoka — Project Architecture

> **📖 This page is a quick reference.** For the full architecture documentation, see [Architecture Overview](architecture-overview.md).

## Overview

Hyoka is a Go CLI tool that evaluates AI agents generating Azure SDK code. It sends prompts through the Copilot SDK, collects generated code, optionally builds it, then runs a multi-model review panel to score the output.

## Eval Pipeline

```
Prompt + Config → Copilot Generation → Build Verification → Multi-Model Review → Reports
```

Each evaluation creates isolated workspaces, runs three reviewer models in parallel, and consolidates scores using majority voting.

## Documentation

- **[Architecture Overview](architecture-overview.md)** — Full pipeline, package structure, workspace isolation
- **[CLI Reference](cli-reference.md)** — All commands and flags
- **[Configuration Guide](configuration.md)** — Config YAML schema, skills, MCP servers
- **[Prompt Authoring](prompt-authoring.md)** — Frontmatter schema, evaluation criteria
- **[Evaluation Criteria](evaluation-criteria.md)** — Scoring methodology, review panel
- **[Reports & Trends](reports-and-trends.md)** — Report format, trend analysis
- **[Guardrails & Safety](guardrails-and-safety.md)** — Limits, safety boundaries
- **[Troubleshooting](troubleshooting.md)** — Common issues, diagnostics
- **[Contributing](contributing.md)** — Build, test, extend
