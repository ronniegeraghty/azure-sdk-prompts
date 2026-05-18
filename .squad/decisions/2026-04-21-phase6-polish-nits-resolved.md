# Phase 6 Polish Nits — Round 3 Resolution

**Date:** 2026-04-21
**Author:** Scribe (recording from PR #613 / #614 review threads + merge outcomes)
**Status:** Decided — original deferred set RESOLVED; new (smaller) follow-up set named below
**Supersedes (in part):** [`2026-04-21-phase6-polish-nits-deferred.md`](./2026-04-21-phase6-polish-nits-deferred.md) — see "Resolution status" below
**Context:** Ronnie reversed the "defer to Phase 7 issues" call and asked the team to address the deferred polish items in-phase. Two parallel workstreams shipped:
- **PR #613** — Trinity: `test(#608): MultiSelectFilter toggle/onChange/aria/inside-click coverage` (squash-merged into `phase-6` at `d05855df`)
- **PR #614** — Morpheus: `chore(#608): site-embed freshness CI hardening (concurrency, untracked detection, prune)` (squash-merged into `phase-6` at `e0b72c63`)

Both reviewed dual-track (test + arch) with author/reviewer isolation:
- #613: Switch (test) ✅ APPROVE + Morpheus (arch) ⚠️ APPROVE WITH NOTES
- #614: Switch (test) ✅ APPROVE + Neo (arch, subbing for Morpheus on author-isolation grounds) ⚠️ APPROVE WITH NOTES

`site/` test count: 122 → 133 (+11, all green). `go build ./hyoka/...` and `go test -race ./hyoka/... -timeout 3m` both clean. PR #607 stayed `MERGEABLE / CLEAN` after the merges.

## Resolution status — original deferred items

### From PR #609 (Switch on Trinity's MultiSelectFilter tests) → all RESOLVED in PR #613
- ✅ Toggle / `onChange` behavior — 3 tests including a controlled-state `Wrapper` with `toHaveBeenNthCalledWith` validating each transition (not just final state).
- ✅ Summary-text branches (single / multi / overflow) — 5 tests, all paths.
- ✅ `aria-expanded` (dynamic toggle) and `aria-selected` (per-option) asserted.
- ✅ Inside-click counterpart to outside-click — listbox stays mounted on inside-click.

Keyboard listbox navigation remains **product-side gap** for Trinity (component doesn't implement it yet); routes separately and is not part of this resolution.

### From PR #611 (Neo on Morpheus's CI freshness gate) → all RESOLVED in PR #614
- ✅ `concurrency:` group added (`${{ github.workflow }}-${{ github.ref }}` with `cancel-in-progress: true`).
- ✅ `git status --porcelain` replaces `git diff --quiet` — catches new untracked asset filenames (vite content-hashing scenario).
- ✅ `rm -rf $(EMBED_DIR)/*` replaces `rm -rf $(EMBED_DIR)/assets` — wholesale prune handles future vite outputs (favicons/manifests) without Makefile edits.
- ✅ `phase-6` removed from push triggers with self-explanatory inline comment naming the future-pruning pattern.
- ⏸ **Duplicate-work consolidation deferred-by-design.** Morpheus + Neo agreed to keep `verify-embed` discrete from `site-build-and-test` rather than fold them: ~1–2 min/run cost is real but bounded; named-required-check signal-clarity is durable. Decision documented in PR body and in-source comment so the next reviewer doesn't re-litigate.

## NEW deferred set (post-merge follow-ups)

These are **non-blocking** for `phase-6 → dev` and for #607. They were surfaced by reviewers during round 3 and named here so they don't evaporate.

### N1. MultiSelectFilter value-vs-label UX bug
**Where:** `site/src/components/multi-select-filter.tsx:57`
**What:** Single-select summary renders `selected[0]` (the raw value, e.g. `"a"`) instead of the matching `opt.label` (e.g. `"Alpha"`). Inconsistent with the option list (line 110, uses `opt.label`) and with the multi-selected branch.
**Confirmed by:** Trinity (PR body), Morpheus (arch review #613), Switch (test review #613).
**Disposition:** Trinity correctly LOCKED IN current behavior with an inline `// Note:` comment in the test rather than silently fixing in a tests-only PR. Fix is one line: `options.find(o => o.value === selected[0])?.label ?? selected[0]`.
**Suggested owner:** Trinity (file + fix as a separate one-line product PR with visual confirmation).

### N2. Makefile `EMBED_DIR` shell-injection guard
**Where:** `Makefile`, the `site-embed` target's `rm -rf $(EMBED_DIR)/*`
**What:** `EMBED_DIR := hyoka/internal/serve/site` is hardcoded and safe at parse time, but `make EMBED_DIR= site-embed` would expand to `rm -rf /*`. Pre-existing risk pattern — NOT introduced by #614.
**Surfaced by:** Switch on #614.
**Suggested fix:** `ifeq ($(strip $(EMBED_DIR)),)` guard at the top of the target, with `$(error ...)` halt.
**Suggested owner:** Morpheus or Neo (tiny defensive PR).

### N3. embedded-asset-freshness skill prose stale
**Where:** `.squad/skills/embedded-asset-freshness/SKILL.md` (~line 57, "refresh procedure" section)
**What:** Skill prose still describes the OLD refresh procedure (manual `rm -rf assets && cp -r`) instead of the now-canonical `make site-embed` / `make verify-embed`.
**Surfaced by:** Neo on #614.
**Suggested owner:** any agent (hygiene PR, doc-only).

### N4. `.gitignore`'d files in EMBED_DIR are invisible to the freshness gate
**Where:** Interaction between `.gitignore` rules covering `EMBED_DIR`, `git status --porcelain` (used by `verify-embed`), and `//go:embed all:site` (the `all:` prefix embeds dotfiles/underscore-files).
**What:** A gitignored file in EMBED_DIR would silently ship via `//go:embed all:site` while being invisible to the freshness check. Pre-existing, NOT introduced by #614.
**Surfaced by:** Neo on #614.
**Suggested fix:** Optional defense-in-depth — pass `--ignored` to `git status` in `verify-embed`. Not worth doing now without a triggering incident.
**Suggested owner:** whoever picks N2 (related Makefile work).

## Rationale

Round 2 deferred these polish items to "Phase 7 issues" (see superseded decision). Ronnie reversed that call: the cost of finishing the round in-phase (two small PRs, four reviews, ~one cycle) was lower than the cost of carrying open polish into the post-merge window. This trade is recorded so the heuristic ("when the deferred set is small and well-bounded, finish in-phase") informs future deferral decisions.

The new deferred set (N1–N4) is genuinely smaller and more peripheral than the original four:
- N1 is a one-line UX bug behind a feature flag of "did anyone notice".
- N2/N4 are pre-existing risks that #614 made visible but did not introduce.
- N3 is doc hygiene.

None of N1–N4 block `phase-6 → dev` or PR #607.

## Owners

- N1 (value-vs-label UX) → Trinity
- N2 (Makefile guard) → Morpheus or Neo
- N3 (skill prose refresh) → any agent
- N4 (`.gitignore` blind spot) → bundles with N2
