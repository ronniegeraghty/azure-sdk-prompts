package criteria

import (
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria/graders"
)

// TestMatchingUnifiedEntries_GatedByToolFilters verifies a grader gated on
// `tool: [{name: azure, source: mcp}]` runs when the eval config includes the
// azure MCP tool and is skipped otherwise. This is the headline use case for
// the Phase 2 config-aware `when:` feature.
func TestMatchingUnifiedEntries_GatedByToolFilters(t *testing.T) {
	bundle := &Bundle{Configs: []UnifiedGraderConfig{{
		Graders: []UnifiedGraderEntry{
			{
				Type: graders.KindTool,
				Name: "uses-azure-list-resources",
				When: WhenClause{
					Tool: []ToolFilter{
						{Name: "azure", Source: "mcp"},
					},
				},
			},
			{
				Type: graders.KindTool,
				Name: "uses-some-skill",
				When: WhenClause{
					Tool: []ToolFilter{
						{Name: "reviewer-skills", Source: "skill"},
					},
				},
			},
			{
				// Always runs — no when block.
				Type: graders.KindTool,
				Name: "always-on",
			},
		},
	}}}

	cases := []struct {
		name  string
		ctx   MatchContext
		want  []string
	}{
		{
			name: "azure-mcp config loads server → grader runs",
			ctx: MatchContext{
				Props: map[string]string{
					"language":  "python",
					"generator": "claude-opus-4.6",
					"config":    "azure-mcp/claude-opus-4.6",
				},
				Tools: []ToolIdentity{
					{Name: "azure", Source: "mcp"},
				},
			},
			want: []string{"uses-azure-list-resources", "always-on"},
		},
		{
			name: "baseline config (no azure mcp) → mcp-gated grader skipped",
			ctx: MatchContext{
				Props: map[string]string{
					"language":  "python",
					"generator": "claude-opus-4.6",
					"config":    "baseline/claude-opus-4.6",
				},
				Tools: nil, // No tools
			},
			want: []string{"always-on"},
		},
		{
			name: "config with skill but no mcp → only skill-gated runs",
			ctx: MatchContext{
				Props: map[string]string{
					"language":  "python",
					"generator": "claude-opus-4.6",
					"config":    "skills-only/claude-opus-4.6",
				},
				Tools: []ToolIdentity{
					{Name: "reviewer-skills", Source: "skill"},
				},
			},
			want: []string{"uses-some-skill", "always-on"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			matched := MatchingUnifiedEntries(bundle, tc.ctx)
			got := make([]string, 0, len(matched))
			for _, m := range matched {
				got = append(got, m.Entry.Name)
			}
			if len(got) != len(tc.want) {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
			for i := range got {
				if got[i] != tc.want[i] {
					t.Errorf("index %d: got %q, want %q", i, got[i], tc.want[i])
				}
			}
		})
	}
}
