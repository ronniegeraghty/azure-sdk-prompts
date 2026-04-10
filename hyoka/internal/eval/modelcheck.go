package eval

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
)

// ModelCheckResult holds the outcome of a pre-flight model availability check.
type ModelCheckResult struct {
	// Available lists the model IDs that were found in the backend.
	Available []string
	// Unavailable lists the model IDs that were NOT found.
	Unavailable []string
	// AllModels is the complete set of models returned by the backend.
	AllModels []string
}

// OK returns true if all requested models are available.
func (r *ModelCheckResult) OK() bool {
	return len(r.Unavailable) == 0
}

// Error returns a human-readable summary of unavailable models.
func (r *ModelCheckResult) Error() string {
	if r.OK() {
		return ""
	}
	return fmt.Sprintf("unavailable model(s): %s", strings.Join(r.Unavailable, ", "))
}

// CheckModelAvailability queries the Copilot backend for available models and
// checks that every model referenced by the given configs (generator +
// reviewer) is present. This pre-flight check prevents mid-eval failures
// when a model like "gemini-3-pro-preview" is configured but not available
// in the backend (#264).
func (e *CopilotSDKEvaluator) CheckModelAvailability(ctx context.Context, configs []config.ToolConfig) (*ModelCheckResult, error) {
	// Create a temporary workspace for the client — the ListModels RPC
	// doesn't access the filesystem, but the SDK requires a valid CWD.
	tmpDir, err := os.MkdirTemp("", "hyoka-modelcheck-*")
	if err != nil {
		return nil, fmt.Errorf("creating temp dir for model check: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	checkCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	client, err := e.Client(checkCtx, tmpDir)
	if err != nil {
		return nil, fmt.Errorf("starting client for model check: %w", err)
	}
	defer func() {
		done := make(chan struct{})
		go func() { client.Stop(); close(done) }()
		select {
		case <-done:
		case <-time.After(5 * time.Second):
			client.ForceStop()
		}
	}()

	models, err := client.ListModels(checkCtx)
	if err != nil {
		return nil, fmt.Errorf("listing models: %w", err)
	}

	// Build a set of available model IDs.
	availableSet := make(map[string]bool, len(models))
	allModelIDs := make([]string, 0, len(models))
	for _, m := range models {
		availableSet[m.ID] = true
		allModelIDs = append(allModelIDs, m.ID)
	}

	// Collect all unique model IDs referenced by configs.
	requiredModels := collectRequiredModels(configs)

	result := &ModelCheckResult{AllModels: allModelIDs}
	for _, model := range requiredModels {
		if availableSet[model] {
			result.Available = append(result.Available, model)
		} else {
			result.Unavailable = append(result.Unavailable, model)
			slog.Warn("Model not available", "model", model)
		}
	}

	return result, nil
}

// collectRequiredModels extracts all unique model IDs from a set of configs,
// including both generator and reviewer models.
func collectRequiredModels(configs []config.ToolConfig) []string {
	seen := make(map[string]bool)
	var models []string
	for _, cfg := range configs {
		if cfg.Generator != nil {
			for _, gm := range cfg.Generator.ResolveModels() {
				if !seen[gm] {
					seen[gm] = true
					models = append(models, gm)
				}
			}
		}
		if cfg.Reviewer != nil {
			for _, m := range cfg.Reviewer.Models {
				if !seen[m] {
					seen[m] = true
					models = append(models, m)
				}
			}
		}
	}
	return models
}

// ValidateModelAvailability is a convenience wrapper that runs
// CheckModelAvailability and returns an error if any models are unavailable.
// Intended for use at run start to fail fast (#264).
func (e *CopilotSDKEvaluator) ValidateModelAvailability(ctx context.Context, configs []config.ToolConfig) error {
	result, err := e.CheckModelAvailability(ctx, configs)
	if err != nil {
		return fmt.Errorf("model availability check failed: %w", err)
	}
	if !result.OK() {
		return fmt.Errorf("pre-flight model check: %s — available models: %s",
			result.Error(), strings.Join(result.AllModels, ", "))
	}
	slog.Info("Pre-flight model check passed", "models_checked", len(result.Available))
	return nil
}
