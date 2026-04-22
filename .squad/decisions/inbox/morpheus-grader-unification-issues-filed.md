### 2026-04-23: Grader Unification — Issues filed, Phase 1 ready for Neo

**Author:** Morpheus 🕶️ (Lead Architect)
**Status:** EXECUTING (Phase 1 unblocked)
**Supersedes:** #622

All ten open questions on the unified grader architecture proposal (`morpheus-grader-unification-proposal.md`) are answered. Schema is locked: flat `type` discriminator at entry level (no `kind`, no nested block), prompt graders share the same shape as every other grader, no `Gate` field anywhere (graders never gate an eval — they run and report), file-level validation errors only fail evals when that file matches the prompt, no deprecation shim for `internal/criteria/`. Four phase issues are filed against `ronniegeraghty/hyoka` and labeled `squad` + `squad:neo`. Issue #622 has been closed with a comment linking the new plan. **Phase 1 (#624) is ready for Neo to start.** Phases 2-4 are sequential and blocked until each predecessor lands.

**Issues:**

| # | Title | URL |
|---|-------|-----|
| 624 | [Grader Unification] Phase 1: Unified schema + back-compat loader in internal/graders/ | https://github.com/ronniegeraghty/hyoka/issues/624 |
| 625 | [Grader Unification] Phase 2: Unified execution path in engine (cut over to internal/graders/) | https://github.com/ronniegeraghty/hyoka/issues/625 |
| 626 | [Grader Unification] Phase 3: Delete internal/criteria/ package | https://github.com/ronniegeraghty/hyoka/issues/626 |
| 627 | [Grader Unification] Phase 4: Default output_check criteria + per-grader docs | https://github.com/ronniegeraghty/hyoka/issues/627 |

**Next action:** Coordinator hands #624 to Neo. Reference docs Neo will need: the full proposal, the directive batch file, and this decision.
