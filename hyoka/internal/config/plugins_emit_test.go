package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config/tool"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/progress"
)

// TestEmitPluginResolutions_MissingPluginIsFailed asserts that a plugin name
// with no registry entry and no installed-plugins dir is reported as Failed
// with a non-empty Reason. This is the case renderers care about most — it
// drives the "❌ Failed to load (not found)" line in the Tools block.
func TestEmitPluginResolutions_MissingPluginIsFailed(t *testing.T) {
	// Isolate from the user's ~/.copilot/installed-plugins/ so the
	// fallback resolver can't accidentally find the plugin name.
	t.Setenv("HOME", t.TempDir())
	// Also cd to a clean dir so the registry's "./plugins" fallback finds
	// no YAMLs.
	origWD, _ := os.Getwd()
	t.Cleanup(func() { _ = os.Chdir(origWD) })
	if err := os.Chdir(t.TempDir()); err != nil {
		t.Fatal(err)
	}

	var events []progress.ProgressEvent
	emit := func(e progress.ProgressEvent) { events = append(events, e) }

	cfg := &ToolConfig{
		Name:    "test",
		Plugins: []string{"does-not-exist-plugin"},
	}
	cfg.EmitPluginResolutions(emit)

	if len(events) != 2 {
		t.Fatalf("want 2 events (Start+Result), got %d: %+v", len(events), events)
	}
	start := events[0]
	result := events[1]
	if start.Type != progress.EventToolResolutionStart {
		t.Errorf("events[0] Type = %v, want EventToolResolutionStart", start.Type)
	}
	if start.ToolKind != progress.ToolKindPlugin {
		t.Errorf("events[0] Kind = %q, want %q", start.ToolKind, progress.ToolKindPlugin)
	}
	if start.ToolName != "does-not-exist-plugin" {
		t.Errorf("events[0] Name = %q", start.ToolName)
	}
	if result.Type != progress.EventToolResolutionResult {
		t.Errorf("events[1] Type = %v", result.Type)
	}
	if result.Status != progress.ToolStatusFailed {
		t.Errorf("events[1] Status = %q, want Failed", result.Status)
	}
	if result.Reason == "" {
		t.Errorf("events[1] Reason must be non-empty for Failed plugin")
	}
}

// TestEmitPluginResolutions_FoundInRegistryIsLoaded writes a minimal plugin
// YAML, points the config loader at it, and asserts the emission path
// reports Loaded.
func TestEmitPluginResolutions_FoundInRegistryIsLoaded(t *testing.T) {
	baseDir := t.TempDir()
	pluginsDir := filepath.Join(baseDir, "plugins")
	if err := os.Mkdir(pluginsDir, 0o755); err != nil {
		t.Fatal(err)
	}
	// Minimal plugin YAML recognised by the registry loader.
	pluginYAML := `
name: demo-plugin
description: test
skills:
  - type: local
    name: demo-skill
    path: ./skill
`
	if err := os.WriteFile(filepath.Join(pluginsDir, "demo-plugin.yaml"), []byte(pluginYAML), 0o644); err != nil {
		t.Fatal(err)
	}

	// EmitPluginResolutions uses DiscoverFromCWD() -> ResolveCandidates("plugins", ...).
	// Simplest way to steer it: chdir into baseDir so "./plugins" resolves here.
	origWD, _ := os.Getwd()
	t.Cleanup(func() { _ = os.Chdir(origWD) })
	if err := os.Chdir(baseDir); err != nil {
		t.Fatal(err)
	}
	// Also isolate HOME so installed-plugin fallback can't inject noise.
	t.Setenv("HOME", t.TempDir())

	var events []progress.ProgressEvent
	emit := func(e progress.ProgressEvent) { events = append(events, e) }

	cfg := &ToolConfig{
		Name:    "test",
		Plugins: []string{"demo-plugin"},
	}
	cfg.EmitPluginResolutions(emit)

	if len(events) != 2 {
		t.Fatalf("want 2 events, got %d: %+v", len(events), events)
	}
	if events[1].Status != progress.ToolStatusLoaded {
		t.Errorf("events[1] Status = %q, want Loaded (reason=%q)", events[1].Status, events[1].Reason)
	}
	if events[1].ToolKind != progress.ToolKindPlugin {
		t.Errorf("events[1] Kind = %q, want %q", events[1].ToolKind, progress.ToolKindPlugin)
	}
}

func TestEmitPluginResolutions_NilEmitNoop(t *testing.T) {
	// Must not panic with nil emit.
	cfg := &ToolConfig{Plugins: []string{"anything"}}
	cfg.EmitPluginResolutions(nil)
}

func TestEmitPluginResolutions_NoPluginsIsNoop(t *testing.T) {
	// Config with no plugins produces no events.
	var calls int
	emit := tool.ProgressEmitter(func(progress.ProgressEvent) { calls++ })
	cfg := &ToolConfig{}
	cfg.EmitPluginResolutions(emit)
	if calls != 0 {
		t.Errorf("want 0 emit calls, got %d", calls)
	}
}

// TestEmitPluginResolutions_MultipleSequentialOrder asserts that when a
// config declares multiple plugins, every Start is paired with its Result
// before the next plugin's Start. Same last-line-only rule as skills.
func TestEmitPluginResolutions_MultipleSequentialOrder(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	origWD, _ := os.Getwd()
	t.Cleanup(func() { _ = os.Chdir(origWD) })
	if err := os.Chdir(t.TempDir()); err != nil {
		t.Fatal(err)
	}

	var events []progress.ProgressEvent
	emit := func(e progress.ProgressEvent) { events = append(events, e) }

	cfg := &ToolConfig{
		Plugins: []string{"p-one", "p-two", "p-three"},
	}
	cfg.EmitPluginResolutions(emit)

	if len(events) != 6 {
		t.Fatalf("want 6 events, got %d: %+v", len(events), events)
	}
	wantOrder := []struct {
		typ  progress.EventType
		name string
	}{
		{progress.EventToolResolutionStart, "p-one"},
		{progress.EventToolResolutionResult, "p-one"},
		{progress.EventToolResolutionStart, "p-two"},
		{progress.EventToolResolutionResult, "p-two"},
		{progress.EventToolResolutionStart, "p-three"},
		{progress.EventToolResolutionResult, "p-three"},
	}
	for i, w := range wantOrder {
		if events[i].Type != w.typ || events[i].ToolName != w.name {
			t.Errorf("events[%d] = (%v,%q), want (%v,%q)", i, events[i].Type, events[i].ToolName, w.typ, w.name)
		}
		if events[i].ToolKind != progress.ToolKindPlugin {
			t.Errorf("events[%d] Kind = %q", i, events[i].ToolKind)
		}
	}
}
