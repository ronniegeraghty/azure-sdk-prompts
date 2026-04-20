import { useState, useEffect, useMemo } from "react";
import { Link } from "react-router";
import { BarChart, Bar, XAxis, YAxis, Tooltip, ResponsiveContainer, LineChart, Line } from "recharts";
import { CheckCircle2, XCircle, Clock, FileCode2, Cpu, TrendingUp, Loader2, Calendar } from "lucide-react";
import { fetchRuns } from "../data/api";
import type { RunSummary } from "../data/types";

const mono = { fontFamily: "'JetBrains Mono', monospace" };

export function DashboardPage() {
  const [runs, setRuns] = useState<RunSummary[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [activeChart, setActiveChart] = useState<"service" | "language">("service");

  useEffect(() => {
    let cancelled = false;
    async function load() {
      try {
        const data = await fetchRuns();
        if (cancelled) return;
        setRuns(data);
      } catch (e) {
        if (!cancelled) setError(e instanceof Error ? e.message : "Failed to load runs");
      } finally {
        if (!cancelled) setLoading(false);
      }
    }
    load();
    return () => { cancelled = true; };
  }, []);

  // Compute dashboard metrics from real run data
  const metrics = useMemo(() => {
    if (runs.length === 0) {
      return {
        totalEvals: 0,
        passRate: 0,
        avgDuration: 0,
        modelsCount: 0,
        passRateByService: [],
        passRateByLang: [],
        durationTrend: [],
        recentEvals: [],
        lastUpdated: null,
      };
    }

    let totalEvals = 0;
    let totalPassed = 0;
    let totalDuration = 0;
    const models = new Set<string>();
    const serviceStats: Record<string, { total: number; passed: number }> = {};
    const langStats: Record<string, { total: number; passed: number }> = {};
    const recentEvalsList: Array<{
      id: string;
      runId: string;
      prompt: string;
      lang: string;
      config: string;
      score: number;
      pass: boolean;
      duration: string;
      files: number;
    }> = [];

    // Sort runs by timestamp (most recent first) for recent evals
    const sortedRuns = [...runs].sort((a, b) => 
      new Date(b.timestamp).getTime() - new Date(a.timestamp).getTime()
    );

    for (const run of runs) {
      totalEvals += run.total_evaluations;
      totalPassed += run.passed;
      totalDuration += run.duration_seconds;

      for (const result of run.results || []) {
        const meta = result.prompt_metadata;
        
        // Track models (use config_name as proxy for model if not available)
        models.add(result.config_name);

        // Service stats
        const service = meta.service || "unknown";
        if (!serviceStats[service]) serviceStats[service] = { total: 0, passed: 0 };
        serviceStats[service].total++;
        if (result.success) serviceStats[service].passed++;

        // Language stats
        const lang = meta.language || "unknown";
        if (!langStats[lang]) langStats[lang] = { total: 0, passed: 0 };
        langStats[lang].total++;
        if (result.success) langStats[lang].passed++;
      }
    }

    // Recent evals (last 10)
    for (const run of sortedRuns) {
      for (const result of run.results || []) {
        if (recentEvalsList.length >= 10) break;
        recentEvalsList.push({
          id: `${run.run_id.slice(0, 8)}-${result.prompt_id.slice(0, 12)}`,
          runId: run.run_id,
          prompt: result.prompt_id,
          lang: result.prompt_metadata.language || "unknown",
          config: result.config_name,
          score: result.review?.overall_score || 0,
          pass: result.success,
          duration: `${result.duration_seconds?.toFixed(1) || 0}s`,
          files: result.generated_files?.length || 0,
        });
      }
      if (recentEvalsList.length >= 10) break;
    }

    const overallPassRate = totalEvals > 0 ? parseFloat(((totalPassed / totalEvals) * 100).toFixed(1)) : 0;
    const avgDuration = runs.length > 0 ? parseFloat((totalDuration / runs.length).toFixed(1)) : 0;

    const passRateByService = Object.entries(serviceStats)
      .map(([name, stats]) => ({
        name,
        rate: parseFloat(((stats.passed / stats.total) * 100).toFixed(1)),
        total: stats.total,
      }))
      .sort((a, b) => b.rate - a.rate);

    const passRateByLang = Object.entries(langStats)
      .map(([name, stats]) => ({
        name,
        rate: parseFloat(((stats.passed / stats.total) * 100).toFixed(1)),
        total: stats.total,
      }))
      .sort((a, b) => b.rate - a.rate);

    // Duration trend (last 10 runs)
    const durationTrend = sortedRuns.slice(0, 10).reverse().map((run, idx) => ({
      run: run.run_id.slice(0, 8),
      duration: parseFloat(run.duration_seconds.toFixed(1)),
      gen: parseFloat((run.avg_generation_duration_seconds || 0).toFixed(1)),
      review: parseFloat((run.avg_review_duration_seconds || 0).toFixed(1)),
    }));

    const lastUpdated = sortedRuns.length > 0 ? sortedRuns[0].timestamp : null;

    return {
      totalEvals,
      passRate: overallPassRate,
      avgDuration,
      modelsCount: models.size,
      passRateByService,
      passRateByLang,
      durationTrend,
      recentEvals: recentEvalsList,
      lastUpdated,
    };
  }, [runs]);

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

  const stats = [
    { label: "Total Evaluations", value: metrics.totalEvals.toLocaleString(), icon: FileCode2, color: "text-blue-400" },
    { label: "Overall Pass Rate", value: `${metrics.passRate}%`, icon: CheckCircle2, color: metrics.passRate >= 80 ? "text-emerald-400" : metrics.passRate >= 60 ? "text-amber-400" : "text-red-400" },
    { label: "Avg Duration", value: `${metrics.avgDuration}s`, icon: Clock, color: "text-amber-400" },
    { label: "Models Tested", value: metrics.modelsCount, icon: Cpu, color: "text-purple-400" },
  ];

  return (
    <div className="min-h-screen bg-[#0a0a0f] px-4 py-8 sm:px-6" style={{ fontFamily: "'Inter', sans-serif" }}>
      <div className="mx-auto max-w-7xl">
        <div className="mb-8 flex flex-col gap-4 sm:flex-row sm:items-center sm:justify-between">
          <div>
            <h1 className="text-white" style={{ ...mono, fontSize: "clamp(1.5rem, 3vw, 2rem)" }}>
              Evaluation Dashboard
            </h1>
            {metrics.lastUpdated ? (
              <div className="flex items-center gap-1.5 text-white/40" style={{ fontSize: 14 }}>
                <Calendar className="h-3.5 w-3.5" />
                <span>Last updated: {new Date(metrics.lastUpdated).toLocaleString()}</span>
              </div>
            ) : (
              <p className="text-white/40" style={{ fontSize: 14 }}>No evaluation data available</p>
            )}
          </div>
        </div>

        {/* Stats */}
        <div className="mb-8 grid grid-cols-2 gap-4 lg:grid-cols-4">
          {stats.map((s) => (
            <div key={s.label} className="rounded-xl border border-white/8 bg-white/[0.03] p-5">
              <div className="mb-3 flex items-center gap-2">
                <s.icon className={`h-4 w-4 ${s.color}`} />
                <span className="text-white/40" style={{ fontSize: 12 }}>{s.label}</span>
              </div>
              <span className={`text-white ${s.color}`} style={{ ...mono, fontSize: 24 }}>{s.value}</span>
            </div>
          ))}
        </div>

        {/* Charts */}
        <div className="mb-8 grid gap-6 lg:grid-cols-2">
          {/* Pass Rate */}
          <div className="rounded-xl border border-white/8 bg-white/[0.03] p-6">
            <div className="mb-4 flex items-center justify-between">
              <h3 className="text-white" style={{ fontSize: 15 }}>Pass Rate</h3>
              <div className="flex gap-1 rounded-lg bg-white/5 p-0.5">
                {(["service", "language"] as const).map((t) => (
                  <button
                    key={t}
                    onClick={() => setActiveChart(t)}
                    className={`rounded-md px-3 py-1 capitalize transition ${activeChart === t ? "bg-emerald-500/20 text-emerald-400" : "text-white/40 hover:text-white/60"}`}
                    style={{ fontSize: 12 }}
                  >
                    {t}
                  </button>
                ))}
              </div>
            </div>
            <ResponsiveContainer width="100%" height={250}>
              <BarChart data={activeChart === "service" ? metrics.passRateByService : metrics.passRateByLang}>
                <XAxis dataKey="name" tick={{ fill: "rgba(255,255,255,0.35)", fontSize: 11 }} axisLine={false} tickLine={false} />
                <YAxis domain={[0, 100]} tick={{ fill: "rgba(255,255,255,0.35)", fontSize: 11 }} axisLine={false} tickLine={false} />
                <Tooltip contentStyle={{ background: "#1a1a2e", border: "1px solid rgba(255,255,255,0.1)", borderRadius: 8, color: "#fff", fontSize: 13 }} />
                <Bar dataKey="rate" fill="#10b981" radius={[6, 6, 0, 0]} />
              </BarChart>
            </ResponsiveContainer>
          </div>

          {/* Duration Trend */}
          <div className="rounded-xl border border-white/8 bg-white/[0.03] p-6">
            <h3 className="mb-4 text-white" style={{ fontSize: 15 }}>Duration Trends (seconds)</h3>
            <ResponsiveContainer width="100%" height={250}>
              <LineChart data={metrics.durationTrend}>
                <XAxis dataKey="run" tick={{ fill: "rgba(255,255,255,0.35)", fontSize: 11 }} axisLine={false} tickLine={false} />
                <YAxis tick={{ fill: "rgba(255,255,255,0.35)", fontSize: 11 }} axisLine={false} tickLine={false} />
                <Tooltip contentStyle={{ background: "#1a1a2e", border: "1px solid rgba(255,255,255,0.1)", borderRadius: 8, color: "#fff", fontSize: 13 }} />
                <Line type="monotone" dataKey="gen" stroke="#10b981" strokeWidth={2} dot={false} name="Generation" />
                <Line type="monotone" dataKey="review" stroke="#8b5cf6" strokeWidth={2} dot={false} name="Review" />
              </LineChart>
            </ResponsiveContainer>
            <div className="mt-3 flex justify-center gap-5">
              {[{ label: "Generation", color: "#10b981" }, { label: "Review", color: "#8b5cf6" }].map((l) => (
                <div key={l.label} className="flex items-center gap-1.5">
                  <div className="h-2 w-2 rounded-full" style={{ background: l.color }} />
                  <span className="text-white/40" style={{ fontSize: 11 }}>{l.label}</span>
                </div>
              ))}
            </div>
          </div>
        </div>

        {/* Recent Evaluations Table */}
        <div className="mb-8">
          <div className="rounded-xl border border-white/8 bg-white/[0.03] p-6">
            <h3 className="mb-4 text-white" style={{ fontSize: 15 }}>Recent Evaluations</h3>
            {metrics.recentEvals.length > 0 ? (
              <div className="overflow-x-auto">
                <table className="w-full" style={{ fontSize: 13 }}>
                  <thead>
                    <tr className="border-b border-white/8">
                      {["ID", "Prompt", "Lang", "Config", "Score", "Status", "Duration"].map((h) => (
                        <th key={h} className="px-3 py-2.5 text-left text-white/30" style={{ fontWeight: 500, fontSize: 11 }}>
                          {h}
                        </th>
                      ))}
                    </tr>
                  </thead>
                  <tbody>
                    {metrics.recentEvals.map((e) => (
                      <tr key={e.id} className="border-b border-white/5 transition hover:bg-white/[0.02]">
                        <td className="px-3 py-3 text-emerald-400/70" style={{ ...mono, fontSize: 12 }}>{e.id}</td>
                        <td className="max-w-[250px] truncate px-3 py-3 text-white/70" title={e.prompt}>{e.prompt}</td>
                        <td className="px-3 py-3 text-white/50">{e.lang}</td>
                        <td className="px-3 py-3 text-white/50" style={{ ...mono, fontSize: 11 }}>{e.config}</td>
                        <td className="px-3 py-3" style={mono}>
                          <span className={e.score >= 80 ? "text-emerald-400" : e.score >= 60 ? "text-amber-400" : "text-red-400"}>
                            {e.score}
                          </span>
                        </td>
                        <td className="px-3 py-3">
                          {e.pass ? (
                            <CheckCircle2 className="h-4 w-4 text-emerald-400" />
                          ) : (
                            <XCircle className="h-4 w-4 text-red-400" />
                          )}
                        </td>
                        <td className="px-3 py-3 text-white/40" style={{ ...mono, fontSize: 12 }}>{e.duration}</td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            ) : (
              <div className="py-8 text-center text-white/30" style={{ fontSize: 13 }}>
                No recent evaluations
              </div>
            )}
          </div>
        </div>

        {/* Note about AI insights */}
        <div className="rounded-xl border border-amber-500/20 bg-amber-500/[0.05] p-5">
          <div className="flex items-start gap-3">
            <TrendingUp className="mt-0.5 h-4 w-4 shrink-0 text-amber-400" />
            <div className="space-y-1">
              <h3 className="text-amber-400" style={{ fontSize: 14, fontWeight: 500 }}>AI-Generated Insights</h3>
              <p className="text-white/50" style={{ fontSize: 13, lineHeight: 1.6 }}>
                This section will display AI-generated insights from the evaluation pipeline's summary/insights output.
                The grading pipeline generates analysis of trends, common failure patterns, and recommendations.
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
