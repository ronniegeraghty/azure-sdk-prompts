package tool

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/progress"
)

// TestResolveSkillsWithReporter_MultiSkillSequentialOrder asserts that for
// each skill entry we see a Start immediately followed by a Result before
// the next skill's Start. Renderers rely on this last-line-only update rule.
func TestResolveSkillsWithReporter_MultiSkillSequentialOrder(t *testing.T) {
	dir := t.TempDir()
	// Two good skills, one missing-SKILL.md so we cover a mix of Loaded/Failed.
	mkSkill := func(name string, withSkillMD bool) string {
		p := filepath.Join(dir, name)
		if err := os.Mkdir(p, 0o755); err != nil {
			t.Fatal(err)
		}
		if withSkillMD {
			if err := os.WriteFile(filepath.Join(p, "SKILL.md"), []byte("# "+name), 0o644); err != nil {
				t.Fatal(err)
			}
		}
		return p
	}
	s1 := mkSkill("alpha", true)
	s2 := mkSkill("beta", false) // no SKILL.md
	s3 := mkSkill("gamma", true)

	var events []progress.ProgressEvent
	emit := func(e progress.ProgressEvent) { events = append(events, e) }

	entries := []Entry{
		{Type: "skill", Source: "local", Path: s1, Name: "alpha"},
		{Type: "skill", Source: "local", Path: s2, Name: "beta"},
		{Type: "skill", Source: "local", Path: s3, Name: "gamma"},
	}
	if _, err := ResolveSkillsWithReporter(context.Background(), entries, dir, emit); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(events) != 6 {
		t.Fatalf("want 6 events (3 pairs), got %d: %+v", len(events), events)
	}
	wantOrder := []struct {
		typ    progress.EventType
		name   string
		status string // empty for Start events
	}{
		{progress.EventToolResolutionStart, "alpha", ""},
		{progress.EventToolResolutionResult, "alpha", progress.ToolStatusLoaded},
		{progress.EventToolResolutionStart, "beta", ""},
		{progress.EventToolResolutionResult, "beta", progress.ToolStatusFailed},
		{progress.EventToolResolutionStart, "gamma", ""},
		{progress.EventToolResolutionResult, "gamma", progress.ToolStatusLoaded},
	}
	for i, w := range wantOrder {
		ev := events[i]
		if ev.Type != w.typ || ev.ToolName != w.name || ev.ToolKind != progress.ToolKindSkill {
			t.Errorf("events[%d] = (%v,%q,%q), want (%v,%q,%q)",
				i, ev.Type, ev.ToolName, ev.ToolKind, w.typ, w.name, progress.ToolKindSkill)
		}
		if w.typ == progress.EventToolResolutionResult && ev.Status != w.status {
			t.Errorf("events[%d] Status = %q, want %q", i, ev.Status, w.status)
		}
		if ev.Status == progress.ToolStatusFailed && ev.Reason == "" {
			t.Errorf("events[%d] Failed result must carry a Reason", i)
		}
	}
}
