// Package compat provides compatibility shims for Copilot SDK protocol changes.
package compat

import copilot "github.com/github/copilot-sdk/go"

// ApproveAll is a permission handler that auto-approves all requests.
//
// Copilot CLI ≥1.0.36 uses a v3 permission protocol where "approved" was
// replaced by "approve-once". The upstream SDK constant still reads "approved"
// in v0.2.x, so this shim sends the value the server actually expects.
//
// When we upgrade to copilot-sdk/go v0.3.0+ (which ships the correct
// "approve-once" constant), this shim can be removed and callers can
// revert to copilot.PermissionHandler.ApproveAll.
var ApproveAll copilot.PermissionHandlerFunc = func(_ copilot.PermissionRequest, _ copilot.PermissionInvocation) (copilot.PermissionRequestResult, error) {
	return copilot.PermissionRequestResult{Kind: "approve-once"}, nil
}
