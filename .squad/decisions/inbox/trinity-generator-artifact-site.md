# Decision: Surface generator.json artifact on eval-detail page

**Date:** 2026-04-24  
**Agent:** Trinity  
**Status:** Implemented  

## Context

Neo shipped `GeneratorArtifact` (commit d1ed5f61) to capture full generator session state (prompt, final response, workspace delta, actions, timing, termination) and write it to `{reportDir}/generator.json` for grader consumption. User directive: "This generator.json file should also be part of the data we use on the site."

## Decision

Wire `GeneratorArtifact` into the report layer and render it on the eval-detail page as a collapsible "Generator Session" panel.

### Go Layer (Phase 1)

1. **Report integration:**
   - Add `GeneratorArtifact *artifact.GeneratorArtifact \`json:"generator_artifact,omitempty"\`` to `report.EvalReport`
   - Use type alias pattern (consistent with `WorkspaceDelta`) to avoid import cycles

2. **Artifact lifecycle:**
   - **Write:** After workspace delta computed, before graders run (line 530 in `engine_eval.go`)
   - **Read:** After FileContents populated, before WriteReport (line 717 in `engine_eval.go`)
   - Helper: `buildGeneratorArtifact()` constructs artifact from eval state (prompt, config, result, timing, termination)

3. **Schema v3:**
   - Neo already bumped `CurrentSchemaVersion` from 2 → 3
   - `generator_artifact` field is `omitempty` so v2 reports remain valid

### TypeScript Layer (Phase 2)

1. **Type definitions (`site/src/app/data/types.ts`):**
   - Mirror Go structs with snake_case JSON tags:
     - `GeneratorArtifact` (prompt_id, config_name, generator_model, original_prompt, final_response, workspace_delta, actions_summary, started_at, ended_at, duration_ms, terminated_by, error)
     - `ArtifactWorkspaceDelta` (bytes_added/removed/net, file counts, created/modified/deleted file arrays)
     - `ActionsSummary` (total_actions, tool_calls, reasoning_steps, truncated)
     - `ArtifactFileInfo` (path, size)

2. **EvalReport extension:**
   - Add `generator_artifact?: GeneratorArtifact` with doc comment explaining v3 addition

### UI Layer (Phase 3)

1. **Panel placement:** ABOVE "Generated Files" panel (between reviewer timelines and files)

2. **Collapsed by default:** New state `showGenSession` (false by default)

3. **Panel sections:**
   - **Termination badge:** Color-coded status (green=completed, yellow=max_actions, orange=timeout/guardrail, red=error)
   - **Timing:** Duration as "Xm Ys", started timestamp
   - **Actions summary:** 3-column grid (total/tool-calls/reasoning), truncation flag
   - **Workspace delta:** Created/modified/deleted file counts with colored badges (emerald/amber/red)
   - **Final response:** Truncated to 500 chars if >500 AND files generated; full text otherwise; copy button

4. **Conditional render:** Only show panel if `generator_artifact` exists (handles v2 reports gracefully with no error state)

## Rationale

- **Write timing:** Artifact must capture complete generation state (including workspace delta) but be written BEFORE graders run (graders may consume it via `GraderInput.GeneratorArtifactPath`).
- **Read timing:** Populate artifact AFTER file contents (Bug 2 fix) to keep related functionality grouped.
- **Collapsed default:** Avoids clutter for users who primarily care about grader scores; power users can expand to see session details.
- **Truncation logic:** Show full response if <500 chars OR no files generated (response-only evals); truncate if long + files exist (file-focused evals).
- **Backward compat:** `omitempty` + conditional render means v2 reports display identically (no artifact panel).

## Consequences

- **Positive:**
  - Site now displays full generator session context (response, actions, timing, termination)
  - Graders and site consume same source of truth (`generator.json`)
  - v2/v3 reports coexist safely (v2 = no artifact panel, v3 = full panel)
  
- **Negative:**
  - One more panel to render (mitigated by collapse-by-default)
  - Artifact file added to each eval's report directory (+1 file per eval)

## Alternatives Considered

1. **Inline response in timeline:** Rejected; timeline is tool-focused, response is distinct
2. **Merge into "Generated Files" panel:** Rejected; semantically different (session metadata vs file contents)
3. **Always expanded:** Rejected; clutters UI for majority of users

## Follow-up

- [ ] Add artifact download link (let users export `generator.json` directly)
- [ ] Syntax highlight final_response if it looks like code/markdown
- [ ] Add artifact file-list expansion (show full file paths + sizes)
