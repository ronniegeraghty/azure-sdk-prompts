package pairwise

import (
	"math"
	"testing"

	"github.com/ronniegeraghty/hyoka/internal/config"
)

func TestComputeImpacts(t *testing.T) {
	results := []VariantResult{
		{ConfigName: "baseline", RemovedTool: "", Score: 8, MaxScore: 10, Success: true},
		{ConfigName: "no-tool-a", RemovedTool: "tool-a", Score: 5, MaxScore: 10, Success: true},
		{ConfigName: "no-tool-b", RemovedTool: "tool-b", Score: 9, MaxScore: 10, Success: true},
		{ConfigName: "no-tool-c", RemovedTool: "tool-c", Score: 8, MaxScore: 10, Success: true},
	}

	report, err := ComputeImpacts("test-prompt", results)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if report.PromptID != "test-prompt" {
		t.Errorf("expected prompt ID test-prompt, got %s", report.PromptID)
	}
	if report.Baseline.ConfigName != "baseline" {
		t.Errorf("expected baseline config name, got %s", report.Baseline.ConfigName)
	}
	if len(report.Variants) != 3 {
		t.Fatalf("expected 3 variants, got %d", len(report.Variants))
	}
	if len(report.Impacts) != 3 {
		t.Fatalf("expected 3 impacts, got %d", len(report.Impacts))
	}

	// Sorted by impact descending: tool-a (30.0), tool-c (0.0), tool-b (-10.0)
	// baseline = 80.0, tool-a = 50.0 → impact = 30.0
	// baseline = 80.0, tool-b = 90.0 → impact = -10.0
	// baseline = 80.0, tool-c = 80.0 → impact = 0.0
	if report.Impacts[0].ToolName != "tool-a" {
		t.Errorf("expected tool-a first, got %s", report.Impacts[0].ToolName)
	}
	if report.Impacts[0].Impact != 30.0 {
		t.Errorf("expected impact 30.0, got %.1f", report.Impacts[0].Impact)
	}
	if report.Impacts[1].ToolName != "tool-c" {
		t.Errorf("expected tool-c second, got %s", report.Impacts[1].ToolName)
	}
	if report.Impacts[1].Impact != 0.0 {
		t.Errorf("expected impact 0.0, got %.1f", report.Impacts[1].Impact)
	}
	if report.Impacts[2].ToolName != "tool-b" {
		t.Errorf("expected tool-b third, got %s", report.Impacts[2].ToolName)
	}
	if report.Impacts[2].Impact != -10.0 {
		t.Errorf("expected impact -10.0, got %.1f", report.Impacts[2].Impact)
	}
}

func TestComputeImpactsNoBaseline(t *testing.T) {
	results := []VariantResult{
		{ConfigName: "no-tool-a", RemovedTool: "tool-a", Score: 5, MaxScore: 10},
	}

	_, err := ComputeImpacts("test-prompt", results)
	if err == nil {
		t.Fatal("expected error for missing baseline")
	}
}

func TestComputeImpactsSingleTool(t *testing.T) {
	results := []VariantResult{
		{ConfigName: "baseline", RemovedTool: "", Score: 7, MaxScore: 10, Success: true},
		{ConfigName: "no-tool-x", RemovedTool: "tool-x", Score: 3, MaxScore: 10, Success: false},
	}

	report, err := ComputeImpacts("single", results)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(report.Impacts) != 1 {
		t.Fatalf("expected 1 impact, got %d", len(report.Impacts))
	}
	if report.Impacts[0].ToolName != "tool-x" {
		t.Errorf("expected tool-x, got %s", report.Impacts[0].ToolName)
	}
	if report.Impacts[0].Impact != 40.0 {
		t.Errorf("expected impact 40.0, got %.1f", report.Impacts[0].Impact)
	}
	if report.Impacts[0].BaselinePass != true {
		t.Error("expected baseline to pass")
	}
	if report.Impacts[0].WithoutPass != false {
		t.Error("expected without to fail")
	}
}

func TestComputeImpactsAllSameScore(t *testing.T) {
	results := []VariantResult{
		{ConfigName: "baseline", RemovedTool: "", Score: 10, MaxScore: 10, Success: true},
		{ConfigName: "no-a", RemovedTool: "a", Score: 10, MaxScore: 10, Success: true},
		{ConfigName: "no-b", RemovedTool: "b", Score: 10, MaxScore: 10, Success: true},
		{ConfigName: "no-c", RemovedTool: "c", Score: 10, MaxScore: 10, Success: true},
	}

	report, err := ComputeImpacts("all-same", results)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, imp := range report.Impacts {
		if imp.Impact != 0.0 {
			t.Errorf("expected 0 impact for %s, got %.1f", imp.ToolName, imp.Impact)
		}
	}
	// With all same impact, should be sorted alphabetically
	if report.Impacts[0].ToolName != "a" {
		t.Errorf("expected a first (alphabetical tiebreak), got %s", report.Impacts[0].ToolName)
	}
}

func TestComputeImpactsZeroMaxScore(t *testing.T) {
	results := []VariantResult{
		{ConfigName: "baseline", RemovedTool: "", Score: 0, MaxScore: 0, Success: false},
		{ConfigName: "no-tool", RemovedTool: "tool", Score: 0, MaxScore: 0, Success: false},
	}

	report, err := ComputeImpacts("zero-max", results)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if report.Impacts[0].Impact != 0.0 {
		t.Errorf("expected 0 impact for zero max score, got %.1f", report.Impacts[0].Impact)
	}
}

func TestSortByImpact(t *testing.T) {
	impacts := []ToolImpact{
		{ToolName: "z-tool", Impact: 5.0},
		{ToolName: "a-tool", Impact: 5.0},
		{ToolName: "m-tool", Impact: 10.0},
		{ToolName: "low", Impact: -3.0},
	}

	SortByImpact(impacts)

	expected := []string{"m-tool", "a-tool", "z-tool", "low"}
	for i, exp := range expected {
		if impacts[i].ToolName != exp {
			t.Errorf("position %d: expected %s, got %s", i, exp, impacts[i].ToolName)
		}
	}
}

func TestNormalizeScore(t *testing.T) {
	tests := []struct {
		score, max int
		want       float64
	}{
		{8, 10, 80.0},
		{10, 10, 100.0},
		{0, 10, 0.0},
		{3, 7, 42.9},
		{0, 0, 0.0},
		{5, 0, 0.0},
	}
	for _, tt := range tests {
		got := normalizeScore(tt.score, tt.max)
		if math.Abs(got-tt.want) > 0.1 {
			t.Errorf("normalizeScore(%d, %d) = %.1f, want %.1f", tt.score, tt.max, got, tt.want)
		}
	}
}

func TestAggregateImpacts(t *testing.T) {
	reports := []*PairwiseReport{
		{
			PromptID: "p1",
			Impacts: []ToolImpact{
				{ToolName: "tool-a", Impact: 20.0, BaselineScore: 80.0, WithoutScore: 60.0, BaselinePass: true, WithoutPass: true},
				{ToolName: "tool-b", Impact: -10.0, BaselineScore: 80.0, WithoutScore: 90.0, BaselinePass: true, WithoutPass: true},
			},
		},
		{
			PromptID: "p2",
			Impacts: []ToolImpact{
				{ToolName: "tool-a", Impact: 10.0, BaselineScore: 70.0, WithoutScore: 60.0, BaselinePass: true, WithoutPass: false},
				{ToolName: "tool-b", Impact: 5.0, BaselineScore: 70.0, WithoutScore: 65.0, BaselinePass: true, WithoutPass: true},
			},
		},
	}

	agg := AggregateImpacts(reports)
	if len(agg) != 2 {
		t.Fatalf("expected 2 aggregated tools, got %d", len(agg))
	}

	// tool-a: avg impact = (20+10)/2 = 15.0, tool-b: avg impact = (-10+5)/2 = -2.5
	if agg[0].ToolName != "tool-a" {
		t.Errorf("expected tool-a first, got %s", agg[0].ToolName)
	}
	if agg[0].Impact != 15.0 {
		t.Errorf("expected tool-a impact 15.0, got %.1f", agg[0].Impact)
	}
	if agg[1].ToolName != "tool-b" {
		t.Errorf("expected tool-b second, got %s", agg[1].ToolName)
	}
	if agg[1].Impact != -2.5 {
		t.Errorf("expected tool-b impact -2.5, got %.1f", agg[1].Impact)
	}

	// tool-a: WithoutPass should be false (not all passed)
	if agg[0].WithoutPass != false {
		t.Error("expected tool-a WithoutPass false (not all prompts passed without it)")
	}
	// tool-b: WithoutPass should be true (all passed)
	if agg[1].WithoutPass != true {
		t.Error("expected tool-b WithoutPass true")
	}
}

// --- ExpandPairwise / collectTogglable / removeTool tests ---

func TestExpandPairwiseBasicTools(t *testing.T) {
	cfg := config.ToolConfig{
		Name: "test",
		Generator: &config.GeneratorConfig{
			Tools: []config.ToolEntry{
				{Name: "tool-a"},
				{Name: "tool-b", AlwaysOn: true},
			},
			AvailableTools: []string{"tool-c"},
		},
	}
	variants := ExpandPairwise(cfg)
	// baseline + tool-a + tool-c (tool-b is AlwaysOn)
	if len(variants) != 3 {
		t.Fatalf("expected 3 variants, got %d", len(variants))
	}
	if variants[0].Name != "test/baseline" {
		t.Errorf("expected baseline, got %s", variants[0].Name)
	}
	if variants[1].Name != "test/without-tool-a" {
		t.Errorf("expected without-tool-a, got %s", variants[1].Name)
	}
	if variants[2].Name != "test/without-tool-c" {
		t.Errorf("expected without-tool-c, got %s", variants[2].Name)
	}
}

func TestExpandPairwiseMCPServers(t *testing.T) {
	cfg := config.ToolConfig{
		Name: "azure-mcp/claude-opus-4.6",
		Generator: &config.GeneratorConfig{
			MCPServers: map[string]*config.MCPServer{
				"azure": {Type: "stdio", Command: "npx"},
				"redis": {Type: "stdio", Command: "redis-mcp"},
			},
		},
	}
	variants := ExpandPairwise(cfg)
	// baseline + azure + redis = 3
	if len(variants) != 3 {
		t.Fatalf("expected 3 variants, got %d", len(variants))
	}
	if variants[0].Name != "azure-mcp/claude-opus-4.6/baseline" {
		t.Errorf("expected baseline, got %s", variants[0].Name)
	}
	// Sorted alphabetically: azure before redis
	if variants[1].Name != "azure-mcp/claude-opus-4.6/without-mcp:azure" {
		t.Errorf("expected without-mcp:azure, got %s", variants[1].Name)
	}
	if variants[2].Name != "azure-mcp/claude-opus-4.6/without-mcp:redis" {
		t.Errorf("expected without-mcp:redis, got %s", variants[2].Name)
	}
	// Verify azure MCP server is removed from variant 1
	if _, ok := variants[1].Generator.MCPServers["azure"]; ok {
		t.Error("expected azure MCP server to be removed in variant 1")
	}
	if _, ok := variants[1].Generator.MCPServers["redis"]; !ok {
		t.Error("expected redis MCP server to be kept in variant 1")
	}
	// Verify redis MCP server is removed from variant 2
	if _, ok := variants[2].Generator.MCPServers["redis"]; ok {
		t.Error("expected redis MCP server to be removed in variant 2")
	}
	if _, ok := variants[2].Generator.MCPServers["azure"]; !ok {
		t.Error("expected azure MCP server to be kept in variant 2")
	}
}

func TestExpandPairwiseSkills(t *testing.T) {
	cfg := config.ToolConfig{
		Name: "baseline-skills/opus",
		Generator: &config.GeneratorConfig{
			Skills: []config.Skill{
				{Type: "local", Path: "skills/generator/azure-sdk.md"},
				{Type: "remote", Name: "copilot-best-practices", Repo: "github/copilot-skills"},
			},
		},
	}
	variants := ExpandPairwise(cfg)
	// baseline + 2 skills = 3
	if len(variants) != 3 {
		t.Fatalf("expected 3 variants, got %d", len(variants))
	}
	if variants[1].Name != "baseline-skills/opus/without-skill:skills/generator/azure-sdk.md" {
		t.Errorf("unexpected variant name: %s", variants[1].Name)
	}
	if variants[2].Name != "baseline-skills/opus/without-skill:copilot-best-practices" {
		t.Errorf("unexpected variant name: %s", variants[2].Name)
	}
	// Verify local skill removed in variant 1
	if len(variants[1].Generator.Skills) != 1 {
		t.Fatalf("expected 1 skill remaining, got %d", len(variants[1].Generator.Skills))
	}
	if variants[1].Generator.Skills[0].Name != "copilot-best-practices" {
		t.Errorf("expected copilot-best-practices kept, got %s", variants[1].Generator.Skills[0].Name)
	}
	// Verify remote skill removed in variant 2
	if len(variants[2].Generator.Skills) != 1 {
		t.Fatalf("expected 1 skill remaining, got %d", len(variants[2].Generator.Skills))
	}
	if variants[2].Generator.Skills[0].Path != "skills/generator/azure-sdk.md" {
		t.Errorf("expected azure-sdk.md kept, got %s", variants[2].Generator.Skills[0].Path)
	}
}

func TestExpandPairwiseMixed(t *testing.T) {
	cfg := config.ToolConfig{
		Name: "full",
		Generator: &config.GeneratorConfig{
			Tools:          []config.ToolEntry{{Name: "tool-x"}},
			AvailableTools: []string{"tool-y"},
			MCPServers: map[string]*config.MCPServer{
				"azure": {Type: "stdio", Command: "npx"},
			},
			Skills: []config.Skill{
				{Type: "local", Path: "skills/gen.md"},
			},
		},
	}
	variants := ExpandPairwise(cfg)
	// baseline + tool-x + tool-y + mcp:azure + skill:skills/gen.md = 5
	if len(variants) != 5 {
		t.Fatalf("expected 5 variants, got %d", len(variants))
	}
	names := make([]string, len(variants))
	for i, v := range variants {
		names[i] = v.Name
	}
	expected := []string{
		"full/baseline",
		"full/without-tool-x",
		"full/without-tool-y",
		"full/without-mcp:azure",
		"full/without-skill:skills/gen.md",
	}
	for i, exp := range expected {
		if names[i] != exp {
			t.Errorf("variant %d: expected %s, got %s", i, exp, names[i])
		}
	}
}

func TestExpandPairwiseNilGenerator(t *testing.T) {
	cfg := config.ToolConfig{Name: "empty"}
	variants := ExpandPairwise(cfg)
	// Only baseline when nothing to toggle
	if len(variants) != 1 {
		t.Fatalf("expected 1 variant (baseline only), got %d", len(variants))
	}
}

func TestRemoveToolMCP(t *testing.T) {
	cfg := config.ToolConfig{
		Generator: &config.GeneratorConfig{
			MCPServers: map[string]*config.MCPServer{
				"azure": {Type: "stdio"},
				"redis": {Type: "stdio"},
			},
		},
	}
	removeTool(&cfg, "mcp:azure")
	if _, ok := cfg.Generator.MCPServers["azure"]; ok {
		t.Error("expected azure removed")
	}
	if _, ok := cfg.Generator.MCPServers["redis"]; !ok {
		t.Error("expected redis kept")
	}
}

func TestRemoveToolSkill(t *testing.T) {
	cfg := config.ToolConfig{
		Generator: &config.GeneratorConfig{
			Skills: []config.Skill{
				{Type: "local", Path: "a.md"},
				{Type: "local", Path: "b.md"},
			},
		},
	}
	removeTool(&cfg, "skill:a.md")
	if len(cfg.Generator.Skills) != 1 {
		t.Fatalf("expected 1 skill, got %d", len(cfg.Generator.Skills))
	}
	if cfg.Generator.Skills[0].Path != "b.md" {
		t.Errorf("expected b.md kept, got %s", cfg.Generator.Skills[0].Path)
	}
}

func TestCloneDoesNotMutateOriginalMCP(t *testing.T) {
	orig := config.ToolConfig{
		Name: "orig",
		Generator: &config.GeneratorConfig{
			MCPServers: map[string]*config.MCPServer{
				"azure": {Type: "stdio", Command: "npx", Args: []string{"@azure/mcp"}},
			},
			Skills: []config.Skill{{Type: "local", Path: "s.md"}},
		},
	}
	cloned := cloneToolConfig(orig)
	removeTool(&cloned, "mcp:azure")
	removeTool(&cloned, "skill:s.md")
	if _, ok := orig.Generator.MCPServers["azure"]; !ok {
		t.Error("original MCP servers should not be affected by clone mutation")
	}
	if len(orig.Generator.Skills) != 1 {
		t.Error("original skills should not be affected by clone mutation")
	}
}

func TestAggregateImpactsEmpty(t *testing.T) {
	agg := AggregateImpacts(nil)
	if len(agg) != 0 {
		t.Errorf("expected empty result, got %d", len(agg))
	}
}
