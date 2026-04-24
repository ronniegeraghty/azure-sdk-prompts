# Squad Decisions

## Active Decisions

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

