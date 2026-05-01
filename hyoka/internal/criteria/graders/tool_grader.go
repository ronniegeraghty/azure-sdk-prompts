package graders

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
)

// ToolGrader is the unified canonical tool-perspective grader that consolidates
// the functionality of behavior, tool_constraint, and tool_usage graders.
//
// It evaluates a list of check rules against the action log and environment.
// Each rule emits its own GraderCheck. Supported check kinds:
//   - specific_tool: a named tool was used at least once
//   - tool_not_used: a named tool was NOT used
//   - any_of_group: at least one tool from a named group was used
//   - group_not_used: no tool from a group was used
//   - turn_limit: turn count <= max_turns
//   - min_calls: tool was called >= N times
//   - max_calls: tool was called <= N times
//
// Groups are identified by name and can reference:
//   - mcp: all MCP tools
//   - skill_plugin: all skill plugins
//   - skill_repo: skills from a specific repo
//   - tool_name_glob: tools matching a glob pattern
type ToolGrader struct {
	name   string
	checks []ToolCheckRule
}

// NewToolGrader constructs a ToolGrader from a name and ToolConfig.
func NewToolGrader(name string, cfg *ToolConfig) (*ToolGrader, error) {
	if cfg == nil {
		return nil, fmt.Errorf("tool grader %q: config is required", name)
	}
	if len(cfg.Checks) == 0 {
		return nil, fmt.Errorf("tool grader %q: at least one check is required", name)
	}

	for i, rule := range cfg.Checks {
		if err := validateToolCheckRule(rule, i, name); err != nil {
			return nil, err
		}
	}

	return &ToolGrader{name: name, checks: cfg.Checks}, nil
}

func validateToolCheckRule(rule ToolCheckRule, idx int, graderName string) error {
	pos := fmt.Sprintf("checks[%d]", idx)
	switch rule.Kind {
	case "specific_tool", "tool_not_used":
		if strings.TrimSpace(rule.Name) == "" {
			return fmt.Errorf("tool grader %q: %s kind=%s requires name", graderName, pos, rule.Kind)
		}
	case "any_of_group", "group_not_used":
		if strings.TrimSpace(rule.Group) == "" {
			return fmt.Errorf("tool grader %q: %s kind=%s requires group", graderName, pos, rule.Kind)
		}
	case "turn_limit":
		if rule.N <= 0 {
			return fmt.Errorf("tool grader %q: %s kind=turn_limit requires n > 0", graderName, pos)
		}
	case "min_calls", "max_calls":
		if strings.TrimSpace(rule.Name) == "" {
			return fmt.Errorf("tool grader %q: %s kind=%s requires name", graderName, pos, rule.Kind)
		}
		if rule.N < 0 {
			return fmt.Errorf("tool grader %q: %s kind=%s requires n >= 0", graderName, pos, rule.Kind)
		}
	default:
		return fmt.Errorf("tool grader %q: %s unknown kind %q", graderName, pos, rule.Kind)
	}
	return nil
}

func (g *ToolGrader) Kind() string { return KindTool }
func (g *ToolGrader) Name() string { return g.name }

func (g *ToolGrader) Grade(_ context.Context, input GraderInput) (GraderResult, error) {
	toolSet := collectToolSet(input.ActionLog)
	toolCounts := countTools(input.ActionLog)
	maxTurn := maxTurnNumber(input.ActionLog)

	var checks []GraderCheck
	for i, rule := range g.checks {
		checkID := fmt.Sprintf("check_%d", i+1)
		check := evaluateToolCheck(rule, checkID, toolSet, toolCounts, maxTurn, input.EnvironmentTools)
		checks = append(checks, check)
	}

	msg := summarizeToolChecks(checks)
	return NewResult(KindTool, g.name, input.Config, checks, msg, nil), nil
}

func evaluateToolCheck(rule ToolCheckRule, checkID string, toolSet map[string]bool, toolCounts map[string]int, maxTurn int, envTools []EnvironmentTool) GraderCheck {
	switch rule.Kind {
	case "specific_tool":
		used := toolSet[rule.Name]
		return GraderCheck{
			Label:   fmt.Sprintf("tool used: %s", rule.Name),
			Pass:    used,
			Message: ternary(used, "", fmt.Sprintf("tool %q not found", rule.Name)),
		}

	case "tool_not_used":
		notUsed := !toolSet[rule.Name]
		return GraderCheck{
			Label:   fmt.Sprintf("tool not used: %s", rule.Name),
			Pass:    notUsed,
			Message: ternary(notUsed, "", fmt.Sprintf("tool %q was used", rule.Name)),
		}

	case "any_of_group":
		groupTools := resolveGroup(rule.Group, envTools)
		anyUsed := false
		var usedTools []string
		for _, t := range groupTools {
			if toolSet[t] {
				anyUsed = true
				usedTools = append(usedTools, t)
			}
		}
		label := fmt.Sprintf("any tool from group %s used", rule.Group)
		msg := ""
		if !anyUsed {
			msg = fmt.Sprintf("no tool from group %s found", rule.Group)
		} else {
			msg = fmt.Sprintf("used: %s", strings.Join(usedTools, ", "))
		}
		return GraderCheck{
			Label:   label,
			Pass:    anyUsed,
			Message: msg,
		}

	case "group_not_used":
		groupTools := resolveGroup(rule.Group, envTools)
		noneUsed := true
		var usedTools []string
		for _, t := range groupTools {
			if toolSet[t] {
				noneUsed = false
				usedTools = append(usedTools, t)
			}
		}
		label := fmt.Sprintf("no tool from group %s used", rule.Group)
		msg := ""
		if !noneUsed {
			msg = fmt.Sprintf("used: %s", strings.Join(usedTools, ", "))
		}
		return GraderCheck{
			Label:   label,
			Pass:    noneUsed,
			Message: msg,
		}

	case "turn_limit":
		within := maxTurn <= rule.N
		return GraderCheck{
			Label:   fmt.Sprintf("turn limit ≤ %d", rule.N),
			Pass:    within,
			Message: ternary(within, "", fmt.Sprintf("turn count %d exceeds limit %d", maxTurn, rule.N)),
		}

	case "min_calls":
		count := toolCounts[rule.Name]
		ok := count >= rule.N
		return GraderCheck{
			Label:   fmt.Sprintf("tool %s called ≥ %d", rule.Name, rule.N),
			Pass:    ok,
			Message: ternary(ok, "", fmt.Sprintf("called %d time(s), required ≥ %d", count, rule.N)),
		}

	case "max_calls":
		count := toolCounts[rule.Name]
		ok := count <= rule.N
		return GraderCheck{
			Label:   fmt.Sprintf("tool %s called ≤ %d", rule.Name, rule.N),
			Pass:    ok,
			Message: ternary(ok, "", fmt.Sprintf("called %d time(s), limit %d", count, rule.N)),
		}

	default:
		return GraderCheck{
			Label:   fmt.Sprintf("unknown check kind: %s", rule.Kind),
			Pass:    false,
			Message: fmt.Sprintf("unknown check kind %q", rule.Kind),
		}
	}
}

// resolveGroup returns the list of tool names that belong to the specified group.
func resolveGroup(group string, envTools []EnvironmentTool) []string {
	if group == "mcp" {
		var tools []string
		for _, et := range envTools {
			if et.Kind == "mcp" {
				tools = append(tools, et.Name)
			}
		}
		return tools
	}

	if group == "skill_plugin" {
		var tools []string
		for _, et := range envTools {
			if et.Kind == "skill" {
				tools = append(tools, et.Name)
			}
		}
		return tools
	}

	if strings.HasPrefix(group, "skill_repo:") {
		repo := strings.TrimPrefix(group, "skill_repo:")
		var tools []string
		for _, et := range envTools {
			if et.Kind == "skill" && et.Repo == repo {
				tools = append(tools, et.Name)
			}
		}
		return tools
	}

	if strings.HasPrefix(group, "tool_name_glob:") {
		pattern := strings.TrimPrefix(group, "tool_name_glob:")
		var tools []string
		for _, et := range envTools {
			matched, _ := filepath.Match(pattern, et.Name)
			if matched {
				tools = append(tools, et.Name)
			}
		}
		return tools
	}

	return nil
}

func summarizeToolChecks(checks []GraderPoint) string {
	passed := 0
	for _, c := range checks {
		if c.Pass {
			passed++
		}
	}
	return fmt.Sprintf("tool checks: %d/%d passed", passed, len(checks))
}

func ternary(cond bool, ifTrue, ifFalse string) string {
	if cond {
		return ifTrue
	}
	return ifFalse
}
