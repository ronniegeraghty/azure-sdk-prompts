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

export interface ReviewPanelEntry {
  model: string;
  overall_score: number;
  max_score: number;
  summary: string;
  scores?: ReviewScores;
  events?: unknown[];
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

export interface Environment {
  model: string;
  skills_loaded?: string[];
  skills_invoked?: string[];
  available_tools?: string[];
  mcp_servers?: string[];
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
  session_events?: SessionEvent[];
  event_count?: number;
  tool_calls?: string[];
  review: Review;
  review_panel?: ReviewPanelEntry[];
  prompt_metadata: PromptMetadata;
  environment?: Environment;
  config_used?: { model: string; name: string };
  rerunCommand?: string;
  guardrail_abort_reason?: string;
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

export interface ConfigComparison {
  config_a: string;
  config_b: string;
  per_prompt: PromptDiff[];
  summary: ComparisonSummary;
}
