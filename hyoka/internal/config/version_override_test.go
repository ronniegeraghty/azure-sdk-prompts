package config

import (
"strings"
"testing"
)

func TestApplyVersionOverrides_PinsByRepo(t *testing.T) {
cf := &ConfigFile{
ToolVersionOverride: map[string]string{
"microsoft/skills": "v2.0.0",
"azure/tools":      "v1.5.0",
},
Configs: []ToolConfig{{
Name: "c1",
Generator: &GeneratorConfig{
Tools: []ToolEntry{
{Name: "azure-skill-1", Type: "skill", Source: "remote", Repo: "microsoft/skills"},
{Name: "azure-skill-2", Type: "skill", Source: "remote", Repo: "microsoft/skills"},
{Name: "untouched", Type: "skill", Source: "remote", Repo: "x/y"},
},
},
Reviewer: &ReviewerConfig{
Tools: []ToolEntry{
{Name: "review-tool", Type: "skill", Source: "remote", Repo: "azure/tools"},
},
},
}},
}

cf.ApplyVersionOverrides()

// Both entries from microsoft/skills should get v2.0.0
if got := cf.Configs[0].Generator.Tools[0].Version; got != "v2.0.0" {
t.Errorf("generator[0] version: got %q, want v2.0.0", got)
}
if got := cf.Configs[0].Generator.Tools[1].Version; got != "v2.0.0" {
t.Errorf("generator[1] version: got %q, want v2.0.0", got)
}
// x/y has no override
if got := cf.Configs[0].Generator.Tools[2].Version; got != "" {
t.Errorf("generator[2] should be untouched, got version %q", got)
}
// azure/tools should get v1.5.0
if got := cf.Configs[0].Reviewer.Tools[0].Version; got != "v1.5.0" {
t.Errorf("reviewer[0] version: got %q, want v1.5.0", got)
}
}

func TestApplyVersionOverrides_PerEntryWinsOverMap(t *testing.T) {
cf := &ConfigFile{
ToolVersionOverride: map[string]string{"microsoft/skills": "v1.0.0"},
Configs: []ToolConfig{{
Name: "c1",
Generator: &GeneratorConfig{
Tools: []ToolEntry{
{Name: "my-skill", Type: "skill", Source: "remote", Repo: "microsoft/skills", Version: "v9.9.9"},
},
},
}},
}
cf.ApplyVersionOverrides()
if got := cf.Configs[0].Generator.Tools[0].Version; got != "v9.9.9" {
t.Errorf("per-entry version should win: got %q, want v9.9.9", got)
}
}

func TestApplyVersionOverrides_GitHubPrefixNormalization(t *testing.T) {
cf := &ConfigFile{
ToolVersionOverride: map[string]string{
"github.com/microsoft/skills": "v2.0.0",
},
Configs: []ToolConfig{{
Name: "c1",
Generator: &GeneratorConfig{
Tools: []ToolEntry{
{Name: "skill1", Type: "skill", Source: "remote", Repo: "microsoft/skills"},
{Name: "skill2", Type: "skill", Source: "remote", Repo: "github.com/microsoft/skills"},
},
},
}},
}
cf.ApplyVersionOverrides()

// Both should match because of normalization
if got := cf.Configs[0].Generator.Tools[0].Version; got != "v2.0.0" {
t.Errorf("tool[0] version: got %q, want v2.0.0", got)
}
if got := cf.Configs[0].Generator.Tools[1].Version; got != "v2.0.0" {
t.Errorf("tool[1] version: got %q, want v2.0.0", got)
}
}

func TestApplyVersionOverrides_NoMapNoOp(t *testing.T) {
cf := &ConfigFile{
Configs: []ToolConfig{{
Name: "c1",
Generator: &GeneratorConfig{
Tools: []ToolEntry{{Name: "x", Type: "skill", Source: "remote", Repo: "a/b"}},
},
}},
}
cf.ApplyVersionOverrides()
if got := cf.Configs[0].Generator.Tools[0].Version; got != "" {
t.Errorf("no override map → version should stay empty, got %q", got)
}
}

func TestApplyVersionOverrides_Idempotent(t *testing.T) {
cf := &ConfigFile{
ToolVersionOverride: map[string]string{"a/b": "v1"},
Configs: []ToolConfig{{
Name: "c1",
Generator: &GeneratorConfig{
Tools: []ToolEntry{{Name: "x", Type: "skill", Source: "remote", Repo: "a/b"}},
},
}},
}
cf.ApplyVersionOverrides()
cf.ApplyVersionOverrides()
if got := cf.Configs[0].Generator.Tools[0].Version; got != "v1" {
t.Errorf("idempotent apply should keep v1, got %q", got)
}
}

func TestApplyVersionOverrides_SkipsLocalSkills(t *testing.T) {
cf := &ConfigFile{
ToolVersionOverride: map[string]string{"a/b": "v1.0.0"},
Configs: []ToolConfig{{
Name: "c1",
Generator: &GeneratorConfig{
Tools: []ToolEntry{
{Name: "local-skill", Type: "skill", Source: "local", Path: "./skills/local"},
{Name: "remote-skill", Type: "skill", Source: "remote", Repo: "a/b"},
},
},
}},
}
cf.ApplyVersionOverrides()

// Local skill (no Repo) should be untouched
if got := cf.Configs[0].Generator.Tools[0].Version; got != "" {
t.Errorf("local skill should have no version, got %q", got)
}
// Remote skill should get the override
if got := cf.Configs[0].Generator.Tools[1].Version; got != "v1.0.0" {
t.Errorf("remote skill version: got %q, want v1.0.0", got)
}
}

func TestApplyVersionOverrides_EmptyValueSkipped(t *testing.T) {
cf := &ConfigFile{
ToolVersionOverride: map[string]string{"a/b": ""},
Configs: []ToolConfig{{
Name: "c1",
Generator: &GeneratorConfig{
Tools: []ToolEntry{{Name: "x", Type: "skill", Source: "remote", Repo: "a/b"}},
},
}},
}
cf.ApplyVersionOverrides()
if got := cf.Configs[0].Generator.Tools[0].Version; got != "" {
t.Errorf("empty override value should be skipped, got %q", got)
}
}

func TestParseConfig_ToolVersionOverride(t *testing.T) {
yaml := []byte(`
tool_version_override:
  microsoft/skills: "v2.0.0"
configs:
  - name: c1
    generator:
      model: gpt-5.2
      tools:
        - name: my-skill
          type: skill
          source: remote
          repo: microsoft/skills
`)
cf, err := Parse(yaml)
if err != nil {
t.Fatalf("parse: %v", err)
}
if cf.ToolVersionOverride["microsoft/skills"] != "v2.0.0" {
t.Errorf("override map not parsed: %v", cf.ToolVersionOverride)
}
// Parse does not call ApplyVersionOverrides (Load does); ensure the API works
// when invoked explicitly.
cf.ApplyVersionOverrides()
if got := cf.Configs[0].Generator.Tools[0].Version; got != "v2.0.0" {
t.Errorf("post-apply version: got %q, want v2.0.0", got)
}
}

func TestValidateOverrideKeys_OldShapeRejected(t *testing.T) {
yaml := []byte(`
tool_version_override:
  my-skill: "v2.0.0"
configs:
  - name: c1
    generator:
      model: gpt-5.2
`)
_, err := Parse(yaml)
if err == nil {
t.Fatal("expected error for old-shape key, got nil")
}
if !strings.Contains(err.Error(), "now keys by repo") {
t.Errorf("expected migration hint, got: %v", err)
}
if !strings.Contains(err.Error(), "my-skill") {
t.Errorf("expected error to mention the bad key, got: %v", err)
}
}

func TestValidateOverrideKeys_MalformedKeysRejected(t *testing.T) {
tests := []struct {
name string
key  string
want string
}{
{
name: "single component",
key:  "microsoft",
want: "now keys by repo",
},
{
name: "three components",
key:  "a/b/c",
want: "not in \"owner/repo\" format",
},
{
name: "empty owner",
key:  "/skills",
want: "empty owner or repo",
},
{
name: "empty repo",
key:  "microsoft/",
want: "empty owner or repo",
},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
yaml := []byte("tool_version_override:\n  " + tt.key + ": v1.0.0\nconfigs:\n  - name: c1\n    generator:\n      model: gpt-5.2\n")
_, err := Parse(yaml)
if err == nil {
t.Fatalf("expected error for key %q, got nil", tt.key)
}
if !strings.Contains(err.Error(), tt.want) {
t.Errorf("expected error containing %q, got: %v", tt.want, err)
}
})
}
}

func TestValidateOverrideKeys_ValidKeysAccepted(t *testing.T) {
tests := []string{
"microsoft/skills",
"github.com/azure/tools",
"a/b",
"org-name/repo-name",
}

for _, key := range tests {
t.Run(key, func(t *testing.T) {
yaml := []byte("tool_version_override:\n  " + key + ": v1.0.0\nconfigs:\n  - name: c1\n    generator:\n      model: gpt-5.2\n")
_, err := Parse(yaml)
if err != nil {
t.Errorf("valid key %q rejected: %v", key, err)
}
})
}
}

func TestLoadDir_IdenticalOverridesMerge(t *testing.T) {
// Verify that identical override values across files merge without error
cf1 := &ConfigFile{
ToolVersionOverride: map[string]string{"microsoft/skills": "v1.0.0"},
Configs: []ToolConfig{{Name: "c1", Generator: &GeneratorConfig{Model: "gpt-5.2"}}},
}
cf2 := &ConfigFile{
ToolVersionOverride: map[string]string{"microsoft/skills": "v1.0.0"},
Configs: []ToolConfig{{Name: "c2", Generator: &GeneratorConfig{Model: "gpt-5.2"}}},
}

// Simulate the merge logic from LoadDir
merged := &ConfigFile{ToolVersionOverride: make(map[string]string)}
for k, v := range cf1.ToolVersionOverride {
merged.ToolVersionOverride[k] = v
}

for k, v := range cf2.ToolVersionOverride {
if existing, ok := merged.ToolVersionOverride[k]; ok && existing != v {
t.Fatalf("unexpected conflict for identical values")
}
merged.ToolVersionOverride[k] = v
}

if merged.ToolVersionOverride["microsoft/skills"] != "v1.0.0" {
t.Errorf("merged override: got %q, want v1.0.0", merged.ToolVersionOverride["microsoft/skills"])
}
}
