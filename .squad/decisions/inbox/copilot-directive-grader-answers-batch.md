### 2026-04-22T21:22Z: User directives — Grader unification answers (Q2, Q4, Q5, Q6, Q7, Q8, Q10)
**By:** ronniegeraghty (via Copilot)
**Why:** User answers to Morpheus's open questions on the unified grader schema. These lock the Phase 1 schema design.

---

**Q2 — Discriminator field name and shape:** Use a flat `type` field at the entry level (NOT `kind`, NOT a nested `typed:` block). Prompt graders are NOT special — they use the same shape as every other grader. Example schema:

```yaml
graders:
  - name: grader 1
    type: prompt
    prompt: "the LLM-review prompt text..."
  - name: grader 2
    type: output_check
    details: { ... type-specific config ... }
```

**Rationale (verbatim):** "Why would we have the prompt graders be so different from the others? … This is what I'm talking about when I say we need to treat the prompt grader just like any other grader."

---

**Q4 — Misconfigured grader entries (e.g. `isolate: true` on a typed grader, or any structural error):** Reject the FILE at validation time — a criteria file or prompt file with a malformed grader fails validation. BUT the failure only errors out an eval if that file/criteria is actually used in the eval being run. Validation errors on unused criteria do not block unrelated evals.

---

**Q5 — Multiple graders of the same `type` in one file:** Allowed. Just like we already allow multiple prompt graders. Uniqueness is enforced by `name`, not by `type`. (User: "This should be obvious.")

---

**Q6 — Deprecation strategy for `internal/criteria/`:** Delete immediately at end of Phase 2. No deprecation shim. (Combined with Q8 — we don't care about breaking.)

---

**Q7 — `output_check` v1 knobs:** Build directly on the WorkspaceDelta artifact already collected post-eval. v1 must support:
- File-count thresholds: `min_files`, `max_files` (created OR updated, configurable)
- Specific-file checks: "is there a file named X" (presence), "was file X updated" (in delta)
- Per-file size thresholds: `min_bytes_per_file` style checks

The grader output is yes/no per check, with the check details surfaced in the result. Defer richer pattern-matching (globs, content regex) until v2 unless trivial.

---

**Q8 — Verification strategy for the Phase 2 cutover:** No special ceremony needed. "hyoka isn't a stable tool, so we don't care about breaking." Ship the cutover, fix what breaks, move on. No golden-file or parallel-run gate required.

---

**Q10 — Documentation of all grader types:** Yes, document every grader (`prompt`, `output_check`, `action_sequence`, `tool_constraint`, etc.) in user-facing docs. (User: "Why is this not obvious. We should document all the graders available.")
