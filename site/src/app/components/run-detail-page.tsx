import { useParams, Link, useNavigate } from "react-router";
import { useState, useEffect } from "react";
import { fetchRun } from "../data/api";
import type { RunSummary, EvalResult, EvalReport } from "../data/types";
import { CheckCircle2, XCircle, Clock, FileCode2, ArrowLeft, Loader2, Tag } from "lucide-react";
import { GraderResultRow } from "./GraderResultRow";

const mono = { fontFamily: "'JetBrains Mono', monospace" };

function ScoreBadge({ passed, total }: { passed: number; total: number }) {
  const allPassed = passed === total && total > 0;
  const color = allPassed ? "text-emerald-400" : "text-red-400";
  return <span className={color} style={{ ...mono, fontSize: 13 }}>{passed}/{total}</span>;
}

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

export function RunDetailPage() {
  const { runId } = useParams();
  const navigate = useNavigate();
  const [run, setRun] = useState<RunSummary | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [filterStatus, setFilterStatus] = useState<"all" | "pass" | "fail">("all");
  const [filterService, setFilterService] = useState<string>("all");
  const [filterLang, setFilterLang] = useState<string>("all");
  const [viewMode, setViewMode] = useState<"table" | "matrix">("table");

  useEffect(() => {
    if (!runId) return;
    fetchRun(runId)
      .then(setRun)
      .catch(e => setError(e.message))
      .finally(() => setLoading(false));
  }, [runId]);

  if (loading) {
    return (
      <div className="flex min-h-screen items-center justify-center bg-[#0a0a0f]">
        <Loader2 className="h-6 w-6 animate-spin text-emerald-400" />
      </div>
    );
  }

  if (error || !run) {
    return (
      <div className="flex min-h-screen items-center justify-center bg-[#0a0a0f]">
        <div className="text-center">
          <p className="mb-4 text-white/50">{error || `Run not found: ${runId}`}</p>
          <Link to="/runs" className="text-emerald-400">← Back to runs</Link>
        </div>
      </div>
    );
  }

  const passed = run.passed ?? 0;
  const total = run.total_evaluations ?? 0;
  const rate = total > 0 ? ((passed / total) * 100).toFixed(1) : "0.0";
  const results = run.results || [];

  const services = [...new Set(results.map(r => r.prompt_metadata?.service).filter(Boolean))];
  const langs = [...new Set(results.map(r => r.prompt_metadata?.language).filter(Boolean))];

  const filtered = results.filter((r: EvalResult) => {
    if (filterStatus === "pass" && !r.success) return false;
    if (filterStatus === "fail" && r.success) return false;
    if (filterService !== "all" && r.prompt_metadata?.service !== filterService) return false;
    if (filterLang !== "all" && r.prompt_metadata?.language !== filterLang) return false;
    return true;
  });

  const totalFiles = results.reduce((s, r) => s + (r.generated_files?.length || 0), 0);

  const promptIds = [...new Set(results.map(r => r.prompt_id))];
  const configs = [...new Set(results.map(r => r.config_name))];
  const matrixData = new Map<string, Map<string, EvalResult>>();
  
  for (const r of results) {
    if (!matrixData.has(r.prompt_id)) {
      matrixData.set(r.prompt_id, new Map());
    }
    matrixData.get(r.prompt_id)!.set(r.config_name, r);
  }

  return (
    <div className="min-h-screen bg-[#0a0a0f] px-4 py-8 sm:px-6" style={{ fontFamily: "'Inter', sans-serif" }}>
      <div className="mx-auto max-w-7xl">
        <Link to="/runs" className="mb-6 inline-flex items-center gap-1.5 text-white/40 no-underline transition hover:text-emerald-400" style={{ fontSize: 13 }}>
          <ArrowLeft className="h-3.5 w-3.5" /> All Runs
        </Link>

        <div className="mb-8 flex flex-col gap-4 sm:flex-row sm:items-end sm:justify-between">
          <div>
            <h1 className="mb-1 text-white" style={{ ...mono, fontSize: "clamp(1.25rem, 3vw, 1.75rem)" }}>
              {formatTimestamp(run.timestamp)}
            </h1>
            <p className="text-white/40" style={{ fontSize: 13 }}>
              Run ID: {run.run_id} · {run.total_evaluations ?? 0} evaluations · {formatDuration(run.duration_seconds)}
            </p>
          </div>
          <div className="flex items-center gap-2 rounded-lg border border-emerald-500/20 bg-emerald-500/10 px-4 py-2">
            <span className="text-emerald-400/60" style={{ fontSize: 12 }}>Pass Rate</span>
            <span className="text-emerald-400" style={{ ...mono, fontSize: 20 }}>{rate}%</span>
          </div>
        </div>

        <div className={`mb-8 grid gap-3 ${totalFiles > 0 ? "grid-cols-2 md:grid-cols-4" : "grid-cols-3"}`}>
          {[
            { label: "Passed", value: run.passed, icon: CheckCircle2, color: "text-emerald-400" },
            { label: "Failed", value: run.failed, icon: XCircle, color: "text-red-400" },
            { label: "Duration", value: formatDuration(run.duration_seconds), icon: Clock, color: "text-blue-400" },
            ...(totalFiles > 0 ? [{ label: "Files", value: totalFiles, icon: FileCode2, color: "text-amber-400" }] : []),
          ].map(s => (
            <div key={s.label} className="rounded-xl border border-white/8 bg-white/[0.03] p-4">
              <div className="mb-2 flex items-center gap-1.5">
                <s.icon className={`h-3.5 w-3.5 ${s.color}`} />
                <span className="text-white/35" style={{ fontSize: 11 }}>{s.label}</span>
              </div>
              <span className="text-white" style={{ ...mono, fontSize: 20 }}>{s.value}</span>
            </div>
          ))}
        </div>

        {run.analysis && (
          <div className="mb-8 rounded-xl border border-white/8 bg-white/[0.03] p-5">
            <div className="mb-2 text-white/40" style={{ fontSize: 11 }}>Run Summary</div>
            <p className="text-white/70" style={{ fontSize: 13, lineHeight: 1.6 }}>
              {run.analysis}
            </p>
          </div>
        )}

        <div className="mb-4 flex items-center justify-between">
          <div className="flex flex-wrap gap-2">
            <select
              value={filterStatus}
              onChange={e => setFilterStatus(e.target.value as "all" | "pass" | "fail")}
              className="rounded-lg border border-white/10 bg-white/5 px-3 py-1.5 text-white/70"
              style={{ fontSize: 12 }}
            >
              <option value="all">All Status</option>
              <option value="pass">Passed</option>
              <option value="fail">Failed</option>
            </select>
            <select
              value={filterService}
              onChange={e => setFilterService(e.target.value)}
              className="rounded-lg border border-white/10 bg-white/5 px-3 py-1.5 text-white/70"
              style={{ fontSize: 12 }}
            >
              <option value="all">All Services</option>
              {services.map(s => <option key={s} value={s}>{s}</option>)}
            </select>
            <select
              value={filterLang}
              onChange={e => setFilterLang(e.target.value)}
              className="rounded-lg border border-white/10 bg-white/5 px-3 py-1.5 text-white/70"
              style={{ fontSize: 12 }}
            >
              <option value="all">All Languages</option>
              {langs.map(l => <option key={l} value={l}>{l}</option>)}
            </select>
            <span className="self-center text-white/30" style={{ fontSize: 12 }}>{filtered.length} results</span>
          </div>

          <div className="flex gap-1 rounded-lg border border-white/10 bg-white/5 p-1">
            <button
              onClick={() => setViewMode("table")}
              className={`rounded px-3 py-1 transition ${viewMode === "table" ? "bg-emerald-500/20 text-emerald-400" : "text-white/40 hover:text-white/70"}`}
              style={{ fontSize: 12 }}
            >
              Table
            </button>
            <button
              onClick={() => setViewMode("matrix")}
              className={`rounded px-3 py-1 transition ${viewMode === "matrix" ? "bg-emerald-500/20 text-emerald-400" : "text-white/40 hover:text-white/70"}`}
              style={{ fontSize: 12 }}
            >
              Matrix
            </button>
          </div>
        </div>

        {viewMode === "table" && (
          <div className="overflow-x-auto rounded-xl border border-white/8 bg-white/[0.03]">
            <table className="w-full" style={{ fontSize: 13 }}>
              <thead>
                <tr className="border-b border-white/8">
                  {["Score", "Prompt", "Model", "Tools", "Service", "Lang", "Difficulty", "Duration", ""].map(h => (
                    <th key={h} className="px-4 py-3 text-left text-white/30" style={{ fontWeight: 500, fontSize: 11 }}>{h}</th>
                  ))}
                </tr>
              </thead>
              <tbody>
                {filtered.map((r, i) => {
                  const gradersPassed = (r as EvalReport).grader_results?.filter(g => g.pass === true).length ?? 0;
                  const gradersTotal = (r as EvalReport).grader_results?.length ?? 0;
                  
                  const evalReport = r as EvalReport;
                  const model = evalReport.config_used?.model || evalReport.environment?.model || r.config_name;
                  const tools = evalReport.environment?.mcp_servers || [];
                  const skills = evalReport.environment?.skills_loaded || [];
                  
                  return (
                    <tr 
                      key={`${r.prompt_id}-${r.config_name}-${i}`} 
                      onClick={() => navigate(`/runs/${run.run_id}/eval/${encodeURIComponent(r.prompt_id)}/${r.config_name}`)}
                      className="cursor-pointer border-b border-white/5 transition hover:bg-white/[0.02]"
                    >
                      <td className="px-4 py-3">
                        <ScoreBadge passed={gradersPassed} total={gradersTotal} />
                      </td>
                      <td className="px-4 py-3">
                        <span className="text-emerald-400/80" style={{ ...mono, fontSize: 12 }}>
                          {r.prompt_id}
                        </span>
                      </td>
                      <td className="px-4 py-3">
                        <span className="rounded-md bg-blue-500/10 px-2 py-0.5 text-blue-400/80" style={{ fontSize: 11 }}>
                          {model}
                        </span>
                      </td>
                      <td className="px-4 py-3">
                        <div className="flex flex-wrap gap-1">
                          {tools.slice(0, 2).map((t, idx) => (
                            <span key={idx} className="flex items-center gap-1 rounded-md bg-purple-500/10 px-2 py-0.5 text-purple-400/80" style={{ fontSize: 10 }}>
                              <Tag className="h-2.5 w-2.5" />
                              {t}
                            </span>
                          ))}
                          {skills.slice(0, 1).map((s, idx) => (
                            <span key={idx} className="flex items-center gap-1 rounded-md bg-amber-500/10 px-2 py-0.5 text-amber-400/80" style={{ fontSize: 10 }}>
                              {s}
                            </span>
                          ))}
                          {(tools.length + skills.length > 3) && (
                            <span className="text-white/30" style={{ fontSize: 10 }}>+{tools.length + skills.length - 3}</span>
                          )}
                        </div>
                      </td>
                      <td className="px-4 py-3">
                        <span className="rounded-md bg-white/5 px-2 py-0.5 text-white/50" style={{ fontSize: 11 }}>{r.prompt_metadata?.service}</span>
                      </td>
                      <td className="px-4 py-3 text-white/50" style={{ fontSize: 12 }}>{r.prompt_metadata?.language}</td>
                      <td className="px-4 py-3">
                        <span className={`rounded-md px-2 py-0.5 ${
                          r.prompt_metadata?.difficulty === "basic" ? "bg-emerald-500/10 text-emerald-400/70" :
                          r.prompt_metadata?.difficulty === "intermediate" ? "bg-amber-500/10 text-amber-400/70" :
                          "bg-red-500/10 text-red-400/70"
                        }`} style={{ fontSize: 11 }}>
                          {r.prompt_metadata?.difficulty}
                        </span>
                      </td>
                      <td className="px-4 py-3 text-white/40" style={{ ...mono, fontSize: 12 }}>
                        {r.duration_seconds != null && !isNaN(r.duration_seconds) ? `${r.duration_seconds.toFixed(1)}s` : "N/A"}
                      </td>
                      <td className="px-4 py-3 text-right">
                        <span className="text-white/30 text-xs">→</span>
                      </td>
                    </tr>
                  );
                })}
              </tbody>
            </table>
          </div>
        )}

        {viewMode === "matrix" && (
          <div className="space-y-8">
            {promptIds.map(promptId => {
              const promptResults = matrixData.get(promptId)!;
              const firstResult = Array.from(promptResults.values())[0];
              
              return (
                <div key={promptId} className="rounded-xl border border-white/8 bg-white/[0.03] p-5">
                  <div className="mb-4 border-b border-white/5 pb-3">
                    <div className="mb-1 text-emerald-400/90" style={{ ...mono, fontSize: 14 }}>
                      {promptId}
                    </div>
                    <div className="flex flex-wrap gap-2 text-white/40" style={{ fontSize: 11 }}>
                      <span>{firstResult.prompt_metadata?.service}</span>
                      <span>·</span>
                      <span>{firstResult.prompt_metadata?.language}</span>
                      <span>·</span>
                      <span>{firstResult.prompt_metadata?.plane}</span>
                      <span>·</span>
                      <span className={`${
                        firstResult.prompt_metadata?.difficulty === "basic" ? "text-emerald-400/70" :
                        firstResult.prompt_metadata?.difficulty === "intermediate" ? "text-amber-400/70" :
                        "text-red-400/70"
                      }`}>
                        {firstResult.prompt_metadata?.difficulty}
                      </span>
                    </div>
                  </div>

                  <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
                    {configs.map(configName => {
                      const result = promptResults.get(configName);
                      if (!result) {
                        return (
                          <div key={configName} className="rounded-lg border border-white/5 bg-white/[0.01] p-3">
                            <div className="mb-2 text-white/30" style={{ fontSize: 11 }}>{configName}</div>
                            <p className="text-white/20" style={{ fontSize: 12 }}>No data</p>
                          </div>
                        );
                      }

                      const evalReport = result as EvalReport;
                      const graders = evalReport.grader_results || [];
                      const model = evalReport.config_used?.model || configName;

                      return (
                        <div key={configName} className="rounded-lg border border-white/5 bg-white/[0.01]">
                          <div className="border-b border-white/5 p-3">
                            <div className="mb-1 flex items-center gap-2">
                              <span className="rounded-md bg-blue-500/10 px-2 py-0.5 text-blue-400/80" style={{ fontSize: 11 }}>
                                {model}
                              </span>
                              {result.success ? (
                                <CheckCircle2 className="h-3 w-3 text-emerald-400" />
                              ) : (
                                <XCircle className="h-3 w-3 text-red-400" />
                              )}
                            </div>
                            <div className="text-white/30" style={{ fontSize: 10 }}>{configName}</div>
                          </div>

                          <div className="space-y-2 p-3">
                            {graders.length > 0 ? (
                              graders.map((grader, idx) => (
                                <GraderResultRow key={idx} result={grader} />
                              ))
                            ) : (
                              <p className="text-white/30" style={{ fontSize: 11 }}>No grader results</p>
                            )}
                          </div>

                          <div className="border-t border-white/5 p-3">
                            <Link
                              to={`/runs/${run.run_id}/eval/${encodeURIComponent(result.prompt_id)}/${result.config_name}`}
                              className="text-emerald-400/80 no-underline transition hover:text-emerald-400"
                              style={{ fontSize: 12 }}
                            >
                              View full detail →
                            </Link>
                          </div>
                        </div>
                      );
                    })}
                  </div>
                </div>
              );
            })}
          </div>
        )}
      </div>
    </div>
  );
}
