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
Command: "echo",
Args:    []string{"hello", "world"},
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
t.Fatal("expected ProgramDetails to be set")
}
if result.Extras.Program.ExitCode != 0 {
t.Errorf("expected ExitCode=0, got %d", result.Extras.Program.ExitCode)
}
if result.Extras.Program.Stdout != "hello world\n" {
t.Errorf("expected stdout 'hello world\\n', got %q", result.Extras.Program.Stdout)
}
}

func TestProgramGrader_CommandFails(t *testing.T) {
g, err := NewProgramGrader("false-test", &ProgramConfig{
Command: "false",
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
t.Error("expected Pass=false")
}
if result.Score != 0.0 {
t.Errorf("expected Score=0.0, got %f", result.Score)
}
if result.Extras.Program == nil {
t.Fatal("expected ProgramDetails to be set")
}
if result.Extras.Program.ExitCode == 0 {
t.Error("expected non-zero exit code")
}
}

func TestProgramGrader_Timeout(t *testing.T) {
if runtime.GOOS == "windows" {
t.Skip("sleep command differs on windows")
}

g, err := NewProgramGrader("timeout-test", &ProgramConfig{
Command: "sleep",
Args:    []string{"10"},
Timeout: 1,
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
t.Error("expected Pass=false on timeout")
}
if result.Score != 0.0 {
t.Errorf("expected Score=0.0, got %f", result.Score)
}
if elapsed > 5*time.Second {
t.Errorf("timeout not enforced: took %s", elapsed)
}
if result.Extras.Program == nil {
t.Fatal("expected ProgramDetails to be set")
}
if result.Extras.Program.ExitCode != -1 {
t.Errorf("expected ExitCode=-1 for timeout, got %d", result.Extras.Program.ExitCode)
}
}

func TestProgramGrader_CapturesStdoutStderr(t *testing.T) {
g, err := NewProgramGrader("capture-test", &ProgramConfig{
Command: "sh",
Args:    []string{"-c", "echo out-content && echo err-content >&2"},
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

if result.Extras.Program == nil {
t.Fatal("expected ProgramDetails to be set")
}
if result.Extras.Program.Stdout != "out-content\n" {
t.Errorf("stdout: expected 'out-content\\n', got %q", result.Extras.Program.Stdout)
}
if result.Extras.Program.Stderr != "err-content\n" {
t.Errorf("stderr: expected 'err-content\\n', got %q", result.Extras.Program.Stderr)
}
}

func TestProgramGrader_WorkingDirectory(t *testing.T) {
dir := t.TempDir()
marker := "workspace-marker.txt"
if err := os.WriteFile(filepath.Join(dir, marker), []byte("ok"), 0644); err != nil {
t.Fatal(err)
}

g, err := NewProgramGrader("wd-test", &ProgramConfig{
Command: "cat",
Args:    []string{marker},
})
if err != nil {
t.Fatalf("NewProgramGrader: %v", err)
}

input := GraderInput{
WorkspacePath: dir,
Config:        GraderConfig{Weight: 1.0},
}
result, err := g.Grade(context.Background(), input)
if err != nil {
t.Fatalf("Grade: %v", err)
}

if !result.Pass {
t.Error("expected Pass=true when file exists in workspace")
}
if result.Extras.Program == nil {
t.Fatal("expected ProgramDetails to be set")
}
if result.Extras.Program.Stdout != "ok" {
t.Errorf("expected stdout 'ok', got %q", result.Extras.Program.Stdout)
}
}

func TestProgramGrader_ContextCancellation(t *testing.T) {
if runtime.GOOS == "windows" {
t.Skip("sleep command differs on windows")
}

g, err := NewProgramGrader("cancel-test", &ProgramConfig{
Command: "sleep",
Args:    []string{"30"},
Timeout: 60,
})
if err != nil {
t.Fatalf("NewProgramGrader: %v", err)
}

ctx, cancel := context.WithCancel(context.Background())
go func() {
time.Sleep(500 * time.Millisecond)
cancel()
}()

input := GraderInput{
WorkspacePath: t.TempDir(),
Config:        GraderConfig{Weight: 1.0},
}
start := time.Now()
result, err := g.Grade(ctx, input)
elapsed := time.Since(start)
if err != nil {
t.Fatalf("Grade: %v", err)
}

if result.Pass {
t.Error("expected Pass=false on cancellation")
}
if elapsed > 5*time.Second {
t.Errorf("cancellation not respected: took %s", elapsed)
}
}

func TestNewProgramGrader_Validation(t *testing.T) {
tests := []struct {
name    string
grName  string
cfg     *ProgramConfig
wantErr bool
}{
{
name:    "missing command",
grName:  "no-cmd",
cfg:     &ProgramConfig{},
wantErr: true,
},
{
name:    "empty name",
grName:  "",
cfg:     &ProgramConfig{Command: "echo"},
wantErr: true,
},
{
name:    "valid minimal",
grName:  "ok",
cfg:     &ProgramConfig{Command: "echo"},
wantErr: false,
},
{
name:    "valid full",
grName:  "ok-full",
cfg:     &ProgramConfig{Command: "echo", Args: []string{"hi"}, Timeout: 10},
wantErr: false,
},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
_, err := NewProgramGrader(tt.grName, tt.cfg)
if (err != nil) != tt.wantErr {
t.Errorf("NewProgramGrader() error = %v, wantErr %v", err, tt.wantErr)
}
})
}
}

func TestProgramGrader_CommandNotFound(t *testing.T) {
g, err := NewProgramGrader("not-found", &ProgramConfig{
Command: "this-command-definitely-does-not-exist-xyz123",
})
if err != nil {
t.Fatal(err)
}
input := GraderInput{
WorkspacePath: t.TempDir(),
Config:        GraderConfig{Weight: 1.0},
}
_, err = g.Grade(context.Background(), input)
if err == nil {
t.Error("expected error for command not found")
}
}

func TestProgramGrader_ImplementsGraderInterface(t *testing.T) {
g, err := NewProgramGrader("iface-test", &ProgramConfig{Command: "echo"})
if err != nil {
t.Fatal(err)
}
var _ Grader = g // compile-time interface check
}

func TestProgramGrader_DefaultTimeout(t *testing.T) {
g, err := NewProgramGrader("default-timeout", &ProgramConfig{Command: "echo"})
if err != nil {
t.Fatal(err)
}
if g.timeout != defaultProgramTimeout {
t.Errorf("expected default timeout %s, got %s", defaultProgramTimeout, g.timeout)
}
}

func TestProgramGrader_FailExitCode(t *testing.T) {
g, err := NewProgramGrader("exit-code-test", &ProgramConfig{
Command: "sh",
Args:    []string{"-c", "exit 42"},
})
if err != nil {
t.Fatal(err)
}

input := GraderInput{
WorkspacePath: t.TempDir(),
Config:        GraderConfig{Weight: 1.0},
}
result, err := g.Grade(context.Background(), input)
if err != nil {
t.Fatalf("Grade: %v", err)
}

if result.Extras.Program == nil {
t.Fatal("expected ProgramDetails to be set")
}
if result.Extras.Program.ExitCode != 42 {
t.Errorf("expected exit code 42, got %d", result.Extras.Program.ExitCode)
}
}
