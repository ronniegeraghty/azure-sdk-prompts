# Project Context

- **Project:** hyoka — Go evaluation tool that runs prompts through the Copilot SDK, reviews code via a multi-model panel, produces criteria-based pass/fail reports
- **Stack:** Go 1.26.1+, GitHub Copilot CLI/SDK, React + Recharts on frontend
- **User:** Ronnie Geraghty
- **Created:** 2026-04-03
- **Repo:** /home/rgeraghty/projects/hyoka

## Core Context

Agent Trinity initialized as Site / UI architect. Charter: frontend components, data visualization, React patterns. Expertise: Vitest + React Testing Library, Recharts integration, TypeScript strict mode, URL-driven state.

### Condensed History (Phase 0–4)

**Phase 0 (2026-04-03):** Team initialization. Trinity tasked with site work in parallel to core Go test infrastructure.

**Phase 1–3 (2026-04-04 → 2026-04-17):** Built initial site infrastructure (home page, runs page, SPA setup with Vitest). Identified and worked around ResizeObserver/IntersectionObserver gaps in jsdom for recharts compat.

**Phase 4 Wave 1–2 (2026-04-17 → 2026-04-20):** Site component work on eval-detail page, run-detail page, basic comparison UI. Phase 4 scope management: 5 items shipped, overflow plan established for burnout risk.

**Phase 5 (2026-04-20):** Designed compare-page with multi-select filters. Locked out of #364 artifacts per reviewer-protocol; unlocked by Morpheus. All Phase 5 work approved and merged.

**Key pattern:** URL persistence via useSearchParams; pure-function lib pattern for filter logic; Recharts integration via composition.

### Phase 6 CLI Invocation Convention (2026-04-21)

**Note:** As of Phase 5, main.go was moved to repo root. Any site docs that reference `hyoka` commands should use:
```bash
go run . <command>     # ✅ CORRECT
```

NOT:
```bash
go run ./hyoka ...     # ❌ STALE (Phase 5 regression)
```

Oracle found 47 stale references in phase-6 docs and fixed them (commits b5c4782c–874bedf9). When writing docs related to CLI interaction or getting-started guides, use `go run .` going forward.

## Recent Sessions

### Session 2026-04-21 (#600 — R146/R147 Run-Level Filter System)

Built the filter UI for the runs page. PR targets `phase-6` from worktree `hyoka-600`.

- **Run-level filtering, not eval-level.** The runs page lists runs; flattening it into eval cards would change its identity. So a run matches when each active filter dimension finds at least one matching eval inside it. The run-detail page still does its own per-eval filtering — that's the right surface for that.
- **Filter semantics inherited from `group-based-comparison-ui`.** OR within a dimension, AND across dimensions, empty = match-all. Documenting this once in the skill saves re-litigating it per page.
- **Status precedence: errors > failing > passing.** A run with both errors and failures is `errors` only. This avoids double-counting in the Status filter and matches how users think ("show me the broken runs").
- **URL is the source of truth, not React state.** `useSearchParams` + `setSearchParams(..., { replace: true })` so reload/share both work and history doesn't fill up with intermediate filter states. Keys: `config`, `lang`, `status` (comma-joined).
- **Reusable `MultiSelectFilter` primitive** at `components/ui/multi-select-filter.tsx` — outside-click + Escape close, ARIA listbox roles, "N selected" summary when >1. Other pages (prompts, dashboard) can adopt without copying state machinery.
- **Pure-function lib pattern again.** All matching logic in `lib/run-filters.ts`. 16 lib tests + extended runs-page DOM tests = 25 total new tests. Lib tests run in milliseconds with no React tree.
- **Catalog built from real data.** `buildCatalog` walks every run's `results[]` and only surfaces statuses that actually appear. No phantom filter chips.
- **Tests gotcha:** When asserting filter results in the DOM, the timestamp text (`Mar 28, 2026`) is the most stable identifier per card — `run_id` doesn't appear in the run card UI.
- **Decision filed:** `.squad/decisions/inbox/trinity-issue-600-filters.md`.

## Session 2026-04-21 (Phase 6 Round-1 Approval: Compare Page + Run Filters)

**Mission:** Test & architectural review of Phase 6 Round-1 batch (PRs #601, #602, #603)

**PR #601 Verdicts:**
- **Switch (test):** ✅ APPROVE — 31 new tests, 99/99 green, edge cases (top-bin overflow, malformed JSON, AND/OR filter semantics)
- **Morpheus (arch):** ✅ APPROVE — layering matches site convention, versioned localStorage, follow-up on dead `fetchCompareConfigs` endpoint

**PR #600 (Run-level filters) embedded in #601 review:**
- **Design:** Filter at run level (not per-eval); run matches if every active dimension finds ≥1 matching eval
- **Semantics:** Within-dim OR, across-dim AND, empty = match-all
- **Module:** `run-filters.ts` (pure lib) + `multi-select-filter.tsx` (reusable UI primitive) + runs-page orchestration
- **URL persistence:** `useSearchParams` as source of truth, `replace: true` for no-history pollution
- **Catalog:** Built from real data; only surfaces statuses that actually appear

**Status:** PR #601 approved, ready to merge with #602 + #603 (pending Tank's wiring tests).

**Phase 6 Round-1:** All three PRs approved. Embedded-asset refresh completed (a1a3c95d). Ready to merge to phase-6.

### Session 2026-04-21 (#608 — #604 Filter Test Polish)

**Branch:** `ronniegeraghty/issue-608-604-filter-tests` → **PR #609** (target `phase-6`)

Added three `MultiSelectFilter` component tests flagged by Switch on PR #604:
- Outside-click dismissal via `fireEvent.mouseDown` on a sibling element (component listens on `mousedown`, not `click` — the useEffect registers `mousedown`, so `fireEvent.click` would be a false-positive pattern).
- Escape key dismissal via `fireEvent.keyDown(document, { key: "Escape" })`.
- Empty-options state asserts the `No options` placeholder renders and `queryAllByRole("option")` is empty.

122/122 site tests green. No production code changes.

**Note:** `site/node_modules` did not exist in the worktree; `npm install` required before first test run. Normal for fresh worktrees.

### Session 2026-04-21 (#608 Round 3 — PR #613 — MultiSelectFilter follow-up tests)

**Branch:** issue-608 follow-up worktree → **PR #613** (target `phase-6`, squash-merged at `d05855df`)

Closed all four deferred test gaps Switch flagged on PR #609. Test-only diff, +232 LOC in `site/src/__tests__/multi-select-filter.test.tsx`, +11 tests (122 → 133 site total, all green in 3.28s).

**Coverage added:**
- **Toggle / onChange (3 tests):** Real `userEvent.setup()` + `await user.click()`. Controlled-state `Wrapper` (local `useState`) mirrors how `<FilterBar>` in `runs-page.tsx` actually wires the primitive. `toHaveBeenNthCalledWith(1..4, …)` validates each transition, not just final state. Exact-payload assertions (`toHaveBeenCalledWith(["a","b"])`).
- **Summary text (5 tests):** All branches — empty (renders placeholder), single, two, multi-overflow ("+N more").
- **ARIA (2 tests):** `aria-expanded` toggles dynamically open/closed; `aria-selected` per `option`.
- **Inside-click (1 test):** Counterpart to #609's outside-click — listbox stays mounted on inside-click. Component untouched.

**Behavioral lock-in (deliberate, not a bug-introduction):** Single-select summary at `multi-select-filter.tsx:57` renders `selected[0]` (the raw value, e.g. `"a"`), NOT the matching `opt.label`. Inconsistent with the option list and the multi-selected branch. **Test asserts current behavior with an explicit `// Note:` comment** rather than silently fixing in a tests-only PR. Filed as N1 in `2026-04-21-phase6-polish-nits-resolved.md` — own this as a separate one-line product PR with visual confirmation.

**Reviews:**
- Switch (test): ✅ APPROVE — confirmed real userEvent, exact-payload assertions, sequential-transition validation, mouseDown retained for outside-click.
- Morpheus (arch): ⚠️ APPROVE WITH NOTES — controlled-primitive boundary (D-2026-04-21) reinforced; Wrapper idiom appropriate; pattern portable to Compare-page filter bar at near-zero per-consumer cost.

**Pattern reinforced:** A tests-only PR should never silently change behavior. Locking in suspect behavior with an inline `// Note:` is the audit trail; the fix goes in a separate PR.
