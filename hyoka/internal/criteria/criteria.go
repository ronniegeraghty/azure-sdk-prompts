// Package criteria implements a grader-config-based evaluation criteria system.
//
// Criteria YAML files define grader configs with hierarchical when conditions:
//   - File-level when: applies to entire file
//   - Group-level when: applies to graders within a group
//   - Grader-level when: applies to individual grader
//
// At eval time, matching grader configs are collected and merged with any
// prompt-specific criteria to form the final evaluation rubric.
package criteria

import (
"bytes"
"fmt"
"log/slog"
"os"
"path/filepath"
"strings"

"gopkg.in/yaml.v3"
)

// GraderEntry defines a single grader with its evaluation prompt and weight.
// Supports hierarchical when conditions for prompt-based matching.
//
// Isolate, when true and the engine is run with --review-mode isolated, causes
// this grader to be reviewed in a dedicated Copilot session rather than sharing
// a session with other criteria. Has no effect in the default combined mode.
type GraderEntry struct {
Name    string            `yaml:"name" json:"name"`
Weight  float64           `yaml:"weight" json:"weight"`
Prompt  string            `yaml:"prompt" json:"prompt"`
When    map[string]string `yaml:"when,omitempty" json:"when,omitempty"` // Grader-level when conditions
Isolate bool              `yaml:"isolate,omitempty" json:"isolate,omitempty"`
}

// GraderGroup is a named collection of graders with optional when conditions.
// Groups allow hierarchical when matching: group-level conditions apply to
// all graders in the group unless overridden by grader-level conditions.
//
// Isolate, when true and the engine is run with --review-mode isolated, causes
// the entire group to be reviewed in a single dedicated Copilot session
// (separate from other criteria but shared across the group's graders).
// Per-grader Isolate within an isolated group is ignored — the group-level
// flag wins. Has no effect in the default combined mode.
type GraderGroup struct {
Name    string            `yaml:"name,omitempty" json:"name,omitempty"`
When    map[string]string `yaml:"when,omitempty" json:"when,omitempty"`
Graders []GraderEntry     `yaml:"graders" json:"graders"`
Isolate bool              `yaml:"isolate,omitempty" json:"isolate,omitempty"`
}

// GraderConfig is a collection of graders with conditions for when they apply.
// Supports three levels of when conditions:
//   1. File-level when (applies to entire file)
//   2. Group-level when (applies to graders within a group)
//   3. Grader-level when (applies to individual grader)
//
// Resolution: grader-level when overrides group-level, which overrides file-level.
// A grader matches only if ALL applicable when conditions match the prompt properties.
type GraderConfig struct {
When    map[string]string `yaml:"when,omitempty" json:"when,omitempty"` // File-level when
Graders []GraderEntry     `yaml:"graders,omitempty" json:"graders,omitempty"` // Top-level graders
Groups  []GraderGroup     `yaml:"groups,omitempty" json:"groups,omitempty"` // Grouped graders
Source  string            `yaml:"-" json:"-"`
}

// matchesWhen returns true when every key-value pair in when matches the
// properties map (case-insensitive values). An empty when map always matches.
func matchesWhen(when map[string]string, props map[string]string) bool {
for k, v := range when {
if !strings.EqualFold(props[k], v) {
return false
}
}
return true
}

// mergeWhen merges hierarchical when conditions from parent (file or group) and
// child (group or grader) levels. Child conditions override parent conditions for
// the same key. Returns the merged when map.
func mergeWhen(parent, child map[string]string) map[string]string {
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

// LoadDir loads all grader config YAML files from a directory tree.
func LoadDir(dir string) ([]GraderConfig, error) {
var configs []GraderConfig

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

gc, err := loadFile(path)
if err != nil {
slog.Warn("Skipping invalid grader config file", "path", path, "error", err)
return nil
}
gc.Source = path
configs = append(configs, *gc)
slog.Debug("Loaded grader config", "path", path, "grader_count", len(gc.Graders))
return nil
})
if err != nil {
return nil, fmt.Errorf("walking criteria directory %s: %w", dir, err)
}
return configs, nil
}

func loadFile(path string) (*GraderConfig, error) {
data, err := os.ReadFile(path)
if err != nil {
return nil, err
}
var gc GraderConfig
dec := yaml.NewDecoder(bytes.NewReader(data))
dec.KnownFields(true)
if err := dec.Decode(&gc); err != nil {
return nil, fmt.Errorf("parsing %s: %w", path, err)
}
if len(gc.Graders) == 0 && len(gc.Groups) == 0 {
return nil, fmt.Errorf("%s: no graders or groups defined", path)
}
return &gc, nil
}

// MatchingGraders returns all grader entries from configs whose when-conditions
// match the given prompt properties. Respects hierarchical when resolution:
// file-level → group-level → grader-level (most specific wins).
func MatchingGraders(configs []GraderConfig, props map[string]string) []GraderEntry {
	matched := MatchingGradersWithIsolation(configs, props)
	out := make([]GraderEntry, 0, len(matched))
	for _, m := range matched {
		out = append(out, m.Entry)
	}
	return out
}

// FormatGraders formats a list of grader entries as a text block suitable for
// injection into a review prompt.
func FormatGraders(graders []GraderEntry) string {
if len(graders) == 0 {
return ""
}
var b strings.Builder
for i, g := range graders {
fmt.Fprintf(&b, "%d. **%s**", i+1, g.Name)
if g.Prompt != "" {
fmt.Fprintf(&b, " — %s", strings.TrimSpace(g.Prompt))
}
b.WriteString("\n")
}
return b.String()
}

// MergeCriteria combines attribute-matched grader entries with prompt-specific
// criteria text. Returns the merged string suitable for passing to the reviewer.
func MergeCriteria(graders []GraderEntry, promptCriteria string) string {
parts := make([]string, 0, 2)

formatted := FormatGraders(graders)
if formatted != "" {
parts = append(parts, "### Attribute-Matched Criteria\n\n"+formatted)
}

promptCriteria = strings.TrimSpace(promptCriteria)
if promptCriteria != "" {
parts = append(parts, "### Prompt-Specific Criteria\n\n"+promptCriteria)
}

return strings.Join(parts, "\n")
}
