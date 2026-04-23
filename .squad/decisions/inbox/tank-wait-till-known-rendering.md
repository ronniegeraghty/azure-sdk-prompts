# Decision: Wait-Till-Known Tool Rendering (Tank, 2026-04-24)

**Status:** 📥 Inbox — awaits Scribe merge into decisions.md
**Author:** Tank (Backend/Progress Dev)
**Relates to:** WU-A3, commit 18d105c3

## Context

Post-eval, the Tools section showed stuck "🔄 Loading…" rows (never resolved) and missing fan-out for plugins/skills-dirs. Root cause was the interactive renderer emitting a transient loading line at every `EventToolResolutionStart` without guaranteeing a matching `EventToolResolutionResult`. Additionally, plugin parents were double-rendered (flat top-level row + grouped parent header).

## Decision

The progress renderer now uses a **wait-till-known** rule: Start events update internal bookkeeping only; Result events commit a single final line. Parent headers (plugin, skills-dir) are emitted exactly once on the first matching child. Containers are detected by kind (plugin) or by reference (any entry appearing as a child's `ParentName`). Orphan "loading" entries (Start without Result) are filtered from the rendered output rather than surfaced as failures.

Rationale:
1. Eliminates the stuck-pending class of rendering bugs.
2. Makes the renderer resilient to resolver quirks (e.g. `validate.go` emits a Start for skill_dir parents but no Result; emitting children with `ParentName=entry.Path` rather than `entry.Name`).
3. Deduplicates plugin rendering (one grouped block instead of top-level leaf + header).

## Known upstream bug (not fixed here — Neo's territory)

`validate.go`'s `validateSkillDirEntry` emits `emitStart(emit, entry.Name, …)` but children use `ParentName=entry.Path`. The renderer renders the header using whatever `ParentName` arrives (absolute path, ugly). Recommend Neo either skip the parent Start OR have children reference `entry.Name` consistently.

## Renderer → resolver contract

Fields consumed by the renderer:
- `ToolName`, `ToolKind`, `Status`, `Reason`
- `ParentName`, `ParentKind` (both set together; non-empty = grouped child)

Container detection rule: a top-level entry is a container iff `ToolKind == "plugin"` OR its `ToolName` matches the `ParentName` of at least one other entry with non-empty `ParentKind`.

Adding new fields (e.g. `Source: remote|local`) is additive — the renderer can surface them in `renderToolLine` without changing grouping logic.

## What was not done

- Live `go run . run …` validation: blocked by unrelated in-flight build breakage in `internal/config/plugins.go` (Neo's schema sweep still in progress). Progress package builds and tests pass in isolation. Re-verify end-to-end once Neo's sweep lands.
