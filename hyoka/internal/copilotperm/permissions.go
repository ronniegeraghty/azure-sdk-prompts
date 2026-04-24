// Package copilotperm provides shared permission handlers for use with the
// Copilot Go SDK.
package copilotperm

import (
	copilot "github.com/github/copilot-sdk/go"
)

// ApproveAll is a permission handler that approves every request.
//
// It exists because copilot.PermissionHandler.ApproveAll (SDK v0.2.0) returns
// {kind: "approved"}, which the Copilot CLI v0.0.415-1 (1.0.35) does not
// recognize — the CLI's permission switch only accepts "approve-once",
// "approve-for-session", "approve-for-location", "reject", and
// "user-not-available". Anything else triggers
// "Error: unexpected user permission response" and the tool call fails with
// the LLM seeing only a generic permission error.
//
// Returning Kind "approve-once" maps cleanly to the CLI's expected
// approval path for a single tool invocation.
//
// TODO: Remove once the SDK and CLI agree on the wire format and
// copilot.PermissionHandler.ApproveAll works against the installed CLI.
func ApproveAll(_ copilot.PermissionRequest, _ copilot.PermissionInvocation) (copilot.PermissionRequestResult, error) {
	return copilot.PermissionRequestResult{Kind: "approve-once"}, nil
}
