---
name: embedded-asset-freshness
description: Verify `go:embed`-backed frontend bundles are refreshed before merging site changes. A distribution-layer variant of observable-wiring-tests.
confidence: high
applies-to: [site, serve, rollup-reviews]
---

# Embedded Asset Freshness

## The problem

hyoka's site is shipped via `go:embed all:site` in `hyoka/internal/serve/embed.go`. The canonical embedded tree lives at **`hyoka/internal/serve/site/`**. Vite builds to `site/dist/`; a manual copy step is required. Vitest tests hit `site/src/**` (source) — they do **not** exercise the embedded bundle.

**Failure mode:** A site PR modifies `site/src/**`, adds component tests that pass, and lands. `hyoka/internal/serve/site/` is never refreshed. The merged binary ships the **previous** UI. All CI is green.

This is the [#587-trap](../observable-wiring-tests/SKILL.md) at the distribution layer: source + tests prove the feature exists in the codebase, but the shipped artifact doesn't contain it.

**Real recurrence:** PR #607 (Phase 6 rollup). Six PRs touched `site/src/**` (#601 Compare redesign, #604 filter bar, et al.). Zero refreshed `hyoka/internal/serve/site/`. `go run . serve` on `phase-6` HEAD served the Phase 4 bundle.

## Checks any rollup reviewer must run

### 1. Git log the embed path

```bash
git log --oneline <base>..<head> -- hyoka/internal/serve/site/
```

If this is empty **and** the same range touches `site/src/**`, the bundle is stale. Reject.

### 2. Run the actual server and diff the bundle hash

```bash
# On the PR branch (fresh checkout, no manual build):
go run . serve --port 8080 &
curl -sS http://localhost:8080/ | grep -oE "index-[A-Za-z0-9]+\.js"
# Compare to what a fresh build produces:
cd site && npm run build
grep -oE "index-[A-Za-z0-9]+\.js" dist/index.html
```

Different hashes → stale embed.

### 3. Drive the new UI with Playwright against the served binary

If a reviewer loads the served site and the new component (filter bar, redesigned page, etc.) is missing from the DOM, the source tests are lying.

```bash
playwright-cli open http://localhost:8080/<new-feature-route>
playwright-cli --raw eval "document.body.innerHTML.includes('<marker-for-new-feature>')"
```

## The refresh procedure (1 commit)

```bash
cd site && npm run build
rm -rf ../hyoka/internal/serve/site/assets
cp -r dist/* ../hyoka/internal/serve/site/
cd .. && git add hyoka/internal/serve/site/ \
  && git commit -m "chore: refresh embedded site bundle"
```

Must include `index.html` (new script/css hashes) **and** both files in `assets/`.

## Permanent fixes (Phase 7 candidates)

Choose one — this shouldn't rely on reviewer diligence:

- **Make target**: `make build` runs `vite build` + copy + `go build ./...` atomically.
- **CI hash-diff**: in any PR touching `site/**`, compute `sha256sum site/dist/index.html` after rebuild and fail if it differs from the committed `hyoka/internal/serve/site/index.html`.
- **Pre-push hook**: shipped via `hyoka init` or committed under `.githooks/`, rebuilds + stages the embed on any `site/**` change.

## When to invoke this skill

- Reviewing any **rollup PR** that aggregates site-touching sub-PRs.
- Reviewing any site PR where the reviewer can't reproduce the new UI via `go run . serve`.
- Diagnosing "I merged the feature but `hyoka serve` still shows the old thing" reports from users.

## See also

- `.squad/skills/observable-wiring-tests/SKILL.md` — source-layer sibling
- `.squad/skills/playwright-cli/SKILL.md` — live verification tool
- Switch's Phase 4 history item #6 (first documented recurrence of this foot-gun)
- PR #607 review (first rollup-level catch)
