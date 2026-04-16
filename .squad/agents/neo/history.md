# Project Context

- **Project:** hyoka — Go evaluation tool that runs prompts through the Copilot SDK, reviews code via a multi-model panel, produces criteria-based pass/fail reports
- **Stack:** Go 1.26.1+, GitHub Copilot CLI/SDK, MCP servers (Azure MCP via npx)
- **User:** Ronnie Geraghty
- **Created:** 2026-04-03
- **Repo:** /home/rgeraghty/projects/hyoka
- **Key paths:** hyoka/internal/ (core engine), hyoka/main.go (entry point), prompts/ (evaluation prompts), criteria/ (pass/fail criteria)

## Core Context

Agent Neo initialized as Core Dev for hyoka. Owns the evaluation engine, review panel, criteria logic, and Copilot SDK integration. The tool has guardrails (max turns: 25, max files: 50, max output: 1MB, max session actions: 50). Safety boundaries prevent real Azure resource provisioning by default (`--allow-cloud` to opt out).

## Recent Updates

📌 Team initialized on 2026-04-03

📋 **Morpheus Audit (2026-04-03):** Comprehensive codebase health assessment complete. Key finding: **reviewer model bug in main.go:469-473** (P0) — multi-config evaluations share one reviewer panel. See `.squad/decisions.md` for full P0/P1/P2 action items.

## Learnings

### Issue #92: Per-Task Reviewer Panel Creation (2025-01-19)
**Branch:** `ronniegeraghty/issue-92-reviewer-model-bug`  
**PR:** [#170](https://github.com/ronniegeraghty/hyoka/pull/170)  
**Status:** ✅ Complete

Fixed critical bug where multi-config evaluations reused the reviewer panel from the FIRST config for ALL configs, causing every evaluation to use incorrect reviewer models.

**Implementation:**
- Introduced `ReviewerFactory` function type that creates reviewers per-config
- Replaced `Engine.reviewer/panelReviewer` fields with `reviewerFactory` field
- Moved reviewer creation from main.go into `runSingleEval()` using `task.Config`
- Each config now gets its own reviewer panel with correct models
- Added `NewEngineWithReviewerFactory()` constructor
- Maintained backward compatibility with deprecated `NewEngineWithReviewer()`

**Testing:**
- Created `reviewer_factory_test.go` with 3 tests verifying correct behavior
- All existing tests pass
- Build and vet clean

**Learnings:**
1. **Reviewer Factory Pattern**: When multiple configs need different reviewer settings, create reviewers lazily per-task rather than once at engine creation. Use a factory function that closes over shared resources (clientOpts) but creates instances based on task.Config.

2. **Backward Compatibility**: When refactoring constructors, wrap deprecated APIs to call new implementation. Preserves existing call sites while enabling new patterns.

3. **Testing Concurrent Tasks**: Don't assert on execution order. Use maps to track outcomes by task ID when testing engines with concurrent workers.

Initial setup complete. Architecture is sound. Main engineering focus should be: (1) fix reviewer model bug, (2) refactor main.go into cmd/ package, (3) add integration tests.

### Session 2026-04-04T00-05 (Morpheus Evolution Plan)

Evolution plan assigns you Phase 1 core model changes (generic properties, criteria filters, tool filters) and Phase 2 pairwise testing. Read `.squad/decisions.md` for full plan. Also assigned: reviewer model bug (P0), discarded error logging, early auth check.

### Session 2026-04-04T19:45 (Phase 0 Execution — Reviewer Factory Fix)

**Status:** COMPLETE  
**Issue:** #92  
**PR:** #170

Implemented ReviewerFactory pattern to fix multi-config reviewer panel bug. Each config now receives correct reviewer models instead of all configs using first config's reviewers.

**Key outcome:** Lazy per-task reviewer creation replaces engine-scoped setup. Factory pattern enables clean separation of concerns and tested backward compatibility.

**Cross-agent dependency:** Tank's config migration (#96, PR #171) enabled clean Generator/Reviewer schema that makes this fix viable. Switch's flaky test fix (#99, PR #167) and Tank's CI pipeline (#91, PR #168) ensure test reliability in review panel code.

**Files:** engine.go, main.go, reviewer_factory_test.go

### Session 2026-04-07T03:47 (P0 Config Hardening)

**Status:** COMPLETE  
**PR:** [#256](https://github.com/ronniegeraghty/hyoka/pull/256)  
**Branch:** `ronniegeraghty/p0-config-hardening`

Implemented P0 config hardening from codebase audit (#252, items 1-2):

1. **Validation Before Logging:** Moved `cfg.Validate()` before the config-loaded logging loop in `Parse()` to prevent nil-pointer panic when accessing `c.Generator.Model` on invalid configs.

2. **Negative Limit Rejection:** Added validation to reject negative values for all SessionLimits fields (max_turns, max_files, max_output_size, max_session_actions). Zero is explicitly allowed (means "use default").

**Implementation:**
- Fixed validation order in `config.go:125-133` (Parse function)
- Added negative-value checks in `config.go:183-197` (Validate function)
- Created 6 new tests: negative limits (4 tests), zero limits acceptance, nil generator panic prevention
- Fixed UTF-8 encoding (em-dash in comment) and Go formatting

**Testing:** All 69 config package tests pass. Full build succeeds.

**Key Learning:** Validation order matters for early error detection. Always validate config structure BEFORE accessing fields that may not exist. This prevents cryptic panics and provides clear error messages to users.

**Files:** hyoka/internal/config/config.go, hyoka/internal/config/config_test.go


### Session 2026-04-08 (Config Unification #252)

**Status:** COMPLETE

Unified generator/reviewer tooling into a single `tools` array with typed entries (`tool`, `mcp`, `skill`). Updated config validation, tool resolution, MCP wiring, pairwise ablation logic, plugin merging, and skill path resolution to operate on typed ToolEntry records.

**Implementation:**
- Extended `ToolEntry` with MCP and skill metadata plus type normalization
- Reworked eval/session setup to derive MCP servers, skills, and available tools from `generator.tools`
- Updated configs and docs to the unified `tools` schema and migrated tests

**Testing:** `go build ./hyoka/...`, `go vet ./hyoka/...`, `go test ./hyoka/... -count=1`

**Key Learning:** Centralizing tool configuration simplifies downstream consumers — filter by type once and reuse the same entries for MCP, skills, and tool allowlists without duplicating schema fields.

### Phase 3 Integration with Hotfix #567 (2026-04-16)
**Branch:** Merged PR #562 into `ronniegeraghty/dev`  
**PR:** [#562](https://github.com/ronniegeraghty/hyoka/pull/562) (Phase 3: Advanced Core & CLI Polish)  
**Merge commits:** `1ef6081d` (main→dev), `4b4e95f9` (Phase 3→dev)  
**Status:** ✅ Complete

Successfully integrated hotfix #567 (starter-aware guardrails) with Phase 3 work. The challenge was that Phase 2 split engine.go into engine.go + engine_eval.go, while main still had everything in engine.go.

**Conflict Resolution:**
- Main tried to add `runSingleEval()` to engine.go, but dev already had it in engine_eval.go
- Resolution: kept dev's engine.go (ends at line ~700), updated engine_eval.go with hotfix guardrail logic
- Added `snapshotStarterSizes()` call after `CopyStarterFiles()` in engine_eval.go (line 125)
- Replaced old guardrail logic (lines 422-448) with starter-aware helpers:
  - `computeAgentFileCount()` for file count guardrail
  - `computeAgentOutputSize()` for output size guardrail
- New files from hotfix (guardrail.go, guardrail_test.go) merged cleanly

**Testing:** Full test suite passed with race detector (`go test -race ./... -timeout 3m`). All 15 guardrail test cases pass, including zero-byte edge cases.

**Key Learnings:**
1. **File splits require careful merge attention**: When one branch splits a file and another modifies it, auto-merge may fail. Solution: understand the split intent, keep the split structure, port changes to the correct file.
2. **Guardrail helpers are testable**: Extracting pure functions (snapshotStarterSizes, computeAgentOutputSize, computeAgentFileCount) to guardrail.go made the logic directly unit-testable (15 table-driven cases).
3. **Phase 3 now includes hotfix**: The dev branch has both Phase 3 features AND the starter-aware guardrail fix. Future merges to main will include both.

## 2026-04-16: Cross-Agent Update — Remote Skill Bug Flagged

**From:** Tank 📡 (PR #573)  
**Relevance:** Config/tool territory

Tank discovered bug during remote skill example work: `internal/skills/fetcher.go::fetchRemote` shells out to `npx skills add` without `--yes`, causing interactive prompt to block under non-TTY and yield 0 skills selected (repo clones fine, but manifest resolution fails).

**Impact on your work:**
- **#566 (WorkspaceDelta):** Not directly blocked; delta is independent of skill fetching.
- **Phase 4 config/tool work (#355–#357):** If you touch `FetchRemote` or skill fetcher, this is a known issue worth fixing: add `--yes` to `npx skills add` invocation.
- **Real remote-skill usage:** Will fail in CI without fix.

**Decision captured:** `.squad/decisions.md` → "Where example configs live + how to invoke them" (Tank decision, Tank caveat section).

