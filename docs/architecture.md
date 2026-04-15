# Hyoka — Project Architecture

## Overview

Hyoka is a Go CLI tool that evaluates AI agents generating Azure SDK code. It sends prompts through the Copilot SDK, collects generated code, then runs a multi-model review panel to score the output.

## Directory Layout

```
hyoka/                         # Go module (github.com/ronniegeraghty/hyoka)
├── cmd/                       # Cobra command definitions (root, run, list, etc.)
└── internal/
    ├── config/                # Loads YAML config files (models, skills, MCP servers)
    │   └── tool/              #   Unified tool entry resolution (skills, MCP, tools)
    ├── prompt/                # Loads prompt .md/.yaml files, parses frontmatter
    │   ├── loader.go          #   Discovers prompts from prompts/ directory tree
    │   ├── parser.go          #   Parses frontmatter + body from markdown
    │   └── types.go           #   Prompt struct definition
    │
    ├── eval/                  # CORE — orchestrates the entire eval pipeline
    │   ├── engine.go          #   Runs: generate → review → report
    │   ├── copilot.go         #   Talks to Copilot SDK — sends prompts, handles events
    │   └── workspace.go       #   Temp dirs, file recovery, cleanup
    │
    ├── process/               # Process lifecycle management
    │   ├── proctracker.go     #   Tracks child processes, kills orphans
    │   ├── resourcemonitor.go #   CPU/memory sampling for Copilot sessions
    │   └── signal_*.go        #   Platform-specific signal handling
    │
    ├── review/                # Multi-model code review panel
    │   ├── reviewer.go        #   Creates reviewer sessions, runs panel, consolidates scores
    │   ├── rubric.go          #   Scoring rubric template injected into review prompts
    │   └── types.go           #   ReviewResult, CriterionResult structs
    │
    ├── report/                # Writes evaluation results to disk
    │   ├── generator.go       #   Writes JSON reports
    │   ├── markdown.go        #   Generates Markdown summary
    │   ├── summary_stats.go   #   Aggregate statistics
    │   └── types.go           #   EvalReport struct
    │
    ├── graders/               # Pluggable grader system (file, program, behavior, etc.)
    ├── comparison/            # Cross-config and cross-run comparison engine
    ├── criteria/              # Tiered evaluation criteria (attribute-matched YAML)
    ├── plugin/                # Composable plugin system (bundles skills + MCP servers)
    ├── pairwise/              # Tool-ablation pairwise expansion
    ├── pidfile/               # PID file management for session tracking
    ├── logging/logging.go     # slog setup, EvalLogger helper, CLI flag integration
    ├── progress/              # Live terminal display (TUI during eval runs)
    ├── trends/                # Cross-run trend analysis
    ├── validate/              # Prompt frontmatter validation
    ├── rerender/              # Re-render Markdown reports from existing JSON
    ├── checkenv/              # Environment prerequisite checks
    ├── clean/                 # Session state & orphan process cleanup
    ├── utils/                 # Shared helpers
    └── serve/                 # Local web server for browsing eval reports (React dashboard)
```

## Eval Pipeline

A single evaluation flows through these stages:

```
 ┌─────────────┐
 │ Prompt + Config │
 └──────┬──────┘
        ▼
 ┌─────────────┐     Copilot SDK session generates code
 │  eval/copilot │──→ files land in isolated workspace
 └──────┬──────┘
        ▼
 ┌─────────────┐     Multiple LLMs score code against rubric
 │  review/      │──→ each model reviews independently,
 └──────┬──────┘     then scores are consolidated
        ▼
 ┌─────────────┐
 │  report/      │──→ JSON + Markdown output
 └─────────────┘
```

## Key Concepts

### Configs
YAML files in `configs/` define an evaluation environment: which model generates code, which models review it, what skills and MCP servers to attach. Multiple configs can run against the same prompts for comparison.

### Prompts
Markdown files in `prompts/` organized by `{language}/{service}/`. Each has YAML frontmatter (`id`, `service`, `language`, `plane`, `category`, `difficulty`) and a body containing the task description plus optional evaluation criteria.

### Multi-Model Review Panel
The review phase sends generated code to multiple LLMs (e.g., Claude, Gemini, GPT) independently. Each scores against a rubric of criteria (pass/fail per criterion). A consolidator model merges the individual reviews into a final consensus score.

### Skills
Copilot skills (SKILL.md files) provide domain knowledge to the generator and reviewer sessions. Skills can be local (in `skills/`) or remote (fetched from GitHub repos before eval starts).

### Project Directory (`.hyoka`)
The `hyoka init` command scaffolds a `.hyoka` project directory containing `configs/`, `prompts/`, `criteria/`, `skills/`, and `reports/` subdirectories. This allows teams to maintain their own evaluation setups outside the main repository.

### Graders
Pluggable grading criteria defined in YAML. Five deterministic grader types are implemented (file, program, behavior, action_sequence, tool_constraint) with an AI prompt grader coming. Graders support property-based applicability (`when` conditions) and gate semantics for hard pass/fail requirements. See [grader-config-schema.md](grader-config-schema.md).

### Guardrails
- **Turn limit**: 25 assistant turns per generation (prevents runaway sessions)
- **File limit**: 50 generated files max
- **Output size limit**: 1 MB total
- **Process tracking**: Child processes are registered and killed on timeout/abort
- **Workspace isolation**: Generated code lands in temp dirs, misplaced files are recovered

## CLI

```bash
hyoka run           # Run evaluations (main command)
hyoka list          # List available prompts
hyoka init          # Scaffold a .hyoka project directory
hyoka compare       # Compare evaluation results between configs/runs
hyoka tools         # List available tools and plugins (alias: hyoka plugins)
hyoka configs       # List available configurations
hyoka validate      # Validate prompt frontmatter
hyoka check-env     # Verify prerequisites
hyoka trends        # Analyze trends across runs
hyoka report        # Re-generate reports from JSON
hyoka serve         # Launch local web UI for browsing reports (React dashboard)
hyoka plugins       # List registered plugins
hyoka clean         # Remove stale session state and orphaned processes
hyoka version       # Print version
```

Key flags for `hyoka run`:
- `--prompt-id` — Run a single prompt
- `--language` / `--service` — Filter prompts
- `--config-file` — Use a specific config
- `--pairwise` / `-P` — Expand configs into pairwise tool-ablation variants
- `--log-level debug` — Structured debug logging
- `--log-file path` — Redirect logs to file
- `--skip-review` — Skip the review phase
- `--criteria-dir` — Directory with tiered criteria YAML files
- `--strict-cleanup` — Fail run if orphaned processes detected after cleanup

### Reports and Action Timeline

Each evaluation produces a report containing:
- **Summary statistics** — pass/fail counts, average scores, per-criteria breakdowns
- **Action timeline** — chronological log of all generator actions (file creates, tool calls, reasoning steps)
- **Generated files** — full content of all files produced by the generator
- **Review results** — individual reviewer scores and the consolidated consensus

Reports are written in JSON and Markdown formats to `reports/<run-id>/`. Use `hyoka serve` to browse reports in a web dashboard.
