// Package pairwise generates config variants for pairwise tool-ablation testing.
// Given a base config with N togglable tools, it produces N+1 variants:
// one baseline with all tools enabled, and N variants each with one tool disabled.
package pairwise

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/ronniegeraghty/hyoka/internal/config"
)

// ExpandPairwise generates N+1 config variants from a base config.
//   - Variant 0 (baseline): all tools enabled, named "{base}/baseline"
//   - Variant 1..N: each disables one togglable tool, named "{base}/without-{tool}"
//
// Tools with AlwaysOn: true in Generator.Tools are never toggled.
// Both Generator.Tools ([]ToolEntry) and Generator.AvailableTools ([]string)
// are considered; duplicates across the two lists are unified.
func ExpandPairwise(base config.ToolConfig) []config.ToolConfig {
	togglable := collectTogglable(base)

	variants := make([]config.ToolConfig, 0, len(togglable)+1)

	// Baseline: deep copy with all tools intact.
	baseline := cloneToolConfig(base)
	baseline.Name = base.Name + "/baseline"
	variants = append(variants, baseline)

	// One variant per togglable tool, with that tool removed.
	for _, tool := range togglable {
		v := cloneToolConfig(base)
		v.Name = fmt.Sprintf("%s/without-%s", base.Name, tool)
		removeTool(&v, tool)
		variants = append(variants, v)
	}

	return variants
}

// collectTogglable returns deduplicated item names eligible for toggling.
// Order: ToolEntry names first (preserving order), then AvailableTools,
// then MCP servers (sorted alphabetically), then generator skills.
// MCP servers are prefixed with "mcp:" and skills with "skill:" to
// distinguish them from regular tools.
func collectTogglable(cfg config.ToolConfig) []string {
	if cfg.Generator == nil {
		return nil
	}

	seen := make(map[string]bool)
	var items []string

	for _, te := range cfg.Generator.Tools {
		if te.AlwaysOn || seen[te.Name] {
			continue
		}
		seen[te.Name] = true
		items = append(items, te.Name)
	}

	for _, name := range cfg.Generator.AvailableTools {
		if seen[name] {
			continue
		}
		seen[name] = true
		items = append(items, name)
	}

	// MCP servers sorted for deterministic variant order.
	mcpKeys := make([]string, 0, len(cfg.Generator.MCPServers))
	for k := range cfg.Generator.MCPServers {
		mcpKeys = append(mcpKeys, k)
	}
	sort.Strings(mcpKeys)
	for _, k := range mcpKeys {
		key := "mcp:" + k
		if !seen[key] {
			seen[key] = true
			items = append(items, key)
		}
	}

	// Generator skills.
	for _, s := range cfg.Generator.Skills {
		key := "skill:" + skillKey(s)
		if key == "skill:" || seen[key] {
			continue
		}
		seen[key] = true
		items = append(items, key)
	}

	return items
}

// skillKey returns a stable identifier for a Skill entry.
func skillKey(s config.Skill) string {
	if s.Path != "" {
		return s.Path
	}
	return s.Name
}

// removeTool removes a named item from the config. Items prefixed with
// "mcp:" are removed from Generator.MCPServers; items prefixed with
// "skill:" are removed from Generator.Skills; all others are removed
// from Generator.Tools and Generator.AvailableTools.
func removeTool(cfg *config.ToolConfig, name string) {
	if cfg.Generator == nil {
		return
	}

	if strings.HasPrefix(name, "mcp:") {
		delete(cfg.Generator.MCPServers, strings.TrimPrefix(name, "mcp:"))
		return
	}

	if strings.HasPrefix(name, "skill:") {
		id := strings.TrimPrefix(name, "skill:")
		var kept []config.Skill
		for _, s := range cfg.Generator.Skills {
			if skillKey(s) != id {
				kept = append(kept, s)
			}
		}
		cfg.Generator.Skills = kept
		return
	}

	var tools []config.ToolEntry
	for _, te := range cfg.Generator.Tools {
		if te.Name != name {
			tools = append(tools, te)
		}
	}
	cfg.Generator.Tools = tools

	var avail []string
	for _, t := range cfg.Generator.AvailableTools {
		if t != name {
			avail = append(avail, t)
		}
	}
	cfg.Generator.AvailableTools = avail
}

// cloneToolConfig returns a deep copy of a ToolConfig so mutations to the
// clone never affect the original.
func cloneToolConfig(src config.ToolConfig) config.ToolConfig {
	dst := src // copies value fields (Name, Description, Plugins header)

	if src.Generator != nil {
		gen := *src.Generator

		if len(src.Generator.Skills) > 0 {
			gen.Skills = make([]config.Skill, len(src.Generator.Skills))
			copy(gen.Skills, src.Generator.Skills)
		}

		if len(src.Generator.Tools) > 0 {
			gen.Tools = make([]config.ToolEntry, len(src.Generator.Tools))
			copy(gen.Tools, src.Generator.Tools)
			for i, te := range gen.Tools {
				if te.When != nil {
					m := make(map[string]string, len(te.When))
					for k, v := range te.When {
						m[k] = v
					}
					gen.Tools[i].When = m
				}
			}
		}

		if len(src.Generator.AvailableTools) > 0 {
			gen.AvailableTools = make([]string, len(src.Generator.AvailableTools))
			copy(gen.AvailableTools, src.Generator.AvailableTools)
		}

		if len(src.Generator.ExcludedTools) > 0 {
			gen.ExcludedTools = make([]string, len(src.Generator.ExcludedTools))
			copy(gen.ExcludedTools, src.Generator.ExcludedTools)
		}

		if len(src.Generator.MCPServers) > 0 {
			gen.MCPServers = make(map[string]*config.MCPServer, len(src.Generator.MCPServers))
			for k, v := range src.Generator.MCPServers {
				srv := *v
				if len(v.Args) > 0 {
					srv.Args = make([]string, len(v.Args))
					copy(srv.Args, v.Args)
				}
				if len(v.Tools) > 0 {
					srv.Tools = make([]string, len(v.Tools))
					copy(srv.Tools, v.Tools)
				}
				gen.MCPServers[k] = &srv
			}
		}

		dst.Generator = &gen
	}

	if src.Reviewer != nil {
		rev := *src.Reviewer
		if len(src.Reviewer.Skills) > 0 {
			rev.Skills = make([]config.Skill, len(src.Reviewer.Skills))
			copy(rev.Skills, src.Reviewer.Skills)
		}
		if len(src.Reviewer.Models) > 0 {
			rev.Models = make([]string, len(src.Reviewer.Models))
			copy(rev.Models, src.Reviewer.Models)
		}
		dst.Reviewer = &rev
	}

	if len(src.Plugins) > 0 {
		dst.Plugins = make([]string, len(src.Plugins))
		copy(dst.Plugins, src.Plugins)
	}

	return dst
}

// VariantResult holds the evaluation outcome for a single pairwise variant.
type VariantResult struct {
ConfigName  string `json:"config_name"`
RemovedTool string `json:"removed_tool,omitempty"`
Score       int    `json:"score"`
MaxScore    int    `json:"max_score"`
Success     bool   `json:"success"`
}

// ToolImpact holds the computed impact of a single tool.
type ToolImpact struct {
ToolName      string  `json:"tool_name"`
Impact        float64 `json:"impact"`
BaselineScore float64 `json:"baseline_score"`
WithoutScore  float64 `json:"without_score"`
BaselinePass  bool    `json:"baseline_pass"`
WithoutPass   bool    `json:"without_pass"`
}

// PairwiseReport holds the complete pairwise comparison results for a prompt.
type PairwiseReport struct {
PromptID string          `json:"prompt_id"`
Baseline VariantResult   `json:"baseline"`
Variants []VariantResult `json:"variants"`
Impacts  []ToolImpact    `json:"impacts"`
}

// normalizeScore converts a raw score/maxScore pair to a 0-100 scale.
func normalizeScore(score, maxScore int) float64 {
if maxScore <= 0 {
return 0
}
return math.Round(float64(score)/float64(maxScore)*1000) / 10
}

// ComputeImpacts calculates per-tool impact from a baseline and a set of variants.
func ComputeImpacts(promptID string, results []VariantResult) (*PairwiseReport, error) {
var baseline *VariantResult
var variants []VariantResult

for i := range results {
if results[i].RemovedTool == "" {
baseline = &results[i]
} else {
variants = append(variants, results[i])
}
}

if baseline == nil {
return nil, fmt.Errorf("pairwise: no baseline found for prompt %s", promptID)
}

baselineNorm := normalizeScore(baseline.Score, baseline.MaxScore)

impacts := make([]ToolImpact, 0, len(variants))
for _, v := range variants {
withoutNorm := normalizeScore(v.Score, v.MaxScore)
impacts = append(impacts, ToolImpact{
ToolName:      v.RemovedTool,
Impact:        math.Round((baselineNorm-withoutNorm)*10) / 10,
BaselineScore: baselineNorm,
WithoutScore:  withoutNorm,
BaselinePass:  baseline.Success,
WithoutPass:   v.Success,
})
}

SortByImpact(impacts)

return &PairwiseReport{
PromptID: promptID,
Baseline: *baseline,
Variants: variants,
Impacts:  impacts,
}, nil
}

// SortByImpact sorts impacts by impact score descending.
func SortByImpact(impacts []ToolImpact) {
sort.Slice(impacts, func(i, j int) bool {
if impacts[i].Impact != impacts[j].Impact {
return impacts[i].Impact > impacts[j].Impact
}
return impacts[i].ToolName < impacts[j].ToolName
})
}

// AggregateImpacts merges impacts from multiple prompts into a single per-tool summary.
func AggregateImpacts(reports []*PairwiseReport) []ToolImpact {
type accum struct {
totalImpact   float64
totalBaseline float64
totalWithout  float64
count         int
baselinePass  int
withoutPass   int
}

byTool := make(map[string]*accum)
for _, r := range reports {
for _, imp := range r.Impacts {
a, ok := byTool[imp.ToolName]
if !ok {
a = &accum{}
byTool[imp.ToolName] = a
}
a.totalImpact += imp.Impact
a.totalBaseline += imp.BaselineScore
a.totalWithout += imp.WithoutScore
a.count++
if imp.BaselinePass {
a.baselinePass++
}
if imp.WithoutPass {
a.withoutPass++
}
}
}

result := make([]ToolImpact, 0, len(byTool))
for tool, a := range byTool {
result = append(result, ToolImpact{
ToolName:      tool,
Impact:        math.Round(a.totalImpact/float64(a.count)*10) / 10,
BaselineScore: math.Round(a.totalBaseline/float64(a.count)*10) / 10,
WithoutScore:  math.Round(a.totalWithout/float64(a.count)*10) / 10,
BaselinePass:  a.baselinePass == a.count,
WithoutPass:   a.withoutPass == a.count,
})
}

SortByImpact(result)
return result
}
