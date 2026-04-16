# Trinity — History

## Project Context

- **Project:** hyoka — Go evaluation tool for AI-generated Azure SDK code, powered by Copilot SDK and multi-model review panels.
- **Stack:** Go 1.26.1+, GitHub Copilot CLI/SDK, MCP servers
- **User:** Ronnie Geraghty
- **My domain:** Frontend — serve/, report/, rerender/, templates/, site/, trends/

## Learnings

(New agent — no learnings yet.)

### Session 2026-04-04T00-05 (Morpheus Evolution Plan)

Evolution plan assigns you Phase 3 serve site evolution, comparison engine, transparent review panel UI, enhanced trends. Read `.squad/decisions.md` for full plan. Also assigned: runID validation in serve handler, HTML template extraction.

### Session 2026-04-07T21-20 (Morpheus Architecture Proposal)

Morpheus completed architectural proposal for embedding React SPA into Go binary using `go:embed`. Recommends pre-building in CI, committing `site/dist/` to repo, serving from embedded filesystem. No runtime auto-build. Decision recorded in `.squad/decisions.md`. Awaits Trinity implementation planning.

### Session 2026-04-07T21-28 (Issue #288 — Embed Site SPA)

Implemented site embedding per Morpheus's architecture proposal. Key decisions and patterns:

- **Embed file location:** `hyoka/internal/serve/embed.go` with `//go:embed all:site` directive. Uses `all:` prefix to include dotfiles if any exist.
- **Site copy pattern:** Built React SPA copied from `site/dist/` → `hyoka/internal/serve/site/` because `go:embed` requires paths relative to the Go source file and cannot use `..` paths.
- **Dual-source `fs.FS` pattern:** `spaHandler()` uses `fs.Sub(embeddedSite, "site")` for embedded mode, `os.DirFS(siteDir)` for dev override. Both share identical SPA fallback logic via `fs.FS` interface.
- **SPA fallback:** Opens file in `fs.FS`, checks `Stat()` for existence, serves via `http.FileServerFS`. Falls back to reading `index.html` via `fs.ReadFile` for client-side routing.
- **Test update:** `TestSPANoSiteDir` changed from expecting 404 to 200 since embedded site is always available.
- **PR:** #289 on branch `ronniegeraghty/issue-288-embed-site-spa`.

## 2026-04-16 — Phase 3 Merged to Dev (Neo)

Neo completed Phase 3 merge sequence: main→dev (hotfix #567 integrated), dev→Phase3 (clean), Phase3→dev (PR #562 squash-merged). Dev branch now has both Phase 3 features and starter-aware guardrail fix. All tests pass, CI green.
