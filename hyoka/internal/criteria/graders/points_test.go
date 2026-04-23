package graders

import (
	"context"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/workspace"
)

// TestOutputCheckGraderPoints verifies the Phase 2 invariant: every grader
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
	pointByName := make(map[string]GraderPoint, len(res.Points))
	for _, p := range res.Points {
		pointByName[p.Name] = p
	}
	if p, ok := pointByName["min_files"]; !ok || !p.Pass {
		t.Errorf("min_files point: got %+v, want present and Pass=true", p)
	}
	if p, ok := pointByName["require_files"]; !ok || !p.Pass {
		t.Errorf("require_files point: got %+v, want present and Pass=true", p)
	}
	if p, ok := pointByName["forbid_files"]; !ok || p.Pass {
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
	if res.Points[0].Name != "exit code 0" || !res.Points[0].Pass {
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

