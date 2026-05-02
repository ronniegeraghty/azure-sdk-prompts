package eval

import (
	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/config/tool"
)

// injectConfigProps adds eval-config-derived properties to the props map used
// for grader `when:` matching. Phase 1 keys:
//
//   - "generator"          → cfg.Generator.Model
//   - "config"             → cfg.Name
//   - "skill:<name>"       → "true" for each ToolEntry with Type=="skill"
//   - "mcp_server:<name>"  → "true" for each ToolEntry with Type=="mcp"
//   - "plugin:<name>"      → "true" for each ToolEntry with Type=="plugin"
//
// Entries with an empty Name or unknown Type are skipped silently. Keys are
// written unconditionally — they overwrite any prompt frontmatter clash
// (the `:`-prefixed namespace is reserved for engine-injected props).
//
// Phase 2 (mcp_tool:<server>/<tool>, value negation) is intentionally not
// implemented here; see decision drop neo-config-aware-when-phase1.md.
func injectConfigProps(props map[string]string, cfg config.ToolConfig) {
	if props == nil {
		return
	}
	props["config"] = cfg.Name
	if cfg.Generator != nil {
		props["generator"] = cfg.Generator.Model
		for _, te := range cfg.Generator.Tools {
			if te.Name == "" {
				continue
			}
			switch te.ResolvedType() {
			case tool.TypeSkill:
				props["skill:"+te.Name] = "true"
			case tool.TypeMCP:
				props["mcp_server:"+te.Name] = "true"
			case tool.TypePlugin:
				props["plugin:"+te.Name] = "true"
			}
		}
	}
}
