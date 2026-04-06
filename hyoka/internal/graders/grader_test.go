package graders

import (
	"math"
	"testing"
)

func TestAggregateResultsWeightedAverage(t *testing.T) {
	results := []GraderResult{
		{Kind: KindFile, Name: "file_check", Score: 1.0, Weight: 1.0, Pass: true},
		{Kind: KindPrompt, Name: "review", Score: 0.8, Weight: 2.0, Pass: true},
		{Kind: KindBehavior, Name: "behavior", Score: 0.6, Weight: 1.0, Pass: true},
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

func TestAggregateResultsGateFailureOverrides(t *testing.T) {
	results := []GraderResult{
		{Kind: KindFile, Name: "file_check", Score: 1.0, Weight: 1.0, Pass: true, Gate: true},
		{Kind: KindProgram, Name: "build", Score: 0.0, Weight: 1.0, Pass: false, Gate: true},
		{Kind: KindPrompt, Name: "review", Score: 0.9, Weight: 2.0, Pass: true},
	}

	agg, err := AggregateResults(results)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if agg.Score != 0 {
		t.Errorf("score = %f, want 0 (gate failure)", agg.Score)
	}
	if agg.Pass {
		t.Error("expected pass = false (gate failure)")
	}
	if !agg.GateFailed {
		t.Error("expected gate_failed = true")
	}
}

func TestAggregateResultsGatePassDoesNotOverride(t *testing.T) {
	results := []GraderResult{
		{Kind: KindFile, Name: "file_check", Score: 1.0, Weight: 1.0, Pass: true, Gate: true},
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
		{Kind: KindFile, Name: "a", Score: 1.0, Weight: 0, Pass: true},
		{Kind: KindFile, Name: "b", Score: 0.5, Weight: 0, Pass: true},
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
		{Kind: KindFile, Name: "check", Score: 0.0, Weight: 1.0, Pass: false, Gate: false},
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
	if agg.GateFailed {
		t.Error("gate_failed should be false for non-gate failures")
	}
}

func TestAggregateResultsMultipleGateFailures(t *testing.T) {
	results := []GraderResult{
		{Kind: KindFile, Name: "gate1", Score: 0.0, Weight: 1.0, Pass: false, Gate: true},
		{Kind: KindProgram, Name: "gate2", Score: 0.0, Weight: 1.0, Pass: false, Gate: true},
		{Kind: KindPrompt, Name: "review", Score: 1.0, Weight: 1.0, Pass: true},
	}

	agg, err := AggregateResults(results)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if agg.Score != 0 {
		t.Errorf("score = %f, want 0", agg.Score)
	}
	if !agg.GateFailed {
		t.Error("expected gate_failed = true")
	}
}
