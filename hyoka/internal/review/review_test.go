package review

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	copilot "github.com/github/copilot-sdk/go"
)

// ---------------------------------------------------------------------------
// BuildReviewPrompt tests
// ---------------------------------------------------------------------------

func TestBuildReviewPrompt(t *testing.T) {
	prompt := "Write Azure Blob Storage auth code"
	generated := map[string]string{
		"Program.cs": "using Azure.Storage.Blobs;\n// ...",
	}
	reference := map[string]string{
		"Program.cs": "using Azure.Storage.Blobs;\n// reference",
	}

	result := BuildReviewPrompt(prompt, generated, reference, nil, nil)

	if result == "" {
		t.Fatal("expected non-empty review prompt")
	}

	checks := []string{
		"Original Prompt",
		"Generated Code",
		"Reference Answer",
		"Scoring Instructions",
		"passed",
		"Program.cs",
	}
	for _, check := range checks {
		if !strings.Contains(result, check) {
			t.Errorf("review prompt missing %q", check)
		}
	}
}

func TestBuildReviewPromptNoReference(t *testing.T) {
	prompt := "Write code"
	generated := map[string]string{"main.go": "package main"}

	result := BuildReviewPrompt(prompt, generated, nil, nil, nil)

	if !strings.Contains(result, "No reference answer provided") {
		t.Error("expected 'No reference answer provided' when no reference given")
	}
}

func TestBuildReviewPromptEmptyReference(t *testing.T) {
	result := BuildReviewPrompt("prompt", map[string]string{"a.go": "code"}, map[string]string{}, nil, nil)
	if !strings.Contains(result, "No reference answer provided") {
		t.Error("empty reference map should show 'No reference answer provided'")
	}
}

func TestBuildReviewPromptWithEvaluationCriteria(t *testing.T) {
	prompt := "Write Azure code"
	generated := map[string]string{"main.go": "package main"}
	checks := []ReviewCheck{
		{ID: "check_1", Text: "Must use DefaultAzureCredential"},
		{ID: "check_2", Text: "Must handle errors"},
	}

	result := BuildReviewPrompt(prompt, generated, nil, checks, nil)

	if !strings.Contains(result, "Evaluation Criteria") {
		t.Error("expected evaluation criteria section")
	}
	if !strings.Contains(result, "check_1") || !strings.Contains(result, "DefaultAzureCredential") {
		t.Error("expected criteria content in prompt")
	}
}

func TestBuildReviewPromptNoCriteria(t *testing.T) {
	result := BuildReviewPrompt("prompt", map[string]string{"a.go": "code"}, nil, nil, nil)
	// When no criteria are passed, the "Evaluation Criteria" section should not appear.
	if strings.Contains(result, "## Evaluation Criteria") {
		t.Error("should not contain criteria section header when criteria is empty")
	}
}

func TestBuildReviewPromptMultipleFiles(t *testing.T) {
	generated := map[string]string{
		"main.go":   "package main",
		"helper.go": "package helper",
		"util.go":   "package util",
	}
	reference := map[string]string{
		"ref_main.go": "package main // ref",
		"ref_help.go": "package helper // ref",
	}

	result := BuildReviewPrompt("prompt", generated, reference, nil, nil)

	for name := range generated {
		if !strings.Contains(result, name) {
			t.Errorf("prompt missing generated file %q", name)
		}
	}
	for name := range reference {
		if !strings.Contains(result, name) {
			t.Errorf("prompt missing reference file %q", name)
		}
	}
}

func TestBuildReviewPromptEmptyGeneratedFiles(t *testing.T) {
	result := BuildReviewPrompt("prompt", map[string]string{}, nil, nil, nil)
	if !strings.Contains(result, "Generated Code") {
		t.Error("should still contain Generated Code header even with empty files")
	}
}

func TestBuildReviewPromptContainsScoringInstructions(t *testing.T) {
	result := BuildReviewPrompt("p", map[string]string{"f": "c"}, nil, nil, nil)
	if !strings.Contains(result, "Scoring Instructions") {
		t.Error("prompt should contain scoring instructions")
	}
}

func TestBuildReviewPromptPreservesOriginalPrompt(t *testing.T) {
	original := "Write a Python script that uses azure-identity DefaultAzureCredential"
	result := BuildReviewPrompt(original, map[string]string{"main.py": "pass"}, nil, nil, nil)
	if !strings.Contains(result, original) {
		t.Error("prompt should contain the original prompt verbatim")
	}
}

// ---------------------------------------------------------------------------
// parseReviewResponse tests
// ---------------------------------------------------------------------------

func TestParseReviewResponse(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantErr   bool
		score     int
		maxScore  int
		criteria  int
		summary   string
		issues    int
		strengths int
	}{
		{
			name:      "clean json with criteria",
			input:     `{"scores":{"criteria":[{"name":"Code Builds","passed":true,"reason":"OK"},{"name":"Best Practices","passed":true,"reason":"Good"},{"name":"Error Handling","passed":false,"reason":"Missing"}]},"overall_score":2,"max_score":3,"summary":"Good code","issues":["Missing retry"],"strengths":["Clean"]}`,
			score:     2,
			maxScore:  3,
			criteria:  3,
			summary:   "Good code",
			issues:    1,
			strengths: 1,
		},
		{
			name:     "wrapped in markdown json fence",
			input:    "```json\n" + `{"scores":{"criteria":[{"name":"Code Builds","passed":true}]},"overall_score":1,"max_score":1,"summary":"Good","issues":[],"strengths":[]}` + "\n```",
			score:    1,
			maxScore: 1,
			criteria: 1,
			summary:  "Good",
		},
		{
			name:     "wrapped in plain markdown fence",
			input:    "```\n" + `{"scores":{"criteria":[{"name":"X","passed":false}]},"overall_score":0,"max_score":1,"summary":"Bad","issues":["everything"],"strengths":[]}` + "\n```",
			score:    0,
			maxScore: 1,
			criteria: 1,
			issues:   1,
		},
		{
			name:    "no json",
			input:   "I cannot review this code because...",
			wantErr: true,
		},
		{
			name:    "empty",
			input:   "",
			wantErr: true,
		},
		{
			name:    "only whitespace",
			input:   "   \n\t  \n  ",
			wantErr: true,
		},
		{
			name:     "auto-fill max_score from criteria count",
			input:    `{"scores":{"criteria":[{"name":"A","passed":true},{"name":"B","passed":true},{"name":"C","passed":false}]},"overall_score":0,"max_score":0,"summary":"test","issues":[],"strengths":[]}`,
			score:    2,
			maxScore: 3,
			criteria: 3,
		},
		{
			name:     "auto-fill overall_score from criteria",
			input:    `{"scores":{"criteria":[{"name":"A","passed":true},{"name":"B","passed":false}]},"summary":"test","issues":[],"strengths":[]}`,
			score:    1,
			maxScore: 2,
			criteria: 2,
		},
		{
			name:     "json with surrounding text",
			input:    `Here is my review: {"scores":{"criteria":[{"name":"Build","passed":true}]},"overall_score":1,"max_score":1,"summary":"ok","issues":[],"strengths":[]} End of review.`,
			score:    1,
			maxScore: 1,
			criteria: 1,
		},
		{
			name:      "all criteria passed",
			input:     `{"scores":{"criteria":[{"name":"A","passed":true},{"name":"B","passed":true}]},"overall_score":2,"max_score":2,"summary":"Perfect","issues":[],"strengths":["Great"]}`,
			score:     2,
			maxScore:  2,
			criteria:  2,
			strengths: 1,
		},
		{
			name:     "all criteria failed",
			input:    `{"scores":{"criteria":[{"name":"A","passed":false},{"name":"B","passed":false}]},"overall_score":0,"max_score":2,"summary":"Bad","issues":["A failed","B failed"],"strengths":[]}`,
			score:    0,
			maxScore: 2,
			criteria: 2,
			issues:   2,
		},
		{
			name:    "invalid json structure",
			input:   `{"scores": "not an object"}`,
			wantErr: true,
		},
		{
			name:     "empty criteria list",
			input:    `{"scores":{"criteria":[]},"overall_score":0,"max_score":0,"summary":"Nothing to evaluate","issues":[],"strengths":[]}`,
			score:    0,
			maxScore: 0,
			criteria: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseReviewResponse(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.OverallScore != tt.score {
				t.Errorf("OverallScore = %d, want %d", result.OverallScore, tt.score)
			}
			if tt.maxScore > 0 && result.MaxScore != tt.maxScore {
				t.Errorf("MaxScore = %d, want %d", result.MaxScore, tt.maxScore)
			}
			if tt.criteria > 0 && len(result.Scores.Criteria) != tt.criteria {
				t.Errorf("Criteria count = %d, want %d", len(result.Scores.Criteria), tt.criteria)
			}
			if tt.summary != "" && result.Summary != tt.summary {
				t.Errorf("Summary = %q, want %q", result.Summary, tt.summary)
			}
			if tt.issues > 0 && len(result.Issues) != tt.issues {
				t.Errorf("Issues count = %d, want %d", len(result.Issues), tt.issues)
			}
			if tt.strengths > 0 && len(result.Strengths) != tt.strengths {
				t.Errorf("Strengths count = %d, want %d", len(result.Strengths), tt.strengths)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// ReviewScores tests
// ---------------------------------------------------------------------------

func TestReviewScoresPassedCount(t *testing.T) {
	tests := []struct {
		name     string
		criteria []CriterionResult
		want     int
	}{
		{"all passed", []CriterionResult{
			{Name: "A", Passed: true},
			{Name: "B", Passed: true},
		}, 2},
		{"none passed", []CriterionResult{
			{Name: "A", Passed: false},
			{Name: "B", Passed: false},
		}, 0},
		{"mixed", []CriterionResult{
			{Name: "A", Passed: true},
			{Name: "B", Passed: false},
			{Name: "C", Passed: true},
		}, 2},
		{"empty", nil, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ReviewScores{Criteria: tt.criteria}
			if got := s.PassedCount(); got != tt.want {
				t.Errorf("PassedCount() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestReviewScoresTotalCount(t *testing.T) {
	tests := []struct {
		name     string
		criteria []CriterionResult
		want     int
	}{
		{"three criteria", []CriterionResult{
			{Name: "A"}, {Name: "B"}, {Name: "C"},
		}, 3},
		{"empty", nil, 0},
		{"one", []CriterionResult{{Name: "A"}}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ReviewScores{Criteria: tt.criteria}
			if got := s.TotalCount(); got != tt.want {
				t.Errorf("TotalCount() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestReviewScoresAllPassed(t *testing.T) {
	tests := []struct {
		name     string
		criteria []CriterionResult
		want     bool
	}{
		{"all passed", []CriterionResult{
			{Name: "A", Passed: true},
			{Name: "B", Passed: true},
		}, true},
		{"one failed", []CriterionResult{
			{Name: "A", Passed: true},
			{Name: "B", Passed: false},
		}, false},
		{"none passed", []CriterionResult{
			{Name: "A", Passed: false},
		}, false},
		{"empty returns false", nil, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ReviewScores{Criteria: tt.criteria}
			if got := s.AllPassed(); got != tt.want {
				t.Errorf("AllPassed() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// StubReviewer tests
// ---------------------------------------------------------------------------

func TestStubReviewer(t *testing.T) {
	s := &StubReviewer{}
	result, err := s.Review(nil, "test prompt", "some-dir", "", "", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Summary != "Review skipped (stub evaluator)" {
		t.Errorf("unexpected summary: %s", result.Summary)
	}
}

func TestStubReviewerScores(t *testing.T) {
	s := &StubReviewer{}
	result, err := s.Review(nil, "prompt", "dir", "", "", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.OverallScore != 1 {
		t.Errorf("OverallScore = %d, want 1", result.OverallScore)
	}
	if result.MaxScore != 1 {
		t.Errorf("MaxScore = %d, want 1", result.MaxScore)
	}
	if len(result.Scores.Criteria) != 1 {
		t.Fatalf("Criteria count = %d, want 1", len(result.Scores.Criteria))
	}
	c := result.Scores.Criteria[0]
	if c.Name != "stub_criterion" {
		t.Errorf("criterion name = %q, want %q", c.Name, "stub_criterion")
	}
	if !c.Passed {
		t.Error("stub criterion should pass")
	}
	if result.Issues == nil {
		t.Error("Issues should not be nil")
	}
	if result.Strengths == nil {
		t.Error("Strengths should not be nil")
	}
}

func TestStubReviewerIgnoresInputs(t *testing.T) {
	s := &StubReviewer{}
	r1, _ := s.Review(nil, "prompt1", "dir1", "ref1", "criteria1", nil)
	r2, _ := s.Review(nil, "prompt2", "dir2", "ref2", "criteria2", nil)

	if r1.Summary != r2.Summary {
		t.Error("stub reviewer should return identical results regardless of inputs")
	}
	if r1.OverallScore != r2.OverallScore {
		t.Error("stub reviewer should return identical scores regardless of inputs")
	}
}

// ---------------------------------------------------------------------------
// NewCopilotReviewer tests
// ---------------------------------------------------------------------------

func TestNewCopilotReviewerDefaultModel(t *testing.T) {
	r := NewCopilotReviewer(nil, "", 50)
	if r.model != "claude-sonnet-4.5" {
		t.Errorf("default model = %q, want %q", r.model, "claude-sonnet-4.5")
	}
}

func TestNewCopilotReviewerCustomModel(t *testing.T) {
	r := NewCopilotReviewer(nil, "gpt-4o", 100)
	if r.model != "gpt-4o" {
		t.Errorf("model = %q, want %q", r.model, "gpt-4o")
	}
	if r.maxSessionActions != 100 {
		t.Errorf("maxSessionActions = %d, want 100", r.maxSessionActions)
	}
}

func TestCopilotReviewerSetSkillDirectories(t *testing.T) {
	r := NewCopilotReviewer(nil, "", 50)
	dirs := []string{"/skills/gen", "/skills/rev"}
	r.SetSkillDirectories(dirs)
	if len(r.skillDirectories) != 2 {
		t.Errorf("skillDirectories count = %d, want 2", len(r.skillDirectories))
	}
	for i, d := range dirs {
		if r.skillDirectories[i] != d {
			t.Errorf("skillDirectories[%d] = %q, want %q", i, r.skillDirectories[i], d)
		}
	}
}

func TestCopilotReviewerSetSessionTimeout(t *testing.T) {
	r := NewCopilotReviewer(nil, "", 50)
	r.SetSessionTimeout(5 * time.Minute)
	if r.sessionTimeout != 5*time.Minute {
		t.Errorf("sessionTimeout = %v, want %v", r.sessionTimeout, 5*time.Minute)
	}
}

// ---------------------------------------------------------------------------
// PanelReviewer construction tests
// ---------------------------------------------------------------------------

func TestNewPanelReviewer(t *testing.T) {
	models := []string{"model-a", "model-b", "model-c"}
	p := NewPanelReviewer(nil, models, 25)

	if len(p.models) != 3 {
		t.Fatalf("model count = %d, want 3", len(p.models))
	}
	if p.maxSessionActions != 25 {
		t.Errorf("maxSessionActions = %d, want 25", p.maxSessionActions)
	}
}

func TestPanelReviewerModels(t *testing.T) {
	models := []string{"a", "b"}
	p := NewPanelReviewer(nil, models, 10)
	got := p.Models()
	if len(got) != len(models) {
		t.Fatalf("Models() returned %d items, want %d", len(got), len(models))
	}
	for i, m := range models {
		if got[i] != m {
			t.Errorf("Models()[%d] = %q, want %q", i, got[i], m)
		}
	}
}

func TestPanelReviewerSetSkillDirectories(t *testing.T) {
	p := NewPanelReviewer(nil, []string{"m"}, 10)
	dirs := []string{"/a", "/b"}
	p.SetSkillDirectories(dirs)
	if len(p.skillDirectories) != 2 {
		t.Errorf("skillDirectories = %d, want 2", len(p.skillDirectories))
	}
}

func TestPanelReviewerSetSessionTimeout(t *testing.T) {
	p := NewPanelReviewer(nil, []string{"m"}, 10)
	p.SetSessionTimeout(3 * time.Minute)
	if p.sessionTimeout != 3*time.Minute {
		t.Errorf("sessionTimeout = %v, want %v", p.sessionTimeout, 3*time.Minute)
	}
}

// ---------------------------------------------------------------------------
// averageReview tests
// ---------------------------------------------------------------------------

func TestAverageReviewEmpty(t *testing.T) {
	result := averageReview(nil, nil)
	if result.Summary != "No reviews to consolidate" {
		t.Errorf("Summary = %q, want %q", result.Summary, "No reviews to consolidate")
	}
}

func TestAverageReviewSingleReviewer(t *testing.T) {
	panel := []ReviewResult{{
		Model: "model-a",
		Scores: ReviewScores{Criteria: []CriterionResult{
			{Name: "Build", Passed: true, Reason: "ok"},
			{Name: "Style", Passed: false, Reason: "messy"},
		}},
		OverallScore: 1,
		MaxScore:     2,
		Summary:      "Decent",
		Issues:       []string{"messy code"},
		Strengths:    []string{"compiles"},
	}}

	result := averageReview(panel, nil)

	if result.OverallScore != 1 {
		t.Errorf("OverallScore = %d, want 1", result.OverallScore)
	}
	if result.MaxScore != 2 {
		t.Errorf("MaxScore = %d, want 2", result.MaxScore)
	}
	// With 1 reviewer: 1/1 > 1/2 = true for Build, 0/1 > 0 = false for Style
	buildPassed := false
	stylePassed := false
	for _, c := range result.Scores.Criteria {
		if c.Name == "Build" {
			buildPassed = c.Passed
		}
		if c.Name == "Style" {
			stylePassed = c.Passed
		}
	}
	if !buildPassed {
		t.Error("Build should pass with 1/1 majority")
	}
	if stylePassed {
		t.Error("Style should fail with 0/1 majority")
	}
}

func TestAverageReviewMajorityVoting(t *testing.T) {
	panel := []ReviewResult{
		{
			Model: "m1",
			Scores: ReviewScores{Criteria: []CriterionResult{
				{Name: "Build", Passed: true},
				{Name: "Style", Passed: true},
				{Name: "Errors", Passed: false},
			}},
			Issues:    []string{"no retries"},
			Strengths: []string{"clean"},
		},
		{
			Model: "m2",
			Scores: ReviewScores{Criteria: []CriterionResult{
				{Name: "Build", Passed: true},
				{Name: "Style", Passed: false},
				{Name: "Errors", Passed: true},
			}},
			Issues:    []string{"inconsistent style"},
			Strengths: []string{"handles errors"},
		},
		{
			Model: "m3",
			Scores: ReviewScores{Criteria: []CriterionResult{
				{Name: "Build", Passed: true},
				{Name: "Style", Passed: false},
				{Name: "Errors", Passed: true},
			}},
			Issues:    []string{"no retries"},
			Strengths: []string{"clean"},
		},
	}

	result := averageReview(panel, nil)

	criteriaMap := map[string]bool{}
	for _, c := range result.Scores.Criteria {
		criteriaMap[c.Name] = c.Passed
	}

	// Build: 3/3 pass → pass
	if !criteriaMap["Build"] {
		t.Error("Build should pass (0/3 failed)")
	}
	// Style: 1/3 pass → fail (strict: any fail = fail)
	if criteriaMap["Style"] {
		t.Error("Style should fail (2/3 failed, strict voting)")
	}
	// Errors: 2/3 pass → fail (strict: any fail = fail)
	if criteriaMap["Errors"] {
		t.Error("Errors should fail (1/3 failed, strict any-fail voting)")
	}

	// Verify correct overall score: only Build passes = 1
	if result.OverallScore != 1 {
		t.Errorf("OverallScore = %d, want 1", result.OverallScore)
	}
	if result.MaxScore != 3 {
		t.Errorf("MaxScore = %d, want 3", result.MaxScore)
	}
}

func TestAverageReviewDeduplicatesIssuesAndStrengths(t *testing.T) {
	panel := []ReviewResult{
		{
			Scores:    ReviewScores{Criteria: []CriterionResult{{Name: "A", Passed: true}}},
			Issues:    []string{"dup issue", "unique1"},
			Strengths: []string{"dup strength", "unique_s1"},
		},
		{
			Scores:    ReviewScores{Criteria: []CriterionResult{{Name: "A", Passed: true}}},
			Issues:    []string{"dup issue", "unique2"},
			Strengths: []string{"dup strength", "unique_s2"},
		},
	}

	result := averageReview(panel, nil)

	if len(result.Issues) != 3 {
		t.Errorf("Issues count = %d, want 3 (dedup 'dup issue')", len(result.Issues))
	}
	if len(result.Strengths) != 3 {
		t.Errorf("Strengths count = %d, want 3 (dedup 'dup strength')", len(result.Strengths))
	}
}

func TestAverageReviewDisjointCriteria(t *testing.T) {
	panel := []ReviewResult{
		{
			Scores: ReviewScores{Criteria: []CriterionResult{
				{Name: "Build", Passed: true},
			}},
		},
		{
			Scores: ReviewScores{Criteria: []CriterionResult{
				{Name: "Style", Passed: false},
			}},
		},
	}

	result := averageReview(panel, nil)

	if len(result.Scores.Criteria) != 2 {
		t.Errorf("Criteria count = %d, want 2 (union of disjoint sets)", len(result.Scores.Criteria))
	}
	// Build: 1/1 → pass; Style: 0/1 → fail
	criteriaMap := map[string]bool{}
	for _, c := range result.Scores.Criteria {
		criteriaMap[c.Name] = c.Passed
	}
	if !criteriaMap["Build"] {
		t.Error("Build should pass (1/1)")
	}
	if criteriaMap["Style"] {
		t.Error("Style should fail (0/1)")
	}
}

func TestAverageReviewSummaryFormat(t *testing.T) {
	panel := []ReviewResult{
		{
			Scores: ReviewScores{Criteria: []CriterionResult{
				{Name: "A", Passed: true},
			}},
		},
		{
			Scores: ReviewScores{Criteria: []CriterionResult{
				{Name: "A", Passed: true},
			}},
		},
	}

	result := averageReview(panel, nil)

	if !strings.Contains(result.Summary, "2 reviewers") {
		t.Errorf("Summary should mention reviewer count, got: %q", result.Summary)
	}
	if result.Model != "consensus (strict-vote)" {
		t.Errorf("Model = %q, want %q", result.Model, "consensus (strict-vote)")
	}
}

func TestAverageReviewPreservesCriteriaOrder(t *testing.T) {
	panel := []ReviewResult{{
		Scores: ReviewScores{Criteria: []CriterionResult{
			{Name: "Build", Passed: true},
			{Name: "Style", Passed: true},
			{Name: "Errors", Passed: true},
			{Name: "Docs", Passed: true},
		}},
	}}

	result := averageReview(panel, nil)

	expected := []string{"Build", "Style", "Errors", "Docs"}
	for i, c := range result.Scores.Criteria {
		if c.Name != expected[i] {
			t.Errorf("Criteria[%d].Name = %q, want %q", i, c.Name, expected[i])
		}
	}
}

func TestAverageReviewEvenSplitFailsByCriteria(t *testing.T) {
	// With 2 reviewers, 1 pass + 1 fail → strict any-fail = fail
	panel := []ReviewResult{
		{Scores: ReviewScores{Criteria: []CriterionResult{{Name: "X", Passed: true}}}},
		{Scores: ReviewScores{Criteria: []CriterionResult{{Name: "X", Passed: false}}}},
	}

	result := averageReview(panel, nil)

	if len(result.Scores.Criteria) != 1 {
		t.Fatal("expected 1 criterion")
	}
	if result.Scores.Criteria[0].Passed {
		t.Error("tie (1/2) should fail — strict any-fail voting")
	}
}

// ---------------------------------------------------------------------------
// copyDirToTemp tests
// ---------------------------------------------------------------------------

func TestCopyDirToTemp(t *testing.T) {
	src := t.TempDir()
	os.WriteFile(filepath.Join(src, "main.go"), []byte("package main"), 0644)
	sub := filepath.Join(src, "pkg")
	os.MkdirAll(sub, 0755)
	os.WriteFile(filepath.Join(sub, "lib.go"), []byte("package pkg"), 0644)

	dst, err := copyDirToTemp(src, "hyoka-test-*")
	if err != nil {
		t.Fatalf("copyDirToTemp failed: %v", err)
	}
	defer os.RemoveAll(dst)

	data, err := os.ReadFile(filepath.Join(dst, "main.go"))
	if err != nil {
		t.Fatalf("failed to read copied main.go: %v", err)
	}
	if string(data) != "package main" {
		t.Errorf("main.go content = %q, want %q", string(data), "package main")
	}

	data, err = os.ReadFile(filepath.Join(dst, "pkg", "lib.go"))
	if err != nil {
		t.Fatalf("failed to read copied pkg/lib.go: %v", err)
	}
	if string(data) != "package pkg" {
		t.Errorf("pkg/lib.go content = %q, want %q", string(data), "package pkg")
	}
}

func TestCopyDirToTempSkipsDotDirs(t *testing.T) {
	src := t.TempDir()
	os.WriteFile(filepath.Join(src, "main.go"), []byte("package main"), 0644)
	hidden := filepath.Join(src, ".git")
	os.MkdirAll(hidden, 0755)
	os.WriteFile(filepath.Join(hidden, "config"), []byte("gitconfig"), 0644)

	dst, err := copyDirToTemp(src, "hyoka-test-*")
	if err != nil {
		t.Fatalf("copyDirToTemp failed: %v", err)
	}
	defer os.RemoveAll(dst)

	if _, err := os.Stat(filepath.Join(dst, ".git")); !os.IsNotExist(err) {
		t.Error("hidden .git directory should not be copied")
	}
}

func TestCopyDirToTempSkipsBuildArtifactDirs(t *testing.T) {
	src := t.TempDir()
	os.WriteFile(filepath.Join(src, "main.go"), []byte("package main"), 0644)
	nm := filepath.Join(src, "node_modules")
	os.MkdirAll(nm, 0755)
	os.WriteFile(filepath.Join(nm, "pkg.json"), []byte("{}"), 0644)

	dst, err := copyDirToTemp(src, "hyoka-test-*")
	if err != nil {
		t.Fatalf("copyDirToTemp failed: %v", err)
	}
	defer os.RemoveAll(dst)

	if _, err := os.Stat(filepath.Join(dst, "node_modules")); !os.IsNotExist(err) {
		t.Error("node_modules should be skipped as build artifact dir")
	}
}

func TestCopyDirToTempEmptyDir(t *testing.T) {
	src := t.TempDir()

	dst, err := copyDirToTemp(src, "hyoka-test-*")
	if err != nil {
		t.Fatalf("copyDirToTemp failed: %v", err)
	}
	defer os.RemoveAll(dst)

	entries, err := os.ReadDir(dst)
	if err != nil {
		t.Fatalf("failed to read dst: %v", err)
	}
	if len(entries) != 0 {
		t.Errorf("expected empty directory, got %d entries", len(entries))
	}
}

// ---------------------------------------------------------------------------
// ReviewResult / ReviewEvent structural tests
// ---------------------------------------------------------------------------

func TestReviewResultZeroValue(t *testing.T) {
	var r ReviewResult
	if r.OverallScore != 0 {
		t.Errorf("zero-value OverallScore = %d", r.OverallScore)
	}
	if r.Scores.PassedCount() != 0 {
		t.Error("zero-value PassedCount should be 0")
	}
	if r.Scores.AllPassed() {
		t.Error("zero-value AllPassed should be false")
	}
}

func TestReviewEventFields(t *testing.T) {
	evt := ReviewEvent{
		Type:     "tool_execution_complete",
		ToolName: "read_file",
		ToolArgs: `{"path": "main.go"}`,
		Content:  "file content here",
		Result:   "success",
		Error:    "",
		Duration: 123.45,
	}
	if evt.Type != "tool_execution_complete" {
		t.Error("Type mismatch")
	}
	if evt.Duration != 123.45 {
		t.Errorf("Duration = %f, want 123.45", evt.Duration)
	}
}

// ---------------------------------------------------------------------------
// Reviewer interface compliance
// ---------------------------------------------------------------------------

func TestStubReviewerImplementsReviewer(t *testing.T) {
	var _ Reviewer = &StubReviewer{}
}

func TestPanelReviewerImplementsReviewer(t *testing.T) {
	var _ Reviewer = &PanelReviewer{}
}

func TestCopilotReviewerImplementsReviewer(t *testing.T) {
	var _ Reviewer = &CopilotReviewer{}
}

func TestCopilotReviewer_SetSystemPrompt(t *testing.T) {
	r := NewCopilotReviewer(nil, "claude-sonnet-4.5", 50)

	if r.systemPrompt != "" {
		t.Errorf("expected empty default systemPrompt, got %q", r.systemPrompt)
	}

	r.SetSystemPrompt("You are a strict reviewer.")
	if r.systemPrompt != "You are a strict reviewer." {
		t.Errorf("expected custom systemPrompt, got %q", r.systemPrompt)
	}

	r.SetSystemPrompt("")
	if r.systemPrompt != "" {
		t.Errorf("expected empty systemPrompt after clear, got %q", r.systemPrompt)
	}
}

func TestPanelReviewer_SetSystemPrompt(t *testing.T) {
	p := NewPanelReviewer(nil, []string{"model-a", "model-b"}, 50)

	if p.systemPrompt != "" {
		t.Errorf("expected empty default systemPrompt, got %q", p.systemPrompt)
	}

	p.SetSystemPrompt("You are a review judge.")
	if p.systemPrompt != "You are a review judge." {
		t.Errorf("expected custom systemPrompt, got %q", p.systemPrompt)
	}
}

// ---------------------------------------------------------------------------
// CopilotReviewer.Review error-path tests
// ---------------------------------------------------------------------------

func TestCopilotReviewerReviewNoGeneratedFiles(t *testing.T) {
	emptyDir := t.TempDir()
	r := NewCopilotReviewer(nil, "test-model", 50)
	_, err := r.Review(context.Background(), "prompt", emptyDir, "", "", nil)
	if err == nil {
		t.Fatal("expected error for empty workDir")
	}
	if !strings.Contains(err.Error(), "no generated files") {
		t.Errorf("error = %q, want to contain 'no generated files'", err.Error())
	}
}

func TestCopilotReviewerReviewNonexistentWorkDir(t *testing.T) {
	r := NewCopilotReviewer(nil, "test-model", 50)
	_, err := r.Review(context.Background(), "prompt", "/nonexistent/dir/abc", "", "", nil)
	if err == nil {
		t.Fatal("expected error for nonexistent workDir")
	}
}

func TestCopilotReviewerSettersEdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		check func(t *testing.T)
	}{
		{
			name: "SetSkillDirectories nil",
			check: func(t *testing.T) {
				r := NewCopilotReviewer(nil, "", 50)
				r.SetSkillDirectories(nil)
				if r.skillDirectories != nil {
					t.Error("expected nil skillDirectories")
				}
			},
		},
		{
			name: "SetSkillDirectories empty",
			check: func(t *testing.T) {
				r := NewCopilotReviewer(nil, "", 50)
				r.SetSkillDirectories([]string{})
				if len(r.skillDirectories) != 0 {
					t.Error("expected empty skillDirectories")
				}
			},
		},
		{
			name: "SetSessionTimeout zero",
			check: func(t *testing.T) {
				r := NewCopilotReviewer(nil, "", 50)
				r.SetSessionTimeout(0)
				if r.sessionTimeout != 0 {
					t.Errorf("expected zero timeout, got %v", r.sessionTimeout)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.check(t)
		})
	}
}

// ---------------------------------------------------------------------------
// PanelReviewer.ReviewPanel error-path tests
// ---------------------------------------------------------------------------

func TestPanelReviewerReviewPanelNoModels(t *testing.T) {
	p := NewPanelReviewer(nil, []string{}, 50)
	_, _, err := p.ReviewPanel(context.Background(), "prompt", t.TempDir(), "", "", nil)
	if err == nil {
		t.Fatal("expected error for empty models")
	}
	if !strings.Contains(err.Error(), "no reviewer models configured") {
		t.Errorf("error = %q, want 'no reviewer models configured'", err.Error())
	}
}

func TestPanelReviewerReviewPanelNoGeneratedFiles(t *testing.T) {
	emptyDir := t.TempDir()
	p := NewPanelReviewer(nil, []string{"model-a"}, 50)
	_, _, err := p.ReviewPanel(context.Background(), "prompt", emptyDir, "", "", nil)
	if err == nil {
		t.Fatal("expected error for empty workDir")
	}
	if !strings.Contains(err.Error(), "no generated files to review") {
		t.Errorf("error = %q, want 'no generated files to review'", err.Error())
	}
}

func TestPanelReviewerReviewDelegatesToReviewPanel(t *testing.T) {
	p := NewPanelReviewer(nil, []string{}, 50)
	_, err := p.Review(context.Background(), "prompt", t.TempDir(), "", "", nil)
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "no reviewer models configured") {
		t.Errorf("Review should delegate to ReviewPanel, got: %v", err)
	}
}

func TestPanelReviewerReviewPanelCancelledContext(t *testing.T) {
	workDir := t.TempDir()
	os.WriteFile(filepath.Join(workDir, "main.go"), []byte("package main"), 0644)
	p := NewPanelReviewer(nil, []string{"model-a", "model-b"}, 50)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, _, err := p.ReviewPanel(ctx, "prompt", workDir, "", "", nil)
	if err == nil {
		t.Fatal("expected error for cancelled context")
	}
	if !strings.Contains(err.Error(), "all reviewers failed") {
		t.Errorf("error = %q, want 'all reviewers failed'", err.Error())
	}
}

// ---------------------------------------------------------------------------
// Additional parseReviewResponse edge cases
// ---------------------------------------------------------------------------

func TestParseReviewResponseEdgeCases(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		check   func(t *testing.T, r *ReviewResult)
	}{
		{
			name:  "extra fields ignored",
			input: `{"scores":{"criteria":[{"name":"A","passed":true}]},"overall_score":1,"max_score":1,"summary":"ok","issues":[],"strengths":[],"extra_field":"ignored"}`,
			check: func(t *testing.T, r *ReviewResult) {
				if r.OverallScore != 1 {
					t.Errorf("OverallScore = %d, want 1", r.OverallScore)
				}
			},
		},
		{
			name:  "criteria with empty reason",
			input: `{"scores":{"criteria":[{"name":"Build","passed":true,"reason":""}]},"overall_score":1,"max_score":1,"summary":"good","issues":[],"strengths":[]}`,
			check: func(t *testing.T, r *ReviewResult) {
				if r.Scores.Criteria[0].Reason != "" {
					t.Errorf("expected empty reason, got %q", r.Scores.Criteria[0].Reason)
				}
			},
		},
		{
			name:    "truncated json",
			input:   `{"scores":{"criteria":[{"name":"A","pas`,
			wantErr: true,
		},
		{
			name:  "explicit max_score preserved",
			input: `{"scores":{"criteria":[{"name":"A","passed":true}]},"overall_score":1,"max_score":5,"summary":"test","issues":[],"strengths":[]}`,
			check: func(t *testing.T, r *ReviewResult) {
				if r.MaxScore != 5 {
					t.Errorf("MaxScore = %d, want 5", r.MaxScore)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseReviewResponse(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if tt.check != nil {
				tt.check(t, result)
			}
		})
	}
}

// ---------------------------------------------------------------------------
// Additional averageReview edge cases
// ---------------------------------------------------------------------------

func TestAverageReviewAllEmptyCriteria(t *testing.T) {
	panel := []ReviewResult{
		{Scores: ReviewScores{Criteria: nil}},
		{Scores: ReviewScores{Criteria: nil}},
	}
	result := averageReview(panel, nil)
	if len(result.Scores.Criteria) != 0 {
		t.Errorf("expected 0 criteria, got %d", len(result.Scores.Criteria))
	}
	if result.OverallScore != 0 {
		t.Errorf("OverallScore = %d, want 0", result.OverallScore)
	}
}

func TestAverageReviewNilIssuesAndStrengths(t *testing.T) {
	panel := []ReviewResult{
		{
			Scores:    ReviewScores{Criteria: []CriterionResult{{Name: "A", Passed: true}}},
			Issues:    nil,
			Strengths: nil,
		},
	}
	result := averageReview(panel, nil)
	if result.OverallScore != 1 {
		t.Errorf("OverallScore = %d, want 1", result.OverallScore)
	}
}

// ---------------------------------------------------------------------------
// Additional copyDirToTemp edge cases
// ---------------------------------------------------------------------------

func TestCopyDirToTempWithNestedDirs(t *testing.T) {
	src := t.TempDir()
	nested := filepath.Join(src, "a", "b", "c")
	os.MkdirAll(nested, 0755)
	os.WriteFile(filepath.Join(nested, "deep.txt"), []byte("deep"), 0644)

	dst, err := copyDirToTemp(src, "hyoka-test-*")
	if err != nil {
		t.Fatalf("copyDirToTemp failed: %v", err)
	}
	defer os.RemoveAll(dst)

	data, err := os.ReadFile(filepath.Join(dst, "a", "b", "c", "deep.txt"))
	if err != nil {
		t.Fatalf("failed to read deep file: %v", err)
	}
	if string(data) != "deep" {
		t.Errorf("content = %q, want %q", string(data), "deep")
	}
}

func TestCopyDirToTempNonexistentSource(t *testing.T) {
	_, err := copyDirToTemp("/nonexistent/source/dir", "hyoka-test-*")
	if err == nil {
		t.Fatal("expected error for nonexistent source")
	}
}

// ---------------------------------------------------------------------------
// ReviewResult JSON round-trip
// ---------------------------------------------------------------------------

func TestReviewResultJSONMarshalRoundTrip(t *testing.T) {
	original := ReviewResult{
		Model: "test-model",
		Scores: ReviewScores{Criteria: []CriterionResult{
			{Name: "Build", Passed: true, Reason: "compiles"},
			{Name: "Style", Passed: false, Reason: "needs work"},
		}},
		OverallScore: 1,
		MaxScore:     2,
		Summary:      "Mixed results",
		Issues:       []string{"style issue"},
		Strengths:    []string{"builds"},
		Events:       []ReviewEvent{{Type: "message", Content: "hello"}},
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var decoded ReviewResult
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if decoded.Model != original.Model {
		t.Errorf("Model = %q, want %q", decoded.Model, original.Model)
	}
	if decoded.OverallScore != original.OverallScore {
		t.Errorf("OverallScore = %d, want %d", decoded.OverallScore, original.OverallScore)
	}
	if len(decoded.Scores.Criteria) != len(original.Scores.Criteria) {
		t.Errorf("Criteria count = %d, want %d", len(decoded.Scores.Criteria), len(original.Scores.Criteria))
	}
	if len(decoded.Events) != 1 {
		t.Errorf("Events count = %d, want 1", len(decoded.Events))
	}
}

// ---------------------------------------------------------------------------
// BuildReviewPrompt additional edge cases
// ---------------------------------------------------------------------------

func TestBuildReviewPromptSpecialChars(t *testing.T) {
	prompt := "Write code with backticks and bold and special chars"
	generated := map[string]string{
		"test.go": "package main\nfunc main() { fmt.Println(\"hello\") }",
	}
	result := BuildReviewPrompt(prompt, generated, nil, nil, nil)
	if !strings.Contains(result, "backticks") {
		t.Error("should preserve special characters in prompt")
	}
}

// ---------------------------------------------------------------------------
// NewCopilotReviewer additional edge cases
// ---------------------------------------------------------------------------

func TestNewCopilotReviewerZeroMaxActions(t *testing.T) {
	r := NewCopilotReviewer(nil, "model", 0)
	if r.maxSessionActions != 0 {
		t.Errorf("maxSessionActions = %d, want 0", r.maxSessionActions)
	}
}

func TestNewPanelReviewerSingleModel(t *testing.T) {
	p := NewPanelReviewer(nil, []string{"only-model"}, 10)
	if len(p.Models()) != 1 {
		t.Fatalf("Models() count = %d, want 1", len(p.Models()))
	}
	if p.Models()[0] != "only-model" {
		t.Errorf("Models()[0] = %q, want %q", p.Models()[0], "only-model")
	}
}

func TestNewPanelReviewerNilModels(t *testing.T) {
	p := NewPanelReviewer(nil, nil, 10)
	if p.Models() != nil {
		t.Errorf("expected nil Models(), got %v", p.Models())
	}
}

// ---------------------------------------------------------------------------
// PanelReviewer setter edge cases
// ---------------------------------------------------------------------------

func TestPanelReviewerSetSkillDirectoriesNil(t *testing.T) {
	p := NewPanelReviewer(nil, []string{"m"}, 10)
	p.SetSkillDirectories(nil)
	if p.skillDirectories != nil {
		t.Error("expected nil skillDirectories")
	}
}

func TestPanelReviewerSetSessionTimeoutZero(t *testing.T) {
	p := NewPanelReviewer(nil, []string{"m"}, 10)
	p.SetSessionTimeout(0)
	if p.sessionTimeout != 0 {
		t.Errorf("expected zero timeout, got %v", p.sessionTimeout)
	}
}

func TestPanelReviewerReviewPanelWithReferenceDir(t *testing.T) {
	workDir := t.TempDir()
	os.WriteFile(filepath.Join(workDir, "main.go"), []byte("package main"), 0644)

	refDir := t.TempDir()
	os.WriteFile(filepath.Join(refDir, "ref.go"), []byte("package ref"), 0644)

	p := NewPanelReviewer(nil, []string{"model-a"}, 50)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, _, err := p.ReviewPanel(ctx, "prompt", workDir, refDir, "some criteria", nil)
	if err == nil {
		t.Fatal("expected error (cancelled context)")
	}
}

func TestPanelReviewerReviewPanelWithInvalidReferenceDir(t *testing.T) {
	workDir := t.TempDir()
	os.WriteFile(filepath.Join(workDir, "main.go"), []byte("package main"), 0644)

	p := NewPanelReviewer(nil, []string{"model-a"}, 50)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	// Non-fatal: reference read failure should not prevent the run
	_, _, err := p.ReviewPanel(ctx, "prompt", workDir, "/nonexistent/ref/dir", "criteria", nil)
	if err == nil {
		t.Fatal("expected error (cancelled context, not ref failure)")
	}
	// Should fail with "all reviewers failed" not reference error
	if !strings.Contains(err.Error(), "all reviewers failed") {
		t.Errorf("error = %q, want 'all reviewers failed'", err.Error())
	}
}

func TestPanelReviewerReviewPanelWithEmptyReferenceDir(t *testing.T) {
	workDir := t.TempDir()
	os.WriteFile(filepath.Join(workDir, "main.go"), []byte("package main"), 0644)

	refDir := t.TempDir() // empty reference dir

	p := NewPanelReviewer(nil, []string{"model-a"}, 50)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, _, err := p.ReviewPanel(ctx, "prompt", workDir, refDir, "", nil)
	if err == nil {
		t.Fatal("expected error (cancelled context)")
	}
}

func TestCopilotReviewerReviewWithOnlyHiddenFiles(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, ".gitignore"), []byte("*.pyc"), 0644)

	r := NewCopilotReviewer(nil, "test-model", 50)
	_, err := r.Review(context.Background(), "prompt", dir, "", "", nil)
	if err == nil {
		t.Fatal("expected error for dir with only hidden files")
	}
	if !strings.Contains(err.Error(), "no generated files") {
		t.Errorf("error = %q, want 'no generated files'", err.Error())
	}
}

func TestPanelReviewerReviewPanelWithOnlyHiddenFiles(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, ".gitignore"), []byte("*.pyc"), 0644)

	p := NewPanelReviewer(nil, []string{"model-a"}, 50)
	_, _, err := p.ReviewPanel(context.Background(), "prompt", dir, "", "", nil)
	if err == nil {
		t.Fatal("expected error for dir with only hidden files")
	}
	if !strings.Contains(err.Error(), "no generated files to review") {
		t.Errorf("error = %q, want 'no generated files to review'", err.Error())
	}
}

// ---------------------------------------------------------------------------
// eventCollector tests
// ---------------------------------------------------------------------------

func TestEventCollectorHandleAssistantMessage(t *testing.T) {
	cancelled := false
	cancel := func() { cancelled = true }
	c := newEventCollector("test-model", 100, cancel)

	content := "Hello, world!"
	c.handleEvent(copilot.SessionEvent{
		Type: copilot.SessionEventTypeAssistantMessage,
		Data: copilot.Data{Content: &content},
	})

	text, events := c.response()
	if text != "Hello, world!" {
		t.Errorf("assistantContent = %q, want %q", text, "Hello, world!")
	}
	if len(events) != 1 {
		t.Fatalf("events count = %d, want 1", len(events))
	}
	if events[0].Type != string(copilot.SessionEventTypeAssistantMessage) {
		t.Errorf("event type = %q", events[0].Type)
	}
	if events[0].Content != "Hello, world!" {
		t.Errorf("event content = %q", events[0].Content)
	}
	if cancelled {
		t.Error("should not cancel under action limit")
	}
}

func TestEventCollectorAccumulatesContent(t *testing.T) {
	c := newEventCollector("model", 100, func() {})

	part1 := "Hello, "
	part2 := "world!"
	c.handleEvent(copilot.SessionEvent{
		Type: copilot.SessionEventTypeAssistantMessage,
		Data: copilot.Data{Content: &part1},
	})
	c.handleEvent(copilot.SessionEvent{
		Type: copilot.SessionEventTypeAssistantMessage,
		Data: copilot.Data{Content: &part2},
	})

	text, events := c.response()
	if text != "Hello, world!" {
		t.Errorf("accumulated content = %q, want %q", text, "Hello, world!")
	}
	if len(events) != 2 {
		t.Errorf("events count = %d, want 2", len(events))
	}
}

func TestEventCollectorActionLimit(t *testing.T) {
	cancelled := false
	cancel := func() { cancelled = true }
	c := newEventCollector("model", 2, cancel)

	// Send 3 action events — limit is 2, so 3rd should trigger cancel
	for i := 0; i < 3; i++ {
		c.handleEvent(copilot.SessionEvent{
			Type: copilot.SessionEventTypeAssistantMessage,
			Data: copilot.Data{},
		})
	}

	if !cancelled {
		t.Error("expected cancel after exceeding action limit")
	}
	if !c.actionLimitHit {
		t.Error("expected actionLimitHit to be true")
	}
}

func TestEventCollectorNoLimitWhenZero(t *testing.T) {
	cancelled := false
	cancel := func() { cancelled = true }
	c := newEventCollector("model", 0, cancel)

	for i := 0; i < 100; i++ {
		c.handleEvent(copilot.SessionEvent{
			Type: copilot.SessionEventTypeAssistantMessage,
			Data: copilot.Data{},
		})
	}

	if cancelled {
		t.Error("should not cancel when maxActions is 0")
	}
}

func TestEventCollectorToolEvents(t *testing.T) {
	c := newEventCollector("model", 100, func() {})

	toolName := "read_file"
	args := map[string]string{"path": "main.go"}
	resultContent := "file content"
	dur := 42.5
	c.handleEvent(copilot.SessionEvent{
		Type: copilot.SessionEventTypeToolExecutionStart,
		Data: copilot.Data{
			ToolName:  &toolName,
			Arguments: args,
		},
	})
	c.handleEvent(copilot.SessionEvent{
		Type: copilot.SessionEventTypeToolExecutionComplete,
		Data: copilot.Data{
			ToolName: &toolName,
			Result:   &copilot.Result{Content: &resultContent},
			Duration: &dur,
		},
	})

	_, events := c.response()
	if len(events) != 2 {
		t.Fatalf("events count = %d, want 2", len(events))
	}
	if events[0].ToolName != "read_file" {
		t.Errorf("start event tool = %q", events[0].ToolName)
	}
	if events[0].ToolArgs == "" {
		t.Error("start event should have tool args")
	}
	if events[1].Result != "file content" {
		t.Errorf("complete event result = %q", events[1].Result)
	}
	if events[1].Duration != 42.5 {
		t.Errorf("complete event duration = %f", events[1].Duration)
	}
}

func TestEventCollectorErrorEvents(t *testing.T) {
	tests := []struct {
		name      string
		errorData *copilot.ErrorUnion
		wantError string
	}{
		{
			name: "error class",
			errorData: &copilot.ErrorUnion{
				ErrorClass: &copilot.ErrorClass{Message: "something broke"},
			},
			wantError: "something broke",
		},
		{
			name:      "error string",
			errorData: &copilot.ErrorUnion{String: strPtr("string error")},
			wantError: "string error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newEventCollector("model", 100, func() {})
			c.handleEvent(copilot.SessionEvent{
				Type: copilot.SessionEventTypeToolExecutionComplete,
				Data: copilot.Data{Error: tt.errorData},
			})

			_, events := c.response()
			if len(events) != 1 {
				t.Fatalf("events = %d, want 1", len(events))
			}
			if events[0].Error != tt.wantError {
				t.Errorf("error = %q, want %q", events[0].Error, tt.wantError)
			}
		})
	}
}

func TestEventCollectorTurnEvents(t *testing.T) {
	c := newEventCollector("model", 100, func() {})

	c.handleEvent(copilot.SessionEvent{
		Type: copilot.SessionEventTypeAssistantTurnStart,
		Data: copilot.Data{},
	})
	c.handleEvent(copilot.SessionEvent{
		Type: copilot.SessionEventTypeAssistantTurnEnd,
		Data: copilot.Data{},
	})

	_, events := c.response()
	if len(events) != 2 {
		t.Fatalf("events = %d, want 2", len(events))
	}
	if events[0].Type != string(copilot.SessionEventTypeAssistantTurnStart) {
		t.Errorf("first event type = %q", events[0].Type)
	}
	if events[1].Type != string(copilot.SessionEventTypeAssistantTurnEnd) {
		t.Errorf("second event type = %q", events[1].Type)
	}
}

func TestEventCollectorUsageEvent(t *testing.T) {
	c := newEventCollector("model", 100, func() {})
	c.handleEvent(copilot.SessionEvent{
		Type: copilot.SessionEventTypeAssistantUsage,
		Data: copilot.Data{},
	})

	_, events := c.response()
	if len(events) != 1 {
		t.Fatalf("events = %d, want 1", len(events))
	}
}

func TestEventCollectorReasoningCountsAsAction(t *testing.T) {
	cancelled := false
	c := newEventCollector("model", 1, func() { cancelled = true })

	c.handleEvent(copilot.SessionEvent{
		Type: copilot.SessionEventTypeAssistantReasoning,
		Data: copilot.Data{},
	})
	c.handleEvent(copilot.SessionEvent{
		Type: copilot.SessionEventTypeAssistantReasoning,
		Data: copilot.Data{},
	})

	if !cancelled {
		t.Error("reasoning events should count toward action limit")
	}
}

func TestEventCollectorToolStartCountsAsAction(t *testing.T) {
	cancelled := false
	c := newEventCollector("model", 1, func() { cancelled = true })

	c.handleEvent(copilot.SessionEvent{
		Type: copilot.SessionEventTypeToolExecutionStart,
		Data: copilot.Data{},
	})
	c.handleEvent(copilot.SessionEvent{
		Type: copilot.SessionEventTypeToolExecutionStart,
		Data: copilot.Data{},
	})

	if !cancelled {
		t.Error("tool start events should count toward action limit")
	}
}

func TestEventCollectorNilContentNotAccumulated(t *testing.T) {
	c := newEventCollector("model", 100, func() {})

	c.handleEvent(copilot.SessionEvent{
		Type: copilot.SessionEventTypeAssistantMessage,
		Data: copilot.Data{Content: nil},
	})

	text, _ := c.response()
	if text != "" {
		t.Errorf("expected empty content, got %q", text)
	}
}

func TestEventCollectorResponseCopiesEvents(t *testing.T) {
	c := newEventCollector("model", 100, func() {})
	content := "test"
	c.handleEvent(copilot.SessionEvent{
		Type: copilot.SessionEventTypeAssistantMessage,
		Data: copilot.Data{Content: &content},
	})

	_, events1 := c.response()
	_, events2 := c.response()

	// Verify they are separate slices
	if len(events1) != len(events2) {
		t.Error("response should return consistent results")
	}
}

// ---------------------------------------------------------------------------
// buildConsolidationPrompt tests
// ---------------------------------------------------------------------------

func TestBuildConsolidationPrompt(t *testing.T) {
	panel := []ReviewResult{
		{
			Model:        "model-a",
			OverallScore: 2,
			MaxScore:     3,
			Summary:      "Good overall",
			Scores: ReviewScores{Criteria: []CriterionResult{
				{Name: "Build", Passed: true},
				{Name: "Style", Passed: true},
				{Name: "Errors", Passed: false},
			}},
		},
		{
			Model:        "model-b",
			OverallScore: 1,
			MaxScore:     3,
			Summary:      "Needs work",
			Scores: ReviewScores{Criteria: []CriterionResult{
				{Name: "Build", Passed: true},
				{Name: "Style", Passed: false},
				{Name: "Errors", Passed: false},
			}},
		},
	}

	prompt := buildConsolidationPrompt("Write Azure code", panel)

	checks := []string{
		"senior review consolidator",
		"Original Prompt",
		"Write Azure code",
		"Individual Reviews",
		"Reviewer 1 (model-a)",
		"Reviewer 2 (model-b)",
		"Instructions",
		"consensus review",
		"majority",
		"JSON object",
	}

	for _, check := range checks {
		if !strings.Contains(prompt, check) {
			t.Errorf("consolidation prompt missing %q", check)
		}
	}
}

func TestBuildConsolidationPromptSingleReviewer(t *testing.T) {
	panel := []ReviewResult{{
		Model:   "model-a",
		Summary: "Test",
		Scores:  ReviewScores{Criteria: []CriterionResult{{Name: "A", Passed: true}}},
	}}

	prompt := buildConsolidationPrompt("prompt", panel)
	if !strings.Contains(prompt, "Reviewer 1 (model-a)") {
		t.Error("should contain reviewer label")
	}
}

func TestBuildConsolidationPromptEmpty(t *testing.T) {
	prompt := buildConsolidationPrompt("prompt", nil)
	if !strings.Contains(prompt, "Original Prompt") {
		t.Error("should contain prompt section even with empty panel")
	}
	if !strings.Contains(prompt, "Individual Reviews") {
		t.Error("should contain reviews section even with empty panel")
	}
}

// helper
func strPtr(s string) *string { return &s }

// ---------------------------------------------------------------------------
// ReviewerResponse / new JSON schema tests (#343)
// ---------------------------------------------------------------------------

func TestParseReviewResponseNewSchema(t *testing.T) {
	input := `{"criteria":[{"criterion":"Uses DefaultAzureCredential","passed":true,"reasoning":"Correctly imports and uses DefaultAzureCredential"},{"criterion":"Handles errors","passed":false,"reasoning":"Missing error handling for auth failures"}],"summary":"Partial pass","issues":["No error handling"],"strengths":["Correct auth"]}`

	result, err := parseReviewResponse(input)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Scores.Criteria) != 2 {
		t.Fatalf("expected 2 criteria, got %d", len(result.Scores.Criteria))
	}
	if result.Scores.Criteria[0].Name != "Uses DefaultAzureCredential" {
		t.Errorf("criterion[0].Name = %q, want %q", result.Scores.Criteria[0].Name, "Uses DefaultAzureCredential")
	}
	if !result.Scores.Criteria[0].Passed {
		t.Error("criterion[0] should pass")
	}
	if result.Scores.Criteria[1].Passed {
		t.Error("criterion[1] should fail")
	}
	if result.OverallScore != 1 {
		t.Errorf("OverallScore = %d, want 1", result.OverallScore)
	}
	if result.MaxScore != 2 {
		t.Errorf("MaxScore = %d, want 2", result.MaxScore)
	}
}

func TestParseReviewResponseNewSchemaInMarkdown(t *testing.T) {
	input := "```json\n" + `{"criteria":[{"criterion":"Build","passed":true,"reasoning":"OK"}],"summary":"Good","issues":[],"strengths":["Clean"]}` + "\n```"
	result, err := parseReviewResponse(input)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Scores.Criteria) != 1 {
		t.Fatalf("expected 1 criterion, got %d", len(result.Scores.Criteria))
	}
	if result.Scores.Criteria[0].Reason != "OK" {
		t.Errorf("reason = %q, want %q", result.Scores.Criteria[0].Reason, "OK")
	}
}

func TestValidateReviewerResponseValid(t *testing.T) {
	result := &ReviewResult{
		Scores: ReviewScores{Criteria: []CriterionResult{
			{Name: "A", Passed: true},
			{Name: "B", Passed: false, Reason: "missing"},
		}},
	}
	errs := validateReviewerResponse(result)
	if len(errs) > 0 {
		t.Errorf("expected no errors, got %v", errs)
	}
}

func TestValidateReviewerResponseNoCriteria(t *testing.T) {
	result := &ReviewResult{Scores: ReviewScores{}}
	errs := validateReviewerResponse(result)
	if len(errs) == 0 {
		t.Error("expected error for missing criteria")
	}
}

func TestValidateReviewerResponseEmptyName(t *testing.T) {
	result := &ReviewResult{
		Scores: ReviewScores{Criteria: []CriterionResult{
			{Name: "", Passed: true},
		}},
	}
	errs := validateReviewerResponse(result)
	if len(errs) == 0 {
		t.Error("expected error for empty criterion name")
	}
}

func TestValidateReviewerResponseNil(t *testing.T) {
	errs := validateReviewerResponse(nil)
	if len(errs) == 0 {
		t.Error("expected error for nil result")
	}
}

func TestDeterministicVoteStrictFailure(t *testing.T) {
	panel := []ReviewResult{
		{
			Model: "m1",
			Scores: ReviewScores{Criteria: []CriterionResult{
				{Name: "Auth", Passed: true},
				{Name: "Build", Passed: true},
			}},
		},
		{
			Model: "m2",
			Scores: ReviewScores{Criteria: []CriterionResult{
				{Name: "Auth", Passed: true},
				{Name: "Build", Passed: false, Reason: "compile error"},
			}},
		},
	}

	result := deterministicVote(panel, nil)

	criteriaMap := map[string]bool{}
	for _, c := range result.Scores.Criteria {
		criteriaMap[c.Name] = c.Passed
	}
	// Auth: both pass → pass
	if !criteriaMap["Auth"] {
		t.Error("Auth should pass (0 failures)")
	}
	// Build: one fail → fail (any-fail voting)
	if criteriaMap["Build"] {
		t.Error("Build should fail (any-fail voting)")
	}
	if result.OverallScore != 1 {
		t.Errorf("OverallScore = %d, want 1", result.OverallScore)
	}
}

func TestCriterionJudgmentJSONRoundTrip(t *testing.T) {
	resp := ReviewerResponse{
		Criteria: []CriterionJudgment{
			{Criterion: "test criterion", Passed: true, Reasoning: "looks good"},
		},
	}
	data, err := json.Marshal(resp)
	if err != nil {
		t.Fatal(err)
	}
	var decoded ReviewerResponse
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if len(decoded.Criteria) != 1 {
		t.Fatalf("expected 1 criterion, got %d", len(decoded.Criteria))
	}
	if decoded.Criteria[0].Criterion != "test criterion" {
		t.Error("criterion text should round-trip")
	}
}

// ---------------------------------------------------------------------------
// GeneratorArtifact integration tests
// ---------------------------------------------------------------------------

// TestBuildReviewPrompt_WithArtifactAndFiles verifies that when both generated
// files AND a final response artifact are present, the prompt includes BOTH
// sections unconditionally.
func TestBuildReviewPrompt_WithArtifactAndFiles(t *testing.T) {
prompt := "Write a Python script"
files := map[string]string{
"main.py": "print('hello')",
}
artifact := &GeneratorArtifact{
FinalResponse: "I have created the script as requested.",
}

result := BuildReviewPrompt(prompt, files, nil, criteriaStringToChecks("Must run without errors"), artifact)

// Both sections must appear
if !strings.Contains(result, "## Generated Code") {
t.Error("prompt must include Generated Code section when files present")
}
if !strings.Contains(result, "main.py") {
t.Error("prompt must include file content")
}
if !strings.Contains(result, "## Agent's Final Response") {
t.Error("prompt must include Agent's Final Response section when artifact present")
}
if !strings.Contains(result, "I have created the script") {
t.Error("prompt must include artifact's final response text")
}
}

// TestBuildReviewPrompt_WithArtifactNoFiles verifies that when no files are
// generated but an artifact with a final response exists, the prompt includes
// the agent's response and indicates no files were created.
func TestBuildReviewPrompt_WithArtifactNoFiles(t *testing.T) {
prompt := "Explain how to use DefaultAzureCredential"
artifact := &GeneratorArtifact{
FinalResponse: "DefaultAzureCredential is a chained credential...",
}

result := BuildReviewPrompt(prompt, nil, nil, criteriaStringToChecks("Must be accurate"), artifact)

if !strings.Contains(result, "## Generated Code") {
t.Error("prompt must include Generated Code section header")
}
if !strings.Contains(result, "No files were created") {
t.Error("prompt must indicate no files when workspace is empty")
}
if !strings.Contains(result, "## Agent's Final Response") {
t.Error("prompt must include Agent's Final Response section")
}
if !strings.Contains(result, "DefaultAzureCredential is a chained") {
t.Error("prompt must include artifact response")
}
}

// TestBuildReviewPrompt_NoArtifactWithFiles verifies that when files are
// generated but no artifact is provided, the prompt still works (legacy behavior).
func TestBuildReviewPrompt_NoArtifactWithFiles(t *testing.T) {
prompt := "Write code"
files := map[string]string{"test.py": "code"}

result := BuildReviewPrompt(prompt, files, nil, nil, nil)

if !strings.Contains(result, "## Generated Code") {
t.Error("prompt must include Generated Code section")
}
if !strings.Contains(result, "test.py") {
t.Error("prompt must include file content")
}
if strings.Contains(result, "## Agent's Final Response") {
t.Error("prompt should not include Agent's Final Response when artifact is nil")
}
}

// TestBuildReviewPrompt_EmptyArtifactResponse verifies that an artifact with
// an empty FinalResponse field does not add the Agent's Final Response section.
func TestBuildReviewPrompt_EmptyArtifactResponse(t *testing.T) {
prompt := "Write code"
files := map[string]string{"test.py": "code"}
artifact := &GeneratorArtifact{
FinalResponse: "", // empty
}

result := BuildReviewPrompt(prompt, files, nil, nil, artifact)

if strings.Contains(result, "## Agent's Final Response") {
t.Error("prompt should not include Agent's Final Response when FinalResponse is empty")
}
}

// TestCopilotReviewer_EmptyWorkspaceWithArtifact verifies that the reviewer
// accepts an empty workspace when a non-nil artifact with a response is provided.
func TestCopilotReviewer_EmptyWorkspaceWithArtifact(t *testing.T) {
	artifact := &GeneratorArtifact{
		FinalResponse: "Here is my explanation of how to use the SDK...",
	}

	// We can't actually invoke the real reviewer without Copilot, but we can
	// test that BuildReviewPrompt doesn't error and includes the response
	prompt := BuildReviewPrompt("Explain Azure SDK", map[string]string{}, nil, criteriaStringToChecks("Must be clear"), artifact)

	if !strings.Contains(prompt, "Here is my explanation") {
		t.Error("prompt must include artifact response when no files present")
	}
}

// TestCopilotReviewer_EmptyWorkspaceNoArtifact verifies that the reviewer
// errors when BOTH workspace is empty AND no artifact is provided (nothing to review).
func TestCopilotReviewer_EmptyWorkspaceNoArtifact(t *testing.T) {
emptyDir := t.TempDir()
r := NewCopilotReviewer(nil, "test-model", 50)

// This should error because there's nothing to review
_, err := r.Review(context.Background(), "prompt", emptyDir, "", "criteria", nil)
if err == nil {
t.Fatal("expected error when workspace is empty and no artifact provided")
}
if !strings.Contains(err.Error(), "no generated files") && !strings.Contains(err.Error(), "no agent response") {
t.Errorf("error should mention missing files or response, got: %v", err)
}
}
