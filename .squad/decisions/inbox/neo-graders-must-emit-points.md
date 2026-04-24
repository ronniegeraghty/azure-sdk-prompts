# Decision: Every Grader Must Emit ≥ 1 GraderPoint

**Author:** Neo 💊  
**Date:** 2026-04-25  
**Status:** Proposed (inbox)  
**Scope:** `hyoka/internal/criteria/graders/`, `hyoka/internal/criteria/exec.go`, `hyoka/internal/eval/engine_eval.go`

## Invariant

> **Every `graders.GraderResult` produced by the engine — whether by a normal Grade() path, an error fallback, or a panic-recovery path — MUST carry at least one `GraderPoint`. A Points-less result is a bug.**

## Rationale

The site renders per-grader UI off the `Points` slice. When `len(Points) == 0` the renderer historically fell back to displaying the grader's overall `Pass` and `Score` as a "PASS" badge with "100%" — even when the grader had silently failed or been skipped. This produced false-positive UI on every error path.

Trinity's site fix adds defensive rendering (no PASS/100% headers, collapsed-by-default rows, blank-passing-point-label). Neo's engine fix guarantees the renderer never receives a Points-less result in the first place. Both layers together: belt and braces.

## How the invariant is enforced

1. **Constructor-level (`graders.NewResult`)** — panics if `len(points) == 0`. This is the single canonical way to build a successful Result.
2. **Error fallback (`graders.NewErrorResult`)** — synthesizes a single failing `"grader executed"` Point and routes through `NewResult`. Use this everywhere a Result must be constructed outside a normal Grade() execution (engine error paths, panic-recovery hooks, future skipped-grader paths).
3. **Defensive converter (`engine_eval.convertGraderResults`)** — if a Points-less result somehow reaches the report layer, log `slog.Warn` and synthesize a fallback Point so the JSON is well-formed. This branch should be unreachable; the warning makes new bypasses visible.
4. **Per-grader fallback Points** — graders that allow zero-knob configuration (`BehaviorGrader`, `ToolConstraintGrader`, `OutputCheckGrader`) emit a single trivially-passing `"no_constraints"` / `"no_knobs"` Point when no constraints are configured.

## What this means for new graders

When writing a new grader:

- Always go through `graders.NewResult(kind, name, cfg, points, msg, extras)`.
- Never construct `graders.GraderResult{...}` directly.
- If your grader has configurable sub-checks (knobs), ensure that the "no knobs configured" branch still emits a single trivially-passing Point.
- Each Point needs a stable, snake_case `Label` (the identifier — used by tests and trend aggregation), a `Pass` bool, and a `Message` describing what was checked. Pass-case messages are encouraged but not required; fail-case messages are required.
- If you need to construct an error or skipped-grader result outside Grade(), use `graders.NewErrorResult(kind, name, cfg, msg)`.

## What this means for engine code

When invoking graders or constructing fallback results outside a Grade() call:

- Use `graders.NewErrorResult` — never assemble a `graders.GraderResult{}` literal.
- The engine's panic-recovery and error-fallback paths in `criteria/exec.go` and `eval/engine_eval.go` are the canonical examples.

## Test coverage

- `graders/points_test.go::TestEveryGraderEmitsPointsOnPassAndFail` — exercises every concrete grader kind in pass and fail scenarios; asserts `len(Points) >= 1` and non-empty Labels.
- `graders/points_test.go::TestNewResult_PanicsOnEmptyPoints` — locks the constructor invariant.
- `graders/points_test.go::TestNewErrorResult_AlwaysEmitsPoint` — locks the error-fallback invariant.

## For Trinity

The graders that previously emitted Points-less results — and triggered the "PASS"/"100%" rendering — were **engine error fallbacks** in:

- `hyoka/internal/criteria/exec.go::RunGradersWithHooks` (grader returned error or panicked)
- `hyoka/internal/eval/engine_eval.go` review-grader error fallback

Both now use `NewErrorResult`. Your defensive site code is still valuable for legacy on-disk reports (pre-v4 schema), but freshly-generated v4 reports will never lack Points.

## Verification

Real eval (2026-04-25): `hyoka run --prompt-id key-vault-dp-python-crud --config baseline/claude-opus-4.6` → `reports/20260424-195854/.../report.json`. `jq '.grader_results | map(select(.points == null or (.points | length) == 0))'` returns `[]`. Zero graderless graders.
