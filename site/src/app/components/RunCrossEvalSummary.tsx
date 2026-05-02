import { useMemo, useState } from "react";
import type { RunSummary, EvalResult, GraderResult, GraderPoint } from "../data/types";
import { evalPassFromPoints, evalPointTotals } from "../lib/evalPass";
import { CheckCircle2, XCircle, ChevronDown, ChevronRight } from "lucide-react";
import { BarChart, Bar, XAxis, YAxis, Cell, ResponsiveContainer, LabelList } from "recharts";

const mono = { fontFamily: "'JetBrains Mono', monospace" };

interface CrossEvalProps {
  run: RunSummary;
}

interface Check {
  graderName: string;
  label: string;
  fullKey: string;
}

interface MatrixRow {
  promptId: string;
  configName: string;
  result: EvalResult;
}

interface CheckResult {
  pass: boolean | null; // null = not applicable
}

function formatDuration(s: number | undefined | null): string {
  if (s == null || isNaN(s)) return "N/A";
  if (s < 60) return `${s.toFixed(1)}s`;
  const m = Math.floor(s / 60);
  const sec = (s % 60).toFixed(0);
  return `${m}m ${sec}s`;
}

export function RunCrossEvalSummary({ run }: CrossEvalProps) {
  const [expandedGraders, setExpandedGraders] = useState<Set<string>>(new Set());
  const [globalExpanded, setGlobalExpanded] = useState(false);

  const results = run.results || [];

  // Compute all aggregates from results[].grader_results[].points[]
  const summary = useMemo(() => {
    let totalChecks = 0;
    let passedChecks = 0;
    let totalDuration = 0;
    let totalEvalsPass = 0;
    let totalEvals = results.length;

    const configStats = new Map<string, {
      checksPass: number;
      checksTotal: number;
      evalsPass: number;
      evalsTotal: number;
      durations: number[];
      order: boolean[];
    }>();

    const graderTypeStats = new Map<string, { pass: number; fail: number }>();
    const checkFailures = new Map<string, { passed: number; total: number }>();

    for (const r of results) {
      const graderResults = r.grader_results || [];
      const evalPassed = evalPassFromPoints(r);
      if (evalPassed) totalEvalsPass++;
      
      totalDuration += r.duration_seconds || 0;

      // Per-config rollup
      if (!configStats.has(r.config_name)) {
        configStats.set(r.config_name, {
          checksPass: 0,
          checksTotal: 0,
          evalsPass: 0,
          evalsTotal: 0,
          durations: [],
          order: [],
        });
      }
      const cfg = configStats.get(r.config_name)!;
      cfg.evalsTotal++;
      if (evalPassed) cfg.evalsPass++;
      cfg.durations.push(r.duration_seconds || 0);
      cfg.order.push(evalPassed);

      for (const grader of graderResults) {
        const points = grader.points || [];
        if (points.length === 0) {
          // Synthesize single point from grader verdict
          totalChecks++;
          cfg.checksTotal++;
          if (grader.pass) {
            passedChecks++;
            cfg.checksPass++;
          }

          // Grader-type stats
          if (!graderTypeStats.has(grader.grader_type)) {
            graderTypeStats.set(grader.grader_type, { pass: 0, fail: 0 });
          }
          const gts = graderTypeStats.get(grader.grader_type)!;
          if (grader.pass) gts.pass++;
          else gts.fail++;
        } else {
          for (const point of points) {
            totalChecks++;
            cfg.checksTotal++;
            if (point.pass) {
              passedChecks++;
              cfg.checksPass++;
            }

            // Grader-type stats
            if (!graderTypeStats.has(grader.grader_type)) {
              graderTypeStats.set(grader.grader_type, { pass: 0, fail: 0 });
            }
            const gts = graderTypeStats.get(grader.grader_type)!;
            if (point.pass) gts.pass++;
            else gts.fail++;

            // Check failure tracking
            const checkKey = `${grader.grader_name} · ${point.label}`;
            if (!checkFailures.has(checkKey)) {
              checkFailures.set(checkKey, { passed: 0, total: 0 });
            }
            const cf = checkFailures.get(checkKey)!;
            cf.total++;
            if (point.pass) cf.passed++;
          }
        }
      }
    }

    // Find hardest check
    let hardestCheck = "";
    let hardestRate = 1.0;
    for (const [key, stats] of checkFailures.entries()) {
      const rate = stats.total > 0 ? stats.passed / stats.total : 1.0;
      if (rate < hardestRate) {
        hardestRate = rate;
        hardestCheck = `${key} (${stats.passed}/${stats.total})`;
      }
    }

    return {
      configs: new Set(results.map(r => r.config_name)).size,
      checksPass: passedChecks,
      checksTotal: totalChecks,
      passRate: totalChecks > 0 ? ((passedChecks / totalChecks) * 100).toFixed(1) : "0.0",
      hardestCheck: hardestCheck || "N/A",
      avgDuration: totalEvals > 0 ? totalDuration / totalEvals : 0,
      configStats,
      graderTypeStats,
      totalEvalsPass,
      totalEvals,
    };
  }, [results]);

  // Matrix data structure
  const matrixData = useMemo(() => {
    const prompts = [...new Set(results.map(r => r.prompt_id))];
    
    const data: {
      prompt: string;
      checks: Check[];
      rows: MatrixRow[];
      matrix: Map<string, Map<string, CheckResult>>;
      graderGroups: Map<string, Check[]>;
    }[] = [];

    for (const promptId of prompts) {
      const promptResults = results.filter(r => r.prompt_id === promptId);
      
      // Collect all unique checks for this prompt
      const checksSet = new Map<string, Check>();
      const graderGroups = new Map<string, Check[]>();
      
      for (const r of promptResults) {
        const graderResults = r.grader_results || [];
        for (const grader of graderResults) {
          const points = grader.points || [];
          if (points.length === 0) {
            // Synthesize single check from grader
            const key = `${grader.grader_name} · [overall]`;
            if (!checksSet.has(key)) {
              const check = {
                graderName: grader.grader_name,
                label: "[overall]",
                fullKey: key,
              };
              checksSet.set(key, check);
              if (!graderGroups.has(grader.grader_name)) {
                graderGroups.set(grader.grader_name, []);
              }
              graderGroups.get(grader.grader_name)!.push(check);
            }
          } else {
            for (const point of points) {
              const key = `${grader.grader_name} · ${point.label}`;
              if (!checksSet.has(key)) {
                const check = {
                  graderName: grader.grader_name,
                  label: point.label,
                  fullKey: key,
                };
                checksSet.set(key, check);
                if (!graderGroups.has(grader.grader_name)) {
                  graderGroups.set(grader.grader_name, []);
                }
                graderGroups.get(grader.grader_name)!.push(check);
              }
            }
          }
        }
      }

      const checks = Array.from(checksSet.values());

      // Build matrix
      const matrix = new Map<string, Map<string, CheckResult>>();
      const rows: MatrixRow[] = [];

      for (const r of promptResults) {
        const rowKey = `${r.prompt_id}/${r.config_name}`;
        matrix.set(rowKey, new Map());
        rows.push({ promptId: r.prompt_id, configName: r.config_name, result: r });

        const graderResults = r.grader_results || [];
        for (const grader of graderResults) {
          const points = grader.points || [];
          if (points.length === 0) {
            const key = `${grader.grader_name} · [overall]`;
            matrix.get(rowKey)!.set(key, { pass: grader.pass });
          } else {
            for (const point of points) {
              const key = `${grader.grader_name} · ${point.label}`;
              matrix.get(rowKey)!.set(key, { pass: point.pass });
            }
          }
        }

        // Fill in null for missing checks
        for (const check of checks) {
          if (!matrix.get(rowKey)!.has(check.fullKey)) {
            matrix.get(rowKey)!.set(check.fullKey, { pass: null });
          }
        }
      }

      data.push({ prompt: promptId, checks, rows, matrix, graderGroups });
    }

    return data;
  }, [results]);

  // Grader-type bar chart data
  const graderTypeChartData = useMemo(() => {
    const data = [];
    for (const [type, stats] of summary.graderTypeStats.entries()) {
      const total = stats.pass + stats.fail;
      const passPercent = total > 0 ? ((stats.pass / total) * 100).toFixed(0) : "0";
      data.push({
        type,
        pass: stats.pass,
        fail: stats.fail,
        passPercent: parseInt(passPercent),
        label: `${stats.pass} pass / ${stats.fail} fail (${passPercent}%)`,
      });
    }
    return data.sort((a, b) => b.passPercent - a.passPercent);
  }, [summary]);

  const toggleGrader = (graderName: string) => {
    setExpandedGraders(prev => {
      const next = new Set(prev);
      if (next.has(graderName)) {
        next.delete(graderName);
      } else {
        next.add(graderName);
      }
      return next;
    });
  };

  const toggleGlobalExpand = () => {
    if (globalExpanded) {
      setExpandedGraders(new Set());
      setGlobalExpanded(false);
    } else {
      const allGraders = new Set<string>();
      for (const { graderGroups } of matrixData) {
        for (const grader of graderGroups.keys()) {
          allGraders.add(grader);
        }
      }
      setExpandedGraders(allGraders);
      setGlobalExpanded(true);
    }
  };

  return (
    <div className="mb-8 space-y-6">
      <div className="border-b border-white/5 pb-2">
        <h2 className="text-white/60" style={{ fontSize: 14, fontWeight: 500 }}>
          Cross-Evaluation Summary
        </h2>
      </div>

      {/* 1. Top-line summary band */}
      <div className="grid gap-3 sm:grid-cols-2 lg:grid-cols-5">
        {[
          { label: "Configs", value: summary.configs },
          { label: "Checks", value: `${summary.checksPass} / ${summary.checksTotal}` },
          { label: "Pass Rate", value: `${summary.passRate}%` },
          { label: "Hardest Check", value: summary.hardestCheck },
          { label: "Avg Duration", value: formatDuration(summary.avgDuration) },
        ].map(stat => (
          <div key={stat.label} className="rounded-lg border border-white/8 bg-white/[0.03] p-3">
            <div className="mb-1 text-white/35" style={{ fontSize: 11 }}>{stat.label}</div>
            <div className="text-white" style={{ ...mono, fontSize: 13 }}>
              {stat.value}
            </div>
          </div>
        ))}
      </div>

      {/* 2. Per-config rollup strip */}
      <div className="space-y-3">
        <h3 className="text-white/50" style={{ fontSize: 13, fontWeight: 500 }}>
          Per-Config Summary
        </h3>
        <div className="grid gap-3 sm:grid-cols-2 lg:grid-cols-3">
          {Array.from(summary.configStats.entries()).map(([configName, stats]) => {
            const checksPercent = stats.checksTotal > 0
              ? Math.round((stats.checksPass / stats.checksTotal) * 100)
              : 0;
            const avgDur = stats.durations.length > 0
              ? stats.durations.reduce((a, b) => a + b, 0) / stats.durations.length
              : 0;

            return (
              <div key={configName} className="rounded-lg border border-white/8 bg-white/[0.03] p-4">
                <div className="mb-3 border-b border-white/5 pb-2">
                  <div className="text-white" style={{ fontSize: 13, fontWeight: 500 }}>
                    {configName}
                  </div>
                </div>

                <div className="space-y-2">
                  <div>
                    <div className="mb-1 flex items-center justify-between">
                      <span className="text-white/40" style={{ fontSize: 11 }}>Checks</span>
                      <span className="text-white/70" style={{ ...mono, fontSize: 11 }}>
                        {stats.checksPass}/{stats.checksTotal}
                      </span>
                    </div>
                    <div className="h-1.5 overflow-hidden rounded-full bg-white/5">
                      <div
                        className="h-full bg-emerald-500"
                        style={{ width: `${checksPercent}%` }}
                      />
                    </div>
                    <div className="mt-0.5 text-right text-white/30" style={{ fontSize: 10 }}>
                      {checksPercent}%
                    </div>
                  </div>

                  <div className="flex items-center justify-between">
                    <span className="text-white/40" style={{ fontSize: 11 }}>Evals</span>
                    <div className="flex items-center gap-1">
                      <span className="text-white/70" style={{ ...mono, fontSize: 11 }}>
                        {stats.evalsPass}/{stats.evalsTotal}
                      </span>
                      <div className="flex gap-0.5">
                        {stats.order.map((pass, idx) => (
                          <div
                            key={idx}
                            className={`h-1.5 w-1.5 rounded-full ${
                              pass ? "bg-emerald-400" : "bg-red-400"
                            }`}
                            aria-label={pass ? "passed" : "failed"}
                          />
                        ))}
                      </div>
                    </div>
                  </div>

                  <div className="flex items-center justify-between">
                    <span className="text-white/40" style={{ fontSize: 11 }}>Avg Duration</span>
                    <span className="text-white/70" style={{ ...mono, fontSize: 11 }}>
                      {formatDuration(avgDur)}
                    </span>
                  </div>
                </div>
              </div>
            );
          })}
        </div>
      </div>

      {/* 3. Evals × Checks matrix */}
      {matrixData.length > 0 && (
        <div className="space-y-6">
          <div className="flex items-center justify-between">
            <h3 className="text-white/50" style={{ fontSize: 13, fontWeight: 500 }}>
              Evals × Checks Matrix
            </h3>
            <button
              onClick={toggleGlobalExpand}
              className="rounded-md border border-white/10 bg-white/5 px-3 py-1 text-white/60 transition hover:bg-white/10 hover:text-white/80"
              style={{ fontSize: 11 }}
            >
              {globalExpanded ? "Collapse All" : "Expand All"}
            </button>
          </div>

          {matrixData.map(({ prompt, checks, rows, matrix, graderGroups }) => (
            <div key={prompt} className="rounded-lg border border-white/8 bg-white/[0.03] p-4">
              {matrixData.length > 1 && (
                <div className="mb-3 border-b border-white/5 pb-2">
                  <div className="text-emerald-400/90" style={{ ...mono, fontSize: 12 }}>
                    {prompt}
                  </div>
                </div>
              )}

              <div className="overflow-x-auto">
                <table className="w-full border-collapse" style={{ fontSize: 11 }}>
                  <thead>
                    <tr className="border-b border-white/8">
                      <th className="sticky left-0 z-10 bg-white/[0.03] px-3 py-2 text-left text-white/40">
                        Config
                      </th>
                      {Array.from(graderGroups.entries()).map(([graderName, graderChecks]) => {
                        const isExpanded = expandedGraders.has(graderName);
                        const hasMultiple = graderChecks.length > 1;

                        if (!hasMultiple || isExpanded) {
                          return graderChecks.map(check => (
                            <th
                              key={check.fullKey}
                              className="px-2 py-2 text-center text-white/40"
                              style={{ fontSize: 10, minWidth: 80 }}
                            >
                              <div className="break-words">
                                {graderName}
                              </div>
                              <div className="mt-0.5 text-white/25">
                                {check.label}
                              </div>
                            </th>
                          ));
                        } else {
                          return (
                            <th
                              key={graderName}
                              className="cursor-pointer px-2 py-2 text-center text-white/40 transition hover:bg-white/5"
                              onClick={() => toggleGrader(graderName)}
                              style={{ fontSize: 10, minWidth: 80 }}
                            >
                              <div className="flex items-center justify-center gap-1">
                                <ChevronRight className="h-3 w-3" />
                                <span className="break-words">{graderName}</span>
                              </div>
                            </th>
                          );
                        }
                      })}
                    </tr>
                  </thead>
                  <tbody>
                    {rows.map(row => {
                      const rowKey = `${row.promptId}/${row.configName}`;
                      const rowData = matrix.get(rowKey)!;

                      return (
                        <tr key={rowKey} className="border-b border-white/5">
                          <td className="sticky left-0 z-10 bg-white/[0.03] px-3 py-2 text-white/70">
                            {row.configName}
                          </td>
                          {Array.from(graderGroups.entries()).map(([graderName, graderChecks]) => {
                            const isExpanded = expandedGraders.has(graderName);
                            const hasMultiple = graderChecks.length > 1;

                            if (!hasMultiple || isExpanded) {
                              return graderChecks.map(check => {
                                const result = rowData.get(check.fullKey);
                                return (
                                  <td key={check.fullKey} className="px-2 py-2 text-center">
                                    {result?.pass === true ? (
                                      <CheckCircle2 className="inline h-3.5 w-3.5 text-emerald-400" aria-label="passed" />
                                    ) : result?.pass === false ? (
                                      <XCircle className="inline h-3.5 w-3.5 text-red-400" aria-label="failed" />
                                    ) : (
                                      <span className="text-white/20" aria-label="not applicable">—</span>
                                    )}
                                  </td>
                                );
                              });
                            } else {
                              // Collapsed: aggregate across all checks in this grader
                              const allPass = graderChecks.every(check => {
                                const result = rowData.get(check.fullKey);
                                return result?.pass === true;
                              });
                              const anyFail = graderChecks.some(check => {
                                const result = rowData.get(check.fullKey);
                                return result?.pass === false;
                              });
                              const allNull = graderChecks.every(check => {
                                const result = rowData.get(check.fullKey);
                                return result?.pass === null;
                              });

                              return (
                                <td
                                  key={graderName}
                                  className="cursor-pointer px-2 py-2 text-center transition hover:bg-white/5"
                                  onClick={() => toggleGrader(graderName)}
                                >
                                  {allNull ? (
                                    <span className="text-white/20" aria-label="not applicable">—</span>
                                  ) : allPass ? (
                                    <CheckCircle2 className="inline h-3.5 w-3.5 text-emerald-400" aria-label="all passed" />
                                  ) : anyFail && !allPass ? (
                                    <div className="inline-flex items-center gap-0.5" aria-label="partial">
                                      <CheckCircle2 className="h-3 w-3 text-emerald-400/50" />
                                      <XCircle className="h-3 w-3 text-red-400/50" />
                                    </div>
                                  ) : (
                                    <XCircle className="inline h-3.5 w-3.5 text-red-400" aria-label="failed" />
                                  )}
                                </td>
                              );
                            }
                          })}
                        </tr>
                      );
                    })}
                    <tr className="border-t-2 border-white/10 font-semibold">
                      <td className="sticky left-0 z-10 bg-white/[0.03] px-3 py-2 text-white/50">
                        Total
                      </td>
                      {Array.from(graderGroups.entries()).map(([graderName, graderChecks]) => {
                        const isExpanded = expandedGraders.has(graderName);
                        const hasMultiple = graderChecks.length > 1;

                        if (!hasMultiple || isExpanded) {
                          return graderChecks.map(check => {
                            let passed = 0;
                            let total = 0;
                            for (const row of rows) {
                              const rowKey = `${row.promptId}/${row.configName}`;
                              const result = matrix.get(rowKey)!.get(check.fullKey);
                              if (result?.pass !== null) {
                                total++;
                                if (result?.pass) passed++;
                              }
                            }
                            return (
                              <td key={check.fullKey} className="px-2 py-2 text-center text-white/60">
                                <span style={{ ...mono, fontSize: 10 }}>{passed}/{total}</span>
                              </td>
                            );
                          });
                        } else {
                          return (
                            <td
                              key={graderName}
                              className="cursor-pointer px-2 py-2 text-center text-white/60 transition hover:bg-white/5"
                              onClick={() => toggleGrader(graderName)}
                            >
                              <ChevronRight className="inline h-3 w-3" />
                            </td>
                          );
                        }
                      })}
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          ))}
        </div>
      )}

      {/* 4. Per-grader-type stacked bars */}
      {graderTypeChartData.length > 0 && (
        <div className="space-y-3">
          <h3 className="text-white/50" style={{ fontSize: 13, fontWeight: 500 }}>
            Check Pass Rate by Grader Type
          </h3>
          <div className="rounded-lg border border-white/8 bg-white/[0.03] p-4">
            <ResponsiveContainer width="100%" height={graderTypeChartData.length * 40 + 20}>
              <BarChart
                data={graderTypeChartData}
                layout="vertical"
                margin={{ top: 5, right: 30, left: 120, bottom: 5 }}
              >
                <XAxis type="number" domain={[0, 100]} hide />
                <YAxis
                  type="category"
                  dataKey="type"
                  tick={{ fill: "rgba(255, 255, 255, 0.5)", fontSize: 11 }}
                  axisLine={false}
                  tickLine={false}
                />
                <Bar dataKey="passPercent" fill="#10b981" radius={[0, 4, 4, 0]}>
                  <LabelList
                    dataKey="label"
                    position="right"
                    fill="rgba(255, 255, 255, 0.6)"
                    style={{ fontSize: 10, fontFamily: mono.fontFamily }}
                  />
                </Bar>
              </BarChart>
            </ResponsiveContainer>
          </div>
        </div>
      )}

      <div className="border-t border-white/5 pt-4">
        <p className="text-white/30" style={{ fontSize: 11 }}>
          Aggregated from {results.length} evaluation{results.length === 1 ? "" : "s"} across {summary.configs} config{summary.configs === 1 ? "" : "s"}.
        </p>
      </div>
    </div>
  );
}
