package graders

import (
	"context"
	"testing"
)

func TestToolUsageGrader_NewValidation(t *testing.T) {
	cases := []struct {
		name    string
		cfg     *ToolUsageConfig
		wantErr bool
	}{
		{
			name:    "nil config",
			cfg:     nil,
			wantErr: true,
		},
		{
			name: "empty type",
			cfg: &ToolUsageConfig{Rules: []ToolUsageRule{
				{Name: "x", Expect: "at_least_one_tool_call"},
			}},
			wantErr: true,
		},
		{
			name: "mcp_server missing name",
			cfg: &ToolUsageConfig{Rules: []ToolUsageRule{
				{Type: "mcp_server", Expect: "at_least_one_tool_call"},
			}},
			wantErr: true,
		},
		{
			name: "skill_repo missing repo",
			cfg: &ToolUsageConfig{Rules: []ToolUsageRule{
				{Type: "skill_repo", Expect: "any_skill_invoked"},
			}},
			wantErr: true,
		},
		{
			name: "unsupported type",
			cfg: &ToolUsageConfig{Rules: []ToolUsageRule{
				{Type: "wat", Name: "x", Expect: "any_skill_invoked"},
			}},
			wantErr: true,
		},
		{
			name: "missing expect",
			cfg: &ToolUsageConfig{Rules: []ToolUsageRule{
				{Type: "mcp_server", Name: "x"},
			}},
			wantErr: true,
		},
		{
			name:    "no rules ok",
			cfg:     &ToolUsageConfig{},
			wantErr: false,
		},
		{
			name: "valid mcp_server",
			cfg: &ToolUsageConfig{Rules: []ToolUsageRule{
				{Type: "mcp_server", Name: "azure", Expect: "at_least_one_tool_call"},
			}},
			wantErr: false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewToolUsageGrader("u", tc.cfg)
			if (err != nil) != tc.wantErr {
				t.Fatalf("err=%v wantErr=%v", err, tc.wantErr)
			}
		})
	}
}

func TestToolUsageGrader_Grade(t *testing.T) {
	cases := []struct {
		name        string
		cfg         ToolUsageConfig
		input       GraderInput
		wantPoints  int
		wantPassing int
		wantLabel0  string
	}{
		{
			name: "no applicable rules trivially passes",
			cfg: ToolUsageConfig{Rules: []ToolUsageRule{
				{Type: "mcp_server", Name: "absent-mcp", Expect: "at_least_one_tool_call"},
			}},
			input:       GraderInput{},
			wantPoints:  1,
			wantPassing: 1,
			wantLabel0:  "no_applicable_rules",
		},
		{
			name: "mcp_server used → pass",
			cfg: ToolUsageConfig{Rules: []ToolUsageRule{
				{Type: "mcp_server", Name: "azure", Expect: "at_least_one_tool_call"},
			}},
			input: GraderInput{
				EnvironmentTools: []EnvironmentTool{{Kind: "mcp", Name: "azure"}},
				MCPServersUsed:   []string{"azure"},
			},
			wantPoints:  1,
			wantPassing: 1,
		},
		{
			name: "mcp_server not used → fail",
			cfg: ToolUsageConfig{Rules: []ToolUsageRule{
				{Type: "mcp_server", Name: "azure", Expect: "at_least_one_tool_call"},
			}},
			input: GraderInput{
				EnvironmentTools: []EnvironmentTool{{Kind: "mcp", Name: "azure"}},
			},
			wantPoints:  1,
			wantPassing: 0,
		},
		{
			name: "skill_plugin invoked → pass",
			cfg: ToolUsageConfig{Rules: []ToolUsageRule{
				{Type: "skill_plugin", Name: "azure-sdk-python", Expect: "any_skill_invoked"},
			}},
			input: GraderInput{
				EnvironmentTools: []EnvironmentTool{{Kind: "skill", Name: "azure-sdk-python", Path: "skills/reviewer/azure-sdk-python"}},
				SkillsInvoked:    []string{"azure-sdk-python"},
			},
			wantPoints:  1,
			wantPassing: 1,
		},
		{
			name: "skill_repo with matching repo + invoked → pass",
			cfg: ToolUsageConfig{Rules: []ToolUsageRule{
				{Type: "skill_repo", Repo: "mauromedda/agent-toolkit", Expect: "any_skill_invoked"},
			}},
			input: GraderInput{
				EnvironmentTools: []EnvironmentTool{{Kind: "skill", Name: "python-style", Repo: "mauromedda/agent-toolkit", Path: ".hyoka/skills/python-style"}},
				SkillsInvoked:    []string{"python-style"},
			},
			wantPoints:  1,
			wantPassing: 1,
		},
		{
			name: "generator-dir skill is excluded",
			cfg: ToolUsageConfig{Rules: []ToolUsageRule{
				{Type: "skill_plugin", Name: "gen-only", Expect: "any_skill_invoked"},
			}},
			input: GraderInput{
				EnvironmentTools: []EnvironmentTool{{Kind: "skill", Name: "gen-only", Path: "skills/generator/gen-only"}},
				SkillsInvoked:    []string{"gen-only"},
			},
			wantPoints:  1, // falls through to no_applicable_rules
			wantPassing: 1,
			wantLabel0:  "no_applicable_rules",
		},
		{
			name: "mixed pass/fail",
			cfg: ToolUsageConfig{Rules: []ToolUsageRule{
				{Type: "mcp_server", Name: "azure", Expect: "at_least_one_tool_call"},
				{Type: "skill_plugin", Name: "azure-sdk-python", Expect: "any_skill_invoked"},
			}},
			input: GraderInput{
				EnvironmentTools: []EnvironmentTool{
					{Kind: "mcp", Name: "azure"},
					{Kind: "skill", Name: "azure-sdk-python", Path: "skills/reviewer/azure-sdk-python"},
				},
				MCPServersUsed: []string{"azure"},
			},
			wantPoints:  2,
			wantPassing: 1,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g, err := NewToolUsageGrader("u", &tc.cfg)
			if err != nil {
				t.Fatalf("new: %v", err)
			}
			res, err := g.Grade(context.Background(), tc.input)
			if err != nil {
				t.Fatalf("grade: %v", err)
			}
			if len(res.Points) != tc.wantPoints {
				t.Fatalf("points=%d want=%d (%+v)", len(res.Points), tc.wantPoints, res.Points)
			}
			passing := 0
			for _, p := range res.Points {
				if p.Pass {
					passing++
				}
			}
			if passing != tc.wantPassing {
				t.Errorf("passing=%d want=%d", passing, tc.wantPassing)
			}
			if tc.wantLabel0 != "" && res.Points[0].Label != tc.wantLabel0 {
				t.Errorf("points[0].Label=%q want=%q", res.Points[0].Label, tc.wantLabel0)
			}
		})
	}
}
