package graders

import (
	"context"
	"strings"
	"testing"
)

func TestToolGraderToolUsed(t *testing.T) {
	cfg := &ToolConfig{
		Checks: []ToolCheckRule{
			{Kind: "tool_used", Tool: "bash"},
			{Kind: "tool_used", Tool: "edit"},
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
			{Kind: "tool_not_used", Tool: "dangerous_tool"},
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

func TestToolGraderAnyFromGroup(t *testing.T) {
	cfg := &ToolConfig{
		Checks: []ToolCheckRule{
			{Kind: "any_from_group", Group: "mcp"},
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

func TestToolGraderNoneFromGroup(t *testing.T) {
	cfg := &ToolConfig{
		Checks: []ToolCheckRule{
			{Kind: "none_from_group", Group: "mcp"},
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

func TestToolGraderToolUsedWithMinCalls(t *testing.T) {
	minCalls := 3
	cfg := &ToolConfig{
		Checks: []ToolCheckRule{
			{Kind: "tool_used", Tool: "bash", MinCalls: &minCalls},
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

func TestToolGraderToolUsedWithMaxCalls(t *testing.T) {
	maxCalls := 2
	cfg := &ToolConfig{
		Checks: []ToolCheckRule{
			{Kind: "tool_used", Tool: "bash", MaxCalls: &maxCalls},
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
			{Kind: "any_from_group", Group: "skill_repo:github.com/org/repo"},
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
	minNegative := -1
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
			name: "tool_used missing tool",
			cfg: &ToolConfig{
				Checks: []ToolCheckRule{
					{Kind: "tool_used"},
				},
			},
			expectError: true,
		},
		{
			name: "any_from_group missing group",
			cfg: &ToolConfig{
				Checks: []ToolCheckRule{
					{Kind: "any_from_group"},
				},
			},
			expectError: true,
		},
		{
			name: "tool_used negative min_calls",
			cfg: &ToolConfig{
				Checks: []ToolCheckRule{
					{Kind: "tool_used", Tool: "bash", MinCalls: &minNegative},
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
					{Kind: "tool_used", Tool: "bash"},
					{Kind: "tool_not_used", Tool: "edit"},
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

func TestToolGraderLegacyKindMigration(t *testing.T) {
	tests := []struct {
		name          string
		kind          string
		expectError   bool
		errorContains string
	}{
		{
			name:          "specific_tool",
			kind:          "specific_tool",
			expectError:   true,
			errorContains: "tool_used",
		},
		{
			name:          "any_of_group",
			kind:          "any_of_group",
			expectError:   true,
			errorContains: "any_from_group",
		},
		{
			name:          "group_not_used",
			kind:          "group_not_used",
			expectError:   true,
			errorContains: "none_from_group",
		},
		{
			name:          "turn_limit",
			kind:          "turn_limit",
			expectError:   true,
			errorContains: "activity grader",
		},
		{
			name:          "min_calls",
			kind:          "min_calls",
			expectError:   true,
			errorContains: "tool_used",
		},
		{
			name:          "max_calls",
			kind:          "max_calls",
			expectError:   true,
			errorContains: "tool_used",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &ToolConfig{
				Checks: []ToolCheckRule{
					{Kind: tt.kind, Tool: "bash", Group: "mcp"},
				},
			}
			_, err := NewToolGrader("test", cfg)
			if tt.expectError {
				if err == nil {
					t.Errorf("expected error for legacy kind %q, got none", tt.kind)
				} else if tt.errorContains != "" && !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("expected error containing %q, got: %v", tt.errorContains, err)
				}
			} else if err != nil {
				t.Errorf("expected no error, got: %v", err)
			}
		})
	}
}

func TestToolGraderInterface(t *testing.T) {
	cfg := &ToolConfig{
		Checks: []ToolCheckRule{
			{Kind: "tool_used", Tool: "bash"},
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
			{Kind: "any_from_group", Group: "tool_name_glob:azure-*"},
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

// Test source field filtering
func TestToolGraderToolUsedWithSource(t *testing.T) {
	tests := []struct {
		name       string
		rule       ToolCheckRule
		events     []ActionEvent
		shouldPass bool
	}{
		{
			name: "source=skill matches skill event",
			rule: ToolCheckRule{Kind: "tool_used", Tool: "auth", Source: "skill"},
			events: []ActionEvent{
				{Type: "skill", Tool: "auth"},
			},
			shouldPass: true,
		},
		{
			name: "source=skill filters out mcp",
			rule: ToolCheckRule{Kind: "tool_used", Tool: "auth", Source: "skill"},
			events: []ActionEvent{
				{Type: "mcp_call", Tool: "auth"},
			},
			shouldPass: false,
		},
		{
			name: "source=mcp matches mcp_call",
			rule: ToolCheckRule{Kind: "tool_used", Tool: "list-resources", Source: "mcp"},
			events: []ActionEvent{
				{Type: "mcp_call", Tool: "list-resources"},
			},
			shouldPass: true,
		},
		{
			name: "source=mcp filters out skill",
			rule: ToolCheckRule{Kind: "tool_used", Tool: "list-resources", Source: "mcp"},
			events: []ActionEvent{
				{Type: "skill", Tool: "list-resources"},
			},
			shouldPass: false,
		},
		{
			name: "source=builtin matches tool_call",
			rule: ToolCheckRule{Kind: "tool_used", Tool: "bash", Source: "builtin"},
			events: []ActionEvent{
				{Type: "bash", Tool: "bash"},
			},
			shouldPass: true,
		},
		{
			name: "source=builtin filters out skill",
			rule: ToolCheckRule{Kind: "tool_used", Tool: "bash", Source: "builtin"},
			events: []ActionEvent{
				{Type: "skill", Tool: "bash"},
			},
			shouldPass: false,
		},
		{
			name: "no source matches any type",
			rule: ToolCheckRule{Kind: "tool_used", Tool: "edit"},
			events: []ActionEvent{
				{Type: "skill", Tool: "edit"},
			},
			shouldPass: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &ToolConfig{Checks: []ToolCheckRule{tt.rule}}
			g, err := NewToolGrader("test", cfg)
			if err != nil {
				t.Fatalf("NewToolGrader: %v", err)
			}

			input := GraderInput{ActionLog: tt.events}
			res, err := g.Grade(context.Background(), input)
			if err != nil {
				t.Fatalf("Grade: %v", err)
			}

			if len(res.Checks) != 1 {
				t.Fatalf("expected 1 check, got %d", len(res.Checks))
			}

			if res.Checks[0].Pass != tt.shouldPass {
				t.Errorf("expected Pass=%v, got %v", tt.shouldPass, res.Checks[0].Pass)
			}
		})
	}
}

// Test mcp_server field filtering
func TestToolGraderToolUsedWithMCPServer(t *testing.T) {
	tests := []struct {
		name       string
		rule       ToolCheckRule
		events     []ActionEvent
		shouldPass bool
	}{
		{
			name: "mcp_server matches",
			rule: ToolCheckRule{Kind: "tool_used", Tool: "list-resources", Source: "mcp", MCPServer: "azure-mcp"},
			events: []ActionEvent{
				{Type: "mcp_call", Tool: "list-resources", MCPServer: "azure-mcp"},
			},
			shouldPass: true,
		},
		{
			name: "mcp_server mismatch",
			rule: ToolCheckRule{Kind: "tool_used", Tool: "list-resources", Source: "mcp", MCPServer: "azure-mcp"},
			events: []ActionEvent{
				{Type: "mcp_call", Tool: "list-resources", MCPServer: "aws-mcp"},
			},
			shouldPass: false,
		},
		{
			name: "no mcp_server matches any server",
			rule: ToolCheckRule{Kind: "tool_used", Tool: "list-resources", Source: "mcp"},
			events: []ActionEvent{
				{Type: "mcp_call", Tool: "list-resources", MCPServer: "azure-mcp"},
			},
			shouldPass: true,
		},
		{
			name: "mcp_server with multiple events",
			rule: ToolCheckRule{Kind: "tool_used", Tool: "list-resources", Source: "mcp", MCPServer: "azure-mcp"},
			events: []ActionEvent{
				{Type: "mcp_call", Tool: "list-resources", MCPServer: "aws-mcp"},
				{Type: "mcp_call", Tool: "list-resources", MCPServer: "azure-mcp"},
			},
			shouldPass: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &ToolConfig{Checks: []ToolCheckRule{tt.rule}}
			g, err := NewToolGrader("test", cfg)
			if err != nil {
				t.Fatalf("NewToolGrader: %v", err)
			}

			input := GraderInput{ActionLog: tt.events}
			res, err := g.Grade(context.Background(), input)
			if err != nil {
				t.Fatalf("Grade: %v", err)
			}

			if len(res.Checks) != 1 {
				t.Fatalf("expected 1 check, got %d", len(res.Checks))
			}

			if res.Checks[0].Pass != tt.shouldPass {
				t.Errorf("expected Pass=%v, got %v", tt.shouldPass, res.Checks[0].Pass)
			}
		})
	}
}

func TestToolGraderMCPServerWideCheck(t *testing.T) {
	tests := []struct {
		name       string
		rule       ToolCheckRule
		events     []ActionEvent
		shouldPass bool
	}{
		{
			name: "matches any tool from server",
			rule: ToolCheckRule{Kind: "tool_used", Source: "mcp", MCPServer: "azure"},
			events: []ActionEvent{
				{Type: "tool_call", Tool: "azure-documentation", MCPServer: "azure"},
			},
			shouldPass: true,
		},
		{
			name: "rejects a different server",
			rule: ToolCheckRule{Kind: "tool_used", Source: "mcp", MCPServer: "azure"},
			events: []ActionEvent{
				{Type: "mcp_call", Tool: "github-search", MCPServer: "github"},
			},
			shouldPass: false,
		},
		{
			name: "tool_not_used fails for any tool from server",
			rule: ToolCheckRule{Kind: "tool_not_used", Source: "mcp", MCPServer: "azure"},
			events: []ActionEvent{
				{Type: "mcp_call", Tool: "azure-documentation", MCPServer: "azure"},
			},
			shouldPass: false,
		},
		{
			name:       "tool_not_used passes when server is absent",
			rule:       ToolCheckRule{Kind: "tool_not_used", Source: "mcp", MCPServer: "azure"},
			events:     []ActionEvent{{Type: "bash", Tool: "bash"}},
			shouldPass: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := NewToolGrader("test", &ToolConfig{Checks: []ToolCheckRule{tt.rule}})
			if err != nil {
				t.Fatalf("NewToolGrader: %v", err)
			}

			res, err := g.Grade(context.Background(), GraderInput{ActionLog: tt.events})
			if err != nil {
				t.Fatalf("Grade: %v", err)
			}
			if got := res.Checks[0].Pass; got != tt.shouldPass {
				t.Errorf("Pass = %v, want %v; message: %s", got, tt.shouldPass, res.Checks[0].Message)
			}
		})
	}
}

// Test tool_not_used with source filtering
func TestToolGraderToolNotUsedWithSource(t *testing.T) {
	tests := []struct {
		name       string
		rule       ToolCheckRule
		events     []ActionEvent
		shouldPass bool
	}{
		{
			name: "tool not used with source filter passes",
			rule: ToolCheckRule{Kind: "tool_not_used", Tool: "dangerous", Source: "skill"},
			events: []ActionEvent{
				{Type: "mcp_call", Tool: "safe"},
			},
			shouldPass: true,
		},
		{
			name: "tool used but wrong source passes",
			rule: ToolCheckRule{Kind: "tool_not_used", Tool: "dangerous", Source: "skill"},
			events: []ActionEvent{
				{Type: "mcp_call", Tool: "dangerous"},
			},
			shouldPass: true,
		},
		{
			name: "tool used with matching source fails",
			rule: ToolCheckRule{Kind: "tool_not_used", Tool: "dangerous", Source: "skill"},
			events: []ActionEvent{
				{Type: "skill", Tool: "dangerous"},
			},
			shouldPass: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &ToolConfig{Checks: []ToolCheckRule{tt.rule}}
			g, err := NewToolGrader("test", cfg)
			if err != nil {
				t.Fatalf("NewToolGrader: %v", err)
			}

			input := GraderInput{ActionLog: tt.events}
			res, err := g.Grade(context.Background(), input)
			if err != nil {
				t.Fatalf("Grade: %v", err)
			}

			if len(res.Checks) != 1 {
				t.Fatalf("expected 1 check, got %d", len(res.Checks))
			}

			if res.Checks[0].Pass != tt.shouldPass {
				t.Errorf("expected Pass=%v, got %v", tt.shouldPass, res.Checks[0].Pass)
			}
		})
	}
}

// Test source validation
func TestToolGraderSourceValidation(t *testing.T) {
	tests := []struct {
		name          string
		rule          ToolCheckRule
		expectError   bool
		errorContains string
	}{
		{
			name:        "valid source skill",
			rule:        ToolCheckRule{Kind: "tool_used", Tool: "auth", Source: "skill"},
			expectError: false,
		},
		{
			name:        "valid source mcp",
			rule:        ToolCheckRule{Kind: "tool_used", Tool: "auth", Source: "mcp"},
			expectError: false,
		},
		{
			name:        "valid source builtin",
			rule:        ToolCheckRule{Kind: "tool_used", Tool: "bash", Source: "builtin"},
			expectError: false,
		},
		{
			name:          "invalid source",
			rule:          ToolCheckRule{Kind: "tool_used", Tool: "auth", Source: "invalid"},
			expectError:   true,
			errorContains: "source must be one of: skill, mcp, builtin",
		},
		{
			name:          "mcp_server without source",
			rule:          ToolCheckRule{Kind: "tool_used", Tool: "auth", MCPServer: "azure-mcp"},
			expectError:   true,
			errorContains: "mcp_server requires source=mcp",
		},
		{
			name:          "mcp_server with source=skill",
			rule:          ToolCheckRule{Kind: "tool_used", Tool: "auth", Source: "skill", MCPServer: "azure-mcp"},
			expectError:   true,
			errorContains: "mcp_server requires source=mcp",
		},
		{
			name:        "valid mcp_server with source=mcp",
			rule:        ToolCheckRule{Kind: "tool_used", Tool: "auth", Source: "mcp", MCPServer: "azure-mcp"},
			expectError: false,
		},
		{
			name:        "valid server-wide MCP check",
			rule:        ToolCheckRule{Kind: "tool_used", Source: "mcp", MCPServer: "azure-mcp"},
			expectError: false,
		},
		{
			name:          "missing tool and mcp_server",
			rule:          ToolCheckRule{Kind: "tool_used", Source: "mcp"},
			expectError:   true,
			errorContains: "requires 'tool' or source=mcp with 'mcp_server'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &ToolConfig{Checks: []ToolCheckRule{tt.rule}}
			_, err := NewToolGrader("test", cfg)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error, got none")
				} else if tt.errorContains != "" && !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("expected error containing %q, got: %v", tt.errorContains, err)
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got: %v", err)
				}
			}
		})
	}
}
