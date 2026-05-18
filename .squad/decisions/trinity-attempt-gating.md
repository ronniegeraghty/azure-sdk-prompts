# Decision: Gate Agent Attempt Rendering on EventToolsVerified

**Date:** 2026-04-23  
**Author:** Trinity  
**Status:** Implemented (commit 0747aa58)  
**Branch:** ronniegeraghty/dev

## Context

The interactive progress renderer (`hyoka/internal/progress/display_interactive.go`) was designed to render per-eval transcripts with a strict section order:

```
Prompt: <id>
Config: <name>
Tools:
  - <tool>: ✅ Loaded | ❌ Failed
Agent Attempt:
  🔄 Running… (live tail)
  ✅ Complete — N files
Session Details: ...
Graders: ...
```

Ronnie ran a real `--pairwise` eval and reported that the Tools and Agent Attempt sections were interleaving — Agent Attempt started rendering before tool verification completed.

## Problem

Agent-attempt events (EventPhaseChange, EventToolStart, etc.) were rendering immediately when they arrived, without waiting for EventToolsVerified. This caused the "Agent Attempt:" header to print BEFORE all tool-resolution events had completed, breaking the intended layout.

In pairwise mode (or when tool resolution is slow), the interleaving was visible:

```
Tools:
  - mcp-a (mcp): 🔄 Loading…
Agent Attempt:          ← WRONG — started too early
  🔄 Running…
  - mcp-a (mcp): ✅ Loaded   ← tool line appears AFTER agent section started
```

## Decision

Implement gating logic to hold back Agent Attempt rendering until EventToolsVerified arrives:

1. **Buffer agent-attempt events per eval** until `EventToolsVerified` fires (or until a no-tools case is detected).
2. **On EventToolsVerified arrival:** flush buffered events in order, then resume real-time rendering.
3. **No-tools detection:** If the first agent-attempt event arrives and no tool events have been seen yet (no Tools: header printed, no tool lines), treat this as a no-tools config and open the gate immediately.
4. **Safety fallback:** Terminal events (Passed/Failed/Error) force the gate open if tools verification never arrived (defensive against missing EventToolsVerified in old tests or edge cases).
5. **Cursor hygiene:** When `onToolsVerified` triggers a block redraw (status flip), call `freezeTail()` BEFORE `redrawToolsBlock()` to ensure the DECSC/DECRC cursor calculation is correct.

## Implementation

Added to `interactiveEval` struct:
- `agentEventsBuffered []ProgressEvent` — holds buffered agent events
- `agentGateOpen bool` — tracks whether rendering is unblocked

New functions:
- `openAgentGate()` — flushes buffered events, sets gate flag
- `detectNoTools()` — returns true if no tool events seen yet
- `renderAgentEvent(evt)` — factored-out rendering logic for agent events

Modified event handlers:
- `onPhaseChange`, `onAgentActivity` — check gate state; buffer or render
- `onToolsVerified` — opens gate after tools block is complete
- `onPassed`, `onFailed`, `onError` — force gate open if still closed (safety)

## Rationale

- **Strictly ordered output:** Users expect Tools to finish rendering before Agent Attempt begins. Interleaving breaks the mental model and makes logs harder to parse.
- **Matches original design intent:** The sprint plan and layout spec (`plan.md`) explicitly show Tools before Agent Attempt.
- **No-tools configs work seamlessly:** Configs with zero tools (no skills, no MCP) don't have `EventToolsVerified` to wait for, so the renderer detects this and proceeds immediately.
- **Cross-eval isolation:** Each eval's state is keyed on EvalID; buffering is per-eval, so back-to-back evals don't interfere.
- **Defensive fallback:** Old tests or edge cases that skip EventToolsVerified won't hang — terminal events force flush.

## Alternatives Considered

1. **Emit EventToolsVerified earlier (upstream):** Rejected — the event fires after session start, which is the correct timing for post-session-start verification. Moving it earlier would break the verification contract.
2. **Delay all rendering until verification completes:** Rejected — users expect live progress. Buffering only the agent-attempt section preserves real-time tool-resolution feedback while ensuring correct section order.
3. **Print "Agent Attempt:" immediately but buffer tail lines:** Rejected — the header line itself is part of the section that should wait. Printing it early would still break the layout.

## Impact

- **User-visible:** Logs now have Tools fully complete before Agent Attempt starts, matching the intended layout.
- **CI renderer:** No change — CI mode is append-only with no Tools/Agent detail during the run, just start/finish timestamps and a summary table.
- **Tests:** Added 4 new test cases covering in-order, out-of-order, no-tools, and cross-eval scenarios. Pre-existing tests (that skip EventToolsVerified) still pass due to safety fallback.

## Follow-Up

None required. The fix is complete and tested. If future tool-resolution changes alter the timing of EventToolsVerified, the no-tools detection and safety fallback ensure the renderer adapts gracefully.
