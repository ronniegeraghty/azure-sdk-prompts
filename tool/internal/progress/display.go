package progress

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
	"time"
)

// DisplayConfig configures the multi-line progress display.
type DisplayConfig struct {
	Total     int       // Total evaluations
	Workers   int       // Parallel workers (= display slots)
	Writer    io.Writer // Output writer (default: os.Stdout)
	Disabled  bool      // Force disabled (debug, dry-run, piped output)
	ReportDir string    // Report directory path (shown in final output)
}

type slot struct {
	evalID     string
	promptID   string
	configName string
	phase      Phase  // Current phase (generating/verifying/reviewing)
	icon       string // Activity icon within the phase
	activity   string // Activity description within the phase
	startTime  time.Time
	active     bool // currently running
}

// completedEval stores the final result of a finished eval.
type completedEval struct {
	promptID    string
	configName  string
	passed      bool
	errored     bool
	fileCount   int
	reviewScore int
	message     string // failure/error message
	duration    time.Duration
}

// Display renders multi-line per-eval progress with live status updates.
// Each worker gets a dedicated line that updates in-place using ANSI escapes.
// Completed evals are stored separately so all results persist in final output.
type Display struct {
	slots          []slot
	completedEvals []completedEval // ALL finished evals (Issue 2 & 3)
	total          int
	completed      int
	passed         int
	failed         int
	errors         int
	mu             sync.Mutex
	w              io.Writer
	rendered       bool
	disabled       bool
	evalSlots      map[string]int // evalID → slot index
	width          int
	startTime      time.Time
	reportDir      string
}

// NewDisplay creates a multi-line progress display.
// Automatically disabled when stdout is not a terminal.
func NewDisplay(cfg DisplayConfig) *Display {
	w := cfg.Writer
	if w == nil {
		w = os.Stdout
	}

	disabled := cfg.Disabled
	if !disabled {
		disabled = !IsTerminal(os.Stdout)
	}

	d := &Display{
		slots:     make([]slot, cfg.Workers),
		total:     cfg.Total,
		w:         w,
		disabled:  disabled,
		evalSlots: make(map[string]int),
		width:     TermWidth(),
		startTime: time.Now(),
		reportDir: cfg.ReportDir,
	}

	if !d.disabled {
		fmt.Fprintf(d.w, "\n%sRunning %d evaluations (%d workers)%s\n\n",
			ColorBold, cfg.Total, cfg.Workers, ColorReset)
		d.redraw()
	}

	return d
}

// phaseIcon returns the primary status icon for the current eval phase.
func phaseIcon(p Phase) string {
	switch p {
	case PhaseGenerating:
		return "🔄"
	case PhaseVerifying:
		return "🔍"
	case PhaseReviewing:
		return "📝"
	default:
		return "⏳"
	}
}

// phaseLabel returns the short label for the current eval phase.
func phaseLabel(p Phase) string {
	switch p {
	case PhaseGenerating:
		return "Generating"
	case PhaseVerifying:
		return "Verifying"
	case PhaseReviewing:
		return "Reviewing"
	default:
		return "Starting"
	}
}

// HandleEvent processes a progress event and redraws the display.
func (d *Display) HandleEvent(evt ProgressEvent) {
	if d.disabled {
		return
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	switch evt.Type {
	case EventStarting:
		idx := d.claimSlot()
		if idx >= 0 {
			d.slots[idx] = slot{
				evalID:     evt.EvalID,
				promptID:   evt.PromptID,
				configName: evt.ConfigName,
				phase:      PhaseGenerating,
				icon:       "⏳",
				activity:   "Waiting for session...",
				startTime:  time.Now(),
				active:     true,
			}
			d.evalSlots[evt.EvalID] = idx
		}

	case EventPhaseChange:
		if idx, ok := d.evalSlots[evt.EvalID]; ok {
			d.slots[idx].phase = evt.Phase
			d.slots[idx].icon = phaseIcon(evt.Phase)
			if evt.Message != "" {
				d.slots[idx].activity = evt.Message
			} else {
				d.slots[idx].activity = phaseLabel(evt.Phase) + "..."
			}
		}

	case EventSendingPrompt:
		if idx, ok := d.evalSlots[evt.EvalID]; ok {
			d.slots[idx].icon = "→"
			d.slots[idx].activity = evt.Message
		}

	case EventReasoning:
		if idx, ok := d.evalSlots[evt.EvalID]; ok {
			d.slots[idx].icon = "💭"
			d.slots[idx].activity = evt.Message
		}

	case EventToolStart:
		if idx, ok := d.evalSlots[evt.EvalID]; ok {
			d.slots[idx].icon = "⚙"
			d.slots[idx].activity = evt.Message
		}

	case EventToolComplete:
		// Issue 1: Don't persist ✓ — revert to the current phase icon
		if idx, ok := d.evalSlots[evt.EvalID]; ok {
			d.slots[idx].icon = phaseIcon(d.slots[idx].phase)
			d.slots[idx].activity = evt.Message
		}

	case EventWritingFile:
		if idx, ok := d.evalSlots[evt.EvalID]; ok {
			d.slots[idx].icon = "📝"
			d.slots[idx].activity = evt.Message
		}

	case EventWaiting:
		if idx, ok := d.evalSlots[evt.EvalID]; ok {
			d.slots[idx].icon = "⏳"
			d.slots[idx].activity = evt.Message
		}

	case EventPassed:
		d.completed++
		d.passed++
		if idx, ok := d.evalSlots[evt.EvalID]; ok {
			elapsed := time.Since(d.slots[idx].startTime)
			d.completedEvals = append(d.completedEvals, completedEval{
				promptID:    d.slots[idx].promptID,
				configName:  d.slots[idx].configName,
				passed:      true,
				fileCount:   evt.FileCount,
				reviewScore: evt.ReviewScore,
				duration:    elapsed,
			})
			d.slots[idx] = slot{} // release slot
			delete(d.evalSlots, evt.EvalID)
		}

	case EventFailed:
		d.completed++
		d.failed++
		if idx, ok := d.evalSlots[evt.EvalID]; ok {
			elapsed := time.Since(d.slots[idx].startTime)
			msg := "verification failed"
			if evt.Message != "" {
				msg = evt.Message
			}
			d.completedEvals = append(d.completedEvals, completedEval{
				promptID:   d.slots[idx].promptID,
				configName: d.slots[idx].configName,
				passed:     false,
				message:    msg,
				duration:   elapsed,
			})
			d.slots[idx] = slot{} // release slot
			delete(d.evalSlots, evt.EvalID)
		}

	case EventError:
		d.completed++
		d.errors++
		if idx, ok := d.evalSlots[evt.EvalID]; ok {
			elapsed := time.Since(d.slots[idx].startTime)
			msg := "ERROR"
			if evt.Message != "" {
				msg = evt.Message
			}
			d.completedEvals = append(d.completedEvals, completedEval{
				promptID:   d.slots[idx].promptID,
				configName: d.slots[idx].configName,
				errored:    true,
				message:    msg,
				duration:   elapsed,
			})
			d.slots[idx] = slot{} // release slot
			delete(d.evalSlots, evt.EvalID)
		}
	}

	d.redraw()
}

// Finish stops the ANSI refresh loop and prints all final results as static output.
// This replaces Done() for the final display — all completed eval lines persist.
func (d *Display) Finish() {
	if d.disabled {
		return
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	// Clear the live display region
	if d.rendered {
		lines := len(d.slots) + 2
		fmt.Fprintf(d.w, "\033[%dA", lines)
		for i := 0; i < lines; i++ {
			fmt.Fprintf(d.w, "\033[2K\n")
		}
		fmt.Fprintf(d.w, "\033[%dA", lines)
	}

	// Print all completed evals as static lines (no ANSI cursor movement)
	for _, ce := range d.completedEvals {
		name := d.formatName(ce.promptID, ce.configName)
		if ce.errored {
			fmt.Fprintf(d.w, "  %-40s ❌ %s  %s\n", name, ce.message, fmtDuration(ce.duration))
		} else if !ce.passed {
			fmt.Fprintf(d.w, "  %-40s ❌ FAILED  %s  %s\n", name, ce.message, fmtDuration(ce.duration))
		} else {
			score := ""
			if ce.reviewScore > 0 {
				score = fmt.Sprintf("  %d/10", ce.reviewScore)
			}
			fmt.Fprintf(d.w, "  %-40s ✅ PASSED  %d files%s  %s\n", name, ce.fileCount, score, fmtDuration(ce.duration))
		}
	}

	// Summary line
	elapsed := time.Since(d.startTime)
	fmt.Fprintf(d.w, "\n%sSummary: %d/%d passed%s", ColorBold, d.passed, d.total, ColorReset)
	fmt.Fprintf(d.w, "  %s✅ %d%s", ColorGreen, d.passed, ColorReset)
	if d.failed > 0 {
		fmt.Fprintf(d.w, "  %s❌ %d%s", ColorRed, d.failed, ColorReset)
	}
	if d.errors > 0 {
		fmt.Fprintf(d.w, "  %s❌ %d errors%s", ColorRed, d.errors, ColorReset)
	}
	fmt.Fprintf(d.w, "  Duration: %s\n", fmtDuration(elapsed))

	if d.reportDir != "" {
		fmt.Fprintf(d.w, "Reports: %s\n", d.reportDir)
	}
}

// Done finalizes the display with a summary line (backward compat — prefer Finish).
func (d *Display) Done() {
	d.Finish()
}

// CompletedEvalCount returns the number of completed evals (for testing).
func (d *Display) CompletedEvalCount() int {
	d.mu.Lock()
	defer d.mu.Unlock()
	return len(d.completedEvals)
}

// claimSlot finds a free slot for a new eval.
func (d *Display) claimSlot() int {
	// Prefer empty (never-used or released) slots
	for i, s := range d.slots {
		if !s.active && s.evalID == "" {
			return i
		}
	}
	// Then any inactive slot
	for i, s := range d.slots {
		if !s.active {
			return i
		}
	}
	return -1
}

func (d *Display) redraw() {
	if d.rendered {
		lines := len(d.slots) + 2 // slots + blank + summary
		fmt.Fprintf(d.w, "\033[%dA", lines)
	}

	actW := d.activityWidth()
	for i := range d.slots {
		s := &d.slots[i]
		fmt.Fprintf(d.w, "\033[2K") // clear line
		if s.active {
			name := d.formatName(s.promptID, s.configName)
			pLabel := phaseLabel(s.phase)
			pIcon := phaseIcon(s.phase)
			activity := d.truncateActivity(s.activity)
			elapsed := fmtDuration(time.Since(s.startTime))
			fmt.Fprintf(d.w, "  %-40s %s %-12s %s %-*s %6s",
				name, pIcon, pLabel, s.icon, actW-16, activity, elapsed)
		}
		fmt.Fprint(d.w, "\n")
	}

	// Blank line
	fmt.Fprintf(d.w, "\033[2K\n")

	// Summary line
	fmt.Fprintf(d.w, "\033[2KCompleted: %d/%d", d.completed, d.total)
	if d.passed > 0 {
		fmt.Fprintf(d.w, "  %s✅ %d%s", ColorGreen, d.passed, ColorReset)
	}
	if d.failed > 0 {
		fmt.Fprintf(d.w, "  %s❌ %d%s", ColorRed, d.failed, ColorReset)
	}
	if d.errors > 0 {
		fmt.Fprintf(d.w, "  %s❌ %d errors%s", ColorRed, d.errors, ColorReset)
	}
	fmt.Fprint(d.w, "\n")

	d.rendered = true
}

func (d *Display) formatName(promptID, configName string) string {
	name := promptID + "/" + configName
	const maxLen = 38
	if len(name) > maxLen {
		name = name[:maxLen-2] + ".."
	}
	return name
}

func (d *Display) activityWidth() int {
	// Layout: 2 indent + 40 name + 1 space + ~3 icon + 1 space + activity + 2 spaces + 6 elapsed
	w := d.width - 55
	if w < 20 {
		w = 20
	}
	return w
}

func (d *Display) truncateActivity(s string) string {
	maxLen := d.activityWidth()
	if len(s) > maxLen {
		return s[:maxLen-3] + "..."
	}
	return s
}

// TermWidth returns the terminal width from COLUMNS env var, defaulting to 120.
func TermWidth() int {
	if cols := os.Getenv("COLUMNS"); cols != "" {
		if n, err := strconv.Atoi(cols); err == nil && n > 0 {
			return n
		}
	}
	return 120
}

// IsTerminal reports whether f is connected to a terminal device.
func IsTerminal(f *os.File) bool {
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}
