package eval

import (
	"context"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria/graders"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/progress"
)

// fakeGrader is a test double implementing graders.Grader. It returns a
// canned GraderResult — useful for exercising the hook wiring without
// spinning up a real file/program/prompt_review grader.
type fakeGrader struct {
	kind    string
	name    string
	result  graders.GraderResult
	gradeFn func() (graders.GraderResult, error)
}

func (f *fakeGrader) Kind() string { return f.kind }
func (f *fakeGrader) Name() string { return f.name }
func (f *fakeGrader) Grade(_ context.Context, _ graders.GraderInput) (graders.GraderResult, error) {
	if f.gradeFn != nil {
		return f.gradeFn()
	}
	return f.result, nil
}

func TestEmitGraderStart_PopulatesIDAndKind(t *testing.T) {
	var r reporter
	g := &fakeGrader{kind: graders.KindFile, name: "my-file-check"}
	emitGraderStart(r.emit, g, "", "")
	if len(r.events) != 1 {
		t.Fatalf("want 1 event, got %d", len(r.events))
	}
	got := r.events[0]
	if got.Type != progress.EventGraderStart {
		t.Errorf("Type = %v, want EventGraderStart", got.Type)
	}
	if got.GraderID != "my-file-check" {
		t.Errorf("GraderID = %q, want my-file-check", got.GraderID)
	}
	if got.GraderKind != graders.KindFile {
		t.Errorf("GraderKind = %q, want %q", got.GraderKind, graders.KindFile)
	}
}

func TestEmitGraderStart_NilSenderNoPanic(t *testing.T) {
	g := &fakeGrader{kind: graders.KindFile, name: "x"}
	// Must not panic on nil sender or nil grader.
	emitGraderStart(nil, g, "", "")
	emitGraderStart(func(progress.ProgressEvent) { t.Fatal("called with nil grader") }, nil, "", "")
}

func TestEmitGraderComplete_ScorePolicy(t *testing.T) {
	cases := []struct {
		name        string
		kind        string
		pass        bool
		score       float64
		message     string
		wantResult  string
		wantScorePtr bool
	}{
		{"file pass", graders.KindFile, true, 1.0, "main.py present", progress.GraderResultPass, false},
		{"file fail", graders.KindFile, false, 0, "main.py missing", progress.GraderResultFail, false},
		{"program fail", graders.KindProgram, false, 0, "build failed", progress.GraderResultFail, false},
		{"behavior pass", graders.KindBehavior, true, 1.0, "ok", progress.GraderResultPass, false},
		{"action_sequence pass", graders.KindActionSequence, true, 1.0, "seq ok", progress.GraderResultPass, false},
		{"tool_constraint fail", graders.KindToolConstraint, false, 0, "blocked tool", progress.GraderResultFail, false},
		{"output_check fail", graders.KindOutputCheck, false, 0, "too few files", progress.GraderResultFail, false},
		// Only prompt / prompt_review populate Score:
		{"prompt_review pass", graders.KindPromptReview, true, 0.85, "panel 8.5/10", progress.GraderResultPass, true},
		{"prompt pass", graders.KindPrompt, true, 0.5, "judge 5/10", progress.GraderResultPass, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var r reporter
			g := &fakeGrader{kind: tc.kind, name: "g"}
			res := graders.GraderResult{
				Kind:    tc.kind,
				Name:    "g",
				Pass:    tc.pass,
				Score:   tc.score,
				Message: tc.message,
			}
			emitGraderComplete(r.emit, g, res)
			if len(r.events) != 1 {
				t.Fatalf("want 1 event, got %d", len(r.events))
			}
			got := r.events[0]
			if got.Type != progress.EventGraderComplete {
				t.Errorf("Type = %v, want EventGraderComplete", got.Type)
			}
			if got.Result != tc.wantResult {
				t.Errorf("Result = %q, want %q", got.Result, tc.wantResult)
			}
			if got.GraderID != "g" || got.GraderKind != tc.kind {
				t.Errorf("ID/Kind = (%q,%q)", got.GraderID, got.GraderKind)
			}
			if got.Message == "" {
				t.Errorf("Message must be non-empty for pass and fail cases; got empty")
			}
			if tc.wantScorePtr {
				if got.Score == nil {
					t.Errorf("Score ptr must be non-nil for %s", tc.kind)
				} else if *got.Score != tc.score {
					t.Errorf("Score = %v, want %v", *got.Score, tc.score)
				}
			} else if got.Score != nil {
				t.Errorf("Score ptr must be nil for binary grader kind %q; got %v", tc.kind, *got.Score)
			}
		})
	}
}

func TestEmitGraderComplete_NilSenderNoPanic(t *testing.T) {
	g := &fakeGrader{kind: graders.KindFile, name: "x"}
	emitGraderComplete(nil, g, graders.GraderResult{Pass: true, Message: "m"})
	// also nil grader
	emitGraderComplete(func(progress.ProgressEvent) { t.Fatal("unexpected") }, nil, graders.GraderResult{})
}

func TestBuildGraderHooks_NilSenderReturnsZeroHooks(t *testing.T) {
	h := buildGraderHooks(nil)
	if h.OnStart != nil || h.OnComplete != nil {
		t.Errorf("nil sender should produce zero-value hooks, got %+v", h)
	}
}

// TestRunGradersWithHooks_EmitsStartAndCompletePerGrader drives
// criteria.RunGradersWithHooks with 2 typed graders and 1 prompt_review
// grader, all backed by fakeGrader, and asserts one Start + one Complete
// event per grader in order, with matching IDs/kinds.
func TestRunGradersWithHooks_EmitsStartAndCompletePerGrader(t *testing.T) {
	var r reporter
	hooks := buildGraderHooks(r.emit)

	instances := []graders.Grader{
		&fakeGrader{kind: graders.KindFile, name: "file-check", result: graders.GraderResult{
			Kind: graders.KindFile, Name: "file-check", Pass: true, Message: "ok",
		}},
		&fakeGrader{kind: graders.KindBehavior, name: "behavior-check", result: graders.GraderResult{
			Kind: graders.KindBehavior, Name: "behavior-check", Pass: false, Message: "missing required tool",
		}},
		&fakeGrader{kind: graders.KindPromptReview, name: "ai_review", result: graders.GraderResult{
			Kind: graders.KindPromptReview, Name: "ai_review", Pass: true, Score: 0.9, Message: "panel 9/10",
		}},
	}
	configs := []graders.GraderConfig{
		{Kind: graders.KindFile, Name: "file-check"},
		{Kind: graders.KindBehavior, Name: "behavior-check"},
		{Kind: graders.KindPromptReview, Name: "ai_review"},
	}
	results := criteria.RunGradersWithHooks(context.Background(), instances, configs, graders.GraderInput{}, hooks)
	if len(results) != 3 {
		t.Fatalf("want 3 results, got %d", len(results))
	}
	// Expect event stream: Start(file), Complete(file), Start(behavior),
	// Complete(behavior), Start(ai_review), Complete(ai_review).
	if len(r.events) != 6 {
		t.Fatalf("want 6 events, got %d: %+v", len(r.events), r.events)
	}
	wantOrder := []struct {
		typ  progress.EventType
		id   string
		kind string
	}{
		{progress.EventGraderStart, "file-check", graders.KindFile},
		{progress.EventGraderComplete, "file-check", graders.KindFile},
		{progress.EventGraderStart, "behavior-check", graders.KindBehavior},
		{progress.EventGraderComplete, "behavior-check", graders.KindBehavior},
		{progress.EventGraderStart, "ai_review", graders.KindPromptReview},
		{progress.EventGraderComplete, "ai_review", graders.KindPromptReview},
	}
	for i, w := range wantOrder {
		got := r.events[i]
		if got.Type != w.typ || got.GraderID != w.id || got.GraderKind != w.kind {
			t.Errorf("events[%d] = (%v,%q,%q), want (%v,%q,%q)",
				i, got.Type, got.GraderID, got.GraderKind, w.typ, w.id, w.kind)
		}
	}
	// Score populated only on prompt_review; binary kinds nil.
	fileComplete := r.events[1]
	behaviorComplete := r.events[3]
	promptComplete := r.events[5]
	if fileComplete.Score != nil {
		t.Errorf("file Score should be nil, got %v", *fileComplete.Score)
	}
	if behaviorComplete.Score != nil {
		t.Errorf("behavior Score should be nil, got %v", *behaviorComplete.Score)
	}
	if promptComplete.Score == nil {
		t.Errorf("prompt_review Score should be non-nil")
	} else if *promptComplete.Score != 0.9 {
		t.Errorf("prompt_review Score = %v, want 0.9", *promptComplete.Score)
	}
	// Result field is populated for both pass and fail cases.
	if fileComplete.Result != progress.GraderResultPass {
		t.Errorf("file Result = %q, want pass", fileComplete.Result)
	}
	if behaviorComplete.Result != progress.GraderResultFail {
		t.Errorf("behavior Result = %q, want fail", behaviorComplete.Result)
	}
	// Messages are passed through unchanged.
	if fileComplete.Message == "" || behaviorComplete.Message == "" || promptComplete.Message == "" {
		t.Errorf("Message must be non-empty for pass AND fail; got file=%q behavior=%q prompt=%q",
			fileComplete.Message, behaviorComplete.Message, promptComplete.Message)
	}
}

// TestRunGradersWithHooks_NilReporterNoPanic ensures the whole grader path
// stays silent (and panic-free) when the engine isn't wired to a reporter.
func TestRunGradersWithHooks_NilReporterNoPanic(t *testing.T) {
	hooks := buildGraderHooks(nil)
	instances := []graders.Grader{
		&fakeGrader{kind: graders.KindFile, name: "f1", result: graders.GraderResult{Pass: true, Message: "ok"}},
	}
	configs := []graders.GraderConfig{{Kind: graders.KindFile, Name: "f1"}}
	results := criteria.RunGradersWithHooks(context.Background(), instances, configs, graders.GraderInput{}, hooks)
	if len(results) != 1 {
		t.Fatalf("want 1 result, got %d", len(results))
	}
}
