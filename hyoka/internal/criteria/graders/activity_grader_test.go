package graders

import (
"context"
"testing"

"github.com/ronniegeraghty/hyoka/hyoka/internal/artifact"
)

func TestActivityGrader_AllChecksPass(t *testing.T) {
	cfg := &ActivityConfig{
		Checks: []ActivityCheck{
			{Kind: "turn_limit", Max: intPtr(10)},
			{Kind: "action_count", Min: intPtr(1), Max: intPtr(10)},
			{Kind: "tool_call_count", Min: intPtr(1)},
			{Kind: "contains_subsequence", Tools: []string{"view", "edit"}},
			{Kind: "contains_action", Tool: "view", Min: intPtr(1)},
			{Kind: "excludes_action", Type: "error"},
			{Kind: "terminated_by", Equals: "completed"},
		},
	}
	grader, err := NewActivityGrader("test-activity", cfg)
	if err != nil {
		t.Fatalf("NewActivityGrader failed: %v", err)
	}

	input := GraderInput{
		ActionLog: []ActionEvent{
			{Type: "tool_call", Tool: "view", TurnNumber: 1},
			{Type: "tool_call", Tool: "edit", TurnNumber: 2},
		},
		GeneratorArtifact: &artifact.GeneratorArtifact{
			ActionsSummary: artifact.ActionsSummary{
				TotalActions: 5,
				ToolCalls:    2,
				Truncated:    false,
			},
			TerminatedBy: "completed",
		},
		Config: GraderConfig{Name: "test-activity", Kind: "activity"},
	}

	result, err := grader.Grade(context.Background(), input)
	if err != nil {
		t.Fatalf("Grade failed: %v", err)
	}

	if !result.Pass {
		t.Errorf("Pass = %v, want true (msg=%s)", result.Pass, result.Message)
	}
	if len(result.Checks) != 7 {
		t.Errorf("Checks count = %d, want 7", len(result.Checks))
	}
}

func TestActivityGrader_TurnLimitFail(t *testing.T) {
cfg := &ActivityConfig{
Checks: []ActivityCheck{
{Kind: "turn_limit", Max: intPtr(5)},
},
}
grader, _ := NewActivityGrader("test", cfg)

input := GraderInput{
ActionLog: []ActionEvent{
{Tool: "view", TurnNumber: 10},
},
GeneratorArtifact: &artifact.GeneratorArtifact{},
Config:            GraderConfig{},
}

result, _ := grader.Grade(context.Background(), input)
if result.Pass {
t.Error("expected Pass=false when turn limit exceeded")
}
}

func TestActivityGrader_ContainsSubsequencePass(t *testing.T) {
cfg := &ActivityConfig{
Checks: []ActivityCheck{
{Kind: "contains_subsequence", Tools: []string{"view", "edit"}},
},
}
grader, _ := NewActivityGrader("test", cfg)

input := GraderInput{
ActionLog: []ActionEvent{
{Tool: "view"},
{Tool: "grep"},
{Tool: "edit"},
},
GeneratorArtifact: &artifact.GeneratorArtifact{},
Config:            GraderConfig{},
}

result, _ := grader.Grade(context.Background(), input)
if !result.Pass {
t.Errorf("expected Pass=true; got message: %s", result.Message)
}
}

func TestActivityGrader_ContainsActionByMessageText(t *testing.T) {
	cfg := &ActivityConfig{
		Checks: []ActivityCheck{
			{Kind: "contains_action", Type: "message", Contains: "deployment"},
		},
	}
	grader, err := NewActivityGrader("test", cfg)
	if err != nil {
		t.Fatalf("NewActivityGrader: %v", err)
	}
	input := GraderInput{
		ActionLog: []ActionEvent{
			{Type: "message", Text: "Starting deployment process"},
			{Type: "tool_call", Tool: "view"},
		},
		GeneratorArtifact: &artifact.GeneratorArtifact{},
	}
	result, _ := grader.Grade(context.Background(), input)
	if !result.Pass {
		t.Errorf("expected pass, got fail: %s", result.Message)
	}
}

func TestActivityGrader_ExcludesAction(t *testing.T) {
	cfg := &ActivityConfig{
		Checks: []ActivityCheck{
			{Kind: "excludes_action", Type: "error"},
		},
	}
	grader, _ := NewActivityGrader("test", cfg)

	// No errors → pass
	input := GraderInput{
		ActionLog:         []ActionEvent{{Type: "message", Text: "hi"}},
		GeneratorArtifact: &artifact.GeneratorArtifact{},
	}
	if r, _ := grader.Grade(context.Background(), input); !r.Pass {
		t.Errorf("expected pass with no errors, got fail")
	}

	// Has an error → fail
	input.ActionLog = []ActionEvent{{Type: "error", Error: "boom"}}
	if r, _ := grader.Grade(context.Background(), input); r.Pass {
		t.Errorf("expected fail when error event present")
	}
}

func TestActivityGrader_DropsNotTruncated(t *testing.T) {
	cfg := &ActivityConfig{Checks: []ActivityCheck{{Kind: "not_truncated"}}}
	if _, err := NewActivityGrader("test", cfg); err == nil {
		t.Error("expected not_truncated to be rejected as unknown kind")
	}
}

func TestActivityGrader_TerminatedBy(t *testing.T) {
tests := []struct {
name          string
check         ActivityCheck
terminatedBy  string
wantPass      bool
}{
{"equals match", ActivityCheck{Kind: "terminated_by", Equals: "completed"}, "completed", true},
{"equals mismatch", ActivityCheck{Kind: "terminated_by", Equals: "completed"}, "error", false},
{"not_in pass", ActivityCheck{Kind: "terminated_by", NotIn: []string{"error", "max_actions"}}, "completed", true},
{"not_in fail", ActivityCheck{Kind: "terminated_by", NotIn: []string{"error", "max_actions"}}, "error", false},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
cfg := &ActivityConfig{Checks: []ActivityCheck{tt.check}}
grader, _ := NewActivityGrader("test", cfg)

input := GraderInput{
ActionLog: []ActionEvent{},
GeneratorArtifact: &artifact.GeneratorArtifact{
TerminatedBy: tt.terminatedBy,
},
Config: GraderConfig{},
}

result, _ := grader.Grade(context.Background(), input)
if result.Pass != tt.wantPass {
t.Errorf("Pass = %v, want %v", result.Pass, tt.wantPass)
}
})
}
}

func TestActivityGrader_NoChecks(t *testing.T) {
cfg := &ActivityConfig{Checks: []ActivityCheck{}}
grader, _ := NewActivityGrader("test", cfg)

input := GraderInput{
ActionLog:         []ActionEvent{},
GeneratorArtifact: &artifact.GeneratorArtifact{},
Config:            GraderConfig{},
}

result, _ := grader.Grade(context.Background(), input)
if !result.Pass {
t.Error("expected Pass=true when no checks")
}
if result.Message != "no checks configured" {
t.Errorf("unexpected message: %s", result.Message)
}
}

func intPtr(v int) *int {
return &v
}
