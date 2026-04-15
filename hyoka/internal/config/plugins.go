package config

import (
	"log/slog"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config/tool"
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
					entries = append(entries, tool.ConvertPluginEntry(entry))
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
		c.Generator.Tools = tool.AppendEntries(c.Generator.Tools, entries)
	}
	if c.Reviewer != nil {
		c.Reviewer.Tools = tool.AppendEntries(c.Reviewer.Tools, entries)
	}
	return nil
}

func resolvePluginsDir() string {
	proj := DiscoverFromCWD()
	candidates := ResolveCandidates(proj, "plugins", "./plugins", "../plugins")
	if len(candidates) > 0 {
		return candidates[0]
	}
	return "./plugins"
}
