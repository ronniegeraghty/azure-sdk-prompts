package tool

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/progress"
)

func TestResolveSkillsWithReporter_EmitsStartAndLoaded(t *testing.T) {
	dir := t.TempDir()
	skillDir := filepath.Join(dir, "my-skill")
	if err := os.Mkdir(skillDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(skillDir, "SKILL.md"), []byte("# Skill"), 0644); err != nil {
		t.Fatal(err)
	}

	var events []progress.ProgressEvent
	emit := func(e progress.ProgressEvent) { events = append(events, e) }

	_, err := ResolveSkillsWithReporter(context.Background(), []Entry{
		{Type: "skill", Source: "local", Path: skillDir, Name: "s1"},
	}, dir, emit)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(events) != 2 {
		t.Fatalf("want 2 events (Start+Result), got %d: %+v", len(events), events)
	}
	if events[0].Type != progress.EventToolResolutionStart || events[0].ToolName != "s1" || events[0].ToolKind != progress.ToolKindSkill {
		t.Errorf("unexpected start event: %+v", events[0])
	}
	if events[1].Type != progress.EventToolResolutionResult || events[1].Status != progress.ToolStatusLoaded {
		t.Errorf("unexpected result event: %+v", events[1])
	}
}

func TestResolveSkillsWithReporter_EmitsFailedOnMissingSKILLMD(t *testing.T) {
	dir := t.TempDir()
	skillDir := filepath.Join(dir, "my-skill")
	if err := os.Mkdir(skillDir, 0755); err != nil {
		t.Fatal(err)
	}
	// no SKILL.md

	var events []progress.ProgressEvent
	emit := func(e progress.ProgressEvent) { events = append(events, e) }

	_, err := ResolveSkillsWithReporter(context.Background(), []Entry{
		{Type: "skill", Source: "local", Path: skillDir, Name: "missing"},
	}, dir, emit)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(events) != 2 {
		t.Fatalf("want 2 events, got %d", len(events))
	}
	if events[1].Status != progress.ToolStatusFailed {
		t.Errorf("want Failed for missing SKILL.md, got %q (reason=%q)", events[1].Status, events[1].Reason)
	}
}

func TestResolveSkillsWithReporter_NilEmitSilent(t *testing.T) {
	dir := t.TempDir()
	skillDir := filepath.Join(dir, "my-skill")
	if err := os.Mkdir(skillDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(skillDir, "SKILL.md"), []byte("# Skill"), 0644); err != nil {
		t.Fatal(err)
	}
	// Should not panic and should match ResolveSkills output.
	dirs, err := ResolveSkillsWithReporter(context.Background(), []Entry{
		{Type: "skill", Source: "local", Path: skillDir, Name: "s1"},
	}, dir, nil)
	if err != nil || len(dirs) != 1 {
		t.Fatalf("unexpected: dirs=%v err=%v", dirs, err)
	}
}

func TestEmitMCPResolutions_LocalAndRemote(t *testing.T) {
	var events []progress.ProgressEvent
	emit := func(e progress.ProgressEvent) { events = append(events, e) }

	entries := []Entry{
		{Type: TypeMCP, Name: "local-ok", Command: "npx", Args: []string{"something"}},
		{Type: TypeMCP, Name: "local-bad"}, // missing command
		{Type: TypeMCP, Name: "remote-ok", MCPType: "remote", URL: "http://example"},
		{Type: TypeMCP, Name: "remote-bad", MCPType: "remote"}, // missing url
		{Type: TypeSkill, Name: "not-mcp"},                     // should be skipped
	}
	EmitMCPResolutions(entries, emit)

	// 4 MCP entries × 2 events each = 8
	if len(events) != 8 {
		t.Fatalf("want 8 events, got %d: %+v", len(events), events)
	}
	// Assert result statuses in pairs.
	want := []struct {
		name   string
		status string
	}{
		{"local-ok", progress.ToolStatusLoaded},
		{"local-bad", progress.ToolStatusFailed},
		{"remote-ok", progress.ToolStatusLoaded},
		{"remote-bad", progress.ToolStatusFailed},
	}
	for i, w := range want {
		start := events[i*2]
		result := events[i*2+1]
		if start.Type != progress.EventToolResolutionStart || start.ToolName != w.name || start.ToolKind != progress.ToolKindMCP {
			t.Errorf("events[%d] unexpected start: %+v", i*2, start)
		}
		if result.Type != progress.EventToolResolutionResult || result.ToolName != w.name || result.Status != w.status {
			t.Errorf("events[%d] unexpected result: %+v (want status=%q)", i*2+1, result, w.status)
		}
	}
}

func TestEmitMCPResolutions_NilEmitNoop(t *testing.T) {
	// Must not panic.
	EmitMCPResolutions([]Entry{{Type: TypeMCP, Name: "x"}}, nil)
}
