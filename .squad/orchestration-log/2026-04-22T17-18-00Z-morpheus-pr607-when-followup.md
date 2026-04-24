# Orchestration Log: Morpheus — PR #607 Review Comment Follow-up

**Date:** 2026-04-22T17:18:00Z
**Agent:** Morpheus (🏗️)
**Mission:** Investigate inline review comment on PR #607 (comment 3125681737) regarding hierarchical criteria schema.

## Scope

PR #607 review comment on `examples/criteria/hierarchical-when-example.yaml:66`: Does the schema support two group-level `when` blocks declared with YAML `---` doc separator?

## Finding

**No. The example is misleading and the loader has a silent truncation bug.**

The file uses `---` document separator (line 46) to suggest two sibling group-level `when` blocks are valid. However:

- `hyoka/internal/criteria/criteria.go:134-136`: `yaml.NewDecoder` + single `Decode()` call processes **only the first YAML document**. Second document (Rust block, lines 46-66) silently discarded.
- `hyoka validate` doesn't flag it because validation operates on parsed structure, not raw bytes.

The schema's correct mechanism for multiple group-level `when`s is the top-level `groups:` list (each `GraderGroup` has optional `when`). Canonical example in `hyoka/internal/criteria/hierarchical_test.go:208-247`.

## Risk

Anyone copying the example pattern loses half their criteria silently. Worst kind of bug — no error, looks like it works.

## Recommended follow-ups (file as separate issues)

1. **Fix the example** — rewrite using `groups:` list instead of `---` separator.
2. **Fix the loader** — either strict multi-doc rejection OR merge documents (strict rejection safer first move).
3. **Validate coverage gap** — detect trailing YAML documents even before loader fix.

## Outcome

❌ Example is incorrect and loader silently truncates YAML documents. Issues recommended for separate PRs.

## Learnings captured

Appended to Neo history.md: "Loader silently drops YAML docs after first --- in criteria files." Since Neo would own the loader fix.

Also appended to Oracle history.md: "Examples can have misleading patterns (e.g., hierarchical-when-example.yaml uses discarded YAML docs); audit during docs maintenance."
