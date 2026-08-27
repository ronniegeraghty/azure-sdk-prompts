package graders

import (
	"context"
	"fmt"
	"strings"
)

// ToolGrader is the canonical tool-perspective grader that evaluates tool usage
// patterns against the action log and environment. Supports exactly four check kinds:
//   - tool_used: a named tool was called (optional min_calls / max_calls)
//   - tool_not_used: a named tool was NOT called
//   - any_from_group: at least one tool from a group was used (optional except: list)
//   - none_from_group: no tool from a group was used (optional except: list)
//
// Groups resolve by entry Name from the tool topology:
//   - skill_dir name → child skills
//   - plugin name → plugin's exported tools
//   - mcp_server name → server's registered tools
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

	// Legacy kind migration errors
	legacyKinds := map[string]string{
		"specific_tool":  "tool_used",
		"any_of_group":   "any_from_group",
		"group_not_used": "none_from_group",
		"turn_limit":     "REMOVED: turn_limit now belongs to the activity grader",
		"min_calls":      "FOLDED INTO tool_used: use tool_used with min_calls/max_calls fields",
		"max_calls":      "FOLDED INTO tool_used: use tool_used with min_calls/max_calls fields",
	}

	if migration, ok := legacyKinds[rule.Kind]; ok {
		return fmt.Errorf("tool grader %q: %s uses deprecated kind %q → %s",
			graderName, pos, rule.Kind, migration)
	}

	switch rule.Kind {
	case "tool_used":
		if strings.TrimSpace(rule.Tool) == "" && !isMCPServerRule(rule) {
			return fmt.Errorf("tool grader %q: %s kind=tool_used requires 'tool' or source=mcp with 'mcp_server'", graderName, pos)
		}
		if rule.MinCalls != nil && *rule.MinCalls < 0 {
			return fmt.Errorf("tool grader %q: %s min_calls must be >= 0", graderName, pos)
		}
		if rule.MaxCalls != nil && *rule.MaxCalls < 0 {
			return fmt.Errorf("tool grader %q: %s max_calls must be >= 0", graderName, pos)
		}
		if rule.MinCalls != nil && rule.MaxCalls != nil && *rule.MinCalls > *rule.MaxCalls {
			return fmt.Errorf("tool grader %q: %s min_calls cannot exceed max_calls", graderName, pos)
		}
		// Validate source field
		if rule.Source != "" && rule.Source != "skill" && rule.Source != "mcp" && rule.Source != "builtin" {
			return fmt.Errorf("tool grader %q: %s source must be one of: skill, mcp, builtin", graderName, pos)
		}
		// Validate mcp_server requires source=mcp
		if rule.MCPServer != "" && rule.Source != "mcp" {
			return fmt.Errorf("tool grader %q: %s mcp_server requires source=mcp", graderName, pos)
		}

	case "tool_not_used":
		if strings.TrimSpace(rule.Tool) == "" && !isMCPServerRule(rule) {
			return fmt.Errorf("tool grader %q: %s kind=tool_not_used requires 'tool' or source=mcp with 'mcp_server'", graderName, pos)
		}
		// Validate source field
		if rule.Source != "" && rule.Source != "skill" && rule.Source != "mcp" && rule.Source != "builtin" {
			return fmt.Errorf("tool grader %q: %s source must be one of: skill, mcp, builtin", graderName, pos)
		}
		// Validate mcp_server requires source=mcp
		if rule.MCPServer != "" && rule.Source != "mcp" {
			return fmt.Errorf("tool grader %q: %s mcp_server requires source=mcp", graderName, pos)
		}

	case "any_from_group", "none_from_group":
		if strings.TrimSpace(rule.Group) == "" {
			return fmt.Errorf("tool grader %q: %s kind=%s requires 'group' field", graderName, pos, rule.Kind)
		}

	default:
		return fmt.Errorf("tool grader %q: %s unknown kind %q (valid: tool_used, tool_not_used, any_from_group, none_from_group)",
			graderName, pos, rule.Kind)
	}
	return nil
}

func (g *ToolGrader) Kind() string { return KindTool }
func (g *ToolGrader) Name() string { return g.name }

func isMCPServerRule(rule ToolCheckRule) bool {
	return rule.Source == "mcp" && strings.TrimSpace(rule.MCPServer) != ""
}

func (g *ToolGrader) Grade(_ context.Context, input GraderInput) (GraderResult, error) {
	toolSet := collectToolSet(input.ActionLog)
	toolCounts := countTools(input.ActionLog)

	var checks []GraderCheck
	for i, rule := range g.checks {
		checkID := fmt.Sprintf("check_%d", i+1)
		check := evaluateToolCheck(rule, checkID, toolSet, toolCounts, input.ActionLog, input.EnvironmentTools)
		checks = append(checks, check)
	}

	msg := summarizeToolChecks(checks)
	return NewResult(KindTool, g.name, input.Config, checks, msg, nil), nil
}

func evaluateToolCheck(rule ToolCheckRule, checkID string, toolSet map[string]bool, toolCounts map[string]int, actionLog []ActionEvent, envTools []EnvironmentTool) GraderCheck {
	switch rule.Kind {
	case "tool_used":
		// Filter by source and mcp_server if specified
		count := countToolsFiltered(actionLog, rule.Tool, rule.Source, rule.MCPServer)
		used := count > 0

		// Apply min/max constraints
		pass := used
		var constraints []string
		var failReason string

		if rule.MinCalls != nil {
			constraints = append(constraints, fmt.Sprintf("min=%d", *rule.MinCalls))
			if count < *rule.MinCalls {
				pass = false
				failReason = fmt.Sprintf("called %d time(s), required >= %d", count, *rule.MinCalls)
			}
		}
		if rule.MaxCalls != nil {
			constraints = append(constraints, fmt.Sprintf("max=%d", *rule.MaxCalls))
			if count > *rule.MaxCalls {
				pass = false
				failReason = fmt.Sprintf("called %d time(s), limit %d", count, *rule.MaxCalls)
			}
		}

		toolLabel := rule.Tool
		if toolLabel == "" {
			toolLabel = "any tool"
		}
		label := fmt.Sprintf("tool_used: %s", toolLabel)
		if rule.Source != "" {
			label += fmt.Sprintf(" (source=%s", rule.Source)
			if rule.MCPServer != "" {
				label += fmt.Sprintf(", server=%s", rule.MCPServer)
			}
			label += ")"
		}
		if len(constraints) > 0 {
			if rule.Source != "" {
				label = strings.TrimSuffix(label, ")")
				label += fmt.Sprintf(", %s)", strings.Join(constraints, ", "))
			} else {
				label += fmt.Sprintf(" (%s)", strings.Join(constraints, ", "))
			}
		}

		message := ""
		if !used && rule.Tool == "" {
			message = fmt.Sprintf("MCP server %q was not used", rule.MCPServer)
		} else if !used {
			message = fmt.Sprintf("tool %q not found", rule.Tool)
		} else if !pass {
			message = failReason
		} else {
			message = fmt.Sprintf("called %d time(s)", count)
		}

		return GraderCheck{
			Label:   label,
			Pass:    pass,
			Message: message,
		}

	case "tool_not_used":
		// Filter by source and mcp_server if specified
		count := countToolsFiltered(actionLog, rule.Tool, rule.Source, rule.MCPServer)
		used := count > 0
		notUsed := !used
		message := ""
		if used && rule.Tool == "" {
			message = fmt.Sprintf("MCP server %q was used (%d time(s))", rule.MCPServer, count)
		} else if used {
			message = fmt.Sprintf("tool %q was used (%d time(s))", rule.Tool, count)
		}
		toolLabel := rule.Tool
		if toolLabel == "" {
			toolLabel = "any tool"
		}
		label := fmt.Sprintf("tool_not_used: %s", toolLabel)
		if rule.Source != "" {
			label += fmt.Sprintf(" (source=%s", rule.Source)
			if rule.MCPServer != "" {
				label += fmt.Sprintf(", server=%s", rule.MCPServer)
			}
			label += ")"
		}
		return GraderCheck{
			Label:   label,
			Pass:    notUsed,
			Message: message,
		}

	case "any_from_group":
		groupTools := resolveGroup(rule.Group, envTools)
		// Filter out exceptions
		if len(rule.Except) > 0 {
			filtered := []string{}
			for _, t := range groupTools {
				if !containsString(rule.Except, t) {
					filtered = append(filtered, t)
				}
			}
			groupTools = filtered
		}

		anyUsed := false
		var usedTools []string
		for _, t := range groupTools {
			if toolSet[t] {
				anyUsed = true
				usedTools = append(usedTools, t)
			}
		}

		label := fmt.Sprintf("any_from_group: %s", rule.Group)
		if len(rule.Except) > 0 {
			label += fmt.Sprintf(" (except: %s)", strings.Join(rule.Except, ", "))
		}

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

	case "none_from_group":
		groupTools := resolveGroup(rule.Group, envTools)
		// Filter out exceptions
		if len(rule.Except) > 0 {
			filtered := []string{}
			for _, t := range groupTools {
				if !containsString(rule.Except, t) {
					filtered = append(filtered, t)
				}
			}
			groupTools = filtered
		}

		noneUsed := true
		var usedTools []string
		for _, t := range groupTools {
			if toolSet[t] {
				noneUsed = false
				usedTools = append(usedTools, t)
			}
		}

		label := fmt.Sprintf("none_from_group: %s", rule.Group)
		if len(rule.Except) > 0 {
			label += fmt.Sprintf(" (except: %s)", strings.Join(rule.Except, ", "))
		}

		msg := ""
		if !noneUsed {
			msg = fmt.Sprintf("used: %s", strings.Join(usedTools, ", "))
		}
		return GraderCheck{
			Label:   label,
			Pass:    noneUsed,
			Message: msg,
		}

	default:
		return GraderCheck{
			Label:   fmt.Sprintf("unknown check kind: %s", rule.Kind),
			Pass:    false,
			Message: fmt.Sprintf("unknown check kind %q", rule.Kind),
		}
	}
}

// Helper functions

func collectToolSet(log []ActionEvent) map[string]bool {
	toolSet := make(map[string]bool)
	for _, e := range log {
		if e.Tool != "" {
			toolSet[e.Tool] = true
		}
	}
	return toolSet
}

// resolveGroup returns the list of tool names that belong to the specified group.
// Groups resolve by entry Name from the tool topology. For now, this is a simple
// implementation that returns all tools (the full Environment Tools will later
// need Parent linkage to properly resolve groups).
func resolveGroup(group string, envTools []EnvironmentTool) []string {
	// TODO: Once EnvironmentTool gains Parent/ParentKind fields (from ToolLoadItem),
	// this should filter by et.Parent == group. For now, return all tools of
	// matching type as a placeholder.
	var tools []string
	for _, et := range envTools {
		tools = append(tools, et.Name)
	}
	return tools
}

func containsString(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func summarizeToolChecks(checks []GraderCheck) string {
	passed := 0
	for _, c := range checks {
		if c.Pass {
			passed++
		}
	}
	return fmt.Sprintf("tool checks: %d/%d passed", passed, len(checks))
}

// countToolsFiltered counts tool calls matching the given name and optional source/mcp_server filters.
// source can be "skill", "mcp", or "builtin" (maps to event.Type).
// mcpServer is only used when source="mcp".
func countToolsFiltered(log []ActionEvent, toolName string, source string, mcpServer string) int {
	count := 0
	for _, e := range log {
		if toolName != "" && e.Tool != toolName {
			continue
		}

		// If source is specified, filter by event.Type
		if source != "" {
			matched := false
			switch source {
			case "skill":
				matched = e.Type == "skill"
			case "mcp":
				matched = e.Type == "mcp_call" || e.MCPServer != ""
			case "builtin":
				// Builtin tools are tool_call, file_read, file_write, bash (not skill or mcp_call)
				matched = e.MCPServer == "" && (e.Type == "tool_call" || e.Type == "file_read" ||
					e.Type == "file_write" || e.Type == "bash")
			}
			if !matched {
				continue
			}
		}

		// If mcpServer is specified (requires source=mcp), filter by e.MCPServer
		if mcpServer != "" && e.MCPServer != mcpServer {
			continue
		}

		count++
	}
	return count
}
