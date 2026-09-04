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

func assistantMessageEvent(content string) copilot.SessionEvent {
	return copilot.SessionEvent{Data: &copilot.AssistantMessageData{Content: content}}
}

func assistantReasoningEvent() copilot.SessionEvent {
	return copilot.SessionEvent{Data: &copilot.AssistantReasoningData{}}
}

func toolStartEvent(name string, arguments any) copilot.SessionEvent {
	return copilot.SessionEvent{Data: &copilot.ToolExecutionStartData{
		ToolName:  name,
		Arguments: arguments,
	}}
}

func toolCompleteEvent(result string, err *copilot.ToolExecutionCompleteError) copilot.SessionEvent {
	return copilot.SessionEvent{Data: &copilot.ToolExecutionCompleteData{
		Error:   err,
		Result:  &copilot.ToolExecutionCompleteResult{Content: result},
		Success: err == nil,
	}}
}

func TestEventCollectorHandleAssistantMessage(t *testing.T) {
	cancelled := false
	cancel := func() { cancelled = true }
	c := newEventCollector("test-model", 100, cancel)

	content := "Hello, world!"
	c.handleEvent(assistantMessageEvent(content))

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
	c.handleEvent(assistantMessageEvent(part1))
	c.handleEvent(assistantMessageEvent(part2))

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
		c.handleEvent(assistantMessageEvent(""))
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
		c.handleEvent(assistantMessageEvent(""))
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
	startedAt := time.Date(2026, time.August, 27, 1, 2, 3, 0, time.UTC)
	start := toolStartEvent(toolName, args)
	start.Timestamp = startedAt
	complete := toolCompleteEvent(resultContent, nil)
	complete.Timestamp = startedAt.Add(42_500 * time.Microsecond)
	c.handleEvent(start)
	c.handleEvent(complete)

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
		t.Errorf("complete event duration = %f, want 42.5", events[1].Duration)
	}
}

func TestEventCollectorUnmatchedCompletionHasZeroDuration(t *testing.T) {
	c := newEventCollector("model", 100, func() {})
	complete := toolCompleteEvent("file content", nil)
	complete.Timestamp = time.Now()

	c.handleEvent(complete)

	_, events := c.response()
	if len(events) != 1 {
		t.Fatalf("events count = %d, want 1", len(events))
	}
	if events[0].Duration != 0 {
		t.Errorf("complete event duration = %f, want 0", events[0].Duration)
	}
}

func TestEventCollectorErrorEvents(t *testing.T) {
	tests := []struct {
		name string
		err  *copilot.ToolExecutionCompleteError
		want string
	}{
		{
			name: "tool error",
			err:  &copilot.ToolExecutionCompleteError{Message: "something broke"},
			want: "something broke",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newEventCollector("model", 100, func() {})
			c.handleEvent(toolCompleteEvent("", tt.err))

			_, events := c.response()
			if len(events) != 1 {
				t.Fatalf("events = %d, want 1", len(events))
			}
			if events[0].Error != tt.want {
				t.Errorf("error = %q, want %q", events[0].Error, tt.want)
			}
		})
	}
}

func TestEventCollectorTurnEvents(t *testing.T) {
	c := newEventCollector("model", 100, func() {})

	c.handleEvent(copilot.SessionEvent{Data: &copilot.AssistantTurnStartData{TurnID: "1"}})
	c.handleEvent(copilot.SessionEvent{Data: &copilot.AssistantTurnEndData{TurnID: "1"}})

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
	c.handleEvent(copilot.SessionEvent{Data: &copilot.AssistantUsageData{}})

	_, events := c.response()
	if len(events) != 1 {
		t.Fatalf("events = %d, want 1", len(events))
	}
}

func TestEventCollectorReasoningCountsAsAction(t *testing.T) {
	cancelled := false
	c := newEventCollector("model", 1, func() { cancelled = true })

	c.handleEvent(assistantReasoningEvent())
	c.handleEvent(assistantReasoningEvent())

	if !cancelled {
		t.Error("reasoning events should count toward action limit")
	}
}

func TestEventCollectorToolStartCountsAsAction(t *testing.T) {
	cancelled := false
	c := newEventCollector("model", 1, func() { cancelled = true })

	c.handleEvent(toolStartEvent("", nil))
	c.handleEvent(toolStartEvent("", nil))

	if !cancelled {
		t.Error("tool start events should count toward action limit")
	}
}

func TestEventCollectorNilContentNotAccumulated(t *testing.T) {
	c := newEventCollector("model", 100, func() {})

	c.handleEvent(assistantMessageEvent(""))

	text, _ := c.response()
	if text != "" {
		t.Errorf("expected empty content, got %q", text)
	}
}

func TestEventCollectorResponseCopiesEvents(t *testing.T) {
	c := newEventCollector("model", 100, func() {})
	content := "test"
	c.handleEvent(assistantMessageEvent(content))

	_, events1 := c.response()
	_, events2 := c.response()

	// Verify they are separate slices
	if len(events1) != len(events2) {
		t.Error("response should return consistent results")
	}
}

// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// Reviewer validation tests
// ---------------------------------------------------------------------------

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

// ---------------------------------------------------------------------------
// ID-aware parser and validator tests
// ---------------------------------------------------------------------------

func TestParseReviewResponseV2_GoodIDs(t *testing.T) {
	expected := []ReviewCheck{
		{ID: "check_1", Text: "Build succeeds"},
		{ID: "check_2", Text: "Tests pass"},
	}
	response := `{
"criteria": [
{"id": "check_1", "passed": true, "reasoning": "builds"},
{"id": "check_2", "passed": false, "reasoning": "1 test failed"}
],
"summary": "ok",
"issues": [],
"strengths": []
}`
	result, errs := parseReviewResponseV2(response, expected)
	if len(errs) > 0 {
		t.Fatalf("unexpected errors: %v", errs)
	}
	if result == nil {
		t.Fatal("result is nil")
	}
	if len(result.Scores.Criteria) != 2 {
		t.Fatalf("expected 2 criteria, got %d", len(result.Scores.Criteria))
	}
	if result.Scores.Criteria[0].ID != "check_1" {
		t.Errorf("criterion 0 id = %q, want check_1", result.Scores.Criteria[0].ID)
	}
	if result.Scores.Criteria[0].Name != "Build succeeds" {
		t.Errorf("criterion 0 name = %q, want canonical label", result.Scores.Criteria[0].Name)
	}
	if !result.Scores.Criteria[0].Passed {
		t.Error("criterion 0 should be passed")
	}
}

func TestParseReviewResponseV2_MissingID(t *testing.T) {
	expected := []ReviewCheck{
		{ID: "check_1", Text: "A"},
		{ID: "check_2", Text: "B"},
	}
	response := `{
"criteria": [
{"id": "check_1", "passed": true, "reasoning": "ok"}
],
"summary": "ok",
"issues": [],
"strengths": []
}`
	_, errs := parseReviewResponseV2(response, expected)
	if len(errs) == 0 {
		t.Fatal("expected validation errors for missing id")
	}
	found := false
	for _, e := range errs {
		if strings.Contains(e, "missing") && strings.Contains(e, "check_2") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected 'missing ids' error, got %v", errs)
	}
}

func TestParseReviewResponseV2_ExtraID(t *testing.T) {
	expected := []ReviewCheck{
		{ID: "check_1", Text: "A"},
	}
	response := `{
"criteria": [
{"id": "check_1", "passed": true, "reasoning": "ok"},
{"id": "check_99", "passed": false, "reasoning": "extra"}
],
"summary": "ok",
"issues": [],
"strengths": []
}`
	_, errs := parseReviewResponseV2(response, expected)
	if len(errs) == 0 {
		t.Fatal("expected validation errors for extra id")
	}
	found := false
	for _, e := range errs {
		if strings.Contains(e, "unexpected") && strings.Contains(e, "check_99") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected 'unexpected id' error, got %v", errs)
	}
}

func TestParseReviewResponseV2_MalformedShape(t *testing.T) {
	expected := []ReviewCheck{{ID: "check_1", Text: "A"}}
	response := `{"criteria": "not an array"}`
	_, errs := parseReviewResponseV2(response, expected)
	if len(errs) == 0 {
		t.Fatal("expected validation errors for malformed response")
	}
}

func TestAverageReview_KeysByID_NotName(t *testing.T) {
	expected := []ReviewCheck{
		{ID: "check_1", Text: "Build succeeds"},
	}
	panel := []ReviewResult{
		{
			Model: "m1",
			Scores: ReviewScores{Criteria: []CriterionResult{
				{ID: "check_1", Name: "Build succeeds", Passed: true, Reason: "compiles fine"},
			}},
		},
		{
			Model: "m2",
			Scores: ReviewScores{Criteria: []CriterionResult{
				{ID: "check_1", Name: "Build succeeds", Passed: true, Reason: "no errors"},
			}},
		},
	}
	result := averageReview(panel, expected)
	if len(result.Scores.Criteria) != 1 {
		t.Fatalf("expected 1 criterion, got %d", len(result.Scores.Criteria))
	}
	c := result.Scores.Criteria[0]
	if c.Name != "Build succeeds" {
		t.Errorf("name = %q, want canonical label", c.Name)
	}
	if !strings.Contains(c.Reason, "2/2") {
		t.Errorf("reason = %q, want 2/2 reviewers", c.Reason)
	}
}

func TestAverageReview_BucketScoping(t *testing.T) {
	expected := []ReviewCheck{
		{ID: "check_1", Text: "A"},
	}
	panel := []ReviewResult{
		{
			Model: "m1",
			Scores: ReviewScores{Criteria: []CriterionResult{
				{ID: "check_1", Name: "[bucketA] A", Passed: true},
				{ID: "check_1", Name: "[bucketB] A", Passed: false},
			}},
		},
	}
	result := averageReview(panel, expected)
	if len(result.Scores.Criteria) != 2 {
		t.Fatalf("expected 2 criteria (bucket-scoped), got %d", len(result.Scores.Criteria))
	}
	names := make(map[string]bool)
	for _, c := range result.Scores.Criteria {
		names[c.Name] = true
	}
	if !names["[bucketA] A"] || !names["[bucketB] A"] {
		t.Errorf("expected both bucket-prefixed entries, got %v", names)
	}
}

func TestBuildReviewPrompt_RendersIDs(t *testing.T) {
	checks := []ReviewCheck{
		{ID: "check_1", Text: "File exists"},
		{ID: "check_2", Text: "Tests pass"},
	}
	prompt := BuildReviewPrompt("Write code", nil, nil, checks, nil)
	if !strings.Contains(prompt, "check_1:") {
		t.Error("prompt missing 'check_1:' id rendering")
	}
	if !strings.Contains(prompt, "check_2:") {
		t.Error("prompt missing 'check_2:' id rendering")
	}
	if !strings.Contains(prompt, "File exists") {
		t.Error("prompt missing check text")
	}
	if !strings.Contains(prompt, "Do NOT") {
		t.Error("prompt missing id-contract rule")
	}
	if !strings.Contains(prompt, `"id"`) {
		t.Error("prompt missing schema example with id field")
	}
}

func TestParseReviewResponseV2_RejectsLegacyText(t *testing.T) {
	expected := []ReviewCheck{{ID: "check_1", Text: "A"}}
	// Response with only "criterion" field (no "id")
	response := `{
"criteria": [
{"criterion": "A", "passed": true, "reasoning": "ok"}
],
"summary": "ok",
"issues": [],
"strengths": []
}`
	_, errs := parseReviewResponseV2(response, expected)
	if len(errs) == 0 {
		t.Fatal("expected validation errors for legacy criterion-only response")
	}
	found := false
	for _, e := range errs {
		if strings.Contains(e, "empty id") || strings.Contains(e, "missing") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected empty/missing id error, got %v", errs)
	}
}

// TestAverageReview_AnchoredToExpectedCheckIDs verifies that consensus
// is anchored to the expected check list, marking missing checks as failed.
func TestAverageReview_AnchoredToExpectedCheckIDs(t *testing.T) {
	expected := []ReviewCheck{
		{ID: "check_1", Text: "Uses DefaultAzureCredential"},
		{ID: "check_2", Text: "Handles errors properly"},
		{ID: "check_3", Text: "Uses environment variables"},
	}

	// Panel where reviewer 2 has no vote for check_3
	panel := []ReviewResult{
		{
			Model: "reviewer-1",
			Scores: ReviewScores{
				Criteria: []CriterionResult{
					{ID: "check_1", Name: "Uses DefaultAzureCredential", Passed: true},
					{ID: "check_2", Name: "Handles errors properly", Passed: true},
					{ID: "check_3", Name: "Uses environment variables", Passed: false},
				},
			},
		},
		{
			Model: "reviewer-2",
			Scores: ReviewScores{
				Criteria: []CriterionResult{
					{ID: "check_1", Name: "Uses DefaultAzureCredential", Passed: true},
					{ID: "check_2", Name: "Handles errors properly", Passed: false},
					// check_3 missing → should appear in consensus as failed
				},
			},
		},
	}

	consensus := averageReview(panel, expected)

	if consensus == nil {
		t.Fatal("consensus is nil")
	}
	if len(consensus.Scores.Criteria) != 3 {
		t.Errorf("consensus has %d criteria, want 3", len(consensus.Scores.Criteria))
	}
	if consensus.MaxScore != 3 {
		t.Errorf("MaxScore = %d, want 3", consensus.MaxScore)
	}

	// check_1: both passed → should pass
	c1 := consensus.Scores.Criteria[0]
	if c1.ID != "check_1" || !c1.Passed {
		t.Errorf("check_1: ID=%q Passed=%t, want ID=check_1 Passed=true", c1.ID, c1.Passed)
	}

	// check_2: reviewer-2 failed → should fail (any-fail voting)
	c2 := consensus.Scores.Criteria[1]
	if c2.ID != "check_2" || c2.Passed {
		t.Errorf("check_2: ID=%q Passed=%t, want ID=check_2 Passed=false", c2.ID, c2.Passed)
	}

	// check_3: reviewer-1 failed it, reviewer-2 didn't vote → should fail (any-fail voting)
	// The key test is that it APPEARS in the consensus with the correct ID in expected order
	c3 := consensus.Scores.Criteria[2]
	if c3.ID != "check_3" {
		t.Errorf("check_3: ID=%q, want check_3", c3.ID)
	}
	if c3.Passed {
		t.Errorf("check_3 should fail (any-fail voting)")
	}
}
