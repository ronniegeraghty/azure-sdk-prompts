# Discrepancy: 82cd8590 tool-verification wiring not on `ronniegeraghty/dev`

**Author:** Switch 🤍
**Date:** 2026-04-22
**Severity:** Medium — decisions ledger claims "Shipped" but commit is not in branch.

## What

`.squad/decisions.md` (active decisions, round 1–2 consolidation, 2026-04-22) lists:

| Item | Agent | Commit | Status |
|------|-------|--------|--------|
| Tool-verification emission wiring (SDK post-session) | Neo 💊 | `82cd8590` | ✅ Shipped |

`git merge-base --is-ancestor 82cd8590 HEAD` returns non-zero on
`ronniegeraghty/dev` (HEAD was `a0105a9d` when verified). The commit
sits on a parallel branch that diverged from `bffd0c40` (grader
serialization) — dev went on to include `e06ead61` (tool-resolution
emit), while `82cd8590` never merged.

Before this commit, `hyoka/internal/eval/copilot.go` contained the
expected-set building but nothing that emitted `EventToolsVerified`.
No test would have caught it, because verification was never
unit-testable — it lived entirely as a closure inside `Run()`.

## What I did about it

While writing `tests-events-wiring`, I re-landed the emission in a
**more testable shape**:

- New file `hyoka/internal/eval/tool_verification.go` exports a
  `toolVerifier` struct with `newToolVerifier`, `onSkillsLoaded`,
  `onMCPLoaded`, `emitIfReady`.
- `copilot.go` OnEvent handler now constructs one `toolVerifier` per
  eval, calls the `onXLoaded` methods inside the `mu.Lock()` region,
  and invokes `e.progressFn` with the resulting `EventToolsVerified`
  payload *after* `mu.Unlock()` — preserving the "build under lock,
  dispatch outside lock" guarantee documented in the round 1–2 spec.
- Behaviour matches the original 82cd8590 contract: at-most-once,
  fires only when every configured kind has produced its SDK load
  event, deterministic sort, configured-only payload, plugins not
  covered.
- Nine table-driven tests in
  `hyoka/internal/eval/tool_verification_test.go` cover every bullet.

## Request

Scribe — when you merge the next round, please:

1. Update the round 1–2 table to mark `82cd8590` as **Re-landed (Switch)**
   with the commit from this testing sprint, or strike through the
   row if it's no longer semantically accurate.
2. Confirm with Neo that no additional behaviour from `82cd8590`
   (e.g. logging side effects) was lost. I preserved the existing
   `lg.Warn("Expected MCP server not loaded", ...)` and
   `"No MCP servers loaded despite configuration"` slog paths.

No blockers for the round-3 renderer work — renderers see
`EventToolsVerified` events as expected; the interactive renderer
snapshot tests (`display_interactive_test.go`) still pass.
