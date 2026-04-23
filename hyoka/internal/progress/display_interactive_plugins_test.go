package progress

import (
	"bytes"
	"strings"
	"testing"
)

// TestInteractive_TwoPluginsDistinctHeaders verifies that two plugins
// each render their own grouped header and their children stay
// attributed to the correct parent. Regression guard against the old
// bug where plugin fan-out shared a single flat list.
func TestInteractive_TwoPluginsDistinctHeaders(t *testing.T) {
	var buf bytes.Buffer
	d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, Mode: ModeInteractive})
	d.HandleEvent(ProgressEvent{EvalID: "e1", PromptID: "p", ConfigName: "c", Type: EventStarting})

	// Plugin A parent + two children.
	d.HandleEvent(ProgressEvent{EvalID: "e1", Type: EventToolResolutionStart, ToolName: "plugin-A", ToolKind: ToolKindPlugin})
	d.HandleEvent(ProgressEvent{EvalID: "e1", Type: EventToolResolutionResult, ToolName: "plugin-A", ToolKind: ToolKindPlugin, Status: ToolStatusLoaded})
	d.HandleEvent(ProgressEvent{EvalID: "e1", Type: EventToolResolutionResult, ToolName: "a-child-1", ToolKind: ToolKindSkill, Status: ToolStatusLoaded, ParentName: "plugin-A", ParentKind: ToolParentKindPlugin})
	d.HandleEvent(ProgressEvent{EvalID: "e1", Type: EventToolResolutionResult, ToolName: "a-child-2", ToolKind: ToolKindSkill, Status: ToolStatusLoaded, ParentName: "plugin-A", ParentKind: ToolParentKindPlugin})

	// Plugin B parent + one child.
	d.HandleEvent(ProgressEvent{EvalID: "e1", Type: EventToolResolutionStart, ToolName: "plugin-B", ToolKind: ToolKindPlugin})
	d.HandleEvent(ProgressEvent{EvalID: "e1", Type: EventToolResolutionResult, ToolName: "plugin-B", ToolKind: ToolKindPlugin, Status: ToolStatusLoaded})
	d.HandleEvent(ProgressEvent{EvalID: "e1", Type: EventToolResolutionResult, ToolName: "b-child-1", ToolKind: ToolKindMCP, Status: ToolStatusLoaded, ParentName: "plugin-B", ParentKind: ToolParentKindPlugin})

	d.HandleEvent(ProgressEvent{EvalID: "e1", Type: EventPassed, FileCount: 0})
	d.Finish()

	out := buf.String()
	for _, want := range []string{
		"plugin-A (plugin):",
		"plugin-B (plugin):",
		"a-child-1",
		"a-child-2",
		"b-child-1",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("missing %q in output:\n%s", want, out)
		}
	}

	// Children must appear after their own plugin header.
	aHdr := strings.Index(out, "plugin-A (plugin):")
	aC1 := strings.Index(out, "a-child-1")
	aC2 := strings.Index(out, "a-child-2")
	if aHdr < 0 || aC1 < 0 || aC2 < 0 || !(aHdr < aC1 && aHdr < aC2) {
		t.Errorf("plugin-A header must precede its children; aHdr=%d c1=%d c2=%d\n%s", aHdr, aC1, aC2, out)
	}
	bHdr := strings.Index(out, "plugin-B (plugin):")
	bC1 := strings.Index(out, "b-child-1")
	if bHdr < 0 || bC1 < 0 || bHdr >= bC1 {
		t.Errorf("plugin-B header must precede its child; bHdr=%d c1=%d\n%s", bHdr, bC1, out)
	}

	// Children of A must NOT appear after B's header-only (i.e. the
	// two groups should be contiguous, not interleaved).
	if aC1 > bHdr {
		t.Errorf("plugin-A children must not appear after plugin-B header:\n%s", out)
	}

	// Neither plugin should also appear as a flat leaf status row.
	for _, bad := range []string{
		"plugin-A (plugin): ✅",
		"plugin-B (plugin): ✅",
	} {
		if strings.Contains(out, bad) {
			t.Errorf("plugin container must not render a flat leaf status; saw %q in:\n%s", bad, out)
		}
	}
}

// TestInteractive_WaitTillKnown_FailedEmitsReason complements Tank's
// WaitTillKnown test by verifying the failed-transition path emits
// both the failed marker AND the reason text on the committed line.
func TestInteractive_WaitTillKnown_FailedEmitsReason(t *testing.T) {
	var buf bytes.Buffer
	d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, Mode: ModeInteractive})
	d.HandleEvent(ProgressEvent{EvalID: "wk", PromptID: "p", ConfigName: "c", Type: EventStarting})

	// Start only — must not emit anything tool-related.
	d.HandleEvent(ProgressEvent{EvalID: "wk", Type: EventToolResolutionStart, ToolName: "doomed", ToolKind: ToolKindMCP})
	mid := buf.String()
	if strings.Contains(mid, "doomed") {
		t.Fatalf("tool name leaked before Result:\n%s", mid)
	}

	// Now a failed Result with a reason.
	d.HandleEvent(ProgressEvent{
		EvalID: "wk", Type: EventToolResolutionResult,
		ToolName: "doomed", ToolKind: ToolKindMCP,
		Status: ToolStatusFailed, Reason: "command not found",
	})
	d.HandleEvent(ProgressEvent{EvalID: "wk", Type: EventPassed, FileCount: 0})
	d.Finish()

	out := buf.String()
	for _, want := range []string{"doomed", "❌ Failed", "command not found"} {
		if !strings.Contains(out, want) {
			t.Errorf("missing %q in final output:\n%s", want, out)
		}
	}
	if strings.Contains(out, "Loading") {
		t.Errorf("transient Loading text leaked into final output:\n%s", out)
	}
}

// TestInteractive_PluginParentNoLoadedFailedBadge regresses issue (b):
// the plugin parent header must never carry a Loaded or Failed badge in
// the final output, even after EventToolsVerified flips
// loaded-but-not-SDK-reported tools to Failed. The plugin parent is a
// container — only its leaf children are SDK-reported.
func TestInteractive_PluginParentNoLoadedFailedBadge(t *testing.T) {
var buf bytes.Buffer
d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, Mode: ModeInteractive})
d.HandleEvent(ProgressEvent{EvalID: "e1", PromptID: "p", ConfigName: "c", Type: EventStarting})

// Plugin Start (mirrors validate.go emitStart for the parent), then
// only child Results arrive (the success path no longer emits a
// parent Result).
d.HandleEvent(ProgressEvent{EvalID: "e1", Type: EventToolResolutionStart, ToolName: "azure-sdk-python", ToolKind: ToolKindPlugin})
d.HandleEvent(ProgressEvent{EvalID: "e1", Type: EventToolResolutionResult, ToolName: "default-azure-credential", ToolKind: ToolKindSkill, Status: ToolStatusLoaded, ParentName: "azure-sdk-python", ParentKind: ToolParentKindPlugin})
d.HandleEvent(ProgressEvent{EvalID: "e1", Type: EventToolResolutionResult, ToolName: "azure-mcp", ToolKind: ToolKindMCP, Status: ToolStatusLoaded, ParentName: "azure-sdk-python", ParentKind: ToolParentKindPlugin})

// SDK verification: reports only the leaf children. The plugin parent
// must NOT be flipped to Failed by the not-reported-by-SDK rule.
d.HandleEvent(ProgressEvent{EvalID: "e1", Type: EventToolsVerified, Tools: []ToolStatus{
{ToolName: "default-azure-credential", ToolKind: ToolKindSkill, Status: ToolStatusLoaded, ParentName: "azure-sdk-python", ParentKind: ToolParentKindPlugin},
{ToolName: "azure-mcp", ToolKind: ToolKindMCP, Status: ToolStatusLoaded, ParentName: "azure-sdk-python", ParentKind: ToolParentKindPlugin},
}})

d.HandleEvent(ProgressEvent{EvalID: "e1", Type: EventPassed, FileCount: 0})
d.Finish()

out := buf.String()

if !strings.Contains(out, "azure-sdk-python (plugin):") {
t.Fatalf("expected plugin header in output:\n%s", out)
}
// Plugin parent row must NOT carry any status badge.
for _, bad := range []string{
"azure-sdk-python (plugin): ✅",
"azure-sdk-python (plugin): ❌",
"azure-sdk-python (plugin): Loaded",
"azure-sdk-python (plugin): Failed",
"azure-sdk-python (plugin): not reported by SDK",
} {
if strings.Contains(out, bad) {
t.Errorf("plugin parent must have no status badge; saw %q in:\n%s", bad, out)
}
}
// Children should still report Loaded.
if !strings.Contains(out, "default-azure-credential") || !strings.Contains(out, "Loaded") {
t.Errorf("expected children to report Loaded:\n%s", out)
}
}

// TestInteractive_ChildKindLabelShown regresses issue (c): children
// rendered under a plugin/skill_dir header must still display their
// kind label so callers can distinguish skill children from MCP
// children at a glance.
func TestInteractive_ChildKindLabelShown(t *testing.T) {
var buf bytes.Buffer
d := NewDisplay(DisplayConfig{Total: 1, Workers: 1, Writer: &buf, Mode: ModeInteractive})
d.HandleEvent(ProgressEvent{EvalID: "kc", PromptID: "p", ConfigName: "c", Type: EventStarting})

// One plugin with mixed children: a skill and an MCP.
d.HandleEvent(ProgressEvent{EvalID: "kc", Type: EventToolResolutionStart, ToolName: "azure-sdk-python", ToolKind: ToolKindPlugin})
d.HandleEvent(ProgressEvent{EvalID: "kc", Type: EventToolResolutionResult, ToolName: "default-azure-credential", ToolKind: ToolKindSkill, Status: ToolStatusLoaded, ParentName: "azure-sdk-python", ParentKind: ToolParentKindPlugin})
d.HandleEvent(ProgressEvent{EvalID: "kc", Type: EventToolResolutionResult, ToolName: "azure-mcp", ToolKind: ToolKindMCP, Status: ToolStatusLoaded, ParentName: "azure-sdk-python", ParentKind: ToolParentKindPlugin})

d.HandleEvent(ProgressEvent{EvalID: "kc", Type: EventPassed, FileCount: 0})
d.Finish()

out := buf.String()
for _, want := range []string{
"default-azure-credential (skill)",
"azure-mcp (mcp)",
} {
if !strings.Contains(out, want) {
t.Errorf("expected child to include kind label %q in:\n%s", want, out)
}
}
}
