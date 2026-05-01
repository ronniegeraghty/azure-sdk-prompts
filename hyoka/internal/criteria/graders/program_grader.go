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
// Per v4 spec: emits single point "exit code 0", with stderr tail in Message on fail.
func (g *ProgramGrader) Grade(ctx context.Context, input GraderInput) (GraderResult, error) {
	ctx, cancel := context.WithTimeout(ctx, g.timeout)
	defer cancel()

	var stdout, stderr bytes.Buffer

	cmd := exec.CommandContext(ctx, g.command, g.args...)
	cmd.Dir = input.WorkspacePath
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	start := time.Now()
	err := cmd.Run()
	elapsed := time.Since(start)

	extras := &ProgramExtras{
		Command:    g.command,
		Args:       g.args,
		ExitCode:   0,
		Stdout:     stdout.String(),
		Stderr:     stderr.String(),
		DurationMs: elapsed.Milliseconds(),
	}

	var checks []GraderCheck
	var msg string

	if err != nil {
		if ctx.Err() != nil {
			extras.ExitCode = -1
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				msg = fmt.Sprintf("command timed out after %s", g.timeout)
			} else {
				msg = fmt.Sprintf("command cancelled: %v", ctx.Err())
			}
			checks = append(checks, GraderCheck{
				Label:   "exit code 0",
				Pass:    false,
				Message: msg,
			})
			return NewResult(KindProgram, g.name, input.Config, checks, msg, &GraderExtras{Program: extras}), nil
		}

		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			extras.ExitCode = exitErr.ExitCode()
			msg = fmt.Sprintf("command exited with code %d (took %s)", exitErr.ExitCode(), elapsed)
			// Include stderr tail in point Message for debugging
			stderrTail := stderr.String()
			if len(stderrTail) > 500 {
				stderrTail = "..." + stderrTail[len(stderrTail)-500:]
			}
			pointMsg := fmt.Sprintf("exited with code %d", exitErr.ExitCode())
			if stderrTail != "" {
				pointMsg = fmt.Sprintf("exited with code %d; stderr: %s", exitErr.ExitCode(), stderrTail)
			}
			checks = append(checks, GraderCheck{
				Label:   "exit code 0",
				Pass:    false,
				Message: pointMsg,
			})
			return NewResult(KindProgram, g.name, input.Config, checks, msg, &GraderExtras{Program: extras}), nil
		}

		return GraderResult{}, fmt.Errorf("program grader %q: %w", g.name, err)
	}

	msg = fmt.Sprintf("command succeeded (took %s)", elapsed)
	checks = append(checks, GraderCheck{
		Label:   "exit code 0",
		Pass:    true,
		Message: "",
	})
	return NewResult(KindProgram, g.name, input.Config, checks, msg, &GraderExtras{Program: extras}), nil
}
