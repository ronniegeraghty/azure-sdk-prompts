import { Link, useSearchParams } from "react-router";
import { useState, useEffect, useMemo } from "react";
import { fetchRuns } from "../data/api";
import type { RunSummary } from "../data/types";
import { CheckCircle2, XCircle, AlertTriangle, Clock, ChevronRight, Loader2, X } from "lucide-react";
import { motion } from "motion/react";
import { MultiSelectFilter } from "./ui/multi-select-filter";
import {
  EMPTY_FILTERS,
  STATUS_LABEL,
  activeFilterCount,
  applyFilters,
  applyFiltersToSearchParams,
  buildCatalog,
  filtersFromSearchParams,
  hasActiveFilters,
  type RunFilters,
  type RunStatus,
} from "../lib/run-filters";

function formatDuration(s: number | undefined | null): string {
  if (s == null || isNaN(s)) return "N/A";
  if (s < 60) return `${s.toFixed(1)}s`;
  const m = Math.floor(s / 60);
  const sec = (s % 60).toFixed(0);
  return `${m}m ${sec}s`;
}

function formatTimestamp(ts: string | undefined | null): string {
  if (!ts) return "Unknown";
  const d = new Date(ts);
  if (isNaN(d.getTime())) return "Unknown";
  return d.toLocaleDateString("en-US", { 
    month: "short", 
    day: "numeric", 
    year: "numeric", 
    hour: "2-digit", 
    minute: "2-digit",
    hour12: true
  });
}

export function RunsPage() {
  const [runs, setRuns] = useState<RunSummary[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [searchParams, setSearchParams] = useSearchParams();

  useEffect(() => {
    fetchRuns()
      .then(setRuns)
      .catch(e => setError(e.message))
      .finally(() => setLoading(false));
  }, []);

  // Filters live in the URL so reload + share preserves state.
  const filters: RunFilters = useMemo(
    () => filtersFromSearchParams(searchParams),
    [searchParams],
  );
  const catalog = useMemo(() => buildCatalog(runs), [runs]);
  const filteredRuns = useMemo(() => applyFilters(runs, filters), [runs, filters]);

  function updateFilters(next: RunFilters) {
    const params = new URLSearchParams(searchParams);
    applyFiltersToSearchParams(params, next);
    setSearchParams(params, { replace: true });
  }

  function resetFilters() {
    updateFilters(EMPTY_FILTERS);
  }

  if (loading) {
    return (
      <div className="flex min-h-screen items-center justify-center bg-[#0a0a0f]">
        <Loader2 className="h-6 w-6 animate-spin text-emerald-400" />
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex min-h-screen items-center justify-center bg-[#0a0a0f]">
        <div className="text-center">
          <p className="mb-2 text-red-400">Failed to load runs</p>
          <p className="text-white/40" style={{ fontSize: 13 }}>{error}</p>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-[#0a0a0f] px-4 py-8 sm:px-6" style={{ fontFamily: "'Inter', sans-serif" }}>
      <div className="mx-auto max-w-5xl">
        <div className="mb-8">
          <h1 className="mb-2 text-white" style={{ fontFamily: "'JetBrains Mono', monospace", fontSize: "clamp(1.5rem, 3vw, 2rem)" }}>
            Evaluation Runs
          </h1>
          <p className="text-white/40" style={{ fontSize: 14 }}>
            Browse all evaluation runs and drill into individual results.
          </p>
        </div>

        {runs.length === 0 ? (
          <div className="rounded-xl border border-white/8 bg-white/[0.03] p-8 text-center">
            <p className="text-white/40">No runs found.</p>
          </div>
        ) : (
          <>
            <FilterBar
              catalog={catalog}
              filters={filters}
              onChange={updateFilters}
              onReset={resetFilters}
              filteredCount={filteredRuns.length}
              totalCount={runs.length}
            />
            {filteredRuns.length === 0 ? (
              <div
                role="status"
                className="rounded-xl border border-white/8 bg-white/[0.03] p-8 text-center"
              >
                <p className="mb-2 text-white/60">No runs match the current filters.</p>
                <button
                  onClick={resetFilters}
                  className="text-emerald-400 hover:text-emerald-300"
                  style={{ fontSize: 13 }}
                >
                  Reset filters
                </button>
              </div>
            ) : (
              <div className="space-y-4">
                {filteredRuns.map((run, i) => {
              const passed = run.passed ?? 0;
              const total = run.total_evaluations ?? 0;
              const errors = run.errors ?? 0;
              // Errored evaluations never produced a result — exclude from
              // the denominator so an all-errored run doesn't render as a
              // smug emerald `0.0%`. Falls back to total when no errors.
              const effectiveTotal = Math.max(total - errors, 0);
              const rate = effectiveTotal > 0
                ? ((passed / effectiveTotal) * 100).toFixed(1)
                : "0.0";
              const hasErrors = errors > 0;
              const barColor = hasErrors ? "bg-amber-500" : "bg-emerald-500";
              const rateColor = hasErrors ? "text-amber-300" : "text-white/50";
              return (
                <motion.div
                  key={run.run_id}
                  initial={{ opacity: 0, y: 12 }}
                  animate={{ opacity: 1, y: 0 }}
                  transition={{ delay: i * 0.05 }}
                >
                  <Link
                    to={`/runs/${run.run_id}`}
                    className="group block rounded-xl border border-white/8 bg-white/[0.03] p-5 no-underline transition hover:border-emerald-500/20 hover:bg-white/[0.05]"
                  >
                    <div className="flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
                      <div className="flex-1">
                        <div className="mb-1 flex items-center gap-2 text-white/80" style={{ fontSize: 15 }}>
                          {formatTimestamp(run.timestamp)}
                          {hasErrors && (
                            <span
                              className="inline-flex items-center gap-1 rounded-md border border-amber-500/30 bg-amber-500/[0.08] px-1.5 py-0.5 text-amber-300"
                              style={{ fontFamily: "'JetBrains Mono', monospace", fontSize: 10 }}
                            >
                              <AlertTriangle className="h-2.5 w-2.5" />
                              run errored
                            </span>
                          )}
                        </div>
                        <p className="text-white/40" style={{ fontSize: 12 }}>
                          {run.total_evaluations} evaluations
                        </p>
                      </div>

                      <div className="flex items-center gap-5">
                        <div className="flex items-center gap-4">
                          <div className="flex items-center gap-1.5">
                            <CheckCircle2 className="h-3.5 w-3.5 text-emerald-400" />
                            <span className="text-emerald-400" style={{ fontFamily: "'JetBrains Mono', monospace", fontSize: 13 }}>{run.passed}</span>
                          </div>
                          <div className="flex items-center gap-1.5">
                            <XCircle className="h-3.5 w-3.5 text-red-400" />
                            <span className="text-red-400" style={{ fontFamily: "'JetBrains Mono', monospace", fontSize: 13 }}>{run.failed}</span>
                          </div>
                          {run.errors > 0 && (
                            <div className="flex items-center gap-1.5">
                              <AlertTriangle className="h-3.5 w-3.5 text-amber-400" />
                              <span className="text-amber-400" style={{ fontFamily: "'JetBrains Mono', monospace", fontSize: 13 }}>{run.errors}</span>
                            </div>
                          )}
                        </div>

                        <div className="hidden items-center gap-2 sm:flex">
                          <div className="h-2 w-24 overflow-hidden rounded-full bg-white/10">
                            <div className={`h-full rounded-full ${barColor}`} style={{ width: `${rate}%` }} />
                          </div>
                          <span className={rateColor} style={{ fontFamily: "'JetBrains Mono', monospace", fontSize: 12 }}>
                            {rate}%
                          </span>
                        </div>

                        <div className="hidden items-center gap-1.5 text-white/30 sm:flex">
                          <Clock className="h-3.5 w-3.5" />
                          <span style={{ fontSize: 12 }}>{formatDuration(run.duration_seconds)}</span>
                        </div>

                        <ChevronRight className="h-4 w-4 text-white/20 transition group-hover:text-emerald-400" />
                      </div>
                    </div>
                  </Link>
                </motion.div>
              );
            })}
              </div>
            )}
          </>
        )}
      </div>
    </div>
  );
}

interface FilterBarProps {
  catalog: ReturnType<typeof buildCatalog>;
  filters: RunFilters;
  onChange: (next: RunFilters) => void;
  onReset: () => void;
  filteredCount: number;
  totalCount: number;
}

function FilterBar({
  catalog,
  filters,
  onChange,
  onReset,
  filteredCount,
  totalCount,
}: FilterBarProps) {
  const active = hasActiveFilters(filters);
  const count = activeFilterCount(filters);
  return (
    <div className="mb-5 rounded-xl border border-white/8 bg-white/[0.02] p-3">
      <div className="flex flex-wrap items-center gap-2">
        <MultiSelectFilter
          label="Config"
          options={catalog.configs.map((c) => ({ value: c, label: c }))}
          selected={filters.configs}
          onChange={(configs) => onChange({ ...filters, configs })}
        />
        <MultiSelectFilter
          label="Language"
          options={catalog.languages.map((l) => ({ value: l, label: l }))}
          selected={filters.languages}
          onChange={(languages) => onChange({ ...filters, languages })}
        />
        <MultiSelectFilter
          label="Status"
          options={catalog.statuses.map((s) => ({ value: s, label: STATUS_LABEL[s] }))}
          selected={filters.statuses}
          onChange={(values) =>
            onChange({ ...filters, statuses: values as RunStatus[] })
          }
        />

        <div className="ml-auto flex items-center gap-3">
          <span
            className="text-white/40"
            style={{ fontFamily: "'JetBrains Mono', monospace", fontSize: 12 }}
          >
            {active ? `${filteredCount} of ${totalCount}` : `${totalCount} runs`}
          </span>
          {active && (
            <button
              type="button"
              onClick={onReset}
              aria-label="Reset filters"
              className="flex items-center gap-1 rounded-lg border border-white/10 bg-white/5 px-2.5 py-1.5 text-white/70 transition hover:border-white/20 hover:text-white"
              style={{ fontSize: 12 }}
            >
              <X className="h-3 w-3" />
              Reset{count > 0 ? ` (${count})` : ""}
            </button>
          )}
        </div>
      </div>
    </div>
  );
}
