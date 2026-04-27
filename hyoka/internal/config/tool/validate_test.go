package tool

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/plugin"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/progress"
)

func TestValidateAndExpand_HappyPath(t *testing.T) {
	dir := t.TempDir()
	
	// Create valid plugin YAML
	pluginDir := filepath.Join(dir, "plugins")
	if err := os.Mkdir(pluginDir, 0755); err != nil {
		t.Fatal(err)
	}
	pluginYAML := `name: test-plugin
skills:
  - name: plugin-skill
    path: ./plugin-skill
`
	if err := os.WriteFile(filepath.Join(pluginDir, "test-plugin.yaml"), []byte(pluginYAML), 0644); err != nil {
		t.Fatal(err)
	}
	
	// Create the skill referenced by plugin
	pluginSkillDir := filepath.Join(dir, "plugin-skill")
	if err := os.Mkdir(pluginSkillDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(pluginSkillDir, "SKILL.md"), []byte("# Plugin Skill"), 0644); err != nil {
		t.Fatal(err)
	}
	
	// Create valid skill_dir
	skillsDir := filepath.Join(dir, "skills")
	if err := os.Mkdir(skillsDir, 0755); err != nil {
		t.Fatal(err)
	}
	skill1 := filepath.Join(skillsDir, "skill1")
	if err := os.Mkdir(skill1, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(skill1, "SKILL.md"), []byte("# Skill 1"), 0644); err != nil {
		t.Fatal(err)
	}
	
	// Create valid single skill
	singleSkill := filepath.Join(dir, "single-skill")
	if err := os.Mkdir(singleSkill, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(singleSkill, "SKILL.md"), []byte("# Single Skill"), 0644); err != nil {
		t.Fatal(err)
	}
	
	report, err := ValidateAndExpand(context.Background(), ValidationInput{
		GeneratorTools: []Entry{
			{Type: "plugin", Name: "test-plugin", Source: "local"},
			{Type: "skill", Source: "local", Path: "skills", Name: "gen-skills", SkillDir: true},
			{Type: "skill", Source: "local", Path: "single-skill", Name: "single"},
			{Type: "mcp", Name: "test-mcp", Command: "npx test"},
		},
		ConfigDir:  dir,
		PluginsDir: pluginDir,
	})
	
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if report.Failed() {
		t.Fatalf("expected no failures, report: %+v", report)
	}
	
	// Verify items
	if len(report.Items) == 0 {
		t.Fatal("expected at least one item")
	}
	
	// Check plugin child loaded
	foundPluginSkill := false
	for _, item := range report.Items {
		if item.Kind == progress.ToolKindSkill && item.Name == "plugin-skill" && item.Status == progress.ToolStatusLoaded {
			foundPluginSkill = true
			if item.ParentKind != progress.ToolParentKindPlugin {
				t.Errorf("plugin skill should have ParentKind=plugin, got %q", item.ParentKind)
			}
		}
	}
	if !foundPluginSkill {
		t.Error("expected plugin-skill to be loaded")
	}
	
	// Check skill_dir child loaded
	foundSkill1 := false
	for _, item := range report.Items {
		if item.Kind == progress.ToolKindSkill && item.Name == "skill1" && item.Status == progress.ToolStatusLoaded {
			foundSkill1 = true
			if item.ParentKind != progress.ToolParentKindSkillDir {
				t.Errorf("skill1 should have ParentKind=skill_dir, got %q", item.ParentKind)
			}
			// Parent must be the config-file `name:` value (issue (a) regression).
			if item.Parent != "gen-skills" {
				t.Errorf("skill1 Parent should be entry.Name=%q, got %q", "gen-skills", item.Parent)
			}
		}
	}
	if !foundSkill1 {
		t.Error("expected skill1 from skill_dir to be loaded")
	}
	
	// Check single skill loaded
	foundSingle := false
	for _, item := range report.Items {
		if item.Kind == progress.ToolKindSkill && item.Name == "single" && item.Status == progress.ToolStatusLoaded {
			foundSingle = true
			if item.Parent != "" {
				t.Errorf("single skill should have no parent, got %q", item.Parent)
			}
		}
	}
	if !foundSingle {
		t.Error("expected single skill to be loaded")
	}
	
	// Check MCP loaded
	foundMCP := false
	for _, item := range report.Items {
		if item.Kind == progress.ToolKindMCP && item.Name == "test-mcp" && item.Status == progress.ToolStatusLoaded {
			foundMCP = true
		}
	}
	if !foundMCP {
		t.Error("expected test-mcp to be loaded")
	}
}

func TestValidateAndExpand_MissingPlugin(t *testing.T) {
	dir := t.TempDir()
	pluginDir := filepath.Join(dir, "plugins")
	if err := os.Mkdir(pluginDir, 0755); err != nil {
		t.Fatal(err)
	}
	
	report, err := ValidateAndExpand(context.Background(), ValidationInput{
		GeneratorTools: []Entry{
			{Type: "plugin", Name: "nonexistent-plugin", Source: "local"},
		},
		ConfigDir:  dir,
		PluginsDir: pluginDir,
	})
	
	if err == nil {
		t.Fatal("expected error for missing plugin")
	}
	
	toolErr, ok := err.(*ToolLoadError)
	if !ok {
		t.Fatalf("expected *ToolLoadError, got %T", err)
	}
	if toolErr.Kind != progress.ToolKindPlugin {
		t.Errorf("expected Kind=plugin, got %q", toolErr.Kind)
	}
	if toolErr.Name != "nonexistent-plugin" {
		t.Errorf("expected Name=nonexistent-plugin, got %q", toolErr.Name)
	}
	if toolErr.Reason == "" {
		t.Error("expected non-empty reason")
	}
	
	if !report.Failed() {
		t.Error("expected report.Failed() == true")
	}
	
	// Check the report item
	foundFailed := false
	for _, item := range report.Items {
		if item.Kind == progress.ToolKindPlugin && item.Name == "nonexistent-plugin" && item.Status == progress.ToolStatusFailed {
			foundFailed = true
			if item.Reason == "" {
				t.Error("expected non-empty reason in item")
			}
		}
	}
	if !foundFailed {
		t.Error("expected failed plugin item in report")
	}
}

func TestValidateAndExpand_MissingSkillDir(t *testing.T) {
	dir := t.TempDir()
	
	report, err := ValidateAndExpand(context.Background(), ValidationInput{
		GeneratorTools: []Entry{
			{Type: "skill", Source: "local", Path: "./nonexistent", Name: "missing-dir", SkillDir: true},
		},
		ConfigDir: dir,
	})
	
	if err == nil {
		t.Fatal("expected error for missing skill dir")
	}
	
	toolErr, ok := err.(*ToolLoadError)
	if !ok {
		t.Fatalf("expected *ToolLoadError, got %T", err)
	}
	if toolErr.Kind != progress.ToolKindSkill {
		t.Errorf("expected Kind=skill, got %q", toolErr.Kind)
	}
	
	if !report.Failed() {
		t.Error("expected report.Failed() == true")
	}
}

func TestValidateAndExpand_MalformedPluginYAML(t *testing.T) {
	dir := t.TempDir()
	pluginDir := filepath.Join(dir, "plugins")
	if err := os.Mkdir(pluginDir, 0755); err != nil {
		t.Fatal(err)
	}
	
	// Write invalid YAML
	badYAML := `name: bad-plugin
skills: [this is not valid yaml
`
	if err := os.WriteFile(filepath.Join(pluginDir, "bad-plugin.yaml"), []byte(badYAML), 0644); err != nil {
		t.Fatal(err)
	}
	
	// Plugin registry load should fail or return empty
	// ValidateAndExpand should report plugin not found
	report, err := ValidateAndExpand(context.Background(), ValidationInput{
		GeneratorTools: []Entry{
			{Type: "plugin", Name: "bad-plugin", Source: "local"},
		},
		ConfigDir:  dir,
		PluginsDir: pluginDir,
	})
	
	if err == nil {
		t.Fatal("expected error for malformed plugin")
	}
	
	if !report.Failed() {
		t.Error("expected report.Failed() == true")
	}
}

func TestValidateAndExpand_PluginChildMissingSKILLMD(t *testing.T) {
	dir := t.TempDir()
	pluginDir := filepath.Join(dir, "plugins")
	if err := os.Mkdir(pluginDir, 0755); err != nil {
		t.Fatal(err)
	}
	
	pluginYAML := `name: broken-plugin
skills:
  - name: broken-skill
    path: ./broken-skill
`
	if err := os.WriteFile(filepath.Join(pluginDir, "broken-plugin.yaml"), []byte(pluginYAML), 0644); err != nil {
		t.Fatal(err)
	}
	
	// Create skill dir but no SKILL.md
	brokenSkill := filepath.Join(dir, "broken-skill")
	if err := os.Mkdir(brokenSkill, 0755); err != nil {
		t.Fatal(err)
	}
	
	report, err := ValidateAndExpand(context.Background(), ValidationInput{
		GeneratorTools: []Entry{
			{Type: "plugin", Name: "broken-plugin", Source: "local"},
		},
		ConfigDir:  dir,
		PluginsDir: pluginDir,
	})
	
	if err == nil {
		t.Fatal("expected error for plugin child missing SKILL.md")
	}
	
	if !report.Failed() {
		t.Error("expected report.Failed() == true")
	}
	
	// Verify child is marked failed
	foundFailedChild := false
	for _, item := range report.Items {
		if item.Kind == progress.ToolKindSkill && item.Name == "broken-skill" && item.Status == progress.ToolStatusFailed {
			foundFailedChild = true
			if item.ParentKind != progress.ToolParentKindPlugin {
				t.Errorf("expected ParentKind=plugin, got %q", item.ParentKind)
			}
		}
	}
	if !foundFailedChild {
		t.Error("expected broken-skill child to be marked failed")
	}
}

func TestValidateAndExpand_EmptySkillDir(t *testing.T) {
	dir := t.TempDir()
	emptyDir := filepath.Join(dir, "empty-skills")
	if err := os.Mkdir(emptyDir, 0755); err != nil {
		t.Fatal(err)
	}
	// Add a non-skill file
	if err := os.WriteFile(filepath.Join(emptyDir, ".gitkeep"), nil, 0644); err != nil {
		t.Fatal(err)
	}
	
	report, err := ValidateAndExpand(context.Background(), ValidationInput{
		GeneratorTools: []Entry{
			{Type: "skill", Source: "local", Path: "empty-skills", Name: "empty", SkillDir: true},
		},
		ConfigDir: dir,
	})
	
	if err == nil {
		t.Fatal("expected error for empty skill_dir")
	}
	
	if !report.Failed() {
		t.Error("expected report.Failed() == true")
	}
	
	// Verify reason mentions zero skills
	foundEmpty := false
	for _, item := range report.Items {
		if item.Kind == progress.ToolKindSkill && item.Name == "empty" && item.Status == progress.ToolStatusFailed {
			foundEmpty = true
			if item.Reason == "" || len(item.Reason) == 0 {
				t.Error("expected non-empty reason for empty skill_dir")
			}
		}
	}
	if !foundEmpty {
		t.Error("expected empty skill_dir item to be marked failed")
	}
}

func TestValidateAndExpand_EmptyConfig(t *testing.T) {
	dir := t.TempDir()
	
	report, err := ValidateAndExpand(context.Background(), ValidationInput{
		GeneratorTools: nil,
		ReviewerTools:  nil,
		ConfigDir:      dir,
	})
	
	if err != nil {
		t.Fatalf("expected no error for empty config, got: %v", err)
	}
	if report.Failed() {
		t.Errorf("expected no failures for empty config")
	}
	if len(report.Items) != 0 {
		t.Errorf("expected 0 items for empty config, got %d", len(report.Items))
	}
}

func TestValidateAndExpand_RelativeVsAbsolutePaths(t *testing.T) {
	dir := t.TempDir()
	
	// Create skill with relative path
	relSkill := filepath.Join(dir, "rel-skill")
	if err := os.Mkdir(relSkill, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(relSkill, "SKILL.md"), []byte("# Rel"), 0644); err != nil {
		t.Fatal(err)
	}
	
	// Create skill with absolute path
	absSkill := filepath.Join(dir, "abs-skill")
	if err := os.Mkdir(absSkill, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(absSkill, "SKILL.md"), []byte("# Abs"), 0644); err != nil {
		t.Fatal(err)
	}
	
	report, err := ValidateAndExpand(context.Background(), ValidationInput{
		GeneratorTools: []Entry{
			{Type: "skill", Source: "local", Path: "rel-skill", Name: "rel"},
			{Type: "skill", Source: "local", Path: absSkill, Name: "abs"},
		},
		ConfigDir: dir,
	})
	
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if report.Failed() {
		t.Fatal("expected no failures")
	}
	
	foundRel := false
	foundAbs := false
	for _, item := range report.Items {
		if item.Kind == progress.ToolKindSkill && item.Name == "rel" && item.Status == progress.ToolStatusLoaded {
			foundRel = true
			if !filepath.IsAbs(item.Path) {
				t.Errorf("expected absolute path for rel skill, got %q", item.Path)
			}
		}
		if item.Kind == progress.ToolKindSkill && item.Name == "abs" && item.Status == progress.ToolStatusLoaded {
			foundAbs = true
			if !filepath.IsAbs(item.Path) {
				t.Errorf("expected absolute path for abs skill, got %q", item.Path)
			}
		}
	}
	if !foundRel {
		t.Error("expected rel-skill to be loaded")
	}
	if !foundAbs {
		t.Error("expected abs-skill to be loaded")
	}
}

func TestValidateAndExpand_MCPMissingCommand(t *testing.T) {
	dir := t.TempDir()
	
	report, err := ValidateAndExpand(context.Background(), ValidationInput{
		GeneratorTools: []Entry{
			{Type: "mcp", Name: "broken-mcp", Command: ""},
		},
		ConfigDir: dir,
	})
	
	if err == nil {
		t.Fatal("expected error for MCP missing command")
	}
	
	if !report.Failed() {
		t.Error("expected report.Failed() == true")
	}
	
	foundFailed := false
	for _, item := range report.Items {
		if item.Kind == progress.ToolKindMCP && item.Name == "broken-mcp" && item.Status == progress.ToolStatusFailed {
			foundFailed = true
			if item.Reason == "" {
				t.Error("expected non-empty reason")
			}
		}
	}
	if !foundFailed {
		t.Error("expected broken-mcp to be marked failed")
	}
}

func TestValidateAndExpand_ReviewerRolePartitioning(t *testing.T) {
	dir := t.TempDir()
	
	// Create generator skill
	genSkill := filepath.Join(dir, "gen-skill")
	if err := os.Mkdir(genSkill, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(genSkill, "SKILL.md"), []byte("# Gen"), 0644); err != nil {
		t.Fatal(err)
	}
	
	// Create reviewer skill
	revSkill := filepath.Join(dir, "rev-skill")
	if err := os.Mkdir(revSkill, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(revSkill, "SKILL.md"), []byte("# Rev"), 0644); err != nil {
		t.Fatal(err)
	}
	
	report, err := ValidateAndExpand(context.Background(), ValidationInput{
		GeneratorTools: []Entry{
			{Type: "skill", Source: "local", Path: "gen-skill", Name: "gen"},
		},
		ReviewerTools: []Entry{
			{Type: "skill", Source: "local", Path: "rev-skill", Name: "rev"},
		},
		ConfigDir: dir,
	})
	
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if report.Failed() {
		t.Fatal("expected no failures")
	}
	
	// Verify role attribution
	genDirs := report.GeneratorSkillDirs()
	revDirs := report.ReviewerSkillDirs()
	
	if len(genDirs) != 1 {
		t.Errorf("expected 1 generator skill dir, got %d", len(genDirs))
	}
	if len(revDirs) != 1 {
		t.Errorf("expected 1 reviewer skill dir, got %d", len(revDirs))
	}
	
	// Verify they don't overlap
	if len(genDirs) > 0 && len(revDirs) > 0 && genDirs[0] == revDirs[0] {
		t.Error("generator and reviewer skill dirs should be different")
	}
}

func TestToolLoadReport_FirstError(t *testing.T) {
	report := &ToolLoadReport{
		Items: []ToolLoadItem{
			{Kind: progress.ToolKindSkill, Name: "good", Status: progress.ToolStatusLoaded},
			{Kind: progress.ToolKindPlugin, Name: "bad", Status: progress.ToolStatusFailed, Reason: "not found"},
			{Kind: progress.ToolKindMCP, Name: "also-bad", Status: progress.ToolStatusFailed, Reason: "missing command"},
		},
	}
	
	err := report.FirstError()
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if err.Kind != progress.ToolKindPlugin {
		t.Errorf("expected first error Kind=plugin, got %q", err.Kind)
	}
	if err.Name != "bad" {
		t.Errorf("expected first error Name=bad, got %q", err.Name)
	}
	if err.Reason != "not found" {
		t.Errorf("expected first error Reason='not found', got %q", err.Reason)
	}
}

func TestRegistryLookup(t *testing.T) {
	// Create temp plugin dir
	dir := t.TempDir()
	pluginYAML := `name: test-plugin
skills:
  - name: test-skill
    path: ./test-skill
`
	if err := os.WriteFile(filepath.Join(dir, "test-plugin.yaml"), []byte(pluginYAML), 0644); err != nil {
		t.Fatal(err)
	}
	
	reg := plugin.NewRegistry()
	if err := reg.LoadDir(dir); err != nil {
		t.Fatal(err)
	}
	
	// Test found
	p, ok := registryLookup(reg, "test-plugin")
	if !ok {
		t.Error("expected to find test-plugin")
	}
	if p == nil {
		t.Error("expected non-nil plugin")
	}
	
	// Test not found
	_, ok = registryLookup(reg, "nonexistent")
	if ok {
		t.Error("expected not to find nonexistent plugin")
	}
	
	// Test nil registry
	_, ok = registryLookup(nil, "test-plugin")
	if ok {
		t.Error("expected not to find plugin with nil registry")
	}
}

func TestValidateAndExpand_GlobExpansion(t *testing.T) {
	dir := t.TempDir()
	
	// Create multiple matching skills
	for _, name := range []string{"skill-a", "skill-b", "skill-c"} {
		skillDir := filepath.Join(dir, name)
		if err := os.Mkdir(skillDir, 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(skillDir, "SKILL.md"), []byte("# "+name), 0644); err != nil {
			t.Fatal(err)
		}
	}
	
	// Create non-matching skill
	otherDir := filepath.Join(dir, "other")
	if err := os.Mkdir(otherDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(otherDir, "SKILL.md"), []byte("# other"), 0644); err != nil {
		t.Fatal(err)
	}
	
	report, err := ValidateAndExpand(context.Background(), ValidationInput{
		GeneratorTools: []Entry{
			{Type: "skill", Source: "local", Path: "skill-*", Name: "glob-skills"},
		},
		ConfigDir: dir,
	})
	
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if report.Failed() {
		t.Fatal("expected no failures")
	}
	
	// Should have 3 matched skills
	matchedCount := 0
	for _, item := range report.Items {
		if item.Kind == progress.ToolKindSkill && item.Status == progress.ToolStatusLoaded && item.ParentKind == progress.ToolParentKindSkillDir {
			matchedCount++
		}
	}
	if matchedCount != 3 {
		t.Errorf("expected 3 glob-matched skills, got %d", matchedCount)
	}
}

// TestValidateAndExpand_RemotePluginMissingRepo verifies that a plugin
// entry with source: remote but no repo: field is rejected with a clear
// error message. The @marketplace shorthand has been removed entirely;
// every remote plugin must declare its source repo explicitly so hyoka
// knows where to fetch it from.
func TestValidateAndExpand_RemotePluginMissingRepo(t *testing.T) {
dir := t.TempDir()

report, err := ValidateAndExpand(context.Background(), ValidationInput{
GeneratorTools: []Entry{
{Type: "plugin", Name: "azure-sdk-python", Source: "remote"},
},
ConfigDir: dir,
})

if err == nil {
t.Fatal("expected error for source: remote plugin without repo: field")
}
toolErr, ok := err.(*ToolLoadError)
if !ok {
t.Fatalf("expected *ToolLoadError, got %T", err)
}
if toolErr.Kind != progress.ToolKindPlugin {
t.Errorf("expected Kind=plugin, got %q", toolErr.Kind)
}
if toolErr.Name != "azure-sdk-python" {
t.Errorf("expected Name=azure-sdk-python, got %q", toolErr.Name)
}
// Reason should mention the fix: add a repo: field.
for _, want := range []string{"repo:", "microsoft/skills"} {
if !strings.Contains(toolErr.Reason, want) {
t.Errorf("expected reason to contain %q, got: %s", want, toolErr.Reason)
}
}
if !report.Failed() {
t.Error("expected report.Failed() == true")
}
}

// TestValidateAndExpand_RemotePluginNameWithAt verifies that the retired
// @marketplace shorthand is rejected with a clear migration message —
// names must be plain identifiers, never name@marketplace.
func TestValidateAndExpand_PluginNameWithAt_Rejected(t *testing.T) {
dir := t.TempDir()

report, err := ValidateAndExpand(context.Background(), ValidationInput{
GeneratorTools: []Entry{
{Type: "plugin", Name: "azure-sdk-python@skills", Source: "remote"},
},
ConfigDir: dir,
})

if err == nil {
t.Fatal("expected error for plugin name containing '@'")
}
toolErr, ok := err.(*ToolLoadError)
if !ok {
t.Fatalf("expected *ToolLoadError, got %T", err)
}
if toolErr.Kind != progress.ToolKindPlugin {
t.Errorf("expected Kind=plugin, got %q", toolErr.Kind)
}
// Reason must explain the @marketplace shorthand was removed and point
// callers at repo:.
for _, want := range []string{"@marketplace shorthand has been removed", "repo:"} {
if !strings.Contains(toolErr.Reason, want) {
t.Errorf("expected reason to contain %q, got: %s", want, toolErr.Reason)
}
}
if !report.Failed() {
t.Error("expected report.Failed() == true")
}
}

func TestReadSkillFrontmatterName(t *testing.T) {
cases := []struct {
name     string
content  string
want     string
wantOk   bool
}{
{
name:    "well-formed frontmatter with name",
content: "---\nname: python\ndescription: foo\n---\n# body\n",
want:    "python",
wantOk:  true,
},
{
name:    "quoted name",
content: "---\nname: \"python-best-practices\"\n---\n",
want:    "python-best-practices",
wantOk:  true,
},
{
name:    "no frontmatter",
content: "# Just a heading\n",
want:    "",
wantOk:  false,
},
{
name:    "frontmatter without name",
content: "---\ndescription: foo\n---\n",
want:    "",
wantOk:  false,
},
}
for _, tc := range cases {
t.Run(tc.name, func(t *testing.T) {
dir := t.TempDir()
if err := os.WriteFile(filepath.Join(dir, "SKILL.md"), []byte(tc.content), 0o644); err != nil {
t.Fatal(err)
}
got, ok := readSkillFrontmatterName(dir)
if got != tc.want || ok != tc.wantOk {
t.Errorf("got (%q, %v), want (%q, %v)", got, ok, tc.want, tc.wantOk)
}
})
}
t.Run("missing SKILL.md", func(t *testing.T) {
dir := t.TempDir()
got, ok := readSkillFrontmatterName(dir)
if got != "" || ok {
t.Errorf("got (%q, %v), want empty", got, ok)
}
})
}
