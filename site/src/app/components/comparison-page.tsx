import { useState, useEffect, useMemo } from "react";
import { fetchRuns, fetchCompareConfigs } from "../data/api";
import type { RunSummary, ConfigComparison, PromptDiff, GraderDiff } from "../data/types";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "./ui/select";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "./ui/table";
import {
  ArrowUpRight,
  ArrowDownRight,
  Minus,
  Loader2,
  GitCompareArrows,
  TrendingUp,
  TrendingDown,
  Equal,
  ChevronDown,
  ChevronRight,
  Filter,
} from "lucide-react";
import { motion, AnimatePresence } from "motion/react";

function formatScore(score: number): string {
  return (score * 100).toFixed(1) + "%";
}

function formatDelta(delta: number): string {
  const pct = (delta * 100).toFixed(1);
  if (delta > 0) return `+${pct}%`;
  return `${pct}%`;
}

function deltaColor(delta: number): string {
  if (delta > 0.001) return "text-emerald-400";
  if (delta < -0.001) return "text-red-400";
  return "text-white/40";
}

function deltaBg(delta: number): string {
  if (delta > 0.001) return "bg-emerald-500/10";
  if (delta < -0.001) return "bg-red-500/10";
  return "bg-white/5";
}

function DeltaIcon({ delta }: { delta: number }) {
  if (delta > 0.001)
    return <ArrowUpRight className="h-3.5 w-3.5 text-emerald-400" />;
  if (delta < -0.001)
    return <ArrowDownRight className="h-3.5 w-3.5 text-red-400" />;
  return <Minus className="h-3.5 w-3.5 text-white/30" />;
}

function extractConfigNames(runs: RunSummary[]): string[] {
  const names = new Set<string>();
  for (const run of runs) {
    if (!run.results) continue;
    for (const r of run.results) {
      if (r.config_name) names.add(r.config_name);
    }
  }
  return Array.from(names).sort();
}

function extractFilterValues(
  diffs: PromptDiff[]
): { languages: string[]; services: string[] } {
  const languages = new Set<string>();
  const services = new Set<string>();
  for (const d of diffs) {
    const parts = d.prompt_id.split("-");
    // Pattern: {service}-{plane}-{language}-{rest}
    if (parts.length >= 3) {
      services.add(parts[0]);
      // Language is after the plane abbreviation (dp/mp)
      languages.add(parts[2]);
    }
  }
  return {
    languages: Array.from(languages).sort(),
    services: Array.from(services).sort(),
  };
}

function GraderDiffRow({ diff }: { diff: GraderDiff }) {
  return (
    <TableRow className="border-white/5 bg-white/[0.01]">
      <TableCell className="pl-10 text-white/50" style={{ fontSize: 12 }}>
        {diff.name}
      </TableCell>
      <TableCell
        className="text-white/60"
        style={{ fontFamily: "'JetBrains Mono', monospace", fontSize: 12 }}
      >
        {diff.score_a.toFixed(2)}
        <span className="ml-1.5 text-white/30">
          {diff.pass_a ? "✓" : "✗"}
        </span>
      </TableCell>
      <TableCell
        className="text-white/60"
        style={{ fontFamily: "'JetBrains Mono', monospace", fontSize: 12 }}
      >
        {diff.score_b.toFixed(2)}
        <span className="ml-1.5 text-white/30">
          {diff.pass_b ? "✓" : "✗"}
        </span>
      </TableCell>
      <TableCell>
        <span
          className={`inline-flex items-center gap-1 ${deltaColor(diff.delta)}`}
          style={{ fontFamily: "'JetBrains Mono', monospace", fontSize: 12 }}
        >
          <DeltaIcon delta={diff.delta} />
          {diff.delta > 0 ? "+" : ""}
          {diff.delta.toFixed(2)}
        </span>
      </TableCell>
    </TableRow>
  );
}

function PromptDiffRow({ diff }: { diff: PromptDiff }) {
  const [expanded, setExpanded] = useState(false);
  const hasGraders = diff.grader_diffs && diff.grader_diffs.length > 0;
  const isOneSided = diff.only_in_a || diff.only_in_b;

  return (
    <>
      <TableRow
        className={`cursor-pointer border-white/5 transition-colors hover:bg-white/[0.04] ${
          hasGraders ? "" : "cursor-default"
        }`}
        onClick={() => hasGraders && setExpanded(!expanded)}
      >
        <TableCell>
          <div className="flex items-center gap-2">
            {hasGraders ? (
              expanded ? (
                <ChevronDown className="h-3.5 w-3.5 text-white/30" />
              ) : (
                <ChevronRight className="h-3.5 w-3.5 text-white/30" />
              )
            ) : (
              <span className="w-3.5" />
            )}
            <span
              className="text-emerald-400"
              style={{
                fontFamily: "'JetBrains Mono', monospace",
                fontSize: 13,
              }}
            >
              {diff.prompt_id}
            </span>
            {isOneSided && (
              <span className="rounded bg-amber-500/10 px-1.5 py-0.5 text-amber-400" style={{ fontSize: 10 }}>
                {diff.only_in_a ? "Only in A" : "Only in B"}
              </span>
            )}
          </div>
        </TableCell>
        <TableCell
          className="text-white/60"
          style={{ fontFamily: "'JetBrains Mono', monospace", fontSize: 13 }}
        >
          {isOneSided && diff.only_in_b ? "—" : formatScore(diff.score_a)}
        </TableCell>
        <TableCell
          className="text-white/60"
          style={{ fontFamily: "'JetBrains Mono', monospace", fontSize: 13 }}
        >
          {isOneSided && diff.only_in_a ? "—" : formatScore(diff.score_b)}
        </TableCell>
        <TableCell>
          {!isOneSided ? (
            <span
              className={`inline-flex items-center gap-1 rounded-md px-2 py-0.5 ${deltaBg(diff.delta)} ${deltaColor(diff.delta)}`}
              style={{
                fontFamily: "'JetBrains Mono', monospace",
                fontSize: 13,
              }}
            >
              <DeltaIcon delta={diff.delta} />
              {formatDelta(diff.delta)}
            </span>
          ) : (
            <span className="text-white/20">—</span>
          )}
        </TableCell>
      </TableRow>
      <AnimatePresence>
        {expanded && hasGraders && (
          <motion.tr
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.15 }}
          >
            <td colSpan={4} className="p-0">
              <Table>
                <TableBody>
                  {diff.grader_diffs!.map((gd) => (
                    <GraderDiffRow key={gd.name} diff={gd} />
                  ))}
                </TableBody>
              </Table>
            </td>
          </motion.tr>
        )}
      </AnimatePresence>
    </>
  );
}

export function ComparisonPage() {
  const [runs, setRuns] = useState<RunSummary[]>([]);
  const [configNames, setConfigNames] = useState<string[]>([]);
  const [configA, setConfigA] = useState<string>("");
  const [configB, setConfigB] = useState<string>("");
  const [comparison, setComparison] = useState<ConfigComparison | null>(null);
  const [loading, setLoading] = useState(false);
  const [runsLoading, setRunsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  // Filters
  const [filterLanguage, setFilterLanguage] = useState<string>("all");
  const [filterService, setFilterService] = useState<string>("all");

  useEffect(() => {
    fetchRuns()
      .then((data) => {
        setRuns(data);
        setConfigNames(extractConfigNames(data));
      })
      .catch((e) => setError(e.message))
      .finally(() => setRunsLoading(false));
  }, []);

  useEffect(() => {
    if (!configA || !configB || configA === configB) {
      setComparison(null);
      return;
    }
    setLoading(true);
    setError(null);
    fetchCompareConfigs(configA, configB)
      .then(setComparison)
      .catch((e) => setError(e.message))
      .finally(() => setLoading(false));
  }, [configA, configB]);

  const filterValues = useMemo(
    () => extractFilterValues(comparison?.per_prompt ?? []),
    [comparison]
  );

  const filteredDiffs = useMemo(() => {
    if (!comparison) return [];
    return comparison.per_prompt.filter((d) => {
      const parts = d.prompt_id.split("-");
      if (filterLanguage !== "all" && parts.length >= 3 && parts[2] !== filterLanguage)
        return false;
      if (filterService !== "all" && parts.length >= 1 && parts[0] !== filterService)
        return false;
      return true;
    });
  }, [comparison, filterLanguage, filterService]);

  const filteredSummary = useMemo(() => {
    let improved = 0,
      regressed = 0,
      unchanged = 0,
      totalDelta = 0,
      paired = 0;
    for (const d of filteredDiffs) {
      if (d.only_in_a || d.only_in_b) continue;
      paired++;
      totalDelta += d.delta;
      if (d.delta > 0.001) improved++;
      else if (d.delta < -0.001) regressed++;
      else unchanged++;
    }
    return {
      avg_delta: paired > 0 ? totalDelta / paired : 0,
      improved,
      regressed,
      unchanged,
    };
  }, [filteredDiffs]);

  if (runsLoading) {
    return (
      <div className="flex min-h-screen items-center justify-center bg-[#0a0a0f]">
        <Loader2 className="h-6 w-6 animate-spin text-emerald-400" />
      </div>
    );
  }

  return (
    <div
      className="min-h-screen bg-[#0a0a0f] px-4 py-8 sm:px-6"
      style={{ fontFamily: "'Inter', sans-serif" }}
    >
      <div className="mx-auto max-w-6xl">
        {/* Header */}
        <div className="mb-8">
          <div className="mb-2 flex items-center gap-3">
            <GitCompareArrows className="h-6 w-6 text-emerald-400" />
            <h1
              className="text-white"
              style={{
                fontFamily: "'JetBrains Mono', monospace",
                fontSize: "clamp(1.5rem, 3vw, 2rem)",
              }}
            >
              Config Comparison
            </h1>
          </div>
          <p className="text-white/40" style={{ fontSize: 14 }}>
            Compare evaluation scores between two configurations side-by-side.
          </p>
        </div>

        {/* Config selectors */}
        <div className="mb-6 grid grid-cols-1 gap-4 sm:grid-cols-2">
          <div>
            <label
              className="mb-1.5 block text-white/50"
              style={{ fontSize: 12 }}
            >
              Config A (baseline)
            </label>
            <Select value={configA} onValueChange={setConfigA}>
              <SelectTrigger className="border-white/10 bg-white/[0.03] text-white">
                <SelectValue placeholder="Select config A…" />
              </SelectTrigger>
              <SelectContent className="border-white/10 bg-[#1a1a2e] text-white">
                {configNames.map((name) => (
                  <SelectItem key={name} value={name} disabled={name === configB}>
                    {name}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>
          <div>
            <label
              className="mb-1.5 block text-white/50"
              style={{ fontSize: 12 }}
            >
              Config B (comparison)
            </label>
            <Select value={configB} onValueChange={setConfigB}>
              <SelectTrigger className="border-white/10 bg-white/[0.03] text-white">
                <SelectValue placeholder="Select config B…" />
              </SelectTrigger>
              <SelectContent className="border-white/10 bg-[#1a1a2e] text-white">
                {configNames.map((name) => (
                  <SelectItem key={name} value={name} disabled={name === configA}>
                    {name}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>
        </div>

        {/* Loading */}
        {loading && (
          <div className="flex items-center justify-center py-16">
            <Loader2 className="h-6 w-6 animate-spin text-emerald-400" />
          </div>
        )}

        {/* Error */}
        {error && !loading && (
          <div className="rounded-xl border border-red-500/20 bg-red-500/5 p-6 text-center">
            <p className="mb-1 text-red-400">Comparison failed</p>
            <p className="text-white/40" style={{ fontSize: 13 }}>
              {error}
            </p>
          </div>
        )}

        {/* Empty state */}
        {!loading && !error && !comparison && configA && configB && configA === configB && (
          <div className="rounded-xl border border-white/8 bg-white/[0.03] p-8 text-center">
            <p className="text-white/40">Please select two different configurations to compare.</p>
          </div>
        )}

        {!loading && !error && !comparison && (!configA || !configB) && (
          <div className="rounded-xl border border-white/8 bg-white/[0.03] p-8 text-center">
            <GitCompareArrows className="mx-auto mb-3 h-8 w-8 text-white/20" />
            <p className="text-white/40">
              Select two configurations above to see a side-by-side comparison.
            </p>
          </div>
        )}

        {/* Results */}
        {comparison && !loading && (
          <motion.div
            initial={{ opacity: 0, y: 12 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.2 }}
          >
            {/* Summary stats */}
            <div className="mb-6 grid grid-cols-2 gap-3 sm:grid-cols-4">
              <div className="rounded-xl border border-white/8 bg-white/[0.03] p-4">
                <div className="mb-1 flex items-center gap-2">
                  <TrendingUp className="h-4 w-4 text-emerald-400" />
                  <span className="text-white/40" style={{ fontSize: 12 }}>
                    Avg Delta
                  </span>
                </div>
                <span
                  className={`${deltaColor(filteredSummary.avg_delta)}`}
                  style={{
                    fontFamily: "'JetBrains Mono', monospace",
                    fontSize: 20,
                  }}
                >
                  {formatDelta(filteredSummary.avg_delta)}
                </span>
              </div>
              <div className="rounded-xl border border-white/8 bg-white/[0.03] p-4">
                <div className="mb-1 flex items-center gap-2">
                  <ArrowUpRight className="h-4 w-4 text-emerald-400" />
                  <span className="text-white/40" style={{ fontSize: 12 }}>
                    Improved
                  </span>
                </div>
                <span
                  className="text-emerald-400"
                  style={{
                    fontFamily: "'JetBrains Mono', monospace",
                    fontSize: 20,
                  }}
                >
                  {filteredSummary.improved}
                </span>
              </div>
              <div className="rounded-xl border border-white/8 bg-white/[0.03] p-4">
                <div className="mb-1 flex items-center gap-2">
                  <ArrowDownRight className="h-4 w-4 text-red-400" />
                  <span className="text-white/40" style={{ fontSize: 12 }}>
                    Regressed
                  </span>
                </div>
                <span
                  className="text-red-400"
                  style={{
                    fontFamily: "'JetBrains Mono', monospace",
                    fontSize: 20,
                  }}
                >
                  {filteredSummary.regressed}
                </span>
              </div>
              <div className="rounded-xl border border-white/8 bg-white/[0.03] p-4">
                <div className="mb-1 flex items-center gap-2">
                  <Equal className="h-4 w-4 text-white/30" />
                  <span className="text-white/40" style={{ fontSize: 12 }}>
                    Unchanged
                  </span>
                </div>
                <span
                  className="text-white/50"
                  style={{
                    fontFamily: "'JetBrains Mono', monospace",
                    fontSize: 20,
                  }}
                >
                  {filteredSummary.unchanged}
                </span>
              </div>
            </div>

            {/* Filter controls */}
            <div className="mb-4 flex flex-wrap items-center gap-3">
              <Filter className="h-4 w-4 text-white/30" />
              <Select value={filterService} onValueChange={setFilterService}>
                <SelectTrigger className="w-40 border-white/10 bg-white/[0.03] text-white" size="sm">
                  <SelectValue />
                </SelectTrigger>
                <SelectContent className="border-white/10 bg-[#1a1a2e] text-white">
                  <SelectItem value="all">All Services</SelectItem>
                  {filterValues.services.map((s) => (
                    <SelectItem key={s} value={s}>
                      {s}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
              <Select value={filterLanguage} onValueChange={setFilterLanguage}>
                <SelectTrigger className="w-40 border-white/10 bg-white/[0.03] text-white" size="sm">
                  <SelectValue />
                </SelectTrigger>
                <SelectContent className="border-white/10 bg-[#1a1a2e] text-white">
                  <SelectItem value="all">All Languages</SelectItem>
                  {filterValues.languages.map((l) => (
                    <SelectItem key={l} value={l}>
                      {l}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
              {(filterLanguage !== "all" || filterService !== "all") && (
                <button
                  onClick={() => {
                    setFilterLanguage("all");
                    setFilterService("all");
                  }}
                  className="rounded-md px-2 py-1 text-white/40 transition-colors hover:bg-white/5 hover:text-white/60"
                  style={{ fontSize: 12 }}
                >
                  Clear filters
                </button>
              )}
              <span className="ml-auto text-white/30" style={{ fontSize: 12 }}>
                {filteredDiffs.length} prompt{filteredDiffs.length !== 1 ? "s" : ""}
              </span>
            </div>

            {/* Comparison table */}
            <div className="overflow-hidden rounded-xl border border-white/8 bg-white/[0.02]">
              <Table>
                <TableHeader>
                  <TableRow className="border-white/8 hover:bg-transparent">
                    <TableHead className="text-white/50" style={{ fontSize: 12 }}>
                      Prompt
                    </TableHead>
                    <TableHead className="text-white/50" style={{ fontSize: 12 }}>
                      {comparison.config_a}
                    </TableHead>
                    <TableHead className="text-white/50" style={{ fontSize: 12 }}>
                      {comparison.config_b}
                    </TableHead>
                    <TableHead className="text-white/50" style={{ fontSize: 12 }}>
                      Delta (B − A)
                    </TableHead>
                  </TableRow>
                </TableHeader>
                <TableBody>
                  {filteredDiffs.length === 0 ? (
                    <TableRow>
                      <TableCell colSpan={4} className="py-8 text-center text-white/30">
                        No prompts match the current filters.
                      </TableCell>
                    </TableRow>
                  ) : (
                    filteredDiffs.map((d) => (
                      <PromptDiffRow key={d.prompt_id} diff={d} />
                    ))
                  )}
                </TableBody>
              </Table>
            </div>
          </motion.div>
        )}
      </div>
    </div>
  );
}
