# Tool Load Hard-Fail — Implementation Decisions

**Author:** Neo
**Date:** 2026-04-23
**Status:** Implemented (branch `ronniegeraghty/dev`, commits acd36cde..0131f35d)
**Relates to:** Morpheus plugin-loader-plan WU-1, WU-2. Tank WU-3 depends on this.

## Scope

Implemented Morpheus's WU-1 (static pre-session tool validation) and WU-2
(reviewer skill resolution parity). Four commits, all on `ronniegeraghty/dev`.

## Strict vs Lenient Contract

The plan asked for an opt-in strict flag on `resolveLocal` / `resolveSkillDir`
so legacy lenient callers could stay silent. I took a different shape that
I believe achieves the same goal with less disruption:

**Strict path (new, hard-fail):**
- `tool.ValidateAndExpand(ctx, ValidationInput) (*ToolLoadReport, error)` —
  the single entry point for pre-session validation. Returns a non-nil error
  on any failed item; callers return an EvalResult with
  `ErrorCategory="tool_load_failure"` and abort before CreateSession.

**Lenient path (unchanged):**
- `tool.ResolveSkills` / `ResolveSkillsWithReporter` — still return
  `(nil, nil)` on missing paths with `slog.Warn`. One caller remains:
  `engine_eval.go:268` (the post-run reporting path that populates
  `env.SkillDirectories` even when the main eval failed — losing this would
  produce confusing "zero skills" reports in legitimate error cases).

**Not changed (yet):**
- `resolveLocal`, `resolveSkillDir`, `loadPlugin` internal helpers.
  ValidateAndExpand does its own `os.Stat`-based checks rather than calling
  through to them. Rationale: reimplementing in the validator kept the diff
  small AND avoided a subtle API break for tests in
  `internal/config/tool/resolve_test.go` that assert the lenient
  `(nil, nil)` return path.

Callers that want strict semantics should use `ValidateAndExpand`, not
flag the old helpers.

## Where ValidateAndExpand Runs

Two sites:

1. **`copilot.go` Run()**: called after `NewIsolatedConfigDir` and before
   `buildSessionConfigForEval`. Validates generator tools + plugins.
   `ReviewerTools` left empty here because reviewer validation happens
   per-config in cmd/run.go (point 2 below) and validating both would
   double-emit progress events.

2. **`cmd/run.go` reviewerFactory**: each call to the closure validates
   only `cfg.Reviewer.Tools` for that specific config. This is what kills
   cross-config leakage — there is no pooled slice to leak from.

## Report Role Attribution

Every `ToolLoadItem` carries a `Role` field: `"generator"`, `"reviewer"`,
or `"plugin"`. The report's `GeneratorSkillDirs()` / `ReviewerSkillDirs()`
helpers filter by role so cmd/run.go can extract reviewer skills without
seeing generator skills even when a single future call site validates both.

## Event-Shape Contract for Tank (WU-3)

Committed standalone as `acd36cde` so Tank can rebase his branch from that
SHA without pulling in WU-1/WU-2. Adds two fields to `progress.ProgressEvent`
and `progress.ToolStatus`:

```go
ParentName string // plugin name, or skills-dir path, or "" (no container)
ParentKind string // ToolParentKindPlugin, ToolParentKindSkillDir, or ""
```

`ValidateAndExpand` emits these on plugin children (ParentKind=plugin,
ParentName=plugin-name) and skill_dir children (ParentKind=skill_dir,
ParentName=skills-dir-path). Glob expansions are treated as skill_dir-style
parents (ParentName=glob-pattern). Top-level entries have empty parent.

## What I Didn't Do

- **Didn't re-enable the SDK post-SendAndWait verification gate** (still
  disabled from `4b593d3b`). Static validation is Phase 1; Phase 2 is the
  SDK-event gate. Out of scope for this PR per plan section 5.
- **Didn't add `optional: true` escape hatch** on tool entries. Plan
  marked this Phase 2.
- **Didn't dedupe the 13× plugin-registry re-walk** on each invocation.
  Cosmetic.
- **Didn't rename the existing `resolve*` helpers to make them "internal".**
  Their tests are still green; leave them for later.

## What Switch Needs (WU-4)

Table-driven tests for `ValidateAndExpand`:
- Each failure mode F1–F7, F9 with exact `Reason` string assertions.
- Happy-path per-plugin / per-skill_dir expansion count.
- Integration test: missing plugin via `StubRunner`-adjacent harness →
  eval report shows `error_category: "tool_load_failure"`.

Tests for the reviewer factory closure (cmd/run.go): exercise two matched
configs, pass only one via --config, assert only that config's
reviewer.Tools appear in the factory's ValidateAndExpand input.

## Open Questions

1. Should we bump `ExpandPlugins` at config-load time to also be strict?
   Currently it stays lenient so `hyoka list` succeeds on a contributor's
   machine that's missing plugin installs. Pro-strict: single source of
   truth. Pro-lenient: `list` ≠ `run`; let `list` tolerate, let `run`
   reject. I chose the latter.

2. `validateSingleSkill` errors on `source` unset. Today `SkillSource()`
   infers from Path/Repo, so this matters only for entries with neither —
   which is a malformed config. Should ValidateAndExpand produce a
   clearer "entry must specify path or repo" message? Probably yes, but
   deferred.
