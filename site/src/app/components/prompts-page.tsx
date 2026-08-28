import { useState, useMemo, useEffect } from "react";
import { Link } from "react-router";
import { fetchPrompts, type PromptInfo } from "../data/api";
import type { RunSummary } from "../data/types";
import { Search, ChevronRight, Loader2, TrendingUp, TrendingDown, Clock } from "lucide-react";
import { motion } from "motion/react";
import { pointsPassRate, evalPassFromPoints } from "../lib/evalPass";
import { useRuns } from "../hooks/useRuns";

const mono = { fontFamily: "'JetBrains Mono', monospace" };

type SortBy = "recent" | "alphabetical" | "best" | "worst";

interface PromptWithStats {
  prompt_id: string;
  metadata: {
    service: string;
    language: string;
    difficulty: string;
    plane: string;
    tags: string[];
    category: string;
    sdk_package: string;
  };
  evalCount: number;
  passRate: number;
  lastEvaluated?: string;
  recentScores: number[];
}

export function PromptsPage() {
  const { runs, loading: runsLoading, error: runsError } = useRuns();
  const [allPrompts, setAllPrompts] = useState<PromptWithStats[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [search, setSearch] = useState("");
  const [filterService, setFilterService] = useState("all");
  const [filterLang, setFilterLang] = useState("all");
  const [filterDifficulty, setFilterDifficulty] = useState("all");
  const [filterPlane, setFilterPlane] = useState("all");
  const [onlyWithEvals, setOnlyWithEvals] = useState(true);
  const [sortBy, setSortBy] = useState<SortBy>("recent");

  useEffect(() => {
    // If runs are still loading or errored, wait
    if (runsLoading) {
      setLoading(true);
      return;
    }
    if (runsError) {
      setError(runsError);
      setLoading(false);
      return;
    }

    let cancelled = false;
    async function load() {
      try {
        const prompts = await fetchPrompts();
        if (cancelled) return;

        // Compute eval stats per prompt from runs
        const promptStats = new Map<string, { 
          evals: number; 
          passed: number; 
          lastEvaluated: string;
          recentScores: number[];
        }>();
        
        // Sort runs by timestamp to get recent scores in order
        const sortedRuns = [...runs].sort((a, b) => 
          new Date(b.timestamp).getTime() - new Date(a.timestamp).getTime()
        );

        for (const run of sortedRuns) {
          for (const result of run.results || []) {
            const key = result.prompt_id;
            const stats = promptStats.get(key) || { 
              evals: 0, 
              passed: 0,
              lastEvaluated: run.timestamp,
              recentScores: []
            };
            stats.evals++;
            if (evalPassFromPoints(result)) stats.passed++;
            
            // Track recent scores (last 10) for sparkline. Use fractional
            // grader-point pass rate, not legacy review.overall_score.
            if (stats.recentScores.length < 10) {
              stats.recentScores.push(pointsPassRate(result));
            }
            
            // Update last evaluated if this run is more recent
            if (new Date(run.timestamp) > new Date(stats.lastEvaluated)) {
              stats.lastEvaluated = run.timestamp;
            }
            
            promptStats.set(key, stats);
          }
        }

        const merged: PromptWithStats[] = prompts.map((p: PromptInfo) => {
          const stats = promptStats.get(p.id);
          return {
            prompt_id: p.id,
            metadata: {
              service: p.service,
              language: p.language,
              difficulty: p.difficulty,
              plane: p.plane,
              tags: p.tags || [],
              category: p.category,
              sdk_package: p.sdk_package,
            },
            evalCount: stats?.evals || 0,
            passRate: stats?.evals
              ? Math.round((stats.passed / stats.evals) * 100)
              : 0,
            lastEvaluated: stats?.lastEvaluated,
            recentScores: stats?.recentScores || [],
          };
        });

        setAllPrompts(merged);
      } catch (e) {
        if (!cancelled) setError(e instanceof Error ? e.message : "Failed to load prompts");
      } finally {
        if (!cancelled) setLoading(false);
      }
    }
    load();
    return () => { cancelled = true; };
  }, [runs, runsLoading, runsError]);

  const services = useMemo(() => [...new Set(allPrompts.map(p => p.metadata.service))].sort(), [allPrompts]);
  const langs = useMemo(() => [...new Set(allPrompts.map(p => p.metadata.language))].sort(), [allPrompts]);

  // Filter prompts
  const filtered = useMemo(() => {
    let result = allPrompts.filter(p => {
      if (onlyWithEvals && p.evalCount === 0) return false;
      if (search && !p.prompt_id.toLowerCase().includes(search.toLowerCase())) return false;
      if (filterService !== "all" && p.metadata.service !== filterService) return false;
      if (filterLang !== "all" && p.metadata.language !== filterLang) return false;
      if (filterDifficulty !== "all" && p.metadata.difficulty !== filterDifficulty) return false;
      if (filterPlane !== "all" && p.metadata.plane !== filterPlane) return false;
      return true;
    });

    // Apply sorting
    switch (sortBy) {
      case "recent":
        result.sort((a, b) => {
          if (!a.lastEvaluated) return 1;
          if (!b.lastEvaluated) return -1;
          return new Date(b.lastEvaluated).getTime() - new Date(a.lastEvaluated).getTime();
        });
        break;
      case "alphabetical":
        result.sort((a, b) => a.prompt_id.localeCompare(b.prompt_id));
        break;
      case "best":
        result.sort((a, b) => b.passRate - a.passRate);
        break;
      case "worst":
        result.sort((a, b) => a.passRate - b.passRate);
        break;
    }

    return result;
  }, [allPrompts, search, filterService, filterLang, filterDifficulty, filterPlane, onlyWithEvals, sortBy]);

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
        <div className="rounded-xl border border-red-500/20 bg-red-500/5 px-6 py-4 text-red-400" style={{ fontSize: 14 }}>
          {error}
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-[#0a0a0f] px-4 py-8 sm:px-6" style={{ fontFamily: "'Inter', sans-serif" }}>
      <div className="mx-auto max-w-6xl">
        <div className="mb-8">
          <h1 className="mb-2 text-white" style={{ ...mono, fontSize: "clamp(1.5rem, 3vw, 2rem)" }}>
            Prompt Explorer
          </h1>
          <p className="text-white/40" style={{ fontSize: 14 }}>
            Browse and filter all evaluation prompts. Click any prompt to see its history across runs.
          </p>
        </div>

        {/* Search & Filters */}
        <div className="mb-6 space-y-4">
          {/* Search bar */}
          <div className="relative">
            <Search className="absolute left-3 top-1/2 h-4 w-4 -translate-y-1/2 text-white/30" />
            <input
              type="text"
              value={search}
              onChange={e => setSearch(e.target.value)}
              placeholder="Search prompts..."
              className="w-full rounded-lg border border-white/10 bg-white/5 py-2 pl-10 pr-4 text-white placeholder-white/30 outline-none focus:border-emerald-500/30"
              style={{ fontSize: 13 }}
            />
          </div>

          {/* Filter controls */}
          <div className="flex flex-wrap items-center gap-3">
            <select value={filterService} onChange={e => setFilterService(e.target.value)}
              className="rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white/70 outline-none focus:border-emerald-500/30" style={{ fontSize: 12 }}>
              <option value="all">All Services</option>
              {services.map(s => <option key={s} value={s}>{s}</option>)}
            </select>
            <select value={filterLang} onChange={e => setFilterLang(e.target.value)}
              className="rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white/70 outline-none focus:border-emerald-500/30" style={{ fontSize: 12 }}>
              <option value="all">All Languages</option>
              {langs.map(l => <option key={l} value={l}>{l}</option>)}
            </select>
            <select value={filterDifficulty} onChange={e => setFilterDifficulty(e.target.value)}
              className="rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white/70 outline-none focus:border-emerald-500/30" style={{ fontSize: 12 }}>
              <option value="all">All Difficulty</option>
              <option value="basic">Basic</option>
              <option value="intermediate">Intermediate</option>
              <option value="advanced">Advanced</option>
            </select>
            <select value={filterPlane} onChange={e => setFilterPlane(e.target.value)}
              className="rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white/70 outline-none focus:border-emerald-500/30" style={{ fontSize: 12 }}>
              <option value="all">All Planes</option>
              <option value="data-plane">Data Plane</option>
              <option value="management-plane">Management Plane</option>
            </select>
            
            <div className="ml-auto flex items-center gap-3">
              {/* Only show with evals toggle */}
              <label className="flex cursor-pointer items-center gap-2 rounded-lg border border-white/10 bg-white/5 px-3 py-2 transition hover:bg-white/10">
                <input
                  type="checkbox"
                  checked={onlyWithEvals}
                  onChange={e => setOnlyWithEvals(e.target.checked)}
                  className="h-3.5 w-3.5 accent-emerald-500"
                />
                <span className="text-white/70" style={{ fontSize: 12 }}>Only show with evals</span>
              </label>

              {/* Sort dropdown */}
              <select value={sortBy} onChange={e => setSortBy(e.target.value as SortBy)}
                className="rounded-lg border border-white/10 bg-white/5 px-3 py-2 text-white/70 outline-none focus:border-emerald-500/30" style={{ fontSize: 12 }}>
                <option value="recent">Most Recently Evaluated</option>
                <option value="alphabetical">Alphabetically</option>
                <option value="best">Best Performing</option>
                <option value="worst">Worst Performing</option>
              </select>
            </div>
          </div>
        </div>

        <p className="mb-4 text-white/30" style={{ fontSize: 12 }}>
          {filtered.length} prompt{filtered.length === 1 ? '' : 's'} found
          {onlyWithEvals && ` (${allPrompts.filter(p => p.evalCount === 0).length} hidden without evals)`}
        </p>

        {/* Prompt grid */}
        <div className="grid gap-3 md:grid-cols-2 lg:grid-cols-3">
          {filtered.map((p, i) => {
            const rateColor = p.passRate >= 80 ? "text-emerald-400" : p.passRate >= 60 ? "text-amber-400" : "text-red-400";
            const diffColor = p.metadata.difficulty === "basic" ? "bg-emerald-500/10 text-emerald-400/70" :
              p.metadata.difficulty === "intermediate" ? "bg-amber-500/10 text-amber-400/70" : "bg-red-500/10 text-red-400/70";

            return (
              <motion.div
                key={p.prompt_id}
                initial={{ opacity: 0, y: 10 }}
                animate={{ opacity: 1, y: 0 }}
                transition={{ delay: Math.min(i * 0.02, 0.3) }}
              >
                <Link
                  to={`/prompts/${encodeURIComponent(p.prompt_id)}`}
                  className="group block rounded-xl border border-white/8 bg-white/[0.03] p-4 no-underline transition hover:border-emerald-500/20 hover:bg-white/[0.05]"
                >
                  <div className="mb-3 flex items-start justify-between gap-2">
                    <span className="text-emerald-400/80" style={{ ...mono, fontSize: 12 }}>
                      {p.prompt_id}
                    </span>
                    <ChevronRight className="h-3.5 w-3.5 shrink-0 text-white/15 transition group-hover:text-emerald-400" />
                  </div>

                  <div className="mb-3 flex flex-wrap gap-1.5">
                    <span className="rounded-md bg-white/5 px-2 py-0.5 text-white/50" style={{ fontSize: 10 }}>{p.metadata.service}</span>
                    <span className="rounded-md bg-white/5 px-2 py-0.5 text-white/50" style={{ fontSize: 10 }}>{p.metadata.language}</span>
                    <span className={`rounded-md px-2 py-0.5 ${diffColor}`} style={{ fontSize: 10 }}>{p.metadata.difficulty}</span>
                  </div>

                  {/* Tags - more prominent */}
                  {p.metadata.tags.length > 0 && (
                    <div className="mb-3 flex flex-wrap gap-1">
                      {p.metadata.tags.slice(0, 5).map(t => (
                        <span key={t} className="rounded bg-white/[0.06] px-2 py-0.5 text-white/40" style={{ fontSize: 10 }}>{t}</span>
                      ))}
                    </div>
                  )}

                  {/* Pass rate with sparkline */}
                  <div className="mb-2 flex items-center justify-between gap-3">
                    <div className="flex items-center gap-2">
                      <span className={rateColor} style={{ ...mono, fontSize: 13, fontWeight: 600 }}>{p.passRate}%</span>
                      {p.passRate >= 80 ? (
                        <TrendingUp className="h-3.5 w-3.5 text-emerald-400" />
                      ) : p.passRate < 60 ? (
                        <TrendingDown className="h-3.5 w-3.5 text-red-400" />
                      ) : null}
                    </div>
                    
                    {/* Mini sparkline if we have score data */}
                    {p.recentScores.length > 0 && (
                      <svg width="60" height="20" className="opacity-50">
                        <polyline
                          fill="none"
                          stroke={p.passRate >= 80 ? "#34d399" : p.passRate >= 60 ? "#fbbf24" : "#f87171"}
                          strokeWidth="1.5"
                          points={p.recentScores.map((score, idx) => {
                            const x = (idx / Math.max(p.recentScores.length - 1, 1)) * 60;
                            const y = 20 - (score / 100) * 18;
                            return `${x},${y}`;
                          }).join(' ')}
                        />
                      </svg>
                    )}
                  </div>

                  {/* Eval count - more prominent */}
                  <div className="flex items-center gap-1.5 text-white/40">
                    <Clock className="h-3 w-3" />
                    <span style={{ fontSize: 11 }}>{p.evalCount} evaluation{p.evalCount === 1 ? '' : 's'}</span>
                    {p.lastEvaluated && (
                      <span className="text-white/20" style={{ fontSize: 10 }}>
                        · {new Date(p.lastEvaluated).toLocaleDateString()}
                      </span>
                    )}
                  </div>
                </Link>
              </motion.div>
            );
          })}
        </div>
      </div>
    </div>
  );
}
