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

// ── Grader system types (Phase 3) ────────────────────────────────

export interface FileCheckDetail {
  path: string;
  exists: boolean;
  pattern_matched?: boolean | null;
  pattern?: string;
}

export interface FileGraderDetail {
  checked_files: FileCheckDetail[];
}

export interface ProgramGraderDetail {
  command: string;
  exit_code: number;
  stdout: string;
  stderr: string;
}

export interface PromptGraderDetail {
  model: string;
  rubric: string;
  reasoning: string;
  raw_score?: number;
  max_score?: number;
}

export interface BehaviorGraderDetail {
  tools_used?: string[];
  missing_tools?: string[];
  forbidden_used?: string[];
  turn_count?: number;
  max_turns?: number;
  actual_turns?: number;
  total_actions?: number;
  turn_limit_hit?: boolean;
  violations?: string[];
  sequence_match?: boolean;
  expected_sequence?: string[];
  actual_sequence?: string[];
  matched_actions?: number;
  constraints_met?: boolean;
  tool_counts?: Record<string, number>;
}

export interface ReviewGraderDetail {
  model?: string;
  overall_score: number;
  max_score: number;
  summary: string;
  issues?: string[];
  strengths?: string[];
  is_consensus?: boolean;
  criteria?: ReviewCriteria[];
  panel_results?: ReviewPanelEntry[];
}

// GraderPoint mirrors hyoka/internal/graders.GraderPoint and the report-side
// `points` array introduced in schema v3 (Phase 2). Each grader emits one or
// more sub-checks; the site renders these directly instead of re-deriving
// pass/fail from the legacy expanded review entries. Empty for v2 reports.
export interface GraderPoint {
  name: string;
  pass: boolean;
  message?: string;
}

export interface GraderResult {
  grader_name: string;
  grader_type: string;
  model?: string;
  scores?: ReviewScores;
  overall_score?: number;
  max_score?: number;
  summary?: string;
  issues?: string[];
  strengths?: string[];
  duration_seconds?: number;
  is_consensus?: boolean;
  score?: number;
  weight?: number;
  pass?: boolean | null;
  gate?: boolean;
  file_details?: FileGraderDetail;
  program_details?: ProgramGraderDetail;
  prompt_details?: PromptGraderDetail;
  behavior_details?: BehaviorGraderDetail;
  review_details?: ReviewGraderDetail;
  /**
   * Per-sub-check rows from the grader. Populated in v3 reports. The site
   * AND-rolls these to derive overall pass/fail (see lib/evalPass.ts) so
   * eval-detail and run-detail never disagree on a grader's verdict.
   */
  points?: GraderPoint[];
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
