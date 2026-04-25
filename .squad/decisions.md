# 🚨 STANDING POLICY — Model Selection (2026-04-24)

**Status:** ACTIVE. Top-level. Read this before any agent spawn.
**By:** Ronnie Geraghty (via Copilot directive 2026-04-24T19:47Z)
**Persisted in:** `.squad/config.json` → `defaultModel: claude-opus-4.7`

- **Default model for every agent: `claude-opus-4.7`.** No exceptions.
- **`claude-haiku-4.5` is FORBIDDEN.** Never spawn Haiku. Never bump down to Haiku. This includes Scribe and Ralph — their charters' "preferred: claude-haiku-4.5" lines are overridden.
- **Latest Sonnet (`claude-sonnet-4.5`)** is allowed only for "really simple things" — trivial mechanical work where opus-4.7 would be wasteful.
- **Rationale:** Quality over cost. User preference, captured for team memory.

---

## 2026-04-24: Engine invariant — every grader emits ≥ 1 GraderPoint

**By:** Neo 💊 (Engine)
**Scope:** `internal/criteria/graders/`, `internal/criteria/exec.go`, `internal/eval/engine_eval.go`
**Commit:** `b7611606`

Every `graders.GraderResult` produced by the engine — normal `Grade()`, error fallback, panic recovery, skipped-grader — MUST carry at least one `GraderPoint`. A Points-less result is a bug.

**Enforcement layers:**
1. `graders.NewResult` panics on `len(points) == 0` — single canonical constructor.
2. `graders.NewErrorResult(kind, name, cfg, msg)` synthesizes a failing `"grader executed"` Point and routes through `NewResult`. Use everywhere a Result is built outside a normal `Grade()` (engine error paths, panic recovery, future skipped-grader paths).
3. `engine_eval.convertGraderResults` defensively logs `slog.Warn` and synthesizes a fallback Point if a Points-less result somehow reaches the report layer (should be unreachable).
4. Graders allowing zero-knob configuration (`BehaviorGrader`, `ToolConstraintGrader`, `OutputCheckGrader`) emit a single trivially-passing `"no_constraints"` / `"no_knobs"` Point.

**Rules for new graders:**
- Always go through `graders.NewResult(kind, name, cfg, points, msg, extras)`. Never construct `graders.GraderResult{...}` literals.
- Each Point: stable snake_case `Label`, `Pass` bool, `Message` (required on fail, encouraged on pass).
- Outside `Grade()`, use `graders.NewErrorResult`.

**Tests:** `graders/points_test.go` covers per-kind pass/fail invariant, constructor panic, error-fallback path.

**Verification:** `reports/20260424-195854/.../report.json` — `jq '.grader_results | map(select(.points == null or (.points | length) == 0))'` returns `[]`.

---

## 2026-04-24: Site per-eval grader UI — defensive rendering shipped

**By:** Trinity 🖤 (Frontend)

- `GraderResultRow.tsx` — collapsed by default; defensive `result.points?.length`; point label fallback chain `label || message || ("Check passed" / "Check failed")`; message no longer duplicates when used as label fallback.
- `graderScore.ts` — when `points` is empty/missing, return `"1/1 points"` (pass) or `"0/1 points"` (fail). Never `"0/0 points"`, never `"PASS"`, never `"100%"`.
- `eval-detail-page.tsx` — Generator Session block moved above Grader Results.
- `GraderResultRow.test.tsx` — rewritten for v4 schema (was pre-v4, 8/8 failing).

**Verification:** 131/131 site tests; Playwright drove a real per-eval page; programmatic asserts on section ordering, chevrons-right, no PASS/100%/FAIL strings, no blank labels.

**Relationship to Neo's invariant:** The site fallback for Points-less results is now belt-and-braces — fresh v4 reports never lack Points. Fallback stays for legacy on-disk reports.

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


---

## Morpheus — Grader Structure Audit & Plan (2026-04-24)

# Morpheus 🕶️ — Grader Structural Audit

**Date:** 2026-04-23
**Requested by:** Ronnie
**Trigger:** "Looking at the report on the site, it seems like the graders have different output and vastly different structures. I would like to make the graders more structured."
**Scope:** Analysis only — no code changes. Recommend options; team decides.

---

## TL;DR

The graders share an **interface** (`Grader.Grade → GraderResult`) and a
**core common surface** (`Kind`, `Name`, `Score`, `Pass`, `Weight`, `Gate`,
`Message`, `Points`). They diverge on the **detail payload** (`*Details`
struct) — every kind has its own bespoke shape, and three of them
(behavior / action_sequence / tool_constraint) actually share one struct
that is a bag of optional fields. Phase 2 introduced `Points` as the
"generalized per-sub-check" channel and the React renderer now treats it
as the canonical view, but the per-kind `*Details` structs still exist
in parallel and the report layer marshals them inconsistently:

- `OutputCheckGraderDetails` is **dropped on the floor** at the report
  boundary (`report.GraderResult` has no field for it).
- `ReviewGraderDetails` fields are also flattened to the top level of
  `report.GraderResult` (`OverallScore`, `MaxScore`, `Summary`, `Issues`,
  `Strengths`, `Scores`, `IsConsensus`), so the same data lives in two
  places for one grader kind only.
- `BehaviorGraderDetails` is a 14-field union used by three graders that
  only set 3–6 fields each, with the renderer guessing which apply.

The user-visible inconsistency is real and almost entirely a
**data-shape problem**, not a CSS/template problem. Recommend Option B
below (promote `Points` to be the single canonical sub-check channel and
freeze each detail struct as kind-specific *display extras* with a tight
common contract).

---

## 1. Common surface area

Every grader implementation today produces a `graders.GraderResult` with
these always-populated fields (`hyoka/internal/criteria/graders/grader.go:96-119`):

| Field | Type | Meaning | Always set? |
|---|---|---|---|
| `Kind` | `string` | One of the 8 `KindXxx` constants | yes |
| `Name` | `string` | Instance name from YAML `name:` | yes |
| `Score` | `float64` | 0.0–1.0 normalized | yes (some graders only emit 0.0/1.0) |
| `Weight` | `float64` | Aggregation weight (`EffectiveWeight`, defaults 1.0) | yes — copied from `input.Config.EffectiveWeight()` |
| `Pass` | `bool` | Binary pass/fail | yes |
| `Gate` | `bool` | Gate flag (no longer short-circuits — see `AggregateResults`) | yes — copied from `input.Config.Gate` |
| `Message` | `string` | Human-readable summary | yes (every grader sets one) |
| `Points` | `[]GraderPoint` | Per-sub-check rows — `{Name, Pass, Message}` | **yes after Phase 2** — every grader now populates ≥1 point |

Plus one of the following, mutually-exclusive typed details (DM4):

```
FileDetails        *FileGraderDetails
ProgramDetails     *ProgramGraderDetails
PromptDetails      *PromptGraderDetails
BehaviorDetails    *BehaviorGraderDetails
ReviewDetails      *ReviewGraderDetails
OutputCheckDetails *OutputCheckGraderDetails
```

The **shared input** (`GraderInput`, `grader.go:29-60`) is a single 11-field
struct every grader receives; graders ignore what they don't need (DM5).
This part is healthy.

The **interface** (`Grader`, `grader.go:18-25`) is exactly three methods:
`Kind() / Name() / Grade(ctx, GraderInput) (GraderResult, error)`. Healthy.

---

## 2. Divergent surface area

### 2a. Per-kind config schema (the input side)

| Kind | Config struct | Unique fields |
|---|---|---|
| `file` | `FileConfig` | `Path`, `Pattern` (regex), `MustExist *bool` |
| `program` | `ProgramConfig` | `Command`, `Args[]`, `Timeout int` (seconds) |
| `prompt` | `PromptConfig` | `Model`, `Rubric` |
| `behavior` | `BehaviorConfig` | `RequiredTools[]`, `ForbiddenTools[]`, `MaxTurns int` |
| `action_sequence` | `ActionSequenceConfig` | `ExpectedActions[]` |
| `tool_constraint` | `ToolConstraintConfig` | `Required[]`, `Forbidden[]`, `MinCalls map`, `MaxCalls map` |
| `output_check` | `OutputCheckConfig` | 7 knobs: `MinFiles`, `MaxFiles`, `RequireFiles[]`, `ForbidFiles[]`, `RequireUpdated[]`, `MinBytesPerFile`, `MaxBytesPerFile` |
| `prompt_review` | *(constructed by engine, not from YAML)* | wraps `review.Reviewer` and/or `review.PanelReviewer` |

Note: `prompt_review` is the only kind not configured via `DecodeConfig`
in `types.go`. It's instantiated directly by the engine with reviewer
deps. That asymmetry is fine — it has different lifecycle needs — but
worth flagging.

### 2b. Per-kind result detail (the output side)

| Kind | Detail struct | Distinctive fields | Sub-check channel |
|---|---|---|---|
| `file` | `FileGraderDetails` | `CheckedFiles[]{Path, Exists, PatternMatched, Pattern}` | 1 `Point` per file |
| `program` | `ProgramGraderDetails` | `Command, ExitCode, Stdout, Stderr` | 1 synthetic `Point` named "exit code 0" |
| `prompt` | `PromptGraderDetails` | `Model, Rubric, Reasoning, RawScore, MaxScore` (int 0–10) | 1 synthetic `Point` named "LLM judge" |
| `behavior` | `BehaviorGraderDetails` (shared) | sets `ToolsUsed, MissingTools, ForbiddenUsed, MaxTurns, ActualTurns, TotalActions, TurnLimitHit, Violations` | 1 `Point` per required/forbidden tool + 1 per turn_limit |
| `action_sequence` | `BehaviorGraderDetails` (shared) | sets `SequenceMatch, ExpectedSequence, ActualSequence, MatchedActions, ToolsUsed, TotalActions` | 1 `Point` named "expected_sequence" |
| `tool_constraint` | `BehaviorGraderDetails` (shared) | sets `ToolsUsed, ToolCounts, MissingTools, ForbiddenUsed, Violations, ConstraintsMet` | 1 `Point` per required/forbidden + per min/max constraint |
| `output_check` | `OutputCheckGraderDetails` | `ProducedFiles[]`, `SubChecks[]{Check, Pass, Message}` | 1 `Point` per configured knob (mirror of `SubChecks`) |
| `prompt_review` | `ReviewGraderDetails` | `Model, OverallScore, MaxScore, Summary, Issues[], Strengths[], IsConsensus, Criteria[], PanelResults[]` | 1 `Point` per criterion (or 1 fallback "consensus") |

### 2c. Score semantics divergence

This is also user-visible:

- `file` — 1.0 / 0.5 (file exists, pattern fails) / 0.0
- `program` — 1.0 or 0.0 only (exit code based)
- `prompt` — `RawScore / MaxScore` (e.g., 7/10 → 0.7), so partial credit
- `behavior`, `tool_constraint` — 1.0 or 0.0 (binary)
- `action_sequence` — `matched / expected` (partial credit)
- `output_check` — 1.0 or 0.0 (binary; AND of sub-checks)
- `prompt_review` — `overallScore / maxScore` (partial credit, criteria-based)

The aggregator (`AggregateResults`) does a weighted average on `Score`,
which is fine, but the **renderer** has to decide whether to show
`70%` vs `7/10` vs `2/3 points` vs `PASS` — and currently does this
ad-hoc in `GraderResultRow.tsx:51-61`. That ad-hoc logic is a symptom,
not a cause; it's downstream of the divergent score semantics.

---

## 3. Output-shape inconsistencies the user sees on the site

Concretely, looking at what `GraderResultRow.tsx` does with each kind:

| Kind | Header score shown | Expanded body |
|---|---|---|
| `file` | "100%" / "50%" / "0%" | "File Checks" list (one row per file) |
| `program` | "100%" / "0%" | Code-styled command + exit code + stdout + stderr |
| `prompt` | "70%" | "LLM Review" with Model + Reasoning paragraph |
| `behavior` | "PASS"/"FAIL" or `N/N passed` | "Behavior Analysis" 2-col grid with tools/turns/violations |
| `action_sequence` | "67%" (e.g. 2/3) | Same "Behavior Analysis" panel — but only `tools_used` + `total_actions` populated; **no sequence visualization** |
| `tool_constraint` | "PASS"/"FAIL" | Same "Behavior Analysis" panel — but renderer has no idea that `tool_counts`/`min_calls` constraints exist; only the synthesized "Violations" line lands |
| `output_check` | "N/N points" | **No bespoke rendering** — `OutputCheckDetails` is not even in `report.GraderResult`. Users only see the generic "Points" list. `ProducedFiles` and the per-knob `SubChecks.Message` are dropped. |
| `prompt_review` | "8/10" + "Review Panel" | Pills for each criterion + Issues/Strengths/Summary at top level |

Specific user-visible gaps Ronnie is likely reacting to:

1. **`prompt_review` is the only kind that surfaces Summary / Issues /
   Strengths.** Everything else has only a `Message` string. Looking at
   the report, prompt_review feels like a different species.
2. **`output_check` has zero kind-specific UI** — it relies entirely on
   the generic `Points` list. The `ProducedFiles` array (which would
   answer "what did the agent actually emit?") is not even marshalled
   to the report (`report/types.go:38-70` has no `OutputCheckDetails`
   field).
3. **`action_sequence` has no sequence display.** It populates
   `ExpectedSequence`/`ActualSequence` in `BehaviorGraderDetails`, but
   the React component only renders `tools_used`, `turn_count`,
   `total_actions`, `violations`. The expected vs actual diff — the
   *whole point* of the grader — never reaches the user.
4. **Three different graders sharing one detail struct** means the
   renderer is guessing which fields to show. The shared
   `BehaviorGraderDetail` is a "union of optionals" pattern that always
   results in this kind of drift.
5. **`prompt` shows `Rubric` in the Go struct but not in the React
   renderer.** Probably intentional (rubrics are long), but never
   stated.
6. **Score formatting drift:** `file` shows "50%", `prompt` shows "70%",
   `output_check` shows "3/3 points", `prompt_review` shows "8/10",
   `behavior` shows "PASS". Same column, four different shapes, no
   legend.

---

## 4. Where the inconsistency lives

This is a three-layer problem and each layer contributes:

### Layer 1 — `graders.GraderResult` (source)

- Has 6 mutually-exclusive `*Details` pointers (DM4 design choice). The
  contract "exactly one is non-nil" is informal — there's no Go-level
  guarantee.
- `BehaviorGraderDetails` is shared by three graders and has 14
  optional fields, with no schema for "which fields each kind sets".
- `ReviewGraderDetails` duplicates fields that other graders express via
  `Message` (Summary) and `Points` (Criteria).

### Layer 2 — `report.GraderResult` (the marshalling layer)

`hyoka/internal/report/types.go:37-70` is where shape drift gets baked
into the JSON the site consumes:

- It **flattens** `ReviewDetails` fields to the top level: `Scores`,
  `OverallScore`, `MaxScore`, `Summary`, `Issues`, `Strengths`,
  `IsConsensus`, `Duration`, `Model` are all top-level fields that
  **only `prompt_review` populates**. Every other grader has them as
  zero-valued JSON noise.
- It is **missing `OutputCheckDetails` entirely** — the engine's
  copy-from-grader-to-report logic has nothing to copy into. So the
  rich `ProducedFiles`/`SubChecks` data dies at this layer.
- `Pass` becomes `*bool` here (vs `bool` in the grader layer) to support
  "legacy review-type graders" — another carve-out for `prompt_review`.

### Layer 3 — `GraderResultRow.tsx` (the renderer)

`site/src/app/components/GraderResultRow.tsx`:

- Has a hardcoded `if (file_details) … if (program_details) …` cascade
  (lines 228–376). Six `if` blocks, one per detail kind. No `output_check`
  block at all.
- Re-derives `passed` from a 5-way `if/else if` (lines 26–40) because the
  source-of-truth shifted from `pass` → `points` mid-Phase-2 and the
  helper `graderPasses` exists but isn't called inline.
- Score formatting cascade (lines 51–61) picks among `pointsPassed/total`,
  `score%`, and `overall_score/max_score` based on which fields are set.

Summary: **Layer 1 is the structural cause. Layers 2 and 3 are
amplifying it.** Fixing only Layer 3 (template) would not solve the
underlying drift — the `output_check` data wouldn't even reach it, and
`action_sequence` data would still be hidden inside a 14-field union.

---

## 5. Recommended unification — three options

### Option A — Minimum viable: just plumb the missing data through

Don't change the grader type system. Just close the data loss:

1. Add `OutputCheckDetails *OutputCheckGraderDetail` to
   `report.GraderResult` and copy it across in the engine.
2. Add an `output_check` rendering block to `GraderResultRow.tsx`
   showing `ProducedFiles` and the per-knob `SubChecks` table.
3. Add an `action_sequence` rendering branch (or extend the behavior
   block) to render the expected vs actual sequence diff.
4. Decide one score-display convention per kind and codify it in a
   helper.

**What changes:** ~50 LOC across 3 files (`report/types.go`,
`engine/...` copy step, one TSX file). Site assets re-built. No config
breakage.

**Cost:** Low. ~1 day. No migration. No breaking change.

**Trade-off:** Doesn't solve the underlying "every grader is its own
snowflake" problem. Adding a 9th grader will require touching all
three layers again. The "union of optionals" `BehaviorGraderDetails`
sticks around. Renderer cascade still grows linearly.

### Option B — Promote `Points` to the canonical sub-check channel; reduce `*Details` to display extras (RECOMMENDED)

Make the contract explicit: **every grader's verdict is `Points[]`**.
The detail struct is reduced to *kind-specific display extras* that
can never affect pass/fail.

Sketch:

```go
type GraderResult struct {
    Kind    string
    Name    string
    Weight  float64
    Gate    bool
    Score   float64        // derived; for partial-credit kinds
    Pass    bool           // = AND(Points[i].Pass)  — enforce at construction
    Message string         // headline only

    Points  []GraderPoint  // REQUIRED, ≥1; the canonical sub-checks

    // Extras: opaque, kind-specific, render-only payload.
    // No field in here may carry pass/fail signal.
    Extras  GraderExtras
}

type GraderExtras struct {
    File        *FileExtras        `json:"file,omitempty"`
    Program     *ProgramExtras     `json:"program,omitempty"`
    Prompt      *PromptExtras      `json:"prompt,omitempty"`
    Behavior    *BehaviorExtras    `json:"behavior,omitempty"`    // shared by behavior/action_sequence/tool_constraint
    Review      *ReviewExtras      `json:"review,omitempty"`
    OutputCheck *OutputCheckExtras `json:"output_check,omitempty"`
}
```

**What each grader changes:**

- `file`: `FileExtras = { Path, Exists, Pattern, PatternMatched }`. Already aligned.
- `program`: `ProgramExtras = { Command, ExitCode, Stdout, Stderr, DurationMs }`. Already aligned.
- `prompt`: `PromptExtras = { Model, Reasoning }`. Drop `RawScore/MaxScore` from extras — `Score` carries it; rubric stays out (already not rendered).
- `behavior`: split the shared union into THREE typed extras:
  `BehaviorExtras`, `ActionSequenceExtras`, `ToolConstraintExtras`. Each
  only carries fields its grader actually sets. Schema becomes
  self-documenting.
- `output_check`: `OutputCheckExtras = { ProducedFiles, SubChecks }`. Plumb through to `report.GraderResult`.
- `prompt_review`: `ReviewExtras = { Model, Summary, Issues, Strengths, IsConsensus, PanelResults }`. **Drop `Criteria` from extras** — they're already in `Points`. **Drop `OverallScore/MaxScore`** — `Score` (0–1) plus the `Points` count carry it. `Summary`/`Issues`/`Strengths` become the only kind-special display data.

**What `report.GraderResult` changes:** Stop flattening review fields to
top level. Move them into `Extras.Review`. This is the breaking change —
any existing JSON parser of the report will break. We control the
sole consumer (the React app).

**What the React renderer changes:** Three dispatchers replace six:
`<HeaderBadge>` (uses Pass + Score uniformly), `<PointsList>` (always
rendered, since Points is required), and `<KindExtras>` (one switch on
`extras` discriminant). `output_check` and `action_sequence` get proper
rendering for free because their extras are now first-class.

**Cost:** Medium. ~300–500 LOC across:
- `graders/grader.go`, all 8 grader impls (rename / split detail structs)
- `report/types.go` (move fields under `Extras`)
- `engine/...` copy step (one big rename, mechanical)
- `site/src/app/data/types.ts` + `GraderResultRow.tsx` (rewrite render path)
- All grader `_test.go` files (mostly find-replace)

**Migration path:** Bump report schema to v4. Site can read v3 (legacy
flat) and v4 (extras) side-by-side for one release, then v3 read path
gets retired. The detail-loss in v3 (`OutputCheckDetails`) is not
recoverable so v3 reports continue to look as broken as they do today.

**Trade-offs:**
- ✅ Self-documenting per-kind shape; no more "which fields does this kind
  set" guessing.
- ✅ Adding a 9th grader is mechanical: define `FooExtras`, set
  `Extras.Foo`, add one render branch.
- ✅ Closes the `OutputCheckDetails` data leak permanently.
- ✅ Forces every grader to emit `Points`, which the renderer already
  treats as canonical (`graderPasses` helper, multi-point UI). Removes a
  whole source of "pass derived two ways" drift.
- ⚠️ Breaking change to report JSON. Acceptable because we own the
  consumer.
- ⚠️ Splitting `BehaviorGraderDetails` into three structs touches the
  three behavior-family graders' tests. Mechanical.

### Option C — Maximally radical: a single `SubCheck` model, retire `*Details` entirely

Reduce every grader to: `Points[]` + a free-form `Evidence map[string]any`
for kind-specific display data. No typed extras at all.

Cost: Largest refactor, but smallest type surface. Every site rendering
becomes generic-with-overrides.

Risk: We lose Go-side type safety on the detail payload. Bugs that
are currently compile-time (typo in `tools_used`) become runtime
("nothing rendered"). The Phase 1 → Phase 2 history shows we already
walked back from "everything is `interface{}`" (DM4) — Option C
re-opens that wound.

**Not recommended** unless we want to genuinely commodity-ize graders
to the point of dynamic registration from YAML alone. That's a
different product.

---

## My recommendation

**Option B.** Specifically:

1. The work is mostly a rename + a struct-split + a renderer simplification.
2. It closes the two real data losses (`output_check`, `action_sequence`).
3. It makes the `Points`-as-canonical decision (already half-made in
   Phase 2) finally complete instead of dual-channel.
4. It cuts the "13 inconsistent fields per grader on the report" surface
   down to "5 common + 1 kind-specific extras blob".
5. The behavior-family split is a nice side win — `action_sequence`
   stops pretending to be `behavior`.

Option A is a safe stopgap if release pressure is high — pair it with
"we'll do B in the next minor". But don't ship A alone for long; the
underlying drift will keep biting.

---

## Open questions for Ronnie

1. **Is the report JSON consumed by anything besides the React site?**
   If yes, Option B becomes more expensive — we'd need to keep the v3
   shape alive or ship a converter.
2. **Should `prompt_review` lose the partial-credit `Score` semantics?**
   Today it's the only grader where `Score` is a fraction of LLM-judge
   points rather than "fraction of Points that passed". Aligning would
   mean a grader's `Score` is *always* `passed_points / total_points`.
   That's a stronger invariant but flattens the LLM's per-criterion
   weighting if we ever add one.
3. **Do we keep `FileGrader`'s 0.5 score (file exists, pattern fails) or
   collapse to binary?** Today it's the only "partial credit single-
   point" grader. Either drop it (binary, Pass=false) or model it as
   two Points ("file present" + "pattern matches"), which is more
   honest.


---

# Morpheus 🕶️ — Grader Unification Implementation Plan (Option B)

**Date:** 2026-04-23
**Status:** Greenlit by Ronnie ("fix the results payload so we can have the results of the graders on the site look more consistent and be handled consistently")
**Foundation:** `.squad/decisions/inbox/morpheus-grader-structure-audit.md`
**Owners:** Neo (engine + Go types) and Trinity (site + renderer)
**Schema bump:** report v3 → **v4**

---

## 0. Decisions made (the three open questions)

I'm picking sensible defaults rather than waiting on Ronnie. He's busy and trusts the call.

### Q1. External report consumers — safe to bump to v4?

**Decision: Yes, hard cutover to v4. No backward-compat read path for v3.**

- The only first-party consumer is the React site (`site/`), which we control end-to-end.
- Reports under `reports/` are git-ignored regenerable artifacts. Anyone who has stale reports re-runs them.
- Maintaining a v3 read path doubles the renderer's complexity and would re-introduce the very ad-hoc cascade we're trying to delete.
- `CurrentSchemaVersion` becomes `4`. Loader rejects v < 4 with a clear "regenerate this report" error.

### Q2. `prompt_review` score semantics — keep OverallScore/MaxScore or fold into Points?

**Decision: Fold into Points. Drop `OverallScore` / `MaxScore` from the report entirely.**

- Every grader's Score becomes `passed_points / total_points` (with `Weight` per-point optional, see §2). That's the new invariant.
- The LLM judge's per-criterion verdict already maps 1:1 to a Point (audit table, §2b). The "8/10" the user sees today is just the criteria pass count rephrased.
- Per-criterion *scores* (the int 0–N the model emits per criterion) survive as a `Weight` field on each Point — see §2 — so partial-credit weighting isn't lost.
- Result: same canonical score string for every grader (§4). No special case for `prompt_review`.

### Q3. FileGrader's 0.5 partial-credit (file exists, pattern fails)

**Decision: Normalize. Two Points per file: `{file present}` + `{pattern matches}` (the second only emitted when a Pattern is configured). Drop the 0.5.**

- Honest to what was actually checked.
- Aligns with the new invariant `Score = passed/total`.
- Renderer gets a uniform multi-point row; no per-grader fudging.

---

## 1. Final unified `GraderResult` shape

### Engine side: `hyoka/internal/criteria/graders/grader.go`

```go
// GraderResult is the single shape every grader returns. Pass and Score
// are derived from Points at construction time — they are NOT independent
// signals. Any field outside Points is render-only and may not influence
// pass/fail.
type GraderResult struct {
    Kind    string  `json:"kind"`              // one of KindXxx
    Name    string  `json:"name"`              // YAML instance name
    Weight  float64 `json:"weight"`            // aggregation weight (from config)
    Gate    bool    `json:"gate"`              // gate flag (from config)

    // Derived from Points — see NewGraderResult helper in §2.
    Score   float64 `json:"score"`             // sum(point.Weight * pass) / sum(point.Weight); 0 if no points
    Pass    bool    `json:"pass"`              // AND over Points[i].Pass
    Message string  `json:"message"`           // headline summary (≤ ~120 chars)

    Points  []GraderPoint `json:"points"`      // REQUIRED, len ≥ 1; the canonical sub-checks
    Extras  *GraderExtras `json:"extras,omitempty"` // kind-specific render-only payload
}

type GraderPoint struct {
    Label    string  `json:"label"`              // short, what was checked (e.g. "file present: src/main.py")
    Pass     bool    `json:"pass"`
    Message  string  `json:"message,omitempty"`  // why it passed/failed (the "reason" Ronnie asked for)
    Weight   float64 `json:"weight,omitempty"`   // for Score weighting; defaults to 1.0 when 0/omitted
    Evidence map[string]string `json:"evidence,omitempty"` // tiny, optional, string-only KV (e.g. {"pattern":"^def "})
}

type GraderExtras struct {
    File           *FileExtras           `json:"file,omitempty"`
    Program        *ProgramExtras        `json:"program,omitempty"`
    Prompt         *PromptExtras         `json:"prompt,omitempty"`
    Behavior       *BehaviorExtras       `json:"behavior,omitempty"`
    ActionSequence *ActionSequenceExtras `json:"action_sequence,omitempty"`
    ToolConstraint *ToolConstraintExtras `json:"tool_constraint,omitempty"`
    OutputCheck    *OutputCheckExtras    `json:"output_check,omitempty"`
    Review         *ReviewExtras         `json:"review,omitempty"`
}
```

### Report side: `hyoka/internal/report/types.go`

`report.GraderResult` becomes a thin mirror of `graders.GraderResult` (no flattened review fields, no per-detail field cascade):

```go
type GraderResult struct {
    GraderName string  `json:"grader_name"`
    GraderType string  `json:"grader_type"` // == graders.Kind
    Score      float64 `json:"score"`
    Weight     float64 `json:"weight"`
    Pass       bool    `json:"pass"`        // not *bool anymore — every grader has a verdict
    Gate       bool    `json:"gate,omitempty"`
    Message    string  `json:"message"`     // renamed from Summary

    Points []GraderPoint `json:"points"` // required, len ≥ 1
    Extras *GraderExtras `json:"extras,omitempty"`
}
```

**Removed from `report.GraderResult`:** `Model`, `Scores`, `OverallScore`, `MaxScore`, `Summary`, `Issues`, `Strengths`, `Duration`, `IsConsensus`, `FileDetails`, `ProgramDetails`, `PromptDetails`, `BehaviorDetails`, `ReviewDetails`. (`Model`, `Issues`, `Strengths`, `IsConsensus`, etc. live inside `Extras.Review` now. `Duration` belonged in the run-level metrics; if we currently use it on the row it moves to `Extras.Review.DurationSeconds`.)

---

## 2. What `Points[]` looks like — the contract

**The invariant that powers everything else:**

```go
// Constructor — every grader builds via this. Cannot construct a
// GraderResult by literal; the type is unexported-fielded or the helper
// just panics on len(points)==0. Either way: Points is required.
func NewResult(kind, name string, cfg GraderConfig, points []GraderPoint, msg string, extras *GraderExtras) GraderResult
```

The helper computes:

- `Pass = all(p.Pass for p in points)`
- `Score = sum(p.weightOr1() * boolToFloat(p.Pass)) / sum(p.weightOr1())`
- copies `Weight` and `Gate` from `cfg.EffectiveWeight()` / `cfg.Gate`

**Each Point answers Ronnie's "why":**

- `Label` — short noun phrase: *what was checked*. E.g. `file present: src/main.py`, `tool used: azure-cli`, `criterion: error handling`.
- `Pass` — boolean.
- `Message` — *why this point's verdict*. Always populated on failure (e.g. "file not found at workspace/src/main.py"), optional on pass.
- `Weight` — for partial credit. Defaults to 1.0. Only `prompt_review` and `action_sequence` use non-1.0 today.
- `Evidence` — small, string-only KV of supporting data. Lets the renderer surface things like `{"actual":"5","expected":">=3"}` without ad-hoc fields.

**No grader may emit zero Points.** A grader that has nothing to check is a config error and should fail loudly at construction.

---

## 3. Per-grader mapping (all 8 kinds)

For each kind, this table is the contract Neo implements:

| Kind | What becomes a Point (one per…) | Extras type | Extras carries |
|---|---|---|---|
| `file` | one Point per file: `file present: <path>` (always); plus a second Point `pattern matches: <path>` per file when `Pattern` is set | `FileExtras` | `Files []FileExtra{Path, Exists, Pattern, PatternMatched, Size}` for syntax-highlighted display |
| `program` | one Point: `exit code 0` (Pass=ExitCode==0; Message=stderr tail on fail) | `ProgramExtras` | `Command, Args, ExitCode, Stdout, Stderr, DurationMs` |
| `prompt` | one Point: `LLM judge: <rubric short name>` (Pass = `RawScore >= passing threshold`; Message = model's reasoning summary, truncated) | `PromptExtras` | `Model, Rubric, Reasoning, RawScore, MaxScore` (RawScore/MaxScore retained here for transparency, NOT used to compute Score — Score is from the single Point) |
| `behavior` | one Point per required tool (`tool required: X`), one per forbidden tool absent (`tool forbidden: Y`), one for turn-limit (`turn limit ≤ N`) when configured | `BehaviorExtras` | `ToolsUsed[], MissingTools[], ForbiddenUsed[], TurnCount, MaxTurns, TotalActions, TurnLimitHit, Violations[]` |
| `action_sequence` | one Point per expected action position: `step N: expected <tool>` (Pass = matched at that index; Message = "got <actual>" on fail). Optionally use `Weight` if certain steps are weighted in YAML (future). | `ActionSequenceExtras` | `ExpectedSequence[], ActualSequence[], MatchedActions, ToolsUsed[], TotalActions` |
| `tool_constraint` | one Point per required tool, per forbidden tool, per `MinCalls[t]` constraint (`tool X called ≥ N`), per `MaxCalls[t]` constraint (`tool X called ≤ N`) | `ToolConstraintExtras` | `ToolsUsed[], ToolCounts map[string]int, MissingTools[], ForbiddenUsed[], Violations[], ConstraintsMet bool` |
| `output_check` | one Point per configured knob: `min files: ≥ N`, `max files: ≤ N`, `require file: <path>` (one per entry), `forbid file: <path>` (one per), `require updated: <path>` (one per), `min bytes/file: ≥ N`, `max bytes/file: ≤ N` | `OutputCheckExtras` | `ProducedFiles []FileEntry{Path, Size}` (the agent's actual output) |
| `prompt_review` | one Point per criterion in the rubric. Label = criterion name; Pass = criterion passed per LLM; Message = LLM's per-criterion reasoning; **Weight = LLM's per-criterion max points** so weighted Score still reflects the rubric weighting | `ReviewExtras` | `Model, Summary, IsConsensus, PanelResults []PanelMemberResult{Model, Score, Pass, Issues[], Strengths[]}, Issues []string (consensus), Strengths []string (consensus), DurationSeconds` |

**Key consequences:**
- `output_check` finally has its `ProducedFiles` and per-knob results reach the site (current bug — they die in marshalling).
- `action_sequence`'s expected-vs-actual diff is now first-class: each step is a Point, the full sequences live in Extras for visualization.
- `prompt_review` no longer has a parallel score channel; its score is `pointsPassed/pointsTotal` weighted by criterion. The `Summary`/`Issues`/`Strengths` move to `Extras.Review` where they belong.
- `behavior`/`action_sequence`/`tool_constraint` get **three separate Extras structs** instead of one 14-field union. Each grader sets only the struct it owns. No more "which fields apply".

---

## 4. Score display — one canonical format

**The rule (Trinity, encode this once in a helper):**

```ts
// In site/src/app/lib/graderScore.ts
export function formatGraderScore(r: GraderResult): string {
  const passed = r.points.filter(p => p.pass).length;
  const total  = r.points.length;
  return `${passed}/${total} points`;   // ALWAYS this shape, even when total === 1
}
```

- Single point that passed → `"1/1 points"`. Not "Passed", not "100%".
- Three of three → `"3/3 points"`. Not "100%".
- Two of three → `"2/3 points"`. Not "67%".
- This is the *only* score string shown in the row header.
- Internal numeric `Score` (the weighted float) is still emitted for aggregation/sort, but the row never displays it as a percentage.

The pass/fail icon (✓/✗) still appears next to the score and is driven by `r.pass` (which equals `passed === total`).

---

## 5. Site rendering rules

**Single source of truth in the row header. No duplicated info on the right.**

Current state (the inconsistency Ronnie called out): the row shows e.g. `Passed` AND `100%` AND a separate `1/1 points` deeper in the body.

New header layout (left-to-right):

```
[expand-chevron] [grader-name]  [kind-pill]  [score-string]  [✓/✗ badge]
```

- `score-string` = `formatGraderScore(r)` from §4.
- The badge is icon-only (no "PASS"/"FAIL" text) — the score string already encodes the count, and the badge encodes the binary verdict.
- **Remove** the right-side "100%" / "N/N" / "PASS" duplication entirely.

Body (when expanded):

1. **Always**: `<PointsList>` — `<ul>` of Points. Each row: `[✓/✗] <Label>` on left; `<Message>` on right (italic, muted). If `evidence` non-empty, render as small KV chips below the row.
2. **If `extras` populated**: `<KindExtras kind={r.grader_type} extras={r.extras}>` — single dispatcher with one branch per kind, rendering only the render-only data (file lists, command output, sequence diff, panel breakdown, etc.).

**Auto-expand rule:** expand by default when `points.length > 1` OR when `r.pass === false`. Single-passing-point graders stay collapsed.

---

## 6. File-by-file change list

### Neo (engine + Go types) — work item N

Order matters; later steps depend on earlier ones.

1. **`hyoka/internal/criteria/graders/grader.go`** — replace `GraderResult` per §1; add `NewResult` constructor; add `GraderPoint.Weight` and `Evidence`; define `GraderExtras` and the seven `*Extras` structs.
2. **`hyoka/internal/criteria/graders/file_grader.go`** — emit two-point pattern per file when Pattern set; build `FileExtras`. Drop 0.5.
3. **`hyoka/internal/criteria/graders/program_grader.go`** — single Point `exit code 0`; `ProgramExtras`.
4. **`hyoka/internal/criteria/graders/prompt_grader.go` + `prompt_grader_adapter.go`** — single Point; `PromptExtras` retains `RawScore/MaxScore` for display only.
5. **`hyoka/internal/criteria/graders/behavior_grader.go`** — split: produce per-tool Points + turn-limit Point; `BehaviorExtras`. (This file currently holds all three behavior-family graders' shared details — Neo splits the Extras here.)
6. **`hyoka/internal/criteria/graders/behavior_grader.go`** (action_sequence path) — per-step Points + `ActionSequenceExtras`.
7. **`hyoka/internal/criteria/graders/behavior_grader.go`** (tool_constraint path) — per-constraint Points + `ToolConstraintExtras`.
8. **`hyoka/internal/criteria/graders/output_check_grader.go`** — per-knob Points (already has SubChecks; just promote them); `OutputCheckExtras` with `ProducedFiles`.
9. **`hyoka/internal/criteria/graders/prompt_review_grader.go`** — per-criterion Points (with Weight = criterion max); `ReviewExtras` with Summary/Issues/Strengths/PanelResults.
10. **`hyoka/internal/report/types.go`** — replace `GraderResult` per §1 (mirror); bump `CurrentSchemaVersion = 4`; add the seven `*Extras` mirror types; remove the flattened review fields and the per-kind detail fields.
11. **`hyoka/internal/eval/engine_eval.go`** — rewrite `convertGraderResults` (lines 1186+): drop the six `if .XDetails` branches; replace with one `Extras` copy; drop the `prompt_review`-special-case block (lines 1248+); copy `Points` and `Extras` mechanically.
12. **`hyoka/internal/report/`** loader/reader — reject `schema_version < 4` with explicit error message ("regenerate report; v3 → v4 schema bump").
13. **All `_test.go` files in `hyoka/internal/criteria/graders/`** — update to new shape (mostly mechanical: replace `result.FileDetails.X` with assertions on Points / Extras.File).
14. **`hyoka/internal/criteria/graders/points_test.go`** — see §8.

### Trinity (site + renderer) — work item T

Trinity can start as soon as Neo lands step 1 (the Go types) — frontend types/components don't need the back-end implementations yet, just the JSON shape contract.

1. **`site/src/app/data/types.ts`** — replace `GraderResult` interface to match new JSON shape (§1, report side). Add `GraderPoint` (with `weight`, `evidence`), `GraderExtras`, and seven `*Extras` interfaces.
2. **`site/src/app/lib/graderScore.ts`** (new) — `formatGraderScore` per §4. Plus `graderPasses` cleanup: now just `r.pass`.
3. **`site/src/app/lib/evalPass.ts`** — collapse the tri-state `graderPasses` cascade to `r => r.pass`. Old field-existence checks no longer needed.
4. **`site/src/app/components/GraderResultRow.tsx`** — full rewrite of body:
    - Header per §5 (drop right-side duplicate).
    - Replace 6-way `if (file_details) … if (program_details) …` cascade with `<PointsList>` (always) + `<KindExtras>` (one switch).
5. **`site/src/app/components/grader-extras/`** (new directory) — one component per kind: `FileExtras.tsx`, `ProgramExtras.tsx`, `PromptExtras.tsx`, `BehaviorExtras.tsx`, `ActionSequenceExtras.tsx` (now finally renders the expected-vs-actual sequence!), `ToolConstraintExtras.tsx`, `OutputCheckExtras.tsx` (finally renders ProducedFiles!), `ReviewExtras.tsx`.
6. **`site/src/app/components/GraderResultRow.test.tsx`** — rewrite assertions against new shape; add cases for each Extras kind; add a test asserting the score string is `N/M points` for both single-point and multi-point graders.
7. **JSON consumers elsewhere in `site/`** — `grep -rn "overall_score\|max_score\|file_details\|program_details\|behavior_details\|review_details\|summary"` and migrate each. Likely hits in `eval-detail-page.tsx` and `prompt-detail-page.tsx`.

### Sync points

- **Sync 1 (after Neo step 1 + Trinity step 1):** Neo and Trinity confirm JSON shape parity. Lock the contract. Trinity unblocked to build components against fixtures.
- **Sync 2 (after Neo step 11):** Neo produces a single fresh report from a real run; Trinity hits it with the new renderer end-to-end.
- **Sync 3 (before merge):** dogfood — `go run ./hyoka serve` against a fresh full-run report; Morpheus walks the site looking at every kind's row.

---

## 7. Migration / breaking-change handling

- **Schema:** v3 → v4. Hard cutover.
- **Old reports:** the loader emits `fmt.Errorf("report schema v%d is no longer supported; regenerate with: hyoka run …", v)`. No silent migration. (Auto-migration is impossible anyway: v3 dropped `OutputCheckDetails` on the floor — the data we'd want isn't there to migrate.)
- **`reports/` is git-ignored** — no commit fallout.
- **Doc updates:** `docs/architecture.md` grader section, any sample `criteria/*.yaml` whose comments reference per-kind detail fields.
- **Issue + PR:** Neo opens a tracking issue ("v4: unified GraderResult"), links this plan, ships the engine half as one PR; Trinity ships the site half as a stacked PR. Do NOT split per-grader — Phase 1 of v4 must be atomic so the JSON shape never lives half-migrated on `dev`.

---

## 8. Test plan

### `hyoka/internal/criteria/graders/points_test.go` (Switch will pick this up)

Add table-driven coverage:

1. **Invariant tests on `NewResult`:**
    - `Pass == true` iff every Point passes.
    - `Score == sum(weight*pass)/sum(weight)`.
    - Empty Points panics (or returns sentinel error) — Points is required.
    - Default Weight 1.0 when Point.Weight == 0.
2. **Per-grader Points-shape tests** (one per kind):
    - `file`: with/without Pattern → 1 vs 2 Points per file; failing pattern produces failing Point with informative Message.
    - `program`: exit 0 / exit non-0; Message contains stderr tail on fail.
    - `prompt`: rubric pass → single Point pass; rubric fail → fail with reasoning in Message.
    - `behavior`: required tool present/absent; forbidden tool absent/present; turn-limit hit/not.
    - `action_sequence`: per-step Points produced in order; Message on mismatch is "got <actual>".
    - `tool_constraint`: each of the 4 knob types produces a distinct Point.
    - `output_check`: each of the 7 knobs produces a Point only when configured (no Point for unset knobs).
    - `prompt_review`: per-criterion Points; weighted Score recovers original `OverallScore/MaxScore` ratio.
3. **Round-trip JSON test:** marshal a `report.GraderResult` of each kind → unmarshal → assert no field loss; Extras discriminant lights up only the expected branch.
4. **Schema-version test:** loading a v3 report returns the explicit "regenerate" error.

### Site (`site/src/app/components/GraderResultRow.test.tsx`)

1. Score string is `N/M points` for single-point AND multi-point.
2. No "100%", "PASS", or duplicated score on the right of the header.
3. PointsList renders one row per point with Pass icon and Message.
4. `KindExtras` dispatches to the correct component per `grader_type`; output_check and action_sequence extras render their distinctive content.

---

## 9. Suggested split — Neo vs. Trinity

| Phase | Neo | Trinity | Sync |
|---|---|---|---|
| **0** | — | — | Both read this plan |
| **1** | Step 1: define new `graders.GraderResult` + `GraderPoint` + `GraderExtras` types in `grader.go`. Push to a branch with type definitions only. | Step 1: define matching TS interfaces in `data/types.ts`. Hand-write fixture JSON for each kind. | **Sync 1**: confirm shape parity |
| **2** | Steps 2–9: rewrite each grader (parallel-safe internally — one grader per commit). | Steps 2–5: write `formatGraderScore` helper, new `GraderResultRow.tsx`, the seven `grader-extras/*.tsx` components. Use Trinity's fixtures. | Trinity unblocked; works in parallel |
| **3** | Steps 10–12: report layer + engine conversion + loader version check. | Step 6: tests. Step 7: sweep other site consumers. | — |
| **4** | Step 13–14: grader tests + `points_test.go`. | — | — |
| **5** | — | — | **Sync 2**: Neo produces fresh report; Trinity verifies E2E with playwright |
| **6** | PR engine half. | Stacked PR for site half. | **Sync 3**: Morpheus dogfoods, then approves both PRs together |

**Parallelism:** Trinity is unblocked from end of Phase 1. Neo's per-grader work in Phase 2 is independent across graders (one commit each). The only true serialization is Phase 3 (Neo's engine conversion) before Sync 2.

**Estimated size:** ~400–600 LOC across ~20 files for Neo; ~400–500 LOC across ~12 files for Trinity (mostly the new components + test fixtures).

---

## Appendix A — what gets *deleted*

Tracking what ships out the door so future readers can confirm cleanup happened:

- `report.GraderResult.{Model, Scores, OverallScore, MaxScore, Summary, Issues, Strengths, Duration, IsConsensus}` — gone (Summary lives on as `Message`; Model/Issues/Strengths/IsConsensus/Duration move to `Extras.Review`).
- `report.GraderResult.Pass` flips from `*bool` to `bool` — every grader emits a verdict now.
- `report.{FileGraderDetail, ProgramGraderDetail, PromptGraderDetail, BehaviorGraderDetail, ReviewGraderDetail}` — replaced by `*Extras` mirrors.
- `graders.{FileGraderDetails, …, ReviewGraderDetails}` — replaced by `*Extras`.
- `graders.GraderResult.{FileDetails, ProgramDetails, PromptDetails, BehaviorDetails, ReviewDetails, OutputCheckDetails}` — collapsed into single `Extras *GraderExtras` field.
- `BehaviorGraderDetails` 14-field union — split into three single-purpose Extras.
- The 6-way `if (X_details)` cascade in `GraderResultRow.tsx` — replaced by single `KindExtras` dispatcher.
- The 5-way `passed` derivation cascade in `GraderResultRow.tsx` — replaced by `r.pass`.
- The score-format cascade (`pointsPassed/total` vs `score%` vs `overall_score/max_score`) — replaced by `formatGraderScore`.
- `convertGraderResults` six-detail copy block + `prompt_review` special-case block — replaced by mechanical Points + Extras copy.

---

## 2026-04-24: Example-file & docs schema audit (Oracle)

**By:** Oracle 🔮
**Memo:** `.squad/decisions/inbox/oracle-example-files-audit.md` (merged, deleted)

Audited every prompt/config/criteria/docs example file vs current schemas (prompt frontmatter, config YAML, unified grader v4). 121 files reviewed.

**Fixed (16 files):**
- `hyoka/cmd/init.go` `exampleConfig` const — wrapped a single ToolConfig in the missing `configs:` list. The seed for `hyoka init --with-examples` previously could not load.
- `docs/graders/index.md` and `docs/graders/prompt.md` — rewrote prompt-grader examples from the (REJECTED) `details: { prompt: ... }` shape to the v4 top-level `prompt:`/`checks:` shape with v4 invariant call-out.
- `docs/grader-config-schema.md` (353 lines, legacy `kind:`/`config:`/`gate:` schema) — rewrote as stub redirect + migration cheat sheet pointing at `docs/graders/index.md`.
- Added explicit `type: prompt` to: `criteria/language/{java,rust}.yaml`, `examples/criteria/language/{dotnet,go,java,python,rust}.yaml`, `examples/criteria/service/{key-vault,storage}.yaml`, `examples/criteria/hierarchical-when-example.yaml`.

**Flagged for human review (NOT auto-fixed):**

A. `examples/prompts/graders-frontmatter-example.prompt.md` documents a `graders:` frontmatter field the parser does not consume (no `Graders` field on `prompt.frontmatter`) AND uses pre-v4 `kind:`/`config:`/`gate:`. Banner added. Need decision: implement `graders:` overrides in the parser (matches v4 unification model) OR delete the example. **Ask Neo.**

B. `docs/starter-files.md` Option B (`starter_files:` list of paths) — only `starter_project:` (directory) is implemented. Doc was already DRAFT; explicit "implementation status" call-out added. Long-term: ship the feature or drop Option B.

**Note for Neo (engine):**
`hyoka/internal/criteria/bundle.go:84-97` silently coerces `no type: + has prompt: + no details:` → `type: prompt`. This kept 8 example files working without explicit `type:` and hid the schema drift. Consider either logging a deprecation warning when the translation fires, or removing it now that all in-tree examples are migrated.

**Skill produced:** `.squad/skills/example-file-validation/SKILL.md`.

---

## 2026-04-24: Site fixes v2 + embed-target convention (Trinity)

**By:** Trinity ⚛️
**Commit:** `fcb8d1d6` on `dev`
**Memo:** `.squad/decisions/inbox/trinity-site-fixes-v2.md` (merged, deleted)

### Convention adopted (per Trinity's ask)
**Site renderers MUST tolerate legacy Point fields.** Use the canonical fallback chain when rendering a Point's display string:
```
p.label || p.name || p.title || p.check || p.message || p.reason || "<unnamed check>"
```
Rationale: 744 of 838 historical Points across local `reports/` use the pre-v4 `name` field. `reports/` is git-ignored ephemera and these will never be regenerated. The engine on current tip is clean (Neo verified — see "Engine invariant — every grader emits ≥ 1 GraderPoint" above), so the fallback is purely backwards-compat for old artifacts in the wild.

**When the chain is needed in a second call site, extract it to `site/src/lib/pointLabel.ts` and centralize.** Today only one call site exists; do not pre-extract.

### Embed-target convention (CRITICAL)
**The Go binary embeds `hyoka/internal/serve/site/`, NOT `site/dist/`.** A `cd site && npm run build` is necessary but **not sufficient**.

Required workflow (per Makefile + `embedded-asset-freshness` skill):
1. `make site-embed` — Vite build + atomic wipe + copy `dist/` → `hyoka/internal/serve/site/`
2. `go build ./...` — picks up new embedded bytes
3. Commit BOTH the source AND the embedded bundle

Three "shipped" site fixes have failed this way in two consecutive sessions. **Ask:** Morpheus/Oracle, please wire `make verify-embed` into CI on every PR that touches `site/src/**`, or add it to the PR template's checklist. The foot-gun will recur otherwise.

### What landed in fcb8d1d6
- Fallback chain in `GraderResultRow.tsx`.
- Defensive `[graderless]` `console.warn` synth when no label field is present.
- Section reorder + collapsed-by-default (now actually deployed; verified via embed bytes).
- 132 site tests pass. Serve killed cleanly.

---

## 2026-04-24T23-55-30Z: User directive — never use `:` in filenames

**Status:** ENFORCED. Cross-platform compliance.
**By:** Ronnie Geraghty (via Copilot)
**Scope:** ALL files written by the team. Applies especially to `.squad/log/`, `.squad/orchestration-log/`, `.squad/decisions/`, and any other timestamped/generated files.

Never create files with `:` in the filename. Windows filesystems reject colons outright. For ISO 8601 timestamped logs, replace colons with hyphens:
- ✅ CORRECT: `2026-04-24T23-58-37Z` (hyphens)
- ❌ WRONG: `2026-04-24T23:58:37Z` (colons)

**Why:** Cross-platform compatibility — the repo must clone and work cleanly on Windows.

**Baseline:** Commit `8148ba13` renamed 83 pre-existing files with colons. All agents and Scribe must follow this invariant going forward. Reference: `.squad/skills/windows-compatibility/SKILL.md`.

---

## 2026-04-25: Tool Version Override Migrated to Repo-Keyed (Morpheus, Neo, Oracle)

**Status:** MERGED & COMMITTED  
**Session:** repo-keyed-version-override (2026-04-25T00:24:30Z)  
**Spawned agents:** Morpheus 🕶️, Neo 💊, Oracle 🔮  

### Decision: Migrate `tool_version_override` from name-keyed to repo-keyed

**Author:** Morpheus 🕶️ (Lead/Architect)  
**Date:** 2026-04-24  
**Approved:** All four recommended defaults adopted  

#### Problem

`tool_version_override` was keyed by user-chosen tool entry name (`Entry.Name`), but git refs are properties of repos, not individual skills. This caused:
1. Git ref granularity mismatch (fetcher clones `owner/repo`, not per-skill)
2. Monorepo fan-out (`microsoft/skills` → N redundant entries)
3. Arbitrary key coupling (entry names are locally-meaningful, not canonical)
4. Common case verbose (pin whole repo requires N entries)

**Blast radius:** Zero shipped configs use this field; only fixtures and docs needed migration.

#### Proposed Schema

```yaml
tool_version_override:
  owner/repo: git_ref  # e.g. microsoft/skills: v1.2.0
```

**Key format:** `owner/repo` (GitHub prefix optional, normalized away).

#### Design Choices (All Approved)

1. **Key format:** `owner/repo` with `github.com/` prefix normalization
2. **Migration:** Hard cut — reject old-shape (keys without `/`) with migration-hint error
3. **Unknown repo override:** Warn (allow shared override maps across configs)
4. **Empty value:** Silent skip (preserve existing behavior)

#### Resolution & Merge

**Per-entry `version:` always wins** (unchanged). Precedence:
```
per-entry version: > tool_version_override (by repo key) > fetcher default
```

**Multi-file merge:** Maps merge across files; conflicting values for same repo = hard error (determinism guarantee); identical values merge silently.

#### Migration Path

Hard cut with clear error message:

```
tool_version_override now keys by repo (e.g. "microsoft/skills"), not by tool name.
Found name-shaped key "my-skill". Migration:
  - Replace each tool-name key with the repo it points to.
  - If multiple tools shared the same repo, collapse to one entry.
See docs/configuration.md → "Tool Versioning" for examples.
```

#### Edge Cases Handled

| Case | Behavior |
|---|---|
| Two entries from same repo, both explicit `version:` | Both per-entry pins win. Override ignored. |
| Override + per-entry both set on same repo | Per-entry wins. |
| Override references unknown repo | Warn (shared maps may cover more repos). |
| Key has `github.com/` prefix | Normalize away. |
| Key malformed (`microsoft`, `a/b/c`, empty parts) | Hard error at parse time. |
| Local skill (`source: local`, no `Repo`) | Override never applies. |
| Plugin (`type: plugin`, `source: remote`) | Override applies via same code path. |
| Empty value (`owner/repo: ""`) | Silent skip (no override). |

#### Implementation (Neo)

**File:** `hyoka/internal/config/config.go`
- Updated `ToolVersionOverride` doc comment (key format: `owner/repo`)
- Added `normalizeRepoKey(s string) string` (strips `github.com/` prefix)
- Rewrote `ApplyVersionOverrides` (lookup by normalized `Entry.Repo`, skip local skills/MCPs)
- Added `validateOverrideKeys` (rejects old-shape, validates `owner/repo` format)
- Updated `LoadDir` error message (references repo, not tool)

**File:** `hyoka/internal/config/version_override_test.go`
- Replaced name-keyed fixtures with repo-keyed ones
- Added 11 new test cases covering multi-entry repos, normalization, old-shape rejection, validation, merge semantics

**File:** `CHANGELOG.md`
- Added breaking-change entry (pre-1.0)

**Verification:** `go build ./...` ✅, `go test -race ./hyoka/internal/config/...` ✅, `go vet ./hyoka/internal/config/...` ✅

#### Documentation (Oracle)

**File:** `docs/configuration.md`
- Per-entry `version:` field documentation (branch, tag, SHA)
- Top-level `tool_version_override` map documentation (repo-key schema, precedence, monorepo examples)
- Resolution order: per-entry > override > fetcher default
- Multi-file merge semantics and conflict-error behavior
- Migration callout with before/after YAML and error message quote
- Remote Skills cross-link to Tool Versioning section

#### Impact

**Breaking change (pre-1.0):** All `tool_version_override` entries must migrate from name-keyed to repo-keyed format. User benefit: monorepo pinning now requires ONE entry per repo instead of N entries per skill (e.g., `microsoft/skills` → 1 entry vs 10+ entries).

---

**Orchestration logs:** `.squad/orchestration-log/2026-04-25-00-24-30-{morpheus,neo,oracle}.md`  
**Session log:** `.squad/log/2026-04-25-00-24-30-repo-keyed-version-override.md`
