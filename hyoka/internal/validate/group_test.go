package validate

import (
	"strings"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/prompt"
)

func TestIsValidGroupName(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"crud-operations", true},
		{"auth", true},
		{"a", true},
		{"auth-flows-v2", true},
		{"abc123-def", true},
		{"", false},
		{"-leading", false},
		{"trailing-", false},
		{"double--dash", false},
		{"UPPER", false},
		{"camelCase", false},
		{"snake_case", false},
		{"1starts-with-digit", false},
		{"has space", false},
		{"slash/group", false},
		{strings.Repeat("a", 65), false},
		{strings.Repeat("a", 64), true},
	}
	for _, c := range cases {
		if got := IsValidGroupName(c.in); got != c.want {
			t.Errorf("IsValidGroupName(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}

func TestValidatePromptStruct_Group(t *testing.T) {
	base := func(group string) *prompt.Prompt {
		return &prompt.Prompt{
			ID:         "test-dp-python-sample",
			PromptText: "do something",
			Group:      group,
			Properties: map[string]string{
				"service":  "test-service",
				"language": "python",
				"plane":    "data-plane",
			},
		}
	}

	if err := ValidatePromptStruct(base("")); err != nil {
		t.Errorf("empty group should be valid (ungrouped): %v", err)
	}
	if err := ValidatePromptStruct(base("crud-operations")); err != nil {
		t.Errorf("valid group rejected: %v", err)
	}
	err := ValidatePromptStruct(base("Bad Group!"))
	if err == nil {
		t.Fatalf("invalid group should fail validation")
	}
	if !strings.Contains(err.Error(), "group") {
		t.Errorf("error should mention group: %v", err)
	}
}
