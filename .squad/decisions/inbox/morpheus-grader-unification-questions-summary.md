# Grader Unification — Open Questions Summary

**Source:** `.squad/decisions/inbox/morpheus-grader-unification-proposal.md` § 6
**Author:** Morpheus 🕶️
**Date:** 2026-04-23 (recovery)

These 10 questions need Ronnie's decision before Phase 1 begins. Each has options and a recommendation.

1. **CLI flag naming** — Keep `--criteria-dir` (rec: A, keep as-is) or rename to `--graders-dir` with alias?
2. **`Kind` placement** — Bare optional field on entry (rec: A) or nested `typed:` block?
3. **`Gate` on typed graders** — Hard fail (rec: A, consistent with existing `AggregateResults`) or soft warn?
4. **`Isolate` on typed graders** — Silent ignore, load-time warning (rec: B), or load-time error?
5. **Multiple same-Kind in one file** — Allowed with unique names (rec: A) or rejected?
6. **`internal/criteria/` deletion timing** — Immediate in Phase 3 (rec: A) or deprecation shim for one release?
7. **`output_check` v1 knobs** — Ship `min_files` + `min_bytes_per_file` (rec). Defer `required_extensions`, `glob_filter`.
8. **Test strategy** — Both golden-file + parallel-run (rec: C) during transition.
9. **`Gate` on prompt graders** — Reject at load time (rec: B). LLM scores too noisy for hard gates.
10. **`action_sequence`/`tool_constraint` in user docs** — Document in Advanced section (rec: A).

Full rationale and trade-off tables in the proposal §6.
