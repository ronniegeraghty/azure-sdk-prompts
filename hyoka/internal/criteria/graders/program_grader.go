package graders

import (
"bytes"
"context"
"errors"
"fmt"
"os/exec"
"time"
)

// Default timeout for program graders when none is configured.
const defaultProgramTimeout = 30 * time.Second

// ProgramGrader runs a command and grades based on exit code.
// Exit code 0 → Pass, non-zero → Fail.
type ProgramGrader struct {
name    string
command string
args    []string
timeout time.Duration
}

// NewProgramGrader constructs a ProgramGrader from a parsed ProgramConfig.
func NewProgramGrader(name string, cfg *ProgramConfig) (*ProgramGrader, error) {
if name == "" {
return nil, fmt.Errorf("program grader: name is required")
}
if cfg.Command == "" {
return nil, fmt.Errorf("program grader %q: command is required", name)
}

timeout := defaultProgramTimeout
if cfg.Timeout > 0 {
timeout = time.Duration(cfg.Timeout) * time.Second
}

return &ProgramGrader{
name:    name,
command: cfg.Command,
args:    cfg.Args,
timeout: timeout,
}, nil
}

// Kind returns the grader type identifier.
func (g *ProgramGrader) Kind() string { return KindProgram }

// Name returns the grader's name.
func (g *ProgramGrader) Name() string { return g.name }

// Grade executes the configured command in the workspace directory.
// Exit code 0 produces Pass=true/Score=1.0; non-zero produces Pass=false/Score=0.0.
func (g *ProgramGrader) Grade(ctx context.Context, input GraderInput) (GraderResult, error) {
ctx, cancel := context.WithTimeout(ctx, g.timeout)
defer cancel()

var stdout, stderr bytes.Buffer

cmd := exec.CommandContext(ctx, g.command, g.args...)
cmd.Dir = input.WorkspacePath
cmd.Stdout = &stdout
cmd.Stderr = &stderr

result := GraderResult{
Kind:   KindProgram,
Name:   g.name,
Weight: input.Config.EffectiveWeight(),
Gate:   input.Config.Gate,
}

start := time.Now()
err := cmd.Run()
elapsed := time.Since(start)

details := &ProgramGraderDetails{
Command:  g.command,
ExitCode: 0,
Stdout:   stdout.String(),
Stderr:   stderr.String(),
}

if err != nil {
if ctx.Err() != nil {
details.ExitCode = -1
result.Score = 0.0
result.Pass = false
if errors.Is(ctx.Err(), context.DeadlineExceeded) {
result.Message = fmt.Sprintf("command timed out after %s", g.timeout)
} else {
result.Message = fmt.Sprintf("command cancelled: %v", ctx.Err())
}
result.ProgramDetails = details
return result, nil
}

var exitErr *exec.ExitError
if errors.As(err, &exitErr) {
details.ExitCode = exitErr.ExitCode()
result.Score = 0.0
result.Pass = false
result.Message = fmt.Sprintf("command exited with code %d (took %s)", exitErr.ExitCode(), elapsed)
result.ProgramDetails = details
return result, nil
}

return GraderResult{}, fmt.Errorf("program grader %q: %w", g.name, err)
}

result.Score = 1.0
result.Pass = true
result.Message = fmt.Sprintf("command succeeded (took %s)", elapsed)
result.ProgramDetails = details
return result, nil
}
