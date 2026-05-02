// Package pairwise generates config variants for pairwise tool-ablation testing.
// Given a base config with N togglable tools, it produces N+1 variants:
// one baseline with all tools enabled, and N variants each with one tool disabled.
package pairwise

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/plugin"
)

// ExpandPairwise generates N+1 config variants from a base config.
//   - Variant 0 (baseline): all tools enabled, named "{base}/baseline"
//   - Variant 1..N: each disables one togglable tool, named "{base}/without-{tool}"
//
// Tools with AlwaysOn: true in Generator.Tools are never toggled.
// Entries marked pairwise: off are excluded. MCP entries can opt into deep
// toggling, which expands their mcp_tools list into individual variants.
// Skill_dir entries with pairwise: deep expand each subdirectory skill into
// its own variant.
//
// The baseDir parameter is used to resolve relative paths in skill_dir entries.
func ExpandPairwise(base config.ToolConfig, baseDir string) []config.ToolConfig {
	togglable := collectTogglable(base, baseDir)

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

// collectTogglable returns deduplicated tool names eligible for toggling.
// For skill_dir entries with pairwise: deep, each subdirectory skill is
// enumerated and returned as "{entry-name}/{skill-name}".
func collectTogglable(cfg config.ToolConfig, baseDir string) []string {
	if cfg.Generator == nil {
		return nil
	}

	seen := make(map[string]bool)
	var tools []string

	for _, te := range cfg.Generator.Tools {
		if te.ResolvedPairwise() == "off" {
			continue
		}
		// MCP deep mode: enumerate mcp_tools
		if te.ResolvedPairwise() == "deep" && te.ResolvedType() == "mcp" {
			if len(te.MCPTools) == 0 || containsWildcard(te.MCPTools) {
				if !seen[te.Name] {
					seen[te.Name] = true
					tools = append(tools, te.Name)
				}
				continue
			}
			for _, tool := range te.MCPTools {
				name := fmt.Sprintf("%s/%s", te.Name, tool)
				if seen[name] {
					continue
				}
				seen[name] = true
				tools = append(tools, name)
			}
			continue
		}
		// Skill_dir deep mode: enumerate subdirectory skills
		if te.ResolvedPairwise() == "deep" && te.ResolvedType() == "skill" && te.SkillDir {
			skills := enumerateSkillDir(te, baseDir)
			for _, skill := range skills {
				name := fmt.Sprintf("%s/%s", te.Name, skill)
				if seen[name] {
					continue
				}
				seen[name] = true
				tools = append(tools, name)
			}
			continue
		}
		// Plugin deep mode: enumerate plugin tools (skills and MCP servers)
		if te.ResolvedPairwise() == "deep" && te.ResolvedType() == "tool" {
			pluginTools := enumeratePluginTools(te, baseDir)
			for _, toolName := range pluginTools {
				name := fmt.Sprintf("%s/%s", te.Name, toolName)
				if seen[name] {
					continue
				}
				seen[name] = true
				tools = append(tools, name)
			}
			continue
		}
		// Shallow mode (default): toggle the entire entry
		if !seen[te.Name] {
			seen[te.Name] = true
			tools = append(tools, te.Name)
		}
	}

	return tools
}

// enumerateSkillDir returns the list of skill subdirectory names in the given
// skill_dir entry. Each subdirectory containing SKILL.md is treated as a skill.
// The returned names are subdirectory basenames (e.g., "markdown-headings").
func enumerateSkillDir(entry config.ToolEntry, baseDir string) []string {
	if entry.Path == "" {
		return nil
	}

	// Resolve path relative to baseDir
	path := entry.Path
	if !filepath.IsAbs(path) {
		path = filepath.Join(baseDir, path)
	}

	// Read directory
	entries, err := os.ReadDir(path)
	if err != nil {
		// Silently skip — resolution failures are handled elsewhere
		return nil
	}

	var skills []string
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		subDir := filepath.Join(path, e.Name())
		if _, err := os.Stat(filepath.Join(subDir, "SKILL.md")); err == nil {
			skills = append(skills, e.Name())
		}
	}

	return skills
}

// enumeratePluginTools returns the list of tool names (skills + MCP servers)
// exposed by a plugin. Each tool is returned as its bare name (e.g., "azure-search").
// Resolution silently skips missing plugins.
func enumeratePluginTools(entry config.ToolEntry, baseDir string) []string {
	// Try loading the plugin from the registry
	reg := plugin.NewRegistry()
	if entry.Path != "" {
		path := entry.Path
		if !filepath.IsAbs(path) {
			path = filepath.Join(baseDir, path)
		}
		if err := reg.LoadDir(path); err != nil {
			return nil
		}
	}
	
	p, err := reg.Get(entry.Name)
	if err != nil || p == nil {
		return nil
	}
	
	var tools []string
	for _, toolEntry := range p.ToToolEntries() {
		tools = append(tools, toolEntry.Name)
	}
	return tools
}

// removeTool removes a named tool from Generator.Tools in the given config.
// For skill_dir deep variants ("{entry}/{skill}"), it adds the skill to the
// entry's ExcludedSkills list instead of removing the entire entry.
func removeTool(cfg *config.ToolConfig, name string) {
	if cfg.Generator == nil {
		return
	}

	if strings.Contains(name, "/") {
		entryName, subName, ok := strings.Cut(name, "/")
		if !ok {
			return
		}
		// Check if this is an MCP deep variant
		for i, te := range cfg.Generator.Tools {
			if te.ResolvedType() == "mcp" && te.Name == entryName {
				te.MCPTools = removeMCPTool(te.MCPTools, subName)
				cfg.Generator.Tools[i] = te
				return
			}
		}
		// Check if this is a skill_dir deep variant
		for i, te := range cfg.Generator.Tools {
			if te.ResolvedType() == "skill" && te.SkillDir && te.Name == entryName {
				// Add to exclusion list
				te.ExcludedSkills = append(te.ExcludedSkills, subName)
				cfg.Generator.Tools[i] = te
				return
			}
		}
		// Check if this is a plugin deep variant
		for i, te := range cfg.Generator.Tools {
			if te.ResolvedType() == "tool" && te.Name == entryName {
				// Add to exclusion list
				te.ExcludedTools = append(te.ExcludedTools, subName)
				cfg.Generator.Tools[i] = te
				return
			}
		}
	}

	// Shallow removal: remove the entire entry
	var tools []config.ToolEntry
	for _, te := range cfg.Generator.Tools {
		if te.Name != name {
			tools = append(tools, te)
		}
	}
	cfg.Generator.Tools = tools
}

// cloneToolConfig returns a deep copy of a ToolConfig so mutations to the
// clone never affect the original.
func cloneToolConfig(src config.ToolConfig) config.ToolConfig {
	dst := src // copies value fields (Name, Description, Plugins header)

	if src.Generator != nil {
		gen := *src.Generator

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
				if te.Args != nil {
					args := make([]string, len(te.Args))
					copy(args, te.Args)
					gen.Tools[i].Args = args
				}
				if te.MCPTools != nil {
					tools := make([]string, len(te.MCPTools))
					copy(tools, te.MCPTools)
					gen.Tools[i].MCPTools = tools
				}
				if te.ExcludedSkills != nil {
					excluded := make([]string, len(te.ExcludedSkills))
					copy(excluded, te.ExcludedSkills)
					gen.Tools[i].ExcludedSkills = excluded
				}
				if te.ExcludedTools != nil {
					excluded := make([]string, len(te.ExcludedTools))
					copy(excluded, te.ExcludedTools)
					gen.Tools[i].ExcludedTools = excluded
				}
			}
		}

		if len(src.Generator.ExcludedTools) > 0 {
			gen.ExcludedTools = make([]string, len(src.Generator.ExcludedTools))
			copy(gen.ExcludedTools, src.Generator.ExcludedTools)
		}
		if len(src.Generator.Models) > 0 {
			gen.Models = make([]string, len(src.Generator.Models))
			copy(gen.Models, src.Generator.Models)
		}

		dst.Generator = &gen
	}

	if src.Reviewer != nil {
		rev := *src.Reviewer
		if len(src.Reviewer.Tools) > 0 {
			rev.Tools = make([]config.ToolEntry, len(src.Reviewer.Tools))
			copy(rev.Tools, src.Reviewer.Tools)
			for i, te := range rev.Tools {
				if te.When != nil {
					m := make(map[string]string, len(te.When))
					for k, v := range te.When {
						m[k] = v
					}
					rev.Tools[i].When = m
				}
				if te.Args != nil {
					args := make([]string, len(te.Args))
					copy(args, te.Args)
					rev.Tools[i].Args = args
				}
				if te.MCPTools != nil {
					tools := make([]string, len(te.MCPTools))
					copy(tools, te.MCPTools)
					rev.Tools[i].MCPTools = tools
				}
				if te.ExcludedSkills != nil {
					excluded := make([]string, len(te.ExcludedSkills))
					copy(excluded, te.ExcludedSkills)
					rev.Tools[i].ExcludedSkills = excluded
				}
				if te.ExcludedTools != nil {
					excluded := make([]string, len(te.ExcludedTools))
					copy(excluded, te.ExcludedTools)
					rev.Tools[i].ExcludedTools = excluded
				}
			}
		}
		if len(src.Reviewer.Models) > 0 {
			rev.Models = make([]string, len(src.Reviewer.Models))
			copy(rev.Models, src.Reviewer.Models)
		}
		dst.Reviewer = &rev
	}

	return dst
}

func containsWildcard(values []string) bool {
	for _, v := range values {
		if v == "*" {
			return true
		}
	}
	return false
}

func removeMCPTool(tools []string, remove string) []string {
	if len(tools) == 0 {
		return tools
	}
	if containsWildcard(tools) {
		// Wildcard lists require enumerating all tools to remove one entry.
		// Until the SDK exposes tool discovery, leave the wildcard in place.
		return tools
	}
	kept := tools[:0]
	for _, tool := range tools {
		if tool != remove {
			kept = append(kept, tool)
		}
	}
	if len(kept) == 0 {
		return []string{}
	}
	return append([]string(nil), kept...)
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

// PairwiseCheckDiff represents a single check comparison between baseline and variant.
type PairwiseCheckDiff struct {
	GraderName     string `json:"grader_name"`  // grader the check belongs to
	GraderType     string `json:"grader_type"`  // kind of grader
	CheckID        string `json:"check_id"`     // stable id (e.g., check_1)
	CheckLabel     string `json:"check_label"`  // human label
	BaselinePassed bool   `json:"baseline_passed"`
	VariantPassed  bool   `json:"variant_passed"`
	Movement       string `json:"movement"`     // "improved" | "regressed" | "unchanged"
	Reasoning      string `json:"reasoning,omitempty"` // optional — variant reviewer's reasoning if it failed
}

// PairwiseReport holds the complete pairwise comparison results for a prompt.
type PairwiseReport struct {
	PromptID   string                       `json:"prompt_id"`
	Baseline   VariantResult                `json:"baseline"`
	Variants   []VariantResult              `json:"variants"`
	Impacts    []ToolImpact                 `json:"impacts"`
	CheckDiffs map[string][]PairwiseCheckDiff `json:"check_diffs,omitempty"` // keyed by variant config name
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

// EvalReportData is the minimal data needed from an EvalReport to compute check diffs.
// This avoids an import cycle (pairwise → report → pairwise).
type EvalReportData struct {
	ConfigName string
	Graders    []GraderData
}

// GraderData mirrors the essential fields from report.GraderResult.
type GraderData struct {
	Name   string
	Type   string
	Checks []PointData
}

// PointData mirrors the essential fields from report.GraderCheck.
type PointData struct {
	Label   string
	Pass    bool
	Message string
}

// ComputeCheckDiffs compares grader points between baseline and variants.
// Returns a map keyed by variant config name, each containing per-check diffs.
func ComputeCheckDiffs(baseline *EvalReportData, variants []*EvalReportData) map[string][]PairwiseCheckDiff {
	if baseline == nil || len(variants) == 0 {
		return nil
	}

	// Build baseline check index: graderName → checkIndex → point
	baselineChecks := indexPoints(baseline)
	
	result := make(map[string][]PairwiseCheckDiff)
	
	for _, variant := range variants {
		variantChecks := indexPoints(variant)
		var diffs []PairwiseCheckDiff
		
		// Sort grader names for deterministic iteration order
		var graderNames []string
		for name := range baselineChecks {
			graderNames = append(graderNames, name)
		}
		sort.Strings(graderNames)
		
		// Compare each grader in sorted order
		for _, graderName := range graderNames {
			baselinePoints := baselineChecks[graderName]
			variantPoints, hasVariant := variantChecks[graderName]
			
			for checkIdx, basePoint := range baselinePoints {
				varPoint, hasCheck := variantPoints[checkIdx]
				
				var movement string
				varPassed := false
				varReasoning := ""
				
				if hasCheck {
					varPassed = varPoint.Pass
					varReasoning = varPoint.Message
					
					// Determine movement
					if !basePoint.Pass && varPassed {
						movement = "improved"
					} else if basePoint.Pass && !varPassed {
						movement = "regressed"
					} else {
						movement = "unchanged"
					}
				} else {
					// Missing check in variant — treat as unchanged
					movement = "unchanged"
				}
				
				diffs = append(diffs, PairwiseCheckDiff{
					GraderName:     graderName,
					GraderType:     basePoint.Type,
					CheckID:        fmt.Sprintf("check_%d", checkIdx),
					CheckLabel:     basePoint.Label,
					BaselinePassed: basePoint.Pass,
					VariantPassed:  varPassed,
					Movement:       movement,
					Reasoning:      varReasoning,
				})
			}
			
			// Handle checks that exist in variant but not in baseline (rare)
			if hasVariant {
				for checkIdx, varPoint := range variantPoints {
					if _, exists := baselinePoints[checkIdx]; !exists {
						diffs = append(diffs, PairwiseCheckDiff{
							GraderName:     graderName,
							GraderType:     varPoint.Type,
							CheckID:        fmt.Sprintf("check_%d", checkIdx),
							CheckLabel:     varPoint.Label,
							BaselinePassed: false,
							VariantPassed:  varPoint.Pass,
							Movement:       "unchanged", // New check, treat as unchanged
							Reasoning:      varPoint.Message,
						})
					}
				}
			}
		}
		
		result[variant.ConfigName] = diffs
	}
	
	return result
}

// indexedPointData holds point data for comparison.
type indexedPointData struct {
	Label   string
	Pass    bool
	Message string
	Type    string
}

// indexPoints builds a map of graderName → checkIndex → point data.
func indexPoints(data *EvalReportData) map[string]map[int]indexedPointData {
	if data == nil {
		return nil
	}
	
	result := make(map[string]map[int]indexedPointData)
	
	for _, grader := range data.Graders {
		if _, exists := result[grader.Name]; !exists {
			result[grader.Name] = make(map[int]indexedPointData)
		}
		
		for i, point := range grader.Checks {
			result[grader.Name][i] = indexedPointData{
				Label:   point.Label,
				Pass:    point.Pass,
				Message: point.Message,
				Type:    grader.Type,
			}
		}
	}
	
	return result
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
