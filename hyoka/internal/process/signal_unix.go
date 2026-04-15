//go:build !windows

package process

import (
	"os"
	"os/signal"
	"syscall"
)

// NotifyShutdownSignals registers ch to receive SIGINT and SIGTERM.
func NotifyShutdownSignals(ch chan<- os.Signal) {
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
}
