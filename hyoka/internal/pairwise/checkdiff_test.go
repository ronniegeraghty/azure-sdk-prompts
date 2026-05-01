package pairwise

import (
	"testing"
)

func TestComputeCheckDiffs(t *testing.T) {
	baseline := &EvalReportData{
		ConfigName: "baseline",
		Graders: []GraderData{
			{
				Name: "file_check",
				Type: "file",
				Checks: []PointData{
					{Label: "main.py exists", Pass: false, Message: "file not found"},
					{Label: "config.json exists", Pass: true, Message: "found"},
				},
			},
			{
				Name: "build_test",
				Type: "program",
				Checks: []PointData{
					{Label: "build succeeds", Pass: true, Message: "built successfully"},
				},
			},
		},
	}

	variant1 := &EvalReportData{
		ConfigName: "without-tool-a",
		Graders: []GraderData{
			{
				Name: "file_check",
				Type: "file",
				Checks: []PointData{
					{Label: "main.py exists", Pass: true, Message: "file found"},
					{Label: "config.json exists", Pass: true, Message: "found"},
				},
			},
			{
				Name: "build_test",
				Type: "program",
				Checks: []PointData{
					{Label: "build succeeds", Pass: false, Message: "build failed: syntax error"},
				},
			},
		},
	}

	variant2 := &EvalReportData{
		ConfigName: "without-tool-b",
		Graders: []GraderData{
			{
				Name: "file_check",
				Type: "file",
				Checks: []PointData{
					{Label: "main.py exists", Pass: false, Message: "still not found"},
					{Label: "config.json exists", Pass: true, Message: "found"},
				},
			},
			{
				Name: "build_test",
				Type: "program",
				Checks: []PointData{
					{Label: "build succeeds", Pass: true, Message: "built successfully"},
				},
			},
		},
	}

	diffs := ComputeCheckDiffs(baseline, []*EvalReportData{variant1, variant2})

	if len(diffs) != 2 {
		t.Fatalf("expected 2 variant diffs, got %d", len(diffs))
	}

	// Check variant1 diffs
	v1Diffs := diffs["without-tool-a"]
	if len(v1Diffs) != 3 {
		t.Fatalf("expected 3 check diffs for variant1, got %d", len(v1Diffs))
	}

	// Test improved: baseline failed, variant passed
	if v1Diffs[0].CheckLabel != "main.py exists" {
		t.Errorf("expected check_0 to be 'main.py exists', got %q", v1Diffs[0].CheckLabel)
	}
	if v1Diffs[0].Movement != "improved" {
		t.Errorf("expected 'improved', got %q", v1Diffs[0].Movement)
	}
	if !v1Diffs[0].VariantPassed || v1Diffs[0].BaselinePassed {
		t.Error("expected baseline=fail, variant=pass")
	}

	// Test unchanged-pass: both passed
	if v1Diffs[1].CheckLabel != "config.json exists" {
		t.Errorf("expected check_1 to be 'config.json exists', got %q", v1Diffs[1].CheckLabel)
	}
	if v1Diffs[1].Movement != "unchanged" {
		t.Errorf("expected 'unchanged', got %q", v1Diffs[1].Movement)
	}
	if !v1Diffs[1].BaselinePassed || !v1Diffs[1].VariantPassed {
		t.Error("expected both to pass")
	}

	// Test regressed: baseline passed, variant failed
	if v1Diffs[2].CheckLabel != "build succeeds" {
		t.Errorf("expected check_2 to be 'build succeeds', got %q", v1Diffs[2].CheckLabel)
	}
	if v1Diffs[2].Movement != "regressed" {
		t.Errorf("expected 'regressed', got %q", v1Diffs[2].Movement)
	}
	if v1Diffs[2].BaselinePassed && v1Diffs[2].VariantPassed {
		t.Error("expected baseline=pass, variant=fail")
	}
	if v1Diffs[2].Reasoning != "build failed: syntax error" {
		t.Errorf("expected reasoning to be captured, got %q", v1Diffs[2].Reasoning)
	}

	// Check variant2 diffs
	v2Diffs := diffs["without-tool-b"]
	if len(v2Diffs) != 3 {
		t.Fatalf("expected 3 check diffs for variant2, got %d", len(v2Diffs))
	}

	// Test unchanged-fail: both failed
	if v2Diffs[0].Movement != "unchanged" {
		t.Errorf("expected 'unchanged', got %q", v2Diffs[0].Movement)
	}
	if v2Diffs[0].BaselinePassed || v2Diffs[0].VariantPassed {
		t.Error("expected both to fail")
	}

	// Test grader type propagation
	if v1Diffs[0].GraderType != "file" {
		t.Errorf("expected grader_type='file', got %q", v1Diffs[0].GraderType)
	}
	if v1Diffs[2].GraderType != "program" {
		t.Errorf("expected grader_type='program', got %q", v1Diffs[2].GraderType)
	}
}

func TestComputeCheckDiffsNilInputs(t *testing.T) {
	result := ComputeCheckDiffs(nil, nil)
	if result != nil {
		t.Errorf("expected nil for nil inputs, got %v", result)
	}

	baseline := &EvalReportData{ConfigName: "baseline"}
	result = ComputeCheckDiffs(baseline, nil)
	if result != nil {
		t.Errorf("expected nil for nil variants, got %v", result)
	}

	result = ComputeCheckDiffs(nil, []*EvalReportData{baseline})
	if result != nil {
		t.Errorf("expected nil for nil baseline, got %v", result)
	}
}

func TestComputeCheckDiffsMissingChecks(t *testing.T) {
	baseline := &EvalReportData{
		ConfigName: "baseline",
		Graders: []GraderData{
			{
				Name: "grader_a",
				Type: "file",
				Checks: []PointData{
					{Label: "check_1", Pass: true},
					{Label: "check_2", Pass: false},
				},
			},
		},
	}

	// Variant missing second check
	variant := &EvalReportData{
		ConfigName: "without-tool-x",
		Graders: []GraderData{
			{
				Name: "grader_a",
				Type: "file",
				Checks: []PointData{
					{Label: "check_1", Pass: false},
					// check_2 missing
				},
			},
		},
	}

	diffs := ComputeCheckDiffs(baseline, []*EvalReportData{variant})
	vDiffs := diffs["without-tool-x"]

	if len(vDiffs) != 2 {
		t.Fatalf("expected 2 diffs (including missing check), got %d", len(vDiffs))
	}

	// First check should be regressed
	if vDiffs[0].Movement != "regressed" {
		t.Errorf("expected 'regressed' for check_1, got %q", vDiffs[0].Movement)
	}

	// Second check should be unchanged (missing in variant)
	if vDiffs[1].CheckLabel != "check_2" {
		t.Errorf("expected check_2, got %q", vDiffs[1].CheckLabel)
	}
	if vDiffs[1].Movement != "unchanged" {
		t.Errorf("expected 'unchanged' for missing check, got %q", vDiffs[1].Movement)
	}
	if vDiffs[1].VariantPassed {
		t.Error("expected missing check to default to not passed")
	}
}

func TestComputeCheckDiffsExtraChecksInVariant(t *testing.T) {
	baseline := &EvalReportData{
		ConfigName: "baseline",
		Graders: []GraderData{
			{
				Name: "grader_a",
				Type: "file",
				Checks: []PointData{
					{Label: "check_1", Pass: true},
				},
			},
		},
	}

	// Variant has additional check
	variant := &EvalReportData{
		ConfigName: "without-tool-x",
		Graders: []GraderData{
			{
				Name: "grader_a",
				Type: "file",
				Checks: []PointData{
					{Label: "check_1", Pass: true},
					{Label: "check_2_new", Pass: true}, // new check not in baseline
				},
			},
		},
	}

	diffs := ComputeCheckDiffs(baseline, []*EvalReportData{variant})
	vDiffs := diffs["without-tool-x"]

	if len(vDiffs) != 2 {
		t.Fatalf("expected 2 diffs (including new check), got %d", len(vDiffs))
	}

	// New check should be present and marked unchanged
	if vDiffs[1].CheckLabel != "check_2_new" {
		t.Errorf("expected check_2_new, got %q", vDiffs[1].CheckLabel)
	}
	if vDiffs[1].Movement != "unchanged" {
		t.Errorf("expected 'unchanged' for new check, got %q", vDiffs[1].Movement)
	}
}

func TestIndexPoints(t *testing.T) {
	data := &EvalReportData{
		Graders: []GraderData{
			{
				Name: "grader_a",
				Type: "file",
				Checks: []PointData{
					{Label: "check_1", Pass: true, Message: "msg1"},
					{Label: "check_2", Pass: false, Message: "msg2"},
				},
			},
			{
				Name: "grader_b",
				Type: "program",
				Checks: []PointData{
					{Label: "test_1", Pass: true},
				},
			},
		},
	}

	index := indexPoints(data)

	if len(index) != 2 {
		t.Fatalf("expected 2 graders indexed, got %d", len(index))
	}

	graderA := index["grader_a"]
	if len(graderA) != 2 {
		t.Fatalf("expected 2 points for grader_a, got %d", len(graderA))
	}

	if graderA[0].Label != "check_1" || !graderA[0].Pass {
		t.Error("expected check_1 to be indexed correctly")
	}
	if graderA[1].Label != "check_2" || graderA[1].Pass {
		t.Error("expected check_2 to be indexed correctly")
	}
	if graderA[0].Type != "file" {
		t.Errorf("expected grader_type='file', got %q", graderA[0].Type)
	}

	graderB := index["grader_b"]
	if len(graderB) != 1 {
		t.Fatalf("expected 1 point for grader_b, got %d", len(graderB))
	}
}

func TestIndexPointsNil(t *testing.T) {
	index := indexPoints(nil)
	if index != nil {
		t.Errorf("expected nil for nil input, got %v", index)
	}
}
