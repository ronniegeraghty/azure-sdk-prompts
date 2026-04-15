package review

// CriterionResult holds the pass/fail outcome for a single evaluation criterion.
type CriterionResult struct {
	Name   string `json:"name"`
	Passed bool   `json:"passed"`
	Reason string `json:"reason,omitempty"`
}

// ReviewerResponse is the structured JSON schema that reviewers must return.
// Each criterion from the prompt is judged individually as pass/fail.
type ReviewerResponse struct {
	Criteria []CriterionJudgment `json:"criteria"`
}

// CriterionJudgment holds a single criterion evaluation from one reviewer.
type CriterionJudgment struct {
	Criterion string `json:"criterion"` // original criteria text
	Passed    bool   `json:"passed"`
	Reasoning string `json:"reasoning"`
}

// ReviewScores holds pass/fail results for each evaluation criterion.
type ReviewScores struct {
	Criteria []CriterionResult `json:"criteria"`
}

// PassedCount returns the number of criteria that passed.
func (s ReviewScores) PassedCount() int {
	n := 0
	for _, c := range s.Criteria {
		if c.Passed {
			n++
		}
	}
	return n
}

// TotalCount returns the total number of criteria evaluated.
func (s ReviewScores) TotalCount() int {
	return len(s.Criteria)
}

// AllPassed returns true if every criterion passed.
func (s ReviewScores) AllPassed() bool {
	return s.PassedCount() == s.TotalCount() && s.TotalCount() > 0
}

// ReviewEvent captures a single event from the review Copilot session.
type ReviewEvent struct {
	Type     string  `json:"type"`
	ToolName string  `json:"tool_name,omitempty"`
	ToolArgs string  `json:"tool_args,omitempty"`
	Content  string  `json:"content,omitempty"`
	Result   string  `json:"result,omitempty"`
	Error    string  `json:"error,omitempty"`
	Duration float64 `json:"duration_ms,omitempty"`
}

// ReviewResult holds the full output from an LLM-as-judge code review.
type ReviewResult struct {
	Model        string        `json:"model,omitempty"`
	Scores       ReviewScores  `json:"scores"`
	OverallScore int           `json:"overall_score"` // count of passed criteria
	MaxScore     int           `json:"max_score"`     // total criteria count
	Summary      string        `json:"summary"`
	Issues       []string      `json:"issues"`
	Strengths    []string      `json:"strengths"`
	Events       []ReviewEvent `json:"events,omitempty"`
}
