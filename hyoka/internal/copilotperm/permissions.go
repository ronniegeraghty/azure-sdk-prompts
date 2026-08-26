// Package copilotperm provides shared permission handlers for use with the
// Copilot Go SDK.
package copilotperm

import (
	copilot "github.com/github/copilot-sdk/go"
	"github.com/github/copilot-sdk/go/rpc"
)

// ApproveAll is a permission handler that approves every request.
//
// It explicitly returns the v1 typed approve-once decision so every tool
// request is approved without persisting the approval beyond that invocation.
func ApproveAll(_ copilot.PermissionRequest, _ copilot.PermissionInvocation) (rpc.PermissionDecision, error) {
	return &rpc.PermissionDecisionApproveOnce{}, nil
}
