package cmd

import "fmt"

// validateReviewMode enforces the allowed values for --review-mode (#580).
//
// Accepted values:
//   - ""         : treated as "combined" (legacy default before the flag existed)
//   - "combined" : single review session per panel model with merged criteria
//   - "isolated" : one session per grader/group marked isolate: true
//
// Any other value returns an error with the offending value quoted so users
// see exactly what they typed in the message.
//
// Extracted from runCmd's RunE so it can be unit-tested without spinning up
// the full prompt/config loading pipeline (#603, requested by Switch).
func validateReviewMode(mode string) error {
	switch mode {
	case "", "combined", "isolated":
		return nil
	default:
		return fmt.Errorf("invalid --review-mode %q: must be \"combined\" or \"isolated\"", mode)
	}
}
