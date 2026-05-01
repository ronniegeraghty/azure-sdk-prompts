// Unified grader loader (Phase 1 of Grader Unification — issue #624).
//
// ParseUnified, LoadUnifiedFile, and LoadUnifiedDir read the new flat-`type`
// schema from YAML and also translate legacy criteria.yaml files (no `type`
// field — implicit prompt graders) into the unified shape in-memory.
//
// Per Q4 of the locked decisions, malformed files are NOT fatal at load time.
// LoadUnifiedDir collects per-file validation errors into Bundle.FileErrors;
// the engine (Phase 2) is responsible for failing an eval only when it
// actually selects a file that failed to load. Loader-level fatal errors
// are reserved for I/O failures (e.g. directory walk).
package criteria

import (
	"bytes"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria/graders"
)

// Bundle is the result of loading a directory of unified grader configs.
// Configs is the list of files that parsed and validated successfully.
// FileErrors maps the absolute file path to the failure that prevented it
// from being included. Each FileError preserves the file-level `when:` block
// (when parseable), so MatchingErrors can report failures only for files an
// eval would actually have used.
type Bundle struct {
	Configs    []UnifiedGraderConfig
	FileErrors map[string]FileError
}

// FileError captures a per-file load/validation failure together with the
// file-level `when:` block, used to decide whether the failure is relevant
// to a given eval (Q4).
type FileError struct {
	Path string
	When map[string]string
	Err  error
}

// Error satisfies the error interface so FileError can be surfaced directly.
func (f FileError) Error() string {
	if f.Err == nil {
		return fmt.Sprintf("%s: unknown error", f.Path)
	}
	return fmt.Sprintf("%s: %v", f.Path, f.Err)
}

// Unwrap exposes the underlying cause for errors.Is/errors.As.
func (f FileError) Unwrap() error { return f.Err }

// ParseUnified decodes raw YAML bytes into a UnifiedGraderConfig, applying
// legacy back-compat translation and structural validation. Unknown YAML keys
// (including `gate` and `kind`) are rejected by yaml.v3's KnownFields(true).
//
// Source is left empty; callers (LoadUnifiedFile/LoadUnifiedDir) populate it.
func ParseUnified(data []byte) (*UnifiedGraderConfig, error) {
	var gc UnifiedGraderConfig
	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true)
	if err := dec.Decode(&gc); err != nil {
		return nil, fmt.Errorf("parsing unified grader config: %w", err)
	}
	translateLegacy(&gc)
	if err := validateConfig(&gc); err != nil {
		return nil, err
	}
	return &gc, nil
}

// translateLegacy walks every entry and infers Type="prompt" for any entry
// that omits `type:` but provides `prompt:`. This is the back-compat hook
// for existing internal/criteria YAML files which never carried a `type`
// discriminator. Entries that already have a type are left untouched.
//
// We do NOT silently coerce entries with neither type nor prompt — those
// will fail validation downstream with a clear "type is required" message.
func translateLegacy(gc *UnifiedGraderConfig) {
	translateSlice(gc.Graders)
	for gi := range gc.Groups {
		translateSlice(gc.Groups[gi].Graders)
	}
}

func translateSlice(entries []UnifiedGraderEntry) {
	for i := range entries {
		if entries[i].Type == "" && entries[i].Prompt != "" && !hasChecks(entries[i].Checks) {
			entries[i].Type = graders.KindPrompt
		}
	}
}

// LoadUnifiedFile reads, parses, and validates a single YAML file. Returns
// the populated UnifiedGraderConfig with Source set to path on success.
func LoadUnifiedFile(path string) (*UnifiedGraderConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}
	gc, err := ParseUnified(data)
	if err != nil {
		return nil, fmt.Errorf("in %s: %w", path, err)
	}
	gc.Source = path
	return gc, nil
}

// LoadUnifiedDir walks dir and loads every *.yaml / *.yml file under it.
// Per-file load failures are deferred via Bundle.FileErrors. Only I/O errors
// from the walk itself are returned as the second value.
func LoadUnifiedDir(dir string) (*Bundle, error) {
	bundle := &Bundle{FileErrors: map[string]FileError{}}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		if ext != ".yaml" && ext != ".yml" {
			return nil
		}

		gc, loadErr := LoadUnifiedFile(path)
		if loadErr != nil {
			when := peekFileWhen(path)
			bundle.FileErrors[path] = FileError{Path: path, When: when, Err: loadErr}
			slog.Warn("Deferred unified grader file error", "path", path, "error", loadErr)
			return nil
		}
		bundle.Configs = append(bundle.Configs, *gc)
		slog.Debug("Loaded unified grader config",
			"path", path,
			"top_graders", len(gc.Graders),
			"groups", len(gc.Groups),
		)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walking unified grader directory %s: %w", dir, err)
	}
	return bundle, nil
}

// peekFileWhen makes a best-effort read of the file's top-level `when:` block
// using a permissive (non-strict) decoder. Returns nil if the file is so
// broken that even the file-level when can't be extracted — in that case
// MatchingErrors will surface the failure to every eval (safe default).
func peekFileWhen(path string) map[string]string {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var probe struct {
		When map[string]string `yaml:"when"`
	}
	dec := yaml.NewDecoder(bytes.NewReader(data))
	// Permissive on purpose: we want the when block even if the rest is malformed.
	if err := dec.Decode(&probe); err != nil {
		return nil
	}
	return probe.When
}

// MatchingErrors returns every FileError whose file-level `when:` would have
// matched props, OR whose `when:` could not be peeked (nil map → assume
// relevant, fail-loud). Use this to satisfy Q4: "the eval fails only if the
// malformed file is actually used in that eval run."
//
// The empty case (no relevant errors) returns nil so callers can write
// `if err := bundle.MatchingErrors(props); err != nil`.
func (b *Bundle) MatchingErrors(props map[string]string) error {
	if b == nil || len(b.FileErrors) == 0 {
		return nil
	}
	var relevant []error
	for _, fe := range b.FileErrors {
		if fe.When == nil || matchesUnifiedWhen(fe.When, props) {
			relevant = append(relevant, fe)
		}
	}
	if len(relevant) == 0 {
		return nil
	}
	return errors.Join(relevant...)
}
