# Squad Discussion — human-review.md Analysis

---

## Oracle 🔮 — Docs & Quality

### 1. Work Item Grouping

I've sorted the documentation-related review notes into logical work items, ordered by dependency and impact.

#### WI-1: Skill Cleanup (R1, R3, R5, R6, R7, R8, R9, R10, R11, R12, R13)
The single largest batch — 8 skills to remove, 1 to update, 1 to fill, 2 to consolidate. See §4 below for sequencing.

#### WI-2: `.copilot/` Directory Removal (R18)
Confirmed: `.copilot/skills/` contains 26 skills that are copies of `.agents/skills/`. The MCP config is redundant. Entire directory can be deleted. **This should happen AFTER WI-1** to avoid doing cleanup in both locations.

#### WI-3: Contributing Guide Consolidation (R4, R26)
Three locations mention contributing info:
- `.agents/skills/contributor-guide/` (212 lines, detailed)
- `docs/contributing.md` (119 lines, build/test focused)
- No root `CONTRIBUTING.md` exists yet

Action: Create root `CONTRIBUTING.md` by merging content from both sources, delete the skill, update `docs/contributing.md` to redirect or become the detailed dev guide. Fix `ronniegeraghty/dev` vs `main` branch discrepancy.

#### WI-4: AGENTS.md Overhaul (R23)
9 specific sub-items covering every section. The core tension: AGENTS.md tries to be both a reference doc AND a teaching tool. Proposed approach: slim it to pointers (link to docs/, link to skills) with only truly stable content inline.

#### WI-5: README.md Restructure (R24)
Structural overhaul of the 540-line README. Verified issues: `hyoka tools list`/`hyoka tools add` commands don't exist (it's `hyoka plugins`), repo structure tree is stale, `--max-session-actions` duplicated. This is a significant rewrite — not a patch job.

#### WI-6: `hyoka/README.md` Refocus (R40)
Confirmed: 451-line file is a stale copy of root README. Needs full rewrite as a developer-oriented doc covering package architecture, how to add commands/graders, debugging tips.

#### WI-7: Docs Cleanup (R25, R27, R28)
- Remove `docs/cleanup-plan.md` (372 lines, planning artifact)
- Remove `docs/eval-tool-plan.md` (980 lines, planning artifact)
- Remove `docs/tool-registry.md` (142 lines, dead feature — see WI-9)
- Document `plugins/` in `docs/configuration.md`
- Full docs audit (every doc validated against current CLI behavior)

#### WI-8: Examples Update (R30, R31, R33, R41, R83)
- Update `examples/configs/example-full.yaml` (uses outdated `mcp_servers:` key)
- Remove `examples/configs/example-registry.yaml` (dead feature)
- Update `examples/prompts/prompt-template.prompt.md` (uses old flat frontmatter)
- Add `.prompt.yaml` example
- Add prompt-with-starter-files example

#### WI-9: Dead Code / Dead Feature Docs (R32, R87, R138, R139)
- Remove `docs/tool-registry.md` (no configs use registry)
- Remove `hyoka/internal/tools/` package (confirmed: zero imports)
- Remove `hyoka/internal/history/` package (confirmed: zero imports)
- Remove `hyoka/internal/manifest/` package (confirmed: zero imports)
- Remove `hyoka/internal/build/` package (confirmed: zero imports, line 119)

#### WI-10: Rubric Removal (R19)
`hyoka/internal/review/rubric.go` (50 lines) embeds `rubric.md` (35 lines) via `//go:embed`. This is actively compiled into the binary. Removal requires verifying all call sites of `embeddedRubric` and `GetDefaultRubric()` and replacing them with prompt-specific/criteria-based evaluation. **This is code work, not just docs** — needs coordination with whoever owns the review package.

#### WI-11: Criteria Cleanup (R14, R15, R17)
- Remove `criteria/language/dotnet.yaml`, `go.yaml`, `java.yaml`, `rust.yaml` and `criteria/service/storage.yaml`
- Investigate why criteria files aren't being used in evals (this is a bug, not just docs)

#### WI-12: Validation Improvements (R64, R66, R67, R137)
See §5 below for approach.

#### WI-13: CLI Help & Flag Audit (R43, R44, R45, R46, R47, R48)
Audit all `--help` text, remove/add specific flags. Docs follow code here — do the flag changes first, then update docs.

---

### 2. Contradictions & Tensions

#### Contributing info — three-way split
- Line 5 says move `.agents/skills/contributor-guide/` → root `CONTRIBUTING.md`
- Line 51 says move `docs/contributing.md` → root `CONTRIBUTING.md`
- Both say "consolidate" but neither acknowledges the other exists

**Resolution:** These are complementary, not contradictory. The skill has agent-focused guidance (git workflow, branch naming, PR process). The doc has build/test commands. Both feed into one root `CONTRIBUTING.md`. The skill should be deleted (it's teaching agents repo-specific process that belongs in AGENTS.md or CONTRIBUTING.md, not a skill). Keep `docs/contributing.md` as a redirect or merge entirely.

#### Skills: remove vs update vs fill
This looks contradictory but isn't — each skill was assessed individually:
- **Remove** (8): `ci-validation-gates`, `config-system`, `copilot-sdk-integration`, `process-lifecycle`, `property-migration`, `report-generation`, `serve-patterns`, `grader-system` — all too implementation-specific for a fast-changing codebase
- **Update** (1): `eval-pipeline` — concept is stable, details are wrong
- **Fill** (1): `project-conventions` — confirmed it's a template with placeholders
- **Consolidate** (2): `prompt-authoring` + `prompt-conventions` → merge

No conflict here. The principle is clear: skills should teach stable patterns, not document volatile implementation details.

#### AGENTS.md: remove detail vs keep detail
Line 25 says either remove the repo structure section or commit to maintaining it. Line 27 says replace hardcoded config tables with pointers to docs. But line 30 says replace inline conventions with pointers to skills. This creates a potential circular reference problem: AGENTS.md → docs → skills → AGENTS.md?

**Resolution:** Establish a clear hierarchy: AGENTS.md is the agent entry point (minimal, links out). Skills teach patterns. Docs are the authoritative reference. No circular links — each level only points down, never back up.

#### Docs on site vs docs in repo
Line 315 says remove developer docs from the site (architecture, contributing). But line 52 says Oracle should audit all docs. If we split user docs (site) from dev docs (repo), the audit scope changes.

**Resolution:** Audit all docs regardless of where they live. The site should show user-facing docs; dev docs stay in repo as markdown. The audit covers both.

---

### 3. Investigation Items — Ongoing Responsibilities

The review assigns me four ongoing review responsibilities. Here's my assessment of current state and audit plan.

#### AGENTS.md (R23)
- **Current state:** 217 lines, 9 sections. Hardcoded absolute paths, stale config table, inline conventions that duplicate skills. Missing packages (`pidfile/`, `site/`). Hardcoded username `ronniegeraghty`.
- **Scope:** Medium rewrite. ~60% of content needs changing. The structure is sound; the content is stale.
- **Audit plan:** Rewrite after WI-1 (skill cleanup) completes so I can reference the correct skills. Validate every command/path in the file against current CLI output.

#### README.md (R24)
- **Current state:** 540 lines. Documents nonexistent commands (`hyoka tools list/add`). Duplicate flag entries. Stale repo tree. Roadmap references things that already exist.
- **Scope:** Large rewrite to proposed 6-section structure. Requires knowing final CLI command surface (depends on WI-13 flag cleanup).
- **Audit plan:** Wait for flag removals/additions (lines 69–73) to settle, then rewrite. Verify every command example runs successfully.

#### Full Docs Audit (R27)
- **Current state:** 13 docs, 3,542 total lines. At least 3 are dead (`cleanup-plan.md`, `eval-tool-plan.md`, `tool-registry.md`). `grader-config-schema.md` is DRAFT with 2 unimplemented grader types. `configuration.md` doesn't mention plugins. `cli-reference.md` needs verification against actual `--help` output.
- **Scope:** Large. Every doc needs line-by-line validation against current behavior.
- **Audit plan:**
  1. Delete dead docs first (3 files)
  2. Run `hyoka --help` and every subcommand `--help`, diff against `cli-reference.md`
  3. Validate `configuration.md` against actual config parsing code
  4. Validate `grader-config-schema.md` against implemented graders
  5. Validate `getting-started.md` by following it step-by-step
  6. Review `architecture.md`, `guardrails.md`, `prompt-authoring.md`, `starter-files.md`, `tool-filter-schema.md` for accuracy
  7. Add `plugins/` documentation to `configuration.md`

#### `hyoka/README.md` (R40)
- **Current state:** 451 lines, stale copy of root README. Not fit for purpose — should be a dev guide.
- **Scope:** Full rewrite. Need to understand package architecture to write this well.
- **Audit plan:** Write after dead code removal (WI-9) so the package list is accurate. Cover: package map, how to add a command, how to add a grader, debug tips, test patterns.

**Proposed cadence for ongoing reviews:** After any PR that adds/removes packages, changes CLI flags, modifies config schema, or changes the evaluation pipeline, I'll review and update AGENTS.md + README.md + relevant docs. This should be part of PR review checklist, not a separate async task.

---

### 4. Skill Cleanup Plan — Sequencing

**Phase 1: Safe removals (no dependencies, parallel)**
These skills have no dependents and can be removed in any order:
1. `ci-validation-gates` — Squad-product-specific, not hyoka
2. `config-system` — outdated structs
3. `copilot-sdk-integration` — pseudo-code doesn't match real SDK
4. `process-lifecycle` — stale pseudo-code
5. `property-migration` — references removed code
6. `report-generation` — implementation-specific
7. `serve-patterns` — references chi router (hyoka uses stdlib)
8. `grader-system` — implementation-specific

All 8 can be done in a single commit.

**Phase 2: Update `eval-pipeline` (after Phase 1)**
Correct the pipeline description: Generate → Grade → Review → Report. Remove "Build Verification Phase" reference. Keep it high-level to avoid future staleness.

**Phase 3: Consolidate `prompt-authoring` + `prompt-conventions` (after Phase 2)**
Merge into one `prompt-authoring` skill. Ensure it covers:
- Both flat and `properties:` map frontmatter formats
- Correct tool filtering (`When` not `properties` on ToolEntry)
- `## Evaluation Criteria` section format
- ID naming conventions

Delete `prompt-conventions` after merge.

**Phase 4: Fill `project-conventions` (after Phases 1–3)**
Currently a template with placeholders. Populate with stable conventions:
- Go stdlib preferred, `log/slog`, return errors with `%w`
- Table-driven tests with `-race`
- Cobra CLI with kebab-case flags
- Build/test commands
- Git workflow

Do this last because it should reference the surviving skills (error-handling, logging-conventions, cli-patterns, testing-patterns) and avoid duplicating them.

**Phase 5: Delete `.copilot/` directory**
After all `.agents/skills/` changes are final, delete the entire `.copilot/` directory since it's a stale copy.

**Phase 6: Move `contributor-guide` → root `CONTRIBUTING.md`**
Create `CONTRIBUTING.md`, merge with `docs/contributing.md`, delete the skill.

---

### 5. Validation Improvements — Approach

Three validation enhancements are called for:

#### 5a. Schema-based validation (R137)
Currently `hyoka validate` checks prompts against hardcoded value lists. Moving to schema-based:
- Define prompt schema as a Go struct with validation tags (or a JSON Schema document)
- Define criteria YAML schema similarly
- Define config YAML schema
- Validation becomes: parse YAML → validate against schema → report errors
- Benefit: format changes only require schema updates, not validation code changes

**Approach:** Start with Go struct validation tags (`validate:"required,oneof=..."`) using a lightweight validator. Avoid adding a JSON Schema dependency unless we need external schema files for editor integration.

#### 5b. Criteria validation (R66)
`hyoka validate` currently skips criteria files entirely. Add:
- Parse all `criteria/**/*.yaml` files
- Validate structure (must have `when`, `graders` list)
- Validate each grader entry has required fields for its `kind`
- Cross-reference: warn if criteria `when` conditions don't match any existing prompt properties

#### 5c. Prompt format detection (R67)
Flag prompts using old flat frontmatter vs `properties:` map format:
- Detect by checking if top-level keys include `service`, `language`, `plane` (flat) vs `properties` (new)
- Report as a warning with migration suggestion
- Could add `--fix` flag to auto-migrate (the parser already handles both formats, so migration is straightforward)

**Sequencing:** 5c first (quick win, informs scope of 5a), then 5a (schema definitions), then 5b (criteria validation using same schema patterns).

---

### 6. Disagreements

#### Line 25 — "Remove repo structure entirely" option
I disagree with the "remove entirely" option for the repo structure in AGENTS.md. Agents need a mental model of the codebase to navigate it. Without the structure section, agents will waste tokens exploring with `find`/`ls`. **However**, I agree it shouldn't list every file. Keep it to top-level packages only (2 levels deep max), and auto-generate it with a script to prevent staleness.

#### Line 30 — "Replace inline conventions with pointers to skills"
Partially disagree. AGENTS.md should contain the most critical conventions inline (error handling pattern, logging choice, test approach) as a quick reference. Pointing to 4 separate skills for basic conventions adds unnecessary indirection for agents. Skills should elaborate; AGENTS.md should summarize.

#### Line 48 — "Remove roadmap"
Disagree with full removal. A high-level roadmap (3–5 bullet points of what's planned next) is valuable for contributors and agents deciding whether to implement something or wait. Move detailed roadmap to a separate doc, keep a one-liner summary in README.

#### Line 199 — "Remove old flat frontmatter backward compatibility"
I'd add a caveat: don't remove until `hyoka validate --fix` can auto-migrate. The review says "all 89 prompts already use `properties:` format" — if true, removal is safe. But we should verify this claim by running the validator first. If even one prompt uses flat format, removing backward compat breaks it silently.

#### Line 69 — "Remove `--max-sessions` flag"
Agree with the intent but want to flag: if we ever add reviewer parallelism (running multiple reviewer models concurrently), we'd want session limits separate from worker counts. Suggest deprecating with a warning message rather than hard-removing, to give users time to adjust configs.

---

### Summary — Priority Order

1. **Skill cleanup** (WI-1) — highest impact, unblocks everything
2. **Dead code removal** (WI-9) — shrinks codebase, simplifies audit
3. **`.copilot/` removal** (WI-2) — quick win after WI-1
4. **Contributing consolidation** (WI-3) — creates standard repo file
5. **Criteria cleanup** (WI-11) — small file changes + bug investigation
6. **Validation improvements** (WI-12) — builds confidence in remaining content
7. **Examples update** (WI-8) — ensures examples match reality
8. **Docs audit** (WI-7) — comprehensive but depends on code settling
9. **AGENTS.md overhaul** (WI-4) — depends on WI-1, WI-3, WI-7
10. **README restructure** (WI-5) — depends on WI-13 flag changes
11. **`hyoka/README.md` rewrite** (WI-6) — depends on WI-9
12. **Rubric removal** (WI-10) — code change, needs review system coordination
13. **CLI flag audit** (WI-13) — code changes, docs follow

---

## Tank 📡 — CLI & Config

### 1. Work Item Grouping

After reading all 323 review lines and investigating the relevant code, here's how I group the CLI/config-related notes into named work items:

---

#### WI-1: `flag-cleanup` (R44, R45, R46, R47, R48, R70) — Remove/fix deprecated CLI flags
**Review lines:** 69–73, 99

- **Remove `--max-sessions`** (L69): Confirmed redundant with `--workers`. Both exist in `runFlags` (L33–34 in run.go). `MaxSessions` is passed to engine but could default to `workers × 3` internally.
- **Remove `--model`** (L70): Exists at `run.go:81`. Overrides all config models at L204–211. Useful for dev testing but agreed it's an unexpected override path — models belong in config YAML.
- **Remove `--skip-tests`** (L71): Exists at `run.go:83`, passed as `SkipTests` to engine. Likely vestigial from when a build phase existed.
- **Remove `--stub` flag and `StubEvaluator`** (L72): Exists at `run.go:87`. `StubEvaluator` at `engine.go:58` just writes a single `stub_output.txt`. Not a proper test harness.
- **Re-add `--max-turns`** (L73): Missing from CLI flags but engine supports `MaxTurns` (default 25). Should be added alongside `--max-session-actions` and `--max-files` for consistency.
- **Fix `--no-analyze` pattern in `trends.go`** (L99): Confirmed fragile — wraps `RunE` with another `RunE` at L94–100. Standard Cobra pattern: use `PreRunE` + `cmd.Flags().Changed()`.

**Dependencies:** None — can be done independently.

---

#### WI-2: `check-env-fix` (R68, R69) — Fix check-env command
**Review lines:** 97–98

- **Exit code bug (L97):** CONFIRMED. `check_env.go:14` uses `Run:` (not `RunE:`), so the command always exits 0. The `checkenv.Run()` function at `checkenv.go:18` returns nothing — it just prints. Fix: change to `RunE:`, have `checkenv.Run()` return an error (or bool), exit 1 if required tools are missing.
- **Binary mismatch (L98):** CONFIRMED. `checkenv.go:189–194` checks for a `copilot` binary via `exec.Command("copilot", "--version")`. The Copilot SDK uses `CLIPath` on `ClientOptions` with its own discovery logic. These could find different binaries. Fix: use the SDK's own `copilot.NewClient().Start()` to verify the binary the evaluator will actually use.

**Dependencies:** None — can be done early and independently.

---

#### WI-3: `path-resolution-unify` (R52, R56, R58, R60, R64) — Consolidate path resolution
**Review lines:** 77, 81, 83, 85, 89

- **Scattered resolution (L77, L81, L85):** `run.go` resolves prompts/output at L161–163 (step 2), configs at L167–181 (step 3), and criteria passed as raw string. Timeout parsing at L273–277 happens mid-flow between evaluator setup and reviewer setup.
- **Validate config dir (L89):** `validate.go:49` derives config dir as `filepath.Join(filepath.Dir(promptsDir), "configs")` — hardcoded sibling assumption. Should use `resolveConfigDir()` from `helpers.go:58`.
- **Local plugin paths not resolved (L83):** `resolveConfigSkillDirs()` in `helpers.go:82–117` only resolves skills, not plugins. If a config references a local plugin with a relative path, and hyoka is invoked from a different directory, the plugin won't be found.
- **Proposal:** Create a single `resolveAllPaths()` at the top of `RunE` that resolves prompts, configs, criteria, plugins, and output dirs together. Move all flag validation (timeout parsing, byte size parsing) there too.

**Dependencies:** Should be done before WI-5 (resource loading refactor).

---

#### WI-4: `dead-code-removal` (R32, R87, R138, R139) — Remove dead packages and commands
**Review lines:** 57, 119, 204–205, 212–213

- **Remove `internal/tools/`** (L57): CONFIRMED dead code — `registry.go`, `remote.go` + tests. Zero imports across entire codebase. Also remove `examples/configs/example-registry.yaml` and `docs/tool-registry.md`.
- **Remove `internal/build/`** (L119): CONFIRMED dead code — zero imports. Engine never calls it.
- **Remove `internal/history/`** (L204, L212): CONFIRMED dead code — zero imports, no command registered.
- **Remove `internal/manifest/`** (L205, L213): CONFIRMED dead code — zero imports, `hyoka manifest` not in `root.go`.

**Dependencies:** None — pure cleanup, do early.

---

#### WI-5: `run-cmd-refactor` (R52, R57, R60, R61, R95) — Refactor run.go resource loading
**Review lines:** 77, 82, 85, 86, 127

- **Load resources together (L82):** Currently prompts load at L235, configs at L166–181, criteria is just a string passed to engine. Should be: resolve all paths → load all resources → apply filters → calculate eval matrix.
- **Move flag processing up (L85):** Session timeout parsing at L273 happens mid-flow. `parseByteSize` at L373. Both should be at the top.
- **Move pairwise expansion into engine (L127):** Currently in `run.go:214–224`. Should be in `Engine.Run()` alongside `expandGeneratorModels`.

**Dependencies:** WI-3 (path resolution) should land first.

---

#### WI-6: `list-enhance` (R62) — Enhance `hyoka list` to show configs and criteria
**Review line:** 87

- **Current state:** `list.go` only lists prompts. `configs.go` is a separate command.
- **Proposal:** Make `hyoka list` the one-stop discovery command. Show prompts (default), `--configs` to show configs, `--criteria` to show criteria, or `--all` for everything. Consider merging `hyoka configs` into `hyoka list --configs` and deprecating the standalone command.

**Dependencies:** None, but benefits from WI-3 (unified path resolution).

---

#### WI-7: `validate-enhance` (R64, R66, R67, R74, R75, R76) — Enhance validation command
**Review lines:** 89, 95, 96, 103–105

- **Fix config dir assumption (L89):** `validate.go:49` hardcodes `filepath.Dir(promptsDir) + "/configs"`. Should use `resolveConfigDir()`.
- **Add criteria validation (L95):** Currently validates only prompts and configs. Should also validate `criteria/*.yaml` files.
- **Add prompt format detection (L96):** Flag prompts using old flat frontmatter vs `properties:` map format.
- **Fix `new-prompt` format (L103):** `new_prompt.go:62` generates old flat format. Should generate `properties:` map format.
- **Fix `askFreeText()` bug (L104):** CONFIRMED. `new_prompt.go:98` uses `fmt.Scanln(&input)` which stops at whitespace. Multi-word descriptions get truncated. Fix: use `bufio.Scanner`.
- **Add `--prompts` flag to `new-prompt` (L105):** `new_prompt.go:26` uses `resolvePromptsDir(cmd)` but doesn't register its own `--prompts` flag.

**Dependencies:** WI-3 for path resolution fix.

---

#### WI-8: `init-enhance` (R86) — Enhance `hyoka init`
**Review lines:** 115–118

- **Add `plugins/` to subdirs (L116):** `init.go` creates configs, prompts, criteria, skills, reports. Missing plugins.
- **Add `--with-examples` flag (L117):** Copy starter prompt + minimal config.
- **Accept optional path argument (L118):** Currently hardcoded to CWD.

**Dependencies:** None.

---

#### WI-9: `version-build-time` (R85) — Use Go build-time version
**Review line:** 114

- **Current state:** `root.go:9` hardcodes `Version = "0.3.0"`.
- **Fix:** Use `runtime/debug.ReadBuildInfo()` with hardcoded value as fallback for `go run` dev builds. Also set via `-ldflags` in CI.

**Dependencies:** None.

---

#### WI-10: `plugins-to-tools-rename` (R78) — Rename `hyoka plugins` to `hyoka tools`
**Review line:** 107

- **Current state:** `plugins.go` already has `Aliases: []string{"tools"}` so `hyoka tools` works. But the primary name should be `tools` with `plugins` as the alias (reversed).
- Should show all tool types (skills, MCP, plugins) from `.hyoka/` and cache.

**Dependencies:** WI-4 (remove dead `internal/tools/` first to avoid name collision).

---

#### WI-11: `progress-redesign` (R51, R137, R141) — Redesign progress bar output
**Review lines:** 76, 208, 215–229

- Research simpler approaches. Current live mode uses complex ANSI escape code region tracking. The proposed section-based format (L215–229) is cleaner.
- Consider `github.com/schollz/progressbar` or a line-by-line approach.

**Dependencies:** Engine refactoring (WI-5) informs what phases to display.

---

#### WI-12: `report-cmd-rename` (R72, R73) — Rename or clarify `hyoka report`
**Review lines:** 101–102

- `rerender.go` currently uses `Use: "report [run-id]"`. Name is misleading — it re-renders, not generates.
- Consider renaming to `hyoka rerender` to match internal naming, or improve help text.

**Dependencies:** None.

---

#### WI-13: `help-text-audit` (R43) — Audit all CLI help text
**Review line:** 68

- Run `hyoka --help` and `hyoka <command> --help` for every command.
- Verify descriptions, flag names, defaults are accurate.
- Root command still says "Azure SDK Prompt Evaluation Tool" at `root.go:21` — should be generalized.

**Dependencies:** Do after flag cleanup (WI-1) lands.

---

#### WI-14: `guardrail-tiers` (R79) — Split guardrail limits into three tiers
**Review line:** 108

- Currently turns, actions, time are one flat set. Split into generator limits (larger), reviewer limits (smaller), overall per-config limits.
- Affects `runFlags`, engine options, and session config construction.

**Dependencies:** WI-5 (run.go refactor), review system redesign.

---

#### WI-15: `sdk-client-factory` (R61, R71) — Share Copilot SDK client creation
**Review lines:** 86, 100

- SDK clients created in 5 places with duplicated `ClientOptions` setup. Create a `newHyokaClient()` shared factory.
- `trends.go` (L100) creates its own independent client for AI analysis.

**Dependencies:** None — good foundational work for other items.

---

### 2. Contradictions and Tensions

#### A. Remove `--model` vs Model Override Capability
L70 says remove `--model`. But during dev/testing, overriding the model without editing config YAML is extremely useful for quick iteration. **Tension:** simplifying the flag surface vs developer ergonomics. **My take:** Keep it but mark it hidden (`cmd.Flags().MarkHidden("model")`). Power users can still discover it.

#### B. Remove `--stub` vs Testing Without SDK
L72 says remove `--stub` and `StubEvaluator`. But L67 asks to investigate `--allow-cloud` flag behavior. Without `--stub`, there's no way to test CLI plumbing without a working Copilot SDK connection. **Tension:** removing dead-ish code vs needing a local test path. **My take:** The current `StubEvaluator` is too minimal to be useful. Either remove it entirely OR build a proper mock evaluator that replays recorded sessions — don't keep the half-implemented version.

#### C. Overlapping Path Resolution Notes
L77, L81, L83, L85 all say variations of "consolidate path resolution." These aren't contradictions but they overlap significantly. I've grouped them into WI-3 above. The key insight: all should become one `resolveAllPaths()` call.

#### D. `--max-turns` Re-add vs Flag Simplification
L73 wants `--max-turns` re-added, but L69–72 want flags removed. **Not a real contradiction** — `--max-turns` controls a meaningful guardrail (turn count) that has no other CLI path, while the removed flags are either redundant (`--max-sessions`) or bypass config (`--model`).

---

### 3. Investigation Findings

#### L54: Verify Local Plugin Support (R29)
**FINDING: Partially supported.** The code path exists and works:
1. `config/plugins.go:37–78` — `ExpandPlugins()` checks the plugin registry (local `plugins/` dir YAML files) first, then falls back to `~/.copilot/installed-plugins/`.
2. `config/plugins.go:122–129` — `resolvePluginsDir()` uses `.hyoka/plugins` → `./plugins` → `../plugins` candidate chain.
3. `cmd/plugins.go:18–51` — The `plugins` command loads from a `--plugins-dir` flag.
4. **Gap:** `resolveConfigSkillDirs()` in `helpers.go:82–117` resolves local skill paths but NOT local plugin paths. If a config references a local plugin with a relative path, and hyoka is invoked from a different directory, the plugin won't be found. Confirms the note at L83.
5. **Gap:** The `run` command calls `InstallSkillsAndPlugins()` which processes `Plugins` entries needing `npx skills add`. Local plugins from the `plugins/` dir are expanded via `ConfigFile.ExpandPlugins()` — but `run.go` never calls `ExpandPlugins()`. Local plugin YAML definitions in `plugins/` are only loaded inside `InstallSkillsAndPlugins` for skip-check purposes, not for actual expansion into tool entries during the eval.

#### L59: Verify Prompt Discovery Handles Deep Nesting (R34)
**FINDING: Works correctly.** `loader.go:27` uses `filepath.Walk(root, ...)` which recursively traverses all subdirectories with no depth limit. Any `.prompt.md`, `.prompt.yaml`, or `.prompt.yml` file at any nesting depth will be discovered. The `isPromptFile()` check at L13–16 only looks at the filename suffix, not the path structure. Confirmed by test at `loader_test.go:190` which creates prompts in a subdirectory.

#### L97: Fix `check-env` Exit Code (R68)
**FINDING: Confirmed broken.** `check_env.go:14` uses `Run:` callback (not `RunE:`). Cobra's `Run` doesn't propagate errors — the command always exits 0 regardless of check results. The `checkenv.Run()` function returns nothing; it just prints emoji-decorated output. Fix requires:
1. Change `check_env.go` from `Run:` to `RunE:`.
2. Change `checkenv.Run()` to return `error` (or at least `bool`).
3. Track whether any required checks failed, return error if so.

#### L98: Verify `check-env` Copilot CLI Binary Matches SDK (R69)
**FINDING: Confirmed mismatch risk.** `checkenv.go:189–194` runs `exec.Command("copilot", "--version")` — this finds whatever `copilot` binary is first on `$PATH`. The SDK uses `copilot.ClientOptions.CLIPath` (if set) or its own internal discovery. These could resolve to different binaries if multiple versions are installed. The evaluator at `copilot.go:89` calls `copilot.NewClient(&opts)` where `opts.CLIPath` is empty by default. Fix: `check-env` should instantiate a `copilot.NewClient()` and attempt `.Start()` to validate the SDK's own binary discovery path.

#### L67: `--allow-cloud` Flag Investigation (R42)
**FINDING: Flag is NOT wired through.** The `allowCloud` field is stored on `CopilotSDKEvaluator` (copilot.go:27) and set from options (L75). But searching `buildSessionConfig()` (L665–815), the `allowCloud` field is **never referenced**. The system prompt construction at L680–695 only adds skills hints — there's no conditional safety boundary text based on `allowCloud`. The field is stored but never used to modify the session config or system prompt. Reports saying "cloud disabled" are correct — the flag literally doesn't do anything.

---

### 4. Dependencies (Execution Order)

```
Phase 1 (no deps — do first):
  WI-4  (dead-code-removal)
  WI-2  (check-env-fix)
  WI-9  (version-build-time)
  WI-8  (init-enhance)
  WI-15 (sdk-client-factory)
  WI-1  (flag-cleanup)
  WI-12 (report-cmd-rename)

Phase 2 (path/loading foundation):
  WI-3  (path-resolution-unify)

Phase 3 (builds on phase 2):
  WI-5  (run-cmd-refactor)
  WI-6  (list-enhance)
  WI-7  (validate-enhance)
  WI-10 (plugins-to-tools-rename) — also after WI-4

Phase 4 (builds on phase 3):
  WI-11 (progress-redesign) — after WI-5
  WI-13 (help-text-audit) — after WI-1
  WI-14 (guardrail-tiers) — after WI-5 + review redesign
```

---

### 5. Disagreements

#### L70: Remove `--model` — Partial Disagreement
I'd keep it but hide it. During development and debugging, being able to test `go run ./hyoka run --prompt-id X --config Y --model claude-sonnet-4` without editing YAML is invaluable. Hidden flags don't clutter `--help` but are still available. Removal is a DX regression for developers.

#### L102: Switch `report` positional arg to `--run-id` — Disagree
Positional args for primary identifiers are idiomatic in CLIs (`git show <sha>`, `docker logs <container>`). `hyoka report <run-id>` reads naturally. The kebab-case convention from `cli-patterns` applies to flags, not args. Adding `--run-id` makes the command more verbose for zero benefit. Keep as-is but improve help text.

#### L106: Make `new-prompt` directory structure configurable — Low Priority / Scope Creep
The current `{service}/{plane}/{language}/` structure is fine for now. A `--flat` flag adds complexity for a scenario that doesn't exist yet (other users with different layouts). I'd defer this until someone actually needs it. Keep `new-prompt` simple.

#### L69: Remove `--max-sessions` — Agree With Caveat
Agree it should be removed — but verify the engine's default `workers × 3` formula works well empirically first. If reviewers don't run in parallel (they don't), then sessions = workers is correct. The `× 3` multiplier may be overprovisioning. Investigate before removing.

---

## Trinity 🖤 — Frontend & Reports

### 1. Work Item Grouping

I've grouped the review notes touching my domain into **six logical work items**, ordered from foundational to incremental.

#### WI-1: Dead Code Purge (R32, R38, R138, R139) — site + backend packages
*Lines: 63, 57, 204, 205, 212, 213*

**Investigation result:** I audited every component in `site/src/`. Of the 45 shadcn/ui components, **43 are completely unused** — only `select.tsx` and `table.tsx` are imported outside of `ui/`. One custom component (`ImageWithFallback.tsx` in `figma/`) is also dead. All 12 page components are routed, so no dead pages. This is a quick, high-confidence cleanup — delete the 43 unused UI files, delete `ImageWithFallback.tsx`, and remove the backend dead code (`internal/tools/`, `internal/history/`, `internal/manifest/`).

#### WI-2: Eliminate Static HTML Reports + Trends HTML (R131, R134, R135, R148) — Consolidate on SPA
*Lines: 175–181, 192–197, 276–277*

**Investigation result:** `report/html.go` (744 lines) + two `.gohtml` templates generate standalone HTML files. The SPA at `hyoka serve` renders the exact same data from JSON via `/api/runs/{runId}` and `/api/runs/{runId}/eval`. The static HTML reports add zero functionality the SPA doesn't have — they're a parallel rendering pipeline with a maintenance cost.

The trends package is worse: `trends.go` has ~280 lines of raw `fmt.Fprintf` / `WriteString` calls building inline HTML with embedded CSS and injected `marked.js`. Meanwhile, the serve package already exposes `/api/trends` returning rich structured JSON (`TrendReport` with `PromptTrend`, `RunResult`, `GraderScores`). The site can consume this directly.

**Recommendation:**
- Remove `report/html.go`, both `.gohtml` templates, and the `writeHTMLReport` function in `trends.go`.
- Keep `report/markdown.go` as the portable offline fallback (CI artifacts, sharing).
- Keep JSON as the canonical data format. The SPA is the viewer.
- This aligns with the review note on line 197: *"reports are data (JSON), the site is the viewer, markdown is the portable fallback, static HTML goes away."*

#### WI-3: Eval Detail Page Fix + Redesign (R148)
*Lines: 257–276*

This is the most impactful single page change. The eval detail page currently:
- **Has a routing bug** (line 258): clicking an eval from run detail produces a 404 — URL format mismatch.
- Shows old general rubric criteria ("Code Builds", "Latest Package Versions") instead of actual grader results (line 275).
- Has an "AI Consolidated Review" section using the consolidator we're removing (line 276).

Priority order within this WI:
1. Fix the 404 routing bug (unblocks all other work on this page)
2. Replace rubric criteria badges with grader results table
3. Add summary/insights agent output section
4. Split environment into Env block + Run stats cards
5. Redesign reviewer timeline for new grader sessions
6. Make generated files expandable with inline content

#### WI-4: Run Detail + Runs Page Improvements (R146, R147)
*Lines: 244–256*

- Convert timestamps to human-readable run names
- Replace status column with color-coded score column
- Show full prompt IDs (don't truncate)
- Replace config name with model + tools tags inline representation
- Make rows clickable (full-row navigation)
- Remove "Files" summary card (not relevant for non-code-gen)
- Add prompt × config vs eval criteria matrix table

#### WI-5: Prompt Pages + Dashboard Rethink (R150, R151, R154)
*Lines: 278–297, 308*

- Prompts page: add "only show prompts with evals" filter (default on), ordering options, remove non-functional dropdowns
- Prompt detail: fix score trend chart x-axis to days, show ALL models not top 3, add tool usage toggle
- Dashboard: currently shows 100% mock data (hardcoded "1,247 evaluations", fake model names). Must be connected to real report data. Radar chart won't work without default criteria — needs rethinking.

#### WI-6: Compare Page Redesign + Docs Page Improvements (R153, R155)
*Lines: 302–307, 309–322*

- Compare page: replace simple A vs B with configurable group-based comparison (referencing devex-reviews pattern)
- Docs page: add sidebar groupings (Getting Started, CLI Reference, Configuration, Concepts), fix "Get Started" button link, remove dead docs from sidebar

---

### 2. Contradictions & Tensions

#### Static HTML vs SPA — resolved, not contradictory
Lines 178 and 197 both raise this question, and the answer is consistent across the review notes: **eliminate static HTML, keep SPA + Markdown**. The review note on line 197 is explicit: *"static HTML goes away."* I agree. The only tension is in `getting-started.md` which currently documents `open reports/<run-id>/summary.html` as the primary workflow — that doc will need updating when HTML generation is removed.

#### "Eliminate HTML generation" vs "Keep Markdown"
No actual conflict. These are complementary — HTML is the redundant rendering, Markdown is the portable fallback. JSON remains canonical.

#### Dashboard page purpose
Line 308 asks whether the dashboard should be a "global overview" or "consolidated into runs/compare pages." My take: **keep it as a landing page with real aggregate data**, but only after the grader system is finalized. The current mock data makes it misleading. Priority should be: fix eval detail (WI-3) → fix runs (WI-4) → then dashboard once the data model stabilizes.

#### Report tools "configured" vs "loaded" (line 156) vs site tool display (lines 260, 294)
Slight tension: line 156 says report tools as "loaded" once SDK events confirm it, but lines 260 and 294 want the site to show tools as "available vs actually used." These are complementary — the report data model needs three states: configured → loaded → used. The site renders from that.

---

### 3. Investigation Findings

#### Line 63: Site dead code audit (R38) — 43 of 45 shadcn/ui components unused
**Done.** Full audit results above in WI-1. The 43 unused components total ~8,500 lines of dead code. Only `select.tsx` and `table.tsx` survive. This is expected — shadcn/ui encourages installing components individually, but someone ran a bulk install. Safe to delete; re-add as needed.

#### Line 176: Two `BuildActionTimeline()` implementations (R131)
**Done.** These are **intentionally separate, not duplicates:**

| Aspect | `eval/action.go` | `report/types.go` |
|--------|------------------|-------------------|
| Returns | `*ActionTimeline` | `*ActionTimelineReport` |
| Event pairing | Flat list, no pairing | Pairs `tool.execution_start` + `tool.execution_complete` into merged entries |
| Session setup | Not handled | Enriched with setup metadata (skills/MCP loaded) |
| Purpose | Internal eval processing / grader input | Structured report generation |
| Callers | `eval/copilot.go` (3 calls) | `report/generator.go`, `serve/dashboard.go` |

**Verdict:** Keep both. The review note's concern is addressed — they serve different purposes. The naming overlap is confusing though; I'd suggest renaming the eval version to `ClassifyEvents()` to make the distinction obvious.

#### Line 178: HTML reports vs embedded SPA (R131) — are both needed?
**Done.** Full analysis above. Summary:
- Static HTML: 744 lines in `html.go` + 2 Go templates. Self-contained files, no server needed.
- SPA: Full React app with 12 pages, 6+ API endpoints for cross-run analysis, interactive timelines.
- **Overlap is ~90%** for individual eval/summary views.
- **SPA-only features:** dashboard analytics, comparisons, trends, prompt history, docs viewer.

**Verdict:** Remove static HTML. Keep JSON (data) + Markdown (portable) + SPA (interactive viewer). The offline use case is served by Markdown output.

#### Lines 192–197: Trends inline HTML + duplicate matchesProperties (R134, R135)
**Done.** The trends package has ~280 lines of inline HTML string building via `fmt.Fprintf`/`WriteString`. Meanwhile `/api/trends` already exists in `serve/dashboard.go` and returns structured `TrendReport` JSON. The `matchesProperties` function in `trends/slice.go` is character-for-character identical in logic to `matchesWhen` in `config/tool_filter.go` — only parameter names differ. This should use a shared utility.

---

### 4. Site Page Redesign Priority

Based on impact, dependencies, and the fact that the grader system is being overhauled, here's my recommended order:

| Priority | Page | Rationale |
|----------|------|-----------|
| **P0** | Eval Detail | Blocked by a **404 routing bug** — literally broken. Also the page where grader results, reviewer votes, and summary/insights all land. Must be redesigned to match the new grader system. |
| **P1** | Run Detail | Gateway page to eval details. Needs score column, clickable rows, criteria matrix. High visibility. |
| **P2** | Runs List | Timestamp formatting, duration clarity. Quick wins with high user impact. |
| **P3** | Prompts + Prompt Detail | Filtering, ordering, tool usage charts. Important for repeat users but less critical than fixing the core eval viewing flow. |
| **P4** | Pairwise | Already works well per the review. Add explanation tooltips and tool usage frequency chart. |
| **P5** | Compare | Full redesign to group-based comparison. Large effort, lower urgency — current A vs B works. |
| **P6** | Dashboard | **Block until grader system stabilizes.** Current mock data is misleading. Need real aggregate data, and the visualizations depend on what the finalized grader output looks like. |
| **P7** | How It Works | Content rewrite — depends on finalized pipeline description. Do after grader system is done. |
| **P8** | Homepage + Footer | Cosmetic text changes. Low effort, do anytime. |
| **P9** | Docs Page | Sidebar restructuring. Depends on Oracle's docs audit completing first. |

**Most impactful single change:** Fix the eval detail 404 + redesign that page for the new grader system. This is where users spend the most time understanding results, and it's currently broken.

---

### 5. Disagreements

#### I partially disagree with removing the "Files" summary card from run detail (line 253)
The note says "Remove 'Files' summary card — not relevant for non-code-gen prompts." True for non-code-gen, but **most current prompts ARE code-gen**. Rather than removing it, I'd make it **conditional** — show it when any eval in the run produced files, hide it when none did. Same principle applies to other code-gen-specific UI elements: conditionally show based on actual data, don't hardcode assumptions about prompt types.

#### I'd deprioritize the Compare page redesign (line 302–307)
The review envisions a full group-based comparison system inspired by devex-reviews. That's a significant engineering effort (group model, builder UI, localStorage persistence, configurable charts). The current simple A vs B comparison works. I'd rather invest that effort in getting eval detail and the grader results table right first — those are foundational. Compare page redesign should be a stretch goal after the core viewing experience is solid.

#### "Remove Pass Rate by Config bar chart" on prompt detail (line 291)
The note says it's "not useful with many pairwise variants." I think the chart is useful — it just needs **grouping by base config** with pairwise variants collapsed underneath. Removing it entirely loses a quick visual overview. Better to fix the visualization than remove it.

#### Dashboard shouldn't be P6 forever
The review note (line 308) questions whether the dashboard should exist or be folded into runs/compare. I think a dashboard landing page is valuable — users want a quick "how are things going" view. But it must show real data. Once the grader system is stable and we have a few runs of real data, dashboard should jump to P2 priority.

---



#### WI-7: Site Content & Branding Updates (R143, R144, R145, R149, R152)

These are content/copy changes that don't require structural code changes:
- **R143** — Design new logo (replace generic >_ with something representative)
- **R144** — Revise homepage content (remove code-gen/Azure-specific messaging)
- **R145** — Rewrite "How It Works" page (update pipeline stages, fix visual bugs, generalize language)
- **R149** — Update site footer (remove Azure SDK reference)
- **R152** — Pairwise page improvements (score clarity, tool usage frequency chart)

These can be done in parallel with the structural site work and should be tackled after the grader system redesign since the How It Works page needs to reflect the new pipeline.

## Neo 💊 — Core Eval Framework

### 1. Work Item Grouping

From my domain (eval engine, review panel, criteria, graders, SDK integration, skills/plugins), I propose these logical work items:

---

#### WI-1: Criteria & Grader Unification (R15, R17, R19, R37, R84, R92, R93, R94, R97, R100, R101, R102, R142)
**Notes:** 16, 18, 20, 62, 113, 124, 125, 126, 129, 132–136

This is the single most impactful body of work. The criteria system is *implemented but silently unused* (details in §3 below). The general rubric masks the problem. These notes all converge on one goal: **make the grader system the single source of truth for evaluation**.

Sub-items:
- **WI-1a: Fix criteria auto-discovery** — `--criteria-dir` defaults to `""` and `loadCriteria()` silently returns. Add a warning log, improve auto-discovery in `resolveCriteriaDir()`.
- **WI-1b: Remove general rubric** — Delete `rubric.go` + `rubric.md` (note 20). Force review panel to use only prompt-specific + attribute-matched criteria.
- **WI-1c: Automate criteria extraction from prompts** — Parse `## Evaluation Criteria` into individual grader entries (note 126). Each `-` bullet becomes a prompt grader.
- **WI-1d: Structured JSON responses from reviewers** — Enforce JSON schema for reviewer output, validate per-criterion coverage, retry on malformed responses (note 113, 230).
- **WI-1e: AI prompt grader session modes** — Add `--review-mode combined|isolated` and per-grader `isolate: true` (notes 132–136).
- **WI-1f: Hierarchical `when` on criteria** — Support `when` at file, group, and individual grader levels (note 125).
- **WI-1g: Unify criteria + graders into one pipeline** — Make the AI review panel one grader type alongside file/behavior/program/build graders (note 129).

---

#### WI-2: Tool System Consolidation (R49, R50, R53, R54, R55, R78, R80, R88, R89, R90, R91)
**Notes:** 74, 75, 78, 79, 80, 107, 109, 120, 121, 122, 123

All of these converge on: **one unified tool resolution → acquisition → caching → session-injection pipeline**.

Sub-items:
- **WI-2a: Extract `internal/tool/` package** — Centralize resolution of skills, MCP servers, and plugins (note 123).
- **WI-2b: Unify remote skill and plugin installation** — Single pre-flight step for all remote deps (note 74).
- **WI-2c: Isolate from user environment** — Use `~/.hyoka/cache/` instead of `~/.copilot/installed-plugins/` (note 75).
- **WI-2d: Per-session skill isolation** — Copy only declared skills into temp dir (note 121).
- **WI-2e: Add versioning** — Track installed versions, check for updates (note 79).
- **WI-2f: Build custom fetcher** — Replace `npx skills add` (note 80).
- **WI-2g: Add remote MCP server support** — Extend config struct + validation + session builder (note 122, findings in §3 below).
- **WI-2h: Local plugin path resolution** — Unify with skill path resolution (note 78).

---

#### WI-3: Engine Refactoring (R52, R95, R96, R103, R104, R106, R108, R115)
**Notes:** 77, 127, 128, 137, 138, 140, 142, 154

These are code quality improvements to `engine.go` and related orchestration:

- **WI-3a: Split `engine.go`** — Separate orchestrator (`Run()`) from single-eval lifecycle (`runSingleEval()`) (note 137).
- **WI-3b: Consolidate task expansion** — Move pairwise expansion into `Engine.Run()` (note 127).
- **WI-3c: Rename `CopilotEvaluator` → `PromptRunner`** — Better name for what it does (note 128).
- **WI-3d: Refactor `run.go` resource loading** — One resolution step for all paths/resources (note 77).
- **WI-3e: Extract report data population** — Centralize report field population (note 138).
- **WI-3f: Move workspace management out of eval flow** — Report workspace managed by report system (note 140).
- **WI-3g: Extract early-return error helper** — `evalReport.failSetup()` (note 142).
- **WI-3h: Generalize zero-file diagnostics** — Support non-code-gen prompts (note 154).

---

#### WI-4: Workspace Containment Hardening (R98, R109, R110, R111, R112)
**Notes:** 143–151

Security-critical work around sandboxing the agent workspace:

- **WI-4a: Fix bash command containment** — Currently bash is logged but NOT restricted. This is the main breach vector.
- **WI-4b: Remove home/CWD snapshot-and-recovery** — Fix root cause instead of scanning user's home dir (note 143). Privacy concern.
- **WI-4c: Investigate SDK built-in containment** — Test whether `WorkingDirectory` enforces containment (note 149, findings in §3 below).
- **WI-4d: Normalize paths** — Add `filepath.Abs()` to prevent `../` bypass (note 143).
- **WI-4e: Expand `isFileWriteTool` list** — Maintain dynamically or get from SDK (note 143).
- **WI-4f: Verify tool-loading events** — Listen for `mcp_server_status_changed`, `extensions_loaded`, fail fast on missing tools (note 130, findings in §3 below).

---

#### WI-5: `--allow-cloud` Fix (R42)
**Note:** 67

Two bugs found (details in §3 below):
- **WI-5a:** `buildSessionConfig()` never uses `e.allowCloud` to modify system prompt.
- **WI-5b:** `EnvironmentInfo.AllowCloud` is hardcoded to `false` in engine.go:981. Engine can't access evaluator's flag due to interface boundary.

---

#### WI-6: CLI Flag Cleanup (R44, R45, R46, R47, R48)
**Notes:** 69–73

- **WI-6a:** Remove `--max-sessions` (note 69).
- **WI-6b:** Remove `--model` (note 70). *(I disagree — see §5.)*
- **WI-6c:** Remove `--skip-tests` (note 71).
- **WI-6d:** Remove `--stub` and `StubEvaluator` (note 72).
- **WI-6e:** Re-add `--max-turns` (note 73).

---

#### WI-7: Dead Code Removal (R32, R87, R138, R139)
**Notes:** 57, 119, 204, 205

- Remove `internal/tools/` package (note 57) — dead ToolRegistry.
- Remove `internal/build/` package (note 119) — unused build verification.
- Remove `internal/history/` package (note 204) — no command registered.
- Remove `internal/manifest/` package (note 205) — no command registered.

---

#### WI-8: SDK Session & Cleanup (R71, R99, R107, R113)
**Notes:** 100, 131, 141, 152

- **WI-8a:** Simplify Copilot session cleanup — test if plain `Stop()` suffices (note 152).
- **WI-8b:** Build timeline from events during report generation, not during eval (note 131). *(I disagree — see §5.)*
- **WI-8c:** Verify temp workspace cleanup reliability (note 141).
- **WI-8d:** Share SDK client setup for trends.go (note 100).

---

#### WI-9: Guardrail System (R79, R105)
**Notes:** 108, 139

- **WI-9a:** Split guardrail limits into generator/reviewer tiers (note 108). *(I'd skip the overall tier — see §5.)*
- **WI-9b:** Document priority chain (prompt > config > CLI > default) in user-facing docs (note 139).
- **WI-9c:** Allow all guardrails to be overridden at all levels (note 139).

---

### 2. Contradictions and Tensions

#### Tension A: Rubric Removal vs. Working System
Note 20 says "remove `rubric.go` and `rubric.md`" but note 113 says "harden up the review criteria being asked." If we remove the rubric *before* fixing the criteria system (WI-1a), reviewers will have **zero criteria** for prompts that lack an `## Evaluation Criteria` section and where `--criteria-dir` isn't resolved. **Resolution needed:** WI-1a and WI-1c must be completed *before* WI-1b (rubric removal).

#### Tension B: `when` Filter Won't-Fix vs. New Tool System
Note 109 says `when` filtering for skills/MCP is a "won't fix" because the plan is to move everything under `tools` with `type: "tool"`. But note 120 says to generalize tool handling so plugins flatten to skills + MCP at the SDK level. These aren't contradictory but the end-state needs clarification: will there be one `tools` section with `type: tool|mcp|skill`, or will everything resolve to SDK primitives (SkillDirectories, MCPServers, AvailableTools)?

#### Tension C: Build Package — Remove vs. Reimagine
Note 119 says "remove `internal/build/` package — dead code" but also says "reimplemented as a proper build grader." The removal is fine, but note 129 wants to unify criteria and graders. The build grader design depends on whether we go with unified grading (note 129) or keep graders separate from criteria. **My recommendation:** Remove the dead code now, design the build grader as part of WI-1g (unified pipeline).

#### Tension D: Plugin Concept vs. Tool Unification
Note 107 says rename `hyoka plugins` → `hyoka tools`. Note 120 says plugins are "just a hyoka abstraction that gets flattened." Note 75 says isolate plugin installation. These all agree plugins should go away as a first-class concept, but the migration path matters. **My recommendation:** Keep plugins in config YAML as syntactic sugar (a named bundle of skills + MCP entries), but resolve them at load time into the unified tool system. `hyoka tools` lists the resolved view.

---

### 3. Investigation Findings

#### Investigation: Criteria System (R15, R17)

**Status: Confirmed broken in production, working in tests.**

Full trace from CLI to review prompt:

1. `run.go:97` — `--criteria-dir` defaults to `""`.
2. `helpers.go` — `resolveCriteriaDir()` checks `./criteria`, `../criteria`. Returns `""` if not found.
3. `run.go:405` — `EngineOptions{CriteriaDir: f.criteriaDir}` passes potentially empty string.
4. `engine.go:376` — `e.loadCriteria()` called in `Run()`.
5. `engine.go:196-211` — **Silent return when `CriteriaDir == ""`**. No warning logged.
6. `engine.go:1251` — `mergedCriteria()` calls `criteria.MatchingGraders(e.graderConfigs, props)` — returns nothing because `e.graderConfigs` is empty.
7. `review/rubric.go:15-48` — `BuildReviewPrompt()` always appends `embeddedRubric` (line 48) regardless of whether criteria were loaded.

**Root causes:**
- Empty default + silent failure = criteria never loaded in typical usage.
- General rubric always appended = problem is masked; reviews look like they work.
- Test `TestCriteriaMergedIntoReview` (engine_test.go:1043) passes because it explicitly sets `CriteriaDir`.

**Fix plan:**
1. Log a warning when `CriteriaDir == ""` and no auto-discovery succeeds.
2. Improve auto-discovery to check `.hyoka/criteria` and project-relative paths.
3. Once criteria loading is reliable, remove the general rubric.

---

#### Investigation: `--allow-cloud` Flag (R42)

**Status: Two confirmed bugs.**

**Bug 1 — System prompt not modified (`copilot.go:665-815`):**
`buildSessionConfig()` has access to `e.allowCloud` but never uses it. No safety boundary prompt is added when cloud is disallowed, and no restriction is removed when it's allowed. The flag is stored but has zero effect on session behavior.

**Bug 2 — Report hardcoded (`engine.go:975-983`):**
```go
SafetyBoundaries: true,   // hardcoded
AllowCloud:       false,  // hardcoded
```
The Engine uses the `CopilotEvaluator` interface which doesn't expose `allowCloud`. The flag can't flow from evaluator to report without either:
- Adding `AllowCloud` to `EngineOptions`, or
- Adding a method to the `CopilotEvaluator` interface.

**Fix plan:** Add `AllowCloud` to `EngineOptions`, pass from `run.go`, use in both `buildSessionConfig()` (to set system prompt) and `runSingleEval()` (to set report field).

---

#### Investigation: SDK Events for Tool Verification (R98)

**Status: Partially implemented — events captured but not used for verification.**

Currently captured in `copilot.go`:
- ✅ `SessionEventTypeSessionSkillsLoaded` (lines 301-309) — extracts skill names.
- ✅ `SessionEventTypeSessionMcpServersLoaded` (lines 310-318) — extracts server names.
- ✅ `SessionEventTypeSessionToolsUpdated` (lines 319-320) — logged.
- ❌ `mcp_server_status_changed` — **not handled**.
- ❌ `extensions_loaded` — **not handled**.

Events are stored in `SessionEventRecord` structs in the report but **no verification logic exists**. After `CreateSession()`, the code doesn't check whether all requested tools actually loaded.

**Fix plan:**
1. Add case handlers for the missing event types.
2. After `CreateSession()`, wait for `skills_loaded` and `mcp_servers_loaded` events.
3. Compare loaded tool names against requested tools from config.
4. Fail fast with actionable error if any tool failed to load.

---

#### Investigation: SDK Workspace Containment (R110)

**Status: Unknown — needs runtime testing.**

Current configuration in `copilot.go`:
- `WorkingDirectory: workDir` (line 700) — set to isolated temp workspace.
- `OnPermissionRequest: copilot.PermissionHandler.ApproveAll` (line 701) — approves everything.

The SDK docs say `WorkingDirectory` makes "tool operations relative to this directory" and `OnPermissionRequest` defaults to deny-all if nil. However:
- We override with `ApproveAll`, which may bypass SDK containment.
- No runtime test has confirmed whether the SDK enforces path restrictions when `WorkingDirectory` is set.
- Bash/shell commands bypass the `PreToolUseHook` entirely (logged but not restricted).

**Fix plan:** Run a test eval with debug logging to observe whether SDK-level containment catches writes outside `WorkingDirectory`. If it does, we can simplify `PreToolUseHook`. If not, we need to harden the hook for bash commands.

---

#### Investigation: Remote MCP Server Support (R90)

**Status: Not supported — three blockers identified.**

1. **Config struct** (`tool_filter.go:9-24`) — `ToolEntry` only has `Command` + `Args` for MCP. No fields for URL, socket path, or server type.
2. **Validation** (`tool_filter.go:109-112`) — Requires `Command` field for all MCP entries. Remote servers don't have a command.
3. **Session builder** (`copilot.go:784`) — Hardcodes `"type": "local"` for all MCP servers.

The SDK's `MCPServerConfig` is `map[string]interface{}` so it can accept any key-value pairs including remote server configs. The limitation is entirely in hyoka's config layer.

**Fix plan:**
1. Add `ServerType` field to `ToolEntry` (`"local"` default, `"http"`, `"socket"`).
2. Add `URL` and/or `SocketPath` fields.
3. Update validation: require `Command` only for `local` type.
4. Update `buildSessionConfig()`: pass through the configured type instead of hardcoding.

---

### 4. Dependencies & Critical Path

```
WI-7 (Dead Code Removal)         ← No deps, do first for clean baseline
    │
    ▼
WI-1a (Fix criteria auto-discovery)
    │
    ▼
WI-1c (Automate criteria extraction from prompts)
    │
    ▼
WI-1b (Remove general rubric)    ← MUST come after 1a + 1c
    │
    ▼
WI-1d (Structured JSON responses from reviewers)
    │
    ▼
WI-1e–1g (Session modes, hierarchical when, unified pipeline)
```

```
WI-2a (Extract internal/tool/ package) ← Foundation for all tool work
    │
    ├─▶ WI-2b (Unify remote skill/plugin install)
    ├─▶ WI-2c (Isolate from user env)
    ├─▶ WI-2g (Remote MCP support)
    │
    ▼
WI-2d (Per-session skill isolation)  ← Depends on 2a
    │
    ▼
WI-2e,2f (Versioning, custom fetcher)  ← Nice-to-have, later
```

```
WI-5 (--allow-cloud fix)         ← Standalone, can be done anytime
WI-6 (CLI flag cleanup)          ← Standalone, can be done anytime
WI-4 (Workspace containment)     ← Standalone but WI-4c informs 4a/4b
WI-3 (Engine refactoring)        ← After WI-1 to avoid merge conflicts
```

**Critical path:** WI-7 → WI-1a → WI-1c → WI-1b → WI-1d → WI-3a

The criteria system is the highest-priority fix because it affects evaluation quality — every eval run until this is fixed produces results graded against the wrong criteria.

---

### 5. Disagreements

#### Note 70 — Remove `--model` flag: **I disagree.**

`--model` is useful for quick one-off experiments ("how does this prompt do on Opus vs Sonnet?") without creating separate config files. The config YAML should be the canonical source, but `--model` as a CLI override is a standard pattern in eval tools. I'd keep it but document it as an override, not a primary way to set the model.

#### Note 108 — Split guardrails into three tiers: **Partially disagree.**

Generator and reviewer tiers make sense. An "overall per-config" tier adds complexity for unclear benefit — if you set generator=50 actions and reviewer=25 actions, the overall limit is just a redundant cap. I'd implement generator + reviewer tiers only and let the sum be the effective overall limit.

#### Note 128 — Rename `CopilotEvaluator` → `PromptRunner`: **Agree with caveat.**

The rename is good conceptually, but `PromptRunner.Run()` collides with the common Go `Runner` interface pattern. I'd suggest `SessionRunner` — it runs a Copilot session for a prompt. More specific, less likely to collide.

#### Note 131 — Build timeline during report generation, not eval: **Disagree.**

The behavior grader needs the timeline *during* the eval to produce grading results. Moving timeline construction to report generation would require the behavior grader to work off raw events, which is a significant refactor with little benefit. The timeline is lightweight to build — the real cost is the SDK session, not event classification. Keep timeline construction during eval.

#### Note 129 — Unify criteria and graders into one system: **Agree, with phasing.**

This is the right end-state but it's a large architectural change. Phase it:
1. First: fix criteria loading (WI-1a) and remove rubric (WI-1b).
2. Second: automate criteria extraction (WI-1c) and enforce structured responses (WI-1d).
3. Third: unify into one grading pipeline where AI review is one grader type.

Trying to do it all at once risks destabilizing the eval engine.
