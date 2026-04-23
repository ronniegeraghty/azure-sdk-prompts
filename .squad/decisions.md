# Squad Decisions

## Active Decisions

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

