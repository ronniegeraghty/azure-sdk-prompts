package config

import (
	"strings"
	"testing"
)

// TestParse_PluginTypeEntry_SourceOmitted verifies that `type: plugin`
// entries parse without a `source:` key. The resolver (not the parser)
// is responsible for inferring the default source at validation time;
// at parse time, the field is preserved as empty.
func TestParse_PluginTypeEntry_SourceOmitted(t *testing.T) {
	data := []byte(`
configs:
  - name: src-defaults
    description: "plugin without explicit source"
    generator:
      model: "gpt-4"
      tools:
        - name: foo
          type: plugin
`)
	cfg, err := Parse(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	tools := cfg.Configs[0].Generator.Tools
	if len(tools) != 1 {
		t.Fatalf("expected 1 tool, got %d", len(tools))
	}
	if tools[0].ResolvedType() != "plugin" {
		t.Errorf("expected ResolvedType=plugin, got %q", tools[0].ResolvedType())
	}
	if tools[0].Name != "foo" {
		t.Errorf("expected Name=foo, got %q", tools[0].Name)
	}
	if tools[0].Source != "" {
		t.Errorf("expected Source to be preserved as empty at parse time; got %q", tools[0].Source)
	}
}

// TestParse_PluginTypeEntry_ExplicitLocalSource verifies that an explicit
// `source: local` round-trips through the parser and is preserved on the
// tool entry for the resolver to consume.
func TestParse_PluginTypeEntry_ExplicitLocalSource(t *testing.T) {
	data := []byte(`
configs:
  - name: local-src
    description: "plugin with explicit source: local"
    generator:
      model: "gpt-4"
      tools:
        - name: foo
          type: plugin
          source: local
`)
	cfg, err := Parse(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	tools := cfg.Configs[0].Generator.Tools
	if len(tools) != 1 {
		t.Fatalf("expected 1 tool, got %d", len(tools))
	}
	if tools[0].Source != "local" {
		t.Errorf("expected Source=local, got %q", tools[0].Source)
	}
}

// TestParse_PluginInGeneratorOnly_NotAutoAppendedToReviewer locks in the
// post-migration contract: a plugin declared only under generator.tools
// must NOT appear in reviewer.tools after Parse. The pre-migration
// ExpandPlugins code auto-appended plugin children to both roles; if
// anyone reintroduces that behavior this test fails.
func TestParse_PluginInGeneratorOnly_NotAutoAppendedToReviewer(t *testing.T) {
	data := []byte(`
configs:
  - name: gen-only
    description: "plugin in generator.tools only"
    generator:
      model: "gpt-4"
      tools:
        - name: foo-plugin
          type: plugin
          source: remote
    reviewer:
      models:
        - "gpt-4"
      tools:
        - name: reviewer-skill
          type: skill
          source: local
          path: "./skills/reviewer"
`)
	cfg, err := Parse(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	c := cfg.Configs[0]

	// Generator should contain the plugin entry.
	foundInGenerator := false
	for _, tl := range c.Generator.Tools {
		if tl.Name == "foo-plugin" && tl.ResolvedType() == "plugin" {
			foundInGenerator = true
		}
	}
	if !foundInGenerator {
		t.Error("expected foo-plugin in generator.tools")
	}

	// Reviewer must NOT contain the plugin entry.
	for _, tl := range c.Reviewer.Tools {
		if tl.Name == "foo-plugin" {
			t.Errorf("foo-plugin must not be auto-appended to reviewer.tools; got entry %+v", tl)
		}
		if tl.ResolvedType() == "plugin" {
			t.Errorf("no plugin entries should be auto-added to reviewer.tools; got %+v", tl)
		}
	}
}

// TestParse_PluginInBothRoles_BothPreserved verifies that an operator
// can explicitly list a plugin under BOTH generator.tools AND
// reviewer.tools and both copies survive the parse intact.
func TestParse_PluginInBothRoles_BothPreserved(t *testing.T) {
	data := []byte(`
configs:
  - name: dual-role
    description: "explicit plugin in both roles"
    generator:
      model: "gpt-4"
      tools:
        - name: shared-plugin
          type: plugin
          source: remote
    reviewer:
      models:
        - "gpt-4"
      tools:
        - name: shared-plugin
          type: plugin
          source: remote
`)
	cfg, err := Parse(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	c := cfg.Configs[0]

	genHas := false
	for _, tl := range c.Generator.Tools {
		if tl.Name == "shared-plugin" && tl.ResolvedType() == "plugin" {
			genHas = true
		}
	}
	revHas := false
	for _, tl := range c.Reviewer.Tools {
		if tl.Name == "shared-plugin" && tl.ResolvedType() == "plugin" {
			revHas = true
		}
	}
	if !genHas {
		t.Error("expected shared-plugin in generator.tools")
	}
	if !revHas {
		t.Error("expected shared-plugin in reviewer.tools (explicit dual-role)")
	}
}

// TestParse_RejectsRetiredTopLevelPluginsField_PointsToMigration pins the
// exact migration-hint shape in the rejection error. If anyone softens
// the message (e.g. drops the `type: plugin` / `source:` guidance), this
// test fires.
func TestParse_RejectsRetiredTopLevelPluginsField_PointsToMigration(t *testing.T) {
	data := []byte(`
configs:
  - name: legacy
    description: "legacy schema"
    generator:
      model: "gpt-4"
    plugins:
      - "azure-sdk-python"
`)
	_, err := Parse(data)
	if err == nil {
		t.Fatal("expected error for top-level plugins: field")
	}
	msg := err.Error()
	for _, want := range []string{
		"retired",
		"generator.tools",
		"type: plugin",
		"source:",
	} {
		if !strings.Contains(msg, want) {
			t.Errorf("expected migration hint to contain %q; got: %s", want, msg)
		}
	}
}
