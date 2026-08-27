// Package graders defines the configuration schema for the pluggable grader
// system that replaces the three-tier criteria approach.
//
// Each grader is a single-concern evaluator (file check, build verification,
// LLM review, etc.) defined in YAML config and executed independently.
// Results are aggregated by weighted scoring, with gate graders acting as
// hard pass/fail constraints (DM3).
//
// See docs/grader-config-schema.md for the full schema specification.
package graders

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// Supported grader kinds.
const (
	KindProgram      = "program"
	KindPrompt       = "prompt"
	KindPromptReview = "prompt_review"
	KindTool         = "tool"      // Canonical tool-perspective grader
	KindWorkspace    = "workspace" // Canonical workspace-delta grader
	KindActivity     = "activity"  // Canonical session-activity grader
)

// validKinds is the set of recognized grader kind values.
var validKinds = map[string]bool{
	KindProgram:      true,
	KindPrompt:       true,
	KindPromptReview: true,
	KindTool:         true,
	KindWorkspace:    true,
	KindActivity:     true,
}

// GraderConfigFile is the top-level YAML structure containing a list of graders.
type GraderConfigFile struct {
	Graders []GraderConfig `yaml:"graders" json:"graders"`
}

// GraderConfig defines a single grader instance in the evaluation pipeline.
type GraderConfig struct {
	Kind   string    `yaml:"kind" json:"kind"`
	Name   string    `yaml:"name" json:"name"`
	Config yaml.Node `yaml:"config" json:"config"`
	Weight float64   `yaml:"weight,omitempty" json:"weight,omitempty"`

	// Deprecated: Gate semantics are no longer enforced. All graders run and contribute
	// their score to the weighted aggregate. Use the consolidated 'tool' grader's check
	// kinds or separate explicit graders to express pass/fail requirements instead.
	Gate bool    `yaml:"gate,omitempty" json:"gate,omitempty"`
	When WhenMap `yaml:"when,omitempty" json:"when,omitempty"`
}

// WhenMap holds property-based applicability conditions.
// All entries must match for the grader to apply (AND logic).
// Matching is case-insensitive.
type WhenMap map[string]string

// Matches returns true if all conditions in the map match the given properties.
// An empty WhenMap matches everything.
func (w WhenMap) Matches(props map[string]string) bool {
	for k, v := range w {
		pv, ok := props[k]
		if !ok || !strings.EqualFold(v, pv) {
			return false
		}
	}
	return true
}

// FileConfig holds configuration for the "file" grader kind.
type FileConfig struct {
	Path      string `yaml:"path" json:"path"`
	Pattern   string `yaml:"pattern,omitempty" json:"pattern,omitempty"`
	MustExist *bool  `yaml:"must_exist,omitempty" json:"must_exist,omitempty"`
}

// ProgramConfig holds configuration for the "program" grader kind.
// The grader runs one or more commands sequentially, each producing a check.
type ProgramConfig struct {
	Checks []ProgramCheck `yaml:"checks" json:"checks"`
}

// ProgramCheck defines a single command check within a program grader.
type ProgramCheck struct {
	Kind    string   `yaml:"kind" json:"kind"`                           // "command" (only supported kind)
	Command string   `yaml:"command" json:"command"`                     // Command to run
	Args    []string `yaml:"args,omitempty" json:"args,omitempty"`       // Command arguments
	Timeout int      `yaml:"timeout,omitempty" json:"timeout,omitempty"` // Timeout in seconds (default: 30)
}

// PromptConfig holds configuration for the "prompt" grader kind.
// Each prompt grader runs exactly one model (DM19).
type PromptConfig struct {
	Model  string `yaml:"model" json:"model"`
	Rubric string `yaml:"rubric" json:"rubric"`
}

// ActivityConfig holds configuration for the "activity" grader kind.
// This grader evaluates session activity using ActionLog, ActionsSummary, and TerminatedBy.
type ActivityConfig struct {
	Checks []ActivityCheck `yaml:"checks" json:"checks"`
}

// ActivityCheck defines a single check within an activity grader.
type ActivityCheck struct {
	Kind string `yaml:"kind" json:"kind"` // One of: turn_limit, action_count, tool_call_count, contains_subsequence, contains_action, excludes_action, terminated_by

	// For turn_limit
	Max *int `yaml:"max,omitempty" json:"max,omitempty"`

	// For action_count, tool_call_count, contains_action (count bounds)
	Min *int `yaml:"min,omitempty" json:"min,omitempty"`

	// For contains_subsequence
	Tools []string `yaml:"tools,omitempty" json:"tools,omitempty"`

	// For contains_action / excludes_action
	//   Type:     event type filter (tool_call, file_read, file_write, bash,
	//             mcp_call, skill, message, reasoning, intent, error, warning,
	//             file_change, truncation, compaction, turn_start, turn_end,
	//             abort). Empty matches any type.
	//   Tool:     tool/skill name filter (matches ev.Tool exactly). Empty = any.
	//   Contains: substring that must appear in the event's text payload
	//             (Output for messages/reasoning/intent/warning/tool output;
	//             Error for error events; Path for file_change).
	//   Excludes: substring that must NOT appear in the event's text payload.
	//   For excludes_action: any matching event causes failure (count must be 0).
	Type     string `yaml:"type,omitempty" json:"type,omitempty"`
	Tool     string `yaml:"tool,omitempty" json:"tool,omitempty"`
	Contains string `yaml:"contains,omitempty" json:"contains,omitempty"`
	Excludes string `yaml:"excludes,omitempty" json:"excludes,omitempty"`

	// Deprecated: retained as YAML aliases so existing fixtures keep parsing.
	// MinCalls maps to Min, MaxCalls maps to Max for contains_action.
	MinCalls *int `yaml:"min_calls,omitempty" json:"min_calls,omitempty"`
	MaxCalls *int `yaml:"max_calls,omitempty" json:"max_calls,omitempty"`

	// For terminated_by
	Equals string   `yaml:"equals,omitempty" json:"equals,omitempty"` // completed | max_actions | max_turns | guardrail | error
	NotIn  []string `yaml:"not_in,omitempty" json:"not_in,omitempty"` // array of forbidden values
}

// OutputCheckConfig holds configuration for the "output_check" grader kind.
// DEPRECATED: Use WorkspaceConfig with type: workspace instead.
// This type is kept for backwards compatibility but will emit a parse error.
type OutputCheckConfig struct {
	MinFiles        int      `yaml:"min_files,omitempty" json:"min_files,omitempty"`
	MaxFiles        int      `yaml:"max_files,omitempty" json:"max_files,omitempty"`
	RequireFiles    []string `yaml:"require_files,omitempty" json:"require_files,omitempty"`
	ForbidFiles     []string `yaml:"forbid_files,omitempty" json:"forbid_files,omitempty"`
	RequireUpdated  []string `yaml:"require_updated,omitempty" json:"require_updated,omitempty"`
	MinBytesPerFile int64    `yaml:"min_bytes_per_file,omitempty" json:"min_bytes_per_file,omitempty"`
	MaxBytesPerFile int64    `yaml:"max_bytes_per_file,omitempty" json:"max_bytes_per_file,omitempty"`
}

// WorkspaceConfig holds configuration for the "workspace" grader kind.
//
// This grader operates on WorkspaceDelta (NewFiles, ModifiedFiles, DeletedFiles)
// and validates workspace state against configured checks. Six check kinds:
//   - require_to_create: path must be in NewFiles
//   - forbidden_to_create: path must NOT be in NewFiles
//   - required_to_update: path must be in ModifiedFiles
//   - required_to_delete: path must be in DeletedFiles
//   - forbidden_to_delete: if files:["*"], no deletions; else specific paths must not be deleted
//   - file: state present (exists on disk + optional size/content checks) or absent (not on disk)
type WorkspaceConfig struct {
	Checks []WorkspaceCheck `yaml:"checks" json:"checks"`
}

// WorkspaceCheck defines a single check within a workspace grader.
type WorkspaceCheck struct {
	Kind     string   `yaml:"kind" json:"kind"`                               // One of: require_to_create, forbidden_to_create, required_to_update, required_to_delete, forbidden_to_delete, file
	Files    []string `yaml:"files,omitempty" json:"files,omitempty"`         // For kinds other than "file"
	Name     string   `yaml:"name,omitempty" json:"name,omitempty"`           // For kind: file
	State    string   `yaml:"state,omitempty" json:"state,omitempty"`         // For kind: file; "present" or "absent"
	MinBytes *int64   `yaml:"min_bytes,omitempty" json:"min_bytes,omitempty"` // For kind: file, state: present
	MaxBytes *int64   `yaml:"max_bytes,omitempty" json:"max_bytes,omitempty"` // For kind: file, state: present
	Contains string   `yaml:"contains,omitempty" json:"contains,omitempty"`   // For kind: file, state: present
	Excludes string   `yaml:"excludes,omitempty" json:"excludes,omitempty"`   // For kind: file, state: present
}

// ToolConfig holds configuration for the "tool" grader kind (canonical).
// The tool grader consolidates behavior, tool_constraint, and tool_usage.
type ToolConfig struct {
	Checks []ToolCheckRule `yaml:"checks" json:"checks"`
}

// ToolCheckRule defines one check within a ToolConfig. See tool_grader.go for semantics.
type ToolCheckRule struct {
	Kind      string   `yaml:"kind" json:"kind"`
	Tool      string   `yaml:"tool,omitempty" json:"tool,omitempty"`             // For tool_used, tool_not_used; optional for an MCP server-wide check
	Group     string   `yaml:"group,omitempty" json:"group,omitempty"`           // For any_from_group, none_from_group
	Except    []string `yaml:"except,omitempty" json:"except,omitempty"`         // Optional exclusion list for group checks
	MinCalls  *int     `yaml:"min_calls,omitempty" json:"min_calls,omitempty"`   // Optional for tool_used
	MaxCalls  *int     `yaml:"max_calls,omitempty" json:"max_calls,omitempty"`   // Optional for tool_used
	Source    string   `yaml:"source,omitempty" json:"source,omitempty"`         // Optional: skill|mcp|builtin (filters by event.Type)
	MCPServer string   `yaml:"mcp_server,omitempty" json:"mcp_server,omitempty"` // Optional: MCP server name (requires source=mcp)
}

// EffectiveWeight returns the grader's weight, defaulting to 1.0 if unset.
func (g *GraderConfig) EffectiveWeight() float64 {
	if g.Weight == 0 {
		return 1.0
	}
	return g.Weight
}

// DecodeConfig decodes the raw YAML config node into the appropriate
// kind-specific config struct. Returns an error if the kind is unknown
// or the config doesn't match the expected schema.
func (g *GraderConfig) DecodeConfig() (any, error) {
	switch g.Kind {
	case KindProgram:
		var c ProgramConfig
		if err := g.Config.Decode(&c); err != nil {
			return nil, fmt.Errorf("decoding program config for %q: %w", g.Name, err)
		}
		return &c, nil
	case KindPrompt:
		var c PromptConfig
		if err := g.Config.Decode(&c); err != nil {
			return nil, fmt.Errorf("decoding prompt config for %q: %w", g.Name, err)
		}
		return &c, nil
	case "activity":
		var c ActivityConfig
		if err := g.Config.Decode(&c); err != nil {
			return nil, fmt.Errorf("decoding activity config for %q: %w", g.Name, err)
		}
		return &c, nil
	case "workspace":
		var c WorkspaceConfig
		if err := g.Config.Decode(&c); err != nil {
			return nil, fmt.Errorf("decoding workspace config for %q: %w", g.Name, err)
		}
		return &c, nil
	case KindTool:
		var c ToolConfig
		if err := g.Config.Decode(&c); err != nil {
			return nil, fmt.Errorf("decoding tool config for %q: %w", g.Name, err)
		}
		return &c, nil
	default:
		return nil, fmt.Errorf("unknown grader kind %q for %q", g.Kind, g.Name)
	}
}

// Parse decodes YAML bytes into a GraderConfigFile and validates it.
func Parse(data []byte) (*GraderConfigFile, error) {
	var gcf GraderConfigFile
	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true)
	if err := dec.Decode(&gcf); err != nil {
		return nil, fmt.Errorf("parsing grader config: %w", err)
	}
	if err := Validate(&gcf); err != nil {
		return nil, err
	}
	return &gcf, nil
}

// LoadFile loads and parses a single grader config YAML file.
func LoadFile(path string) (*GraderConfigFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading grader config %s: %w", path, err)
	}
	gcf, err := Parse(data)
	if err != nil {
		return nil, fmt.Errorf("in %s: %w", path, err)
	}
	return gcf, nil
}

// LoadDir loads all grader config YAML files from a directory tree.
func LoadDir(dir string) (*GraderConfigFile, error) {
	merged := &GraderConfigFile{}

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
		gcf, err := LoadFile(path)
		if err != nil {
			return err
		}
		merged.Graders = append(merged.Graders, gcf.Graders...)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walking grader config directory %s: %w", dir, err)
	}
	return merged, nil
}

// Validate checks a GraderConfigFile for structural correctness.
func Validate(gcf *GraderConfigFile) error {
	if len(gcf.Graders) == 0 {
		return fmt.Errorf("no graders defined")
	}

	names := make(map[string]bool, len(gcf.Graders))
	for i, g := range gcf.Graders {
		if g.Name == "" {
			return fmt.Errorf("grader at index %d: name is required", i)
		}
		if g.Kind == "" {
			return fmt.Errorf("grader %q: kind is required", g.Name)
		}
		if !validKinds[g.Kind] {
			return fmt.Errorf("grader %q: unknown kind %q", g.Name, g.Kind)
		}
		if names[g.Name] {
			return fmt.Errorf("grader %q: duplicate name", g.Name)
		}
		names[g.Name] = true
		if g.Weight < 0 || g.Weight > 1 {
			return fmt.Errorf("grader %q: weight must be between 0.0 and 1.0, got %f", g.Name, g.Weight)
		}
		if _, err := g.DecodeConfig(); err != nil {
			return err
		}
	}
	return nil
}

// ApplicableGraders filters graders by the given prompt properties,
// returning only those whose When conditions match.
func ApplicableGraders(graders []GraderConfig, props map[string]string) []GraderConfig {
	var applicable []GraderConfig
	for _, g := range graders {
		if g.When.Matches(props) {
			applicable = append(applicable, g)
		}
	}
	return applicable
}
