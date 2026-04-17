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

### Critical path

**#566 → #355 + #356 → #357 → #358 → #359**

- Neo's #566 stabilizes core types, then #355/#356 (parallel), then #357.
- Trinity can start #358 immediately (depends only on #344, which is done), but #357's `ComparisonResult` type will feed into #361's API endpoints. 
- #359 is hard-blocked on #358 — same page patterns, same component library.
- #360, #361, #362 are independently startable, but #361's comparison endpoints should wait for #357's unified interface.

### Does Trinity block on Neo?

**Partially.** #358 (Eval Detail) does NOT block on Neo's Phase 4 work — it depends on #344 (done). But #361 (Serve Updates) needs Neo's #357 for comparison endpoints. Trinity should start #358 and #362 immediately.

### Does Oracle block on anything?

**No.** #363 depends on #353 (done). Oracle can start immediately.

---

## 3. Per-Agent Launch Plan

### Neo 💊

**Start NOW:** #566 (WorkspaceDelta) — 2-day time-box. See §4 for rationale.

**Then:** #355 (Review Session Modes) and #356 (Hierarchical When) — parallel or sequence.

**Last:** #357 (Comparison Unification) — depends on #355/#356.

**Architecture guidance:**
- #566: New package `hyoka/internal/workspace/`. Add `WorkspaceDelta` field to both `EvalReport` and `GraderInput`. Keep grader changes out of scope.
- #355: Review session mode logic in `internal/review/`. The `--review-mode` flag wires through `cmd/run.go` → engine options → review orchestration.
- #356: Hierarchical `when` in `internal/criteria/criteria.go`. Merge semantics (grader > group > file) deterministic and documented.
- #357: Unify into `internal/comparison/`. Expose `ComparisonResult` as public type — Trinity's serve endpoints will import it.

### Trinity 🖤

**Start NOW:** #358 (Eval Detail Redesign, XL) + #362 (Site Content & Branding).

**Parallel:** #360 (Pairwise in Site), #361 (Serve Updates — caching layer first; defer comparison endpoints until #357).

**After #358:** #359 (Run Detail + Runs Page).

**Architecture guidance:**
- #358: Eval detail page reads from `EvalReport` JSON. Establish `GraderResultRow` component pattern — #359 and #360 will reuse.
- #361: In-memory cache with `os.Stat` mtime checking. Wait for Neo's `ComparisonResult` type from #357.
- #360: Pairwise integration through serve API. Tool usage frequency chart is key visualization.
- #362: Logo, homepage, footer, How It Works. Zero references to "code generation" or "Azure SDK".

### Oracle 🔮

**Start NOW:** #363 (Examples Update) — no blockers.

**Guidance:**
- `examples/configs/example-full.yaml` must use `tools:`, include `reviewer:` section.
- Delete `examples/configs/example-registry.yaml`.
- YAML prompt example should demo `properties:` map with `when` conditions.
- Run `go run ./hyoka validate` on all files before done.

### Switch 🤍

**Ongoing:** #375 (Phase 4 Test Review).

**Immediate:** Tests for #566 (delta computation, guardrails, output presence).

**Then:** Tests for #355/#356 as Neo moves to those.

**Gate:** `go test -race ./hyoka/...` and `npm test` after every merge.

---

## 4. #566 WorkspaceDelta Decision

**Recommendation: Option A — Neo does #566 first, before #355–#357.**

**Rationale:** #566 adds `WorkspaceDelta` to both `EvalReport` and `GraderInput` — core types that Neo's and Trinity's Phase 4 work all read from. Landing #566 first stabilizes these types before broader phase begins. Scope is tight (new package + wiring, no grader behavior changes) and time-boxable to 2 days.

**Time-box:** 2 days hard cap. If guardrail softening proves complex, ship type + wiring only.

---

## 5. Risk Register

| # | Risk | Mitigation |
|---|------|------------|
| 1 | **Trinity carries 5/9 items.** Burnout if #358 runs long. | #362 and #360 can shift to Tank as overflow. |
| 2 | **#358 is XL and linchpin.** #359 and Phase 5 depend on patterns. | Morpheus reviews at 50% mark. |
| 3 | **Serve API mismatch.** Trinity consumes endpoints; Neo changes model. | Neo publishes `ComparisonResult` signature first. |
| 4 | **#566 overrun.** If >2 days, phase slips. | Hard cap; defer guardrail thresholds to Phase 4.5. |
| 5 | **#357 scope creep.** Touches 3 comparison modes. | Neo defines interface first, gets approval, then implements. |

---

## 6. Go-Live Gate

Phase 4 complete when ALL true:

1. All 9 items closed (#355–#363, + #566)
2. Switch's test review green (#375 checklist; `go test -race`, `npm test`)
3. Site renders real eval data (not mock)
4. `hyoka compare` CLI and site page produce identical results
5. Branding updated (general AI evaluation, zero Azure SDK)
6. Examples pass `go run ./hyoka validate`
7. All Phase 4 work merged to `ronniegeraghty/dev`, CI passing

---

*Brief archived from decisions inbox.*
