package graders

import (
"context"
"os"
"path/filepath"
"runtime"
"testing"
"time"
)

func TestProgramGrader_CommandSucceeds(t *testing.T) {
g, err := NewProgramGrader("echo-test", &ProgramConfig{
Checks: []ProgramCheck{
{Kind: "command", Command: "echo", Args: []string{"hello", "world"}},
},
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

if !result.Pass {
t.Error("expected Pass=true")
}
if result.Score != 1.0 {
t.Errorf("expected Score=1.0, got %f", result.Score)
}
if result.Name != "echo-test" {
t.Errorf("expected Name=echo-test, got %s", result.Name)
}
if result.Extras.Program == nil {
t.Fatal("expected ProgramExtras to be set")
}
if len(result.Extras.Program.CheckResults) != 1 {
t.Fatalf("expected 1 check result, got %d", len(result.Extras.Program.CheckResults))
}
checkRes := result.Extras.Program.CheckResults[0]
if checkRes.ExitCode != 0 {
t.Errorf("expected ExitCode=0, got %d", checkRes.ExitCode)
}
if checkRes.Stdout != "hello world\n" {
t.Errorf("expected stdout 'hello world\\n', got %q", checkRes.Stdout)
}
}

func TestProgramGrader_CommandFails(t *testing.T) {
g, err := NewProgramGrader("fail-test", &ProgramConfig{
Checks: []ProgramCheck{
{Kind: "command", Command: "false"},
},
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

if result.Pass {
t.Error("expected Pass=false for failing command")
}
if result.Score != 0.0 {
t.Errorf("expected Score=0.0, got %f", result.Score)
}
if len(result.Extras.Program.CheckResults) != 1 {
t.Fatalf("expected 1 check result, got %d", len(result.Extras.Program.CheckResults))
}
checkRes := result.Extras.Program.CheckResults[0]
if checkRes.ExitCode == 0 {
t.Errorf("expected non-zero exit code, got 0")
}
}

func TestProgramGrader_Timeout(t *testing.T) {
if runtime.GOOS == "windows" {
t.Skip("skipping timeout test on windows")
}
g, err := NewProgramGrader("timeout-test", &ProgramConfig{
Checks: []ProgramCheck{
{Kind: "command", Command: "sleep", Args: []string{"10"}, Timeout: 1},
},
})
if err != nil {
t.Fatalf("NewProgramGrader: %v", err)
}

input := GraderInput{
WorkspacePath: t.TempDir(),
Config:        GraderConfig{Weight: 1.0},
}

start := time.Now()
result, err := g.Grade(context.Background(), input)
elapsed := time.Since(start)

if err != nil {
t.Fatalf("Grade: %v", err)
}
if result.Pass {
t.Error("expected Pass=false for timeout")
}
if elapsed > 3*time.Second {
t.Errorf("expected timeout around 1s, took %v", elapsed)
}
}

func TestProgramGrader_WorkingDirectory(t *testing.T) {
tmp := t.TempDir()
testFile := filepath.Join(tmp, "testfile.txt")
if err := os.WriteFile(testFile, []byte("content"), 0644); err != nil {
t.Fatalf("setup: %v", err)
}

g, err := NewProgramGrader("pwd-test", &ProgramConfig{
Checks: []ProgramCheck{
{Kind: "command", Command: "test", Args: []string{"-f", "testfile.txt"}},
},
})
if err != nil {
t.Fatalf("NewProgramGrader: %v", err)
}

input := GraderInput{
WorkspacePath: tmp,
Config:        GraderConfig{Weight: 1.0},
}
result, err := g.Grade(context.Background(), input)
if err != nil {
t.Fatalf("Grade: %v", err)
}

if !result.Pass {
t.Errorf("expected Pass=true (file exists in workspace), got Pass=false")
}
}

func TestProgramGrader_MultipleChecks(t *testing.T) {
g, err := NewProgramGrader("multi-test", &ProgramConfig{
Checks: []ProgramCheck{
{Kind: "command", Command: "true"},
{Kind: "command", Command: "echo", Args: []string{"test"}},
{Kind: "command", Command: "false"}, // intentional failure
},
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

// 2 pass + 1 fail = 2/3 pass, so overall Pass should be false
if result.Pass {
t.Error("expected Pass=false when any check fails")
}
// Score should be 2/3 ≈ 0.666...
expectedScore := 2.0 / 3.0
if result.Score < expectedScore-0.01 || result.Score > expectedScore+0.01 {
t.Errorf("expected Score≈%.2f, got %.2f", expectedScore, result.Score)
}
if len(result.Checks) != 3 {
t.Fatalf("expected 3 checks, got %d", len(result.Checks))
}
// Check 1 and 2 should pass, check 3 should fail
if !result.Checks[0].Pass {
t.Error("check 1 should pass")
}
if !result.Checks[1].Pass {
t.Error("check 2 should pass")
}
if result.Checks[2].Pass {
t.Error("check 3 should fail")
}
}
