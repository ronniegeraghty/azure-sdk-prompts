package graders

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func makeYAMLNode(t *testing.T, data string) yaml.Node {
	t.Helper()
	var node yaml.Node
	if err := yaml.Unmarshal([]byte(data), &node); err != nil {
		t.Fatalf("failed to parse YAML: %v", err)
	}
	// yaml.Unmarshal wraps in a document node; return the first content node
	if node.Kind == yaml.DocumentNode && len(node.Content) > 0 {
		return *node.Content[0]
	}
	return node
}

func TestNewGraderFileKind(t *testing.T) {
	gc := GraderConfig{
		Kind:   KindFile,
		Name:   "test_file",
		Config: makeYAMLNode(t, `path: "main.py"`),
	}
	g, err := NewGrader(gc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if g.Name() != "test_file" || g.Kind() != KindFile {
		t.Errorf("unexpected name=%q kind=%q", g.Name(), g.Kind())
	}
}

func TestNewGraderProgramKind(t *testing.T) {
	gc := GraderConfig{
		Kind:   KindProgram,
		Name:   "test_prog",
		Config: makeYAMLNode(t, `command: "echo"`),
	}
	g, err := NewGrader(gc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if g.Kind() != KindProgram {
		t.Errorf("expected kind %q, got %q", KindProgram, g.Kind())
	}
}

func TestNewGraderPromptKind(t *testing.T) {
	gc := GraderConfig{
		Kind:   KindPrompt,
		Name:   "test_prompt",
		Config: makeYAMLNode(t, "model: opus\nrubric: test"),
	}
	g, err := NewGrader(gc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if g.Kind() != KindPrompt {
		t.Errorf("expected kind %q, got %q", KindPrompt, g.Kind())
	}
}

func TestNewGraderBehaviorKind(t *testing.T) {
	gc := GraderConfig{
		Kind:   KindBehavior,
		Name:   "test_behavior",
		Config: makeYAMLNode(t, `required_tools: ["azure-mcp"]`),
	}
	g, err := NewGrader(gc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if g.Kind() != KindBehavior {
		t.Errorf("expected kind %q, got %q", KindBehavior, g.Kind())
	}
}

func TestNewGraderActionSequenceKind(t *testing.T) {
	gc := GraderConfig{
		Kind:   KindActionSequence,
		Name:   "test_as",
		Config: makeYAMLNode(t, `expected_actions: ["read_file"]`),
	}
	g, err := NewGrader(gc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if g.Kind() != KindActionSequence {
		t.Errorf("expected kind %q, got %q", KindActionSequence, g.Kind())
	}
}

func TestNewGraderToolConstraintKind(t *testing.T) {
	gc := GraderConfig{
		Kind:   KindToolConstraint,
		Name:   "test_tc",
		Config: makeYAMLNode(t, "required: [\"a\"]\nmin_calls:\n  a: 1"),
	}
	g, err := NewGrader(gc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if g.Kind() != KindToolConstraint {
		t.Errorf("expected kind %q, got %q", KindToolConstraint, g.Kind())
	}
}

func TestNewGraderUnknownKind(t *testing.T) {
	gc := GraderConfig{
		Kind:   "unknown",
		Name:   "bad",
		Config: makeYAMLNode(t, `path: "x"`),
	}
	_, err := NewGrader(gc)
	if err == nil {
		t.Fatal("expected error for unknown kind")
	}
}
