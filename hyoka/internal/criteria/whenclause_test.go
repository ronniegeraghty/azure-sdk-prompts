package criteria

import (
	"testing"
)

func TestWhenClause_Matches(t *testing.T) {
	tests := []struct {
		name  string
		when  WhenClause
		ctx   MatchContext
		want  bool
	}{
		{
			name: "empty when matches everything",
			when: WhenClause{},
			ctx: MatchContext{
				Props: map[string]string{"language": "python"},
			},
			want: true,
		},
		{
			name: "scalar match",
			when: WhenClause{
				Language: StringOrSlice{"python"},
			},
			ctx: MatchContext{
				Props: map[string]string{"language": "python"},
			},
			want: true,
		},
		{
			name: "scalar case-insensitive match",
			when: WhenClause{
				Language: StringOrSlice{"Python"},
			},
			ctx: MatchContext{
				Props: map[string]string{"language": "python"},
			},
			want: true,
		},
		{
			name: "scalar no match",
			when: WhenClause{
				Language: StringOrSlice{"python"},
			},
			ctx: MatchContext{
				Props: map[string]string{"language": "java"},
			},
			want: false,
		},
		{
			name: "list OR match first",
			when: WhenClause{
				Language: StringOrSlice{"python", "java"},
			},
			ctx: MatchContext{
				Props: map[string]string{"language": "python"},
			},
			want: true,
		},
		{
			name: "list OR match second",
			when: WhenClause{
				Language: StringOrSlice{"python", "java"},
			},
			ctx: MatchContext{
				Props: map[string]string{"language": "java"},
			},
			want: true,
		},
		{
			name: "list OR no match",
			when: WhenClause{
				Language: StringOrSlice{"python", "java"},
			},
			ctx: MatchContext{
				Props: map[string]string{"language": "go"},
			},
			want: false,
		},
		{
			name: "multiple fields AND",
			when: WhenClause{
				Language: StringOrSlice{"python"},
				Service:  StringOrSlice{"key-vault"},
			},
			ctx: MatchContext{
				Props: map[string]string{
					"language": "python",
					"service":  "key-vault",
				},
			},
			want: true,
		},
		{
			name: "multiple fields AND one fails",
			when: WhenClause{
				Language: StringOrSlice{"python"},
				Service:  StringOrSlice{"key-vault"},
			},
			ctx: MatchContext{
				Props: map[string]string{
					"language": "python",
					"service":  "storage",
				},
			},
			want: false,
		},
		{
			name: "OR within field, AND across fields",
			when: WhenClause{
				Language: StringOrSlice{"python", "java"},
				Service:  StringOrSlice{"key-vault", "storage"},
			},
			ctx: MatchContext{
				Props: map[string]string{
					"language": "java",
					"service":  "storage",
				},
			},
			want: true,
		},
		{
			name: "tool filter single match",
			when: WhenClause{
				Tool: []ToolFilter{
					{Name: "azure", Source: "mcp"},
				},
			},
			ctx: MatchContext{
				Tools: []ToolIdentity{
					{Name: "azure", Source: "mcp"},
				},
			},
			want: true,
		},
		{
			name: "tool filter no match (name)",
			when: WhenClause{
				Tool: []ToolFilter{
					{Name: "azure", Source: "mcp"},
				},
			},
			ctx: MatchContext{
				Tools: []ToolIdentity{
					{Name: "other", Source: "mcp"},
				},
			},
			want: false,
		},
		{
			name: "tool filter no match (source)",
			when: WhenClause{
				Tool: []ToolFilter{
					{Name: "azure", Source: "mcp"},
				},
			},
			ctx: MatchContext{
				Tools: []ToolIdentity{
					{Name: "azure", Source: "skill"},
				},
			},
			want: false,
		},
		{
			name: "tool filter with MCP server",
			when: WhenClause{
				Tool: []ToolFilter{
					{Name: "azure", Source: "mcp", MCPServer: "azure"},
				},
			},
			ctx: MatchContext{
				Tools: []ToolIdentity{
					{Name: "azure", Source: "mcp", MCPServer: "azure"},
				},
			},
			want: true,
		},
		{
			name: "tool filter MCP server mismatch",
			when: WhenClause{
				Tool: []ToolFilter{
					{Name: "azure", Source: "mcp", MCPServer: "azure"},
				},
			},
			ctx: MatchContext{
				Tools: []ToolIdentity{
					{Name: "azure", Source: "mcp", MCPServer: "other"},
				},
			},
			want: false,
		},
		{
			name: "tool filter AND across entries",
			when: WhenClause{
				Tool: []ToolFilter{
					{Name: "azure", Source: "mcp"},
					{Name: "markdown-headings", Source: "skill"},
				},
			},
			ctx: MatchContext{
				Tools: []ToolIdentity{
					{Name: "azure", Source: "mcp"},
					{Name: "markdown-headings", Source: "skill"},
				},
			},
			want: true,
		},
		{
			name: "tool filter AND across entries, one missing",
			when: WhenClause{
				Tool: []ToolFilter{
					{Name: "azure", Source: "mcp"},
					{Name: "markdown-headings", Source: "skill"},
				},
			},
			ctx: MatchContext{
				Tools: []ToolIdentity{
					{Name: "azure", Source: "mcp"},
				},
			},
			want: false,
		},
		{
			name: "tool filter negate=true (tool absent)",
			when: WhenClause{
				Tool: []ToolFilter{
					{Name: "azure", Source: "mcp", Negate: true},
				},
			},
			ctx: MatchContext{
				Tools: []ToolIdentity{
					{Name: "other", Source: "skill"},
				},
			},
			want: true,
		},
		{
			name: "tool filter negate=true (tool present)",
			when: WhenClause{
				Tool: []ToolFilter{
					{Name: "azure", Source: "mcp", Negate: true},
				},
			},
			ctx: MatchContext{
				Tools: []ToolIdentity{
					{Name: "azure", Source: "mcp"},
				},
			},
			want: false,
		},
		{
			name: "scalars + tool filter all match",
			when: WhenClause{
				Language: StringOrSlice{"python"},
				Service:  StringOrSlice{"key-vault"},
				Tool: []ToolFilter{
					{Name: "azure", Source: "mcp"},
				},
			},
			ctx: MatchContext{
				Props: map[string]string{
					"language": "python",
					"service":  "key-vault",
				},
				Tools: []ToolIdentity{
					{Name: "azure", Source: "mcp"},
				},
			},
			want: true,
		},
		{
			name: "scalars + tool filter, scalar fails",
			when: WhenClause{
				Language: StringOrSlice{"python"},
				Tool: []ToolFilter{
					{Name: "azure", Source: "mcp"},
				},
			},
			ctx: MatchContext{
				Props: map[string]string{
					"language": "java",
				},
				Tools: []ToolIdentity{
					{Name: "azure", Source: "mcp"},
				},
			},
			want: false,
		},
		{
			name: "scalars + tool filter, tool fails",
			when: WhenClause{
				Language: StringOrSlice{"python"},
				Tool: []ToolFilter{
					{Name: "azure", Source: "mcp"},
				},
			},
			ctx: MatchContext{
				Props: map[string]string{
					"language": "python",
				},
				Tools: []ToolIdentity{
					{Name: "other", Source: "skill"},
				},
			},
			want: false,
		},
		{
			name: "all scalar fields",
			when: WhenClause{
				Language:   StringOrSlice{"python"},
				Service:    StringOrSlice{"key-vault"},
				Plane:      StringOrSlice{"data-plane"},
				Category:   StringOrSlice{"crud"},
				SDK:        StringOrSlice{"azure-sdk"},
				Difficulty: StringOrSlice{"easy"},
				Generator:  StringOrSlice{"claude-opus-4.6"},
				Config:     StringOrSlice{"azure-mcp/claude-opus-4.6"},
			},
			ctx: MatchContext{
				Props: map[string]string{
					"language":   "python",
					"service":    "key-vault",
					"plane":      "data-plane",
					"category":   "crud",
					"sdk":        "azure-sdk",
					"difficulty": "easy",
					"generator":  "claude-opus-4.6",
					"config":     "azure-mcp/claude-opus-4.6",
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.when.Matches(tt.ctx)
			if got != tt.want {
				t.Errorf("Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWhenClause_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		when WhenClause
		want bool
	}{
		{
			name: "zero value",
			when: WhenClause{},
			want: true,
		},
		{
			name: "with language",
			when: WhenClause{Language: StringOrSlice{"python"}},
			want: false,
		},
		{
			name: "with tool filter",
			when: WhenClause{Tool: []ToolFilter{{Name: "azure", Source: "mcp"}}},
			want: false,
		},
		{
			name: "all fields empty slices",
			when: WhenClause{
				Language: StringOrSlice{},
				Service:  StringOrSlice{},
				Tool:     []ToolFilter{},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.when.IsEmpty()
			if got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
