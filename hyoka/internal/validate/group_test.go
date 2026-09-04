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
		// Boundary conditions (#608 polish):
		{strings.Repeat("a", 63), true},                    // one under max
		{"ab" + strings.Repeat("-c", 31), true},            // 64-char hyphenated
		{"ab" + strings.Repeat("-c", 31) + "d", false},     // 65-char hyphenated
		{" ", false},                                       // whitespace only
		{"   ", false},                                     // multi whitespace
		{"\t", false},                                      // tab
		{"\n", false},                                      // newline
		{"a\nb", false},                                    // embedded newline
		{"leading space", false},
		{" leading-space", false},                          // leading actual space
		{"trailing-space ", false},
		{"-", false},                                       // hyphen only
		{"--", false},                                      // only hyphens
		{"a-", false},                                      // trailing hyphen after letter
		{"-a", false},                                      // leading hyphen before letter
		{"a--b", false},                                    // consecutive hyphens mid
		{"a---b", false},                                   // triple consecutive hyphens
		{"a-b-c-d-e-f", true},                              // many single hyphens
		{"a1", true},                                       // letter + digit
		{"a-1", true},                                      // letter-digit segment
		{"a-1b", true},                                     // mixed alphanumeric segment
		{"1", false},                                       // single digit
		{"1a", false},                                      // starts with digit
		{"0-abc", false},                                   // starts with digit + hyphen
		{"abc.def", false},                                 // dot separator
		{"abc_def", false},                                 // underscore
		{"abc+def", false},                                 // plus
		{"abc def", false},                                 // internal space
		{"abc@def", false},                                 // at sign
		{"abc#def", false},                                 // hash
		{"abc!def", false},                                 // bang
		{"ABC", false},                                     // all uppercase
		{"aBc", false},                                     // mixed case
		{"ñoño", false},                                    // non-ASCII lowercase
		{"café", false},                                    // accented lowercase
		{"emoji🙂", false},                                  // emoji
		{"null\x00byte", false},                            // embedded null
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
