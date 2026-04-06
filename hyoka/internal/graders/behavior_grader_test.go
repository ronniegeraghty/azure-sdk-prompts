package graders

import (
"context"
"testing"
)

func TestBehaviorGrader_AllConstraintsSatisfied(t *testing.T) {
g, err := NewBehaviorGrader("test", map[string]any{
"required_tools":  []any{"read_file", "edit_file"},
"forbidden_tools": []any{"rm", "sudo"},
"max_turns":       25,
})
if err != nil {
t.Fatalf("NewBehaviorGrader: %v", err)
}
result, err := g.Grade(context.Background(), GraderInput{
ActionLog: []ActionEvent{
{Tool: "read_file", TurnNumber: 1},
{Tool: "edit_file", TurnNumber: 2},
{Tool: "azure-mcp", TurnNumber: 3},
},
})
if err != nil {
t.Fatalf("Grade: %v", err)
}
if !result.Pass {
t.Errorf("expected Pass=true, message: %s", result.Message)
}
if result.Score != 1.0 {
t.Errorf("expected Score=1.0, got %f", result.Score)
}
if result.Kind != KindBehavior {
t.Errorf("expected Kind=%s, got %s", KindBehavior, result.Kind)
}
}

func TestBehaviorGrader_RequiredToolMissing(t *testing.T) {
g, _ := NewBehaviorGrader("test", map[string]any{
"required_tools": []any{"read_file", "azure-mcp"},
})
result, err := g.Grade(context.Background(), GraderInput{
ActionLog: []ActionEvent{
{Tool: "read_file", TurnNumber: 1},
{Tool: "edit_file", TurnNumber: 2},
},
})
if err != nil {
t.Fatal(err)
}
if result.Pass {
t.Error("expected Pass=false when required tool is missing")
}
if result.Score != 0.0 {
t.Errorf("expected Score=0.0, got %f", result.Score)
}
details := result.BehaviorDetails
if details == nil {
t.Fatal("expected BehaviorDetails to be set")
}
if len(details.MissingTools) != 1 || details.MissingTools[0] != "azure-mcp" {
t.Errorf("expected MissingTools=[azure-mcp], got %v", details.MissingTools)
}
}

func TestBehaviorGrader_ForbiddenToolUsed(t *testing.T) {
g, _ := NewBehaviorGrader("test", map[string]any{
"forbidden_tools": []any{"rm", "sudo"},
})
result, err := g.Grade(context.Background(), GraderInput{
ActionLog: []ActionEvent{
{Tool: "read_file", TurnNumber: 1},
{Tool: "rm", TurnNumber: 2},
},
})
if err != nil {
t.Fatal(err)
}
if result.Pass {
t.Error("expected Pass=false when forbidden tool is used")
}
details := result.BehaviorDetails
if details == nil {
t.Fatal("expected BehaviorDetails")
}
if len(details.ForbiddenUsed) != 1 || details.ForbiddenUsed[0] != "rm" {
t.Errorf("expected ForbiddenUsed=[rm], got %v", details.ForbiddenUsed)
}
}

func TestBehaviorGrader_TurnLimitExceeded(t *testing.T) {
g, _ := NewBehaviorGrader("test", map[string]any{"max_turns": 3})
result, err := g.Grade(context.Background(), GraderInput{
ActionLog: []ActionEvent{
{Tool: "read_file", TurnNumber: 1},
{Tool: "edit_file", TurnNumber: 4},
},
})
if err != nil {
t.Fatal(err)
}
if result.Pass {
t.Error("expected Pass=false when turn limit exceeded")
}
if result.BehaviorDetails == nil || !result.BehaviorDetails.TurnLimitHit {
t.Error("expected TurnLimitHit=true")
}
}

func TestBehaviorGrader_TurnLimitNotExceeded(t *testing.T) {
g, _ := NewBehaviorGrader("test", map[string]any{"max_turns": 10})
result, err := g.Grade(context.Background(), GraderInput{
ActionLog: []ActionEvent{{Tool: "read_file", TurnNumber: 5}},
})
if err != nil {
t.Fatal(err)
}
if !result.Pass {
t.Errorf("expected Pass=true, message: %s", result.Message)
}
}

func TestBehaviorGrader_EmptyActionLog(t *testing.T) {
g, _ := NewBehaviorGrader("test", map[string]any{"required_tools": []any{"read_file"}})
result, err := g.Grade(context.Background(), GraderInput{})
if err != nil {
t.Fatal(err)
}
if result.Pass {
t.Error("expected Pass=false with empty log and required tools")
}
}

func TestBehaviorGrader_NoConstraints(t *testing.T) {
g, _ := NewBehaviorGrader("test", map[string]any{})
result, err := g.Grade(context.Background(), GraderInput{
ActionLog: []ActionEvent{{Tool: "anything", TurnNumber: 100}},
})
if err != nil {
t.Fatal(err)
}
if !result.Pass {
t.Error("expected Pass=true with no constraints")
}
}

func TestBehaviorGrader_ValidationErrors(t *testing.T) {
tests := []struct {
name string
gn   string
cfg  map[string]any
}{
{"empty name", "", map[string]any{}},
{"bad required_tools", "x", map[string]any{"required_tools": "not-a-slice"}},
{"bad forbidden_tools", "x", map[string]any{"forbidden_tools": 42}},
{"bad max_turns type", "x", map[string]any{"max_turns": "fast"}},
{"negative max_turns", "x", map[string]any{"max_turns": -1}},
}
for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
_, err := NewBehaviorGrader(tt.gn, tt.cfg)
if err == nil {
t.Error("expected error")
}
})
}
}

func TestBehaviorGrader_ImplementsGraderInterface(t *testing.T) {
g, _ := NewBehaviorGrader("iface", map[string]any{})
var _ Grader = g
if g.Kind() != KindBehavior {
t.Errorf("Kind() = %s, want %s", g.Kind(), KindBehavior)
}
}

// --- ActionSequenceGrader tests ---

func TestActionSequenceGrader_FullMatch(t *testing.T) {
g, _ := NewActionSequenceGrader("test", map[string]any{
"expected_actions": []any{"read_file", "edit_file", "run_tests"},
})
result, err := g.Grade(context.Background(), GraderInput{
ActionLog: []ActionEvent{
{Tool: "read_file"}, {Tool: "analyze"}, {Tool: "edit_file"},
{Tool: "format"}, {Tool: "run_tests"},
},
})
if err != nil {
t.Fatal(err)
}
if !result.Pass {
t.Errorf("expected Pass=true, message: %s", result.Message)
}
if result.Score != 1.0 {
t.Errorf("expected Score=1.0, got %f", result.Score)
}
if result.Kind != KindActionSequence {
t.Errorf("expected Kind=%s, got %s", KindActionSequence, result.Kind)
}
details := result.BehaviorDetails
if details == nil || !details.SequenceMatch {
t.Error("expected SequenceMatch=true")
}
if details.MatchedActions != 3 {
t.Errorf("expected MatchedActions=3, got %d", details.MatchedActions)
}
}

func TestActionSequenceGrader_ExactMatch(t *testing.T) {
g, _ := NewActionSequenceGrader("test", map[string]any{
"expected_actions": []any{"read_file", "edit_file"},
})
result, _ := g.Grade(context.Background(), GraderInput{
ActionLog: []ActionEvent{{Tool: "read_file"}, {Tool: "edit_file"}},
})
if !result.Pass {
t.Error("expected Pass=true for exact match")
}
}

func TestActionSequenceGrader_MissingStep(t *testing.T) {
g, _ := NewActionSequenceGrader("test", map[string]any{
"expected_actions": []any{"read_file", "edit_file", "run_tests"},
})
result, _ := g.Grade(context.Background(), GraderInput{
ActionLog: []ActionEvent{{Tool: "read_file"}, {Tool: "edit_file"}},
})
if result.Pass {
t.Error("expected Pass=false when step is missing")
}
details := result.BehaviorDetails
if details == nil || details.MatchedActions != 2 {
t.Errorf("expected MatchedActions=2, got %v", details)
}
expected := 2.0 / 3.0
if result.Score < expected-0.01 || result.Score > expected+0.01 {
t.Errorf("expected Score~=%f, got %f", expected, result.Score)
}
}

func TestActionSequenceGrader_WrongOrder(t *testing.T) {
g, _ := NewActionSequenceGrader("test", map[string]any{
"expected_actions": []any{"edit_file", "read_file"},
})
result, _ := g.Grade(context.Background(), GraderInput{
ActionLog: []ActionEvent{{Tool: "read_file"}, {Tool: "edit_file"}},
})
if result.Pass {
t.Error("expected Pass=false for wrong order")
}
}

func TestActionSequenceGrader_ExtraActionsOK(t *testing.T) {
g, _ := NewActionSequenceGrader("test", map[string]any{
"expected_actions": []any{"read_file"},
})
result, _ := g.Grade(context.Background(), GraderInput{
ActionLog: []ActionEvent{
{Tool: "setup"}, {Tool: "read_file"}, {Tool: "analyze"},
{Tool: "edit_file"}, {Tool: "cleanup"},
},
})
if !result.Pass {
t.Error("expected Pass=true, extra actions should not matter")
}
}

func TestActionSequenceGrader_EmptyLog(t *testing.T) {
g, _ := NewActionSequenceGrader("test", map[string]any{"expected_actions": []any{"a"}})
result, _ := g.Grade(context.Background(), GraderInput{})
if result.Pass {
t.Error("expected Pass=false with empty log")
}
if result.Score != 0.0 {
t.Errorf("expected Score=0.0, got %f", result.Score)
}
}

func TestActionSequenceGrader_ValidationErrors(t *testing.T) {
tests := []struct {
name string
gn   string
cfg  map[string]any
}{
{"empty name", "", map[string]any{"expected_actions": []any{"a"}}},
{"missing key", "x", map[string]any{}},
{"empty slice", "x", map[string]any{"expected_actions": []any{}}},
{"bad type", "x", map[string]any{"expected_actions": "not-a-slice"}},
}
for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
_, err := NewActionSequenceGrader(tt.gn, tt.cfg)
if err == nil {
t.Error("expected error")
}
})
}
}

func TestActionSequenceGrader_ImplementsGraderInterface(t *testing.T) {
g, _ := NewActionSequenceGrader("iface", map[string]any{"expected_actions": []any{"a"}})
var _ Grader = g
if g.Kind() != KindActionSequence {
t.Errorf("Kind() = %s, want %s", g.Kind(), KindActionSequence)
}
}

// --- ToolConstraintGrader tests ---

func TestToolConstraintGrader_AllSatisfied(t *testing.T) {
g, _ := NewToolConstraintGrader("test", map[string]any{
"required":  []any{"azure-mcp", "read_file"},
"forbidden": []any{"rm"},
"min_calls": map[string]any{"azure-mcp": 2},
"max_calls": map[string]any{"azure-mcp": 10},
})
result, err := g.Grade(context.Background(), GraderInput{
ActionLog: []ActionEvent{
{Tool: "azure-mcp"}, {Tool: "read_file"}, {Tool: "azure-mcp"},
{Tool: "edit_file"}, {Tool: "azure-mcp"},
},
})
if err != nil {
t.Fatal(err)
}
if !result.Pass {
t.Errorf("expected Pass=true, message: %s", result.Message)
}
if result.Score != 1.0 {
t.Errorf("expected Score=1.0, got %f", result.Score)
}
if result.Kind != KindToolConstraint {
t.Errorf("expected Kind=%s, got %s", KindToolConstraint, result.Kind)
}
details := result.BehaviorDetails
if details == nil || details.ToolCounts["azure-mcp"] != 3 {
t.Errorf("expected azure-mcp count=3")
}
if !details.ConstraintsMet {
t.Error("expected ConstraintsMet=true")
}
}

func TestToolConstraintGrader_RequiredMissing(t *testing.T) {
g, _ := NewToolConstraintGrader("test", map[string]any{"required": []any{"azure-mcp"}})
result, _ := g.Grade(context.Background(), GraderInput{
ActionLog: []ActionEvent{{Tool: "read_file"}},
})
if result.Pass {
t.Error("expected Pass=false when required tool missing")
}
}

func TestToolConstraintGrader_ForbiddenUsed(t *testing.T) {
g, _ := NewToolConstraintGrader("test", map[string]any{"forbidden": []any{"dangerous"}})
result, _ := g.Grade(context.Background(), GraderInput{
ActionLog: []ActionEvent{{Tool: "read_file"}, {Tool: "dangerous"}},
})
if result.Pass {
t.Error("expected Pass=false when forbidden tool used")
}
details := result.BehaviorDetails
if details == nil || len(details.Violations) == 0 {
t.Error("expected violations")
}
}

func TestToolConstraintGrader_BelowMinCalls(t *testing.T) {
g, _ := NewToolConstraintGrader("test", map[string]any{
"min_calls": map[string]any{"azure-mcp": 3},
})
result, _ := g.Grade(context.Background(), GraderInput{
ActionLog: []ActionEvent{{Tool: "azure-mcp"}, {Tool: "azure-mcp"}},
})
if result.Pass {
t.Error("expected Pass=false when below min calls")
}
}

func TestToolConstraintGrader_ExceedsMaxCalls(t *testing.T) {
g, _ := NewToolConstraintGrader("test", map[string]any{
"max_calls": map[string]any{"azure-mcp": 2},
})
result, _ := g.Grade(context.Background(), GraderInput{
ActionLog: []ActionEvent{{Tool: "azure-mcp"}, {Tool: "azure-mcp"}, {Tool: "azure-mcp"}},
})
if result.Pass {
t.Error("expected Pass=false when exceeding max calls")
}
}

func TestToolConstraintGrader_NoConstraints(t *testing.T) {
g, _ := NewToolConstraintGrader("test", map[string]any{})
result, _ := g.Grade(context.Background(), GraderInput{
ActionLog: []ActionEvent{{Tool: "anything"}},
})
if !result.Pass {
t.Error("expected Pass=true with no constraints")
}
}

func TestToolConstraintGrader_ValidationErrors(t *testing.T) {
tests := []struct {
name string
gn   string
cfg  map[string]any
}{
{"empty name", "", map[string]any{}},
{"bad required", "x", map[string]any{"required": "not-a-slice"}},
{"bad forbidden", "x", map[string]any{"forbidden": 42}},
{"bad min_calls", "x", map[string]any{"min_calls": "not-a-map"}},
{"bad max_calls", "x", map[string]any{"max_calls": 42}},
}
for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
_, err := NewToolConstraintGrader(tt.gn, tt.cfg)
if err == nil {
t.Error("expected error")
}
})
}
}

func TestToolConstraintGrader_ImplementsGraderInterface(t *testing.T) {
g, _ := NewToolConstraintGrader("iface", map[string]any{})
var _ Grader = g
if g.Kind() != KindToolConstraint {
t.Errorf("Kind() = %s, want %s", g.Kind(), KindToolConstraint)
}
}
