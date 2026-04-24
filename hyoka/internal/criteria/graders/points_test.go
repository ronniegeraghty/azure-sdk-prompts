package graders

import (
	"context"
	"os"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/workspace"
)

// TestOutputCheckGraderPoints verifies the Phase 3 invariant: every grader
// emits at least one Point, and a grader with N sub-checks emits N Points
// whose collective Pass status matches the overall result. Specifically: 3
// configured knobs where 1 fails → result.Pass == false and len(Points) == 3.
func TestOutputCheckGraderPoints(t *testing.T) {
	cfg := &OutputCheckConfig{
		MinFiles:     1,
		RequireFiles: []string{"present.go"},
		ForbidFiles:  []string{"present.go"}, // intentional fail: present file is also forbidden.
	}
	g, err := NewOutputCheckGrader("oc", cfg)
	if err != nil {
		t.Fatalf("NewOutputCheckGrader: %v", err)
	}
	delta := &WorkspaceDelta{
		NewFiles: []workspace.NewFile{{Path: "present.go", Size: 100}},
	}
	res, err := g.Grade(context.Background(), GraderInput{
		WorkspaceDelta: delta,
		Config:         GraderConfig{Weight: 1.0},
	})
	if err != nil {
		t.Fatalf("Grade: %v", err)
	}
	if res.Pass {
		t.Errorf("expected Pass=false (forbid_files violated), got true; message=%q", res.Message)
	}
	if got, want := len(res.Points), 3; got != want {
		t.Fatalf("len(Points) = %d, want %d (one per configured knob); points=%+v", got, want, res.Points)
	}
	pointByLabel := make(map[string]GraderPoint, len(res.Points))
	for _, p := range res.Points {
		pointByLabel[p.Label] = p
	}
	if p, ok := pointByLabel["min_files"]; !ok || !p.Pass {
		t.Errorf("min_files point: got %+v, want present and Pass=true", p)
	}
	if p, ok := pointByLabel["require_files"]; !ok || !p.Pass {
		t.Errorf("require_files point: got %+v, want present and Pass=true", p)
	}
	if p, ok := pointByLabel["forbid_files"]; !ok || p.Pass {
		t.Errorf("forbid_files point: got %+v, want present and Pass=false", p)
	}
}

// TestProgramGraderEmitsSinglePoint ensures the program grader maintains the
// one-Point invariant for its single "exit code 0" check.
func TestProgramGraderEmitsSinglePoint(t *testing.T) {
	g, err := NewProgramGrader("ok", &ProgramConfig{Command: "true"})
	if err != nil {
		t.Fatalf("NewProgramGrader: %v", err)
	}
	res, err := g.Grade(context.Background(), GraderInput{
		WorkspacePath: t.TempDir(),
		Config:        GraderConfig{Weight: 1.0},
	})
	if err != nil {
		t.Fatalf("Grade: %v", err)
	}
	if !res.Pass {
		t.Fatalf("expected Pass=true for exit 0, got false; message=%q", res.Message)
	}
	if got, want := len(res.Points), 1; got != want {
		t.Fatalf("len(Points) = %d, want %d", got, want)
	}
	if res.Points[0].Label != "exit code 0" || !res.Points[0].Pass {
		t.Errorf("Point[0] = %+v, want Name=\"exit code 0\" Pass=true", res.Points[0])
	}
}

// TestBehaviorGraderPointsPerConstraint ensures the behavior grader emits one
// Point per configured constraint (required/forbidden tools + max_turns).
func TestBehaviorGraderPointsPerConstraint(t *testing.T) {
	g, err := NewBehaviorGrader("b", map[string]any{
		"required_tools":  []string{"read", "write"},
		"forbidden_tools": []string{"network"},
		"max_turns":       10,
	})
	if err != nil {
		t.Fatalf("NewBehaviorGrader: %v", err)
	}
	log := []ActionEvent{
		{Tool: "read", TurnNumber: 1},
		{Tool: "write", TurnNumber: 2},
	}
	res, err := g.Grade(context.Background(), GraderInput{ActionLog: log, Config: GraderConfig{Weight: 1.0}})
	if err != nil {
		t.Fatalf("Grade: %v", err)
	}
	if !res.Pass {
		t.Fatalf("expected Pass=true, got false; message=%q", res.Message)
	}
	// 2 required + 1 forbidden + 1 max_turns = 4 points.
	if got, want := len(res.Points), 4; got != want {
		t.Fatalf("len(Points) = %d, want %d; points=%+v", got, want, res.Points)
	}
	for _, p := range res.Points {
		if !p.Pass {
			t.Errorf("expected all points to pass, got %+v", p)
		}
	}
}


// TestEveryGraderEmitsPointsOnPassAndFail is the Phase 3 cross-grader
// invariant test: every concrete grader kind (file, program, behavior,
// action_sequence, tool_constraint, output_check) must emit at least one
// GraderPoint in BOTH a passing and a failing scenario. Without this the
// site renderer falls back to "PASS"/"100%" stubs whenever a grader forgets
// to populate Points. The prompt and prompt_review kinds are exercised
// through their own dedicated tests because they require a stub LLM caller
// and a stub reviewer respectively.
func TestEveryGraderEmitsPointsOnPassAndFail(t *testing.T) {
tmp := t.TempDir()
if err := writeFile(tmp+"/exists.txt", "hello"); err != nil {
t.Fatalf("seed file: %v", err)
}

delta := &WorkspaceDelta{
NewFiles: []workspace.NewFile{{Path: "exists.txt", Size: 5}},
}

cases := []struct {
name  string
build func() Grader
input GraderInput
}{
{
name: "file_pass",
build: func() Grader {
g, err := NewFileGrader("f", &FileConfig{Path: "exists.txt"})
if err != nil {
t.Fatalf("file grader: %v", err)
}
return g
},
input: GraderInput{WorkspacePath: tmp},
},
{
name: "file_fail",
build: func() Grader {
g, err := NewFileGrader("f", &FileConfig{Path: "missing.txt"})
if err != nil {
t.Fatalf("file grader: %v", err)
}
return g
},
input: GraderInput{WorkspacePath: tmp},
},
{
name: "program_pass",
build: func() Grader {
g, err := NewProgramGrader("p", &ProgramConfig{Command: "true"})
if err != nil {
t.Fatalf("program grader: %v", err)
}
return g
},
input: GraderInput{WorkspacePath: tmp},
},
{
name: "program_fail",
build: func() Grader {
g, err := NewProgramGrader("p", &ProgramConfig{Command: "false"})
if err != nil {
t.Fatalf("program grader: %v", err)
}
return g
},
input: GraderInput{WorkspacePath: tmp},
},
{
name: "behavior_pass",
build: func() Grader {
g, err := NewBehaviorGrader("b", map[string]any{
"required_tools": []string{"read_file"},
})
if err != nil {
t.Fatalf("behavior grader: %v", err)
}
return g
},
input: GraderInput{ActionLog: []ActionEvent{{Tool: "read_file", TurnNumber: 1}}},
},
{
name: "behavior_fail",
build: func() Grader {
g, err := NewBehaviorGrader("b", map[string]any{
"required_tools": []string{"missing_tool"},
})
if err != nil {
t.Fatalf("behavior grader: %v", err)
}
return g
},
input: GraderInput{ActionLog: []ActionEvent{{Tool: "read_file", TurnNumber: 1}}},
},
{
name: "behavior_no_constraints",
build: func() Grader {
g, err := NewBehaviorGrader("b", map[string]any{})
if err != nil {
t.Fatalf("behavior grader: %v", err)
}
return g
},
input: GraderInput{ActionLog: []ActionEvent{{Tool: "any", TurnNumber: 1}}},
},
{
name: "action_sequence_pass",
build: func() Grader {
g, err := NewActionSequenceGrader("as", map[string]any{
"expected_actions": []string{"read", "write"},
})
if err != nil {
t.Fatalf("action_sequence grader: %v", err)
}
return g
},
input: GraderInput{ActionLog: []ActionEvent{{Tool: "read"}, {Tool: "write"}}},
},
{
name: "action_sequence_fail",
build: func() Grader {
g, err := NewActionSequenceGrader("as", map[string]any{
"expected_actions": []string{"read", "write"},
})
if err != nil {
t.Fatalf("action_sequence grader: %v", err)
}
return g
},
input: GraderInput{ActionLog: []ActionEvent{{Tool: "read"}}},
},
{
name: "tool_constraint_pass",
build: func() Grader {
g, err := NewToolConstraintGrader("tc", map[string]any{
"required": []string{"read"},
})
if err != nil {
t.Fatalf("tool_constraint grader: %v", err)
}
return g
},
input: GraderInput{ActionLog: []ActionEvent{{Tool: "read"}}},
},
{
name: "tool_constraint_fail",
build: func() Grader {
g, err := NewToolConstraintGrader("tc", map[string]any{
"forbidden": []string{"bad"},
})
if err != nil {
t.Fatalf("tool_constraint grader: %v", err)
}
return g
},
input: GraderInput{ActionLog: []ActionEvent{{Tool: "bad"}}},
},
{
name: "tool_constraint_no_constraints",
build: func() Grader {
g, err := NewToolConstraintGrader("tc", map[string]any{})
if err != nil {
t.Fatalf("tool_constraint grader: %v", err)
}
return g
},
input: GraderInput{ActionLog: []ActionEvent{{Tool: "x"}}},
},
{
name: "output_check_pass",
build: func() Grader {
g, err := NewOutputCheckGrader("oc", &OutputCheckConfig{MinFiles: 1})
if err != nil {
t.Fatalf("output_check grader: %v", err)
}
return g
},
input: GraderInput{WorkspaceDelta: delta},
},
{
name: "output_check_fail",
build: func() Grader {
g, err := NewOutputCheckGrader("oc", &OutputCheckConfig{MinFiles: 99})
if err != nil {
t.Fatalf("output_check grader: %v", err)
}
return g
},
input: GraderInput{WorkspaceDelta: delta},
},
{
name: "output_check_no_knobs",
build: func() Grader {
g, err := NewOutputCheckGrader("oc", &OutputCheckConfig{})
if err != nil {
t.Fatalf("output_check grader: %v", err)
}
return g
},
input: GraderInput{WorkspaceDelta: delta},
},
}

for _, tc := range cases {
t.Run(tc.name, func(t *testing.T) {
g := tc.build()
res, err := g.Grade(context.Background(), tc.input)
if err != nil {
t.Fatalf("Grade: %v", err)
}
if len(res.Points) == 0 {
t.Fatalf("grader %s emitted zero Points — Phase 3 invariant violated", tc.name)
}
for i, p := range res.Points {
if p.Label == "" {
t.Errorf("grader %s Point[%d] has empty Label", tc.name, i)
}
}
})
}
}

// TestNewErrorResult_AlwaysEmitsPoint verifies the Phase 3 fallback
// constructor always synthesizes a single failing Point.
func TestNewErrorResult_AlwaysEmitsPoint(t *testing.T) {
r := NewErrorResult(KindFile, "boom", GraderConfig{Weight: 1.0}, "kaboom")
if len(r.Points) != 1 {
t.Fatalf("expected 1 Point, got %d", len(r.Points))
}
if r.Points[0].Pass {
t.Errorf("expected failing Point, got Pass=true")
}
if r.Pass {
t.Errorf("expected overall Pass=false, got true")
}
if r.Score != 0 {
t.Errorf("expected Score=0, got %v", r.Score)
}
}

// TestNewResult_PanicsOnEmptyPoints documents the loud failure mode for
// graders that forget to emit a Point.
func TestNewResult_PanicsOnEmptyPoints(t *testing.T) {
defer func() {
if r := recover(); r == nil {
t.Fatal("expected panic on empty Points, got nil")
}
}()
NewResult(KindFile, "x", GraderConfig{}, nil, "msg", nil)
}

// writeFile is a tiny test helper so we don't need to import os in only one place.
func writeFile(path, content string) error {
return os.WriteFile(path, []byte(content), 0644)
}
