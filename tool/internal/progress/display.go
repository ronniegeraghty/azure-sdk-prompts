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
	Total    int       // Total evaluations
	Workers  int       // Parallel workers (= display slots)
	Writer   io.Writer // Output writer (default: os.Stdout)
	Disabled bool      // Force disabled (debug, dry-run, piped output)
}

type slot struct {
	evalID     string
	promptID   string
	configName string
	icon       string
	activity   string
	startTime  time.Time
	active     bool // currently running
	completed  bool // shows result until next eval claims slot
}

// Display renders multi-line per-eval progress with live status updates.
// Each worker gets a dedicated line that updates in-place using ANSI escapes.
type Display struct {
	slots     []slot
	total     int
	completed int
	passed    int
	failed    int
	errors    int
	mu        sync.Mutex
	w         io.Writer
	rendered  bool
	disabled  bool
	evalSlots map[string]int // evalID → slot index
	width     int
	startTime time.Time
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
	}

	if !d.disabled {
		fmt.Fprintf(d.w, "\n%sRunning %d evaluations (%d workers)%s\n\n",
			ColorBold, cfg.Total, cfg.Workers, ColorReset)
		d.redraw()
	}

	return d
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
				icon:       "⏳",
				activity:   "Waiting for session...",
				startTime:  time.Now(),
				active:     true,
			}
			d.evalSlots[evt.EvalID] = idx
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
		if idx, ok := d.evalSlots[evt.EvalID]; ok {
			d.slots[idx].icon = "✓"
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
			d.slots[idx].icon = "✅"
			d.slots[idx].activity = fmt.Sprintf("PASSED  %d files  %s", evt.FileCount, fmtDuration(elapsed))
			d.slots[idx].active = false
			d.slots[idx].completed = true
			delete(d.evalSlots, evt.EvalID)
		}

	case EventFailed:
		d.completed++
		d.failed++
		if idx, ok := d.evalSlots[evt.EvalID]; ok {
			elapsed := time.Since(d.slots[idx].startTime)
			d.slots[idx].icon = "❌"
			d.slots[idx].activity = fmt.Sprintf("FAILED  %s", fmtDuration(elapsed))
			d.slots[idx].active = false
			d.slots[idx].completed = true
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
			d.slots[idx].icon = "❌"
			d.slots[idx].activity = fmt.Sprintf("%s  %s", msg, fmtDuration(elapsed))
			d.slots[idx].active = false
			d.slots[idx].completed = true
			delete(d.evalSlots, evt.EvalID)
		}
	}

	d.redraw()
}

// Done finalizes the display with a summary line.
func (d *Display) Done() {
	if d.disabled {
		return
	}

	d.mu.Lock()
	defer d.mu.Unlock()

	// Clear the display region
	if d.rendered {
		lines := len(d.slots) + 2
		fmt.Fprintf(d.w, "\033[%dA", lines)
		for i := 0; i < lines; i++ {
			fmt.Fprintf(d.w, "\033[2K\n")
		}
		fmt.Fprintf(d.w, "\033[%dA", lines)
	}

	elapsed := time.Since(d.startTime)
	fmt.Fprintf(d.w, "\n%s━━━ Complete: %d/%d%s", ColorBold, d.completed, d.total, ColorReset)
	fmt.Fprintf(d.w, "  %s%d passed%s", ColorGreen, d.passed, ColorReset)
	if d.failed > 0 {
		fmt.Fprintf(d.w, "  %s%d failed%s", ColorRed, d.failed, ColorReset)
	}
	if d.errors > 0 {
		fmt.Fprintf(d.w, "  %s%d errors%s", ColorRed, d.errors, ColorReset)
	}
	fmt.Fprintf(d.w, "  %s\n", fmtDuration(elapsed))
}

// claimSlot finds a free slot for a new eval.
func (d *Display) claimSlot() int {
	// Prefer empty (never-used) slots
	for i, s := range d.slots {
		if !s.active && !s.completed && s.evalID == "" {
			return i
		}
	}
	// Then completed slots (replace finished result)
	for i, s := range d.slots {
		if !s.active && s.completed {
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
		if s.active || s.completed {
			name := d.formatName(s.promptID, s.configName)
			activity := d.truncateActivity(s.activity)
			if s.active {
				elapsed := fmtDuration(time.Since(s.startTime))
				fmt.Fprintf(d.w, "  %-40s %s %-*s %6s", name, s.icon, actW, activity, elapsed)
			} else {
				fmt.Fprintf(d.w, "  %-40s %s %s", name, s.icon, activity)
			}
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
