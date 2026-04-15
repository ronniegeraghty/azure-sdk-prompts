package tool

import (
	"github.com/ronniegeraghty/hyoka/hyoka/internal/plugin"
)

// ConvertPluginEntry converts a plugin.ToolEntry to a tool.Entry.
func ConvertPluginEntry(entry plugin.ToolEntry) Entry {
	return Entry{
		Name:     entry.Name,
		Type:     entry.Type,
		Source:   entry.Source,
		Path:     entry.Path,
		Repo:     entry.Repo,
		Command:  entry.Command,
		Args:     CloneStringSlice(entry.Args),
		MCPTools: CloneStringSlice(entry.MCPTools),
	}
}

// AppendEntries appends extras to existing, deduplicating by type:name key.
func AppendEntries(existing []Entry, extras []Entry) []Entry {
	if len(extras) == 0 {
		return existing
	}
	seen := make(map[string]bool, len(existing))
	for _, entry := range existing {
		key := entry.ResolvedType() + ":" + entry.Name
		seen[key] = true
	}
	for _, entry := range extras {
		key := entry.ResolvedType() + ":" + entry.Name
		if seen[key] {
			continue
		}
		seen[key] = true
		existing = append(existing, entry)
	}
	return existing
}

// CloneStringSlice returns a copy of the input slice, or nil if empty.
func CloneStringSlice(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	clone := make([]string, len(values))
	copy(clone, values)
	return clone
}
