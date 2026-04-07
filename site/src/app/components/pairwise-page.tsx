import { useEffect, useState, useMemo } from "react";
import {
  BarChart,
  Bar,
  XAxis,
  YAxis,
  Tooltip,
  ResponsiveContainer,
  Cell,
  CartesianGrid,
} from "recharts";
import {
  ArrowUpRight,
  ArrowDownRight,
  Minus,
  Grid3X3,
  BarChart3,
  Zap,
  AlertTriangle,
} from "lucide-react";
import { fetchRuns } from "../data/api";
import type { RunSummary, PairwiseReport, ToolImpact } from "../data/types";

// ── Helpers ──────────────────────────────────────────────────────────

function impactColor(impact: number): string {
  if (impact > 5) return "#10b981";
  if (impact > 0) return "#6ee7b7";
  if (impact === 0) return "#6b7280";
  if (impact > -5) return "#fca5a5";
  return "#ef4444";
}

function impactBg(impact: number): string {
  if (impact > 5) return "rgba(16,185,129,0.25)";
  if (impact > 0) return "rgba(16,185,129,0.12)";
  if (impact === 0) return "rgba(107,114,128,0.15)";
  if (impact > -5) return "rgba(239,68,68,0.12)";
  return "rgba(239,68,68,0.25)";
}

function classifyTool(impact: number): "helper" | "hurter" | "neutral" {
  if (impact > 1) return "helper";
  if (impact < -1) return "hurter";
  return "neutral";
}

// ── Heatmap Cell ─────────────────────────────────────────────────────

function HeatmapCell({ impact }: { impact: number | null }) {
  if (impact === null) {
    return (
      <div className="flex h-10 w-full items-center justify-center rounded bg-white/[0.02] text-white/15" style={{ fontSize: 10 }}>
        —
      </div>
    );
  }
  return (
    <div
      className="flex h-10 w-full items-center justify-center rounded transition-colors"
      style={{ background: impactBg(impact), fontSize: 12, fontFamily: "'JetBrains Mono', monospace", color: impactColor(impact) }}
      title={`Impact: ${impact > 0 ? "+" : ""}${impact}`}
    >
      {impact > 0 ? "+" : ""}{impact}
    </div>
  );
}

// ── Impact Summary Card ──────────────────────────────────────────────

function ImpactSummaryCard({ impacts }: { impacts: ToolImpact[] }) {
  const helpers = impacts.filter((t) => classifyTool(t.impact) === "helper");
  const hurters = impacts.filter((t) => classifyTool(t.impact) === "hurter");
  const neutral = impacts.filter((t) => classifyTool(t.impact) === "neutral");

  return (
    <div className="grid gap-4 sm:grid-cols-3">
      {/* Helpers */}
      <div className="rounded-xl border border-emerald-500/15 bg-emerald-500/[0.03] p-5">
        <div className="mb-3 flex items-center gap-2">
          <ArrowUpRight className="h-4 w-4 text-emerald-400" />
          <span className="text-emerald-400" style={{ fontSize: 13, fontWeight: 600 }}>Top Helpers</span>
          <span className="ml-auto rounded-full bg-emerald-500/20 px-2 py-0.5 text-emerald-400" style={{ fontSize: 11, fontFamily: "'JetBrains Mono', monospace" }}>
            {helpers.length}
          </span>
        </div>
        {helpers.length === 0 ? (
          <p className="text-white/30" style={{ fontSize: 13 }}>No tools with positive impact</p>
        ) : (
          <div className="space-y-2">
            {helpers.slice(0, 5).map((t) => (
              <div key={t.tool_name} className="flex items-center justify-between">
                <span className="truncate text-white/70" style={{ fontSize: 13, maxWidth: "70%" }}>{t.tool_name}</span>
                <span className="text-emerald-400" style={{ fontSize: 12, fontFamily: "'JetBrains Mono', monospace" }}>
                  +{t.impact}
                </span>
              </div>
            ))}
          </div>
        )}
      </div>

      {/* Hurters */}
      <div className="rounded-xl border border-red-500/15 bg-red-500/[0.03] p-5">
        <div className="mb-3 flex items-center gap-2">
          <ArrowDownRight className="h-4 w-4 text-red-400" />
          <span className="text-red-400" style={{ fontSize: 13, fontWeight: 600 }}>Top Hurters</span>
          <span className="ml-auto rounded-full bg-red-500/20 px-2 py-0.5 text-red-400" style={{ fontSize: 11, fontFamily: "'JetBrains Mono', monospace" }}>
            {hurters.length}
          </span>
        </div>
        {hurters.length === 0 ? (
          <p className="text-white/30" style={{ fontSize: 13 }}>No tools with negative impact</p>
        ) : (
          <div className="space-y-2">
            {hurters.slice(0, 5).map((t) => (
              <div key={t.tool_name} className="flex items-center justify-between">
                <span className="truncate text-white/70" style={{ fontSize: 13, maxWidth: "70%" }}>{t.tool_name}</span>
                <span className="text-red-400" style={{ fontSize: 12, fontFamily: "'JetBrains Mono', monospace" }}>
                  {t.impact}
                </span>
              </div>
            ))}
          </div>
        )}
      </div>

      {/* Neutral */}
      <div className="rounded-xl border border-white/8 bg-white/[0.03] p-5">
        <div className="mb-3 flex items-center gap-2">
          <Minus className="h-4 w-4 text-white/40" />
          <span className="text-white/50" style={{ fontSize: 13, fontWeight: 600 }}>Neutral</span>
          <span className="ml-auto rounded-full bg-white/10 px-2 py-0.5 text-white/50" style={{ fontSize: 11, fontFamily: "'JetBrains Mono', monospace" }}>
            {neutral.length}
          </span>
        </div>
        {neutral.length === 0 ? (
          <p className="text-white/30" style={{ fontSize: 13 }}>All tools have measurable impact</p>
        ) : (
          <div className="space-y-2">
            {neutral.slice(0, 5).map((t) => (
              <div key={t.tool_name} className="flex items-center justify-between">
                <span className="truncate text-white/70" style={{ fontSize: 13, maxWidth: "70%" }}>{t.tool_name}</span>
                <span className="text-white/40" style={{ fontSize: 12, fontFamily: "'JetBrains Mono', monospace" }}>
                  {t.impact > 0 ? "+" : ""}{t.impact}
                </span>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}

// ── Contribution Bar Chart ───────────────────────────────────────────

function ContributionChart({ impacts }: { impacts: ToolImpact[] }) {
  const sorted = useMemo(
    () => [...impacts].sort((a, b) => b.impact - a.impact),
    [impacts],
  );

  const data = sorted.map((t) => ({
    name: t.tool_name,
    impact: t.impact,
    baseline: t.baseline_score,
    without: t.without_score,
  }));

  const maxAbs = Math.max(10, ...impacts.map((t) => Math.abs(t.impact)));
  const chartHeight = Math.max(250, sorted.length * 36);

  return (
    <ResponsiveContainer width="100%" height={chartHeight}>
      <BarChart data={data} layout="vertical" margin={{ left: 10, right: 20, top: 5, bottom: 5 }}>
        <CartesianGrid horizontal={false} stroke="rgba(255,255,255,0.04)" />
        <XAxis
          type="number"
          domain={[-maxAbs, maxAbs]}
          tick={{ fill: "rgba(255,255,255,0.35)", fontSize: 11 }}
          axisLine={false}
          tickLine={false}
        />
        <YAxis
          type="category"
          dataKey="name"
          width={140}
          tick={{ fill: "rgba(255,255,255,0.5)", fontSize: 11 }}
          axisLine={false}
          tickLine={false}
        />
        <Tooltip
          contentStyle={{
            background: "#1a1a2e",
            border: "1px solid rgba(255,255,255,0.1)",
            borderRadius: 8,
            color: "#fff",
            fontSize: 13,
          }}
          formatter={(value: number, name: string) => {
            if (name === "impact") return [`${value > 0 ? "+" : ""}${value}`, "Impact"];
            return [value, name];
          }}
          labelFormatter={(label) => `Tool: ${label}`}
        />
        <Bar dataKey="impact" radius={[0, 4, 4, 0]} maxBarSize={24}>
          {data.map((entry, i) => (
            <Cell key={i} fill={impactColor(entry.impact)} />
          ))}
        </Bar>
      </BarChart>
    </ResponsiveContainer>
  );
}

// ── Tool Impact Heatmap ──────────────────────────────────────────────

function ToolImpactHeatmap({ reports }: { reports: PairwiseReport[] }) {
  const toolNames = useMemo(() => {
    const set = new Set<string>();
    for (const r of reports) {
      for (const imp of r.impacts) set.add(imp.tool_name);
    }
    return [...set].sort();
  }, [reports]);

  const promptIds = useMemo(() => reports.map((r) => r.prompt_id), [reports]);

  // Build lookup: promptId -> toolName -> impact
  const lookup = useMemo(() => {
    const m = new Map<string, Map<string, number>>();
    for (const r of reports) {
      const inner = new Map<string, number>();
      for (const imp of r.impacts) inner.set(imp.tool_name, imp.impact);
      m.set(r.prompt_id, inner);
    }
    return m;
  }, [reports]);

  // Shorten prompt ID for display
  const shortId = (id: string) => {
    const parts = id.split("-");
    if (parts.length > 3) return parts.slice(-2).join("-");
    return id;
  };

  return (
    <div className="overflow-x-auto">
      <table style={{ fontSize: 12, borderCollapse: "separate", borderSpacing: 3 }}>
        <thead>
          <tr>
            <th className="sticky left-0 bg-[#0a0a0f] px-2 py-2 text-left text-white/30" style={{ fontSize: 11, fontWeight: 500, minWidth: 120 }}>
              Tool \ Prompt
            </th>
            {promptIds.map((pid) => (
              <th
                key={pid}
                className="px-1 py-2 text-center text-white/30"
                style={{ fontSize: 10, fontWeight: 400, minWidth: 60, maxWidth: 80, writingMode: "vertical-rl", textOrientation: "mixed", whiteSpace: "nowrap" }}
                title={pid}
              >
                {shortId(pid)}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {toolNames.map((tool) => (
            <tr key={tool}>
              <td className="sticky left-0 bg-[#0a0a0f] truncate px-2 py-1 text-white/60" style={{ fontSize: 12, maxWidth: 160 }} title={tool}>
                {tool}
              </td>
              {promptIds.map((pid) => {
                const impact = lookup.get(pid)?.get(tool) ?? null;
                return (
                  <td key={pid} className="px-0.5 py-0.5">
                    <HeatmapCell impact={impact} />
                  </td>
                );
              })}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

// ── Main Page ────────────────────────────────────────────────────────

export function PairwisePage() {
  const [runs, setRuns] = useState<RunSummary[]>([]);
  const [selectedRunId, setSelectedRunId] = useState<string>("");
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetchRuns()
      .then((data) => {
        setRuns(data);
        // Auto-select the most recent run that has pairwise data
        const withPairwise = data.filter((r) => r.pairwise_results && r.pairwise_results.length > 0);
        if (withPairwise.length > 0) setSelectedRunId(withPairwise[0].run_id);
      })
      .catch((e) => setError(e.message))
      .finally(() => setLoading(false));
  }, []);

  const selectedRun = runs.find((r) => r.run_id === selectedRunId);
  const pairwiseReports = selectedRun?.pairwise_results ?? [];

  // Aggregate impacts across all prompts in this run
  const aggregatedImpacts = useMemo(() => {
    if (pairwiseReports.length === 0) return [];

    const byTool = new Map<string, { total: number; baseTotal: number; withoutTotal: number; count: number; basePasses: number; withoutPasses: number }>();
    for (const report of pairwiseReports) {
      for (const imp of report.impacts) {
        const existing = byTool.get(imp.tool_name) ?? { total: 0, baseTotal: 0, withoutTotal: 0, count: 0, basePasses: 0, withoutPasses: 0 };
        existing.total += imp.impact;
        existing.baseTotal += imp.baseline_score;
        existing.withoutTotal += imp.without_score;
        existing.count += 1;
        if (imp.baseline_pass) existing.basePasses += 1;
        if (imp.without_pass) existing.withoutPasses += 1;
        byTool.set(imp.tool_name, existing);
      }
    }

    const result: ToolImpact[] = [];
    for (const [toolName, a] of byTool) {
      result.push({
        tool_name: toolName,
        impact: Math.round((a.total / a.count) * 10) / 10,
        baseline_score: Math.round((a.baseTotal / a.count) * 10) / 10,
        without_score: Math.round((a.withoutTotal / a.count) * 10) / 10,
        baseline_pass: a.basePasses === a.count,
        without_pass: a.withoutPasses === a.count,
      });
    }
    result.sort((a, b) => b.impact - a.impact || a.tool_name.localeCompare(b.tool_name));
    return result;
  }, [pairwiseReports]);

  const runsWithPairwise = runs.filter((r) => r.pairwise_results && r.pairwise_results.length > 0);

  if (loading) {
    return (
      <div className="flex min-h-screen items-center justify-center bg-[#0a0a0f]">
        <div className="text-white/40" style={{ fontSize: 14 }}>Loading pairwise data…</div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex min-h-screen items-center justify-center bg-[#0a0a0f]">
        <div className="text-red-400" style={{ fontSize: 14 }}>Error: {error}</div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-[#0a0a0f] px-4 py-8 sm:px-6" style={{ fontFamily: "'Inter', sans-serif" }}>
      <div className="mx-auto max-w-7xl">
        {/* Header */}
        <div className="mb-8 flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
          <div>
            <h1 className="flex items-center gap-3 text-white" style={{ fontFamily: "'JetBrains Mono', monospace", fontSize: "clamp(1.5rem, 3vw, 2rem)" }}>
              <Zap className="h-6 w-6 text-emerald-400" />
              Pairwise Impact
            </h1>
            <p className="mt-1 text-white/40" style={{ fontSize: 14 }}>
              Tool-by-tool impact analysis — what helps and what hurts
            </p>
          </div>

          {/* Run selector */}
          {runsWithPairwise.length > 0 && (
            <select
              value={selectedRunId}
              onChange={(e) => setSelectedRunId(e.target.value)}
              className="rounded-lg border border-white/10 bg-white/5 px-4 py-2 text-white/80 outline-none transition focus:border-emerald-500/30"
              style={{ fontSize: 13, fontFamily: "'JetBrains Mono', monospace" }}
            >
              {runsWithPairwise.map((r) => (
                <option key={r.run_id} value={r.run_id} className="bg-[#1a1a2e]">
                  {r.run_id} ({r.pairwise_results!.length} prompts)
                </option>
              ))}
            </select>
          )}
        </div>

        {/* Empty state */}
        {runsWithPairwise.length === 0 ? (
          <div className="flex flex-col items-center justify-center rounded-xl border border-white/8 bg-white/[0.02] py-20">
            <AlertTriangle className="mb-4 h-10 w-10 text-white/20" />
            <p className="mb-2 text-white/50" style={{ fontSize: 16 }}>No pairwise data available</p>
            <p className="text-white/30" style={{ fontSize: 13 }}>
              Run evaluations with pairwise mode enabled to see tool impact analysis
            </p>
          </div>
        ) : (
          <>
            {/* Impact Summary Cards */}
            <div className="mb-8">
              <ImpactSummaryCard impacts={aggregatedImpacts} />
            </div>

            {/* Contribution Bar Chart */}
            <div className="mb-8 rounded-xl border border-white/8 bg-white/[0.03] p-6">
              <div className="mb-4 flex items-center gap-2">
                <BarChart3 className="h-4 w-4 text-emerald-400" />
                <h3 className="text-white" style={{ fontSize: 15 }}>Tool Contribution</h3>
                <span className="ml-2 text-white/30" style={{ fontSize: 12 }}>
                  Average impact per tool (positive = helps, negative = hurts)
                </span>
              </div>
              {aggregatedImpacts.length > 0 ? (
                <ContributionChart impacts={aggregatedImpacts} />
              ) : (
                <p className="py-8 text-center text-white/30" style={{ fontSize: 13 }}>No impact data</p>
              )}
              {/* Legend */}
              <div className="mt-4 flex flex-wrap justify-center gap-4">
                {[
                  { label: "Strong positive (>5)", color: "#10b981" },
                  { label: "Positive", color: "#6ee7b7" },
                  { label: "Neutral", color: "#6b7280" },
                  { label: "Negative", color: "#fca5a5" },
                  { label: "Strong negative (<-5)", color: "#ef4444" },
                ].map((l) => (
                  <div key={l.label} className="flex items-center gap-1.5">
                    <div className="h-2.5 w-2.5 rounded-sm" style={{ background: l.color }} />
                    <span className="text-white/35" style={{ fontSize: 11 }}>{l.label}</span>
                  </div>
                ))}
              </div>
            </div>

            {/* Tool × Prompt Heatmap */}
            <div className="mb-8 rounded-xl border border-white/8 bg-white/[0.03] p-6">
              <div className="mb-4 flex items-center gap-2">
                <Grid3X3 className="h-4 w-4 text-emerald-400" />
                <h3 className="text-white" style={{ fontSize: 15 }}>Tool × Prompt Heatmap</h3>
                <span className="ml-2 text-white/30" style={{ fontSize: 12 }}>
                  Per-prompt impact scores for each tool
                </span>
              </div>
              {pairwiseReports.length > 0 ? (
                <ToolImpactHeatmap reports={pairwiseReports} />
              ) : (
                <p className="py-8 text-center text-white/30" style={{ fontSize: 13 }}>No heatmap data</p>
              )}
            </div>

            {/* Detailed Table */}
            <div className="rounded-xl border border-white/8 bg-white/[0.03] p-6">
              <h3 className="mb-4 text-white" style={{ fontSize: 15 }}>Detailed Impact Scores</h3>
              <div className="overflow-x-auto">
                <table className="w-full" style={{ fontSize: 13 }}>
                  <thead>
                    <tr className="border-b border-white/8">
                      {["Tool", "Avg Impact", "Baseline", "Without", "Δ Pass", "Classification"].map((h) => (
                        <th key={h} className="px-3 py-2.5 text-left text-white/30" style={{ fontWeight: 500, fontSize: 11 }}>
                          {h}
                        </th>
                      ))}
                    </tr>
                  </thead>
                  <tbody>
                    {aggregatedImpacts.map((t) => {
                      const cls = classifyTool(t.impact);
                      return (
                        <tr key={t.tool_name} className="border-b border-white/5 transition hover:bg-white/[0.02]">
                          <td className="px-3 py-3 text-white/70" style={{ fontFamily: "'JetBrains Mono', monospace", fontSize: 12 }}>
                            {t.tool_name}
                          </td>
                          <td className="px-3 py-3" style={{ fontFamily: "'JetBrains Mono', monospace" }}>
                            <span style={{ color: impactColor(t.impact) }}>
                              {t.impact > 0 ? "+" : ""}{t.impact}
                            </span>
                          </td>
                          <td className="px-3 py-3 text-white/50" style={{ fontFamily: "'JetBrains Mono', monospace", fontSize: 12 }}>
                            {t.baseline_score}
                          </td>
                          <td className="px-3 py-3 text-white/50" style={{ fontFamily: "'JetBrains Mono', monospace", fontSize: 12 }}>
                            {t.without_score}
                          </td>
                          <td className="px-3 py-3">
                            {t.baseline_pass && !t.without_pass ? (
                              <span className="rounded bg-red-500/15 px-1.5 py-0.5 text-red-400" style={{ fontSize: 11 }}>Breaks</span>
                            ) : !t.baseline_pass && t.without_pass ? (
                              <span className="rounded bg-amber-500/15 px-1.5 py-0.5 text-amber-400" style={{ fontSize: 11 }}>Fixes</span>
                            ) : (
                              <span className="text-white/25" style={{ fontSize: 11 }}>—</span>
                            )}
                          </td>
                          <td className="px-3 py-3">
                            <span
                              className="inline-flex items-center gap-1 rounded-full px-2 py-0.5"
                              style={{
                                fontSize: 11,
                                background: cls === "helper" ? "rgba(16,185,129,0.15)" : cls === "hurter" ? "rgba(239,68,68,0.15)" : "rgba(107,114,128,0.15)",
                                color: cls === "helper" ? "#10b981" : cls === "hurter" ? "#ef4444" : "#9ca3af",
                              }}
                            >
                              {cls === "helper" && <ArrowUpRight className="h-3 w-3" />}
                              {cls === "hurter" && <ArrowDownRight className="h-3 w-3" />}
                              {cls === "neutral" && <Minus className="h-3 w-3" />}
                              {cls.charAt(0).toUpperCase() + cls.slice(1)}
                            </span>
                          </td>
                        </tr>
                      );
                    })}
                  </tbody>
                </table>
              </div>
            </div>
          </>
        )}
      </div>
    </div>
  );
}
