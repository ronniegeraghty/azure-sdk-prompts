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

	// When narrows applicability by prompt properties and config state.
	// All fields AND together. Empty matches everything.
	When WhenClause `yaml:"when,omitempty" json:"when,omitempty"`

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
	//
	// For typed graders (workspace, tool, activity, program), Checks is a
	// yaml.Node that will be decoded into the type-specific []XxxCheck slice.
	Checks yaml.Node `yaml:"checks,omitempty" json:"checks,omitempty"`
}

// UnifiedGraderGroup is a named collection of grader entries with optional
// group-level when conditions. Groups support hierarchical when resolution
// (file → group → grader; most specific wins).
type UnifiedGraderGroup struct {
	Name    string               `yaml:"name,omitempty" json:"name,omitempty"`
	When    WhenClause           `yaml:"when,omitempty" json:"when,omitempty"`
	Graders []UnifiedGraderEntry `yaml:"graders" json:"graders"`
	Isolate bool                 `yaml:"isolate,omitempty" json:"isolate,omitempty"`
}

// UnifiedGraderConfig is the top-level shape of a unified criteria YAML file.
// A file may define top-level graders, groups, or both. The Source field is
// populated by the loader and is never read from YAML.
type UnifiedGraderConfig struct {
	When    WhenClause           `yaml:"when,omitempty" json:"when,omitempty"`
	Graders []UnifiedGraderEntry `yaml:"graders,omitempty" json:"graders,omitempty"`
	Groups  []UnifiedGraderGroup `yaml:"groups,omitempty" json:"groups,omitempty"`
	Source  string               `yaml:"-" json:"-"`
}

// validTypedKinds enumerates the non-prompt grader types accepted by the
// unified schema. Mirrors validKinds in types.go minus graders.KindPrompt.
// KindPromptReview is NOT included — it's the kind of manually-created
// PromptReviewGrader instances, not a valid criteria-file type.
var validTypedKinds = map[string]bool{
	graders.KindProgram:   true,
	graders.KindTool:      true,
	graders.KindWorkspace: true,
	graders.KindActivity:  true,
}

// IsValidUnifiedType returns true if t is a recognized unified-schema type
// (graders.KindPrompt or any typed kind).
func IsValidUnifiedType(t string) bool {
	if t == graders.KindPrompt {
		return true
	}
	return validTypedKinds[t]
}

// hasChecks reports whether n carries a non-empty YAML payload for checks.
// yaml.v3 returns Kind==0 for absent fields; an explicit empty mapping has Kind!=0.
func hasChecks(n yaml.Node) bool {
	if n.Kind == 0 {
		return false
	}
	// A null scalar (`checks: ~`) is treated as empty.
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
	
	// Reject deprecated types with loud migration errors.
	deprecatedTypes := map[string]string{
		"file":            "workspace",
		"behavior":        "tool or activity",
		"action_sequence": "activity",
		"tool_constraint": "tool",
		"output_check":    "workspace",
		"tool_usage":      "tool",
	}
	if replacement, ok := deprecatedTypes[e.Type]; ok {
		return fmt.Errorf("grader %q: type %q is no longer supported — use %q instead (see docs/graders.md)", tag, e.Type, replacement)
	}
	
	if !IsValidUnifiedType(e.Type) {
		return fmt.Errorf("grader %q: unknown type %q", tag, e.Type)
	}
	if e.Weight < 0 {
		return fmt.Errorf("grader %q: weight must be >= 0, got %f", tag, e.Weight)
	}
	if e.Type == graders.KindPrompt {
		hasPrompt := strings.TrimSpace(e.Prompt) != ""
		// For prompt graders, we need to decode Checks as []string
		var promptChecks []string
		if hasChecks(e.Checks) {
			if err := e.Checks.Decode(&promptChecks); err != nil {
				return fmt.Errorf("grader %q: failed to decode checks as []string: %w", tag, err)
			}
		}
		hasChecksArray := len(promptChecks) > 0
		if !hasPrompt && !hasChecksArray {
			return fmt.Errorf("grader %q: type=prompt requires non-empty prompt or checks", tag)
		}
		if hasPrompt && !hasChecksArray {
			slog.Warn("grader type=prompt has prompt but no checks: will synthesize a single pass/fail check at runtime; prefer adding explicit checks",
				"grader", tag)
		}
		if hasChecksArray {
			for i, c := range promptChecks {
				if strings.TrimSpace(c) == "" {
					return fmt.Errorf("grader %q: checks[%d] must be a non-empty string", tag, i)
				}
			}
		}
		return nil
	}
	// Typed grader — require checks.
	if e.Prompt != "" {
		return fmt.Errorf("grader %q: type=%s must not set prompt", tag, e.Type)
	}
	if !hasChecks(e.Checks) {
		return fmt.Errorf("grader %q: type=%s requires non-empty checks", tag, e.Type)
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

// ─────────────────────────────────────────────────────────────────────────────
// Phase 2: Config-aware when: clause with structured tool filters
// ─────────────────────────────────────────────────────────────────────────────

// StringOrSlice decodes either a YAML scalar or a YAML sequence of scalars
// into a normalized []string. Empty/nil means "no constraint".
//
// JSON/YAML marshalling always emits a list (even for single-element slices),
// providing a stable shape for downstream consumers.
type StringOrSlice []string

// UnmarshalYAML accepts either a YAML scalar string or a sequence of strings.
func (s *StringOrSlice) UnmarshalYAML(node *yaml.Node) error {
	switch node.Kind {
	case yaml.ScalarNode:
		// Only accept string-typed scalars. Reject ints/bools/etc so a
		// stray `language: 42` is a loud error rather than the silent
		// string `"42"`. Untagged plain scalars resolve to !!str for
		// text (yaml.v3 sets the tag during parsing); an empty/null
		// scalar is treated as "no constraint".
		switch node.Tag {
		case "!!null":
			*s = nil
			return nil
		case "", "!!str":
			// fall through
		default:
			return fmt.Errorf("when: field at line %d must be a string or a list of strings, got scalar with tag %s",
				node.Line, node.Tag)
		}
		var v string
		if err := node.Decode(&v); err != nil {
			return err
		}
		if v == "" {
			*s = nil
			return nil
		}
		*s = StringOrSlice{v}
		return nil
	case yaml.SequenceNode:
		var v []string
		if err := node.Decode(&v); err != nil {
			return err
		}
		*s = StringOrSlice(v)
		return nil
	default:
		return fmt.Errorf("when: field at line %d must be a string or a list of strings, got %v",
			node.Line, node.Kind)
	}
}

// Matches reports whether candidate equals any entry (case-insensitive).
// An empty/nil StringOrSlice matches everything.
func (s StringOrSlice) Matches(candidate string) bool {
	if len(s) == 0 {
		return true
	}
	for _, v := range s {
		if strings.EqualFold(v, candidate) {
			return true
		}
	}
	return false
}

// WhenClause narrows grader applicability. All fields AND together.
// An empty WhenClause matches everything.
//
// Every scalar prompt/config field is a StringOrSlice — it accepts either
// a single string or a YAML list of strings, normalized to []string.
// Within a field, entries are OR'd (any-of match). Across fields, AND.
type WhenClause struct {
	// Scalar prompt props (case-insensitive any-of equality).
	Language   StringOrSlice `yaml:"language,omitempty" json:"language,omitempty"`
	Service    StringOrSlice `yaml:"service,omitempty" json:"service,omitempty"`
	Plane      StringOrSlice `yaml:"plane,omitempty" json:"plane,omitempty"`
	Category   StringOrSlice `yaml:"category,omitempty" json:"category,omitempty"`
	SDK        StringOrSlice `yaml:"sdk,omitempty" json:"sdk,omitempty"`
	Difficulty StringOrSlice `yaml:"difficulty,omitempty" json:"difficulty,omitempty"`

	// Scalar config props.
	Generator StringOrSlice `yaml:"generator,omitempty" json:"generator,omitempty"`
	Config    StringOrSlice `yaml:"config,omitempty" json:"config,omitempty"`

	// Structured tool filter; AND across entries.
	Tool []ToolFilter `yaml:"tool,omitempty" json:"tool,omitempty"`
}

// ToolFilter matches one entry in the eval config's resolved tool list.
// Field names mirror ToolCheckRule (graders/types.go:197) deliberately.
type ToolFilter struct {
	Name      string `yaml:"name" json:"name"`
	Source    string `yaml:"source" json:"source"` // skill | mcp | builtin | plugin
	MCPServer string `yaml:"mcp_server,omitempty" json:"mcp_server,omitempty"`
	Negate    bool   `yaml:"negate,omitempty" json:"negate,omitempty"`
}

// MatchContext bundles everything a WhenClause matches against.
// Built once per (prompt, config) pair before evaluating graders.
//
// Props stays map[string]string deliberately: the prompt/config side is 1:1
// (a prompt has one language, one service, etc.). Only the WhenClause side
// accepts lists, because only the gate has an "any of these" notion.
type MatchContext struct {
	// Scalar props derived from prompt frontmatter + eval config.
	// Includes language/service/plane/category/sdk/difficulty/generator/config.
	Props map[string]string

	// Resolved tool list from cfg.Generator.Tools, with type already
	// disambiguated via ToolEntry.ResolvedType().
	Tools []ToolIdentity
}

// ToolIdentity is the canonical (name, source, server) triple.
type ToolIdentity struct {
	Name      string
	Source    string // skill | mcp | builtin | plugin
	MCPServer string // populated only when Source == "mcp"
}

// IsEmpty reports whether the clause has no constraints.
func (w WhenClause) IsEmpty() bool {
	return len(w.Language) == 0 &&
		len(w.Service) == 0 &&
		len(w.Plane) == 0 &&
		len(w.Category) == 0 &&
		len(w.SDK) == 0 &&
		len(w.Difficulty) == 0 &&
		len(w.Generator) == 0 &&
		len(w.Config) == 0 &&
		len(w.Tool) == 0
}

// Matches evaluates the clause against the resolved match context.
// All fields AND together; a constraint passes if it's empty or matches.
func (w WhenClause) Matches(ctx MatchContext) bool {
	// Scalar fields: AND across fields, OR within each field (via StringOrSlice.Matches).
	if !w.Language.Matches(ctx.Props["language"]) {
		return false
	}
	if !w.Service.Matches(ctx.Props["service"]) {
		return false
	}
	if !w.Plane.Matches(ctx.Props["plane"]) {
		return false
	}
	if !w.Category.Matches(ctx.Props["category"]) {
		return false
	}
	if !w.SDK.Matches(ctx.Props["sdk"]) {
		return false
	}
	if !w.Difficulty.Matches(ctx.Props["difficulty"]) {
		return false
	}
	if !w.Generator.Matches(ctx.Props["generator"]) {
		return false
	}
	if !w.Config.Matches(ctx.Props["config"]) {
		return false
	}

	// Tool filters: AND across entries. Every ToolFilter must match some tool.
	for _, filter := range w.Tool {
		matched := false
		for _, tool := range ctx.Tools {
			if matchesToolFilter(filter, tool) {
				matched = true
				break
			}
		}
		// Apply negation: if negate=true, invert the match result.
		if filter.Negate {
			matched = !matched
		}
		if !matched {
			return false
		}
	}

	return true
}

// matchesToolFilter checks if a single ToolFilter matches a ToolIdentity.
func matchesToolFilter(filter ToolFilter, tool ToolIdentity) bool {
	// Name must match (case-insensitive).
	if !strings.EqualFold(filter.Name, tool.Name) {
		return false
	}
	// Source must match (case-insensitive).
	if !strings.EqualFold(filter.Source, tool.Source) {
		return false
	}
	// If MCPServer is set on the filter, it must match.
	if filter.MCPServer != "" {
		if !strings.EqualFold(filter.MCPServer, tool.MCPServer) {
			return false
		}
	}
	return true
}

// mergeWhenClause merges parent and child WhenClause; child REPLACES parent
// for every field (scalars and tool alike). An explicit empty list at child
// level removes the parent's constraint.
func mergeWhenClause(parent, child WhenClause) WhenClause {
	out := parent // value copy

	// Uniform rule: child slice non-nil wins.
	if child.Language != nil {
		out.Language = child.Language
	}
	if child.Service != nil {
		out.Service = child.Service
	}
	if child.Plane != nil {
		out.Plane = child.Plane
	}
	if child.Category != nil {
		out.Category = child.Category
	}
	if child.SDK != nil {
		out.SDK = child.SDK
	}
	if child.Difficulty != nil {
		out.Difficulty = child.Difficulty
	}
	if child.Generator != nil {
		out.Generator = child.Generator
	}
	if child.Config != nil {
		out.Config = child.Config
	}
	if child.Tool != nil {
		out.Tool = child.Tool
	}
	return out
}
