package prompt

import "testing"

func TestParsePromptFile_GroupField(t *testing.T) {
	tests := []struct {
		name      string
		content   string
		wantGroup string
	}{
		{
			name: "group present",
			content: `---
id: test-dp-python-sample
group: crud-operations
properties:
  service: test
  language: python
  plane: data-plane
---

## Prompt

Write code.
`,
			wantGroup: "crud-operations",
		},
		{
			name: "group absent defaults empty",
			content: `---
id: test-dp-python-sample
properties:
  service: test
  language: python
  plane: data-plane
---

## Prompt

Write code.
`,
			wantGroup: "",
		},
		{
			name: "group whitespace trimmed",
			content: `---
id: test-dp-python-sample
group: "  auth-flows  "
properties:
  service: test
  language: python
  plane: data-plane
---

## Prompt

Write code.
`,
			wantGroup: "auth-flows",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := ParsePromptFile([]byte(tt.content), "test.prompt.md")
			if err != nil {
				t.Fatalf("ParsePromptFile: %v", err)
			}
			if p.Group != tt.wantGroup {
				t.Errorf("Group = %q, want %q", p.Group, tt.wantGroup)
			}
		})
	}
}

func TestParsePromptYAML_GroupField(t *testing.T) {
	content := []byte(`id: test-dp-python-sample
group: pagination
prompt_text: "do the thing"
properties:
  service: test
  language: python
  plane: data-plane
`)
	p, err := ParsePromptYAML(content, "test.prompt.yaml")
	if err != nil {
		t.Fatalf("ParsePromptYAML: %v", err)
	}
	if p.Group != "pagination" {
		t.Errorf("Group = %q, want %q", p.Group, "pagination")
	}
}
