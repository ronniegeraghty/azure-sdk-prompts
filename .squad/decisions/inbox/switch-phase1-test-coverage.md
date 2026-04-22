# Decision: Switch — Phase 1 Unified Grader Loader Test Coverage (#624)

**Author:** Switch 🤍
**Date:** 2026-04-22
**Branch:** `ronniegeraghty/dev` (direct commits, no PR)
**Status:** ✅ Complete, all green.

## Summary

Acceptance tests for the unified grader loader (Phase 1 of the Grader Unification rollout) landed on `ronniegeraghty/dev` alongside Neo's implementation. Ten test functions (plus 6 malformed-input subcases) and 12 YAML fixtures lock in the observable contract from issue #624 — flat `type` discriminator, no `Gate`/`kind`, uniqueness by name, legacy `criteria.yaml` back-compat, and deferred-error `Bundle` semantics from `LoadUnifiedDir`.

## Commits

| SHA | Purpose |
|---|---|
| `f66ab1bb` | Initial TDD scaffold, gated behind `//go:build phase1_pending` so CI stayed green while Neo's loader was in flight. |
| `f3915739` | Dropped the build tag, adapted to Neo's final API (`UnifiedGraderConfig` / `LoadUnifiedFile` / `LoadUnifiedDir` / `Bundle`). All green. |

## Test File

`hyoka/internal/graders/phase1_loader_test.go`

## Fixtures

`hyoka/internal/graders/testdata/phase1/`
- `mixed_prompt_and_typed.yaml`
- `two_same_type_unique_names.yaml`
- `duplicate_name.yaml`
- `malformed_missing_type.yaml` — neither `type` nor `prompt`
- `malformed_unknown_type.yaml`
- `malformed_prompt_missing_prompt.yaml`
- `malformed_typed_missing_details.yaml`
- `malformed_gate_field.yaml`
- `malformed_kind_field.yaml`
- `legacy_criteria.yaml` + `legacy_criteria_unified_equivalent.yaml` (back-compat pair)
- `empty_graders.yaml`
- `only_prompt_graders.yaml`
- `only_typed_graders.yaml`

## Coverage Map

| # | Requirement (issue #624) | Test | ✅ |
|---|---|---|---|
| 1 | Mixed prompt + output_check loads | `TestPhase1Loader_MixedPromptAndTyped` | ✅ |
| 2 | Same-type, different-name → success | `TestPhase1Loader_SameTypeDifferentNames` | ✅ |
| 3 | Duplicate name → error names both name + path | `TestPhase1Loader_DuplicateNameRejected` | ✅ |
| 4a | Missing `type` (no prompt fallback) | `…/missing_type_field` | ✅ |
| 4b | Unknown `type` value | `…/unknown_type_value` | ✅ |
| 4c | `type: prompt` without `prompt:` body | `…/prompt_type_missing_prompt_body` | ✅ |
| 4d | Typed grader without `details:` | `…/typed_missing_details` | ✅ |
| 4e | `gate:` rejected by KnownFields(true) | `…/gate_field_rejected_by_known_fields` | ✅ |
| 4f | `kind:` rejected by KnownFields(true) | `…/kind_field_rejected_by_known_fields` | ✅ |
| 5 | Legacy criteria.yaml back-compat | `TestPhase1Loader_LegacyBackCompat` | ✅ |
| 6 | Empty graders list — fail-loud (corrected) | `TestPhase1Loader_EmptyGraders` | ✅ |
| 7 | Prompt-only file | `TestPhase1Loader_OnlyPromptGraders` | ✅ |
| 8 | Typed-only file | `TestPhase1Loader_OnlyTypedGraders` | ✅ |
| + | `LoadUnifiedDir` deferred-error Bundle (Q4) | `TestPhase1Loader_LoadDirDeferredErrors` | ✅ |
| + | Nonexistent-file error sanity | `TestPhase1Loader_NonexistentFile` | ✅ |

## Findings That Shifted The Spec

1. **Empty graders must be rejected, not silently accepted.** The original task spec said "loads cleanly", but that would let a mis-indented `graders:` key silently produce a no-op criteria file. Neo correctly preserves the fail-loud behavior from `internal/criteria/criteria.go:139-141`. Test now locks that in.

2. **Back-compat translator promotes `{name, weight, prompt, when, isolate}` (no `type`) → `type: prompt`.** The first "missing type" fixture had a `prompt:` body, so it was a valid legacy entry, not malformed. Fixture rewritten to omit both `type` AND `prompt` — now correctly exercises the "type is required" validator.

## Non-Blocking Follow-Ups (for Phase 2)

- **No test for `Isolate` field back-compat.** Issue #624 spec says `isolate: true` is silently ignored on typed graders but passed through on prompt graders. Worth a table-driven test in Phase 2 once the engine actually consults `Isolate`.
- **No parallel-run test against real `criteria/*.yaml`.** Neo shipped `TestLoadUnifiedDir_RealCriteriaFixtures` in `unified_realfixtures_test.go` which exercises the 8-file real library, but it doesn't do a cross-check vs `internal/criteria/MatchingGraders()` for identical prompt-property resolution. That's the Phase 2 gate.
- **No test for `MatchingErrors(props)` Q4 semantics** (unused-criteria errors should NOT block unrelated evals). Neo's `unified_entry_test.go` covers this directly; our acceptance layer doesn't duplicate.

## Verification

```
go build ./...                              — clean
go test -race ./hyoka/internal/graders/...  — all green (2.7s)
```

## TDD Process Note

When implementation is landing in parallel and exact identifier names aren't locked, `//go:build <tag>_pending` on the test file is the right first move — keeps CI green, lands the spec, and the two-commit pattern (gated spec → drop-tag-when-impl-lands) makes the adaptation diff crisp and reviewable.
