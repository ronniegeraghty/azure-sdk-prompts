# Project Context

- **Project:** hyoka — Go evaluation tool that runs prompts through the Copilot SDK, reviews code via a multi-model panel, produces criteria-based pass/fail reports
- **Stack:** Go 1.26.1+, GitHub Copilot CLI/SDK, MCP servers (Azure MCP via npx)
- **User:** Ronnie Geraghty
- **Created:** 2026-04-03
- **Repo:** /home/rgeraghty/projects/hyoka
- **Key paths:** hyoka/ (Go tool), prompts/ (evaluation prompts), criteria/ (pass/fail criteria), configs/ (run configs), reports/ (output), templates/, site/

## Core Context

Agent Morpheus initialized as Lead/Architect for hyoka. The project has guardrails for generation (turn count, file count, output size, session actions) and safety boundaries preventing real Azure resource provisioning by default. CLI supports list, run with filters (--service, --language), and smart path detection.

## Recent Updates

📌 Team initialized on 2026-04-03

## Learnings

**Key patterns learned across audits:** Factory pattern > singleton for task-specific configuration. Closure-based factories can share resources while creating per-task instances. Config normalization handles legacy→new format migration elegantly. Two-stage shutdown (SIGTERM→wait→SIGKILL) for clean process lifecycle.

### Comprehensive Audit Summaries (2026-07-14, 2026-10-14) — ARCHIVED

**July 2014 baseline:** ~20K lines Go / 21 packages / 1329-line main.go monolith / 264 test functions / 87 prompts / 8 configs. Clean eval→review→report pipeline. Reviewer factory pattern implemented (PR #170). 10 open issues. Untested: pidfile, clean/kill, serve handlers, trends.

**October hardening delta:** Zero structural changes. Reviewer model bug still present (main.go:469-473). No CI pipeline found (.github/workflows/ only has squad orchestration). Config validation gap: empty Generator.Model, duplicate shadowing. Serve.go:171 runID not validated (low risk). Error handling strong except reviewer.go:352, copilot.go:83. Test quality high but gaps remain (264 tests, flaky time.Sleep in resourcemonitor_test.go). Decision: CI is the single biggest hardening priority.

### React SPA Embedding Architecture (2026-04-07)

**Context:** User reported blank page when running `hyoka serve` without pre-built React site. Site build output (`site/dist/`, ~1.3 MB) is gitignored. Current serve command probes filesystem for `site/dist/` and fails with plaintext error if missing. Breaks `go install` users who get binary only.

### Detailed Session Summaries (2026-04-07 through 2026-10-16) — ARCHIVED

**2026-04-07 React SPA Embedding:** Analyzed embedding site using `go:embed`. Recommendation: Embed in binary, commit to repo, add CI build step. Follow Waza's proven pattern. Key principle: CLI tools must not depend on npm/node in production.

**2026-10-14 Comprehensive Audit & Evolution Plan:** ~20K lines Go / 21 packages / 1329-line main.go monolith. Created comprehensive 5-phase evolution plan (40+ tasks) incorporating Ronnie's 15 requirements. Critical dependency chain: CI → Generic Properties → Everything else. Recommended 14 project-specific skills (highest-priority: eval-pipeline, error-handling, contributor-guide).

**2026-10-15 Anchoring Bias Review:** Found 3 significant biases: Review system (should adopt Waza's pluggable grader architecture), config legacy fields (big-bang migrate), Prompt Properties (make sole source of truth). Ronnie approved all 3 pivots. Updated all 5 plan documents accordingly.

**2026-10-16 Session Limits Architecture:** 4-layer resolution — prompt > config YAML > CLI flag > engine default. Established with #284 (Neo assigned). Optional difficulty-based auto-scaling (basic=30, intermediate=50, advanced=100).

## 2026-04-16 — Phase 3 Merged to Dev (Neo)

Neo completed Phase 3 merge sequence: main→dev (hotfix #567 integrated), dev→Phase3 (clean), Phase3→dev (PR #562 squash-merged). Dev branch now has both Phase 3 features and starter-aware guardrail fix. All tests pass, CI green.

## 2026-04-17 — Phase 4 Kickoff Brief

**Brief:** `.squad/decisions/inbox/morpheus-phase4-kickoff.md`

### Dependency insights surfaced

- **Critical path is Neo-gated:** #566 → #355/#356 → #357 feeds into Trinity's #361 (serve comparison endpoints). Trinity's #358 (Eval Detail) is independently startable against Phase 3 data, but #359 is hard-blocked on #358.
- **`ComparisonResult` is the key shared boundary.** Neo owns the type definition in `internal/comparison/`, Trinity consumes it in `internal/serve/`. Misalignment here blocks site comparison features. Called this out as a risk with an explicit interface review gate.
- **`report/summary_stats.go` is dead code walking.** #357 should delete it and consolidate into `internal/comparison/`. Currently two code paths can produce different comparison results — a correctness bug waiting to happen.
- **Oracle's #363 has zero blockers.** All prereqs (#353, #344) shipped in Phase 3. Cleanest start of any agent.

### #566 WorkspaceDelta decision

**Option A: Neo does #566 first, before Phase 4 proper.** Rationale: it adds fields to `EvalReport` and `GraderInput` — the core types everything else reads from. Landing it first stabilizes the schema before 8 other issues touch surrounding code. 2-day time-box with fallback to ship type+wiring only if guardrail softening is complex.

### Phase 3 observations affecting Phase 4

- Phase 3 shipped 98 files changed, +4804/-7851 lines — a massive refactor. The unified grading pipeline (#344) is the foundation everything in Phase 4 builds on.
- `workspace.go` and `workspace_test.go` already exist in `hyoka/internal/eval/` from the hotfix path. #566 should decide: promote to `internal/workspace/` or extend in place. I recommended the new package for separation of concerns.
- The site structure uses React Router with component-per-page in `site/src/app/components/`. Trinity's #358 redesign should establish reusable component patterns (e.g., `GraderResultRow`) that #359 and #360 can inherit — called this out in the launch plan.
