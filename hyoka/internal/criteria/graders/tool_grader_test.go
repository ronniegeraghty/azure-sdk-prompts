package graders

import (
	"context"
	"testing"
)

func TestToolGraderSpecificTool(t *testing.T) {
	cfg := &ToolConfig{
		Checks: []ToolCheckRule{
			{Kind: "specific_tool", Name: "bash"},
			{Kind: "specific_tool", Name: "edit"},
		},
	}
	g, err := NewToolGrader("test", cfg)
	if err != nil {
		t.Fatalf("NewToolGrader: %v", err)
	}

	input := GraderInput{
		ActionLog: []ActionEvent{
			{Tool: "bash", TurnNumber: 1},
			{Tool: "bash", TurnNumber: 2},
		},
	}

	res, err := g.Grade(context.Background(), input)
	if err != nil {
		t.Fatalf("Grade: %v", err)
	}

	if len(res.Checks) != 2 {
		t.Fatalf("expected 2 checks, got %d", len(res.Checks))
	}

	// bash was used — should pass
	if !res.Checks[0].Pass {
		t.Errorf("check 0 (bash) should pass")
	}

	// edit was not used — should fail
	if res.Checks[1].Pass {
		t.Errorf("check 1 (edit) should fail")
	}
}

func TestToolGraderToolNotUsed(t *testing.T) {
	cfg := &ToolConfig{
		Checks: []ToolCheckRule{
			{Kind: "tool_not_used", Name: "dangerous_tool"},
		},
	}
	g, err := NewToolGrader("test", cfg)
	if err != nil {
		t.Fatalf("NewToolGrader: %v", err)
	}

	input := GraderInput{
		ActionLog: []ActionEvent{
			{Tool: "bash", TurnNumber: 1},
		},
	}

	res, err := g.Grade(context.Background(), input)
	if err != nil {
		t.Fatalf("Grade: %v", err)
	}

	if len(res.Checks) != 1 {
		t.Fatalf("expected 1 check, got %d", len(res.Checks))
	}

	// dangerous_tool not used — should pass
	if !res.Checks[0].Pass {
		t.Errorf("check should pass (tool not used)")
	}

	// Test when tool WAS used
	input2 := GraderInput{
		ActionLog: []ActionEvent{
			{Tool: "dangerous_tool", TurnNumber: 1},
		},
	}

	res2, err := g.Grade(context.Background(), input2)
	if err != nil {
		t.Fatalf("Grade2: %v", err)
	}

	// dangerous_tool was used — should fail
	if res2.Checks[0].Pass {
		t.Errorf("check should fail (tool was used)")
	}
}

func TestToolGraderAnyOfGroup(t *testing.T) {
	cfg := &ToolConfig{
		Checks: []ToolCheckRule{
			{Kind: "any_of_group", Group: "mcp"},
		},
	}
	g, err := NewToolGrader("test", cfg)
	if err != nil {
		t.Fatalf("NewToolGrader: %v", err)
	}

	input := GraderInput{
		ActionLog: []ActionEvent{
			{Tool: "azure-mcp-server-list_resources", TurnNumber: 1},
		},
		EnvironmentTools: []EnvironmentTool{
			{Name: "azure-mcp-server-list_resources", Kind: "mcp"},
		},
	}

	res, err := g.Grade(context.Background(), input)
	if err != nil {
		t.Fatalf("Grade: %v", err)
	}

	if len(res.Checks) != 1 {
		t.Fatalf("expected 1 check, got %d", len(res.Checks))
	}

	// MCP tool was used — should pass
	if !res.Checks[0].Pass {
		t.Errorf("check should pass (MCP tool used)")
	}
}

func TestToolGraderGroupNotUsed(t *testing.T) {
	cfg := &ToolConfig{
		Checks: []ToolCheckRule{
			{Kind: "group_not_used", Group: "mcp"},
		},
	}
	g, err := NewToolGrader("test", cfg)
	if err != nil {
		t.Fatalf("NewToolGrader: %v", err)
	}

	input := GraderInput{
		ActionLog: []ActionEvent{
			{Tool: "bash", TurnNumber: 1},
		},
		EnvironmentTools: []EnvironmentTool{
			{Name: "azure-mcp", Kind: "mcp"},
		},
	}

	res, err := g.Grade(context.Background(), input)
	if err != nil {
		t.Fatalf("Grade: %v", err)
	}

	if len(res.Checks) != 1 {
		t.Fatalf("expected 1 check, got %d", len(res.Checks))
	}

	// No MCP tool used — should pass
	if !res.Checks[0].Pass {
		t.Errorf("check should pass (no MCP tool used)")
	}
}

func TestToolGraderTurnLimit(t *testing.T) {
	cfg := &ToolConfig{
		Checks: []ToolCheckRule{
			{Kind: "turn_limit", N: 5},
		},
	}
	g, err := NewToolGrader("test", cfg)
	if err != nil {
		t.Fatalf("NewToolGrader: %v", err)
	}

	// Within limit
	input := GraderInput{
		ActionLog: []ActionEvent{
			{Tool: "bash", TurnNumber: 1},
			{Tool: "bash", TurnNumber: 3},
		},
	}

	res, err := g.Grade(context.Background(), input)
	if err != nil {
		t.Fatalf("Grade: %v", err)
	}

	if len(res.Checks) != 1 {
		t.Fatalf("expected 1 check, got %d", len(res.Checks))
	}

	// Turn 3 <= 5 — should pass
	if !res.Checks[0].Pass {
		t.Errorf("check should pass (within limit)")
	}

	// Over limit
	input2 := GraderInput{
		ActionLog: []ActionEvent{
			{Tool: "bash", TurnNumber: 6},
		},
	}

	res2, err := g.Grade(context.Background(), input2)
	if err != nil {
		t.Fatalf("Grade2: %v", err)
	}

	// Turn 6 > 5 — should fail
	if res2.Checks[0].Pass {
		t.Errorf("check should fail (over limit)")
	}
}

func TestToolGraderMinCalls(t *testing.T) {
	cfg := &ToolConfig{
		Checks: []ToolCheckRule{
			{Kind: "min_calls", Name: "bash", N: 3},
		},
	}
	g, err := NewToolGrader("test", cfg)
	if err != nil {
		t.Fatalf("NewToolGrader: %v", err)
	}

	// Meets minimum
	input := GraderInput{
		ActionLog: []ActionEvent{
			{Tool: "bash", TurnNumber: 1},
			{Tool: "bash", TurnNumber: 2},
			{Tool: "bash", TurnNumber: 3},
		},
	}

	res, err := g.Grade(context.Background(), input)
	if err != nil {
		t.Fatalf("Grade: %v", err)
	}

	if len(res.Checks) != 1 {
		t.Fatalf("expected 1 check, got %d", len(res.Checks))
	}

	// 3 >= 3 — should pass
	if !res.Checks[0].Pass {
		t.Errorf("check should pass (meets minimum)")
	}

	// Below minimum
	input2 := GraderInput{
		ActionLog: []ActionEvent{
			{Tool: "bash", TurnNumber: 1},
		},
	}

	res2, err := g.Grade(context.Background(), input2)
	if err != nil {
		t.Fatalf("Grade2: %v", err)
	}

	// 1 < 3 — should fail
	if res2.Checks[0].Pass {
		t.Errorf("check should fail (below minimum)")
	}
}

func TestToolGraderMaxCalls(t *testing.T) {
	cfg := &ToolConfig{
		Checks: []ToolCheckRule{
			{Kind: "max_calls", Name: "bash", N: 2},
		},
	}
	g, err := NewToolGrader("test", cfg)
	if err != nil {
		t.Fatalf("NewToolGrader: %v", err)
	}

	// Within max
	input := GraderInput{
		ActionLog: []ActionEvent{
			{Tool: "bash", TurnNumber: 1},
			{Tool: "bash", TurnNumber: 2},
		},
	}

	res, err := g.Grade(context.Background(), input)
	if err != nil {
		t.Fatalf("Grade: %v", err)
	}

	if len(res.Checks) != 1 {
		t.Fatalf("expected 1 check, got %d", len(res.Checks))
	}

	// 2 <= 2 — should pass
	if !res.Checks[0].Pass {
		t.Errorf("check should pass (within max)")
	}

	// Exceeds max
	input2 := GraderInput{
		ActionLog: []ActionEvent{
			{Tool: "bash", TurnNumber: 1},
			{Tool: "bash", TurnNumber: 2},
			{Tool: "bash", TurnNumber: 3},
		},
	}

	res2, err := g.Grade(context.Background(), input2)
	if err != nil {
		t.Fatalf("Grade2: %v", err)
	}

	// 3 > 2 — should fail
	if res2.Checks[0].Pass {
		t.Errorf("check should fail (exceeds max)")
	}
}

func TestToolGraderSkillRepoGroup(t *testing.T) {
	cfg := &ToolConfig{
		Checks: []ToolCheckRule{
			{Kind: "any_of_group", Group: "skill_repo:github.com/org/repo"},
		},
	}
	g, err := NewToolGrader("test", cfg)
	if err != nil {
		t.Fatalf("NewToolGrader: %v", err)
	}

	input := GraderInput{
		ActionLog: []ActionEvent{
			{Tool: "my-skill", TurnNumber: 1},
		},
		EnvironmentTools: []EnvironmentTool{
			{Name: "my-skill", Kind: "skill", Repo: "github.com/org/repo"},
		},
	}

	res, err := g.Grade(context.Background(), input)
	if err != nil {
		t.Fatalf("Grade: %v", err)
	}

	if len(res.Checks) != 1 {
		t.Fatalf("expected 1 check, got %d", len(res.Checks))
	}

	// Skill from repo was used — should pass
	if !res.Checks[0].Pass {
		t.Errorf("check should pass (skill from repo used)")
	}
}

func TestToolGraderValidation(t *testing.T) {
	tests := []struct {
		name        string
		cfg         *ToolConfig
		expectError bool
	}{
		{
			name: "no checks",
			cfg: &ToolConfig{
				Checks: []ToolCheckRule{},
			},
			expectError: true,
		},
		{
			name: "specific_tool missing name",
			cfg: &ToolConfig{
				Checks: []ToolCheckRule{
					{Kind: "specific_tool"},
				},
			},
			expectError: true,
		},
		{
			name: "any_of_group missing group",
			cfg: &ToolConfig{
				Checks: []ToolCheckRule{
					{Kind: "any_of_group"},
				},
			},
			expectError: true,
		},
		{
			name: "turn_limit zero N",
			cfg: &ToolConfig{
				Checks: []ToolCheckRule{
					{Kind: "turn_limit", N: 0},
				},
			},
			expectError: true,
		},
		{
			name: "min_calls negative N",
			cfg: &ToolConfig{
				Checks: []ToolCheckRule{
					{Kind: "min_calls", Name: "bash", N: -1},
				},
			},
			expectError: true,
		},
		{
			name: "unknown kind",
			cfg: &ToolConfig{
				Checks: []ToolCheckRule{
					{Kind: "unknown_kind"},
				},
			},
			expectError: true,
		},
		{
			name: "valid config",
			cfg: &ToolConfig{
				Checks: []ToolCheckRule{
					{Kind: "specific_tool", Name: "bash"},
					{Kind: "turn_limit", N: 5},
				},
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewToolGrader("test", tt.cfg)
			if tt.expectError && err == nil {
				t.Errorf("expected error, got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("expected no error, got: %v", err)
			}
		})
	}
}

func TestToolGraderInterface(t *testing.T) {
	cfg := &ToolConfig{
		Checks: []ToolCheckRule{
			{Kind: "specific_tool", Name: "bash"},
		},
	}
	g, _ := NewToolGrader("test", cfg)

	var _ Grader = g

	if g.Kind() != KindTool {
		t.Errorf("Kind() = %q, want %q", g.Kind(), KindTool)
	}
	if g.Name() != "test" {
		t.Errorf("Name() = %q, want %q", g.Name(), "test")
	}
}

func TestToolGraderToolNameGlob(t *testing.T) {
	cfg := &ToolConfig{
		Checks: []ToolCheckRule{
			{Kind: "any_of_group", Group: "tool_name_glob:azure-*"},
		},
	}
	g, err := NewToolGrader("test", cfg)
	if err != nil {
		t.Fatalf("NewToolGrader: %v", err)
	}

	input := GraderInput{
		ActionLog: []ActionEvent{
			{Tool: "azure-mcp-list", TurnNumber: 1},
		},
		EnvironmentTools: []EnvironmentTool{
			{Name: "azure-mcp-list", Kind: "mcp"},
			{Name: "aws-mcp-list", Kind: "mcp"},
		},
	}

	res, err := g.Grade(context.Background(), input)
	if err != nil {
		t.Fatalf("Grade: %v", err)
	}

	if len(res.Checks) != 1 {
		t.Fatalf("expected 1 check, got %d", len(res.Checks))
	}

	// Tool matching azure-* was used — should pass
	if !res.Checks[0].Pass {
		t.Errorf("check should pass (glob matched tool used)")
	}
}
