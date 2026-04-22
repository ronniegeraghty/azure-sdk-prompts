package criteria

import (
	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria/graders"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// --- entry-level validation -------------------------------------------------

func TestValidateEntry_PromptOK(t *testing.T) {
	e := UnifiedGraderEntry{Type: graders.KindPrompt, Name: "x", Prompt: "say hi"}
	if err := validateEntry(e); err != nil {
		t.Fatalf("expected valid prompt entry, got %v", err)
	}
}

func TestValidateEntry_TypedOK(t *testing.T) {
	gc, err := ParseUnified([]byte(`graders:
  - type: output_check
    name: must_have_files
    details:
      min_files: 1
`))
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if got := gc.Graders[0].Type; got != graders.KindOutputCheck {
		t.Fatalf("type = %q want %q", got, graders.KindOutputCheck)
	}
}

func TestValidateEntry_MissingType(t *testing.T) {
	// No type, no prompt → cannot infer; should error.
	_, err := ParseUnified([]byte(`graders:
  - name: nope
    details: { min_files: 1 }
`))
	if err == nil || !strings.Contains(err.Error(), "type is required") {
		t.Fatalf("want 'type is required', got %v", err)
	}
}

func TestValidateEntry_UnknownType(t *testing.T) {
	_, err := ParseUnified([]byte(`graders:
  - type: telepathy
    name: woo
    details: { x: 1 }
`))
	if err == nil || !strings.Contains(err.Error(), "unknown type") {
		t.Fatalf("want unknown type error, got %v", err)
	}
}

func TestValidateEntry_PromptWithDetails(t *testing.T) {
	_, err := ParseUnified([]byte(`graders:
  - type: prompt
    name: bad
    prompt: "do it"
    details: { min_files: 1 }
`))
	if err == nil || !strings.Contains(err.Error(), "must not set details") {
		t.Fatalf("want details-on-prompt error, got %v", err)
	}
}

func TestValidateEntry_TypedWithPrompt(t *testing.T) {
	_, err := ParseUnified([]byte(`graders:
  - type: output_check
    name: bad
    prompt: "do it"
    details: { min_files: 1 }
`))
	if err == nil || !strings.Contains(err.Error(), "must not set prompt") {
		t.Fatalf("want prompt-on-typed error, got %v", err)
	}
}

func TestValidateEntry_TypedMissingDetails(t *testing.T) {
	_, err := ParseUnified([]byte(`graders:
  - type: output_check
    name: bad
`))
	if err == nil || !strings.Contains(err.Error(), "requires non-empty details") {
		t.Fatalf("want missing-details error, got %v", err)
	}
}

func TestValidateEntry_PromptMissingPrompt(t *testing.T) {
	_, err := ParseUnified([]byte(`graders:
  - type: prompt
    name: bad
`))
	if err == nil || !strings.Contains(err.Error(), "requires non-empty prompt") {
		t.Fatalf("want missing-prompt error, got %v", err)
	}
}

func TestValidateEntry_NegativeWeight(t *testing.T) {
	_, err := ParseUnified([]byte(`graders:
  - type: prompt
    name: bad
    weight: -0.1
    prompt: "x"
`))
	if err == nil || !strings.Contains(err.Error(), "weight must be >= 0") {
		t.Fatalf("want weight error, got %v", err)
	}
}

// --- file-level rules -------------------------------------------------------

func TestValidate_DuplicateNamesAcrossGroupAndTopLevel(t *testing.T) {
	_, err := ParseUnified([]byte(`graders:
  - type: prompt
    name: dup
    prompt: "a"
groups:
  - name: g1
    graders:
      - type: prompt
        name: dup
        prompt: "b"
`))
	if err == nil || !strings.Contains(err.Error(), "duplicate grader name") {
		t.Fatalf("want duplicate-name error, got %v", err)
	}
}

func TestValidate_EmptyFile(t *testing.T) {
	_, err := ParseUnified([]byte(`when: { language: python }`))
	if err == nil || !strings.Contains(err.Error(), "no graders or groups") {
		t.Fatalf("want empty-file error, got %v", err)
	}
}

// --- KnownFields rejects gate / kind ---------------------------------------

func TestParse_RejectsGateField(t *testing.T) {
	_, err := ParseUnified([]byte(`graders:
  - type: prompt
    name: x
    prompt: "y"
    gate: true
`))
	if err == nil || !strings.Contains(err.Error(), "gate") {
		t.Fatalf("want gate-rejected error, got %v", err)
	}
}

func TestParse_RejectsKindField(t *testing.T) {
	_, err := ParseUnified([]byte(`graders:
  - kind: prompt
    name: x
    prompt: "y"
`))
	if err == nil || !strings.Contains(err.Error(), "kind") {
		t.Fatalf("want kind-rejected error, got %v", err)
	}
}

// --- legacy back-compat translation ----------------------------------------

func TestParse_LegacyImpliesPromptType(t *testing.T) {
	gc, err := ParseUnified([]byte(`when:
  language: java
graders:
  - name: Uses azure-identity
    weight: 1.0
    prompt: "Uses DefaultAzureCredential."
  - name: Builds clean
    prompt: "Code compiles."
`))
	if err != nil {
		t.Fatalf("legacy parse should succeed, got %v", err)
	}
	if len(gc.Graders) != 2 {
		t.Fatalf("graders=%d want 2", len(gc.Graders))
	}
	for i, g := range gc.Graders {
		if g.Type != graders.KindPrompt {
			t.Errorf("graders[%d].Type=%q want %q", i, g.Type, graders.KindPrompt)
		}
	}
}

func TestParse_LegacyMixedWithGroups(t *testing.T) {
	gc, err := ParseUnified([]byte(`groups:
  - name: auth
    when: { category: auth }
    graders:
      - name: DAC
        prompt: "uses DefaultAzureCredential"
      - name: NoSecrets
        prompt: "no hardcoded secrets"
`))
	if err != nil {
		t.Fatalf("legacy groups parse failed: %v", err)
	}
	if len(gc.Groups) != 1 || len(gc.Groups[0].Graders) != 2 {
		t.Fatalf("unexpected shape: %+v", gc)
	}
	for _, g := range gc.Groups[0].Graders {
		if g.Type != graders.KindPrompt {
			t.Errorf("group grader %q got Type=%q want %q", g.Name, g.Type, graders.KindPrompt)
		}
	}
}

// --- unified mixed file -----------------------------------------------------

func TestParse_MixedPromptAndTyped(t *testing.T) {
	gc, err := ParseUnified([]byte(`when: { language: python }
graders:
  - type: prompt
    name: review-1
    prompt: "Code is idiomatic Python."
  - type: output_check
    name: produced-files
    details:
      min_files: 1
      min_bytes_per_file: 1
`))
	if err != nil {
		t.Fatalf("parse failed: %v", err)
	}
	if gc.Graders[0].Type != graders.KindPrompt || gc.Graders[1].Type != graders.KindOutputCheck {
		t.Fatalf("types wrong: %+v", gc.Graders)
	}
}

// --- LoadDir + Bundle deferred errors --------------------------------------

func TestLoadUnifiedDir_DeferredErrorIsRelevantOnlyForMatchingProps(t *testing.T) {
	dir := t.TempDir()
	// Good file: matches python, no problems.
	if err := os.WriteFile(filepath.Join(dir, "python.yaml"), []byte(`when: { language: python }
graders:
  - type: prompt
    name: ok
    prompt: "fine"
`), 0o644); err != nil {
		t.Fatal(err)
	}
	// Bad file: matches java, but has a `gate` field which is rejected.
	if err := os.WriteFile(filepath.Join(dir, "java.yaml"), []byte(`when: { language: java }
graders:
  - type: prompt
    name: bad
    prompt: "x"
    gate: true
`), 0o644); err != nil {
		t.Fatal(err)
	}

	bundle, err := LoadUnifiedDir(dir)
	if err != nil {
		t.Fatalf("walk failure should not happen: %v", err)
	}
	if len(bundle.Configs) != 1 {
		t.Fatalf("want 1 good config, got %d", len(bundle.Configs))
	}
	if len(bundle.FileErrors) != 1 {
		t.Fatalf("want 1 deferred error, got %d", len(bundle.FileErrors))
	}

	// Python eval: bad java file is irrelevant.
	if err := bundle.MatchingErrors(map[string]string{"language": "python"}); err != nil {
		t.Fatalf("python eval should not surface java error, got %v", err)
	}
	// Java eval: bad java file IS relevant.
	if err := bundle.MatchingErrors(map[string]string{"language": "java"}); err == nil {
		t.Fatalf("java eval must surface deferred error")
	}
}

func TestLoadUnifiedDir_UnreadableWhenSurfacesUniversally(t *testing.T) {
	dir := t.TempDir()
	// Completely malformed: not even valid YAML, can't peek `when:`.
	if err := os.WriteFile(filepath.Join(dir, "broken.yaml"),
		[]byte("graders: [\n  - this is not yaml\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	bundle, err := LoadUnifiedDir(dir)
	if err != nil {
		t.Fatalf("walk: %v", err)
	}
	if len(bundle.FileErrors) != 1 {
		t.Fatalf("want 1 error, got %d", len(bundle.FileErrors))
	}
	// With nil When, error should surface for every eval (fail-loud default).
	if err := bundle.MatchingErrors(map[string]string{"language": "go"}); err == nil {
		t.Fatalf("expected fail-loud surfacing")
	}
}

func TestLoadUnifiedFile_SetsSource(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "x.yaml")
	if err := os.WriteFile(p, []byte(`graders:
  - type: prompt
    name: a
    prompt: "b"
`), 0o644); err != nil {
		t.Fatal(err)
	}
	gc, err := LoadUnifiedFile(p)
	if err != nil {
		t.Fatal(err)
	}
	if gc.Source != p {
		t.Fatalf("Source=%q want %q", gc.Source, p)
	}
}

// --- when helpers ----------------------------------------------------------

func TestMatchesUnifiedWhen_CaseInsensitiveAndEmpty(t *testing.T) {
	if !matchesUnifiedWhen(nil, map[string]string{"a": "b"}) {
		t.Error("nil when should match anything")
	}
	if !matchesUnifiedWhen(map[string]string{"language": "Python"},
		map[string]string{"language": "python"}) {
		t.Error("case-insensitive match failed")
	}
	if matchesUnifiedWhen(map[string]string{"language": "go"},
		map[string]string{"language": "rust"}) {
		t.Error("non-match should not match")
	}
}

func TestMergeUnifiedWhen_ChildOverridesParent(t *testing.T) {
	got := mergeUnifiedWhen(
		map[string]string{"language": "python", "service": "kv"},
		map[string]string{"language": "go"},
	)
	if got["language"] != "go" {
		t.Errorf("child should override: %v", got)
	}
	if got["service"] != "kv" {
		t.Errorf("parent should survive: %v", got)
	}
	if mergeUnifiedWhen(nil, nil) != nil {
		t.Error("merging two empties should return nil")
	}
}

// --- Effective weight ------------------------------------------------------

func TestEffectiveWeight_ZeroDefaultsToOne(t *testing.T) {
	if w := (UnifiedGraderEntry{}).EffectiveWeight(); w != 1.0 {
		t.Errorf("zero weight should default to 1.0, got %v", w)
	}
	if w := (UnifiedGraderEntry{Weight: 0.5}).EffectiveWeight(); w != 0.5 {
		t.Errorf("explicit weight should pass through, got %v", w)
	}
}

// --- Real legacy fixture compatibility -------------------------------------

func TestParse_RealLegacyJavaCriteria(t *testing.T) {
	// Mirrors the shape of criteria/language/java.yaml — implicit prompt graders.
	src := []byte(`when:
  language: java
graders:
  - name: Correct Dependencies
    weight: 1.0
    prompt: >
      Uses com.azure group ID for all Azure SDK packages.
  - name: BOM
    weight: 1.0
    prompt: >
      Uses azure-sdk-bom in dependencyManagement.
`)
	gc, err := ParseUnified(src)
	if err != nil {
		t.Fatalf("real legacy fixture should parse: %v", err)
	}
	if len(gc.Graders) != 2 {
		t.Fatalf("want 2 graders, got %d", len(gc.Graders))
	}
	for _, g := range gc.Graders {
		if g.Type != graders.KindPrompt {
			t.Errorf("legacy entry %q should translate to type=prompt, got %q", g.Name, g.Type)
		}
	}
}
