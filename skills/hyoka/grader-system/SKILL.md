---
name: "grader-system"
description: "Pluggable grader architecture (6 types, gate semantics)"
domain: "architecture"
confidence: "high"
source: "hyoka/internal/graders/grader.go, hyoka/internal/graders/registry.go"
---

## Context

Hyoka's grading system is pluggable and multi-layered. Six independent grader types inspect generated code and action timelines from different angles, then consolidate into a holistic assessment. Graders are **advisory** — they report findings, they don't gate evaluation completion.

## Grader Architecture

All graders implement:
```go
type Grader interface {
    Kind() string
    Name() string
    Grade(ctx context.Context, input GraderInput) (GraderResult, error)
}
```

### GraderInput
```go
type GraderInput struct {
    Code           string          // Generated code
    Language       string          // e.g., "python"
    ActionLog      []ActionEvent   // Timeline of agent actions
    BuildStatus    string          // "success", "failed", "skipped"
    BuildOutput    string          // Compiler/interpreter output
}
```

### GraderResult
```go
type GraderResult struct {
    Kind    string                  // e.g., "behavior", "lint"
    Name    string                  // Grader instance name
    Pass    bool                    // Critical gate (true = safe to deploy)
    Score   float64                 // 0.0-1.0 numeric score
    Message string                  // Human-readable summary
    Details interface{}             // Type-specific details
}
```

## Six Grader Types

### 1. Behavior Grader
Inspects action timeline for required/forbidden tool usage and turn limits.

```yaml
graders:
  - kind: behavior
    name: tool_compliance
    required_tools: [file_write, read_file]
    forbidden_tools: [rm, sudo]
    max_turns: 25
```

**Details:** `BehaviorGraderDetails` with ToolsUsed, MaxTurns, Violations

### 2. Lint Grader
Runs language-specific linters on generated code.

```yaml
graders:
  - kind: lint
    name: python_lint
    linters: [pylint, black, mypy]
    threshold: 0.8  # Must pass 80% of linters
```

**Details:** `LintGraderDetails` with per-linter pass/fail, warnings

### 3. Build Grader
Verifies code builds (or interprets) without errors.

```yaml
graders:
  - kind: build
    name: cargo_build
```

**Details:** `BuildGraderDetails` with exit code, stderr excerpt

### 4. File Grader
Checks generated file structure (count, naming, organization).

```yaml
graders:
  - kind: file
    name: file_structure
    min_files: 2
    max_files: 50
    required_files: [main.py, tests.py]
```

**Details:** `FileGraderDetails` with file list, violations

### 5. Program Grader
Runs generated code and checks output against expected results.

```yaml
graders:
  - kind: program
    name: integration_test
    test_command: python tests.py
    expected_output: "All tests passed"
```

**Details:** `ProgramGraderDetails` with actual vs. expected output

### 6. Prompt Grader
Uses an LLM to score code against semantic criteria (a.k.a. "LLM-as-judge").

```yaml
graders:
  - kind: prompt
    name: semantic_correctness
    rubric: "Does the code correctly implement the requested feature?"
    model: claude-opus-4.6
```

**Details:** `PromptGraderDetails` with rubric reasoning, score breakdown

## Gate Semantics

**Soft gates (reporting):**
- Graders run independently in parallel
- Timeout on one grader doesn't block others
- Failure on one grader doesn't prevent report generation

**Hard gates (evaluation completion):**
- If generation or review phases **hard-fail** (e.g., timeout, SDK crash), eval stops
- Grader failures do NOT stop evaluation (graders are advisory)

## Pluggable Registry

Graders are registered via factory functions:

```go
type GraderFactory func(name string, cfg map[string]any) (Grader, error)

var registry = map[string]GraderFactory{
    "behavior": NewBehaviorGrader,
    "lint":     NewLintGrader,
    "build":    NewBuildGrader,
    "file":     NewFileGrader,
    "program":  NewProgramGrader,
    "prompt":   NewPromptGrader,
}

// New grader types can be added by updating registry
```

## Configuration

Graders are defined in config YAML:

```yaml
graders:
  - kind: behavior
    name: required_tools
    required_tools: [file_write, bash]
  
  - kind: lint
    name: python_style
    linters: [pylint]
    threshold: 0.9

  - kind: prompt
    name: correctness
    model: gpt-5.4
```

## Error Handling

Each grader catches its own errors:

```go
func (g *LintGrader) Grade(ctx context.Context, input GraderInput) (GraderResult, error) {
    // Run linter
    cmd := exec.CommandContext(ctx, "pylint", ...)
    
    // Timeout?
    if ctx.Err() != nil {
        return GraderResult{
            Pass: false,
            Message: "Linter timeout",
        }, nil  // Return error object, not error value
    }
}
```

Grader errors are **not** fatal — they're reported in the grader result.

## Code Locations

- **Grader interface and registry:** `hyoka/internal/graders/grader.go`
- **Individual grader implementations:** `hyoka/internal/graders/{kind}_grader.go`
- **Example grader tests:** `hyoka/internal/graders/*_test.go`

## Anti-Patterns

- Using grader failures as eval blockers (they're advisory only)
- Hardcoding grader lists in engine (add via config)
- Ignoring grader timeout errors (report them)
- Assuming all graders finish synchronously (they may timeout)
