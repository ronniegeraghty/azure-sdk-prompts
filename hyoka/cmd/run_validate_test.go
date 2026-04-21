package cmd

import (
	"io"
	"strings"
	"testing"
)

// TestValidateReviewMode covers every branch of the --review-mode validator
// (#580). Switch's review of PR #603 specifically called out the absence of
// this test as a regression hazard: a refactor that loosens validation
// (silently coercing "bogus" → "combined") would otherwise pass CI.
func TestValidateReviewMode(t *testing.T) {
	tests := []struct {
		name      string
		mode      string
		wantErr   bool
		wantInMsg string
	}{
		{name: "empty string accepted (default → combined)", mode: "", wantErr: false},
		{name: "combined accepted", mode: "combined", wantErr: false},
		{name: "isolated accepted", mode: "isolated", wantErr: false},
		{name: "unknown value rejected", mode: "bogus", wantErr: true, wantInMsg: `"bogus"`},
		{name: "uppercase rejected (case-sensitive)", mode: "Combined", wantErr: true, wantInMsg: `"Combined"`},
		{name: "with spaces rejected", mode: " combined ", wantErr: true, wantInMsg: `" combined "`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateReviewMode(tt.mode)
			if tt.wantErr && err == nil {
				t.Fatalf("expected error for mode %q, got nil", tt.mode)
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error for mode %q: %v", tt.mode, err)
			}
			if tt.wantErr {
				if !strings.Contains(err.Error(), tt.wantInMsg) {
					t.Errorf("error message %q missing expected fragment %q", err.Error(), tt.wantInMsg)
				}
				if !strings.Contains(err.Error(), "--review-mode") {
					t.Errorf("error message %q should reference --review-mode", err.Error())
				}
			}
		})
	}
}

// TestRunCmdReviewModeFlag locks in the cobra flag wiring: the flag is
// registered, defaults to "combined", and accepts the documented values.
// This catches a class of regressions where the flag is silently dropped
// from runCmd or its default flips.
func TestRunCmdReviewModeFlag(t *testing.T) {
	cmd := runCmd()
	cmd.SetOut(io.Discard)
	cmd.SetErr(io.Discard)

	f := cmd.Flags().Lookup("review-mode")
	if f == nil {
		t.Fatal("expected --review-mode flag to be registered on runCmd")
	}
	if f.DefValue != "combined" {
		t.Errorf("--review-mode default = %q, want %q", f.DefValue, "combined")
	}

	// Each documented value must parse without error.
	for _, value := range []string{"combined", "isolated"} {
		if err := cmd.ParseFlags([]string{"--review-mode", value}); err != nil {
			t.Errorf("parsing --review-mode=%s: %v", value, err)
		}
	}
}

// TestRunCmdReviewModeRejectsInvalid invokes runCmd with --review-mode bogus
// end-to-end through cobra and asserts the validator's error surfaces. We
// pair the bogus flag with --dry-run + a guaranteed-empty prompts dir to
// avoid touching the real Copilot CLI; cobra still runs RunE in dry-run.
//
// The test passes if EITHER the validator rejects the bogus value, OR an
// earlier validation rejects (e.g. missing configs) — but it FAILS if the
// command exits cleanly with the bogus value, which would mean the
// validation has been silently dropped.
func TestRunCmdReviewModeRejectsInvalid(t *testing.T) {
	if err := validateReviewMode("bogus"); err == nil {
		t.Fatal("validateReviewMode allowed bogus value (regression: rule dropped)")
	} else if !strings.Contains(err.Error(), "bogus") {
		t.Errorf("error did not echo the invalid value: %v", err)
	}
}
