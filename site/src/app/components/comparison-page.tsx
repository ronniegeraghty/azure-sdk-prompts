// Comparison page — group-based comparison tool (#365 / WI-047).
//
// Replaces the old A vs B config picker with a flexible system:
//   • Users define named ComparisonGroups by filtering on any prompt or
//     config attribute (catalog is computed from real eval data).
//   • Multiple charts can be toggled to visualize differences across groups.
//   • Group definitions and chart selections persist in localStorage.

import { useEffect, useMemo, useState } from "react";
import {
  BarChart,
  Bar,
  Cell,
  XAxis,
  YAxis,
  Tooltip,
  Legend,
  ResponsiveContainer,
  CartesianGrid,
} from "recharts";
import { GitCompareArrows, Loader2, Plus, Settings2 } from "lucide-react";
import { fetchRuns } from "../data/api";
import type { RunSummary } from "../data/types";
import {
  buildCatalog,
  CHART_OPTIONS,
  computeMetrics,
  type ChartId,
  type ComparisonGroup,
  type GroupMetrics,
  filterGroup,
  flattenResults,
  loadState,
  newGroupId,
  nextGroupColor,
  saveState,
} from "../lib/comparison-groups";
import { GroupBuilder } from "./group-builder";

const mono = { fontFamily: "'JetBrains Mono', monospace" };

const DEFAULT_CHARTS: ChartId[] = ["pass_rate", "avg_score", "by_service"];

function pct(v: number): string {
  return `${(v * 100).toFixed(1)}%`;
}

interface NamedMetrics {
  group: ComparisonGroup;
  metrics: GroupMetrics;
}

export function ComparisonPage() {
  const [runs, setRuns] = useState<RunSummary[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const [groups, setGroups] = useState<ComparisonGroup[]>([]);
  const [charts, setCharts] = useState<ChartId[]>(DEFAULT_CHARTS);
  const [hydrated, setHydrated] = useState(false);

  // Hydrate from localStorage on mount.
  useEffect(() => {
    const persisted = loadState();
    if (persisted) {
      setGroups(persisted.groups);
      if (persisted.charts.length > 0) setCharts(persisted.charts);
    }
    setHydrated(true);
  }, []);

  // Persist whenever groups or charts change (after initial hydration).
  useEffect(() => {
    if (!hydrated) return;
    saveState({ groups, charts });
  }, [groups, charts, hydrated]);

  useEffect(() => {
    fetchRuns()
      .then((data) => setRuns(data))
      .catch((e) => setError(e.message))
      .finally(() => setLoading(false));
  }, []);

  const allResults = useMemo(() => flattenResults(runs), [runs]);
  const catalog = useMemo(() => buildCatalog(allResults), [allResults]);

  const namedMetrics: NamedMetrics[] = useMemo(
    () =>
      groups.map((g) => ({
        group: g,
        metrics: computeMetrics(filterGroup(allResults, g)),
      })),
    [groups, allResults]
  );

  function addGroup() {
    const next: ComparisonGroup = {
      id: newGroupId(),
      name: `Group ${groups.length + 1}`,
      color: nextGroupColor(groups),
      filters: {},
    };
    setGroups([...groups, next]);
  }
  function updateGroup(id: string, next: ComparisonGroup) {
    setGroups(groups.map((g) => (g.id === id ? next : g)));
  }
  function removeGroup(id: string) {
    setGroups(groups.filter((g) => g.id !== id));
  }
  function toggleChart(id: ChartId) {
    setCharts((prev) => (prev.includes(id) ? prev.filter((c) => c !== id) : [...prev, id]));
  }

  if (loading) {
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
      <div className="mx-auto max-w-7xl">
        {/* Header */}
        <div className="mb-8">
          <div className="mb-2 flex items-center gap-3">
            <GitCompareArrows className="h-6 w-6 text-emerald-400" />
            <h1 className="text-white" style={{ ...mono, fontSize: "clamp(1.5rem, 3vw, 2rem)" }}>
              Compare
            </h1>
          </div>
          <p className="text-white/40" style={{ fontSize: 14 }}>
            Define one or more groups by filtering on any property — model, service, language, plane,
            category, or difficulty — then compare metrics across them. Group definitions persist in
            your browser.
          </p>
        </div>

        {error && (
          <div className="mb-6 rounded-xl border border-red-500/20 bg-red-500/5 p-4">
            <p className="text-red-400" style={{ fontSize: 13 }}>
              {error}
            </p>
          </div>
        )}

        {allResults.length === 0 && !error && (
          <div className="rounded-xl border border-white/8 bg-white/[0.03] p-8 text-center">
            <p className="text-white/40">No evaluations available yet. Run an evaluation first.</p>
          </div>
        )}

        {allResults.length > 0 && (
          <div className="grid grid-cols-1 gap-6 lg:grid-cols-[400px_1fr]">
            {/* Left column: groups + chart toggles */}
            <div className="space-y-4">
              <div className="flex items-center justify-between">
                <h2 className="text-white/80" style={{ ...mono, fontSize: 14 }}>
                  Groups ({groups.length})
                </h2>
                <button
                  onClick={addGroup}
                  className="inline-flex items-center gap-1 rounded-md border border-emerald-400/40 bg-emerald-500/10 px-3 py-1 text-emerald-300 transition-colors hover:bg-emerald-500/20"
                  style={{ fontSize: 12 }}
                >
                  <Plus className="h-3.5 w-3.5" /> Add group
                </button>
              </div>

              {groups.length === 0 ? (
                <div className="rounded-xl border border-dashed border-white/10 bg-white/[0.02] p-6 text-center">
                  <p className="mb-3 text-white/50" style={{ fontSize: 13 }}>
                    No groups yet. Add one to start comparing.
                  </p>
                  <button
                    onClick={addGroup}
                    className="inline-flex items-center gap-1 rounded-md bg-emerald-500/20 px-3 py-1.5 text-emerald-300 hover:bg-emerald-500/30"
                    style={{ fontSize: 12 }}
                  >
                    <Plus className="h-3.5 w-3.5" /> Create first group
                  </button>
                </div>
              ) : (
                groups.map((g) => (
                  <GroupBuilder
                    key={g.id}
                    group={g}
                    catalog={catalog}
                    matchCount={namedMetrics.find((nm) => nm.group.id === g.id)?.metrics.count ?? 0}
                    totalCount={allResults.length}
                    onChange={(next) => updateGroup(g.id, next)}
                    onRemove={() => removeGroup(g.id)}
                  />
                ))
              )}

              {/* Chart toggles */}
              <div className="rounded-xl border border-white/8 bg-white/[0.02] p-4">
                <div className="mb-3 flex items-center gap-2">
                  <Settings2 className="h-4 w-4 text-white/40" />
                  <h3 className="text-white/70" style={{ ...mono, fontSize: 13 }}>
                    Visualizations
                  </h3>
                </div>
                <div className="space-y-2">
                  {CHART_OPTIONS.map((c) => {
                    const on = charts.includes(c.id);
                    return (
                      <label
                        key={c.id}
                        className="flex cursor-pointer items-start gap-2 rounded-md p-1 hover:bg-white/[0.02]"
                      >
                        <input
                          type="checkbox"
                          checked={on}
                          onChange={() => toggleChart(c.id)}
                          aria-label={c.label}
                          className="mt-0.5 h-3.5 w-3.5 accent-emerald-400"
                        />
                        <span className="flex-1">
                          <span className="block text-white/80" style={{ fontSize: 12 }}>
                            {c.label}
                          </span>
                          <span className="block text-white/40" style={{ fontSize: 11 }}>
                            {c.description}
                          </span>
                        </span>
                      </label>
                    );
                  })}
                </div>
              </div>
            </div>

            {/* Right column: comparison output */}
            <div className="space-y-6">
              <ComparisonOutput groups={groups} metrics={namedMetrics} charts={charts} />
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

// ── Comparison output ────────────────────────────────────────────────

function ComparisonOutput({
  groups,
  metrics,
  charts,
}: {
  groups: ComparisonGroup[];
  metrics: NamedMetrics[];
  charts: ChartId[];
}) {
  if (groups.length === 0) {
    return (
      <div className="rounded-xl border border-white/8 bg-white/[0.03] p-8 text-center">
        <GitCompareArrows className="mx-auto mb-3 h-8 w-8 text-white/20" />
        <p className="text-white/40">Add a comparison group to begin.</p>
      </div>
    );
  }

  const counts = metrics.map((m) => m.metrics.count);
  const allEmpty = counts.every((c) => c === 0);
  const someEmpty = counts.some((c) => c === 0);
  const uneven =
    counts.length > 1 &&
    Math.max(...counts) > 0 &&
    Math.min(...counts) / Math.max(...counts) < 0.5;

  return (
    <>
      {/* Summary cards */}
      <div className="grid grid-cols-1 gap-3 sm:grid-cols-2 xl:grid-cols-3">
        {metrics.map(({ group, metrics: m }) => (
          <SummaryCard key={group.id} group={group} metrics={m} />
        ))}
      </div>

      {/* Notices */}
      {groups.length === 1 && (
        <div
          className="rounded-md border border-sky-400/30 bg-sky-500/5 px-3 py-2 text-sky-200"
          style={{ fontSize: 12 }}
        >
          Only one group defined — add another to see side-by-side comparisons.
        </div>
      )}
      {allEmpty && (
        <div
          className="rounded-md border border-amber-500/30 bg-amber-500/5 px-3 py-2 text-amber-200"
          style={{ fontSize: 12 }}
        >
          None of the defined groups match any evals. Adjust filters to see comparisons.
        </div>
      )}
      {!allEmpty && someEmpty && (
        <div
          className="rounded-md border border-amber-500/30 bg-amber-500/5 px-3 py-2 text-amber-200"
          style={{ fontSize: 12 }}
        >
          Some groups have no matching evals. Their bars are hidden from charts.
        </div>
      )}
      {!allEmpty && !someEmpty && uneven && (
        <div
          className="rounded-md border border-white/10 bg-white/[0.03] px-3 py-2 text-white/50"
          style={{ fontSize: 12 }}
        >
          Groups have uneven eval counts ({counts.join(" / ")}). Pass-rate and average-score charts
          are normalized, but raw counts vary — interpret with care.
        </div>
      )}

      {/* Charts */}
      {charts.length === 0 ? (
        <div
          className="rounded-xl border border-dashed border-white/10 bg-white/[0.02] p-6 text-center text-white/40"
          style={{ fontSize: 13 }}
        >
          No visualizations selected. Pick at least one from the Visualizations panel.
        </div>
      ) : (
        charts.map((c) => <ChartPanel key={c} chartId={c} metrics={metrics} />)
      )}
    </>
  );
}

function SummaryCard({ group, metrics: m }: { group: ComparisonGroup; metrics: GroupMetrics }) {
  return (
    <div className="rounded-xl border border-white/8 bg-white/[0.03] p-4" data-testid="summary-card">
      <div className="mb-2 flex items-center gap-2">
        <span className="h-2.5 w-2.5 rounded-full" style={{ background: group.color }} />
        <span className="truncate text-white" style={{ ...mono, fontSize: 13 }}>
          {group.name}
        </span>
      </div>
      <div className="grid grid-cols-2 gap-3">
        <Metric label="Pass rate" value={pct(m.pass_rate)} mono />
        <Metric label="Avg score" value={pct(m.avg_score)} mono />
        <Metric label="Evals" value={String(m.count)} mono />
        <Metric label="Avg duration" value={`${m.avg_duration.toFixed(1)}s`} mono />
      </div>
    </div>
  );
}

function Metric({ label, value, mono: useMono }: { label: string; value: string; mono?: boolean }) {
  return (
    <div>
      <div className="text-white/40" style={{ fontSize: 11 }}>
        {label}
      </div>
      <div className="text-white" style={useMono ? { ...mono, fontSize: 16 } : { fontSize: 16 }}>
        {value}
      </div>
    </div>
  );
}

// ── Chart panels ────────────────────────────────────────────────────

interface ChartRow {
  name: string;
  [groupId: string]: string | number;
}

const tooltipStyle = {
  background: "#1a1a2e",
  border: "1px solid rgba(255,255,255,0.1)",
  borderRadius: 8,
  color: "#fff",
  fontSize: 13,
};

function ChartPanel({ chartId, metrics }: { chartId: ChartId; metrics: NamedMetrics[] }) {
  const opt = CHART_OPTIONS.find((c) => c.id === chartId)!;
  const populated = metrics.filter((m) => m.metrics.count > 0);

  let chart: React.ReactNode = null;
  if (populated.length === 0) {
    chart = (
      <div className="flex h-32 items-center justify-center text-white/30" style={{ fontSize: 12 }}>
        No data — none of the groups match any evals.
      </div>
    );
  } else if (chartId === "pass_rate") {
    chart = renderSingleMetricBar(populated, (m) => m.pass_rate * 100, "Pass %", true);
  } else if (chartId === "avg_score") {
    chart = renderSingleMetricBar(populated, (m) => m.avg_score * 100, "Score %", true);
  } else if (chartId === "eval_count") {
    chart = renderSingleMetricBar(populated, (m) => m.count, "Evals", false);
  } else if (chartId === "by_service") {
    chart = renderBreakdown(populated, "by_service");
  } else if (chartId === "by_language") {
    chart = renderBreakdown(populated, "by_language");
  } else if (chartId === "score_distribution") {
    chart = renderDistribution(populated);
  }

  return (
    <div
      className="rounded-xl border border-white/8 bg-white/[0.03] p-5"
      data-testid={`chart-${chartId}`}
    >
      <h3 className="mb-1 text-white" style={{ ...mono, fontSize: 14 }}>
        {opt.label}
      </h3>
      <p className="mb-4 text-white/40" style={{ fontSize: 12 }}>
        {opt.description}
      </p>
      {chart}
    </div>
  );
}

function renderSingleMetricBar(
  populated: NamedMetrics[],
  selector: (m: GroupMetrics) => number,
  label: string,
  percent: boolean
) {
  const data = populated.map((m) => ({
    name: m.group.name,
    value: Number(selector(m.metrics).toFixed(2)),
    color: m.group.color,
  }));
  return (
    <ResponsiveContainer width="100%" height={240}>
      <BarChart data={data} margin={{ top: 8, right: 8, left: 0, bottom: 8 }}>
        <CartesianGrid stroke="rgba(255,255,255,0.05)" vertical={false} />
        <XAxis
          dataKey="name"
          tick={{ fill: "rgba(255,255,255,0.5)", fontSize: 11 }}
          axisLine={false}
          tickLine={false}
        />
        <YAxis
          domain={percent ? [0, 100] : ["auto", "auto"]}
          tick={{ fill: "rgba(255,255,255,0.4)", fontSize: 11 }}
          axisLine={false}
          tickLine={false}
        />
        <Tooltip
          contentStyle={tooltipStyle}
          formatter={(v: number) => (percent ? `${v.toFixed(1)}%` : String(v))}
        />
        <Bar dataKey="value" name={label} radius={[6, 6, 0, 0]}>
          {data.map((d, i) => (
            <Cell key={i} fill={d.color} />
          ))}
        </Bar>
      </BarChart>
    </ResponsiveContainer>
  );
}

// Grouped breakdown chart (pass-rate per category, with one series per group).
function renderBreakdown(populated: NamedMetrics[], dim: "by_service" | "by_language") {
  const cats = new Set<string>();
  for (const m of populated) for (const k of Object.keys(m.metrics[dim])) cats.add(k);
  const sortedCats = Array.from(cats).sort();

  const data: ChartRow[] = sortedCats.map((cat) => {
    const row: ChartRow = { name: cat };
    for (const m of populated) {
      const stat = m.metrics[dim][cat];
      row[m.group.id] = stat ? Number(((stat.passed / stat.total) * 100).toFixed(1)) : 0;
    }
    return row;
  });

  return (
    <ResponsiveContainer width="100%" height={Math.max(260, sortedCats.length * 36)}>
      <BarChart data={data} margin={{ top: 8, right: 8, left: 0, bottom: 8 }} layout="vertical">
        <CartesianGrid stroke="rgba(255,255,255,0.05)" horizontal={false} />
        <XAxis
          type="number"
          domain={[0, 100]}
          tick={{ fill: "rgba(255,255,255,0.4)", fontSize: 11 }}
          axisLine={false}
          tickLine={false}
        />
        <YAxis
          type="category"
          dataKey="name"
          tick={{ fill: "rgba(255,255,255,0.5)", fontSize: 11 }}
          axisLine={false}
          tickLine={false}
          width={120}
        />
        <Tooltip contentStyle={tooltipStyle} formatter={(v: number) => `${v.toFixed(1)}%`} />
        <Legend wrapperStyle={{ color: "rgba(255,255,255,0.6)", fontSize: 12 }} />
        {populated.map((m) => (
          <Bar
            key={m.group.id}
            dataKey={m.group.id}
            name={m.group.name}
            fill={m.group.color}
            radius={[0, 4, 4, 0]}
          />
        ))}
      </BarChart>
    </ResponsiveContainer>
  );
}

function renderDistribution(populated: NamedMetrics[]) {
  const bins = [
    "0–10%",
    "10–20%",
    "20–30%",
    "30–40%",
    "40–50%",
    "50–60%",
    "60–70%",
    "70–80%",
    "80–90%",
    "90–100%",
  ];
  const data: ChartRow[] = bins.map((label, i) => {
    const row: ChartRow = { name: label };
    for (const m of populated) row[m.group.id] = m.metrics.score_distribution[i];
    return row;
  });
  return (
    <ResponsiveContainer width="100%" height={260}>
      <BarChart data={data} margin={{ top: 8, right: 8, left: 0, bottom: 8 }}>
        <CartesianGrid stroke="rgba(255,255,255,0.05)" vertical={false} />
        <XAxis
          dataKey="name"
          tick={{ fill: "rgba(255,255,255,0.5)", fontSize: 10 }}
          axisLine={false}
          tickLine={false}
        />
        <YAxis
          tick={{ fill: "rgba(255,255,255,0.4)", fontSize: 11 }}
          axisLine={false}
          tickLine={false}
          allowDecimals={false}
        />
        <Tooltip contentStyle={tooltipStyle} />
        <Legend wrapperStyle={{ color: "rgba(255,255,255,0.6)", fontSize: 12 }} />
        {populated.map((m) => (
          <Bar
            key={m.group.id}
            dataKey={m.group.id}
            name={m.group.name}
            fill={m.group.color}
            radius={[4, 4, 0, 0]}
          />
        ))}
      </BarChart>
    </ResponsiveContainer>
  );
}
