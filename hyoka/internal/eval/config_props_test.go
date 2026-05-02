package eval

import (
"testing"

"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
"github.com/ronniegeraghty/hyoka/hyoka/internal/config/tool"
"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria"
"github.com/ronniegeraghty/hyoka/hyoka/internal/report"
)

func TestInjectConfigProps(t *testing.T) {
tests := []struct {
name string
cfg  config.ToolConfig
want map[string]string
}{
{
name: "baseline fields",
cfg: config.ToolConfig{
Name: "baseline/claude-opus-4.6",
Generator: &config.GeneratorConfig{
Model: "claude-opus-4.6",
},
},
want: map[string]string{
"config":    "baseline/claude-opus-4.6",
"generator": "claude-opus-4.6",
},
},
{
name: "nil generator",
cfg: config.ToolConfig{
Name: "baseline/claude-opus-4.6",
},
want: map[string]string{
"config": "baseline/claude-opus-4.6",
},
},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
props := make(map[string]string)
injectConfigProps(props, tt.cfg)
for k, want := range tt.want {
if got := props[k]; got != want {
t.Errorf("props[%q] = %q, want %q", k, got, want)
}
}
if len(props) != len(tt.want) {
t.Errorf("got %d keys, want %d: %v", len(props), len(tt.want), props)
}
})
}
}

func TestBuildToolIdentities(t *testing.T) {
tests := []struct {
name string
cfg  config.ToolConfig
want []criteria.ToolIdentity
}{
{
name: "nil generator",
cfg:  config.ToolConfig{},
want: nil,
},
{
name: "skill",
cfg: config.ToolConfig{
Generator: &config.GeneratorConfig{
Tools: []config.ToolEntry{
{Name: "markdown-headings", Type: tool.TypeSkill},
},
},
},
want: []criteria.ToolIdentity{
{Name: "markdown-headings", Source: tool.TypeSkill},
},
},
{
name: "mcp server",
cfg: config.ToolConfig{
Generator: &config.GeneratorConfig{
Tools: []config.ToolEntry{
{Name: "azure", Type: tool.TypeMCP},
},
},
},
want: []criteria.ToolIdentity{
{Name: "azure", Source: tool.TypeMCP, MCPServer: "azure"},
},
},
{
name: "mixed tools",
cfg: config.ToolConfig{
Generator: &config.GeneratorConfig{
Tools: []config.ToolEntry{
{Name: "markdown-headings", Type: tool.TypeSkill},
{Name: "azure", Type: tool.TypeMCP},
},
},
},
want: []criteria.ToolIdentity{
{Name: "markdown-headings", Source: tool.TypeSkill},
{Name: "azure", Source: tool.TypeMCP, MCPServer: "azure"},
},
},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
got := buildToolIdentities(tt.cfg, nil)
if !toolIdentitiesEqual(got, tt.want) {
t.Errorf("buildToolIdentities() mismatch:\ngot:  %+v\nwant: %+v", got, tt.want)
}
})
}
}

func TestBuildToolIdentities_EnvDerivedSkills(t *testing.T) {
// When env is provided, leaf skills come from env.SkillsLoaded
// (the SDK-resolved set, with skill_dir expansion + pairwise
// exclusions already applied). Top-level skill_dir entries on cfg
// should NOT also appear as identities — env wins.
cfg := config.ToolConfig{
Generator: &config.GeneratorConfig{
Tools: []config.ToolEntry{
{Name: "test-skills", Type: tool.TypeSkill}, // skill_dir wrapper
{Name: "filesystem", Type: tool.TypeTool},   // built-in passes through
},
},
}
env := &report.EnvironmentInfo{
SkillsLoaded: []string{"markdown-headings", "markdown-lists"},
MCPServers:   []string{"azure"},
}
got := buildToolIdentities(cfg, env)
want := []criteria.ToolIdentity{
{Name: "markdown-headings", Source: tool.TypeSkill},
{Name: "markdown-lists", Source: tool.TypeSkill},
{Name: "azure", Source: tool.TypeMCP, MCPServer: "azure"},
{Name: "filesystem", Source: tool.TypeTool},
}
if !toolIdentitiesEqual(got, want) {
t.Errorf("buildToolIdentities(cfg, env) mismatch:\ngot:  %+v\nwant: %+v", got, want)
}
}

func toolIdentitiesEqual(a, b []criteria.ToolIdentity) bool {
if len(a) != len(b) {
return false
}
for i := range a {
if a[i].Name != b[i].Name ||
a[i].Source != b[i].Source ||
a[i].MCPServer != b[i].MCPServer {
return false
}
}
return true
}
