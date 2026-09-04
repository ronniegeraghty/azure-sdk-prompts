package criteria

import (
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestMatchSet_UnmarshalYAML(t *testing.T) {
	tests := []struct {
		name    string
		yaml    string
		want    MatchSet
		wantErr string
	}{
		{
			name: "scalar string",
			yaml: "language: python",
			want: MatchSet{Is: StringOrSlice{"python"}},
		},
		{
			name: "sequence",
			yaml: "language: [python, java]",
			want: MatchSet{Is: StringOrSlice{"python", "java"}},
		},
		{
			name: "map with is only",
			yaml: "language:\n  is: python",
			want: MatchSet{Is: StringOrSlice{"python"}},
		},
		{
			name: "map with is list",
			yaml: "language:\n  is: [python, java]",
			want: MatchSet{Is: StringOrSlice{"python", "java"}},
		},
		{
			name: "map with not only",
			yaml: "language:\n  not: python",
			want: MatchSet{Not: StringOrSlice{"python"}},
		},
		{
			name: "map with not list",
			yaml: "language:\n  not: [python, java]",
			want: MatchSet{Not: StringOrSlice{"python", "java"}},
		},
		{
			name: "map with both is and not",
			yaml: "language:\n  is: [python, java]\n  not: rust",
			want: MatchSet{
				Is:  StringOrSlice{"python", "java"},
				Not: StringOrSlice{"rust"},
			},
		},
		{
			name: "map empty",
			yaml: "language: {}",
			want: MatchSet{},
		},
		{
			name: "null",
			yaml: "language: null",
			want: MatchSet{},
		},
		{
			name:    "map with unknown key",
			yaml:    "language:\n  is: python\n  maybe: java",
			wantErr: "unknown key \"maybe\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			type wrapper struct {
				Language MatchSet `yaml:"language"`
			}
			var w wrapper
			err := yaml.Unmarshal([]byte(tt.yaml), &w)
			if tt.wantErr != "" {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tt.wantErr)
				}
				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("error %q does not contain %q", err.Error(), tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !matchSetEqual(w.Language, tt.want) {
				t.Errorf("got %+v, want %+v", w.Language, tt.want)
			}
		})
	}
}

func TestMatchSet_Matches(t *testing.T) {
	tests := []struct {
		name      string
		ms        MatchSet
		candidate string
		want      bool
	}{
		{
			name:      "empty matches anything",
			ms:        MatchSet{},
			candidate: "python",
			want:      true,
		},
		{
			name:      "is only - positive match",
			ms:        MatchSet{Is: StringOrSlice{"python"}},
			candidate: "python",
			want:      true,
		},
		{
			name:      "is only - negative match",
			ms:        MatchSet{Is: StringOrSlice{"python"}},
			candidate: "java",
			want:      false,
		},
		{
			name:      "is only - case insensitive",
			ms:        MatchSet{Is: StringOrSlice{"Python"}},
			candidate: "python",
			want:      true,
		},
		{
			name:      "is list - match first",
			ms:        MatchSet{Is: StringOrSlice{"python", "java"}},
			candidate: "python",
			want:      true,
		},
		{
			name:      "is list - match second",
			ms:        MatchSet{Is: StringOrSlice{"python", "java"}},
			candidate: "java",
			want:      true,
		},
		{
			name:      "is list - no match",
			ms:        MatchSet{Is: StringOrSlice{"python", "java"}},
			candidate: "rust",
			want:      false,
		},
		{
			name:      "not only - not in list",
			ms:        MatchSet{Not: StringOrSlice{"python"}},
			candidate: "java",
			want:      true,
		},
		{
			name:      "not only - in list",
			ms:        MatchSet{Not: StringOrSlice{"python"}},
			candidate: "python",
			want:      false,
		},
		{
			name:      "not only - case insensitive",
			ms:        MatchSet{Not: StringOrSlice{"Python"}},
			candidate: "python",
			want:      false,
		},
		{
			name:      "both is and not - passes",
			ms:        MatchSet{Is: StringOrSlice{"python", "java"}, Not: StringOrSlice{"rust"}},
			candidate: "python",
			want:      true,
		},
		{
			name:      "both is and not - in is but also in not",
			ms:        MatchSet{Is: StringOrSlice{"python", "java"}, Not: StringOrSlice{"python"}},
			candidate: "python",
			want:      false,
		},
		{
			name:      "both is and not - not in is",
			ms:        MatchSet{Is: StringOrSlice{"python", "java"}, Not: StringOrSlice{"rust"}},
			candidate: "go",
			want:      false,
		},
		{
			name:      "both is and not - in not",
			ms:        MatchSet{Is: StringOrSlice{"python", "java"}, Not: StringOrSlice{"rust"}},
			candidate: "rust",
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.ms.Matches(tt.candidate)
			if got != tt.want {
				t.Errorf("Matches(%q) = %v, want %v", tt.candidate, got, tt.want)
			}
		})
	}
}

func TestMatchSet_MatchesAny(t *testing.T) {
	tests := []struct {
		name       string
		ms         MatchSet
		candidates []string
		want       bool
	}{
		{
			name:       "empty matches anything",
			ms:         MatchSet{},
			candidates: []string{"auth", "crud"},
			want:       true,
		},
		{
			name:       "empty candidates with empty MatchSet",
			ms:         MatchSet{},
			candidates: []string{},
			want:       true,
		},
		{
			name:       "is only - has intersection",
			ms:         MatchSet{Is: StringOrSlice{"auth"}},
			candidates: []string{"auth", "crud"},
			want:       true,
		},
		{
			name:       "is only - no intersection",
			ms:         MatchSet{Is: StringOrSlice{"pagination"}},
			candidates: []string{"auth", "crud"},
			want:       false,
		},
		{
			name:       "is only - case insensitive",
			ms:         MatchSet{Is: StringOrSlice{"Auth"}},
			candidates: []string{"auth", "crud"},
			want:       true,
		},
		{
			name:       "is list - intersection with first",
			ms:         MatchSet{Is: StringOrSlice{"auth", "pagination"}},
			candidates: []string{"auth", "crud"},
			want:       true,
		},
		{
			name:       "is list - intersection with second",
			ms:         MatchSet{Is: StringOrSlice{"pagination", "crud"}},
			candidates: []string{"auth", "crud"},
			want:       true,
		},
		{
			name:       "is list - no intersection",
			ms:         MatchSet{Is: StringOrSlice{"pagination", "streaming"}},
			candidates: []string{"auth", "crud"},
			want:       false,
		},
		{
			name:       "not only - no intersection (passes)",
			ms:         MatchSet{Not: StringOrSlice{"pagination"}},
			candidates: []string{"auth", "crud"},
			want:       true,
		},
		{
			name:       "not only - has intersection (fails)",
			ms:         MatchSet{Not: StringOrSlice{"auth"}},
			candidates: []string{"auth", "crud"},
			want:       false,
		},
		{
			name:       "not only - case insensitive",
			ms:         MatchSet{Not: StringOrSlice{"Auth"}},
			candidates: []string{"auth", "crud"},
			want:       false,
		},
		{
			name:       "both is and not - passes",
			ms:         MatchSet{Is: StringOrSlice{"auth", "crud"}, Not: StringOrSlice{"pagination"}},
			candidates: []string{"auth", "streaming"},
			want:       true,
		},
		{
			name:       "both is and not - is intersection but also not intersection (fails)",
			ms:         MatchSet{Is: StringOrSlice{"auth", "crud"}, Not: StringOrSlice{"auth"}},
			candidates: []string{"auth", "streaming"},
			want:       false,
		},
		{
			name:       "both is and not - no is intersection",
			ms:         MatchSet{Is: StringOrSlice{"pagination", "crud"}, Not: StringOrSlice{"auth"}},
			candidates: []string{"auth", "streaming"},
			want:       false,
		},
		{
			name:       "both is and not - multiple candidates pass",
			ms:         MatchSet{Is: StringOrSlice{"auth", "crud"}, Not: StringOrSlice{"deprecated"}},
			candidates: []string{"auth", "crud", "pagination"},
			want:       true,
		},
		{
			name:       "both is and not - one in not fails all",
			ms:         MatchSet{Is: StringOrSlice{"auth", "crud"}, Not: StringOrSlice{"deprecated"}},
			candidates: []string{"auth", "crud", "deprecated"},
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.ms.MatchesAny(tt.candidates)
			if got != tt.want {
				t.Errorf("MatchesAny(%v) = %v, want %v", tt.candidates, got, tt.want)
			}
		})
	}
}

func TestMatchSet_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		ms   MatchSet
		want bool
	}{
		{
			name: "empty",
			ms:   MatchSet{},
			want: true,
		},
		{
			name: "is only",
			ms:   MatchSet{Is: StringOrSlice{"python"}},
			want: false,
		},
		{
			name: "not only",
			ms:   MatchSet{Not: StringOrSlice{"python"}},
			want: false,
		},
		{
			name: "both",
			ms:   MatchSet{Is: StringOrSlice{"python"}, Not: StringOrSlice{"java"}},
			want: false,
		},
		{
			name: "empty slices",
			ms:   MatchSet{Is: StringOrSlice{}, Not: StringOrSlice{}},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.ms.IsEmpty()
			if got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func matchSetEqual(a, b MatchSet) bool {
	return stringSliceEqual(a.Is, b.Is) && stringSliceEqual(a.Not, b.Not)
}

func stringSliceEqual(a, b []string) bool {
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
