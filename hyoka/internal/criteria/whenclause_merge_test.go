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
				Language: MatchSet{Is: StringOrSlice{"python"}},
				Service:  MatchSet{Is: StringOrSlice{"key-vault"}},
			},
			child: WhenClause{},
			want: WhenClause{
				Language: MatchSet{Is: StringOrSlice{"python"}},
				Service:  MatchSet{Is: StringOrSlice{"key-vault"}},
			},
		},
		{
			name:   "child only scalars",
			parent: WhenClause{},
			child: WhenClause{
				Language: MatchSet{Is: StringOrSlice{"java"}},
				Plane:    MatchSet{Is: StringOrSlice{"data-plane"}},
			},
			want: WhenClause{
				Language: MatchSet{Is: StringOrSlice{"java"}},
				Plane:    MatchSet{Is: StringOrSlice{"data-plane"}},
			},
		},
		{
			name: "child replaces parent scalar",
			parent: WhenClause{
				Language: MatchSet{Is: StringOrSlice{"python"}},
				Service:  MatchSet{Is: StringOrSlice{"key-vault"}},
			},
			child: WhenClause{
				Language: MatchSet{Is: StringOrSlice{"java"}},
			},
			want: WhenClause{
				Language: MatchSet{Is: StringOrSlice{"java"}},
				Service:  MatchSet{Is: StringOrSlice{"key-vault"}},
			},
		},
		{
			name: "child scalar-or-list replaces parent",
			parent: WhenClause{
				Language: MatchSet{Is: StringOrSlice{"python"}},
			},
			child: WhenClause{
				Language: MatchSet{Is: StringOrSlice{"java", "go"}},
			},
			want: WhenClause{
				Language: MatchSet{Is: StringOrSlice{"java", "go"}},
			},
		},
		{
			name: "child empty list clears parent constraint",
			parent: WhenClause{
				Language: MatchSet{Is: StringOrSlice{"python"}},
				Service:  MatchSet{Is: StringOrSlice{"key-vault"}},
			},
			child: WhenClause{
				Language: MatchSet{Is: StringOrSlice{}},
			},
			want: WhenClause{
				Language: MatchSet{Is: StringOrSlice{}},
				Service:  MatchSet{Is: StringOrSlice{"key-vault"}},
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
				Language: MatchSet{Is: StringOrSlice{"python", "java"}},
				Service:  MatchSet{Is: StringOrSlice{"key-vault"}},
				Tool: []ToolFilter{
					{Name: "azure", Source: "mcp"},
				},
			},
			child: WhenClause{
				Language: MatchSet{Is: StringOrSlice{"python"}}, // narrows to python only
				Plane:    MatchSet{Is: StringOrSlice{"data-plane"}},
			},
			want: WhenClause{
				Language: MatchSet{Is: StringOrSlice{"python"}},
				Service:  MatchSet{Is: StringOrSlice{"key-vault"}},
				Plane:    MatchSet{Is: StringOrSlice{"data-plane"}},
				Tool: []ToolFilter{
					{Name: "azure", Source: "mcp"},
				},
			},
		},
		{
			name: "all fields replaced",
			parent: WhenClause{
				Language:   MatchSet{Is: StringOrSlice{"python"}},
				Service:    MatchSet{Is: StringOrSlice{"key-vault"}},
				Plane:      MatchSet{Is: StringOrSlice{"data-plane"}},
				Category:   MatchSet{Is: StringOrSlice{"crud"}},
				SDK:        MatchSet{Is: StringOrSlice{"azure-sdk"}},
				Difficulty: MatchSet{Is: StringOrSlice{"easy"}},
				Generator:  MatchSet{Is: StringOrSlice{"claude-opus-4.6"}},
				Config:     MatchSet{Is: StringOrSlice{"baseline/claude-opus-4.6"}},
				Tool:       []ToolFilter{{Name: "azure", Source: "mcp"}},
			},
			child: WhenClause{
				Language:   MatchSet{Is: StringOrSlice{"java"}},
				Service:    MatchSet{Is: StringOrSlice{"storage"}},
				Plane:      MatchSet{Is: StringOrSlice{"management-plane"}},
				Category:   MatchSet{Is: StringOrSlice{"auth"}},
				SDK:        MatchSet{Is: StringOrSlice{"azure-mgmt"}},
				Difficulty: MatchSet{Is: StringOrSlice{"hard"}},
				Generator:  MatchSet{Is: StringOrSlice{"gpt-5.4"}},
				Config:     MatchSet{Is: StringOrSlice{"azure-mcp/gpt-5.4"}},
				Tool:       []ToolFilter{{Name: "markdown-lists", Source: "skill"}},
			},
			want: WhenClause{
				Language:   MatchSet{Is: StringOrSlice{"java"}},
				Service:    MatchSet{Is: StringOrSlice{"storage"}},
				Plane:      MatchSet{Is: StringOrSlice{"management-plane"}},
				Category:   MatchSet{Is: StringOrSlice{"auth"}},
				SDK:        MatchSet{Is: StringOrSlice{"azure-mgmt"}},
				Difficulty: MatchSet{Is: StringOrSlice{"hard"}},
				Generator:  MatchSet{Is: StringOrSlice{"gpt-5.4"}},
				Config:     MatchSet{Is: StringOrSlice{"azure-mcp/gpt-5.4"}},
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
	return matchSetEqual(a.Language, b.Language) &&
		matchSetEqual(a.Service, b.Service) &&
		matchSetEqual(a.Plane, b.Plane) &&
		matchSetEqual(a.Category, b.Category) &&
		matchSetEqual(a.SDK, b.SDK) &&
		matchSetEqual(a.Difficulty, b.Difficulty) &&
		matchSetEqual(a.Generator, b.Generator) &&
		matchSetEqual(a.Config, b.Config) &&
		matchSetEqual(a.Tags, b.Tags) &&
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
