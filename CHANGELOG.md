# Changelog

All notable changes to hyoka are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- **WorkspaceDelta field in EvalReport JSON** — captures file-level changes (created, modified, deleted) per evaluation run, enabling graders and dashboards to reason about what the agent actually changed
- **Eval Detail page workspace delta display** — reports now render file-level changes (created/modified/deleted) when available
- **GraderResultRow component** — new reusable TypeScript component for rendering individual grader results with pass/fail badges, scores, and expandable details; standardizes grader display across report views
- **Starter-aware MaxOutputSize and MaxFiles guardrails** — guardrails now correctly account for workspace state, only charging the agent for delta output (new files or modifications), not starter-project files copied in before generation
- **Review session bucketing (Phase 6)** — reviewers now organize results by criteria buckets for improved multi-criterion grading
- **Embedded asset freshness policy** — site/dist must be rebuilt and committed whenever site/src changes
- **Installed-binary documentation form** — all documentation updated to use `hyoka` command form instead of `go run` for clarity
- **Post-architecture examples audit (WI-058)** — new examples demonstrating hierarchical `when:` syntax, prompt-level `graders:` frontmatter, and overall examples documentation (examples/README.md)
- **Tool load failure error category** — evals now report `error_category: "tool_load_failure"` when any configured tool fails to resolve before session creation
- **Grouped Tools progress output** — plugins and skill directories now expand to show each child tool/skill individually in the Tools progress section (e.g., `plugin-name (2 tools) ├── tool-1 └── tool-2`), enabling faster diagnostics

### Changed

- **Tool load validation now hard-fails before code generation** — evals abort immediately when any required plugin, skill directory, or MCP server fails to resolve, preventing silent failures with incomplete tool stacks
- **Site rebranded** — from Azure SDK code-generation evaluation to general-purpose AI agent evaluation platform
- **Example configs migrated** — `examples/configs/example-full.yaml` now reflects Phase 3 unified grading architecture (single results list, no separate review block)
- **Homepage, features, and documentation** — rewritten to describe general-purpose AI agent evaluation rather than Azure SDK code generation
- **Eval detail page** — refactored to render `grader_results` table instead of legacy rubric badges; includes graceful backward compatibility notice for pre-Phase-3 reports
- **TypeScript type alignment** — `GraderResult` types now match Go struct definitions, ensuring all grader detail types are properly represented in reports
- **Docs: installed-binary form** — all examples in docs/ now use `hyoka <cmd>` instead of `go run . <cmd>` for consistency with end-user perspective
- **getting-started.md** — clarified installation instructions with `go install ./...` and removed stale reference to old `hyoka/` directory structure

### Breaking Changes

- **Retired top-level `plugins:` field** — Pre-1.0, no deprecation path. All plugin declarations must migrate to `generator.tools` and `reviewer.tools` with `type: plugin`. Example migration:
  ```yaml
  # OLD (no longer supported)
  plugins:
    - "azure-sdk-python@skills"
  
  # NEW
  generator:
    tools:
      - name: azure-sdk-python
        type: plugin
        source: remote
  ```
  Affected configs (auto-migrated in this commit):
  - `configs/baseline-sonnet-skills.yaml`
  - `configs/python-pairwise.yaml`

### Fixed

- **Reviewer skill resolution cross-config leakage** — reviewer tool validation is now scoped per-config, preventing skills from one config accidentally being used in another during multi-config runs
- **False-positive MaxOutputSize failures on large starter projects** — large starter codebases no longer incorrectly trigger guardrail failures when copied into workspace before generation
- **Silent zero-render bug from Go↔TypeScript field drift** — Phase 3 unified grading now properly serialized to reports; eval detail page gracefully handles legacy reports with missing grader_results field
- **README and AGENTS docs:** — updated build/test commands to use `./...` glob instead of old `./hyoka/...` paths
- **Stale directory references in docs** — removed obsolete references to `hyoka/internal/` and `cd hyoka/` patterns from documentation

## [0.3.1] — Phase 3: Advanced Core & CLI Polish

### Added

- **Unified grading architecture** — AI review is now a standard `PromptReviewGrader` implementing the `Grader` interface; no more separate review phase
- **Structured JSON reviewer responses** — reviewers return validated JSON with per-criterion judgments and deterministic voting (any-fail across models = criterion fails)
- **Remote MCP server support** — `mcp_type: remote` with URL field for connecting to MCP servers over HTTP
- **Workspace containment hardening** — expanded PreToolUse validation, bash command containment, all file tools checked
- **Real-time guardrail enforcement** — turn/file limits enforced during session, not just post-hoc
- **Report tool usage tracking** — reports compare available vs. used tools per evaluation
- **Progress redesign** — section-based per-eval layout for clearer output
- **File exclusion unification** — single configurable system across all file operations
- **Prompt package updates** — removed flat format compatibility, added heading-aware splitter

### Changed

- **Results overwrite bug fixed** — all grader types contribute to single results list
- **Help text audit** — fixed stale references, split filter/run flags for clarity
- **Docs audit & cleanup** — updated all documentation, removed stale content

### Removed

- **Snapshot hack** — workspace snapshot logic cleaned up in favor of containment validation

---

## Legend

- **Added:** new features and capabilities
- **Changed:** changes to existing functionality
- **Deprecated:** soon-to-be-removed functionality
- **Removed:** removed functionality
- **Fixed:** bug fixes
- **Security:** security-related changes
