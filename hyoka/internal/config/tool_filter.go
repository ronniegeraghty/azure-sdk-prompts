package config

import "fmt"

// ToolEntry represents a tool, MCP server, or skill with optional conditions.
// When the When map is empty, the entry is unconditionally included.
// When the When map has entries, all key-value pairs must match the
// prompt's properties for the entry to be included.
type ToolEntry struct {
	Name     string            `yaml:"name" json:"name"`
	Type     string            `yaml:"type,omitempty" json:"type,omitempty"` // "tool" (default), "mcp", "skill"
	When     map[string]string `yaml:"when,omitempty" json:"when,omitempty"`
	AlwaysOn bool              `yaml:"always_on,omitempty" json:"always_on,omitempty"`
	Pairwise string            `yaml:"pairwise,omitempty" json:"pairwise,omitempty"` // "off", "shallow" (default), "deep"
	// MCP-specific fields
	Command  string   `yaml:"command,omitempty" json:"command,omitempty"`
	Args     []string `yaml:"args,omitempty" json:"args,omitempty"`
	MCPTools []string `yaml:"mcp_tools,omitempty" json:"mcp_tools,omitempty"`
	// Skill-specific fields
	Source string `yaml:"source,omitempty" json:"source,omitempty"` // "local" or "remote"
	Path   string `yaml:"path,omitempty" json:"path,omitempty"`
	Repo   string `yaml:"repo,omitempty" json:"repo,omitempty"`
}

func resolvedToolType(entry ToolEntry) string {
	if entry.Type == "" {
		return "tool"
	}
	return entry.Type
}

// ResolvedType returns the normalized entry type ("tool" by default).
func (e ToolEntry) ResolvedType() string {
	return resolvedToolType(e)
}

// ResolvedPairwise returns the effective pairwise setting ("shallow" default).
// AlwaysOn overrides any explicit setting to "off".
func (e ToolEntry) ResolvedPairwise() string {
	if e.AlwaysOn {
		return "off"
	}
	if e.Pairwise == "" {
		return "shallow"
	}
	return e.Pairwise
}

// SkillSource returns the normalized skill source ("local" or "remote").
// If Source is unset, it infers from Path/Repo when possible.
func (e ToolEntry) SkillSource() string {
	if e.Source != "" {
		return e.Source
	}
	if e.Path != "" {
		return "local"
	}
	if e.Repo != "" {
		return "remote"
	}
	return ""
}

// ResolveTools evaluates tool entries against prompt properties and returns
// the names of tools whose conditions are satisfied. An empty entries slice
// returns nil (meaning "all default tools"). Duplicate names are deduplicated
// while preserving first-seen order.
func ResolveTools(entries []ToolEntry, properties map[string]string) []string {
	if len(entries) == 0 {
		return nil
	}
	seen := make(map[string]bool, len(entries))
	var resolved []string
	for _, e := range entries {
		if resolvedToolType(e) != "tool" {
			continue
		}
		if matchesWhen(e.When, properties) && !seen[e.Name] {
			seen[e.Name] = true
			resolved = append(resolved, e.Name)
		}
	}
	return resolved
}

// matchesWhen returns true when every key-value pair in when matches the
// properties map. An empty when map always matches.
func matchesWhen(when map[string]string, props map[string]string) bool {
	for k, v := range when {
		if props[k] != v {
			return false
		}
	}
	return true
}

// validateToolEntry checks that a ToolEntry has valid fields.
func validateToolEntry(entry ToolEntry, configName string, idx int) error {
	if entry.Name == "" {
		return fmt.Errorf("config %q: tools[%d] missing name", configName, idx)
	}
	if entry.Pairwise != "" && entry.Pairwise != "off" && entry.Pairwise != "shallow" && entry.Pairwise != "deep" {
		return fmt.Errorf("config %q: tools[%d] has invalid pairwise %q", configName, idx, entry.Pairwise)
	}
	switch resolvedToolType(entry) {
	case "tool":
		return nil
	case "mcp":
		if entry.Command == "" {
			return fmt.Errorf("config %q: tools[%d] mcp entry missing command", configName, idx)
		}
	case "skill":
		if entry.Path == "" && entry.Repo == "" {
			return fmt.Errorf("config %q: tools[%d] skill entry missing path or repo", configName, idx)
		}
		if entry.Source != "" && entry.Source != "local" && entry.Source != "remote" {
			return fmt.Errorf("config %q: tools[%d] skill entry has invalid source %q", configName, idx, entry.Source)
		}
	default:
		return fmt.Errorf("config %q: tools[%d] has unknown type %q", configName, idx, entry.Type)
	}
	return nil
}

// validateSkillRef checks that a SkillRef has valid fields.
func validateSkillRef(ref SkillRef, configName string, idx int) error {
	if ref.Type != "" && ref.Type != "local" && ref.Type != "remote" {
		return fmt.Errorf("config %q: skills[%d] has invalid type %q (must be \"local\" or \"remote\")", configName, idx, ref.Type)
	}
	resolvedType := ref.Type
	if resolvedType == "" {
		if ref.Path != "" {
			resolvedType = "local"
		} else if ref.Repo != "" {
			resolvedType = "remote"
		}
	}
	switch resolvedType {
	case "local":
		if ref.Path == "" {
			return fmt.Errorf("config %q: skills[%d] local skill missing path", configName, idx)
		}
	case "remote":
		if ref.Repo == "" {
			return fmt.Errorf("config %q: skills[%d] remote skill missing repo", configName, idx)
		}
	default:
		return fmt.Errorf("config %q: skills[%d] missing type, path, or repo", configName, idx)
	}
	return nil
}
