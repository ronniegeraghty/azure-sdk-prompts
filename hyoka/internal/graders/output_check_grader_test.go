package graders

import (
	"context"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestOutputCheckGrader_Grade(t *testing.T) {
	tests := []struct {
		name     string
		cfg      *OutputCheckConfig
		files    []FileEntry
		wantPass bool
	}{
		{
			name:     "default config passes with one non-empty file",
			cfg:      &OutputCheckConfig{},
			files:    []FileEntry{{Path: "main.py", Size: 42}},
			wantPass: true,
		},
		{
			name:     "default config fails with zero files",
			cfg:      &OutputCheckConfig{},
			files:    nil,
			wantPass: false,
		},
		{
			name:     "default config fails when only empty files exist",
			cfg:      &OutputCheckConfig{},
			files:    []FileEntry{{Path: "main.py", Size: 0}, {Path: "requirements.txt", Size: 0}},
			wantPass: false,
		},
		{
			name:     "min_files=3 passes when three qualifying files present",
			cfg:      &OutputCheckConfig{MinFiles: 3},
			files:    []FileEntry{{Path: "a", Size: 1}, {Path: "b", Size: 1}, {Path: "c", Size: 1}},
			wantPass: true,
		},
		{
			name:     "min_files=3 fails with two qualifying files",
			cfg:      &OutputCheckConfig{MinFiles: 3},
			files:    []FileEntry{{Path: "a", Size: 1}, {Path: "b", Size: 1}},
			wantPass: false,
		},
		{
			name:     "min_bytes_per_file filters out tiny files",
			cfg:      &OutputCheckConfig{MinFiles: 1, MinBytesPerFile: 100},
			files:    []FileEntry{{Path: "tiny", Size: 10}, {Path: "small", Size: 50}},
			wantPass: false,
		},
		{
			name:     "min_total_bytes enforced",
			cfg:      &OutputCheckConfig{MinFiles: 1, MinTotalBytes: 1000},
			files:    []FileEntry{{Path: "a", Size: 100}, {Path: "b", Size: 200}},
			wantPass: false,
		},
		{
			name:     "min_total_bytes met across multiple files",
			cfg:      &OutputCheckConfig{MinFiles: 1, MinTotalBytes: 250},
			files:    []FileEntry{{Path: "a", Size: 100}, {Path: "b", Size: 200}},
			wantPass: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := NewOutputCheckGrader(tt.name, tt.cfg)
			if err != nil {
				t.Fatalf("NewOutputCheckGrader: %v", err)
			}
			res, err := g.Grade(context.Background(), GraderInput{Files: tt.files})
			if err != nil {
				t.Fatalf("Grade: %v", err)
			}
			if res.Pass != tt.wantPass {
				t.Errorf("Pass=%v want %v (msg=%q)", res.Pass, tt.wantPass, res.Message)
			}
			if res.Kind != KindOutputCheck {
				t.Errorf("Kind=%q want %q", res.Kind, KindOutputCheck)
			}
		})
	}
}

func TestOutputCheckGrader_RegistryAndDecode(t *testing.T) {
	yamlSrc := `
graders:
  - name: produced-files
    kind: output_check
    config:
      min_files: 2
      min_bytes_per_file: 5
`
	gcf, err := Parse([]byte(yamlSrc))
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	if len(gcf.Graders) != 1 {
		t.Fatalf("expected 1 grader, got %d", len(gcf.Graders))
	}
	g, err := NewGrader(gcf.Graders[0])
	if err != nil {
		t.Fatalf("NewGrader: %v", err)
	}
	if g.Kind() != KindOutputCheck {
		t.Errorf("Kind=%q want %q", g.Kind(), KindOutputCheck)
	}

	// Sanity: bad min_total_bytes should fail validation.
	var n yaml.Node
	if err := yaml.Unmarshal([]byte("min_total_bytes: -1"), &n); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	bad := GraderConfig{Kind: KindOutputCheck, Name: "bad", Config: n}
	if _, err := NewGrader(bad); err == nil {
		t.Errorf("expected error for negative min_total_bytes")
	}
}
