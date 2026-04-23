package progress

// CI-mode renderer: append-only, timestamped start/finish lines during the
// run, followed by an end-of-run summary table. Designed for redirected
// output and log aggregators — never moves the cursor.
//
// Activated by ProgressMode "ci" (preferred) or the legacy alias "log",
// which used to drive a more verbose append-only renderer. The new CI
// renderer replaces the legacy log behavior: the old per-phase lines are
// gone in favor of one start line and one finish line per eval.

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"time"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/progress/style"
)

// ciEvalState captures the info needed to render the per-eval start/finish
// lines during the run and the row in the end-of-run summary table.
type ciEvalState struct {
	promptID    string
	configName  string
	startTime   time.Time
	duration    time.Duration
	result      string // "pass", "fail", "error"
	message     string // failure reason (one line)
	graderPass  int
	graderTotal int
	finished    bool
	tools       []ToolStatus // tools loaded for this eval
	toolsPrinted bool        // guard: only print tools section once
}

// ciRenderer is the append-only renderer used for ModeCI and ModeLog.
// All methods must be called with the enclosing Display's mutex held.
type ciRenderer struct {
	w         io.Writer
	st        *style.Styler
	startTime time.Time
	total     int
	workers   int
	configs   int
	reportDir string
	useEmoji  bool

	evals map[string]*ciEvalState
	order []string // eval IDs in first-seen order for the summary table
}

// newCIRenderer constructs a CI renderer and prints the intro line.
// workers/total/configs are advisory — the intro line falls back gracefully
// when they are unknown.
func newCIRenderer(w io.Writer, total, workers, configs int, reportDir string) *ciRenderer {
	st := style.New(w)
	r := &ciRenderer{
		w:         w,
		st:        st,
		startTime: time.Now(),
		total:     total,
		workers:   workers,
		configs:   configs,
		reportDir: reportDir,
		useEmoji:  st.Enabled, // drop emoji when color is disabled (NO_COLOR / non-TTY)
		evals:     make(map[string]*ciEvalState),
	}
	r.writeIntro()
	return r
}

func (r *ciRenderer) writeIntro() {
	var intro string
	switch {
	case r.total > 0 && r.configs > 0:
		intro = fmt.Sprintf("Running %d evals across %d configs with %d workers…", r.total, r.configs, r.workers)
	case r.total > 0:
		intro = fmt.Sprintf("Running %d evals with %d workers…", r.total, r.workers)
	default:
		intro = fmt.Sprintf("Running evals with %d workers…", r.workers)
	}
	fmt.Fprintln(r.w, intro)
	fmt.Fprintln(r.w)
}

// timestamp renders a [HH:MM:SS] marker relative to run start. Wrapped in
// Muted style when colors are enabled; bare brackets otherwise.
func (r *ciRenderer) timestamp(at time.Time) string {
	d := at.Sub(r.startTime)
	if d < 0 {
		d = 0
	}
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	return r.st.Muted(fmt.Sprintf("[%02d:%02d:%02d]", h, m, s))
}

// startGlyph / passGlyph / failGlyph produce the leading glyph for the
// per-event lines. Emoji are dropped when colors are disabled so log
// aggregators that choke on unicode get clean text.
func (r *ciRenderer) startGlyph() string {
	if r.useEmoji {
		return r.st.Info("▶ start")
	}
	return "START"
}

func (r *ciRenderer) passGlyph() string {
	if r.useEmoji {
		return r.st.OK("✅ pass ")
	}
	return r.st.OK("PASS ")
}

func (r *ciRenderer) failGlyph() string {
	if r.useEmoji {
		return r.st.Fail("❌ fail ")
	}
	return r.st.Fail("FAIL ")
}

// handle routes a progress event into either a start/finish append line or
// internal per-eval state (for the summary table). Expected to be invoked
// from Display.HandleEvent under its mutex.
func (r *ciRenderer) handle(evt ProgressEvent) {
	switch evt.Type {
	case EventStarting:
		r.onStart(evt)
	case EventToolsVerified:
		r.onToolsVerified(evt)
	case EventGraderStart:
		if s := r.evals[evt.EvalID]; s != nil {
			s.graderTotal++
		}
	case EventGraderComplete:
		if s := r.evals[evt.EvalID]; s != nil {
			if evt.Result == GraderResultPass {
				s.graderPass++
			}
		}
	case EventPassed:
		r.onFinish(evt, "pass")
	case EventFailed:
		r.onFinish(evt, "fail")
	case EventError:
		r.onFinish(evt, "error")
	}
}

func (r *ciRenderer) onStart(evt ProgressEvent) {
	if _, ok := r.evals[evt.EvalID]; ok {
		return
	}
	s := &ciEvalState{
		promptID:   evt.PromptID,
		configName: evt.ConfigName,
		startTime:  time.Now(),
	}
	r.evals[evt.EvalID] = s
	r.order = append(r.order, evt.EvalID)

	fmt.Fprintf(r.w, "%s %s  %s  %s  %s\n",
		r.timestamp(s.startTime),
		r.startGlyph(),
		evt.PromptID,
		r.st.Muted("|"),
		evt.ConfigName,
	)
}

func (r *ciRenderer) onToolsVerified(evt ProgressEvent) {
	s := r.evals[evt.EvalID]
	if s == nil || s.toolsPrinted || len(evt.Tools) == 0 {
		return
	}
	s.toolsPrinted = true
	s.tools = evt.Tools
	
	// Print Tools: header
	fmt.Fprintln(r.w, "Tools:")
	
	// Group and render tools
	grouped := r.groupToolsCI(evt.Tools)
	for _, line := range grouped {
		fmt.Fprintln(r.w, line)
	}
	fmt.Fprintln(r.w) // blank line after tools section
}

// groupToolsCI groups tools by parent for CI output (plain text, no ANSI)
func (r *ciRenderer) groupToolsCI(tools []ToolStatus) []string {
	type parentKey struct {
		name string
		kind string
	}
	
	var parentOrder []parentKey
	parentSeen := make(map[parentKey]bool)
	parentChildren := make(map[parentKey][]ToolStatus)
	var topLevel []ToolStatus
	
	for _, t := range tools {
		if t.ParentKind == "" {
			topLevel = append(topLevel, t)
		} else {
			pk := parentKey{name: t.ParentName, kind: t.ParentKind}
			if !parentSeen[pk] {
				parentOrder = append(parentOrder, pk)
				parentSeen[pk] = true
			}
			parentChildren[pk] = append(parentChildren[pk], t)
		}
	}
	
	var result []string
	
	// Top-level tools
	for _, t := range topLevel {
		result = append(result, r.renderToolLineCI(t, false))
	}
	
	// Parent groups
	for _, pk := range parentOrder {
		children := parentChildren[pk]
		if len(children) == 0 {
			continue
		}
		
		kindLabel := "plugin"
		if pk.kind == ToolParentKindSkillDir {
			kindLabel = "skills dir"
		}
		result = append(result, fmt.Sprintf("  - %s (%s):", pk.name, kindLabel))
		
		for _, child := range children {
			result = append(result, r.renderToolLineCI(child, true))
		}
	}
	
	return result
}

// renderToolLineCI renders a tool status line for CI output (no ANSI)
func (r *ciRenderer) renderToolLineCI(t ToolStatus, indented bool) string {
	indent := "  "
	if indented {
		indent = "      "
	}
	
	statusLabel := "Loaded"
	if t.Status == ToolStatusFailed {
		statusLabel = "Failed"
		if t.Reason != "" {
			statusLabel += " (" + t.Reason + ")"
		}
	}
	
	return fmt.Sprintf("%s- %s: %s", indent, t.ToolName, statusLabel)
}


func (r *ciRenderer) onFinish(evt ProgressEvent, result string) {
	s := r.evals[evt.EvalID]
	if s == nil {
		// Finish without a Start — fabricate a minimal row so we don't drop it.
		s = &ciEvalState{
			promptID:   evt.PromptID,
			configName: evt.ConfigName,
			startTime:  r.startTime,
		}
		r.evals[evt.EvalID] = s
		r.order = append(r.order, evt.EvalID)
	}
	if s.finished {
		return
	}
	s.finished = true
	now := time.Now()
	s.duration = now.Sub(s.startTime)
	s.result = result
	if result != "pass" {
		s.message = oneLine(evt.Message)
		if s.message == "" {
			s.message = defaultFailReason(result)
		}
	}

	glyph := r.passGlyph()
	if result != "pass" {
		glyph = r.failGlyph()
	}

	// Suffix: "(duration, G/T graders) — reason" (reason only on failure).
	suffix := fmt.Sprintf("(%s, %d/%d graders)", fmtCIDuration(s.duration), s.graderPass, s.graderTotal)
	if result != "pass" {
		suffix += " — " + s.message
	}

	fmt.Fprintf(r.w, "%s %s %s  %s  %s  %s\n",
		r.timestamp(now),
		glyph,
		s.promptID,
		r.st.Muted("|"),
		s.configName,
		r.st.Muted(suffix),
	)
}

func defaultFailReason(result string) string {
	if result == "error" {
		return "eval errored"
	}
	return "graders failed"
}

// finish renders the end-of-run summary table and the footer line.
func (r *ciRenderer) finish() {
	fmt.Fprintln(r.w)
	fmt.Fprintln(r.w, r.st.Bold("Summary"))

	rows := r.summaryRows()
	if len(rows) > 0 {
		writeTable(r.w, []string{"Prompt", "Config", "Result", "Graders", "Duration"}, rows, r.st)
	}

	passed := 0
	total := len(r.order)
	for _, id := range r.order {
		if s := r.evals[id]; s != nil && s.result == "pass" {
			passed++
		}
	}

	fmt.Fprintln(r.w)
	footer := fmt.Sprintf("%d/%d passed", passed, total)
	if r.reportDir != "" {
		footer += " · report: " + r.reportDir
	}
	fmt.Fprintln(r.w, footer)
}

// summaryRows produces summary-table rows in first-seen eval order (matching
// the order in which Start lines appeared, so the table reads chronologically).
func (r *ciRenderer) summaryRows() [][]string {
	rows := make([][]string, 0, len(r.order))
	ids := append([]string(nil), r.order...)
	// Stable sort: keep first-seen order; preserved by append above. Only
	// formal call to sort to make intent explicit.
	sort.SliceStable(ids, func(i, j int) bool { return false })

	for _, id := range ids {
		s := r.evals[id]
		if s == nil {
			continue
		}
		resultCell := "PASS"
		if s.result != "pass" {
			resultCell = "FAIL"
		}
		graders := fmt.Sprintf("%d/%d", s.graderPass, s.graderTotal)
		dur := fmtCIDuration(s.duration)
		rows = append(rows, []string{s.promptID, s.configName, resultCell, graders, dur})
	}
	return rows
}

// --- Helpers ---

// fmtCIDuration renders a duration as "Ns" below a minute and "MmSSs" above.
// Kept separate from fmtDuration in display.go so the CI output stays stable
// if the older renderer's formatter shifts.
func fmtCIDuration(d time.Duration) string {
	if d <= 0 {
		return "0s"
	}
	secs := int(d.Seconds() + 0.5)
	if secs < 60 {
		return fmt.Sprintf("%ds", secs)
	}
	return fmt.Sprintf("%dm%02ds", secs/60, secs%60)
}

// oneLine collapses any whitespace run in msg to a single space and trims.
// Failure reasons render on a single line even if the underlying message
// carried newlines or tabs (common with wrapped error strings).
func oneLine(msg string) string {
	var buf bytes.Buffer
	inSpace := false
	for _, r := range msg {
		if r == '\n' || r == '\r' || r == '\t' || r == ' ' {
			if !inSpace && buf.Len() > 0 {
				buf.WriteByte(' ')
				inSpace = true
			}
			continue
		}
		inSpace = false
		buf.WriteRune(r)
	}
	out := buf.String()
	// Trim trailing space produced by a message that ended in whitespace.
	for len(out) > 0 && out[len(out)-1] == ' ' {
		out = out[:len(out)-1]
	}
	return out
}

// --- Table rendering ---

// writeTable renders an ASCII-art table using unicode box-drawing characters.
// Columns auto-size to the max width of header / cell contents. Styler is
// used for the header row only — cell coloring is the caller's job (or left
// plain for snapshot-friendly output).
func writeTable(w io.Writer, headers []string, rows [][]string, st *style.Styler) {
	if len(headers) == 0 {
		return
	}
	widths := make([]int, len(headers))
	for i, h := range headers {
		widths[i] = len(h)
	}
	for _, row := range rows {
		for i, cell := range row {
			if i >= len(widths) {
				break
			}
			if w := len(cell); w > widths[i] {
				widths[i] = w
			}
		}
	}

	writeBorder(w, widths, '┌', '┬', '┐')
	writeRow(w, headers, widths, st, true)
	writeBorder(w, widths, '├', '┼', '┤')
	for _, row := range rows {
		writeRow(w, row, widths, st, false)
	}
	writeBorder(w, widths, '└', '┴', '┘')
}

func writeBorder(w io.Writer, widths []int, left, mid, right rune) {
	var buf bytes.Buffer
	buf.WriteRune(left)
	for i, width := range widths {
		for j := 0; j < width+2; j++ {
			buf.WriteRune('─')
		}
		if i < len(widths)-1 {
			buf.WriteRune(mid)
		} else {
			buf.WriteRune(right)
		}
	}
	buf.WriteByte('\n')
	w.Write(buf.Bytes())
}

func writeRow(w io.Writer, cells []string, widths []int, st *style.Styler, header bool) {
	var buf bytes.Buffer
	buf.WriteRune('│')
	for i, width := range widths {
		cell := ""
		if i < len(cells) {
			cell = cells[i]
		}
		padded := " " + cell + repeat(' ', width-len(cell)) + " "
		if header {
			padded = " " + st.Bold(cell) + repeat(' ', width-len(cell)) + " "
		}
		buf.WriteString(padded)
		buf.WriteRune('│')
	}
	buf.WriteByte('\n')
	w.Write(buf.Bytes())
}

func repeat(r rune, n int) string {
	if n <= 0 {
		return ""
	}
	out := make([]rune, n)
	for i := range out {
		out[i] = r
	}
	return string(out)
}
