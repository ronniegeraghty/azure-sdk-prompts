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
func (a *PromptGraderAdapter) Grade(ctx context.Context, input GraderInput) (GraderResult, error) {
result := GraderResult{
Kind:   KindPrompt,
Name:   a.inner.Name,
Weight: input.Config.EffectiveWeight(),
Gate:   input.Config.Gate,
}

pr, err := a.inner.Grade(ctx, input.WorkspacePath)
if err != nil {
return GraderResult{}, fmt.Errorf("prompt grader %q: %w", a.inner.Name, err)
}

result.Score = pr.Score
result.Pass = pr.Passed
result.Message = fmt.Sprintf("LLM review: %d/%d", pr.Details.RawScore, pr.Details.MaxScore)
result.PromptDetails = &PromptGraderDetails{
Model:     pr.Details.Model,
Rubric:    pr.Details.Rubric,
Reasoning: pr.Details.Reasoning,
RawScore:  pr.Details.RawScore,
MaxScore:  pr.Details.MaxScore,
}

return result, nil
}
