# Phase 4 Kickoff Brief

**Author:** Morpheus 🕶️  
**Date:** 2026-04-17  
**Tracking issue:** #310  
**Phase 3 landed via:** PR #562 → `ronniegeraghty/dev` (commit `4b4e95f9`)

---

## 1. Phase 4 North Star

Rebuild the site and reporting experience to consume the unified grading pipeline shipped in Phase 3. Win condition: every site page renders real grader data from a recent eval run, the comparison engine is a single code path used by both CLI and site, and hyoka's public-facing content presents it as a general-purpose AI agent evaluation tool — not an Azure SDK code-gen checker.

---

## 2. Dependency Graph

### What's already done (Phase 2/3 prerequisites — all merged to `ronniegeraghty/dev`)

| Prereq | Status | Unblocks |
|--------|--------|----------|
| #344 WI-023 Unified Grading Pipeline | ✅ Phase 3 | #355, #356, #357, #358, #360, #361, #362 |
| #341 WI-043 Eliminate HTML Reports | ✅ Phase 2 | #361 |
| #353 WI-042 Prompt Package Updates | ✅ Phase 3 | #363 |

### Phase 4 DAG

```
                ┌──────────────────────────────────────┐
                │  #566 WorkspaceDelta (Neo, pre-phase) │
                └──────────┬───────────────────────────┘
                           │ stabilizes EvalReport/GraderInput
                           ▼
            ┌─────────┐  ┌─────────┐
            │ #355    │  │ #356    │
            │ Review  │  │ Hier.   │    ┌─────────┐
            │ Session │  │ When    │    │ #363    │
            │ Modes   │  │ Criteria│    │ Examples│  ← Oracle, independent
            └────┬────┘  └────┬────┘    └─────────┘
                 │            │
                 └─────┬──────┘
                       ▼
                 ┌──────────┐
                 │ #357     │
                 │ Compare  │
                 │ Unify    │
                 └──────────┘
                       │ exposes ComparisonResult for site
                       ▼
              ┌─────────────────────────────────────────┐
              │         Trinity's site work              │
              ├────────────┬──────────┬────────┬────────┤
              │ #358       │ #360    │ #361   │ #362   │
              │ Eval Detail│ Pairwise│ Serve  │ Content│
              │ (XL, focus)│ in Site │ Updates│ Brand  │
              └─────┬──────┘         │        │        │
                    │                │        │        │
                    ▼                │        │        │
              ┌──────────┐          │        │        │
              │ #359     │          │        │        │
              │ Run Detail│         │        │        │
              │ + Runs Pg │         │        │        │
              └──────────┘          │        │        │
                                    ▼        ▼        ▼
                              (all can start in parallel)
```

### Critical path

**#566 → #355 + #356 → #357 → #358 → #359**

- Neo's #566 stabilizes core types, then #355/#356 (parallel), then #357.
- Trinity can start #358 immediately (depends only on #344, which is done), but #357's `ComparisonResult` type will feed into #361's API endpoints. Trinity should design #358 against current report data and adapt when comparison lands.
- #359 is hard-blocked on #358 — same page patterns, same component library.
- #360, #361, #362 are independently startable once #344 is done (now), but #361's comparison endpoints should wait for #357's unified interface.

### Does Trinity block on Neo?

**Partially.** #358 (Eval Detail) does NOT block on Neo's Phase 4 work — it depends on #344 (done). But #361 (Serve Updates) needs Neo's #357 (Comparison Unification) for the comparison API endpoints. Trinity should start #358 and #362 immediately, do #361's caching layer now but defer comparison endpoint updates until #357 lands.

### Does Oracle block on anything?

**No.** #363 depends on #353 (done). Oracle can start immediately.

---

## 3. Per-Agent Launch Plan

### Neo 💊

**Start NOW:** #566 (WorkspaceDelta) — 2-day time-box. See §4 for rationale.

**Then:** #355 (Review Session Modes) and #356 (Hierarchical When) — these can be worked in parallel or sequence, both unblocked.

**Last:** #357 (Comparison Unification) — depends on #355/#356 completing per the tracking issue dependency chain.

**Architecture guidance:**
- #566: New package `hyoka/internal/workspace/`. There's already `workspace.go` / `workspace_test.go` in `hyoka/internal/eval/` from the hotfix — decide whether to promote those to the new package or replace them. Add `WorkspaceDelta` field to both `EvalReport` (in `report/report_data.go`) and `GraderInput` (in `graders/`). Keep grader changes out of scope — data availability only.
- #355: Review session mode logic goes in `internal/review/`. The `--review-mode` flag wires through `cmd/run.go` → engine options → review orchestration. Don't touch the `Grader` interface — this is about how the review grader organizes its sessions internally.
- #356: Hierarchical `when` lives in `internal/criteria/criteria.go`. The merge semantics (grader > group > file) must be deterministic and documented in a comment block. Other packages should not need to change.
- #357: Unify into `internal/comparison/`. Delete the comparison logic in `report/summary_stats.go` (currently `summary_stats.go` + `summary_stats_test.go`). Expose a `ComparisonResult` struct as a **public type** — Trinity's serve endpoints will import it. Auto-generate comparison data during multi-config runs by calling the unified engine from `engine_eval.go`. Coordinate with Trinity on the `ComparisonResult` shape before implementing.

### Trinity 🖤

**Start NOW:** #358 (Eval Detail Redesign) — this is XL and the critical path for site work. Also #362 (Site Content & Branding) — pure content, no API dependency.

**Parallel (once #358 design is stable):** #360 (Pairwise in Site), #361 (Serve Updates — start with caching layer; defer comparison endpoints until Neo ships #357).

**After #358:** #359 (Run Detail + Runs Page) — reuses component patterns from #358.

**Architecture guidance:**
- #358: The eval detail page reads from `EvalReport` JSON. The grader results table replaces the old rubric badges. Components go in `site/src/app/components/`. Establish a `GraderResultRow` component pattern here — #359 and #360 will reuse it.
- #361: Serve package is at `hyoka/internal/serve/`. Add an in-memory cache keyed by file path with `os.Stat` mtime checking for invalidation. The API endpoints must return the new grader data structures from Phase 3's unified pipeline. For comparison endpoints, wait for Neo's `ComparisonResult` type from #357 — don't invent a parallel data model.
- #360: Pairwise data currently lives in separate JSON files. Integration means reading it through the serve API and displaying on the eval detail and pairwise pages. The tool usage frequency chart (available/unused/used) is the key new visualization.
- #362: Logo, homepage, footer, How It Works. The How It Works pipeline is now: Prompt → Agent Session → Graders → Summary/Insights → Reports. Remove all references to "code generation", "Azure SDK", "rubric criteria", "AI consolidation".

**Shared boundary with Neo:** The `ComparisonResult` type from `internal/comparison/` (Neo's #357) will be consumed by serve endpoints (Trinity's #361). Neo defines the type, Trinity consumes it. Agree on the struct shape before either side implements.

### Oracle 🔮

**Start NOW:** #363 (Examples Update) — no blockers.

**Guidance:**
- `examples/configs/example-full.yaml` must use `tools:` (not `mcp_servers:`), include `reviewer:` section, and pass `go run ./hyoka validate`.
- Delete `examples/configs/example-registry.yaml`.
- The YAML prompt example (`.prompt.yaml`) should demonstrate the `properties:` map format with `when` conditions.
- The starter files example should follow the pattern from commit `92bfa067` (existing-files prompt example already merged to main).
- Run `go run ./hyoka validate` against all example files before considering done.

### Switch 🤍

**Ongoing:** #375 (Phase 4 Test Review) — start writing tests as work items land.

**Immediate:** Write tests for #566 (WorkspaceDelta) since Neo starts there. Cover: delta computation (new/modified/deleted files, byte counts), guardrail warnings (not hard-fails), and `WorkspaceDelta` field presence in `EvalReport` JSON output.

**Then:** Tests for #355/#356 as Neo moves to those. For #355: test both `combined` and `isolated` modes, per-grader `isolate: true` override. For #356: test three-level `when` inheritance with override semantics.

**For site work:** Write Vitest component tests as Trinity establishes component patterns in #358. Key: `GraderResultRow` component renders all grader types; `EvalDetailPage` handles empty/missing data gracefully.

**Gate:** Run `go test -race ./hyoka/...` and `cd site && npm test` after every merge to dev. Report gaps immediately — don't accumulate.

---

## 4. #566 WorkspaceDelta Decision

**Recommendation: Option A — Neo does #566 first, before #355–#357.**

**Rationale:** #566 adds `WorkspaceDelta` to both `EvalReport` and `GraderInput` — the two core types that Neo's Phase 4 work (#355, #357) and Trinity's site work (#358) all read from. Landing #566 first means these types are stable before the broader phase begins. If we defer, #566 lands mid-phase and forces everyone to rebase around schema changes to `EvalReport`. The scope is tight (new package + wiring + guardrail softening, no grader behavior changes) and time-boxable to 2 days. It also subsumes the starter-aware `MaxOutputSize` hotfix already on dev — completing the guardrail story cleanly.

**Time-box:** 2 days hard cap. If guardrail softening proves complex, ship the type + wiring first and defer threshold changes to a follow-up.

---

## 5. Risk Register

| # | Risk | Mitigation |
|---|------|------------|
| 1 | **Trinity carries 5/9 items.** Burnout or bottleneck if #358 runs long. | #362 (content) and #360 (pairwise) are independent and can be deprioritized or handed to Tank as overflow. |
| 2 | **#358 (Eval Detail) is XL and the linchpin.** #359 and Phase 5 pages depend on component patterns established here. | Trinity starts #358 day 1. Morpheus reviews component architecture at the 50% mark before Trinity builds more pages on top. |
| 3 | **Serve API contract mismatch.** Trinity's site consumes serve endpoints; Neo changes the underlying data model. | Neo publishes `ComparisonResult` struct definition before implementing #357. Trinity codes against that contract. Any schema change requires a Morpheus-approved interface review. |
| 4 | **#566 time-box overrun.** If WorkspaceDelta takes >2 days, entire phase slips. | Hard time-box. If guardrail softening is complex, ship type + wiring only. Guardrail thresholds move to Phase 4.5. |
| 5 | **#357 scope creep.** Comparison Unification touches config, run, and temporal comparison — three modes in one issue. | Neo defines the unified interface first (types + function signatures), gets Morpheus sign-off, THEN implements. No cowboy coding on this one. |

---

## 6. Go-Live Gate

Phase 4 is done when ALL of the following are true:

1. **All 9 work items closed:** #355, #356, #357, #358, #359, #360, #361, #362, #363. Plus #566 if adopted per §4.
2. **Switch's test review green:** #375 checklist complete, `go test -race ./hyoka/...` passes, `cd site && npm test` passes.
3. **Site renders real data:** Eval detail, run detail, runs list, and pairwise pages all render data from an actual eval run on `ronniegeraghty/dev`. Not mock data — real report JSON from `go run ./hyoka run`.
4. **Comparison works end-to-end:** `hyoka compare` CLI command and site comparison page both produce identical results for the same two-config run.
5. **Branding updated:** Homepage, How It Works, footer, and logo reflect general-purpose AI agent evaluation — zero mentions of "Azure SDK code generation".
6. **Examples pass validation:** `go run ./hyoka validate` succeeds on all files in `examples/`.
7. **PR merged to dev:** All Phase 4 work on `ronniegeraghty/dev` branch, passing CI.

---

*Brief ready for review. Awaiting Ronnie's approval before fanning out to the team.*
