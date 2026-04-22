# 2026-04-22T21:28:00Z — Grader Unification Phase 1 Unblocked

**Context:** User (Ronnie) answered all 10 reformulated questions on grader unification schema. Morpheus filed Issues #624–#627 (Phase 1–4). Neo is ready to begin Phase 1 (unified schema + back-compat loader).

**Key Decisions Locked:**
- Flat `type` discriminator (no `kind`, no nested block)
- Prompt graders use same shape as other graders  
- File-level validation errors only fail evals using that file
- No deprecation shim for internal/criteria/ (delete immediately in Phase 3)
- output_check v1: `min_files`, `max_files`, file presence, per-file size checks

**Next:** Neo picks up #624. Tank and Trinity standby for Phase 2 execution upon Phase 1 completion.
