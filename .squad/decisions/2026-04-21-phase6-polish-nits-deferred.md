# Phase 6 Polish Nits — Deferred to Follow-Up Issues

**Date:** 2026-04-21
**Author:** Scribe (recording from Switch + Neo review comments)
**Status:** Decided — deferred
**Context:** Phase 6 polish round 2 (PRs #609 / #610 / #611 / #612, all merged into `phase-6`)

## Decision

Reviewers (Switch on #609, Neo on #611) flagged non-blocking findings during the polish round. None block #607 (rollup) or `phase-6 → dev`. They are deferred to dedicated follow-up issues to be filed against the `phase-6 → dev` merge target, NOT carried as silent debt.

## Items deferred

### From PR #609 (Switch on Trinity's `MultiSelectFilter` tests)

- Toggle / `onChange` behavior is not exercised — highest-value miss for a "MultiSelect" component. Ship a follow-up that drives selection through `userEvent.click` and asserts the `onChange` payload.
- Summary-text branches (single / multi / overflow) untested.
- `aria-expanded` and `aria-selected` not asserted.
- "Click inside the listbox keeps it open" counterpart to the outside-click test is missing.

Keyboard listbox navigation is a **product gap** for Trinity (component behavior), not a test gap, and routes separately.

### From PR #611 (Neo on Morpheus's CI freshness gate)

- `verify-embed` overlaps with `ci.yml`'s `site-build-and-test` job (both run `npm ci && npm run build`). Either fold the verify into that job to save ~1–2 min per run, or accept the discrete PR check status as worth the duplication. Decision pending after a real-world cost observation.
- Workflow's `push` triggers list `phase-6` — must be pruned once the branch retires.
- Add `concurrency:` group to cancel superseded runs on the same PR.
- Morpheus's own three nits on the Makefile: `git diff --quiet` won't catch new untracked files inside the embed dir; `rm -rf assets/` won't prune stale root-level files in the embed tree; same `concurrency:` note.

## Rationale

All four polish PRs cleared the bar (8 ✅ verdicts, 1 ⚠️ APPROVE WITH NOTES — none REQUEST CHANGES). Holding the merges to address style-grade and follow-up-grade findings would have risked stalling #607. The findings are real and named here so they don't evaporate; the next person filing Phase 7 issues should turn the lists above into tickets.

## Owners

- MultiSelectFilter test gaps → Trinity (test) + product gap (Trinity, separate)
- CI freshness workflow polish → whoever picks up `#608` follow-up (likely Morpheus or Neo)
