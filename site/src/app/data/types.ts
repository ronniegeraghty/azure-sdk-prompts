// ── Types matching the API data shapes ──────────────────────────────

export interface ReviewCriteria {
  name: string;
  passed: boolean;
  reason: string;
}

export interface ReviewScores {
  criteria: ReviewCriteria[];
}

export interface Review {
  overall_score: number;
  max_score: number;
  summary: string;
  strengths?: string[];
  issues?: string[];
  scores?: ReviewScores;
}

export interface ReviewedFile {
  path: string;
  content: string;
}

export interface ReviewPanelEntry {
  model: string;
  overall_score: number;
  max_score: number;
  summary: string;
  scores?: ReviewScores;
  events?: unknown[];
}

// ── Grader system types (v4 unified) ────────────────────────────────

/**
 * GraderPoint is the canonical sub-check unit. Every grader emits at least
 * one Point. Pass and Score on GraderResult are derived from Points.
 * Schema v4.
 */
export interface GraderPoint {
  label: string;
  pass: boolean;
  message?: string;
  weight?: number;
  evidence?: Record<string, string>;
  /** Legacy field name from pre-v4 reports — kept for backward-compat fallback. */
  name?: string;
  /** Some older renderers used `title`. */
  title?: string;
  /** Some older renderers used `check`. */
  check?: string;
  /** Older synonym for message used by some graders. */
  reason?: string;
}

/**
 * File grader extras — per-file details for file existence and pattern checks.
 */
export interface FileExtra {
  path: string;
  exists: boolean;
  pattern?: string;
  pattern_matched?: boolean;
  size?: number;
}

export interface FileExtras {
  files: FileExtra[];
}

/**
 * Program grader extras — command execution details.
 */
export interface ProgramExtras {
  command: string;
  args?: string[];
  exit_code: number;
  stdout: string;
  stderr: string;
  duration_ms?: number;
}

/**
 * Prompt grader extras — LLM-as-judge details.
 */
export interface PromptExtras {
  model: string;
  rubric: string;
  reasoning: string;
  raw_score: number;
  max_score: number;
}

/**
 * Behavior grader extras — tool usage and turn analysis.
 */
export interface BehaviorExtras {
  tools_used: string[];
  missing_tools?: string[];
  forbidden_used?: string[];
  turn_count: number;
  max_turns?: number;
  total_actions: number;
  turn_limit_hit?: boolean;
  violations?: string[];
}

/**
 * Action sequence grader extras — expected vs actual sequence diff.
 */
export interface ActionSequenceExtras {
  expected_sequence: string[];
  actual_sequence: string[];
  matched_actions: number;
  tools_used: string[];
  total_actions: number;
}

/**
 * Tool constraint grader extras — per-tool call constraints.
 */
export interface ToolConstraintExtras {
  tools_used: string[];
  tool_counts: Record<string, number>;
  missing_tools?: string[];
  forbidden_used?: string[];
  violations?: string[];
  constraints_met: boolean;
}

/**
 * Output check grader extras — produced files list.
 */
export interface FileEntry {
  path: string;
  size: number;
}

export interface OutputCheckExtras {
  produced_files: FileEntry[];
}

/**
 * Review grader extras — multi-model panel breakdown.
 */
export interface ReviewExtras {
  model: string;
  summary: string;
  is_consensus?: boolean;
  panel_results?: ReviewPanelEntry[];
  issues?: string[];
  strengths?: string[];
  duration_seconds?: number;
}

/**
 * GraderExtras is a discriminated union — only one branch is populated per grader.
 */
export interface GraderExtras {
  file?: FileExtras;
  program?: ProgramExtras;
  prompt?: PromptExtras;
  behavior?: BehaviorExtras;
  action_sequence?: ActionSequenceExtras;
  tool_constraint?: ToolConstraintExtras;
  output_check?: OutputCheckExtras;
  review?: ReviewExtras;
}

/**
 * GraderResult is the unified shape every grader returns. Pass and Score
 * are derived from Points. Schema v4.
 */
export interface GraderResult {
  grader_name: string;
  grader_type: string;
  score: number;
  weight: number;
  pass: boolean;
  gate?: boolean;
  message: string;
  points: GraderPoint[];
  extras?: GraderExtras;
}

// ── Workspace Delta types (#566) ──────────────────────────────────

export interface NewFile {
  path: string;
  size: number;
  hash: string;
}

export interface ModifiedFile {
  path: string;
  size_before: number;
  size_after: number;
  hash_after: string;
}

export interface DeletedFile {
  path: string;
  original_size: number;
}

export interface WorkspaceDelta {
  bytes_added: number;
  bytes_removed: number;
  bytes_net: number;
  new_file_count: number;
  modified_file_count: number;
  deleted_file_count: number;
  new_files: NewFile[];
  modified_files: ModifiedFile[];
  deleted_files: DeletedFile[];
}

export interface SessionEvent {
  type: string;
  tool_name?: string;
  tool_args?: string;
  tool_result?: string;
  tool_success?: boolean;
  content?: string;
  duration_ms?: number;
  turnNumber?: number;
  input_tokens?: number;
  output_tokens?: number;
  inputTokens?: number;
  outputTokens?: number;
  file_path?: string;
  file_operation?: string;
  file_size?: number;
  mcp_server_name?: string;
  mcp_tool_name?: string;
  error?: string;
}

export interface PromptMetadata {
  service: string;
  plane: string;
  language: string;
  category: string;
  difficulty: string;
  tags?: string[];
  sdk_package?: string;
}

/**
 * One row of `eval.tool_availability`. Records whether a specific
 * tool/skill/MCP server was available to the agent and whether it was
 * actually used. Mirrors hyoka/internal/report.ToolAvailabilityEntry.
 *
 * `parent` / `parent_kind` / `kind` are forward-compat optionals: today the
 * Go side only carries linkage on `environment.skill_groups`, but the site
 * accepts them on this shape so when the Go layer extends the row we won't
 * have to re-ship types. Until then the eval-detail "Available Tools"
 * panel joins these rows against `environment.skill_groups` by name to
 * pick up parent/plugin context.
 */
export interface ToolAvailabilityEntry {
  name: string;
  type: string; // "tool" | "skill" | "mcp"
  available: boolean;
  used: boolean;
  parent?: string;
  parent_kind?: string; // "plugin" | "skill_dir"
  kind?: string;        // "skill" | "mcp" | "plugin" | "skill_dir"
}

export interface SkillGroupEntry {
  /** Skill or MCP server name as the SDK reported it loaded. */
  name: string;
  /** Plugin or skill_dir this entry was contributed by; empty for top-level entries. */
  parent?: string;
  /** "skill" | "mcp" — what kind of leaf this is. */
  kind?: string;
  /** "plugin" | "skill_dir" — kind of the parent container. */
  parent_kind?: string;
}

export interface Environment {
  model: string;
  /**
   * Flat list of skill names the SDK loaded for the session.
   * Kept as `string[]` for backward compatibility with v2 reports and
   * existing site components. v3 reports also populate `skill_groups`
   * (sibling field) with the same set plus parent linkage — adopt that
   * when rendering grouped views.
   */
  skills_loaded?: string[];
  /** Schema v2 alias (Go emits camelCase). Same data as `skills_loaded`. */
  skillsLoaded?: string[];
  /**
   * Schema v3: structured view of `skills_loaded` with parent linkage
   * (plugin / skill_dir). Empty for v2 reports and any v3 run where the
   * tool validator did not emit topology (stub mode etc.).
   */
  skill_groups?: SkillGroupEntry[];
  skills_invoked?: string[];
  /** Schema v2 alias (Go emits camelCase). Same data as `skills_invoked`. */
  skillsInvoked?: string[];
  available_tools?: string[];
  mcp_servers?: string[];
  /** Schema v2 alias (Go emits camelCase). Same data as `mcp_servers`. */
  mcpServers?: string[];
  totalInputTokens?: number;
  totalOutputTokens?: number;
  total_input_tokens?: number;
  total_output_tokens?: number;
  turnCount?: number;
  turn_count?: number;
}

export interface EvalResult {
  prompt_id: string;
  config_name: string;
  success: boolean;
  error?: string;
  review: Review;
  duration_seconds: number;
  generated_files?: string[];
  prompt_metadata: PromptMetadata;
}

// ── Pairwise impact types (matching Go pairwise package) ──────────────

export interface VariantResult {
  config_name: string;
  removed_tool?: string;
  score: number;
  max_score: number;
  success: boolean;
}

export interface ToolImpact {
  tool_name: string;
  impact: number;
  baseline_score: number;
  without_score: number;
  baseline_pass: boolean;
  without_pass: boolean;
}

export interface PairwiseReport {
  prompt_id: string;
  baseline: VariantResult;
  variants: VariantResult[];
  impacts: ToolImpact[];
}

// Mirrors hyoka/internal/report.PairwiseRunReport — the payload returned by
// GET /api/runs/{runId}/pairwise. `aggregate_impacts` is optional because
// runs without any togglable tools produce no aggregate entries.
export interface PairwiseRunReport {
  run_id: string;
  timestamp: string;
  reports: PairwiseReport[];
  aggregate_impacts?: ToolImpact[];
}

// ── Generator Artifact (schema v3 addition for site display) ──────────

export interface ArtifactFileInfo {
  path: string;
  size: number;
}

export interface ArtifactWorkspaceDelta {
  bytes_added: number;
  bytes_removed: number;
  bytes_net: number;
  new_file_count: number;
  modified_file_count: number;
  deleted_file_count: number;
  created_files: ArtifactFileInfo[];
  modified_files: ArtifactFileInfo[];
  deleted_files: ArtifactFileInfo[];
}

export interface ActionsSummary {
  total_actions: number;
  tool_calls: number;
  reasoning_steps: number;
  truncated: boolean;
}

export interface GeneratorArtifact {
  prompt_id: string;
  eval_id?: string;
  config_name: string;
  generator_model: string;
  original_prompt: string;
  final_response: string;
  workspace_delta: ArtifactWorkspaceDelta;
  actions_summary: ActionsSummary;
  started_at: string;
  ended_at: string;
  duration_ms: number;
  terminated_by: string;
  error?: string;
}

// ── Run summary and eval result types ──────────────────────────────

export interface RunSummary {
  run_id: string;
  timestamp: string;
  total_prompts?: number;
  total_configs?: number;
  total_evaluations: number;
  passed: number;
  failed: number;
  errors: number;
  duration_seconds: number;
  avg_generation_duration_seconds?: number;
  avg_review_duration_seconds?: number;
  avg_build_duration_seconds?: number;
  analysis?: string;
  results: EvalResult[];
  pairwise_results?: PairwiseReport[];
}

export interface EvalReport {
  prompt_id: string;
  config_name: string;
  timestamp: string;
  success: boolean;
  error?: string;
  duration_seconds: number;
  generation_duration_seconds?: number;
  review_duration_seconds?: number;
  generated_files?: string[];
  reviewed_files?: ReviewedFile[];
  /**
   * `file_contents` maps generated file paths to their full text contents.
   * Populated by the engine for every eval that produced files (#bug-2).
   * Used as the fallback content source on the eval-detail page when the
   * reviewer did not annotate a file (`reviewed_files` lookup misses).
   */
  file_contents?: Record<string, string>;
  session_events?: SessionEvent[];
  event_count?: number;
  tool_calls?: string[];
  grader_results?: GraderResult[];
  review?: Review; // Legacy, kept for backward compat
  review_panel?: ReviewPanelEntry[];
  workspace_delta?: WorkspaceDelta;
  prompt_metadata: PromptMetadata;
  environment?: Environment;
  config_used?: { model: string; name: string };
  rerunCommand?: string;
  guardrail_abort_reason?: string;
  /**
   * `generator_artifact` captures the complete generator session state.
   * Populated at report-build time from generator.json when it exists.
   * Schema v3 addition for site display of session details, final response,
   * workspace delta, timing, and termination reason.
   */
  generator_artifact?: GeneratorArtifact;
  /**
   * `tool_availability` rows from the Go report (#348). Populated in v2
   * and v3. v3 may eventually include parent linkage on each entry.
   */
  tool_availability?: ToolAvailabilityEntry[];
  /**
   * Schema v3 roll-up totals computed engine-side from the unified
   * grader aggregate. When present, the site reads these directly so it
   * never has to recompute pass/total from grader_results — eliminating
   * the entire class of roll-up-divergence bugs by construction. Zero
   * value (no graders ran) is omitted, so v2 reports unmarshal safely.
   */
  graders_passed?: number;
  graders_total?: number;
  /**
   * Report schema marker. v3 = single ai_review entry + Points + roll-up
   * totals. v2 reports omit this field entirely. Number on the wire
   * (matches the Go side which emits an int).
   */
  schema_version?: number;
}

export interface DocEntry {
  slug: string;
  title: string;
}

export interface DocDetail {
  slug: string;
  title: string;
  content: string;
}

// ── Comparison types ─────────────────────────────────────────────────

export interface GraderDiff {
  name: string;
  score_a: number;
  score_b: number;
  delta: number;
  pass_a: boolean;
  pass_b: boolean;
}

export interface PromptDiff {
  prompt_id: string;
  score_a: number;
  score_b: number;
  delta: number;
  grader_diffs?: GraderDiff[];
  only_in_a?: boolean;
  only_in_b?: boolean;
}

export interface ComparisonSummary {
  avg_delta: number;
  improved: number;
  regressed: number;
  unchanged: number;
  top_improved?: PromptDiff[];
  top_regressed?: PromptDiff[];
}

export type ComparisonKind = "configs" | "runs" | "temporal";

export interface ComparisonResult {
  kind: ComparisonKind;
  label_a: string;
  label_b: string;
  config?: string;
  since?: string;
  per_prompt: PromptDiff[];
  summary: ComparisonSummary;
}
