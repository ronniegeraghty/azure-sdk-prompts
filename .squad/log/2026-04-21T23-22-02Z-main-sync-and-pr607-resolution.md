# Session Log: Main Sync and PR #607 Conflict Resolution (2026-04-21T23:22:02Z)

**Date:** 2026-04-21  
**User Request:** Pull origin/main into both dev and phase-6 branches; switch docs/ to installed-binary command form  
**Agents:** Tank (dev merge), Neo (PR #607 resolution)  
**Duration:** Two-agent parallel orchestration  

## Executive Summary

User Ronnie requested main branch be pulled into both active development branches and docs/ be converted to installed-binary commands. Tank merged main into dev (independent conflict resolution); Neo then resolved PR #607 by merging dev into phase-6, keeping phase-6's architectural improvements while adopting dev's corrected paths.

## Outcomes

✅ **both branches now 19+ commits ahead of main**
✅ **PR #607 transitioned from DIRTY/CONFLICTING to CLEAN/MERGEABLE**
✅ **docs/ all switched to `hyoka` (installed-binary form) — 28 replacements in getting-started.md**
✅ **build clean, all 24 test packages pass with -race**

## Conflict Resolution Pattern

Tank and Neo independently resolved the same 9 upstream conflicts with different semantic choices, creating a secondary conflict in PR #607. Neo's resolution:

| Aspect | Tank (dev) | Neo (phase-6, final) | Reason |
|--------|-----------|-----------------|--------|
| Skill fetching | Direct npx | Pluggable Fetcher + context.Context | Extensibility + cancellation propagation |
| README path | `go run ./hyoka` (fixed) | `go run . ` (fixed via dev) | Stability vs internal restructure |
| Test style | Various | Cleaner cosmetics (phase-6) | Pre-existing phase-6 pattern |

**Key insight:** Two-branch independent merges of the same upstream requires semantic resolution during down-merge. Blind conflict-picker tools fail here.

## Decisions Captured

1. **Docs installed-binary directive:** All docs/ examples now use `hyoka <cmd>`, never `go run .`. Source-dev commands belong in CONTRIBUTING.md only.
2. **PR #607 resolution:** Kept phase-6's extensible architecture while correcting paths from dev.
3. **Routing note (informal):** Future docs work should route to Oracle by default (not Tank).

## Git State

- **ronniegeraghty/dev:** 8bfc4da2 (Merge main)
- **phase-6:** 25675461 (Merge dev; now ahead of main by 19 commits)
- **CI status:** All checks passing

## Files Changed (Summary)

- **Merge files:** 6 conflicts resolved (engine.go, copilot.go, workspace.go, go.work*, go.sum)
- **Docs:** 28 `go run . ` → `hyoka ` replacements in docs/getting-started.md
- **.squad/:** Decision and orchestration logs recorded

## Related

- **Orchestration logs:** 
  - `.squad/orchestration-log/2026-04-21T23-22-02Z-tank.md`
  - `.squad/orchestration-log/2026-04-21T23-22-02Z-neo.md`
- **Decisions:**
  - `.squad/decisions.md` (PR #607 resolution strategy + docs installed-binary directive)
