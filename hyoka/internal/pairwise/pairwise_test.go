package pairwise

import (
	"math"
	"sort"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
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

func TestExpandPairwise_ShallowMCP(t *testing.T) {
	cfg := config.ToolConfig{
		Name: "baseline",
		Generator: &config.GeneratorConfig{
			Model: "gpt-4",
			Tools: []config.ToolEntry{
				{Name: "azure", Type: "mcp", Command: "npx"},
			},
		},
	}

	variants := ExpandPairwise(cfg, ".")
	if len(variants) != 2 {
		t.Fatalf("expected 2 variants, got %d", len(variants))
	}
	if variants[0].Name != "baseline/baseline" {
		t.Errorf("expected baseline variant name, got %q", variants[0].Name)
	}
	if variants[1].Name != "baseline/without-azure" {
		t.Errorf("expected without-azure variant, got %q", variants[1].Name)
	}
	if len(variants[1].Generator.Tools) != 0 {
		t.Errorf("expected MCP entry removed, got %d tools", len(variants[1].Generator.Tools))
	}
}

func TestExpandPairwise_DeepMCPTools(t *testing.T) {
	cfg := config.ToolConfig{
		Name: "baseline",
		Generator: &config.GeneratorConfig{
			Model: "gpt-4",
			Tools: []config.ToolEntry{
				{
					Name:     "azure",
					Type:     "mcp",
					Command:  "npx",
					Pairwise: "deep",
					MCPTools: []string{"storage", "key-vault"},
				},
			},
		},
	}

	variants := ExpandPairwise(cfg, ".")
	if len(variants) != 3 {
		t.Fatalf("expected 3 variants, got %d", len(variants))
	}

	checkVariant := func(name string, expected []string) {
		t.Helper()
		for _, v := range variants {
			if v.Name == name {
				if len(v.Generator.Tools) != 1 {
					t.Fatalf("expected 1 tool entry for %s, got %d", name, len(v.Generator.Tools))
				}
				got := v.Generator.Tools[0].MCPTools
				if len(got) != len(expected) {
					t.Fatalf("expected %d mcp tools for %s, got %d", len(expected), name, len(got))
				}
				for i, tool := range expected {
					if got[i] != tool {
						t.Fatalf("expected tool %q at %d for %s, got %q", tool, i, name, got[i])
					}
				}
				return
			}
		}
		t.Fatalf("variant %q not found", name)
	}

	checkVariant("baseline/without-azure/storage", []string{"key-vault"})
	checkVariant("baseline/without-azure/key-vault", []string{"storage"})
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

func TestAggregateImpactsEmpty(t *testing.T) {
	agg := AggregateImpacts(nil)
	if len(agg) != 0 {
		t.Errorf("expected empty result, got %d", len(agg))
	}
}

func TestExpandPairwise_DeepSkillDir(t *testing.T) {
	// Use the test skills directory (absolute path from repo root)
	cfg := config.ToolConfig{
		Name: "test/baseline",
		Generator: &config.GeneratorConfig{
			Models: []string{"gpt-4"},
			Tools: []config.ToolEntry{
				{
					Name:     "test-skills",
					Type:     "skill",
					Source:   "local",
					Path:     "skills/test",
					SkillDir: true,
					Pairwise: "deep",
				},
			},
		},
	}

	// Use repo root as baseDir
	variants := ExpandPairwise(cfg, "/home/rgeraghty/projects/hyoka")
	if len(variants) != 3 {
		t.Fatalf("expected 3 variants (baseline + 2 skills), got %d", len(variants))
	}

	// Check baseline
	if variants[0].Name != "test/baseline/baseline" {
		t.Errorf("expected baseline variant, got %q", variants[0].Name)
	}
	if len(variants[0].Generator.Tools) != 1 {
		t.Errorf("baseline should have 1 tool entry, got %d", len(variants[0].Generator.Tools))
	}

	// Check that we have the two skill ablation variants
	names := []string{variants[1].Name, variants[2].Name}
	sort.Strings(names)
	expectedNames := []string{
		"test/baseline/without-test-skills/markdown-headings",
		"test/baseline/without-test-skills/markdown-lists",
	}
	for i, name := range expectedNames {
		if names[i] != name {
			t.Errorf("expected variant %q, got %q", name, names[i])
		}
	}

	// Verify that each variant has the skill excluded
	for _, v := range variants[1:] {
		if len(v.Generator.Tools) != 1 {
			t.Fatalf("variant %s should have 1 tool entry, got %d", v.Name, len(v.Generator.Tools))
		}
		entry := v.Generator.Tools[0]
		if entry.Name != "test-skills" {
			t.Errorf("expected test-skills entry, got %q", entry.Name)
		}
		if len(entry.ExcludedSkills) != 1 {
			t.Fatalf("variant %s should exclude 1 skill, got %d", v.Name, len(entry.ExcludedSkills))
		}
	}
}

func TestExpandPairwise_DeepSkillDirEmpty(t *testing.T) {
	cfg := config.ToolConfig{
		Name: "baseline",
		Generator: &config.GeneratorConfig{
			Models: []string{"gpt-4"},
			Tools: []config.ToolEntry{
				{
					Name:     "empty-skills",
					Type:     "skill",
					Source:   "local",
					Path:     "/nonexistent",
					SkillDir: true,
					Pairwise: "deep",
				},
			},
		},
	}

	variants := ExpandPairwise(cfg, ".")
	// Empty skill_dir produces only baseline (no skills to ablate)
	if len(variants) != 1 {
		t.Fatalf("expected 1 variant (baseline only), got %d", len(variants))
	}
	if variants[0].Name != "baseline/baseline" {
		t.Errorf("expected baseline variant, got %q", variants[0].Name)
	}
}

func TestExpandPairwise_ShallowSkillDir(t *testing.T) {
	cfg := config.ToolConfig{
		Name: "baseline",
		Generator: &config.GeneratorConfig{
			Models: []string{"gpt-4"},
			Tools: []config.ToolEntry{
				{
					Name:     "test-skills",
					Type:     "skill",
					Source:   "local",
					Path:     "skills/test",
					SkillDir: true,
					// Pairwise defaults to "shallow"
				},
			},
		},
	}

	variants := ExpandPairwise(cfg, "/home/rgeraghty/projects/hyoka")
	// Shallow mode: baseline + remove entire entry
	if len(variants) != 2 {
		t.Fatalf("expected 2 variants (baseline + without-entry), got %d", len(variants))
	}
	if variants[0].Name != "baseline/baseline" {
		t.Errorf("expected baseline variant, got %q", variants[0].Name)
	}
	if variants[1].Name != "baseline/without-test-skills" {
		t.Errorf("expected without-test-skills variant, got %q", variants[1].Name)
	}
	if len(variants[1].Generator.Tools) != 0 {
		t.Errorf("shallow removal should remove entire entry, got %d tools", len(variants[1].Generator.Tools))
	}
}
