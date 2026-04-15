package eval

import (
	"context"
	"log/slog"

	copilot "github.com/github/copilot-sdk/go"
)

// BuildBaseClientOpts returns a shared set of Copilot SDK client options
// with the hyoka session environment tag and debug logging when enabled.
// All code that creates a copilot.ClientOptions should use this function
// to ensure consistent configuration.
func BuildBaseClientOpts() *copilot.ClientOptions {
	opts := &copilot.ClientOptions{
		Env: HyokaBaseEnv(),
	}
	if slog.Default().Enabled(context.Background(), slog.LevelDebug) {
		opts.LogLevel = "debug"
	}
	return opts
}
