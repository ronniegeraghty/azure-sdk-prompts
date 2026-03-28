# Reports & Trends

This document explains hyoka's report format, directory structure, and trend analysis system.

## Report Directory Structure

Each evaluation run produces a hierarchical report tree:

```
reports/
└── {runID}/                                         # e.g., 20250728-143022
    ├── summary.json                                 # Aggregate run statistics
    ├── summary.html                                 # HTML dashboard
    ├── summary.md                                   # Markdown summary
    └── results/
        └── {service}/
            └── {plane}/
                └── {language}/
                    └── {category}/
                        └── {promptID}/
                            └── {configName}/
                                ├── report.json      # Full evaluation data
                                ├── report.html      # Detailed HTML report
                                ├── report.md        # Markdown report
                                └── generated-code/  # All generated files
```

The `runID` is a timestamp-based identifier (e.g., `20250728-143022`).

## Run Summary

The summary files (`summary.json`, `summary.html`, `summary.md`) contain aggregate statistics across all evaluations in a run.

### Summary Fields

| Field | Type | Description |
|-------|------|-------------|
| `RunID` | string | Timestamp-based run identifier |
| `Timestamp` | string | RFC3339 timestamp |
| `TotalPrompts` | int | Number of unique prompts evaluated |
| `TotalConfigs` | int | Number of configs used |
| `TotalEvals` | int | Total evaluations (prompts × configs) |
| `Passed` | int | Evaluations where all criteria passed |
| `Failed` | int | Evaluations with review failures |
| `Errors` | int | Evaluations with SDK/timeout/setup errors |
| `Duration` | float64 | Total run duration in seconds |
| `Results` | list | All individual `EvalReport` objects |
| `Analysis` | string | AI-generated trend analysis (if enabled) |

### Summary HTML Dashboard

The HTML summary includes:

- **Prompt × Config Matrix** — pass/fail status with scores for every combination
- **Duration Analysis** — min/avg/max per config and per prompt
- **Config Comparison** — side-by-side pass rates
- **Tool Usage** — aggregate tool call statistics
- **Links to detailed reports** — click through to individual evaluations

## Individual Evaluation Report

Each `report.json` contains the complete data for one prompt × config evaluation.

### EvalReport Fields

| Field | Type | Description |
|-------|------|-------------|
| `PromptID` | string | Prompt identifier |
| `ConfigName` | string | Config name used |
| `Timestamp` | string | RFC3339 timestamp |
| `Duration` | float64 | Evaluation duration in seconds |
| `Success` | bool | `true` if all criteria passed and no errors |
| `Error` | string | Error message if evaluation failed |
| `ErrorCategory` | string | Category: `timeout`, `sdk_error`, `review_failure`, `no_files`, `generation_failure` |
| `FailureReason` | string | Human-readable failure explanation |
| `GeneratedFiles` | list | Relative paths of generated code files |
| `Review` | object | Final consolidated `ReviewResult` |
| `ReviewPanel` | list | Individual reviewer results (if panel used) |
| `Build` | object | Build verification result (if `--verify-build`) |
| `ToolUsage` | object | Expected vs actual tool comparison |
| `SessionEvents` | list | Detailed event timeline |
| `Environment` | object | Model, skills, MCP servers, token usage |

### Guardrail Fields

| Field | Type | Description |
|-------|------|-------------|
| `GuardrailMaxTurns` | int | Turn limit that was applied |
| `GuardrailMaxFiles` | int | File count limit that was applied |
| `GuardrailMaxOutputSize` | int64 | Output size limit in bytes |
| `GuardrailAbortReason` | string | Which guardrail triggered (if any) |

### Environment Info

| Field | Type | Description |
|-------|------|-------------|
| `Model` | string | LLM model used for generation |
| `SkillDirectories` | list | Configured skill paths |
| `SkillsLoaded` | list | Skills loaded at session start |
| `SkillsInvoked` | list | Skills actually used during generation |
| `MCPServers` | list | MCP servers available |
| `TotalInputTokens` | int | Total input tokens consumed |
| `TotalOutputTokens` | int | Total output tokens produced |
| `TurnCount` | int | Number of conversation turns |

### Build Result

| Field | Type | Description |
|-------|------|-------------|
| `Language` | string | Detected language |
| `Command` | string | Build command executed |
| `ExitCode` | int | 0 = success |
| `Stdout` | string | Build output |
| `Stderr` | string | Error output |
| `Duration` | duration | Build time |
| `Success` | bool | Whether build succeeded |

### Session Events

The `SessionEvents` array captures the full timeline of the Copilot session:

| Event Field | Description |
|-------------|-------------|
| `Type` | Event type (e.g., `tool.execution_start`, `assistant.message`) |
| `ToolName` | Tool invoked (e.g., `create_file`, `bash`) |
| `ToolArgs` | JSON arguments to the tool |
| `Content` | Message content |
| `Duration` | Event duration in milliseconds |
| `TurnNumber` | Conversation turn index |
| `InputTokens` | Token usage for this event |
| `OutputTokens` | Token usage for this event |

## Report Formats

### JSON (`report.json`)

The canonical data format. Contains all fields listed above. Used as the source of truth for re-rendering and trend analysis.

### HTML (`report.html`)

Rich interactive report with:

- Prompt metadata and configuration details
- Generated code with syntax highlighting
- Per-reviewer scoring breakdown with criterion-level detail
- Consolidated review with pass/fail badges
- Session event timeline (expandable per-reviewer action history)
- Tool usage comparison (expected vs actual)
- Build verification results

### Markdown (`report.md`)

Readable format with tables. Suitable for GitHub issue comments or documentation. Contains the same information as HTML but in plain Markdown.

## Re-Rendering Reports

After updating HTML/Markdown templates, re-render existing reports without re-running evaluations:

```bash
# Re-render a specific run
go run ./hyoka report 20250728-143022

# Re-render all runs
go run ./hyoka report --all
```

This reads `report.json` files and regenerates `report.html` and `report.md` using current templates.

## Trend Analysis

The `trends` command scans all past runs and produces trend reports showing how evaluation quality changes over time.

### Running Trends

```bash
# Generate trend report for all past runs
go run ./hyoka trends

# Filter by prompt or service
go run ./hyoka trends --prompt-id storage-dp-python-crud
go run ./hyoka trends --service storage

# Skip AI analysis
go run ./hyoka trends --no-analyze

# Auto-open HTML report
go run ./hyoka trends --open
```

### Trend Data

Each run contributes entries to the trend dataset:

| Field | Description |
|-------|-------------|
| `RunID` | Which run this entry came from |
| `Timestamp` | When the run occurred |
| `ConfigName` | Which config was used |
| `PromptID` | Which prompt was evaluated |
| `Success` | Whether all criteria passed |
| `Score` | Number of passed criteria |
| `MaxScore` | Total criteria count |
| `Duration` | Evaluation duration |
| `ToolCalls` | Tools used during generation |

### Trend Classifications

Per-prompt trends are classified as:

| Classification | Meaning |
|---------------|---------|
| **Improving** | Pass rate is increasing over time |
| **Regressing** | Pass rate is decreasing over time |
| **Stable** | Pass rate is consistent |
| **Flaky** | Pass rate fluctuates unpredictably |
| **New** | Only one data point exists |

### AI-Powered Insights

When `--analyze` is enabled (default), hyoka sends trend data to a Copilot session for analysis. The AI produces:

- Tool usage patterns and their impact on quality
- Cross-config comparison insights
- Regression detection and root cause hypotheses
- Recommendations for improving evaluation quality
- Resource utilization observations

The analysis text is included in both the trend report and the run summary.

### Trend Report Output

Trend reports are written to `reports/trends/` (configurable with `--output`):

- `trends.json` — raw trend data
- `trends.html` — interactive HTML with charts
- `trends.md` — Markdown summary with tables
