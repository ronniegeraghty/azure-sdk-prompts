# Unified Grader Architecture Proposal

**Author:** Morpheus 🕶️ (Lead Architect)
**Date:** 2026-04-22 (recovered 2026-04-23)
**Status:** PROPOSED
**Supersedes:** Issue #622
**Recovery Note:** Original 513-line proposal was deleted before commit. This is a faithful regeneration from code analysis and the locked direction in `decisions.md` (line 3229).

---

## 1. Current State Assessment

hyoka's grading system has two disconnected halves that evolved independently. They share no schema, no loader, and no execution path. This section documents both with file:line citations.

### 1.1 The Criteria Pipeline (`internal/criteria/`)

**Schema** — `hyoka/internal/criteria/criteria.go:29-66`

```go
// criteria.go:29-35
type GraderEntry struct {
    Name    string            `yaml:"name"`
    Weight  float64           `yaml:"weight"`
    Prompt  string            `yaml:"prompt"`
    When    map[string]string `yaml:"when,omitempty"`
    Isolate bool              `yaml:"isolate,omitempty"`
}

// criteria.go:46-51
type GraderGroup struct {
    Name    string            `yaml:"name,omitempty"`
    When    map[string]string `yaml:"when,omitempty"`
    Graders []GraderEntry     `yaml:"graders"`
    Isolate bool              `yaml:"isolate,omitempty"`
}

// criteria.go:61-66
type GraderConfig struct {
    When    map[string]string `yaml:"when,omitempty"`
    Graders []GraderEntry     `yaml:"graders,omitempty"`
    Groups  []GraderGroup     `yaml:"groups,omitempty"`
    Source  string            `yaml:"-"`
}
```

**Loader** — `criteria.go:97-143`:
- `LoadDir()` walks `--criteria-dir`, calls `loadFile()` per YAML.
- `loadFile()` uses `yaml.NewDecoder` with `KnownFields(true)` (line 135) — strict schema enforcement.
- Only the first YAML document is decoded (silent truncation at `---` separators).
- Validation: requires at least one grader or group (line 139-141).

**Matching** — `criteria.go:70-77, 82-94, 148-155`:
- `matchesWhen()` — AND logic, case-insensitive value comparison.
- `mergeWhen()` — hierarchical merge (parent + child, child wins).
- `MatchingGraders()` → `MatchingGradersWithIsolation()` — collects entries that match prompt properties.

**Engine integration** — `engine.go:110, 140, 214-233`:
- `CriteriaDir` option (line 110) — wired to `--criteria-dir` CLI flag.
- `graderConfigs []criteria.GraderConfig` field (line 140) — loaded at engine init.
- `loadCriteria()` (line 214-233) — loads once, stores on engine.
- `mergedCriteria()` / `reviewBuckets()` (engine.go:257-289) — feeds into AI review.

**What it supports:** LLM-prompt graders only. No typed evaluation (file checks, program runs, etc.). Every grader is a name + weight + prompt string sent to the review panel.

**What it owns that must survive:** Hierarchical `when` (file → group → grader), `Isolate`, `Groups`, `KnownFields(true)` strictness, `Source` tracking.

### 1.2 The Typed Graders Pipeline (`internal/graders/`)

**Schema** — `hyoka/internal/graders/types.go:46-59`

```go
// types.go:46-49
type GraderConfigFile struct {
    Graders []GraderConfig `yaml:"graders"`
}

// types.go:52-59
type GraderConfig struct {
    Kind   string    `yaml:"kind"`
    Name   string    `yaml:"name"`
    Config yaml.Node `yaml:"config"`
    Weight float64   `yaml:"weight,omitempty"`
    Gate   bool      `yaml:"gate,omitempty"`
    When   WhenMap   `yaml:"when,omitempty"`
}
```

**Supported kinds** — `types.go:23-32`: `file`, `program`, `prompt`, `behavior`, `action_sequence`, `tool_constraint`, `prompt_review`, `output_check` (8 kinds).

**Factory** — `registry.go:10-97`:
- `NewGrader()` switch-dispatches on `Kind` to construct concrete grader instances.
- `InstantiateGraders()` / `RunGraders()` (registry.go:100-149) — batch execution.

**Type-specific configs** — `types.go:79-132`:
- `FileConfig`, `ProgramConfig`, `PromptConfig`, `BehaviorConfig`, `ActionSequenceConfig`, `ToolConstraintConfig`, `OutputCheckConfig`.
- Each decoded from the `Config yaml.Node` via `DecodeConfig()` (types.go:145-192).

**Loader** — `types.go:195-247`:
- `Parse()` — strict `KnownFields(true)` decoding (line 198).
- `LoadFile()` / `LoadDir()` — walks directory, merges into single `GraderConfigFile`.

**Engine integration** — `engine.go:117, 141, 237-252`:
- `GradersDir` option (line 117) — **not wired to any CLI flag**. Dead code path.
- `pluginGraders []graders.GraderConfig` field (line 141) — loaded at engine init.
- `loadGraders()` (line 237-252) — loads once, stores on engine.

**Execution** — `engine_eval.go:447-462`:
- Phase 1: `ApplicableGraders()` → `InstantiateGraders()` → `RunGraders()`.
- Results appended to `allGraderResults`.

**What it supports:** All typed evaluations (file existence, program execution, LLM prompt, behavior analysis, output checks, etc.). Full `Gate` semantics. Per-grader `When` matching.

**What it lacks:** No hierarchical `when` (file-level, group-level). No `Groups`. No `Isolate`. No CLI flag to load configs. Users cannot reach it.

### 1.3 The Execution Split — `engine_eval.go:421-543`

The engine runs grading in two sequential phases inside `runSingleEval()`:

```
// engine_eval.go:421-424 (comment header)
// Unified grading pipeline (WI-023) — all graders (pluggable + AI review)
// run in a single phase.

// Phase 1: Pluggable graders (engine_eval.go:447-462)
e.pluginGraders → ApplicableGraders → InstantiateGraders → RunGraders → allGraderResults

// Phase 2: AI review grader (engine_eval.go:464-516)
e.reviewerFactory → PromptReviewGrader.Grade() → allGraderResults

// Aggregation (engine_eval.go:518-543)
AggregateResults(allGraderResults)
```

Despite the "unified grading pipeline" comment (WI-023), these remain two distinct code paths with different:
- **Schema** — `criteria.GraderConfig` vs `graders.GraderConfig`
- **Loader** — `loadCriteria()` vs `loadGraders()`
- **Storage** — `e.graderConfigs` vs `e.pluginGraders`
- **Matching** — `criteria.MatchingGraders()` vs `graders.ApplicableGraders()`
- **CLI surface** — `--criteria-dir` (wired) vs `GradersDir` (unwired)

### 1.4 Summary of Pain Points

| Concern | Criteria Pipeline | Typed Graders Pipeline |
|---------|-------------------|----------------------|
| User-reachable | ✅ `--criteria-dir` | ❌ No CLI flag |
| Schema | `GraderEntry` (name/weight/prompt) | `GraderConfig` (kind/name/config/weight/gate) |
| Hierarchical `when` | ✅ 3-level | ❌ Flat only |
| Groups | ✅ `GraderGroup` | ❌ None |
| Isolation | ✅ `Isolate` field | ❌ Not applicable |
| Gate semantics | ❌ Not supported | ✅ `Gate` field |
| Typed graders | ❌ Prompt-only | ✅ 8 kinds |
| `KnownFields(true)` | ✅ | ✅ |

---

## 2. Target End State

### Package Layout

```
internal/graders/          # ONE package — absorbs criteria
  criteria.go              # GraderConfig, GraderGroup, UnifiedGraderEntry (was criteria/criteria.go)
  criteria_compat.go       # Back-compat loader, mergeWhen, matchesWhen
  buckets.go               # Review bucket building (was criteria/buckets.go)
  types.go                 # Kind constants, typed config structs (FileConfig, etc.)
  registry.go              # NewGrader() factory
  grader.go                # Grader interface, GraderInput, GraderResult, AggregateResults
  file_grader.go           # FileGrader implementation
  program_grader.go        # ProgramGrader implementation
  output_check_grader.go   # OutputCheckGrader implementation
  behavior_grader.go       # BehaviorGrader implementation
  prompt_grader.go         # PromptGrader implementation
  prompt_review_grader.go  # AI review panel grader (WI-023)
  ...                      # Other grader implementations unchanged
```

**Deleted:** `internal/criteria/` package (Phase 3).

### Schema: ONE `UnifiedGraderEntry`

Single struct replaces both `criteria.GraderEntry` and `graders.GraderConfig`. The `Kind` field is the discriminator: empty = LLM-prompt grader (backward-compatible), non-empty = typed grader.

### Execution: ONE path

The engine loads all graders from `--criteria-dir`, partitions by Kind, and dispatches:
- Empty Kind → existing review panel (criteria text injection)
- Non-empty Kind → `NewGrader()` → `Grade()`
- All results → `AggregateResults()`

### CLI: ONE flag

`--criteria-dir` remains the single user-facing flag. `GradersDir` engine option is removed.

---

## 3. Unified Schema Design

### 3.1 `UnifiedGraderEntry` struct

```go
// UnifiedGraderEntry defines a single grader in the evaluation pipeline.
// The Kind field discriminates between LLM-prompt graders (empty Kind)
// and typed graders (non-empty Kind like "file", "output_check", etc.).
//
// When Kind is empty, the entry behaves exactly like the legacy
// criteria.GraderEntry: Name + Weight + Prompt are sent to the AI review panel.
//
// When Kind is non-empty, the entry is dispatched to NewGrader() and
// executed programmatically. The Config field carries kind-specific YAML.
type UnifiedGraderEntry struct {
    Name    string            `yaml:"name" json:"name"`
    Weight  float64           `yaml:"weight" json:"weight"`
    Prompt  string            `yaml:"prompt,omitempty" json:"prompt,omitempty"`
    When    map[string]string `yaml:"when,omitempty" json:"when,omitempty"`
    Isolate bool              `yaml:"isolate,omitempty" json:"isolate,omitempty"`

    // Typed grader fields — only used when Kind is non-empty.
    Kind   string    `yaml:"kind,omitempty" json:"kind,omitempty"`
    Config yaml.Node `yaml:"config,omitempty" json:"config,omitempty"`
    Gate   bool      `yaml:"gate,omitempty" json:"gate,omitempty"`
}
```

**Design rationale:**
- `Kind` is optional (`omitempty`) — backward compatible with every existing criteria file.
- `Prompt` is optional — typed graders don't have one.
- `Config` is optional — prompt graders don't have one.
- `Gate` is optional — only meaningful for typed graders (but allowed on prompt graders for future use).
- `Isolate` is optional — only meaningful for prompt graders (silently ignored for typed graders).
- `Weight` retains the same default behavior (1.0 if unset).

### 3.2 `GraderGroup` (unchanged)

```go
type GraderGroup struct {
    Name    string               `yaml:"name,omitempty"`
    When    map[string]string    `yaml:"when,omitempty"`
    Graders []UnifiedGraderEntry `yaml:"graders"`
    Isolate bool                 `yaml:"isolate,omitempty"`
}
```

### 3.3 `GraderConfig` (top-level file)

```go
type GraderConfig struct {
    When    map[string]string    `yaml:"when,omitempty"`
    Graders []UnifiedGraderEntry `yaml:"graders,omitempty"`
    Groups  []GraderGroup        `yaml:"groups,omitempty"`
    Source  string               `yaml:"-"`
}
```

### 3.4 YAML Examples

**Existing criteria file — zero changes required:**

```yaml
# criteria/language/java.yaml — works unchanged
when:
  language: java
graders:
  - name: Correct Dependencies (com.azure)
    weight: 1.0
    prompt: >
      Uses com.azure group ID for all Azure SDK packages.
  - name: DefaultAzureCredential Authentication
    weight: 1.0
    prompt: >
      Uses DefaultAzureCredential or another com.azure.identity credential.
```

**New typed grader entry:**

```yaml
# criteria/quality/output.yaml — new file
graders:
  - name: Output Files Exist
    kind: output_check
    weight: 0.5
    gate: true
    config:
      min_files: 1
      min_bytes_per_file: 1
  - name: Main File Present
    kind: file
    weight: 0.5
    config:
      path: "main.py"
      must_exist: true
    when:
      language: python
```

**Mixed file — prompt + typed graders together:**

```yaml
# criteria/service/key-vault-quality.yaml
when:
  service: key-vault

groups:
  - name: Structural Checks
    graders:
      - name: Output Files Exist
        kind: output_check
        gate: true
        config:
          min_files: 1

  - name: SDK Quality
    when:
      language: python
    graders:
      - name: Azure SDK Best Practices
        weight: 1.0
        prompt: >
          Uses azure-identity DefaultAzureCredential. No hardcoded secrets.
      - name: Async Pattern
        weight: 0.8
        prompt: >
          Uses aio async clients where appropriate.
        isolate: true
```

---

## 4. Execution Path Unification

### 4.1 Engine changes

**Before (dual path):**

```
Engine.loadCriteria()  → e.graderConfigs []criteria.GraderConfig   ← --criteria-dir
Engine.loadGraders()   → e.pluginGraders []graders.GraderConfig    ← GradersDir (unwired)
```

**After (single path):**

```
Engine.loadGraders()   → e.graderConfigs []graders.GraderConfig    ← --criteria-dir
```

The engine stores a single `[]graders.GraderConfig` loaded from `--criteria-dir`. No second loader, no second field, no second option.

### 4.2 Partition-and-dispatch in `runSingleEval()`

```
                    ┌──────────────────────┐
                    │  e.graderConfigs      │
                    │  (loaded from         │
                    │   --criteria-dir)     │
                    └──────────┬───────────┘
                               │
                     MatchingGraders(configs, props)
                               │
                    ┌──────────┴───────────┐
                    │  matched entries      │
                    └──────────┬───────────┘
                               │
              ┌────────────────┼────────────────┐
              │                │                 │
        Kind == ""      Kind == "prompt"    Kind != "" (typed)
        (LLM-prompt)    (standalone LLM)
              │                │                 │
    ┌─────────┴──────┐  ┌─────┴──────┐  ┌──────┴──────┐
    │ Build criteria  │  │ NewGrader()│  │ NewGrader() │
    │ text → review   │  │ → Grade()  │  │ → Grade()   │
    │ panel           │  └─────┬──────┘  └──────┬──────┘
    └────────┬───────┘        │                 │
             │                 │                 │
             └────────────────┼─────────────────┘
                              │
                    AggregateResults(allResults)
```

**Pseudocode:**

```go
func (e *Engine) runGraders(ctx context.Context, matched []UnifiedGraderEntry, ...) {
    var promptEntries []UnifiedGraderEntry
    var typedEntries  []UnifiedGraderEntry

    for _, entry := range matched {
        if entry.Kind == "" {
            promptEntries = append(promptEntries, entry)
        } else {
            typedEntries = append(typedEntries, entry)
        }
    }

    var allResults []graders.GraderResult

    // Typed graders: instantiate and run
    for _, entry := range typedEntries {
        gc := entry.ToGraderConfig()
        g, err := graders.NewGrader(gc)
        // ... grade and collect result
        allResults = append(allResults, result)
    }

    // Prompt graders: build criteria text → review panel
    if len(promptEntries) > 0 && !e.opts.SkipReview {
        criteriaText := FormatGraders(promptEntries)
        // ... run review panel with criteria
        allResults = append(allResults, reviewResult)
    }

    agg, _ := graders.AggregateResults(allResults)
    // ... update report
}
```

### 4.3 What `ToGraderConfig()` does

The `UnifiedGraderEntry` needs a conversion method to produce the `graders.GraderConfig` expected by `NewGrader()`:

```go
func (u *UnifiedGraderEntry) ToGraderConfig() graders.GraderConfig {
    return graders.GraderConfig{
        Kind:   u.Kind,
        Name:   u.Name,
        Config: u.Config,
        Weight: u.Weight,
        Gate:   u.Gate,
        When:   graders.WhenMap(u.When),
    }
}
```

---

## 5. Backward-Compatibility Guarantees

### 5.1 Zero-change files

Every existing `criteria/*.yaml` file works without modification on day 1 of Phase 1. This is guaranteed because:

1. **Schema superset:** `UnifiedGraderEntry` is a strict superset of `criteria.GraderEntry`. Fields `name`, `weight`, `prompt`, `when`, `isolate` are identical in name, type, and YAML tag.

2. **`KnownFields(true)` preserved:** The new loader uses the same `yaml.NewDecoder` + `KnownFields(true)` pattern (currently at `criteria.go:134-135`). Any YAML key not in the struct is still a hard error.

3. **Empty `Kind` = prompt grader:** Files with no `kind` field produce entries with `Kind == ""`, which are routed to the review panel exactly as today.

4. **`Isolate` behavior unchanged:** Prompt graders with `isolate: true` still get dedicated review sessions. The review bucket building (`buckets.go`) is moved but not modified.

5. **Hierarchical `when` unchanged:** `mergeWhen()`, `matchesWhen()`, and the file → group → grader resolution chain are moved to `internal/graders/` but remain byte-identical.

6. **Weight defaults unchanged:** `Weight == 0` still defaults to 1.0 in aggregation.

### 5.2 What changes for users

- **Nothing** in Phase 1-2. The CLI, flag names, criteria file format, and scoring behavior are identical.
- **Phase 3** deletes `internal/criteria/` — purely internal cleanup, no user-facing change.
- **Phase 4** adds new criteria files using typed graders — additive only.

---

## 6. Open Questions for Ronnie

### Q1: CLI flag naming — keep `--criteria-dir` or rename?

**Question:** The flag `--criteria-dir` was named when the directory only held LLM-prompt criteria. With typed graders living alongside, should we rename to `--graders-dir`?

**Options:**

| Option | Pros | Cons |
|--------|------|------|
| A) Keep `--criteria-dir` | Zero breaking change. Existing scripts, docs, CI all work. | Semantically inaccurate — the dir now holds graders, not just criteria. |
| B) Rename to `--graders-dir` + alias `--criteria-dir` | Semantically correct. Alias preserves backward compat. | Adds complexity. Two names for one flag in `--help` output. |
| C) Rename to `--graders-dir`, deprecate `--criteria-dir` with warning | Clean long-term. Clear migration path. | Breaking change (even if soft) for existing users. |

**Recommendation:** Option A — keep `--criteria-dir`. The name is already established, and "criteria" is still accurate (typed graders are evaluation criteria). Rename is noise for zero user value.

### Q2: `Kind` discriminator placement

**Question:** Should `Kind` be a bare optional field on every `UnifiedGraderEntry`, or should typed grader fields live in a separate `typed:` block?

**Options:**

| Option | Pros | Cons |
|--------|------|------|
| A) Bare field (`kind:` at entry level) | Flat YAML, easy to read. `KnownFields(true)` validates naturally. | Prompt graders see `kind`, `config`, `gate` as valid-but-unused fields. |
| B) Separate `typed:` block | Clean separation. Prompt graders don't see typed fields. | Nested YAML, more complex schema, harder to explain, breaks `KnownFields(true)` cleanly. |

**Recommendation:** Option A — bare field. The `omitempty` tags already handle this cleanly. Users writing pure prompt criteria never see `kind` in their files (it's not required). `KnownFields(true)` continues to reject truly unknown fields. This matches how Go's `encoding/json` and `gopkg.in/yaml.v3` handle optional fields industry-wide.

### Q3: `Gate` semantics for typed graders

**Question:** When a typed grader has `gate: true` and fails, does it hard-fail the entire evaluation (score = 0), or soft-warn?

**Options:**

| Option | Pros | Cons |
|--------|------|------|
| A) Hard fail (score = 0, eval fails) | Consistent with current `AggregateResults` gate logic (`grader.go:204-210`). Predictable. | A misconfigured file grader could block all evals silently. |
| B) Soft warn (log warning, reduce score but don't zero) | Safer for experimentation. | Inconsistent with existing gate semantics. Users expect "gate" to mean hard fail. |
| C) Hard fail with `--strict-gates` opt-in | Safe default, power-user escape hatch. | Extra flag complexity. |

**Recommendation:** Option A — hard fail. `Gate` already means hard fail in `AggregateResults` (line 204-210). Changing semantics would surprise users of existing typed graders. The fix for misconfiguration is validation (`hyoka validate`), not weakened gates.

### Q4: `Isolate` on typed graders

**Question:** Typed graders (file, program, output_check) don't use Copilot sessions. What happens if someone puts `isolate: true` on a typed grader?

**Options:**

| Option | Pros | Cons |
|--------|------|------|
| A) Silent ignore | Simple. No-op fields are common in YAML configs. | User might think isolation is happening when it isn't. |
| B) Load-time warning (log, continue) | Alerts to misconfiguration. Non-breaking. | Log noise for intentional mixed files. |
| C) Load-time error (reject file) | Strictest validation. No ambiguity. | Prevents mixed files where groups contain both prompt + typed entries. |

**Recommendation:** Option B — load-time warning. Log `slog.Warn("isolate has no effect on typed grader", "name", entry.Name, "kind", entry.Kind)` during validation. This follows the existing pattern of `criteria.go:113-114` where invalid files get a warning but don't block the run.

### Q5: Multiple instances of the same `Kind` in one file

**Question:** Should a single criteria file be allowed to have two `output_check` graders (different configs, different `when` conditions)?

**Options:**

| Option | Pros | Cons |
|--------|------|------|
| A) Allowed (names must be unique) | Flexible. Enables per-language output checks in one file. | Possible user confusion. |
| B) Rejected (one instance per Kind per file) | Simple to reason about. | Arbitrary restriction. Forces file-per-grader for multi-condition setups. |

**Recommendation:** Option A — allowed, with unique name enforcement. The current `graders.Validate()` (types.go:256-268) already enforces name uniqueness. Kind uniqueness would be an artificial limitation with no schema benefit.

### Q6: `internal/criteria/` deletion timing

**Question:** Phase 3 deletes `internal/criteria/`. Should this happen immediately after Phase 2, or should we keep a deprecation shim for one release?

**Options:**

| Option | Pros | Cons |
|--------|------|------|
| A) Immediate deletion in Phase 3 | Clean. No shim maintenance. | Any external code importing `internal/criteria` breaks (but it's `internal/`, so no external consumers). |
| B) Deprecation shim for one release | Safer if anyone copies internal code patterns. | Maintenance burden. Shim code is dead weight. |

**Recommendation:** Option A — immediate deletion. The `internal/` path prevents external imports per Go convention. All consumers are within `hyoka/`. Phase 2 already moves all callers. There are no external dependents to protect.

### Q7: `output_check` v1 schema knobs

**Question:** The current `OutputCheckConfig` (types.go:112-124) has `MinFiles`, `MinBytesPerFile`, `MinTotalBytes`. For the user-facing `criteria/quality/output.yaml`, what knobs should we expose?

**Options to include/exclude:**

| Knob | Current | Ship in v1? | Notes |
|------|---------|-------------|-------|
| `min_files` | ✅ (default 1) | ✅ Yes | Core use case — "did the agent produce files?" |
| `min_bytes_per_file` | ✅ (default 1) | ✅ Yes | Catches empty stubs |
| `min_total_bytes` | ✅ (default 0) | ❓ Maybe | Niche — rarely configured |
| `required_extensions` | ❌ Not yet | ❓ Maybe | e.g., must produce `.py` files |
| `glob_filter` | ❌ Not yet | ❓ Maybe | e.g., only count `src/**/*.py` |
| `non_empty_content` | ❌ Not yet | ❌ Defer | Overlap with `min_bytes_per_file >= 1` |

**Recommendation:** Ship `min_files` + `min_bytes_per_file` as the v1 schema. Defer `required_extensions` and `glob_filter` to a follow-up issue. `min_total_bytes` can stay in the schema but is not promoted in docs. Keep the schema extensible — future knobs are additive under `KnownFields(true)`.

### Q8: Test strategy — golden-file regression vs new test harness

**Question:** How do we verify that the unified loader produces byte-identical results to the old criteria loader for every existing criteria file?

**Options:**

| Option | Pros | Cons |
|--------|------|------|
| A) Golden-file regression tests | Direct comparison. Catches any drift. Easy to maintain. | Requires capturing current output as golden files. |
| B) Parallel-run comparison (old loader vs new loader in same test) | Proves equivalence dynamically. No golden files to maintain. | Requires keeping old loader code alive during Phase 1-2. |
| C) Both | Maximum confidence. | More test code. |

**Recommendation:** Option C — both. Phase 1 uses parallel-run (old loader + new loader, assert equal output). Phase 3 removes the old loader and switches to golden-file snapshots captured during Phase 2. This gives us dynamic equivalence proof during transition and static regression protection afterward.

### Q9: `Gate` on prompt graders (future consideration)

**Question:** Should `gate: true` be supported on prompt graders (Kind == "")? If a prompt grader with `gate: true` scores below threshold, should it hard-fail the eval?

**Options:**

| Option | Pros | Cons |
|--------|------|------|
| A) Support now | Unified semantics. All graders support all flags. | Review panel scoring is fuzzy — hard-failing on a subjective LLM score is risky. |
| B) Reject at load time | Simple. Clear boundary: gates are for objective checks. | Limits future use cases. |
| C) Accept but warn, defer implementation | Schema-ready. No behavior change. | "gate: true" does nothing, which is confusing. |

**Recommendation:** Option B for now — reject `gate: true` on prompt graders at validation time. Gate semantics make sense for deterministic graders (file exists, program passes). LLM review scores are inherently noisy. Revisit when we have confidence calibration for prompt graders.

### Q10: `action_sequence` and `tool_constraint` in user-facing criteria

**Question:** These grader kinds analyze agent behavior (tool call sequences, tool usage constraints). Should they be documented for user-facing criteria files, or kept as internal-only?

**Options:**

| Option | Pros | Cons |
|--------|------|------|
| A) Document for users | Full power. Users can enforce "agent must use az CLI before creating resources." | Complex. Requires understanding session action logs. |
| B) Internal only (no docs, no examples) | Simpler user experience. | Limits power users. |

**Recommendation:** Option A — document them, but in an "Advanced" section. The schema already supports them. Users who need behavioral checks shouldn't have to fork hyoka. Gating behind docs complexity is sufficient.

---

## 7. Phased Rollout

### Phase 1: Unified Schema + Back-Compat Loader

**Deliverables:**
- `UnifiedGraderEntry` struct in `internal/graders/criteria.go`
- `GraderGroup` and `GraderConfig` (top-level) in `internal/graders/criteria.go`
- Back-compat `LoadDir()` that accepts both old criteria format and new typed format
- `mergeWhen()`, `matchesWhen()`, `MatchingGraders()` moved to `internal/graders/`
- Review bucket building moved to `internal/graders/buckets.go`
- Parallel-run test: old `criteria.LoadDir()` vs new `graders.LoadDir()` produce identical matching results for all files in `criteria/`
- `KnownFields(true)` strictness verified in loader tests

**Definition of done:** `go test ./internal/graders/... -run TestBackCompat` passes. All existing criteria files load without error. Existing engine behavior is identical.

### Phase 2: Unified Execution Path

**Deliverables:**
- Engine stores single `[]graders.GraderConfig` (not dual `graderConfigs` + `pluginGraders`)
- `loadCriteria()` replaced by unified `loadGraders()` using `--criteria-dir`
- `GradersDir` engine option removed
- `runSingleEval()` uses partition-and-dispatch (prompt entries → panel, typed entries → `NewGrader()`)
- `mergedCriteria()` and `reviewBuckets()` rewritten to use unified types
- End-to-end eval test: run a prompt with both criteria and typed graders in the same file

**Definition of done:** `go test ./... -race` passes. Live eval `hyoka run --prompt-id key-vault-dp-python-crud --config baseline/claude-opus-4.6` produces equivalent report to pre-change baseline.

### Phase 3: Delete `internal/criteria/`

**Deliverables:**
- Delete `internal/criteria/` package entirely (5 files)
- Update all imports across `internal/eval/`, `cmd/`, etc.
- Remove parallel-run tests (no old loader to compare against)
- Add golden-file snapshot tests for criteria loading
- Verify `go build ./...` and `go test ./... -race` clean

**Definition of done:** `internal/criteria/` directory does not exist. All tests pass. No compile errors.

### Phase 4: Ship `criteria/quality/output.yaml` + Docs

**Deliverables:**
- `criteria/quality/output.yaml` — default output_check grader (gate, min_files: 1)
- Updated docs: `docs/grader-config-schema.md` expanded with unified schema reference
- New doc: `docs/criteria-authoring.md` — how to write criteria files with typed graders
- Examples in `examples/criteria/` for mixed prompt + typed files
- `CHANGELOG.md` entry

**Definition of done:** `hyoka validate` accepts all new criteria files. Docs reviewed. Example files load without error.

---

## 8. Test Strategy

### 8.1 Golden-File Regression

Capture the output of `criteria.MatchingGraders()` for every file in `criteria/` with a standard set of prompt properties (one per language × service combination). Store as JSON golden files. After Phase 2, assert the unified path produces byte-identical JSON.

### 8.2 `KnownFields(true)` Preservation

Test that:
- Files with unknown fields are rejected (e.g., adding `foo: bar` to a criteria file).
- Files with valid new fields (`kind`, `config`, `gate`) are accepted.
- Files with only legacy fields (`name`, `weight`, `prompt`, `when`, `isolate`) are accepted.

### 8.3 Schema Validation Tests

- Empty `Kind` requires `Prompt` to be non-empty.
- Non-empty `Kind` requires `Config` to be non-empty.
- Non-empty `Kind` must be in `validKinds` set.
- `Gate: true` rejected on empty `Kind` entries (per Q9 recommendation).
- `Isolate: true` on non-empty `Kind` emits warning (per Q4 recommendation).
- Duplicate names within a file are rejected.

### 8.4 End-to-End Eval Comparison

Run a real evaluation with the unified engine against the same prompt + config as a known-good baseline report. Compare:
- `GraderResults` field: same grader names, same scores (within tolerance for LLM review).
- `ScoreBreakdown` field: same structure.
- `ReviewPanel` field: present when expected.

### 8.5 Hierarchical `when` Tests

Migrate all existing tests from `internal/criteria/hierarchical_test.go` to `internal/graders/`. Add new tests for typed graders with hierarchical `when`:
- File-level `when` + typed grader in group → typed grader inherits file-level when.
- Group-level `when` + typed grader → typed grader inherits group-level when.
- Grader-level `when` overrides group + file for typed graders.

---

## 9. Issue Plan

This plan supersedes issue #622. File the following issues:

| # | Title | Scope |
|---|-------|-------|
| 1 | **Unified grader schema + back-compat loader (Phase 1)** | `UnifiedGraderEntry` struct, moved loader, parallel-run tests, `KnownFields(true)` validation. |
| 2 | **Unified execution path in engine (Phase 2)** | Partition-and-dispatch in `runSingleEval()`, single `--criteria-dir` loader, remove `GradersDir`, end-to-end eval test. |
| 3 | **Delete `internal/criteria/` package (Phase 3)** | Remove package, update imports, golden-file snapshot tests, build verification. |
| 4 | **Ship `criteria/quality/output.yaml` + docs (Phase 4)** | Default output_check criteria file, schema docs, authoring guide, examples, CHANGELOG. |
| 5 | **`output_check` v1 schema finalization** | Decide on `required_extensions` and `glob_filter` knobs. Document final schema. Standalone issue for schema design review. |

Each issue should reference this proposal and list its phase. Close #622 with a comment linking to the new issues.

---

## 10. Risk & Mitigation

### R1: Scoring regression in existing criteria

**Risk:** The unified loader produces different matching results than the old `criteria.LoadDir()`, causing score changes.

**Mitigation:** Phase 1 parallel-run tests compare old vs new loader output for every criteria file. Phase 2 end-to-end eval comparison catches scoring drift. Golden-file snapshots in Phase 3 lock in expected behavior.

**Containment:** Phased rollout. Phase 1 changes no runtime behavior. Phase 2 is a single PR with A/B eval comparison.

### R2: `KnownFields(true)` breaks on new fields

**Risk:** Adding `kind`, `config`, `gate` to the struct makes `KnownFields(true)` accept them in legacy files, changing the validation surface.

**Mitigation:** This is intentional and correct — the new fields are valid YAML keys. The risk is that users accidentally use `kind:` in a legacy file and get unexpected behavior. Mitigated by validation: if `kind` is non-empty, `config` must also be present.

### R3: Review panel behavior change

**Risk:** Moving criteria text formatting from `criteria.FormatGraders()` to `graders.FormatGraders()` introduces subtle text differences that affect LLM review scores.

**Mitigation:** `FormatGraders()` is moved byte-for-byte. Test with string comparison of formatted output before and after move.

### R4: Import cycle during migration

**Risk:** During Phase 1-2, `internal/eval/` imports both `internal/criteria/` (old) and `internal/graders/` (new), creating potential type conflicts.

**Mitigation:** Phase 1 does NOT change `internal/eval/`. It only adds new code to `internal/graders/`. Phase 2 switches `internal/eval/` from criteria to graders in a single PR. Phase 3 deletes criteria. No phase has both imports active.

### R5: `output_check` gate too strict for some prompts

**Risk:** Shipping `criteria/quality/output.yaml` with `gate: true` on `output_check` hard-fails prompts that intentionally produce no files (e.g., analysis-only prompts).

**Mitigation:** The output_check criteria file uses `when:` conditions to scope to prompts that expect files. Analysis-only prompts get a different `category` and are excluded. Additionally, users can override with prompt-level criteria.

---

*End of proposal. Questions in Section 6 are decision-forcing — each needs a yes/no/option-letter from Ronnie before Phase 1 implementation begins.*
