# WorkspaceDelta Test Plan (#566)

**Author:** Switch 🤍  
**Date:** 2026-04-17  
**Issue:** #566  
**Status:** Test plan ready for Neo's implementation

---

## Overview

This document enumerates all test scenarios for the WorkspaceDelta feature (#566). Once Neo's branch `squad/566-workspace-delta` exists with the new types, these scenarios will be codified as table-driven tests in `hyoka/internal/workspace/delta_test.go`.

**Reading audience:** Neo (implement the type to satisfy these), Switch (codify these into _test.go), Morpheus (audit coverage).

---

## 1. Delta Computation Correctness

### Scenario 1.1: Fresh workspace (no starter)

**Setup:**
- Empty starter snapshot (no files)
- Agent creates 3 files: `main.py` (100 bytes), `lib/utils.py` (200 bytes), `README.md` (50 bytes)

**Expected Delta:**
```
NewFiles:         [main.py (100B), lib/utils.py (200B), README.md (50B)]
ModifiedFiles:    []
DeletedFiles:     []
NewFileCount:     3
ModifiedFileCount: 0
DeletedFileCount: 0
BytesAdded:       350
BytesRemoved:     0
BytesNet:         +350
```

**Test Implementation:**
- Create workspace with no starter
- Write 3 files
- Compute delta
- Assert all files in `NewFiles` list
- Assert no files in `Modified` or `Deleted` lists
- Assert byte counts match file sizes

---

### Scenario 1.2: Starter file unchanged by agent

**Setup:**
- Starter: `setup.py` (100 bytes)
- Agent: does not modify `setup.py`

**Expected Delta:**
```
NewFiles:         []
ModifiedFiles:    []
DeletedFiles:     []
NewFileCount:     0
ModifiedFileCount: 0
DeletedFileCount: 0
BytesAdded:       0
BytesRemoved:     0
BytesNet:         0
```

**Rationale:** Unchanged starter files do NOT appear in any delta list — they are not agent work product.

**Test Implementation:**
- Snapshot starter with 1 file
- Leave file unmodified
- Compute delta
- Assert all lists empty
- Assert all counts zero

---

### Scenario 1.3: Starter file modified by agent (growth)

**Setup:**
- Starter: `app.py` (50 bytes, content: "# starter")
- Agent: appends content → `app.py` (200 bytes, content: "# starter\nprint('hello')\n...")

**Expected Delta:**
```
NewFiles:         []
ModifiedFiles:    [app.py (sizeBefore=50, sizeAfter=200, delta=+150)]
DeletedFiles:     []
NewFileCount:     0
ModifiedFileCount: 1
DeletedFileCount: 0
BytesAdded:       150
BytesRemoved:     0
BytesNet:         +150
```

**Test Implementation:**
- Create starter file with known content
- Snapshot
- Overwrite file with larger content
- Compute delta
- Assert file appears in `ModifiedFiles` with correct before/after sizes
- Assert `BytesNet` = delta only

---

### Scenario 1.4: Starter file shrunk by agent

**Setup:**
- Starter: `config.json` (500 bytes)
- Agent: overwrites → `config.json` (100 bytes)

**Expected Delta:**
```
NewFiles:         []
ModifiedFiles:    [config.json (sizeBefore=500, sizeAfter=100, delta=-400)]
DeletedFiles:     []
NewFileCount:     0
ModifiedFileCount: 1
DeletedFileCount: 0
BytesAdded:       0
BytesRemoved:     400
BytesNet:         -400
```

**Rationale:** Shrinking counts as agent work (they edited it). `BytesNet` is negative.

**Test Implementation:**
- Create large starter file
- Snapshot
- Overwrite with smaller content
- Compute delta
- Assert negative delta in `ModifiedFiles`
- Assert `BytesNet` is negative

---

### Scenario 1.5: Starter file deleted by agent

**Setup:**
- Starter: `legacy.py` (300 bytes)
- Agent: deletes `legacy.py`

**Expected Delta:**
```
NewFiles:         []
ModifiedFiles:    []
DeletedFiles:     [legacy.py (originalSize=300)]
DeletedFileCount: 1
BytesAdded:       0
BytesRemoved:     300
BytesNet:         -300
```

**Test Implementation:**
- Create starter file
- Snapshot
- Delete file
- Compute delta
- Assert file in `DeletedFiles` with correct original size
- Assert `BytesNet` is negative

---

### Scenario 1.6: Agent creates then deletes file in same session

**Setup:**
- No starter
- Agent creates `temp.py` (100 bytes), then deletes it

**Expected Delta:**
```
NewFiles:         []
ModifiedFiles:    []
DeletedFiles:     []
NewFileCount:     0
ModifiedFileCount: 0
DeletedFileCount: 0
BytesAdded:       0
BytesRemoved:     0
BytesNet:         0
```

**Rationale:** Final state == initial state → no net delta. File never existed in starter and doesn't exist in final workspace.

**Test Implementation:**
- Empty starter snapshot
- Final workspace state: empty
- Compute delta
- Assert all lists empty (agent created and deleted in same session = no net change)

---

### Scenario 1.7: Mixed operations (new + modified + deleted)

**Setup:**
- Starter: `main.py` (100 bytes), `lib.py` (50 bytes)
- Agent:
  - Modifies `main.py` → 250 bytes (+150)
  - Deletes `lib.py` (-50)
  - Creates `test.py` (80 bytes)
  - Creates `README.md` (40 bytes)

**Expected Delta:**
```
NewFiles:         [test.py (80B), README.md (40B)]
ModifiedFiles:    [main.py (sizeBefore=100, sizeAfter=250, delta=+150)]
DeletedFiles:     [lib.py (originalSize=50)]
NewFileCount:     2
ModifiedFileCount: 1
DeletedFileCount: 1
BytesAdded:       80 + 40 + 150 = 270
BytesRemoved:     50
BytesNet:         +220
```

**Test Implementation:**
- Create starter with 2 files
- Snapshot
- Perform all operations above
- Compute delta
- Assert correct categorization in each list
- Assert byte math is correct

---

### Scenario 1.8: Zero-byte files

**Setup:**
- Starter: `__init__.py` (0 bytes)
- Agent:
  - Leaves `__init__.py` unchanged (0 bytes)
  - Creates `empty.txt` (0 bytes)

**Expected Delta:**
```
NewFiles:         [empty.txt (0B)]
ModifiedFiles:    []
DeletedFiles:     []
NewFileCount:     1
ModifiedFileCount: 0
DeletedFileCount: 0
BytesAdded:       0
BytesRemoved:     0
BytesNet:         0
```

**Rationale:** Zero-byte files are tracked but contribute zero bytes. Unchanged `__init__.py` does not appear in delta.

**Test Implementation:**
- Create zero-byte starter file
- Snapshot
- Create new zero-byte file
- Compute delta
- Assert new file appears in `NewFiles` with size 0
- Assert `BytesNet` = 0

---

### Scenario 1.9: Empty workspace after agent run

**Setup:**
- Starter: `main.py` (100 bytes), `lib.py` (50 bytes)
- Agent: deletes everything

**Expected Delta:**
```
NewFiles:         []
ModifiedFiles:    []
DeletedFiles:     [main.py (100B), lib.py (50B)]
DeletedFileCount: 2
BytesAdded:       0
BytesRemoved:     150
BytesNet:         -150
```

**Test Implementation:**
- Create starter with 2 files
- Snapshot
- Delete all files
- Compute delta
- Assert both files in `DeletedFiles`
- Assert negative `BytesNet`

---

## 2. JSON Output Integration

### Scenario 2.1: EvalReport JSON contains workspace_delta

**Setup:**
- Run eval with mixed delta (new + modified + deleted)
- Serialize `EvalReport` to JSON

**Expected:**
- JSON contains `"workspace_delta": { "bytes_added": N, "bytes_removed": M, ... }`
- All fields present with correct values
- Field names stable (snake_case per JSON convention)

**Test Implementation:**
- Populate `EvalReport.WorkspaceDelta` with test data
- Call `json.Marshal(report)`
- Parse JSON and assert field names and values match
- No omitempty tags on critical fields (all must serialize even if zero)

---

### Scenario 2.2: Missing/nil WorkspaceDelta does not break decoding

**Setup:**
- JSON from older hyoka version (no `workspace_delta` field)
- Decode into current `EvalReport` struct

**Expected:**
- No decode error
- `report.WorkspaceDelta` is nil (not panic)
- Other fields decode correctly

**Test Implementation:**
- Create JSON string without `workspace_delta` field
- Unmarshal into `EvalReport`
- Assert no error
- Assert `report.WorkspaceDelta == nil`
- Assert other fields (e.g., `PromptID`, `Success`) decode correctly

---

### Scenario 2.3: WorkspaceDelta with zero values serializes correctly

**Setup:**
- `WorkspaceDelta` with all counts and sizes = 0 (no agent changes)

**Expected:**
- JSON serializes all fields as `0` (not omitted)

**Test Implementation:**
- Create `WorkspaceDelta{}` (all zero values)
- Marshal to JSON
- Assert fields are present with value `0`, not omitted

---

## 3. Grader Integration

### Scenario 3.1: GraderInput exposes WorkspaceDelta

**Setup:**
- Create `GraderInput` with populated `WorkspaceDelta`

**Expected:**
- Grader can access `input.WorkspaceDelta`
- No nil panics if delta is absent (graceful nil handling)

**Test Implementation:**
- Construct `GraderInput` with non-nil `WorkspaceDelta`
- Pass to a test grader
- Grader reads `input.WorkspaceDelta.NewFiles`
- Assert no panic

---

### Scenario 3.2: Grader handles nil WorkspaceDelta gracefully

**Setup:**
- `GraderInput` with `WorkspaceDelta = nil` (pre-#566 report)

**Expected:**
- Grader checks `if input.WorkspaceDelta != nil` before accessing
- No panic, returns meaningful result or skips delta-based logic

**Test Implementation:**
- Create `GraderInput` with `nil` delta
- Pass to grader
- Grader should not panic
- Document pattern: graders must nil-check `WorkspaceDelta`

---

## 4. Guardrail Interaction

### Scenario 4.1: Large delta triggers warning (not hard-fail)

**Setup:**
- Config: `MaxOutputSize = 10 MB` (warning threshold)
- Agent delta: `BytesNet = 12 MB`

**Expected:**
- Eval succeeds (`Success = true`)
- `GuardrailWarnings` contains: "Agent output size 12.00 MB exceeds warning threshold 10.00 MB"
- Eval does NOT hard-fail

**Test Implementation:**
- Create workspace with large delta (exceeds threshold)
- Run guardrail check
- Assert `Success = true`
- Assert warning message in `EvalReport.GuardrailWarnings`

---

### Scenario 4.2: Delta-based size vs total size

**Setup:**
- Starter: 5 MB of files
- Agent: adds 2 MB of new files
- Threshold: `MaxOutputSize = 8 MB`

**Expected:**
- **Old behavior (total size):** 7 MB total → under threshold
- **New behavior (delta size):** 2 MB delta → under threshold
- Both pass, but new behavior is more accurate (starter size not charged to agent)

**Test Implementation:**
- Compute agent output size using `computeAgentOutputSize` helper
- Assert it counts delta only (2 MB), not total workspace size (7 MB)
- Assert guardrail check uses delta size

---

### Scenario 4.3: Starter deleted reduces delta size

**Setup:**
- Starter: 10 MB of files
- Agent: deletes all starter files, creates 1 MB of new files
- Threshold: `MaxOutputSize = 5 MB`

**Expected:**
- `BytesNet = +1 MB - 10 MB = -9 MB` (net negative)
- Guardrail check: delta size < threshold → passes

**Test Implementation:**
- Create large starter
- Delete all, add small files
- Compute `BytesNet`
- Assert negative value
- Assert guardrail does not trigger (negative delta = agent removed more than added)

---

### Scenario 4.4: MaxNewFiles guardrail (new)

**Setup:**
- Config: `MaxNewFiles = 100` (warning threshold)
- Agent: creates 120 new files

**Expected:**
- Eval succeeds (`Success = true`)
- `GuardrailWarnings` contains: "Agent created 120 new files, exceeds warning threshold 100"

**Test Implementation:**
- Create workspace with 120 new files (no starter)
- Run guardrail check
- Assert warning in `GuardrailWarnings`
- Assert `Success = true`

---

## 5. Edge Cases

### Scenario 5.1: Binary files in delta

**Setup:**
- Agent creates `image.png` (5 KB binary data)

**Expected:**
- `NewFiles` contains `image.png` with size tracked
- **No content diffing** — binary files are tracked by path and size only
- No panic or decode error

**Test Implementation:**
- Write binary file to workspace
- Compute delta
- Assert file appears in `NewFiles` with correct size
- Assert no attempt to diff content (size-only tracking)

---

### Scenario 5.2: Unexpected build artifacts (e.g., .pyc files)

**Setup:**
- Starter: `main.py`
- Agent: runs code, creates `__pycache__/main.cpython-311.pyc`

**Expected:**
- `__pycache__/` is excluded from delta computation (matches `DefaultIgnoreDirs`)
- `.pyc` files do NOT appear in `NewFiles`

**Test Implementation:**
- Create workspace with starter
- Write `.pyc` file in `__pycache__/`
- Compute delta with exclusion filter
- Assert `.pyc` not in `NewFiles`
- Document: delta computation respects `DefaultIgnoreDirs`

---

### Scenario 5.3: Symlinks

**Setup:**
- Agent creates symlink `link.py` → `main.py`

**Expected:**
- Symlinks are skipped (per existing `copyDir` logic)
- **If** symlinks are tracked: follow symlink target, report target file size
- **If** symlinks are skipped: not in delta

**Test Implementation:**
- Create symlink in workspace
- Compute delta
- Assert symlink handling matches existing `copyDir` behavior (skipped, logged as warning)
- Document: symlinks are excluded from delta (safety — prevent escape)

---

### Scenario 5.4: Very large file (GB-scale)

**Setup:**
- Agent creates `data.bin` (2 GB)

**Expected:**
- Delta computation does not OOM (size-only tracking, no content load)
- `NewFiles` contains entry with `size = 2147483648` (2 GB)
- Warning triggered if exceeds threshold

**Test Implementation:**
- Create large file using `fallocate` or sparse file
- Compute delta
- Assert size is tracked correctly (int64, no overflow)
- Assert no memory spike (size-only, not content-based)

---

### Scenario 5.5: Unicode filenames

**Setup:**
- Agent creates `文档.txt`, `résumé.md`

**Expected:**
- Files tracked correctly in delta
- No encoding issues in JSON serialization

**Test Implementation:**
- Create files with non-ASCII names
- Compute delta
- Assert files appear in `NewFiles` with correct paths
- Marshal to JSON, assert no errors

---

### Scenario 5.6: File permissions change (no size change)

**Setup:**
- Starter: `script.sh` (755, 100 bytes)
- Agent: changes to 644, same size

**Expected:**
- **Not tracked as modification** — permission changes are not size deltas
- File does not appear in `ModifiedFiles`

**Rationale:** WorkspaceDelta is size-based, not metadata-based. Permission changes are out of scope.

**Test Implementation:**
- Create file with one permission
- Snapshot
- Change permission, keep content identical
- Compute delta
- Assert file not in `ModifiedFiles`

---

## 6. Integration with Existing Tests

### Scenario 6.1: Workspace lifecycle tests pass with delta

**Setup:**
- Existing `workspace_test.go` tests (`TestCopyStarterFiles`, `TestWorkspaceCleanup`, etc.)

**Expected:**
- All existing tests pass unchanged
- Delta computation is opt-in (does not break existing workspace flows)

**Test Implementation:**
- Run `go test ./hyoka/internal/eval/ -v`
- Assert zero regressions

---

### Scenario 6.2: Guardrail tests updated

**Setup:**
- Existing `guardrail_test.go` table-driven tests

**Expected:**
- Tests updated to use `WorkspaceDelta` instead of raw file lists
- All 15 existing cases still pass

**Test Implementation:**
- Refactor guardrail tests to construct `WorkspaceDelta`
- Run tests, assert no regressions

---

## Test Organization

### File Structure

```
hyoka/internal/workspace/
  delta.go               # WorkspaceDelta type + computation logic
  delta_test.go          # All scenarios above (table-driven)
  snapshot.go            # snapshotStarterSizes helper
  snapshot_test.go       # Snapshot logic tests

hyoka/internal/eval/
  guardrail.go           # Updated to use WorkspaceDelta
  guardrail_test.go      # Updated test cases
```

### Test Patterns

- **Table-driven tests** for all delta computation scenarios (1.1–1.9, 5.1–5.6)
- **JSON marshal/unmarshal tests** for serialization (2.1–2.3)
- **Integration tests** with `GraderInput` (3.1–3.2)
- **Guardrail tests** with warning assertions (4.1–4.4)

---

## Acceptance Criteria

A test case passes when:

1. **Setup** state is reproducible (temp dirs, fixture files)
2. **Expected delta** matches computed delta (all fields)
3. **No panics** on nil/missing delta
4. **No regressions** in existing tests
5. **JSON round-trips** correctly (marshal → unmarshal → equal)
6. **Guardrail warnings** appear in `GuardrailWarnings`, not hard-fail
7. **Edge cases** (binary, symlinks, large files) handled gracefully

---

## Handoff to Neo

**Read this plan before implementing.** Each scenario describes:
- **Why** the case matters (rationale)
- **What** the expected output is (assertions)
- **How** to structure the test (pattern)

Once your branch `squad/566-workspace-delta` has the `WorkspaceDelta` type defined, I'll write the actual `delta_test.go` file implementing all these scenarios. We'll coordinate via decision inbox — ping me when types are ready.

---

**Status:** ✅ Test plan complete. Awaiting Neo's implementation branch.

**Next steps:**
1. Neo: Define `WorkspaceDelta` struct + computation logic
2. Switch: Code `delta_test.go` against Neo's branch
3. Switch: Run tests, report gaps
4. Neo: Fix gaps, iterate
5. Merge: All tests green → PR ready
