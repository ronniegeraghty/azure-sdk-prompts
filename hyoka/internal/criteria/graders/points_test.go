package graders

import (
"context"
"os"
"testing"

"github.com/ronniegeraghty/hyoka/hyoka/internal/workspace"
)

// TestWorkspaceGraderPoints verifies the Phase 3 invariant: every grader
// emits at least one Point, and a grader with N sub-checks emits N Points
// whose collective Pass status matches the overall result. Specifically: 3
// configured checks where 1 fails → result.Pass == false and len(Points) == 3.
func TestWorkspaceGraderPoints(t *testing.T) {
cfg := &WorkspaceConfig{
Checks: []WorkspaceCheck{
{Kind: "require_to_create", Files: []string{"present.go"}},
{Kind: "forbidden_to_create", Files: []string{"present.go"}}, // intentional fail
{Kind: "file", Name: "present.go", State: "present", MinBytes: new(int64)},
},
}
*cfg.Checks[2].MinBytes = 1
g, err := NewWorkspaceGrader("ws", cfg)
if err != nil {
t.Fatalf("NewWorkspaceGrader: %v", err)
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
t.Errorf("expected Pass=false (forbidden_to_create violated), got true; message=%q", res.Message)
}
if got, want := len(res.Checks), 3; got != want {
t.Fatalf("len(Points) = %d, want %d (one per configured check); points=%+v", got, want, res.Checks)
}
failedCount := 0
for _, p := range res.Checks {
if !p.Pass {
failedCount++
}
}
if failedCount != 1 {
t.Errorf("expected 1 failed point, got %d", failedCount)
}
}

// TestProgramGraderEmitsSinglePoint ensures the program grader maintains the
// one-Point invariant for its single "exit code 0" check.
func TestProgramGraderEmitsSinglePoint(t *testing.T) {
g, err := NewProgramGrader("ok", &ProgramConfig{
Checks: []ProgramCheck{
{Kind: "command", Command: "true"},
},
})
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
if got, want := len(res.Checks), 1; got != want {
t.Fatalf("len(Points) = %d, want %d", got, want)
}
if res.Checks[0].Label != "check 1: true" || !res.Checks[0].Pass {
t.Errorf("Point[0] = %+v, want Pass=true", res.Checks[0])
}
}

func writeFile(path, content string) error {
return os.WriteFile(path, []byte(content), 0644)
}
