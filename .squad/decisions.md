# Squad Decisions

## Active Decisions
### Decision: Bucket-Per-Entry Structure for AI Grader Display (2026-04-24T04:55:03Z)

**Agent:** Tank 📡  
**Branch:** ronniegeraghty/dev  
**Commit:** 9e2d8100  
**Status:** Implemented

## Context

AI grader display was showing a single "combined" bucket grouping all criteria-file entries, instead of one top-level grader per entry. This was the fourth fix in a chain addressing the bucket display bug:

- `4adc9288` — initial per-bucket refactor
- `84b1606d` — remove KindPromptReview from validTypedKinds
- `609ff869` — fix bucket contents (clear EvalCriteria per-bucket)
- `9e2d8100` — **fix bucket structure (one bucket per entry)**

## Problem

User expected:
```
- Output Files Exist (output_check)
- Criteria from prompt file (prompt)
- DefaultAzureCredential Authentication (prompt)
```

But saw:
```
- Output Files Exist (output_check)
- Criteria from prompt file (prompt)
- combined (prompt)              ← grouping all entries
  - DefaultAzureCredential Authentication
```

## Decision

**Change `BuildUnifiedReviewBuckets` combined mode behavior to create one bucket per criteria-file entry.**

- Each bucket uses the entry's `name` field as the bucket name
- Prompt-frontmatter criteria continue to live in their own "Criteria from prompt file" bucket
- Isolated mode behavior unchanged (entries marked `isolate: true` get their own bucket, rest go to "combined")

## Implementation

Modified `hyoka/internal/criteria/buckets.go` lines 184-189:

**Before:**
```go
if mode != ReviewModeIsolated || !HasUnifiedIsolation(matched) {
    if len(matched) > 0 {
        buckets = append(buckets, combinedCriteriaFileBucket(matched))
    }
    return buckets
}
```

**After:**
```go
if mode != ReviewModeIsolated || !HasUnifiedIsolation(matched) {
    for _, m := range matched {
        buckets = append(buckets, graders.ReviewBucket{
            Name:     bucketName(m.Entry.Name, len(buckets)),
            Criteria: MergeUnifiedCriteria([]UnifiedGraderEntry{m.Entry}, ""),
        })
    }
    return buckets
}
```

## Impact

- **Display:** Each criteria-file grader entry now renders as its own top-level grader with individual sub-criteria
- **Test updates:** 6 test files updated to expect N+1 buckets (prompt + N entries) instead of 2 (prompt + combined)
- **Backward compat:** The "combined" bucket name still exists in isolated mode for leftover entries, preserving its special handling in `mergeBucketResults`

## Rationale

1. **User expectation:** One grader per criteria entry matches the mental model ("I defined 3 criteria, I should see 3 graders")
2. **Clarity:** Each criteria entry's name becomes a top-level display element
3. **Consistency:** Matches the existing behavior for prompt-frontmatter criteria (separate bucket)

## Alternatives Considered

1. **Keep "combined" and rename it:** Would still group entries, defeating the purpose
2. **Add a flag to toggle behavior:** Adds complexity; the per-entry behavior is always desired
3. **Change display layer to unpack "combined":** Would require display logic to understand bucket internals

## Related

- Issue: User verbatim complaint about "combined" bucket
- Commits: `4adc9288`, `84b1606d`, `609ff869`, `9e2d8100`
- Files: `buckets.go`, `engine.go`, 6 test files

---

### Decision: Test Fixture Language Pattern (2026-04-24)

**Agent:** Switch 🤍  
**Status:** Implemented

## Context

Tank iterating on grader display bugs. Real Python Azure evals take 2-3 min each because Copilot generates substantial SDK code. Needed trivially fast fixture for grader rendering iteration.

## Decision

Built a "test fixture" prompt (`storage-dp-python-hello-markdown-test`) that:
- Uses existing `language: python` (not a new "test" language)
- Inherits `criteria/language/python.yaml` graders for realistic multi-grader test
- Writes a single trivial markdown file (completes in 1 turn)
- Runs in **29 seconds** end-to-end (83% faster than Azure prompts)

## Rationale

1. **Reuse existing language:** Avoids modifying prompt validation logic to add "test" to allowed languages
2. **Inherit real graders:** Exercises the full grader pipeline with actual language criteria
3. **Expected failures are OK:** DefaultAzureCredential grader fails (not applicable to markdown) — this is correct behavior
4. **Fast iteration:** 29-second evals vs 2-3 min = rapid feedback loop for display bugs

## Implementation

- `prompts/test/hello-markdown.prompt.md` — trivial markdown-writing prompt
- `criteria/language/test.yaml` — 4 graders (prompt, output_check, file, behavior)
- `configs/test-baseline.yaml` — haiku model, no tools

## Invocation

```bash
hyoka run --prompt-id storage-dp-python-hello-markdown-test \
  --config test/haiku
```

## Status

✅ Implemented and verified end-to-end. Validation passes. All configured graders execute.

## Related

- Tank's concurrent work: `criteria/language/python.yaml`, `internal/eval/engine_eval.go` (staged, not touched)
- Future: Consider extracting "universally applicable" graders (output_check, file, behavior) into a `criteria/language/_common.yaml` to avoid per-language duplication

---

### Decision: Test Fixture Language Allowlist Extension (2026-04-24)

**Agent:** Switch 🤍  
**Status:** Implemented

## Context

Round 2 cleanup of test fixture (fixing language field divergence).

## Decision

Extended validation allowlists in `internal/validate/validate.go` to include `"test"` as a valid value for:
- `ValidServices` (added `"test"`)
- `ValidLanguages` (added `"test"`)
- `ValidCategories` (added `"test"`)

## Rationale

1. **User spec:** Original test fixture request explicitly asked for `language: test` (quote: *"Maybe a prompt files in a test dir under the prompt dir that has a language of `test`"*)

2. **Dual validation paths:** hyoka has two separate validation implementations:
   - `schema.go:ValidatePromptStruct()` — Has `isTestValue()` escape hatch that accepts "test" prefix/suffix
   - `validate.go:validatePrompt()` — Used by `hyoka validate` command, NO escape hatch
   
   The first pass assumed the escape hatch would work, but `hyoka validate` rejected `language: test`.

3. **User preference:** When the divergence was discovered, the user said: *"If there's a hardcoded allowlist, ADD `test` to it (don't remove the allowlist — extend it). If `validate` complains, fix the validation rule rather than the prompt."*

4. **Test fixture semantic:** The `test` value is explicitly for test/fixture purposes, not production Azure prompts. Adding it to the allowlist makes it a first-class fixture language alongside `python`, `dotnet`, etc.

## Impact

- Test prompts with `service: test`, `language: test`, `category: test` now pass validation
- No impact on production prompts (they don't use these values)
- Enables fast fixture iteration without validation gymnastics
- Does NOT remove or weaken existing allowlists — purely additive

## Files Modified

- `hyoka/internal/validate/validate.go` (lines 15-31) — Extended `ValidServices`, `ValidLanguages`, `ValidCategories`

## Verification

- `go test -race ./internal/validate/...` — all green
- `hyoka validate` — all 90 prompts valid (including test fixture)
- End-to-end eval with `language: test` — runs cleanly, loads `criteria/language/test.yaml` graders

---


### Decision: Per-Bucket Grader Input Isolation (Grader Display Fix V2) (2026-04-24T04:36:24Z)

**Agent:** Tank 🎖️  
**Branch:** ronniegeraghty/dev  
**Commits:** `609ff869` (third attempt; chain: `4adc9288` → `84b1606d` → `609ff869`)

## Context

After commit `84b1606d` removed `KindPromptReview` from `validTypedKinds`, users were still seeing duplicate/mixed AI grader output. Both review buckets ("Criteria from prompt file" and "combined") were displaying the SAME 8 criteria, instead of being properly separated.

## Investigation

Live eval analysis revealed:
- "Criteria from prompt file" bucket: 8 criteria (should be 5)
- "combined" bucket: 8 criteria (should be 1)
- Both buckets contained ALL criteria from both the prompt and attribute-matched files

Root cause: In `engine_eval.go` (lines 632-634), the per-bucket grader input construction was:
```go
bucketInput := graderInput
bucketInput.EvalCriteriaBuckets = []graders.ReviewBucket{bucket}
```

This copied the parent `graderInput`, which included the merged `EvalCriteria` field containing ALL criteria. When `PromptReviewGrader.gradePanel()` processed a single bucket, it used the merged `EvalCriteria` instead of the bucket-specific `EvalCriteriaBuckets[0].Criteria`.

## Decision

Clear the merged `EvalCriteria` field after setting `EvalCriteriaBuckets` to enforce bucket isolation:
```go
bucketInput := graderInput
bucketInput.EvalCriteriaBuckets = []graders.ReviewBucket{bucket}
bucketInput.EvalCriteria = ""  // Force use of bucket-specific criteria
```

## Rationale

1. **Bucket Isolation:** Each bucket should see ONLY its own criteria, not the merged view
2. **Fallback Semantics:** The `EvalCriteria` field is a fallback when `EvalCriteriaBuckets` is empty or has one bucket; clearing it forces the grader to use the bucket's criteria
3. **Minimal Change:** One-line fix; no changes to the bucket construction logic or PromptReviewGrader
4. **Testable:** Live eval immediately shows the correct separation of criteria

## Verification

BEFORE fix (run 20260424-042347):
```
"Criteria from prompt file" bucket:
  - 8 criteria (all prompt + all attribute-matched)
  
"combined" bucket:
  - 8 criteria (duplicate)
```

AFTER fix (run 20260424-043129):
```
"Criteria from prompt file" bucket:
  - 5 criteria (prompt-only):
    • Installing azure-keyvault-secrets...
    • Creating a SecretClient...
    • set_secret(), get_secret()...
    • Handling soft-delete...
    • Exception handling...

"combined" bucket:
  - 1 criterion (attribute-matched only):
    • DefaultAzureCredential Authentication — Check the following criteria...
```

## Additional Changes

Updated `criteria/language/python.yaml` DefaultAzureCredential grader prompt to list two distinct criteria (authentication + async/await) to make any future mis-grouping immediately visible in test runs.

## Impact

- ✅ Per-bucket grader display now works correctly
- ✅ No duplication of criteria across buckets
- ✅ Each grader shows only its intended criteria
- ✅ Report JSON structure is correct
- ⚠️ AI reviewer still treats multi-point prompts as single criteria (future work if needed)

## Alternatives Considered

1. **Modify PromptReviewGrader.gradePanel():** Check `EvalCriteriaBuckets` first, then fall back to `EvalCriteria`. Rejected: More complex, adds conditional logic to the grader.
2. **Change bucket construction:** Separate prompt and attribute-matched into different inputs. Rejected: Bigger change, affects multiple files.

## Related

- Issue: User report of duplicate AI grader display
- Previous fix: Commit `84b1606d` (removed KindPromptReview from validTypedKinds)
- Phase 2 per-bucket work: Commit `4adc9288` (created per-bucket graders)
- **Pattern:** ALWAYS clear merged `EvalCriteria` when setting `EvalCriteriaBuckets` in per-bucket grader input construction

---

### Decision: KindPromptReview Removed from validTypedKinds (2026-04-24T03:59:28Z)

**Agent:** Tank 🎖️  
**Branch:** ronniegeraghty/dev  
**Commits:** `84b1606d`, `a37763f3`

## Context

After commit 4adc9288 (per-grader display refactor), user reported potential duplicate AI grader execution. Investigation revealed that `graders.KindPromptReview` was incorrectly listed in `validTypedKinds` in `hyoka/internal/criteria/config.go`.

## Problem

`KindPromptReview` being in `validTypedKinds` suggested that `type: prompt_review` was a valid criteria-file type, but:

1. `NewGrader` in `graders/registry.go` doesn't handle `KindPromptReview` — it would error on instantiation
2. `PromptReviewGrader` instances are created manually by the engine in `engine_eval.go` Phase 2 (lines 596-671), not via the criteria-file system
3. Criteria YAML files should only use `type: prompt` for LLM-review graders, never `type: prompt_review`

## Decision

**Removed `graders.KindPromptReview` from the `validTypedKinds` map.**

The valid typed kinds are now:
- `file` — file existence/content graders
- `program` — external program graders
- `behavior` — tool usage behavioral graders
- `action_sequence` — action timeline graders
- `tool_constraint` — tool call constraint graders
- `output_check` — output validation graders

`KindPromptReview` is NOT a criteria-file type — it's the kind returned by runtime `PromptReviewGrader` instances created by the engine's per-bucket review loop.

## Rationale

1. **Schema correctness:** Criteria files define grader configuration, not runtime grader kinds. `type: prompt` entries are partitioned into `promptEntries` by `PartitionMatched` and used to build review buckets, not instantiated via `NewGrader`.

2. **Error prevention:** If someone added a `type: prompt_review` entry to a criteria file, it would pass validation (because it was in `validTypedKinds`) but fail during instantiation (because `NewGrader` doesn't handle it).

3. **Clear separation:** The typed graders path (Phase 1) uses `NewGrader` for instantiation. The AI review path (Phase 2) manually creates `PromptReviewGrader` instances, one per bucket.

## Implementation

**File:** `hyoka/internal/criteria/config.go`  
**Change:** Removed `graders.KindPromptReview` from the `validTypedKinds` map, added comment clarifying that `KindPromptReview` is not included because it's the kind of manually-created instances, not a valid criteria-file type.

## Verification

- Build: ✅ `go build ./...`
- Tests: ✅ `go test ./hyoka/internal/criteria/... ./hyoka/internal/eval/...`
- Live eval: ✅ Completed successfully (no errors)

## Related

- Commit 4adc9288: Per-grader display refactor (created per-bucket review loop)
- `engine_eval.go` lines 596-671: Phase 2 per-bucket grader creation
- `buckets.go`: `BuildUnifiedReviewBuckets` — creates review buckets from `type: prompt` entries
- `registry.go`: `NewGrader` — handles typed graders (does NOT handle `KindPromptReview`)

---

### Decision: Graders Run on Every Eval; Generator Response Threaded Through (2026-04-24T00:56:09Z)

**Agent:** Neo 💊
**Branch:** ronniegeraghty/dev
**Commit:** `8794e70b`

## Context

The grading pipeline was historically guarded with `if len(generatedFiles) > 0` (engine_eval.go:500), preventing graders from running when the agent produced no files. This was a legacy artifact from when the engine itself enforced a "no files generated" failure condition. This guard created two problems:

1. **Bug 1 root cause:** When agents generated zero files, the grader block was skipped entirely, even if graders had validation logic (`output_check` with `min_files: 1` to fail intentionally).
2. **No response-only evals:** Prompts evaluating agent reasoning, planning, or explanation text could not run graders because the guard prevented them from executing.

**User directive:** "Graders should run on every eval, even when no files are produced; pass generator's final response through to graders for response-only evaluations."

## Decision

**Graders now run on every eval, regardless of file count.**

1. **Removed guard in engine_eval.go:500** — Deleted the conditional that skipped grader execution when generated files were empty.
2. **Added `AgentFinalResponse` to `GraderInput`** — hyoka/internal/criteria/graders/types.go now carries the agent's final assistant message through the grading pipeline.
3. **Implemented `extractLastAssistantMessage` helper** — Scans session events backwards to capture the last assistant message, populates `EvalResult.FinalResponse` in success/action-limit/error paths.
4. **Graders have autonomy** — Individual graders decide whether to use:
   - `GraderInput.WorkspaceDelta` (file changes)
   - `GraderInput.AgentFinalResponse` (agent's last message)
   - Both
   - Neither (e.g., action-sequence graders that only look at ActionLog)

## Rationale

1. **Configurable graders, not engine guards:** The `output_check` grader could never enforce file-count rules on empty workspaces because the guard prevented it from running. Engine-level guards must not pre-empt configurable graders.

2. **Agent responses are first-class artifacts:** Some prompts evaluate the agent's reasoning, planning, or explanation *text*, not files. Examples:
   - "Explain the trade-offs between approach A and B"
   - "Recommend a design pattern for this scenario"
   - "Analyze these requirements and propose a plan"

3. **Grader autonomy:** Each grader is responsible for its own scoring logic. If a grader doesn't care about empty workspaces, it passes. If it does care, it fails.

## Implementation

**Files changed:**
- `hyoka/internal/eval/engine_eval.go` — Removed `len(generatedFiles) > 0` guard
- `hyoka/internal/criteria/graders/types.go` — Added `AgentFinalResponse string` to GraderInput
- `hyoka/internal/eval/copilot.go` — Added `extractLastAssistantMessage` helper, populates `EvalResult.FinalResponse` in all terminal paths
- PromptReviewGrader and output_check — Now have access to AgentFinalResponse

## Verification

- All tests pass: `go test -race ./...`
- Build clean
- Live eval verified: generator response threaded through to graders
- No regressions on existing file-based evals

## Impact

- **Fixes Bug 1 root cause:** Graders now run on empty workspaces, enabling `output_check` and similar rules to fire
- **Enables new eval modes:** Prompts can now evaluate pure response quality without generating files
- **Backward compatible:** Existing evals unchanged — graders that don't use `AgentFinalResponse` ignore it

## Exception

The bundle-error check (`e.graderBundle.MatchingErrors(props)`) remains in place. Config errors still skip grading entirely — we don't trust partial results when the grader bundle couldn't be fully parsed.

## Future Enhancement

The `PromptReviewGrader` passes files to the review panel via `workDir`. When the workspace is empty, the panel has nothing to review. Future enhancement could write `AgentFinalResponse` to a file in `workDir` (e.g., `_agent_response.txt`) when no files exist, labeled as "Agent's final response" for the review panel.

---

### Decision: Plugins May Be Single-Skill OR Container; Container Plugins Fan Out Per Child (2026-04-23T19:42Z)

**Agent:** Neo 💊
**Branch:** ronniegeraghty/dev
**Commits:** `4a8c4a0d` (this fix); extends `2c1de1c0` and `3b306c9` (explicit `repo:` schema).
**Driver:** After the explicit-`repo:` migration, evals still failed with `tool_load_failure: plugin "azure-sdk-python" ... not found` — even with the cache populated. Root cause was a shape mismatch between the resolver (looking for one skill) and the upstream layout (a container of many skills).

**Finding:** `plugin.ResolveInstalled` (`hyoka/internal/plugin/installed.go:43`) used `isSkillDir`, which required a top-level `SKILL.md`. Container-style plugins (e.g. `azure-sdk-python` in `microsoft/skills`) have no top-level `SKILL.md`; they hold `skills/<child>/SKILL.md` for many children (41+ in this case). Resolver returned `""` → fan-out never ran → verifier had nothing to match against → hard fail.

**Decision:**
1. **A plugin is a directory of skills, not a single skill.** `ResolveInstalled` now accepts both shapes via a widened `isPluginDir` check: a top-level `SKILL.md` (single-skill plugin) OR a `skills/` subdirectory containing at least one `SKILL.md`-bearing child (container plugin).
2. **New helper `plugin.EnumerateChildSkills`** returns the absolute paths of each `<dir>/skills/<child>/SKILL.md`-bearing subdir, sorted lexicographically.
3. **`validatePluginEntry` fans containers out.** When children exist, it emits one `ToolLoadItem` per child with `ParentKind=plugin`, `ParentName=<plugin name>`, `Path=<child dir>`. Single-skill plugins emit one item as before.
4. **The verifier matches by child basename** — which is what the SDK actually reports in `SessionSkillsLoaded`. Never match against the parent container directory; it has no `SKILL.md` and never appears in the loaded set.

**Tests:** `TestResolveInstalled_ContainerPluginFanOut`, `TestResolveInstalled_SingleSkillPluginStillWorks`, `TestEnumerateChildSkills_IgnoresChildrenWithoutSkillMd`, `TestValidateAndExpand_RemoteContainerPlugin_FansOutChildren`. Full `go test ./hyoka/...` green.

**Verified live:** Pre-fix `hyoka run --prompt-id key-vault-dp-python-crud --config python-pairwise` → 3/3 evals errored. Post-fix → 0 errors, all generators report `Skills loaded: ..., azure-keyvault-py, azure-identity-py, azure-storage-blob-py, ...`.

**Relationship to prior decisions:** Extends — does NOT supersede — the explicit-`repo:` decision (commit `2c1de1c0`, dated 2026-04-23T18:50Z). That decision fixed the *locator* shape (where to fetch plugins from). This decision fixes the *content* shape (what a plugin actually contains on disk). Both contracts now hold simultaneously.

**Follow-ups (filed as issues):**
1. **No `hyoka plugin install` command.** When a remote plugin is missing, the validator's error message points users at Copilot CLI's `/plugin install`, which is misleading — that's a different tool. Today the cache must be populated by hand (`git clone github.com/<owner>/<repo> ~/.hyoka/cache/default/<owner>/<repo>`). Either ship a `hyoka plugin install` command or rewrite the error to document the manual steps.
2. **`pluginCheckedPaths` enumerates only parent dirs.** When a partial cache exists (parent present but children missing), the error lists candidate parent paths but not the child shape now expected. Refine the diagnostic.

**Reusable rule:** *Whatever the verifier checks must be the leaves the SDK actually loads.* If you ever introduce a new container shape, fan it out at validation time and verify the leaves — never the container.

---

### Decision: Default Model Pinned to claude-opus-4.7 for All Squad Agents (2026-04-23)

**By:** ronniegeraghty (via Coordinator)
**What:** Set `defaultModel: "claude-opus-4.7"` in `.squad/config.json`. All squad agent spawns use Claude Opus 4.7 regardless of role/task auto-selection, until the user clears the preference.
**Why:** User directive — "Can we update our squad characters so they are always using the Claude Opus 4.7 model for their work."
**Scope:** Layer 0 persistent config. Overrides default task-aware selection (Layer 3). Scribe (normally haiku) and Ralph also run on Opus 4.7 under this override.
**Reversal:** User can say "switch back to automatic" or "clear model preference" to remove.

---

### Decision: Remote Plugin Source Requires `@marketplace` Locator in Name (2026-04-23)

> **⚠️ SUPERSEDED (2026-04-23T18:50Z) by "Explicit `repo:` required on remote plugin entries; `@skills` magic removed (BREAKING)" — see entry at end of this file. The `@marketplace` locator convention has been removed; `repo:` is now the required form.**


**Agent:** Neo 💊
**Branch:** ronniegeraghty/dev
**Commit:** `769dea69`
**Driver:** Ronnie asked: "How does `source: remote` without a URL tell hyoka where to pull the plugin from?"

**Finding:** `configs/python-pairwise.yaml` declared `{name: azure-sdk-python, type: plugin, source: remote}` with no locator field. There is no dedicated `repo:` / `url:` field on plugin entries — the locator is encoded as a `@marketplace` suffix on `name` (e.g. `azure-sdk-python@skills` → microsoft/skills marketplace cache), parsed by `plugin.ResolveInstalled`. Without the suffix, resolution fell through to bare-name lookups under `~/.hyoka/cache/default/<name>/skills` and `~/.copilot/installed-plugins/<name>/skills` — producing the confusing "Checked:" path dump Ronnie flagged. Schema gap, not missing feature.

**Decision:**
1. **Keep the `@marketplace` locator convention.** Matches `/plugin install name@skills` UX; adding a full `repo:` field on plugin entries would duplicate the skill-fetcher pipeline without need.
2. **Reject bare `source: remote` at validation time.** `validatePluginEntry` now fails fast with a fix-it message.
3. **Fix the two broken configs** to use `azure-sdk-<lang>@skills`: `configs/python-pairwise.yaml` (1 entry), `configs/baseline-sonnet-skills.yaml` (5 entries).
4. **Future work (non-blocking):** If auto-fetch for plugins is ever wanted, mirror the skill flow with explicit `repo:`/`ref:` fields. Today `@marketplace` is the only supported locator and is sufficient.

**Files changed:** `hyoka/internal/config/tool/validate.go`, `hyoka/internal/config/tool/validate_test.go` (new `TestValidateAndExpand_RemotePluginMissingLocator`), `configs/python-pairwise.yaml`, `configs/baseline-sonnet-skills.yaml`.

**Verification:** `go build ./...` clean; `go test ./...` all green; `hyoka validate` — 89 prompts, 13 configs, 3 criteria files valid.

**Reusable pattern:** Any tool entry referencing remote content **must** carry an explicit locator. Skills: `repo:` (+ optional `ref:`/`version:`). Plugins: `@marketplace` suffix on `name`. Validation rejects remote entries missing their locator rather than letting failures surface downstream.

---

### Decision: Issue #305 (Hyoka v0.3.1) Status — Keep Open (2026-04-23)

**Agent:** Morpheus 🕶️  
**Finding:** Release does not exist (no v0.3.1 tag, no GitHub Release). Phase 0–2 work merged to main; Phases 3–6 complete on dev but NOT merged.  
**Decision:** Keep open. Issue is legitimate accountability mechanism for final integration/release steps (merge, tag, GitHub Release).

#### Rationale

Issue #305 was flagged as "probably stale" in prior audit. Correction: it's an active umbrella for unreleased work on dev. The absence of a tag is the core finding, not evidence of staleness. Code is phase-complete and ready for release, but the release ceremony has not happened.

#### Next Steps

Ronnie decides release intent:
- **Option A:** Ship v0.3.1 now (merge dev → main, tag, GitHub Release)
- **Option B:** Defer release to batch with Phase 7

See `.squad/decisions/inbox/morpheus-issue-305-status.md` for full audit trail.

---

### Decision: Issue #595 (useRuns Hook) — Left Open (2026-04-23)

**Agent:** Trinity 🖤  
**Issue:** [#595 — Extract useRuns hook for dashboard/prompts pages](https://github.com/ronniegeraghty/hyoka/issues/595)  
**Finding:** Hook does not exist. Duplicate fetch + cancellation code present in both `dashboard-page.tsx` and `prompts-page.tsx`.  
**Decision:** Leave open. Valid pending work item with clear acceptance criteria.

#### Acceptance Criteria

- [ ] Create `site/src/app/hooks/useRuns.ts` with shared hook returning `{ runs, loading, error }`
- [ ] Refactor both components to use the hook
- [ ] `cd site && npm test` passes
- [ ] No behavior change

Hook extraction is straightforward — ready for implementation as a follow-up task.

---

### Decision: Issue #290 (Criteria Table Layout) — Ready to Close (2026-04-23)

**Agent:** Trinity 🖤  
**Issue:** [#290 — Criteria table: put baseline config on the left](https://github.com/ronniegeraghty/hyoka/issues/290)  
**Finding:** Comparison matrix layout is correct. Baseline configs appear on the left (leftmost columns) in rendered reports.  
**Decision:** Close the issue. Layout requirement satisfied.

#### Evidence

- Example report: `/reports/20260423-172207/summary.md` (9 configs)
- Column order: baseline variants first, then non-baseline variants
- Phase 4/5 eval detail page redesign (issues #358, #572, #590) achieved the goal
- Code: `internal/report/markdown.go:410–451`, `internal/report/report_data.go:33–75`

#### Proposed Close Comment

> ✅ Verified complete. Phase 4/5 eval detail page redesign (#358, #572, #590) implemented criteria table improvements. Rendered reports confirm baseline configs now appear on the left (leftmost columns) in comparison tables.

---

### Decision: Agent Attempt Three-State Display — Streaming Replaced (2026-04-23)

**Status:** ✅ Implemented. Commits b17f1ef5 (code) + e9d590e6 (docs). Branch: `ronniegeraghty/dev`.  
**Directive:** User request after 4 failed attempts to fix live-tail rendering (commits 6b3d3d48, 42ea88fb, fe6efebf, 670c5dbf).  
**Driver:** Tank 📡 (implementation) + Ronnie Geraghty (directive).

#### Rationale

The Agent Attempt section in the interactive renderer used a streaming tail that displayed live activity ("Running… turn X/Y, N tool calls") but suffered from persistent line-wrapping leaks caused by:
- Wide characters (emoji, CJK) occupying 2 terminal cells
- slog stderr foreign writes between tail updates
- Terminal width edge cases
- Complex multi-row clearing math

Four sequential fix attempts addressed individual symptoms (multi-row clearing, truncation, cell-width, terminal margin) but failed to achieve stability. Root issue: **streaming variable-length content is fragile at terminal boundaries.**

**User decision:** Accept loss of real-time detail (activity messages, duration counter, tool call count) in exchange for UX stability. Replace with simple, bounded state machine.

#### Solution

Three-state display:

| State | When | Display |
|---|---|---|
| `Running` | Generator session started, no terminal event | `🔄 Running` |
| `Completed` | Session ended successfully (no guardrail) | `✅ Completed` |
| `Guardrail hit — {reason}` | A guardrail terminated the run | e.g., `Guardrail hit — turn limit (25)` |

Single-line bounded content (max ~35 chars) eliminates line-wrapping risk. No truncation math, no multi-row tracking, no ticker dependency. State transitions are event-driven.

#### Implementation

**Code changes (commit b17f1ef5):**
- Added `agentAttemptState` enum to `display_interactive.go` (Running, Completed, Guardrail)
- Removed streaming fields from `interactiveEval`: `agentActivity`, `agentToolCalls`, `agentFileCount`, `agentTurns`, `agentDuration`
- Added `agentState` + `agentGuardrailMsg` fields
- Simplified `renderAgentEvent()` and `renderAgentStateLine()` (state-driven, not streaming)
- Removed per-event updates from `tickLoop()`; three-state model is event-driven only
- Updated `display_interactive_test.go`: changed expected output "✅ Complete" → "✅ Completed"
- Added `GuardrailReason` field to `ProgressEvent` (events.go)
- Added `extractGuardrailShortReason()` helper in `engine.go` to convert verbose guardrail text ("guardrail: turn count 26 exceeded limit of 25") to compact form ("turn limit (25)")

**Verification:**
- `go build ./...` — clean
- `go test -race ./...` — all pass (24 packages)
- Manual test: `go run ./hyoka run --prompt-id key-vault-dp-python-crud --config "baseline/claude-opus-4.6" --log-level error` — state transitions work correctly

**Not yet tested (user should verify):**
- Triggering an actual guardrail (MaxTurns, MaxFiles) to confirm message renders
- Narrow terminal (stty cols 60) to verify no wrap

#### Docs Update (commit e9d590e6)

Added "Lesson: Agent Attempt UX Redesign" to architecture docs summarizing when bounded-state vs streaming is appropriate.

#### Supersedes

This decision **resolves and archives** all prior tail-leak fix work:
- Tail leak v1 fix: commit 6b3d3d48 (multi-row clearing)
- Tail leak v1 fix: commit 42ea88fb (truncation)
- Tail leak v2 fix: commit fe6efebf (wide character cell width)
- Tail leak v2 fix: commit 670c5dbf (foreign writes + terminal width margin)

Streaming approach is officially abandoned in favor of three-state machine.

---

### Decision: CLI Output UX Overhaul — Round 1–2 Stable Contracts (2026-04-22)

**Status:** ✅ Schema + emitters landed on `ronniegeraghty/dev`. Renderers (round 3) in flight.
**Sprint:** CLI Output UX Overhaul (interactive vs CI modes).
**Consolidated from 5 inbox entries** merged 2026-04-22T23-25-55Z:

| Item | Agent | Commit | Status |
|------|-------|--------|--------|
| ProgressEvent schema extension (6 new event types + fields + string consts) | Neo 💊 | `61d830c6` | ✅ Locked |
| Style helper package `internal/progress/style/` | Trinity 🖤 | `21636fdd` | ✅ Locked |
| Tool-resolution emission wiring (plugin → MCP → skill) | Neo 💊 | `e06ead61` | ✅ Shipped |
| Tool-verification emission wiring (SDK post-session) | Neo 💊 → Switch 🤍 | `82cd8590` (never merged) → re-landed via `25ce00a7` | ⚠️ Re-landed — see round-3/4 reconciliation below |
| Grader serialization + per-grader lifecycle events | Neo 💊 | `bffd0c40` | ✅ Shipped |
| Workers default flipped to 1 | Tank 📡 | `3b9cbab9` | ✅ Shipped |
| Progress mode auto-selection (TTY + worker count) | Tank 📡 | `d6fd0a59` | ✅ Shipped |

#### ProgressEvent schema (permanent contract)

New `EventType` values (appended to existing `iota` block — append-only):
`EventToolResolutionStart`, `EventToolResolutionResult`, `EventToolsVerified`, `EventGraderStart`, `EventGraderComplete`, `EventSessionDetails`.

New string constants (exported so emitters don't hardcode):
- `ToolKindSkill = "skill"`, `ToolKindPlugin = "plugin"`, `ToolKindMCP = "mcp"`
- `ToolStatusLoaded = "loaded"`, `ToolStatusFailed = "failed"`
- `GraderResultPass = "pass"`, `GraderResultFail = "fail"`

New fields on `ProgressEvent` (all in-process only, no JSON tags):
`ToolName`, `ToolKind`, `Status`, `Reason`, `Tools []ToolStatus`, `GraderID`, `GraderKind`, `Result`, `Score *float64`, `Files []string`, `Turns int`, `ToolCalls int`, `Cost float64`.

Helper type `ToolStatus{ToolName, ToolKind, Status, Reason}`. Fat-union pattern matches existing `ProgressEvent`. No constructor helpers — existing emitters use raw struct literals. `Score` is a pointer so "grader didn't report" is distinguishable from legitimate `0.0`.

#### Grader serialization & event policy

- Emission is gated on **reporter presence** (`sendRawEvent != nil`), not on worker count. Graders are already sequential in both modes (single `for` loop in `criteria.RunGraders`; `prompt_review` runs after typed graders in the same function).
- `GraderID` is stable between `Start` and `Complete`. For the review panel, `GraderID = "ai_review"`. For typed graders, `GraderID = Grader.Name()`.
- **`Score` policy:** populated (non-nil) only for `prompt_review` and `prompt` LLM-judge graders. All binary graders (`file`, `program`, `output_check`, `behavior`, `action_sequence`, `tool_constraint`) leave `Score` nil — rendering "(0/10)" for a binary fail would mislead.
- `Result` is always populated (`GraderResultPass` / `GraderResultFail`). `Message` is whatever the grader already put in `GraderResult.Message` — no fabrication.
- New API: `criteria.GraderHooks` + `criteria.RunGradersWithHooks`. `criteria.RunGraders` preserved as a zero-hooks shim. `engine.runSingleEval` gained a `sendRawEvent` sixth arg that auto-fills `EvalID`/`PromptID`/`ConfigName`.
- Report JSON (`GraderResults`, `ScoreBreakdown`, aggregate `Pass`/`Score`) unchanged — events are in-process UX only.

#### Tool-resolution wiring (what renderers can assume)

Emitted from `CopilotPromptRunner.buildSessionConfig` before the Copilot session starts, gated on `e.progressFn != nil`:

| Tool kind | Emitter | Loaded rule |
|-----------|---------|-------------|
| `plugin` | `config.ToolConfig.EmitPluginResolutions` | Resolves via plugin registry **or** installed Copilot CLI plugins; else Failed+"not found" |
| `mcp` | `tool.EmitMCPResolutions` | Local: `Command` set. Remote: `URL` set. Else Failed with reason |
| `skill` | `tool.ResolveSkillsWithReporter` | Resolves to ≥1 skill directory; else Failed (missing SKILL.md / empty dir) |

- **Sequential pairing:** every `EventToolResolutionStart(ToolName, ToolKind)` is followed by exactly one matching `EventToolResolutionResult` before any other tool-resolution event.
- **Order within a config:** plugins (declaration order under `plugins:`), then MCPs (declaration order in `generator.tools`), then skills (declaration order).
- **No concurrency** between events for a single eval — all synchronous on the eval goroutine.
- Nil reporter = silent no-op. `ResolveSkills` becomes a nil-emitter shim over `ResolveSkillsWithReporter`.

#### Tool-verification wiring (ordering guarantees)

Exactly one `EventToolsVerified` per eval is emitted after SDK `SessionSkillsLoaded` + `SessionMcpServersLoaded`:

1. **At-most-once** per eval (guarded by `verifiedEmitted` under OnEvent mutex).
2. **Fires after both load events** when both skills and MCP are configured; **fires after the single relevant event** when only one kind is configured; **never fires** when neither is configured.
3. **Never fires** when reporter is nil.
4. **Always fires before generation begins** — `CreateSession` finishes before `SendAndWait`. Renderers can treat `EventToolsVerified` as terminal for the Tools block; subsequent `EventToolStart` / `EventWritingFile` / `EventGraderStart` arrive after.
5. `emitToolsVerified` builds the slice under lock but invokes `progressFn` post-unlock — no deadlock risk for renderers holding their own locks.
6. `Tools` sorted by `(ToolKind, ToolName)` ascending — deterministic for snapshot tests.
7. Every configured tool appears exactly once, tagged Loaded if SDK confirmed, Failed otherwise with populated `Reason`. Tools reported by SDK but not configured are dropped (intent: "did what I asked for succeed?", not "what did SDK load?").
8. Plugin kind is **not** covered by verification — no SDK post-start signal; plugin status from config-time resolution remains authoritative.
9. Skill match is by basename (`filepath.Base(dir)`) against SDK-reported `Name`.

#### Style helper API (`internal/progress/style/`)

- **Package:** `github.com/ronniegeraghty/hyoka/hyoka/internal/progress/style`. Stdlib-only.
- **Constructors:** `New(w io.Writer) *Styler` (auto-detects TTY + `NO_COLOR`), `NewFromEnabled(bool) *Styler` (bypass for tests).
- **Detection:** disabled if `NO_COLOR` set, if writer isn't `*os.File`, or if file is not a char device. Else enabled.
- **Zero value / nil receiver are safe** — methods return raw text when disabled.
- **Color methods:** `Green/Red/Yellow/Cyan/Blue/Dim/Bold/Reset`.
- **Semantic helpers (preferred):** `OK` (green), `Fail` (red), `Warn` (yellow), `Info` (cyan), `Muted` (dim). Renderer code should use these so palette changes stay single-source.
- Does **not** strip existing ANSI from input, **not** handle 256/truecolor, **not** expose writer helpers, **not** auto-enable Windows VT mode.
- Tests: force `NewFromEnabled(false)` for plain-text goldens; `NewFromEnabled(true)` to assert ANSI presence. Never rely on `New(&bytes.Buffer{})` (always disabled).

#### Backward-compatibility guarantees (all rounds)

- `ProgressEvent` schema extension is additive only — no renames, reorders, or removals. Existing emitters and display paths compile + behave unchanged.
- `criteria.RunGraders` signature unchanged; existing callers and tests untouched.
- Report JSON format unchanged.
- All verifications: `go build ./...`, `go vet ./hyoka/...`, `go test -race ./hyoka/internal/{progress,eval,criteria,config/tool}/...` green on each commit.

---

### Decision: CLI Output UX Overhaul — Round 3 & 4 (Renderers, Tests, Docs, Verification) (2026-04-23)

**Status:** ✅ Sprint complete. Shipped on `ronniegeraghty/dev` at HEAD `2d38533f`.
**Consolidated from 4 inbox entries + 2 bug reports** merged 2026-04-23T00:05:04Z:
`neo-interactive-renderer.md`, `trinity-ci-renderer.md`, `switch-renderer-tests.md`, `switch-tool-verification-rerelease.md`, `switch-bug-ci-mode-suppressed-when-piped.md`, `switch-bug-clean-blocks-non-interactive.md`.

| Item | Agent | Commit | Status |
|------|-------|--------|--------|
| CI append-only renderer + summary table | Trinity 🖤 | `63e2c11f` | ✅ Shipped |
| Interactive renderer (tail-only tool+grader layout) | Neo 💊 | `a0105a9d` | ✅ Shipped |
| Docs refresh (README, getting-started, cli-reference) | Oracle 📖 | `32f4e6c9` | ✅ Shipped |
| Renderer snapshot tests (13 cases) | Switch 🤍 | `142da225` | ✅ Shipped |
| Event-wiring tests (35 cases) + ToolsVerified re-release | Switch 🤍 | `25ce00a7` | ✅ Shipped |
| Hot-fix: `--progress auto` order for piped CI | Tank 📡 | `2d38533f` | ✅ Shipped |

Total sprint: 15 commits, 48 new test cases, 2 regressions caught, 1 ledger discrepancy reconciled.

#### Interactive renderer (`display_interactive.go`)

- Mode string `"interactive"`; dispatched via `NewDisplay` mirroring Trinity's CI delegation pattern.
- Auto-mode: `workers==1` → `"interactive"`, `workers>1` → `"ci"`. Explicit `--progress live|log|ci|off` still overrides. Debug/info log level without `--log-file` downgrades `interactive`→`ci` to keep stderr slog out of cursor moves.
- **Tail-update protocol:** `writeLine` (immutable append), `writeTail` (freeze previous tail, write without newline), `rewriteTail` (`\r\x1b[2K` + text, same physical row), `freezeTail` (`\n`, clear `tailKind`).
- **One sanctioned exception:** `redrawToolsBlock` triggered only from `onToolsVerified` when ≥1 tool status flips. Sequence: DECSC `\x1b7` → `\x1b[<N>A\r` → rewrite N lines → DECRC `\x1b8`. `toolsVerified` flag guards against double redraws.
- Per-eval layout: `Prompt` / `Config` / `Tools:` / `Agent Attempt:` / `Session Details:` / `Graders:`. Sections omitted when their events never arrive.
- Ticker: 1 Hz, refreshes only the Agent Attempt duration counter.
- Multi-eval: interactive mode only selected at `workers==1`; queued evals print sequentially with a blank separator.
- Counters (`completed/passed/failed/errors`) updated in dispatch shim; final `Summary` written by `interactiveRenderer.finish()`.
- All styled output via `style.Styler` (`sty.OK/Fail/Muted/Info`); `bytes.Buffer` writers get plain text.

#### CI renderer (`display_ci.go`)

- Mode string `"ci"` (preferred); `"log"` kept as a non-breaking alias — existing CI scripts get the new output with no flag change. Legacy per-phase/inline log behavior is gone.
- **Event → line mapping:** only `EventStarting` / `EventPassed` / `EventFailed` / `EventError` produce output. `EventGraderStart`/`Complete` update per-eval tallies (`graderPass` / `graderTotal`) keyed on `evt.EvalID` — interleaved graders across parallel evals are attributed correctly. All other events are deliberately silent in CI mode.
- **Line format:** `[HH:MM:SS] ▶ start <promptID> | <configName>` (cyan/dim on TTY), `[HH:MM:SS] ✅ pass … (<dur>, G/T graders)`, `[HH:MM:SS] ❌ fail … (<dur>, G/T graders) — <reason>`. `EventError` renders as FAIL with default reason `"eval errored"`; `EventFailed` with empty `Message` defaults to `"graders failed"`. Reasons collapsed to one line via `oneLine()`.
- **Timestamps:** `[HH:MM:SS]` relative to renderer construction (`time.Since(startTime).Round(time.Second)`). Styled `Muted` when enabled.
- **Summary table:** rendered at `Finish()` after blank line + bold `Summary` header. Row order = first-seen eval order (`order []string` keyed on evalID). Column widths auto-sized to `max(header, cells)` via `len()`. Unicode box-drawing (`┌─┐│├┼┤└┴┘`) rendered unconditionally — valid UTF-8, works in GitHub Actions / Datadog / Splunk / `less`. Result column is literal `PASS` / `FAIL` (no emoji, no ANSI) so snapshot goldens stay clean.
- Footer: `N/M passed · report: <reportDir>` (plain text; `report:` omitted when empty).
- Glyph/color tied to `Styler.Enabled` — NO_COLOR or non-TTY drops both emoji and SGR codes; box borders survive.

#### Auto-mode resolution (after hot-fix `2d38533f`)

Pure function `resolveAutoProgress(workers, isTerminal, logLevel, logFile) string` in `cmd/run.go`. Case order (after the fix):

1. `workers > 1` → `"ci"` (the CI renderer is exactly what should engage in piped/CI contexts).
2. `!isTerminal(os.Stdout)` → `"off"` (single-eval non-TTY is silent unless forced).
3. Single-eval TTY → `"interactive"`.
4. `--log-file` exception preserved verbatim from `3b9cbab9` as a post-pass.

Explicit `--progress` flag always overrides. Regression guarded by table-driven tests in `cmd/cmd_test.go`.

#### Tool-verification reconciliation (82cd8590 never merged)

The round-1/2 ledger marked `82cd8590` as ✅ Shipped, but `git merge-base --is-ancestor 82cd8590 HEAD` returns non-zero on `ronniegeraghty/dev`. The commit sits on a parallel branch that diverged from `bffd0c40`; dev went on to include `e06ead61` while `82cd8590` never merged.

Switch re-landed the emission inside `25ce00a7` in a more testable shape:

- New `hyoka/internal/eval/tool_verification.go` exporting `toolVerifier` with `newToolVerifier`, `onSkillsLoaded`, `onMCPLoaded`, `emitIfReady`.
- `copilot.go` OnEvent handler constructs one per eval, calls `onXLoaded` under `mu.Lock()`, invokes `progressFn(EventToolsVerified)` **after** `mu.Unlock()` — preserves the "build under lock, dispatch outside lock" guarantee from round 1–2.
- Contract preserved: at-most-once per eval; fires after both SDK load events when both kinds configured / after the single relevant event when one is configured / never when neither; never when reporter nil; before any generation event; deterministic `(ToolKind, ToolName)` sort; configured-only payload; plugins excluded; skill match by basename.
- Preserved slog paths: `lg.Warn("Expected MCP server not loaded", ...)` and `"No MCP servers loaded despite configuration"`.
- 9 table-driven tests in `hyoka/internal/eval/tool_verification_test.go`.

#### Test coverage (Switch, 48 new cases total)

**Renderer snapshots — `142da225` (13 cases):**

- `display_interactive_test.go` extended: 6 scenario cases (`happy_path_one_tool_two_graders`, `tool_load_failure_at_resolution`, `tools_verified_flip_loaded_to_failed`, `grader_fail_one_pass_one_fail`, `error_path_generation_error`, + `NoColorEnvDropsColor` via `t.Setenv`) and dedicated `ANSIMarkers` test asserting tail-update `\r\x1b[2K` and DECSC/DECRC bracket sequencing. Retains Neo's 3 pre-existing happy-path tests.
- `display_ci_test.go` new: 5 scenario cases (`happy_path_three_evals_all_pass`, `mixed_two_pass_one_fail_with_reason`, `multi_eval_interleaved_graders`, `no_color_drops_emoji_keeps_box_borders`, `zero_evals_empty_summary_does_not_crash`) + `TestCIRenderer_HappyPathSnapshot` full-output golden.
- Infra: `normalizeCI` regex stripper (`[HH:MM:SS]`, `(Ns, G/T graders)`, Duration-column cells) with placeholders `[HH:MM:SS]` / `DUR`; `feedInteractive` / `feedCI` helpers; `floatPtr` for grader scores. No `testdata/` needed — inline string literals + regex normalization sufficed.

**Event wiring — `25ce00a7` (35 cases):** unit tests for tool-resolution pairing/ordering (plugins → MCPs → skills, each group in declaration order, one Start/Result pair per tool before any next), grader event attribution across interleaved evals, `tool_verification.go` behavior (9 cases), session-details propagation, terminal events (`EventPassed` / `EventFailed` / `EventError`) all routed through counter dispatch.

#### Docs refresh (`32f4e6c9`)

- README, `docs/getting-started.md`, `docs/cli-reference.md` updated.
- Documented: `workers=1` default, `--progress interactive|ci|live|log|auto|off` values, `live`→`interactive` and `log`→`ci` aliases (kept as aliases rather than hidden so existing CI scripts remain greppable), auto-selection matrix (TTY × worker count).
- NO_COLOR behavior documented as **OR** condition: disabled if `NO_COLOR=1` **or** stdout is non-TTY. Common "both required" misreading explicitly prevented.
- Sample layout blocks are verbatim from the sprint plan — future renderer tweaks detectable via diff against the golden text.
- Suggested future: `docs/progress.md` with "Renderer snapshot coverage" subsection pointing at `display_interactive_test.go` + `display_ci_test.go` as canonical renderer contract (deferred — not blocking).

#### Resolved sprint regressions

| Row | Repro | Root cause | Fix |
|-----|-------|------------|-----|
| Matrix rows 6 & 8 | `--workers 4 > out.log` — no timestamped start lines, no summary table | `--progress auto` resolution in `cmd/run.go` checked `!IsTerminal` **before** worker count, so any piped invocation fell through to `off` regardless of worker count | Tank `2d38533f` — refactored to pure `resolveAutoProgress(...)`, reordered so `workers>1` wins first. Regression test in `cmd/cmd_test.go`. |

#### Verification

- Final `go build ./...`, `go vet ./hyoka/...`, `go test -race ./hyoka/...` green at HEAD `2d38533f`.
- 8-row manual verification matrix (workers × TTY × log-file) run by Switch on the built binary. 6/8 passed initially → 7/8 after Tank's hot-fix. The remaining fail row is unrelated (see Known Issues).

---

### Known Issues (non-blocking, out-of-sprint)

#### `hyoka clean` blocks in non-interactive contexts

**Filed by:** Switch 🤍 during matrix verification setup. **Status:** OPEN. **Not a sprint deliverable.**

Running `./hyoka-bin clean` with stdin not attached to a terminal hangs indefinitely at the `Kill these N process(es)? [y/N]` prompt. No input is possible; the call blocks the next scripted step.

Impact:

- `AGENTS.md` instructs agents to run `hyoka clean` after each test run, but agent shells have no interactive stdin.
- Any CI workflow that chains `hyoka clean` as a cleanup step hangs until timeout.

Suggested fix: add `-y / --yes` flag and/or auto-confirm when stdin is not a TTY; emit a deterministic exit code when there's nothing to clean so scripted callers can branch. Preexisting bug — not introduced by this sprint.

---

### Decision: Grader Unification — Phases 1–4 Shipped + Option A Restructure (2026-04-22)

**Status:** ✅ Code-complete across Phases 1–4. Option A package restructure merged.
**Branch:** `ronniegeraghty/dev` (direct commits, no PRs).

Consolidated from 8 inbox entries merged 2026-04-22T23:05:53Z:

| Item | Agent | Commit / Issue | Status |
|------|-------|----------------|--------|
| Morpheus grader unification proposal (schema redesign + roadmap) | Morpheus 🕶️ | — | Proposed → accepted |
| User directive: grader package layout = **Option A** (nested under `hyoka/internal/criteria/`) | Ronnie (via Copilot) | 2026-04-22T22:11Z | Locked |
| Morpheus Option A replan | Morpheus 🕶️ | — | Approved |
| Phase 1: Unified schema + back-compat loader (#624) | Neo 💊 | `faf556eb` | ✅ Shipped |
| Phase 1 test coverage | Switch 🤍 | — | ✅ All green |
| Phase 2: Engine cutover (#625) | Neo 💊 | `a8a6d2d4` | ✅ Shipped |
| Phase 3: `internal/criteria/` deleted (#626) | Neo 💊 | `46b624fb` | ✅ Shipped |
| Phase 4: `output_check` workspace-delta grader | Tank 📡 | `ad2a8ce7` | ✅ Shipped |
| Option A package restructure (supersedes phase 2/3 flat layout) | Neo 💊 | `46ddda2e` | ✅ Shipped |

**Package layout (Option A, locked):**
- `hyoka/internal/criteria/` — file-level concerns: `Bundle`, `LoadUnifiedDir`, `UnifiedGraderConfig`, `UnifiedGraderEntry`, `MatchingErrors`, `PartitionMatched`, `BuildUnifiedReviewBuckets`, `InstantiateGraders`, `RunGraders`.
- `hyoka/internal/criteria/graders/` — typed grader implementations.

Full per-agent shipping notes were merged from `.squad/decisions/inbox/` and the inbox cleared. See git history on `ronniegeraghty/dev` for commit-level detail.

---

### Decision: Grader Unification — Locked Answers + Phase 1–4 Issues Filed (2026-04-22)

**Author:** User (Ronnie) + Morpheus 🕶️  
**Date:** 2026-04-22  
**Status:** ✅ Schema locked. Phase 1 (#624) ready for Neo. Phases 2–4 sequential.

**Context:** Morpheus filed 744-line grader unification proposal covering schema redesign and implementation roadmap. All 10 open questions now answered by user; schema is final and ready for implementation. Issues #624–#627 filed and labeled `squad` + `squad:neo`.

**Locked Answers (Q1–Q10):**

1. **CLI flag naming:** Keep `--criteria-dir` (no rename).
2. **`type` placement:** Flat `type` field at entry level (not `kind`, not nested block). Prompt graders use same shape as other graders.
3. **`Gate` on typed graders:** Hard fail (consistent with `AggregateResults`).
4. **`isolate` on typed graders:** Load-time warning (silent ignore + warn).
5. **Multiple same-type graders:** Allowed. Uniqueness enforced by `name`, not `type`.
6. **`internal/criteria/` deletion:** Immediate in Phase 3. No deprecation shim.
7. **`output_check` v1 knobs:** Ship `min_files`, `max_files`, file presence, per-file size checks. Defer globs/regex.
8. **Verification strategy:** No golden-file or parallel-run gate. "hyoka isn't stable, so we don't care about breaking."
9. **`Gate` on prompt graders:** Reject at load time (LLM scores too noisy).
10. **Grader docs:** Document all graders (`prompt`, `output_check`, `action_sequence`, `tool_constraint`, etc.) in user-facing docs.

**Issues filed (all labeled `squad` + `squad:neo`):**

| Issue | Title | 
|-------|-------|
| #624 | [Grader Unification] Phase 1: Unified schema + back-compat loader in internal/graders/ |
| #625 | [Grader Unification] Phase 2: Unified execution path in engine (cut over to internal/graders/) |
| #626 | [Grader Unification] Phase 3: Delete internal/criteria/ package |
| #627 | [Grader Unification] Phase 4: Default output_check criteria + per-grader docs |

**Full proposal reference:** `.squad/decisions/inbox/morpheus-grader-unification-proposal.md` (744 lines — kept until Neo completes Phase 1).

**Next:** Coordinator hands #624 to Neo. Tank and Trinity standby for Phases 2–4.

---


### Decision: Morpheus — Examples Validation Audit & PR #607 Hierarchical-When Investigation (2026-04-22)

**Author:** Morpheus 🏗️
**Date:** 2026-04-22
**Scope:** Two parallel investigations

#### Run 1: Examples Directory Validation

**Status:** ✅ Audit complete. All 8 real artifacts valid. 1 intentional skeleton (`prompt-template.prompt.md`).

**Findings:**

| Kind | File | Status |
|---|---|---|
| prompt | `example.prompt.yaml` | ✅ |
| prompt | `graders-frontmatter-example.prompt.md` | ✅ |
| prompt | `existing-files-example.prompt.md` (+ `.starters/`) | ✅ |
| prompt | `prompt-template.prompt.md` | ❌ intentional skeleton |
| config | `example-full.yaml` | ✅ |
| config | `example-generator-skills.yaml` | ✅ |
| config | `example-remote-skill.yaml` | ✅ |
| criteria | `hierarchical-when-example.yaml` | ✅ |
| criteria | `language/*.yaml` (5 files) | ✅ |
| criteria | `service/*.yaml` (2 files) | ✅ |

**Recommended follow-up (low priority):** Rename `prompt-template.prompt.md` → `prompt-template.md` to skip validate scope. Keeps template human-readable; no schema fix needed. No urgent action — `examples/` is documentation, not active library.

**Learnings appended to Morpheus history.md:** Staging quirk — `starter_project` paths resolve relative to prompt file's directory; when staging with prefix, starter dir symlink must exist under unprefixed name.

#### Run 2: PR #607 Hierarchical-When Investigation

**Status:** ❌ Investigation surfaced critical bugs.

**Finding:** `examples/criteria/hierarchical-when-example.yaml` uses YAML `---` doc separator to suggest two group-level `when` blocks are valid. **They are not.** Root cause: `hyoka/internal/criteria/criteria.go:134-136` uses `yaml.NewDecoder` + single `Decode()` call, processing only the first YAML document. The Rust block (lines 46-66) is silently discarded at load time. Validate doesn't flag it because validation operates on parsed structure, not raw bytes.

**Correct schema:** Top-level `groups:` list (each `GraderGroup` has optional `when`). Canonical example in `hyoka/internal/criteria/hierarchical_test.go:208-247`.

**Risk:** Anyone copying the example loses half their criteria silently. No error, no warning.

**Recommended follow-ups (file as separate GitHub issues):**

1. **Fix the example** — rewrite `examples/criteria/hierarchical-when-example.yaml` to use `groups:` list instead of `---` separator. Keep file-level + grader-level demonstration on Python side; move Rust scope into group entry.
2. **Fix the loader** — `hyoka/internal/criteria/criteria.go:loadFile` (lines 134-136) should either:
   - Strict rejection: loop decoder, reject if more than one document (safest first move).
   - OR merge documents (user-friendly but changes semantics — risky without demand signal).
3. **Validate coverage gap** — `hyoka validate` should detect trailing YAML documents in criteria files as a footgun, independent of loader fix.

**Threaded reply posted** on PR #607 (comment 3125721580) documenting the silent-truncation bug and recommended fixes.

**Learnings appended:**
- Neo history.md: "Loader silently drops YAML docs after first --- in criteria files; affects `hyoka/internal/criteria/criteria.go:134-136`." (Neo owns the fix)
- Oracle history.md: "Examples can have misleading patterns (e.g., hierarchical-when-example.yaml uses discarded YAML docs); audit examples during docs maintenance cycles."


---

## Archive

Decisions prior to 2026-04-22 have been moved to `.squad/decisions/archive/`. See `pre-2026-04-22-decisions-archived-2026-04-23T18-29-46Z.md` for the most recent archival.
### 2026-04-23T18:50Z: Explicit `repo:` required on remote plugin entries; `@skills` magic removed (BREAKING)

**By:** Neo (per Ronnie directive)

**What:** Remote plugin entries in tool configs (`generator.tools` / `reviewer.tools`) MUST now declare a `repo:` field. The `name` field is a plain identifier. The `@marketplace` shorthand (e.g. `azure-sdk-python@skills`) and the hardcoded `microsoft/skills` alias inside `plugin.ResolveInstalled` have both been removed.

**New shape:**
```yaml
- name: azure-sdk-python
  type: plugin
  source: remote
  repo: github.com/microsoft/skills   # REQUIRED for source: remote
  version: main                        # OPTIONAL git ref
```

**Validator contract:**
- `source: remote` + missing `repo:` → hard-fail with: *"plugin … declares source: remote but has no repo: field. Add repo: github.com/microsoft/skills (or your fork) so hyoka knows where to fetch it from."*
- Any plugin name containing `@` → hard-fail with: *"plugin name … contains '@' — the @marketplace shorthand has been removed. … declare the source repo explicitly via repo: …"*

**Resolver contract:** `plugin.ResolveInstalled(repo, name)` looks under `~/.hyoka/cache/default/<owner>/<repo>/{.github/plugins,.github/skills,skills}/<name>` and the legacy `~/.copilot/installed-plugins/<owner>-<repo>/<name>/skills`. There is no implicit `microsoft/skills` fallback. `repo` accepts `owner/repo`, `github.com/owner/repo`, or `https://github.com/owner/repo[.git]` (normalized via the new exported `plugin.SplitOwnerRepo`).

**Why (Ronnie's framing, paraphrased):** A `source` field tells hyoka the *kind* of source; a `repo` field tells it the *location*. Both are required for any remote entry. Implicit aliases — even ones that "everybody knows" — are landmines for future readers of the config.

**Reverses:** The `@marketplace`-locator validator from commit `769dea69`. That commit closed a real schema gap but with the wrong primitive — it entrenched the magic alias instead of removing it. This decision is the corrective.

**Breaking change semantics:** Pre-1.0, no deprecation path. The cure for an implicit-magic schema isn't a softer warning; it's removal. Configs in this repo (`python-pairwise.yaml`, `baseline-sonnet-skills.yaml`) are migrated in the same commit. Downstream / wild configs WILL fail validation with a clear migration message — that's the point.

**Follow-ups:**
1. **Wild-config audit.** Any config not in this repo that previously used `name@skills` will fail validation. The error message tells users exactly what to change. Worth a CHANGELOG callout (Oracle).
2. **Plugin auto-fetch.** Today plugins must be pre-installed (via Copilot CLI `/plugin install` or manual placement) — `ResolveInstalled` is read-only. Now that `repo:` is explicit, a future `gitFetcher`-style auto-clone for plugins is a clean addition; the schema is ready for it.
3. **Skill `@owner/repo` form.** The `name@owner/repo` shorthand for `type: skill` is preserved (it's at least explicit), but it's less idiomatic than `name + repo:` separately. Consider deprecating in a follow-up if it causes confusion.
4. **`Version` is the git ref.** Reaffirmed: there is no separate `ref:` field. The `Version` field on `Entry` covers branches/tags/commits for both skills and plugins. Documented in `entry.go`.
### 2026-04-23: Canonical `repo:` form is `owner/repo`

**By:** Neo (requested by Ronnie)
**What:** In configs and docs, the `repo:` field for remote plugins is now written in the short canonical form `owner/repo` (e.g. `microsoft/skills`). The longer `github.com/owner/repo` form continues to validate and resolve correctly — `SplitOwnerRepo` strips the prefix — so existing user configs are unaffected. The change is purely about which shape we recommend and ship as examples.
**Why:** Remote plugins are GitHub-only, so the `github.com/` prefix is redundant noise. Keeping the long form supported preserves backward compatibility.


---

## Session 2026-04-23: Grader Points Rethink

Decisions merged from inbox during cleanup pass after 6-phase autopilot session.

### Decision: # Report data model review — fan-out + grader-Points alignment (2026-04-23T20:45Z)

**By:** Morpheus 🏗️


**Date:** 2026-04-23
**Branch:** `ronniegeraghty/dev` (read-only review, no edits)
**Requested by:** ronniegeraghty
**Trinity inbox:** _not yet present at write time — link when filed_

---

## TL;DR

The "all passed at the top, inconsistent per-row/per-page" symptom is a real bug, fully reproducible against `reports/20260423-195948`. Root cause is **two divergent roll-up paths**:

1. The **engine** computes `EvalReport.Success` from a single unified `agg.Pass` over the *internal* grader list (one entry per grader, `Pass bool`).
2. The **site** computes a per-row "graders passed" by counting `r.grader_results.filter(g => g.pass === true)` over the *report* list — which has been **expanded** by `convertGraderResults` so that one passing AI-review grader becomes 3 entries (panel members + consensus) all with `pass: null`.

Result: engine-truth `success=true`, site-derived `1/4` displayed in red on every row. Plus the "no files → grading skipped → empty grader_results, but Success=true" path produces `0/0` red badges for any config that produced no files.

The plugin / skill_dir fan-out makes the data-model gap worse: `report.SessionSetup` and `Environment` are both flat lists of names with no parent/child linkage, so the new "no plugin status, only child statuses" model cannot be expressed faithfully on disk yet.

---

## 1. Discrepancy reproduction

**Run:** `reports/20260423-195948` — 12 evals across 4 config groups, summary says `passed: 12, failed: 0, errors: 0` (i.e. 100%).

**Pages that disagree** (screenshots in `/tmp/morpheus-site-review/`):

| File | What it shows |
|---|---|
| `01-run-detail-header-vs-rows.png` | Run detail page: header shows `12 Passed`, but every visible table row shows a red `1/4` or `0/0` score badge. |
| `02-eval-detail-success-but-na-graders.png` | Eval detail page: green `✅ key-vault-dp-python-crud` headline (engine truth), then a `Grader Results (4)` section with `PASS / N/A / N/A / N/A` badges. |
| `03-run-detail-top.png` | Run detail header summary cards (12 passed, 100% pass rate) above the inconsistent rows. |

Concrete grader payload from one eval (`baseline/claude-sonnet-4.5/report.json`):

```json
"success": true,
"grader_results": [
  {"grader_name":"Output Files Exist", "grader_type":"output_check", "pass":true,  "score":1},
  {"grader_name":"claude-opus-4.6",    "grader_type":"review",       "pass":null, "score":null},
  {"grader_name":"gpt-4.1",            "grader_type":"review",       "pass":null, "score":null},
  {"grader_name":"consensus",          "grader_type":"review",       "pass":null, "score":null}
]
```

For the three `gpt-5.3-codex` configs in this run, `grader_results` is `[]` while `success` is still `true` (no files were generated, so the grading pipeline was skipped — see `engine_eval.go:433`).

---

## 2. Root cause(s)

### 2a. The site invents its own roll-up that disagrees with the engine

`site/src/app/components/run-detail-page.tsx:236`:

```ts
const gradersPassed = (r as EvalReport).grader_results?.filter(g => g.pass === true).length ?? 0;
const gradersTotal  = (r as EvalReport).grader_results?.length ?? 0;
```

`<ScoreBadge passed={gradersPassed} total={gradersTotal} />` → green only when `passed === total && total > 0`. With `pass:null` review entries this condition is mathematically impossible whenever an AI-review grader was run.

This is a roll-up that **does not exist on the engine side** — the engine's truth is `EvalReport.Success`, set from `agg.Pass` (`hyoka/internal/criteria/graders/grader.go:222-252`) over the *internal* grader list which has one entry per grader (`Pass bool`, never null).

### 2b. The expansion from internal-grader → report-grader is destructive

`hyoka/internal/eval/engine_eval.go:840-953`:

- `convertGraderResults` iterates `agg.Results` (each entry `Pass bool`, e.g. `ai_review` with `Pass=true`).
- For each `KindPromptReview` entry it calls `expandReviewGraderResult`, which **discards** the unified `Pass` and emits one `report.GraderResult` per panel member + one for consensus, **all with `Pass: nil`** and `Score: 0` — only `OverallScore` / `MaxScore` carry the actual review numbers.
- For typed graders (output_check, file, program, behavior) the conversion preserves `Pass` correctly (`engine_eval.go:847-855`).

So the on-disk shape of a single passing `ai_review` grader is *N indistinguishable-from-failing rows*. The engine's `agg.Pass=true` survives only as the top-level `EvalReport.Success` boolean — it is not represented inside `grader_results` at all.

### 2c. Empty-graders path leaks "success" without evidence

`engine_eval.go:433` (`if len(generatedFiles) > 0`): if no files were generated, the entire grading block is skipped. `EvalReport.Success` defaults to whatever the generator set — typically `true` from `engine.go:74` unless something later flipped it. Result: `success=true`, `grader_results=[]`. The site row renders `0/0` (red) and the eval-detail header renders ✅. Both are arguably wrong for different reasons.

### 2d. Where the roll-up is computed (audit)

| Location | Source of truth | Honors `EvalReport.Success`? |
|---|---|---|
| `engine_eval.go:191, 396, 423, 442, 564` | sets `Success` from `agg.Pass` and guardrails | n/a (this is the writer) |
| `report/summary_stats.go:116, 123, 131` | counts `r.Success` for config/prompt pass rates | ✅ |
| `report/markdown.go:29, 153, 437, 459` | renders pass/fail icons on Markdown reports | ✅ |
| `report/report_data.go:56` | builds `MatrixCell.Success` for cross-config matrix | ✅ |
| `serve/dashboard.go` | API endpoints: passes report objects through unchanged | ✅ (no roll-up) |
| `site/run-detail-page.tsx:76` (header summary) | `run.passed` from `summary.json` | ✅ |
| `site/run-detail-page.tsx:85-86` (filter) | `r.success` | ✅ |
| `site/run-detail-page.tsx:236` (row badge) | `grader_results.filter(g => g.pass === true)` | ❌ **DISAGREES** |
| `site/run-detail-page.tsx:360` (per-config card) | `result.success` | ✅ |
| `site/eval-detail-page.tsx:423, 441, 448` | `r.success` | ✅ |
| `site/GraderResultRow.tsx:16` (per-grader badge) | `result.pass` (null → "N/A") | ⚠ correct given the data, but the data is wrong |

**Only `run-detail-page.tsx:236-237` invents its own roll-up.** Everywhere else trusts `EvalReport.Success`.

---

## 3. Schema impact of the plugin / skill_dir fan-out

The Phase-1 display change (Tank, in flight) plus the already-shipped artifact-graph fan-out (commits `4a8c4a0d`, `370295d0`) push the report schema past what it can model.

### 3a. Today's `report.SessionSetup` is flat, parent-less

`hyoka/internal/report/types.go:326-341`:

```go
type ToolLoadResult struct {
    Name    string `json:"name"`
    Status  string `json:"status"`  // "loaded", "failed", "configured"
    Error   string `json:"error,omitempty"`
    Details string `json:"details,omitempty"`
}

type SessionSetupEvent struct {
    MCPServers   []ToolLoadResult `json:"mcp_servers,omitempty"`
    Skills       []ToolLoadResult `json:"skills,omitempty"`
    Tools        []string         `json:"tools_available,omitempty"`
    SystemPrompt string           `json:"system_prompt_status"`
    StarterFiles []string         `json:"starter_files,omitempty"`
}
```

Concrete observation from a real report: `session_setup.skills` has `[{name:"generator-skills", status:"configured"}]` (the parent skill_dir) while `environment.skillsLoaded` has 44 child skill names. The two lists are **not cross-referenced anywhere** — there's no record that the 44 children belong to `generator-skills`.

After Tank's Phase 1, plugin parents will have no Loaded/Failed status (only children carry status). The current `ToolLoadResult.Status` for a plugin parent has no valid value to write — `"configured"` is the closest fit but already gets used for skill_dirs.

### 3b. Required schema additions to `ToolLoadResult`

```go
type ToolLoadResult struct {
    Name       string `json:"name"`
    Status     string `json:"status,omitempty"`     // ← omitempty: parents can omit
    Error      string `json:"error,omitempty"`
    Details    string `json:"details,omitempty"`

    // NEW: parent/child linkage for fan-out
    Kind       string `json:"kind,omitempty"`       // "skill" | "mcp" | "plugin" | "skill_dir"
    Parent     string `json:"parent,omitempty"`     // name of the container this is a child of
    ParentKind string `json:"parent_kind,omitempty"`// "plugin" | "skill_dir"
}
```

Children carry `Status` and a `Parent`/`ParentKind` back-pointer; container parents emit a single row with `Status` empty (or a new `"container"` value) and no error semantics. Old reports remain valid because all new fields are `omitempty`.

### 3c. `Environment.skills_loaded` (site-side mirror) needs the same shape

`site/src/app/data/types.ts:180-192` — `skills_loaded` is `string[]`. The site classifies tools via `environment.skills_loaded.includes(tool)` (`eval-detail-page.tsx:501`). Once parent and child names both appear, exact-match classification loses the parent → child relationship.

Two viable options:

- **Option A (preferred):** change `skills_loaded` to `Array<{name: string, parent?: string, kind?: string}>` and bump the Go struct in `report/types.go` (`EnvironmentInfo.SkillsLoaded`) similarly. Requires site-side migration but expresses the truth.
- **Option B (cheap, less truthful):** keep `skills_loaded` flat, add a sibling `skill_groups: Array<{parent: string, children: string[]}>`. Site prefers `skill_groups` when present, falls back to flat list for old reports.

Either way, the migration story for already-on-disk reports is "fields default to empty / no grouping" — the site must handle absence gracefully.

### 3d. Migration path

- No `CurrentSchemaVersion` bump required for additive optional fields. Old reports load and render the same way they do today.
- A `MigrateToV2`-style migrator is **not** needed for the fan-out additions — there's no information in old reports that we can synthesize parent linkage from (parents are determined by config-time validation, which old reports didn't capture).

---

## 4. Schema impact of the grader `Points` field (Phase 2)

> Coordinating: the data model and the renderer; Trinity owns the visual presentation. This proposal is the data-side decision; Trinity will likely want to negotiate how Points render inside `GraderResultRow`.

### 4a. Drop-in field is safe and back-compat

```go
type GraderResult struct {
    // ... existing fields preserved ...
    Points []GraderPoint `json:"points,omitempty"`
}

type GraderPoint struct {
    Name    string `json:"name"`
    Pass    bool   `json:"pass"`
    Message string `json:"message,omitempty"`
}
```

`omitempty` keeps old reports byte-identical when re-encoded. No `SchemaVersion` bump strictly required.

### 4b. **However — Phase 2 should also fix the "expand into N entries" anti-pattern**

If Phase 2 just adds `Points` *next to* the existing expansion, we still emit 4 row-shaped entries per `ai_review` (with `Pass:nil` and `Score:0`). The 1/4 bug stays. The Points field becomes ornamental.

**Recommendation:** in Phase 2, change `convertGraderResults` for `KindPromptReview`:

- Emit **one** `report.GraderResult{GraderName: "ai_review", GraderType: "prompt_review", Pass: &aggregatedPass, Score: aggregatedScore, Points: <one per criterion>, ReviewDetails: <existing struct unchanged>}`.
- `ReviewDetails.PanelResults` and `ReviewDetails.Criteria` stay populated for the static Markdown/HTML templates that already iterate them.
- The site reads `Points` when present and falls back to `ReviewDetails.Criteria` for old reports.

This ALSO fixes the "1/4 in red" bug structurally, because `ai_review` becomes one passing row instead of three nil-pass rows.

### 4c. `EvalReport.SchemaVersion` bump?

If we change the *number of entries* in `grader_results` for the same input data (1 entry instead of N), that is a semantics change that older code paths could trip on. I'd bump to **v3** specifically to mark "review graders are no longer expanded; per-criterion data is in `Points`." `MigrateToV2` becomes `MigrateToV3` and old v2 reports keep their expanded shape on read (no de-expansion attempted — too lossy).

### 4d. Engine wiring for Points

Each `graders.GraderResult` needs to grow a `Points []GraderPoint` field upstream of the report layer (per the plan doc). The conversion in `convertGraderResults` then copies it into `report.GraderResult.Points`. Typed graders that today have only a single binary outcome populate one Point; output_check populates one per knob, file populates one per file checked, behavior one per constraint, prompt_review one per criterion. This matches the plan doc's table verbatim.

---

## 5. Recommended fixes (prioritized)

### Must-fix bugs (separate from Phase 2 — short, surgical commits)

1. **Site row-badge roll-up disagrees with engine truth.** `site/src/app/components/run-detail-page.tsx:236-237`. Either:
   - Show `r.success ? "✓" : "✗"` instead of a fraction (simplest, mirrors eval-detail header), OR
   - Count `g => g.pass !== false` (treats null as "not failed") so review rows don't drag the count down, OR
   - Surface `score_breakdown.final_score_pct` directly (most truthful — the engine already computed it).
   My pick: **show pass/fail icon mirroring `r.success`**, and if a numeric "X/N graders passed" is desired keep it as a *secondary* tooltip/sub-text using `pass !== false`. Single source of truth wins.

2. **Empty-graders → red `0/0` badge.** Same site location. When `grader_results.length === 0`, render `r.success ? "✓" : "✗"` rather than `0/0`. Optional second-step: have the engine refuse to set `Success=true` when no graders ran AND no files were generated — that's a separate engine fix worth filing.

3. **`GraderResultRow` showing `N/A` for review entries.** Cosmetic, but: until Phase 2 consolidates the expansion, render review-type entries with the `OverallScore/MaxScore` pair as a numeric badge (e.g. `8/10`) rather than `N/A`. The data is there.

### Schema evolutions (Phase 2 work — coordinate with Tank + Trinity)

4. **Stop expanding `ai_review` into N report entries.** Replace `expandReviewGraderResult` with a single-entry mapping that sets `Pass`, `Score`, `Points`, and keeps `ReviewDetails` populated. Bump `CurrentSchemaVersion` to 3. Site renders `Points` when present.

5. **Add `Points []GraderPoint` to `report.GraderResult`** and to the upstream `graders.GraderResult`. Each grader implementation populates at least one Point (per the plan doc table). Existing typed-detail structs stay alongside Points.

6. **Add `Parent`, `ParentKind`, `Kind` to `report.ToolLoadResult`.** Parents emit a row with `Status` empty (or new `"container"`), children carry runtime status + back-pointer. Old reports remain valid. Coordinate the rollout with Tank's display Phase 1 so the live renderer and the persisted schema land together.

7. **Enrich `report.EnvironmentInfo.SkillsLoaded` with parent linkage** (Option A or B from §3c). Surface in `site/src/app/data/types.ts` Environment type. Site classification (`eval-detail-page.tsx:501-507`) becomes group-aware.

### Nice-to-haves

8. **Single source of truth for the per-eval pass count.** Add `EvalReport.GradersPassed int` and `EvalReport.GradersTotal int` populated at engine time from `agg.Results` (BEFORE expansion). The site reads these directly instead of recomputing. Eliminates the entire class of roll-up-divergence bugs by construction. Trivial to add and `omitempty`-safe.

9. **Audit `success=true` semantics for the no-files case.** `engine_eval.go:433` skips grading when `generatedFiles == 0`. There's no positive evidence of success in that case — `Success` is only `true` because nothing flipped it `false`. Worth a separate decision: do we treat "agent produced nothing" as a hard fail, or do we keep the soft-pass for prompts that legitimately have no expected output? Today it leaks through as "passing."

---

## 6. Open questions for the user

1. **Run-detail row badge — what number do you actually want there?** Options: (a) just a pass/fail icon, (b) the score breakdown percentage, (c) a `passed/total` fraction that treats nil-pass review entries as passing. I'd go with (a) for clarity but (b) is the most informative.

2. **`success=true` with zero graders run** — bug or feature? Today, configs that produce no files end up green at the engine level. If you want them to fail loudly, that's a one-line engine fix; if you want to keep "no expected output → neutral pass," we should at least mark them in the report (`graders_skipped: true`) so the site can render them distinctly.

3. **Phase 2 ordering** — are we OK bumping `SchemaVersion` to v3 and breaking de-expansion for old `v2` reports (i.e. v2 reports keep their N-entry shape forever, v3 reports use the new 1-entry-with-Points shape)? The alternative (write a migrator that re-collapses N entries) is doable but loses the panel-member detail unless we carefully reconstruct it from `panel_results`.

4. **Plugin parent in `session_setup`** — when Tank's Phase 1 lands and the parent emits no status, do you want the parent to appear in the JSON at all (as a `Status:""` row that the site can use to draw the group header), or only in the children's `Parent` back-pointer? The former is friendlier to consumers that want to render the group as a tree; the latter is simpler.

---

## Boundary with Trinity

Per spawn instructions: I own the data model and roll-up logic. Trinity owns site UX (templates, CSS, page-by-page presentation). Where we overlap:

- The **fix to `run-detail-page.tsx:236`** is mine to recommend, hers to implement (it's a presentation choice masquerading as a data choice — the underlying fix is "stop expanding," but until that ships the site needs a behavior).
- The **`GraderResultRow` rendering of Points** is hers — I'm only saying that the Points field will exist and what it carries.
- The **session_setup → site Environment** structural change requires both: I propose the JSON shape, she chooses how to render the parent → child grouping in the tools-available list and any session-setup card.

Will link to her inbox file when it appears (none present at write time).

---

## Files cited

- `hyoka/internal/eval/engine_eval.go:191, 396, 423, 433, 442, 564, 826-953`
- `hyoka/internal/criteria/graders/grader.go:206-252`
- `hyoka/internal/report/types.go:22-48, 326-341, 700-711`
- `hyoka/internal/report/report_data.go:56`
- `hyoka/internal/report/summary_stats.go:116, 123, 131`
- `hyoka/internal/report/markdown.go:29, 153, 437, 459`
- `hyoka/internal/serve/dashboard.go:41-61`
- `site/src/app/components/run-detail-page.tsx:76, 85-86, 236-237, 360`
- `site/src/app/components/eval-detail-page.tsx:382, 423, 441, 448, 501-507`
- `site/src/app/components/GraderResultRow.tsx:16`
- `site/src/app/data/types.ts:180-192, 259-282`

Screenshots: `/tmp/morpheus-site-review/01-run-detail-header-vs-rows.png`, `02-eval-detail-success-but-na-graders.png`, `03-run-detail-top.png`.

### Decision: # Plugin tool-load assertions check leaves, not parents (2026-04-23T19:40Z)

**By:** Neo 💊


**Date:** 2026-04-23

**Decision:** When a remote plugin has the standard Copilot container layout (`<plugin>/skills/<child>/SKILL.md`), the validator MUST fan out into one report row per child skill. Tool-load assertions then check that each child loaded — never the parent plugin directory, which has no SKILL.md of its own and will never appear in the SDK's `SessionSkillsLoaded` event.

**Why:** The microsoft/skills `azure-sdk-python` plugin is a directory of 41 child skills, not a single skill. Asserting "did the plugin load?" by looking for the plugin's name in the SDK's loaded-skills list would always fail: the SDK loads the children, by their child basenames. The fan-out is what makes the assertion meaningful.

**Scope:** Applies to `plugin.ResolveInstalled` + `validatePluginEntry`. Single-skill plugins (top-level SKILL.md) keep the one-row-per-plugin behavior. Container plugins fan out.

**Reference:** Fix commit, plus `TestValidateAndExpand_RemoteContainerPlugin_FansOutChildren` for the regression guard.

### 2026-04-23T21:09Z: Phase 1 tool-loading display polish shipped
**By:** Tank (CLI/UX)
**What:** Five fixes shipped to `ronniegeraghty/dev` (3635a09f, 582ab59f, efe18373, 9f994107, fbcd9f38) covering: skill_dir parent name, plugin parent badge removal, child kind labels, frozen agent/grader row in-place rewrite, and removal of redundant tool events around the AI review grader. Phase 1 design intentionally keeps the grader handler open to Phase 2's `Points []GraderPoint` extension — Neo's work is not blocked.
**Why:** Make the live-eval interactive renderer match the design spec in `plan/2026-04/tool-loading-display-polish.md`. Validated via unit tests (race-clean) plus a 3/3 live eval run.


### Decision: # Site UX Review — pass/fail rendering against real "all passed" run (2026-04-23T20:46Z)

**By:** Trinity 🖤


**Date:** 2026-04-23
**Anchor run:** `reports/20260423-195948` — `passed=12 / failed=0 / errors=0` at the top, but per-eval pages contradict the verdict in three different places.
**Method:** Read `hyoka/internal/serve/` + `site/src/`, ran `go run . serve --port 8088`, drove with `playwright-cli`. Screenshots in `/tmp/trinity-site-review/`.
**Scope:** Presentation/UX only. Data-model and roll-up logic critique is Morpheus's territory — link to `morpheus-report-architecture-review.md` once it lands; flagged data shape gaps are listed in §5 below.

---

## 1. Discrepancy walk-through — one run, three contradictions

The anchor run's `summary.json` is unambiguous: `"passed": 12, "failed": 0, "errors": 0`, and every result's `success: true`. The site agrees at the top of the funnel and contradicts itself on the way down.

### 1a. `/runs` — the run card looks correct
*Screenshot:* `01-runs-list.png`, `07-runs-page.png`

```
Apr 23, 2026, 07:59 PM    12 evaluations    ✓ 12  ✗ 0    100.0%
```
Fine. `runs-page.tsx:136-138` divides `passed/total` from `summary.json` top-level. Honest.

### 1b. `/runs/20260423-195948` — every row claims **1/4** or **0/0** in red
*Screenshot:* `02-run-detail-table.png`

| Score (rendered) | Prompt | Model | Duration |
|---|---|---|---|
| **1/4** 🟥 | key-vault-dp-python-crud | claude-sonnet-4.5 | 101.4s |
| **0/0** 🟥 | key-vault-dp-python-crud | gpt-5.3-codex | 42.3s |
| **0/0** 🟥 | key-vault-dp-python-crud | gpt-5.3-codex | 19.6s |
| **1/4** 🟥 | key-vault-dp-python-crud | claude-opus-4.6 | 66.5s |
| … (rest mixed 1/4 and 0/0) | | | |

**Cause** — `run-detail-page.tsx:236-237`:
```ts
const gradersPassed = (r as EvalReport).grader_results?.filter(g => g.pass === true).length ?? 0;
const gradersTotal  = (r as EvalReport).grader_results?.length ?? 0;
```
Two compounding bugs:
1. **Tri-state collapse.** Strict `g.pass === true` excludes `pass: null`. In the JSON, `output_check` is the only grader whose backend sets `pass`. The three `review`-type graders (`claude-opus-4.6`, `gpt-4.1`, `consensus`) all have `pass: null` even though every criterion under them is `passed: true` and `overall_score === max_score`. So a fully-green eval renders as **1/4**.
2. **Schema drift.** The `gpt-5.3-codex` rows are stored without a `grader_results` array on the summary endpoint (`/api/runs/20260423-195948`) — confirmed via curl (see §5). Filter returns 0, total returns 0, ScoreBadge renders `0/0` and colours it red because `passed === total && total > 0` is false.

`ScoreBadge` (`run-detail-page.tsx:10-13`) is the wrong colour metaphor here regardless: it's already disagreeing with the per-row `r.success` boolean that lives one column to the right (er, would, if the table even showed `success` — see 1c).

### 1c. `/runs/.../eval/.../baseline/claude-sonnet-4.5` — header says 12/12 ✅, grader rows say **N/A** ⚪
*Screenshot:* `03-eval-detail-header.png`, `08-eval-detail-full.png`, `08b-eval-grader-section.png`

The hero score box shows `12 / 12` on an emerald-bordered card with a green check (line 423/441), driven by `r.success` and `review.overall_score / review.max_score`. Then below:

```
Grader Results (4)
  [PASS] Output Files Exist                Output Check        100%
  [N/A]  claude-opus-4.6                   Review • opus-4.6   6/6
  [N/A]  gpt-4.1                           Review • gpt-4.1    6/6
  [N/A]  consensus                         Review • consensus  12/12
```

Same root cause as 1b plus a UX choice: `GraderResultRow.tsx:16,29-39,68` has explicit tri-state (`true/false/null`) → `PASS/FAIL/N/A`. With three `pass: null` graders that scored `6/6` and `12/12`, the row literally renders `N/A` in grey. The user reads "PASS for output check, no opinion on the review graders" — but the score column directly to the right says `12/12`. The page is internally inconsistent in two consecutive flexbox cells.

### 1d. `/dashboard` — uncaught crash
*Screenshot:* `06-dashboard.png`, `09-dashboard-crash.png`

```
Unexpected Application Error!
Cannot read properties of undefined (reading 'toFixed')
TypeError: Cannot read properties of undefined (reading 'toFixed')
   at … index-Br3u5NeB.js:365:33160
   at Array.map …
```
React Router's default ErrorBoundary catches it; the user sees a stack trace. Reproducible by clicking "Dashboard" from the navbar at any moment — the navbar advertises a page that crashes. Not directly the same root bug as 1a–1c but it's the user's third click and shapes the impression that pass/fail data is unstable across the site.

### 1e. `/runs` rate column — error rows render as `0.0%` red
*Screenshot:* `11-runs-listing-bad-rows.png*

```
Apr 23, 2026, 07:21 PM    3 evaluations    0  0  3    0.0%
Apr 23, 2026, 07:19 PM    3 evaluations    0  0  3    0.0%
```
Three rows of `passed=0 failed=0 errors=3` — runs that errored before any grader ran. The card shows `0.0%` in the same emerald-fill bar as a successful run, just empty. Errors are aggregated into the denominator (`passed/total`), so an all-error run is indistinguishable visually from an all-fail run. There's also no badge/tag on the card saying "errored" — only a small amber count next to the X icon.

---

## 2. Page inventory — every site surface that touches grader/tools data

| Page | Component | Reads `success` | Reads `pass` | Reads `grader_results` | Reads tools | Plugin/skill-dir aware? | Issues |
|------|-----------|-----------------|--------------|-----------------------|-------------|-------------------------|--------|
| `/runs` | `runs-page.tsx` | indirectly via `summary.passed` | no | no | no | n/a | 1e: errors fold into denominator with no visual signal |
| `/runs/:id` | `run-detail-page.tsx` | yes (filter, line 85-86) | **yes (strict `=== true`)** | yes (per-row count) | shows top-2 mcp + top-1 skill flat | **no** | 1b: bad denominator + tri-state collapse |
| `/runs/:id/eval/...` | `eval-detail-page.tsx` | yes (header card, line 423/441) | via GraderResultRow | yes (full list) | flat skill/MCP/builtin tag per tool (line 495-535) | **no** | 1c: header agrees with `success`, rows disagree |
| `/runs/:id/eval/...` (Grader rows) | `GraderResultRow.tsx` | no | yes, tri-state | yes | n/a | n/a | tri-state UX correct in isolation, but every review-type grader currently lands on N/A |
| `/prompts` | `prompts-page.tsx` | computed pass-rate | no | no | no | n/a | leaks one-shot config names like `neo-test/missing-plugin` into "Worst" stat |
| `/prompts/:id` | `prompt-detail-page.tsx` | uses summary.passed | no | no | aggregates `Pass Rate by Tool Used` from invoked tools | **no** | "Pass Rate by Tool Used" lists tools as flat names — no plugin grouping (`azure.sdk_*` siblings show as separate rows) |
| `/dashboard` | `dashboard-page.tsx` | n/a | n/a | n/a | n/a | n/a | 1d: hard crash on `.toFixed` of undefined |
| `/pairwise` | `pairwise-page.tsx` | uses `success` | no | no | no | n/a | renders, but uses same `success` semantics as above; would inherit any roll-up change |
| `/compare` | `comparison-page.tsx` | uses `success` | no | no | no | n/a | same as pairwise |
| `/runs/:id` (tools cell) | inside `run-detail-page.tsx:264-279` | n/a | n/a | n/a | yes — top-2 mcp + top-1 skill | **no** | flat name list; would need plugin parent grouping post-Phase-1 too |

**Two structural gaps the templates can't currently see:**
1. Nothing in `site/src/app/data/types.ts` describes a plugin parent or skill-dir parent. `tool_availability` from the JSON is a flat array (`{name,type:"skill"|"mcp",available,used}`). No `parent_kind`, no `parent_name`, no children grouping. The renderer has no input from which to draw the parent/child distinction Phase-2 demands.
2. Grader rows have no `Points []` field on the wire today. Two existing structs (`OutputCheckGraderDetails.SubChecks`, `ReviewGraderDetails.criteria`) carry the underlying truth, but the unified `Points` shape from plan.md Phase 2 isn't represented in the TS types yet, and no template surfaces the per-criterion list anywhere except the giant criterion-by-criterion render at the bottom of the eval page (lines 615-660).

---

## 3. Pre-Phase-2 quick wins (independent of data-model changes)

These ship before Points lands and remove the false-negative impressions today.

### Q1. `run-detail-page.tsx:236-237` — stop counting `pass === true`; use the same truth the header uses

Today:
```ts
const gradersPassed = (r as EvalReport).grader_results?.filter(g => g.pass === true).length ?? 0;
const gradersTotal  = (r as EvalReport).grader_results?.length ?? 0;
```
**Fix (presentation-only, mirrors what `success` already represents):**
```ts
const isPass = (g: GraderResult) =>
  g.pass === true ||
  // tri-state-null with full criteria → AND of criteria
  (g.pass == null && (g.scores?.criteria?.length ?? 0) > 0 &&
                     g.scores!.criteria!.every(c => c.passed)) ||
  // tri-state-null with overall score === max → pass
  (g.pass == null && g.overall_score != null && g.max_score != null &&
                     g.max_score > 0 && g.overall_score === g.max_score);

const gradersPassed = (r.grader_results ?? []).filter(isPass).length;
const gradersTotal  = (r.grader_results ?? []).length;
// fall back to r.success when no graders ran (gpt-5.3-codex case)
const display = gradersTotal > 0 ? `${gradersPassed}/${gradersTotal}` : (r.success ? "—" : "✗");
```
Same `ScoreBadge` colour rule, but now anchored to the same `success` truth as the header card. `0/0` red boxes go away.

> **Note:** this is a stop-gap for the rendering layer. The right fix is for graders to set `pass` correctly server-side — but that's Morpheus's call.

### Q2. `GraderResultRow.tsx:16` — same fallback so review graders stop showing N/A

```ts
// before
const passed = result.pass !== null && result.pass !== undefined ? result.pass : null;
// after
const passed =
  result.pass != null ? result.pass :
  (result.scores?.criteria?.length ?? 0) > 0
    ? result.scores!.criteria!.every(c => c.passed)
    : (result.overall_score != null && result.max_score != null && result.max_score > 0
        ? result.overall_score === result.max_score
        : null);
```
Three review rows on the anchor eval flip from `N/A` (grey) to `PASS` (emerald) without touching backend.

### Q3. `dashboard-page.tsx` — guard the `.toFixed` site

Add `?? 0` (or skip the row) where the runtime hits `undefined.toFixed`. Even a friendlier ErrorBoundary message is better than the React default. I haven't traced the exact line because the bundle is minified — a 5-min un-minified rebuild + repro would pinpoint it. **Asking Switch or Tank to bisect since the dashboard is shipping a stack trace today.**

### Q4. `runs-page.tsx:136-138` — distinguish "errored" from "failed"

Two narrow choices, pick one:
- (a) Compute `effectiveTotal = total - errors` and `rate = passed / effectiveTotal`. Make the bar amber instead of empty-emerald when `errors > 0`.
- (b) Add a tag/strip to the card: `⚠ run errored (3/3)` instead of a `0.0%` rate.

Either keeps the user from confusing "this run was bad" with "this run never finished".

### Q5. `runs-page.tsx` — surface the in-progress / partial run dir

The newest entry on `/runs` was the not-yet-finished `20260423-203921` and rendered as `Unknown … evaluations … 0.0% N/A`. Either filter dirs that have no `summary.json` yet, or render `In progress` with a spinner — the current "Unknown / N/A" in the same shape as a finished run muddies the list.

### Q6. `eval-detail-page.tsx:441-444` — score-card legend

The big `12 / 12` card in the header is `review.overall_score / review.max_score`, which is the *review-grader-only* score, not the unified eval verdict. After Q1+Q2 land, `review.overall_score` will still be present but conceptually subsumed by Points. Two options for the interim:
- (a) Re-label the card "Review Score 12/12" so the grader rows below stop looking like they're contradicting it.
- (b) Compute a `pointsPassed/pointsTotal` synthesis from existing criteria (`output_check.SubChecks` count + `review.scores.criteria` count) and show that under the score number as `(15/15 points across 4 graders)`.

I recommend (a) for now — it's the smaller change and aligns with Morpheus's likely structural fix.

---

## 4. Phase-2 UX proposals (after Points lands)

All wireframes below stay inside the existing visual language: same `mono` font, `rounded-xl border border-white/8 bg-white/[0.03]` card chrome, emerald/red/amber accents from Tailwind tokens already in use.

### 4a. Plugin parent — header-only group, children carry status

**Today** (eval-detail "Available Tools", lines 495-535) — flat row of pill tags:
```
[azure (MCP)] [azure-keyvault-py (skill)] [azure-cosmos-py (skill)] … (40 more)
```

**Proposed** — group by `parent_kind/parent_name`, parent renders as a header without a status pill:

```
Available Tools

  azure-sdk-python  (plugin)            ← header only, no pass/fail badge
  ├─ azure-keyvault-py        (skill)   ✓ used
  ├─ azure-cosmos-py          (skill)
  ├─ azure-identity-py        (skill)
  └─ azure-eventhub-py        (skill)
  …41 children…

  generator-skills  (skills dir)        ← header uses config `name`, not path
  ├─ python-script-quality    (skill)
  └─ python-readme            (skill)

  azure              (mcp)              ← top-level, unchanged from today
```

Markup sketch (reusing existing `inline-flex … rounded-md px-2 py-1 bg-…/10 text-…/40` tag chrome):

```tsx
{Object.entries(toolsByParent).map(([parentKey, group]) => (
  <div key={parentKey} className="mb-3">
    {group.parent && (
      <div className="mb-1.5 flex items-center gap-1.5 text-white/50" style={{ ...mono, fontSize: 11 }}>
        <span>{group.parent.name}</span>
        <span className="text-white/30">({group.parent.kind})</span>
        {/* deliberately NO status pill on parent */}
      </div>
    )}
    <div className={`flex flex-wrap gap-1.5 ${group.parent ? "ml-4 border-l border-white/5 pl-3" : ""}`}>
      {group.children.map(child => <ToolTag key={child.name} tool={child} />)}
    </div>
  </div>
))}
```

The `ml-4 border-l pl-3` indent matches what the run-detail-page table already does for nested filter rows — visual language stays consistent.

**Skill-dir parent** uses `entry.Name` from config (which the validator change in plan.md Phase 1 already routes through) — so the header reads `generator-skills (skills dir)` not `./skills/generator (skills dir)`. No template change needed for that — purely a data-shape ask: the JSON needs `parent_name` populated from config.Name (Morpheus territory).

### 4b. Grader card — header + N indented points

**Today** (`GraderResultRow` collapsed):
```
[N/A] consensus                Review • consensus            12/12
```

**Proposed post-Points** (collapsed):
```
[✓ 12/12 passed] consensus                Review • consensus
```

**Expanded:**
```
[✓ 12/12 passed] consensus                Review • consensus       ▼
  ✓ DefaultAzureCredential Authentication
  ✓ Installing azure-keyvault-secrets and azure-identity packages
  ✓ Creating a SecretClient with vault URL and credential
  ✓ set_secret(), get_secret(), begin_delete_secret(), purge_deleted_secret()
  …
  ✗ paginates list_secrets                                           (if any failed)
```

Markup, reusing GraderResultRow's existing chrome + the criterion pill style already used on line 615-660 for review criteria:

```tsx
const passCount = (result.points ?? []).filter(p => p.pass).length;
const totalCount = (result.points ?? []).length;
const allPass = totalCount > 0 && passCount === totalCount;
const noneRan = totalCount === 0;

// Header badge:
<div className={`flex items-center gap-1.5 rounded-md border px-2 py-1 ${badgeColor}`} style={{ fontSize: 10 }}>
  {badgeIcon}
  <span>
    {noneRan ? (passed === true ? "PASS" : passed === false ? "FAIL" : "N/A")
             : `${passCount}/${totalCount} passed`}
  </span>
</div>

// Expanded body — one row per point, indented:
{(result.points ?? []).map(p => (
  <div key={p.name} className="ml-6 flex items-start gap-2 py-1" style={{ fontSize: 12 }}>
    {p.pass
      ? <CheckCircle2 className="mt-0.5 h-3 w-3 shrink-0 text-emerald-400" />
      : <XCircle className="mt-0.5 h-3 w-3 shrink-0 text-red-400" />}
    <div className="flex-1">
      <div className={p.pass ? "text-white/70" : "text-red-400/80"}>{p.name}</div>
      {p.message && <div className="text-white/40" style={{ fontSize: 11 }}>{p.message}</div>}
    </div>
  </div>
))}
```

Existing `ChevronDown / ChevronRight` toggle stays. The `ml-6` indent matches the criterion list rendering already shipped at eval-detail-page.tsx:615-660.

When `len(points) <= 1`, fall back to today's flat row with a real PASS/FAIL badge — matches what the CLI renderer plans to do per plan.md.

### 4c. Roll-up — one rule, applied everywhere

Add a single helper `isEvalPass(eval)` and use it on **every** surface that paints emerald-vs-red:

```ts
function evalPassFromPoints(r: EvalReport): boolean {
  // Phase 2: every grader emits Points; eval passes iff every point passes
  if (r.grader_results?.length) {
    return r.grader_results.every(g =>
      (g.points?.length ?? 0) === 0
        ? g.pass === true
        : g.points!.every(p => p.pass));
  }
  // No graders ran (e.g. errored generation) → fall back to existing flag
  return r.success === true;
}

function runPassFromEvals(run: RunSummary): { passed: number; failed: number; errored: number } {
  // ... aggregates evalPassFromPoints across run.results
}
```

Apply to:
- `runs-page.tsx` rate calculation
- `run-detail-page.tsx` `gradersPassed/gradersTotal` cell (replace strict `=== true` filter)
- `eval-detail-page.tsx` header score card (replace `r.success` direct read)
- `prompt-detail-page.tsx` "Pass Rate by Config" / "by Tool" aggregations
- `pairwise-page.tsx` and `comparison-page.tsx` per-cell win/loss

Once the rule lives in one place, no surface can drift again.

### 4d. Score-card and counts the user actually wants

For the eval-detail header (replacing the `12/12` review-only card):

```
   ┌────────────────────────────┐
   │     ✓ 15 / 15  points      │   ← AND of every Point across all graders
   │     across 4 graders       │
   └────────────────────────────┘
```

For the run-detail row (replacing `1/4` red):

```
   Score
   ┌──────┐
   │ 4/4  │   ← graders, not points (concise enough for a table cell)
   └──────┘
   tooltip: "15/15 points (4 graders) — output_check, opus-4.6, gpt-4.1, consensus"
```

Two-tier display: granularity at the eval, summary at the run. Both honest, both green when the answer is green.

---

## 5. Coordination with Morpheus

Morpheus has `phase4-pw-final-summary.md` in his folder but no `morpheus-report-architecture-review.md` in `.squad/decisions/inbox/` as of writing. **Will link when it lands.** Boundary respected: this report doesn't propose data-model changes — only flags what the templates need from the wire.

**Data fields the site needs from the JSON that aren't there today** (Morpheus to weigh in on shape):
1. **`parent_kind` + `parent_name` on every `tool_availability` entry.** Today's flat array can't express plugin/skill-dir grouping. Without this the renderer in §4a is impossible.
2. **`Points []GraderPoint{Name, Pass, Message}` on `GraderResult`** — already in plan.md Phase 2; confirming the site needs it on the JSON for §4b.
3. **`pass: boolean` populated correctly on every grader** — including review-type graders. Today the strict `pass === true` filter excludes review graders entirely. If the backend continues returning `pass: null` for review graders, the site will keep needing the fallback in Q1/Q2; cleaner if the engine sets `pass = points.every(p => p.pass)` once Points lands.
4. **`overall_pass: boolean` on `EvalReport`** — currently we lean on `success` (success of *generation*, I think?) and re-derive from grader_results. A canonical eval-level boolean computed by the engine from Points would let every site surface stop reinventing the rule.
5. **`overall_passed` / `pass_rate_excluding_errors` on `RunSummary`** — to fix §1e without each surface re-deriving. Already pre-computed server-side in `summary.go`?

If any of these change shape, ping Trinity — site types live in `site/src/app/data/types.ts`.

---

## 6. Open questions for the user (Ronnie)

1. **For review graders (LLM-as-panel), what's the canonical pass definition?** Today: `criteria.every(c => c.passed)` is implicitly true on the anchor run, but `pass` on the wire is `null`. Should the engine just write `pass = AND(criteria)` for review graders? Or do you want the consensus grader to use a different rule (e.g. quorum)? This determines whether Q1/Q2 are temporary or load-bearing.

2. **Score-card metaphor on the eval page.** Pre-Points the `12/12` is "reviewer agreement". Post-Points it could be "points passed / points total" (granular, true) or "graders passed / graders total" (concise, less true). Which do you want as the dominant number?

3. **In-progress runs on `/runs`.** The list currently shows the partial `20260423-203921` as "Unknown 0.0% N/A". Hide it, or render it as `In progress`?

4. **Errored runs on `/runs`.** Should I split the rate column into `passed | failed | errored` instead of folding errors into the denominator?

5. **Dashboard page priority.** It's been crashing for some unknown number of commits. Pull a fix into the same wave as the grader render fixes, or carve a separate session?

6. **Plugin parent click target.** Should the parent header in §4a be expandable/collapsable (hide the 41 children behind a `▶` toggle), or always-expanded? On a prompt that loads the full `azure-sdk-python` plugin the children take ~6 visual rows of pills.

---

## Appendix — quick reference

- **Anchor JSON**: `/home/rgeraghty/projects/hyoka/reports/20260423-195948/results/key-vault/data-plane/python/crud/key-vault-dp-python-crud/python-pairwise/baseline/claude-sonnet-4.5/report.json`
- **Screenshots**: `/tmp/trinity-site-review/01-runs-list.png` … `11-runs-listing-bad-rows.png`
- **Files cited**:
  - `hyoka/internal/serve/serve.go:142-178` (route map)
  - `site/src/app/components/runs-page.tsx:136-198` (run cards, errored-row issue)
  - `site/src/app/components/run-detail-page.tsx:236-256` (table score cell — primary bug)
  - `site/src/app/components/eval-detail-page.tsx:382-461,495-560,615-660` (header card, tools, graders, criteria render)
  - `site/src/app/components/GraderResultRow.tsx:16,29-39,68` (tri-state badge logic)
  - `site/src/app/components/dashboard-page.tsx` (crash site, line not yet bisected)
  - `site/src/app/data/types.ts` (no `parent_kind`/`parent_name`/`points` types yet)
- **Plan reference**: session-state plan at `/home/rgeraghty/.copilot/session-state/87e98ab8-…/plan.md` Phase 1 + Phase 2.

Read-only review — no templates, CSS, types, or data files were modified. No commits, no pushes.


---

## Latest Decisions (2026-04-24: Generator.json Artifact Arc)

### 2026-04-24T02:50:00Z: Grader Inputs Always Provided, Never Conditional

**Agent:** Ronnie Geraghty (via Copilot)

**What:** Graders (typed AND AI review) always receive BOTH:
1. A copy of the generator's working directory files
2. `generator.json` containing all collected session info (prompt, final response, workspace delta, action summary, termination, etc.)

Neither is conditional on the other being empty. The reviewer prompt unconditionally includes both the Generated Code section AND the Agent Session section. If files are missing, the Generated Code section says "(no files were created)" but the section is still present.

**Why:** User clarification: "we should always pass anything that is available. We should copy over all the files from the generator's working dir and the generator.json that contains all the info we collect from the generator agent's work that is provided to the graders."

**Implications:**
1. GraderInput carries both `WorkspacePath` and `GeneratorArtifactPath`
2. AI review graders always get a temp-dir copy of workspace files (existing PromptReviewGrader behavior continues)
3. Typed graders read workspace directly from `WorkspacePath`
4. Reviewer prompt rendering always includes both sections (never conditional)

**Related:** `coordinator-grader-input-model.md`

---

### 2026-04-24T02:56:00Z: Grader Input Model — generator.json + AI Reviewer Workspace Copy

**Agent:** Ronnie Geraghty (via Copilot)

**What:**
1. ALL graders (typed AND AI review) receive the same input data via the `generator.json` data model — single canonical schema describing the generator session (prompt, final response, workspace delta, action summary, termination, etc.). The artifact is persisted to disk per eval and also passed in-memory via GraderInput. Loose convenience fields on GraderInput remain for back-compat but generator.json is the authoritative source.

2. AI review graders ADDITIONALLY receive a copy of all workspace files in their reviewer working dir (existing PromptReviewGrader.copyDirToTemp behavior). This is the only grader-class-specific input difference. Typed graders read the workspace from `WorkspacePath` directly (no copy).

3. The reviewer prompt always includes both the Generated Code section (rendered from the workspace copy) and the Agent Session section (rendered from generator.json) — never conditional on either being empty. When files are absent, the Generated Code section says so explicitly but is still present.

**Why:** User clarification: "we should be giving all the graders the same information via the data model of the generator.json and we should additionally give AI reviewer graders a copy of all workspace files in their working dir."

**Implementation:** Captured in orchestration logs 2026-04-24T09:15:00Z-neo.md and 2026-04-24T10:30:00Z-trinity.md.

---

### 2026-04-24T02:56Z: Generator.json Artifact Must Be Surfaced on the Site

**Agent:** Ronnie (via Coordinator)

**What:** The `generator.json` artifact (prompt_id, original_prompt, final_response, workspace_delta, actions_summary, timing, terminated_by, error) that the engine emits for graders MUST also be ingested into the report layer and rendered on the site (eval-detail page at minimum). Same data model — graders and site consume one canonical artifact.

**Why:** User directive — single source of truth for everything the generator session produced. Site users want to see the agent's final response, what files changed, and how the session terminated, alongside file contents and grader results.

**Implications:**
1. `report.EvalReport` gets a `GeneratorArtifact *GeneratorArtifact` field (or inline equivalent). Populated alongside `FileContents`.
2. JSON schema bumps to v3 (Neo Phase 1).
3. Site `data/types.ts` mirrors the new shape; eval-detail template renders a "Generator Session" panel showing final_response (collapsible), workspace_delta summary, actions/turns/duration, terminated_by/error.
4. No conditional rendering — if artifact is present (v3+), show it. v2 reports just don't have it.

**Status:** ✅ Implemented (Trinity Phase 2, commit 9f34f072)

---

### 2026-04-24 (Date TBD): Eval Detail Pages Include Generated File Contents

**Agent:** Trinity 🌐

**Status:** ✅ Implemented (Commit c06ca9e2)

**What:** Generated file contents are now captured in `report.EvalReport.FileContents` at report-build time.

**Root Cause of Bug:** The eval detail page served by `hyoka serve` at `/api/runs/{runId}/eval?path=...` returns the `report.json` for that eval, which contained the `GeneratedFiles` array (a list of file paths). However, the file **contents** were never populated in the report JSON.

**Decision:**
1. Added `FileContents map[string]string` field to `EvalReport` (marked `json:"file_contents,omitempty"`).
2. Added `readGeneratedFileContents()` helper in `engine_eval.go`.
   - Called right before `WriteReport()`, reads each file from `ws.Dir` (the workspace directory).
   - **Size cap:** Files exceeding 1MB are capped with a message: `[File too large to display (N bytes) — view on disk at {path}]`.
   - **Binary detection:** Files with binary extensions are skipped.
   - **Error handling:** Files that can't be read show: `[Error reading file: {error}]`.
3. Populated `evalReport.FileContents` before calling `report.WriteReport()`.

**Binary Extensions Detected:**  
`.png`, `.jpg`, `.jpeg`, `.gif`, `.bmp`, `.pdf`, `.zip`, `.tar`, `.gz`, `.7z`, `.exe`, `.dll`, `.so`, `.dylib`, `.bin`, `.dat`, `.db`, `.sqlite`

**Verification:**
- All tests pass: `go test ./hyoka/...`
- Live test: `hyoka serve` on existing report directory shows file contents in eval JSON

**Impact:**
- Site can now display generated file contents on eval detail pages
- Size-safe: 1MB cap per file prevents JSON bloat
- Backward compatible: Existing reports without `file_contents` continue to work

**Reusable Pattern:** Capture report artifacts at report-build time, not serve time. Reading when the report is written (workspace still exists) is more reliable than serve-time reads.

---

### 2026-04-24: Surface generator.json Artifact on Eval-Detail Page

**Agent:** Trinity 🌐

**Status:** ✅ Implemented (Commit 9f34f072)

**What:** Wire `GeneratorArtifact` into the report layer and render it on the eval-detail page as a collapsible "Generator Session" panel.

**Implementation:**

**Go Layer (Phase 1):**
1. Add `GeneratorArtifact *artifact.GeneratorArtifact \`json:"generator_artifact,omitempty"\`` to `report.EvalReport`
2. Type alias pattern (consistent with `WorkspaceDelta`) to avoid import cycles
3. **Write:** After workspace delta computed, before graders run
4. **Read:** After FileContents populated, before WriteReport
5. Helper: `buildGeneratorArtifact()` constructs artifact from eval state

**TypeScript Layer (Phase 2):**
1. Type definitions (`site/src/app/data/types.ts`) mirror Go structs with snake_case
2. Add `generator_artifact?: GeneratorArtifact` to EvalReport

**UI Layer (Phase 3):**
1. Panel placement: ABOVE "Generated Files" panel
2. Collapsed by default with state `showGenSession` (false by default)
3. Panel sections:
   - **Termination badge:** Color-coded status (green=completed, yellow=max_actions, orange=timeout/guardrail, red=error)
   - **Timing:** Duration as "Xm Ys", started timestamp
   - **Actions summary:** 3-column grid (total/tool-calls/reasoning), truncation flag
   - **Workspace delta:** Created/modified/deleted file counts with colored badges
   - **Final response:** Truncated to 500 chars if >500 AND files generated; full text otherwise; copy button
4. Conditional render: Only show panel if `generator_artifact` exists (handles v2 reports gracefully)

**Rationale:**
- Write timing: Artifact must capture complete generation state before graders run
- Read timing: Populate artifact after file contents to keep related functionality grouped
- Collapsed default: Avoids clutter for users who primarily care about grader scores
- Truncation logic: Show full response if <500 chars OR no files generated; truncate if long + files exist
- Backward compat: `omitempty` + conditional render means v2 reports display identically

---

### 2026-04-24: Grader Score Denominator Counts All Grader Points

**Agent:** Trinity 🌐

**Status:** ✅ Implemented (Commit c06ca9e2)

**What:** The total score denominator is the sum of all grader points across all graders, not the count of graders.

**Root Cause of Bug:** In `engine_eval.go:636-637`, `GradersTotal` and `GradersPassed` fields were computed using `len(agg.Results)` and `countPassed(agg.Results)`, which counted the number of graders, not the number of grader points.

**Example:**
- `file_check` grader with 3 points (file1, file2, file3)
- `output_check` grader with 2 points (min_files, require_files)
- Old denominator: 2 (number of graders)
- Correct denominator: 5 (number of grader points)
- Site incorrectly displayed `3/2` instead of `3/5`

**Decision:**
1. Added `countTotalPoints()` helper that sums `len(g.Points)` for each grader in `agg.Results`
   - Graders with `len(Points) == 0` treated as 1 point (backward compatibility)
2. Added `countPassedPoints()` helper that counts passed points across all graders
   - For graders with Points, count points where `Point.Pass == true`
   - For graders with no Points (legacy), use grader's overall `Pass` field (1 if true, 0 if false)
3. Updated `engine_eval.go:636-637` to use point-based helpers

**Verification:**
- All tests pass: `go test ./hyoka/...`
- Table-driven tests cover: multiple points per grader, zero-point graders (legacy), mixed scenarios, empty results

**Impact:**
- Site now displays accurate `X/Y` scores where Y is total grader points
- Backward compatible: Legacy graders treated as 1 point
- Consistent with grader-points architecture

**Reusable Rule:** When aggregating grader scores, always count grader points, not graders. Denominator is `Σ len(g.Points) for g in graders`, with fallback of 1 for legacy graders.

---

### 2026-04-24: Prompt-Frontmatter Criteria Always Get Their Own Review-Grader Bucket

**Agent:** Neo 💊

**Status:** ✅ Implemented (Commit 27c04c71)

**Branch:** ronniegeraghty/dev

**What:** Prompt-frontmatter criteria and criteria-file entries are ALWAYS bucketed separately, regardless of review mode (combined or isolated).

**Root Cause of Bug:** User report: "I'm only seeing one group of ai review graders running but I thought we decided that if we wanted grader points to be graded in the same review session they would have to be grader points on the same grader."

`BuildUnifiedReviewBuckets` previously merged `promptCriteria` (from prompt frontmatter) with matched criteria-file entries into a single `combined` bucket, resulting in one AI review grader. This violated the source-separation principle.

**Decision:**
Each source produces its own `PromptReviewGrader` instance → its own Copilot review session → its own grader display entry.

**Bucket Naming:**
- Prompt-frontmatter criteria → bucket named `"Criteria from prompt file"`
- Criteria-file entries → bucket(s) named based on mode:
  - Combined mode: one `"combined"` bucket for all criteria-file entries
  - Isolated mode: separate buckets per isolated entry/group, plus a `"combined"` bucket for leftovers

**Edge Cases:**
- If `promptCriteria` is empty AND matched criteria-file entries exist → only criteria-file bucket(s)
- If `promptCriteria` is non-empty AND no criteria-file entries → only prompt-frontmatter bucket
- If both are empty → zero buckets (no review graders)

**Rationale:**
**Source of truth determines grader identity.** Prompt-frontmatter criteria come from the `.prompt.md` file's YAML frontmatter. Criteria-file entries come from `criteria/*.yaml` files. These are fundamentally different sources and must be evaluated in separate Copilot sessions:

1. Grader results are distinguishable — users see which bucket contributed which failing criteria
2. Isolation is truly isolated — failure in one doesn't block the other
3. Bucket names are meaningful — "Criteria from prompt file" is explicit

The review mode (combined vs isolated) controls how criteria-FILE entries are bucketed among themselves; it does NOT control whether prompt-frontmatter criteria are separated. Source-separation is a harder rule than mode-separation.

**Implementation:**
- `buckets.go`: Refactored `BuildUnifiedReviewBuckets` to always prepend prompt-frontmatter bucket
- `buckets_test.go`: Updated all combined-mode tests to expect 2 buckets; added edge-case tests
- `engine_reviewbuckets_test.go` and `engine_reviewmode_runtime_test.go`: Updated integration tests

**Verification:**
- ✅ All tests pass: `go test -race ./hyoka/...`
- ✅ Build clean
- 🔄 Live verification: observe TWO separate AI review grader entries in output

**Reusable Principle:** Grader source-of-truth determines grader identity. When multiple criteria sources contribute (prompt frontmatter, criteria files, future: remote criteria), each source must produce a separate grader instance if we want distinguishable pass/fail reporting. The `--review-mode` flag is SECONDARY — it partitions WITHIN a single source's bucket set.

---

### 2026-04-24: Phase-State Guard for Reviewer Event Suppression

**Agent:** Tank ⚛️

**Status:** ✅ Implemented (Commit 6f2e1f03)

**What:** Add a **phase-state guard** at the top of `renderAgentEvent` to suppress activity events from downstream reviewer sessions once the generation phase is complete.

**Root Cause of Bug 3:** The PromptReviewGrader runs a real Copilot SDK session that emits the same generation events (EventReasoning, EventToolStart, EventToolComplete) through the shared progress event channel. After EventSessionDetails flipped Agent Attempt → Completed, subsequent reviewer activity events were still landing in `renderAgentEvent` and creating duplicate "Agent Attempt: ✅ Completed" rows.

**Decision:**
```go
// Agent Attempt is already finalized — generation phase is over. Ignore
// activity events from downstream sessions (reviewer Copilot sessions
// emit the same EventReasoning/EventToolStart/etc. through the shared
// event channel, but they belong to grader rows, not the agent tail).
if r.cur != nil && (r.cur.agentState == agentStateCompleted || r.cur.agentState == agentStateGuardrail) {
    return
}
```

**Rationale:**
- The agent tail belongs to GENERATION phase only
- Once `agentState` is Completed or Guardrail, generation phase is closed
- Downstream sessions (reviewers) emit same event types but belong to grader rows, NOT agent tail
- Renderer must filter events by phase-ownership, not just event type

**Implementation:**
- `display_interactive.go` (+5 lines): phase-state guard in `renderAgentEvent`
- `display_interactive_test.go` (+56 lines): regression test `TestInteractive_ReviewerEventsAfterCompletionIgnored`

**Test Coverage:**
`TestInteractive_ReviewerEventsAfterCompletionIgnored` drives full sequence: generation → EventSessionDetails (completes agent) → typed grader → AI review grader start → EventReasoning/EventToolStart/EventToolComplete from reviewer → grader complete → passed. Asserts exactly ONE "Agent Attempt:" row.

**Verification:**
- `go test -race ./hyoka/internal/progress/...` — all tests pass
- Live eval (key-vault-dp-python-crud / python-pairwise) — zero duplicate rows

**Alternatives Considered:**
1. **Filter at event emission:** Rejected — couples eval package to display semantics
2. **Add "review phase" state:** Rejected — over-engineered; agent tail is generation-only
3. **Change event types:** Rejected — would break existing event consumers

**Lessons:**
- Phase-state guards are critical in event-driven terminal renderers
- When multiple concurrent processes emit through shared channel, filter by phase, not event type
- Test for phase isolation — events from one phase must not bleed into another's rendering

---


### Decision: Grader Output Format — Per-Grader Display with Point Breakdown (2026-04-24T03:26Z)

**Agent:** Coordinator (user direction via Ronnie)  
**Branch:** ronniegeraghty/dev  
**Implementation:** Tank (commit `4adc9288`)  

## Context

Current progress display lumped all AI-review grader points under a single `- ai_review` line, making it impossible to see which specific grader passed or failed which criteria point.

## Decision

**Grader output now displays per-grader with individual point breakdown.**

Format:
```
- <grader-name> (<grader-type>): <score>
  - <point.name or first ~60 chars of point.prompt>: <pass|fail>
  - <point.name or first ~60 chars of point.prompt>: <pass|fail>
```

Examples:
- `- DefaultAzureCredential Authentication (prompt): 3/4`
  - `- Uses DefaultAzureCredential from azure.identity: pass`
  - `- Wraps client construction in try/except: fail`
- `- Criteria from prompt file (prompt): 5/5`
- `- file_exists/keyvault_crud.py (file): pass`
- `- behavior/runs_without_error (behavior): pass`

## Rules

1. **Grader name** comes from the grader entry's `name:` field. For prompt-frontmatter bucket, hard-coded to `Criteria from prompt file`.
2. **Grader type** rendered in parens after name: `prompt` (AI review), `file`, `behavior`, `program`, `output_check`.
3. **Points** indented as sub-bullets. Use point's `name` if present; else first ~60 chars of `prompt` with ellipsis.
4. **Single-point graders** render on one line — no sub-bullets.
5. **Score format:** `pass`/`fail` for binary; `X/Y` for multi-point counts.

## Scope

- Interactive CLI display
- JSON report (`graders[].name`, `graders[].type`)
- Site eval-detail page (Trinity's frontend panel)

## Implementation

Tank refactored engine to create one PromptReviewGrader per ReviewBucket instead of one "ai_review" grader for all buckets. Result:
- `engine_eval.go`: Per-bucket grader iteration
- `display_interactive.go`: `displayKind()` helper for type mapping
- `display_interactive_points_test.go`: Updated assertions for multi-grader rendering

## Rationale

Single `- ai_review` row with nested points provides no signal about which grader (e.g., "DefaultAzureCredential") passed or failed. Per-grader rows with type labels and point-level breakdown enables users to:
- Identify which criteria bucket failed
- See exactly which individual points passed/failed
- Correlate failures to grader source (prompt, file, behavior, etc.)

## Verification

✅ All tests pass  
✅ Live evals display per-bucket graders with point breakdown  
✅ Site renders grader names and types correctly

---


---

### Decision: Prompt Grader `checks:` Redesign — Scope (2026-04-24T05:37:56Z)

**Agent:** Morpheus 🕶️  
**For:** Neo (engine/grader logic) and Tank (CLI/report rendering) collaboration  
**Status:** Scoped, ready for implementation  
**Branch:** `ronniegeraghty/prompt-grader-checks`

## Problem

Prompt-file-driven prompt graders look fine in the run output: each top-level bullet under `## Evaluation Criteria` becomes its own pass/fail row nested under the parent grader (handled by `prompt.ParseEvaluationCriteria` → `ParsedCriteria` → bucket rendering → LLM returns one judgement per criterion → `result.Points`).

YAML criteria-file prompt graders do **not** look fine. The current schema has a single `prompt:` field per `UnifiedGraderEntry`. Authors who want multiple checks today smuggle them inside that string ("Check the following criteria: 1. … 2. …"). The bucket renderer (`FormatUnifiedPromptEntries` in `internal/criteria/buckets.go:119`) emits that whole blob as a **single** numbered criterion to the review LLM, so the LLM returns **one** pass/fail for the whole blob and we lose per-check granularity. The terminal renderer already supports multi-Point graders (`internal/progress/display_interactive.go:1003-1062`) — there's just nothing for it to render.

The user wants YAML prompt graders to declare checks as first-class structured items so we get the same nested per-check output the prompt-file path already produces.

## Solution

**Add `Checks []string` field to `UnifiedGraderEntry`; migrate two criteria files to use it; update bucket rendering and grader logic; verify end-to-end through report layer.**

### Schema change

Add to `UnifiedGraderEntry` in `internal/criteria/config.go`:

```go
Checks []string `yaml:"checks,omitempty" json:"checks,omitempty"`
```

YAML shape:

```yaml
graders:
  - name: Markdown Structure
    type: prompt
    weight: 1.0
    prompt: Check the following criteria:    # optional preamble
    checks:
      - File hello.md exists and contains a level-1 heading.
      - File contains exactly three bullet list items.
```

### Validation (extends `validateEntry`)

For `type: prompt`:
- One of `prompt:` or `checks:` MUST be non-empty. Both may be set.
- If `checks:` is set, each entry must be a non-empty trimmed string. Empty/whitespace-only entries fail with index.
- If `checks:` is unset/empty, backward compat: single `prompt:` becomes one criterion → one Point.
- For non-prompt types: `checks` must be empty (mirrors existing `prompt`-must-be-empty rule).

### Bucket rendering (`internal/criteria/buckets.go`)

`FormatUnifiedPromptEntries` now honors `Checks`:

- **With checks:** Parent line `N. **Name**`, optional preamble from `Prompt` field (judge-only, rendered indented), then nested checks as numbered items:
  ```
  3. **DefaultAzureCredential Authentication**
     Check the following criteria:
       1. Uses DefaultAzureCredential from azure.identity.
       2. Wraps client construction in try/except.
  ```
- **Without checks:** Legacy single-line behavior preserved byte-for-byte: `N. **Name** — Prompt`.

### Grader logic (`internal/criteria/graders/prompt_review_grader.go`)

- Count expected criteria: parse the rendered bucket text to find leaf-numbered items (prefers indented when present); compare against returned criteria count.
- Emit `slog.Debug` when count differs (non-fatal; helps track outliers).

### YAML Migrations

Hard-migrate two criterion files per §2:

- **`criteria/language/python.yaml`** — `DefaultAzureCredential Authentication`: split embedded "1. … 2. …" into preamble `"Check the following criteria:"` + 2 `checks:` items (auth-credential, async/await).
- **`criteria/language/test.yaml`** — `Markdown Structure`: matches the user's worked example exactly (preamble + 2 checks for hello.md heading and bullet count).

Other files (`java.yaml`, `rust.yaml`, testdata) unchanged (no embedded numbering; single-prompt path covers them).

### Tests

- `hyoka/internal/criteria/config_test.go`: Replaced `TestValidateEntry_PromptMissingPrompt`; added 6-case `TestValidateEntry_PromptChecks` table covering checks-only, preamble+checks, legacy prompt-only, empty string in checks (with index), whitespace-only check, checks on wrong type.
- `hyoka/internal/criteria/buckets_test.go`: Added `TestFormatUnifiedPromptEntries_Shapes` (5 cases): empty, legacy single-prompt, checks+preamble, checks-no-preamble, mixed legacy+checks. All assert exact rendered text.

### Verification

```
go build ./...
go test ./... -timeout 5m
go test -race ./hyoka/internal/criteria/...
hyoka run --prompt-id key-vault-dp-python-crud --config baseline/claude-opus-4.6
```

Expected: Per-check Points propagate end-to-end. Debug log fires on criterion-count mismatch.

---

## Ownership

- **Neo (§1, §4, §5):** Schema change, validation, `FormatUnifiedPromptEntries` renderer, grader logic.
- **Tank (§6, §7):** Display badge format update (`✅ Pass (X/Y)` / `❌ Fail (X/Y)`), verify Points end-to-end in report layer.

File sets are disjoint; can land independently.

---

### Decision: Neo — Prompt Grader `checks:` Implementation (2026-04-24T05:37:56Z)

**Agent:** Neo 💊  
**For:** Coordinator / Scribe (merge into decisions.md)  
**Status:** Landed on branch `ronniegeraghty/prompt-grader-checks`  
**Scope ref:** Morpheus's scope §1, §4, §5  
**Commit:** `2949f578`

## What landed

Implemented Morpheus's YAML prompt grader `checks:` field per scope §1, §4, §5. YAML prompt graders can now declare structured per-check items that produce one `GraderPoint` (and one nested Pass/Fail row in renderers) per check, instead of smuggling numbered lists inside `prompt:` and getting one combined verdict.

## Files touched

**Engine / schema:**
- `hyoka/internal/criteria/config.go` — added `Checks []string` to `UnifiedGraderEntry`; `validateEntry` enforces:
  - `type=prompt`: at least one of `prompt` / `checks` non-empty; each `checks` entry must trim non-empty (errors include the bad index); `details` still forbidden.
  - non-prompt types: `checks` must be empty (mirrors the existing `prompt`-must-be-empty rule).
- `hyoka/internal/criteria/buckets.go` — `FormatUnifiedPromptEntries` honors `Checks`. With checks: parent line `N. **Name**`, optional preamble (entry's `Prompt` field, judge-only), then nested `   1. …   2. …` checks. Without checks: legacy single-line `N. **Name** — Prompt` shape preserved byte-for-byte.
- `hyoka/internal/criteria/graders/prompt_review_grader.go` — added `expectedCriteriaCount` (regex-based leaf-numbered-item count, prefers indented when present) and `logCriteriaCountMismatch` helper. Both `gradePanel` and `gradeSingle` now emit a `slog.Debug` when `len(returnedCriteria) != expected`. No fail; no structural changes.

**YAML migrations (hard-migrated per §2):**
- `criteria/language/python.yaml` — `DefaultAzureCredential Authentication`: split smuggled `1. … 2. …` into preamble `"Check the following criteria:"` + 2 `checks:` items (auth-credential, async/await).
- `criteria/language/test.yaml` — `Markdown Structure`: matches the user's worked example exactly (preamble + 2 checks for hello.md heading and bullet count).
- `criteria/language/java.yaml`, `criteria/language/rust.yaml`, and `hyoka/internal/criteria/testdata/**` — left alone (no embedded numbering; single-`prompt` path covers them).

**Tests:**
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

---

### Decision: Tank — Grader Points Badge Format Alignment & End-to-End Verification (2026-04-24T05:37:56Z)

**Agent:** Tank 📡  
**Scope ref:** Morpheus's scope §6, §7  
**Branch:** `ronniegeraghty/prompt-grader-checks` (shared with Neo)  
**Status:** Implemented, tested, end-to-end verified  
**Commit:** `a47cb97d`

## What changed

1. **Badge format** in `hyoka/internal/progress/display_interactive.go` (`renderGraderWithPoints`):
   - Old: `✅ 2/2 passed` / `❌ 1/3 passed`
   - New: `✅ Pass (2/2)` / `❌ Fail (1/3)` (matches user spec).
   - One-line behavior change; godoc updated.

2. **Soft truncation** of per-Point names in the interactive renderer:
   - `name = truncateToWidth(p.Name, 50)` — long check strings get a `…` ellipsis at ~50 cols.
   - Reuses existing `truncateToWidth` (ANSI-aware, wide-char-aware) — no new helpers, no width plumbing needed.
   - Full text remains in `report.GraderResult.Points[].Name`.
   - Rationale for fixed 50: terminal width isn't piped down to this renderer today, and 50 cols looks clean inside the `    - <name>: ✅ Pass` indent without breaking on standard 100-col terminals. Easy to swap for terminal-width-derived later if Ronnie wants.

3. **Test updates** in `display_interactive_points_test.go`:
   - Three assertions migrated from `"X/Y passed"` to `"Pass (X/Y)"` / `"Fail (X/Y)"`.
   - Single-Point flat-row negative check updated to the new format so it still catches regressions.

## Verifications

- `go build ./...` — clean.
- `go test -race ./internal/progress/... ./internal/report/... -timeout 3m` — green (progress 1.07s, report 29.5s).
- **End-to-end smoke** against Neo's migrated `criteria/language/test.yaml` using `prompts/test/hello-markdown.prompt.md` + `test/haiku` config. Inspected `reports/20260424-052849/.../report.json`:
  - `Markdown Structure` (YAML grader, `type: prompt`, `checks: [2 items]`) → `points: 2` ✅
  - `Criteria from prompt file` (prompt-file path, 3 bullets under `## Evaluation Criteria`) → `points: 3` ✅
  - Both flow through `convertGraderResults` (`hyoka/internal/eval/engine_eval.go:1199-1204`) into `report.GraderResult.Points`. No schema bump needed — Phase 2 v3 already covered this.

## Renderers verified (no changes required)

- `display_interactive.go` — updated (this PR).
- `display_ci.go` — only tracks aggregate pass/total per eval; doesn't render Points individually. Correct as-is.
- No separate `display_json*.go` / `display_quiet*.go` exist; report JSON serialization is owned by `internal/report` and was verified above.

## Out of scope (per Morpheus's scope §6, §7)

- Site / TS template Points rendering — Trinity's domain (Phase 6).
- Schema bump — not needed; v3 already carries Points.
- Terminal-width-aware truncation — fixed 50-col cap is sufficient for now.

## Co-author

Co-authored with Neo on the same branch (`ronniegeraghty/prompt-grader-checks`).
Tank's commits are file-disjoint from Neo's: `hyoka/internal/progress/display_interactive*.go` only.

---

### Decision: Move Taxonomy from Hardcoded to Discovered (2026-04-24T05:37:56Z)

**Agent:** Morpheus 🕶️  
**Issue:** [#635](https://github.com/ronniegeraghty/hyoka/issues/635)  
**Status:** Proposed (issue filed; awaiting implementation)

## Direction

Hyoka is moving away from package-level hardcoded `Valid*` slices in `internal/validate/` toward **runtime discovery** of taxonomy values (`service`, `plane`, `language`, `category`, `difficulty`) from the corpus itself (`prompts/`, `criteria/`, `configs/`).

A new `internal/taxonomy` package will own the walker, the union-of-observed-values set, and a Levenshtein-based "did you mean" suggester. An optional `taxonomy.yaml` at the project root provides forward-declared values (corpus ∪ declared = full allowlist).

## Why future agents should care

- **Don't add new values to `internal/validate/validate.go` or `schema.go`.** Those lists are slated for deletion. If you need to add a value, add a prompt/criteria/config that uses it (or wait for `taxonomy.yaml` to land and declare it there).
- **`isValidLanguage`/`isValidService`/etc. in `schema.go:176-213` will be deleted** along with `isTestValue`. Don't extend them.
- **`cmd/new_prompt.go:14-32` is the only in-repo consumer of the exported `Valid*` slices.** Any new consumer will be deleted with the slices.
- **Site (Trinity) impact is zero for v1** — `internal/serve/dashboard.go` already accepts opaque filter strings. Follow-up: Trinity may want to consume `hyoka taxonomy --json` for empty-state UX, but that's out of scope for #635.

## Suggested owner

Neo (engine/validate work); Morpheus reviews; Oracle docs the `taxonomy.yaml` format once shape is final.

## Out of scope (non-goals for #635)

- Renaming or deprecating existing taxonomy values
- Changing prompt frontmatter schema
- Removing the ID-prefix naming convention check (`schema.go:166`)
- Server-side enforcement on the serve API

---
