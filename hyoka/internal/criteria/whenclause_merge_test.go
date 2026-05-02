package criteria

import (
	"testing"
)

func TestMergeWhenClause(t *testing.T) {
	tests := []struct {
		name   string
		parent WhenClause
		child  WhenClause
		want   WhenClause
	}{
		{
			name:   "both empty",
			parent: WhenClause{},
			child:  WhenClause{},
			want:   WhenClause{},
		},
		{
			name: "parent only scalars",
			parent: WhenClause{
				Language: StringOrSlice{"python"},
				Service:  StringOrSlice{"key-vault"},
			},
			child: WhenClause{},
			want: WhenClause{
				Language: StringOrSlice{"python"},
				Service:  StringOrSlice{"key-vault"},
			},
		},
		{
			name:   "child only scalars",
			parent: WhenClause{},
			child: WhenClause{
				Language: StringOrSlice{"java"},
				Plane:    StringOrSlice{"data-plane"},
			},
			want: WhenClause{
				Language: StringOrSlice{"java"},
				Plane:    StringOrSlice{"data-plane"},
			},
		},
		{
			name: "child replaces parent scalar",
			parent: WhenClause{
				Language: StringOrSlice{"python"},
				Service:  StringOrSlice{"key-vault"},
			},
			child: WhenClause{
				Language: StringOrSlice{"java"},
			},
			want: WhenClause{
				Language: StringOrSlice{"java"},
				Service:  StringOrSlice{"key-vault"},
			},
		},
		{
			name: "child scalar-or-list replaces parent",
			parent: WhenClause{
				Language: StringOrSlice{"python"},
			},
			child: WhenClause{
				Language: StringOrSlice{"java", "go"},
			},
			want: WhenClause{
				Language: StringOrSlice{"java", "go"},
			},
		},
		{
			name: "child empty list clears parent constraint",
			parent: WhenClause{
				Language: StringOrSlice{"python"},
				Service:  StringOrSlice{"key-vault"},
			},
			child: WhenClause{
				Language: StringOrSlice{},
			},
			want: WhenClause{
				Language: StringOrSlice{},
				Service:  StringOrSlice{"key-vault"},
			},
		},
		{
			name: "parent with tool filter",
			parent: WhenClause{
				Tool: []ToolFilter{
					{Name: "azure", Source: "mcp"},
				},
			},
			child: WhenClause{},
			want: WhenClause{
				Tool: []ToolFilter{
					{Name: "azure", Source: "mcp"},
				},
			},
		},
		{
			name: "child replaces parent tool filter",
			parent: WhenClause{
				Tool: []ToolFilter{
					{Name: "azure", Source: "mcp"},
				},
			},
			child: WhenClause{
				Tool: []ToolFilter{
					{Name: "markdown-headings", Source: "skill"},
				},
			},
			want: WhenClause{
				Tool: []ToolFilter{
					{Name: "markdown-headings", Source: "skill"},
				},
			},
		},
		{
			name: "child empty tool list clears parent",
			parent: WhenClause{
				Tool: []ToolFilter{
					{Name: "azure", Source: "mcp"},
				},
			},
			child: WhenClause{
				Tool: []ToolFilter{},
			},
			want: WhenClause{
				Tool: []ToolFilter{},
			},
		},
		{
			name: "file → group → grader cascade",
			parent: WhenClause{
				Language: StringOrSlice{"python", "java"},
				Service:  StringOrSlice{"key-vault"},
				Tool: []ToolFilter{
					{Name: "azure", Source: "mcp"},
				},
			},
			child: WhenClause{
				Language: StringOrSlice{"python"}, // narrows to python only
				Plane:    StringOrSlice{"data-plane"},
			},
			want: WhenClause{
				Language: StringOrSlice{"python"},
				Service:  StringOrSlice{"key-vault"},
				Plane:    StringOrSlice{"data-plane"},
				Tool: []ToolFilter{
					{Name: "azure", Source: "mcp"},
				},
			},
		},
		{
			name: "all fields replaced",
			parent: WhenClause{
				Language:   StringOrSlice{"python"},
				Service:    StringOrSlice{"key-vault"},
				Plane:      StringOrSlice{"data-plane"},
				Category:   StringOrSlice{"crud"},
				SDK:        StringOrSlice{"azure-sdk"},
				Difficulty: StringOrSlice{"easy"},
				Generator:  StringOrSlice{"claude-opus-4.6"},
				Config:     StringOrSlice{"baseline/claude-opus-4.6"},
				Tool:       []ToolFilter{{Name: "azure", Source: "mcp"}},
			},
			child: WhenClause{
				Language:   StringOrSlice{"java"},
				Service:    StringOrSlice{"storage"},
				Plane:      StringOrSlice{"management-plane"},
				Category:   StringOrSlice{"auth"},
				SDK:        StringOrSlice{"azure-mgmt"},
				Difficulty: StringOrSlice{"hard"},
				Generator:  StringOrSlice{"gpt-5.4"},
				Config:     StringOrSlice{"azure-mcp/gpt-5.4"},
				Tool:       []ToolFilter{{Name: "markdown-lists", Source: "skill"}},
			},
			want: WhenClause{
				Language:   StringOrSlice{"java"},
				Service:    StringOrSlice{"storage"},
				Plane:      StringOrSlice{"management-plane"},
				Category:   StringOrSlice{"auth"},
				SDK:        StringOrSlice{"azure-mgmt"},
				Difficulty: StringOrSlice{"hard"},
				Generator:  StringOrSlice{"gpt-5.4"},
				Config:     StringOrSlice{"azure-mcp/gpt-5.4"},
				Tool:       []ToolFilter{{Name: "markdown-lists", Source: "skill"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeWhenClause(tt.parent, tt.child)
			if !whenClauseEqual(got, tt.want) {
				t.Errorf("mergeWhenClause() mismatch:\ngot:  %+v\nwant: %+v", got, tt.want)
			}
		})
	}
}

func whenClauseEqual(a, b WhenClause) bool {
	return equalSlices(a.Language, b.Language) &&
		equalSlices(a.Service, b.Service) &&
		equalSlices(a.Plane, b.Plane) &&
		equalSlices(a.Category, b.Category) &&
		equalSlices(a.SDK, b.SDK) &&
		equalSlices(a.Difficulty, b.Difficulty) &&
		equalSlices(a.Generator, b.Generator) &&
		equalSlices(a.Config, b.Config) &&
		toolFiltersEqual(a.Tool, b.Tool)
}

func toolFiltersEqual(a, b []ToolFilter) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Name != b[i].Name ||
			a[i].Source != b[i].Source ||
			a[i].MCPServer != b[i].MCPServer ||
			a[i].Negate != b[i].Negate {
			return false
		}
	}
	return true
}
