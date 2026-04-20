// Group-based comparison model and pure helpers.
//
// A ComparisonGroup is a named, filtered subset of evaluation results.
// Users define multiple groups and the comparison page renders configurable
// visualizations comparing metrics across those groups.
//
// All functions here are pure — they take RunSummary[] (already loaded by
// the page) and produce derived data. No fetching, no React.

import type { RunSummary, EvalResult } from "../data/types";

export type FilterDimension =
  | "configs"
  | "services"
  | "languages"
  | "planes"
  | "categories"
  | "difficulties";

export const FILTER_DIMENSIONS: { key: FilterDimension; label: string }[] = [
  { key: "configs", label: "Config" },
  { key: "services", label: "Service" },
  { key: "languages", label: "Language" },
  { key: "planes", label: "Plane" },
  { key: "categories", label: "Category" },
  { key: "difficulties", label: "Difficulty" },
];

export type GroupFilters = Partial<Record<FilterDimension, string[]>>;

export interface ComparisonGroup {
  id: string;
  name: string;
  color: string;
  filters: GroupFilters;
}

export type ChartId =
  | "pass_rate"
  | "avg_score"
  | "by_service"
  | "by_language"
  | "score_distribution"
  | "eval_count";

export const CHART_OPTIONS: { id: ChartId; label: string; description: string }[] = [
  { id: "pass_rate", label: "Pass Rate", description: "Percent of evals where pass=true." },
  { id: "avg_score", label: "Average Score", description: "Mean review score (0–100%)." },
  { id: "by_service", label: "Pass Rate by Service", description: "Pass rate broken down per service." },
  { id: "by_language", label: "Pass Rate by Language", description: "Pass rate broken down per language." },
  { id: "score_distribution", label: "Score Distribution", description: "Histogram of eval scores in 10% bins." },
  { id: "eval_count", label: "Eval Count", description: "Number of matching evals per group." },
];

// Visually distinct accent palette. Cycled when more groups than colors.
export const GROUP_COLORS = [
  "#10b981", // emerald
  "#38bdf8", // sky
  "#a78bfa", // violet
  "#f59e0b", // amber
  "#f472b6", // pink
  "#22d3ee", // cyan
  "#fb7185", // rose
  "#84cc16", // lime
];

export function nextGroupColor(existing: ComparisonGroup[]): string {
  return GROUP_COLORS[existing.length % GROUP_COLORS.length];
}

// ── Catalog extraction ─────────────────────────────────────────────

export interface FilterCatalog {
  configs: string[];
  services: string[];
  languages: string[];
  planes: string[];
  categories: string[];
  difficulties: string[];
}

export function flattenResults(runs: RunSummary[]): EvalResult[] {
  const out: EvalResult[] = [];
  for (const run of runs) {
    if (!run.results) continue;
    for (const r of run.results) out.push(r);
  }
  return out;
}

export function buildCatalog(results: EvalResult[]): FilterCatalog {
  const configs = new Set<string>();
  const services = new Set<string>();
  const languages = new Set<string>();
  const planes = new Set<string>();
  const categories = new Set<string>();
  const difficulties = new Set<string>();

  for (const r of results) {
    if (r.config_name) configs.add(r.config_name);
    const m = r.prompt_metadata;
    if (m?.service) services.add(m.service);
    if (m?.language) languages.add(m.language);
    if (m?.plane) planes.add(m.plane);
    if (m?.category) categories.add(m.category);
    if (m?.difficulty) difficulties.add(m.difficulty);
  }
  const sort = (s: Set<string>) => Array.from(s).sort();
  return {
    configs: sort(configs),
    services: sort(services),
    languages: sort(languages),
    planes: sort(planes),
    categories: sort(categories),
    difficulties: sort(difficulties),
  };
}

// ── Group filtering ────────────────────────────────────────────────

function passesFilter(values: string[] | undefined, candidate: string | undefined): boolean {
  if (!values || values.length === 0) return true; // empty filter = match-all
  if (!candidate) return false;
  return values.includes(candidate);
}

export function evalMatchesGroup(r: EvalResult, group: ComparisonGroup): boolean {
  const f = group.filters;
  if (!passesFilter(f.configs, r.config_name)) return false;
  const m = r.prompt_metadata;
  if (!passesFilter(f.services, m?.service)) return false;
  if (!passesFilter(f.languages, m?.language)) return false;
  if (!passesFilter(f.planes, m?.plane)) return false;
  if (!passesFilter(f.categories, m?.category)) return false;
  if (!passesFilter(f.difficulties, m?.difficulty)) return false;
  return true;
}

export function filterGroup(results: EvalResult[], group: ComparisonGroup): EvalResult[] {
  return results.filter((r) => evalMatchesGroup(r, group));
}

// ── Metric computation ─────────────────────────────────────────────

function safeScore(r: EvalResult): number {
  const m = r.review?.max_score ?? 0;
  if (!m) return 0;
  return r.review.overall_score / m;
}

export interface GroupMetrics {
  count: number;
  pass_count: number;
  pass_rate: number; // 0..1
  avg_score: number; // 0..1
  avg_duration: number; // seconds
  by_service: Record<string, { total: number; passed: number }>;
  by_language: Record<string, { total: number; passed: number }>;
  score_distribution: number[]; // length 10, bins 0–10%, …, 90–100%
}

export function computeMetrics(results: EvalResult[]): GroupMetrics {
  const byService: Record<string, { total: number; passed: number }> = {};
  const byLanguage: Record<string, { total: number; passed: number }> = {};
  const dist = new Array(10).fill(0) as number[];

  let scoreSum = 0;
  let durationSum = 0;
  let passCount = 0;

  for (const r of results) {
    const score = safeScore(r);
    scoreSum += score;
    durationSum += r.duration_seconds || 0;
    if (r.success) passCount++;

    const svc = r.prompt_metadata?.service || "unknown";
    const lang = r.prompt_metadata?.language || "unknown";
    byService[svc] ??= { total: 0, passed: 0 };
    byService[svc].total++;
    if (r.success) byService[svc].passed++;
    byLanguage[lang] ??= { total: 0, passed: 0 };
    byLanguage[lang].total++;
    if (r.success) byLanguage[lang].passed++;

    // Histogram bin 0..9. Score 1.0 falls in last bin.
    let bin = Math.floor(score * 10);
    if (bin < 0) bin = 0;
    if (bin > 9) bin = 9;
    dist[bin]++;
  }

  const n = results.length;
  return {
    count: n,
    pass_count: passCount,
    pass_rate: n > 0 ? passCount / n : 0,
    avg_score: n > 0 ? scoreSum / n : 0,
    avg_duration: n > 0 ? durationSum / n : 0,
    by_service: byService,
    by_language: byLanguage,
    score_distribution: dist,
  };
}

// ── Persistence ────────────────────────────────────────────────────

export const STORAGE_KEY = "hyoka:comparison:v1";

export interface PersistedState {
  groups: ComparisonGroup[];
  charts: ChartId[];
}

export function loadState(storage: Storage = safeLocalStorage()): PersistedState | null {
  try {
    const raw = storage.getItem(STORAGE_KEY);
    if (!raw) return null;
    const parsed = JSON.parse(raw) as PersistedState;
    if (!parsed || !Array.isArray(parsed.groups) || !Array.isArray(parsed.charts)) return null;
    // Lightweight shape validation; drop anything malformed.
    const groups = parsed.groups.filter(
      (g) =>
        g &&
        typeof g.id === "string" &&
        typeof g.name === "string" &&
        typeof g.color === "string" &&
        g.filters &&
        typeof g.filters === "object"
    );
    return { groups, charts: parsed.charts.filter((c) => typeof c === "string") as ChartId[] };
  } catch {
    return null;
  }
}

export function saveState(state: PersistedState, storage: Storage = safeLocalStorage()): void {
  try {
    storage.setItem(STORAGE_KEY, JSON.stringify(state));
  } catch {
    // Quota / privacy mode — silently ignore. Comparison still works in-memory.
  }
}

function safeLocalStorage(): Storage {
  if (typeof window === "undefined" || !window.localStorage) {
    // Test/SSR fallback: an in-memory shim.
    const store = new Map<string, string>();
    return {
      getItem: (k) => store.get(k) ?? null,
      setItem: (k, v) => void store.set(k, v),
      removeItem: (k) => void store.delete(k),
      clear: () => store.clear(),
      key: (i) => Array.from(store.keys())[i] ?? null,
      get length() {
        return store.size;
      },
    } as Storage;
  }
  return window.localStorage;
}

// ── ID helpers ─────────────────────────────────────────────────────

export function newGroupId(): string {
  return `g-${Date.now().toString(36)}-${Math.random().toString(36).slice(2, 7)}`;
}

export function describeFilters(g: ComparisonGroup): string {
  const parts: string[] = [];
  for (const dim of FILTER_DIMENSIONS) {
    const vals = g.filters[dim.key];
    if (vals && vals.length > 0) {
      parts.push(`${dim.label}: ${vals.join(", ")}`);
    }
  }
  return parts.length > 0 ? parts.join(" · ") : "No filters (matches all evals)";
}
