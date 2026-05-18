//go:build windows

package tool

import "context"

// acquireRepoLock is a no-op on Windows. Concurrent hyoka runs against the
// same repo cache may race; the cost is a possible failed git fetch that the
// caller surfaces as a tool-load error and the user retries. A real
// LockFileEx implementation can land later if Windows usage warrants it.
func acquireRepoLock(_ context.Context, _ string) (func() error, error) {
	return func() error { return nil }, nil
}
