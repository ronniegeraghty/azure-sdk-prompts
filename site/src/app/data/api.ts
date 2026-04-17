import type { RunSummary, EvalReport, DocEntry, DocDetail, PairwiseRunReport } from "./types";
import type { ComparisonResult } from "./types";

const API_BASE = "";

async function fetchJSON<T>(url: string): Promise<T> {
  const res = await fetch(`${API_BASE}${url}`);
  if (!res.ok) {
    throw new Error(`API error ${res.status}: ${res.statusText}`);
  }
  return res.json();
}

export interface PromptInfo {
  id: string;
  service: string;
  plane: string;
  language: string;
  category: string;
  difficulty: string;
  description: string;
  sdk_package: string;
  doc_url: string;
  tags: string[];
  created: string;
  author: string;
  prompt_text: string;
  evaluation_criteria: string;
  file_path: string;
}

export async function fetchRuns(): Promise<RunSummary[]> {
  return fetchJSON<RunSummary[]>("/api/runs");
}

export async function fetchRun(runId: string): Promise<RunSummary> {
  return fetchJSON<RunSummary>(`/api/runs/${encodeURIComponent(runId)}`);
}

export async function fetchEval(runId: string, evalPath: string): Promise<EvalReport> {
  return fetchJSON<EvalReport>(
    `/api/runs/${encodeURIComponent(runId)}/eval?path=${encodeURIComponent(evalPath)}`
  );
}

export async function fetchDocs(): Promise<DocEntry[]> {
  return fetchJSON<DocEntry[]>("/api/docs");
}

export async function fetchDoc(slug: string): Promise<DocDetail> {
  return fetchJSON<DocDetail>(`/api/docs/${encodeURIComponent(slug)}`);
}

export async function fetchPrompts(): Promise<PromptInfo[]> {
  return fetchJSON<PromptInfo[]>("/api/prompts");
}

export async function fetchPrompt(promptId: string): Promise<PromptInfo> {
  return fetchJSON<PromptInfo>(`/api/prompts/${encodeURIComponent(promptId)}`);
}

export async function fetchCompareConfigs(configA: string, configB: string): Promise<ComparisonResult> {
  return fetchJSON<ComparisonResult>(
    `/api/compare/configs?a=${encodeURIComponent(configA)}&b=${encodeURIComponent(configB)}`
  );
}

/**
 * Fetch pairwise tool-ablation results for a run. Returns `null` if the run
 * has no pairwise data (404 from the API), so callers can render an empty
 * state without a try/catch.
 */
export async function fetchPairwise(runId: string): Promise<PairwiseRunReport | null> {
  const res = await fetch(`/api/runs/${encodeURIComponent(runId)}/pairwise`);
  if (res.status === 404) return null;
  if (!res.ok) {
    throw new Error(`API error ${res.status}: ${res.statusText}`);
  }
  return res.json();
}
