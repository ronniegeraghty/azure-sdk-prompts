package graders

import (
"context"
"fmt"
)

// PromptGraderAdapter wraps PromptGrader to satisfy the Grader interface.
type PromptGraderAdapter struct {
inner *PromptGrader
}

// Kind returns the grader type identifier.
func (a *PromptGraderAdapter) Kind() string { return KindPrompt }

// Name returns the grader name.
func (a *PromptGraderAdapter) Name() string { return a.inner.Name }

// Grade adapts the PromptGrader's standalone Grade method to the Grader interface.
// Per v4 spec: emits single point "LLM judge: <rubric>", PromptExtras carries RawScore/MaxScore for display.
func (a *PromptGraderAdapter) Grade(ctx context.Context, input GraderInput) (GraderResult, error) {
	pr, err := a.inner.Grade(ctx, input.WorkspacePath)
	if err != nil {
		return GraderResult{}, fmt.Errorf("prompt grader %q: %w", a.inner.Name, err)
	}

	var checks []GraderCheck
	label := fmt.Sprintf("LLM judge: %s", a.inner.Name)
	pointMsg := ""
	if !pr.Passed {
		// Include reasoning summary on failure (truncate if needed)
		reasoning := pr.Details.Reasoning
		if len(reasoning) > 200 {
			reasoning = reasoning[:200] + "..."
		}
		pointMsg = fmt.Sprintf("score %d/%d: %s", pr.Details.RawScore, pr.Details.MaxScore, reasoning)
	}
	checks = append(checks, GraderCheck{
		Label:   label,
		Pass:    pr.Passed,
		Message: pointMsg,
		Evidence: map[string]string{
			"raw_score": fmt.Sprintf("%d", pr.Details.RawScore),
			"max_score": fmt.Sprintf("%d", pr.Details.MaxScore),
		},
	})

	msg := fmt.Sprintf("LLM review: %d/%d", pr.Details.RawScore, pr.Details.MaxScore)
	extras := &GraderExtras{
		Prompt: &PromptExtras{
			Model:     pr.Details.Model,
			Rubric:    pr.Details.Rubric,
			Reasoning: pr.Details.Reasoning,
			RawScore:  pr.Details.RawScore,
			MaxScore:  pr.Details.MaxScore,
		},
	}

	return NewResult(KindPrompt, a.inner.Name, input.Config, checks, msg, extras), nil
}
