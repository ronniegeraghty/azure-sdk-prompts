# Decision: Grader Coverage UX Gap (2026-04-23)

## Status
✅ Investigation complete. No code bug found. Observability logs added. Documentation recommendations provided.

## Context
User reported: "Graders aren't running on all evals."

Neo investigated end-to-end from CLI → engine → grader execution and found the system is working as designed, but users lack visibility into which graders match and why some are skipped.

## Problem Analysis

### What's Actually Happening (Post-#625 Unified System)

The grader system has two types:
1. **Typed graders** (`output_check`, `file`, `program`, etc.) — Run directly via `criteria.RunGradersWithHooks()`
2. **Prompt-type graders** (`type: prompt`) — Fed into AI review panel as evaluation criteria

Both types are defined in `criteria/*.yaml` files with attribute matching via `when:` blocks (e.g., `when: language: python`).

### Why Users Think "Graders Aren't Running"

**Scenario A: Old Reports (Pre-#625)**
- Reports from before April 20, 2026 used the OLD dual-pipeline system
- Those reports may have zero graders if the old system wasn't fully configured
- **Solution:** User should look at recent reports (post-#625)

**Scenario B: Using `--skip-review`**
- Typed graders run normally
- Prompt-type graders are SKIPPED (they're embedded in review criteria)
- User sees only 1-2 typed graders and thinks "not all graders ran"
- **Solution:** Omit `--skip-review` to run all graders

**Scenario C: Zero Graders Matched**
- Grader file has `when: language: python`
- Prompt has `language: go`
- No graders match → correctly produces `grader_results: null`
- **Solution:** This is correct behavior (attribute matching)

**Scenario D: Generation Failed**
- Eval failed during generation phase (no files produced)
- Never reached grading phase → zero graders
- **Solution:** Check `error` field in report

### Auto-Discovery Already Works

The `--criteria-dir` flag has automatic fallback resolution:
```go
candidates := config.ResolveCandidates(proj, "criteria", "./criteria", "../criteria")
```

Users do NOT need to explicitly pass `--criteria-dir` if `./criteria` exists.

## Changes Made (Commit 0c20df51)

Added observability logs in `internal/eval/engine_eval.go`:

```go
// Log grader matching summary for observability.
glg.Info("Matched graders for eval",
    "total", len(matched),
    "typed", len(typedMatched),
    "prompt", len(promptMatched))
if len(matched) == 0 {
    glg.Warn("No graders matched this prompt's properties",
        "language", props["language"],
        "service", props["service"],
        "plane", props["plane"])
}
if len(promptMatched) > 0 && e.opts.SkipReview {
    glg.Info("Prompt-type graders matched but review is disabled",
        "count", len(promptMatched),
        "hint", "omit --skip-review to run these graders")
}
```

These logs fire at INFO/WARN level during grading phase. Users running with `--log-level info` or higher will now see:
- How many graders matched
- Breakdown of typed vs. prompt graders
- Warning when zero graders match (with prompt properties for debugging)
- Hint when prompt-type graders are skipped due to `--skip-review`

## Recommendations

### For the User (Immediate)
1. **Check recent reports** — Old reports (pre-#625) may not reflect current behavior
2. **Run without `--skip-review`** — Prompt-type graders require review to execute
3. **Check grader `when:` filters** — Ensure criteria files match your prompt's properties
4. **Use `--log-level info`** — New logs will show grader matching details

### For the Team (Future)

**A. Documentation (Oracle)**
- Explain the two grader types (typed vs. prompt) in user docs
- Show example criteria files with `when:` blocks
- Document `--skip-review` behavior: "skips prompt-type graders"
- Add troubleshooting section: "Why aren't my graders running?"

**B. CLI UX (Tank)**
- Add `--show-graders` to `hyoka list` command: show which graders would apply to each prompt
- Example: `hyoka list --prompt-id X --show-graders`
- Output: grader names, types, sources (file path), and match status

**C. Defer (Not Urgent)**
- Issue #622: Typed grader CLI surface (direct instantiation without criteria files)
- Not needed yet — users CAN define all graders in `criteria/*.yaml`

## Conclusion

**No bug exists.** The grader system is working as designed post-#625 unification. The user's observation was likely due to:
- Looking at old reports
- Not understanding typed vs. prompt grader execution
- Using `--skip-review` without realizing it skips half the graders

**Logging improvements** (commit 0c20df51) will surface grader matching behavior. Users can now diagnose:
- Which graders matched
- Why zero graders ran
- When prompt-type graders are skipped

**Next steps:** Documentation updates (Oracle) and optional CLI enhancement (Tank, `--show-graders` flag).

## Commit
- `0c20df51` — feat(eval): add grader matching observability logs
