# Phase 0 Review — Switch 🤍

**Branch:** `ronniegeraghty/hyoka-0.3.1-phase0`
**Reviewed:** 2025-07-17
**Verdict:** ✅ **APPROVED**

---

## 1. Dead reference scan — ✅ PASS

| Pattern | Matches |
|---|---|
| `internal/tools` in `*.go` | 0 |
| `internal/build` in `*.go` (excluding build tags) | 0 |
| `internal/history` in `*.go` | 0 |
| `internal/manifest` in `*.go` | 0 |

All four removed packages (`hyoka/internal/tools`, `hyoka/internal/build`, `hyoka/internal/history`, `hyoka/internal/manifest`) are confirmed deleted with no dangling imports.

## 2. Build & test — ✅ PASS

| Check | Result |
|---|---|
| `go build ./...` | ✅ Clean |
| `go test -race ./... -timeout 3m` | ✅ 22 packages pass (1 no test files) |
| `cd site && npm run build` | ✅ Built in 6.9s |
| `cd site && npm test` | ✅ 8 test files, 43 tests pass |

## 3. Dead UI components — ✅ PASS

**Kept components verified in use:**

| Component | Imported by |
|---|---|
| `select.tsx` | pairwise-page, prompts-page, run-detail-page, eval-detail-page |
| `table.tsx` | prompt-detail-page, pairwise-page, run-detail-page, docs-page, dashboard-page |
| `utils.ts` (`cn()`) | select.tsx, table.tsx (10+ call sites) |

**Spot-check of deleted components:** The grep flagged `checkbox`, `command`, `popover`, `separator`, and `toggle` — all are **false positives**:

- `checkbox` — CSS attribute selector `[role=checkbox]` in table.tsx, not a component import
- `command` — data field name in mock-data.ts (`command: "dotnet build"`)
- `popover`/`separator` — Tailwind CSS class names (`bg-popover`, `data-slot="select-separator"`) in select.tsx, not component imports
- `toggle` — test description string ("toggle buttons for pass rate chart"), not a component import

No actual component imports reference deleted files.

## 4. WI-009 version detection — ✅ PASS

```
$ go run . version
hyoka version dev
```

Shows `dev` (expected for non-tagged builds via `runtime/debug.ReadBuildInfo()`), not hardcoded `0.3.0`.

## 5. WI-014 routing fix — ✅ PASS

Route pattern uses wildcard splat:
```ts
{ path: "runs/:runId/eval/:promptId/*", Component: EvalDetailPage }
```

Config name extracted via:
```ts
const { runId, promptId, "*": configName } = useParams();
```

This correctly handles config names containing `/` (e.g., `baseline/claude-opus-4.6`).

## 6. WI-013 CI workflow — ✅ PASS

`.github/workflows/ci.yml` contains a new `site-build-and-test` job with:
- `actions/checkout@v4`
- `actions/setup-node@v4` with Node 22 and npm cache (`cache-dependency-path: site/package-lock.json`)
- `npm ci` → `npm run build` → `npm test`

Runs on the same triggers as the Go job (`pull_request` all branches, `push` to main/dev).

## 7. Skill changes (WI-003) — ✅ PASS

**9 skills confirmed deleted:**

| Skill | Status |
|---|---|
| ci-validation-gates | ✓ Deleted |
| config-system | ✓ Deleted |
| copilot-sdk-integration | ✓ Deleted |
| grader-system | ✓ Deleted |
| process-lifecycle | ✓ Deleted |
| prompt-conventions | ✓ Deleted (consolidated into prompt-authoring) |
| property-migration | ✓ Deleted |
| report-generation | ✓ Deleted |
| serve-patterns | ✓ Deleted |

**eval-pipeline:** No mentions of "Build Verification Phase" — grep returns zero matches in `.agents/skills/eval-pipeline/SKILL.md`. Pipeline now documents 4 phases: Generation → Grading → Review Panel → Report Generation.

**project-conventions:** 91 lines of real content covering language/deps, logging, error handling, testing, CLI conventions, and code style. Not placeholders.

**prompt-authoring:** Consolidated skill exists at `.agents/skills/prompt-authoring/SKILL.md` with frontmatter format, file structure, and filtering guidance. The `prompt-conventions` directory is gone.

## 8. Additional observations

**WI-005 (.gitignore):** No legacy entries found. `.vscode/mcp.json` is tracked (`git ls-files` confirms).

**WI-006 (criteria):** Root `criteria/` directory removed. `.hyoka/criteria` symlink updated to `../examples/criteria` (valid target with language/ and service/ subdirs). Java criteria still exists in `examples/criteria/language/java.yaml` — this is fine as an example file; the "stale" java criteria at the root level was the one removed.

**WI-001 (dead packages):** `docs/examples/` confirmed deleted. All 4 internal packages confirmed deleted.

---

## Summary

All 8 work items verified. Zero regressions found. Build and tests pass cleanly across both Go and site.

**Phase 0 is approved.** ✅

---

## Morpheus 🕶️ — Architecture Review

**Branch:** `ronniegeraghty/hyoka-0.3.1-phase0`
**Reviewed:** 2025-07-17
**Verdict:** ✅ **APPROVED — with 4 low-severity issues to track for Phase 1**

Build passes. All 22 Go packages pass tests with `-race`. No regressions. The dead code removal is surgically clean in Go source. I'm calling out documentation drift and two minor skill inaccuracies that should be cleaned up but are not merge-blockers.

---

### 1. Dead Code Removal — ✅ PASS (with doc drift)

**Go source: clean.** All 4 removed packages (`tools/`, `build/`, `history/`, `manifest/`) have zero dangling imports. Every remaining `internal/` package has at least 1 external import — no other dead packages exist. `cmd/root.go` registers 14 commands; all resolve to valid `.go` files. No orphaned `_test.go` files or testdata reference removed types.

**Documentation drift (not blocking):**

| File | Line | Stale Reference |
|---|---|---|
| `hyoka/cmd/root.go` | 30 | `Long:` still says "running build verification" |
| `hyoka/README.md` | 52, 135, 431, 436, 445 | References `--verify-build`, `build/`, `manifest/` |
| `docs/architecture.md` | 45, 49, 52 | Tree listing includes `tools/`, `history/`, `manifest/` |
| `docs/contributing.md` | 50, 57, 59, 68 | Tree listing includes `build/`, `history/`, `manifest/`, `tools/` |

**Recommendation:** Track as Phase 1 hygiene. The `root.go` Long description is user-facing — fix that soonest.

---

### 2. Skill Quality — ⚠️ 2 MINOR ISSUES

**eval-pipeline/SKILL.md:**
- **Line 3:** Description says `"generate → grade → report"` — should be `"generate → grade → review → report"` (review phase is missing from the short description)
- **Line 67:** `"Build failure → reported but review proceeds"` — stale reference to removed build phase. Should be removed.
- Pipeline phases (Generation → Grading → Review Panel → Report) are otherwise correctly documented.

**prompt-authoring/SKILL.md:** ✅ Excellent consolidation. Covers both flat and nested `properties:` format with correct field names, valid values, and directory structure. Tool filtering with `When` is accurately documented.

**project-conventions/SKILL.md:** ✅ Accurate. Logging (slog), CLI (cobra), config (yaml.v3), error handling (%w), testing (table-driven, -race), and Git conventions all match the actual codebase.

**Removed skills vs. active code:** Switch's review noted 9 skills deleted. Several of those (grader-system, config-system, report-generation, serve-patterns, copilot-sdk-integration) still have active code in `internal/`. This is acceptable — these were removed as stale *skill files* whose content had drifted from the code, not as removal of the code itself. They can be rewritten accurately in a future pass if agents need them. The AGENTS.md custom instruction already covers the essential architecture for agent context.

---

### 3. Version Detection — ✅ PASS

`runtime/debug.ReadBuildInfo()` is correctly implemented in `cmd/root.go:13-16`:
- Checks `info.Main.Version != ""` and `!= "(devel)"` — correct guard
- Falls back to `"dev"` — appropriate for local `go run` and untagged builds
- No `go.work` file exists, so no workspace complications with `ReadBuildInfo()`
- `versionCmd()` correctly reads the package-level `Version` variable

The `rootCmd.Version` field (cobra's built-in) is not set — the project uses a custom `version` subcommand instead. This is a valid pattern; cobra's `--version` flag is optional.

---

### 4. CI Workflow — ✅ PASS (1 optimization note)

**Node.js 22:** ✅ Current LTS, appropriate for React Router v7 + Vite 6.

**Commands:** ✅ `npm ci` → `npm run build` → `npm test` — correct sequence. `actions/setup-node@v4` with `cache: 'npm'` and `cache-dependency-path: site/package-lock.json` handles caching properly.

**Optimization note (not blocking):** The `site-build-and-test` job runs on every PR regardless of whether `site/` files changed. Consider adding a `paths:` filter or making it conditional via `paths-filter` action. Low priority — the job completes in ~15s.

**Missing lint step:** No `npm run lint` in the workflow. This is acceptable if linting isn't configured in the site's `package.json` scripts. Worth adding when/if ESLint or similar is set up.

---

### 5. Routing Fix — ✅ PASS (1 edge case noted)

The splat route pattern is the correct React Router v7 approach for config names containing `/`:

```ts
// routes.ts
{ path: "runs/:runId/eval/:promptId/*", Component: EvalDetailPage }

// eval-detail-page.tsx
const { runId, promptId, "*": configName } = useParams();
```

**Encoding is handled correctly:** `promptId` is encoded via `encodeURIComponent()` in links and decoded on consumption. `configName` is passed raw through the splat — this is correct because the splat captures the entire remaining path including literal `/` characters, so no encoding/decoding is needed for slash-separated config names.

**Edge case (theoretical, not blocking):** Config names containing `?`, `#`, or `%` would break routing. Current config names (`baseline/claude-opus-4.6`, `azure-mcp/claude-sonnet-4.5`, etc.) use only alphanumeric + `/` + `-` + `.`, so this is safe. If config naming expands to include special characters, the link construction in `run-detail-page.tsx:199` would need `encodeURIComponent()` per segment.

---

### 6. Impact on Phase 1+ — ✅ CLEAN STATE

The codebase is in a healthy state for Phase 1. No structural issues created by Phase 0 changes.

**Items to track for Phase 1:**

| Priority | Item | Effort |
|---|---|---|
| P2 | Fix `root.go:30` Long description — remove "build verification" | 1 line |
| P3 | Update `docs/architecture.md` and `docs/contributing.md` package trees | 10 min |
| P3 | Fix eval-pipeline skill description and remove build failure reference | 2 lines |
| P3 | Update `hyoka/README.md` — remove `--verify-build` flag and stale tree entries | 10 min |
| P4 | Consider `paths:` filter for site CI job | Optional |

None of these block Phase 1 work. The Go source and test suite are clean. The site builds and routes work. Phase 0 achieved its objective: reduce dead weight before feature work begins.

---

### Summary

| Review Area | Verdict | Notes |
|---|---|---|
| Dead code removal (Go source) | ✅ Clean | All 4 packages fully removed, zero dangling imports |
| Dead code removal (docs) | ⚠️ Drift | 4 doc files still reference removed packages |
| Skill accuracy | ⚠️ Minor | eval-pipeline has 2 stale refs; other skills accurate |
| Version detection | ✅ Correct | ReadBuildInfo + "dev" fallback is the right pattern |
| CI workflow | ✅ Good | Node 22 + caching + proper commands |
| Routing fix | ✅ Correct | Splat route is the right react-router pattern |
| Phase 1 readiness | ✅ Ready | Clean foundation, no structural debt introduced |

**Phase 0 is approved.** ✅
