// Package comparison provides config-vs-config, run-vs-run, and temporal diff
// analysis for hyoka evaluation reports.
//
// ComparisonResult is the canonical, unified result type produced by every
// comparison entry point (CLI, serve API, auto-generated run summaries). It
// replaces the previous trio of ConfigComparison/RunComparison/TemporalComparison
// so that the `hyoka compare` CLI and the site comparison page render
// byte-identical data for the same inputs.
package comparison

import "time"

// ComparisonKind identifies which comparison variant produced a result.
type ComparisonKind string

const (
	// KindConfigs is a pairwise comparison of two configs' latest results
	// across their shared prompt set.
	KindConfigs ComparisonKind = "configs"
	// KindRuns is a comparison of results from two specific evaluation runs.
	KindRuns ComparisonKind = "runs"
	// KindTemporal is a before/after comparison for a single config around a
	// cutoff timestamp.
	KindTemporal ComparisonKind = "temporal"
)

// ComparisonResult is the canonical comparison payload. All three comparison
// modes (configs, runs, temporal) produce this struct. Consumers should switch
// on Kind to interpret LabelA/LabelB appropriately:
//
//   - Kind == KindConfigs  → LabelA/LabelB are config names
//   - Kind == KindRuns     → LabelA/LabelB are run IDs
//   - Kind == KindTemporal → LabelA is the base run ID, LabelB is the latest
//     run ID, Config names the tracked config, and Since is the cutoff
//     timestamp.
//
// This type is the shared interface boundary between the comparison engine
// and any consumer (CLI renderer, serve API JSON handler, site frontend,
// auto-generated run summaries).
type ComparisonResult struct {
	Kind      ComparisonKind    `json:"kind"`
	LabelA    string            `json:"label_a"`
	LabelB    string            `json:"label_b"`
	Config    string            `json:"config,omitempty"` // only set for KindTemporal
	Since     *time.Time        `json:"since,omitempty"`  // only set for KindTemporal
	PerPrompt []PromptDiff      `json:"per_prompt"`
	Summary   ComparisonSummary `json:"summary"`
}
