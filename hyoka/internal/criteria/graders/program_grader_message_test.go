package graders

import (
"context"
"testing"
)

func TestProgramGrader_MessageFormat(t *testing.T) {
tests := []struct {
name        string
checks      []ProgramCheck
wantMessage string
}{
{
name: "one check pass",
checks: []ProgramCheck{
{Kind: "command", Command: "true"},
},
wantMessage: "program checks: 1/1 passed",
},
{
name: "two checks mixed",
checks: []ProgramCheck{
{Kind: "command", Command: "true"},
{Kind: "command", Command: "false"},
},
wantMessage: "program checks: 1/2 passed",
},
{
name: "three checks all pass",
checks: []ProgramCheck{
{Kind: "command", Command: "echo", Args: []string{"a"}},
{Kind: "command", Command: "echo", Args: []string{"b"}},
{Kind: "command", Command: "echo", Args: []string{"c"}},
},
wantMessage: "program checks: 3/3 passed",
},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
g, err := NewProgramGrader("test", &ProgramConfig{
Checks: tt.checks,
})
if err != nil {
t.Fatalf("NewProgramGrader: %v", err)
}

input := GraderInput{
WorkspacePath: t.TempDir(),
Config:        GraderConfig{Weight: 1.0},
}
result, err := g.Grade(context.Background(), input)
if err != nil {
t.Fatalf("Grade: %v", err)
}

if result.Message != tt.wantMessage {
t.Errorf("Message = %q, want %q", result.Message, tt.wantMessage)
}
})
}
}
