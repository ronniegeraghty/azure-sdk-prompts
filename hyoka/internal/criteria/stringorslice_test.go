package criteria

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestStringOrSlice_UnmarshalYAML(t *testing.T) {
	tests := []struct {
		name     string
		yaml     string
		want     StringOrSlice
		wantErr  bool
		errMatch string
	}{
		{
			name: "scalar string",
			yaml: `value: python`,
			want: StringOrSlice{"python"},
		},
		{
			name: "empty string",
			yaml: `value: ""`,
			want: nil,
		},
		{
			name: "flow list",
			yaml: `value: [python, java]`,
			want: StringOrSlice{"python", "java"},
		},
		{
			name: "block list",
			yaml: "value:\n  - python\n  - java\n  - go",
			want: StringOrSlice{"python", "java", "go"},
		},
		{
			name: "empty list",
			yaml: `value: []`,
			want: StringOrSlice{},
		},
		{
			name: "single-element list",
			yaml: `value: [python]`,
			want: StringOrSlice{"python"},
		},
		{
			name:     "number scalar",
			yaml:     `value: 42`,
			wantErr:  true,
			errMatch: "must be a string or a list of strings",
		},
		{
			name:     "mapping",
			yaml:     `value: {key: val}`,
			wantErr:  true,
			errMatch: "must be a string or a list of strings",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got struct {
				Value StringOrSlice `yaml:"value"`
			}
			err := yaml.Unmarshal([]byte(tt.yaml), &got)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tt.errMatch)
				}
				if tt.errMatch != "" && !contains(err.Error(), tt.errMatch) {
					t.Fatalf("expected error containing %q, got %q", tt.errMatch, err.Error())
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !equalSlices(got.Value, tt.want) {
				t.Errorf("got %v, want %v", got.Value, tt.want)
			}
		})
	}
}

func TestStringOrSlice_Matches(t *testing.T) {
	tests := []struct {
		name      string
		slice     StringOrSlice
		candidate string
		want      bool
	}{
		{
			name:      "nil matches everything",
			slice:     nil,
			candidate: "anything",
			want:      true,
		},
		{
			name:      "empty matches everything",
			slice:     StringOrSlice{},
			candidate: "anything",
			want:      true,
		},
		{
			name:      "single exact match",
			slice:     StringOrSlice{"python"},
			candidate: "python",
			want:      true,
		},
		{
			name:      "single case-insensitive match",
			slice:     StringOrSlice{"Python"},
			candidate: "python",
			want:      true,
		},
		{
			name:      "single no match",
			slice:     StringOrSlice{"python"},
			candidate: "java",
			want:      false,
		},
		{
			name:      "multi any-of first",
			slice:     StringOrSlice{"python", "java", "go"},
			candidate: "python",
			want:      true,
		},
		{
			name:      "multi any-of middle",
			slice:     StringOrSlice{"python", "java", "go"},
			candidate: "java",
			want:      true,
		},
		{
			name:      "multi any-of last",
			slice:     StringOrSlice{"python", "java", "go"},
			candidate: "go",
			want:      true,
		},
		{
			name:      "multi no match",
			slice:     StringOrSlice{"python", "java", "go"},
			candidate: "rust",
			want:      false,
		},
		{
			name:      "multi case-insensitive",
			slice:     StringOrSlice{"PYTHON", "Java", "GO"},
			candidate: "python",
			want:      true,
		},
		{
			name:      "empty candidate matches only if slice is empty/nil",
			slice:     StringOrSlice{"python"},
			candidate: "",
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.slice.Matches(tt.candidate)
			if got != tt.want {
				t.Errorf("Matches(%q) = %v, want %v", tt.candidate, got, tt.want)
			}
		})
	}
}

func equalSlices(a, b StringOrSlice) bool {
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

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && (s[:len(substr)] == substr || s[len(s)-len(substr):] == substr || containsMiddle(s, substr)))
}

func containsMiddle(s, substr string) bool {
	for i := 0; i+len(substr) <= len(s); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
