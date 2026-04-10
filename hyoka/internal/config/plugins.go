package config

import (
	"log/slog"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/plugin"
)

// ExpandPlugins loads plugins from dir and appends their tool entries to configs.
// Missing plugin directories are skipped silently.
func (cf *ConfigFile) ExpandPlugins(dir string) error {
	if dir == "" {
		return nil
	}
	reg := plugin.NewRegistry()
	if err := reg.LoadDir(dir); err != nil {
		return err
	}
	if reg.Count() == 0 {
		return nil
	}
	for i := range cf.Configs {
		if err := cf.Configs[i].ExpandPlugins(reg); err != nil {
			return err
		}
	}
	slog.Debug("Expanded plugins into tool entries", "plugins", reg.Count())
	return nil
}

// ExpandPlugins appends plugin-derived tool entries to generator/reviewer tools.
func (c *ToolConfig) ExpandPlugins(reg *plugin.Registry) error {
	if reg == nil || len(c.Plugins) == 0 {
		return nil
	}
	var entries []ToolEntry
	for _, name := range c.Plugins {
		p, err := reg.Get(name)
		if err != nil {
			return err
		}
		for _, entry := range p.ToToolEntries() {
			entries = append(entries, convertPluginToolEntry(entry))
		}
	}
	if len(entries) == 0 {
		return nil
	}
	if c.Generator != nil {
		c.Generator.Tools = appendToolEntries(c.Generator.Tools, entries)
	}
	if c.Reviewer != nil {
		c.Reviewer.Tools = appendToolEntries(c.Reviewer.Tools, entries)
	}
	return nil
}

func convertPluginToolEntry(entry plugin.ToolEntry) ToolEntry {
	return ToolEntry{
		Name:     entry.Name,
		Type:     entry.Type,
		Source:   entry.Source,
		Path:     entry.Path,
		Repo:     entry.Repo,
		Command:  entry.Command,
		Args:     cloneStringSlice(entry.Args),
		MCPTools: cloneStringSlice(entry.MCPTools),
	}
}

func appendToolEntries(existing []ToolEntry, extras []ToolEntry) []ToolEntry {
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

func cloneStringSlice(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	clone := make([]string, len(values))
	copy(clone, values)
	return clone
}

func resolvePluginsDir() string {
	proj := DiscoverFromCWD()
	candidates := ResolveCandidates(proj, "plugins", "./plugins", "../plugins")
	if len(candidates) > 0 {
		return candidates[0]
	}
	return "./plugins"
}
