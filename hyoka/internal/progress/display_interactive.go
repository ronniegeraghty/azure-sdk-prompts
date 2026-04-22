package progress

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/progress/style"
)

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
//	Agent Attempt:
//	  🔄 Running… turn X/Y, N tool calls  (MM:SS) (tail; per-second ticker refresh)
//	  ✅ Complete — N files, M turns        (MM:SS)
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
	tailKind tailKind
	tailText string // raw text of tail (for bookkeeping; re-emit on update)

	// Tools section.
	toolsHeaderPrinted bool
	toolsFirstLine     int              // linesWritten index of first tool line ("  - …")
	toolLines          []toolLine       // one entry per emitted tool line, in order
	toolsVerified      bool             // guard so we only redraw once
	toolIndexByName    map[string]int   // name -> index in toolLines

	// Agent Attempt section.
	agentHeaderPrinted bool
	agentActivity      string
	agentStartTime     time.Time
	agentTurn          int
	agentMaxTurns      int
	agentToolCalls     int
	agentDone          bool
	agentFileCount     int
	agentTurns         int
	agentDuration      time.Duration

	// Session Details section.
	sessionPrinted bool

	// Graders section.
	gradersHeaderPrinted bool

	// Terminal state for this eval — set when EventPassed/Failed/Error fires.
	terminalStatus evalStatus
	terminalMsg    string
}

// toolLine captures what we printed for a tool line, so we can compare
// against the bulk verification event and redraw on flips.
type toolLine struct {
	name   string
	kind   string
	status string // "loading" | ToolStatusLoaded | ToolStatusFailed
	reason string
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

// ANSI cursor helpers. Kept private to this file so the rest of the package
// doesn't rely on raw escapes.
const (
	ansiClearLine    = "\x1b[2K"
	ansiCR           = "\r"
	ansiCursorUp1    = "\x1b[1A"
	ansiSaveCursor   = "\x1b7" // DECSC — same convention display.go uses
	ansiRestoreCurs  = "\x1b8" // DECRC
	ansiCursorUpFmt  = "\x1b[%dA"
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

// tickLoop refreshes the Agent Attempt tail line's duration counter once per
// second. All other tail updates are driven by incoming events.
func (r *interactiveRenderer) tickLoop() {
	defer r.wg.Done()
	for {
		select {
		case <-r.stopCh:
			return
		case <-r.ticker.C:
			r.mu.Lock()
			if r.cur != nil && r.tailIsAgent() && !r.cur.agentDone {
				r.rewriteAgentTail()
			}
			r.mu.Unlock()
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
	// Freeze whatever prior tool line held the tail (it's final: loaded/failed
	// was set before the next Start arrived, per the sequential emission
	// contract in neo-tool-resolution-wiring.md).
	r.freezeTail()
	tl := toolLine{
		name:   evt.ToolName,
		kind:   evt.ToolKind,
		status: "loading",
	}
	r.cur.toolLines = append(r.cur.toolLines, tl)
	r.cur.toolIndexByName[evt.ToolName] = len(r.cur.toolLines) - 1
	r.writeTail(tailTool, r.renderToolLine(tl))
}

func (r *interactiveRenderer) onToolResolutionResult(evt ProgressEvent) {
	idx, ok := r.cur.toolIndexByName[evt.ToolName]
	if !ok {
		// Result without a Start — defensive: behave like a fresh start then
		// complete.
		r.ensureToolsHeader()
		r.freezeTail()
		r.cur.toolLines = append(r.cur.toolLines, toolLine{})
		idx = len(r.cur.toolLines) - 1
		r.cur.toolIndexByName[evt.ToolName] = idx
		r.writeTail(tailTool, "")
	}
	tl := &r.cur.toolLines[idx]
	tl.name = evt.ToolName
	tl.kind = evt.ToolKind
	tl.status = evt.Status
	tl.reason = evt.Reason
	// Update the tail in place. The tool line IS the tail here, because
	// ToolResolutionResult always immediately follows its Start in the same
	// sequential tool-resolution pass.
	if r.cur.tailKind == tailTool {
		r.rewriteTail(r.renderToolLine(*tl))
	} else {
		// Tail has moved on (shouldn't happen per schema, but be safe:
		// trigger a block redraw).
		r.redrawToolsBlock()
	}
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
			return
		}
		r.ensureToolsHeader()
		for _, t := range evt.Tools {
			tl := toolLine{name: t.ToolName, kind: t.ToolKind, status: t.Status, reason: t.Reason}
			r.cur.toolLines = append(r.cur.toolLines, tl)
			r.cur.toolIndexByName[t.ToolName] = len(r.cur.toolLines) - 1
			r.freezeTail()
			r.writeLine(r.renderToolLine(tl))
		}
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
	for i, tl := range r.cur.toolLines {
		if !seen[tl.name] && tl.status == ToolStatusLoaded {
			r.cur.toolLines[i].status = ToolStatusFailed
			if r.cur.toolLines[i].reason == "" {
				r.cur.toolLines[i].reason = "not reported by SDK"
			}
			dirty = true
		}
	}
	if dirty {
		r.redrawToolsBlock()
	}
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
	for _, tl := range e.toolLines {
		buf.WriteString(ansiClearLine)
		buf.WriteString(r.renderToolLine(tl))
		buf.WriteByte('\n')
	}
	buf.WriteString(ansiRestoreCurs)
	r.w.Write(buf.Bytes())
}

func (r *interactiveRenderer) renderToolLine(tl toolLine) string {
	name := tl.name
	kind := r.sty.Muted(fmt.Sprintf("(%s)", tl.kind))
	switch tl.status {
	case "", "loading":
		return fmt.Sprintf("  - %s %s: 🔄 %s", name, kind, r.sty.Muted("Loading…"))
	case ToolStatusLoaded:
		return fmt.Sprintf("  - %s %s: %s", name, kind, r.sty.OK("✅ Loaded"))
	case ToolStatusFailed:
		reason := tl.reason
		if reason == "" {
			reason = "failed"
		}
		return fmt.Sprintf("  - %s %s: %s %s", name, kind,
			r.sty.Fail("❌ Failed"), r.sty.Muted("("+reason+")"))
	default:
		return fmt.Sprintf("  - %s %s: %s", name, kind, tl.status)
	}
}

// --- Agent Attempt section ---

func (r *interactiveRenderer) ensureAgentHeader() {
	if r.cur.agentHeaderPrinted {
		return
	}
	r.freezeTail()
	r.writeLine("Agent Attempt:")
	r.cur.agentHeaderPrinted = true
	r.cur.agentStartTime = time.Now()
}

func (r *interactiveRenderer) onPhaseChange(evt ProgressEvent) {
	// Phase events are informational in interactive mode — activity/agent
	// framing is already driven by tool/file/reasoning events.
	if evt.Phase == PhaseGenerating {
		r.ensureAgentHeader()
		if r.cur.tailKind != tailAgent {
			r.writeTail(tailAgent, r.renderAgentTail())
		}
	}
}

func (r *interactiveRenderer) onAgentActivity(evt ProgressEvent) {
	r.ensureAgentHeader()
	// Update rolling activity counters.
	if evt.Type == EventToolStart {
		r.cur.agentToolCalls++
	}
	if evt.Message != "" {
		r.cur.agentActivity = evt.Message
	}
	if r.cur.tailKind != tailAgent {
		r.writeTail(tailAgent, r.renderAgentTail())
	} else {
		r.rewriteAgentTail()
	}
}

// rewriteAgentTail refreshes the live Agent Attempt line in place.
func (r *interactiveRenderer) rewriteAgentTail() {
	r.rewriteTail(r.renderAgentTail())
}

func (r *interactiveRenderer) renderAgentTail() string {
	e := r.cur
	dur := time.Since(e.agentStartTime)
	if e.agentStartTime.IsZero() {
		dur = time.Since(e.startTime)
	}
	act := e.agentActivity
	if act == "" {
		act = "Starting…"
	}
	if e.agentToolCalls > 0 {
		act = fmt.Sprintf("%s · %d tool calls", act, e.agentToolCalls)
	}
	return fmt.Sprintf("  🔄 %s   %s",
		r.sty.Info(act),
		r.sty.Muted("("+fmtClock(dur)+")"))
}

// agentComplete freezes the Agent Attempt tail into a Complete line.
// Called from EventPassed/Failed (where the generation phase has ended).
func (r *interactiveRenderer) agentComplete(fileCount int, success bool) {
	e := r.cur
	if !e.agentHeaderPrinted || e.agentDone {
		return
	}
	e.agentDone = true
	e.agentFileCount = fileCount
	if e.agentStartTime.IsZero() {
		e.agentDuration = time.Since(e.startTime)
	} else {
		e.agentDuration = time.Since(e.agentStartTime)
	}
	var line string
	if success {
		line = fmt.Sprintf("  %s — %d files, %d tool calls   %s",
			r.sty.OK("✅ Complete"),
			e.agentFileCount,
			e.agentToolCalls,
			r.sty.Muted("("+fmtClock(e.agentDuration)+")"))
	} else {
		line = fmt.Sprintf("  %s   %s",
			r.sty.Fail("❌ Failed"),
			r.sty.Muted("("+fmtClock(e.agentDuration)+")"))
	}
	if e.tailKind == tailAgent {
		r.rewriteTail(line)
	} else {
		r.freezeTail()
		r.writeLine(line)
		return
	}
	r.freezeTail()
}

// --- Session Details section ---

func (r *interactiveRenderer) onSessionDetails(evt ProgressEvent) {
	if r.cur.sessionPrinted {
		return
	}
	r.cur.sessionPrinted = true
	r.freezeTail()
	r.writeLine("Session Details:")
	if len(evt.Files) > 0 {
		r.writeLine(fmt.Sprintf("  Files: %s", joinTruncated(evt.Files, 6)))
	}
	r.writeLine(fmt.Sprintf("  Turns: %d   Tool calls: %d   Cost: %s",
		evt.Turns, evt.ToolCalls, fmtCost(evt.Cost)))
}

// --- Graders section ---

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
	line := fmt.Sprintf("  - %s %s: 🔄 %s",
		evt.GraderID,
		r.sty.Muted("("+evt.GraderKind+")"),
		r.sty.Muted("Running…"))
	r.writeTail(tailGrader, line)
}

func (r *interactiveRenderer) onGraderComplete(evt ProgressEvent) {
	r.ensureGradersHeader()
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
	line := fmt.Sprintf("  - %s %s: %s",
		evt.GraderID,
		r.sty.Muted("("+evt.GraderKind+")"),
		outcome)
	if evt.Message != "" && evt.Result == GraderResultFail {
		line += " " + r.sty.Muted("— "+evt.Message)
	}
	if r.cur.tailKind == tailGrader {
		r.rewriteTail(line)
		r.freezeTail()
	} else {
		r.freezeTail()
		r.writeLine(line)
	}
}

// --- Terminal events ---

func (r *interactiveRenderer) onPassed(evt ProgressEvent) {
	r.agentComplete(evt.FileCount, true)
	r.completed++
	r.passed++
	r.cur.terminalStatus = evalPassed
	r.finalizeCurrentEval()
}

func (r *interactiveRenderer) onFailed(evt ProgressEvent) {
	r.agentComplete(evt.FileCount, false)
	r.completed++
	r.failed++
	r.cur.terminalStatus = evalFailed
	r.cur.terminalMsg = evt.Message
	r.freezeTail()
	msg := evt.Message
	if msg == "" {
		msg = "failed"
	}
	r.writeLine(fmt.Sprintf("  %s", r.sty.Fail("❌ "+msg)))
	r.finalizeCurrentEval()
}

func (r *interactiveRenderer) onError(evt ProgressEvent) {
	r.agentComplete(0, false)
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
// so the ticker knows whether to refresh it.
func (r *interactiveRenderer) writeTail(kind tailKind, text string) {
	r.freezeTail()
	r.w.Write([]byte(text))
	r.cur.tailKind = kind
	r.cur.tailText = text
}

// rewriteTail replaces the current tail line's content in place. No-op if
// there is no active tail.
func (r *interactiveRenderer) rewriteTail(text string) {
	if r.cur == nil || r.cur.tailKind == tailNone {
		return
	}
	// Move to column 0, clear the line, rewrite. Cursor stays on the same
	// physical row — linesWritten is unchanged because we never committed
	// a newline.
	var buf bytes.Buffer
	buf.WriteString(ansiCR)
	buf.WriteString(ansiClearLine)
	buf.WriteString(text)
	r.w.Write(buf.Bytes())
	r.cur.tailText = text
}

// freezeTail commits the current tail line with a trailing newline, so it
// becomes immutable. No-op if there is no active tail.
func (r *interactiveRenderer) freezeTail() {
	if r.cur == nil || r.cur.tailKind == tailNone {
		return
	}
	r.w.Write([]byte{'\n'})
	r.cur.linesWritten++
	r.cur.tailKind = tailNone
	r.cur.tailText = ""
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
