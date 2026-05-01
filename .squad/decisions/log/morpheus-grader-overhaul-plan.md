## Status: SHIPPED (2026-05-01)
# Morpheus — Grader Overhaul Plan

**Author:** Morpheus 🕶️
**Date:** 2026-05-01
**Status:** SCOPING (not yet implemented)
**Owners:** Neo (eval/graders/review), Tank (CLI/site backend), Trinity (frontend), Switch (test fixtures)
**Tasked by:** Ronnie

---

## ⚠️ TOP CONCERN — Determinism fix is INCOMPLETE

The `deterministic-llm-panel` skill claims the X/Y check counts are stable
across runs. They are not, and Ronnie is still seeing drift. The stable-ID
work landed correctly, but **two independent sources of non-determinism
survived**:

### Bug A — `averageReview` builds the criterion order from observed votes, not from `expected`

`hyoka/internal/review/reviewer.go:683-755` (`averageReview`) walks
`panel[].Scores.Criteria` to seed `criteriaOrder` and `criteriaMap`. Anything
no reviewer voted on **simply disappears from `MaxScore`**. The `expected
[]ReviewCheck` argument is used only for canonical labels and never to anchor
the result set.

### Bug B — Whole-reviewer drop and per-bucket failure mutate the panel size

Two paths can silently shrink what gets voted on:

1. `runSingleReview` (`reviewer.go:582-650`) **returns an error after
   `maxRetries` (2) validation failures**, dropping the entire reviewer's
   contribution. With LLM variance, run A may parse cleanly while run B
   exhausts retries → different panel sizes → strict any-fail voting flips.
2. `ReviewPanelBuckets` (`buckets.go:128-153`) silently drops a single
   (model, bucket) result on `rerr != nil`. If both panel models fail the
   same bucket, all checks in that bucket vanish from `criteriaOrder`.
   Different runs → different bucket-failure patterns → different `MaxScore`.

### Why the earlier "byte-identical smoke test" missed it

The test ran the simplest path — `combined` mode, both reviewers succeeded,
no retries triggered. The smoke test never exercised the failure paths
that produce drift. This is the gap Ronnie keeps re-hitting.

### Required correction (Ask 1)

`expected []ReviewCheck` is the **single source of truth** for the result
set. The vote stage must:

1. Anchor `criteriaOrder` to `expected`, not to whatever the panel returned.
2. Keep retrying a single reviewer on missing/extra IDs (3 chances, not 2).
3. After 3 strikes, **synthesize a failing `CriterionResult` for each
   missing ID for that reviewer** instead of dropping the reviewer.
4. After all reviewers, any `expected` ID with zero votes is marked
   **failed by default**, not omitted. (Defensive: should not happen once
   step 3 lands, but cheap insurance.)

That kills both bugs at once and makes X/Y deterministic for any combination
of bucket, retry, and panel-failure outcomes.

---

## Investigation findings

### Ask 1 — Determinism (root cause above)

**Files involved:**
- `hyoka/internal/review/reviewer.go` (`runSingleReview`, `averageReview`,
  `parseReviewResponseV2`)
- `hyoka/internal/review/buckets.go` (`ReviewPanelBuckets`,
  `mergeBucketResults`)
- `hyoka/internal/review/review_test.go` (existing tests pass `nil` for
  `expected` to `averageReview` — they do not exercise the new contract)
- `.squad/skills/deterministic-llm-panel/SKILL.md` (must update once fix lands)

### Ask 2 — Grader uniformity & consolidation

**What already exists:**
- `graders.GraderResult` *is* the canonical shape. Every grader returns
  `{Kind, Name, Weight, Gate, Score, Pass, Message, Points[], Extras,
  SourceFile, SourceType}`. The `Points []GraderPoint` invariant is
  enforced (`grader.go:264-275` panics if no points).
- The `Extras` discriminated union holds kind-specific render-only data.
  This is fine — it is render-only and doesn't affect score/pass.

**What is overlapping / redundant:**

*Workspace-output graders (3 overlapping types):*
| Kind | Purpose | Data source | Overlap |
|---|---|---|---|
| `file` | Single-file existence + regex | filesystem walk | Subset of `output_check` |
| `output_check` | Multi-knob delta-based file checks | `WorkspaceDelta` | Canonical |
| `program` (when used as `test -f hello.md`) | Existence via shell exit code | `os/exec` | Hack — exact subset of `output_check.require_files` |

*Tool-perspective graders (4 overlapping types):*
| Kind | Purpose | Data source | Overlap |
|---|---|---|---|
| `behavior` | required/forbidden tool names + max_turns | `ActionLog` | Subset of `tool_constraint` |
| `tool_constraint` | required / forbidden / min_calls / max_calls | `ActionLog` | Closest to canonical, but no group support |
| `tool_usage` | env-aware skill/MCP rules with silent skip | `EnvironmentTools` + `SkillsInvoked` + `MCPServersUsed` | Different signal source — must merge carefully |
| `action_sequence` | ordered prefix of expected actions | `ActionLog` | Distinct (sequence semantic) |

**Recommended consolidation (DEFAULT):**

- **Workspace grader:** keep `output_check` as the one. Deprecate `file`
  and recommend `output_check.require_files`/`min_bytes_per_file` instead.
  Keep `program` as-is — it is genuinely a "run a build/test command"
  grader, not a file-existence shim. Provide a one-time deprecation
  warning when `file` graders are loaded.
- **Tool grader:** introduce one consolidated `tool` grader with
  per-rule sub-checks supporting:
  - `kind: any_of_group` (groups: `mcp`, `skill_plugin`, `skill_repo`,
    `tool_name_glob`)
  - `kind: specific_tool` (one named tool)
  - `kind: tool_not_used` (forbidden)
  - `kind: group_not_used` (no tool from group)
  - `kind: action_sequence` (ordered prefix — absorbs the existing
    grader so we have ONE tool-perspective grader)
  - `kind: turn_limit` (max_turns)
  Deprecate `behavior`, `tool_constraint`, `tool_usage`, `action_sequence`
  with one-release shims that internally translate to the new form.
- Action_sequence is borderline — moving it under `tool` is a stretch.
  **DEFAULT: keep `action_sequence` separate** — it's a different semantic
  (ordering vs presence). Consolidate the other three only.

*Provenance / canonical "report" shape:* `GraderResult` already carries
`SourceFile` + `SourceType` (set by the engine after `Grade()` returns,
`engine_eval.go:784-785`). No new struct needed — Ronnie's "uniform report
data shape" already exists; the work is *naming and documentation*, not
restructuring.

### Ask 3 — Pairwise site clarity

**What exists:**
- `hyoka/internal/pairwise/pairwise.go` — generates `baseline + N
  without-tool` variants. Outputs go through normal eval pipeline.
- `hyoka/internal/serve/serve.go:256-263` (`handleAPIPairwise`) —
  serves the run's `pairwise.json`.
- `site/src/app/components/pairwise-page.tsx` (765 lines!) — already has
  `ToolImpact`, `classifyTool` (helper/hurter/neutral), `HeatmapCell`,
  `ImpactSummaryCard`. The data shape supports the analysis already.

**What's broken / unclear:**
- The pairwise *run detail* view doesn't surface a per-eval comparison
  that says "this check went from PASS in baseline to FAIL in
  without-skill-X". It shows aggregate impact only.
- The diff between two specific evals (baseline vs. one variant) is not
  rendered side-by-side at the check level.
- The classifier thresholds are arbitrary (impact > 1, < -1) and are not
  shown to users.

**Required:** A "tool diff" view per pairwise run that highlights, per
check: improved / hurt / no-effect, sorted by magnitude.

### Ask 4 — Final "X/Y" line UX

**Where it's emitted:**
- Per-eval failed line: `engine.go:777` — `fmt.Sprintf("%d/%d points",
  passed, total)` becomes `evt.Message`.
- Per-eval passed line: no message currently shown (just the green check).
- Run summary: `progress/display.go:348`, `display.go:606`,
  `display_interactive.go:1376` — `Summary: %d/%d passed` (this counts
  *evals*, not checks).
- Per-grader badge: `display_interactive.go:1072-1074` — `❌ Fail (%d/%d)`
  / `✅ Pass (%d/%d)` (counts *points/checks within one grader*).

**Required:**
- Add blank line above the per-eval terminal line.
- Replace `%d/%d points` with `Total checks that passed across all
  graders: %d/%d`.
- Always show the line on PASS too, not only on FAIL (currently passed
  evals silently skip the message).

### Ask 5 — Rename "grader points" → "checks"

**Code touch points (full sweep):**
- `hyoka/internal/criteria/graders/grader.go` —
  `GraderPoint` → `GraderCheck` (avoid colliding with `review.ReviewCheck`),
  `Points` → `Checks`. `NewResult(... points []GraderPoint ...)` signature.
- `hyoka/internal/criteria/graders/{behavior,file,program,output_check,
  prompt_review,prompt_grader_adapter,tool_usage,tool_constraint,
  action_sequence}_grader.go` — every `GraderPoint{...}` literal.
- `hyoka/internal/criteria/graders/grader.go:354` — `AggregateResults`
  comments / log messages.
- `hyoka/internal/eval/engine_eval.go:1240,1255` —
  `countTotalPoints` → `countTotalChecks`,
  `countPassedPoints` → `countPassedChecks`. Comments.
- `hyoka/internal/eval/engine_eval.go:1681-1695`,
  `progress/display_interactive.go`, `progress/display.go`,
  `progress/display_ci.go` — `progress.GraderPoint` → `progress.GraderCheck`,
  `result.Points` → `result.Checks`, all "points" log/format strings.
- `hyoka/internal/report/types.go:58` — `report.GraderPoint` →
  `report.GraderCheck`. Schema bump? **DEFAULT: no schema bump**, keep
  JSON field tag as `"points"` for one release for back-compat (tag
  comment notes the rename, code uses `Checks`). Add a follow-up commit
  to flip the tag once site reads both.
- `hyoka/internal/report/markdown.go`, `report/report_data.go` — display
  strings.
- `hyoka/internal/trends/analysis.go` — display strings.
- All `_test.go` files — update literals + names.
- `site/src/app/data/types.ts` — `points` field renamed to `checks` (with
  back-compat type alias for one release).
- `site/src/app/components/*.tsx` — anywhere reading `result.points`.

**YAML grader taxonomy ("checks:" property):**
- `prompt` graders already use `checks:` (`config.go:60`).
- Typed graders DO NOT have a `checks:` list — their config IS the
  check spec (e.g. `output_check.require_files` → one check per path).
  **DEFAULT:** do not invent a `checks:` list for typed graders. Keep
  the existing typed-config style. The user's intent is unified
  *result-shape* naming, which the rename of `Points` → `Checks`
  achieves. If a future grader genuinely lists discrete sub-checks
  (e.g. the new consolidated `tool` grader), use `checks:` in YAML.

**Other YAML consistency wins (free upgrades while we're touching the schema):**
- `details.min_bytes_per_file` and friends use `snake_case` — consistent.
- `Weight` defaults to 1.0 silently (config.go:223) — surface in docs.
- `Gate` field still on `GraderConfig` even though Phase 2 removed
  short-circuit behavior (`grader.go:347-353`). **DEFAULT: leave the
  field for one more release** with a deprecation notice in
  `criteria/config.go` + `docs/grader-config-schema.md`.
- `program.timeout` is bare seconds; everything else takes durations as
  Go duration strings would be better, but **DEFAULT: don't change** —
  not worth a breaking schema bump for one field.

### Ask 6 — Workflow

Commits must:
- Be small (one ask = one logical group).
- Build cleanly (tests passing) at every commit.
- Have a verification step (build / unit tests / live eval).
- Re-run live evals **after each ask** to confirm no regression.

---

## Phased commit plan (15 commits)

> Conventional commit prefixes: `fix:` for bug fixes, `refactor:` for
> renames, `feat:` for new behavior, `chore:` for tests/docs.

### ASK 1 — Determinism (highest priority — Neo)

**C1 — `fix(review): anchor consensus vote to expected check IDs`**
- Files: `hyoka/internal/review/reviewer.go` (`averageReview`), `review/review_test.go`
- Why: `criteriaOrder` must come from `expected []ReviewCheck`, not from
  the union of observed votes. Missing observations become explicit
  failures.
- Verify: `go test ./hyoka/internal/review/...` and a new regression test
  that simulates a panel where reviewer 2's bucket failed → consensus
  still has all expected checks; missing ones are `Passed: false`.
- Owner: **Neo**

**C2 — `fix(review): retry-3-then-mark-failed for missing check IDs`**
- Files: `hyoka/internal/review/reviewer.go` (`runSingleReview`)
- Why: Replace "drop reviewer after 2 retries" with "3 retries, then
  synthesize failing `CriterionResult` for each missing ID and KEEP the
  reviewer". Eliminates panel-size variance.
- Verify: New unit test: mock a reviewer that returns invalid IDs three
  times → assert reviewer is retained with the missing IDs marked as
  failed. Re-run determinism smoke (`hyoka run --prompt-id
  test-dp-test-hello-markdown --config test/baseline`) **5×** and
  verify byte-identical X/Y across all five.
- Owner: **Neo**

**C3 — `chore(review): drop dead text-keyed legacy parser path`**
- Files: `hyoka/internal/review/reviewer.go` (`parseReviewResponse`,
  `parseReviewResponseV2`'s "fall back to legacy" branch)
- Why: With C1+C2, the legacy parser is unreachable for the panel path
  and is itself non-deterministic (paraphrase keying). Removing it
  prevents future regressions.
- Verify: `go build ./...`, `go test ./...`. Search for callers — none
  outside tests after the change.
- Owner: **Neo**

**C4 — `docs(skill): update deterministic-llm-panel skill with retry-then-fail contract`**
- Files: `.squad/skills/deterministic-llm-panel/SKILL.md`
- Why: Skill currently advertises "retry-then-drop" (see `Implementation
  References` line 51). Update to reflect new contract + the
  `expected`-anchored vote.
- Verify: skill markdown lints (no automated test).
- Owner: **Neo** (skill body) / **Morpheus** signs off.

### ASK 2 — Grader consolidation + uniform shape (Neo)

**C5 — `refactor(graders): introduce GraderCheck (rename GraderPoint), keep Points alias one release`**
- Files: `hyoka/internal/criteria/graders/grader.go`, every grader
  implementation, `report/types.go`, `progress/*.go`,
  `eval/engine_eval.go`, all `_test.go`. Site types get a parallel
  `checks` field that aliases `points` (TS).
- Why: Set up the rename with deprecation aliases so the rest of Ask 5
  can land in pieces without breaking the world.
- Verify: `go build ./...`, `go test ./...`, `cd site && npm test`.
- Owner: **Neo** (Go) + **Tank** (TS site types alias)

**C6 — `feat(graders): consolidate tool + behavior + tool_usage into tool grader`**
- Files: new `hyoka/internal/criteria/graders/tool_grader.go`,
  deprecation shims in `behavior_grader.go`, `tool_constraint_grader.go`,
  `tool_usage_grader.go`, `criteria/config.go` `validTypedKinds`,
  `registry.go`.
- Why: One canonical tool-perspective grader. The three legacy kinds
  remain valid YAML but log a deprecation warning and translate to the
  new schema at load time.
- Verify: `go test ./hyoka/internal/criteria/graders/...`, run one live
  eval against `criteria/language/test.yaml` and confirm tool-related
  output identical to before.
- Owner: **Neo**

**C7 — `feat(graders): mark file grader as deprecated alias of output_check`**
- Files: `hyoka/internal/criteria/graders/file_grader.go`, registry.
- Why: `file` is a strict subset of `output_check.require_files` +
  `min_bytes_per_file`. Add deprecation warning at load; keep code path
  alive one release.
- Verify: `go test`. No live eval changes.
- Owner: **Neo**

**C8 — `feat(report): document SourceFile/SourceType as canonical provenance`**
- Files: `hyoka/internal/criteria/graders/grader.go` (godoc on
  `GraderResult`), `docs/grader-config-schema.md` if present.
- Why: Ronnie's "uniform report data shape" already exists; this commit
  documents it so future contributors know not to invent parallel structs.
- Verify: `go vet ./...`, doc render.
- Owner: **Neo** (or Tank — pure docs)

### ASK 5 — Rename sweep finish + YAML wins (Neo + Tank)

**C9 — `refactor(graders): switch all internal call sites from Points to Checks`**
- Files: `hyoka/internal/criteria/graders/*_grader.go`,
  `eval/engine_eval.go`, `progress/*.go`, `report/types.go`,
  `report/markdown.go`, `report/report_data.go`, `trends/analysis.go`.
- Why: After the alias from C5, drop `Points` references in core code.
  Display strings flip from "points" to "checks". `countTotalPoints` →
  `countTotalChecks`.
- Verify: `go build ./... && go test ./...`. Live eval — confirm CLI
  output reads "checks" everywhere.
- Owner: **Neo**

**C10 — `refactor(report,site): rename JSON field points→checks (keep both keys one release)`**
- Files: `hyoka/internal/report/types.go` (custom marshaler that emits
  *both* `"checks"` and `"points"` for one release), `site/src/app/data/types.ts`,
  `site/src/app/components/*.tsx`.
- Why: Avoid breaking already-on-disk reports. Site reads `checks` first,
  falls back to `points`.
- Verify: `go test`, `npm test`. Open serve UI against an old report dir
  and a new one — both render.
- Owner: **Tank**

**C11 — `chore(yaml): document deprecation of gate field on GraderConfig`**
- Files: `hyoka/internal/criteria/graders/types.go` (godoc), docs.
- Why: Gate has been a no-op since Phase 2. Deprecation note steers
  users away.
- Verify: doc build / no behavior change.
- Owner: **Tank**

### ASK 4 — UX of "X/Y" line (Tank)

**C12 — `feat(progress): show "Total checks ... X/Y" line for every eval`**
- Files: `hyoka/internal/progress/display_interactive.go`
  (`onPassed` / `onFailed` — both must print the line),
  `progress/display.go`, `progress/display_ci.go`, `eval/engine.go`
  (build the message with the new wording).
- Why: User wants explanatory text and a blank line above the line, on
  pass and on fail.
- Verify: Run `hyoka run --prompt-id test-dp-test-hello-markdown
  --config test/baseline --progress interactive` and visually confirm.
  Snapshot tests in `progress/display_interactive_points_test.go` need
  updating.
- Owner: **Tank**

### ASK 3 — Pairwise site clarity (Tank backend, Trinity frontend)

**C13 — `feat(serve): expose per-check pairwise diff in pairwise.json`**
- Files: `hyoka/internal/pairwise/pairwise.go` (or a new
  `pairwise_diff.go`), `hyoka/internal/serve/serve.go`,
  `report/types.go` (new `PairwiseCheckDiff` type).
- Why: Frontend currently only has tool-level impact. To show which
  checks moved between baseline and each variant, the API must surface
  per-check deltas.
- Verify: `go test ./hyoka/internal/pairwise/...`. New unit test
  comparing two synthetic eval reports → diff has expected entries.
- Owner: **Tank**

**C14 — `feat(site): pairwise tool-impact view shows check-level diff`**
- Files: `site/src/app/components/pairwise-page.tsx`,
  `site/src/app/data/types.ts`.
- Why: Render `PairwiseCheckDiff` as a sortable list per variant:
  improved / regressed / unchanged, with check label + reason. Include
  classifier thresholds in a small legend.
- Verify: `npm run build`, `npm test`. Eyeball against a fresh pairwise
  run produced after C1–C13 land.
- Owner: **Trinity**

### ASK 2/5 — Test fixture (Switch)

**C15 — `chore(criteria): rebuild criteria/language/test.yaml with one of each remaining type`**
- Files: `criteria/language/test.yaml`,
  `prompts/test/hello-markdown.prompt.md` (only if criteria changes
  invalidate the prompt).
- Why: After consolidation (C6/C7), the test fixture must use the new
  canonical kinds: one `prompt`, one `output_check`, one `tool`, one
  `program`, one `action_sequence`. (Drop `behavior`, `tool_constraint`,
  `tool_usage`, `file`.)
- Verify: `hyoka run --prompt-id test-dp-test-hello-markdown --config
  test/baseline` 3× → byte-identical X/Y. Confirm new "checks" wording
  appears in CLI output.
- Owner: **Switch**

---

## Sequencing notes

```
C1 ─┐
C2 ─┼─ Ask 1 (must land before any rerun-based verification on Asks 2–5)
C3 ─┤
C4 ─┘

C5 (sets up alias) ─┬─ C6 (tool consolidation)
                    ├─ C7 (file deprecation)
                    └─ C8 (docs only)
                          │
                          ▼
                         C9 (rename sweep)
                          │
                          ▼
                  C10 (JSON dual-emit) ──┐
                                         ├── Tank can run C11/C12
                                         │   in parallel with C13
                                         ▼
                                     C13 (pairwise API)
                                         │
                                         ▼
                                     C14 (Trinity: site UI)

C15 (Switch) — gated on C6+C7 landing; can land any time after.
```

**Parallelism opportunities:**
- C7 and C8 can land alongside C6 (independent files).
- C11 + C12 are independent of C13/C14.
- C14 (Trinity) only needs the API contract from C13 — it can begin
  against a stub before C13 merges.
- C15 (Switch) only depends on C6/C7 being merged.

**Critical-path ordering:** C1 → C2 → C5 → C9 → C10. Everything else can
be reordered around that spine.

---

## Open questions (DEFAULT decisions made)

1. **Should `action_sequence` be folded into the new `tool` grader?**
   - **DEFAULT: NO.** Keep `action_sequence` separate — it has ordering
     semantics that don't fit the presence/absence model of `tool`.
2. **JSON schema bump for `points` → `checks` field rename?**
   - **DEFAULT: NO bump this release.** Emit both keys; consume `checks`
     first, `points` as fallback. Follow-up release drops the old key.
3. **Rename `GraderPoint` to `Check` (short) or `GraderCheck` (qualified)?**
   - **DEFAULT: `GraderCheck`.** Avoids collision with
     `review.ReviewCheck` and reads clearly at call sites.
4. **Should the deprecated `gate` field be removed now?**
   - **DEFAULT: NO.** Document deprecation only; remove next release.
5. **Should the per-eval line print on PASS too, or only on FAIL?**
   - **DEFAULT: YES print on both.** Ronnie wants the explanatory text
     consistently visible.
6. **Site back-compat horizon for `points` vs `checks`?**
   - **DEFAULT: 1 release of dual-emit, then drop `points`.**
7. **Update `deterministic-llm-panel` skill or split into a new skill?**
   - **DEFAULT: update in place.** The pattern is the same; only the
     retry/fail mechanic changed.

---

## Verification matrix

| Commit | Build | Unit tests | Live eval | Site build |
|---|---|---|---|---|
| C1, C2 | ✅ | ✅ + new regression | 5× determinism smoke | — |
| C3, C4 | ✅ | ✅ | — | — |
| C5     | ✅ | ✅ + alias coverage | — | ✅ |
| C6     | ✅ | ✅ + new tool_grader_test | 1× test/baseline | — |
| C7     | ✅ | ✅ | — | — |
| C8     | ✅ | — | — | — |
| C9     | ✅ | ✅ | 1× | — |
| C10    | ✅ | ✅ | — | ✅ |
| C11    | ✅ | — | — | — |
| C12    | ✅ | snapshot updates | 1× (visual) | — |
| C13    | ✅ | ✅ | 1× pairwise (`hyoka run --pairwise`) | — |
| C14    | — | ✅ npm | — | ✅ |
| C15    | — | — | 3× determinism | — |

After all 15 land: one full multi-prompt eval (e.g. `hyoka run --service
key-vault --language python --config test/baseline`) and an end-to-end
pairwise run reviewed in the serve UI.

---

## Hand-off

- **Neo:** owns C1, C2, C3, C4, C5, C6, C7, C8, C9.
- **Tank:** owns C10, C11, C12, C13.
- **Trinity:** owns C14.
- **Switch:** owns C15.
- **Morpheus:** review every PR, sign off on the determinism re-runs in
  C2 and C15, update `deterministic-llm-panel` skill once C2 lands.
