package copilotperm

import (
	"testing"

	copilot "github.com/github/copilot-sdk/go"
	"github.com/github/copilot-sdk/go/rpc"
)

func TestApproveAllReturnsApproveOnceDecision(t *testing.T) {
	decision, err := ApproveAll(nil, copilot.PermissionInvocation{})
	if err != nil {
		t.Fatalf("ApproveAll() error = %v", err)
	}
	if _, ok := decision.(*rpc.PermissionDecisionApproveOnce); !ok {
		t.Fatalf("ApproveAll() decision = %T, want *rpc.PermissionDecisionApproveOnce", decision)
	}
}
