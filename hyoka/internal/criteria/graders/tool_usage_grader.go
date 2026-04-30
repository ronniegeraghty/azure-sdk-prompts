package graders

import (
	"context"
	"fmt"
	"strings"
)

// ToolUsageGrader scores whether the generator actually invoked the
// MCP servers / skills declared in the eval's environment.
//
// Each configured rule is evaluated against EnvironmentTools (what was
// available) and the SkillsInvoked / MCPServersUsed signals (what was
// actually used):
//
//   - If the rule's target is NOT present in EnvironmentTools, the rule
//     is silently skipped — no point is emitted. This makes it safe to
//     declare the same rule across attribute-matched criteria files
//     regardless of which configs include the relevant tool.
//   - Local skills under skills/generator/ are always treated as not
//     present (the generator dir isn't tested by tool_usage rules).
//   - If the rule IS applicable, one Point is emitted whose Pass flag is
//     true iff the corresponding usage signal indicates the tool was
//     actually invoked at least once.
//
// When zero rules are applicable for the eval, the grader emits a single
// trivially-passing Point labeled "no_applicable_rules" so the result still
// satisfies the "every grader emits ≥ 1 Point" invariant.
type ToolUsageGrader struct {
	name string
	cfg  ToolUsageConfig
}

// NewToolUsageGrader constructs a ToolUsageGrader. cfg may declare zero rules
// (the grader will trivially pass). Each rule must carry a non-empty Type and
// the Type-appropriate identifier (Name for mcp_server / skill_plugin, Repo
// for skill_repo).
func NewToolUsageGrader(name string, cfg *ToolUsageConfig) (*ToolUsageGrader, error) {
	if cfg == nil {
		return nil, fmt.Errorf("tool_usage grader %q: config is required", name)
	}
	for i, r := range cfg.Rules {
		if r.Type == "" {
			return nil, fmt.Errorf("tool_usage grader %q: rules[%d].type is required", name, i)
		}
		switch r.Type {
		case "mcp_server", "skill_plugin":
			if strings.TrimSpace(r.Name) == "" {
				return nil, fmt.Errorf("tool_usage grader %q: rules[%d] type=%s requires name", name, i, r.Type)
			}
		case "skill_repo":
			if strings.TrimSpace(r.Repo) == "" {
				return nil, fmt.Errorf("tool_usage grader %q: rules[%d] type=skill_repo requires repo", name, i)
			}
		default:
			return nil, fmt.Errorf("tool_usage grader %q: rules[%d].type %q is not supported", name, i, r.Type)
		}
		if strings.TrimSpace(r.Expect) == "" {
			return nil, fmt.Errorf("tool_usage grader %q: rules[%d].expect is required", name, i)
		}
	}
	return &ToolUsageGrader{name: name, cfg: *cfg}, nil
}

// Kind returns the grader kind.
func (g *ToolUsageGrader) Kind() string { return KindToolUsage }

// Name returns the grader instance name.
func (g *ToolUsageGrader) Name() string { return g.name }

// Grade evaluates each rule against the input's EnvironmentTools and usage
// signals. Returns a GraderResult whose Points record applicable rules only.
func (g *ToolUsageGrader) Grade(_ context.Context, input GraderInput) (GraderResult, error) {
	skillSet := stringSet(input.SkillsInvoked)
	mcpSet := stringSet(input.MCPServersUsed)

	var points []GraderPoint
	for _, rule := range g.cfg.Rules {
		match := findEnvTool(input.EnvironmentTools, rule)
		if match == nil {
			continue
		}
		if isGeneratorDirSkill(*match) {
			continue
		}
		p := evaluateRule(rule, *match, skillSet, mcpSet)
		points = append(points, p)
	}

	if len(points) == 0 {
		points = []GraderPoint{{
			Label:   "no_applicable_rules",
			Pass:    true,
			Message: "no tool_usage rules were applicable to this environment",
		}}
	}
	msg := summarizePoints(points)
	return NewResult(KindToolUsage, g.name, input.Config, points, msg, nil), nil
}

func stringSet(xs []string) map[string]bool {
	out := make(map[string]bool, len(xs))
	for _, x := range xs {
		out[x] = true
	}
	return out
}

// findEnvTool locates an EnvironmentTool entry that matches the given rule.
// Returns nil when no matching env entry exists.
func findEnvTool(env []EnvironmentTool, rule ToolUsageRule) *EnvironmentTool {
	for i := range env {
		et := env[i]
		switch rule.Type {
		case "mcp_server":
			if et.Kind == "mcp" && et.Name == rule.Name {
				return &env[i]
			}
		case "skill_plugin":
			if et.Kind == "skill" && et.Name == rule.Name {
				return &env[i]
			}
		case "skill_repo":
			if et.Kind != "skill" || et.Repo != rule.Repo {
				continue
			}
			if rule.Skill != "" && et.Name != rule.Skill {
				continue
			}
			return &env[i]
		}
	}
	return nil
}

// isGeneratorDirSkill reports whether et is a local skill under the
// skills/generator/ directory. Such skills are excluded from tool_usage
// scoring because they configure the generator itself rather than being
// tools the generator uses.
func isGeneratorDirSkill(et EnvironmentTool) bool {
	if et.Kind != "skill" {
		return false
	}
	if et.Path == "" {
		return false
	}
	return strings.HasPrefix(et.Path, "skills/generator/") ||
		strings.Contains(et.Path, "/skills/generator/")
}

// evaluateRule produces a single GraderPoint for an applicable rule.
func evaluateRule(rule ToolUsageRule, et EnvironmentTool, skillSet, mcpSet map[string]bool) GraderPoint {
	switch rule.Expect {
	case "at_least_one_tool_call":
		used := mcpSet[et.Name]
		return GraderPoint{
			Label:   fmt.Sprintf("mcp_server:%s tool call recorded", et.Name),
			Pass:    used,
			Message: tcMessage(used, "at least one tool call recorded", "no tool calls recorded"),
		}
	case "any_skill_invoked", "skill_invoked":
		used := skillSet[et.Name]
		return GraderPoint{
			Label:   fmt.Sprintf("skill:%s invoked", et.Name),
			Pass:    used,
			Message: tcMessage(used, "skill was invoked", "skill was not invoked"),
		}
	default:
		return GraderPoint{
			Label:   fmt.Sprintf("%s:%s unknown expectation", rule.Type, et.Name),
			Pass:    false,
			Message: fmt.Sprintf("unknown expect value %q", rule.Expect),
		}
	}
}

func tcMessage(pass bool, okMsg, failMsg string) string {
	if pass {
		return okMsg
	}
	return failMsg
}

func summarizePoints(points []GraderPoint) string {
	passed := 0
	for _, p := range points {
		if p.Pass {
			passed++
		}
	}
	return fmt.Sprintf("tool_usage: %d/%d rules passed", passed, len(points))
}
