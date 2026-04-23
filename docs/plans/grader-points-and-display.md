# Plan: Tool-loading display polish + grader-points rethink

**Updated 2026-04-23:** Morpheus appended Phase 4 (site quick wins), Phase 5 (report schema v3), and Phase 6 (site Phase-2 alignment) after synthesizing the report-data + site-UX reviews. Phases 1–3 remain unchanged — Tank is mid-execution. New work strictly extends; nothing above this line was modified.

**Branch:** `ronniegeraghty/dev` (commit directly, no PRs, no GH issues)
**Plan doc location:** committed to repo as `docs/plans/grader-points-and-display.md` (so the team can reference it after this session)
**Source layout reminder:** all Go code lives under `hyoka/internal/...` and `hyoka/cmd/...`; module root is `hyoka/hyoka/`.

## Problem

Live-eval output and the grader model both have rough edges that are now visible because real evals are passing end-to-end again (post container-plugin fix `4a8c4a0d`). Six concrete issues, grouped by area:

### A. Tool-loading section (renderer)
1. **Skill-dir parent uses the full path as its display name.** Should use the config entry's `name` (e.g. `generator-skills`), not `./skills/generator`.
2. **Plugin parent claims it failed to load.** The SDK never reports the plugin itself — only its child skills/MCPs. Today the renderer's "loaded but not reported by SDK → failed" rule (display_interactive.go:506-514) flips the plugin parent to ❌. Children loaded fine; the parent shouldn't carry a pass/fail at all.
3. **Children under a plugin lose their kind label.** `renderToolLine(..., indented=true)` suppresses the `(skill)` / `(mcp)` suffix for children (display_interactive.go:672-675). Should keep it so `(skill)` vs `(mcp)` is obvious under a plugin header.

### B. Live-render lifecycle bugs
4. **Agent Attempt status appears below itself, stuck on Running.** The "Running" line is written as the active tail (`tailKind=tailAgent`). When grader events fire, they `freezeTail()` (committing the Running line as permanent) then take over the tail with `tailKind=tailGrader`. When `agentComplete` later runs, `tailKind != tailAgent` so it falls into the `else` branch (display_interactive.go:817-823): freeze the grader tail, then `writeLine` a brand-new `✅ Completed` line at the very bottom — far below where the original Running line was frozen.
5. **`ai_review` grader entry shows up multiple times in different positions.** Mostly a symptom of #4 + the surrounding `sendEvent(EventToolStart, "Review panel: …")` and `sendEvent(EventToolComplete, …)` calls in the review block (engine_eval.go:514-536) which interact with the agent tail right around the grader tail handoff. Suspected cause: tail kind transitions while a grader Start/Complete pair is in flight, so the grader's `rewriteTail` no-ops and `onGraderComplete` falls through to `writeLine` (display_interactive.go:888-894), producing a second printed entry.

### C. Grader model rethink
6. **Graders are flat: one grader = one Pass/Fail.** Reality is richer:
   - `output_check` already produces N `OutputCheckSubResult` rows internally — but the renderer only sees the rolled-up `Pass`.
   - `prompt_review` (the AI panel) returns N `ReviewCriterion` rows — same fate.
   - Auto-extracted prompt criteria from a prompt's markdown frontmatter are evaluated as one Copilot review session, but conceptually they're N independent pass/fails.

   We want one unified concept: **a grader has one or more grading points.** The grader's overall pass = AND of its points; the rendered output shows the grader header with each point indented under it.

## Approach

One PR-equivalent (single working branch, multiple commits). Land the small renderer fixes first so we get visible wins quickly; then build the grader-points generalization on top. All commits go straight to `ronniegeraghty/dev`.

### Phase 1 — Display polish (small, surgical)

| # | Change | File |
|---|--------|------|
| 1 | `validateSkillDirEntry` writes `Parent: entry.Name` (the config name) instead of `entry.Path`. Update related parent-row emission (currently absent for skill_dir; the renderer reconstructs it from `parentName`). Verify rendered header reads `- generator-skills (skills dir):`. | `hyoka/internal/config/tool/validate.go` |
| 2 | Stop emitting a top-level plugin Loaded/Failed row in `emitPluginLoadedWithChildren` and the remote-plugin success path (validate.go:372, 400-401, 438). Instead, emit only the children with `ParentKind=plugin` + `ParentName=<plugin-name>` so the renderer's container-discovery (display_interactive.go:574-578) still groups them. The plugin header just becomes `- <name> (plugin):` with no status. | `validate.go` + `display_interactive.go` (remove the "loaded but not reported by SDK → failed" flip when the entry is a plugin parent — display_interactive.go:506-514) |
| 3 | In `renderToolLine`, keep `kindStr = " (kind)"` for children too (drop the `if !indented` guard at display_interactive.go:672). Verify the kind tag appears for skill_dir children as well — they should already read as `(skill)`. | `hyoka/internal/progress/display_interactive.go` |
| 4 | **Agent Attempt completion lifecycle.** Track `agentLineFrozen bool` + `agentLineRow int` on `evalRenderState`. When the agent tail is frozen by some other event taking the tail (graders), record its row index. In `agentComplete`, if the agent line was frozen, do a single bracketed save/restore rewrite (same DECSC/DECRC pattern as `redrawToolsBlock`) to update that single line in place rather than writing a new line at the bottom. Add a regression test using the existing `display_interactive_test.go` table-driven harness. | `hyoka/internal/progress/display_interactive.go` |
| 5 | **`ai_review` grader render bug.** Two sub-fixes:<br/>(a) Drop the redundant `sendEvent(EventToolStart, "Review panel: …")` / `sendEvent(EventToolComplete, "Review complete: …")` in engine_eval.go:514-536 — they're the noise source. The grader Start/Complete events already convey this state; the dual emission is what disturbs the tail.<br/>(b) Make `onGraderComplete`'s fallback path (display_interactive.go:891-894) idempotent: if a `Running` row for the same `GraderID` was already frozen above, rewrite that row in place (same save/restore pattern as #4) instead of appending a new one. Track grader rows in a `graderRowByID` map. | `hyoka/internal/eval/engine_eval.go`, `hyoka/internal/progress/display_interactive.go` |

Each item gets a regression test exercising the renderer or validator as appropriate. Existing test files (`display_interactive_test.go`, `display_interactive_plugins_test.go`, `validate_*_test.go`) are the homes.

### Phase 2 — Grader points generalization

Goal: every grader emits 1+ **grading points**. The renderer shows one nested block per grader.

#### Data model

Add a single concept to `criteria/graders/grader.go`:

```go
// GraderPoint is one binary pass/fail check inside a grader. A grader's
// overall Pass is the AND of every Point.Pass.
type GraderPoint struct {
    Name    string `json:"name"`           // "min_files", "uses_default_credential", "criterion: returns Future"
    Pass    bool   `json:"pass"`
    Message string `json:"message,omitempty"`
}

type GraderResult struct {
    // ... existing fields ...
    Points []GraderPoint `json:"points,omitempty"`  // NEW: one row per sub-check
}
```

Existing `OutputCheckSubResult`, `ReviewCriterion`, and the prompt grader's single score collapse into `Points`:

| Grader kind | How `Points` is populated |
|-------------|---------------------------|
| `file` | One point per file checked (existing `FileCheckResult` → `GraderPoint`) |
| `program` | Single point: `Name="exit code 0"`, `Pass=ExitCode==0` |
| `output_check` | One point per configured knob (`min_files`, `require_files: foo.py`, etc.) — already structured this way internally |
| `behavior` | One point per constraint (required tool present, forbidden tool absent, turn limit, …) |
| `prompt` (LLM-as-judge with single score) | Single point: `Name="LLM judge", Pass=score>=threshold` |
| `prompt_review` (AI panel with criteria) | One point per `ReviewCriterion` — already structured this way internally |
| Auto-extracted prompt criteria | The whole prompt's `evaluation_criteria` becomes one `prompt_review` grader instance with N points (one per criterion). One Copilot review session per grader instance. |

The detail structs (`OutputCheckGraderDetails`, `ReviewGraderDetails`, etc.) stay for report fidelity (templates already render them) — `Points` is added alongside, not as a replacement, so existing report templates keep working. The renderer reads `Points` only.

#### Progress events

Extend `progress.ProgressEvent`:

```go
// On EventGraderComplete, optional list of per-point outcomes. Empty for
// single-point graders; the renderer falls back to the existing flat row.
Points []GraderPoint
```

`emitGraderComplete` populates `Points` from `result.Points`. No new event type — points ride with the existing complete event.

#### Renderer

Update `onGraderComplete` to render:

```
  - ai_review (prompt_review):
    - returns DefaultAzureCredential: ✅ Pass
    - exposes get_secret/set_secret/delete_secret: ✅ Pass
    - paginates list_secrets: ❌ Fail
```

When `len(evt.Points) <= 1`, fall back to today's flat row (one line). When `> 1`, freeze the tail, write the grader header line `  - <id> (<kind>):`, then one indented row per point. Aggregate Pass/Fail badge at the end of the header, e.g. `  - ai_review (prompt_review): ❌ 2/3 passed`.

#### Report layer

Add `Points` propagation through `report.GraderResult` (the report-side mirror in `internal/report/types.go`) so JSON reports carry the new field. Existing detail structs stay for backward compat with the static markdown/HTML templates.

#### Engine wiring

No change to engine flow. `RunGradersWithHooks` already invokes each grader once. Each grader's `Grade` method is responsible for filling `Points`. Update each grader implementation in `criteria/graders/` to emit at least one point.

### Phase 3 — Verify

1. `go build ./...` and `go test ./hyoka/...` green.
2. Live eval: `hyoka run --prompt-id key-vault-dp-python-crud --config python-pairwise --log-level debug --log-file /tmp/hyoka-verify.log` then `hyoka clean`. Eyeball the live output for:
   - `- generator-skills (skills dir):` (no path)
   - `- azure-sdk-python (plugin):` with no Loaded/Failed badge, children listed with `(skill)` or `(mcp)`
   - `Agent Attempt:` with `✅ Completed` in the right position (where Running was)
   - `ai_review` shows once, with per-criterion points indented under it
3. Capture before/after screenshots in `docs/plans/grader-points-and-display.md` for posterity.

## Todos

Tracked in SQL (see todo list). Phase 1 todos can run in parallel; Phase 2 todos serialize after Phase 1 lands.

## Notes & decisions

- **No GH issues, no PRs.** User wants direct commits to `ronniegeraghty/dev`.
- **Plan doc location:** committed at `docs/plans/grader-points-and-display.md` for team reference. The session-state copy is the working draft.
- **Backward compat for reports:** detail structs (`OutputCheckGraderDetails` etc.) stay alongside the new `Points` field — existing report templates and JSON consumers keep working.
- **Auto-extracted prompt criteria:** treated as ONE `prompt_review` grader with N points. One Copilot review session per grader. This drops the question of "should each criterion be its own grader / its own session?" entirely — the answer is "no, one session, many points."
- **Renderer save/restore pattern:** the existing `redrawToolsBlock` (display_interactive.go:530) is the template for in-place line rewriting. Reuse the same DECSC/DECRC bracket for agent + grader fixes.

---

## Synthesis note (2026-04-23): the "all passed but pages disagree" root cause

Reproduced against `reports/20260423-195948` (12 evals, `passed: 12, failed: 0`). The site shows red `1/4` and `0/0` per row anyway. Root cause is **two divergent roll-ups**:

- **Engine truth:** `EvalReport.Success` set from `agg.Pass` over the internal grader list (one entry per grader, `Pass bool`).
- **Site re-rolls:** `run-detail-page.tsx:236-237` filters `g.pass === true` over the **expanded** report grader list. That list has been blown up by `expandReviewGraderResult` (`engine_eval.go:903-953`), which takes one passing `ai_review` grader and emits 3 entries (panel members + consensus) all with `Pass: nil`. The strict `=== true` filter excludes all three → `1/4` red.

**Phase 2's grader Points work IS the structural fix.** Once each grader emits one `GraderResult` with `Points[]`, `expandReviewGraderResult` becomes obsolete: a single `ai_review` row with `Pass=true` and N `Points` replaces the 3 nil-pass rows. The site then renders `Points[]` directly. The "1/4 red on every row" bug disappears by construction.

Phase 4 ships the site-side stop-gap immediately. Phase 5 lands the schema bump (kill the expansion, add parent linkage, bump `SchemaVersion` to v3) once Phase 2 grader Points exist. Phase 6 updates the site templates to consume the new shape.

---

## Phase 4 — Site quick wins (independent of Phase 2)

**Owner:** Trinity (frontend / serve)
**Depends on:** nothing — pure presentation fixes that ship today and remove the false-negative impressions while Phases 1–3 finish.
**Why:** Today the site contradicts itself: header says `12 Passed / 100%`, every row shows red `1/4`. Five of these are tiny TS edits; one (dashboard crash) needs a bisect.

| # | Change | File | Notes |
|---|--------|------|-------|
| 4.1 | Replace strict `pass === true` filter with a tri-state-aware `isPass()` helper. Falls back to `criteria.every(c => c.passed)` when `pass` is null and criteria exist; `overall_score === max_score` when only scores exist. When `grader_results.length === 0`, render `r.success ? "—" : "✗"` instead of red `0/0`. Wire through `ScoreBadge`. | `site/src/app/components/run-detail-page.tsx:236-256` | See Trinity §3 Q1 for the exact snippet. Stop-gap until Phase 5 kills the expansion. |
| 4.2 | Same fallback in `GraderResultRow`: when `result.pass == null`, derive from `scores.criteria` AND or `overall_score === max_score`. Three review rows on the anchor eval flip from grey `N/A` to emerald `PASS` with no backend change. | `site/src/app/components/GraderResultRow.tsx:16` | See Trinity §3 Q2. |
| 4.3 | `/dashboard` crash — `Cannot read properties of undefined (reading 'toFixed')`. Bisect the un-minified bundle to find the exact site, add `?? 0` guard, friendlier ErrorBoundary message. | `site/src/app/components/dashboard-page.tsx` (line not yet bisected) | Trinity §1d + §3 Q3. May want Switch or Tank to bisect since the bundle is minified. |
| 4.4 | `/runs` rate column folds errored runs into denominator → all-error runs render as emerald `0.0%`. Fix: compute `effectiveTotal = total - errors`, paint amber when `errors > 0`, OR add `⚠ run errored` tag. | `site/src/app/components/runs-page.tsx:136-198` | Trinity §1e + §3 Q4. |
| 4.5 | `/runs` shows in-progress run dirs as `Unknown … 0.0% N/A`. Either filter dirs without `summary.json` or render `In progress` with a spinner. | `site/src/app/components/runs-page.tsx` | Trinity §3 Q5. |
| 4.6 | Eval-detail header `12 / 12` card is misleading — it's review-only score sitting above grader rows that read `N/A`. Re-label to `Review Score 12/12` until Phase 6 reworks the score-card metaphor. | `site/src/app/components/eval-detail-page.tsx:441-444` | Trinity §3 Q6 option (a). |

**No coordination needed with Tank/Neo** — these are TS-only edits.

---

## Phase 5 — Report schema v3 (Go data-model bump, sibling to Phase 2)

**Owner:** Tank (validator/report layer) and Neo (engine wiring), Morpheus reviews.
**Depends on:** Phase 2.1 (`GraderPoint` in data model) — points must exist before we drop the expansion.
**Why a sibling phase, not extra Phase 2 todos:** these changes touch the **report layer** (`internal/report/types.go`, `convertGraderResults`) and the **persisted JSON shape**, not the live-eval renderer that Phase 2 targets. Bundling them into Phase 2 would mix concerns and make Tank's mid-execution work harder to reason about. They wait on Phase 2 landing the upstream `Points` field, then ship as a focused commit.

### 5.1 — Bump `CurrentSchemaVersion` to v3

Old v2 reports keep their N-entry expanded shape on read (no de-expansion attempted — too lossy). v3 reports use the new 1-entry-with-Points shape. Add `MigrateToV3` no-op stub for forward compatibility.

- File: `hyoka/internal/report/types.go` (look for `CurrentSchemaVersion` and `MigrateToV2` as the template).

### 5.2 — Kill `expandReviewGraderResult`

Replace with a single-entry mapping for `KindPromptReview`:

```go
report.GraderResult{
    GraderName: "ai_review",
    GraderType: "prompt_review",
    Pass:       &aggregatedPass,        // AND of all panel decisions
    Score:      aggregatedScore,
    Points:     []GraderPoint{...},     // one per criterion
    ReviewDetails: <existing struct, unchanged>,
}
```

`ReviewDetails.PanelResults` and `ReviewDetails.Criteria` stay populated for backward compat with the static Markdown/HTML templates.

- File: `hyoka/internal/eval/engine_eval.go:840-953` (the whole `convertGraderResults` + `expandReviewGraderResult` block).

### 5.3 — Add `Points` to `report.GraderResult`

```go
type GraderResult struct {
    // ... existing fields ...
    Points []GraderPoint `json:"points,omitempty"`
}
```

`omitempty` keeps old reports byte-identical when re-encoded. Mirrors the upstream `graders.GraderResult.Points` Phase 2.1 added.

- File: `hyoka/internal/report/types.go` (look for the existing `GraderResult` struct).

### 5.4 — Add parent linkage to `report.ToolLoadResult`

```go
type ToolLoadResult struct {
    Name       string `json:"name"`
    Status     string `json:"status,omitempty"`     // ← omitempty: parents omit
    Error      string `json:"error,omitempty"`
    Details    string `json:"details,omitempty"`
    Kind       string `json:"kind,omitempty"`        // "skill" | "mcp" | "plugin" | "skill_dir"
    Parent     string `json:"parent,omitempty"`      // container this is a child of
    ParentKind string `json:"parent_kind,omitempty"` // "plugin" | "skill_dir"
}
```

Parents emit a row with `Status` empty; children carry runtime status + back-pointer. Old reports remain valid (all new fields `omitempty`). Coordinates with Tank's Phase 1.2 (plugin parent emits no status) — same idea persisted to disk.

- File: `hyoka/internal/report/types.go:326-341`.

### 5.5 — Enrich `EnvironmentInfo.SkillsLoaded` with parent linkage

Decision needed (see Open Questions): **Option A** — change `SkillsLoaded` from `[]string` to `[]struct{Name, Parent, Kind string}` and migrate site types. **Option B** — keep flat list, add sibling `SkillGroups []struct{Parent string; Children []string}`.

Morpheus's recommendation: **Option A** — expresses the truth, single source. Site migration is mechanical.

- Files: `hyoka/internal/report/types.go` (`EnvironmentInfo.SkillsLoaded`), `site/src/app/data/types.ts:180-192`.

### 5.6 — Optional: pre-computed roll-up fields on `EvalReport`

Add `EvalReport.GradersPassed int` and `EvalReport.GradersTotal int`, populated at engine time from `agg.Results` BEFORE expansion. Site reads these directly instead of recomputing. Eliminates the entire class of roll-up-divergence bugs by construction. Trivial, `omitempty`-safe.

- File: `hyoka/internal/report/types.go`, populated in `engine_eval.go` near where `Success` is set.

### 5.7 — Migration & test plan

- Snapshot a v2 report to `testdata/`, assert v3 code reads it without panic and renders the legacy expanded shape.
- Generate a fresh v3 report from a live eval, assert one `ai_review` entry with `Pass=true` and N `Points`.
- `go build ./...` and `go test ./hyoka/...` green.

---

## Phase 6 — Site Phase-2 alignment (depends on Phase 2 + Phase 5)

**Owner:** Trinity.
**Depends on:** Phase 2 complete (Points exist on the wire) AND Phase 5 complete (schema v3 — single `ai_review` entry, `parent_kind/parent_name` populated, optional roll-up fields).
**Why:** Once Points lands and the expansion is killed, the site templates need an upgrade to render `Points[]` properly, group plugin children under parents, and unify roll-up logic across every page.

### 6.1 — Add Points + parent-linkage types to TS

- `site/src/app/data/types.ts` — add `GraderPoint { name: string; pass: boolean; message?: string }`, add `points?: GraderPoint[]` to `GraderResult`, add `parent_kind?` / `parent_name?` / `kind?` to tool entries, mirror `EnvironmentInfo.SkillsLoaded` change from Phase 5.5.

### 6.2 — Render `Points[]` in `GraderResultRow`

Header badge: `✓ N/N passed` when all pass, `✗ M/N passed` otherwise, `N/A` only when no points and `pass` is null. Expanded: one indented row per point with check/X icon, name, optional message. Falls back to today's flat row when `len(points) <= 1`.

- File: `site/src/app/components/GraderResultRow.tsx`. Markup sketch in Trinity §4b — reuse existing `ml-6` indent + criterion pill chrome from `eval-detail-page.tsx:615-660`.

### 6.3 — Plugin parent grouping in Available Tools

Group `tool_availability` entries by `parent_kind/parent_name`. Parent renders as a header line (no status pill); children render as indented pill list with `border-l border-white/5 pl-3` chrome. Skill-dir parent uses `entry.Name` (not path) — already routed through by Phase 1.1 + Phase 5.4.

- File: `site/src/app/components/eval-detail-page.tsx:495-535`. Markup sketch in Trinity §4a.

### 6.4 — Single `isEvalPass()` helper, applied everywhere

Replace per-page roll-ups with one canonical helper:

```ts
function evalPassFromPoints(r: EvalReport): boolean {
  if (r.grader_results?.length) {
    return r.grader_results.every(g =>
      (g.points?.length ?? 0) === 0
        ? g.pass === true
        : g.points!.every(p => p.pass));
  }
  return r.success === true;
}
```

Apply to: `runs-page.tsx`, `run-detail-page.tsx` (replaces the Phase 4.1 stop-gap), `eval-detail-page.tsx` header, `prompt-detail-page.tsx`, `pairwise-page.tsx`, `comparison-page.tsx`. Once the rule lives in one place no surface can drift again.

- Files: above. New helper in `site/src/app/lib/` (Trinity's call on exact location).

### 6.5 — Eval-detail score-card replaces the `12/12` review-only metaphor

Per Trinity §4d: dominant card shows `✓ 15 / 15 points across 4 graders` (AND of every Point). Tooltip on the run-detail row shows the per-grader breakdown. Pre-condition: Open Question #2 answered (graders-passed vs. points-passed as the dominant number).

- File: `site/src/app/components/eval-detail-page.tsx:382-461`.

### 6.6 — Prompt-detail tool grouping

`Pass Rate by Tool Used` aggregates tools as flat names; `azure.sdk_*` siblings appear as separate rows. Group by `parent_name` once Phase 5.4 lands.

- File: `site/src/app/components/prompt-detail-page.tsx`.

### 6.7 — Visual regression check

Drive `playwright-cli` against the same anchor run (`reports/20260423-195948`) post-Phase-6: the run-detail rows should be emerald, the eval-detail grader rows should show points expanded, the dashboard should not crash.

---

## Open questions (collated from both reviews — for the user)

These need answers before Phase 5/6 lock in. None block Phase 4.

1. **Run-detail row badge — what number does the user actually want?** (a) just a pass/fail icon, (b) score breakdown percentage, (c) `passed/total` fraction treating nil-pass review entries as passing, (d) post-Phase-2 `points-passed/points-total`. Morpheus's pick: (a) for clarity, (b) most informative. Trinity's pick: graders-passed at the run level, points-passed at the eval level (two-tier). _(Morpheus §6 Q1, Trinity §6 Q2)_

2. **`success=true` with zero graders run** — bug or feature? `engine_eval.go:433` skips grading when no files generated; `Success` defaults to true. Three configs in the anchor run hit this. Options: hard-fail at engine, OR mark `graders_skipped: true` on the report so site renders distinctly. _(Morpheus §6 Q2)_

3. **Phase 5 schema bump** — OK to bump `SchemaVersion` to v3 and leave v2 reports in their N-entry expanded shape forever (no migrator)? Alternative is a re-collapse migrator that loses panel-member detail unless we carefully reconstruct. _(Morpheus §6 Q3)_

4. **Plugin parent in `session_setup`** — when Tank's Phase 1 lands, does the parent appear in JSON at all (as `Status:""` row consumers can use to draw the group header), or only via children's `Parent` back-pointer? The former is friendlier; the latter is simpler. _(Morpheus §6 Q4)_

5. **Canonical pass for review graders.** Should the engine just write `pass = AND(criteria)` for review graders? Or does consensus use a different rule (e.g. quorum)? Determines whether Phase 4.1/4.2 fallbacks are temporary or load-bearing. _(Trinity §6 Q1)_

6. **Score-card metaphor on the eval page** post-Points: "points passed / points total" (granular, true) or "graders passed / graders total" (concise, less true). Which is the dominant number? _(Trinity §6 Q2 — overlaps with #1)_

7. **In-progress runs on `/runs`** — hide partial dirs without `summary.json`, or render as `In progress`? _(Trinity §6 Q3)_

8. **Errored runs on `/runs`** — split rate column into `passed | failed | errored`, or fold errors into denominator with amber bar? _(Trinity §6 Q4)_

9. **Dashboard page priority** — bundle the crash fix into Phase 4 (already there as 4.3) or carve a separate session for it? _(Trinity §6 Q5 — answered by current plan: bundled into Phase 4.)_

10. **Plugin parent click target** — expandable/collapsable (hide 41 children behind `▶`) or always-expanded? `azure-sdk-python` takes ~6 visual rows when fully expanded. _(Trinity §6 Q6)_

11. **`EnvironmentInfo.SkillsLoaded` shape** — Phase 5.5 Option A (typed objects with parent linkage) vs. Option B (flat list + sibling `skill_groups`). Morpheus's recommendation: A. _(Morpheus §3c)_
