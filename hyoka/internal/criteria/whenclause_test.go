package criteria

import (
	"testing"

	"gopkg.in/yaml.v3"
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
				Language: MatchSet{Is: StringOrSlice{"python"}},
			},
			ctx: MatchContext{
				Props: map[string]string{"language": "python"},
			},
			want: true,
		},
		{
			name: "scalar case-insensitive match",
			when: WhenClause{
				Language: MatchSet{Is: StringOrSlice{"Python"}},
			},
			ctx: MatchContext{
				Props: map[string]string{"language": "python"},
			},
			want: true,
		},
		{
			name: "scalar no match",
			when: WhenClause{
				Language: MatchSet{Is: StringOrSlice{"python"}},
			},
			ctx: MatchContext{
				Props: map[string]string{"language": "java"},
			},
			want: false,
		},
		{
			name: "list OR match first",
			when: WhenClause{
				Language: MatchSet{Is: StringOrSlice{"python", "java"}},
			},
			ctx: MatchContext{
				Props: map[string]string{"language": "python"},
			},
			want: true,
		},
		{
			name: "list OR match second",
			when: WhenClause{
				Language: MatchSet{Is: StringOrSlice{"python", "java"}},
			},
			ctx: MatchContext{
				Props: map[string]string{"language": "java"},
			},
			want: true,
		},
		{
			name: "list OR no match",
			when: WhenClause{
				Language: MatchSet{Is: StringOrSlice{"python", "java"}},
			},
			ctx: MatchContext{
				Props: map[string]string{"language": "go"},
			},
			want: false,
		},
		{
			name: "multiple fields AND",
			when: WhenClause{
				Language: MatchSet{Is: StringOrSlice{"python"}},
				Service:  MatchSet{Is: StringOrSlice{"key-vault"}},
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
				Language: MatchSet{Is: StringOrSlice{"python"}},
				Service:  MatchSet{Is: StringOrSlice{"key-vault"}},
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
				Language: MatchSet{Is: StringOrSlice{"python", "java"}},
				Service:  MatchSet{Is: StringOrSlice{"key-vault", "storage"}},
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
				Language: MatchSet{Is: StringOrSlice{"python"}},
				Service:  MatchSet{Is: StringOrSlice{"key-vault"}},
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
				Language: MatchSet{Is: StringOrSlice{"python"}},
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
				Language: MatchSet{Is: StringOrSlice{"python"}},
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
				Language:   MatchSet{Is: StringOrSlice{"python"}},
				Service:    MatchSet{Is: StringOrSlice{"key-vault"}},
				Plane:      MatchSet{Is: StringOrSlice{"data-plane"}},
				Category:   MatchSet{Is: StringOrSlice{"crud"}},
				SDK:        MatchSet{Is: StringOrSlice{"azure-sdk"}},
				Difficulty: MatchSet{Is: StringOrSlice{"easy"}},
				Generator:  MatchSet{Is: StringOrSlice{"claude-opus-4.6"}},
				Config:     MatchSet{Is: StringOrSlice{"azure-mcp/claude-opus-4.6"}},
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
			when: WhenClause{Language: MatchSet{Is: StringOrSlice{"python"}}},
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
				Language: MatchSet{Is: StringOrSlice{}},
				Service:  MatchSet{Is: StringOrSlice{}},
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

func TestWhenClause_TagMatching(t *testing.T) {
	tests := []struct {
		name string
		when WhenClause
		ctx  MatchContext
		want bool
	}{
		{
			name: "empty tags always matches",
			when: WhenClause{},
			ctx: MatchContext{
				Props: map[string]string{"language": "python"},
				Tags:  []string{"auth", "crud"},
			},
			want: true,
		},
		{
			name: "tags is - match one",
			when: WhenClause{
				Tags: MatchSet{Is: StringOrSlice{"auth"}},
			},
			ctx: MatchContext{
				Props: map[string]string{"language": "python"},
				Tags:  []string{"auth", "crud"},
			},
			want: true,
		},
		{
			name: "tags is - no match",
			when: WhenClause{
				Tags: MatchSet{Is: StringOrSlice{"pagination"}},
			},
			ctx: MatchContext{
				Props: map[string]string{"language": "python"},
				Tags:  []string{"auth", "crud"},
			},
			want: false,
		},
		{
			name: "tags is - case insensitive",
			when: WhenClause{
				Tags: MatchSet{Is: StringOrSlice{"AUTH"}},
			},
			ctx: MatchContext{
				Props: map[string]string{"language": "python"},
				Tags:  []string{"auth", "crud"},
			},
			want: true,
		},
		{
			name: "tags not - excludes",
			when: WhenClause{
				Tags: MatchSet{Not: StringOrSlice{"deprecated"}},
			},
			ctx: MatchContext{
				Props: map[string]string{"language": "python"},
				Tags:  []string{"auth", "crud"},
			},
			want: true,
		},
		{
			name: "tags not - fails on match",
			when: WhenClause{
				Tags: MatchSet{Not: StringOrSlice{"deprecated"}},
			},
			ctx: MatchContext{
				Props: map[string]string{"language": "python"},
				Tags:  []string{"auth", "deprecated"},
			},
			want: false,
		},
		{
			name: "tags is and not - passes",
			when: WhenClause{
				Tags: MatchSet{
					Is:  StringOrSlice{"auth", "crud"},
					Not: StringOrSlice{"deprecated"},
				},
			},
			ctx: MatchContext{
				Props: map[string]string{"language": "python"},
				Tags:  []string{"auth", "pagination"},
			},
			want: true,
		},
		{
			name: "tags is and not - fails on not",
			when: WhenClause{
				Tags: MatchSet{
					Is:  StringOrSlice{"auth", "crud"},
					Not: StringOrSlice{"deprecated"},
				},
			},
			ctx: MatchContext{
				Props: map[string]string{"language": "python"},
				Tags:  []string{"auth", "deprecated"},
			},
			want: false,
		},
		{
			name: "mixed scalar and tags",
			when: WhenClause{
				Language: MatchSet{Is: StringOrSlice{"python"}},
				Tags:     MatchSet{Is: StringOrSlice{"auth"}},
			},
			ctx: MatchContext{
				Props: map[string]string{"language": "python"},
				Tags:  []string{"auth", "crud"},
			},
			want: true,
		},
		{
			name: "mixed scalar and tags - language fails",
			when: WhenClause{
				Language: MatchSet{Is: StringOrSlice{"java"}},
				Tags:     MatchSet{Is: StringOrSlice{"auth"}},
			},
			ctx: MatchContext{
				Props: map[string]string{"language": "python"},
				Tags:  []string{"auth", "crud"},
			},
			want: false,
		},
		{
			name: "mixed scalar and tags - tags fail",
			when: WhenClause{
				Language: MatchSet{Is: StringOrSlice{"python"}},
				Tags:     MatchSet{Is: StringOrSlice{"pagination"}},
			},
			ctx: MatchContext{
				Props: map[string]string{"language": "python"},
				Tags:  []string{"auth", "crud"},
			},
			want: false,
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

func TestWhenClause_Negation(t *testing.T) {
	tests := []struct {
		name string
		when WhenClause
		ctx  MatchContext
		want bool
	}{
		{
			name: "service not - passes when different",
			when: WhenClause{
				Service: MatchSet{Not: StringOrSlice{"identity"}},
			},
			ctx: MatchContext{
				Props: map[string]string{"service": "key-vault"},
			},
			want: true,
		},
		{
			name: "service not - fails when matches",
			when: WhenClause{
				Service: MatchSet{Not: StringOrSlice{"identity"}},
			},
			ctx: MatchContext{
				Props: map[string]string{"service": "identity"},
			},
			want: false,
		},
		{
			name: "service is and not together - passes",
			when: WhenClause{
				Service: MatchSet{
					Is:  StringOrSlice{"key-vault", "storage"},
					Not: StringOrSlice{"cosmos-db"},
				},
			},
			ctx: MatchContext{
				Props: map[string]string{"service": "key-vault"},
			},
			want: true,
		},
		{
			name: "service is and not together - not in is",
			when: WhenClause{
				Service: MatchSet{
					Is:  StringOrSlice{"key-vault", "storage"},
					Not: StringOrSlice{"cosmos-db"},
				},
			},
			ctx: MatchContext{
				Props: map[string]string{"service": "identity"},
			},
			want: false,
		},
		{
			name: "service is and not together - in not",
			when: WhenClause{
				Service: MatchSet{
					Is:  StringOrSlice{"key-vault", "storage"},
					Not: StringOrSlice{"cosmos-db"},
				},
			},
			ctx: MatchContext{
				Props: map[string]string{"service": "cosmos-db"},
			},
			want: false,
		},
		{
			name: "category not list - passes",
			when: WhenClause{
				Category: MatchSet{Not: StringOrSlice{"pagination", "streaming"}},
			},
			ctx: MatchContext{
				Props: map[string]string{"category": "auth"},
			},
			want: true,
		},
		{
			name: "category not list - fails on first",
			when: WhenClause{
				Category: MatchSet{Not: StringOrSlice{"pagination", "streaming"}},
			},
			ctx: MatchContext{
				Props: map[string]string{"category": "pagination"},
			},
			want: false,
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

func TestWhenClause_YAMLRoundtrip(t *testing.T) {
yamlStr := `
language: python
service:
  not: [identity, key-vault]
category:
  is: [auth, crud]
  not: pagination
tags:
  is: [authentication]
  not: [deprecated, experimental]
generator: claude-opus-4.6
`
var when WhenClause
err := yaml.Unmarshal([]byte(yamlStr), &when)
if err != nil {
t.Fatalf("unmarshal failed: %v", err)
}

// Check Language (scalar short form)
if len(when.Language.Is) != 1 || when.Language.Is[0] != "python" {
t.Errorf("Language.Is = %v, want [python]", when.Language.Is)
}
if len(when.Language.Not) != 0 {
t.Errorf("Language.Not = %v, want []", when.Language.Not)
}

// Check Service (negation only)
if len(when.Service.Is) != 0 {
t.Errorf("Service.Is = %v, want []", when.Service.Is)
}
if len(when.Service.Not) != 2 || when.Service.Not[0] != "identity" || when.Service.Not[1] != "key-vault" {
t.Errorf("Service.Not = %v, want [identity key-vault]", when.Service.Not)
}

// Check Category (is + not)
if len(when.Category.Is) != 2 || when.Category.Is[0] != "auth" || when.Category.Is[1] != "crud" {
t.Errorf("Category.Is = %v, want [auth crud]", when.Category.Is)
}
if len(when.Category.Not) != 1 || when.Category.Not[0] != "pagination" {
t.Errorf("Category.Not = %v, want [pagination]", when.Category.Not)
}

// Check Tags
if len(when.Tags.Is) != 1 || when.Tags.Is[0] != "authentication" {
t.Errorf("Tags.Is = %v, want [authentication]", when.Tags.Is)
}
if len(when.Tags.Not) != 2 || when.Tags.Not[0] != "deprecated" || when.Tags.Not[1] != "experimental" {
t.Errorf("Tags.Not = %v, want [deprecated experimental]", when.Tags.Not)
}

// Check Generator (scalar short form)
if len(when.Generator.Is) != 1 || when.Generator.Is[0] != "claude-opus-4.6" {
t.Errorf("Generator.Is = %v, want [claude-opus-4.6]", when.Generator.Is)
}

// Marshal back
out, err := yaml.Marshal(&when)
if err != nil {
t.Fatalf("marshal failed: %v", err)
}

// Unmarshal again to verify round-trip
var when2 WhenClause
err = yaml.Unmarshal(out, &when2)
if err != nil {
t.Fatalf("round-trip unmarshal failed: %v", err)
}

// Verify equality (basic check)
if !matchSetEqual(when.Language, when2.Language) {
t.Errorf("round-trip Language changed")
}
if !matchSetEqual(when.Service, when2.Service) {
t.Errorf("round-trip Service changed")
}
if !matchSetEqual(when.Category, when2.Category) {
t.Errorf("round-trip Category changed")
}
if !matchSetEqual(when.Tags, when2.Tags) {
t.Errorf("round-trip Tags changed")
}
}
