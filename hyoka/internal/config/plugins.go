package config

import (
	"log/slog"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config/tool"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/plugin"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/progress"
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

// EmitPluginResolutions replays plugin-lookup results as ToolResolutionStart /
// ToolResolutionResult progress events. It does not mutate the config — plugin
// expansion has already happened at config.Load time, and the registry lookup
// here is read-only. A nil emit is a no-op.
//
// Kind is always progress.ToolKindPlugin. A plugin that resolves via either
// the local plugin registry or the installed Copilot CLI plugin cache is
// reported as Loaded; otherwise Failed with reason "not found", matching the
// existing slog.Warn in ExpandPlugins.
func (c *ToolConfig) EmitPluginResolutions(emit tool.ProgressEmitter) {
	if emit == nil || len(c.Plugins) == 0 {
		return
	}
	reg := plugin.NewRegistry()
	// Best-effort load; an error here just means registry lookups will miss
	// and we fall through to the installed-plugins check (same as ExpandPlugins).
	if dir := resolvePluginsDir(); dir != "" {
		_ = reg.LoadDir(dir)
	}
	for _, name := range c.Plugins {
		emit(progress.ProgressEvent{
			Type:     progress.EventToolResolutionStart,
			ToolName: name,
			ToolKind: progress.ToolKindPlugin,
		})
		found := false
		if _, err := reg.Get(name); err == nil {
			found = true
		} else if dir := resolveInstalledPlugin(name); dir != "" {
			found = true
		}
		status := progress.ToolStatusLoaded
		reason := ""
		if !found {
			status = progress.ToolStatusFailed
			reason = "not found"
		}
		emit(progress.ProgressEvent{
			Type:     progress.EventToolResolutionResult,
			ToolName: name,
			ToolKind: progress.ToolKindPlugin,
			Status:   status,
			Reason:   reason,
		})
	}
}

func resolvePluginsDir() string {
	proj := DiscoverFromCWD()
	candidates := ResolveCandidates(proj, "plugins", "./plugins", "../plugins")
	if len(candidates) > 0 {
		return candidates[0]
	}
	return "./plugins"
}
