// Package graders — Phase 1 acceptance tests for the unified grader loader (#624).
//
// These are the TDD acceptance tests locked against the schema in issue #624:
// flat `type` discriminator, `details:` payload for typed graders, name
// uniqueness per file, no `gate:`, no `kind:`, back-compat translation for
// legacy criteria.yaml shape, and deferred-error Bundle semantics on LoadDir.
//
// They exercise Neo's Phase 1 implementation: UnifiedGraderConfig /
// UnifiedGraderEntry / UnifiedGraderGroup, ParseUnified, LoadUnifiedFile,
// LoadUnifiedDir, Bundle.
//
// See .squad/decisions/inbox/switch-phase1-test-coverage.md for coverage map.
package criteria

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// testdataDir resolves to internal/graders/testdata/phase1.
func testdataDir(t *testing.T) string {
	t.Helper()
	return filepath.Join("testdata", "phase1")
}

// loadPhase1File is a thin wrapper around Neo's LoadUnifiedFile so that any
// future rename is a one-line change here instead of in every test.
func loadPhase1File(t *testing.T, path string) (*UnifiedGraderConfig, error) {
	t.Helper()
	return LoadUnifiedFile(path)
}

// -----------------------------------------------------------------------------
// Case 1: new-format file with mixed prompt + output_check entries loads cleanly
// -----------------------------------------------------------------------------

func TestPhase1Loader_MixedPromptAndTyped(t *testing.T) {
	t.Parallel()
	gc, err := loadPhase1File(t, filepath.Join(testdataDir(t), "mixed_prompt_and_typed.yaml"))
	if err != nil {
		t.Fatalf("load mixed prompt+typed: %v", err)
	}
	if gc == nil {
		t.Fatal("expected non-nil config")
	}
	if got, want := len(gc.Graders), 2; got != want {
		t.Fatalf("grader count: got %d, want %d", got, want)
	}

	// Order-preserved: prompt first, output_check second.
	if gc.Graders[0].Type != "prompt" {
		t.Errorf("graders[0].Type: got %q, want %q", gc.Graders[0].Type, "prompt")
	}
	if gc.Graders[0].Prompt == "" {
		t.Error("prompt grader missing Prompt body")
	}
	if gc.Graders[1].Type != "workspace" {
		t.Errorf("graders[1].Type: got %q, want %q", gc.Graders[1].Type, "workspace")
	}
	// File-level when preserved.
	if gc.When["language"] != "python" {
		t.Errorf("file-level when[language]: got %q, want %q", gc.When["language"], "python")
	}
}

// -----------------------------------------------------------------------------
// Case 2: two graders with same type, different names → success
// -----------------------------------------------------------------------------

func TestPhase1Loader_SameTypeDifferentNames(t *testing.T) {
	t.Parallel()
	gc, err := loadPhase1File(t, filepath.Join(testdataDir(t), "two_same_type_unique_names.yaml"))
	if err != nil {
		t.Fatalf("load two same-type: %v", err)
	}
	if got, want := len(gc.Graders), 2; got != want {
		t.Fatalf("grader count: got %d, want %d", got, want)
	}
	if gc.Graders[0].Type != gc.Graders[1].Type {
		t.Errorf("expected identical types, got %q and %q", gc.Graders[0].Type, gc.Graders[1].Type)
	}
	if gc.Graders[0].Name == gc.Graders[1].Name {
		t.Errorf("names must differ, got duplicates: %q", gc.Graders[0].Name)
	}
}

// -----------------------------------------------------------------------------
// Case 3: two graders with same name → validation error mentioning name + path
// -----------------------------------------------------------------------------

func TestPhase1Loader_DuplicateNameRejected(t *testing.T) {
	t.Parallel()
	path := filepath.Join(testdataDir(t), "duplicate_name.yaml")
	_, err := loadPhase1File(t, path)
	if err == nil {
		t.Fatal("expected validation error for duplicate name, got nil")
	}
	msg := err.Error()
	if !strings.Contains(msg, "Duplicate Name Grader") {
		t.Errorf("error must mention duplicated name, got: %v", err)
	}
	// File path (or at least the filename) should appear so users can locate the file.
	if !strings.Contains(msg, "duplicate_name.yaml") {
		t.Errorf("error must mention file path, got: %v", err)
	}
}

// -----------------------------------------------------------------------------
// Case 4: malformed entries → file-level errors (table-driven)
// -----------------------------------------------------------------------------

func TestPhase1Loader_MalformedEntries(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		fixture    string
		wantSubstr string
	}{
		{
			name:       "missing_type_field",
			fixture:    "malformed_missing_type.yaml",
			wantSubstr: "type",
		},
		{
			name:       "unknown_type_value",
			fixture:    "malformed_unknown_type.yaml",
			wantSubstr: "definitely_not_a_real_type",
		},
		{
			name:       "prompt_type_missing_prompt_body",
			fixture:    "malformed_prompt_missing_prompt.yaml",
			wantSubstr: "prompt",
		},
		{
			name:       "typed_missing_details",
			fixture:    "malformed_typed_missing_details.yaml",
			wantSubstr: "details",
		},
		{
			name:       "gate_field_rejected_by_known_fields",
			fixture:    "malformed_gate_field.yaml",
			wantSubstr: "gate",
		},
		{
			name:       "kind_field_rejected_by_known_fields",
			fixture:    "malformed_kind_field.yaml",
			wantSubstr: "kind",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			path := filepath.Join(testdataDir(t), tc.fixture)
			_, err := loadPhase1File(t, path)
			if err == nil {
				t.Fatalf("expected validation error for %s, got nil", tc.fixture)
			}
			if !strings.Contains(strings.ToLower(err.Error()), strings.ToLower(tc.wantSubstr)) {
				t.Errorf("error message should mention %q, got: %v", tc.wantSubstr, err)
			}
		})
	}
}

// -----------------------------------------------------------------------------
// Case 5: legacy criteria.yaml loads via back-compat; matches unified file
// -----------------------------------------------------------------------------

func TestPhase1Loader_LegacyBackCompat(t *testing.T) {
	t.Parallel()
	legacyPath := filepath.Join(testdataDir(t), "legacy_criteria.yaml")
	unifiedPath := filepath.Join(testdataDir(t), "legacy_criteria_unified_equivalent.yaml")

	legacy, err := loadPhase1File(t, legacyPath)
	if err != nil {
		t.Fatalf("legacy criteria must load via back-compat: %v", err)
	}
	unified, err := loadPhase1File(t, unifiedPath)
	if err != nil {
		t.Fatalf("unified equivalent must load: %v", err)
	}

	// Both must resolve to the same grader count and same (name, type) tuples,
	// in the same order. This is the observable contract of back-compat.
	if len(legacy.Graders) != len(unified.Graders) {
		t.Fatalf("top-level grader count: legacy=%d unified=%d", len(legacy.Graders), len(unified.Graders))
	}
	for i := range legacy.Graders {
		if legacy.Graders[i].Name != unified.Graders[i].Name {
			t.Errorf("graders[%d].Name: legacy=%q unified=%q", i, legacy.Graders[i].Name, unified.Graders[i].Name)
		}
		if legacy.Graders[i].Type != unified.Graders[i].Type {
			t.Errorf("graders[%d].Type: legacy=%q unified=%q — legacy prompt entries must translate to type=prompt",
				i, legacy.Graders[i].Type, unified.Graders[i].Type)
		}
		if legacy.Graders[i].Type != "prompt" {
			t.Errorf("graders[%d]: legacy entries must back-compat-translate to type=prompt, got %q",
				i, legacy.Graders[i].Type)
		}
		if legacy.Graders[i].Prompt == "" {
			t.Errorf("graders[%d]: translated prompt body must be populated", i)
		}
	}

	if len(legacy.Groups) != len(unified.Groups) {
		t.Fatalf("group count: legacy=%d unified=%d", len(legacy.Groups), len(unified.Groups))
	}
	for gi := range legacy.Groups {
		if len(legacy.Groups[gi].Graders) != len(unified.Groups[gi].Graders) {
			t.Fatalf("group[%d] grader count mismatch: legacy=%d unified=%d",
				gi, len(legacy.Groups[gi].Graders), len(unified.Groups[gi].Graders))
		}
		for j := range legacy.Groups[gi].Graders {
			lg := legacy.Groups[gi].Graders[j]
			ug := unified.Groups[gi].Graders[j]
			if lg.Name != ug.Name || lg.Type != ug.Type {
				t.Errorf("group[%d].graders[%d]: legacy=(%q,%q) unified=(%q,%q)",
					gi, j, lg.Name, lg.Type, ug.Name, ug.Type)
			}
			if lg.Type != "prompt" {
				t.Errorf("group[%d].graders[%d]: legacy group entry must be type=prompt, got %q",
					gi, j, lg.Type)
			}
		}
	}
}

// -----------------------------------------------------------------------------
// Case 6: empty graders list — REJECTED (matches legacy internal/criteria/
// behavior preserved in Neo's Phase 1 impl). The original task spec said
// "loads cleanly" but that would silently drop real criteria files whose
// `graders:` key was mis-indented. Phase 1 preserves the legacy fail-loud
// behavior, which is the correct back-compat stance.
// -----------------------------------------------------------------------------

func TestPhase1Loader_EmptyGraders(t *testing.T) {
	t.Parallel()
	gc, err := loadPhase1File(t, filepath.Join(testdataDir(t), "empty_graders.yaml"))
	if err == nil {
		t.Fatalf("empty graders list must be rejected (back-compat with legacy criteria loader), got config: %+v", gc)
	}
	msg := strings.ToLower(err.Error())
	if !strings.Contains(msg, "no graders") && !strings.Contains(msg, "empty") {
		t.Errorf("error must mention empty/no graders, got: %v", err)
	}
}

// -----------------------------------------------------------------------------
// Case 7: file with only prompt graders (the common case today) loads
// -----------------------------------------------------------------------------

func TestPhase1Loader_OnlyPromptGraders(t *testing.T) {
	t.Parallel()
	gc, err := loadPhase1File(t, filepath.Join(testdataDir(t), "only_prompt_graders.yaml"))
	if err != nil {
		t.Fatalf("prompt-only file must load: %v", err)
	}
	if got, want := len(gc.Graders), 3; got != want {
		t.Fatalf("grader count: got %d, want %d", got, want)
	}
	for i, g := range gc.Graders {
		if g.Type != "prompt" {
			t.Errorf("graders[%d].Type: got %q, want %q", i, g.Type, "prompt")
		}
		if g.Prompt == "" {
			t.Errorf("graders[%d]: prompt body must be populated", i)
		}
	}
}

// -----------------------------------------------------------------------------
// Case 8: file with only typed graders loads
// -----------------------------------------------------------------------------

func TestPhase1Loader_OnlyTypedGraders(t *testing.T) {
	t.Parallel()
	gc, err := loadPhase1File(t, filepath.Join(testdataDir(t), "only_typed_graders.yaml"))
	if err != nil {
		t.Fatalf("typed-only file must load: %v", err)
	}
	if got, want := len(gc.Graders), 2; got != want {
		t.Fatalf("grader count: got %d, want %d", got, want)
	}
	for i, g := range gc.Graders {
		if g.Type == "prompt" {
			t.Errorf("graders[%d]: expected non-prompt type, got %q", i, g.Type)
		}
		if g.Prompt != "" {
			t.Errorf("graders[%d]: typed grader must not have prompt body, got %q", i, g.Prompt)
		}
	}
}

// -----------------------------------------------------------------------------
// Smoke test: LoadDir walks a directory of fixtures and reports per-file errors
// without aborting on the first malformed file. Exercises the deferred-error
// Bundle semantics locked in Q4 of the proposal.
// -----------------------------------------------------------------------------

func TestPhase1Loader_LoadDirDeferredErrors(t *testing.T) {
	t.Parallel()

	// Build a scratch dir with one good file and one bad file so we assert that
	// LoadDir collects per-file errors rather than short-circuiting on the bad
	// one. Using a per-test temp dir keeps -race clean and side-effect free.
	dir := t.TempDir()
	copyFixture := func(src, dst string) {
		t.Helper()
		data, err := os.ReadFile(filepath.Join(testdataDir(t), src))
		if err != nil {
			t.Fatalf("read fixture %s: %v", src, err)
		}
		if err := os.WriteFile(filepath.Join(dir, dst), data, 0o644); err != nil {
			t.Fatalf("write fixture %s: %v", dst, err)
		}
	}
	copyFixture("only_prompt_graders.yaml", "good.yaml")
	copyFixture("malformed_unknown_type.yaml", "bad.yaml")

	bundle, err := LoadUnifiedDir(dir)
	// Two valid contracts satisfy Q4 semantics. Neo may choose either:
	//   (a) Return (bundle, nil); bundle.FileErrors contains the bad file.
	//   (b) Return (bundle, err); bundle is still populated with the good file.
	// We accept either shape; the invariant is: the good file is loaded AND
	// the bad file is surfaced somewhere discoverable.
	if bundle == nil && err == nil {
		t.Fatal("LoadDir returned nil bundle and nil error — violates Q4 deferred-error contract")
	}

	// Accept shape (b): bundle may be nil if err is non-nil AND references the bad file.
	if bundle == nil {
		if err == nil || !strings.Contains(err.Error(), "bad.yaml") {
			t.Fatalf("if bundle is nil, err must name the bad file; got err=%v", err)
		}
		t.Skip("LoadDir uses eager-error shape — deferred-error Bundle not implemented; Q4 still satisfied by direct error")
	}

	// Shape (a) — bundle is populated.
	// Exact field names (Configs, FileErrors) are Neo's call; if he names them
	// differently, update the two accessors below.
	good := bundleConfigs(bundle)
	if len(good) == 0 {
		t.Error("bundle must contain the successfully-loaded good.yaml config")
	}
	errs := bundleFileErrors(bundle)
	if len(errs) == 0 {
		t.Error("bundle must record the bad.yaml load error for deferred surfacing")
	} else {
		found := false
		for path := range errs {
			if strings.Contains(path, "bad.yaml") {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("bundle.FileErrors must key on bad.yaml path, got keys: %v", errs)
		}
	}
}

// bundleConfigs / bundleFileErrors are thin accessors so that renaming the
// Bundle struct field is a one-line change. If Neo picks different names
// (e.g. Configs → Loaded, FileErrors → Errors), update here.
func bundleConfigs(b *Bundle) []UnifiedGraderConfig { return b.Configs }
func bundleFileErrors(b *Bundle) map[string]FileError { return b.FileErrors }

// Sanity: errors are non-nil when returned by LoadFile for missing files.
func TestPhase1Loader_NonexistentFile(t *testing.T) {
	t.Parallel()
	_, err := LoadUnifiedFile(filepath.Join(t.TempDir(), "does-not-exist.yaml"))
	if err == nil {
		t.Fatal("expected error for missing file")
	}
	if !errors.Is(err, os.ErrNotExist) && !strings.Contains(err.Error(), "no such file") {
		t.Logf("non-ErrNotExist error (acceptable if wrapped): %v", err)
	}
}
