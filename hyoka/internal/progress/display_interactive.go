package progress

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"sync"
	"time"

	"github.com/mattn/go-runewidth"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/progress/style"
)

// ansiSeqRE matches ANSI CSI escape sequences (color/style/cursor codes) so we
// can compute the visible width of a styled string and truncate it without
// chopping in the middle of an escape sequence.
var ansiSeqRE = regexp.MustCompile(`\x1b\[[0-9;]*[A-Za-z]`)

// visibleWidth returns the visible width of s in terminal columns, stripping
// ANSI escape sequences and using proper cell width calculation (accounts for
// wide characters like emoji and CJK that occupy 2 cells).
func visibleWidth(s string) int {
	// Strip ANSI sequences first.
	stripped := ansiSeqRE.ReplaceAllString(s, "")
	return runewidth.StringWidth(stripped)
}

// truncateToWidth returns s truncated so its visible width is at most max
// columns. ANSI escape sequences are preserved (zero width). Uses proper cell
// width calculation (wide chars like emoji count as 2). If s already fits, it's
// returned unchanged. The truncation appends an ellipsis "…" when content was
// dropped.
func truncateToWidth(s string, max int) string {
	if max <= 0 {
		return s
	}
	// Fast path: if there are no ANSI codes and length already fits.
	if len(s) <= max && !ansiSeqRE.MatchString(s) {
		return s
	}
	visible := 0
	out := make([]byte, 0, len(s))
	i := 0
	for i < len(s) {
		// Match ANSI sequence at position i.
		if loc := ansiSeqRE.FindStringIndex(s[i:]); loc != nil && loc[0] == 0 {
			out = append(out, s[i:i+loc[1]]...)
			i += loc[1]
			continue
		}
		// Decode the next rune and get its display width.
		r, size := decodeRune(s[i:])
		w := runewidth.RuneWidth(r)
		// Check if adding this rune would exceed the limit (leaving 1 col for ellipsis).
		if visible+w > max-1 {
			out = append(out, []byte("…")...)
			// Append ANSI reset to avoid bleeding styles past the truncation.
			out = append(out, []byte("\x1b[0m")...)
			return string(out)
		}
		out = append(out, []byte(string(r))...)
		visible += w
		i += size
	}
	return string(out)
}

func decodeRune(s string) (rune, int) {
	for i, r := range s {
		_ = i
		return r, len(string(r))
	}
	return 0, 0
}

// interactiveRenderer implements the "interactive" progress display described
// in the sprint plan: a per-eval transcript where only the TAIL line is ever
// updated in place. Older lines are immutable, with ONE documented exception:
// on EventToolsVerified, if any tool's previously-reported status flips
// (Loaded → Failed), the Tools block is redrawn via DECSC/DECRC save/restore.
//
// Layout per eval (matches plan.md):
//
//	Prompt: <prompt-id>
//	Config: <config-name>
//	Tools:
//	  - <name> (<kind>): 🔄 Loading…              (tail; flips to ✅/❌ in place)
//	  - <name> (<kind>): ✅ Loaded
//	Agent Attempt: 🔄 Running                       (single-line tail; flips in place)
//	Agent Attempt: ✅ Completed                      (final state — same row, rewritten)
//	Session Details:
//	  Files: …
//	  Turns: …   Tool calls: …   Cost: …
//	Graders:
//	  - <id> (<kind>): ✅ Pass (X/10)
//	  - <id> (<kind>): 🔄 Running…                 (tail; next grader appends
//	                                                a fresh line)
//
// The "Tools:" section header is omitted entirely when no tools are configured
// for a run; same for "Graders:" if no graders fire.
//
// Multi-eval handling: interactive mode is only selected when workers == 1,
// so there is exactly one eval in flight at a time. Queued evals render one
// after another — eval 1 finalizes, then eval 2 begins in a new block.
type interactiveRenderer struct {
	w         io.Writer
	sty       *style.Styler
	mu        sync.Mutex
	startTime time.Time
	total     int
	reportDir string

	// Outer counters (mirror Display so CompletedEvalCount and final summary
	// work).
	completed int
	passed    int
	failed    int
	errors    int

	// Current eval state — reset between evals.
	cur *interactiveEval

	// Per-second ticker used only to refresh the Agent Attempt tail line's
	// duration counter. Other sections are event-driven.
	ticker *time.Ticker
	stopCh chan struct{}
	wg     sync.WaitGroup
}

// interactiveEval captures state for the single in-flight eval.
type interactiveEval struct {
	evalID     string
	promptID   string
	configName string
	startTime  time.Time

	// Line bookkeeping for the ONE non-tail redraw exception (Tools block).
	// linesWritten counts newline-terminated lines printed so far; the tail
	// (if active) sits on line `linesWritten` with no trailing newline yet.
	linesWritten int

	// Tail state. Only one line at a time can be the tail. Freezing the tail
	// means writing "\n" and setting tailKind = tailNone.
	tailKind     tailKind
	tailText     string // raw text of tail (for bookkeeping; re-emit on update)
	tailRowCount int    // how many physical terminal rows the current tail occupies

	// Tools section.
	toolsHeaderPrinted bool
	toolsFirstLine     int                // linesWritten index of first tool line ("  - …")
	toolLines          []toolLine         // bookkeeping for every Start-seen tool, in order
	toolsVerified      bool               // guard so we only redraw once
	toolIndexByName    map[string]int     // name -> index in toolLines
	emittedParents     map[parentKey]bool // parent headers already written to output

	// Agent Attempt section — simplified to a three-state machine (Running, Completed, Guardrail).
	agentHeaderPrinted bool
	agentState         agentAttemptState
	agentStartTime     time.Time
	agentGuardrailMsg  string // populated when state = agentStateGuardrail

	// agentLineFrozen / agentLineRow track the row of a previously-active
	// agent tail that got committed by some other section (typically a
	// grader Start) taking over the tail before agentComplete fires. When
	// agentComplete arrives later, it can rewrite that exact row in place
	// rather than appending a stale "Completed" line at the bottom of the
	// transcript.
	agentLineFrozen bool
	agentLineRow    int

	// graderRowByID records the row index where each grader's "Running"
	// tail was committed when something else took the tail before its own
	// Complete event fired. onGraderComplete uses this to rewrite the
	// original row in place, keeping each grader as exactly one line in
	// stable order rather than producing a second stale entry at the
	// bottom of the Graders section.
	graderRowByID   map[string]int
	pendingGraderID string

	// Agent attempt buffering — hold back rendering until tools are verified.
	// If no tool events arrive before the first agent-attempt event, we treat
	// this as a no-tools config and flush immediately.
	agentEventsBuffered []ProgressEvent
	agentGateOpen       bool // true once tools verification completes (or no tools detected)

	// Session Details section.
	sessionPrinted bool

	// Graders section.
	gradersHeaderPrinted bool
	lastGraderSourceFile string // source file of the last grader file-header printed

	// Terminal state for this eval — set when EventPassed/Failed/Error fires.
	terminalStatus evalStatus
	terminalMsg    string
}

// toolLine captures what we printed for a tool line, so we can compare
// against the bulk verification event and redraw on flips.
type toolLine struct {
	name       string
	kind       string
	status     string // "loading" | ToolStatusLoaded | ToolStatusFailed
	reason     string
	parentName string // Parent container (plugin name or skills-dir path); empty = no parent
	parentKind string // One of ToolParentKindPlugin, ToolParentKindSkillDir, or empty
}

// parentKey identifies a grouped parent (plugin or skills-dir) by (name, kind).
// Stored in emittedParents so we only print each parent header once.
type parentKey struct {
	name string
	kind string
}

// tailKind enumerates which kind of line currently owns the tail, so ticker
// redraws only touch the Agent Attempt line.
type tailKind int

const (
	tailNone tailKind = iota
	tailTool
	tailAgent
	tailGrader
)

// agentAttemptState represents the three possible states for the Agent Attempt section.
type agentAttemptState int

const (
	agentStateRunning agentAttemptState = iota
	agentStateCompleted
	agentStateGuardrail
)

// ANSI cursor helpers. Kept private to this file so the rest of the package
// doesn't rely on raw escapes.
const (
	ansiClearLine   = "\x1b[2K"
	ansiCR          = "\r"
	ansiCursorUp1   = "\x1b[1A"
	ansiSaveCursor  = "\x1b7" // DECSC — same convention display.go uses
	ansiRestoreCurs = "\x1b8" // DECRC
	ansiCursorUpFmt = "\x1b[%dA"
)

// newInteractiveRenderer constructs the renderer. When total==0 we still
// render — interactive mode is used for ad-hoc single-eval runs too.
func newInteractiveRenderer(w io.Writer, total, workers int, reportDir string) *interactiveRenderer {
	r := &interactiveRenderer{
		w:         w,
		sty:       style.New(w),
		startTime: time.Now(),
		total:     total,
		reportDir: reportDir,
	}
	if total > 0 {
		fmt.Fprintf(r.w, "\nRunning %d evaluation(s) (%d worker)\n\n", total, workers)
	}
	r.stopCh = make(chan struct{})
	r.ticker = time.NewTicker(time.Second)
	r.wg.Add(1)
	go r.tickLoop()
	return r
}

// tickLoop refreshes tail lines once per second. With the three-state Agent
// Attempt model, only the Grader tail needs ticker updates (if we add progress
// indicators there later). For now, this is a no-op but kept for symmetry.
func (r *interactiveRenderer) tickLoop() {
	defer r.wg.Done()
	for {
		select {
		case <-r.stopCh:
			return
		case <-r.ticker.C:
			// Agent Attempt no longer needs ticker updates (three-state model)
			// If we add grader progress indicators, update them here
		}
	}
}

func (r *interactiveRenderer) tailIsAgent() bool {
	return r.cur != nil && r.cur.tailKind == tailAgent
}

// --- Event dispatch ---

func (r *interactiveRenderer) handleEvent(evt ProgressEvent) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// A new EventStarting for a different eval finalizes the previous one's
	// block and opens a fresh block.
	if evt.Type == EventStarting {
		if r.cur != nil && r.cur.evalID != evt.EvalID {
			r.finalizeCurrentEval()
		}
		r.startEval(evt)
		return
	}

	if r.cur == nil || r.cur.evalID != evt.EvalID {
		// Event for an unknown / stale eval — ignore. Interactive mode is
		// single-in-flight, so this is defensive.
		return
	}

	switch evt.Type {
	case EventToolResolutionStart:
		r.onToolResolutionStart(evt)
	case EventToolResolutionResult:
		r.onToolResolutionResult(evt)
	case EventToolsVerified:
		r.onToolsVerified(evt)
	case EventPhaseChange:
		r.onPhaseChange(evt)
	case EventSendingPrompt, EventReasoning, EventToolStart, EventToolComplete,
		EventWritingFile, EventWaiting:
		r.onAgentActivity(evt)
	case EventSessionDetails:
		r.onSessionDetails(evt)
	case EventGraderStart:
		r.onGraderStart(evt)
	case EventGraderComplete:
		r.onGraderComplete(evt)
	case EventPassed:
		r.onPassed(evt)
	case EventFailed:
		r.onFailed(evt)
	case EventError:
		r.onError(evt)
	}
}

// --- Eval lifecycle ---

func (r *interactiveRenderer) startEval(evt ProgressEvent) {
	r.cur = &interactiveEval{
		evalID:          evt.EvalID,
		promptID:        evt.PromptID,
		configName:      evt.ConfigName,
		startTime:       time.Now(),
		toolIndexByName: make(map[string]int),
	}
	// Header lines: Prompt + Config. Both committed (newline-terminated)
	// immediately — they're never the tail.
	r.writeLine(fmt.Sprintf("Prompt: %s", evt.PromptID))
	r.writeLine(fmt.Sprintf("Config: %s", evt.ConfigName))
}

// finalizeCurrentEval freezes whatever's on the tail, prints a blank
// separator line, and clears r.cur. Safe to call multiple times.
func (r *interactiveRenderer) finalizeCurrentEval() {
	if r.cur == nil {
		return
	}
	r.freezeTail()
	// Blank separator between evals.
	fmt.Fprintln(r.w)
	r.cur = nil
}

// --- Tools section ---

func (r *interactiveRenderer) ensureToolsHeader() {
	if r.cur.toolsHeaderPrinted {
		return
	}
	r.freezeTail()
	r.writeLine("Tools:")
	r.cur.toolsHeaderPrinted = true
	r.cur.toolsFirstLine = r.cur.linesWritten
}

func (r *interactiveRenderer) onToolResolutionStart(evt ProgressEvent) {
	r.ensureToolsHeader()
	// Wait-till-known: a Start event carries NO terminal state, so we DO NOT
	// render anything for it. We only record the tool in internal bookkeeping
	// so the matching Result can find it. If a Start never gets a Result
	// (e.g. the skill_dir parent Start in validate.go, which emits children
	// results but no parent result), the entry stays "loading" and is
	// filtered out of rendering.
	if _, exists := r.cur.toolIndexByName[evt.ToolName]; exists {
		// Re-start (rare): leave the existing entry alone.
		return
	}
	tl := toolLine{
		name:       evt.ToolName,
		kind:       evt.ToolKind,
		status:     "loading",
		parentName: evt.ParentName,
		parentKind: evt.ParentKind,
	}
	r.cur.toolLines = append(r.cur.toolLines, tl)
	r.cur.toolIndexByName[evt.ToolName] = len(r.cur.toolLines) - 1
}

func (r *interactiveRenderer) onToolResolutionResult(evt ProgressEvent) {
	idx, ok := r.cur.toolIndexByName[evt.ToolName]
	if !ok {
		// Result without a matching Start — defensive: create the entry now.
		r.ensureToolsHeader()
		r.cur.toolLines = append(r.cur.toolLines, toolLine{})
		idx = len(r.cur.toolLines) - 1
		r.cur.toolIndexByName[evt.ToolName] = idx
	}
	tl := &r.cur.toolLines[idx]
	tl.name = evt.ToolName
	tl.kind = evt.ToolKind
	tl.status = evt.Status
	tl.reason = evt.Reason
	tl.parentName = evt.ParentName
	tl.parentKind = evt.ParentKind

	// Any active tail (e.g. a leftover agent line) must be frozen before we
	// commit a new Tools-section line.
	r.freezeTail()

	// Leaf child: ensure its parent header is emitted, then write the
	// indented child row.
	if tl.parentKind != "" {
		r.emitParentHeaderOnce(tl.parentName, tl.parentKind)
		r.writeLine(r.renderToolLine(*tl, true))
		return
	}

	// Top-level plugin: it is a container. When it loads successfully we
	// emit a parent header (children will follow as they resolve). When it
	// fails, we emit a single flat failed row — no children are expected.
	if tl.kind == ToolKindPlugin {
		if tl.status == ToolStatusLoaded {
			r.emitParentHeaderOnce(tl.name, ToolParentKindPlugin)
			return
		}
		r.writeLine(r.renderToolLine(*tl, false))
		return
	}

	// Top-level leaf (skill or MCP).
	r.writeLine(r.renderToolLine(*tl, false))
}

// emitParentHeaderOnce writes the "  - <name> (<kind>):" header line exactly
// once per (name, kind) combination for the current eval.
func (r *interactiveRenderer) emitParentHeaderOnce(name, kind string) {
	if r.cur.emittedParents == nil {
		r.cur.emittedParents = make(map[parentKey]bool)
	}
	pk := parentKey{name: name, kind: kind}
	if r.cur.emittedParents[pk] {
		return
	}
	r.cur.emittedParents[pk] = true
	label := "plugin"
	if kind == ToolParentKindSkillDir {
		label = "skills dir"
	}
	r.freezeTail()
	r.writeLine(fmt.Sprintf("  - %s (%s):", name, label))
}

// onToolsVerified handles the bulk post-session-start verification. This is
// the ONE sanctioned exception to the tail-only update rule: any tool whose
// previously-reported status flips (Loaded→Failed or Failed→Loaded or new
// tool appears / disappears) triggers a single block redraw of the Tools
// section using DECSC/DECRC save/restore. See plan.md "Tools re-verification".
func (r *interactiveRenderer) onToolsVerified(evt ProgressEvent) {
	if r.cur.toolsVerified {
		return
	}
	r.cur.toolsVerified = true
	if !r.cur.toolsHeaderPrinted || len(r.cur.toolLines) == 0 {
		// No Tools section was ever rendered — nothing to redraw. We still
		// emit a header + lines for any reported tools if there are any.
		if len(evt.Tools) == 0 {
			// Open the agent gate now that tools verification is complete.
			r.openAgentGate()
			return
		}
		r.ensureToolsHeader()
		r.freezeTail()
		// When tools are first reported in bulk, render them grouped
		grouped := r.groupToolLinesFromStatus(evt.Tools)
		for _, line := range grouped {
			r.writeLine(line)
		}
		// Update internal state
		for _, t := range evt.Tools {
			tl := toolLine{
				name:       t.ToolName,
				kind:       t.ToolKind,
				status:     t.Status,
				reason:     t.Reason,
				parentName: t.ParentName,
				parentKind: t.ParentKind,
			}
			r.cur.toolLines = append(r.cur.toolLines, tl)
			r.cur.toolIndexByName[t.ToolName] = len(r.cur.toolLines) - 1
		}
		// Open the agent gate now that tools verification is complete.
		r.openAgentGate()
		return
	}
	// Merge verification results into existing toolLines. Any flips mark the
	// block as dirty. Tools reported by the SDK that we never saw at
	// resolution time are appended (rare edge case).
	dirty := false
	seen := make(map[string]bool, len(evt.Tools))
	for _, t := range evt.Tools {
		seen[t.ToolName] = true
		if idx, ok := r.cur.toolIndexByName[t.ToolName]; ok {
			prev := r.cur.toolLines[idx]
			if prev.status != t.Status || prev.reason != t.Reason {
				r.cur.toolLines[idx].status = t.Status
				r.cur.toolLines[idx].reason = t.Reason
				dirty = true
			}
		}
	}
	// Tools we printed as "loaded" but the SDK didn't report must flip to
	// failed. (See plan.md: "Loaded but not reported by SDK → Failed".)
	// EXCEPT for plugin parents and any container entry with at least one
	// child reference — those are organizational headers, not SDK-loadable
	// tools. The SDK only ever reports leaf skills/MCPs, so flipping a
	// container would always mark it Failed even when every child loaded.
	containerName := make(map[string]bool)
	for _, tl := range r.cur.toolLines {
		if tl.parentName != "" {
			containerName[tl.parentName] = true
		}
	}
	for i, tl := range r.cur.toolLines {
		if seen[tl.name] || tl.status != ToolStatusLoaded {
			continue
		}
		if tl.kind == ToolKindPlugin || containerName[tl.name] {
			// Container parent — never SDK-reported. Skip the flip.
			continue
		}
		r.cur.toolLines[i].status = ToolStatusFailed
		if r.cur.toolLines[i].reason == "" {
			r.cur.toolLines[i].reason = "not reported by SDK"
		}
		dirty = true
	}
	if dirty {
		// Freeze any active tail before redrawing the tools block. The block
		// redraw logic assumes the cursor is at a known position (column 0 of
		// the line after the last committed line), and an active tail would
		// break that calculation.
		r.freezeTail()
		r.redrawToolsBlock()
	}
	// Open the agent gate now that tools verification is complete.
	r.openAgentGate()
}

// redrawToolsBlock performs a DECSC/DECRC-bracketed rewrite of every tool
// line. This is the single documented exception to the tail-only rule.
// Precondition: the Tools section has been rendered (toolsHeaderPrinted).
func (r *interactiveRenderer) redrawToolsBlock() {
	e := r.cur
	if !e.toolsHeaderPrinted || len(e.toolLines) == 0 {
		return
	}
	// Cursor is either sitting at column 0 of the tail line (no trailing
	// newline yet) OR — if no tail is active — at column 0 of the line AFTER
	// the last committed line. In both cases, the distance from the cursor's
	// row to the first tool line's row is:
	//   tailOffset + (linesWritten - toolsFirstLine)
	// where tailOffset = 0 when no tail, 0 when tail (tail line = linesWritten).
	// Both reduce to the same expression.
	up := e.linesWritten - e.toolsFirstLine
	if up <= 0 {
		return
	}
	var buf bytes.Buffer
	buf.WriteString(ansiSaveCursor)
	fmt.Fprintf(&buf, ansiCursorUpFmt, up)
	buf.WriteString(ansiCR)

	// Group tools by (ParentKind, ParentName). Preserve insertion order.
	grouped := r.groupToolLines(e.toolLines)
	for _, line := range grouped {
		buf.WriteString(ansiClearLine)
		buf.WriteString(line)
		buf.WriteByte('\n')
	}

	buf.WriteString(ansiRestoreCurs)
	r.w.Write(buf.Bytes())
}

// groupToolLines groups tool lines by parent and returns a flat list of
// formatted lines (parent headers + indented children). Preserves insertion
// order. A top-level entry is treated as a *container* if its kind is
// ToolKindPlugin OR if some child entry references it by ParentName — in
// that case its own status is subsumed into the parent header (loaded →
// just the header; failed → a single flat failed row, no children group).
// Orphan "loading" entries (Start with no Result) are filtered out.
func (r *interactiveRenderer) groupToolLines(toolLines []toolLine) []string {
	// First pass: discover which names act as containers (referenced by at
	// least one child's ParentName).
	isContainer := make(map[string]bool)
	for _, tl := range toolLines {
		if tl.parentKind != "" && tl.parentName != "" {
			isContainer[tl.parentName] = true
		}
	}

	var parentOrder []parentKey
	parentSeen := make(map[parentKey]bool)
	parentChildren := make(map[parentKey][]toolLine)
	parentOwnRow := make(map[parentKey]toolLine) // the container's own row (kind=plugin, etc.)
	var topLevel []toolLine

	for _, tl := range toolLines {
		if tl.parentKind != "" {
			pk := parentKey{name: tl.parentName, kind: tl.parentKind}
			if !parentSeen[pk] {
				parentOrder = append(parentOrder, pk)
				parentSeen[pk] = true
			}
			parentChildren[pk] = append(parentChildren[pk], tl)
			continue
		}
		// Top-level. A plugin-kind entry is always a container.
		// Non-plugin entries are containers only if a child references them
		// (e.g. skill_dir parent emitted with kind=skill in validate.go).
		if tl.kind == ToolKindPlugin || isContainer[tl.name] {
			kind := ToolParentKindPlugin
			if tl.kind != ToolKindPlugin {
				kind = ToolParentKindSkillDir
			}
			pk := parentKey{name: tl.name, kind: kind}
			if !parentSeen[pk] {
				parentOrder = append(parentOrder, pk)
				parentSeen[pk] = true
			}
			parentOwnRow[pk] = tl
			continue
		}
		topLevel = append(topLevel, tl)
	}

	var result []string
	for _, tl := range topLevel {
		if tl.status == "" || tl.status == "loading" {
			// Orphan: never got a Result. Skip — don't surface as output.
			continue
		}
		result = append(result, r.renderToolLine(tl, false))
	}
	for _, pk := range parentOrder {
		// If the container itself failed, render it as a single flat failed
		// row — its children (if any) are not meaningful when the container
		// didn't load.
		if own, ok := parentOwnRow[pk]; ok && own.status == ToolStatusFailed {
			result = append(result, r.renderToolLine(own, false))
			continue
		}
		// Parent header.
		kindLabel := "plugin"
		if pk.kind == ToolParentKindSkillDir {
			kindLabel = "skills dir"
		}
		result = append(result, fmt.Sprintf("  - %s (%s):", pk.name, kindLabel))
		for _, child := range parentChildren[pk] {
			if child.status == "" || child.status == "loading" {
				continue
			}
			result = append(result, r.renderToolLine(child, true))
		}
	}
	return result
}

// groupToolLines overload for []ToolStatus (from EventToolsVerified)
func (r *interactiveRenderer) groupToolLinesFromStatus(tools []ToolStatus) []string {
	converted := make([]toolLine, len(tools))
	for i, t := range tools {
		converted[i] = toolLine{
			name:       t.ToolName,
			kind:       t.ToolKind,
			status:     t.Status,
			reason:     t.Reason,
			parentName: t.ParentName,
			parentKind: t.ParentKind,
		}
	}
	return r.groupToolLines(converted)
}

func (r *interactiveRenderer) renderToolLine(tl toolLine, indented bool) string {
	name := tl.name
	// Indent prefix for children under a parent group
	indent := "  "
	if indented {
		indent = "      " // Extra 4 spaces for grouped children
	}

	// Always render the kind label, including for grouped children. Without
	// this, children under a plugin/skill_dir header read as bare names with
	// no indication of skill vs MCP — visually ambiguous when the parent
	// container mixes both.
	kindStr := " " + r.sty.Muted(fmt.Sprintf("(%s)", tl.kind))

	switch tl.status {
	case "", "loading":
		return fmt.Sprintf("%s- %s%s: 🔄 %s", indent, name, kindStr, r.sty.Muted("Loading…"))
	case ToolStatusLoaded:
		return fmt.Sprintf("%s- %s%s: %s", indent, name, kindStr, r.sty.OK("✅ Loaded"))
	case ToolStatusFailed:
		reason := tl.reason
		if reason == "" {
			reason = "failed"
		}
		return fmt.Sprintf("%s- %s%s: %s %s", indent, name, kindStr,
			r.sty.Fail("❌ Failed"), r.sty.Muted("("+reason+")"))
	default:
		return fmt.Sprintf("%s- %s%s: %s", indent, name, kindStr, tl.status)
	}
}

// --- Agent Attempt section ---

// openAgentGate unblocks the Agent Attempt section rendering and flushes any
// buffered agent-attempt events. Called when EventToolsVerified arrives OR
// when we detect a no-tools config (first agent event with no prior tool events).
func (r *interactiveRenderer) openAgentGate() {
	if r.cur.agentGateOpen {
		return
	}
	r.cur.agentGateOpen = true
	// Flush buffered events in order.
	for _, e := range r.cur.agentEventsBuffered {
		r.renderAgentEvent(e)
	}
	r.cur.agentEventsBuffered = nil
}

// detectNoTools returns true if no tool events have been seen yet for this
// eval AND no tools section has been printed. This signals a no-tools config.
func (r *interactiveRenderer) detectNoTools() bool {
	return !r.cur.toolsHeaderPrinted && len(r.cur.toolLines) == 0
}

// ensureAgentHeader marks the Agent Attempt section as opened and freezes any
// active tail so the next writeTail starts on a fresh row. It does NOT write a
// standalone "Agent Attempt:" header — the section is rendered as a single
// combined line ("Agent Attempt: 🔄 Running") by renderAgentStateLine.
// agentHeaderPrinted means "section opened; the single-line tail has been
// (or is about to be) written at least once."
func (r *interactiveRenderer) ensureAgentHeader() {
	if r.cur.agentHeaderPrinted {
		return
	}
	r.freezeTail()
	r.cur.agentHeaderPrinted = true
	r.cur.agentStartTime = time.Now()
}

func (r *interactiveRenderer) onPhaseChange(evt ProgressEvent) {
	// Phase events are pure metadata in interactive mode. They fire BEFORE
	// tool resolution begins (engine_eval.go calls sendPhase(PhaseGenerating)
	// before evaluator.Run, which is where buildSessionConfig emits tool
	// events). If we treat PhaseChange as agent activity it opens the gate
	// prematurely via detectNoTools() and causes the Agent Attempt header
	// to print before the Tools section. Activity framing is driven by
	// real session events (sending prompt, reasoning, tool calls) — not
	// phase transitions. So this handler is intentionally a no-op.
	_ = evt
}

func (r *interactiveRenderer) onAgentActivity(evt ProgressEvent) {
	// Buffer or render immediately based on whether gate is open.
	if !r.cur.agentGateOpen {
		// No tools verification yet. Check if this is a no-tools config.
		if r.detectNoTools() {
			r.openAgentGate()
		} else {
			// Tools expected — buffer this event.
			r.cur.agentEventsBuffered = append(r.cur.agentEventsBuffered, evt)
			return
		}
	}
	r.renderAgentEvent(evt)
}

// renderAgentEvent performs the actual rendering for an agent-attempt event.
// Factored out so both buffered-flush and direct rendering use the same logic.
// With the three-state model, we just transition to Running on first event.
func (r *interactiveRenderer) renderAgentEvent(evt ProgressEvent) {
	// Agent Attempt is already finalized — generation phase is over. Ignore
	// activity events from downstream sessions (reviewer Copilot sessions
	// emit the same EventReasoning/EventToolStart/etc. through the shared
	// event channel, but they belong to grader rows, not the agent tail).
	if r.cur != nil && (r.cur.agentState == agentStateCompleted || r.cur.agentState == agentStateGuardrail) {
		return
	}
	r.ensureAgentHeader()
	// First agent event transitions to Running state if not already set.
	if r.cur.agentState == 0 { // zero value means not initialized
		r.cur.agentState = agentStateRunning
		r.cur.agentStartTime = time.Now()
	}
	// Write or update the tail to show Running state
	if r.cur.tailKind != tailAgent {
		r.writeTail(tailAgent, r.renderAgentStateLine())
	} else {
		r.rewriteAgentTail()
	}
}

// rewriteAgentTail refreshes the live Agent Attempt line in place.
func (r *interactiveRenderer) rewriteAgentTail() {
	r.rewriteTail(r.renderAgentStateLine())
}

// renderAgentStateLine renders the single-line Agent Attempt status (header
// and state combined on ONE row) based on current state.
func (r *interactiveRenderer) renderAgentStateLine() string {
	e := r.cur
	switch e.agentState {
	case agentStateRunning:
		return fmt.Sprintf("Agent Attempt: 🔄 %s", r.sty.Info("Running"))
	case agentStateCompleted:
		return fmt.Sprintf("Agent Attempt: %s", r.sty.OK("✅ Completed"))
	case agentStateGuardrail:
		msg := "⚠ Guardrail hit"
		if e.agentGuardrailMsg != "" {
			msg = fmt.Sprintf("⚠ Guardrail hit — %s", e.agentGuardrailMsg)
		}
		return fmt.Sprintf("Agent Attempt: %s", r.sty.Warn(msg))
	default:
		return "Agent Attempt: 🔄 Running"
	}
}

// agentComplete transitions the Agent Attempt state to Completed or Guardrail.
// Called from EventPassed/Failed (where the generation phase has ended).
func (r *interactiveRenderer) agentComplete(fileCount int, success bool, guardrailReason string) {
	e := r.cur
	if !e.agentHeaderPrinted {
		return
	}
	// Determine the final state
	if guardrailReason != "" {
		e.agentState = agentStateGuardrail
		e.agentGuardrailMsg = guardrailReason
	} else if success {
		e.agentState = agentStateCompleted
	} else {
		// Non-guardrail failure (e.g., review failed) — treat as completed
		// since the generation itself finished
		e.agentState = agentStateCompleted
	}
	// Update the tail to show final state
	line := r.renderAgentStateLine()
	if e.tailKind == tailAgent {
		r.rewriteTail(line)
		r.freezeTail()
		return
	}
	// Tail no longer ours. If the agent's "Running" line was frozen above
	// (typically because a grader took over the tail), rewrite that exact
	// row in place so the Completed status replaces the stale Running text
	// rather than getting appended at the bottom of the transcript.
	if e.agentLineFrozen {
		r.freezeTail()
		r.rewriteFrozenLine(e.agentLineRow, line)
		e.agentLineFrozen = false
		return
	}
	// No prior agent line at all — append a fresh row.
	r.freezeTail()
	r.writeLine(line)
}

// --- Session Details section ---

func (r *interactiveRenderer) onSessionDetails(evt ProgressEvent) {
	// EventSessionDetails signals that the generation phase has completed.
	// First, flip the Agent Attempt line to "Completed" (if it's still
	// active) so the completion state is frozen BEFORE graders take over
	// the tail. This fixes Bug 2: without this early call to agentComplete,
	// the Agent Attempt line would stay as "Running" until the terminal
	// event (EventPassed/Failed) arrived after ALL graders completed, at
	// which point the frozen row index would be stale.
	if r.cur.agentHeaderPrinted && r.cur.agentState == agentStateRunning {
		r.agentComplete(evt.FileCount, true, "")
	}
	// Now render the Session Details section.
	if r.cur.sessionPrinted {
		return
	}
	r.cur.sessionPrinted = true
	r.freezeTail()
	r.writeLine("Session Details:")
	if len(evt.Files) > 0 {
		r.writeLine(fmt.Sprintf("  Files: %s", joinTruncated(evt.Files, 6)))
	}
	r.writeLine(fmt.Sprintf("  Turns: %d   Tool calls: %d   Tokens: %s in / %s out",
		evt.Turns, evt.ToolCalls, fmtTokens(evt.InputTokens), fmtTokens(evt.OutputTokens)))
}

// --- Graders section ---

// displayKind maps internal grader kinds to user-facing labels.
// Prompt-review and prompt graders both display as "prompt".
func displayKind(kind string) string {
	switch kind {
	case "prompt_review", "prompt":
		return "prompt"
	default:
		return kind
	}
}

// graderIndent returns the indentation prefix for grader header lines.
// When a file grouping header has been printed, graders are indented one
// extra level (4 spaces). Otherwise they use the legacy 2-space prefix.
func graderIndent(hasFileGroup bool) string {
	if hasFileGroup {
		return "    "
	}
	return "  "
}

// pointIndent returns the indentation prefix for point lines.
func pointIndent(hasFileGroup bool) string {
	if hasFileGroup {
		return "        "
	}
	return "    "
}

// ensureGraderFileHeader prints a file-level grouping header when the
// GraderSourceFile on the incoming event differs from the last one printed.
// It is a no-op when sourceFile is empty (graceful degradation for pre-Neo data).
func (r *interactiveRenderer) ensureGraderFileHeader(sourceFile, sourceType string) {
	if sourceFile == "" || sourceFile == r.cur.lastGraderSourceFile {
		return
	}
	r.cur.lastGraderSourceFile = sourceFile

	base := sourceFile
	// Use just the filename component for display.
	if idx := len(sourceFile) - 1; idx >= 0 {
		for i := len(sourceFile) - 1; i >= 0; i-- {
			if sourceFile[i] == '/' || sourceFile[i] == '\\' {
				base = sourceFile[i+1:]
				break
			}
		}
	}

	typeLabel := ""
	switch sourceType {
	case "prompt_file":
		typeLabel = " (prompt file)"
	case "criteria_file":
		typeLabel = " (criteria file)"
	}

	r.writeLine(fmt.Sprintf("  - %s%s:", base, typeLabel))
}

func (r *interactiveRenderer) ensureGradersHeader() {
	if r.cur.gradersHeaderPrinted {
		return
	}
	r.freezeTail()
	r.writeLine("Graders:")
	r.cur.gradersHeaderPrinted = true
}

func (r *interactiveRenderer) onGraderStart(evt ProgressEvent) {
	r.ensureGradersHeader()
	// Serialized: previous grader's line (if any) is already frozen via its
	// own GraderComplete. The Running line becomes the new tail.
	r.freezeTail()
	r.ensureGraderFileHeader(evt.GraderSourceFile, evt.GraderSourceType)
	indent := graderIndent(evt.GraderSourceFile != "")
	line := fmt.Sprintf("%s- %s %s: 🔄 %s",
		indent,
		evt.GraderID,
		r.sty.Muted("("+displayKind(evt.GraderKind)+")"),
		r.sty.Muted("Running…"))
	r.writeTail(tailGrader, line)
	// Track which grader owns the current tail so that, if some other
	// section commits the tail before this grader's Complete event fires,
	// freezeTail can record the row index for later in-place rewrite.
	r.cur.pendingGraderID = evt.GraderID
}

func (r *interactiveRenderer) onGraderComplete(evt ProgressEvent) {
	r.ensureGradersHeader()

	// Multi-point graders render as a header line + one indented row per
	// point. Zero-point graders fall back to the legacy flat row to
	// preserve existing UX for graders without point-level detail.
	if len(evt.Points) >= 1 {
		r.renderGraderWithPoints(evt)
		return
	}

	score := ""
	if evt.Score != nil {
		score = fmt.Sprintf(" (%s/10)", formatScore(*evt.Score))
	}
	var outcome string
	switch evt.Result {
	case GraderResultPass:
		outcome = r.sty.OK("✅ Pass") + score
	case GraderResultFail:
		outcome = r.sty.Fail("❌ Fail") + score
	default:
		outcome = evt.Result + score
	}
	indent := graderIndent(evt.GraderSourceFile != "")
	line := fmt.Sprintf("%s- %s %s: %s",
		indent,
		evt.GraderID,
		r.sty.Muted("("+displayKind(evt.GraderKind)+")"),
		outcome)
	if evt.Message != "" && evt.Result == GraderResultFail {
		line += " " + r.sty.Muted("— "+evt.Message)
	}
	// Happy path: this grader still owns the tail.
	if r.cur.tailKind == tailGrader && r.cur.pendingGraderID == evt.GraderID {
		r.rewriteTail(line)
		r.freezeTail()
		// freezeTail will have cleared pendingGraderID and recorded the
		// row in graderRowByID. Drop the row entry — the line is now in
		// its final state and shouldn't be rewritten again.
		if r.cur.graderRowByID != nil {
			delete(r.cur.graderRowByID, evt.GraderID)
		}
		return
	}
	// Tail no longer ours. If a Running row for this grader was frozen
	// above (because another section took over the tail), rewrite that
	// exact row in place — never append a duplicate entry at the bottom.
	if row, ok := r.cur.graderRowByID[evt.GraderID]; ok {
		r.freezeTail()
		r.rewriteFrozenLine(row, line)
		delete(r.cur.graderRowByID, evt.GraderID)
		return
	}
	// No prior row for this grader at all — append fresh.
	r.freezeTail()
	r.writeLine(line)
}

// renderGraderWithPoints emits a multi-line block for a grader that reported
// multiple Points: a header line summarizing the aggregate (e.g.
// "❌ Fail (2/3)") followed by one indented row per Point. Reuses the same
// in-place rewrite scaffolding as the single-row path so the header replaces
// the original "Running…" row when one was previously frozen.
func (r *interactiveRenderer) renderGraderWithPoints(evt ProgressEvent) {
	passed := 0
	for _, p := range evt.Points {
		if p.Pass {
			passed++
		}
	}
	total := len(evt.Points)
	allPassed := passed == total
	badge := r.sty.Fail(fmt.Sprintf("❌ Fail (%d/%d)", passed, total))
	if allPassed {
		badge = r.sty.OK(fmt.Sprintf("✅ Pass (%d/%d)", passed, total))
	}
	hasFile := evt.GraderSourceFile != ""
	indent := graderIndent(hasFile)
	header := fmt.Sprintf("%s- %s %s: %s",
		indent,
		evt.GraderID,
		r.sty.Muted("("+displayKind(evt.GraderKind)+")"),
		badge)

	// Place the header. Three cases mirror the flat path:
	//   1. This grader still owns the tail — rewrite tail then freeze.
	//   2. The Running row was frozen above — rewrite in place.
	//   3. No prior row — append fresh.
	switch {
	case r.cur.tailKind == tailGrader && r.cur.pendingGraderID == evt.GraderID:
		r.rewriteTail(header)
		r.freezeTail()
		if r.cur.graderRowByID != nil {
			delete(r.cur.graderRowByID, evt.GraderID)
		}
	default:
		if row, ok := r.cur.graderRowByID[evt.GraderID]; ok {
			r.freezeTail()
			r.rewriteFrozenLine(row, header)
			delete(r.cur.graderRowByID, evt.GraderID)
		} else {
			r.freezeTail()
			r.writeLine(header)
		}
	}

	// One indented row per point under the header. Point names are
	// soft-truncated for terminal output so very long check strings don't
	// wrap awkwardly; full text remains in the report.
	const maxPointNameWidth = 50
	pIndent := pointIndent(hasFile)
	for _, p := range evt.Points {
		var status string
		if p.Pass {
			status = r.sty.OK("✅ Pass")
		} else {
			status = r.sty.Fail("❌ Fail")
		}
		name := truncateToWidth(p.Label, maxPointNameWidth)
		line := fmt.Sprintf("%s- %s: %s", pIndent, name, status)
		if !p.Pass && p.Message != "" {
			line += " " + r.sty.Muted("— "+p.Message)
		}
		r.writeLine(line)
	}
}

// --- Terminal events ---

func (r *interactiveRenderer) onPassed(evt ProgressEvent) {
	// Safety: if tools verification never arrived, open the gate now so any
	// buffered agent events are flushed before we finalize.
	if !r.cur.agentGateOpen {
		r.openAgentGate()
	}
	// Only call agentComplete if the Agent Attempt hasn't already transitioned
	// to Completed. EventSessionDetails now triggers agentComplete early (before
	// graders run), so by the time EventPassed arrives, the agent line is
	// already frozen in its final state. Calling agentComplete again would
	// attempt a stale rewriteFrozenLine or append a duplicate row.
	if r.cur.agentHeaderPrinted && r.cur.agentState != agentStateCompleted && r.cur.agentState != agentStateGuardrail {
		r.agentComplete(evt.FileCount, true, evt.GuardrailReason)
	}
	r.completed++
	r.passed++
	r.cur.terminalStatus = evalPassed
	// Print blank line + total checks line when grader points are available
	if evt.GraderChecksTotal > 0 {
		r.writeLine("")
		r.writeLine(fmt.Sprintf("  %s", r.sty.OK(fmt.Sprintf("✅ Total checks that passed across all graders: %d/%d", evt.GraderChecksPassed, evt.GraderChecksTotal))))
	}
	r.finalizeCurrentEval()
}

func (r *interactiveRenderer) onFailed(evt ProgressEvent) {
	// Safety: if tools verification never arrived, open the gate now.
	if !r.cur.agentGateOpen {
		r.openAgentGate()
	}
	// Only call agentComplete if the Agent Attempt hasn't already transitioned.
	// EventSessionDetails triggers agentComplete early, so this is a no-op in
	// the happy path where generation succeeded but graders failed.
	if r.cur.agentHeaderPrinted && r.cur.agentState != agentStateCompleted && r.cur.agentState != agentStateGuardrail {
		r.agentComplete(evt.FileCount, false, evt.GuardrailReason)
	}
	r.completed++
	r.failed++
	r.cur.terminalStatus = evalFailed
	r.cur.terminalMsg = evt.Message
	r.freezeTail()
	// Print blank line + total checks line when grader points are available
	if evt.GraderChecksTotal > 0 {
		r.writeLine("")
		r.writeLine(fmt.Sprintf("  %s", r.sty.Fail(fmt.Sprintf("❌ Total checks that passed across all graders: %d/%d", evt.GraderChecksPassed, evt.GraderChecksTotal))))
	} else if evt.Message != "" {
		// Fallback: print the message when no grader points are available (legacy behavior)
		msg := evt.Message
		if msg == "" {
			msg = "failed"
		}
		r.writeLine(fmt.Sprintf("  %s", r.sty.Fail("❌ "+msg)))
	}
	r.finalizeCurrentEval()
}

func (r *interactiveRenderer) onError(evt ProgressEvent) {
	// Safety: if tools verification never arrived, open the gate now.
	if !r.cur.agentGateOpen {
		r.openAgentGate()
	}
	// Only call agentComplete if the Agent Attempt hasn't already transitioned.
	// For errors that occur during generation (before EventSessionDetails),
	// this will correctly mark the attempt as completed. For errors after
	// graders (rare), this is a no-op since the state is already final.
	if r.cur.agentHeaderPrinted && r.cur.agentState != agentStateCompleted && r.cur.agentState != agentStateGuardrail {
		r.agentComplete(0, false, evt.GuardrailReason)
	}
	r.completed++
	r.errors++
	r.cur.terminalStatus = evalError
	r.cur.terminalMsg = evt.Message
	r.freezeTail()
	msg := evt.Message
	if msg == "" {
		msg = "ERROR"
	}
	r.writeLine(fmt.Sprintf("  %s", r.sty.Fail("❌ "+msg)))
	r.finalizeCurrentEval()
}

// --- Tail primitives ---

// writeLine writes text followed by a newline. Requires the tail to be frozen
// (no active tail) — caller must freezeTail() first.
func (r *interactiveRenderer) writeLine(s string) {
	fmt.Fprintln(r.w, s)
	if r.cur != nil {
		r.cur.linesWritten++
	}
}

// writeTail writes text as the current tail line (no trailing newline).
// Freezes any existing tail first. kind records which section owns the tail
// so the ticker knows whether to refresh it. Truncates to terminal width minus
// a safety margin to prevent line wrapping (which would break in-place rewrites
// — \r\033[2K only clears one physical row). The margin avoids cursor-wrapping
// ambiguity when text is exactly terminal width.
func (r *interactiveRenderer) writeTail(kind tailKind, text string) {
	r.freezeTail()
	w := TermWidth()
	if w <= 0 {
		w = 80 // fallback
	}
	// Subtract 2 columns as safety margin: avoids the "exactly terminal width"
	// cursor wrapping ambiguity where some terminals wrap immediately while
	// others delay until the next character.
	maxWidth := w - 2
	if maxWidth < 10 {
		maxWidth = w // terminal too narrow, skip margin
	}
	text = truncateToWidth(text, maxWidth)
	r.w.Write([]byte(text))
	r.cur.tailKind = kind
	r.cur.tailText = text
	// Track how many physical rows this tail occupies. visibleWidth counts
	// ANSI-stripped cells, so we compute rows as ceil(visible / width).
	visible := visibleWidth(text)
	r.cur.tailRowCount = (visible + w - 1) / w
	if r.cur.tailRowCount < 1 {
		r.cur.tailRowCount = 1
	}
}

// rewriteTail replaces the current tail line's content in place. No-op if
// there is no active tail. Clears ALL physical rows the previous tail occupied
// before writing the new content. Applies the same safety margin as writeTail.
func (r *interactiveRenderer) rewriteTail(text string) {
	if r.cur == nil || r.cur.tailKind == tailNone {
		return
	}
	w := TermWidth()
	if w <= 0 {
		w = 80 // fallback
	}
	// Apply the same safety margin as writeTail
	maxWidth := w - 2
	if maxWidth < 10 {
		maxWidth = w
	}
	text = truncateToWidth(text, maxWidth)

	var buf bytes.Buffer
	// If the previous tail occupied multiple rows, we need to clear all of them.
	// The cursor is currently at the end of the last row. Move back to the start
	// of the first row and clear each row.
	oldRows := r.cur.tailRowCount
	if oldRows > 1 {
		// Move cursor up (oldRows - 1) lines to the start of the tail.
		fmt.Fprintf(&buf, "\x1b[%dA", oldRows-1)
	}
	// Now at the start row. Clear each row from top to bottom.
	for i := 0; i < oldRows; i++ {
		buf.WriteString(ansiCR)
		buf.WriteString(ansiClearLine)
		if i < oldRows-1 {
			buf.WriteString("\n") // move to next row
		}
	}
	// Cursor is now at column 0 of the last (bottom) row, all old content cleared.
	// Write the new tail.
	buf.WriteString(text)
	r.w.Write(buf.Bytes())
	r.cur.tailText = text

	// Update row count for the new tail.
	visible := visibleWidth(text)
	r.cur.tailRowCount = (visible + w - 1) / w
	if r.cur.tailRowCount < 1 {
		r.cur.tailRowCount = 1
	}
}

// freezeTail commits the current tail line with a trailing newline, so it
// becomes immutable. No-op if there is no active tail. Records the row of
// agent / grader tails so a later Complete event can rewrite the original
// row in place (see rewriteFrozenLine).
func (r *interactiveRenderer) freezeTail() {
	if r.cur == nil || r.cur.tailKind == tailNone {
		return
	}
	switch r.cur.tailKind {
	case tailAgent:
		// The agent line currently sits on terminal row r.cur.linesWritten
		// (the tail row, before the trailing newline is written below).
		r.cur.agentLineFrozen = true
		r.cur.agentLineRow = r.cur.linesWritten
	case tailGrader:
		if r.cur.pendingGraderID != "" {
			if r.cur.graderRowByID == nil {
				r.cur.graderRowByID = make(map[string]int)
			}
			r.cur.graderRowByID[r.cur.pendingGraderID] = r.cur.linesWritten
			r.cur.pendingGraderID = ""
		}
	}
	r.w.Write([]byte{'\n'})
	r.cur.linesWritten++
	r.cur.tailKind = tailNone
	r.cur.tailText = ""
	r.cur.tailRowCount = 0
}

// rewriteFrozenLine replaces the content of an already-committed line at
// terminal row `row` (0-indexed; the same coordinate space as
// linesWritten). Uses the same DECSC/DECRC bracketed save-and-restore
// pattern as redrawToolsBlock, so the cursor returns to its original
// position after the rewrite. Caller must freezeTail() first — the cursor
// must be at column 0 of row r.cur.linesWritten when this is called.
//
// No-op when row >= r.cur.linesWritten (line not actually frozen above
// the cursor) — callers should fall back to writeLine in that case.
func (r *interactiveRenderer) rewriteFrozenLine(row int, text string) {
	if r.cur == nil {
		return
	}
	up := r.cur.linesWritten - row
	if up <= 0 {
		return
	}
	w := TermWidth()
	if w <= 0 {
		w = 80
	}
	maxWidth := w - 2
	if maxWidth < 10 {
		maxWidth = w
	}
	text = truncateToWidth(text, maxWidth)

	var buf bytes.Buffer
	buf.WriteString(ansiSaveCursor)
	fmt.Fprintf(&buf, ansiCursorUpFmt, up)
	buf.WriteString(ansiCR)
	buf.WriteString(ansiClearLine)
	buf.WriteString(text)
	buf.WriteString(ansiRestoreCurs)
	r.w.Write(buf.Bytes())
}

// --- Finish / counters ---

func (r *interactiveRenderer) finish() {
	// Stop ticker first so tickLoop doesn't race with our final output.
	if r.ticker != nil {
		r.ticker.Stop()
		close(r.stopCh)
		r.wg.Wait()
		r.ticker = nil
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.cur != nil {
		r.finalizeCurrentEval()
	}

	elapsed := time.Since(r.startTime)
	fmt.Fprintf(r.w, "Summary: %d/%d passed", r.passed, r.total)
	if r.passed > 0 {
		fmt.Fprintf(r.w, "  %s", r.sty.OK(fmt.Sprintf("✅ %d", r.passed)))
	}
	if r.failed > 0 {
		fmt.Fprintf(r.w, "  %s", r.sty.Fail(fmt.Sprintf("❌ %d", r.failed)))
	}
	if r.errors > 0 {
		fmt.Fprintf(r.w, "  %s", r.sty.Fail(fmt.Sprintf("❌ %d errors", r.errors)))
	}
	fmt.Fprintf(r.w, "  %s\n", r.sty.Muted("Duration: "+fmtClock(elapsed)))
	if r.reportDir != "" {
		fmt.Fprintf(r.w, "Reports: %s\n", r.reportDir)
	}
}

func (r *interactiveRenderer) completedCount() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.completed
}

// --- formatting helpers ---

// fmtClock renders a duration as H:MM:SS or MM:SS depending on length.
// Layout in plan.md shows HH:MM:SS — we use MM:SS for sub-hour runs (typical)
// and fall back to H:MM:SS when needed.
func fmtClock(d time.Duration) string {
	if d < 0 {
		d = 0
	}
	total := int(d.Seconds())
	h := total / 3600
	m := (total % 3600) / 60
	s := total % 60
	if h > 0 {
		return fmt.Sprintf("%d:%02d:%02d", h, m, s)
	}
	return fmt.Sprintf("%02d:%02d", m, s)
}

func fmtCost(c float64) string {
	if c <= 0 {
		return "—"
	}
	return fmt.Sprintf("$%.2f", c)
}

// fmtTokens renders a token count with thousands separators ("12,345"),
// or "—" when zero/unknown. Replaces the deprecated fmtCost in the live
// session-details line.
func fmtTokens(n int) string {
	if n <= 0 {
		return "—"
	}
	s := fmt.Sprintf("%d", n)
	// Insert commas from the right.
	out := make([]byte, 0, len(s)+len(s)/3)
	for i, ch := range []byte(s) {
		if i > 0 && (len(s)-i)%3 == 0 {
			out = append(out, ',')
		}
		out = append(out, ch)
	}
	return string(out)
}

func formatScore(score float64) string {
	if score == float64(int(score)) {
		return fmt.Sprintf("%d", int(score))
	}
	return fmt.Sprintf("%.1f", score)
}

func joinTruncated(items []string, max int) string {
	if len(items) == 0 {
		return "—"
	}
	if len(items) <= max {
		return joinWithSep(items, ", ")
	}
	head := joinWithSep(items[:max], ", ")
	return fmt.Sprintf("%s, … (%d more)", head, len(items)-max)
}

func joinWithSep(items []string, sep string) string {
	var buf bytes.Buffer
	for i, s := range items {
		if i > 0 {
			buf.WriteString(sep)
		}
		buf.WriteString(s)
	}
	return buf.String()
}
