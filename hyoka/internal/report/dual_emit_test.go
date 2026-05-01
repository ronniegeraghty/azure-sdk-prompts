package report

import (
	"encoding/json"
	"testing"
)

// TestGraderResultDualEmit verifies that GraderResult marshals with both "checks" and "points" keys,
// and unmarshals from either key (preferring "checks" if both are present).
func TestGraderResultDualEmit(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantChecks int
		wantErr    bool
	}{
		{
			name: "unmarshal legacy points only",
			input: `{
				"grader_name": "test",
				"grader_type": "file",
				"score": 1.0,
				"weight": 1.0,
				"pass": true,
				"message": "ok",
				"points": [{"label": "check1", "pass": true}]
			}`,
			wantChecks: 1,
		},
		{
			name: "unmarshal new checks only",
			input: `{
				"grader_name": "test",
				"grader_type": "file",
				"score": 1.0,
				"weight": 1.0,
				"pass": true,
				"message": "ok",
				"checks": [{"label": "check1", "pass": true}, {"label": "check2", "pass": false}]
			}`,
			wantChecks: 2,
		},
		{
			name: "unmarshal both checks and points prefers checks",
			input: `{
				"grader_name": "test",
				"grader_type": "file",
				"score": 1.0,
				"weight": 1.0,
				"pass": true,
				"message": "ok",
				"checks": [{"label": "check1", "pass": true}],
				"points": [{"label": "point1", "pass": false}, {"label": "point2", "pass": true}]
			}`,
			wantChecks: 1, // Should use "checks" not "points"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result GraderResult
			err := json.Unmarshal([]byte(tt.input), &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && len(result.Checks) != tt.wantChecks {
				t.Errorf("Got %d checks, want %d", len(result.Checks), tt.wantChecks)
			}
		})
	}
}

// TestGraderResultMarshalDualEmit verifies that marshaling emits both "checks" and "points" keys
// with identical content.
func TestGraderResultMarshalDualEmit(t *testing.T) {
	result := GraderResult{
		GraderName: "test-grader",
		GraderType: "file",
		Score:      0.75,
		Weight:     1.0,
		Pass:       true,
		Message:    "ok",
		Checks: []GraderCheck{
			{Label: "check1", Pass: true, Message: "passed"},
			{Label: "check2", Pass: false, Message: "failed"},
		},
	}

	data, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Parse to a map to inspect both keys
	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("Unmarshal to map failed: %v", err)
	}

	// Both "checks" and "points" should exist
	checksVal, hasChecks := m["checks"]
	pointsVal, hasPoints := m["points"]

	if !hasChecks {
		t.Error("Marshal did not emit 'checks' key")
	}
	if !hasPoints {
		t.Error("Marshal did not emit 'points' key")
	}

	// Both should be arrays of the same length
	if hasChecks && hasPoints {
		checksArr, ok1 := checksVal.([]interface{})
		pointsArr, ok2 := pointsVal.([]interface{})
		if !ok1 || !ok2 {
			t.Fatal("checks or points is not an array")
		}
		if len(checksArr) != len(pointsArr) {
			t.Errorf("checks length %d != points length %d", len(checksArr), len(pointsArr))
		}
		if len(checksArr) != 2 {
			t.Errorf("Expected 2 checks, got %d", len(checksArr))
		}
	}
}
