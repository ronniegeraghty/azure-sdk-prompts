package graders

import (
"context"
"testing"

"gopkg.in/yaml.v3"
)

func TestNewGraderProgramKind(t *testing.T) {
yamlData := `
checks:
  - kind: command
    command: echo
    args: [test]
`
var node yaml.Node
if err := yaml.Unmarshal([]byte(yamlData), &node); err != nil {
t.Fatalf("yaml: %v", err)
}

gc := GraderConfig{
Kind:   KindProgram,
Name:   "test_prog",
Config: node,
Weight: 1.0,
}

g, err := NewGrader(gc)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}

if g.Kind() != KindProgram {
t.Errorf("expected kind %q, got %q", KindProgram, g.Kind())
}
if g.Name() != "test_prog" {
t.Errorf("expected name test_prog, got %q", g.Name())
}
}

func TestNewGraderToolKind(t *testing.T) {
yamlData := `
checks:
  - kind: tool_used
    tool: bash
`
var node yaml.Node
if err := yaml.Unmarshal([]byte(yamlData), &node); err != nil {
t.Fatalf("yaml: %v", err)
}

gc := GraderConfig{
Kind:   KindTool,
Name:   "test_tool",
Config: node,
Weight: 1.0,
}

g, err := NewGrader(gc)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}

if g.Kind() != KindTool {
t.Errorf("expected kind %q, got %q", KindTool, g.Kind())
}
}

func TestNewGraderWorkspaceKind(t *testing.T) {
yamlData := `
checks:
  - kind: file
    name: test.txt
    state: present
`
var node yaml.Node
if err := yaml.Unmarshal([]byte(yamlData), &node); err != nil {
t.Fatalf("yaml: %v", err)
}

gc := GraderConfig{
Kind:   KindWorkspace,
Name:   "test_ws",
Config: node,
Weight: 1.0,
}

g, err := NewGrader(gc)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}

if g.Kind() != KindWorkspace {
t.Errorf("expected kind %q, got %q", KindWorkspace, g.Kind())
}
}

func TestNewGraderUnknownKind(t *testing.T) {
gc := GraderConfig{
Kind:   "unknown",
Name:   "test",
Weight: 1.0,
}

_, err := NewGrader(gc)
if err == nil {
t.Error("expected error for unknown kind, got nil")
}
}

func TestGraderInterfaceCompliance(t *testing.T) {
yamlData := `
checks:
  - kind: command
    command: echo
    args: [test]
`
var node yaml.Node
if err := yaml.Unmarshal([]byte(yamlData), &node); err != nil {
t.Fatalf("yaml: %v", err)
}

gc := GraderConfig{
Kind:   KindProgram,
Name:   "test",
Config: node,
Weight: 1.0,
}

g, err := NewGrader(gc)
if err != nil {
t.Fatalf("NewGrader: %v", err)
}

// Call Grade to verify it implements the interface
_, err = g.Grade(context.Background(), GraderInput{WorkspacePath: "/tmp"})
// We don't care if it succeeds, just that it compiles and runs
_ = err
}
