package eval

import (
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/progress"
)

// reporter is a minimal test double that captures every ProgressEvent it
// receives. Its zero value is ready to use.
type reporter struct {
	events []progress.ProgressEvent
}

func (r *reporter) emit(evt progress.ProgressEvent) { r.events = append(r.events, evt) }

func (r *reporter) typeCount(t progress.EventType) int {
	n := 0
	for _, e := range r.events {
		if e.Type == t {
			n++
		}
	}
	return n
}

func TestToolVerifier_SkillBasenameDerivation(t *testing.T) {
	// Verifier derives expected skill names from the basename of each
	// resolved directory. Trailing slashes and dots are filtered.
	dirs := []string{"/tmp/skills/alpha", "relative/beta/", "./", "/"}
	v := newToolVerifier(dirs, nil)
	// alpha + beta. './' -> '.' (filtered), '/' -> '/' (filtered).
	if _, ok := v.expectedSkills["alpha"]; !ok {
		t.Errorf("expected skill 'alpha' in set: %v", v.expectedSkills)
	}
	if _, ok := v.expectedSkills["beta"]; !ok {
		t.Errorf("expected skill 'beta' in set: %v", v.expectedSkills)
	}
	if len(v.expectedSkills) != 2 {
		t.Errorf("expected exactly 2 skills, got %d: %v", len(v.expectedSkills), v.expectedSkills)
	}
}

func TestToolVerifier_EmitsOnceAfterBothKindsConfigured(t *testing.T) {
	v := newToolVerifier([]string{"/skills/s1"}, map[string]bool{"mcp1": true})
	if got := v.emitIfReady(); got != nil {
		t.Fatalf("premature emit before any load event: %+v", got)
	}
	v.onSkillsLoaded([]string{"s1"})
	if got := v.emitIfReady(); got != nil {
		t.Fatalf("premature emit after skills only: %+v", got)
	}
	v.onMCPLoaded([]string{"mcp1"})
	first := v.emitIfReady()
	if len(first) != 2 {
		t.Fatalf("want 2 tools, got %d: %+v", len(first), first)
	}
	// Second call must be a no-op (at-most-once).
	if got := v.emitIfReady(); got != nil {
		t.Errorf("verifier emitted twice; want at-most-once: %+v", got)
	}
}

func TestToolVerifier_SingleKindOnlyFiresOnThatEvent(t *testing.T) {
	cases := []struct {
		name         string
		skills       []string
		mcps         map[string]bool
		fireSkills   bool
		fireMCP      bool
		wantEmit     bool
		wantToolKind string
	}{
		{"skills only", []string{"/s/a"}, nil, true, false, true, progress.ToolKindSkill},
		{"mcp only", nil, map[string]bool{"m1": true}, false, true, true, progress.ToolKindMCP},
		{"skills only but MCP event fired", []string{"/s/a"}, nil, false, true, false, ""},
		{"neither configured", nil, nil, true, true, false, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			v := newToolVerifier(tc.skills, tc.mcps)
			if tc.fireSkills {
				v.onSkillsLoaded([]string{"a"})
			}
			if tc.fireMCP {
				v.onMCPLoaded([]string{"m1"})
			}
			got := v.emitIfReady()
			if tc.wantEmit && got == nil {
				t.Fatalf("want emit, got nil")
			}
			if !tc.wantEmit && got != nil {
				t.Fatalf("want no emit, got %+v", got)
			}
			if tc.wantEmit {
				for _, ts := range got {
					if ts.ToolKind != tc.wantToolKind {
						t.Errorf("unexpected tool kind %q (want %q) in %+v", ts.ToolKind, tc.wantToolKind, ts)
					}
				}
			}
		})
	}
}

func TestToolVerifier_MissingMCPMarkedFailed(t *testing.T) {
	// SDK reports mcp1 loaded but the caller configured mcp1 + mcp2.
	v := newToolVerifier(nil, map[string]bool{"mcp1": true, "mcp2": true})
	v.onMCPLoaded([]string{"mcp1"})
	got := v.emitIfReady()
	if len(got) != 2 {
		t.Fatalf("want 2 tools, got %d: %+v", len(got), got)
	}
	byName := map[string]progress.ToolStatus{}
	for _, ts := range got {
		byName[ts.ToolName] = ts
	}
	if byName["mcp1"].Status != progress.ToolStatusLoaded {
		t.Errorf("mcp1 should be Loaded, got %+v", byName["mcp1"])
	}
	if byName["mcp2"].Status != progress.ToolStatusFailed {
		t.Errorf("mcp2 should be Failed, got %+v", byName["mcp2"])
	}
	if byName["mcp2"].Reason == "" {
		t.Errorf("mcp2 Failed status must carry a Reason, got %q", byName["mcp2"].Reason)
	}
}

func TestToolVerifier_MissingSkillMarkedFailed(t *testing.T) {
	v := newToolVerifier([]string{"/skills/alpha", "/skills/beta"}, nil)
	// SDK only reports alpha; beta is missing.
	v.onSkillsLoaded([]string{"alpha"})
	got := v.emitIfReady()
	if len(got) != 2 {
		t.Fatalf("want 2 skills, got %d: %+v", len(got), got)
	}
	byName := map[string]progress.ToolStatus{}
	for _, ts := range got {
		byName[ts.ToolName] = ts
		if ts.ToolKind != progress.ToolKindSkill {
			t.Errorf("unexpected kind on %+v", ts)
		}
	}
	if byName["alpha"].Status != progress.ToolStatusLoaded {
		t.Errorf("alpha should be Loaded, got %+v", byName["alpha"])
	}
	if byName["beta"].Status != progress.ToolStatusFailed || byName["beta"].Reason == "" {
		t.Errorf("beta should be Failed with Reason, got %+v", byName["beta"])
	}
}

func TestToolVerifier_UnknownSDKNamesIgnored(t *testing.T) {
	// SDK reports extras that the caller never configured. Contract (round-1-2):
	// unconfigured extras are dropped — the event answers "did what I asked
	// for load?", not "what did the SDK start?".
	v := newToolVerifier([]string{"/skills/alpha"}, map[string]bool{"mcp1": true})
	v.onSkillsLoaded([]string{"alpha", "bonus-skill"})
	v.onMCPLoaded([]string{"mcp1", "extra-mcp"})
	got := v.emitIfReady()
	if len(got) != 2 {
		t.Fatalf("want 2 tools (configured only), got %d: %+v", len(got), got)
	}
	for _, ts := range got {
		if ts.ToolName == "bonus-skill" || ts.ToolName == "extra-mcp" {
			t.Errorf("unconfigured SDK extra leaked into ToolsVerified: %+v", ts)
		}
	}
}

func TestToolVerifier_DeterministicSortOrder(t *testing.T) {
	// Contract: sorted by (kind, name) ascending. "mcp" < "skill" alphabetically.
	v := newToolVerifier(
		[]string{"/skills/zebra", "/skills/alpha"},
		map[string]bool{"zulu": true, "alpha-mcp": true},
	)
	v.onSkillsLoaded([]string{"alpha", "zebra"})
	v.onMCPLoaded([]string{"alpha-mcp", "zulu"})
	got := v.emitIfReady()
	wantOrder := []struct {
		kind, name string
	}{
		{progress.ToolKindMCP, "alpha-mcp"},
		{progress.ToolKindMCP, "zulu"},
		{progress.ToolKindSkill, "alpha"},
		{progress.ToolKindSkill, "zebra"},
	}
	if len(got) != len(wantOrder) {
		t.Fatalf("len=%d want=%d; got=%+v", len(got), len(wantOrder), got)
	}
	for i, w := range wantOrder {
		if got[i].ToolKind != w.kind || got[i].ToolName != w.name {
			t.Errorf("pos %d: got (%q,%q) want (%q,%q)", i, got[i].ToolKind, got[i].ToolName, w.kind, w.name)
		}
	}
}

func TestToolVerifier_NeitherConfigured_NoEmitEver(t *testing.T) {
	v := newToolVerifier(nil, nil)
	v.onSkillsLoaded([]string{"ghost"})
	v.onMCPLoaded([]string{"ghost"})
	if got := v.emitIfReady(); got != nil {
		t.Errorf("verifier emitted despite nothing configured: %+v", got)
	}
}

// TestToolVerifier_EmitIsSeparatedFromStateMutation documents that
// emitIfReady builds its output slice without further mutating the state
// maps — callers can safely hold a lock while calling it and then invoke
// progressFn after releasing the lock without racing on the internal maps.
// This mirrors the guarantee in .squad/decisions.md (round 1-2
// tool-verification wiring, point 5: "emitToolsVerified builds the slice
// under lock but invokes progressFn post-unlock — no deadlock risk").
func TestToolVerifier_EmitIsSeparatedFromStateMutation(t *testing.T) {
	v := newToolVerifier([]string{"/skills/alpha"}, map[string]bool{"mcp1": true})
	v.onSkillsLoaded([]string{"alpha"})
	v.onMCPLoaded([]string{"mcp1"})

	// Snapshot the maps before emit.
	loadedSkillsBefore := len(v.loadedSkills)
	loadedMCPBefore := len(v.loadedMCP)

	got := v.emitIfReady()
	if got == nil {
		t.Fatal("expected emit, got nil")
	}
	if len(v.loadedSkills) != loadedSkillsBefore || len(v.loadedMCP) != loadedMCPBefore {
		t.Error("emitIfReady unexpectedly mutated the loaded maps")
	}
}
