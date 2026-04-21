// Pure model + helpers for the run-level filter system on the runs page.
//
// A RunFilters value is a multi-select per dimension (configs, languages,
// statuses). Each RunSummary is annotated with the union of its evals'
// configs/languages plus an aggregate status, and matches a filter when it
// has at least one matching eval per dimension.
//
// Selection semantics (matches group-based-comparison-ui skill):
//   - Empty array for a dimension = match-all for that dimension.
//   - Within a dimension, values OR together (any match counts).
//   - Across dimensions, results AND together (all dimensions must match).
//
// All functions in this module are pure and side-effect free so they can be
// unit-tested without rendering React.

import type { RunSummary } from "../data/types";

export type RunStatus = "passing" | "failing" | "errors";

export const ALL_STATUSES: RunStatus[] = ["passing", "failing", "errors"];

export const STATUS_LABEL: Record<RunStatus, string> = {
  passing: "All passing",
  failing: "Has failures",
  errors: "Has errors",
};

export interface RunFilters {
  configs: string[];
  languages: string[];
  statuses: RunStatus[];
}

export const EMPTY_FILTERS: RunFilters = {
  configs: [],
  languages: [],
  statuses: [],
};

export interface RunFilterCatalog {
  configs: string[];
  languages: string[];
  statuses: RunStatus[];
}

export interface RunFacets {
  configs: Set<string>;
  languages: Set<string>;
  status: RunStatus;
}

export function deriveStatus(run: RunSummary): RunStatus {
  const errors = run.errors ?? 0;
  if (errors > 0) return "errors";
  const failed = run.failed ?? 0;
  if (failed > 0) return "failing";
  return "passing";
}

export function deriveFacets(run: RunSummary): RunFacets {
  const configs = new Set<string>();
  const languages = new Set<string>();
  for (const r of run.results || []) {
    if (r.config_name) configs.add(r.config_name);
    const lang = r.prompt_metadata?.language;
    if (lang) languages.add(lang);
  }
  return { configs, languages, status: deriveStatus(run) };
}

export function buildCatalog(runs: RunSummary[]): RunFilterCatalog {
  const configs = new Set<string>();
  const languages = new Set<string>();
  const statuses = new Set<RunStatus>();
  for (const run of runs) {
    const facets = deriveFacets(run);
    for (const c of facets.configs) configs.add(c);
    for (const l of facets.languages) languages.add(l);
    statuses.add(facets.status);
  }
  return {
    configs: Array.from(configs).sort(),
    languages: Array.from(languages).sort(),
    statuses: ALL_STATUSES.filter((s) => statuses.has(s)),
  };
}

function dimensionPasses(selected: string[], present: Set<string>): boolean {
  if (selected.length === 0) return true;
  for (const v of selected) if (present.has(v)) return true;
  return false;
}

export function runMatches(run: RunSummary, filters: RunFilters): boolean {
  const facets = deriveFacets(run);
  if (!dimensionPasses(filters.configs, facets.configs)) return false;
  if (!dimensionPasses(filters.languages, facets.languages)) return false;
  if (filters.statuses.length > 0 && !filters.statuses.includes(facets.status)) {
    return false;
  }
  return true;
}

export function applyFilters(runs: RunSummary[], filters: RunFilters): RunSummary[] {
  if (!hasActiveFilters(filters)) return runs;
  return runs.filter((r) => runMatches(r, filters));
}

export function hasActiveFilters(filters: RunFilters): boolean {
  return (
    filters.configs.length > 0 ||
    filters.languages.length > 0 ||
    filters.statuses.length > 0
  );
}

export function activeFilterCount(filters: RunFilters): number {
  return filters.configs.length + filters.languages.length + filters.statuses.length;
}

export function toggleValue<T extends string>(values: T[], value: T): T[] {
  return values.includes(value) ? values.filter((v) => v !== value) : [...values, value];
}

// ── URL serialization ──────────────────────────────────────────────
// Filters round-trip through URLSearchParams as comma-separated lists, e.g.
// ?config=baseline%2Fclaude-opus-4.6,azure-mcp%2Fclaude-opus-4.6&lang=python&status=failing

const PARAM_KEYS = {
  configs: "config",
  languages: "lang",
  statuses: "status",
} as const;

function splitList(raw: string | null): string[] {
  if (!raw) return [];
  return raw.split(",").map((v) => v.trim()).filter(Boolean);
}

export function filtersFromSearchParams(params: URLSearchParams): RunFilters {
  const rawStatuses = splitList(params.get(PARAM_KEYS.statuses));
  const statuses = rawStatuses.filter((s): s is RunStatus =>
    ALL_STATUSES.includes(s as RunStatus),
  );
  return {
    configs: splitList(params.get(PARAM_KEYS.configs)),
    languages: splitList(params.get(PARAM_KEYS.languages)),
    statuses,
  };
}

// Mutates `params` in place: sets keys for non-empty dimensions, removes them
// otherwise. Returns the same object for convenience when building a URL.
export function applyFiltersToSearchParams(
  params: URLSearchParams,
  filters: RunFilters,
): URLSearchParams {
  const entries: [string, string[]][] = [
    [PARAM_KEYS.configs, filters.configs],
    [PARAM_KEYS.languages, filters.languages],
    [PARAM_KEYS.statuses, filters.statuses],
  ];
  for (const [key, values] of entries) {
    if (values.length === 0) {
      params.delete(key);
    } else {
      params.set(key, values.join(","));
    }
  }
  return params;
}
