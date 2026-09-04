# Skill: Release Smoke Test

**Purpose:** Verify hyoka release readiness through live evaluation dogfooding

**When to use:**
- Before promoting a branch to main (e.g., dev → main)
- After merging large feature PRs
- Before cutting a versioned release (e.g., v0.3.1)
- When verifying Phase N completion

**When NOT to use:**
- For individual feature testing (use targeted tests instead)
- During active development (too heavyweight)
- For CI validation (use automated test suite)

## Smoke Test Matrix

A release smoke test exercises the full evaluation pipeline end-to-end:

1. **Build verification** — `go build ./...` must pass
2. **Live eval run** — Small matrix (1 prompt × 2 configs) to verify generation, review, and reporting
3. **Comparison generation** — Verify auto-generated comparison files
4. **Serve endpoints** — Verify API and site serving
5. **Cleanup** — Verify session cleanup works

## Standard Test Command

```bash
# Dogfood eval: 1 prompt × 2 configs (Python prompts fastest)
go run . run --prompt-id key-vault-dp-python-crud \
  --config "baseline/claude-opus-4.6,azure-mcp/claude-opus-4.6" \
  --log-level debug --log-file smoke-test.log
```

**Why this prompt?**
- **key-vault-dp-python-crud** — Basic difficulty, reliable completion, exercises Azure SDK
- **Python** — Fastest language for iteration
- **2 configs** — Minimal matrix to test comparison auto-gen (baseline vs tool-enabled)

**Expected duration:** 5–10 minutes

## Verification Checklist

### 1. Build
- [ ] `go build ./...` passes
- [ ] No compiler errors
- [ ] All packages compile

### 2. Eval Run
- [ ] Command completes without panic
- [ ] At least 1 evaluation passes
- [ ] Run directory created in `reports/`
- [ ] `summary.json` and `summary.md` exist
- [ ] No unexpected errors in log (gemini warning OK)

### 3. Comparison Auto-Gen
- [ ] `comparisons.json` exists in run directory
- [ ] File contains valid JSON
- [ ] Has `kind`, `label_a`, `label_b`, `summary` fields
- [ ] `per_prompt` array populated

### 4. Serve Endpoints
- [ ] Start: `go run . serve` (or `nohup ... &`)
- [ ] Root: `curl http://localhost:8080/` returns HTML
- [ ] API runs: `curl http://localhost:8080/api/runs` returns JSON array
- [ ] Run detail: `curl http://localhost:8080/api/runs/{runId}` returns summary
- [ ] Comparisons: `curl http://localhost:8080/api/runs/{runId}/comparisons` returns JSON
- [ ] No 500 errors

### 5. Cleanup
- [ ] `go run . clean` runs successfully
- [ ] Reports freed sessions (or "No orphaned processes")
- [ ] No lingering Copilot sessions

## Evidence Collection

Collect these artifacts for the verification report:

1. **Commit info:** `git log -1 --oneline`
2. **Build output:** `go build ./... 2>&1`
3. **Eval log:** `smoke-test.log` (debug level)
4. **Eval stdout:** Capture console output
5. **Run directory:** Path to `reports/{runId}/`
6. **Comparison JSON:** Contents of `comparisons.json`
7. **Serve log:** Startup output
8. **API responses:** Sample curl outputs

## Common Issues

### Eval fails
- Check log for errors
- Verify configs exist and are valid
- Verify prompt exists and is valid
- Check Copilot SDK connectivity

### Comparison not generated
- Verify 2+ configs in run
- Check `internal/eval/engine.go` calls `comparison.WriteForRun()`
- Check for errors in log related to comparison

### Serve endpoints fail
- Verify site built (or embedded site present)
- Check `reports/` directory exists and is readable
- Check for port conflicts (8080)

### Cleanup reports "No sessions"
- Expected if eval already cleaned up
- Only concerning if eval crashed and left orphans

## Success Criteria

All 5 verification steps pass:
1. ✅ Build clean
2. ✅ Eval run passes (≥1 eval success)
3. ✅ Comparison auto-generated
4. ✅ Serve endpoints working
5. ✅ Cleanup successful

**Pass rate target:** 100% of evaluation runs should pass (for basic prompts like key-vault crud).

**Duration target:** ≤10 minutes for 1 prompt × 2 configs.

## Output

Produce a verification report with:
- Executive summary (PASS/FAIL/BLOCKERS)
- Per-check status (Build, Eval, Comparison, Serve, Cleanup)
- Evidence (logs, artifacts, API responses)
- Blockers identified (if any)
- Recommendations (release now / fix first / more testing)

## Example Report Structure

```markdown
# Release Smoke Test Report
**Date:** YYYY-MM-DD
**Branch:** {branch}
**Commit:** {hash} {message}

## Executive Summary
✅ PASS / ❌ FAIL / ⚠️ BLOCKERS

## Verification Results
### 1. Build: ✅/❌
### 2. Eval Run: ✅/❌
### 3. Comparison Auto-Gen: ✅/❌
### 4. Serve Endpoints: ✅/❌
### 5. Cleanup: ✅/❌

## Blockers
- None / List blockers

## Recommendations
- Cut release / Fix blockers first / More testing needed
```

## Notes

- **Not a substitute for CI:** Smoke tests verify integration, not unit correctness. All tests must pass in CI.
- **Not exhaustive:** This tests happy path only. Edge cases covered by unit/integration tests.
- **Dogfooding principle:** Use the tool as users would. If smoke test fails, release would fail in production.

## Related

- `.squad/agents/morpheus/history.md` — Phase 4 dogfood verification (2026-04-17)
- `docs/contributing.md` — Development workflow
- `docs/cli-reference.md` — Command documentation
