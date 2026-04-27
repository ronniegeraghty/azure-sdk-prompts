package tool

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/progress"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/toolload"
)

// TestValidateAndExpand_MissingPlugin_ErrorEnumeratesEveryCheckedPath
// pins the fail-fast contract: the reason string must list every
// filesystem path the resolver inspected, not just one. Operators
// should be able to copy the path list from the error and `ls` each.
func TestValidateAndExpand_MissingPlugin_ErrorEnumeratesEveryCheckedPath(t *testing.T) {
	dir := t.TempDir()

	// Use a fresh cwd so hyokaPluginsBase resolves to <dir>/.hyoka/plugins.
	prevWD, _ := os.Getwd()
	t.Cleanup(func() { _ = os.Chdir(prevWD) })
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}

	pluginsDir := filepath.Join(dir, "legacy-plugins")
	if err := os.Mkdir(pluginsDir, 0755); err != nil {
		t.Fatal(err)
	}

	name := "ghost-plugin"
	report, err := ValidateAndExpand(context.Background(), ValidationInput{
		GeneratorTools: []Entry{
			{Type: TypePlugin, Name: name, Source: "local"},
		},
		ConfigDir:  dir,
		PluginsDir: pluginsDir,
	})

	if err == nil {
		t.Fatal("expected error for missing plugin")
	}
	if !report.Failed() {
		t.Fatal("expected report.Failed() == true")
	}

	var reason string
	for _, item := range report.Items {
		if item.Kind == progress.ToolKindPlugin && item.Status == progress.ToolStatusFailed {
			reason = item.Reason
			break
		}
	}
	if reason == "" {
		t.Fatal("no failed plugin item in report")
	}

	home, _ := os.UserHomeDir()
	// source: local with no repo declared — the resolver only checks the
	// local plugin registry locations + the legacy ~/.copilot/installed-plugins/<name>/skills.
	// Cache paths under ~/.hyoka/cache/default/<owner>/<repo>/ are only
	// enumerated when repo: is declared (see the remote-plugin variant).
	wantPaths := []string{
		filepath.Join(dir, ".hyoka", "plugins", name, "plugin.yaml"),
		filepath.Join(dir, ".hyoka", "plugins", name+".yaml"),
		filepath.Join(pluginsDir, name+".yaml"),
		filepath.Join(home, ".copilot", "installed-plugins", name, "skills"),
	}

	for _, p := range wantPaths {
		if !strings.Contains(reason, p) {
			t.Errorf("missing path in enumerated error reason: %q\nFull reason:\n%s", p, reason)
		}
	}
	// Guard: a trivial matcher like "contains /plugins" would pass even
	// if only one path appeared. Require at least N distinct matches.
	distinct := 0
	for _, p := range wantPaths {
		if strings.Contains(reason, p) {
			distinct++
		}
	}
	if distinct < len(wantPaths) {
		t.Errorf("expected all %d paths enumerated, matched %d", len(wantPaths), distinct)
	}
}

// TestValidateAndExpand_MissingRemotePlugin_EnumeratesCachePathsForRepo
// verifies that a remote plugin with an explicit repo: that fails to
// resolve enumerates the per-repo cache paths in the error reason.
// We stub the plugin clone helper so this test stays offline; the fetch
// failure path still produces the enumerated-paths hard-fail message.
func TestValidateAndExpand_MissingRemotePlugin_EnumeratesCachePathsForRepo(t *testing.T) {
	dir := t.TempDir()
	prevHome := os.Getenv("HOME")
	prevWD, _ := os.Getwd()
	t.Cleanup(func() {
		_ = os.Setenv("HOME", prevHome)
		_ = os.Chdir(prevWD)
	})
	cleanHome := t.TempDir()
	_ = os.Setenv("HOME", cleanHome)
	restore := toolload.SetTestRoot(filepath.Join(cleanHome, ".hyoka", "cache"))
	defer restore()
	restoreClone := stubPluginCloneFailing(t)
	defer restoreClone()
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}

	report, err := ValidateAndExpand(context.Background(), ValidationInput{
		GeneratorTools: []Entry{
			{Type: TypePlugin, Name: "ghost-plugin", Source: "remote", Repo: "microsoft/skills"},
		},
		ConfigDir: dir,
	})
	if err == nil {
		t.Fatal("expected hard-fail for uncached remote plugin")
	}
	if !report.Failed() {
		t.Fatal("expected report.Failed() == true")
	}

	var reason string
	for _, item := range report.Items {
		if item.Kind == progress.ToolKindPlugin && item.Status == progress.ToolStatusFailed {
			reason = item.Reason
			break
		}
	}
	if reason == "" {
		t.Fatal("no failed plugin item in report")
	}

	wantPaths := []string{
		filepath.Join(toolload.RepoCacheDir("microsoft", "skills", "default"), ".github", "plugins", "ghost-plugin"),
		filepath.Join(toolload.RepoCacheDir("microsoft", "skills", "default"), ".github", "skills", "ghost-plugin"),
		filepath.Join(toolload.RepoCacheDir("microsoft", "skills", "default"), "skills", "ghost-plugin"),
		filepath.Join(cleanHome, ".copilot", "installed-plugins", "microsoft-skills", "ghost-plugin", "skills"),
	}
	for _, p := range wantPaths {
		if !strings.Contains(reason, p) {
			t.Errorf("missing per-repo cache path in error reason: %q\nFull reason:\n%s", p, reason)
		}
	}
}

// TestValidateAndExpand_PluginFanOut_TwoSkillsOneMCP verifies the
// Scope-2 mission requirement: a plugin containing 2 skills and 1
// MCP server produces exactly 3 child rows in the report, each with
// ParentName=<plugin> and ParentKind=plugin. All three must also be
// emitted as progress events for grouped rendering downstream.
func TestValidateAndExpand_PluginFanOut_TwoSkillsOneMCP(t *testing.T) {
	dir := t.TempDir()
	pluginDir := filepath.Join(dir, "plugins")
	if err := os.Mkdir(pluginDir, 0755); err != nil {
		t.Fatal(err)
	}
	// Plugin with 2 skills + 1 MCP.
	pluginYAML := `name: multi-child
skills:
  - name: skill-alpha
    path: ./skill-alpha
  - name: skill-beta
    path: ./skill-beta
mcp_servers:
  mcp-gamma:
    type: local
    command: npx
    args: ["-y", "something"]
`
	if err := os.WriteFile(filepath.Join(pluginDir, "multi-child.yaml"), []byte(pluginYAML), 0644); err != nil {
		t.Fatal(err)
	}
	// Create the two skill dirs with SKILL.md.
	for _, s := range []string{"skill-alpha", "skill-beta"} {
		sd := filepath.Join(dir, s)
		if err := os.Mkdir(sd, 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(sd, "SKILL.md"), []byte("# "+s), 0644); err != nil {
			t.Fatal(err)
		}
	}

	// Capture emitted events.
	var events []progress.ProgressEvent
	emit := func(ev progress.ProgressEvent) { events = append(events, ev) }

	report, err := ValidateAndExpand(context.Background(), ValidationInput{
		GeneratorTools: []Entry{
			{Type: TypePlugin, Name: "multi-child", Source: "local"},
		},
		ConfigDir:  dir,
		PluginsDir: pluginDir,
		Emit:       emit,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if report.Failed() {
		t.Fatalf("expected no failures, items: %+v", report.Items)
	}

	// Verify the 3 child items in the report.
	wantChildren := map[string]string{
		"skill-alpha": progress.ToolKindSkill,
		"skill-beta":  progress.ToolKindSkill,
		"mcp-gamma":   progress.ToolKindMCP,
	}
	gotChildren := make(map[string]string)
	for _, it := range report.Items {
		if it.Parent != "multi-child" {
			continue
		}
		if it.ParentKind != progress.ToolParentKindPlugin {
			t.Errorf("child %q: ParentKind=%q, want plugin", it.Name, it.ParentKind)
		}
		if it.Status != progress.ToolStatusLoaded {
			t.Errorf("child %q: Status=%q reason=%q, want loaded", it.Name, it.Status, it.Reason)
		}
		gotChildren[it.Name] = it.Kind
	}
	if len(gotChildren) != 3 {
		t.Errorf("expected 3 child items under plugin, got %d: %+v", len(gotChildren), gotChildren)
	}
	for name, wantKind := range wantChildren {
		if got, ok := gotChildren[name]; !ok {
			t.Errorf("missing child %q in report", name)
		} else if got != wantKind {
			t.Errorf("child %q: Kind=%q, want %q", name, got, wantKind)
		}
	}

	// Verify Result events carry ParentName/ParentKind for each child.
	childResults := 0
	for _, ev := range events {
		if ev.Type != progress.EventToolResolutionResult {
			continue
		}
		if ev.ParentName != "multi-child" {
			continue
		}
		if ev.ParentKind != progress.ToolParentKindPlugin {
			t.Errorf("event %q: ParentKind=%q, want plugin", ev.ToolName, ev.ParentKind)
		}
		if wantKind, ok := wantChildren[ev.ToolName]; ok {
			if ev.ToolKind != wantKind {
				t.Errorf("event %q: ToolKind=%q, want %q", ev.ToolName, ev.ToolKind, wantKind)
			}
			childResults++
		}
	}
	if childResults != 3 {
		t.Errorf("expected 3 child Result events with ParentName=multi-child, got %d", childResults)
	}
}

// TestValidateAndExpand_PluginOnlyInGenerator_ReviewerUntouched is the
// resolver-layer twin of the config-layer no-auto-append test: when a
// plugin is validated only through GeneratorTools, no reviewer-role
// items should appear in the report.
func TestValidateAndExpand_PluginOnlyInGenerator_ReviewerUntouched(t *testing.T) {
	dir := t.TempDir()
	pluginDir := filepath.Join(dir, "plugins")
	if err := os.Mkdir(pluginDir, 0755); err != nil {
		t.Fatal(err)
	}
	pluginYAML := `name: gen-only
skills:
  - name: lone-skill
    path: ./lone-skill
`
	if err := os.WriteFile(filepath.Join(pluginDir, "gen-only.yaml"), []byte(pluginYAML), 0644); err != nil {
		t.Fatal(err)
	}
	lone := filepath.Join(dir, "lone-skill")
	if err := os.Mkdir(lone, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(lone, "SKILL.md"), []byte("# lone"), 0644); err != nil {
		t.Fatal(err)
	}

	report, err := ValidateAndExpand(context.Background(), ValidationInput{
		GeneratorTools: []Entry{
			{Type: TypePlugin, Name: "gen-only", Source: "local"},
		},
		// ReviewerTools intentionally empty.
		ConfigDir:  dir,
		PluginsDir: pluginDir,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for _, it := range report.Items {
		if it.Role == "reviewer" {
			t.Errorf("no reviewer items expected, got %+v", it)
		}
	}
	// Sanity: at least one generator or plugin-role item exists.
	foundAny := false
	for _, it := range report.Items {
		if it.Role == "generator" || it.Role == "plugin" {
			foundAny = true
		}
	}
	if !foundAny {
		t.Error("expected at least one generator/plugin-role item")
	}
}

// TestValidateAndExpand_PluginInBothRoles_ChildrenResolveInBoth pins the
// dual-role contract: when an operator explicitly lists a plugin in
// both GeneratorTools AND ReviewerTools, the children resolve under
// each role independently.
func TestValidateAndExpand_PluginInBothRoles_ChildrenResolveInBoth(t *testing.T) {
	dir := t.TempDir()
	pluginDir := filepath.Join(dir, "plugins")
	if err := os.Mkdir(pluginDir, 0755); err != nil {
		t.Fatal(err)
	}
	pluginYAML := `name: shared
skills:
  - name: shared-skill
    path: ./shared-skill
`
	if err := os.WriteFile(filepath.Join(pluginDir, "shared.yaml"), []byte(pluginYAML), 0644); err != nil {
		t.Fatal(err)
	}
	sdir := filepath.Join(dir, "shared-skill")
	if err := os.Mkdir(sdir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(sdir, "SKILL.md"), []byte("# shared"), 0644); err != nil {
		t.Fatal(err)
	}

	report, err := ValidateAndExpand(context.Background(), ValidationInput{
		GeneratorTools: []Entry{
			{Type: TypePlugin, Name: "shared", Source: "local"},
		},
		ReviewerTools: []Entry{
			{Type: TypePlugin, Name: "shared", Source: "local"},
		},
		ConfigDir:  dir,
		PluginsDir: pluginDir,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	genCount, revCount := 0, 0
	for _, it := range report.Items {
		if it.Parent != "shared" {
			continue
		}
		switch it.Role {
		case "generator":
			genCount++
		case "reviewer":
			revCount++
		}
	}
	if genCount == 0 {
		t.Error("expected at least one generator-role child of shared plugin")
	}
	if revCount == 0 {
		t.Error("expected at least one reviewer-role child of shared plugin")
	}
}

// TestValidateAndExpand_SkillDir_ThreeSubdirs_ProducesThreeChildren
// verifies skill-dir fan-out for a literal (non-glob) path. Existing
// TestValidateAndExpand_GlobExpansion covers the glob case; this test
// pins the non-glob skill_dir: true branch.
func TestValidateAndExpand_SkillDir_ThreeSubdirs_ProducesThreeChildren(t *testing.T) {
	dir := t.TempDir()
	root := filepath.Join(dir, "skills-bundle")
	if err := os.Mkdir(root, 0755); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"one", "two", "three"} {
		sd := filepath.Join(root, name)
		if err := os.Mkdir(sd, 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(sd, "SKILL.md"), []byte("# "+name), 0644); err != nil {
			t.Fatal(err)
		}
	}

	var events []progress.ProgressEvent
	emit := func(ev progress.ProgressEvent) { events = append(events, ev) }

	report, err := ValidateAndExpand(context.Background(), ValidationInput{
		GeneratorTools: []Entry{
			{Type: "skill", Source: "local", Path: "skills-bundle", Name: "bundle", SkillDir: true},
		},
		ConfigDir: dir,
		Emit:      emit,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if report.Failed() {
		t.Fatalf("expected no failures, got: %+v", report.Items)
	}

	// Exactly 3 child items with ParentKind=skill_dir.
	childNames := map[string]bool{}
	for _, it := range report.Items {
		if it.Kind == progress.ToolKindSkill && it.ParentKind == progress.ToolParentKindSkillDir && it.Status == progress.ToolStatusLoaded {
			childNames[it.Name] = true
			if it.Parent == "" {
				t.Errorf("child %q: Parent unset, expected skills dir path", it.Name)
			}
		}
	}
	for _, want := range []string{"one", "two", "three"} {
		if !childNames[want] {
			t.Errorf("expected child skill %q, got children: %v", want, childNames)
		}
	}
	if len(childNames) != 3 {
		t.Errorf("expected exactly 3 skill-dir children, got %d", len(childNames))
	}

	// Each child must surface as a progress.EventToolResolutionResult
	// with ParentKind=skill_dir for Tank's grouped renderer.
	childEvents := 0
	for _, ev := range events {
		if ev.Type == progress.EventToolResolutionResult &&
			ev.ParentKind == progress.ToolParentKindSkillDir &&
			ev.Status == progress.ToolStatusLoaded {
			childEvents++
		}
	}
	if childEvents != 3 {
		t.Errorf("expected 3 skill-dir child Result events, got %d", childEvents)
	}
}

// TestValidateAndExpand_LocalPlugin_ResolvesFromHyokaPluginsDir verifies
// the documented default local-plugin location: a plugin.yaml at
// `<cwd>/.hyoka/plugins/<name>/plugin.yaml` resolves without needing an
// explicit PluginsDir. Regression guard for the migration claim that
// `source: local` means "look in .hyoka/plugins/".
func TestValidateAndExpand_LocalPlugin_ResolvesFromHyokaPluginsDir(t *testing.T) {
	dir := t.TempDir()
	prevWD, _ := os.Getwd()
	t.Cleanup(func() { _ = os.Chdir(prevWD) })
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}

	// Build the expected layout: .hyoka/plugins/my-local/plugin.yaml
	base := filepath.Join(dir, ".hyoka", "plugins", "my-local")
	if err := os.MkdirAll(base, 0755); err != nil {
		t.Fatal(err)
	}
	pluginYAML := `name: my-local
skills:
  - name: tucked-skill
    path: ./tucked
`
	if err := os.WriteFile(filepath.Join(base, "plugin.yaml"), []byte(pluginYAML), 0644); err != nil {
		t.Fatal(err)
	}
	// Skill referenced by the plugin is resolved relative to ConfigDir.
	tucked := filepath.Join(dir, "tucked")
	if err := os.Mkdir(tucked, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(tucked, "SKILL.md"), []byte("# tucked"), 0644); err != nil {
		t.Fatal(err)
	}

	report, err := ValidateAndExpand(context.Background(), ValidationInput{
		GeneratorTools: []Entry{
			{Type: TypePlugin, Name: "my-local", Source: "local"},
		},
		ConfigDir: dir,
		// PluginsDir intentionally empty — rely on .hyoka/plugins/ default.
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if report.Failed() {
		t.Fatalf("expected no failures, got: %+v", report.Items)
	}

	found := false
	for _, it := range report.Items {
		if it.Parent == "my-local" && it.Name == "tucked-skill" && it.Status == progress.ToolStatusLoaded {
			found = true
		}
	}
	if !found {
		t.Errorf("expected tucked-skill loaded under my-local plugin; report: %+v", report.Items)
	}
}

// TestValidateAndExpand_RemotePlugin_MissingCache_HardFails pins the
// contract that a `source: remote` plugin which neither resolves from
// cache nor successfully fetches surfaces as a *ToolLoadError before
// any session work happens. The plugin clone is stubbed to fail so the
// test stays offline; the failure still surfaces via the aggregated
// hard-fail path.
func TestValidateAndExpand_RemotePlugin_MissingCache_HardFails(t *testing.T) {
	dir := t.TempDir()
	prevHome := os.Getenv("HOME")
	prevWD, _ := os.Getwd()
	t.Cleanup(func() {
		_ = os.Setenv("HOME", prevHome)
		_ = os.Chdir(prevWD)
	})
	// Redirect HOME to a clean temp dir so the cache lookup cleanly misses.
	cleanHome := t.TempDir()
	_ = os.Setenv("HOME", cleanHome)
	restore := toolload.SetTestRoot(filepath.Join(cleanHome, ".hyoka", "cache"))
	defer restore()
	restoreClone := stubPluginCloneFailing(t)
	defer restoreClone()
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}

	report, err := ValidateAndExpand(context.Background(), ValidationInput{
		GeneratorTools: []Entry{
			{Type: TypePlugin, Name: "never-cached", Source: "remote", Repo: "microsoft/skills"},
		},
		ConfigDir: dir,
	})
	if err == nil {
		t.Fatal("expected hard-fail error for uncached remote plugin")
	}
	toolErr := new(ToolLoadError)
	if !errors.As(err, &toolErr) {
		t.Fatalf("expected wrapped *ToolLoadError, got %T (%v)", err, err)
	}
	if toolErr.Kind != progress.ToolKindPlugin {
		t.Errorf("expected Kind=plugin, got %q", toolErr.Kind)
	}
	if !report.Failed() {
		t.Error("expected report.Failed() == true")
	}
	// And the enumerated-paths contract must hold for remote source too.
	var reason string
	for _, it := range report.Items {
		if it.Kind == progress.ToolKindPlugin && it.Status == progress.ToolStatusFailed {
			reason = it.Reason
		}
	}
	if !strings.Contains(reason, ".hyoka/cache") {
		t.Errorf("remote plugin error should enumerate cache paths; got: %s", reason)
	}
}

// TestValidateAndExpand_RemoteContainerPlugin_FansOutChildren is the
// end-to-end regression for the microsoft/skills container plugin layout
// (`.github/plugins/<name>/skills/<child>/SKILL.md`).
//
// Pre-fix: ResolveInstalled returned "" because the plugin directory had
// no top-level SKILL.md, so the validator hard-failed with "plugin not
// found" even when the cache was correctly populated. Post-fix, the
// validator emits one report row per child skill, with ParentName=plugin
// and ParentKind=plugin, so the SDK loads them and the post-session
// verifier matches by child basename.
func TestValidateAndExpand_RemoteContainerPlugin_FansOutChildren(t *testing.T) {
prevHome := os.Getenv("HOME")
prevWD, _ := os.Getwd()
t.Cleanup(func() {
_ = os.Setenv("HOME", prevHome)
_ = os.Chdir(prevWD)
})

cleanHome := t.TempDir()
_ = os.Setenv("HOME", cleanHome)
restore := toolload.SetTestRoot(filepath.Join(cleanHome, ".hyoka", "cache"))
defer restore()
wd := t.TempDir()
if err := os.Chdir(wd); err != nil {
t.Fatal(err)
}

// Mirror microsoft/skills' layout for azure-sdk-python (truncated to
// 3 children for the test).
pluginDir := filepath.Join(toolload.RepoCacheDir("microsoft", "skills", "default"),
".github", "plugins", "azure-sdk-python")
if err := os.MkdirAll(pluginDir, 0o755); err != nil {
t.Fatal(err)
}
// README at root, no top-level SKILL.md (this is what fooled the old check).
if err := os.WriteFile(filepath.Join(pluginDir, "README.md"), []byte("readme"), 0o644); err != nil {
t.Fatal(err)
}
children := []string{"azure-keyvault-py", "azure-identity-py", "azure-storage-blob-py"}
for _, c := range children {
cd := filepath.Join(pluginDir, "skills", c)
if err := os.MkdirAll(cd, 0o755); err != nil {
t.Fatal(err)
}
if err := os.WriteFile(filepath.Join(cd, "SKILL.md"), []byte("# "+c), 0o644); err != nil {
t.Fatal(err)
}
}

report, err := ValidateAndExpand(context.Background(), ValidationInput{
GeneratorTools: []Entry{
{Type: TypePlugin, Name: "azure-sdk-python", Source: "remote", Repo: "microsoft/skills"},
},
ConfigDir: wd,
})
if err != nil {
t.Fatalf("ValidateAndExpand returned error for valid container plugin: %v", err)
}
if report.Failed() {
t.Fatalf("report has failed items: %+v", report.Items)
}

// Expect: one skill row per child (3 total), all parented to the plugin.
skillRows := []ToolLoadItem{}
for _, it := range report.Items {
if it.Kind == progress.ToolKindSkill {
skillRows = append(skillRows, it)
}
}
if len(skillRows) != len(children) {
t.Fatalf("got %d skill rows, want %d: %+v", len(skillRows), len(children), skillRows)
}

gotNames := map[string]bool{}
for _, row := range skillRows {
if row.Parent != "azure-sdk-python" {
t.Errorf("child %q has Parent=%q, want %q", row.Name, row.Parent, "azure-sdk-python")
}
if row.ParentKind != progress.ToolParentKindPlugin {
t.Errorf("child %q has ParentKind=%q, want %q", row.Name, row.ParentKind, progress.ToolParentKindPlugin)
}
if row.Status != progress.ToolStatusLoaded {
t.Errorf("child %q status=%q, want %q (reason: %s)", row.Name, row.Status, progress.ToolStatusLoaded, row.Reason)
}
if row.Path == "" {
t.Errorf("child %q has empty Path", row.Name)
}
if row.Role != "generator" {
t.Errorf("child %q role=%q, want %q", row.Name, row.Role, "generator")
}
gotNames[row.Name] = true
}
for _, want := range children {
if !gotNames[want] {
t.Errorf("missing child skill row %q in report", want)
}
}

// GeneratorSkillDirs feeds SessionConfig.SkillDirectories — must
// expose the per-child paths so the SDK loads each leaf, not the
// container directory.
dirs := report.GeneratorSkillDirs()
if len(dirs) != len(children) {
t.Errorf("GeneratorSkillDirs returned %d paths, want %d: %v", len(dirs), len(children), dirs)
}
for _, d := range dirs {
if !strings.HasPrefix(d, pluginDir) {
t.Errorf("path %q not under plugin dir %q", d, pluginDir)
}
if filepath.Base(filepath.Dir(d)) != "skills" {
t.Errorf("path %q parent is not 'skills' (likely returned container dir, not child)", d)
}
}
}
