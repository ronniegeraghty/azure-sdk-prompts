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

func TestWhenMapMatches(t *testing.T) {
	cases := []struct {
		name  string
		when  WhenMap
		props map[string]string
		want  bool
	}{
		{"empty when matches anything", WhenMap{}, map[string]string{"language": "python"}, true},
		{"single key match", WhenMap{"language": "python"}, map[string]string{"language": "python"}, true},
		{"single key mismatch", WhenMap{"language": "python"}, map[string]string{"language": "go"}, false},
		{"missing key", WhenMap{"language": "python"}, map[string]string{}, false},
		{"AND across keys all match", WhenMap{"language": "python", "plane": "data-plane"}, map[string]string{"language": "python", "plane": "data-plane"}, true},
		{"AND across keys one mismatch", WhenMap{"language": "python", "plane": "data-plane"}, map[string]string{"language": "python", "plane": "management-plane"}, false},
		{"case-insensitive value", WhenMap{"language": "Python"}, map[string]string{"language": "python"}, true},

		// Prefixed keys (config-aware when, Phase 1)
		{"mcp_server prefix match", WhenMap{"mcp_server:azure": "true"}, map[string]string{"mcp_server:azure": "true"}, true},
		{"mcp_server prefix not loaded", WhenMap{"mcp_server:azure": "true"}, map[string]string{"mcp_server:other": "true"}, false},
		{"skill prefix match", WhenMap{"skill:reviewer-skills": "true"}, map[string]string{"skill:reviewer-skills": "true"}, true},
		{"plugin prefix match", WhenMap{"plugin:foo": "true"}, map[string]string{"plugin:foo": "true"}, true},
		{"generator key match", WhenMap{"generator": "claude-opus-4.6"}, map[string]string{"generator": "claude-opus-4.6"}, true},
		{"generator key mismatch", WhenMap{"generator": "claude-opus-4.6"}, map[string]string{"generator": "gpt-5.4"}, false},
		{"prompt key + prefixed key together", WhenMap{"language": "python", "mcp_server:azure": "true"}, map[string]string{"language": "python", "mcp_server:azure": "true"}, true},
		{"prompt match but prefixed missing", WhenMap{"language": "python", "mcp_server:azure": "true"}, map[string]string{"language": "python"}, false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.when.Matches(tc.props); got != tc.want {
				t.Errorf("Matches(%v, %v) = %v, want %v", tc.when, tc.props, got, tc.want)
			}
		})
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
