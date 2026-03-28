# Architecture Overview

This document describes how hyoka works end-to-end — from prompt loading through code generation, review, and report production.

## High-Level Pipeline

```
┌──────────────────────────────────────────────────────────────────┐
│                    hyoka run --service storage                   │
└──────────────┬───────────────────────────────────────────────────┘
               │
               ▼
┌──────────────────────────┐     ┌──────────────────────────────┐
│  Load Prompts            │     │  Load Configs                │
│  (prompts/*.prompt.md)   │     │  (configs/*.yaml)            │
│  Filter by flags         │     │  Normalize legacy fields     │
└──────────┬───────────────┘     └──────────┬───────────────────┘
           │                                │
           └──────────┬─────────────────────┘
                      │
                      ▼
        ┌──────────────────────────┐
        │  Build Task Matrix       │
        │  prompts × configs       │
        │  Fan-out confirmation    │
        └──────────┬───────────────┘
                   │
                   ▼
        ┌──────────────────────────┐
        │  Engine.Run()            │
        │  Worker pool (parallel)  │
        │  Session limiter         │
        └──────────┬───────────────┘
                   │
    ┌──────────────┼──────────────┐
    │              │              │
    ▼              ▼              ▼
┌────────┐   ┌────────┐   ┌────────┐
│ Eval 1 │   │ Eval 2 │   │ Eval N │   (one per prompt × config)
└────────┘   └────────┘   └────────┘
```

## Single Evaluation Pipeline

Each evaluation runs through four phases:

### Phase 1: Code Generation

1. Create isolated **generator workspace** (temp directory)
2. Snapshot home and CWD (for file recovery)
3. Send prompt to Copilot SDK with:
   - Generator model (from config)
   - Skills (local + remote, resolved at startup)
   - MCP servers (launched as child processes)
   - Safety system prompt (no real Azure resources unless `--allow-cloud`)
4. Agent generates code files in the workspace
5. Recover misplaced files (agents sometimes write to home or CWD)
6. Check guardrails (turn count, file count, output size)

### Phase 2: Build Verification (optional, `--verify-build`)

1. Detect language from generated files
2. Run language-specific build command:
   - **dotnet:** `dotnet restore && dotnet build`
   - **Python:** `python3 -m py_compile` on all `.py` files
   - **Go:** `go build ./...`
   - **Node.js:** `node --check` on all `.js`/`.mjs` files
   - **Java:** `javac` or Maven/Gradle build
   - **Rust:** `cargo build`
3. Record build result (success/failure, stdout, stderr, duration)

### Phase 3: Multi-Model Review

1. Create isolated **reviewer workspace** (copy of generated code)
2. Run 3 reviewer models **in parallel**, each in its own Copilot session:
   - Send `BuildReviewPrompt()` with original prompt, generated code, rubric, and evaluation criteria
   - Each reviewer returns pass/fail per criterion as JSON
3. First reviewer model **consolidates** results using majority voting:
   - For each criterion: PASS if ≥2 of 3 reviewers marked it passed
   - Merge issues and strengths across all reviewers
4. If consolidation fails, fall back to `averageReview()` (majority voting without LLM synthesis)
5. Record consolidated `ReviewResult` with per-criterion scores

### Phase 4: Report Generation

1. Write per-evaluation reports: `report.json`, `report.html`, `report.md`
2. Preserve generated code files in the report directory
3. After all evaluations complete, write run summary: `summary.json`, `summary.html`, `summary.md`
4. Optionally run trend analysis with AI insights

## Package Structure

```
hyoka/
├── main.go                         # CLI entry point (Cobra commands)
└── internal/
    ├── eval/                       # Evaluation engine
    │   ├── engine.go               # Engine.Run(), worker pool, task dispatch
    │   ├── copilot.go              # CopilotSDKEvaluator — Copilot SDK integration
    │   ├── workspace.go            # Workspace isolation, file recovery
    │   └── proctracker.go          # Process tracking, SIGTERM/SIGKILL cleanup
    ├── config/                     # Configuration system
    │   ├── config.go               # ToolConfig, GeneratorConfig, ReviewerConfig
    │   ├── loader.go               # Load/LoadDir, Parse, Normalize
    │   └── skills.go               # Skill resolution (local/remote)
    ├── prompt/                     # Prompt management
    │   ├── parser.go               # ParsePromptFile — frontmatter + sections
    │   ├── loader.go               # LoadPrompts — directory walking + filtering
    │   └── types.go                # Prompt struct, Filter struct
    ├── review/                     # Code review
    │   ├── reviewer.go             # CopilotReviewer, PanelReviewer, consolidation
    │   └── rubric.md               # Embedded scoring rubric (go:embed)
    ├── report/                     # Report generation
    │   ├── generator.go            # WriteReport, WriteSummary
    │   ├── html.go                 # HTML templates
    │   └── markdown.go             # Markdown templates
    ├── build/                      # Build verification
    │   └── verifier.go             # Language-specific build commands
    ├── trends/                     # Trend analysis
    │   ├── trends.go               # Scan reports, build trend data
    │   └── analysis.go             # AI-powered trend insights
    ├── validate/                   # Prompt validation
    │   └── validate.go             # Schema + naming convention checks
    ├── checkenv/                   # Environment verification
    ├── logging/                    # Structured logging (slog)
    ├── progress/                   # Progress display (live/log/off)
    ├── skills/                     # Skill/plugin resolution
    ├── rerender/                   # Report re-rendering
    ├── manifest/                   # Prompt manifest generation
    ├── history/                    # Historical run queries
    └── utils/                      # Filesystem/string utilities
```

## Key Architectural Patterns

### Multi-Model Review Panel

Three independent reviewer models score code in parallel, preventing self-bias. The first model then consolidates via majority voting:

```
Reviewer 1 (Claude Opus)    → [criterion1: PASS, criterion2: FAIL, ...]
Reviewer 2 (Gemini Pro)     → [criterion1: PASS, criterion2: PASS, ...]
Reviewer 3 (GPT-4.1)        → [criterion1: FAIL, criterion2: PASS, ...]
                               ─────────────────────────────────────────
Consolidated (majority)     → [criterion1: PASS (2/3), criterion2: PASS (2/3), ...]
```

### Workspace Isolation

Each evaluation gets its own temporary workspace. The generator writes files there, and the reviewer gets an independent copy. This prevents:

- Cross-evaluation contamination
- Agents modifying the repository
- File conflicts between parallel evaluations

After generation, `recoverMisplacedFiles()` scans home and CWD for files the agent may have written outside the workspace, moving them back.

### Config-Driven Evaluation

Everything is driven by configuration:

- **Which model generates code** → `generator.model`
- **Which models review code** → `reviewer.models`
- **What tools are available** → `generator.mcp_servers`, `available_tools`
- **What domain knowledge is provided** → `generator.skills`, `reviewer.skills`

This makes it easy to compare how different models, tools, and skills affect code quality.

### Process Lifecycle

hyoka tracks all spawned Copilot processes (child PIDs) and ensures cleanup:

1. `ProcessTracker.Register()` — records each spawned process
2. On completion or SIGINT: `TerminateAll()` sends SIGTERM
3. After 5-second grace period: SIGKILL for stragglers
4. Orphan detection: warns about processes that outlive their session

### Session Management

Each Copilot interaction (generation, verification, review) gets its own SDK session with:

- Independent model selection
- Separate skill loading
- Isolated working directory
- Full event capture (tool calls, messages, token usage)

This prevents self-bias (reviewer can't see generation reasoning) and enables per-phase timeout control.

## Data Flow

```
Input:                                    Output:
prompts/*.prompt.md  ──┐                  reports/{runID}/
configs/*.yaml       ──┤                  ├── summary.json
skills/              ──┤                  ├── summary.html
rubric.md (embedded) ──┤                  ├── summary.md
                       │                  └── results/
                       ▼                      └── {service}/{plane}/{language}/{category}/
                   Engine.Run()                   └── {promptID}/{configName}/
                       │                              ├── report.json
                       ▼                              ├── report.html
                   EvalReport[]                       ├── report.md
                       │                              └── generated-code/
                       ▼
                   RunSummary
```

## Dependencies

- **Copilot SDK** (v0.2.0) — AI agent communication via CLI server mode
- **Cobra** — CLI framework for command structure
- **slog** — Structured logging (Go stdlib)
- **YAML v3** — Config and frontmatter parsing
- **html/template** — Report generation
