//go:build windows

package process

import (
	"os"
	"os/signal"
)

// NotifyShutdownSignals registers ch to receive os.Interrupt (Ctrl+C).
// Windows does not have SIGTERM; the closest equivalent is os.Interrupt
// which is delivered for both Ctrl+C and Ctrl+Break.
func NotifyShutdownSignals(ch chan<- os.Signal) {
	signal.Notify(ch, os.Interrupt)
}
