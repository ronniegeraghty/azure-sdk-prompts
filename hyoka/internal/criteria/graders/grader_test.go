package graders

import (
	"math"
	"testing"
)

func TestAggregateResultsWeightedAverage(t *testing.T) {
	results := []GraderResult{
		{Kind: KindWorkspace, Name: "file_check", Score: 1.0, Weight: 1.0, Pass: true},
		{Kind: KindPrompt, Name: "review", Score: 0.8, Weight: 2.0, Pass: true},
		{Kind: KindTool, Name: "behavior", Score: 0.6, Weight: 1.0, Pass: true},
	}

	agg, err := AggregateResults(results)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Expected: (1.0*1 + 0.8*2 + 0.6*1) / (1+2+1) = 3.2/4 = 0.8
	want := 0.8
	if math.Abs(agg.Score-want) > 1e-9 {
		t.Errorf("score = %f, want %f", agg.Score, want)
	}
	if !agg.Pass {
		t.Error("expected pass = true")
	}
	if agg.GateFailed {
		t.Error("expected gate_failed = false")
	}
}

func TestAggregateResultsGateFieldNoLongerShortCircuits(t *testing.T) {
	// Phase 2 cutover (#625): Gate flag is a legacy field that no longer
	// short-circuits aggregation. Every grader result contributes to the
	// weighted score; Pass is the AND of every result's Pass.
	results := []GraderResult{
		{Kind: KindWorkspace, Name: "file_check", Score: 1.0, Weight: 1.0, Pass: true, Gate: true},
		{Kind: KindProgram, Name: "build", Score: 0.0, Weight: 1.0, Pass: false, Gate: true},
		{Kind: KindPrompt, Name: "review", Score: 0.9, Weight: 2.0, Pass: true},
	}

	agg, err := AggregateResults(results)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// (1.0*1 + 0.0*1 + 0.9*2) / (1+1+2) = 2.8/4 = 0.7
	want := 0.7
	if math.Abs(agg.Score-want) > 1e-9 {
		t.Errorf("score = %f, want %f (no gate short-circuit)", agg.Score, want)
	}
	if agg.Pass {
		t.Error("expected pass = false when any result fails")
	}
	if agg.GateFailed {
		t.Error("GateFailed must stay false under Phase 2 no-gate semantics")
	}
}

func TestAggregateResultsGatePassDoesNotOverride(t *testing.T) {
	results := []GraderResult{
		{Kind: KindWorkspace, Name: "file_check", Score: 1.0, Weight: 1.0, Pass: true, Gate: true},
		{Kind: KindPrompt, Name: "review", Score: 0.5, Weight: 1.0, Pass: true},
	}

	agg, err := AggregateResults(results)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Gate passed, so normal weighted average: (1.0*1 + 0.5*1) / 2 = 0.75
	want := 0.75
	if math.Abs(agg.Score-want) > 1e-9 {
		t.Errorf("score = %f, want %f", agg.Score, want)
	}
	if !agg.Pass {
		t.Error("expected pass = true")
	}
}

func TestAggregateResultsEmptyReturnsError(t *testing.T) {
	_, err := AggregateResults(nil)
	if err == nil {
		t.Fatal("expected error for empty results")
	}

	_, err = AggregateResults([]GraderResult{})
	if err == nil {
		t.Fatal("expected error for empty slice")
	}
}

func TestAggregateResultsDefaultWeight(t *testing.T) {
	results := []GraderResult{
		{Kind: KindWorkspace, Name: "a", Score: 1.0, Weight: 0, Pass: true},
		{Kind: KindWorkspace, Name: "b", Score: 0.5, Weight: 0, Pass: true},
	}

	agg, err := AggregateResults(results)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Zero weight defaults to 1.0: (1.0*1 + 0.5*1) / 2 = 0.75
	want := 0.75
	if math.Abs(agg.Score-want) > 1e-9 {
		t.Errorf("score = %f, want %f", agg.Score, want)
	}
}

func TestAggregateResultsSingleResult(t *testing.T) {
	results := []GraderResult{
		{Kind: KindPrompt, Name: "solo", Score: 0.42, Weight: 1.0, Pass: true},
	}

	agg, err := AggregateResults(results)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if math.Abs(agg.Score-0.42) > 1e-9 {
		t.Errorf("score = %f, want 0.42", agg.Score)
	}
	if len(agg.Results) != 1 {
		t.Errorf("expected 1 result, got %d", len(agg.Results))
	}
}

func TestAggregateResultsNonGateFailureDoesNotForceZero(t *testing.T) {
	results := []GraderResult{
		{Kind: KindWorkspace, Name: "check", Score: 0.0, Weight: 1.0, Pass: false, Gate: false},
		{Kind: KindPrompt, Name: "review", Score: 0.8, Weight: 1.0, Pass: true},
	}

	agg, err := AggregateResults(results)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Non-gate failure: (0.0*1 + 0.8*1) / 2 = 0.4
	want := 0.4
	if math.Abs(agg.Score-want) > 1e-9 {
		t.Errorf("score = %f, want %f", agg.Score, want)
	}
	if agg.Pass {
		t.Error("expected pass = false when any result fails")
	}
	if agg.GateFailed {
		t.Error("gate_failed should be false for non-gate failures")
	}
}

func TestAggregateResultsMultipleFailuresWithGateFlag(t *testing.T) {
	// Phase 2 cutover (#625): multiple failing results with Gate=true no
	// longer force Score=0 or GateFailed=true. Every result contributes.
	results := []GraderResult{
		{Kind: KindWorkspace, Name: "gate1", Score: 0.0, Weight: 1.0, Pass: false, Gate: true},
		{Kind: KindProgram, Name: "gate2", Score: 0.0, Weight: 1.0, Pass: false, Gate: true},
		{Kind: KindPrompt, Name: "review", Score: 1.0, Weight: 1.0, Pass: true},
	}

	agg, err := AggregateResults(results)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// (0 + 0 + 1) / 3 = 0.333...
	want := 1.0 / 3.0
	if math.Abs(agg.Score-want) > 1e-9 {
		t.Errorf("score = %f, want %f", agg.Score, want)
	}
	if agg.Pass {
		t.Error("expected pass = false when any grader fails")
	}
	if agg.GateFailed {
		t.Error("expected gate_failed = false under Phase 2 no-gate semantics")
	}
}
