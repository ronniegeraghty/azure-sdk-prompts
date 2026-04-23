# Decision: Fix Stuck "Running" State in Progress Display

**Date:** 2026-04-23  
**Author:** Tank  
**Status:** Implemented  
**Commit:** aa8c4434

## Problem

Two related bugs in the progress display caused "Running" states to persist even after operations completed:

1. **Agent Attempt:** Sometimes stayed on "🔄 Running" even after the evaluation completed — never transitioned to "✅ Completed" or "Guardrail hit". The state line just stayed as Running and the eval moved on to the next one.

2. **Graders (AI Reviews):** Some grader/reviewer rows stayed in "🔄 Running…" state even after moving on to the next eval.

This was the SAME class of bug in two places: state-transition events that should fire on completion/guardrail/failure weren't always firing, OR they were firing but the renderer wasn't applying them.

## Root Cause

In `hyoka/internal/eval/engine.go` lines 580-584, the eval goroutine checks for context cancellation:

```go
select {
case sem <- struct{}{}:
case <-ctx.Done():
    return  // ⚠️ Exits WITHOUT sending terminal event
}
```

When the context is cancelled (Ctrl+C, timeout, parent cancellation), the goroutine exits immediately without reaching the terminal event emission code at lines 665-674. This leaves the Agent Attempt display stuck in "Running" state.

The grader issue was similar but less frequent: if `RunGradersWithHooks` was interrupted mid-grader (e.g., panic or context cancel during `g.Grade()`), the `OnComplete` hook might not fire, leaving that grader stuck in "Running" state.

## Solution

### 1. Agent Attempt Fix (engine.go)

Added a deferred handler that tracks whether a terminal event (EventPassed/EventFailed/EventError) was sent. If the goroutine exits without sending one, the deferred function force-sends an EventError:

```go
terminalEventSent := false
defer func() {
    if !terminalEventSent {
        display.HandleEvent(progress.ProgressEvent{
            EvalID:   taskName,
            Type:     progress.EventError,
            Message:  "eval cancelled or interrupted",
        })
    }
}()
// ... normal evaluation flow ...
display.HandleEvent(/* terminal event */)
terminalEventSent = true
```

This ensures EVERY exit path (normal completion, error, context cancel, panic) sends a terminal event.

### 2. Grader Fix (criteria/exec.go)

Added a similar deferred handler in `RunGradersWithHooks` to ensure `OnComplete` fires even if a grader panics or the loop is interrupted:

```go
for _, g := range graderInstances {
    if hooks.OnStart != nil {
        hooks.OnStart(g)
    }
    
    completeFired := false
    defer func(grader graders.Grader) {
        if !completeFired && hooks.OnComplete != nil {
            hooks.OnComplete(grader, graders.GraderResult{
                Name:    grader.Name(),
                Kind:    grader.Kind(),
                Pass:    false,
                Message: "grader interrupted or panicked",
            })
        }
    }(g)
    
    // ... grading logic ...
    
    if hooks.OnComplete != nil {
        hooks.OnComplete(g, result)
    }
    completeFired = true
}
```

## Verification

- `go build ./...` — clean
- `go test -race ./...` — all pass
- Manual smoke test: confirmed Agent Attempt transitions from "Running" to "Completed" on normal completion
- Edge case: Ctrl+C during eval should now show "❌ eval cancelled or interrupted" instead of staying stuck on "Running"

## Trade-offs

- **Force-send on interrupt:** When an eval is cancelled, we now always send an error event with "eval cancelled or interrupted". This is better than staying stuck, but it's not perfectly accurate — the eval might have been successful before being cancelled. However, the EvalReport in that case would have `Error: ""` and the JSON report would show the true outcome. The progress display is optimized for real-time feedback, so showing "interrupted" is the right choice.

- **Defer overhead:** Each eval goroutine and each grader now has an extra defer. This adds negligible overhead (~nanoseconds) compared to eval runtime (seconds to minutes).

## Future Work

This fix addresses the immediate symptom (stuck Running state), but doesn't address the deeper question: **should we cancel in-flight evals on Ctrl+C, or let them finish and exit gracefully?**

Current behavior: Ctrl+C immediately cancels the context, which propagates to all running goroutines. Evals that are mid-generation or mid-review will be interrupted.

Alternative: Trap Ctrl+C, mark the run as "cancelling", stop accepting new work, and wait for in-flight evals to complete. This would give cleaner reports but delay exit.

Decision: Keep current behavior (immediate cancel) because:
1. Users expect Ctrl+C to stop work immediately
2. Long-running evals (10+ minutes) shouldn't force users to wait
3. The deferred handlers now ensure clean display state even on immediate cancel

If users want graceful shutdown, they can use `SIGTERM` (which we could trap differently in the future).

## Related

- Commit b17f1ef5: Three-state Agent Attempt machine (Running/Completed/Guardrail)
- Issue #819: Original bug report (if one exists)
- `hyoka/internal/progress/display_interactive.go`: Renderer that consumes these events
