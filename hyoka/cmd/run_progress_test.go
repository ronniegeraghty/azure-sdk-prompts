package cmd

import "testing"

// TestResolveAutoProgress covers the --progress auto resolution matrix.
//
// Regression: before the fix, non-TTY stdout forced "off" even for workers>1,
// which suppressed the CI renderer exactly where it's most useful (piped/CI).
// The (workers=4, TTY=false) → "ci" row is the explicit regression guard.
func TestResolveAutoProgress(t *testing.T) {
	tests := []struct {
		name       string
		workers    int
		isTerminal bool
		logLevel   string
		logFile    string
		want       string
	}{
		{"single+tty", 1, true, "warn", "", "interactive"},
		{"single+pipe", 1, false, "warn", "", "off"},
		{"multi+tty", 4, true, "warn", "", "ci"},
		{"multi+pipe regression", 4, false, "warn", "", "ci"},

		// --log-file exception: verbose logging in interactive stays interactive
		// when slog output is routed to a file.
		{"single+tty+debug+logfile", 1, true, "debug", "run.log", "interactive"},
		{"single+tty+debug+no-logfile", 1, true, "debug", "", "ci"},
		{"single+tty+info+no-logfile", 1, true, "info", "", "ci"},
		{"single+tty+warn", 1, true, "warn", "", "interactive"},

		// Multi-worker always picks ci regardless of logging.
		{"multi+tty+debug", 4, true, "debug", "", "ci"},
		{"multi+pipe+debug+logfile", 4, false, "debug", "run.log", "ci"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := resolveAutoProgress(tt.workers, tt.isTerminal, tt.logLevel, tt.logFile)
			if got != tt.want {
				t.Errorf("resolveAutoProgress(workers=%d, tty=%v, logLevel=%q, logFile=%q) = %q, want %q",
					tt.workers, tt.isTerminal, tt.logLevel, tt.logFile, got, tt.want)
			}
		})
	}
}
