# Item D — Pre-session hard-fail collects ALL tool failures

**Author:** Neo 💊
**Status:** Shipped (pending Item A merge resolution from Tank)
**Spec:** `morpheus-tool-load-consolidation.md` § Item D + § 2 gap #5

## What changed

Replaced single-failure short-circuit at the eval boundary with an aggregated
multi-failure summary. Operators with N broken tools now see all N in one
shot rather than fixing-running-fixing-running.

## Surface area

### `hyoka/internal/config/tool/validate.go`
- **Added** `(*ToolLoadReport).AllErrors() []*ToolLoadError` — walks `Items`,
  collects every `ToolStatusFailed` row in report order. Returns nil for clean reports.
- **Added** `(*ToolLoadReport).JoinedError() error` — returns a `*joinedToolLoadError`
  wrapping every per-tool failure, or nil. Implements `Unwrap() []error` so
  `errors.As(err, &target)` traverses to any `*ToolLoadError` leaf.
- **Added** `SummarizeToolLoadErrors(errs []*ToolLoadError) string` — exported
  formatter (capital S so Item E can grab it). Empty slice → "".
- **Removed** `(*ToolLoadReport).FirstError()` — no non-test callers, deleted.
- **Changed** `ValidateAndExpand` to return `report.JoinedError()` instead of
  `report.FirstError()`. Signature unchanged: `(*ToolLoadReport, error)`.
- `validateEntries` stays sequential (Morpheus's call — ordered output > parallel speedup).

### `hyoka/internal/eval/copilot.go` (line ~185)
- `EvalResult.Error` now reads `"tool_load_failure:\n" + toolErr.Error()` so the
  multi-line summary gets its own block under the category prefix.
- `EvalResult.ErrorDetails` is the bare summary (no prefix).
- `ErrorCategory` unchanged: `"tool_load_failure"`.

### `hyoka/cmd/run.go` (line ~407)
- Reviewer-tool wrap message now uses `\n%w` so the per-config prefix sits on
  its own line above the multi-line summary.

### Tests
- Deleted `TestToolLoadReport_FirstError`; replaced with `TestToolLoadReport_AllErrors`
  covering AllErrors + JoinedError (clean and dirty paths).
- Added `TestSummarizeToolLoadErrors` — table-driven, covers empty/single/multiple.
- Added `TestValidateAndExpand_MultipleFailures_AggregatesAll` — full integration:
  3 deliberately broken tools (bad skill_dir, missing plugin, blank MCP command),
  asserts every name + the aggregate header appear in the joined message AND
  that `errors.As` still locates the wrapped `*ToolLoadError` leaf.
- Updated 4 existing tests in `validate_test.go` and 1 in `plugin_migration_test.go`
  from `err.(*ToolLoadError)` → `errors.As(err, &target)` (now wrapped, not bare).

## The format string (Item E: match this exactly)

```
N tool(s) failed to load:
  • {kind} "{name}": {reason}
  • {kind} "{name}": {reason}
```

- Header `%d tool(s) failed to load:` (always plural form, even for N=1 — keeps the parser simple).
- Bullets are `\n  • ` (two spaces, U+2022 bullet, one space).
- Per-line content is `(*ToolLoadError).Error()` which is `{kind} "{name}": {reason}` (printf `%q` on name).
- Empty/nil slice → empty string (NOT a header with zero bullets).

Item E should call `tool.SummarizeToolLoadErrors(...)` directly rather than
re-implementing the format. Don't drift the bullet spacing or quote style.

## Decisions for Ronnie

1. **Custom `joinedToolLoadError` over plain `errors.Join`.** Stdlib `errors.Join`
   formats one-error-per-line with no header. Operators benefit from the
   "N tool(s) failed to load:" lead-in, especially in CLI output where the
   error is often surrounded by other log lines. The custom type still implements
   `Unwrap() []error` so `errors.Is`/`errors.As` work normally.
2. **`%q` for names** so `mcp "azure mcp with spaces"` parses unambiguously.
   The illustrative spec example used single-quotes; chose Go-idiomatic double
   quotes via `fmt.Sprintf("%s %q: %s", ...)`.
3. **Header always plural.** "1 tool(s)" looks slightly off but avoids a
   conditional and matches the SummarizeToolLoadErrors contract Item E will rely on.

## Coordination notes

- **Item A (Tank):** No conflicts with my edits to `validate.go` — Tank touches
  `FetchRemote` signatures; I only touched the `ToolLoadReport` methods + the
  one `FirstError → JoinedError` swap inside `ValidateAndExpand`. Already
  confirmed no overlapping lines.
- **Item E (Neo, depends on this):** Use `tool.SummarizeToolLoadErrors` for
  post-session verifier rendering. Same format, same quoting.
- **Tank's WIP** is currently breaking `TestValidateAndExpand_RemoteContainerPlugin_FansOutChildren`
  and `TestResolveInstalled_*` — confirmed by selective stash; not my code.
