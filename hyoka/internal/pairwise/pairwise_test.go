package pairwise

import (
	"testing"

	"github.com/ronniegeraghty/hyoka/internal/config"
)

func TestExpandPairwise_ThreeTools(t *testing.T) {
	base := config.ToolConfig{
		Name: "test-cfg",
		Generator: &config.GeneratorConfig{
			Model: "model-a",
			Tools: []config.ToolEntry{
				{Name: "tool-a"},
				{Name: "tool-b"},
				{Name: "tool-c"},
			},
		},
	}

	variants := ExpandPairwise(base)

	if got := len(variants); got != 4 {
		t.Fatalf("expected 4 variants, got %d", got)
	}

	// Variant 0: baseline with all tools.
	assertName(t, variants[0], "test-cfg/baseline")
	assertToolNames(t, variants[0], []string{"tool-a", "tool-b", "tool-c"})

	// Variant 1: without tool-a.
	assertName(t, variants[1], "test-cfg/without-tool-a")
	assertToolNames(t, variants[1], []string{"tool-b", "tool-c"})

	// Variant 2: without tool-b.
	assertName(t, variants[2], "test-cfg/without-tool-b")
	assertToolNames(t, variants[2], []string{"tool-a", "tool-c"})

	// Variant 3: without tool-c.
	assertName(t, variants[3], "test-cfg/without-tool-c")
	assertToolNames(t, variants[3], []string{"tool-a", "tool-b"})
}

func TestExpandPairwise_AlwaysOnExemption(t *testing.T) {
	base := config.ToolConfig{
		Name: "with-always-on",
		Generator: &config.GeneratorConfig{
			Model: "model-a",
			Tools: []config.ToolEntry{
				{Name: "core-tool", AlwaysOn: true},
				{Name: "optional-a"},
				{Name: "optional-b"},
			},
		},
	}

	variants := ExpandPairwise(base)

	// 2 togglable tools → 3 variants (baseline + 2).
	if got := len(variants); got != 3 {
		t.Fatalf("expected 3 variants, got %d", got)
	}

	// Baseline keeps all three.
	assertToolNames(t, variants[0], []string{"core-tool", "optional-a", "optional-b"})

	// Variant 1: without optional-a; core-tool still present.
	assertName(t, variants[1], "with-always-on/without-optional-a")
	assertToolNames(t, variants[1], []string{"core-tool", "optional-b"})

	// Variant 2: without optional-b; core-tool still present.
	assertName(t, variants[2], "with-always-on/without-optional-b")
	assertToolNames(t, variants[2], []string{"core-tool", "optional-a"})
}

func TestExpandPairwise_ZeroTools(t *testing.T) {
	base := config.ToolConfig{
		Name: "no-tools",
		Generator: &config.GeneratorConfig{
			Model: "model-a",
		},
	}

	variants := ExpandPairwise(base)

	if got := len(variants); got != 1 {
		t.Fatalf("expected 1 variant (baseline only), got %d", got)
	}
	assertName(t, variants[0], "no-tools/baseline")
}

func TestExpandPairwise_SingleTool(t *testing.T) {
	base := config.ToolConfig{
		Name: "single",
		Generator: &config.GeneratorConfig{
			Model: "model-a",
			Tools: []config.ToolEntry{
				{Name: "only-tool"},
			},
		},
	}

	variants := ExpandPairwise(base)

	if got := len(variants); got != 2 {
		t.Fatalf("expected 2 variants, got %d", got)
	}
	assertName(t, variants[0], "single/baseline")
	assertToolNames(t, variants[0], []string{"only-tool"})

	assertName(t, variants[1], "single/without-only-tool")
	assertToolNames(t, variants[1], nil)
}

func TestExpandPairwise_AvailableTools(t *testing.T) {
	base := config.ToolConfig{
		Name: "avail",
		Generator: &config.GeneratorConfig{
			Model: "model-a",
			AvailableTools: []string{"fetch", "search", "run"},
		},
	}

	variants := ExpandPairwise(base)

	if got := len(variants); got != 4 {
		t.Fatalf("expected 4 variants, got %d", got)
	}

	assertName(t, variants[0], "avail/baseline")
	assertAvailableTools(t, variants[0], []string{"fetch", "search", "run"})

	assertName(t, variants[1], "avail/without-fetch")
	assertAvailableTools(t, variants[1], []string{"search", "run"})

	assertName(t, variants[2], "avail/without-search")
	assertAvailableTools(t, variants[2], []string{"fetch", "run"})

	assertName(t, variants[3], "avail/without-run")
	assertAvailableTools(t, variants[3], []string{"fetch", "search"})
}

func TestExpandPairwise_MixedToolSources(t *testing.T) {
	base := config.ToolConfig{
		Name: "mixed",
		Generator: &config.GeneratorConfig{
			Model: "model-a",
			Tools: []config.ToolEntry{
				{Name: "structured-tool"},
			},
			AvailableTools: []string{"flat-tool"},
		},
	}

	variants := ExpandPairwise(base)

	if got := len(variants); got != 3 {
		t.Fatalf("expected 3 variants, got %d", got)
	}

	// Baseline has both.
	assertName(t, variants[0], "mixed/baseline")
	assertToolNames(t, variants[0], []string{"structured-tool"})
	assertAvailableTools(t, variants[0], []string{"flat-tool"})

	// Without structured-tool.
	assertName(t, variants[1], "mixed/without-structured-tool")
	assertToolNames(t, variants[1], nil)
	assertAvailableTools(t, variants[1], []string{"flat-tool"})

	// Without flat-tool.
	assertName(t, variants[2], "mixed/without-flat-tool")
	assertToolNames(t, variants[2], []string{"structured-tool"})
	assertAvailableTools(t, variants[2], nil)
}

func TestExpandPairwise_DeduplicatesAcrossSources(t *testing.T) {
	base := config.ToolConfig{
		Name: "dedup",
		Generator: &config.GeneratorConfig{
			Model:          "model-a",
			Tools:          []config.ToolEntry{{Name: "shared"}},
			AvailableTools: []string{"shared", "unique"},
		},
	}

	variants := ExpandPairwise(base)

	// "shared" appears in both sources but should produce one toggle, not two.
	// Togglable: shared, unique → 3 variants.
	if got := len(variants); got != 3 {
		t.Fatalf("expected 3 variants, got %d", got)
	}
	assertName(t, variants[1], "dedup/without-shared")
	assertName(t, variants[2], "dedup/without-unique")
}

func TestExpandPairwise_NilGenerator(t *testing.T) {
	base := config.ToolConfig{
		Name: "nil-gen",
	}

	variants := ExpandPairwise(base)

	if got := len(variants); got != 1 {
		t.Fatalf("expected 1 variant, got %d", got)
	}
	assertName(t, variants[0], "nil-gen/baseline")
}

func TestExpandPairwise_CloneIsolation(t *testing.T) {
	base := config.ToolConfig{
		Name: "isolation",
		Generator: &config.GeneratorConfig{
			Model: "model-a",
			Tools: []config.ToolEntry{
				{Name: "alpha"},
				{Name: "beta"},
			},
		},
	}

	variants := ExpandPairwise(base)

	// Mutate a variant's tools and verify the original is untouched.
	variants[1].Generator.Tools = append(variants[1].Generator.Tools, config.ToolEntry{Name: "injected"})

	if len(base.Generator.Tools) != 2 {
		t.Fatalf("original base was mutated: got %d tools, want 2", len(base.Generator.Tools))
	}
	// Also verify other variants are unaffected.
	if len(variants[0].Generator.Tools) != 2 {
		t.Fatalf("baseline variant was mutated: got %d tools, want 2", len(variants[0].Generator.Tools))
	}
}

// --- helpers ---

func assertName(t *testing.T, v config.ToolConfig, want string) {
	t.Helper()
	if v.Name != want {
		t.Errorf("name = %q, want %q", v.Name, want)
	}
}

func assertToolNames(t *testing.T, v config.ToolConfig, want []string) {
	t.Helper()
	var got []string
	if v.Generator != nil {
		for _, te := range v.Generator.Tools {
			got = append(got, te.Name)
		}
	}
	assertStringSlice(t, "Tools", got, want)
}

func assertAvailableTools(t *testing.T, v config.ToolConfig, want []string) {
	t.Helper()
	var got []string
	if v.Generator != nil {
		got = v.Generator.AvailableTools
	}
	assertStringSlice(t, "AvailableTools", got, want)
}

func assertStringSlice(t *testing.T, label string, got, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Errorf("%s length = %d, want %d; got %v, want %v", label, len(got), len(want), got, want)
		return
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("%s[%d] = %q, want %q", label, i, got[i], want[i])
		}
	}
}
