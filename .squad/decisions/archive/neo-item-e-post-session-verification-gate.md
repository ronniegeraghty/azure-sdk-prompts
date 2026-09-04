# Item E — Post-session tool verification gate re-enabled

**Author:** Neo 🧠
**Status:** Shipped
**Spec:** `morpheus-tool-load-consolidation.md` § Item E + § 2 gap #6
**Depends on:** Item D (`tool.SummarizeToolLoadErrors` — error format contract)

## What changed

Closed the false-positive eval window. Pre-Item-E, every eval was running graders against generated code regardless of whether the configured skills/MCP servers actually loaded — the original `waitForToolVerification` call was disabled because it ran *before* `session.SendAndWait` and the SDK only emits load events *after* the first round-trip. The gate now runs at the right point in the lifecycle and short-circuits any eval whose tools didn't load.

## Surface area

### `hyoka/internal/eval/tool_verification.go`
- **Added** `postSessionToolVerification(ctx, *toolVerifier, timeout) string` — the engine-facing entry point. Returns "" on clean verification (or no-tools-configured), or a `tool.SummarizeToolLoadErrors`-formatted multi-line summary when any configured tool failed to load.
- **Added** `expectedAsTimeoutFailures(*toolVerifier, timeout) []progress.ToolStatus` — synthesizes a Failed entry per configured tool with reason `"SDK did not confirm tool load within <timeout>"`. Sort order matches `emitIfReady` so renderers see identical layouts across happy/timeout paths.
- **Imported** `internal/config/tool` so the helper can build `[]*tool.ToolLoadError` and call `tool.SummarizeToolLoadErrors` directly — same wording as Item D's pre-session path.

### `hyoka/internal/eval/copilot.go`
1. **Removed** the `// NOTE: Tool validation gate is DISABLED …` block (12 lines including the TODO citing #347) immediately after `client.CreateSession`.
2. **Inserted** the gate call after `session.SendAndWait` returns successfully and after the captured-events copy under `mu.Lock`, but before `listFiles(workDir)` and the success-return — listing files we will not grade is wasted work.
3. On gate failure, returns an `EvalResult` with:
   - `Success: false`
   - `Error: "tool_load_failure:\n" + summary`  — same prefix as line 192 (pre-session)
   - `ErrorDetails: summary`  — bare summary
   - `ErrorCategory: "tool_load_failure"`  — consumed by `engine_eval.go:159` to set `evalReport.ErrorCategory`
   - All captured `SessionEvents`, `ActionTimeline`, `ToolCalls`, `FinalResponse`, `ToolReport`, `CleanupFn`  — diagnostics survive
   - error: `fmt.Errorf("post-session tool verification: %s", summary)`

### `hyoka/internal/eval/tool_verification_test.go`
- 6 new tests under the `PostSessionVerification_` prefix: `NothingConfigured`, `AllLoaded`, `FailedSkill`, `FailedMCP`, `MixedFailures`, `TimeoutMarksAllFailed`, `NilVerifier`.
- All `-race` clean. Test file imports gained `context`, `strings`, `time`.

## Decisions for Ronnie

### 1. Where the gate lives — *after* the captured-events copy

Placed at `copilot.go:744` (post-`mu.Unlock`, pre-`listFiles`). Three reasons:

- **Diagnostics survive abort.** The failure `EvalResult` includes the full `SessionEvents` + `ActionTimeline` + `ToolCalls` slice so operators investigating a tool-load failure still see what the model tried to do during the (broken) session.
- **No wasted I/O.** `listFiles` walks the workspace; skipping it on the failure branch is free.
- **Single ordering invariant.** Keeping the gate strictly between SendAndWait completion and "Collect results" means a future refactor that splits generation from grading sees one obvious "this is where verification happens" line.

### 2. Plugins **not** verified post-session

The SDK exposes `SessionExtensionsLoaded` (with an `Extension.Status` field including `failed`), but **no `SessionPluginsLoaded` event** exists in `copilot-sdk/go@v0.2.0/generated_session_events.go`. "Extensions" are conceptually adjacent to plugins but no other code in hyoka maps Extensions onto our plugin model — wiring that translation is a separate, scope-creep task.

**This is acceptable because** Item B (Neo, shipped) registered `pluginFetcher`, so plugin remote-fetch failures are caught pre-session via Item D's `ValidateAndExpand` aggregation. The post-session gap is now: "a plugin that resolved on disk pre-session but failed to load mid-session" — a narrow window that doesn't currently produce false positives in any real eval.

**Follow-up:** when we move to SDK v0.3+ or wire `SessionExtensionsLoaded` → `progress.ToolKindPlugin`, plumb plugins through the verifier's third channel (`expectedPlugins map[string]bool`) the same way skills/MCP are wired today. New ticket: "Plumb SDK Extension events through toolVerifier."

### 3. Timeout semantics — hard-fail, no partial success

30-second deadline. On timeout, **every** configured tool is marked Failed (not just the ones we hadn't heard from), with reason `"SDK did not confirm tool load within 30s"`. Two reasons:

- The whole point of Item E is **no false positives**. A partial-success path ("MCP loaded but skills timed out, proceed anyway") is exactly the failure mode we're closing.
- The timeout path's summary names every tool, so operators get the same shape of error message regardless of which timing edge case fired. The `Reason` field disambiguates ("SDK did not confirm…" vs. a real per-tool reason).

The 30s value comes from Morpheus's spec § 5 open-question 5 ("default 30s, configurable via `--tool-verify-timeout`"). I did **not** add the CLI flag — out of scope for Item E. If cold-start MCP servers prove flaky, Tank can add the flag in a follow-up; the constant is at the call site for easy override.

### 4. Opportunistic flush before waiting

Found a second latent bug while writing tests: in `copilot.go`, `verifier.emitIfReady()` is only invoked inside `if e.progressFn != nil { … }`. An eval running with no progress display (e.g., `--progress off`, unit tests, or any non-interactive consumer) would have its `verifier.readyChan` *never close* even after both SDK events fired — meaning `waitForToolVerification` would always time out for those callers.

`postSessionToolVerification` now calls `v.emitIfReady()` once at the top before falling through to `waitForToolVerification`. If both kinds already arrived, it returns the tools immediately and closes `readyChan` as a side effect. Cost: one extra map walk per eval. Benefit: gate works correctly regardless of progress-display configuration.

This **does not change** the `progressFn != nil` guard in `copilot.go`'s OnEvent handler — the bulk `EventToolsVerified` progress event is still display-only and still gated. We just no longer rely on that path to satisfy the readiness channel.

### 5. Issue #347 — can be closed

The TODO at the deleted comment block cited issue #347. Per-symbol audit:

- The disabled gate (lines 643-655) → **removed**.
- The "real-time turn / file / action limits" referenced via inline `(#347)` comments at copilot.go:213, 302, 339 → **already implemented** (turnLimitHit, fileLimitHit, actionLimitHit) and unchanged by this work.
- The "verify all expected MCP servers loaded" comment at copilot.go:432 → **superseded** by the post-session gate (the inline `lg.Warn` it produces is now diagnostic-only since the gate also catches missing MCPs).

**Recommendation:** close #347 with a link to this decision file + the Item D + Item E PRs. Open a new narrow ticket for the SDK Extension/plugin wiring (decision 2 above) if/when we want post-session plugin verification.

## Verification

- `go build ./...` — clean.
- `go test -race ./hyoka/internal/eval/... ./hyoka/internal/config/tool/... -timeout 120s` — passes.
- `go test -race ./...` — pre-existing failures in `report`, `rerender`, `serve` due to Tank's in-flight `FetchRemote` signature change (Item A WIP); **unrelated to Item E**. Confirmed by stashing my changes and re-running — same failures.

## Coordination notes

- **Switch (Item G tests):** the production-side coverage I added is sufficient for Item E acceptance. If Switch wants an end-to-end test that actually constructs a Copilot session and broken MCP entry, that requires SDK mocking infrastructure that doesn't exist in the eval package today — that's a Switch-scoped scaffold, not blocking Item E.
- **Tank (Item A/C/F):** no overlap. I touched `copilot.go` lines around 645/740 and `tool_verification.go`; Tank's WIP is in `internal/config/tool/{fetcher,validate,resolve}.go`. The pre-existing build failures in `report`/`rerender`/`serve` are Tank's `FetchRemote` signature change, will resolve when Tank's branch lands.
- **Trinity (site/report):** `EvalResult.ErrorCategory == "tool_load_failure"` continues to render the same way it does for the pre-session path — no site changes required.
