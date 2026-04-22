package criteria

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria/graders"
)

func makeYAMLNode(t *testing.T, data string) yaml.Node {
	t.Helper()
	var node yaml.Node
	if err := yaml.Unmarshal([]byte(data), &node); err != nil {
		t.Fatalf("failed to parse YAML: %v", err)
	}
	if node.Kind == yaml.DocumentNode && len(node.Content) > 0 {
		return *node.Content[0]
	}
	return node
}

func TestInstantiateGraders(t *testing.T) {
	configs := []graders.GraderConfig{
		{Kind: graders.KindFile, Name: "f1", Config: makeYAMLNode(t, `path: "main.py"`)},
		{Kind: graders.KindBehavior, Name: "b1", Config: makeYAMLNode(t, `required_tools: ["mcp"]`)},
	}
	instances, err := InstantiateGraders(configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(instances) != 2 {
		t.Fatalf("expected 2 instances, got %d", len(instances))
	}
	if instances[0].Kind() != graders.KindFile || instances[1].Kind() != graders.KindBehavior {
		t.Errorf("unexpected kinds: %q, %q", instances[0].Kind(), instances[1].Kind())
	}
}

func TestRunGradersAppliesWeightAndGate(t *testing.T) {
	configs := []graders.GraderConfig{
		{Kind: graders.KindFile, Name: "f1", Config: makeYAMLNode(t, `path: "main.py"`), Weight: 0.5, Gate: true},
	}
	instances, err := InstantiateGraders(configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tmpDir := t.TempDir()
	if err := os.WriteFile(filepath.Join(tmpDir, "main.py"), []byte("pass"), 0644); err != nil {
		t.Fatal(err)
	}
	input := graders.GraderInput{
		WorkspacePath: tmpDir,
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
	if !results[0].Pass {
		t.Error("expected pass")
	}
}

func TestRunGradersWithSessionEvents(t *testing.T) {
	configs := []graders.GraderConfig{
		{Kind: graders.KindBehavior, Name: "b1", Config: makeYAMLNode(t, "required_tools: [\"azure-mcp\"]\nforbidden_tools: [\"rm\"]")},
	}
	instances, err := InstantiateGraders(configs)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	input := graders.GraderInput{
		ActionLog: []graders.ActionEvent{
			{Tool: "azure-mcp", Action: "call"},
			{Tool: "read_file", Action: "call"},
		},
	}
	results := RunGraders(context.Background(), instances, configs, input)
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if !results[0].Pass {
		t.Errorf("expected pass, got: %s", results[0].Message)
	}
}

func TestRunGradersWithHooksInvokesCallbacks(t *testing.T) {
configs := []graders.GraderConfig{
{Kind: graders.KindFile, Name: "f1", Config: makeYAMLNode(t, `path: "main.py"`)},
{Kind: graders.KindFile, Name: "f2", Config: makeYAMLNode(t, `path: "missing.py"`)},
}
instances, err := InstantiateGraders(configs)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}

tmpDir := t.TempDir()
if err := os.WriteFile(filepath.Join(tmpDir, "main.py"), []byte("pass"), 0644); err != nil {
t.Fatal(err)
}

var startNames, completeNames []string
var completePasses []bool
hooks := GraderHooks{
OnStart: func(g graders.Grader) {
startNames = append(startNames, g.Name())
},
OnComplete: func(g graders.Grader, r graders.GraderResult) {
completeNames = append(completeNames, g.Name())
completePasses = append(completePasses, r.Pass)
},
}

results := RunGradersWithHooks(context.Background(), instances, configs, graders.GraderInput{WorkspacePath: tmpDir}, hooks)
if len(results) != 2 {
t.Fatalf("expected 2 results, got %d", len(results))
}
if got, want := startNames, []string{"f1", "f2"}; !equalStrings(got, want) {
t.Errorf("OnStart order = %v, want %v", got, want)
}
if got, want := completeNames, []string{"f1", "f2"}; !equalStrings(got, want) {
t.Errorf("OnComplete order = %v, want %v", got, want)
}
if completePasses[0] != results[0].Pass || completePasses[1] != results[1].Pass {
t.Errorf("OnComplete pass mismatch: hooks=%v results=%v,%v",
completePasses, results[0].Pass, results[1].Pass)
}
}

func TestRunGradersNilHooksStillWorks(t *testing.T) {
configs := []graders.GraderConfig{
{Kind: graders.KindFile, Name: "f1", Config: makeYAMLNode(t, `path: "main.py"`)},
}
instances, err := InstantiateGraders(configs)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
tmpDir := t.TempDir()
if err := os.WriteFile(filepath.Join(tmpDir, "main.py"), []byte("pass"), 0644); err != nil {
t.Fatal(err)
}
// Zero-value hooks — both fields nil. Must not panic.
results := RunGradersWithHooks(context.Background(), instances, configs, graders.GraderInput{WorkspacePath: tmpDir}, GraderHooks{})
if len(results) != 1 {
t.Fatalf("expected 1 result, got %d", len(results))
}
}

func equalStrings(a, b []string) bool {
if len(a) != len(b) {
return false
}
for i := range a {
if a[i] != b[i] {
return false
}
}
return true
}
