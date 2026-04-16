// Package progress provides progress reporting and display for the evaluation tool.
package progress

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// ProgressMode controls the rendering strategy.
type ProgressMode string

const (
	ModeAuto ProgressMode = "auto" // ANSI if TTY, log otherwise
	ModeLive ProgressMode = "live" // Force ANSI (cursor save/restore)
	ModeLog  ProgressMode = "log"  // Append-only phase lines (no cursor movement)
	ModeOff  ProgressMode = "off"  // No progress output
)

// DisplayConfig controls the progress display.
type DisplayConfig struct {
	Total     int
	Workers   int
	Writer    io.Writer
	Disabled  bool
	ReportDir string
	Mode      ProgressMode // "" or "auto" uses auto-detection
}

type evalStatus int

const (
	evalActive evalStatus = iota
	evalPassed
	evalFailed
	evalError
)

// evalSection holds the state for one eval in the section-based display.
// Each eval renders as a multi-line section showing prompt, config, and
// per-phase status lines.
type evalSection struct {
	promptID   string
	configName string
	status     evalStatus
	startTime  time.Time
	phase      Phase
	activity   string

	// Per-phase tracking
	genStartTime time.Time
	genDone      bool
	genDuration  time.Duration
	genActivity  string

	revStartTime time.Time
	revDone      bool
	revDuration  time.Duration
	revActivity  string

	fileCount   int
	reviewScore int
	message     string
	duration    time.Duration
}

// Display renders section-based progress for evaluation runs.
//
// Each eval gets its own multi-line section:
//
//	Prompt: key-vault-dp-python-crud
//	Config: baseline/claude-opus-4.6
//	  ├─ Agent:  ✅ Complete (44.3s)
//	  └─ Review: 🔄 Running... (12.5s)
//
// In ANSI mode (real terminal), the entire region is redrawn on a 500ms
// timer using DECSC/DECRC cursor save/restore.
// In non-ANSI mode (piped output, test buffers), sections print inline
// as evals progress and complete.
type Display struct {
	total     int
	workers   int
	completed int
	passed    int
	failed    int
	errors    int
	mu        sync.Mutex
	w         io.Writer
	disabled  bool
	ansi      bool
	startTime time.Time
	reportDir string

	// Dynamic eval sections — grows as evals start.
	sections     []*evalSection
	sectionIndex map[string]int

	// ANSI redraw timer
	ticker *time.Ticker
	stopCh chan struct{}
	wg     sync.WaitGroup
}

// NewDisplay creates a progress display. When Writer is nil, it writes to
// os.Stdout and enables ANSI rendering if stdout is a terminal.
// Mode overrides auto-detection: "live" forces ANSI, "log" forces append-only,
// "off" disables output entirely.
func NewDisplay(cfg DisplayConfig) *Display {
	w := cfg.Writer
	if w == nil {
		w = os.Stdout
	}

	disabled := cfg.Disabled
	ansi := false

	switch cfg.Mode {
	case ModeOff:
		disabled = true
	case ModeLive:
		if !disabled {
			ansi = true
		}
	case ModeLog:
		ansi = false
	default: // ModeAuto or ""
		if !disabled && cfg.Writer == nil {
			if IsTerminal(os.Stdout) {
				ansi = true
			} else {
				disabled = true
			}
		}
	}

	d := &Display{
		total:        cfg.Total,
		workers:      cfg.Workers,
		w:            w,
		disabled:     disabled,
		ansi:         ansi,
		startTime:    time.Now(),
		reportDir:    cfg.ReportDir,
		sections:     []*evalSection{},
		sectionIndex: make(map[string]int),
	}

	if d.ansi && cfg.Total > 0 {
		fmt.Fprintf(d.w, "\nRunning %d evaluations (%d workers)\n", cfg.Total, cfg.Workers)
		fmt.Fprint(d.w, "\0337") // DECSC: save cursor position
		d.drawRegion()
		d.stopCh = make(chan struct{})
		d.ticker = time.NewTicker(500 * time.Millisecond)
		d.wg.Add(1)
		go d.redrawLoop()
	} else if !d.disabled {
		fmt.Fprintf(d.w, "\nRunning %d evaluations (%d workers)\n\n", cfg.Total, cfg.Workers)
	}

	return d
}

// --- ANSI fixed-region rendering (terminal only) ---

// buildRegion renders all eval sections plus a summary line into a buffer.
func (d *Display) buildRegion() []byte {
	var buf bytes.Buffer
	for i, s := range d.sections {
		if i > 0 {
			buf.WriteByte('\n')
		}
		d.writeSection(&buf, s)
	}
	buf.WriteByte('\n')
	d.writeSummaryLine(&buf)
	return buf.Bytes()
}

func (d *Display) drawRegion() {
	d.w.Write(d.buildRegion())
}

// redrawRegion restores cursor to saved position, clears below, redraws.
func (d *Display) redrawRegion() {
	region := d.buildRegion()
	var buf bytes.Buffer
	buf.WriteString("\0338\033[J")
	buf.Write(region)
	d.w.Write(buf.Bytes())
}

// writeSection renders a multi-line section for one eval.
func (d *Display) writeSection(buf *bytes.Buffer, s *evalSection) {
	fmt.Fprintf(buf, "Prompt: %s\n", s.promptID)
	fmt.Fprintf(buf, "Config: %s\n", s.configName)

	switch s.status {
	case evalActive:
		d.writeActivePhases(buf, s)
	case evalPassed:
		score := ""
		if s.reviewScore > 0 {
			score = fmt.Sprintf("  %d/10", s.reviewScore)
		}
		fmt.Fprintf(buf, "  ✅ Passed  %d files%s  (%s)\n", s.fileCount, score, fmtDuration(s.duration))
	case evalFailed, evalError:
		msg := s.message
		if msg == "" {
			msg = "failed"
		}
		fmt.Fprintf(buf, "  ❌ %s  (%s)\n", msg, fmtDuration(s.duration))
	}
}

// writeActivePhases renders the per-phase status lines for an in-progress eval.
func (d *Display) writeActivePhases(buf *bytes.Buffer, s *evalSection) {
	hasReview := s.phase == PhaseReviewing || s.revDone

	agentConnector := "└─"
	if hasReview {
		agentConnector = "├─"
	}

	// Agent (generation) phase
	if s.genDone {
		fmt.Fprintf(buf, "  %s Agent:  ✅ Complete (%s)\n", agentConnector, fmtDuration(s.genDuration))
	} else if s.phase == PhaseGenerating || s.phase == "" {
		activity := s.genActivity
		if activity == "" {
			activity = s.activity
		}
		if activity == "" {
			activity = "Starting..."
		}
		dur := fmtDuration(time.Since(s.genStartTime))
		if s.genStartTime.IsZero() {
			dur = fmtDuration(time.Since(s.startTime))
		}
		fmt.Fprintf(buf, "  %s Agent:  🔄 %s (%s)\n", agentConnector, activity, dur)
	}

	// Review phase
	if hasReview {
		if s.revDone {
			fmt.Fprintf(buf, "  └─ Review: ✅ Complete (%s)\n", fmtDuration(s.revDuration))
		} else {
			activity := s.revActivity
			if activity == "" {
				activity = s.activity
			}
			if activity == "" {
				activity = "Running..."
			}
			dur := fmtDuration(time.Since(s.revStartTime))
			if s.revStartTime.IsZero() {
				dur = fmtDuration(time.Since(s.startTime))
			}
			fmt.Fprintf(buf, "  └─ Review: 🔄 %s (%s)\n", activity, dur)
		}
	}
}

func (d *Display) writeSummaryLine(buf *bytes.Buffer) {
	if d.completed == d.total && d.total > 0 {
		fmt.Fprintf(buf, "Summary: %d/%d passed", d.passed, d.total)
	} else {
		fmt.Fprintf(buf, "%d/%d completed", d.completed, d.total)
	}
	if d.passed > 0 {
		fmt.Fprintf(buf, "  ✅ %d", d.passed)
	}
	if d.failed > 0 {
		fmt.Fprintf(buf, "  ❌ %d", d.failed)
	}
	if d.errors > 0 {
		fmt.Fprintf(buf, "  ❌ %d errors", d.errors)
	}
	fmt.Fprintf(buf, "  %s\n", fmtDuration(time.Since(d.startTime)))
}

func (d *Display) redrawLoop() {
	defer d.wg.Done()
	for {
		select {
		case <-d.stopCh:
			return
		case <-d.ticker.C:
			d.mu.Lock()
			d.redrawRegion()
			d.mu.Unlock()
		}
	}
}

// --- Slot assignment ---

func (d *Display) getOrAssignSlot(evalID, promptID, configName string) int {
	if idx, ok := d.sectionIndex[evalID]; ok {
		return idx
	}
	idx := len(d.sections)
	d.sections = append(d.sections, &evalSection{
		promptID:   promptID,
		configName: configName,
	})
	d.sectionIndex[evalID] = idx
	return idx
}

// --- Event handling ---

// HandleEvent updates internal state from engine/evaluator events.
// In ANSI mode, rendering happens on the timer — not here.
// In non-ANSI mode, section-based output prints inline.
func (d *Display) HandleEvent(evt ProgressEvent) {
	if d.disabled {
		return
	}
	d.mu.Lock()
	defer d.mu.Unlock()

	switch evt.Type {
	case EventStarting:
		idx := d.getOrAssignSlot(evt.EvalID, evt.PromptID, evt.ConfigName)
		s := d.sections[idx]
		s.status = evalActive
		s.startTime = time.Now()
		s.activity = evt.Message
		if !d.ansi {
			fmt.Fprintf(d.w, "Prompt: %s\nConfig: %s\n  └─ Agent:  🔄 Starting...\n\n", s.promptID, s.configName)
		}

	case EventSendingPrompt, EventReasoning, EventToolStart, EventToolComplete,
		EventWritingFile, EventWaiting:
		if idx, ok := d.sectionIndex[evt.EvalID]; ok {
			s := d.sections[idx]
			s.activity = evt.Message
			if s.phase == PhaseReviewing {
				s.revActivity = evt.Message
			} else {
				s.genActivity = evt.Message
			}
			if !d.ansi && evt.Message != "" {
				prefix := "  "
				switch evt.Type {
				case EventToolStart:
					prefix = "    🔧"
				case EventToolComplete:
					prefix = "    ✓ "
				case EventWritingFile:
					prefix = "    📄"
				case EventSendingPrompt:
					prefix = "    📨"
				case EventReasoning:
					prefix = "    💭"
				default:
					prefix = "    ⏳"
				}
				fmt.Fprintf(d.w, "%s [%s/%s] %s\n", prefix, s.promptID, s.configName, evt.Message)
			}
		}

	case EventPhaseChange:
		if idx, ok := d.sectionIndex[evt.EvalID]; ok {
			s := d.sections[idx]
			// Mark previous phase as done
			if evt.Phase == PhaseReviewing && !s.genDone {
				s.genDone = true
				if !s.genStartTime.IsZero() {
					s.genDuration = time.Since(s.genStartTime)
				} else {
					s.genDuration = time.Since(s.startTime)
				}
			}
			s.phase = evt.Phase
			s.activity = string(evt.Phase)
			if evt.Phase == PhaseGenerating {
				s.genStartTime = time.Now()
			} else if evt.Phase == PhaseReviewing {
				s.revStartTime = time.Now()
			}
			if !d.ansi {
				fmt.Fprintf(d.w, "  ▶ [%s/%s] %s...\n", s.promptID, s.configName, evt.Phase)
			}
		}

	case EventPassed:
		d.completed++
		d.passed++
		if idx, ok := d.sectionIndex[evt.EvalID]; ok {
			s := d.sections[idx]
			s.status = evalPassed
			s.duration = time.Since(s.startTime)
			s.fileCount = evt.FileCount
			s.reviewScore = evt.ReviewScore
			if s.phase == PhaseReviewing && !s.revDone {
				s.revDone = true
				if !s.revStartTime.IsZero() {
					s.revDuration = time.Since(s.revStartTime)
				}
			}
			if !d.ansi {
				score := ""
				if evt.ReviewScore > 0 {
					score = fmt.Sprintf("  %d/10", evt.ReviewScore)
				}
				fmt.Fprintf(d.w, "  ✅ Passed  %d files%s  (%s)\n\n",
					evt.FileCount, score, fmtDuration(s.duration))
			}
		}

	case EventFailed:
		d.completed++
		d.failed++
		if idx, ok := d.sectionIndex[evt.EvalID]; ok {
			s := d.sections[idx]
			s.status = evalFailed
			s.duration = time.Since(s.startTime)
			s.message = evt.Message
			if s.message == "" {
				s.message = "failed"
			}
			if !d.ansi {
				fmt.Fprintf(d.w, "  ❌ %s  (%s)\n\n",
					s.message, fmtDuration(s.duration))
			}
		}

	case EventError:
		d.completed++
		d.errors++
		if idx, ok := d.sectionIndex[evt.EvalID]; ok {
			s := d.sections[idx]
			s.status = evalError
			s.duration = time.Since(s.startTime)
			s.message = evt.Message
			if s.message == "" {
				s.message = "ERROR"
			}
			if !d.ansi {
				fmt.Fprintf(d.w, "  ❌ %s  (%s)\n\n",
					s.message, fmtDuration(s.duration))
			}
		}
	}
}

// Finish stops the redraw timer, renders final state, and prints the summary.
func (d *Display) Finish() {
	if d.disabled {
		return
	}

	// ANSI mode: stop timer, wait for redrawLoop to exit, then final redraw
	if d.ansi && d.ticker != nil {
		d.ticker.Stop()
		close(d.stopCh)
		d.wg.Wait()
		d.mu.Lock()
		if d.total > 0 {
			d.redrawRegion()
		}
		if d.reportDir != "" {
			fmt.Fprintf(d.w, "\nReports: %s\n", d.reportDir)
		} else {
			fmt.Fprintln(d.w)
		}
		d.mu.Unlock()
		return
	}

	// Non-ANSI mode: print summary below inline results
	d.mu.Lock()
	defer d.mu.Unlock()

	elapsed := time.Since(d.startTime)
	fmt.Fprintf(d.w, "\nSummary: %d/%d passed", d.passed, d.total)
	fmt.Fprintf(d.w, "  ✅ %d", d.passed)
	if d.failed > 0 {
		fmt.Fprintf(d.w, "  ❌ %d", d.failed)
	}
	if d.errors > 0 {
		fmt.Fprintf(d.w, "  ❌ %d errors", d.errors)
	}
	fmt.Fprintf(d.w, "  Duration: %s\n", fmtDuration(elapsed))
	if d.reportDir != "" {
		fmt.Fprintf(d.w, "Reports: %s\n", d.reportDir)
	}
}

// Done is an alias for Finish.
func (d *Display) Done() { d.Finish() }

// CompletedEvalCount returns the number of evals that have completed.
func (d *Display) CompletedEvalCount() int {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.completed
}

// IsTerminal reports whether f is connected to a terminal.
func IsTerminal(f *os.File) bool {
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}

// TermWidth returns the assumed terminal width.
func TermWidth() int { return 120 }

func fmtDuration(d time.Duration) string {
	secs := d.Seconds()
	if secs < 60 {
		return fmt.Sprintf("%.0fs", secs)
	}
	return fmt.Sprintf("%.1fm", secs/60)
}
