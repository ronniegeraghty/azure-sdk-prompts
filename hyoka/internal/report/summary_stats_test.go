package report

import (
	"testing"
)

func TestComputeSummaryStats(t *testing.T) {
	boolTrue := true
	boolFalse := false

	s := &RunSummary{
		Results: []*EvalReport{
			{
				PromptID:   "p1",
				ConfigName: "baseline",
				Duration:   10.0,
				Success:    true,
				SessionEvents: []SessionEventRecord{
					{Type: "tool.execution_complete", ToolName: "create", ToolSuccess: &boolTrue},
					{Type: "tool.execution_complete", ToolName: "edit", ToolSuccess: &boolTrue},
				},
			},
			{
				PromptID:   "p1",
				ConfigName: "azure-mcp",
				Duration:   15.0,
				Success:    true,
				SessionEvents: []SessionEventRecord{
					{Type: "tool.execution_complete", ToolName: "create", ToolSuccess: &boolTrue},
					{Type: "tool.execution_complete", ToolName: "azure_mcp", ToolSuccess: &boolFalse},
				},
			},
			{
				PromptID:   "p2",
				ConfigName: "baseline",
				Duration:   5.0,
				Success:    false,
				SessionEvents: []SessionEventRecord{
					{Type: "tool.execution_complete", ToolName: "create", ToolSuccess: &boolTrue},
				},
			},
			{
				PromptID:   "p2",
				ConfigName: "azure-mcp",
				Duration:   12.0,
				Success:    true,
				SessionEvents: []SessionEventRecord{
					{Type: "tool.execution_complete", ToolName: "create", ToolSuccess: &boolTrue},
				},
			},
		},
	}

	stats := ComputeSummaryStats(s)

	// Duration by config
	if len(stats.DurationByConfig) != 2 {
		t.Errorf("expected 2 config duration entries, got %d", len(stats.DurationByConfig))
	}
	bl := stats.DurationByConfig["baseline"]
	if bl.Min != 5.0 || bl.Max != 10.0 {
		t.Errorf("baseline duration: expected min=5 max=10, got min=%.1f max=%.1f", bl.Min, bl.Max)
	}

	// Duration by prompt
	if len(stats.DurationByPrompt) != 2 {
		t.Errorf("expected 2 prompt duration entries, got %d", len(stats.DurationByPrompt))
	}

	// Slowest/fastest
	if stats.SlowestEval != "p1/azure-mcp" {
		t.Errorf("expected slowest p1/azure-mcp, got %s", stats.SlowestEval)
	}
	if stats.FastestEval != "p2/baseline" {
		t.Errorf("expected fastest p2/baseline, got %s", stats.FastestEval)
	}

	// Config pass rates
	if len(stats.ConfigPassRates) != 2 {
		t.Fatalf("expected 2 config pass rates, got %d", len(stats.ConfigPassRates))
	}
	for _, cpr := range stats.ConfigPassRates {
		if cpr.Config == "baseline" && cpr.Rate != 50.0 {
			t.Errorf("baseline pass rate: expected 50%%, got %.1f%%", cpr.Rate)
		}
		if cpr.Config == "azure-mcp" && cpr.Rate != 100.0 {
			t.Errorf("azure-mcp pass rate: expected 100%%, got %.1f%%", cpr.Rate)
		}
	}

	// Prompt deltas (p2 passes on azure-mcp but fails on baseline)
	if len(stats.PromptDeltas) != 1 {
		t.Fatalf("expected 1 prompt delta, got %d", len(stats.PromptDeltas))
	}
	if stats.PromptDeltas[0].PromptID != "p2" {
		t.Errorf("expected delta for p2, got %s", stats.PromptDeltas[0].PromptID)
	}

	// Tool usage
	if len(stats.ToolStats) == 0 {
		t.Fatal("expected tool stats")
	}
	// "create" should be most used (4 times)
	if stats.ToolStats[0].Name != "create" || stats.ToolStats[0].Count != 4 {
		t.Errorf("expected create with count 4, got %s with %d", stats.ToolStats[0].Name, stats.ToolStats[0].Count)
	}
}

func TestComputeSummaryStatsEmpty(t *testing.T) {
	s := &RunSummary{Results: []*EvalReport{}}
	stats := ComputeSummaryStats(s)
	if len(stats.DurationByConfig) != 0 {
		t.Error("expected empty stats for empty results")
	}
	if len(stats.ToolStats) != 0 {
		t.Error("expected empty tool stats for empty results")
	}
}

func TestCalcDurationStats(t *testing.T) {
	ds := calcDurationStats([]float64{3.0, 7.0, 5.0})
	if ds.Min != 3.0 {
		t.Errorf("expected min 3, got %.1f", ds.Min)
	}
	if ds.Max != 7.0 {
		t.Errorf("expected max 7, got %.1f", ds.Max)
	}
	if ds.Avg != 5.0 {
		t.Errorf("expected avg 5, got %.1f", ds.Avg)
	}

	empty := calcDurationStats(nil)
	if empty.Min != 0 || empty.Avg != 0 || empty.Max != 0 {
		t.Error("expected zero stats for nil slice")
	}
}

func TestComputeSummaryStatsTimeline(t *testing.T) {
	s := &RunSummary{
		Results: []*EvalReport{
			{
				PromptID:   "p1",
				ConfigName: "baseline",
				Duration:   10.0,
				Success:    true,
				ActionTimeline: &ActionTimelineReport{
					Summary: ActionTimelineSummary{
						TotalActions:     12,
						TotalToolCalls:   4,
						TotalTurns:       2,
						ToolCallDuration: 400.0,
						ToolSuccesses:    3,
						ToolFailures:     1,
					},
				},
			},
			{
				PromptID:   "p2",
				ConfigName: "baseline",
				Duration:   8.0,
				Success:    true,
				ActionTimeline: &ActionTimelineReport{
					Summary: ActionTimelineSummary{
						TotalActions:     8,
						TotalToolCalls:   2,
						TotalTurns:       1,
						ToolCallDuration: 200.0,
						ToolSuccesses:    2,
						ToolFailures:     0,
					},
				},
			},
			{
				// Eval without timeline (e.g. failed early)
				PromptID:   "p3",
				ConfigName: "baseline",
				Duration:   1.0,
				Success:    false,
			},
		},
	}

	stats := ComputeSummaryStats(s)

	if stats.Timeline == nil {
		t.Fatal("expected timeline summary")
	}

	tl := stats.Timeline
	if tl.TotalActions != 20 {
		t.Errorf("expected total_actions=20, got %d", tl.TotalActions)
	}
	if tl.TotalToolCalls != 6 {
		t.Errorf("expected total_tool_calls=6, got %d", tl.TotalToolCalls)
	}
	if tl.TotalTurns != 3 {
		t.Errorf("expected total_turns=3, got %d", tl.TotalTurns)
	}
	// avg actions per eval: 20/2 = 10.0
	if tl.AvgActionsPerEval != 10.0 {
		t.Errorf("expected avg_actions_per_eval=10.0, got %.1f", tl.AvgActionsPerEval)
	}
	// avg tool calls per eval: 6/2 = 3.0
	if tl.AvgToolCallsPerEval != 3.0 {
		t.Errorf("expected avg_tool_calls_per_eval=3.0, got %.1f", tl.AvgToolCallsPerEval)
	}
	// avg turns per eval: 3/2 = 1.5
	if tl.AvgTurnsPerEval != 1.5 {
		t.Errorf("expected avg_turns_per_eval=1.5, got %.1f", tl.AvgTurnsPerEval)
	}
	// avg tool duration: 600/6 = 100.0
	if tl.AvgToolCallDuration != 100.0 {
		t.Errorf("expected avg_tool_call_duration=100.0, got %.1f", tl.AvgToolCallDuration)
	}
	// tool success rate: 5/(5+1) * 100 = 83.3
	if tl.ToolSuccessRate != 83.3 {
		t.Errorf("expected tool_success_rate=83.3, got %.1f", tl.ToolSuccessRate)
	}
}

func TestComputeSummaryStatsNoTimeline(t *testing.T) {
	s := &RunSummary{
		Results: []*EvalReport{
			{PromptID: "p1", ConfigName: "c1", Duration: 5.0, Success: true},
		},
	}
	stats := ComputeSummaryStats(s)
	if stats.Timeline != nil {
		t.Error("expected nil timeline when no evals have timelines")
	}
}
