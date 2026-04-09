package config

import (
	"log/slog"

	"github.com/ronniegeraghty/hyoka/internal/plugin"
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
// It first checks the plugin registry (plugins/ directory), then falls back to
// resolving installed Copilot CLI plugins from ~/.copilot/installed-plugins/.
// Missing plugins are logged as warnings and skipped — this allows configs that
// reference optional plugins to load gracefully for contributors who don't have
// them installed.
func (c *ToolConfig) ExpandPlugins(reg *plugin.Registry) error {
	if len(c.Plugins) == 0 {
		return nil
	}
	var entries []ToolEntry
	for _, name := range c.Plugins {
		// Try plugin registry first (plugins/ directory YAML definitions)
		if reg != nil {
			if p, err := reg.Get(name); err == nil {
				for _, entry := range p.ToToolEntries() {
					entries = append(entries, convertPluginToolEntry(entry))
				}
				continue
			}
		}
		// Fall back to installed Copilot CLI plugins (~/.copilot/installed-plugins/)
		if dir := resolveInstalledPlugin(name); dir != "" {
			entries = append(entries, ToolEntry{
				Name:   name,
				Type:   "skill",
				Source: "local",
				Path:   dir,
			})
			slog.Info("Resolved plugin from installed Copilot CLI plugins", "plugin", name, "path", dir)
			continue
		}
		slog.Warn("Plugin not found, skipping",
			"plugin", name,
			"config", c.Name,
			"hint", "Install with: /plugin install "+name)
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
