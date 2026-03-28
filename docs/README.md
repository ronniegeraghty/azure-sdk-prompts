# Hyoka Documentation

Welcome to the **hyoka** documentation — a comprehensive reference for the Azure SDK prompt evaluation tool.

## What is Hyoka?

Hyoka is a Go CLI tool that evaluates AI agent code generation quality. It runs curated prompts through the Copilot SDK, reviews generated code via a multi-model review panel, and produces criteria-based pass/fail reports.

## Documentation Index

### Getting Started

- **[Getting Started](getting-started.md)** — Prerequisites, installation, first evaluation run, and common workflows.

### Reference

- **[CLI Reference](cli-reference.md)** — Complete command and flag reference for all 9 commands.
- **[Configuration Guide](configuration.md)** — Config YAML schema, generator/reviewer sections, skills, MCP servers.
- **[Prompt Authoring Guide](prompt-authoring.md)** — How to write prompts, frontmatter schema, evaluation criteria, directory structure.

### Concepts

- **[Architecture Overview](architecture.md)** — End-to-end pipeline, package structure, workspace isolation, session management.
- **[Evaluation Criteria](evaluation-criteria.md)** — General rubric, prompt-specific criteria, scoring methodology, multi-model review panel.
- **[Reports & Trends](reports-and-trends.md)** — Report format, directory structure, trend analysis, regression detection.

### Operations

- **[Guardrails & Safety](guardrails-and-safety.md)** — Turn/file/size limits, safety boundaries, process lifecycle, fan-out confirmation.
- **[Troubleshooting](troubleshooting.md)** — Common issues, diagnostics, debug logging, environment checks.

### Development

- **[Contributing](contributing.md)** — Build, test, add commands, project structure, coding conventions.

---

## Quick Links

| I want to... | Go to |
|---|---|
| Install and run my first eval | [Getting Started](getting-started.md) |
| See all CLI flags | [CLI Reference](cli-reference.md) |
| Write a new prompt | [Prompt Authoring Guide](prompt-authoring.md) |
| Configure models and skills | [Configuration Guide](configuration.md) |
| Understand how scoring works | [Evaluation Criteria](evaluation-criteria.md) |
| Read evaluation reports | [Reports & Trends](reports-and-trends.md) |
| Debug a failed run | [Troubleshooting](troubleshooting.md) |
| Contribute to hyoka | [Contributing](contributing.md) |
