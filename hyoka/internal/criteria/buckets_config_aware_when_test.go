package criteria

import (
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria/graders"
)

// TestMatchingUnifiedEntries_GatedByConfigAwarePrefixedKeys verifies a grader
// gated on `mcp_server:azure: "true"` runs when the eval config exposes the
// azure MCP server and is skipped otherwise. This is the headline use case for
// the Phase 1 config-aware `when:` feature.
func TestMatchingUnifiedEntries_GatedByConfigAwarePrefixedKeys(t *testing.T) {
	bundle := &Bundle{Configs: []UnifiedGraderConfig{{
		Graders: []UnifiedGraderEntry{
			{
				Type: graders.KindTool,
				Name: "uses-azure-list-resources",
				When: map[string]string{"mcp_server:azure": "true"},
			},
			{
				Type: graders.KindTool,
				Name: "uses-some-skill",
				When: map[string]string{"skill:reviewer-skills": "true"},
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
		props map[string]string
		want  []string
	}{
		{
			name: "azure-mcp config loads server → grader runs",
			props: map[string]string{
				"language":         "python",
				"generator":        "claude-opus-4.6",
				"config":           "azure-mcp/claude-opus-4.6",
				"mcp_server:azure": "true",
			},
			want: []string{"uses-azure-list-resources", "always-on"},
		},
		{
			name: "baseline config (no azure mcp) → mcp-gated grader skipped",
			props: map[string]string{
				"language":  "python",
				"generator": "claude-opus-4.6",
				"config":    "baseline/claude-opus-4.6",
			},
			want: []string{"always-on"},
		},
		{
			name: "config with skill but no mcp → only skill-gated runs",
			props: map[string]string{
				"language":              "python",
				"generator":             "claude-opus-4.6",
				"config":                "skills-only/claude-opus-4.6",
				"skill:reviewer-skills": "true",
			},
			want: []string{"uses-some-skill", "always-on"},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			matched := MatchingUnifiedEntries(bundle, tc.props)
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
