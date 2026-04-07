---
name: "testing-patterns"
description: "Table-driven tests, -race flag, stub patterns"
domain: "quality"
confidence: "high"
source: "hyoka/internal/*_test.go files"
---

## Context

Hyoka test suite follows Go idioms: table-driven tests for comprehensive coverage, `-race` flag for detecting concurrency issues, and stub patterns for mocking external dependencies (Copilot SDK, file system).

## Table-Driven Tests

Use table-driven tests to test multiple scenarios with the same logic:

**✓ Correct:**
```go
func TestBehaviorGrader_ConstraintValidation(t *testing.T) {
    tests := []struct {
        name      string
        config    map[string]any
        actionLog []ActionEvent
        expect    bool
        wantErr   string
    }{
        {
            name: "required tools present",
            config: map[string]any{
                "required_tools": []any{"read_file", "edit_file"},
            },
            actionLog: []ActionEvent{
                {Tool: "read_file", TurnNumber: 1},
                {Tool: "edit_file", TurnNumber: 2},
            },
            expect: true,
        },
        {
            name: "required tool missing",
            config: map[string]any{
                "required_tools": []any{"read_file"},
            },
            actionLog: []ActionEvent{
                {Tool: "bash", TurnNumber: 1},
            },
            expect: false,
            wantErr: "required tool not found",
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            g, _ := NewBehaviorGrader("test", tt.config)
            result, err := g.Grade(context.Background(), GraderInput{
                ActionLog: tt.actionLog,
            })
            if (err != nil) != (tt.wantErr != "") {
                t.Fatalf("unexpected error: %v", err)
            }
            if result.Pass != tt.expect {
                t.Errorf("expected Pass=%v, got %v", tt.expect, result.Pass)
            }
        })
    }
}
```

## Race Condition Detection

Always run tests with `-race` flag:

```bash
go test -race ./hyoka/...
```

The `-race` detector catches unsynchronized access to shared memory in concurrent code. This is critical for:
- Parallel grader execution
- Multi-worker prompt evaluation
- Action timeline appends

## Stub Patterns for External Dependencies

### Grader Stubs

When testing engine behavior without real graders:

```go
type StubGrader struct {
    gradeFn func(context.Context, GraderInput) (GraderResult, error)
}

func (sg *StubGrader) Grade(ctx context.Context, input GraderInput) (GraderResult, error) {
    return sg.gradeFn(ctx, input)
}

// Usage in test
stubGrader := &StubGrader{
    gradeFn: func(_ context.Context, _ GraderInput) (GraderResult, error) {
        return GraderResult{Pass: true, Score: 1.0}, nil
    },
}
```

### File System Stubs

Mock file operations without touching disk:

```go
type stubFileReader struct {
    readFn func(path string) (string, error)
}

func (s *stubFileReader) Read(path string) (string, error) {
    return s.readFn(path)
}

// Usage
stubFS := &stubFileReader{
    readFn: func(path string) (string, error) {
        if path == "/test/main.py" {
            return "# Generated code", nil
        }
        return "", os.ErrNotExist
    },
}
```

## Test Organization

- **Unit tests:** Test single function/method (`*_test.go` in same package)
- **Integration tests:** Test multi-package workflows (`*_integration_test.go`)
- **Fixtures:** Use `testdata/` directory for test files

## Running Tests

```bash
# All tests
go test ./hyoka/...

# With race detector
go test -race ./hyoka/...

# Specific test
go test -run TestBehaviorGrader ./hyoka/internal/graders

# Verbose with coverage
go test -v -cover ./hyoka/...
```

## Anti-Patterns

- Not using table-driven tests for multiple scenarios
- Ignoring `-race` failures (concurrency bugs are subtle)
- Mocking internal types instead of using interfaces
- Test files that modify package state between runs
- Hardcoding file paths (use testdata or stubs)

## Related Code Locations

- **Example table-driven tests:** `hyoka/internal/graders/*_test.go`
- **Stub patterns:** `hyoka/internal/eval/*_test.go`
- **Test fixtures:** `hyoka/testdata/`
