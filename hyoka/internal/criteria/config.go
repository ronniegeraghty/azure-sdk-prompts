// Unified grader schema (Phase 1 of Grader Unification — issue #624).
//
// This file defines the new YAML schema that supersedes both
// internal/criteria/.GraderEntry and the legacy GraderConfig in types.go.
//
// Every grader entry has a flat `type` discriminator. Prompt graders
// (`type: prompt`) carry the LLM-review prompt in `prompt:`. All other
// supported types carry their typed configuration in `details:`. There is
// no `gate` field — every grader runs and reports its result.
//
// The schema is loaded by unified_loader.go. The engine wiring is Phase 2
// (issue #625) and is intentionally not done here.
package criteria

import (
	"fmt"
	"log/slog"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria/graders"
)

// UnifiedGraderEntry is one grader in a criteria/graders YAML file under the
// unified schema. The flat `type` field selects the grader implementation.
type UnifiedGraderEntry struct {
	// Type is the grader-kind discriminator. Required. One of "prompt"
	// or any value in validTypedKinds (file, program, behavior,
	// action_sequence, tool_constraint, prompt_review, output_check).
	Type string `yaml:"type" json:"type"`

	// Name uniquely identifies this grader within its file. Required.
	Name string `yaml:"name" json:"name"`

	// Weight is the grader's contribution to the aggregate score.
	// Defaults to 1.0 when zero/unset.
	Weight float64 `yaml:"weight,omitempty" json:"weight,omitempty"`

	// When narrows applicability by prompt properties. AND-matched,
	// case-insensitive. Empty matches everything.
	When map[string]string `yaml:"when,omitempty" json:"when,omitempty"`

	// Isolate, when true and the engine runs in --review-mode isolated,
	// places this grader in its own review session. Only meaningful for
	// type=prompt; silently ignored for typed graders.
	Isolate bool `yaml:"isolate,omitempty" json:"isolate,omitempty"`

	// Prompt is the LLM-review prompt. For type=prompt, either Prompt or
	// Checks must be non-empty (both may be set; when Checks is non-empty
	// Prompt acts as preamble text shown to the judge before the numbered
	// checks). Must be empty for any other type.
	Prompt string `yaml:"prompt,omitempty" json:"prompt,omitempty"`

	// Checks lists the individual pass/fail items the LLM judge must
	// evaluate. Allowed only for type=prompt. Each non-empty string becomes
	// one line in the rendered review criteria and one Point in the
	// resulting GraderResult. When set, Prompt is treated as optional
	// preamble text shown to the judge before the numbered checks.
	Checks []string `yaml:"checks,omitempty" json:"checks,omitempty"`

	// Details is the typed-grader payload (e.g. min_files for output_check,
	// path for file, command for program). Required for any type other
	// than prompt; must be empty for type=prompt.
	Details yaml.Node `yaml:"details,omitempty" json:"details,omitempty"`
}

// UnifiedGraderGroup is a named collection of grader entries with optional
// group-level when conditions. Groups support hierarchical when resolution
// (file → group → grader; most specific wins).
type UnifiedGraderGroup struct {
	Name    string               `yaml:"name,omitempty" json:"name,omitempty"`
	When    map[string]string    `yaml:"when,omitempty" json:"when,omitempty"`
	Graders []UnifiedGraderEntry `yaml:"graders" json:"graders"`
	Isolate bool                 `yaml:"isolate,omitempty" json:"isolate,omitempty"`
}

// UnifiedGraderConfig is the top-level shape of a unified criteria YAML file.
// A file may define top-level graders, groups, or both. The Source field is
// populated by the loader and is never read from YAML.
type UnifiedGraderConfig struct {
	When    map[string]string    `yaml:"when,omitempty" json:"when,omitempty"`
	Graders []UnifiedGraderEntry `yaml:"graders,omitempty" json:"graders,omitempty"`
	Groups  []UnifiedGraderGroup `yaml:"groups,omitempty" json:"groups,omitempty"`
	Source  string               `yaml:"-" json:"-"`
}

// validTypedKinds enumerates the non-prompt grader types accepted by the
// unified schema. Mirrors validKinds in types.go minus graders.KindPrompt.
// KindPromptReview is NOT included — it's the kind of manually-created
// PromptReviewGrader instances, not a valid criteria-file type.
var validTypedKinds = map[string]bool{
	graders.KindFile:           true,
	graders.KindProgram:        true,
	graders.KindBehavior:       true,
	graders.KindActionSequence: true,
	graders.KindToolConstraint: true,
	// graders.KindOutputCheck removed — replaced by "workspace"
	graders.KindToolUsage:      true,
	graders.KindTool:           true,
	"workspace":                true, // Canonical workspace grader
}

// IsValidUnifiedType returns true if t is a recognized unified-schema type
// (graders.KindPrompt or any typed kind).
func IsValidUnifiedType(t string) bool {
	if t == graders.KindPrompt {
		return true
	}
	return validTypedKinds[t]
}

// hasDetails reports whether n carries a non-empty YAML payload. yaml.v3
// returns Kind==0 for absent fields; an explicit empty mapping has Kind!=0.
func hasDetails(n yaml.Node) bool {
	if n.Kind == 0 {
		return false
	}
	// A null scalar (`details: ~`) is treated as empty.
	if n.Kind == yaml.ScalarNode && (n.Tag == "!!null" || n.Value == "") {
		return false
	}
	// An empty mapping or sequence counts as empty too.
	if (n.Kind == yaml.MappingNode || n.Kind == yaml.SequenceNode) && len(n.Content) == 0 {
		return false
	}
	return true
}

// validateEntry checks a single entry against the unified-schema rules.
// Returns nil if valid; otherwise an error including the entry name when
// available.
func validateEntry(e UnifiedGraderEntry) error {
	tag := e.Name
	if tag == "" {
		tag = "<unnamed>"
	}
	if e.Name == "" {
		return fmt.Errorf("grader entry: name is required")
	}
	if e.Type == "" {
		return fmt.Errorf("grader %q: type is required", tag)
	}
	
	// Loud migration error for renamed types
	if e.Type == graders.KindOutputCheck {
		return fmt.Errorf("grader %q: type %q has been renamed to \"workspace\" with new check kinds; see docs/graders.md for migration guide", tag, e.Type)
	}
	
	if !IsValidUnifiedType(e.Type) {
		return fmt.Errorf("grader %q: unknown type %q", tag, e.Type)
	}
	if e.Weight < 0 {
		return fmt.Errorf("grader %q: weight must be >= 0, got %f", tag, e.Weight)
	}
	if e.Type == graders.KindPrompt {
		hasPrompt := strings.TrimSpace(e.Prompt) != ""
		hasChecks := len(e.Checks) > 0
		if !hasPrompt && !hasChecks {
			return fmt.Errorf("grader %q: type=prompt requires non-empty prompt or checks", tag)
		}
		if hasPrompt && !hasChecks {
			slog.Warn("grader type=prompt has prompt but no checks: will synthesize a single pass/fail check at runtime; prefer adding explicit checks",
				"grader", tag)
		}
		if hasChecks {
			for i, c := range e.Checks {
				if strings.TrimSpace(c) == "" {
					return fmt.Errorf("grader %q: checks[%d] must be a non-empty string", tag, i)
				}
			}
		}
		if hasDetails(e.Details) {
			return fmt.Errorf("grader %q: type=prompt must not set details", tag)
		}
		return nil
	}
	// Typed grader.
	if e.Prompt != "" {
		return fmt.Errorf("grader %q: type=%s must not set prompt (use details)", tag, e.Type)
	}
	if len(e.Checks) > 0 {
		return fmt.Errorf("grader %q: type=%s must not set checks (only valid for type=prompt)", tag, e.Type)
	}
	if !hasDetails(e.Details) {
		return fmt.Errorf("grader %q: type=%s requires non-empty details", tag, e.Type)
	}
	return nil
}

// validateConfig validates a fully-translated UnifiedGraderConfig. It enforces
// per-entry validity AND name uniqueness across the file (top-level graders
// plus all group graders share one namespace).
func validateConfig(gc *UnifiedGraderConfig) error {
	seen := make(map[string]string) // name → location (for error messages)
	check := func(e UnifiedGraderEntry, where string) error {
		if err := validateEntry(e); err != nil {
			return err
		}
		if prev, ok := seen[e.Name]; ok {
			return fmt.Errorf("duplicate grader name %q (already defined at %s)", e.Name, prev)
		}
		seen[e.Name] = where
		return nil
	}
	for i, e := range gc.Graders {
		if err := check(e, fmt.Sprintf("graders[%d]", i)); err != nil {
			return err
		}
	}
	for gi, grp := range gc.Groups {
		for ei, e := range grp.Graders {
			loc := fmt.Sprintf("groups[%d].graders[%d]", gi, ei)
			if grp.Name != "" {
				loc = fmt.Sprintf("groups[%q].graders[%d]", grp.Name, ei)
			}
			if err := check(e, loc); err != nil {
				return err
			}
		}
	}
	if len(gc.Graders) == 0 && len(gc.Groups) == 0 {
		return fmt.Errorf("no graders or groups defined")
	}
	return nil
}

// EffectiveWeight returns the entry weight, defaulting to 1.0 when zero.
func (e UnifiedGraderEntry) EffectiveWeight() float64 {
	if e.Weight == 0 {
		return 1.0
	}
	return e.Weight
}

// matchesUnifiedWhen reports whether all when constraints match props
// (case-insensitive). An empty when always matches. Equivalent to the
// matchesWhen helper in internal/criteria — duplicated here to keep the
// new package self-contained for Phase 1.
func matchesUnifiedWhen(when, props map[string]string) bool {
	for k, v := range when {
		if !strings.EqualFold(props[k], v) {
			return false
		}
	}
	return true
}

// mergeUnifiedWhen merges parent and child when maps; child wins on key
// collisions. Returns nil when both inputs are empty.
func mergeUnifiedWhen(parent, child map[string]string) map[string]string {
	if len(parent) == 0 && len(child) == 0 {
		return nil
	}
	merged := make(map[string]string, len(parent)+len(child))
	for k, v := range parent {
		merged[k] = v
	}
	for k, v := range child {
		merged[k] = v
	}
	return merged
}
