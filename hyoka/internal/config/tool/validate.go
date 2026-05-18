package tool

import (
	"bufio"
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/plugin"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/progress"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/toolload"
)

// ToolLoadError is a structured error returned by ValidateAndExpand when
// one or more declared tools fail to resolve. It carries the failing tool's
// Kind/Name and the reason so callers (engine, reviewer factory) can tag
// EvalReport entries with error_category=tool_load_failure.
type ToolLoadError struct {
	Kind   string // One of progress.ToolKindSkill / ToolKindPlugin / ToolKindMCP
	Name   string
	Reason string
}

func (e *ToolLoadError) Error() string {
	return fmt.Sprintf("%s %q: %s", e.Kind, e.Name, e.Reason)
}

// ToolLoadItem is one row in a ToolLoadReport. Each row describes a single
// resolved leaf: a skill directory, a plugin child (skill or MCP), a
// skill_dir child, or a top-level MCP server.
//
// Parent is the container identifier (plugin name or skills-dir path) or
// empty if the leaf has no container. ParentKind is one of
// progress.ToolParentKindPlugin / ToolParentKindSkillDir / empty.
type ToolLoadItem struct {
	Kind       string // "skill" | "plugin" | "mcp"
	Name       string
	Parent     string
	ParentKind string
	Status     string // "loaded" | "failed"
	Reason     string // populated when Status == "failed"
	Path       string // resolved filesystem path for skills (empty for MCP / failed)
	// Role records where the entry came from: "generator", "reviewer",
	// or "plugin". Used by callers that want to partition the report
	// (e.g. cmd/run.go only consuming reviewer skills).
	Role string
}

// ToolLoadReport is the output of ValidateAndExpand. It is a flat list of
// leaves — the renderer groups by (ParentKind, Parent) to display a nested
// view. The report's SkillDirs / MCPNames helpers extract the loaded-subset
// for downstream use (session config, reviewer.SetSkillDirectories).
type ToolLoadReport struct {
	Items []ToolLoadItem
}

// Failed returns true if any item has Status == "failed".
func (r *ToolLoadReport) Failed() bool {
	for _, it := range r.Items {
		if it.Status == progress.ToolStatusFailed {
			return true
		}
	}
	return false
}

// AllErrors returns a ToolLoadError for every failed item, in report order.
// Returns nil when no items are failed. This replaces the old FirstError()
// short-circuit so pre-session validation can surface every broken tool in
// one shot rather than dripping them out one-per-run.
func (r *ToolLoadReport) AllErrors() []*ToolLoadError {
	var out []*ToolLoadError
	for _, it := range r.Items {
		if it.Status == progress.ToolStatusFailed {
			out = append(out, &ToolLoadError{Kind: it.Kind, Name: it.Name, Reason: it.Reason})
		}
	}
	return out
}

// JoinedError returns a single error wrapping every per-tool failure via
// errors.Join, or nil when nothing failed. The Error() string is the
// human-readable summary produced by SummarizeToolLoadErrors so callers
// can pass the error straight through to EvalResult.Error / ErrorDetails.
func (r *ToolLoadReport) JoinedError() error {
	errs := r.AllErrors()
	if len(errs) == 0 {
		return nil
	}
	return &joinedToolLoadError{errs: errs}
}

// joinedToolLoadError carries the slice of per-tool errors AND a stable
// formatted summary. We don't use errors.Join directly because its default
// rendering is one-error-per-line with no header — operators benefit from
// the "N tool(s) failed to load:" lead-in. errors.Is / errors.As still work
// against any wrapped *ToolLoadError via the Unwrap() []error method.
type joinedToolLoadError struct {
	errs []*ToolLoadError
}

func (j *joinedToolLoadError) Error() string {
	return SummarizeToolLoadErrors(j.errs)
}

func (j *joinedToolLoadError) Unwrap() []error {
	out := make([]error, len(j.errs))
	for i, e := range j.errs {
		out[i] = e
	}
	return out
}

// SummarizeToolLoadErrors renders a multi-line summary of per-tool load
// failures suitable for both EvalResult.Error/ErrorDetails (eval engine
// path) and the cmd/run.go reviewer-validation path. Format is stable so
// downstream consumers (Item E post-session verification, log greppers)
// can rely on it.
//
//	2 tool(s) failed to load:
//	  • skill "python-best-practices": repo not found
//	  • mcp "azure-mcp": command "azure-mcp" not on PATH
//
// Returns "" for an empty/nil slice.
func SummarizeToolLoadErrors(errs []*ToolLoadError) string {
	if len(errs) == 0 {
		return ""
	}
	var b strings.Builder
	fmt.Fprintf(&b, "%d tool(s) failed to load:", len(errs))
	for _, e := range errs {
		b.WriteString("\n  • ")
		b.WriteString(e.Error())
	}
	return b.String()
}

// Compile-time assertion that joinedToolLoadError participates in the
// errors.Join unwrapping protocol so errors.Is/As traverse the wrapped set.
var _ interface{ Unwrap() []error } = (*joinedToolLoadError)(nil)

// GeneratorSkillDirs returns the loaded skill directory paths contributed
// by generator-role entries. Caller-facing helper for copilot.go.
func (r *ToolLoadReport) GeneratorSkillDirs() []string {
	return r.skillDirsByRole("generator")
}

// ReviewerSkillDirs returns the loaded skill directory paths contributed
// by reviewer-role entries. Caller-facing helper for cmd/run.go.
func (r *ToolLoadReport) ReviewerSkillDirs() []string {
	return r.skillDirsByRole("reviewer")
}

func (r *ToolLoadReport) skillDirsByRole(role string) []string {
	var out []string
	seen := map[string]bool{}
	for _, it := range r.Items {
		if it.Kind != progress.ToolKindSkill {
			continue
		}
		if it.Status != progress.ToolStatusLoaded {
			continue
		}
		if it.Role != role {
			continue
		}
		if it.Path == "" || seen[it.Path] {
			continue
		}
		seen[it.Path] = true
		out = append(out, it.Path)
	}
	return out
}

// ValidationInput bundles the per-config inputs to ValidateAndExpand.
// It decouples the validator from the config package so the tool package
// has no circular import with config.
type ValidationInput struct {
	// GeneratorTools and ReviewerTools are the tool entries from the
	// config's generator/reviewer sections. Entries with Type == "plugin"
	// are expanded into their child skills + MCP servers; the plugin
	// parent and each child are reported as individual items carrying
	// ParentName/ParentKind for grouped rendering. Role ("generator" /
	// "reviewer") is inherited from which list the plugin was declared in.
	GeneratorTools []Entry
	ReviewerTools  []Entry

	// ConfigDir is the isolated config dir (used as the baseDir for
	// resolving relative local skill paths). May be empty; in that
	// case paths are resolved relative to the current working dir.
	ConfigDir string

	// PluginsDir is a legacy directory containing plugin YAML definitions
	// (historically `./plugins`). May be empty. The resolver also checks
	// `./.hyoka/plugins/` and the remote cache (`~/.hyoka/cache/…`,
	// `~/.copilot/installed-plugins/…`) before failing. Local plugins
	// live at `./.hyoka/plugins/<name>/plugin.yaml` (or
	// `./.hyoka/plugins/<name>.yaml`); remote plugins resolve via the
	// marketplace cache. Fetch/resolution failures are fatal — ValidateAndExpand
	// returns an error before the eval runs (fail-fast contract).
	PluginsDir string

	// Emit receives per-leaf progress events (ToolResolutionStart /
	// ToolResolutionResult). A nil emit is a no-op.
	Emit ProgressEmitter
}

// ValidateAndExpand performs static pre-session validation of all declared
// tools. It expands plugins into their child skills + MCP servers, expands
// skill_dir entries into child skills, and validates MCP entries have the
// fields their mode requires.
//
// Any unresolved plugin, missing skill directory, zero-skill skill_dir,
// missing SKILL.md, or MCP field error produces a Failed item in the
// report AND a non-nil error with error_category=tool_load_failure
// semantics. The caller is expected to treat a non-nil error as a
// hard-fail and abort the eval before any model call.
//
// When every item in the report is Loaded, ValidateAndExpand returns
// (report, nil). When any item is Failed, ValidateAndExpand returns
// (report, error) where the error is a *joinedToolLoadError wrapping
// EVERY per-tool failure (see JoinedError / SummarizeToolLoadErrors).
// Callers should surface err.Error() as-is — it is a pre-formatted
// multi-line summary covering all failures, not just the first.
func ValidateAndExpand(ctx context.Context, in ValidationInput) (*ToolLoadReport, error) {
	report := &ToolLoadReport{}

	// Build plugin registry (best-effort: if the legacy PluginsDir doesn't
	// exist, registry stays empty and we fall through to the local
	// `./.hyoka/plugins` tree and then the remote cache).
	reg := plugin.NewRegistry()
	if in.PluginsDir != "" {
		if _, err := os.Stat(in.PluginsDir); err == nil {
			_ = reg.LoadDir(in.PluginsDir)
		}
	}
	// Also load plugins from `./.hyoka/plugins/<name>/plugin.yaml` and
	// `./.hyoka/plugins/<name>.yaml`. This is the new convention, mirroring
	// the `./.hyoka/` project layout used by skills/configs.
	loadHyokaPluginsDir(reg, in.ConfigDir)

	// Resolve generator + reviewer tool entries. Plugin entries are expanded
	// into parent + children with role inherited from which list they
	// came from (no dual-role auto-append).
	validateEntries(ctx, report, in.GeneratorTools, "generator", in.ConfigDir, in.Emit, reg, in.PluginsDir)
	validateEntries(ctx, report, in.ReviewerTools, "reviewer", in.ConfigDir, in.Emit, reg, in.PluginsDir)

	if err := report.JoinedError(); err != nil {
		return report, err
	}
	return report, nil
}

// loadHyokaPluginsDir populates reg with plugin YAMLs under
// `<configDir>/.hyoka/plugins/` when that directory exists. Each plugin can
// live at `.hyoka/plugins/<name>/plugin.yaml` (directory-style, preferred)
// or `.hyoka/plugins/<name>.yaml` (flat file). Errors are tolerated — the
// next resolution tier (remote cache) will catch truly missing plugins.
func loadHyokaPluginsDir(reg *plugin.Registry, configDir string) {
	base := hyokaPluginsBase(configDir)
	if base == "" {
		return
	}
	info, err := os.Stat(base)
	if err != nil || !info.IsDir() {
		return
	}
	// Flat `.yaml` files at the top of .hyoka/plugins/.
	_ = reg.LoadDir(base)
	// Directory-per-plugin: .hyoka/plugins/<name>/plugin.yaml.
	entries, err := os.ReadDir(base)
	if err != nil {
		return
	}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		candidate := filepath.Join(base, e.Name(), "plugin.yaml")
		if _, err := os.Stat(candidate); err != nil {
			continue
		}
		// Load by pointing a temporary dir at the file — LoadDir walks
		// the tree, so calling it on the plugin subdir picks the YAML up.
		_ = reg.LoadDir(filepath.Join(base, e.Name()))
	}
}

// hyokaPluginsBase returns the absolute path of `.hyoka/plugins` under the
// current working directory (the project root). The isolated per-eval
// configDir is deliberately not consulted here — local plugins live in the
// project, not in ephemeral scratch dirs. Returns empty string when the
// CWD cannot be determined.
func hyokaPluginsBase(configDir string) string {
	_ = configDir // retained for signature stability; project-relative layout is canonical
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return filepath.Join(wd, ".hyoka", "plugins")
}

// pluginCheckedPaths enumerates every filesystem path the plugin resolver
// inspects for `name`. Used in the fail-fast error message so operators
// know exactly where the resolver looked.
func pluginCheckedPaths(name, repo, configDir, pluginsDir string) []string {
	var paths []string
	if base := hyokaPluginsBase(configDir); base != "" {
		paths = append(paths,
			filepath.Join(base, name, "plugin.yaml"),
			filepath.Join(base, name+".yaml"),
		)
	}
	if pluginsDir != "" {
		paths = append(paths, filepath.Join(pluginsDir, name+".yaml"))
	}
	owner, repoName := plugin.SplitOwnerRepo(repo)
	if owner != "" && repoName != "" {
		// Canonical hyoka cache (toolload.CacheRoot) — matches what the
		// gitFetcher writes to and what plugin.ResolveInstalled inspects.
		// Single source of truth for the candidate list lives in
		// plugin.PluginCacheCandidates.
		repoCache := toolload.RepoCacheDir(owner, repoName, "default")
		paths = append(paths, plugin.PluginCacheCandidates(repoCache, name)...)
	}
	if home, err := os.UserHomeDir(); err == nil {
		if owner != "" && repoName != "" {
			paths = append(paths,
				filepath.Join(home, ".copilot", "installed-plugins", owner+"-"+repoName, name, "skills"),
			)
		}
		paths = append(paths, filepath.Join(home, ".copilot", "installed-plugins", name, "skills"))
	}
	return paths
}

func registryLookup(reg *plugin.Registry, name string) (*plugin.Plugin, bool) {
	if reg == nil {
		return nil, false
	}
	p, err := reg.Get(name)
	if err != nil {
		return nil, false
	}
	return p, true
}

func validateEntries(ctx context.Context, report *ToolLoadReport, entries []Entry, role, configDir string, emit ProgressEmitter, reg *plugin.Registry, pluginsDir string) {
	for _, entry := range entries {
		kind := entry.ResolvedType()
		switch kind {
		case progress.ToolKindSkill:
			validateSkillEntry(ctx, report, entry, role, configDir, emit)
		case progress.ToolKindMCP:
			validateMCPEntry(report, entry, role, emit)
		case TypePlugin:
			validatePluginEntry(ctx, report, entry, role, configDir, emit, reg, pluginsDir)
		}
	}
}

// validatePluginEntry resolves a `type: plugin` tool entry, fanning out to
// its child skills + MCP servers. source: local prefers the local plugin
// registry (`./.hyoka/plugins/` or the legacy PluginsDir); source: remote
// uses the explicit `repo:` field to locate the plugin in the marketplace
// cache (`~/.hyoka/cache/default/<owner>/<repo>/...`). On failure, the
// reason enumerates every path the resolver checked so operators can
// diagnose fast. Fetch/resolution failures are hard-fails upstream (caller
// returns the ToolLoadError before the eval session starts).
func validatePluginEntry(ctx context.Context, report *ToolLoadReport, entry Entry, role, configDir string, emit ProgressEmitter, reg *plugin.Registry, pluginsDir string) {
	name := entry.Name
	emitStart(emit, name, progress.ToolKindPlugin)

	// Reject the retired @marketplace shorthand outright. The full repo
	// locator goes in `repo:`, not stitched into the name.
	if strings.Contains(name, "@") {
		reason := fmt.Sprintf(
			"plugin name %q contains '@' — the @marketplace shorthand has been removed. "+
				"Set the plugin name to the bare identifier (e.g. %q) and declare the source repo "+
				"explicitly via repo: (e.g. repo: microsoft/skills).",
			name, strings.SplitN(name, "@", 2)[0],
		)
		report.Items = append(report.Items, ToolLoadItem{
			Kind:   progress.ToolKindPlugin,
			Name:   name,
			Status: progress.ToolStatusFailed,
			Reason: reason,
			Role:   role,
		})
		emitResultWithParent(emit, name, progress.ToolKindPlugin, progress.ToolStatusFailed, reason, "", "")
		return
	}

	src := entry.Source // "" | "local" | "remote"
	if src == "" {
		// Default: local if no repo, remote if repo is set.
		if entry.Repo != "" {
			src = SourceRemote
		} else {
			src = SourceLocal
		}
	}

	// Remote plugins require an explicit repo: locator. There are no magic
	// aliases — hyoka must be told exactly where to fetch from.
	if src == SourceRemote && entry.Repo == "" {
		reason := fmt.Sprintf(
			"plugin %q declares source: remote but has no repo: field. "+
				"Add repo: microsoft/skills (or your fork) so hyoka knows where to fetch it from.",
			name,
		)
		report.Items = append(report.Items, ToolLoadItem{
			Kind:   progress.ToolKindPlugin,
			Name:   name,
			Status: progress.ToolStatusFailed,
			Reason: reason,
			Role:   role,
		})
		emitResultWithParent(emit, name, progress.ToolKindPlugin, progress.ToolStatusFailed, reason, "", "")
		return
	}

	// Try local registry first when source == local (or unset -> inferred local).
	if src == SourceLocal {
		if p, ok := registryLookup(reg, name); ok {
			emitPluginLoadedWithChildren(report, emit, p, name, role, configDir, entry.ExcludedTools)
			return
		}
	}
	// Try remote cache (installed plugins) using the explicit repo.
	if src == SourceRemote || src == SourceLocal {
		if dir := plugin.ResolveInstalled(entry.Repo, name); dir != "" {
			// Container plugins (the standard Copilot
			// `.github/plugins/<name>/skills/<child>/SKILL.md` layout, used
			// by microsoft/skills) fan out to one report row per child
			// skill. The parent directory has no SKILL.md of its own, so
			// recording it as a single skill would cause the SDK to load
			// nothing and the post-session verifier would mark the plugin
			// failed even though every child loaded correctly.
			if children := plugin.EnumerateChildSkills(dir); len(children) > 0 {
				// NOTE: Do NOT emit a top-level Loaded result for the plugin
				// itself. The SDK never reports the plugin (it only reports
				// child skills/MCPs) so any Loaded plugin row would later be
				// flipped to Failed by the "loaded but not reported by SDK"
				// rule in display_interactive.onToolsVerified. The renderer
				// builds the plugin's "  - <name> (plugin):" header purely
				// from each child's ParentName/ParentKind.
				for _, childPath := range children {
					childName := filepath.Base(childPath)
					// Skip if excluded (pairwise deep mode)
					if contains(entry.ExcludedTools, childName) {
						continue
					}
					report.Items = append(report.Items, ToolLoadItem{
						Kind:       progress.ToolKindSkill,
						Name:       childName,
						Parent:     name,
						ParentKind: progress.ToolParentKindPlugin,
						Status:     progress.ToolStatusLoaded,
						Path:       childPath,
						Role:       role,
					})
					emitResultWithParent(emit, childName, progress.ToolKindSkill, progress.ToolStatusLoaded, "", name, progress.ToolParentKindPlugin)
				}
				return
			}
			// Single-skill plugin: the resolved dir has its own SKILL.md
			// (no children). One skill row, named after the plugin.
			item := ToolLoadItem{
				Kind:       progress.ToolKindSkill,
				Name:       name,
				Parent:     name,
				ParentKind: progress.ToolParentKindPlugin,
				Status:     progress.ToolStatusLoaded,
				Path:       dir,
				Role:       role,
			}
			report.Items = append(report.Items, item)
			// Single-skill plugin: only emit the child. The plugin parent is
			// a container, not an SDK-reported tool — emitting a Loaded row
			// for it would be flipped to Failed by onToolsVerified.
			emitResultWithParent(emit, name, progress.ToolKindSkill, progress.ToolStatusLoaded, "", name, progress.ToolParentKindPlugin)
			return
		}
		// If source was explicitly remote, try local as a last-ditch fallback.
		if src == SourceRemote {
			if p, ok := registryLookup(reg, name); ok {
				emitPluginLoadedWithChildren(report, emit, p, name, role, configDir, entry.ExcludedTools)
				return
			}
		}
	}

	// Cache miss for a remote plugin: try fetching into the canonical
	// cache (Item B). On success, re-run ResolveInstalled — which now
	// finds the plugin under <CacheRoot>/repos/<owner>/<repo>/default/...
	// and continues with the standard fan-out path. On fetch failure we
	// fall through to the enumerated-paths hard-fail below; the fetch
	// error is appended to the reason so operators see *why* the fetch
	// didn't help.
	var fetchErr error
	if src == SourceRemote && entry.Repo != "" {
		fetchEntry := entry
		fetchEntry.Source = SourceRemote
		if f := LookupFetcher(fetchEntry); f != nil {
			res, err := f.Fetch(ctx, FetchRequest{Entry: fetchEntry, Version: entry.Version})
			if err != nil {
				fetchErr = err
				slog.Debug("Plugin fetch failed", "plugin", name, "repo", entry.Repo, "error", err)
			} else {
				slog.Debug("Plugin fetched into cache", "plugin", name, "repo", entry.Repo, "dir", res.Dir)
				if dir := plugin.ResolveInstalled(entry.Repo, name); dir != "" {
					if children := plugin.EnumerateChildSkills(dir); len(children) > 0 {
						for _, childPath := range children {
							childName := filepath.Base(childPath)
							// Skip if excluded (pairwise deep mode)
							if contains(entry.ExcludedTools, childName) {
								continue
							}
							report.Items = append(report.Items, ToolLoadItem{
								Kind:       progress.ToolKindSkill,
								Name:       childName,
								Parent:     name,
								ParentKind: progress.ToolParentKindPlugin,
								Status:     progress.ToolStatusLoaded,
								Path:       childPath,
								Role:       role,
							})
							emitResultWithParent(emit, childName, progress.ToolKindSkill, progress.ToolStatusLoaded, "", name, progress.ToolParentKindPlugin)
						}
						return
					}
					report.Items = append(report.Items, ToolLoadItem{
						Kind:       progress.ToolKindSkill,
						Name:       name,
						Parent:     name,
						ParentKind: progress.ToolParentKindPlugin,
						Status:     progress.ToolStatusLoaded,
						Path:       dir,
						Role:       role,
					})
					emitResultWithParent(emit, name, progress.ToolKindSkill, progress.ToolStatusLoaded, "", name, progress.ToolParentKindPlugin)
					return
				}
				// Fetcher returned a dir but ResolveInstalled still missed
				// — record a synthetic fetch error so the failure message
				// makes sense.
				fetchErr = fmt.Errorf("fetched %s but plugin %q still not found in cache", res.Dir, name)
			}
		}
	}

	// Hard-fail with enumerated paths.
	paths := pluginCheckedPaths(name, entry.Repo, configDir, pluginsDir)
	reason := "plugin " + strconv.Quote(name) + " not found (source=" + src
	if entry.Repo != "" {
		reason += ", repo=" + entry.Repo
	}
	reason += "). Checked:"
	for _, p := range paths {
		reason += "\n  - " + p
	}
	if fetchErr != nil {
		reason += "\nFetch attempt failed: " + fetchErr.Error()
	}
	reason += "\nInstall a local plugin at .hyoka/plugins/" + name + "/plugin.yaml, or run: /plugin install " + name
	report.Items = append(report.Items, ToolLoadItem{
		Kind:   progress.ToolKindPlugin,
		Name:   name,
		Status: progress.ToolStatusFailed,
		Reason: reason,
		Role:   role,
	})
	emitResultWithParent(emit, name, progress.ToolKindPlugin, progress.ToolStatusFailed, reason, "", "")
}

// emitPluginLoadedWithChildren records each expanded child (skill or MCP)
// of a successfully-resolved plugin. Children carry ParentName/ParentKind
// for grouped rendering. The plugin parent itself is NOT emitted as a
// Loaded result — it is a container, and emitting Loaded for it would be
// flipped to Failed by display_interactive.onToolsVerified (SDK never
// reports the plugin, only its children). Role is inherited from the
// parent entry's list. ExcludedTools can filter out specific children
// (pairwise deep mode).
func emitPluginLoadedWithChildren(report *ToolLoadReport, emit ProgressEmitter, p *plugin.Plugin, name, role, configDir string, excludedTools []string) {
	for _, child := range p.ToToolEntries() {
		// Skip if excluded (pairwise deep mode)
		if contains(excludedTools, child.Name) {
			continue
		}
		childItem := ToolLoadItem{
			Kind:       child.Type,
			Name:       child.Name,
			Parent:     name,
			ParentKind: progress.ToolParentKindPlugin,
			Role:       role,
		}
		emitStart(emit, child.Name, child.Type)
		switch child.Type {
		case progress.ToolKindSkill:
			entry := Entry{
				Name:   child.Name,
				Type:   child.Type,
				Source: child.Source,
				Path:   child.Path,
				Repo:   child.Repo,
			}
			path, err := validateSingleSkill(entry, configDir)
			if err != nil {
				childItem.Status = progress.ToolStatusFailed
				childItem.Reason = err.Error()
			} else {
				childItem.Status = progress.ToolStatusLoaded
				childItem.Path = path
			}
		case progress.ToolKindMCP:
			if child.Command == "" {
				childItem.Status = progress.ToolStatusFailed
				childItem.Reason = "local MCP entry missing command"
			} else {
				childItem.Status = progress.ToolStatusLoaded
			}
		default:
			childItem.Status = progress.ToolStatusLoaded
		}
		report.Items = append(report.Items, childItem)
		emitResultWithParent(emit, child.Name, child.Type, childItem.Status, childItem.Reason, name, progress.ToolParentKindPlugin)
	}
}

func validateSkillEntry(ctx context.Context, report *ToolLoadReport, entry Entry, role, configDir string, emit ProgressEmitter) {
	emitStart(emit, entry.Name, progress.ToolKindSkill)

	switch entry.SkillSource() {
	case SourceRemote:
		dir, err := FetchRemote(ctx, entry)
		if err != nil {
			reason := fmt.Sprintf("fetching remote skill: %v", err)
			report.Items = append(report.Items, ToolLoadItem{
				Kind:   progress.ToolKindSkill,
				Name:   entry.Name,
				Status: progress.ToolStatusFailed,
				Reason: reason,
				Role:   role,
			})
			emitResultWithParent(emit, entry.Name, progress.ToolKindSkill, progress.ToolStatusFailed, reason, "", "")
			return
		}
		// Sanity check: the SDK uses the SKILL.md `name:` frontmatter when
		// loading a skill, not the config entry name. If they don't match,
		// the user will see the SKILL.md name in load events and reports —
		// not what they typed in YAML. Surface a clear warning so it doesn't
		// look like the skill failed to load when really it loaded under a
		// different label.
		if skillName, ok := readSkillFrontmatterName(dir); ok && skillName != "" && skillName != entry.Name {
			slog.Warn("Remote skill name differs from SKILL.md frontmatter — SDK will load it under the SKILL.md name",
				"entry_name", entry.Name,
				"skill_md_name", skillName,
				"path", dir,
				"hint", fmt.Sprintf("rename your config entry to `name: %s` to match", skillName))
		}
		report.Items = append(report.Items, ToolLoadItem{
			Kind:   progress.ToolKindSkill,
			Name:   entry.Name,
			Status: progress.ToolStatusLoaded,
			Path:   dir,
			Role:   role,
		})
		emitResultWithParent(emit, entry.Name, progress.ToolKindSkill, progress.ToolStatusLoaded, "", "", "")
		return
	case SourceLocal:
	// fall through
	default:
		reason := fmt.Sprintf("unknown skill source %q", entry.Source)
		report.Items = append(report.Items, ToolLoadItem{
			Kind:   progress.ToolKindSkill,
			Name:   entry.Name,
			Status: progress.ToolStatusFailed,
			Reason: reason,
			Role:   role,
		})
		emitResultWithParent(emit, entry.Name, progress.ToolKindSkill, progress.ToolStatusFailed, reason, "", "")
		return
	}

	if entry.SkillDir {
		validateSkillDirEntry(report, entry, role, configDir, emit)
		return
	}

	// Single local skill or glob.
	path := entry.Path
	if !filepath.IsAbs(path) && configDir != "" {
		path = filepath.Join(configDir, path)
	}
	if strings.ContainsAny(entry.Path, "*?[") {
		// Glob: expand, each match is treated as a single skill dir.
		matches, err := filepath.Glob(path)
		if err != nil {
			reason := fmt.Sprintf("invalid glob pattern %q: %v", entry.Path, err)
			report.Items = append(report.Items, ToolLoadItem{
				Kind:   progress.ToolKindSkill,
				Name:   entry.Name,
				Status: progress.ToolStatusFailed,
				Reason: reason,
				Role:   role,
			})
			emitResultWithParent(emit, entry.Name, progress.ToolKindSkill, progress.ToolStatusFailed, reason, "", "")
			return
		}
		var dirs []string
		for _, m := range matches {
			if info, err := os.Stat(m); err == nil && info.IsDir() {
				if abs, err := filepath.Abs(m); err == nil {
					dirs = append(dirs, abs)
				}
			}
		}
		if len(dirs) == 0 {
			reason := fmt.Sprintf("glob %q resolved to zero directories", entry.Path)
			report.Items = append(report.Items, ToolLoadItem{
				Kind:   progress.ToolKindSkill,
				Name:   entry.Name,
				Status: progress.ToolStatusFailed,
				Reason: reason,
				Role:   role,
			})
			emitResultWithParent(emit, entry.Name, progress.ToolKindSkill, progress.ToolStatusFailed, reason, "", "")
			return
		}
		for _, d := range dirs {
			report.Items = append(report.Items, ToolLoadItem{
				Kind:   progress.ToolKindSkill,
				Name:   filepath.Base(d),
				Parent: entry.Path,
				// Glob matches are treated as a skill_dir-style expansion.
				ParentKind: progress.ToolParentKindSkillDir,
				Status:     progress.ToolStatusLoaded,
				Path:       d,
				Role:       role,
			})
			emitResultWithParent(emit, filepath.Base(d), progress.ToolKindSkill, progress.ToolStatusLoaded, "", entry.Path, progress.ToolParentKindSkillDir)
		}
		return
	}

	// Single skill directory.
	resolved, err := validateSingleSkill(entry, configDir)
	if err != nil {
		report.Items = append(report.Items, ToolLoadItem{
			Kind:   progress.ToolKindSkill,
			Name:   entry.Name,
			Status: progress.ToolStatusFailed,
			Reason: err.Error(),
			Role:   role,
		})
		emitResultWithParent(emit, entry.Name, progress.ToolKindSkill, progress.ToolStatusFailed, err.Error(), "", "")
		return
	}
	report.Items = append(report.Items, ToolLoadItem{
		Kind:   progress.ToolKindSkill,
		Name:   entry.Name,
		Status: progress.ToolStatusLoaded,
		Path:   resolved,
		Role:   role,
	})
	emitResultWithParent(emit, entry.Name, progress.ToolKindSkill, progress.ToolStatusLoaded, "", "", "")
}

// validateSingleSkill returns the absolute path to a local skill directory
// after confirming it exists and contains SKILL.md.
func validateSingleSkill(entry Entry, configDir string) (string, error) {
	if entry.SkillSource() != SourceLocal {
		return "", fmt.Errorf("non-local skill requires remote fetch")
	}
	path := entry.Path
	if path == "" {
		return "", fmt.Errorf("skill entry missing path")
	}
	resolved := resolvePath(path, configDir)
	if resolved == "" {
		return "", fmt.Errorf("skill directory %q does not exist", entry.Path)
	}
	if _, err := os.Stat(filepath.Join(resolved, "SKILL.md")); err != nil {
		return "", fmt.Errorf("skill directory %q missing SKILL.md", resolved)
	}
	return resolved, nil
}

func validateSkillDirEntry(report *ToolLoadReport, entry Entry, role, configDir string, emit ProgressEmitter) {
	resolved := resolvePath(entry.Path, configDir)
	if resolved == "" {
		reason := fmt.Sprintf("skill_dir %q does not exist", entry.Path)
		report.Items = append(report.Items, ToolLoadItem{
			Kind:   progress.ToolKindSkill,
			Name:   entry.Name,
			Status: progress.ToolStatusFailed,
			Reason: reason,
			Role:   role,
		})
		emitResultWithParent(emit, entry.Name, progress.ToolKindSkill, progress.ToolStatusFailed, reason, "", "")
		return
	}
	dirEntries, err := os.ReadDir(resolved)
	if err != nil {
		reason := fmt.Sprintf("reading skill_dir %q: %v", resolved, err)
		report.Items = append(report.Items, ToolLoadItem{
			Kind:   progress.ToolKindSkill,
			Name:   entry.Name,
			Status: progress.ToolStatusFailed,
			Reason: reason,
			Role:   role,
		})
		emitResultWithParent(emit, entry.Name, progress.ToolKindSkill, progress.ToolStatusFailed, reason, "", "")
		return
	}
	// Parent row: the skill_dir itself. Marked Loaded if at least one
	// child skill resolves; otherwise Failed.
	var childRows []ToolLoadItem
	for _, e := range dirEntries {
		if !e.IsDir() {
			continue
		}
		subDir := filepath.Join(resolved, e.Name())
		if _, err := os.Stat(filepath.Join(subDir, "SKILL.md")); err != nil {
			continue
		}
		// Skip if excluded (pairwise deep mode)
		if contains(entry.ExcludedSkills, e.Name()) {
			continue
		}
		// Parent is the config-file `name:` value (e.g. "generator-skills"),
		// NOT the on-disk directory path. The interactive renderer builds the
		// parent header from this field; using the path would surface the
		// implementation detail (`./skills/generator`) instead of the
		// human-meaningful identifier the user wrote in their config.
		childRows = append(childRows, ToolLoadItem{
			Kind:       progress.ToolKindSkill,
			Name:       e.Name(),
			Parent:     entry.Name,
			ParentKind: progress.ToolParentKindSkillDir,
			Status:     progress.ToolStatusLoaded,
			Path:       subDir,
			Role:       role,
		})
	}
	if len(childRows) == 0 {
		reason := fmt.Sprintf("skill_dir %q contains no skills (no subdirectory with SKILL.md)", resolved)
		report.Items = append(report.Items, ToolLoadItem{
			Kind:   progress.ToolKindSkill,
			Name:   entry.Name,
			Status: progress.ToolStatusFailed,
			Reason: reason,
			Role:   role,
		})
		emitResultWithParent(emit, entry.Name, progress.ToolKindSkill, progress.ToolStatusFailed, reason, "", "")
		return
	}
	report.Items = append(report.Items, childRows...)
	for _, c := range childRows {
		emitResultWithParent(emit, c.Name, progress.ToolKindSkill, progress.ToolStatusLoaded, "", c.Parent, c.ParentKind)
	}
}

func validateMCPEntry(report *ToolLoadReport, entry Entry, role string, emit ProgressEmitter) {
	emitStart(emit, entry.Name, progress.ToolKindMCP)
	status := progress.ToolStatusLoaded
	reason := ""
	switch entry.ResolvedMCPType() {
	case "remote":
		if entry.URL == "" {
			status = progress.ToolStatusFailed
			reason = "remote MCP entry missing url"
		}
	default:
		if entry.Command == "" {
			status = progress.ToolStatusFailed
			reason = "local MCP entry missing command"
		}
	}
	report.Items = append(report.Items, ToolLoadItem{
		Kind:   progress.ToolKindMCP,
		Name:   entry.Name,
		Status: status,
		Reason: reason,
		Role:   role,
	})
	emitResultWithParent(emit, entry.Name, progress.ToolKindMCP, status, reason, "", "")
}

func emitResultWithParent(emit ProgressEmitter, name, kind, status, reason, parent, parentKind string) {
	if emit == nil {
		return
	}
	emit(progress.ProgressEvent{
		Type:       progress.EventToolResolutionResult,
		ToolName:   name,
		ToolKind:   kind,
		Status:     status,
		Reason:     reason,
		ParentName: parent,
		ParentKind: parentKind,
	})
}

// readSkillFrontmatterName reads the `name:` field from the YAML frontmatter
// of a SKILL.md file in the given directory. Returns ("", false) if SKILL.md
// is missing, has no frontmatter, or has no name field. Used to surface
// mismatches between a config entry's name and the actual skill name the
// SDK will load.
func readSkillFrontmatterName(skillDir string) (string, bool) {
	f, err := os.Open(filepath.Join(skillDir, "SKILL.md"))
	if err != nil {
		return "", false
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	inFrontmatter := false
	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		if !inFrontmatter {
			if trimmed == "---" {
				inFrontmatter = true
			}
			continue
		}
		if trimmed == "---" {
			return "", false
		}
		if rest, ok := strings.CutPrefix(trimmed, "name:"); ok {
			name := strings.TrimSpace(rest)
			name = strings.Trim(name, `"'`)
			if name == "" {
				return "", false
			}
			return name, true
		}
	}
	return "", false
}

// contains checks if a string is in a slice of strings.
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
