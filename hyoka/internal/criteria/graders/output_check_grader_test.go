package graders

import (
	"context"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/workspace"
)

// delta is a small helper to build a WorkspaceDelta from parallel slices
// so the table-driven tests stay readable.
type dFile struct {
	path     string
	size     int64
	modified bool // true → ModifiedFiles, false → NewFiles
}

func buildDelta(files []dFile) *WorkspaceDelta {
	d := &WorkspaceDelta{}
	for _, f := range files {
		if f.modified {
			d.ModifiedFiles = append(d.ModifiedFiles, workspace.ModifiedFile{
				Path:      f.path,
				SizeAfter: f.size,
			})
			d.ModifiedFileCount++
		} else {
			d.NewFiles = append(d.NewFiles, workspace.NewFile{
				Path: f.path,
				Size: f.size,
			})
			d.NewFileCount++
		}
	}
	return d
}

func findPoint(t *testing.T, res GraderResult, labelSubstr string) GraderPoint {
	t.Helper()
	for _, p := range res.Checks {
		if strings.Contains(p.Label, labelSubstr) {
			return p
		}
	}
	t.Fatalf("point with label containing %q not found in result (have %d points)", labelSubstr, len(res.Checks))
	return GraderCheck{}
}

func TestOutputCheckGrader_NoKnobs_TriviallyPasses(t *testing.T) {
	g, err := NewOutputCheckGrader("empty", &OutputCheckConfig{})
	if err != nil {
		t.Fatalf("NewOutputCheckGrader: %v", err)
	}
	res, err := g.Grade(context.Background(), GraderInput{WorkspaceDelta: &WorkspaceDelta{}})
	if err != nil {
		t.Fatalf("Grade: %v", err)
	}
	if !res.Pass || res.Score != 1.0 {
		t.Errorf("want trivial pass, got Pass=%v Score=%v", res.Pass, res.Score)
	}
	// Phase 3 invariant: every grader emits ≥ 1 Point. The "no knobs" case
	// emits a single trivially-passing "no_knobs" Point.
	if len(res.Checks) != 1 {
		t.Errorf("expected exactly one (trivial-pass) sub-check, got %d", len(res.Checks))
	}
	if len(res.Checks) >= 1 && res.Checks[0].Label != "no_knobs" {
		t.Errorf("expected single Point with label 'no_knobs', got %q", res.Checks[0].Label)
	}
}

func TestOutputCheckGrader_NilDelta_TreatedAsEmpty(t *testing.T) {
	g, _ := NewOutputCheckGrader("nil-delta", &OutputCheckConfig{MinFiles: 1})
	res, _ := g.Grade(context.Background(), GraderInput{WorkspaceDelta: nil})
	if res.Pass {
		t.Errorf("nil delta with min_files=1 should fail; msg=%q", res.Message)
	}
	point := findPoint(t, res, "min_files")
	if point.Pass {
		t.Errorf("min_files point should fail, got pass: %q", point.Message)
	}
}

func TestOutputCheckGrader_KnobsIndividually(t *testing.T) {
	files := []dFile{
		{path: "src/main.py", size: 120, modified: false},
		{path: "README.md", size: 40, modified: false},
		{path: "config.yaml", size: 8, modified: true},
	}
	delta := buildDelta(files)

	tests := []struct {
		name       string
		cfg        OutputCheckConfig
		wantPass   bool
		wantCheck  string
		wantSubOK  bool
		wantMsgSub string // substring that must appear in the sub-check's Message
	}{
		{
			name:       "min_files pass (>= met)",
			cfg:        OutputCheckConfig{MinFiles: 3},
			wantPass:   true,
			wantCheck:  "min_files",
			wantSubOK:  true,
			wantMsgSub: "produced 3 file(s)",
		},
		{
			name:       "min_files fail",
			cfg:        OutputCheckConfig{MinFiles: 5},
			wantPass:   false,
			wantCheck:  "min_files",
			wantSubOK:  false,
			wantMsgSub: "need >= 5",
		},
		{
			name:       "max_files pass (exact boundary)",
			cfg:        OutputCheckConfig{MaxFiles: 3},
			wantPass:   true,
			wantCheck:  "max_files",
			wantSubOK:  true,
			wantMsgSub: "<= 3 max",
		},
		{
			name:       "max_files fail",
			cfg:        OutputCheckConfig{MaxFiles: 2},
			wantPass:   false,
			wantCheck:  "max_files",
			wantSubOK:  false,
			wantMsgSub: "exceeds max of 2",
		},
		{
			name:       "require_files pass",
			cfg:        OutputCheckConfig{RequireFiles: []string{"src/main.py", "README.md"}},
			wantPass:   true,
			wantCheck:  "require_files",
			wantSubOK:  true,
			wantMsgSub: "all 2 required",
		},
		{
			name:       "require_files fail (missing)",
			cfg:        OutputCheckConfig{RequireFiles: []string{"src/main.py", "LICENSE"}},
			wantPass:   false,
			wantCheck:  "require_files",
			wantSubOK:  false,
			wantMsgSub: "LICENSE",
		},
		{
			name:       "forbid_files pass",
			cfg:        OutputCheckConfig{ForbidFiles: []string{"secret.txt"}},
			wantPass:   true,
			wantCheck:  "forbid_files",
			wantSubOK:  true,
			wantMsgSub: "none of 1",
		},
		{
			name:       "forbid_files fail (found)",
			cfg:        OutputCheckConfig{ForbidFiles: []string{"README.md"}},
			wantPass:   false,
			wantCheck:  "forbid_files",
			wantSubOK:  false,
			wantMsgSub: "README.md",
		},
		{
			name:       "require_updated pass",
			cfg:        OutputCheckConfig{RequireUpdated: []string{"config.yaml"}},
			wantPass:   true,
			wantCheck:  "require_updated",
			wantSubOK:  true,
			wantMsgSub: "all 1 path(s) appear in modified",
		},
		{
			name:       "require_updated fail (new-only path)",
			cfg:        OutputCheckConfig{RequireUpdated: []string{"src/main.py"}},
			wantPass:   false,
			wantCheck:  "require_updated",
			wantSubOK:  false,
			wantMsgSub: "src/main.py",
		},
		{
			name:       "min_bytes_per_file pass (all >= threshold)",
			cfg:        OutputCheckConfig{MinBytesPerFile: 8},
			wantPass:   true,
			wantCheck:  "min_bytes_per_file",
			wantSubOK:  true,
			wantMsgSub: ">= 8 byte(s)",
		},
		{
			name:       "min_bytes_per_file fail (one offender)",
			cfg:        OutputCheckConfig{MinBytesPerFile: 50},
			wantPass:   false,
			wantCheck:  "min_bytes_per_file",
			wantSubOK:  false,
			wantMsgSub: "config.yaml",
		},
		{
			name:       "max_bytes_per_file pass",
			cfg:        OutputCheckConfig{MaxBytesPerFile: 1000},
			wantPass:   true,
			wantCheck:  "max_bytes_per_file",
			wantSubOK:  true,
			wantMsgSub: "<= 1000",
		},
		{
			name:       "max_bytes_per_file fail",
			cfg:        OutputCheckConfig{MaxBytesPerFile: 50},
			wantPass:   false,
			wantCheck:  "max_bytes_per_file",
			wantSubOK:  false,
			wantMsgSub: "src/main.py",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := NewOutputCheckGrader(tt.name, &tt.cfg)
			if err != nil {
				t.Fatalf("NewOutputCheckGrader: %v", err)
			}
			res, err := g.Grade(context.Background(), GraderInput{WorkspaceDelta: delta})
			if err != nil {
				t.Fatalf("Grade: %v", err)
			}
			if res.Pass != tt.wantPass {
				t.Errorf("overall Pass=%v want %v (msg=%q)", res.Pass, tt.wantPass, res.Message)
			}
			if tt.wantPass && res.Score != 1.0 {
				t.Errorf("Pass=true but Score=%v want 1.0", res.Score)
			}
			if !tt.wantPass && res.Score != 0 {
				t.Errorf("Pass=false but Score=%v want 0", res.Score)
			}
			point := findPoint(t, res, tt.wantCheck)
			if point.Pass != tt.wantSubOK {
				t.Errorf("sub-check %q Pass=%v want %v (msg=%q)",
					tt.wantCheck, point.Pass, tt.wantSubOK, point.Message)
			}
			if !strings.Contains(point.Message, tt.wantMsgSub) {
				t.Errorf("sub-check %q message %q does not contain %q",
					tt.wantCheck, point.Message, tt.wantMsgSub)
			}
			if res.Kind != KindOutputCheck {
				t.Errorf("Kind=%q want %q", res.Kind, KindOutputCheck)
			}
		})
	}
}

func TestOutputCheckGrader_AllChecksReported_NoEarlyExit(t *testing.T) {
	// Config that fails some knobs and passes others — verify every
	// configured knob gets a sub-check entry regardless of earlier failures.
	cfg := OutputCheckConfig{
		MinFiles:        10,                    // fails (we have 3)
		MaxFiles:        2,                     // (would fail too but validator rejects min>max)
		RequireFiles:    []string{"src/main.py"},
		ForbidFiles:     []string{"README.md"}, // fails
		MinBytesPerFile: 5,                     // passes (all files >= 5)
	}
	// Drop the conflicting min/max_files pair — keep min_files (failing)
	// and use an independent fail knob.
	cfg.MaxFiles = 0
	cfg.MaxBytesPerFile = 20 // fails: src/main.py is 120 bytes
	delta := buildDelta([]dFile{
		{path: "src/main.py", size: 120},
		{path: "README.md", size: 40},
		{path: "config.yaml", size: 8, modified: true},
	})
	g, err := NewOutputCheckGrader("mixed", &cfg)
	if err != nil {
		t.Fatalf("NewOutputCheckGrader: %v", err)
	}
	res, err := g.Grade(context.Background(), GraderInput{WorkspaceDelta: delta})
	if err != nil {
		t.Fatalf("Grade: %v", err)
	}
	if res.Pass {
		t.Fatalf("expected overall fail, got pass; msg=%q", res.Message)
	}
	if n := len(res.Checks); n != 5 {
		t.Errorf("expected 5 sub-checks, got %d: %+v", n, res.Checks)
	}
	// min_files, forbid_files, max_bytes_per_file should fail;
	// require_files and min_bytes_per_file should pass.
	wantFail := map[string]bool{"min_files (10)": true, "forbid_files": true, "max_bytes_per_file (20)": true}
	wantPass := map[string]bool{"require_files": true, "min_bytes_per_file (5)": true}
	seen := map[string]bool{}
	for _, sc := range res.Checks {
		seen[sc.Label] = true
		if wantFail[sc.Label] && sc.Pass {
			t.Errorf("sub-check %q expected fail, passed: %q", sc.Label, sc.Message)
		}
		if wantPass[sc.Label] && !sc.Pass {
			t.Errorf("sub-check %q expected pass, failed: %q", sc.Label, sc.Message)
		}
	}
	for k := range wantFail {
		if !seen[k] {
			t.Errorf("sub-check %q missing from results", k)
		}
	}
	for k := range wantPass {
		if !seen[k] {
			t.Errorf("sub-check %q missing from results", k)
		}
	}
}

func TestOutputCheckGrader_AllChecksPass_OverallPass(t *testing.T) {
	cfg := OutputCheckConfig{
		MinFiles:        1,
		MaxFiles:        5,
		RequireFiles:    []string{"README.md"},
		ForbidFiles:     []string{".env"},
		RequireUpdated:  []string{"src/main.py"},
		MinBytesPerFile: 1,
		MaxBytesPerFile: 1000,
	}
	delta := buildDelta([]dFile{
		{path: "README.md", size: 100},
		{path: "src/main.py", size: 500, modified: true},
	})
	g, _ := NewOutputCheckGrader("all-pass", &cfg)
	res, _ := g.Grade(context.Background(), GraderInput{WorkspaceDelta: delta})
	if !res.Pass || res.Score != 1.0 {
		t.Errorf("want overall pass, got Pass=%v Score=%v msg=%q", res.Pass, res.Score, res.Message)
	}
	if len(res.Checks) != 7 {
		t.Errorf("expected 7 points, got %d", len(res.Checks))
	}
	for _, point := range res.Checks {
		if !point.Pass {
			t.Errorf("point %q unexpectedly failed: %q", point.Label, point.Message)
		}
	}
}

func TestOutputCheckGrader_BoundaryConditions(t *testing.T) {
	tests := []struct {
		name     string
		cfg      OutputCheckConfig
		delta    *WorkspaceDelta
		wantPass bool
	}{
		{
			name:     "min_files exact boundary pass",
			cfg:      OutputCheckConfig{MinFiles: 2},
			delta:    buildDelta([]dFile{{path: "a", size: 1}, {path: "b", size: 1}}),
			wantPass: true,
		},
		{
			name:     "max_files exact boundary pass",
			cfg:      OutputCheckConfig{MaxFiles: 2},
			delta:    buildDelta([]dFile{{path: "a", size: 1}, {path: "b", size: 1}}),
			wantPass: true,
		},
		{
			name:     "min_bytes_per_file exact boundary pass",
			cfg:      OutputCheckConfig{MinBytesPerFile: 10},
			delta:    buildDelta([]dFile{{path: "a", size: 10}}),
			wantPass: true,
		},
		{
			name:     "max_bytes_per_file exact boundary pass",
			cfg:      OutputCheckConfig{MaxBytesPerFile: 10},
			delta:    buildDelta([]dFile{{path: "a", size: 10}}),
			wantPass: true,
		},
		{
			name:     "zero-byte file fails min_bytes_per_file=1",
			cfg:      OutputCheckConfig{MinBytesPerFile: 1},
			delta:    buildDelta([]dFile{{path: "empty", size: 0}}),
			wantPass: false,
		},
		{
			name:     "zero-byte file ok when min_bytes_per_file unset",
			cfg:      OutputCheckConfig{MinFiles: 1},
			delta:    buildDelta([]dFile{{path: "empty", size: 0}}),
			wantPass: true,
		},
		{
			name:     "min_bytes_per_file with empty delta fails (no files to check)",
			cfg:      OutputCheckConfig{MinBytesPerFile: 10},
			delta:    &WorkspaceDelta{},
			wantPass: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := NewOutputCheckGrader(tt.name, &tt.cfg)
			if err != nil {
				t.Fatalf("NewOutputCheckGrader: %v", err)
			}
			res, _ := g.Grade(context.Background(), GraderInput{WorkspaceDelta: tt.delta})
			if res.Pass != tt.wantPass {
				t.Errorf("Pass=%v want %v; msg=%q", res.Pass, tt.wantPass, res.Message)
			}
		})
	}
}

func TestNewOutputCheckGrader_ConfigValidation(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *OutputCheckConfig
		wantErr string
	}{
		{"nil config", nil, "config is required"},
		{"negative min_files", &OutputCheckConfig{MinFiles: -1}, "min_files must be >= 0"},
		{"negative max_files", &OutputCheckConfig{MaxFiles: -1}, "max_files must be >= 0"},
		{"min > max files", &OutputCheckConfig{MinFiles: 5, MaxFiles: 2}, "> max_files"},
		{"negative min_bytes", &OutputCheckConfig{MinBytesPerFile: -1}, "min_bytes_per_file must be >= 0"},
		{"negative max_bytes", &OutputCheckConfig{MaxBytesPerFile: -1}, "max_bytes_per_file must be >= 0"},
		{"min > max bytes", &OutputCheckConfig{MinBytesPerFile: 100, MaxBytesPerFile: 10}, "> max_bytes_per_file"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewOutputCheckGrader(tt.name, tt.cfg)
			if err == nil {
				t.Fatalf("expected error containing %q, got nil", tt.wantErr)
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Errorf("error %q does not contain %q", err.Error(), tt.wantErr)
			}
		})
	}
}

func TestOutputCheckGrader_RegistryAndParse(t *testing.T) {
	yamlSrc := `
graders:
  - name: produced-files
    kind: output_check
    config:
      min_files: 1
      max_files: 50
      require_files: [README.md]
      forbid_files: [.env, secrets.json]
      require_updated: [src/main.py]
      min_bytes_per_file: 10
      max_bytes_per_file: 1048576
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

	// Decode the config and verify all knobs round-tripped.
	decoded, err := gcf.Graders[0].DecodeConfig()
	if err != nil {
		t.Fatalf("DecodeConfig: %v", err)
	}
	cfg, ok := decoded.(*OutputCheckConfig)
	if !ok {
		t.Fatalf("decoded type = %T want *OutputCheckConfig", decoded)
	}
	if cfg.MinFiles != 1 || cfg.MaxFiles != 50 {
		t.Errorf("min/max files wrong: %+v", cfg)
	}
	if len(cfg.RequireFiles) != 1 || cfg.RequireFiles[0] != "README.md" {
		t.Errorf("require_files = %v", cfg.RequireFiles)
	}
	if len(cfg.ForbidFiles) != 2 {
		t.Errorf("forbid_files = %v", cfg.ForbidFiles)
	}
	if len(cfg.RequireUpdated) != 1 || cfg.RequireUpdated[0] != "src/main.py" {
		t.Errorf("require_updated = %v", cfg.RequireUpdated)
	}
	if cfg.MinBytesPerFile != 10 || cfg.MaxBytesPerFile != 1048576 {
		t.Errorf("byte knobs wrong: %+v", cfg)
	}

	// Negative min_files should fail in NewGrader.
	var n yaml.Node
	if err := yaml.Unmarshal([]byte("min_files: -1"), &n); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	bad := GraderConfig{Kind: KindOutputCheck, Name: "bad", Config: n}
	if _, err := NewGrader(bad); err == nil {
		t.Errorf("expected error for negative min_files")
	}
}

func TestOutputCheckGrader_NewAndModifiedBothCount(t *testing.T) {
	// min_files counts new + modified as "produced".
	delta := buildDelta([]dFile{
		{path: "new.py", size: 50, modified: false},
		{path: "existing.py", size: 80, modified: true},
	})
	g, _ := NewOutputCheckGrader("both", &OutputCheckConfig{MinFiles: 2})
	res, _ := g.Grade(context.Background(), GraderInput{WorkspaceDelta: delta})
	if !res.Pass {
		t.Errorf("new + modified should both count toward min_files; msg=%q", res.Message)
	}
}
