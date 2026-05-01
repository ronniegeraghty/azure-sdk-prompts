package graders

import (
"context"
"fmt"
"sort"
"strings"
)

// BehaviorGrader checks the action log for required/forbidden tool usage and
// turn limits.
type BehaviorGrader struct {
name           string
requiredTools  []string
forbiddenTools []string
maxTurns       int
}

// NewBehaviorGrader constructs a BehaviorGrader from a name and config map.
func NewBehaviorGrader(name string, cfg map[string]any) (*BehaviorGrader, error) {
if name == "" {
return nil, fmt.Errorf("behavior grader: name is required")
}
g := &BehaviorGrader{name: name}
if v, ok := cfg["required_tools"]; ok {
tools, err := toStringSlice(v)
if err != nil {
return nil, fmt.Errorf("behavior grader %q: required_tools: %w", name, err)
}
g.requiredTools = tools
}
if v, ok := cfg["forbidden_tools"]; ok {
tools, err := toStringSlice(v)
if err != nil {
return nil, fmt.Errorf("behavior grader %q: forbidden_tools: %w", name, err)
}
g.forbiddenTools = tools
}
if v, ok := cfg["max_turns"]; ok {
turns, err := toInt(v)
if err != nil {
return nil, fmt.Errorf("behavior grader %q: max_turns: %w", name, err)
}
if turns <= 0 {
return nil, fmt.Errorf("behavior grader %q: max_turns must be positive", name)
}
g.maxTurns = turns
}
return g, nil
}

func (g *BehaviorGrader) Kind() string { return KindBehavior }
func (g *BehaviorGrader) Name() string { return g.name }

func (g *BehaviorGrader) Grade(_ context.Context, input GraderInput) (GraderResult, error) {
	toolSet := collectToolSet(input.ActionLog)
	maxTurn := maxTurnNumber(input.ActionLog)
	
	extras := &BehaviorExtras{
		ToolsUsed:    sortedStringKeys(toolSet),
		MaxTurns:     g.maxTurns,
		TotalActions: len(input.ActionLog),
		TurnCount:    maxTurn,
	}
	
	var violations []string
	var checks []GraderCheck
	
	// Per-tool Checks for required tools
	for _, tool := range g.requiredTools {
		present := toolSet[tool]
		label := fmt.Sprintf("tool required: %s", tool)
		var pointMsg string
		if !present {
			extras.MissingTools = append(extras.MissingTools, tool)
			violations = append(violations, fmt.Sprintf("required tool %q not found", tool))
			pointMsg = fmt.Sprintf("required tool %q not found", tool)
		}
		checks = append(checks, GraderCheck{
			Label:   label,
			Pass:    present,
			Message: pointMsg,
		})
	}
	
	// Per-tool Checks for forbidden tools
	for _, tool := range g.forbiddenTools {
		used := toolSet[tool]
		label := fmt.Sprintf("tool forbidden: %s", tool)
		var pointMsg string
		if used {
			extras.ForbiddenUsed = append(extras.ForbiddenUsed, tool)
			violations = append(violations, fmt.Sprintf("forbidden tool %q was used", tool))
			pointMsg = fmt.Sprintf("forbidden tool %q was used", tool)
		}
		checks = append(checks, GraderCheck{
			Label:   label,
			Pass:    !used,
			Message: pointMsg,
		})
	}
	
	// Turn limit Check (if configured)
	if g.maxTurns > 0 {
		within := maxTurn <= g.maxTurns
		label := fmt.Sprintf("turn limit ≤ %d", g.maxTurns)
		var pointMsg string
		if !within {
			extras.TurnLimitHit = true
			violations = append(violations, fmt.Sprintf("turn count %d exceeds limit %d", maxTurn, g.maxTurns))
			pointMsg = fmt.Sprintf("turn count %d exceeds limit %d", maxTurn, g.maxTurns)
		}
		checks = append(checks, GraderCheck{
			Label:   label,
			Pass:    within,
			Message: pointMsg,
		})
	}
	
	extras.Violations = violations
	
	msg := "all behavior constraints satisfied"
	if len(violations) > 0 {
		msg = strings.Join(violations, "; ")
	}
	
	// Handle no-constraints case (per spec: must emit at least one Point)
	if len(checks) == 0 {
		checks = []GraderCheck{{
			Label:   "no constraints",
			Pass:    true,
			Message: "no behavior constraints configured — trivially passed",
		}}
	}
	
	return NewResult(KindBehavior, g.name, input.Config, checks, msg, &GraderExtras{Behavior: extras}), nil
}

// ActionSequenceGrader verifies that the action log contains an expected
// ordered subsequence of actions (by Tool name).
type ActionSequenceGrader struct {
name            string
expectedActions []string
}

// NewActionSequenceGrader constructs an ActionSequenceGrader from a name and config.
func NewActionSequenceGrader(name string, cfg map[string]any) (*ActionSequenceGrader, error) {
if name == "" {
return nil, fmt.Errorf("action_sequence grader: name is required")
}
v, ok := cfg["expected_actions"]
if !ok {
return nil, fmt.Errorf("action_sequence grader %q: expected_actions is required", name)
}
actions, err := toStringSlice(v)
if err != nil {
return nil, fmt.Errorf("action_sequence grader %q: expected_actions: %w", name, err)
}
if len(actions) == 0 {
return nil, fmt.Errorf("action_sequence grader %q: expected_actions must not be empty", name)
}
return &ActionSequenceGrader{name: name, expectedActions: actions}, nil
}

func (g *ActionSequenceGrader) Kind() string { return KindActionSequence }
func (g *ActionSequenceGrader) Name() string { return g.name }

func (g *ActionSequenceGrader) Grade(_ context.Context, input GraderInput) (GraderResult, error) {
	actual := make([]string, 0, len(input.ActionLog))
	for _, e := range input.ActionLog {
		actual = append(actual, e.Tool)
	}
	
	// Match expected sequence against actual
	matchIdx := 0
	for _, action := range actual {
		if matchIdx < len(g.expectedActions) && action == g.expectedActions[matchIdx] {
			matchIdx++
		}
	}
	fullMatch := matchIdx == len(g.expectedActions)
	
	extras := &ActionSequenceExtras{
		ExpectedSequence: g.expectedActions,
		ActualSequence:   actual,
		MatchedActions:   matchIdx,
		ToolsUsed:        uniqueTools(actual),
		TotalActions:     len(input.ActionLog),
	}
	
	// Per-step Points: one Point per expected action position
	var checks []GraderCheck
	for i, expectedTool := range g.expectedActions {
		label := fmt.Sprintf("step %d: expected %s", i+1, expectedTool)
		matched := i < matchIdx
		var pointMsg string
		if !matched {
			// Find what was at this position (if anything)
			if i < len(actual) {
				pointMsg = fmt.Sprintf("got %s", actual[i])
			} else {
				pointMsg = "sequence too short"
			}
		}
		checks = append(checks, GraderCheck{
			Label:   label,
			Pass:    matched,
			Message: pointMsg,
		})
	}
	
	msg := "action sequence fully matched"
	if !fullMatch {
		msg = fmt.Sprintf("matched %d/%d expected actions", matchIdx, len(g.expectedActions))
	}
	
	return NewResult(KindActionSequence, g.name, input.Config, checks, msg, &GraderExtras{ActionSequence: extras}), nil
}

// ToolConstraintGrader enforces granular tool usage constraints:
// required/forbidden tools plus per-tool minimum/maximum call counts.
type ToolConstraintGrader struct {
name      string
required  []string
forbidden []string
minCalls  map[string]int
maxCalls  map[string]int
}

// NewToolConstraintGrader constructs a ToolConstraintGrader from a name and config.
func NewToolConstraintGrader(name string, cfg map[string]any) (*ToolConstraintGrader, error) {
if name == "" {
return nil, fmt.Errorf("tool_constraint grader: name is required")
}
g := &ToolConstraintGrader{name: name}
if v, ok := cfg["required"]; ok {
tools, err := toStringSlice(v)
if err != nil {
return nil, fmt.Errorf("tool_constraint grader %q: required: %w", name, err)
}
g.required = tools
}
if v, ok := cfg["forbidden"]; ok {
tools, err := toStringSlice(v)
if err != nil {
return nil, fmt.Errorf("tool_constraint grader %q: forbidden: %w", name, err)
}
g.forbidden = tools
}
if v, ok := cfg["min_calls"]; ok {
m, err := toStringIntMap(v)
if err != nil {
return nil, fmt.Errorf("tool_constraint grader %q: min_calls: %w", name, err)
}
g.minCalls = m
}
if v, ok := cfg["max_calls"]; ok {
m, err := toStringIntMap(v)
if err != nil {
return nil, fmt.Errorf("tool_constraint grader %q: max_calls: %w", name, err)
}
g.maxCalls = m
}
return g, nil
}

func (g *ToolConstraintGrader) Kind() string { return KindToolConstraint }
func (g *ToolConstraintGrader) Name() string { return g.name }

func (g *ToolConstraintGrader) Grade(_ context.Context, input GraderInput) (GraderResult, error) {
	toolCounts := countTools(input.ActionLog)
	
	extras := &ToolConstraintExtras{
		ToolsUsed:    sortedStringKeys(collectToolSet(input.ActionLog)),
		ToolCounts:   toolCounts,
	}
	
	var violations []string
	var checks []GraderCheck
	
	// Required tools
	for _, tool := range g.required {
		called := toolCounts[tool] > 0
		label := fmt.Sprintf("tool required: %s", tool)
		var pointMsg string
		if !called {
			extras.MissingTools = append(extras.MissingTools, tool)
			violations = append(violations, fmt.Sprintf("required tool %q not called", tool))
			pointMsg = fmt.Sprintf("required tool %q not called", tool)
		}
		checks = append(checks, GraderCheck{
			Label:   label,
			Pass:    called,
			Message: pointMsg,
		})
	}
	
	// Forbidden tools
	for _, tool := range g.forbidden {
		used := toolCounts[tool] > 0
		label := fmt.Sprintf("tool forbidden: %s", tool)
		var pointMsg string
		if used {
			extras.ForbiddenUsed = append(extras.ForbiddenUsed, tool)
			violations = append(violations, fmt.Sprintf("forbidden tool %q called %d time(s)", tool, toolCounts[tool]))
			pointMsg = fmt.Sprintf("forbidden tool %q called %d time(s)", tool, toolCounts[tool])
		}
		checks = append(checks, GraderCheck{
			Label:   label,
			Pass:    !used,
			Message: pointMsg,
		})
	}
	
	// Min calls constraints (sorted for stable Points order)
	for _, tool := range sortedMapKeys(g.minCalls) {
		minCount := g.minCalls[tool]
		ok := toolCounts[tool] >= minCount
		label := fmt.Sprintf("tool %s called ≥ %d", tool, minCount)
		var pointMsg string
		if !ok {
			violations = append(violations, fmt.Sprintf("tool %q called %d time(s), minimum is %d", tool, toolCounts[tool], minCount))
			pointMsg = fmt.Sprintf("tool %q called %d time(s), minimum is %d", tool, toolCounts[tool], minCount)
		}
		checks = append(checks, GraderCheck{
			Label:   label,
			Pass:    ok,
			Message: pointMsg,
			Evidence: map[string]string{
				"actual": fmt.Sprintf("%d", toolCounts[tool]),
				"expected": fmt.Sprintf(">=%d", minCount),
			},
		})
	}
	
	// Max calls constraints (sorted for stable Points order)
	for _, tool := range sortedMapKeys(g.maxCalls) {
		maxCount := g.maxCalls[tool]
		ok := toolCounts[tool] <= maxCount
		label := fmt.Sprintf("tool %s called ≤ %d", tool, maxCount)
		var pointMsg string
		if !ok {
			violations = append(violations, fmt.Sprintf("tool %q called %d time(s), maximum is %d", tool, toolCounts[tool], maxCount))
			pointMsg = fmt.Sprintf("tool %q called %d time(s), maximum is %d", tool, toolCounts[tool], maxCount)
		}
		checks = append(checks, GraderCheck{
			Label:   label,
			Pass:    ok,
			Message: pointMsg,
			Evidence: map[string]string{
				"actual": fmt.Sprintf("%d", toolCounts[tool]),
				"expected": fmt.Sprintf("<=%d", maxCount),
			},
		})
	}
	
	extras.Violations = violations
	extras.ConstraintsMet = len(violations) == 0
	
	msg := "all tool constraints satisfied"
	if len(violations) > 0 {
		msg = strings.Join(violations, "; ")
	}
	
	// Handle no-constraints case (per spec: must emit at least one Point)
	if len(checks) == 0 {
		checks = []GraderCheck{{
			Label:   "no constraints",
			Pass:    true,
			Message: "no tool constraints configured — trivially passed",
		}}
	}
	
	return NewResult(KindToolConstraint, g.name, input.Config, checks, msg, &GraderExtras{ToolConstraint: extras}), nil
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func collectToolSet(log []ActionEvent) map[string]bool {
tools := make(map[string]bool)
for _, e := range log {
if e.Tool != "" {
tools[e.Tool] = true
}
}
return tools
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

func maxTurnNumber(log []ActionEvent) int {
max := 0
for _, e := range log {
if e.TurnNumber > max {
max = e.TurnNumber
}
}
return max
}

func sortedStringKeys(m map[string]bool) []string {
keys := make([]string, 0, len(m))
for k := range m {
keys = append(keys, k)
}
sort.Strings(keys)
return keys
}

func sortedMapKeys(m map[string]int) []string {
keys := make([]string, 0, len(m))
for k := range m {
keys = append(keys, k)
}
sort.Strings(keys)
return keys
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

func toStringSlice(v any) ([]string, error) {
switch val := v.(type) {
case []string:
return val, nil
case []any:
result := make([]string, len(val))
for i, elem := range val {
s, ok := elem.(string)
if !ok {
return nil, fmt.Errorf("element %d must be a string", i)
}
result[i] = s
}
return result, nil
default:
return nil, fmt.Errorf("must be a string array")
}
}

func toInt(v any) (int, error) {
switch val := v.(type) {
case int:
return val, nil
case float64:
return int(val), nil
default:
return 0, fmt.Errorf("must be a number")
}
}

func toStringIntMap(v any) (map[string]int, error) {
switch val := v.(type) {
case map[string]int:
return val, nil
case map[string]any:
result := make(map[string]int, len(val))
for k, elem := range val {
n, err := toInt(elem)
if err != nil {
return nil, fmt.Errorf("key %q: %w", k, err)
}
result[k] = n
}
return result, nil
default:
return nil, fmt.Errorf("must be a map of string to int")
}
}
