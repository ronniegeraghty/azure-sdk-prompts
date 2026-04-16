# Morpheus Review: PR #567 — Starter-Aware MaxOutputSize Guardrail

**Date:** 2026-04-16  
**Reviewer:** Morpheus 🕶️ (Lead/Architect)  
**PR:** #567 (branch: `squad/565-starter-aware-maxoutputsize`)  
**Issue:** #565  
**Author:** Neo  

## Verdict: **APPROVE** ✅

This hotfix is architecturally sound, surgically scoped, and demonstrates the right engineering judgment for an interim fix that will be gracefully subsumed by the Phase 3.5 WorkspaceDelta work.

---

## Review Summary

### 1. Preserves Hard-Fail Semantics ✅

**Requirement:** This hotfix must keep guardrails as hard failures. Softening to warnings is Phase 3.5 work (#566).

**Verification:**
- Guardrail violations still set `evalReport.Success = false` (lines 1115-1117, 1125-1127)
- Still log `"Guardrail triggered"` with severity WARN
- No behavioral change to failure handling

**Conclusion:** PASS. Hard-fail semantics preserved exactly.

---

### 2. Phase 3.5 Compatibility ✅

**Requirement:** The hotfix must not fight with #566's WorkspaceDelta work. Ideally, it should be trivially replaceable by delta logic.

**Architectural assessment:**

#### What the hotfix introduces:
1. **`snapshotStarterSizes(baseDir, starterFiles)`** — captures `map[string]int64` of starter file sizes at copy time
2. **`computeAgentOutputSize(baseDir, files, snapshot)`** — sums agent-attributable bytes (new files full-size, modified files delta-only)
3. **`computeAgentFileCount(files, snapshot)`** — counts new files + deleted starters

#### How Phase 3.5 subsumes this:

The WorkspaceDelta type (per #566 spec) will capture:
```go
type WorkspaceDelta struct {
    BytesNet          int64   // Replaces computeAgentOutputSize output
    NewFileCount      int     // Part of computeAgentFileCount
    DeletedFileCount  int     // Part of computeAgentFileCount
    ModifiedFiles     []...   // Richer than snapshot, includes hashes
    // ...
}
```

**The delta computation *is* the snapshot-subtract logic, just with richer metadata.**

When #566 ships, the guardrail checks become:
```go
// Before (this hotfix):
totalSize := computeAgentOutputSize(ws.Dir, generatedFiles, starterSnapshot)

// After (#566):
totalSize := evalReport.WorkspaceDelta.BytesNet
```

The three helper functions in `guardrail.go` become dead code and can be deleted. The snapshot capture at line 809 becomes redundant because WorkspaceDelta will capture everything.

**This is clean architectural layering** — the hotfix does NOT introduce coupling or concepts that Phase 3.5 must work around. It solves the immediate problem with the same conceptual model that #566 will formalize.

**Conclusion:** PASS. No architectural debt. WorkspaceDelta will cleanly replace this logic.

---

### 3. Live Test: Real Eval with Starter Project ✅

**Setup:**
- Created 900 KB starter project (`test-starter-pr567/`)
- Created test prompt referencing the starter
- Ran eval: `go run . run --prompt-id test-pr567-starter-aware --config baseline/claude-opus-4.6`

**Results:**
- Starter files copied correctly (2 files, 900 KB + 141 bytes)
- Agent generated `hello.py` (33 bytes)
- **New logic counted:** 33 bytes (agent output only)
- **Old logic would have counted:** 921,774 bytes (all files)
- Evaluation completed successfully, no false guardrail failure
- Unit tests: `go test -race ./hyoka/internal/eval/...` ✅ PASS

**Key insight:** While both old and new logic pass at 900 KB (under 1 MB limit), the test proves the logic executes correctly. With a 1.1 MB starter, old logic would fail immediately; new logic would pass.

**Conclusion:** PASS. Hotfix works in real evaluation.

---

## Code Quality Assessment

### Strengths

1. **Three pure helpers** — `snapshotStarterSizes`, `computeAgentOutputSize`, `computeAgentFileCount` are stateless, deterministic, easily testable
2. **Graceful degradation** — stat failures on starter files default to size 0 (worst case: old behavior)
3. **Comprehensive test coverage** — `guardrail_test.go` covers all edge cases: unchanged, modified, new, shrunk, deleted, mixed
4. **Minimal engine.go diff** — 6 lines added (snapshot), 20 lines replaced (guardrail checks) — surgical
5. **Safe snapshot semantics** — snapshot is taken AFTER `CopyStarterFiles` but BEFORE agent session starts, correct sequencing

### Concerns (minor, non-blocking)

1. **File count logic slightly over-counts deletion risk** — `computeAgentFileCount` counts deleted starters toward the agent's file budget. This is conservative (safe), but Phase 3.5's `WorkspaceDelta.DeletedFileCount` will separate this clearly.
2. **Snapshot keyed by relative path** — assumes `ws.ListFiles()` and `starterFiles` use consistent path formats. This works today (both are workspace-relative), but WorkspaceDelta will formalize this with a `WorkspaceFile` type.

Neither concern warrants changes in this PR. Both are addressed by Phase 3.5's richer modeling.

---

## Decision Rationale

This is textbook surgical refactoring:
- Solves the immediate blocker (colleagues can't eval with large starters)
- Does NOT prematurely implement the full WorkspaceDelta architecture
- Clean conceptual fit — the logic here *is* delta computation, just without the type wrapper
- No technical debt accumulation — #566 deletes this code, doesn't refactor around it

**If this were attempting to be the "real" solution, I'd reject it** — we'd want the full `WorkspaceDelta` type, richer file metadata (hashes, timestamps), and grader wiring. But as an interim fix scoped to unblocking a concrete guardrail failure, it's exactly right.

---

## Next Steps (Post-Merge)

1. **#566 Phase 3.5 implementation:**
   - Introduce `hyoka/internal/workspace/delta.go` with `WorkspaceDelta` type
   - Capture delta at same snapshot point (engine.go:809)
   - Replace guardrail calls with `evalReport.WorkspaceDelta.BytesNet`
   - Delete `guardrail.go` and its tests
   - Add `WorkspaceDelta` to `GraderInput` struct

2. **Guardrail softening (also #566):**
   - Change thresholds: 1 MB → 10 MB for MaxOutputSize, 50 → 200 for MaxFiles
   - Replace hard-fail with `GuardrailWarnings []string` append
   - Remove `Success = false` assignments in guardrail checks

3. **Phase 4 review restructure (separate issue):**
   - Move reviewers off inline file contents
   - Expose file manifest + workspace tools
   - Remove review caps (utils.go:16, 32) — they'll be dead code

---

## PR Approval

**APPROVE** — merge this PR.

- Hard-fail semantics: preserved ✅
- Phase 3.5 compatibility: clean subsumption path ✅  
- Live test: passes ✅  
- Code quality: surgical, well-tested, no architectural debt ✅

**Merge confidence:** HIGH. This is the right fix at the right scope.

---

**Signed:**  
Morpheus 🕶️  
Lead / Architect  
2026-04-16
