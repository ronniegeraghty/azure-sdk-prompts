# Grader Serialization & Per-Grader Progress Events

**Author:** Neo
**Branch:** ronniegeraghty/dev
**Sprint todo:** #5 — "Serialize grader execution in the interactive path and
emit per-grader events"
**Scope:** `hyoka/internal/criteria/exec.go`, `hyoka/internal/eval/engine.go`,
`hyoka/internal/eval/engine_eval.go`

## Summary

Wired `EventGraderStart` / `EventGraderComplete` progress events around each
grader's `Grade()` call so the interactive display can render a per-grader
lifecycle ("🔄 Running… → ✅ Pass (8/10)") on the tail line. Existing grader
output, aggregate `GraderResult`s, and the report JSON are unchanged — this
is a UX-only signal layer.

## Mode-detection rule

**Rule:** Event emission is gated on *reporter presence*, not on worker count.
When `engine.go` wires a non-nil `sendRawEvent` callback into `runSingleEval`,
grader events are emitted. When it doesn't (e.g. `display == nil` path — not
currently possible but guarded anyway), emission is a no-op.

**Why not gate on `opts.Workers == 1`:**

- Graders are **already sequential** in both modes:
  - Typed graders (`file`, `program`, `output_check`, `behavior`, ...) run
    in a single `for` loop inside `criteria.RunGraders` — no goroutines.
  - The `prompt_review` grader runs *after* typed graders finish, in the
    same function. No concurrency between grader groups.
- Therefore "serialize in interactive mode" is a no-op at the runtime level;
  the only new behavior is event emission.
- If we later parallelize anything (e.g. run review alongside typed graders
  in CI mode), the parallelism gate belongs there — adding a dead
  `if interactive` branch here today would just be noise.

The `interactive` signal *is* available via `opts.Workers` if downstream
display renderers want to pick different layouts, but the emission layer
doesn't consult it.

## Grader ID & Kind convention

The `GraderID` field on each event is populated from `Grader.Name()`:

| Grader kind (`GraderKind`) | `GraderID` source | Example `GraderID`s |
|---|---|---|
| `prompt_review` | `"ai_review"` literal in `engine_eval.go` | `ai_review` |
| `file` | grader's configured name | `require_main_py`, `has_readme` |
| `program` | grader's configured name | `pytest_passes`, `mypy_clean` |
| `output_check` | grader's configured name | `no_secrets`, `min_files_2` |
| `behavior` | grader's configured name | `uses_mcp_required` |
| `action_sequence` | grader's configured name | |
| `tool_constraint` | grader's configured name | |
| `prompt` (LLM judge) | grader's configured name | |

`GraderID` is stable within a single eval — the same string is emitted on
both `GraderStart` and `GraderComplete` so the renderer can match them.

## Score / Message field population

Per the schema memo, `Score` is `*float64` so "not reported" is distinguishable
from a legitimate `0.0`. Policy for populating it:

| Grader `Kind` | `Score` populated? | Reason |
|---|---|---|
| `prompt_review` | **yes** — `result.Score` (0.0–1.0) | LLM panel produces a real weighted score (`criteriaScore(consolidated)`) |
| `prompt` | **yes** — `result.Score` (0.0–1.0) | Single-model LLM judge produces a normalized score |
| `file` | **no** (nil) | `Score` is just `1.0` on pass / `0.0` on fail — binary, rendering "pass (0/10)" would mislead |
| `program` | **no** (nil) | Same — binary pass/fail from exit code |
| `output_check` | **no** (nil) | Binary AND-of-sub-checks |
| `behavior` | **no** (nil) | Binary — all constraints met or not |
| `action_sequence` | **no** (nil) | Binary match |
| `tool_constraint` | **no** (nil) | Binary constraint satisfaction |

**`Message`** is populated with whatever the grader already put in
`GraderResult.Message` — a short human-readable summary (e.g. `"file main.py
not found"`, `"panel score 7/10"`, `"grader execution error: <err>"`). No
fabricated text; if the grader left `Message` empty the event leaves it empty
too.

**`Result`** is always populated — `"pass"` when `result.Pass == true`, else
`"fail"`. Uses the `progress.GraderResult*` constants so renderers don't
hardcode strings.

## API changes

- `criteria.GraderHooks` (new type): optional `OnStart(g Grader)` and
  `OnComplete(g Grader, r GraderResult)` callbacks. Zero-value is valid (both
  nil) — used when emission isn't wanted.
- `criteria.RunGradersWithHooks(ctx, instances, configs, input, hooks)` (new
  function): same body as `RunGraders`, plus pre/post hook invocations.
- `criteria.RunGraders` unchanged (delegates to `RunGradersWithHooks` with
  zero hooks) — preserves existing call sites and tests.
- `engine.runSingleEval` signature gains a `sendRawEvent func(ProgressEvent)`
  callback (sixth arg). The single caller in `engine.go` wires it to
  `display.HandleEvent` after auto-filling `EvalID` / `PromptID` /
  `ConfigName` so callers inside `engine_eval.go` can emit rich events
  without duplicating identity plumbing.

## Backward compatibility

- Existing callers of `RunGraders` untouched.
- Report JSON (`GraderResults`, `ScoreBreakdown`, aggregate `Pass` / `Score`)
  unchanged — grader events are in-process UX only.
- All existing display paths, tests, and emitter call sites compile and
  behave unchanged. New test coverage: `TestRunGradersWithHooksInvokesCallbacks`,
  `TestRunGradersNilHooksStillWorks` in `internal/criteria/exec_test.go`.
- Verified: `go build ./...`, `go test -race ./hyoka/internal/eval/...
  ./hyoka/internal/criteria/... ./hyoka/internal/progress/...`, `go vet
  ./hyoka/...` all pass.

## Downstream todos that depend on this

- **Interactive renderer (sprint todo #6)** — consume `EventGraderStart` as
  the new tail "Running…" line, swap to Pass/Fail on `EventGraderComplete`,
  render score bracket `(8/10)` only when `Score != nil`.
- **CI renderer (sprint todo #7)** — append one timestamped line per grader
  lifecycle (or just per-eval summary with `graders: N/M passed` — designer's
  call, the events carry enough info either way).
