// Package config provides configuration loading and parsing for the evaluation tool.
package config

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// GeneratorConfig holds all configuration for the generator agent.
type GeneratorConfig struct {
	Model         string      `yaml:"model,omitempty" json:"model,omitempty"`
	Models        []string    `yaml:"models,omitempty" json:"models,omitempty"`
	SystemPrompt  string      `yaml:"system_prompt,omitempty" json:"system_prompt,omitempty"`
	Tools         []ToolEntry `yaml:"tools,omitempty" json:"tools,omitempty"`
	ExcludedTools []string    `yaml:"excluded_tools,omitempty" json:"excluded_tools,omitempty"`
}

// ResolveModels returns the generator models to use. Models takes precedence
// over Model when both are set.
func (g *GeneratorConfig) ResolveModels() []string {
	if g == nil {
		return nil
	}
	if len(g.Models) > 0 {
		return g.Models
	}
	if g.Model != "" {
		return []string{g.Model}
	}
	return nil
}

// ReviewerConfig holds all configuration for the review/grading plane.
type ReviewerConfig struct {
	Models       []string    `yaml:"models,omitempty" json:"models,omitempty"`
	SystemPrompt string      `yaml:"system_prompt,omitempty" json:"system_prompt,omitempty"`
	Tools        []ToolEntry `yaml:"tools,omitempty" json:"tools,omitempty"`
}

// SessionLimits configures per-config guardrail limits for evaluation sessions.
// Zero values are ignored and fall back to engine-level defaults.
type SessionLimits struct {
	MaxTurns          int    `yaml:"max_turns,omitempty" json:"max_turns,omitempty"`
	MaxFiles          int    `yaml:"max_files,omitempty" json:"max_files,omitempty"`
	MaxSessionActions int    `yaml:"max_session_actions,omitempty" json:"max_session_actions,omitempty"`
	ToolLoadCeiling   string `yaml:"tool_load_ceiling,omitempty" json:"tool_load_ceiling,omitempty"`
}

// ToolConfig represents a single evaluation configuration.
//
// Note (schema retire, 2026-04-24): the top-level `plugins:` field has been
// removed. Plugins are now declared as tool entries under `generator.tools`
// (or `reviewer.tools`) with `type: plugin` and optional `source: local|remote`.
// Any config that still carries a top-level `plugins:` field is rejected at
// Parse time with a migration-hint error.
type ToolConfig struct {
	Name        string           `yaml:"name" json:"name"`
	Description string           `yaml:"description" json:"description"`
	Generator   *GeneratorConfig `yaml:"generator,omitempty" json:"generator,omitempty"`
	Reviewer    *ReviewerConfig  `yaml:"reviewer,omitempty" json:"reviewer,omitempty"`
	Limits      *SessionLimits   `yaml:"limits,omitempty" json:"limits,omitempty"`
}

// Normalize ensures non-nil Generator and Reviewer sub-configs so downstream
// code can append to Tools without nil-checks. Idempotent.
func (tc *ToolConfig) Normalize() {
	if tc.Generator == nil {
		tc.Generator = &GeneratorConfig{}
	}
	if tc.Reviewer == nil {
		tc.Reviewer = &ReviewerConfig{}
	}
}

// EffectiveGeneratorSkills returns the generator's skill list from the normalized config.
func (tc *ToolConfig) EffectiveGeneratorSkills() []ToolEntry {
	if tc.Generator != nil {
		return tc.Generator.Tools
	}
	return nil
}

// ConfigFile represents the top-level config file structure.
type ConfigFile struct {
	// PromptDirectory is an optional path that overrides the default prompt
	// discovery (.hyoka/prompts → ./prompts → ../prompts). When set in a
	// config YAML loaded from disk, relative paths are resolved against the
	// containing config file's directory by Load/LoadDir.
	PromptDirectory string `yaml:"prompt_directory,omitempty" json:"prompt_directory,omitempty"`
	// ToolVersionOverride pins all tool entries from a given repository
	// (matched by Entry.Repo in owner/repo format) to a specific version.
	// Keys should be in "owner/repo" format (e.g., "microsoft/skills").
	// Leading "github.com/" prefixes are normalized away during lookup.
	// The version is forwarded to the registered Fetcher (for the git fetcher,
	// it becomes a git ref). Per-entry `version:` set directly on a tool entry
	// takes precedence over this map. Empty map (or absent field) means "no
	// overrides" — fetcher defaults are used everywhere.
	ToolVersionOverride map[string]string `yaml:"tool_version_override,omitempty" json:"tool_version_override,omitempty"`
	Configs             []ToolConfig      `yaml:"configs"`
}

// normalizeRepoKey normalizes a repository key by trimming the leading
// "github.com/" prefix, if present. This allows users to write either
// "microsoft/skills" or "github.com/microsoft/skills" and have them match.
func normalizeRepoKey(s string) string {
	return strings.TrimPrefix(s, "github.com/")
}

// validateOverrideKeys validates that all keys in tool_version_override are
// in the new repo-keyed format (owner/repo). Returns a migration-hint error
// if old-shape (name-keyed) entries are detected, or a validation error if
// keys are malformed.
func validateOverrideKeys(overrides map[string]string) error {
	for k := range overrides {
		normalized := normalizeRepoKey(k)
		// Detect old-shape: keys without a slash
		if !strings.Contains(normalized, "/") {
			return fmt.Errorf(
				"tool_version_override now keys by repo (e.g. \"microsoft/skills\"), not by tool name.\n"+
					"Found name-shaped key %q. Migration:\n"+
					"  - Replace each tool-name key with the repo it points to.\n"+
					"  - If multiple tools shared the same repo, collapse them to one entry.\n"+
					"See docs/configuration.md → \"Tool Versioning\" for examples.",
				k,
			)
		}
		// Validate owner/repo format: exactly one slash, non-empty parts
		parts := strings.Split(normalized, "/")
		if len(parts) != 2 {
			return fmt.Errorf("tool_version_override: key %q is not in \"owner/repo\" format", k)
		}
		if parts[0] == "" || parts[1] == "" {
			return fmt.Errorf("tool_version_override: key %q has empty owner or repo", k)
		}
	}
	return nil
}

// ApplyVersionOverrides applies cf.ToolVersionOverride to every tool entry
// in every config (Generator and Reviewer). Entries with a non-empty
// Version field are left untouched (per-entry pin wins). Keys are matched
// against Entry.Repo (normalized to owner/repo format). Idempotent.
func (cf *ConfigFile) ApplyVersionOverrides() {
	if cf == nil || len(cf.ToolVersionOverride) == 0 {
		return
	}
	// Normalize override keys once
	normalized := make(map[string]string, len(cf.ToolVersionOverride))
	for k, v := range cf.ToolVersionOverride {
		normalized[normalizeRepoKey(k)] = v
	}
	// Track which override keys were used
	usedKeys := make(map[string]bool)
	
	apply := func(entries []ToolEntry) {
		for i := range entries {
			if entries[i].Version != "" {
				continue // per-entry pin wins
			}
			if entries[i].Repo == "" {
				continue // skip local skills, MCPs
			}
			normalizedRepo := normalizeRepoKey(entries[i].Repo)
			if v, ok := normalized[normalizedRepo]; ok && v != "" {
				entries[i].Version = v
				usedKeys[normalizedRepo] = true
			}
		}
	}
	for i := range cf.Configs {
		if cf.Configs[i].Generator != nil {
			apply(cf.Configs[i].Generator.Tools)
		}
		if cf.Configs[i].Reviewer != nil {
			apply(cf.Configs[i].Reviewer.Tools)
		}
	}
	
	// Warn about unused override keys (override references a repo not present in any tool entry)
	for k := range normalized {
		if !usedKeys[k] {
			slog.Warn("tool_version_override key matches no tool entries in this config set", "repo", k)
		}
	}
}

// Load reads and parses a configuration file from the given path.
func Load(path string) (*ConfigFile, error) {
	slog.Debug("Loading config file", "path", path)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}
	cfg, err := Parse(data)
	if err != nil {
		return nil, err
	}
	cfg.ApplyVersionOverrides()
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	// Resolve a relative prompt_directory against the config file's directory
	// so a config that sits under .hyoka/configs/ can reference ../my-prompts.
	if cfg.PromptDirectory != "" && !filepath.IsAbs(cfg.PromptDirectory) {
		cfg.PromptDirectory = filepath.Join(filepath.Dir(path), cfg.PromptDirectory)
	}
	return cfg, nil
}

// LoadDir reads all .yaml files in a directory and merges their configs.
// This allows splitting configs across multiple files (e.g., baseline.yaml, azure-mcp.yaml).
func LoadDir(dir string) (*ConfigFile, error) {
	slog.Debug("Loading config directory", "dir", dir)
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("reading config directory %s: %w", dir, err)
	}

	merged := &ConfigFile{}
	nameSource := make(map[string]string) // config name → source filename
	var promptDirSource string             // filename that first set prompt_directory
	for _, e := range entries {
		if e.IsDir() || (filepath.Ext(e.Name()) != ".yaml" && filepath.Ext(e.Name()) != ".yml") {
			continue
		}
		cf, err := Load(filepath.Join(dir, e.Name()))
		if err != nil {
			return nil, fmt.Errorf("loading %s: %w", e.Name(), err)
		}
		if cf.PromptDirectory != "" {
			if merged.PromptDirectory == "" {
				merged.PromptDirectory = cf.PromptDirectory
				promptDirSource = e.Name()
			} else if merged.PromptDirectory != cf.PromptDirectory {
				return nil, fmt.Errorf(
					"conflicting prompt_directory: %q in %s vs %q in %s",
					merged.PromptDirectory, promptDirSource,
					cf.PromptDirectory, e.Name(),
				)
			}
		}
		for _, c := range cf.Configs {
			if prev, ok := nameSource[c.Name]; ok {
				return nil, fmt.Errorf("duplicate config name %q found in files %s and %s", c.Name, prev, e.Name())
			}
			nameSource[c.Name] = e.Name()
		}
		merged.Configs = append(merged.Configs, cf.Configs...)
		// Merge tool_version_override maps. Conflicting values across files
		// are an error — silently last-write-wins would be a footgun.
		for k, v := range cf.ToolVersionOverride {
			if merged.ToolVersionOverride == nil {
				merged.ToolVersionOverride = make(map[string]string)
			}
			if existing, ok := merged.ToolVersionOverride[k]; ok && existing != v {
				return nil, fmt.Errorf("conflicting tool_version_override for repo %q: %q vs %q (in %s)", k, existing, v, e.Name())
			}
			merged.ToolVersionOverride[k] = v
		}
	}

	if len(merged.Configs) == 0 {
		return nil, fmt.Errorf("no configs found in %s", dir)
	}
	return merged, nil
}

// Parse parses configuration from YAML bytes.
func Parse(data []byte) (*ConfigFile, error) {
	// Pre-scan for retired schema fields so users get a migration hint
	// instead of a terse "field plugins not found" yaml error. Pre-1.0
	// there is no back-compat for the top-level `plugins:` field — it
	// must be expressed as `type: plugin` entries under
	// generator.tools (or reviewer.tools) instead.
	if err := rejectRetiredPluginsField(data); err != nil {
		return nil, err
	}
	var cfg ConfigFile
	dec := yaml.NewDecoder(bytes.NewReader(data))
	dec.KnownFields(true)
	if err := dec.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("parsing config YAML: %w", err)
	}
	// Validate tool_version_override keys are in the new repo-keyed format
	if err := validateOverrideKeys(cfg.ToolVersionOverride); err != nil {
		return nil, err
	}
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	for _, c := range cfg.Configs {
		models := c.Generator.ResolveModels()
		slog.Info("Config loaded", "name", c.Name, "models", strings.Join(models, ","))
	}
	return &cfg, nil
}

// rejectRetiredPluginsField scans the YAML for a `plugins:` key at the
// top level of any config entry (i.e. sibling of `generator`/`reviewer`)
// and returns a migration-hint error when one is found. The top-level
// `plugins:` field was retired in favor of `type: plugin` entries under
// `generator.tools` / `reviewer.tools`.
func rejectRetiredPluginsField(data []byte) error {
	var raw struct {
		Configs []map[string]yaml.Node `yaml:"configs"`
	}
	if err := yaml.Unmarshal(data, &raw); err != nil {
		// Let the typed decode surface this — it will produce a proper
		// line-numbered yaml error.
		return nil
	}
	for _, c := range raw.Configs {
		if _, ok := c["plugins"]; ok {
			name, _ := c["name"]
			return fmt.Errorf(
				"config %q: the top-level `plugins:` field has been retired. "+
					"Move each plugin under `generator.tools` (or `reviewer.tools` if needed) as:\n"+
					"    - name: <plugin-name>\n"+
					"      type: plugin\n"+
					"      source: remote   # or 'local' for ./.hyoka/plugins/<name>/plugin.yaml",
				nodeValue(name),
			)
		}
	}
	return nil
}

func nodeValue(n yaml.Node) string {
	if n.Kind == yaml.ScalarNode {
		return n.Value
	}
	return "<unnamed>"
}

// Validate checks all configs for required fields and constraint violations.
func (cf *ConfigFile) Validate() error {
	if len(cf.Configs) == 0 {
		return fmt.Errorf("no configs defined")
	}
	namesSeen := make(map[string]int, len(cf.Configs))
	for i, c := range cf.Configs {
		if c.Name == "" {
			return fmt.Errorf("config at index %d has no name", i)
		}
		// Generator with a model is required for every config.
		if c.Generator == nil {
			return fmt.Errorf("config %q: generator.model or generator.models is required", c.Name)
		}
		models := c.Generator.ResolveModels()
		if len(models) == 0 {
			return fmt.Errorf("config %q: generator.model or generator.models is required", c.Name)
		}
		seenModels := make(map[string]bool, len(models))
		for _, model := range models {
			if model == "" {
				return fmt.Errorf("config %q: generator model must not be empty", c.Name)
			}
			if seenModels[model] {
				return fmt.Errorf("config %q: duplicate generator model %q", c.Name, model)
			}
			seenModels[model] = true
		}
		if prev, ok := namesSeen[c.Name]; ok {
			return fmt.Errorf("duplicate config name %q at index %d and %d", c.Name, prev, i)
		}
		namesSeen[c.Name] = i
		if c.Generator != nil {
			for j, te := range c.Generator.Tools {
				if err := validateToolEntry(te, c.Name, j); err != nil {
					return err
				}
			}
		}
		if c.Reviewer != nil {
			for j, te := range c.Reviewer.Tools {
				if err := validateToolEntry(te, c.Name, j); err != nil {
					return err
				}
			}
			// Check for duplicate reviewer models
			seen := make(map[string]bool, len(c.Reviewer.Models))
			for _, rm := range c.Reviewer.Models {
				if seen[rm] {
					return fmt.Errorf("config %q: duplicate reviewer model %q", c.Name, rm)
				}
				seen[rm] = true
			}
		}
		// Reject negative session limits — they bypass guardrails.
		if c.Limits != nil {
			if c.Limits.MaxTurns < 0 {
				return fmt.Errorf("config %q: limits.max_turns must not be negative", c.Name)
			}
			if c.Limits.MaxFiles < 0 {
				return fmt.Errorf("config %q: limits.max_files must not be negative", c.Name)
			}
			if c.Limits.MaxSessionActions < 0 {
				return fmt.Errorf("config %q: limits.max_session_actions must not be negative", c.Name)
			}
		}
	}
	return nil
}

// GetConfig returns a config by name, or an error if not found.
func (cf *ConfigFile) GetConfig(name string) (*ToolConfig, error) {
	for i := range cf.Configs {
		if cf.Configs[i].Name == name {
			return &cf.Configs[i], nil
		}
	}
	return nil, fmt.Errorf("config %q not found", name)
}

// GetConfigs returns configs matching the given names. If names is empty, returns all.
func (cf *ConfigFile) GetConfigs(names []string) ([]ToolConfig, error) {
	if len(names) == 0 {
		return cf.Configs, nil
	}
	nameSet := make(map[string]bool, len(names))
	for _, n := range names {
		nameSet[n] = true
	}
	var result []ToolConfig
	for _, c := range cf.Configs {
		if nameSet[c.Name] {
			result = append(result, c)
			delete(nameSet, c.Name)
		}
	}
	if len(nameSet) > 0 {
		var missing []string
		for n := range nameSet {
			missing = append(missing, n)
		}
		return nil, fmt.Errorf("configs not found: %v", missing)
	}
	return result, nil
}

// InstallSkillsAndPlugins is a no-op retained for backward compatibility.
// Skill/plugin resolution now happens at eval time via ValidateAndExpand
// (plugins) and the git-clone fetcher (remote skills). This function
// remains callable so legacy call sites keep compiling without requiring
// a coordinated removal.
func InstallSkillsAndPlugins(configs []ToolConfig) error {
	// No-op: plugin + skill resolution is lazy.
	return nil
}
