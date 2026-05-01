package graders

import (
"testing"
"gopkg.in/yaml.v3"
)

func TestProgramConfigDecode(t *testing.T) {
yamlData := `
checks:
  - kind: command
    command: echo
    args: [hello, world]
    timeout: 10
`
var node yaml.Node
if err := yaml.Unmarshal([]byte(yamlData), &node); err != nil {
t.Fatalf("yaml unmarshal: %v", err)
}

var pgc ProgramConfig
if err := node.Decode(&pgc); err != nil {
t.Fatalf("decode: %v", err)
}

if len(pgc.Checks) != 1 {
t.Fatalf("expected 1 check, got %d", len(pgc.Checks))
}
check := pgc.Checks[0]
if check.Command != "echo" {
t.Errorf("expected command=echo, got %q", check.Command)
}
if len(check.Args) != 2 || check.Args[0] != "hello" {
t.Errorf("expected args=[hello world], got %v", check.Args)
}
}

func TestValidKinds(t *testing.T) {
want := map[string]bool{
KindProgram:      true,
KindPrompt:       true,
KindPromptReview: true,
KindTool:         true,
KindWorkspace:    true,
KindActivity:     true,
}
if len(validKinds) != len(want) {
t.Fatalf("validKinds has %d entries, want %d", len(validKinds), len(want))
}
for k := range want {
if !validKinds[k] {
t.Errorf("expected %q to be in validKinds", k)
}
}
}
