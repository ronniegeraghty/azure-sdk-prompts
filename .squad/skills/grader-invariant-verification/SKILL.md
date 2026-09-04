# grader-invariant-verification

**Owner:** Neo
**Last verified:** 2026-04-24 (commit `09d9d9bf` head)

## When to use

Whenever a site/UI report claims a grader is showing "no Points", "PASS/100% with no detail", or "blank labels", run this verification recipe BEFORE editing any grader code. It separates engine-side invariant violations from site/data-shape issues.

## Phase-3 invariant (recap)

Every `graders.GraderResult` MUST satisfy:

1. `len(Points) >= 1` (enforced by panic in `graders.NewResult`, `hyoka/internal/criteria/graders/grader.go:244`).
2. Every `Point.Label` is a non-empty string (a Label like `"criterion: "` is technically non-empty but considered a UX bug — see "Verbose vs blank" below).
3. Error/skip paths must construct results via `graders.NewErrorResult`, never an empty `GraderResult{}`.

The defensive synth in `engine_eval.go:1197` (`convertGraderResults`) catches any escapee with `slog.Warn("synthesizing fallback ...")`. **If you see that warning fire, you found a real bug.**

## Verification recipe (5 minutes)

```bash
# 1. Build current tip
mkdir -p .scratch && (cd hyoka && go build -o ../.scratch/h ./)

# 2. Run a fast Python prompt (key-vault-dp-python-crud is a known-fast canary)
./.scratch/h run --prompt-id key-vault-dp-python-crud \
  --config "baseline/claude-opus-4.6" \
  --log-level debug --log-file .scratch/debug.log

# 3. Locate the freshly-written report
REPORT=$(find reports/ -name 'report.json' -mmin -10 | head -1)

# 4. Three jq audits — ALL must pass
# (a) Every grader has ≥1 Point
jq '.grader_results[] | {name: .grader_name, type: .grader_type, pass, points_len: (.points | length), point_labels: [.points[]?.label]}' $REPORT
# (b) No Point has empty/null Label
jq '[.grader_results[].points[] | select(.label=="" or .label==null)]' $REPORT  # MUST be []
# (c) Defensive synth never fired
grep -i "synthesizing fallback\|graderless\|zero Points" .scratch/debug.log  # MUST be empty

# 5. Cleanup
./.scratch/h clean && rm -rf .scratch
```

## Interpreting results

| Result                              | Meaning                                                                 | Action                                                                  |
| ----------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| All 3 audits clean                  | Engine invariant holds. Bug is site-side (legacy data) or verbose label | Hand off to Trinity/site. Fix site to be defensive on legacy reports.   |
| `points_len == 0` for some grader   | A grader bypassed `NewResult` — real engine bug                         | Trace the grader's `Grade` method, fix to use `NewResult`/`NewErrorResult` |
| Empty/null label rows               | A `GraderPoint{}` literal omitted Label or sourced from empty var       | Fix grader to use `fmt.Sprintf("static-prefix: %s", v)` pattern         |
| `synthesizing fallback` log fired   | The defensive synth saved you — root cause the named grader/kind        | Find the code path that emitted Points-less Result; replace with `NewResult` |

## Path conventions

- **Top-level `.grader_results`** is the canonical array (post-v4 schema).
- **`.review`** at top level is the consolidated review *summary*, NOT graders. Don't audit `.review.grader_results` — it doesn't exist.
- Reports live at `reports/<run-id>/results/<service>/<plane>/<lang>/<category>/<prompt-id>/<config-path>/report.json`.

## Verbose vs blank labels (UX gotcha)

Engine treats these as valid:

- ✅ `"criterion: "` (non-empty string)
- ✅ `"criterion: DefaultAzureCredential Authentication\n   Check...\n   1. Uses..."` (non-empty, but multi-line)

Site renders both as broken. If the audit is clean but the site shows "blank-looking" labels, the LLM-judge response parsing in `internal/review/` is leaking bucket-header text into criterion Name. Normalize with `strings.Split(name, "\n")[0]` before passing to `fmt.Sprintf("criterion: %s", name)` in `prompt_review_grader.go` (lines 114, 208).

## Files relevant to the invariant

- `hyoka/internal/criteria/graders/grader.go` — `NewResult` (panic guard), `NewErrorResult`
- `hyoka/internal/criteria/graders/{file,program,prompt,behavior,output_check,prompt_review,prompt_grader_adapter}_grader.go` — every Point construction site
- `hyoka/internal/eval/engine_eval.go:1197` — `convertGraderResults` defensive synth + slog.Warn
