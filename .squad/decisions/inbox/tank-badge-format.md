# Grader Points — badge format alignment & end-to-end Points verification

**Author:** Tank 📡 — 2026-04-24
**Scope ref:** `.squad/decisions/inbox/morpheus-prompt-grader-checks-scope.md` §6, §7
**Branch:** `ronniegeraghty/prompt-grader-checks` (shared with Neo)
**Status:** Implemented, tested, end-to-end verified.

---

## What changed

1. **Badge format** in `hyoka/internal/progress/display_interactive.go` (`renderGraderWithPoints`):
   - Old: `✅ 2/2 passed` / `❌ 1/3 passed`
   - New: `✅ Pass (2/2)` / `❌ Fail (1/3)` (matches user spec).
   - One-line behavior change; godoc updated.

2. **Soft truncation** of per-Point names in the interactive renderer:
   - `name = truncateToWidth(p.Name, 50)` — long check strings get a `…` ellipsis at ~50 cols.
   - Reuses existing `truncateToWidth` (ANSI-aware, wide-char-aware) — no new helpers, no width plumbing needed.
   - Full text remains in `report.GraderResult.Points[].Name`.
   - Rationale for fixed 50: terminal width isn't piped down to this renderer today, and 50 cols looks clean inside the `    - <name>: ✅ Pass` indent without breaking on standard 100-col terminals. Easy to swap for terminal-width-derived later if Ronnie wants.

3. **Test updates** in `display_interactive_points_test.go`:
   - Three assertions migrated from `"X/Y passed"` to `"Pass (X/Y)"` / `"Fail (X/Y)"`.
   - Single-Point flat-row negative check updated to the new format so it still catches regressions.

## Verifications

- `go build ./...` — clean.
- `go test -race ./internal/progress/... ./internal/report/... -timeout 3m` — green (progress 1.07s, report 29.5s).
- **End-to-end smoke** against Neo's migrated `criteria/language/test.yaml` using `prompts/test/hello-markdown.prompt.md` + `test/haiku` config. Inspected `reports/20260424-052849/.../report.json`:
  - `Markdown Structure` (YAML grader, `type: prompt`, `checks: [2 items]`) → `points: 2` ✅
  - `Criteria from prompt file` (prompt-file path, 3 bullets under `## Evaluation Criteria`) → `points: 3` ✅
  - Both flow through `convertGraderResults` (`hyoka/internal/eval/engine_eval.go:1199-1204`) into `report.GraderResult.Points`. No schema bump needed — Phase 2 v3 already covered this.

## Renderers verified (no changes required)

- `display_interactive.go` — updated (this PR).
- `display_ci.go` — only tracks aggregate pass/total per eval; doesn't render Points individually. Correct as-is.
- No separate `display_json*.go` / `display_quiet*.go` exist; report JSON serialization is owned by `internal/report` and was verified above.

## Out of scope (per Morpheus's scope §6, §7)

- Site / TS template Points rendering — Trinity's domain (Phase 6).
- Schema bump — not needed; v3 already carries Points.
- Terminal-width-aware truncation — fixed 50-col cap is sufficient for now.

## Co-author

Co-authored with Neo on the same branch (`ronniegeraghty/prompt-grader-checks`).
Tank's commits are file-disjoint from Neo's: `hyoka/internal/progress/display_interactive*.go` only.
