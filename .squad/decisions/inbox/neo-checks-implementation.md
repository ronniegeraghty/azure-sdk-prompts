# Neo — Prompt Grader `checks:` Implementation

**Author:** Neo 💊 — 2026-04-24
**For:** Coordinator / Scribe (merge into decisions.md)
**Status:** Landed on branch `ronniegeraghty/prompt-grader-checks`
**Scope ref:** `.squad/decisions/inbox/morpheus-prompt-grader-checks-scope.md` §7 "Neo owns"

## What landed

Implemented Morpheus's YAML prompt grader `checks:` field per scope §1, §4, §5. YAML prompt graders can now declare structured per-check items that produce one `GraderPoint` (and one nested Pass/Fail row in renderers) per check, instead of smuggling numbered lists inside `prompt:` and getting one combined verdict.

## Files touched

Engine / schema:
- `hyoka/internal/criteria/config.go` — added `Checks []string` to `UnifiedGraderEntry`; `validateEntry` enforces:
  - `type=prompt`: at least one of `prompt` / `checks` non-empty; each `checks` entry must trim non-empty (errors include the bad index); `details` still forbidden.
  - non-prompt types: `checks` must be empty (mirrors the existing `prompt`-must-be-empty rule).
- `hyoka/internal/criteria/buckets.go` — `FormatUnifiedPromptEntries` honors `Checks`. With checks: parent line `N. **Name**`, optional preamble (entry's `Prompt` field, judge-only), then nested `   1. …   2. …` checks. Without checks: legacy single-line `N. **Name** — Prompt` shape preserved byte-for-byte.
- `hyoka/internal/criteria/graders/prompt_review_grader.go` — added `expectedCriteriaCount` (regex-based leaf-numbered-item count, prefers indented when present) and `logCriteriaCountMismatch` helper. Both `gradePanel` and `gradeSingle` now emit a `slog.Debug` when `len(returnedCriteria) != expected`. No fail; no structural changes.

YAML migrations (hard-migrated per §2):
- `criteria/language/python.yaml` — `DefaultAzureCredential Authentication`: split smuggled `1. … 2. …` into preamble `"Check the following criteria:"` + 2 `checks:` items (auth-credential, async/await).
- `criteria/language/test.yaml` — `Markdown Structure`: matches the user's worked example exactly (preamble + 2 checks for hello.md heading and bullet count).
- `criteria/language/java.yaml`, `criteria/language/rust.yaml`, and `hyoka/internal/criteria/testdata/**` — left alone (no embedded numbering; single-`prompt` path covers them).

Tests:
- `hyoka/internal/criteria/config_test.go` — replaced `TestValidateEntry_PromptMissingPrompt` with `TestValidateEntry_PromptMissingPromptAndChecks` and added a 6-case table (`TestValidateEntry_PromptChecks`) covering: checks-only, preamble+checks, legacy prompt-only, empty string in checks (with index), whitespace-only check, checks on `output_check`.
- `hyoka/internal/criteria/buckets_test.go` — added `TestFormatUnifiedPromptEntries_Shapes` (5 cases): empty, legacy single-prompt, checks+preamble, checks-no-preamble, mixed legacy+checks. All assert exact rendered text.

## Verification

```
go build ./...                                    # clean
go test ./... -timeout 5m                         # all packages pass
go test -race ./hyoka/internal/criteria/...       # criteria + graders pass
```

Live smoke (`run --prompt-id key-vault-dp-python-crud --config baseline/claude-opus-4.6`, run id `20260424-052914`):

```
grader_results[].grader_name="DefaultAzureCredential Authentication"
  pass=false  score=0.333  points=[
    "Uses DefaultAzureCredential …": pass=true   "1/1 reviewers passed"
    "Uses async/await patterns …":   pass=false  "0/1 reviewers passed"
    "<parent name leaked>":           pass=false  "0/1 reviewers passed"   ← flake
  ]
```

Per-check `Points` propagate end-to-end through `convertGraderResults` → `report.GraderResult.Points`. Tank's interactive renderer should now have data to render the nested rows once the `Pass (X/Y)` badge change lands.

The third "leaked" point is the exact flake §5 anticipated — one judge returned the bucket parent string as a third "criterion". My new debug log fired correctly:

```
DEBUG Review judge returned criterion count differs from sent
  grader="DefaultAzureCredential Authentication" expected=2 returned=3
```

## Decisions kept (didn't re-litigate)

- One LLM call per grader, N checks rendered as N criteria (not per-check sessions).
- Hard-migrate the two YAMLs; no dual-format support.
- Preamble (`prompt:`) is judge-only — nested under the parent line, not surfaced separately to humans.
- Prompt-file path (`ParseEvaluationCriteria` bullet split) untouched; unification lives at the bucket-text layer only.

## Coordination with Tank

File-disjoint. Tank's badge change (`✅ Pass (X/Y)` / `❌ Fail (X/Y)` in `internal/progress/display_interactive.go`) and report-side verification can land independently. Branch `ronniegeraghty/prompt-grader-checks` is ready for Tank to push his commits onto.

## Follow-ups (not in scope)

- Judge sometimes returns the parent grader name as an extra criterion (the smoke flake above). If this becomes noisy, options: (a) tighten review prompt to forbid parent-line scoring, (b) post-filter returned criteria against sent leaf names. Tracking via the new debug log; defer until we have data from more runs.
