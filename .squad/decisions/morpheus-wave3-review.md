# Wave 3 Architectural Review — PRs #581 / #582 / #583

**Author:** Morpheus 🕶️
**Date:** 2026-04-17
**Base:** `squad/phase-4-remainder`
**Scope:** Phase 4 final wave — Trinity #361 (serve cache + pairwise endpoint), Trinity #360 (pairwise page polish + proxy chart), Neo #357 (unified ComparisonResult + auto-gen).

---

## TL;DR

| PR | Verdict | Blocker |
|----|---------|---------|
| **#581** (#361 serve cache + pairwise) | **APPROVE** | — |
| **#582** (#360 pairwise polish)        | **APPROVE** | — |
| **#583** (#357 comparison unification) | **REQUEST_CHANGES** | Go↔TS drift: comparison page will render `undefined` labels after merge. One-file fix. |

**Phase 4 gate: NO-GO as currently stacked.** Flip to GO once #583's TS drift is fixed and Switch re-runs coverage on the rewritten `comparison/` package. ETA to GO: hours, not days.

---

## PR #581 — Trinity #361 serve cache + pairwise endpoint + security docs

**Verdict: APPROVE.**

Clean, scoped, well-tested. Cache design is correct for the workload.

### What's right

- `cache.go:44-61` — `(mtime, size)` fingerprint is the right primitive for a report-dir filesystem. Writers (new eval runs landing on disk) transparently bust entries with zero explicit invalidation. Error path at `cache.go:46-49,57-60` evicts stale entries so a transient failure can't poison subsequent reads. This is the safe design.
- `cache.go:15` — `sync.RWMutex` with a read-lock for lookups, release, re-acquire write-lock for population. No upgrade races — a concurrent writer just re-does the stat+read, which is idempotent. 50×20 concurrent reader test at `cache_test.go:108-139` under `-race` covers the mutex discipline.
- `serve.go:131` — cache is scoped to `buildMux`, NOT package-global. Tests and concurrent `Start()` calls stay isolated. Correct design.
- `serve.go:3-31` — security docblock calls out exactly what it should: no auth, permissive CORS, exposes prompt text + generated source. Startup banner at `serve.go:119-122` makes it impossible to miss. This satisfies the Phase 4 security documentation requirement.
- `api.ts:72-79` — `fetchPairwise` returns `null` on 404 rather than throwing. Correct pattern: "no pairwise data" is expected, not an error. Keeps caller code linear.
- `types.ts:226-234` — `PairwiseRunReport` JSON field names match Go struct tags verbatim (`run_id`, `timestamp`, `reports`, `aggregate_impacts`). No `as unknown as` casts anywhere. This is the Wave 1 lesson landing.

### Minor

- `dashboard.go:24` — `_ = cache` placeholder for future comparison endpoint caching. Fine as a marker, but note that after #583 merges and adds `handleAPIRunComparisons`, that handler will bypass the cache (reads `comparisons.json` directly via `os.ReadFile`). Not a regression, but post-merge rebase should wire it through.

### Merge coordination note (not blocking #581)

This PR conflicts with #583 on `hyoka/internal/serve/serve.go` — both add a new case to the `handleAPIRunDetail` switch. Verified via `git merge-tree`. Whichever merges second needs a trivial rebase to add its case alongside the other's. #581's handler signature change (cache threading) also needs to propagate into #583's `handleAPIRunComparisons` on rebase. **Suggest merge order: #583 first (stable type foundation), then rebase #581 on top** so cache wiring covers the comparisons endpoint in one pass.

---

## PR #582 — Trinity #360 pairwise methodology + tool usage chart + deep-link

**Verdict: APPROVE.** Ship it for 0.3.1.

### Tool Usage Frequency chart — accept the proxy

**Question 2 adjudicated: proxy is acceptable for 0.3.1.**

The reasoning is correct. At `pairwise-page.tsx:368-378`, "used" is defined as `imp.impact !== 0 || imp.baseline_pass !== imp.without_pass`. That's impact signal, not observation. Ground truth lives on per-eval `EvalReport.ToolAvailability`, not aggregated in summary.json.

Three reasons to accept:

1. **Labeled.** `pairwise-page.tsx:347-355` comment block explicitly calls it a proxy and directs readers to eval detail for ground truth. `pairwise-page.tsx:672-675` ships the same disclaimer user-visible. No false-precision claim is being made.
2. **Correct scope.** Adding a backend aggregation endpoint would widen #360 by at least another backend handler + TS type + tests. Out of scope for the pairwise-page polish issue.
3. **Actionable follow-up.** Easy to replace later: when a `GET /api/runs/{id}/tool-availability` endpoint exists, swap `computeToolFrequency` to read from it and delete the proxy logic. Zero-risk future change.

**File a follow-up issue** — "Backend tool-availability aggregation endpoint + ground-truth frequency chart" — target 0.3.2 or 0.4.

### What else is right

- `pairwise-page.tsx:483,502-510` — `useSearchParams` for `?run=<id>` is correctly two-way bound with `{replace: true}`, so back-button history doesn't pollute. Deep-link from run-detail → pairwise works.
- `run-detail-page.tsx:153-172` — conditional render on `run.pairwise_results?.length > 0`. `encodeURIComponent(run.run_id)` correctly handles slash-containing run IDs.
- `MethodologyInfo` at `pairwise-page.tsx:297-353` — collapsible state-driven rather than `<details>`, keeps design-system consistency. Positive/negative/zero impact semantics spelled out with color coding. Good.

### Nit (not blocking)

- PR description says "amber" banner on run-detail but `run-detail-page.tsx:155` uses emerald. Either is fine; description cosmetic.

---

## PR #583 — Neo #357 comparison unification

**Verdict: REQUEST_CHANGES.**

The Go architecture is excellent. The TS side was not updated. One file-fix unblocks merge.

### What's right (architecturally significant)

- `result.go:40-48` — `ComparisonResult` with discriminator `Kind` replaces three wrapper types. Right call. Kind-switched rendering at `compare.go:92-101` collapses three near-identical render functions into one. Net `-347` lines.
- `comparison.go:167-187` — `CompareReports` is the in-memory primitive. Every public entry point (`CompareConfigs`, `CompareRuns`, `TemporalDiff`, `AutoGenerateForRun`) delegates to it. This is exactly the shape I asked for in the kickoff brief §3 Neo guidance.
- `comparison.go:189-220` — `AutoGenerateForRun` emits pairs alphabetically (`sort.Strings(configs)` + `i<j` nested loop). Deterministic. Test at `inmem_test.go:99-126` proves it.
- `autogen.go:24-51` — `WriteForRun` writes `comparisons.json` alongside `summary.json`. The file-adjacent decision over adding `Comparisons []ComparisonResult` to `RunSummary` is correct — the `comparison → report → comparison` import cycle is real, and a file is easier to evolve than a struct field. Noted in the PR body.
- `engine.go:706-721` — auto-gen wiring is defensive. Failure is `slog.Warn`, not a run failure. Correct.
- `dashboard.go:171-222` — comparisons endpoint computes on-demand for legacy runs that pre-date auto-generation. Legacy compatibility done right.
- `inmem_test.go:181-225` — `TestAutoGenerateForRun_UsesSameEngine` proves auto-gen output matches direct-path output byte-for-byte. This is real coverage of the shared-core claim.

### 🔴 BLOCKER: TS drift — comparison page will break on merge

`site/src/app/data/types.ts:309-314` still declares:

```ts
export interface ConfigComparison {
  config_a: string;
  config_b: string;
  per_prompt: PromptDiff[];
  summary: ComparisonSummary;
}
```

After this PR merges, `GET /api/compare/configs` returns `{kind: "configs", label_a, label_b, per_prompt, summary}` — the field rename is breaking. Confirmed at `comparison-page.tsx:546,549`:

```tsx
<TableHead>{comparison.config_a}</TableHead>  // → undefined
<TableHead>{comparison.config_b}</TableHead>  // → undefined
```

**The comparison page will render blank column headers as soon as this merges.** This is the exact class of Go↔TS drift I flagged in Wave 1. Tests don't catch it because `fetchCompareConfigs` is mocked in `__tests__/api.test.ts` and `comparison-page.test.tsx`.

Same drift affects `/api/compare/runs` and `/api/compare/temporal` (see `dashboard_test.go:226-240,295-310` — tests were updated, but TS consumers were not).

### Punch list for Neo (one PR revision)

1. **Update `site/src/app/data/types.ts:309-314`** — replace `ConfigComparison` with a union that matches the Go discriminator:
   ```ts
   export type ComparisonKind = "configs" | "runs" | "temporal";

   export interface ComparisonResult {
     kind: ComparisonKind;
     label_a: string;
     label_b: string;
     config?: string;    // only for kind === "temporal"
     since?: string;     // only for kind === "temporal"
     per_prompt: PromptDiff[];
     summary: ComparisonSummary;
   }
   ```
   Keep `ConfigComparison` as a deprecated alias for one release OR remove and update call sites — either works; pick one and commit.
2. **Update `site/src/app/data/api.ts:62-66`** — change return type to `ComparisonResult`.
3. **Update `site/src/app/components/comparison-page.tsx`** — rename `comparison.config_a` → `comparison.label_a`, `config_b` → `label_b` (lines 231, 257, 546, 549, and state type at line 231).
4. **Update `site/src/__tests__/comparison-page.test.tsx`** — fixture field names.
5. **Add TS mirror for the new `/api/runs/{id}/comparisons` endpoint** — `fetchRunComparisons(runId): Promise<ComparisonResult[]>` in `api.ts`. The site consumer of this endpoint lands in Phase 5, but publishing the TS type now closes the contract.

### Other item: `summary_stats.PromptDeltas` retention

**Question 3 adjudicated: SIGN-OFF, retain as-is for 0.3.1.**

Neo's justification is legitimate on both grounds:

1. **Different semantic.** `summary_stats.PromptDeltas` is a per-prompt pass/fail toggle between configs — a binary transition aggregated across a run. `comparison.PromptDiff` is a numeric score-delta. Unifying them means losing the pass-toggle view or coercing it into score-delta (which reports a 1.0 delta for a pass-flip — lossy and misleading).
2. **Import cycle.** `comparison` already imports `report`. `report/summary_stats.go` importing `comparison` would close the cycle. Breaking it would require a third package, which is scope creep for this PR.

Keep. File a Phase 5 cleanup issue: *"Move pass-toggle aggregate into a shared view package and delete `summary_stats.PromptDeltas`."*

---

## Cross-cutting items

### 1. Gate: "CLI and site produce identical results" — shared-core argument is sufficient

**Accept the construction-based claim for 0.3.1.** All three surfaces funnel through `comparison.CompareReports`:

- CLI: `CompareConfigs(reportsDir, ...)` → `loadReportsByConfig` → `CompareReports` → `renderComparison(out, cmp)`
- Serve `/api/compare/configs`: same `CompareConfigs` → `writeJSON(w, cmp)`
- Serve `/api/runs/{id}/comparisons`: `LoadForRun` OR `loadRunReports + AutoGenerateForRun` → `writeJSON(w, results)`
- End-of-run auto-gen: `eval/engine.go:710` → `comparison.WriteForRun` → `AutoGenerateForRun` → `CompareReports`

All four paths share the same map-build (`latestByPrompt` or `indexByPromptConfig`) and `buildSummary` code. `inmem_test.go:181-225` already proves equivalence between the auto-gen and direct in-memory paths. The disk-backed paths differ only in how reports are loaded — same reports, same result.

**Coordinate with Switch** — recommend a snapshot test in `hyoka/internal/comparison/` that:
1. Builds two fixture runs on disk.
2. Calls `CompareConfigs(dir, "a", "b")` → marshals to JSON (S1).
3. Loads the same reports manually and calls `CompareReports` directly → marshals (S2).
4. `assert.Equal(S1, S2)` byte-for-byte.

This is belt-and-suspenders, not a gate blocker. Ship Wave 3 without it; file as a test-hardening follow-up for Switch.

### 4. Contract coherence across PRs

**Clean on the Go side. Drift on the TS side (see #583 blocker).**

Go boundary:
- #583 publishes `ComparisonResult` in `hyoka/internal/comparison/result.go` and exposes `/api/runs/{id}/comparisons`.
- #581 exposes `/api/runs/{id}/pairwise` with a separate type (`report.PairwiseRunReport`) — no overlap with `ComparisonResult`.
- #582 consumes only existing `PairwiseReport` / `ToolImpact` types from Wave 1 — no new dependency on #583's types.

The two endpoints are orthogonal. No type coupling between #581 and #583 beyond the file they both touch (`serve.go` switch).

TS boundary:
- #581 adds `PairwiseRunReport` TS type correctly matching Go.
- #583 does NOT add a TS mirror for `ComparisonResult` AND does not update the existing `ConfigComparison` type, which will now be stale. **This is the drift.**

Fix #583's punch list, contract is clean.

### 5. Phase 4 go-live gate — status

Running down the kickoff brief §6 criteria:

| # | Criterion | Status |
|---|-----------|--------|
| 1 | All 9 items closed (#355–#363, +#566) | #566 ✅ #355/#356 ✅ #357 🟡 (#583 in review) #358 ✅ #359 ✅ #360 🟡 (#582 in review) #361 🟡 (#581 in review) #362 ✅ #363 ✅ |
| 2 | Switch's test review green; `go test -race`, `npm test` | ⚠️ Switch's #375 audit at `switch-375-test-review.md` reports `comparison 91.6%`. That coverage is against the **pre-#583** package. Post-#583 the package is rewritten (+137/-125 lines on `comparison.go`, two new files). **Switch must re-run coverage after #583 lands.** |
| 3 | Site renders real eval data | ✅ achieved in Wave 2 (#358, #359) |
| 4 | `hyoka compare` CLI and site page produce identical results | ✅ satisfied by #583 shared-core, **once TS drift is fixed** |
| 5 | Branding updated (general AI, zero Azure SDK) | ✅ #362 landed |
| 6 | Examples pass `go run ./hyoka validate` | ✅ #363 landed |
| 7 | All Phase 4 work merged to `ronniegeraghty/dev`, CI passing | 🟡 pending consolidation PR `phase-4-remainder → dev` after Wave 3 |

### Go-live decision

**NO-GO until:**

1. **#583** revises per punch list (TS types + comparison page). Non-negotiable — the comparison page visibly breaks on merge otherwise.
2. **Switch** re-runs `go test -race -cover ./hyoka/internal/comparison/...` after #583 lands on `squad/phase-4-remainder`. Must remain ≥85%. Quick spot-check; not a full audit re-run.
3. **Merge order:** #583 first → Switch re-runs → #581 rebased + merged → #582 merged (trivially clean, site-only).
4. **Consolidation PR** `squad/phase-4-remainder → ronniegeraghty/dev` opened with CI green.

After those: **GO for 0.3.1.** Phase 4 complete.

---

## Appendix: decisions to file

- **A1.** Accept pairwise tool usage chart as proxy for 0.3.1; file follow-up for backend aggregation endpoint (0.3.2 or 0.4).
- **A2.** Retain `summary_stats.PromptDeltas` for 0.3.1; file Phase 5 cleanup for shared view package.
- **A3.** Add Switch follow-up: snapshot test proving CLI JSON ≡ serve API JSON ≡ auto-gen JSON for identical inputs. Belt-and-suspenders over the shared-core argument.
- **A4.** Add Go↔TS drift rule to squad conventions: "When a Go response type's JSON field names change, the matching TS type in `site/src/app/data/types.ts` MUST be updated in the same PR." This is now the second wave where this class of bug appeared (Wave 1: `workspace_delta` field-name mismatch; Wave 3: `config_a`/`config_b` → `label_a`/`label_b`). Pattern, not incident.

—Morpheus
