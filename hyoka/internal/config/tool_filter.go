package config

import (
	"fmt"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config/tool"
)

// ToolEntry is an alias for tool.Entry, preserving backward compatibility.
// See the tool sub-package for the type definition and methods.
type ToolEntry = tool.Entry

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
		if e.ResolvedType() != "tool" {
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
	switch entry.ResolvedType() {
	case "tool":
		return nil
	case "mcp":
		mcpType := entry.ResolvedMCPType()
		if mcpType == "remote" {
			if entry.URL == "" {
				return fmt.Errorf("config %q: tools[%d] remote mcp entry missing url", configName, idx)
			}
		} else {
			if entry.Command == "" {
				return fmt.Errorf("config %q: tools[%d] mcp entry missing command", configName, idx)
			}
		}
		if entry.MCPType != "" && entry.MCPType != "local" && entry.MCPType != "remote" {
			return fmt.Errorf("config %q: tools[%d] mcp entry has invalid mcp_type %q", configName, idx, entry.MCPType)
		}
	case "skill":
		if entry.Path == "" && entry.Repo == "" {
			return fmt.Errorf("config %q: tools[%d] skill entry missing path or repo", configName, idx)
		}
		if entry.Source != "" && entry.Source != "local" && entry.Source != "remote" {
			return fmt.Errorf("config %q: tools[%d] skill entry has invalid source %q", configName, idx, entry.Source)
		}
		if entry.SkillDir && entry.Repo != "" {
			return fmt.Errorf("config %q: tools[%d] skill_dir is only valid for local skills", configName, idx)
		}
	default:
		return fmt.Errorf("config %q: tools[%d] has unknown type %q", configName, idx, entry.Type)
	}
	return nil
}
