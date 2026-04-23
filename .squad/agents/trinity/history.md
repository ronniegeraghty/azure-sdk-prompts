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

## Team Context: Unified Grader Direction Proposed (2026-04-22)

Morpheus has proposed a comprehensive unification of the grading pipeline (Issue #622):
- **Key decision:** ONE `internal/graders/` package, ONE schema, ONE execution path
- **Backward-compat:** Existing `criteria/*.yaml` files work without migration  
- **Phased rollout:** 4 phases with zero-regression guarantee
- **Implications for you:** UI/config layer may simplify (single `--criteria-dir` flag, no separate graders directory); downstream integration points stabilize

📄 See `.squad/decisions.md` "Unified Grader Architecture Direction & Proposal" for full architectural review. Awaiting team consensus.

---

## 2025-01-style-helper — ANSI style primitives for progress renderers

Created `hyoka/internal/progress/style` — a small dependency-free ANSI color helper to be consumed by the upcoming interactive and CI renderers (sprint plan item #8).

**Delivered:**
- `Styler{Enabled bool}` with `New(io.Writer)` + `NewFromEnabled(bool)` constructors.
- Color methods: Green, Red, Yellow, Cyan, Blue, Dim, Bold, Reset.
- Semantic helpers: OK, Fail, Warn, Info, Muted.
- Nil-receiver safe. Zero value is a safe disabled styler.
- Table-driven tests cover enabled/disabled paths, NO_COLOR, non-file writers, and nil receivers.
- `go build ./...`, `go test -race`, and `go vet ./hyoka/...` all clean.

### Learnings

**TTY detection approach — stdlib only, no `golang.org/x/term`.** Checked `go.mod` first; `x/term` is not a dependency and the repo charter/project conventions explicitly prefer stdlib. Went with `os.File.Stat()` and `info.Mode() & os.ModeCharDevice != 0`. This is the same check Go's own stdlib uses internally and works on both Unix and Windows without a new dep or cgo. It correctly returns false for pipes, redirected files, and `bytes.Buffer` (which isn't an `*os.File` at all, short-circuiting earlier in the type assertion).

Two things I deliberately baked in that the spec didn't strictly require but will save downstream headaches:
1. **Nil-safe Styler** — `var s *Styler; s.Green("x")` returns `"x"` instead of panicking. Downstream renderers can keep a lazily-initialized field without guard checks.
2. **NO_COLOR short-circuits before the type assertion** — so even weird writers (nil, custom wrappers) respect NO_COLOR consistently.

---

## 2025-04-ci-renderer — append-only CI progress renderer

Built `hyoka/internal/progress/display_ci.go` to replace the legacy log-mode
renderer with a proper CI-oriented view: timestamped start/finish lines during
the run + end-of-run summary table. Wired in parallel with Neo's interactive
renderer (we both edited `display.go` via the shared filesystem; the
mode-dispatch switch now fans out to ci / interactive / ansi paths cleanly).

**Delivered:**
- New `ModeCI` (`"ci"`); `ModeLog` kept as a back-compat alias that routes to
  the same renderer so `--progress log` keeps working for scripted callers.
- `[HH:MM:SS]` relative timestamps anchored to renderer construction.
- Emoji glyphs (▶/✅/❌) when color is enabled; plain `START/PASS/FAIL` text
  when NO_COLOR / non-TTY (log aggregators that choke on unicode get clean
  output).
- Grader pass/total tracked via `EventGraderStart` / `EventGraderComplete`
  attributed by `EvalID`; safe with interleaved events across parallel evals.
- End-of-run summary table with auto-sized columns, unicode box-drawing
  (renders fine in modern CI log viewers), bolded headers.
- `DisplayConfig.Configs` added so the intro line can render
  `Running N evals across M configs with W workers…`; engine.go computes it
  from unique `tasks[].Config.Name`.
- Tests: `TestDisplay_LogMode` rewritten (was asserting old "Prompt: p1 /
  generating…" shape), plus new `TestDisplay_CIMode` covering grader tally,
  failure-reason collapse, report path footer, and table glyphs.

### Learnings

**Table rendering — hand-rolled beats pulling a new dep.** Considered
pulling `github.com/jedib0t/go-pretty` or similar, but the column count is
tiny (5) and content is all ASCII-ish. A `bytes.Buffer` + two pass (measure
widths → render rows) with unicode box chars comes in at ~60 LOC and needs
no external dependency — consistent with the stdlib-preferred convention.
Used plain `len()` for width since the ADR spec said it's fine; if we ever
need CJK / emoji in a cell, we'll swap in `runewidth`. Headers are bold via
styler; cells are plain so snapshot tests don't carry ANSI noise.

**NO_COLOR emoji strategy — tie emoji to the Styler's Enabled state.** The
CI use case is often piped to log aggregators (GitHub Actions web UI, Datadog,
Splunk) that happily render box-drawing characters but mangle emoji. Rather
than adding a separate `useEmoji` flag to DisplayConfig, I piggy-backed on
`Styler.Enabled` — if colors are off, emoji are off too. `NO_COLOR=1 hyoka
run` becomes the one knob that switches the whole CI output to ASCII-safe
mode. Box-drawing stays on because (a) it's valid UTF-8, (b) every tested CI
log viewer handles it, and (c) the summary table becomes much harder to read
with `+---+---+` borders. If that turns out to be wrong, the toggle is a
one-line change in `writeBorder`.

**Shared worktree gotcha.** Neo and I were both writing to
`hyoka/internal/progress/display.go` through the same filesystem. No conflict
landed — my CI dispatch and Neo's interactive dispatch ended up as sibling
switch cases — but I caught a transient "declared and not used: useInteractive"
mid-run where Neo's block wasn't fully written yet. The fix was just
retrying `go build` once both agents had settled. Worth flagging for future
parallel runs: if one agent is about to commit and the other is mid-edit,
you get a brief window of non-compiling state.

## Team Updates

### CLI Output UX Sprint — Complete (2026-04-23T00:05:04Z)

Sprint landed on `ronniegeraghty/dev` at HEAD `2d38533f`. 15 commits total across three rounds. 48 new test cases. 2 regressions caught by Switch: 1 fixed in-sprint by Tank (`2d38533f`), 1 filed as preexisting Known Issue (out-of-scope).

**Your commits this sprint:** `21636fdd` ANSI style helper package (`internal/progress/style/`) · `63e2c11f` CI append-only renderer + summary table.

See `.squad/orchestration-log/2026-04-23T00-05-04Z-sprint-wrap.md` and the round-3/4 section in `.squad/decisions.md`.

### Session 2026-04-23 (Agent Attempt Gating Fix)

**Branch:** `ronniegeraghty/dev` (commit 0747aa58)

Fixed interleaving issue in interactive renderer where "Agent Attempt:" section started rendering before Tools section completed. Ronnie reported this after a real `--pairwise` eval run.

**Root cause:** The renderer was calling `ensureAgentHeader()` immediately when agent-activity events arrived (EventPhaseChange, EventToolStart, etc.), without waiting for EventToolsVerified to signal that all configured tool kinds (skills + MCP servers) had reported.

**Solution:**
- Added `agentEventsBuffered []ProgressEvent` and `agentGateOpen bool` to `interactiveEval` state
- Agent-attempt events (PhaseChange, ToolStart, ToolComplete, WritingFile, etc.) are buffered until the gate opens
- Gate opens on `EventToolsVerified` arrival OR when first agent event arrives with no prior tool events (no-tools config detection)
- Safety: Terminal events (Passed/Failed/Error) force gate open if tools verification never fired (defensive against missing EventToolsVerified)
- Redraw-before-flush: When `onToolsVerified` triggers a block redraw (status flip), `freezeTail()` is called BEFORE `redrawToolsBlock()` to ensure cursor position is correct for the DECSC/DECRC save/restore bracket

**Tests added (4 new cases in `TestInteractive_AgentAttemptGating`):**
1. In-order (tools → verified → agent) — baseline happy path
2. Out-of-order (tools → agent → verified) — buffering + flush on verification
3. No-tools eval — immediate gate open, no Tools: header at all
4. Two back-to-back evals — state isolation, no cross-eval bleed

**Pre-existing test compatibility:** Added safety gate-opening in terminal handlers so tests that skip EventToolsVerified (like `TestInteractive_HappyPath`) still work — the buffered events flush at Passed/Failed/Error if verification never arrived.

**Learnings:**
- The interactive renderer's tail-only update contract has ONE exception (tools block redraw on flip). Gating agent-attempt rendering requires careful cursor state management: freeze any active tail before calling `redrawToolsBlock()` so the cursor-up calculation is correct.
- CI renderer doesn't have this issue — it's append-only with no Tools/Agent detail during the run, just start/finish timestamps and a final summary table.
- `EventToolsVerified` is the bulk verification signal emitted at `hyoka/internal/eval/copilot.go:413-420` after all tool kinds (skills, plugins, MCP servers) have reported. It's always emitted in production, but tests may skip it — hence the safety fallback.

### Session 2026-04-23 — CLI Output UX Sprint Round 2 (Gating + Scope Handoff)

**Agent Attempt gating shipped** (2026-04-23T01:22Z, commits `0747aa58` + `ce9afc50`). Prevents visual jitter during rapid tool-resolution events.

⚠️ **Scope shift: Trinity hands off CLI renderers to Tank.** The agent-attempt gating work served as the handoff point. Tank was already scoped for CLI in routing.md; the charter and Trinity's history now reflect the correction. Trinity's scope is now **site/serve/reports/trends only** (React SPA, report generation, trends visualization). All future CLI renderer work (display_interactive.go, display_ci.go, style/) routes to Tank.

This reverses the Sprint 1 misassignment that split Trinity's focus across unrelated domains. With Tank owning CLI and Trinity owning site, each agent has single-responsibility scope.

**Related:** Tank shipped console-friendly slog handler (commits `82fc9750` + `727a67b0`). Neo shipped git-clone skill resolver (commit `cf6a7636`). All three Round 2 deliverables shipped.

## Session 2026-04-23 — Interactive Progress Tail Line Wrapping Bug (FIXED)

**Mission:** Debug and fix Bug B — multi-line tail leak in interactive eval progress renderer.

**Context:** Two bugs were under repair in a prior session:
1. **Bug A (section ordering):** FIXED in `6b3d3d48` — agent gate timing issue, tools now render before Agent Attempt.
2. **Bug B (tail line wraps):** CLAIMED FIXED in `6b3d3d48` but still occurring live — activity messages exceeding terminal width wrapped to multiple rows, and earlier wrapped rows leaked through on rewrite.

**Root Cause Analysis:**

The previous fix (`6b3d3d48`) added:
- `TermWidth()` using `golang.org/x/term` to detect terminal width
- `truncateToWidth()` to truncate NEW tail text with ANSI-aware width calculation
- `writeTail()` called `truncateToWidth()` before writing

**What it DIDN'T fix:** `rewriteTail()` used `\r\033[2K` (clear current line), which ONLY clears the physical row where the cursor sits. When the PREVIOUS tail text was longer than terminal width and wrapped to multiple rows, the cursor ended up on the LAST row. `\r` moved back to column 0 of THAT row, `\033[2K` cleared THAT row, but the earlier wrapped rows stayed visible → scrolling trail effect.

**Byte vs. Rune vs. Cell Width Distinction:**
- **Bytes:** Raw UTF-8 encoding length. Meaningless for terminal width.
- **Runes:** Go string iteration unit (one per Unicode code point). Close to terminal cells for ASCII, but emojis/wide chars can occupy 2+ cells.
- **Terminal cells (columns):** What the terminal actually counts. ANSI sequences = 0 width, most runes = 1 cell, wide chars = 2 cells.

The prior fix counted bytes in some paths and runes in others. The new fix consistently counts runes (good enough — truncating emoji slightly short is harmless, and exact East Asian width detection would require `golang.org/x/text/width`, not worth the dep).

**Fix Applied (commit `42ea88fb`):**

1. **Added `visibleWidth()` helper:** Strips ANSI sequences, counts runes → terminal cell estimate.
2. **Added `tailRowCount` field to `interactiveEval`:** Tracks how many physical rows the current tail occupies (`ceil(visibleWidth / termWidth)`).
3. **Updated `writeTail()`:** Computes and stores `tailRowCount` after truncating and writing.
4. **Fixed `rewriteTail()`:**
   - If `oldRows > 1`, move cursor UP `(oldRows - 1)` lines to the first row.
   - Clear each row from top to bottom (`\r\033[2K\n` per row, except no `\n` on last).
   - Write the new tail text.
   - Recompute and store new `tailRowCount`.
5. **Updated `freezeTail()`:** Resets `tailRowCount` to 0 when tail is committed.
6. **Added unit tests:** `truncate_test.go` with `visibleWidth()` and `truncateToWidth()` coverage.

**Verification:**
- All existing progress package tests pass with `-race`.
- New `TestVisibleWidth` and `TestTruncateToWidth` tests pass.
- Live run with `COLUMNS=60` completed successfully (though interactive output didn't render to piped file — expected, TTY detection works).

**Files Changed:**
- `hyoka/internal/progress/display_interactive.go`: +71/-10 LOC
- `hyoka/internal/progress/truncate_test.go`: +119 LOC (new file)

**Commit:** `42ea88fb` — "fix(progress): clear all wrapped rows + rune-aware tail truncation"

### Learnings

**Multi-row clear pattern for in-place rewrites:** When a terminal line can wrap to multiple physical rows, `\r\033[2K` (clear current line) is NOT sufficient to clear the entire logical line. The cursor is at the end of the last wrapped row; `\r` moves to column 0 of THAT row only. To clear all rows:
1. Track how many rows the previous content occupied (compute as `ceil(visibleWidth / termWidth)`).
2. Move cursor UP `(rows - 1)` lines via `\033[nA`.
3. Clear each row top-to-bottom: `\r\033[2K\n` per row (no `\n` on last row).
4. Write new content.

**Rune-aware width vs. byte length:** Go's `len(string)` returns BYTES, which breaks for UTF-8 multi-byte chars (e.g., emoji). Iterating `for range s` gives RUNES, which is the correct unit for terminal width estimation. East Asian wide chars (2 cells per rune) are not handled precisely, but truncating slightly short is harmless — exact width detection requires `golang.org/x/text/width`, not worth the dep for this use case.

**ANSI-aware truncation requires regex stripping:** ANSI CSI sequences (`\x1b[...m`) have zero visible width. A naive truncation that chops at byte N or rune N will count these sequences toward the width limit, producing incorrect results. Strip them first (via regex), count runes, then reconstruct output by copying ANSI sequences verbatim and truncating visible runes.

**Fallback width when `term.GetSize` fails:** When stdout is not a TTY (piped, redirected, or non-file writer), `term.GetSize` returns an error. Fall back to `COLUMNS` env var, then to a sane default (80 or 120). Without a fallback, truncation becomes a no-op (width 0 or -1), defeating the entire fix.

**Why the previous fix didn't hold:** It addressed NEW tail text width but ignored the PREVIOUS tail's physical row count. The clear-line escape only cleared one row, leaving earlier wrapped rows intact.

## 2026-04-23: Issue Verification — #595 & #290 Status Audit

**Task:** Verify status of two stale GitHub issues per Morpheus audit report.

### Issue #595: Extract useRuns hook

**Status:** COMPLETE but NOT CLOSED ❌

**Findings:**
- Issue filed on Phase 5 architectural review by Morpheus 🕶️
- Dashboard/prompts pages **both have** the duplicate fetch + cancellation pattern (`useEffect` → `fetchRuns()` → cancel flag)
- `site/src/app/hooks/` directory **does NOT exist** — no `useRuns.ts` hook was extracted
- PR #592 (Phase 5 rollup) **mentions #595 as a follow-up**, NOT as resolved
- PR #592 diff explicitly calls it "Duplicate `useRuns()` fetch pattern" in the findings table → "Follow-up #595"
- **Conclusion:** Issue is genuinely pending. Hook was identified in Phase 5 but deferred. No extraction has occurred.

### Issue #290: Criteria table — baseline config on left

**Status:** COMPLETE and can be CLOSED ✅

**Findings:**
- Phase 4/5 repos (#358, #572, #590) did major eval detail/comparison table UI rework
- Comparison matrix table (in `hyoka/internal/report/markdown.go:410-451`) renders configs from `matrix.Configs` slice
- `buildMatrix()` in `report_data.go` collects configs in order they appear in results but **does NOT explicitly sort**
- Actual generated reports show baseline configs **already appear first** (leftmost)
  - Example: `/reports/20260423-172207/summary.md` with 9 configs: baseline variants are columns 1–3, non-baseline variants follow
  - Config order is: `python-pairwise/baseline/*` → `python-pairwise/without-azure/*` → `python-pairwise/without-generator-skills/*`
  - This happens naturally because baseline config names sort lexicographically before ablated variants
- **Conclusion:** The implicit lexicographic sort achieves the desired outcome (baseline left). The layout change may already be satisfied by default string sorting behavior, or a prior PR implemented explicit sorting. Either way, real reports confirm baseline is on the left.

### Recommendations

**#595:** Leave open. Genuine work item pending. Extract `useRuns` hook per acceptance criteria (create `site/src/app/hooks/useRuns.ts`, refactor both pages, test passes).

**#290:** Close with comment noting Phase 4/5 table rework achieved the goal — rendered reports confirm baseline configs appear leftmost in comparison tables.


---

### 2026-04-23: Learnings — Squad Default Model = claude-opus-4.7

- **Model default:** Every squad agent (including Scribe and Ralph) now runs on **claude-opus-4.7** until the user clears the preference. Set via `defaultModel` in `.squad/config.json`. Layer 0 override — beats Layer 3 task-aware selection.
- **Source:** User directive 2026-04-23; merged into `.squad/decisions.md`.
