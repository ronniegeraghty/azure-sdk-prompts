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
func (g *FileGrader) Grade(_ context.Context, input GraderInput) (GraderResult, error) {
	fullPath := filepath.Join(input.WorkspacePath, g.path)

	result := GraderResult{
		Kind:   KindFile,
		Name:   g.name,
		Weight: input.Config.EffectiveWeight(),
		Gate:   input.Config.Gate,
		FileDetails: &FileGraderDetails{
			CheckedFiles: make([]FileCheckResult, 0, 1),
		},
	}

	check := FileCheckResult{
		Path:    g.path,
		Pattern: g.rawPat,
	}

	data, err := os.ReadFile(fullPath)
	if err != nil {
		check.Exists = false
		result.FileDetails.CheckedFiles = append(result.FileDetails.CheckedFiles, check)

		if g.mustExist {
			result.Score = 0
			result.Pass = false
			result.Message = fmt.Sprintf("required file %q not found", g.path)
			result.Points = []GraderPoint{{
				Name:    g.path,
				Pass:    false,
				Message: fmt.Sprintf("required file %q not found", g.path),
			}}
			return result, nil
		}
		// File not required and doesn't exist — that's fine.
		result.Score = 1.0
		result.Pass = true
		result.Message = fmt.Sprintf("optional file %q not found (ok)", g.path)
		result.Points = []GraderPoint{{
			Name:    g.path,
			Pass:    true,
			Message: fmt.Sprintf("optional file %q not found (ok)", g.path),
		}}
		return result, nil
	}

	check.Exists = true

	if g.pattern != nil {
		matched := g.pattern.Match(data)
		check.PatternMatched = &matched

		if !matched {
			result.FileDetails.CheckedFiles = append(result.FileDetails.CheckedFiles, check)
			result.Score = 0.5 // File exists but content doesn't match
			result.Pass = false
			result.Message = fmt.Sprintf("file %q exists but pattern %q not matched", g.path, g.rawPat)
			result.Points = []GraderPoint{{
				Name:    g.path,
				Pass:    false,
				Message: fmt.Sprintf("pattern %q not matched", g.rawPat),
			}}
			return result, nil
		}
	}

	result.FileDetails.CheckedFiles = append(result.FileDetails.CheckedFiles, check)
	result.Score = 1.0
	result.Pass = true
	result.Message = fmt.Sprintf("file %q check passed", g.path)
	pointMsg := fmt.Sprintf("file %q present", g.path)
	if g.pattern != nil {
		pointMsg = fmt.Sprintf("file %q present and matches %q", g.path, g.rawPat)
	}
	result.Points = []GraderPoint{{
		Name:    g.path,
		Pass:    true,
		Message: pointMsg,
	}}
	return result, nil
}
