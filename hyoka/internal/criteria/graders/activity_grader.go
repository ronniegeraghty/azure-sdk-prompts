package graders

import (
	"context"
	"fmt"
	"strings"
)

// ActivityGrader evaluates agent activity during a session using ActionLog,
// ActionsSummary, and TerminatedBy. Six canonical check kinds:
//   - turn_limit: max turn number ≤ configured max
//   - action_count: TotalActions in [min, max]
//   - tool_call_count: ToolCalls in [min, max]
//   - contains_subsequence: ordered subsequence of tool names
//   - contains_action: events matching {type, tool, contains, excludes} appear
//     within optional [min, max] count bounds (default min=1)
//   - excludes_action: zero events may match {type, tool, contains, excludes}
//   - terminated_by: TerminatedBy matches expectation (equals or not_in)
type ActivityGrader struct {
	name string
	cfg  ActivityConfig
}

// NewActivityGrader constructs an ActivityGrader from a parsed config.
func NewActivityGrader(name string, cfg *ActivityConfig) (*ActivityGrader, error) {
	if cfg == nil {
		return nil, fmt.Errorf("activity grader %q: config is required", name)
	}
	// Backwards compat: alias deprecated min_calls/max_calls onto min/max.
	for i := range cfg.Checks {
		c := &cfg.Checks[i]
		if c.Kind == "contains_action" {
			if c.MinCalls != nil && c.Min == nil {
				c.Min = c.MinCalls
			}
			if c.MaxCalls != nil && c.Max == nil {
				c.Max = c.MaxCalls
			}
		}
	}
	// Validate check structure
	for i, check := range cfg.Checks {
		if err := validateActivityCheck(check); err != nil {
			return nil, fmt.Errorf("activity grader %q check %d: %w", name, i, err)
		}
	}
	return &ActivityGrader{name: name, cfg: *cfg}, nil
}

func validateActivityCheck(check ActivityCheck) error {
	switch check.Kind {
	case "turn_limit":
		if check.Max == nil || *check.Max <= 0 {
			return fmt.Errorf("turn_limit requires max > 0")
		}
	case "action_count":
		if check.Min != nil && *check.Min < 0 {
			return fmt.Errorf("action_count: min must be >= 0")
		}
		if check.Max != nil && *check.Max < 0 {
			return fmt.Errorf("action_count: max must be >= 0")
		}
		if check.Min != nil && check.Max != nil && *check.Min > *check.Max {
			return fmt.Errorf("action_count: min (%d) > max (%d)", *check.Min, *check.Max)
		}
	case "tool_call_count":
		if check.Min != nil && *check.Min < 0 {
			return fmt.Errorf("tool_call_count: min must be >= 0")
		}
		if check.Max != nil && *check.Max < 0 {
			return fmt.Errorf("tool_call_count: max must be >= 0")
		}
		if check.Min != nil && check.Max != nil && *check.Min > *check.Max {
			return fmt.Errorf("tool_call_count: min (%d) > max (%d)", *check.Min, *check.Max)
		}
	case "contains_subsequence":
		if len(check.Tools) == 0 {
			return fmt.Errorf("contains_subsequence requires non-empty tools array")
		}
	case "contains_action":
		// Backwards compat: alias min_calls/max_calls onto min/max.
		if check.MinCalls != nil && check.Min == nil {
			check.Min = check.MinCalls
		}
		if check.MaxCalls != nil && check.Max == nil {
			check.Max = check.MaxCalls
		}
		if check.Type == "" && check.Tool == "" && check.Contains == "" && check.Excludes == "" {
			return fmt.Errorf("contains_action: at least one of type, tool, contains, excludes is required")
		}
		if check.Type != "" && !validActionType(check.Type) {
			return fmt.Errorf("contains_action: invalid type %q", check.Type)
		}
		if check.Min != nil && *check.Min < 0 {
			return fmt.Errorf("contains_action: min must be >= 0")
		}
		if check.Max != nil && *check.Max < 0 {
			return fmt.Errorf("contains_action: max must be >= 0")
		}
		if check.Min != nil && check.Max != nil && *check.Min > *check.Max {
			return fmt.Errorf("contains_action: min (%d) > max (%d)", *check.Min, *check.Max)
		}
	case "excludes_action":
		if check.Type == "" && check.Tool == "" && check.Contains == "" && check.Excludes == "" {
			return fmt.Errorf("excludes_action: at least one of type, tool, contains, excludes is required")
		}
		if check.Type != "" && !validActionType(check.Type) {
			return fmt.Errorf("excludes_action: invalid type %q", check.Type)
		}
	case "terminated_by":
		if check.Equals == "" && len(check.NotIn) == 0 {
			return fmt.Errorf("terminated_by requires either equals or not_in")
		}
		if check.Equals != "" && len(check.NotIn) > 0 {
			return fmt.Errorf("terminated_by: equals and not_in are mutually exclusive")
		}
		validTerminations := map[string]bool{
			"completed": true, "max_actions": true, "max_turns": true,
			"guardrail": true, "error": true,
		}
		if check.Equals != "" && !validTerminations[check.Equals] {
			return fmt.Errorf("terminated_by: invalid equals value %q", check.Equals)
		}
		for _, v := range check.NotIn {
			if !validTerminations[v] {
				return fmt.Errorf("terminated_by: invalid not_in value %q", v)
			}
		}
	default:
		return fmt.Errorf("unknown activity check kind: %q", check.Kind)
	}
	return nil
}

// validActionType reports whether the given event type is a recognized
// ActionEvent.Type value as produced by eval.BuildActionTimeline.
func validActionType(t string) bool {
	switch t {
	case "tool_call", "file_read", "file_write", "bash", "mcp_call", "skill",
		"message", "reasoning", "intent", "warning", "error", "file_change",
		"truncation", "compaction", "turn_start", "turn_end", "abort", "other":
		return true
	}
	return false
}

// Kind returns the grader type identifier.
func (g *ActivityGrader) Kind() string { return "activity" }

// Name returns the human-readable name.
func (g *ActivityGrader) Name() string { return g.name }

// Grade evaluates every configured check and returns a GraderResult.
func (g *ActivityGrader) Grade(_ context.Context, input GraderInput) (GraderResult, error) {
	// If no checks configured, trivially pass
	if len(g.cfg.Checks) == 0 {
		checks := []GraderCheck{{
			Label:   "no checks",
			Pass:    true,
			Message: "no activity checks configured — trivially passed",
		}}
		return NewResult("activity", g.name, input.Config, checks, "no checks configured", &GraderExtras{
			Activity: &ActivityExtras{},
		}), nil
	}

	artifact := input.GeneratorArtifact
	if artifact == nil {
		return GraderResult{}, fmt.Errorf("activity grader %q: GeneratorArtifact is nil", g.name)
	}

	// Compute derived metrics
	maxTurn := maxTurnNumber(input.ActionLog)
	uniqueToolsUsed := uniqueTools(extractToolNames(input.ActionLog))

	var graderChecks []GraderCheck

	for _, check := range g.cfg.Checks {
		switch check.Kind {
		case "turn_limit":
			pass := maxTurn <= *check.Max
			label := fmt.Sprintf("turn_limit (max=%d)", *check.Max)
			msg := fmt.Sprintf("max turn observed: %d", maxTurn)
			graderChecks = append(graderChecks, GraderCheck{
				Label:   label,
				Pass:    pass,
				Message: msg,
			})

		case "action_count":
			totalActions := artifact.ActionsSummary.TotalActions
			pass := true
			var failMsg string
			if check.Min != nil && totalActions < *check.Min {
				pass = false
				failMsg = fmt.Sprintf("total actions %d < min %d", totalActions, *check.Min)
			}
			if check.Max != nil && totalActions > *check.Max {
				pass = false
				failMsg = fmt.Sprintf("total actions %d > max %d", totalActions, *check.Max)
			}
			label := fmt.Sprintf("action_count")
			if check.Min != nil {
				label += fmt.Sprintf(" (min=%d)", *check.Min)
			}
			if check.Max != nil {
				label += fmt.Sprintf(" (max=%d)", *check.Max)
			}
			msg := failMsg
			if pass {
				msg = fmt.Sprintf("total actions: %d", totalActions)
			}
			graderChecks = append(graderChecks, GraderCheck{
				Label:   label,
				Pass:    pass,
				Message: msg,
			})

		case "tool_call_count":
			toolCalls := artifact.ActionsSummary.ToolCalls
			pass := true
			var failMsg string
			if check.Min != nil && toolCalls < *check.Min {
				pass = false
				failMsg = fmt.Sprintf("tool calls %d < min %d", toolCalls, *check.Min)
			}
			if check.Max != nil && toolCalls > *check.Max {
				pass = false
				failMsg = fmt.Sprintf("tool calls %d > max %d", toolCalls, *check.Max)
			}
			label := fmt.Sprintf("tool_call_count")
			if check.Min != nil {
				label += fmt.Sprintf(" (min=%d)", *check.Min)
			}
			if check.Max != nil {
				label += fmt.Sprintf(" (max=%d)", *check.Max)
			}
			msg := failMsg
			if pass {
				msg = fmt.Sprintf("tool calls: %d", toolCalls)
			}
			graderChecks = append(graderChecks, GraderCheck{
				Label:   label,
				Pass:    pass,
				Message: msg,
			})

		case "contains_subsequence":
			actualSequence := extractToolNames(input.ActionLog)
			matchIdx := matchSubsequence(actualSequence, check.Tools)
			fullMatch := matchIdx == len(check.Tools)
			label := fmt.Sprintf("contains_subsequence: %s", strings.Join(check.Tools, " → "))
			var msg string
			if !fullMatch {
				missing := check.Tools[matchIdx:]
				msg = fmt.Sprintf("matched %d/%d; missing: %s", matchIdx, len(check.Tools), strings.Join(missing, ", "))
			} else {
				msg = fmt.Sprintf("subsequence matched")
			}
			graderChecks = append(graderChecks, GraderCheck{
				Label:   label,
				Pass:    fullMatch,
				Message: msg,
			})

		case "contains_action":
			matches := matchActionEvents(input.ActionLog, check)
			count := len(matches)
			min := 1
			if check.Min != nil {
				min = *check.Min
			}
			pass := count >= min
			if check.Max != nil && count > *check.Max {
				pass = false
			}
			label := "contains_action: " + describeActionFilter(check)
			var msg string
			if pass {
				msg = fmt.Sprintf("matched %d event(s)", count)
			} else if check.Max != nil && count > *check.Max {
				msg = fmt.Sprintf("matched %d event(s), max is %d", count, *check.Max)
			} else {
				msg = fmt.Sprintf("matched %d event(s), min is %d", count, min)
			}
			graderChecks = append(graderChecks, GraderCheck{
				Label:   label,
				Pass:    pass,
				Message: msg,
			})

		case "excludes_action":
			matches := matchActionEvents(input.ActionLog, check)
			count := len(matches)
			pass := count == 0
			label := "excludes_action: " + describeActionFilter(check)
			var msg string
			if pass {
				msg = "no matching events"
			} else {
				msg = fmt.Sprintf("found %d disallowed event(s)", count)
			}
			graderChecks = append(graderChecks, GraderCheck{
				Label:   label,
				Pass:    pass,
				Message: msg,
			})

		case "terminated_by":
			terminatedBy := artifact.TerminatedBy
			var pass bool
			var msg string
			if check.Equals != "" {
				pass = terminatedBy == check.Equals
				label := fmt.Sprintf("terminated_by (equals=%s)", check.Equals)
				if !pass {
					msg = fmt.Sprintf("terminated_by: %s (expected: %s)", terminatedBy, check.Equals)
				} else {
					msg = fmt.Sprintf("terminated_by: %s", terminatedBy)
				}
				graderChecks = append(graderChecks, GraderCheck{
					Label:   label,
					Pass:    pass,
					Message: msg,
				})
			} else {
				// not_in check
				pass = true
				for _, forbiddenValue := range check.NotIn {
					if terminatedBy == forbiddenValue {
						pass = false
						break
					}
				}
				label := fmt.Sprintf("terminated_by (not_in=[%s])", strings.Join(check.NotIn, ", "))
				if !pass {
					msg = fmt.Sprintf("terminated_by: %s (forbidden)", terminatedBy)
				} else {
					msg = fmt.Sprintf("terminated_by: %s", terminatedBy)
				}
				graderChecks = append(graderChecks, GraderCheck{
					Label:   label,
					Pass:    pass,
					Message: msg,
				})
			}
		}
	}

	// Compute pass/fail count
	passed := 0
	for _, c := range graderChecks {
		if c.Pass {
			passed++
		}
	}

	msg := fmt.Sprintf("activity checks: %d/%d passed", passed, len(graderChecks))

	// Build extras
	extras := &GraderExtras{
		Activity: &ActivityExtras{
			TotalActions: artifact.ActionsSummary.TotalActions,
			ToolCalls:    artifact.ActionsSummary.ToolCalls,
			TurnCount:    maxTurn,
			Truncated:    artifact.ActionsSummary.Truncated,
			TerminatedBy: artifact.TerminatedBy,
			ToolsUsed:    uniqueToolsUsed,
		},
	}

	return NewResult("activity", g.name, input.Config, graderChecks, msg, extras), nil
}

// Helper functions

func extractToolNames(log []ActionEvent) []string {
	tools := make([]string, 0, len(log))
	for _, e := range log {
		if e.Tool != "" {
			tools = append(tools, e.Tool)
		}
	}
	return tools
}

func matchSubsequence(actual, expected []string) int {
	matchIdx := 0
	for _, action := range actual {
		if matchIdx < len(expected) && action == expected[matchIdx] {
			matchIdx++
		}
	}
	return matchIdx
}

func maxTurnNumber(log []ActionEvent) int {
	max := 0
	for _, e := range log {
		if e.TurnNumber > max {
			max = e.TurnNumber
		}
	}
	return max
}

func countTools(log []ActionEvent) map[string]int {
	counts := make(map[string]int)
	for _, e := range log {
		if e.Tool != "" {
			counts[e.Tool]++
		}
	}
	return counts
}

// matchActionEvents returns the events in log that satisfy the type/tool/contains/excludes
// filter on an ActivityCheck. Empty filter fields match anything.
func matchActionEvents(log []ActionEvent, check ActivityCheck) []ActionEvent {
	var matches []ActionEvent
	for _, e := range log {
		// Type filter
		if check.Type != "" {
			evType := e.Type
			if evType == "" {
				evType = e.Action // legacy events stored the type in Action
			}
			if evType != check.Type {
				continue
			}
		}
		// Tool filter
		if check.Tool != "" && e.Tool != check.Tool {
			continue
		}
		// Determine the haystack for text matching:
		//   - error type → Error
		//   - file_change → Path
		//   - everything else → Text (which carries Output: message/reasoning/intent/warning/tool output)
		evType := e.Type
		if evType == "" {
			evType = e.Action
		}
		var haystack string
		switch evType {
		case "error":
			haystack = e.Error
		case "file_change", "file_read", "file_write":
			if e.Path != "" {
				haystack = e.Path
			} else {
				haystack = e.Text
			}
		default:
			haystack = e.Text
			if haystack == "" && e.Tool != "" {
				haystack = e.Tool
			}
		}
		if check.Contains != "" && !strings.Contains(haystack, check.Contains) {
			continue
		}
		if check.Excludes != "" && strings.Contains(haystack, check.Excludes) {
			continue
		}
		matches = append(matches, e)
	}
	return matches
}

// describeActionFilter renders a short human-readable summary of the filter
// fields on an ActivityCheck for use in check labels.
func describeActionFilter(check ActivityCheck) string {
	var parts []string
	if check.Type != "" {
		parts = append(parts, "type="+check.Type)
	}
	if check.Tool != "" {
		parts = append(parts, "tool="+check.Tool)
	}
	if check.Contains != "" {
		parts = append(parts, fmt.Sprintf("contains=%q", check.Contains))
	}
	if check.Excludes != "" {
		parts = append(parts, fmt.Sprintf("excludes=%q", check.Excludes))
	}
	if len(parts) == 0 {
		return "(any)"
	}
	return strings.Join(parts, ", ")
}

func uniqueTools(tools []string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, t := range tools {
		if t != "" && !seen[t] {
			seen[t] = true
			result = append(result, t)
		}
	}
	return result
}
