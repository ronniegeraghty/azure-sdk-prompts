package graders

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// FileGrader checks file existence and content patterns in the agent workspace.
type FileGrader struct {
	name      string
	path      string         // Required file path (relative to workspace)
	pattern   *regexp.Regexp // Optional compiled content regex
	rawPat    string         // Original pattern string for reporting
	mustExist bool           // If true, file must exist (default true)
}

// NewFileGrader constructs a FileGrader from a parsed FileConfig.
func NewFileGrader(name string, cfg *FileConfig) (*FileGrader, error) {
	if cfg.Path == "" {
		return nil, fmt.Errorf("file grader %q: path is required", name)
	}

	mustExist := true
	if cfg.MustExist != nil {
		mustExist = *cfg.MustExist
	}

	var compiled *regexp.Regexp
	if cfg.Pattern != "" {
		var err error
		compiled, err = regexp.Compile(cfg.Pattern)
		if err != nil {
			return nil, fmt.Errorf("file grader %q: invalid pattern %q: %w", name, cfg.Pattern, err)
		}
	}

	return &FileGrader{
		name:      name,
		path:      cfg.Path,
		pattern:   compiled,
		rawPat:    cfg.Pattern,
		mustExist: mustExist,
	}, nil
}

// Kind returns the grader type identifier.
func (g *FileGrader) Kind() string { return KindFile }

// Name returns the human-readable name.
func (g *FileGrader) Name() string { return g.name }

// Grade checks file existence and optionally matches content against the pattern.
// Per v4 spec: emits two points when Pattern is set (file present + pattern matches),
// one point otherwise (file present). No 0.5 partial credit.
func (g *FileGrader) Grade(_ context.Context, input GraderInput) (GraderResult, error) {
	fullPath := filepath.Join(input.WorkspacePath, g.path)

	var points []GraderPoint
	var fileExtra FileExtra

	fileExtra.Path = g.path
	fileExtra.Pattern = g.rawPat

	data, err := os.ReadFile(fullPath)
	if err != nil {
		// File doesn't exist
		fileExtra.Exists = false

		if g.mustExist {
			// Required file missing → fail
			points = append(points, GraderPoint{
				Label:   fmt.Sprintf("file present: %s", g.path),
				Pass:    false,
				Message: fmt.Sprintf("file not found at %s", fullPath),
			})
			msg := fmt.Sprintf("required file %q not found", g.path)
			return NewResult(KindFile, g.name, input.Config, points, msg, &GraderExtras{
				File: &FileExtras{Files: []FileExtra{fileExtra}},
			}), nil
		}

		// Optional file missing → pass
		points = append(points, GraderPoint{
			Label:   fmt.Sprintf("file present: %s", g.path),
			Pass:    true,
			Message: fmt.Sprintf("optional file %q not required", g.path),
		})
		msg := fmt.Sprintf("optional file %q not found (ok)", g.path)
		return NewResult(KindFile, g.name, input.Config, points, msg, &GraderExtras{
			File: &FileExtras{Files: []FileExtra{fileExtra}},
		}), nil
	}

	// File exists
	fileExtra.Exists = true
	stat, _ := os.Stat(fullPath)
	if stat != nil {
		fileExtra.Size = stat.Size()
	}

	// First point: file present
	points = append(points, GraderPoint{
		Label:   fmt.Sprintf("file present: %s", g.path),
		Pass:    true,
		Message: "",
	})

	// Second point (if pattern configured): pattern matches
	if g.pattern != nil {
		matched := g.pattern.Match(data)
		fileExtra.PatternMatched = matched

		label := fmt.Sprintf("pattern matches: %s", g.path)
		if matched {
			points = append(points, GraderPoint{
				Label:    label,
				Pass:     true,
				Message:  "",
				Evidence: map[string]string{"pattern": g.rawPat},
			})
		} else {
			points = append(points, GraderPoint{
				Label:    label,
				Pass:     false,
				Message:  fmt.Sprintf("pattern %q not found in file", g.rawPat),
				Evidence: map[string]string{"pattern": g.rawPat},
			})
		}
	}

	msg := fmt.Sprintf("file %q check passed", g.path)
	if g.pattern != nil && !fileExtra.PatternMatched {
		msg = fmt.Sprintf("file %q exists but pattern %q not matched", g.path, g.rawPat)
	}

	return NewResult(KindFile, g.name, input.Config, points, msg, &GraderExtras{
		File: &FileExtras{Files: []FileExtra{fileExtra}},
	}), nil
}
