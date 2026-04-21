# Decision: Morpheus — Phase 6 Round-1 Architectural Review

**Author:** Morpheus 🕶️
**Date:** 2026-04-21
**Phase:** 6 (epic #312)
**Round:** 1

## Verdicts

| PR | Author | Issue | Domain | Verdict |
|----|--------|-------|--------|---------|
| #601 | Trinity | #365 (WI-047) | Compare page redesign (site/) | **APPROVE** |
| #602 | Tank | #598 (R77) | Configurable prompt directory (Go CLI) | **APPROVE** |
| #603 | Neo | #580 | Review session splitting runtime (Go review/eval) | **APPROVE** |

All three PRs cleared the architectural bar on the first round. No request-changes; no agent lockouts triggered.

## #601 — Compare Page Redesign (Trinity)

- Layering matches existing site convention (pure-fn lib + thin component + page orchestration), mirrors eval-detail / prompt-detail patterns from Phase 5.
- Catalog driven from real eval data; filter semantics (within-dim OR / across-dim AND / empty = match-all) explicit and tested.
- Versioned `localStorage` key (`hyoka:comparison:v1`) with shape validation + safe SSR/test fallback.
- 99/99 site tests pass; build clean.
- **Follow-up:** Remove dead `fetchCompareConfigs` from `site/src/app/data/api.ts` and sunset its serve-side endpoint. Author already flagged.

## #602 — Configurable Prompt Directory (Tank)

- Backwards compatibility **enforced in code**, not just commented: empty `PromptDirectory` field delegates to legacy `ResolveCandidates` chain, asserted by `TestResolvePromptDirCandidates_NoOverrideDelegates` and `TestLoad_NoPromptDirectoryDefaultsEmpty`.
- `cmd` vs `internal` split respected; new helpers in `internal/config/discovery.go`.
- Resolution priority consistent across `run`, `validate`, `list`, `serve`: flag > config > defaults.
- `PeekPromptDirectory` is a deliberate tolerant-load helper for commands that need the prompts dir before strict config validation runs — well-justified asymmetry.
- Conflict detection on `LoadDir` for divergent values across files in the same `--config-dir`.
- Relative paths resolve against the config file's directory (correct semantic, matches skills paths).
- 11 new tests, all green.
- **Follow-ups:** (a) defensive nil-check on `cfgFile` in `cmd/run.go:185`; (b) separate bug for the malformed `init --with-examples` template Tank flagged in his history.

## #603 — Review Session Splitting (Neo)

- **Critical: PR #355/#587 regression actively prevented.** The flag drives runtime behavior (`ReviewPanelBuckets` → N sessions per panel model), and when `--review-mode isolated` is requested but no graders/groups are marked `isolate`, `slog.Warn` fires and falls back to combined. Flag is observably no-op rather than silently dead.
- Default behavior preserved: combined mode → 1 bucket → falls through to legacy `ReviewPanel` path. Byte-identical output for users who don't opt in.
- `Bucket`/`ReviewBucket` type duplication across `criteria` / `graders` / `review` packages is awkward but correctly justified by import-cycle constraints; conversion sites are minimal and clearly named.
- `mergeBucketResults` prefixes per-bucket criterion names with `[bucket-name]` to keep deterministic any-fail voting unambiguous across panel models.
- `MultiBucketReviewer` interface + graceful degradation for non-bucket-aware reviewers.
- All criteria/review/eval/graders/cmd tests green under `go test -race`.
- **Follow-ups:** (a) document `[bucket-name]` criterion-name prefix in `docs/configuration.md`; (b) audit downstream criterion-name joins (trends/comparison) for the prefix, owner Trinity.

## Architectural Drift Notes

None. All three PRs respect the project's existing layering (Go `cmd` vs `internal`; site pure-fn lib vs component vs page), error-handling conventions (`%w` wrapping, no log-and-return), and naming patterns. No third-party logging introduced; Tank's PR added one new minor dep usage (`gopkg.in/yaml.v3` for the peek helper) which is already in the dependency manifest.

## Lockout Status

No agents locked out. All three authors free for round-2 work or follow-ups.

— Morpheus 🕶️
