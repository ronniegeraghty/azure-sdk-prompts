---


---

## 4. Phase-2 UX proposals (after Points lands)

All wireframes below stay inside the existing visual language: same `mono` font, `rounded-xl border border-white/8 bg-white/[0.03]` card chrome, emerald/red/amber accents from Tailwind tokens already in use.

### 4a. Plugin parent — header-only group, children carry status

**Today** (eval-detail "Available Tools", lines 495-535) — flat row of pill tags:
```
[azure (MCP)] [azure-keyvault-py (skill)] [azure-cosmos-py (skill)] … (40 more)
```

**Proposed** — group by `parent_kind/parent_name`, parent renders as a header without a status pill:

```
Available Tools

  azure-sdk-python  (plugin)            ← header only, no pass/fail badge
  ├─ azure-keyvault-py        (skill)   ✓ used
  ├─ azure-cosmos-py          (skill)
  ├─ azure-identity-py        (skill)
  └─ azure-eventhub-py        (skill)
  …41 children…

  generator-skills  (skills dir)        ← header uses config `name`, not path
  ├─ python-script-quality    (skill)
  └─ python-readme            (skill)

  azure              (mcp)              ← top-level, unchanged from today
```

Markup sketch (reusing existing `inline-flex … rounded-md px-2 py-1 bg-…/10 text-…/40` tag chrome):

```tsx
{Object.entries(toolsByParent).map(([parentKey, group]) => (
  <div key={parentKey} className="mb-3">
    {group.parent && (
      <div className="mb-1.5 flex items-center gap-1.5 text-white/50" style={{ ...mono, fontSize: 11 }}>
        <span>{group.parent.name}</span>
        <span className="text-white/30">({group.parent.kind})</span>
        {/* deliberately NO status pill on parent */}
      </div>
    )}
    <div className={`flex flex-wrap gap-1.5 ${group.parent ? "ml-4 border-l border-white/5 pl-3" : ""}`}>
      {group.children.map(child => <ToolTag key={child.name} tool={child} />)}
    </div>
  </div>
))}
```

The `ml-4 border-l pl-3` indent matches what the run-detail-page table already does for nested filter rows — visual language stays consistent.

**Skill-dir parent** uses `entry.Name` from config (which the validator change in plan.md Phase 1 already routes through) — so the header reads `generator-skills (skills dir)` not `./skills/generator (skills dir)`. No template change needed for that — purely a data-shape ask: the JSON needs `parent_name` populated from config.Name (Morpheus territory).

### 4b. Grader card — header + N indented points

**Today** (`GraderResultRow` collapsed):
```
[N/A] consensus                Review • consensus            12/12
```

**Proposed post-Points** (collapsed):
```
[✓ 12/12 passed] consensus                Review • consensus
```

**Expanded:**
```
[✓ 12/12 passed] consensus                Review • consensus       ▼
  ✓ DefaultAzureCredential Authentication
  ✓ Installing azure-keyvault-secrets and azure-identity packages
  ✓ Creating a SecretClient with vault URL and credential
  ✓ set_secret(), get_secret(), begin_delete_secret(), purge_deleted_secret()
  …
  ✗ paginates list_secrets                                           (if any failed)
```

Markup, reusing GraderResultRow's existing chrome + the criterion pill style already used on line 615-660 for review criteria:

```tsx
const passCount = (result.points ?? []).filter(p => p.pass).length;
const totalCount = (result.points ?? []).length;
const allPass = totalCount > 0 && passCount === totalCount;
const noneRan = totalCount === 0;

// Header badge:
<div className={`flex items-center gap-1.5 rounded-md border px-2 py-1 ${badgeColor}`} style={{ fontSize: 10 }}>
  {badgeIcon}
  <span>
    {noneRan ? (passed === true ? "PASS" : passed === false ? "FAIL" : "N/A")
             : `${passCount}/${totalCount} passed`}
  </span>
</div>

// Expanded body — one row per point, indented:
{(result.points ?? []).map(p => (
  <div key={p.name} className="ml-6 flex items-start gap-2 py-1" style={{ fontSize: 12 }}>
    {p.pass
      ? <CheckCircle2 className="mt-0.5 h-3 w-3 shrink-0 text-emerald-400" />
      : <XCircle className="mt-0.5 h-3 w-3 shrink-0 text-red-400" />}
    <div className="flex-1">
      <div className={p.pass ? "text-white/70" : "text-red-400/80"}>{p.name}</div>
      {p.message && <div className="text-white/40" style={{ fontSize: 11 }}>{p.message}</div>}
    </div>
  </div>
))}
```

Existing `ChevronDown / ChevronRight` toggle stays. The `ml-6` indent matches the criterion list rendering already shipped at eval-detail-page.tsx:615-660.

When `len(points) <= 1`, fall back to today's flat row with a real PASS/FAIL badge — matches what the CLI renderer plans to do per plan.md.

### 4c. Roll-up — one rule, applied everywhere

Add a single helper `isEvalPass(eval)` and use it on **every** surface that paints emerald-vs-red:

```ts
function evalPassFromPoints(r: EvalReport): boolean {
  // Phase 2: every grader emits Points; eval passes iff every point passes
  if (r.grader_results?.length) {
    return r.grader_results.every(g =>
      (g.points?.length ?? 0) === 0
        ? g.pass === true
        : g.points!.every(p => p.pass));
  }
  // No graders ran (e.g. errored generation) → fall back to existing flag
  return r.success === true;
}

function runPassFromEvals(run: RunSummary): { passed: number; failed: number; errored: number } {
  // ... aggregates evalPassFromPoints across run.results
}
```

Apply to:
- `runs-page.tsx` rate calculation
- `run-detail-page.tsx` `gradersPassed/gradersTotal` cell (replace strict `=== true` filter)
- `eval-detail-page.tsx` header score card (replace `r.success` direct read)
- `prompt-detail-page.tsx` "Pass Rate by Config" / "by Tool" aggregations
- `pairwise-page.tsx` and `comparison-page.tsx` per-cell win/loss

Once the rule lives in one place, no surface can drift again.

### 4d. Score-card and counts the user actually wants

For the eval-detail header (replacing the `12/12` review-only card):

```
   ┌────────────────────────────┐
   │     ✓ 15 / 15  points      │   ← AND of every Point across all graders
   │     across 4 graders       │
   └────────────────────────────┘
```

For the run-detail row (replacing `1/4` red):

```
   Score
   ┌──────┐
   │ 4/4  │   ← graders, not points (concise enough for a table cell)
   └──────┘
   tooltip: "15/15 points (4 graders) — output_check, opus-4.6, gpt-4.1, consensus"
```

Two-tier display: granularity at the eval, summary at the run. Both honest, both green when the answer is green.

---

## 5. Coordination with Morpheus

Morpheus has `phase4-pw-final-summary.md` in his folder but no `morpheus-report-architecture-review.md` in `.squad/decisions/inbox/` as of writing. **Will link when it lands.** Boundary respected: this report doesn't propose data-model changes — only flags what the templates need from the wire.

**Data fields the site needs from the JSON that aren't there today** (Morpheus to weigh in on shape):
1. **`parent_kind` + `parent_name` on every `tool_availability` entry.** Today's flat array can't express plugin/skill-dir grouping. Without this the renderer in §4a is impossible.
2. **`Points []GraderPoint{Name, Pass, Message}` on `GraderResult`** — already in plan.md Phase 2; confirming the site needs it on the JSON for §4b.
3. **`pass: boolean` populated correctly on every grader** — including review-type graders. Today the strict `pass === true` filter excludes review graders entirely. If the backend continues returning `pass: null` for review graders, the site will keep needing the fallback in Q1/Q2; cleaner if the engine sets `pass = points.every(p => p.pass)` once Points lands.
4. **`overall_pass: boolean` on `EvalReport`** — currently we lean on `success` (success of *generation*, I think?) and re-derive from grader_results. A canonical eval-level boolean computed by the engine from Points would let every site surface stop reinventing the rule.
5. **`overall_passed` / `pass_rate_excluding_errors` on `RunSummary`** — to fix §1e without each surface re-deriving. Already pre-computed server-side in `summary.go`?

If any of these change shape, ping Trinity — site types live in `site/src/app/data/types.ts`.

---

## 6. Open questions for the user (Ronnie)

1. **For review graders (LLM-as-panel), what's the canonical pass definition?** Today: `criteria.every(c => c.passed)` is implicitly true on the anchor run, but `pass` on the wire is `null`. Should the engine just write `pass = AND(criteria)` for review graders? Or do you want the consensus grader to use a different rule (e.g. quorum)? This determines whether Q1/Q2 are temporary or load-bearing.

2. **Score-card metaphor on the eval page.** Pre-Points the `12/12` is "reviewer agreement". Post-Points it could be "points passed / points total" (granular, true) or "graders passed / graders total" (concise, less true). Which do you want as the dominant number?

3. **In-progress runs on `/runs`.** The list currently shows the partial `20260423-203921` as "Unknown 0.0% N/A". Hide it, or render it as `In progress`?

4. **Errored runs on `/runs`.** Should I split the rate column into `passed | failed | errored` instead of folding errors into the denominator?

5. **Dashboard page priority.** It's been crashing for some unknown number of commits. Pull a fix into the same wave as the grader render fixes, or carve a separate session?

6. **Plugin parent click target.** Should the parent header in §4a be expandable/collapsable (hide the 41 children behind a `▶` toggle), or always-expanded? On a prompt that loads the full `azure-sdk-python` plugin the children take ~6 visual rows of pills.

---

## Appendix — quick reference

- **Anchor JSON**: `/home/rgeraghty/projects/hyoka/reports/20260423-195948/results/key-vault/data-plane/python/crud/key-vault-dp-python-crud/python-pairwise/baseline/claude-sonnet-4.5/report.json`
- **Screenshots**: `/tmp/trinity-site-review/01-runs-list.png` … `11-runs-listing-bad-rows.png`
- **Files cited**:
  - `hyoka/internal/serve/serve.go:142-178` (route map)
  - `site/src/app/components/runs-page.tsx:136-198` (run cards, errored-row issue)
  - `site/src/app/components/run-detail-page.tsx:236-256` (table score cell — primary bug)
  - `site/src/app/components/eval-detail-page.tsx:382-461,495-560,615-660` (header card, tools, graders, criteria render)
  - `site/src/app/components/GraderResultRow.tsx:16,29-39,68` (tri-state badge logic)
  - `site/src/app/components/dashboard-page.tsx` (crash site, line not yet bisected)
  - `site/src/app/data/types.ts` (no `parent_kind`/`parent_name`/`points` types yet)
- **Plan reference**: session-state plan at `/home/rgeraghty/.copilot/session-state/87e98ab8-…/plan.md` Phase 1 + Phase 2.

Read-only review — no templates, CSS, types, or data files were modified. No commits, no pushes.


---

## Latest Decisions (2026-04-24: Generator.json Artifact Arc)

### 2026-04-24T02:50:00Z: Grader Inputs Always Provided, Never Conditional

**Agent:** Ronnie Geraghty (via Copilot)

**What:** Graders (typed AND AI review) always receive BOTH:
1. A copy of the generator's working directory files
2. `generator.json` containing all collected session info (prompt, final response, workspace delta, action summary, termination, etc.)

Neither is conditional on the other being empty. The reviewer prompt unconditionally includes both the Generated Code section AND the Agent Session section. If files are missing, the Generated Code section says "(no files were created)" but the section is still present.

**Why:** User clarification: "we should always pass anything that is available. We should copy over all the files from the generator's working dir and the generator.json that contains all the info we collect from the generator agent's work that is provided to the graders."

**Implications:**
1. GraderInput carries both `WorkspacePath` and `GeneratorArtifactPath`
2. AI review graders always get a temp-dir copy of workspace files (existing PromptReviewGrader behavior continues)
3. Typed graders read workspace directly from `WorkspacePath`
4. Reviewer prompt rendering always includes both sections (never conditional)

**Related:** `coordinator-grader-input-model.md`

---

### 2026-04-24T02:56:00Z: Grader Input Model — generator.json + AI Reviewer Workspace Copy

**Agent:** Ronnie Geraghty (via Copilot)

**What:**
1. ALL graders (typed AND AI review) receive the same input data via the `generator.json` data model — single canonical schema describing the generator session (prompt, final response, workspace delta, action summary, termination, etc.). The artifact is persisted to disk per eval and also passed in-memory via GraderInput. Loose convenience fields on GraderInput remain for back-compat but generator.json is the authoritative source.

2. AI review graders ADDITIONALLY receive a copy of all workspace files in their reviewer working dir (existing PromptReviewGrader.copyDirToTemp behavior). This is the only grader-class-specific input difference. Typed graders read the workspace from `WorkspacePath` directly (no copy).

3. The reviewer prompt always includes both the Generated Code section (rendered from the workspace copy) and the Agent Session section (rendered from generator.json) — never conditional on either being empty. When files are absent, the Generated Code section says so explicitly but is still present.

**Why:** User clarification: "we should be giving all the graders the same information via the data model of the generator.json and we should additionally give AI reviewer graders a copy of all workspace files in their working dir."

**Implementation:** Captured in orchestration logs 2026-04-24T09:15:00Z-neo.md and 2026-04-24T10:30:00Z-trinity.md.

---

### 2026-04-24T02:56Z: Generator.json Artifact Must Be Surfaced on the Site

**Agent:** Ronnie (via Coordinator)

**What:** The `generator.json` artifact (prompt_id, original_prompt, final_response, workspace_delta, actions_summary, timing, terminated_by, error) that the engine emits for graders MUST also be ingested into the report layer and rendered on the site (eval-detail page at minimum). Same data model — graders and site consume one canonical artifact.

**Why:** User directive — single source of truth for everything the generator session produced. Site users want to see the agent's final response, what files changed, and how the session terminated, alongside file contents and grader results.

**Implications:**
1. `report.EvalReport` gets a `GeneratorArtifact *GeneratorArtifact` field (or inline equivalent). Populated alongside `FileContents`.
2. JSON schema bumps to v3 (Neo Phase 1).
3. Site `data/types.ts` mirrors the new shape; eval-detail template renders a "Generator Session" panel showing final_response (collapsible), workspace_delta summary, actions/turns/duration, terminated_by/error.
4. No conditional rendering — if artifact is present (v3+), show it. v2 reports just don't have it.

**Status:** ✅ Implemented (Trinity Phase 2, commit 9f34f072)

---

### 2026-04-24 (Date TBD): Eval Detail Pages Include Generated File Contents

**Agent:** Trinity 🌐

**Status:** ✅ Implemented (Commit c06ca9e2)

**What:** Generated file contents are now captured in `report.EvalReport.FileContents` at report-build time.

**Root Cause of Bug:** The eval detail page served by `hyoka serve` at `/api/runs/{runId}/eval?path=...` returns the `report.json` for that eval, which contained the `GeneratedFiles` array (a list of file paths). However, the file **contents** were never populated in the report JSON.

**Decision:**
1. Added `FileContents map[string]string` field to `EvalReport` (marked `json:"file_contents,omitempty"`).
2. Added `readGeneratedFileContents()` helper in `engine_eval.go`.
   - Called right before `WriteReport()`, reads each file from `ws.Dir` (the workspace directory).
   - **Size cap:** Files exceeding 1MB are capped with a message: `[File too large to display (N bytes) — view on disk at {path}]`.
   - **Binary detection:** Files with binary extensions are skipped.
   - **Error handling:** Files that can't be read show: `[Error reading file: {error}]`.
3. Populated `evalReport.FileContents` before calling `report.WriteReport()`.

**Binary Extensions Detected:**  
`.png`, `.jpg`, `.jpeg`, `.gif`, `.bmp`, `.pdf`, `.zip`, `.tar`, `.gz`, `.7z`, `.exe`, `.dll`, `.so`, `.dylib`, `.bin`, `.dat`, `.db`, `.sqlite`

**Verification:**
- All tests pass: `go test ./hyoka/...`
- Live test: `hyoka serve` on existing report directory shows file contents in eval JSON

**Impact:**
- Site can now display generated file contents on eval detail pages
- Size-safe: 1MB cap per file prevents JSON bloat
- Backward compatible: Existing reports without `file_contents` continue to work

**Reusable Pattern:** Capture report artifacts at report-build time, not serve time. Reading when the report is written (workspace still exists) is more reliable than serve-time reads.

---

### 2026-04-24: Surface generator.json Artifact on Eval-Detail Page

**Agent:** Trinity 🌐

**Status:** ✅ Implemented (Commit 9f34f072)

**What:** Wire `GeneratorArtifact` into the report layer and render it on the eval-detail page as a collapsible "Generator Session" panel.

**Implementation:**

**Go Layer (Phase 1):**
1. Add `GeneratorArtifact *artifact.GeneratorArtifact \`json:"generator_artifact,omitempty"\`` to `report.EvalReport`
2. Type alias pattern (consistent with `WorkspaceDelta`) to avoid import cycles
3. **Write:** After workspace delta computed, before graders run
4. **Read:** After FileContents populated, before WriteReport
5. Helper: `buildGeneratorArtifact()` constructs artifact from eval state

**TypeScript Layer (Phase 2):**
1. Type definitions (`site/src/app/data/types.ts`) mirror Go structs with snake_case
2. Add `generator_artifact?: GeneratorArtifact` to EvalReport

**UI Layer (Phase 3):**
1. Panel placement: ABOVE "Generated Files" panel
2. Collapsed by default with state `showGenSession` (false by default)
3. Panel sections:
   - **Termination badge:** Color-coded status (green=completed, yellow=max_actions, orange=timeout/guardrail, red=error)
   - **Timing:** Duration as "Xm Ys", started timestamp
   - **Actions summary:** 3-column grid (total/tool-calls/reasoning), truncation flag
   - **Workspace delta:** Created/modified/deleted file counts with colored badges
   - **Final response:** Truncated to 500 chars if >500 AND files generated; full text otherwise; copy button
4. Conditional render: Only show panel if `generator_artifact` exists (handles v2 reports gracefully)

**Rationale:**
- Write timing: Artifact must capture complete generation state before graders run
- Read timing: Populate artifact after file contents to keep related functionality grouped
- Collapsed default: Avoids clutter for users who primarily care about grader scores
- Truncation logic: Show full response if <500 chars OR no files generated; truncate if long + files exist
- Backward compat: `omitempty` + conditional render means v2 reports display identically

---

### 2026-04-24: Grader Score Denominator Counts All Grader Points

**Agent:** Trinity 🌐

**Status:** ✅ Implemented (Commit c06ca9e2)

**What:** The total score denominator is the sum of all grader points across all graders, not the count of graders.

**Root Cause of Bug:** In `engine_eval.go:636-637`, `GradersTotal` and `GradersPassed` fields were computed using `len(agg.Results)` and `countPassed(agg.Results)`, which counted the number of graders, not the number of grader points.

**Example:**
- `file_check` grader with 3 points (file1, file2, file3)
- `output_check` grader with 2 points (min_files, require_files)
- Old denominator: 2 (number of graders)
- Correct denominator: 5 (number of grader points)
- Site incorrectly displayed `3/2` instead of `3/5`

**Decision:**
1. Added `countTotalPoints()` helper that sums `len(g.Points)` for each grader in `agg.Results`
   - Graders with `len(Points) == 0` treated as 1 point (backward compatibility)
2. Added `countPassedPoints()` helper that counts passed points across all graders
   - For graders with Points, count points where `Point.Pass == true`
   - For graders with no Points (legacy), use grader's overall `Pass` field (1 if true, 0 if false)
3. Updated `engine_eval.go:636-637` to use point-based helpers

**Verification:**
- All tests pass: `go test ./hyoka/...`
- Table-driven tests cover: multiple points per grader, zero-point graders (legacy), mixed scenarios, empty results

**Impact:**
- Site now displays accurate `X/Y` scores where Y is total grader points
- Backward compatible: Legacy graders treated as 1 point
- Consistent with grader-points architecture

**Reusable Rule:** When aggregating grader scores, always count grader points, not graders. Denominator is `Σ len(g.Points) for g in graders`, with fallback of 1 for legacy graders.

---

### 2026-04-24: Prompt-Frontmatter Criteria Always Get Their Own Review-Grader Bucket

**Agent:** Neo 💊

**Status:** ✅ Implemented (Commit 27c04c71)

**Branch:** ronniegeraghty/dev

**What:** Prompt-frontmatter criteria and criteria-file entries are ALWAYS bucketed separately, regardless of review mode (combined or isolated).

**Root Cause of Bug:** User report: "I'm only seeing one group of ai review graders running but I thought we decided that if we wanted grader points to be graded in the same review session they would have to be grader points on the same grader."

`BuildUnifiedReviewBuckets` previously merged `promptCriteria` (from prompt frontmatter) with matched criteria-file entries into a single `combined` bucket, resulting in one AI review grader. This violated the source-separation principle.

**Decision:**
Each source produces its own `PromptReviewGrader` instance → its own Copilot review session → its own grader display entry.

**Bucket Naming:**
- Prompt-frontmatter criteria → bucket named `"Criteria from prompt file"`
- Criteria-file entries → bucket(s) named based on mode:
  - Combined mode: one `"combined"` bucket for all criteria-file entries
  - Isolated mode: separate buckets per isolated entry/group, plus a `"combined"` bucket for leftovers

**Edge Cases:**
- If `promptCriteria` is empty AND matched criteria-file entries exist → only criteria-file bucket(s)
- If `promptCriteria` is non-empty AND no criteria-file entries → only prompt-frontmatter bucket
- If both are empty → zero buckets (no review graders)

**Rationale:**
**Source of truth determines grader identity.** Prompt-frontmatter criteria come from the `.prompt.md` file's YAML frontmatter. Criteria-file entries come from `criteria/*.yaml` files. These are fundamentally different sources and must be evaluated in separate Copilot sessions:

1. Grader results are distinguishable — users see which bucket contributed which failing criteria
2. Isolation is truly isolated — failure in one doesn't block the other
3. Bucket names are meaningful — "Criteria from prompt file" is explicit

The review mode (combined vs isolated) controls how criteria-FILE entries are bucketed among themselves; it does NOT control whether prompt-frontmatter criteria are separated. Source-separation is a harder rule than mode-separation.

**Implementation:**
- `buckets.go`: Refactored `BuildUnifiedReviewBuckets` to always prepend prompt-frontmatter bucket
- `buckets_test.go`: Updated all combined-mode tests to expect 2 buckets; added edge-case tests
- `engine_reviewbuckets_test.go` and `engine_reviewmode_runtime_test.go`: Updated integration tests

**Verification:**
- ✅ All tests pass: `go test -race ./hyoka/...`
- ✅ Build clean
- 🔄 Live verification: observe TWO separate AI review grader entries in output

**Reusable Principle:** Grader source-of-truth determines grader identity. When multiple criteria sources contribute (prompt frontmatter, criteria files, future: remote criteria), each source must produce a separate grader instance if we want distinguishable pass/fail reporting. The `--review-mode` flag is SECONDARY — it partitions WITHIN a single source's bucket set.

---

### 2026-04-24: Phase-State Guard for Reviewer Event Suppression

**Agent:** Tank ⚛️

**Status:** ✅ Implemented (Commit 6f2e1f03)

**What:** Add a **phase-state guard** at the top of `renderAgentEvent` to suppress activity events from downstream reviewer sessions once the generation phase is complete.

**Root Cause of Bug 3:** The PromptReviewGrader runs a real Copilot SDK session that emits the same generation events (EventReasoning, EventToolStart, EventToolComplete) through the shared progress event channel. After EventSessionDetails flipped Agent Attempt → Completed, subsequent reviewer activity events were still landing in `renderAgentEvent` and creating duplicate "Agent Attempt: ✅ Completed" rows.

**Decision:**
```go
// Agent Attempt is already finalized — generation phase is over. Ignore
// activity events from downstream sessions (reviewer Copilot sessions
// emit the same EventReasoning/EventToolStart/etc. through the shared
// event channel, but they belong to grader rows, not the agent tail).
if r.cur != nil && (r.cur.agentState == agentStateCompleted || r.cur.agentState == agentStateGuardrail) {
    return
}
```

**Rationale:**
- The agent tail belongs to GENERATION phase only
- Once `agentState` is Completed or Guardrail, generation phase is closed
- Downstream sessions (reviewers) emit same event types but belong to grader rows, NOT agent tail
- Renderer must filter events by phase-ownership, not just event type

**Implementation:**
- `display_interactive.go` (+5 lines): phase-state guard in `renderAgentEvent`
- `display_interactive_test.go` (+56 lines): regression test `TestInteractive_ReviewerEventsAfterCompletionIgnored`

**Test Coverage:**
`TestInteractive_ReviewerEventsAfterCompletionIgnored` drives full sequence: generation → EventSessionDetails (completes agent) → typed grader → AI review grader start → EventReasoning/EventToolStart/EventToolComplete from reviewer → grader complete → passed. Asserts exactly ONE "Agent Attempt:" row.

**Verification:**
- `go test -race ./hyoka/internal/progress/...` — all tests pass
- Live eval (key-vault-dp-python-crud / python-pairwise) — zero duplicate rows

**Alternatives Considered:**
1. **Filter at event emission:** Rejected — couples eval package to display semantics
2. **Add "review phase" state:** Rejected — over-engineered; agent tail is generation-only
3. **Change event types:** Rejected — would break existing event consumers

**Lessons:**
- Phase-state guards are critical in event-driven terminal renderers
- When multiple concurrent processes emit through shared channel, filter by phase, not event type
- Test for phase isolation — events from one phase must not bleed into another's rendering

---


### Decision: Grader Output Format — Per-Grader Display with Point Breakdown (2026-04-24T03:26Z)

**Agent:** Coordinator (user direction via Ronnie)  
**Branch:** ronniegeraghty/dev  
**Implementation:** Tank (commit `4adc9288`)  

## Context

Current progress display lumped all AI-review grader points under a single `- ai_review` line, making it impossible to see which specific grader passed or failed which criteria point.

## Decision

**Grader output now displays per-grader with individual point breakdown.**

Format:
```
- <grader-name> (<grader-type>): <score>
  - <point.name or first ~60 chars of point.prompt>: <pass|fail>
  - <point.name or first ~60 chars of point.prompt>: <pass|fail>
```

Examples:
- `- DefaultAzureCredential Authentication (prompt): 3/4`
  - `- Uses DefaultAzureCredential from azure.identity: pass`
  - `- Wraps client construction in try/except: fail`
- `- Criteria from prompt file (prompt): 5/5`
- `- file_exists/keyvault_crud.py (file): pass`
- `- behavior/runs_without_error (behavior): pass`

## Rules

1. **Grader name** comes from the grader entry's `name:` field. For prompt-frontmatter bucket, hard-coded to `Criteria from prompt file`.
2. **Grader type** rendered in parens after name: `prompt` (AI review), `file`, `behavior`, `program`, `output_check`.
3. **Points** indented as sub-bullets. Use point's `name` if present; else first ~60 chars of `prompt` with ellipsis.
4. **Single-point graders** render on one line — no sub-bullets.
5. **Score format:** `pass`/`fail` for binary; `X/Y` for multi-point counts.

## Scope

- Interactive CLI display
- JSON report (`graders[].name`, `graders[].type`)
- Site eval-detail page (Trinity's frontend panel)

## Implementation

Tank refactored engine to create one PromptReviewGrader per ReviewBucket instead of one "ai_review" grader for all buckets. Result:
- `engine_eval.go`: Per-bucket grader iteration
- `display_interactive.go`: `displayKind()` helper for type mapping
- `display_interactive_points_test.go`: Updated assertions for multi-grader rendering

## Rationale

Single `- ai_review` row with nested points provides no signal about which grader (e.g., "DefaultAzureCredential") passed or failed. Per-grader rows with type labels and point-level breakdown enables users to:
- Identify which criteria bucket failed
- See exactly which individual points passed/failed
- Correlate failures to grader source (prompt, file, behavior, etc.)

## Verification

✅ All tests pass  
✅ Live evals display per-bucket graders with point breakdown  
✅ Site renders grader names and types correctly

---


---

### Decision: Prompt Grader `checks:` Redesign — Scope (2026-04-24T05:37:56Z)

**Agent:** Morpheus 🕶️  
**For:** Neo (engine/grader logic) and Tank (CLI/report rendering) collaboration  
**Status:** Scoped, ready for implementation  
**Branch:** `ronniegeraghty/prompt-grader-checks`

## Problem

Prompt-file-driven prompt graders look fine in the run output: each top-level bullet under `## Evaluation Criteria` becomes its own pass/fail row nested under the parent grader (handled by `prompt.ParseEvaluationCriteria` → `ParsedCriteria` → bucket rendering → LLM returns one judgement per criterion → `result.Points`).

YAML criteria-file prompt graders do **not** look fine. The current schema has a single `prompt:` field per `UnifiedGraderEntry`. Authors who want multiple checks today smuggle them inside that string ("Check the following criteria: 1. … 2. …"). The bucket renderer (`FormatUnifiedPromptEntries` in `internal/criteria/buckets.go:119`) emits that whole blob as a **single** numbered criterion to the review LLM, so the LLM returns **one** pass/fail for the whole blob and we lose per-check granularity. The terminal renderer already supports multi-Point graders (`internal/progress/display_interactive.go:1003-1062`) — there's just nothing for it to render.

The user wants YAML prompt graders to declare checks as first-class structured items so we get the same nested per-check output the prompt-file path already produces.

## Solution

**Add `Checks []string` field to `UnifiedGraderEntry`; migrate two criteria files to use it; update bucket rendering and grader logic; verify end-to-end through report layer.**

### Schema change

Add to `UnifiedGraderEntry` in `internal/criteria/config.go`:

```go
Checks []string `yaml:"checks,omitempty" json:"checks,omitempty"`
```

YAML shape:

```yaml
graders:
  - name: Markdown Structure
    type: prompt
    weight: 1.0
    prompt: Check the following criteria:    # optional preamble
    checks:
      - File hello.md exists and contains a level-1 heading.
      - File contains exactly three bullet list items.
```

### Validation (extends `validateEntry`)

For `type: prompt`:
- One of `prompt:` or `checks:` MUST be non-empty. Both may be set.
- If `checks:` is set, each entry must be a non-empty trimmed string. Empty/whitespace-only entries fail with index.
- If `checks:` is unset/empty, backward compat: single `prompt:` becomes one criterion → one Point.
- For non-prompt types: `checks` must be empty (mirrors existing `prompt`-must-be-empty rule).

### Bucket rendering (`internal/criteria/buckets.go`)

`FormatUnifiedPromptEntries` now honors `Checks`:

- **With checks:** Parent line `N. **Name**`, optional preamble from `Prompt` field (judge-only, rendered indented), then nested checks as numbered items:
  ```
  3. **DefaultAzureCredential Authentication**
     Check the following criteria:
       1. Uses DefaultAzureCredential from azure.identity.
       2. Wraps client construction in try/except.
  ```
- **Without checks:** Legacy single-line behavior preserved byte-for-byte: `N. **Name** — Prompt`.

### Grader logic (`internal/criteria/graders/prompt_review_grader.go`)

- Count expected criteria: parse the rendered bucket text to find leaf-numbered items (prefers indented when present); compare against returned criteria count.
- Emit `slog.Debug` when count differs (non-fatal; helps track outliers).

### YAML Migrations

Hard-migrate two criterion files per §2:

- **`criteria/language/python.yaml`** — `DefaultAzureCredential Authentication`: split embedded "1. … 2. …" into preamble `"Check the following criteria:"` + 2 `checks:` items (auth-credential, async/await).
- **`criteria/language/test.yaml`** — `Markdown Structure`: matches the user's worked example exactly (preamble + 2 checks for hello.md heading and bullet count).

Other files (`java.yaml`, `rust.yaml`, testdata) unchanged (no embedded numbering; single-prompt path covers them).

### Tests

- `hyoka/internal/criteria/config_test.go`: Replaced `TestValidateEntry_PromptMissingPrompt`; added 6-case `TestValidateEntry_PromptChecks` table covering checks-only, preamble+checks, legacy prompt-only, empty string in checks (with index), whitespace-only check, checks on wrong type.
- `hyoka/internal/criteria/buckets_test.go`: Added `TestFormatUnifiedPromptEntries_Shapes` (5 cases): empty, legacy single-prompt, checks+preamble, checks-no-preamble, mixed legacy+checks. All assert exact rendered text.

### Verification

```
go build ./...
go test ./... -timeout 5m
go test -race ./hyoka/internal/criteria/...
hyoka run --prompt-id key-vault-dp-python-crud --config baseline/claude-opus-4.6
```

Expected: Per-check Points propagate end-to-end. Debug log fires on criterion-count mismatch.

---

## Ownership

- **Neo (§1, §4, §5):** Schema change, validation, `FormatUnifiedPromptEntries` renderer, grader logic.
- **Tank (§6, §7):** Display badge format update (`✅ Pass (X/Y)` / `❌ Fail (X/Y)`), verify Points end-to-end in report layer.

File sets are disjoint; can land independently.

---

### Decision: Neo — Prompt Grader `checks:` Implementation (2026-04-24T05:37:56Z)

**Agent:** Neo 💊  
**For:** Coordinator / Scribe (merge into decisions.md)  
**Status:** Landed on branch `ronniegeraghty/prompt-grader-checks`  
**Scope ref:** Morpheus's scope §1, §4, §5  
**Commit:** `2949f578`

## What landed

Implemented Morpheus's YAML prompt grader `checks:` field per scope §1, §4, §5. YAML prompt graders can now declare structured per-check items that produce one `GraderPoint` (and one nested Pass/Fail row in renderers) per check, instead of smuggling numbered lists inside `prompt:` and getting one combined verdict.

## Files touched

**Engine / schema:**
- `hyoka/internal/criteria/config.go` — added `Checks []string` to `UnifiedGraderEntry`; `validateEntry` enforces:
  - `type=prompt`: at least one of `prompt` / `checks` non-empty; each `checks` entry must trim non-empty (errors include the bad index); `details` still forbidden.
  - non-prompt types: `checks` must be empty (mirrors the existing `prompt`-must-be-empty rule).
- `hyoka/internal/criteria/buckets.go` — `FormatUnifiedPromptEntries` honors `Checks`. With checks: parent line `N. **Name**`, optional preamble (entry's `Prompt` field, judge-only), then nested `   1. …   2. …` checks. Without checks: legacy single-line `N. **Name** — Prompt` shape preserved byte-for-byte.
- `hyoka/internal/criteria/graders/prompt_review_grader.go` — added `expectedCriteriaCount` (regex-based leaf-numbered-item count, prefers indented when present) and `logCriteriaCountMismatch` helper. Both `gradePanel` and `gradeSingle` now emit a `slog.Debug` when `len(returnedCriteria) != expected`. No fail; no structural changes.

**YAML migrations (hard-migrated per §2):**
- `criteria/language/python.yaml` — `DefaultAzureCredential Authentication`: split smuggled `1. … 2. …` into preamble `"Check the following criteria:"` + 2 `checks:` items (auth-credential, async/await).
- `criteria/language/test.yaml` — `Markdown Structure`: matches the user's worked example exactly (preamble + 2 checks for hello.md heading and bullet count).
- `criteria/language/java.yaml`, `criteria/language/rust.yaml`, and `hyoka/internal/criteria/testdata/**` — left alone (no embedded numbering; single-`prompt` path covers them).

**Tests:**
- `hyoka/internal/criteria/config_test.go` — replaced `TestValidateEntry_PromptMissingPrompt` with `TestValidateEntry_PromptMissingPromptAndChecks` and added a 6-case table (`TestValidateEntry_PromptChecks`) covering: checks-only, preamble+checks, legacy prompt-only, empty string in checks (with index), whitespace-only check, checks on `output_check`.
- `hyoka/internal/criteria/buckets_test.go` — added `TestFormatUnifiedPromptEntries_Shapes` (5 cases): empty, legacy single-prompt, checks+preamble, checks-no-preamble, mixed legacy+checks. All assert exact rendered text.

## Verification

```
go build ./...                                    # clean
go test ./... -timeout 5m                         # all packages pass
go test -race ./hyoka/internal/criteria/...       # criteria + graders pass
```

Live smoke (`run --prompt-id key-vault-dp-python-crud --config baseline/claude-opus-4.6`, run id `20260424-052914`):

```
grader_results[].grader_name="DefaultAzureCredential Authentication"
  pass=false  score=0.333  points=[
    "Uses DefaultAzureCredential …": pass=true   "1/1 reviewers passed"
    "Uses async/await patterns …":   pass=false  "0/1 reviewers passed"
    "<parent name leaked>":           pass=false  "0/1 reviewers passed"   ← flake
  ]
```

Per-check `Points` propagate end-to-end through `convertGraderResults` → `report.GraderResult.Points`. Tank's interactive renderer should now have data to render the nested rows once the `Pass (X/Y)` badge change lands.

The third "leaked" point is the exact flake §5 anticipated — one judge returned the bucket parent string as a third "criterion". My new debug log fired correctly:

```
DEBUG Review judge returned criterion count differs from sent
  grader="DefaultAzureCredential Authentication" expected=2 returned=3
```

## Decisions kept (didn't re-litigate)

- One LLM call per grader, N checks rendered as N criteria (not per-check sessions).
- Hard-migrate the two YAMLs; no dual-format support.
- Preamble (`prompt:`) is judge-only — nested under the parent line, not surfaced separately to humans.
- Prompt-file path (`ParseEvaluationCriteria` bullet split) untouched; unification lives at the bucket-text layer only.

## Coordination with Tank

File-disjoint. Tank's badge change (`✅ Pass (X/Y)` / `❌ Fail (X/Y)` in `internal/progress/display_interactive.go`) and report-side verification can land independently. Branch `ronniegeraghty/prompt-grader-checks` is ready for Tank to push his commits onto.

## Follow-ups (not in scope)

- Judge sometimes returns the parent grader name as an extra criterion (the smoke flake above). If this becomes noisy, options: (a) tighten review prompt to forbid parent-line scoring, (b) post-filter returned criteria against sent leaf names. Tracking via the new debug log; defer until we have data from more runs.

---

### Decision: Tank — Grader Points Badge Format Alignment & End-to-End Verification (2026-04-24T05:37:56Z)

**Agent:** Tank 📡  
**Scope ref:** Morpheus's scope §6, §7  
**Branch:** `ronniegeraghty/prompt-grader-checks` (shared with Neo)  
**Status:** Implemented, tested, end-to-end verified  
**Commit:** `a47cb97d`

## What changed

1. **Badge format** in `hyoka/internal/progress/display_interactive.go` (`renderGraderWithPoints`):
   - Old: `✅ 2/2 passed` / `❌ 1/3 passed`
   - New: `✅ Pass (2/2)` / `❌ Fail (1/3)` (matches user spec).
   - One-line behavior change; godoc updated.

2. **Soft truncation** of per-Point names in the interactive renderer:
   - `name = truncateToWidth(p.Name, 50)` — long check strings get a `…` ellipsis at ~50 cols.
   - Reuses existing `truncateToWidth` (ANSI-aware, wide-char-aware) — no new helpers, no width plumbing needed.
   - Full text remains in `report.GraderResult.Points[].Name`.
   - Rationale for fixed 50: terminal width isn't piped down to this renderer today, and 50 cols looks clean inside the `    - <name>: ✅ Pass` indent without breaking on standard 100-col terminals. Easy to swap for terminal-width-derived later if Ronnie wants.

3. **Test updates** in `display_interactive_points_test.go`:
   - Three assertions migrated from `"X/Y passed"` to `"Pass (X/Y)"` / `"Fail (X/Y)"`.
   - Single-Point flat-row negative check updated to the new format so it still catches regressions.

## Verifications

- `go build ./...` — clean.
- `go test -race ./internal/progress/... ./internal/report/... -timeout 3m` — green (progress 1.07s, report 29.5s).
- **End-to-end smoke** against Neo's migrated `criteria/language/test.yaml` using `prompts/test/hello-markdown.prompt.md` + `test/haiku` config. Inspected `reports/20260424-052849/.../report.json`:
  - `Markdown Structure` (YAML grader, `type: prompt`, `checks: [2 items]`) → `points: 2` ✅
  - `Criteria from prompt file` (prompt-file path, 3 bullets under `## Evaluation Criteria`) → `points: 3` ✅
  - Both flow through `convertGraderResults` (`hyoka/internal/eval/engine_eval.go:1199-1204`) into `report.GraderResult.Points`. No schema bump needed — Phase 2 v3 already covered this.

## Renderers verified (no changes required)

- `display_interactive.go` — updated (this PR).
- `display_ci.go` — only tracks aggregate pass/total per eval; doesn't render Points individually. Correct as-is.
- No separate `display_json*.go` / `display_quiet*.go` exist; report JSON serialization is owned by `internal/report` and was verified above.

## Out of scope (per Morpheus's scope §6, §7)

- Site / TS template Points rendering — Trinity's domain (Phase 6).
- Schema bump — not needed; v3 already carries Points.
- Terminal-width-aware truncation — fixed 50-col cap is sufficient for now.

## Co-author

Co-authored with Neo on the same branch (`ronniegeraghty/prompt-grader-checks`).
Tank's commits are file-disjoint from Neo's: `hyoka/internal/progress/display_interactive*.go` only.

---

### Decision: Move Taxonomy from Hardcoded to Discovered (2026-04-24T05:37:56Z)

**Agent:** Morpheus 🕶️  
**Issue:** [#635](https://github.com/ronniegeraghty/hyoka/issues/635)  
**Status:** Proposed (issue filed; awaiting implementation)

## Direction

Hyoka is moving away from package-level hardcoded `Valid*` slices in `internal/validate/` toward **runtime discovery** of taxonomy values (`service`, `plane`, `language`, `category`, `difficulty`) from the corpus itself (`prompts/`, `criteria/`, `configs/`).

A new `internal/taxonomy` package will own the walker, the union-of-observed-values set, and a Levenshtein-based "did you mean" suggester. An optional `taxonomy.yaml` at the project root provides forward-declared values (corpus ∪ declared = full allowlist).

## Why future agents should care

- **Don't add new values to `internal/validate/validate.go` or `schema.go`.** Those lists are slated for deletion. If you need to add a value, add a prompt/criteria/config that uses it (or wait for `taxonomy.yaml` to land and declare it there).
- **`isValidLanguage`/`isValidService`/etc. in `schema.go:176-213` will be deleted** along with `isTestValue`. Don't extend them.
- **`cmd/new_prompt.go:14-32` is the only in-repo consumer of the exported `Valid*` slices.** Any new consumer will be deleted with the slices.
- **Site (Trinity) impact is zero for v1** — `internal/serve/dashboard.go` already accepts opaque filter strings. Follow-up: Trinity may want to consume `hyoka taxonomy --json` for empty-state UX, but that's out of scope for #635.

## Suggested owner

Neo (engine/validate work); Morpheus reviews; Oracle docs the `taxonomy.yaml` format once shape is final.

## Out of scope (non-goals for #635)

- Renaming or deprecating existing taxonomy values
- Changing prompt frontmatter schema
- Removing the ID-prefix naming convention check (`schema.go:166`)
- Server-side enforcement on the serve API

---

### Decision: Schema Field Naming — `graders_total` Counts Points (2026-04-24)

**Filed by:** Trinity 🖤  
**Date:** 2026-04-24  
**Status:** For team consideration (not blocking)

## Observation

`hyoka/internal/eval/engine_eval.go:690-691` writes:

```go
evalReport.GradersTotal  = countTotalPoints(agg.Results)
evalReport.GradersPassed = countPassedPoints(agg.Results)
```

The JSON field names are `graders_total` / `graders_passed`, but the values
are POINTS counts (sum of grader sub-checks across all graders), not grader
counts. For the test fixture
`reports/20260424-173723/.../test-dp-test-hello-markdown/test/sonnet/`, the
report has 6 graders and 14 total points. `graders_total = 14`,
`graders_passed = 13`.

## Why it bit us

The site's `evalGraderTotals(r)` returns `{passed: graders_passed, total: graders_total}`
verbatim when the engine totals are present, so the `graders_total` semantic
leaks straight into the UI. The per-eval headline subtitle was rendering
"across 14 graders" — wrong noun for the wrong number. Tank's run-detail
table also names the locals `gradersPassed` / `gradersTotal` (`run-detail-page.tsx:261-262`)
even though they hold points counts.

## Options

1. **Rename schema fields** (v4 bump): `graders_total` → `points_total`, etc.
   Update `report.EvalReport` JSON tags, all callers (engine, site, lib helpers,
   trends, comparison), and bump the schema version. Cleanest but cross-cutting.
2. **Add separate fields**: keep `graders_total` for backward compat (deprecated)
   and add `points_total` / `points_passed` as authoritative. Site reads new
   fields when present.
3. **Status quo + helpers**: leave the wire format alone, document the
   semantic, and standardize on `pointTotals.passed` / `pointTotals.graders`
   from `evalPointTotals(r)` in the React layer (which derives both from
   `grader_results` directly). Trinity used this approach for the immediate
   fix.

## Trinity's recommendation

Option 2 in the medium term — start emitting `points_total` / `points_passed`
alongside the legacy fields, mark the legacy fields deprecated in
`report/types.go`, and migrate the site/trends/comparison readers in one
follow-up PR. Avoids a hard schema break while paying down the naming debt.

## Scope of immediate fix (already shipped on dev)

Only the per-eval detail page subtitle. Run-detail page table left as-is —
its labels say "Score" so the misnamed locals aren't user-visible. Worth a
follow-up rename pass when somebody touches run-detail next.

---

---

# Empirical Verification: Issues #586 and #619

**Agent:** Switch 🤍  
**Date:** 2026-04-24  
**Session:** Verification pass requested by Ronnie  

## Summary

Empirically verified the fix status of two open issues by running live tests:

- **#586 (Builtin skill leakage):** ❌ **NOT FIXED** — `customize-cloud-agent` still loads into eval sessions
- **#619 (Tool load guardrail):** ✅ **VERIFIED FIXED** — 47 green tests, hard-fail path confirmed working

## Issue #586: Builtin Skills Still Leak

### Evidence

Ran: `hyoka run --prompt-id app-configuration-dp-python-crud --config baseline/gpt-5.3-codex --log-level debug --log-file hyoka-586-verify.log`

Log output:
```
time=2026-04-24T18:20:03.928Z level=INFO msg="Skills loaded" ... skills=customize-cloud-agent
```

**Symptom confirmed:** The builtin skill `customize-cloud-agent` loaded into the generation session even though the config doesn't request it.

### Why Morpheus's Analysis Was Wrong

Morpheus cited commit `445fea76` ("Fix user-level skills leaking into eval Copilot sessions") as the fix for #586. That commit:

1. Addresses **USER-LEVEL** skills from `~/.config/github-copilot/` (issue #21)
2. Uses `SessionConfig.ConfigDir` isolation
3. Does NOT address **BUILTIN** skills from `~/.copilot/pkg/universal/{cli-version}/builtin-skills/`

Issue #586 explicitly states:

> "builtin skills are loaded by the CLI binary itself from its install dir, not via ConfigDir — so the isolation is a no-op against them."

The SDK field to use is `SessionConfig.DisabledSkills` — which hyoka never sets (`git grep DisabledSkills` returns zero matches).

### What Needs to Happen

1. Enumerate builtin skill names from CLI install dir at session build time
2. Populate `SessionConfig.DisabledSkills` with all detected builtins
3. Add opt-in surface (via config YAML) if a config wants to allow specific builtins
4. Add tests to verify `customize-cloud-agent` is disabled by default
5. Add debug log line showing what was disabled

**Assignee suggestion:** Neo or Tank (session config builders)

---

## Issue #619: Tool Load Guardrail Working

### Evidence

**Unit tests (23 green):**
```bash
go test ./hyoka/internal/config/tool/... -v -run ValidateAndExpand
```
All pass — covers missing plugins, skills, MCP servers, malformed YAML, remote plugin failures.

**Integration tests (47 green):**
```bash
go test ./hyoka/internal/eval/... -v -run "Test.*[Tt]ool"
```
Including:
- `TestCopilotRunner_ToolLoadFailure_HardFail`
- `TestCopilotRunner_ToolLoadFailure_MissingSkill`
- `TestCopilotRunner_ToolLoadFailure_MCPMissingCommand`
- `TestToolValidationGate_*` suite (8 scenarios)

**Code confirmed:**
- `hyoka/internal/config/tool/validate.go`: `ValidateAndExpand` returns `ToolLoadError`
- `hyoka/internal/eval/copilot.go:175`: calls `ValidateAndExpand` before `CreateSession`
- On failure: sets `ErrorCategory: "tool_load_failure"`, aborts before generation

### Implementation Commits

- `8c947c8a` — hard-fail logic in `buildSessionConfig`
- `5c75b47c` — `ValidateAndExpand` + `ToolLoadReport`
- `557bb83b` — docs
- `05b4f6d8` — table-driven test coverage

**Recommendation:** Close #619 — it's done and thoroughly tested.

---

## Key Lessons

1. **Commit-evidence is not enough:** Reading commit messages without running tests misses semantic distinctions (user-level vs builtin skills).

2. **Log lines are smoking guns:** `skills=customize-cloud-agent` in the log is direct proof of the leak — no ambiguity.

3. **Test discipline wins:** #619 shipped with 47 tests. #586 has no tests because it's not actually fixed yet.

4. **Switch's role is empirical validation:** I don't just read commits or trust summaries — I run the code and prove it works (or doesn't).

---

## Action Items

- [ ] Reopen #586 investigation (or clarify it was never closed)
- [ ] Close #619 with reference to commit `8c947c8a` + test coverage
- [ ] Assign #586 implementation to Neo or Tank (session config work)
- [ ] Add "empirical verification" checkpoint to future issue triage workflows

**Verdict Summary:**

| Issue | Status | Evidence |
|-------|--------|----------|
| #586 | ❌ NOT FIXED | Live eval shows `customize-cloud-agent` still loads |
| #619 | ✅ VERIFIED FIXED | 47 green tests, hard-fail path confirmed |
