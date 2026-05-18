package eval

import (
	"context"
	"strings"
	"testing"
	"time"

	tool "github.com/ronniegeraghty/hyoka/hyoka/internal/config/tool"
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
	// Second call must return the same cached result so post-session callers
	// can still retrieve the verified tool list after an in-line emit.
	second := v.emitIfReady()
	if len(second) != len(first) {
		t.Errorf("verifier did not return cached result on second emit: got %d want %d", len(second), len(first))
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

// --- Tool Validation Gate Tests (WU-2) ---
// These tests validate the enforcement gate that checks tool load status
// after session creation and before sending the prompt (per Neo's fix plan).

// TestToolValidationGate_HappyPath validates that when all expected tools
// report as loaded, the eval proceeds normally without errors.
func TestToolValidationGate_HappyPath(t *testing.T) {
	v := newToolVerifier(
		[]string{"/skills/generator-skills"},
		map[string]bool{"azure-mcp": true},
	)
	v.onSkillsLoaded([]string{"generator-skills"})
	v.onMCPLoaded([]string{"azure-mcp"})

	tools := v.emitIfReady()
	if tools == nil {
		t.Fatal("expected tools to be emitted")
	}

	// Validate that all tools show as loaded
	for _, tool := range tools {
		if tool.Status != progress.ToolStatusLoaded {
			t.Errorf("tool %s (%s) should be loaded, got status=%s reason=%q",
				tool.ToolName, tool.ToolKind, tool.Status, tool.Reason)
		}
	}
}

// TestToolValidationGate_SkillLoadFailure validates that when an expected
// skill reports as Failed, the validation gate should detect it.
func TestToolValidationGate_SkillLoadFailure(t *testing.T) {
	v := newToolVerifier(
		[]string{"/skills/generator-skills", "/skills/helper"},
		nil,
	)
	// SDK only reports generator-skills; helper is missing
	v.onSkillsLoaded([]string{"generator-skills"})

	tools := v.emitIfReady()
	if tools == nil {
		t.Fatal("expected tools to be emitted")
	}

	// Verify that the missing skill is marked as failed
	failures := 0
	for _, tool := range tools {
		if tool.ToolName == "helper" && tool.Status == progress.ToolStatusFailed {
			failures++
			if tool.Reason == "" {
				t.Error("failed tool should have a Reason explaining why")
			}
		}
	}

	if failures == 0 {
		t.Error("expected helper skill to be marked as failed")
	}
}

// TestToolValidationGate_MCPLoadFailure validates that when an expected
// MCP server reports as Failed, the validation gate should detect it.
func TestToolValidationGate_MCPLoadFailure(t *testing.T) {
	v := newToolVerifier(
		nil,
		map[string]bool{"azure-mcp": true, "playwright-mcp": true},
	)
	// SDK only reports azure-mcp; playwright-mcp is missing
	v.onMCPLoaded([]string{"azure-mcp"})

	tools := v.emitIfReady()
	if tools == nil {
		t.Fatal("expected tools to be emitted")
	}

	// Verify that the missing MCP server is marked as failed
	failures := 0
	for _, tool := range tools {
		if tool.ToolName == "playwright-mcp" && tool.Status == progress.ToolStatusFailed {
			failures++
			if tool.Reason == "" {
				t.Error("failed MCP server should have a Reason explaining why")
			}
		}
	}

	if failures == 0 {
		t.Error("expected playwright-mcp to be marked as failed")
	}
}

// TestToolValidationGate_MixedFailure validates that when multiple tools
// are configured and some fail, all failures are reported in the tools slice.
func TestToolValidationGate_MixedFailure(t *testing.T) {
	v := newToolVerifier(
		[]string{"/skills/alpha", "/skills/beta", "/skills/gamma"},
		map[string]bool{"mcp1": true, "mcp2": true},
	)
	// SDK reports: alpha loaded, beta missing, gamma loaded, mcp1 missing, mcp2 loaded
	v.onSkillsLoaded([]string{"alpha", "gamma"})
	v.onMCPLoaded([]string{"mcp2"})

	tools := v.emitIfReady()
	if tools == nil {
		t.Fatal("expected tools to be emitted")
	}

	// Map tools by name for easy checking
	byName := make(map[string]progress.ToolStatus)
	for _, tool := range tools {
		byName[tool.ToolName] = tool
	}

	// Verify loaded tools
	if byName["alpha"].Status != progress.ToolStatusLoaded {
		t.Errorf("alpha should be loaded, got %+v", byName["alpha"])
	}
	if byName["gamma"].Status != progress.ToolStatusLoaded {
		t.Errorf("gamma should be loaded, got %+v", byName["gamma"])
	}
	if byName["mcp2"].Status != progress.ToolStatusLoaded {
		t.Errorf("mcp2 should be loaded, got %+v", byName["mcp2"])
	}

	// Verify failed tools
	if byName["beta"].Status != progress.ToolStatusFailed {
		t.Errorf("beta should be failed, got %+v", byName["beta"])
	}
	if byName["mcp1"].Status != progress.ToolStatusFailed {
		t.Errorf("mcp1 should be failed, got %+v", byName["mcp1"])
	}

	// Count failures (should be 2: beta and mcp1)
	failures := 0
	for _, tool := range tools {
		if tool.Status == progress.ToolStatusFailed {
			failures++
		}
	}
	if failures != 2 {
		t.Errorf("expected 2 failures, got %d", failures)
	}
}

// TestToolValidationGate_NoExpectedTools validates that when no tools are
// configured, the verifier never emits and the validation gate is skipped.
func TestToolValidationGate_NoExpectedTools(t *testing.T) {
	v := newToolVerifier(nil, nil)

	// Even if SDK fires events, verifier shouldn't emit
	v.onSkillsLoaded([]string{"unexpected-skill"})
	v.onMCPLoaded([]string{"unexpected-mcp"})

	tools := v.emitIfReady()
	if tools != nil {
		t.Errorf("verifier should not emit when nothing configured, got %+v", tools)
	}
}

// TestToolValidationGate_TimeoutScenario validates the verifier's behavior
// when SDK events never arrive. This tests the "timeout" path where the
// validation gate would need to abort after waiting.
func TestToolValidationGate_TimeoutScenario(t *testing.T) {
	v := newToolVerifier(
		[]string{"/skills/alpha"},
		map[string]bool{"mcp1": true},
	)

	// Simulate: SDK never fires the load events
	// Call emitIfReady multiple times to simulate polling with timeout
	for i := 0; i < 5; i++ {
		tools := v.emitIfReady()
		if tools != nil {
			t.Fatalf("verifier emitted before receiving events: %+v", tools)
		}
	}

	// After timeout expires, the validation gate should detect this as a failure.
	// The actual timeout logic will be in the waitForToolVerification helper
	// that Neo implements, but this tests that the verifier doesn't emit
	// prematurely.
}

// TestToolValidationGate_PartialEventArrival validates that the verifier
// doesn't emit until ALL expected kinds have reported their events.
func TestToolValidationGate_PartialEventArrival(t *testing.T) {
	v := newToolVerifier(
		[]string{"/skills/alpha"},
		map[string]bool{"mcp1": true},
	)

	// Skills event arrives first
	v.onSkillsLoaded([]string{"alpha"})

	// Should NOT emit yet (still waiting for MCP event)
	if tools := v.emitIfReady(); tools != nil {
		t.Errorf("verifier emitted after only skills loaded: %+v", tools)
	}

	// MCP event arrives
	v.onMCPLoaded([]string{"mcp1"})

	// NOW it should emit
	tools := v.emitIfReady()
	if tools == nil {
		t.Fatal("verifier should emit after both kinds report")
	}
	if len(tools) != 2 {
		t.Errorf("expected 2 tools (1 skill + 1 mcp), got %d", len(tools))
	}
}

// TestToolValidationGate_AllFailures validates the case where ALL expected
// tools fail to load. The validation gate should detect and report all failures.
func TestToolValidationGate_AllFailures(t *testing.T) {
	v := newToolVerifier(
		[]string{"/skills/alpha", "/skills/beta"},
		map[string]bool{"mcp1": true, "mcp2": true},
	)
	// SDK reports no tools loaded (empty arrays)
	v.onSkillsLoaded([]string{})
	v.onMCPLoaded([]string{})

	tools := v.emitIfReady()
	if tools == nil {
		t.Fatal("expected tools to be emitted")
	}

	// All 4 tools should be marked as failed
	if len(tools) != 4 {
		t.Fatalf("expected 4 tools, got %d", len(tools))
	}

	for _, tool := range tools {
		if tool.Status != progress.ToolStatusFailed {
			t.Errorf("tool %s should be failed, got status=%s", tool.ToolName, tool.Status)
		}
		if tool.Reason == "" {
			t.Errorf("failed tool %s should have a reason", tool.ToolName)
		}
	}
}


// --- postSessionToolVerification (Item E) ----------------------------------

// TestPostSessionVerification_NothingConfigured asserts the gate is a no-op
// when no remote tools are declared — evals without skills/MCP must not
// stall for the timeout.
func TestPostSessionVerification_NothingConfigured(t *testing.T) {
v := newToolVerifier(nil, nil)
start := time.Now()
got := postSessionToolVerification(context.Background(), v, 5*time.Second)
elapsed := time.Since(start)
if got != "" {
t.Errorf("expected empty summary when nothing configured, got %q", got)
}
if elapsed > 100*time.Millisecond {
t.Errorf("gate should resolve immediately when nothing to verify, took %v", elapsed)
}
}

// TestPostSessionVerification_AllLoaded asserts a clean evaluation returns
// no failure summary — eval should proceed to grading.
func TestPostSessionVerification_AllLoaded(t *testing.T) {
v := newToolVerifier(
[]string{"/skills/alpha"},
map[string]bool{"azure-mcp": true},
)
v.onSkillsLoaded([]string{"alpha"})
v.onMCPLoaded([]string{"azure-mcp"})

got := postSessionToolVerification(context.Background(), v, 5*time.Second)
if got != "" {
t.Errorf("expected empty summary when all tools loaded, got %q", got)
}
}

// TestPostSessionVerification_FailedSkill asserts a missing skill produces
// a tool.SummarizeToolLoadErrors-formatted summary so the eval engine can
// short-circuit to tool_load_failure with the same wording as pre-session.
func TestPostSessionVerification_FailedSkill(t *testing.T) {
v := newToolVerifier(
[]string{"/skills/alpha", "/skills/beta"},
nil,
)
v.onSkillsLoaded([]string{"alpha"}) // beta missing

got := postSessionToolVerification(context.Background(), v, 5*time.Second)
if got == "" {
t.Fatal("expected non-empty summary for failed skill")
}
wantHeader := "1 tool(s) failed to load:"
if !strings.HasPrefix(got, wantHeader) {
t.Errorf("summary should start with %q, got: %s", wantHeader, got)
}
if !strings.Contains(got, `skill "beta"`) {
t.Errorf("summary should name the failed skill, got: %s", got)
}
}

// TestPostSessionVerification_FailedMCP asserts a missing MCP server produces
// the same shape of summary, with kind="mcp".
func TestPostSessionVerification_FailedMCP(t *testing.T) {
v := newToolVerifier(
nil,
map[string]bool{"azure-mcp": true, "playwright-mcp": true},
)
v.onMCPLoaded([]string{"azure-mcp"}) // playwright missing

got := postSessionToolVerification(context.Background(), v, 5*time.Second)
if !strings.Contains(got, `mcp "playwright-mcp"`) {
t.Errorf("summary should name failed mcp, got: %s", got)
}
if !strings.HasPrefix(got, "1 tool(s) failed to load:") {
t.Errorf("unexpected header in summary: %s", got)
}
}

// TestPostSessionVerification_MixedFailures asserts every failed tool —
// both skills and MCPs — is listed in the aggregated summary, in
// (kind, name) sort order matching emitIfReady.
func TestPostSessionVerification_MixedFailures(t *testing.T) {
v := newToolVerifier(
[]string{"/skills/alpha", "/skills/beta"},
map[string]bool{"mcp1": true, "mcp2": true},
)
v.onSkillsLoaded([]string{"alpha"}) // beta failed
v.onMCPLoaded([]string{"mcp2"})     // mcp1 failed

got := postSessionToolVerification(context.Background(), v, 5*time.Second)
if !strings.HasPrefix(got, "2 tool(s) failed to load:") {
t.Errorf("expected 2-failure header, got: %s", got)
}
for _, want := range []string{`skill "beta"`, `mcp "mcp1"`} {
if !strings.Contains(got, want) {
t.Errorf("summary missing %q\nfull summary:\n%s", want, got)
}
}
}

// TestPostSessionVerification_TimeoutMarksAllFailed asserts the timeout
// path treats every configured tool as Failed (no false positives) and
// the reason is the timeout string. Uses a short timeout to keep the test
// fast.
func TestPostSessionVerification_TimeoutMarksAllFailed(t *testing.T) {
v := newToolVerifier(
[]string{"/skills/alpha"},
map[string]bool{"mcp1": true},
)
// No SDK events ever fire.

timeout := 50 * time.Millisecond
start := time.Now()
got := postSessionToolVerification(context.Background(), v, timeout)
elapsed := time.Since(start)

if got == "" {
t.Fatal("expected timeout to produce a failure summary")
}
if !strings.HasPrefix(got, "2 tool(s) failed to load:") {
t.Errorf("timeout should mark every configured tool as failed, got: %s", got)
}
for _, want := range []string{`skill "alpha"`, `mcp "mcp1"`, "Session did not reach first turn within"} {
if !strings.Contains(got, want) {
t.Errorf("timeout summary missing %q\nfull summary:\n%s", want, got)
}
}
if elapsed < timeout {
t.Errorf("gate returned before timeout elapsed: %v < %v", elapsed, timeout)
}
if elapsed > timeout+500*time.Millisecond {
t.Errorf("gate returned far after timeout: %v >> %v", elapsed, timeout)
}
}

// TestPostSessionVerification_FormatMatchesPreSession asserts byte-for-byte
// equivalence between the post-session summary (Item E) and the pre-session
// SummarizeToolLoadErrors output (Item D). Operators must see the same
// "N tool(s) failed to load:" header, same bullet glyph, same quoting, and
// same kind/name ordering regardless of which gate produced the failure.
// Drift between the two paths would defeat the whole point of routing both
// through tool.SummarizeToolLoadErrors.
func TestPostSessionVerification_FormatMatchesPreSession(t *testing.T) {
	v := newToolVerifier(
		[]string{"/skills/alpha", "/skills/beta"},
		map[string]bool{"mcp1": true, "mcp2": true},
	)
	v.onSkillsLoaded([]string{"alpha"}) // beta failed
	v.onMCPLoaded([]string{"mcp2"})     // mcp1 failed

	postSummary := postSessionToolVerification(context.Background(), v, 5*time.Second)
	if postSummary == "" {
		t.Fatal("expected non-empty post-session summary")
	}

	// Build the equivalent pre-session error list using the same kind names
	// and the empty Reason that the post-session path emits when the SDK
	// did not provide a reason. emitIfReady returns entries sorted by
	// (kind, name): mcp before skill, alpha before beta.
	preErrs := []*tool.ToolLoadError{
		{Kind: "mcp", Name: "mcp1", Reason: "SDK did not report MCP server as loaded"},
		{Kind: "skill", Name: "beta", Reason: "SDK did not report skill as loaded"},
	}
	preSummary := tool.SummarizeToolLoadErrors(preErrs)

	if postSummary != preSummary {
		t.Errorf("post-session summary diverges from pre-session format\nPOST:\n%s\nPRE:\n%s",
			postSummary, preSummary)
	}
}

// TestPostSessionVerification_NilVerifier asserts the helper is safe to call
// with a nil verifier — defensive coverage for future call sites.
func TestPostSessionVerification_NilVerifier(t *testing.T) {
if got := postSessionToolVerification(context.Background(), nil, time.Second); got != "" {
t.Errorf("nil verifier should return empty summary, got %q", got)
}
}
