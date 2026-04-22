package progress

import (
	"bytes"
	"regexp"
	"strings"
	"testing"
)

// normalizeCI strips wall-clock variability out of the CI renderer's output so
// golden snapshots remain stable. Two fields drift between runs:
//
//   - the "[HH:MM:SS]" timestamp prefix on every start/finish line;
//   - the duration cell inside the "(Ns, G/T graders)" suffix and the
//     Duration column of the summary table.
//
// Both are replaced with stable placeholders. The ordering and total length
// of each line is preserved only incidentally — snapshot consumers should
// compare on semantic content (line sequence, table structure, footer).
var (
	reCITimestamp = regexp.MustCompile(`\[\d{2}:\d{2}:\d{2}\]`)
	reCIDuration  = regexp.MustCompile(`\(\d+m?\d*s?, (\d+)/(\d+) graders\)`)
	reCITableDur  = regexp.MustCompile(`\s\d+(m\d+)?s\s*│`)
)

func normalizeCI(s string) string {
	s = reCITimestamp.ReplaceAllString(s, "[HH:MM:SS]")
	s = reCIDuration.ReplaceAllString(s, "(DUR, $1/$2 graders)")
	s = reCITableDur.ReplaceAllString(s, " DUR │")
	return s
}

// feedCI wires a CI-mode Display to buf and drives the event sequence
// through HandleEvent + Finish. total/workers/configs/reportDir are exposed
// so each case can tune the intro line without repeating DisplayConfig
// plumbing.
func feedCI(buf *bytes.Buffer, total, workers, configs int, reportDir string, events []ProgressEvent) {
	d := NewDisplay(DisplayConfig{
		Total:     total,
		Workers:   workers,
		Configs:   configs,
		Writer:    buf,
		Mode:      ModeCI,
		ReportDir: reportDir,
	})
	for _, e := range events {
		d.HandleEvent(e)
	}
	d.Finish()
}

// TestCIRenderer_Cases drives the CI renderer through the scenarios listed in
// the tests-renderer-snapshots task. Each case pins on the normalized buffer
// contents (timestamps / durations replaced with stable placeholders) plus a
// list of required substrings.
func TestCIRenderer_Cases(t *testing.T) {
	cases := []struct {
		name         string
		total        int
		workers      int
		configs      int
		reportDir    string
		events       []ProgressEvent
		wantSubs     []string
		notWantSubs  []string
		wantPassText string // e.g. "3/3 passed" or "2/3 passed"
	}{
		{
			name:      "happy_path_three_evals_all_pass",
			total:     3,
			workers:   2,
			configs:   1,
			reportDir: "reports/x/",
			events: []ProgressEvent{
				{EvalID: "e1", PromptID: "p1", ConfigName: "c1", Type: EventStarting},
				{EvalID: "e2", PromptID: "p2", ConfigName: "c1", Type: EventStarting},
				{EvalID: "e3", PromptID: "p3", ConfigName: "c1", Type: EventStarting},
				{EvalID: "e1", Type: EventGraderStart, GraderID: "g1"},
				{EvalID: "e1", Type: EventGraderComplete, GraderID: "g1", Result: GraderResultPass},
				{EvalID: "e1", Type: EventPassed, FileCount: 1},
				{EvalID: "e2", Type: EventGraderStart, GraderID: "g1"},
				{EvalID: "e2", Type: EventGraderComplete, GraderID: "g1", Result: GraderResultPass},
				{EvalID: "e2", Type: EventPassed, FileCount: 1},
				{EvalID: "e3", Type: EventGraderStart, GraderID: "g1"},
				{EvalID: "e3", Type: EventGraderComplete, GraderID: "g1", Result: GraderResultPass},
				{EvalID: "e3", Type: EventPassed, FileCount: 1},
			},
			wantSubs: []string{
				"Running 3 evals across 1 configs with 2 workers",
				"START", // buffer-writer → color disabled → text glyph
				"PASS",
				"p1",
				"p2",
				"p3",
				"Summary",
				"┌", "┐", "└", "┘", "│", "├", "┤",
				"reports/x/",
			},
			notWantSubs:  []string{"FAIL"},
			wantPassText: "3/3 passed",
		},
		{
			name:    "mixed_two_pass_one_fail_with_reason",
			total:   3,
			workers: 2,
			configs: 1,
			events: []ProgressEvent{
				{EvalID: "e1", PromptID: "ok1", ConfigName: "c", Type: EventStarting},
				{EvalID: "e2", PromptID: "ok2", ConfigName: "c", Type: EventStarting},
				{EvalID: "e3", PromptID: "bad", ConfigName: "c", Type: EventStarting},
				{EvalID: "e1", Type: EventPassed, FileCount: 1},
				{EvalID: "e2", Type: EventPassed, FileCount: 1},
				{EvalID: "e3", Type: EventGraderStart, GraderID: "g1"},
				{EvalID: "e3", Type: EventGraderComplete, GraderID: "g1", Result: GraderResultFail},
				{EvalID: "e3", Type: EventFailed, Message: "grader\noutput_check\nfailed"},
			},
			wantSubs: []string{
				"FAIL",
				"PASS",
				"bad",
				"grader output_check failed", // oneLine() collapses newlines
				"0/1 graders",
			},
			wantPassText: "2/3 passed",
		},
		{
			name:    "multi_eval_interleaved_graders",
			total:   2,
			workers: 2,
			configs: 1,
			events: []ProgressEvent{
				{EvalID: "e1", PromptID: "pA", ConfigName: "c", Type: EventStarting},
				{EvalID: "e2", PromptID: "pB", ConfigName: "c", Type: EventStarting},
				// Interleave grader events across evals — each eval's counters
				// must be isolated.
				{EvalID: "e1", Type: EventGraderStart, GraderID: "ga"},
				{EvalID: "e2", Type: EventGraderStart, GraderID: "gx"},
				{EvalID: "e1", Type: EventGraderStart, GraderID: "gb"},
				{EvalID: "e2", Type: EventGraderStart, GraderID: "gy"},
				{EvalID: "e2", Type: EventGraderComplete, GraderID: "gx", Result: GraderResultPass},
				{EvalID: "e1", Type: EventGraderComplete, GraderID: "ga", Result: GraderResultPass},
				{EvalID: "e1", Type: EventGraderComplete, GraderID: "gb", Result: GraderResultPass},
				{EvalID: "e2", Type: EventGraderComplete, GraderID: "gy", Result: GraderResultFail},
				{EvalID: "e1", Type: EventPassed, FileCount: 1},
				{EvalID: "e2", Type: EventFailed, Message: "one grader failed"},
			},
			wantSubs: []string{
				// e1: 2 graders started, 2 passed
				"2/2 graders",
				// e2: 2 graders started, 1 passed
				"1/2 graders",
				"one grader failed",
			},
			wantPassText: "1/2 passed",
		},
		{
			name:    "no_color_drops_emoji_keeps_box_borders",
			total:   1,
			workers: 1,
			configs: 1,
			events: []ProgressEvent{
				{EvalID: "e1", PromptID: "p", ConfigName: "c", Type: EventStarting},
				{EvalID: "e1", Type: EventPassed, FileCount: 1},
			},
			// Buffer writer already disables color + emoji. Assert the plain-
			// text path renders "START"/"PASS" (not "▶ start" / "✅ pass") and
			// box-drawing chars survive regardless of color state.
			wantSubs: []string{
				"START",
				"PASS",
				"┌", "│", "┘",
				"Summary",
			},
			notWantSubs:  []string{"▶ start", "✅ pass", "\x1b[31m", "\x1b[32m", "\x1b[36m"},
			wantPassText: "1/1 passed",
		},
		{
			name:         "zero_evals_empty_summary_does_not_crash",
			total:        0,
			workers:      1,
			configs:      0,
			events:       []ProgressEvent{},
			wantSubs:     []string{"Summary"},
			notWantSubs:  []string{"┌", "│"}, // no rows → writeTable skipped
			wantPassText: "0/0 passed",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			feedCI(&buf, tc.total, tc.workers, tc.configs, tc.reportDir, tc.events)
			out := buf.String()
			norm := normalizeCI(out)

			for _, sub := range tc.wantSubs {
				if !strings.Contains(out, sub) {
					t.Errorf("want substring %q not found\n--- raw ---\n%s\n--- normalized ---\n%s", sub, out, norm)
				}
			}
			for _, sub := range tc.notWantSubs {
				if strings.Contains(out, sub) {
					t.Errorf("unexpected substring %q found\n--- raw ---\n%s", sub, out)
				}
			}
			if tc.wantPassText != "" && !strings.Contains(out, tc.wantPassText) {
				t.Errorf("want footer %q; output:\n%s", tc.wantPassText, out)
			}
			// Timestamp normalization sanity: after normalizing, no literal
			// [HH:MM:SS] with real digits should remain.
			if reCITimestamp.MatchString(norm) {
				t.Errorf("normalizeCI failed to strip timestamps:\n%s", norm)
			}
		})
	}
}

// TestCIRenderer_HappyPathSnapshot pins the full normalized output for the
// simplest happy-path case so future renderer tweaks produce a visible diff.
// Larger matrices stick to substring assertions — they're less fragile when
// the layout shifts.
func TestCIRenderer_HappyPathSnapshot(t *testing.T) {
	var buf bytes.Buffer
	feedCI(&buf, 2, 2, 1, "reports/y/", []ProgressEvent{
		{EvalID: "e1", PromptID: "p1", ConfigName: "c1", Type: EventStarting},
		{EvalID: "e2", PromptID: "p2", ConfigName: "c1", Type: EventStarting},
		{EvalID: "e1", Type: EventGraderStart, GraderID: "g"},
		{EvalID: "e1", Type: EventGraderComplete, GraderID: "g", Result: GraderResultPass},
		{EvalID: "e1", Type: EventPassed, FileCount: 1},
		{EvalID: "e2", Type: EventGraderStart, GraderID: "g"},
		{EvalID: "e2", Type: EventGraderComplete, GraderID: "g", Result: GraderResultFail},
		{EvalID: "e2", Type: EventFailed, Message: "bad output"},
	})
	want := `Running 2 evals across 1 configs with 2 workers…

[HH:MM:SS] START  p1  |  c1
[HH:MM:SS] START  p2  |  c1
[HH:MM:SS] PASS  p1  |  c1  (DUR, 1/1 graders)
[HH:MM:SS] FAIL  p2  |  c1  (DUR, 0/1 graders) — bad output

Summary
┌────────┬────────┬────────┬─────────┬──────────┐
│ Prompt │ Config │ Result │ Graders │ Duration │
├────────┼────────┼────────┼─────────┼──────────┤
│ p1     │ c1     │ PASS   │ 1/1     │ DUR │
│ p2     │ c1     │ FAIL   │ 0/1     │ DUR │
└────────┴────────┴────────┴─────────┴──────────┘

1/2 passed · report: reports/y/
`
	got := normalizeCI(buf.String())
	if got != want {
		t.Errorf("CI snapshot mismatch\n--- want ---\n%s\n--- got ---\n%s", want, got)
	}
}
