package graders

import (
	"context"
	"testing"

	"github.com/ronniegeraghty/hyoka/internal/report"
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
		Config: makeYAMLNode(t, "required: [\"a\"]\nmin_calls: 1"),
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

func TestInstantiateGraders(t *testing.T) {
	configs := []GraderConfig{
		{Kind: KindFile, Name: "f1", Config: makeYAMLNode(t, `path: "main.py"`)},
		{Kind: KindBehavior, Name: "b1", Config: makeYAMLNode(t, `required_tools: ["mcp"]`)},
	}
	instances, err := InstantiateGraders(configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(instances) != 2 {
		t.Fatalf("expected 2 instances, got %d", len(instances))
	}
	if instances[0].Kind() != KindFile || instances[1].Kind() != KindBehavior {
		t.Errorf("unexpected kinds: %q, %q", instances[0].Kind(), instances[1].Kind())
	}
}

func TestRunGradersAppliesWeightAndGate(t *testing.T) {
	configs := []GraderConfig{
		{Kind: KindFile, Name: "f1", Config: makeYAMLNode(t, `path: "main.py"`), Weight: 0.5, Gate: true},
	}
	instances, err := InstantiateGraders(configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	input := &GraderInput{
		GeneratedFiles: []string{"main.py"},
	}
	results := RunGraders(context.Background(), instances, configs, input)
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if results[0].Weight != 0.5 {
		t.Errorf("expected weight 0.5, got %f", results[0].Weight)
	}
	if !results[0].Gate {
		t.Error("expected gate=true")
	}
	if !results[0].Passed {
		t.Error("expected pass")
	}
}

func TestRunGradersWithSessionEvents(t *testing.T) {
	configs := []GraderConfig{
		{Kind: KindBehavior, Name: "b1", Config: makeYAMLNode(t, "required_tools: [\"azure-mcp\"]\nforbidden_tools: [\"rm\"]")},
	}
	instances, err := InstantiateGraders(configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	input := &GraderInput{
		ToolCalls: []string{"azure-mcp", "read_file"},
		SessionEvents: []report.SessionEventRecord{
			{Type: "assistant.turn_start"},
		},
	}
	results := RunGraders(context.Background(), instances, configs, input)
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if !results[0].Passed {
		t.Errorf("expected pass, got: %s", results[0].Message)
	}
}
