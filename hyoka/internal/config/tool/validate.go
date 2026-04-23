package tool

import (
"context"
"fmt"
"os"
"path/filepath"
"strings"

"github.com/ronniegeraghty/hyoka/hyoka/internal/plugin"
"github.com/ronniegeraghty/hyoka/hyoka/internal/progress"
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

// FirstError returns a ToolLoadError for the first failed item, or nil.
func (r *ToolLoadReport) FirstError() *ToolLoadError {
for _, it := range r.Items {
if it.Status == progress.ToolStatusFailed {
return &ToolLoadError{Kind: it.Kind, Name: it.Name, Reason: it.Reason}
}
}
return nil
}

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
// Plugins is the raw list of plugin names declared in the config
// (config.ToolConfig.Plugins). Each name must resolve to either a
// local plugin YAML in PluginsDir or an installed plugin under
// ~/.copilot/installed-plugins/ (via plugin.ResolveInstalled).
Plugins []string

// GeneratorTools and ReviewerTools are the already-plugin-expanded
// tool entries from the config's generator/reviewer sections.
// Children that originated from a plugin (determined by matching
// (kind, name) against the plugin's ToToolEntries) are skipped by
// the validator — the plugin block reports them instead.
GeneratorTools []Entry
ReviewerTools  []Entry

// ConfigDir is the isolated config dir (used as the baseDir for
// resolving relative local skill paths). May be empty; in that
// case paths are resolved relative to the current working dir.
ConfigDir string

// PluginsDir is the directory containing local plugin YAML definitions
// (typically "./plugins"). May be empty; when empty, only installed
// plugins are considered.
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
// (report, *ToolLoadError) so callers can both render the full report
// AND tag the eval result with the first failure's details.
func ValidateAndExpand(ctx context.Context, in ValidationInput) (*ToolLoadReport, error) {
report := &ToolLoadReport{}

// Build plugin registry (best-effort: if PluginsDir doesn't exist,
// registry stays empty and we fall through to installed-plugins).
reg := plugin.NewRegistry()
if in.PluginsDir != "" {
if _, err := os.Stat(in.PluginsDir); err == nil {
_ = reg.LoadDir(in.PluginsDir)
}
}

// Track (kind, name) pairs that belong to plugin expansions so we
// skip them when iterating Generator/Reviewer tools below.
pluginChildKeys := map[string]bool{}

// Resolve plugins first. Each plugin emits a parent "loaded"/"failed"
// row followed by per-child rows keyed by ParentName=plugin.
for _, name := range in.Plugins {
emitStart(in.Emit, name, progress.ToolKindPlugin)
p, ok := registryLookup(reg, name)
if !ok {
// Try installed plugins as an opaque skill dir.
if dir := plugin.ResolveInstalled(name); dir != "" {
path := dir
item := ToolLoadItem{
Kind:       progress.ToolKindSkill,
Name:       name,
Parent:     name,
ParentKind: progress.ToolParentKindPlugin,
Status:     progress.ToolStatusLoaded,
Path:       path,
Role:       "plugin",
}
report.Items = append(report.Items, item)
emitResultWithParent(in.Emit, name, progress.ToolKindPlugin, progress.ToolStatusLoaded, "", "", "")
emitResultWithParent(in.Emit, name, progress.ToolKindSkill, progress.ToolStatusLoaded, "", name, progress.ToolParentKindPlugin)
pluginChildKeys[progress.ToolKindSkill+":"+name] = true
continue
}
reason := "plugin not found in registry or installed plugins"
report.Items = append(report.Items, ToolLoadItem{
Kind:   progress.ToolKindPlugin,
Name:   name,
Status: progress.ToolStatusFailed,
Reason: reason,
Role:   "plugin",
})
emitResultWithParent(in.Emit, name, progress.ToolKindPlugin, progress.ToolStatusFailed, reason, "", "")
continue
}
emitResultWithParent(in.Emit, name, progress.ToolKindPlugin, progress.ToolStatusLoaded, "", "", "")
// Enumerate children.
for _, child := range p.ToToolEntries() {
pluginChildKeys[child.Type+":"+child.Name] = true
childItem := ToolLoadItem{
Kind:       child.Type,
Name:       child.Name,
Parent:     name,
ParentKind: progress.ToolParentKindPlugin,
Role:       "plugin",
}
emitStart(in.Emit, child.Name, child.Type)
switch child.Type {
case progress.ToolKindSkill:
entry := Entry{
Name:   child.Name,
Type:   child.Type,
Source: child.Source,
Path:   child.Path,
Repo:   child.Repo,
}
path, err := validateSingleSkill(entry, in.ConfigDir)
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
emitResultWithParent(in.Emit, child.Name, child.Type, childItem.Status, childItem.Reason, name, progress.ToolParentKindPlugin)
}
}

// Resolve generator + reviewer tool entries. Skip entries that came
// from plugin expansion (already reported as plugin children).
validateEntries(ctx, report, in.GeneratorTools, "generator", in.ConfigDir, in.Emit, pluginChildKeys)
validateEntries(ctx, report, in.ReviewerTools, "reviewer", in.ConfigDir, in.Emit, pluginChildKeys)

if err := report.FirstError(); err != nil {
return report, err
}
return report, nil
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

func validateEntries(ctx context.Context, report *ToolLoadReport, entries []Entry, role, configDir string, emit ProgressEmitter, pluginKeys map[string]bool) {
for _, entry := range entries {
kind := entry.ResolvedType()
key := kind + ":" + entry.Name
if pluginKeys[key] {
continue
}
switch kind {
case progress.ToolKindSkill:
validateSkillEntry(ctx, report, entry, role, configDir, emit)
case progress.ToolKindMCP:
validateMCPEntry(report, entry, role, emit)
}
}
}

func validateSkillEntry(ctx context.Context, report *ToolLoadReport, entry Entry, role, configDir string, emit ProgressEmitter) {
emitStart(emit, entry.Name, progress.ToolKindSkill)

switch entry.SkillSource() {
case SourceRemote:
dir, err := FetchRemote(ctx, entry, configDir)
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
childRows = append(childRows, ToolLoadItem{
Kind:       progress.ToolKindSkill,
Name:       e.Name(),
Parent:     entry.Path,
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
