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

// ProgramGrader runs one or more commands and grades based on exit codes.
// Each check produces one GraderCheck; the overall score is passed_checks / total_checks.
type ProgramGrader struct {
	name   string
	checks []ProgramCheck
}

// NewProgramGrader constructs a ProgramGrader from a parsed ProgramConfig.
func NewProgramGrader(name string, cfg *ProgramConfig) (*ProgramGrader, error) {
	if name == "" {
		return nil, fmt.Errorf("program grader: name is required")
	}
	if len(cfg.Checks) == 0 {
		return nil, fmt.Errorf("program grader %q: at least one check is required", name)
	}
	
	// Validate each check
	for i, check := range cfg.Checks {
		if check.Kind != "command" {
			return nil, fmt.Errorf("program grader %q: check[%d]: unsupported kind %q (only 'command' is supported)", name, i, check.Kind)
		}
		if check.Command == "" {
			return nil, fmt.Errorf("program grader %q: check[%d]: command is required", name, i)
		}
	}

	return &ProgramGrader{
		name:   name,
		checks: cfg.Checks,
	}, nil
}

// Kind returns the grader type identifier.
func (g *ProgramGrader) Kind() string { return KindProgram }

// Name returns the grader's name.
func (g *ProgramGrader) Name() string { return g.name }

// Grade executes all configured command checks in the workspace directory.
// Exit code 0 produces Pass=true for that check; non-zero produces Pass=false.
// Overall score is passed_checks / total_checks.
func (g *ProgramGrader) Grade(ctx context.Context, input GraderInput) (GraderResult, error) {
	var allChecks []GraderCheck
	var checkResults []ProgramCheckResult
	passed := 0

	for i, check := range g.checks {
		timeout := defaultProgramTimeout
		if check.Timeout > 0 {
			timeout = time.Duration(check.Timeout) * time.Second
		}

		checkCtx, cancel := context.WithTimeout(ctx, timeout)
		checkResult, graderCheck := g.runCheck(checkCtx, check, i+1, input.WorkspacePath)
		cancel()

		checkResults = append(checkResults, checkResult)
		allChecks = append(allChecks, graderCheck)
		if graderCheck.Pass {
			passed++
		}
	}

	extras := &ProgramExtras{
		CheckResults: checkResults,
	}

	msg := fmt.Sprintf("%d/%d checks passed", passed, len(g.checks))
	return NewResult(KindProgram, g.name, input.Config, allChecks, msg, &GraderExtras{Program: extras}), nil
}

// runCheck runs a single command check and returns both the detailed result and the grader check.
func (g *ProgramGrader) runCheck(ctx context.Context, check ProgramCheck, checkNum int, workspacePath string) (ProgramCheckResult, GraderCheck) {
	var stdout, stderr bytes.Buffer

	cmd := exec.CommandContext(ctx, check.Command, check.Args...)
	cmd.Dir = workspacePath
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	start := time.Now()
	err := cmd.Run()
	elapsed := time.Since(start)

	result := ProgramCheckResult{
		CheckNumber: checkNum,
		Command:     check.Command,
		Args:        check.Args,
		ExitCode:    0,
		Stdout:      stdout.String(),
		Stderr:      stderr.String(),
		DurationMs:  elapsed.Milliseconds(),
	}

	label := fmt.Sprintf("check %d: %s", checkNum, check.Command)
	if len(check.Args) > 0 {
		label = fmt.Sprintf("check %d: %s %v", checkNum, check.Command, check.Args)
	}

	if err != nil {
		if ctx.Err() != nil {
			result.ExitCode = -1
			msg := "timed out"
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				msg = fmt.Sprintf("timed out after %s", elapsed)
			}
			return result, GraderCheck{
				Label:   label,
				Pass:    false,
				Message: msg,
			}
		}

		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			result.ExitCode = exitErr.ExitCode()
			stderrTail := stderr.String()
			if len(stderrTail) > 500 {
				stderrTail = "..." + stderrTail[len(stderrTail)-500:]
			}
			msg := fmt.Sprintf("exited with code %d", exitErr.ExitCode())
			if stderrTail != "" {
				msg = fmt.Sprintf("exited with code %d; stderr: %s", exitErr.ExitCode(), stderrTail)
			}
			return result, GraderCheck{
				Label:   label,
				Pass:    false,
				Message: msg,
			}
		}

		// Unexpected error
		return result, GraderCheck{
			Label:   label,
			Pass:    false,
			Message: fmt.Sprintf("unexpected error: %v", err),
		}
	}

	return result, GraderCheck{
		Label:   label,
		Pass:    true,
		Message: fmt.Sprintf("exit code 0 (took %s)", elapsed),
	}
}
