package progress

import (
	"testing"
)

func TestVisibleWidth(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "plain text",
			input: "hello",
			want:  5,
		},
		{
			name:  "text with ANSI color codes",
			input: "\x1b[31mred\x1b[0m",
			want:  3,
		},
		{
			name:  "text with multiple ANSI sequences",
			input: "\x1b[1m\x1b[31mbold red\x1b[0m",
			want:  8,
		},
		{
			name:  "empty string",
			input: "",
			want:  0,
		},
		{
			name:  "only ANSI codes",
			input: "\x1b[31m\x1b[0m",
			want:  0,
		},
		{
			name:  "emoji",
			input: "🔄",
			want:  1,
		},
		{
			name:  "mix of text and emoji with ANSI",
			input: "\x1b[32m✅ Loaded\x1b[0m",
			want:  8, // ✅ (1) + space (1) + Loaded (6) = 8 runes
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := visibleWidth(tt.input)
			if got != tt.want {
				t.Errorf("visibleWidth(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestTruncateToWidth(t *testing.T) {
	tests := []struct {
		name  string
		input string
		max   int
		want  string
	}{
		{
			name:  "no truncation needed",
			input: "hello",
			max:   10,
			want:  "hello",
		},
		{
			name:  "truncate plain text",
			input: "hello world",
			max:   5,
			want:  "hell…\x1b[0m",
		},
		{
			name:  "truncate with ANSI codes preserved",
			input: "\x1b[31mhello world\x1b[0m",
			max:   5,
			want:  "\x1b[31mhell…\x1b[0m",
		},
		{
			name:  "zero max",
			input: "hello",
			max:   0,
			want:  "hello",
		},
		{
			name:  "negative max",
			input: "hello",
			max:   -1,
			want:  "hello",
		},
		{
			name:  "exact fit",
			input: "hello",
			max:   5,
			want:  "hello",
		},
		{
			name:  "truncate to 1 char",
			input: "hello",
			max:   1,
			want:  "…\x1b[0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := truncateToWidth(tt.input, tt.max)
			if got != tt.want {
				t.Errorf("truncateToWidth(%q, %d) = %q, want %q", tt.input, tt.max, got, tt.want)
			}
		})
	}
}
