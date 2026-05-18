---
name: embedded-asset-freshness
description: Verify site/dist/ is up to date before merging site changes (go:embed reads directly from committed bundle).
confidence: high
applies-to: [site, serve]
---

# Embedded Asset Freshness

## Current implementation

hyoka's site is shipped via `go:embed all:dist` in `site/embed.go`. The bundle lives at **`site/dist/`** and is committed to git. Vite builds deterministically (content-hashed filenames), so rebuilding identical source produces byte-identical output.

**CI enforcement:** The `site-bundle-freshness.yml` workflow rebuilds site/dist/ on every PR and fails if `git diff --exit-code site/dist/` detects uncommitted changes.

## Refresh procedure

```bash
cd site && npm run build
git add site/dist/
git commit -m "site: rebuild bundle"
```

CI will verify the bundle is fresh. No manual copy step required — go:embed reads site/dist/ directly.

## When stale bundles occur

- Editing `site/src/**` without running `npm run build`
- Updating dependencies (`npm install`) that affect bundle output without rebuilding

## When to invoke this skill

- Diagnosing "I merged the feature but `hyoka serve` still shows the old thing" reports
- Reviewing PRs that touch `site/src/**` but don't update `site/dist/`
- Understanding why CI failed on `site-bundle-freshness` workflow

## See also

- `.github/workflows/site-bundle-freshness.yml` — CI enforcement
- `site/embed.go` — go:embed directive
