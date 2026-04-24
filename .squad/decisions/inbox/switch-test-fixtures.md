# Switch — test.yaml fixture extension (skill-usage + intentionally-failing check)

**Date:** 2026-04-24
**Branch:** `ronniegeraghty/prompt-grader-checks`
**Scope:** Demo two more grader behaviors in `criteria/language/test.yaml`.

## Decision 1 — `tool_constraint` (not `behavior`) for "skill was used"

Both `behavior` (`required_tools:`) and `tool_constraint` (`required:` + `min_calls:`)
match against the tool *name* recorded in `ActionEvent.Tool`. Empirical check of a
prior `test/haiku` run (`reports/20260424-045906/.../report.json`) showed that
**Copilot skills surface as a single tool named `skill`** — the specific skill name
(`markdown-headings`, `markdown-lists`) is passed as an argument, not as the tool
name. `tool_calls` was: `["skill", "create"]`.

So `required_tools: [markdown-headings]` would always fail. Picked
**`tool_constraint`** with `required: [skill]` + `min_calls: {skill: 2}` — this
expresses "the agent invoked the skills tool at least twice (once per loaded
markdown skill)" and gives us a meaningful Pass when both fire. Inline YAML
comment above the entry explains this for future readers.

## Decision 2 — Deliberately-failing prompt-grader check

Appended to existing `Markdown Structure` `checks:` list:

```yaml
# Intentionally failing — demonstrates failed-check rendering.
- File hello.md contains a fenced code block tagged with the language `rust`.
```

A tiny `hello.md` (H1 + 3 bullets) cannot satisfy this; it will reliably render
as `❌ Fail` while the other two checks render as `✅ Pass` — producing the
desired `❌ Fail (2/3)` aggregate badge from Tank's renderer.

## Smoke output (run `20260424-053907`)

In this run the Haiku generator failed to actually write `hello.md` (flaky on
that prompt — known behavior; we've seen 2/3 success across recent runs). The
graders themselves all wired up cleanly:

```
[behavior]    Efficient Behavior:  ✅ Pass  (turn_limit ✅)
[tool_constr] Skill Usage:         ✅ Pass
   ✅ required: skill   — required tool "skill" called 2 time(s)
   ✅ min_calls: skill>=2 — tool "skill" called 2 time(s) (min 2)
[prompt]      Markdown Structure:  ❌ Fail (0/3)
   ❌ File hello.md exists and contains a level-1 heading…
   ❌ File contains exactly three bullet list items.
   ❌ File hello.md contains a fenced code block tagged with the language `rust`.
[output_chk]  Output Files Exist:  ❌ Fail
[file]        Hello.md File Check: ❌ Fail
```

Log line confirming Tank's badge math:
`bucket="Markdown Structure" passed=0 max=3` → renders as `❌ Fail (0/3)` in the
interactive progress display.

When the prompt *does* produce hello.md, this same set of checks will render as
`❌ Fail (2/3)` (the rust code-block check still fails) — which is the partial-
fail visualization Ronnie asked for.

## Caveats / followups

- `hyoka clean` hung in this session (orphaned in subprocess, had to `kill -9`).
  Not caused by this change — pre-existing flake worth noting for whoever owns
  the cleanup path.
- The Haiku generator on `test-dp-test-hello-markdown` is flaky enough that the
  partial 2/3 demo isn't guaranteed in any single run. Not blocking — the
  failed-check rendering is visible regardless.
