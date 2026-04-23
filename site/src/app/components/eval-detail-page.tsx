import { useParams, Link } from "react-router";
import { useState, useEffect } from "react";
import { fetchRun } from "../data/api";
import type { RunSummary, EvalResult, EvalReport, SessionEvent, ReviewPanelEntry, GraderResult, ReviewedFile } from "../data/types";
import {
  ArrowLeft, CheckCircle2, XCircle, Clock, FileCode2, Cpu,
  MessageSquare, Wrench, Terminal, ChevronDown, ChevronRight,
  AlertTriangle, Zap, Copy, Check, Loader2
} from "lucide-react";
import { GraderResultRow } from "./GraderResultRow";
import { evalPassFromPoints, evalGraderTotals, evalPointTotals, graderPasses } from "../lib/evalPass";

const mono = { fontFamily: "'JetBrains Mono', monospace" };

function CopyButton({ text }: { text: string }) {
  const [copied, setCopied] = useState(false);
  return (
    <button
      onClick={() => { navigator.clipboard.writeText(text); setCopied(true); setTimeout(() => setCopied(false), 2000); }}
      className="text-white/30 transition hover:text-white/60"
    >
      {copied ? <Check className="h-3.5 w-3.5 text-emerald-400" /> : <Copy className="h-3.5 w-3.5" />}
    </button>
  );
}

function Switch({ checked, onChange, label }: { checked: boolean; onChange: () => void; label: string }) {
  const id = `switch-${label.replace(/\s+/g, "-").toLowerCase()}`;
  return (
    <div className="flex items-center gap-2 select-none">
      <button
        type="button"
        id={id}
        role="switch"
        aria-checked={checked}
        aria-label={label}
        onClick={onChange}
        onKeyDown={(e) => {
          if (e.key === " " || e.key === "Enter") {
            e.preventDefault();
            onChange();
          }
        }}
        className={`relative inline-flex h-5 w-9 items-center rounded-full transition-colors focus:outline-none focus:ring-2 focus:ring-emerald-400/60 ${
          checked ? "bg-emerald-500/30" : "bg-white/10"
        }`}
      >
        <span
          aria-hidden="true"
          className={`inline-block h-3.5 w-3.5 transform rounded-full transition-transform ${
            checked ? "translate-x-5 bg-emerald-400" : "translate-x-0.5 bg-white/40"
          }`}
        />
      </button>
      <label htmlFor={id} className="text-white/25 cursor-pointer" style={{ fontSize: 11 }}>
        {label}
      </label>
    </div>
  );
}

// ── Timeline types & helpers ────────────────────────────────────

interface TimelineStep {
  index: number;
  type: "prompt" | "reasoning" | "tool_call" | "message" | "turn_divider" | "system";
  icon: string;
  title: string;
  content?: string;
  detail?: string;
  toolName?: string;
  mcpServer?: string;
  duration?: number;
  success?: boolean;
  error?: string;
  turnNumber?: number;
  inputTokens?: number;
  outputTokens?: number;
}

function formatToolArgs(argsStr: string): string {
  try {
    const args = JSON.parse(argsStr);
    if (args.command) return args.command;
    if (args.path || args.file_path) return args.path || args.file_path;
    return JSON.stringify(args, null, 2);
  } catch {
    return argsStr;
  }
}

function getToolTitle(toolName: string, argsStr?: string, mcpServer?: string, mcpTool?: string): string {
  if (mcpServer && mcpTool) return `${mcpTool} via ${mcpServer}`;
  let shortArg = "";
  if (argsStr) {
    try {
      const args = JSON.parse(argsStr);
      if (args.command) {
        const first = args.command.split("\n")[0];
        shortArg = first.length > 80 ? first.slice(0, 77) + "…" : first;
      } else if (args.path || args.file_path) {
        shortArg = args.path || args.file_path;
      } else if (args.pattern) {
        shortArg = args.pattern;
      }
    } catch { /* ignore */ }
  }
  const titles: Record<string, string> = {
    bash: "Run command", write_file: "Create file", create: "Create file",
    read_file: "Read file", view: "Read file", edit: "Edit file",
    grep: "Search", glob: "Find files",
  };
  const base = titles[toolName] || toolName;
  return shortArg ? `${base}: ${shortArg}` : base;
}

function buildTimeline(events: SessionEvent[]): TimelineStep[] {
  const steps: TimelineStep[] = [];
  let stepIndex = 0;

  const usageByTurn = new Map<number, { inputTokens: number; outputTokens: number }>();
  for (const evt of events) {
    if (evt.type === "assistant.usage" && evt.turnNumber != null) {
      usageByTurn.set(evt.turnNumber, {
        inputTokens: evt.inputTokens ?? evt.input_tokens ?? 0,
        outputTokens: evt.outputTokens ?? evt.output_tokens ?? 0,
      });
    }
  }

  // Pre-pair tool starts with their completions
  const completionOf = new Map<number, number>();
  const consumed = new Set<number>();
  for (let i = 0; i < events.length; i++) {
    if (events[i].type !== "tool.execution_start") continue;
    for (let j = i + 1; j < events.length; j++) {
      if (events[j].type === "tool.execution_complete" && !consumed.has(j)) {
        completionOf.set(i, j);
        consumed.add(j);
        break;
      }
      if (events[j].type === "tool.execution_start") break;
    }
  }

  for (let i = 0; i < events.length; i++) {
    const evt = events[i];
    if (consumed.has(i)) continue;
    if (evt.type === "tool.execution_partial_result" || evt.type === "assistant.turn_end" || evt.type === "assistant.usage") continue;

    if (evt.type === "user.message") {
      steps.push({ index: stepIndex++, type: "prompt", icon: "📝", title: "Prompt", content: evt.content });
    } else if (evt.type === "assistant.reasoning") {
      steps.push({ index: stepIndex++, type: "reasoning", icon: "🤔", title: "Thinking", content: evt.content });
    } else if (evt.type === "tool.execution_start") {
      const tn = evt.tool_name || "unknown";
      const cIdx = completionOf.get(i);
      const c = cIdx != null ? events[cIdx] : undefined;
      steps.push({
        index: stepIndex++, type: "tool_call", icon: "🔧",
        title: getToolTitle(tn, evt.tool_args, evt.mcp_server_name, evt.mcp_tool_name),
        toolName: tn, mcpServer: evt.mcp_server_name, detail: evt.tool_args,
        content: c?.tool_result, duration: c?.duration_ms, success: c?.tool_success, error: c?.error,
      });
    } else if (evt.type === "tool.execution_complete") {
      steps.push({
        index: stepIndex++, type: "tool_call", icon: "🔧",
        title: evt.tool_name || "Tool result", toolName: evt.tool_name,
        content: evt.tool_result, duration: evt.duration_ms, success: evt.tool_success, error: evt.error,
      });
    } else if (evt.type === "assistant.message") {
      steps.push({ index: stepIndex++, type: "message", icon: "💬", title: "Response", content: evt.content });
    } else if (evt.type === "assistant.turn_start") {
      const tn = evt.turnNumber;
      const usage = tn != null ? usageByTurn.get(tn) : undefined;
      steps.push({
        index: stepIndex++, type: "turn_divider", icon: "", title: `Turn ${tn ?? "?"}`,
        turnNumber: tn, inputTokens: usage?.inputTokens, outputTokens: usage?.outputTokens,
      });
    } else if (evt.type === "session.start") {
      steps.push({ index: stepIndex++, type: "system", icon: "⏵", title: "Session started" });
    } else {
      steps.push({ index: stepIndex++, type: "system", icon: "⚙", title: evt.type, content: evt.content });
    }
  }
  return steps;
}

const stepBorderColor: Record<string, string> = {
  prompt: "border-blue-500/20 bg-blue-500/[0.05]",
  reasoning: "border-purple-500/20 bg-purple-500/[0.05]",
  tool_call: "border-amber-500/20 bg-amber-500/[0.05]",
  tool_call_fail: "border-red-500/20 bg-red-500/[0.05]",
  message: "border-cyan-500/20 bg-cyan-500/[0.05]",
  system: "border-white/5 bg-white/[0.02]",
};
const stepTextColor: Record<string, string> = {
  prompt: "text-blue-400",
  reasoning: "text-purple-400",
  tool_call: "text-amber-400",
  tool_call_fail: "text-red-400",
  message: "text-cyan-400",
  system: "text-white/40",
};

function TimelineCard({ step, defaultExpanded = false }: { step: TimelineStep; defaultExpanded?: boolean }) {
  const [expanded, setExpanded] = useState(defaultExpanded);
  const hasContent = !!(step.content || step.detail || step.error);
  const variant = step.type === "tool_call" && step.success === false ? "tool_call_fail" : step.type;

  return (
    <div className={`rounded-lg border ${stepBorderColor[variant] || stepBorderColor.system}`}>
      <div
        onClick={() => hasContent && setExpanded(!expanded)}
        className={`flex items-center gap-2.5 px-3 py-2.5 ${hasContent ? "cursor-pointer select-none" : ""}`}
      >
        <span style={{ fontSize: 14 }}>{step.icon}</span>
        <span className={`min-w-0 font-medium ${stepTextColor[variant] || stepTextColor.system}`} style={{ fontSize: 12 }}>
          {step.title.length > 100 ? step.title.slice(0, 97) + "…" : step.title}
        </span>
        <div className="flex flex-1 flex-wrap items-center gap-2">
          {step.mcpServer && (
            <span className="rounded bg-indigo-500/10 px-1.5 py-0.5 text-indigo-400/60" style={{ fontSize: 10 }}>
              via {step.mcpServer}
            </span>
          )}
          {step.duration != null && step.duration > 0 && (
            <span className="text-white/20" style={{ ...mono, fontSize: 10 }}>
              {step.duration >= 1000 ? `${(step.duration / 1000).toFixed(1)}s` : `${step.duration.toFixed(0)}ms`}
            </span>
          )}
          {step.success !== undefined && (
            step.success
              ? <CheckCircle2 className="h-3 w-3 text-emerald-400/70" />
              : <XCircle className="h-3 w-3 text-red-400/70" />
          )}
        </div>
        {hasContent && (
          expanded
            ? <ChevronDown className="h-3.5 w-3.5 shrink-0 text-white/20" />
            : <ChevronRight className="h-3.5 w-3.5 shrink-0 text-white/20" />
        )}
      </div>
      {expanded && (
        <div className="space-y-3 border-t border-white/5 p-3">
          {step.detail && (
            <div>
              <div className="mb-1 text-white/25" style={{ fontSize: 10 }}>Arguments</div>
              <pre className="overflow-auto whitespace-pre-wrap break-words rounded-md bg-black/30 p-3 text-white/50"
                style={{ ...mono, fontSize: 11, maxHeight: 400 }}>{formatToolArgs(step.detail)}</pre>
            </div>
          )}
          {step.content && (
            <div>
              <div className="mb-1 text-white/25" style={{ fontSize: 10 }}>
                {step.type === "tool_call" ? "Result" : step.type === "reasoning" ? "Reasoning" : "Content"}
              </div>
              <pre className="overflow-auto whitespace-pre-wrap break-words rounded-md bg-black/30 p-3 text-white/50"
                style={{ ...mono, fontSize: 11, maxHeight: 400 }}>{step.content}</pre>
            </div>
          )}
          {step.error && (
            <div className="flex items-start gap-1.5 text-red-400/80" style={{ fontSize: 11 }}>
              <XCircle className="mt-0.5 h-3 w-3 shrink-0" />
              <span>{step.error}</span>
            </div>
          )}
        </div>
      )}
    </div>
  );
}

function Timeline({ events }: { events: SessionEvent[] }) {
  const [showSystem, setShowSystem] = useState(false);
  const steps = buildTimeline(events);
  const visible = showSystem ? steps : steps.filter(s => s.type !== "system");
  const systemCount = steps.filter(s => s.type === "system").length;
  const firstPromptIdx = visible.findIndex(s => s.type === "prompt");

  return (
    <div className="space-y-1.5">
      {systemCount > 0 && (
        <div className="mb-3">
          <Switch
            checked={showSystem}
            onChange={() => setShowSystem(!showSystem)}
            label={`Show system events (${systemCount})`}
          />
        </div>
      )}
      {visible.map((step, idx) =>
        step.type === "turn_divider" ? (
          <div key={step.index} className="flex items-center gap-3 py-2">
            <div className="h-px flex-1 bg-white/10" />
            <span className="text-white/20" style={{ fontSize: 10 }}>
              {step.title}
              {step.inputTokens != null && ` • ${step.inputTokens.toLocaleString()} in / ${(step.outputTokens ?? 0).toLocaleString()} out`}
            </span>
            <div className="h-px flex-1 bg-white/10" />
          </div>
        ) : (
          <TimelineCard key={step.index} step={step} defaultExpanded={idx === firstPromptIdx} />
        )
      )}
    </div>
  );
}

function ReviewerTimeline({ reviewer }: { reviewer: ReviewPanelEntry }) {
  const [expanded, setExpanded] = useState(false);
  const events = (reviewer.events || []) as SessionEvent[];
  if (events.length === 0) return null;
  const pct = reviewer.max_score > 0 ? (reviewer.overall_score / reviewer.max_score) * 100 : 0;
  const color = pct >= 80 ? "text-emerald-400" : pct >= 60 ? "text-amber-400" : "text-red-400";

  return (
    <div className="rounded-lg border border-white/5 bg-white/[0.02]">
      <button onClick={() => setExpanded(!expanded)} className="flex w-full items-center justify-between p-4">
        <span className="text-white/60" style={{ fontSize: 13 }}>🔍 {reviewer.model} Review</span>
        <div className="flex items-center gap-3">
          <span className={color} style={{ ...mono, fontSize: 13 }}>{reviewer.overall_score}/{reviewer.max_score}</span>
          {expanded ? <ChevronDown className="h-3.5 w-3.5 text-white/20" /> : <ChevronRight className="h-3.5 w-3.5 text-white/20" />}
        </div>
      </button>
      {expanded && (
        <div className="border-t border-white/5 p-4">
          <Timeline events={events} />
        </div>
      )}
    </div>
  );
}

export function EvalDetailPage() {
  const { runId, promptId, "*": configName } = useParams();
  const [run, setRun] = useState<RunSummary | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [expandedReviewer, setExpandedReviewer] = useState<number>(0);
  const [showTimeline, setShowTimeline] = useState(true);
  const [showFiles, setShowFiles] = useState(false);
  const [expandedFile, setExpandedFile] = useState<string | null>(null);

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

  const decodedPromptId = decodeURIComponent(promptId || "");
  const decodedConfigName = configName || "";

  const evalResult = run?.results?.find(
    (r: EvalResult) => r.prompt_id === decodedPromptId && r.config_name === decodedConfigName
  );

  if (error || !evalResult) {
    return (
      <div className="flex min-h-screen items-center justify-center bg-[#0a0a0f]">
        <div className="text-center">
          <p className="mb-4 text-white/50">{error || "Evaluation not found"}</p>
          {runId && <Link to={`/runs/${runId}`} className="text-emerald-400">← Back to run</Link>}
        </div>
      </div>
    );
  }

  // Cast to EvalReport for detail-level access (EvalResult is summary-only in types)
  const r = evalResult as unknown as EvalReport;
  
  // Grader results (Phase 3) or legacy review (backward compat)
  const graderResults: GraderResult[] = r.grader_results || [];
  const hasGraders = graderResults.length > 0;

  const review = r.review || {};
  const overallScore = review.overall_score ?? 0;
  const maxScore = review.max_score ?? 100;

  // Phase 6.5 score-card metaphor: AND of every Point across every grader.
  // This is the dominant card; it reads true to the actual evaluation
  // ("did everything we checked pass?") instead of the legacy review-only
  // 12/12 number that ignored file/program/behavior graders.
  const evalPassed = evalPassFromPoints(r);
  const pointTotals = evalPointTotals(r);
  const graderTotals = evalGraderTotals(r);

  // Per-grader breakdown (used in the score card detail row + run-detail
  // tooltip). Each entry: name + points-passed + points-total + pass.
  const perGraderBreakdown = graderResults.map(g => {
    const pts = g.points ?? [];
    return {
      name: g.grader_name,
      pass: graderPasses(g),
      passed: pts.length > 0 ? pts.filter(p => p.pass).length : (graderPasses(g) ? 1 : 0),
      total: pts.length > 0 ? pts.length : 1,
    };
  });

  // Session events may come from the individual eval or from the summary result
  const sessionEvents: SessionEvent[] = r.session_events || [];
  const reviewPanel: ReviewPanelEntry[] = r.review_panel || [];
  const environment = r.environment;
  const generatedFiles = r.generated_files || [];
  const reviewedFiles = r.reviewed_files || [];
  const criteria = review.scores?.criteria || [];
  const rerunCommand = r.rerunCommand || "";
  const guardrailReason = r.guardrail_abort_reason || "";
  const errorMsg = r.error || guardrailReason || "";
  const workspaceDelta = r.workspace_delta;

  const envModel = environment?.model || "";
  const envInputTokens = environment?.totalInputTokens ?? environment?.total_input_tokens ?? 0;
  const envOutputTokens = environment?.totalOutputTokens ?? environment?.total_output_tokens ?? 0;
  const envTurnCount = environment?.turnCount ?? environment?.turn_count ?? 0;

  const durationTotal = r.duration_seconds || 0;
  const durationGen = r.generation_duration_seconds || 0;
  const durationReview = r.review_duration_seconds || 0;

  return (
    <div className="min-h-screen bg-[#0a0a0f] px-4 py-8 sm:px-6" style={{ fontFamily: "'Inter', sans-serif" }}>
      <div className="mx-auto max-w-6xl">
        <Link to={`/runs/${runId}`} className="mb-6 inline-flex items-center gap-1.5 text-white/40 no-underline transition hover:text-emerald-400" style={{ fontSize: 13 }}>
          <ArrowLeft className="h-3.5 w-3.5" /> Back to run
        </Link>

        {/* Header */}
        <div className="mb-6 flex flex-col gap-4 sm:flex-row sm:items-start sm:justify-between">
          <div>
            <div className="mb-2 flex items-center gap-3">
              {evalPassed ? <CheckCircle2 className="h-5 w-5 text-emerald-400" /> : <XCircle className="h-5 w-5 text-red-400" />}
              <h1 className="text-white" style={{ ...mono, fontSize: "clamp(1.1rem, 2.5vw, 1.5rem)" }}>
                {r.prompt_id}
              </h1>
            </div>
            <div className="flex flex-wrap items-center gap-2">
              <span className="rounded-md bg-blue-500/10 px-2.5 py-1 text-blue-400/80" style={{ ...mono, fontSize: 11 }}>{r.config_name}</span>
              {r.prompt_metadata?.service && <span className="rounded-md bg-white/5 px-2.5 py-1 text-white/50" style={{ fontSize: 11 }}>{r.prompt_metadata.service}</span>}
              {r.prompt_metadata?.language && <span className="rounded-md bg-white/5 px-2.5 py-1 text-white/50" style={{ fontSize: 11 }}>{r.prompt_metadata.language}</span>}
              {r.prompt_metadata?.difficulty && (
                <span className={`rounded-md px-2.5 py-1 ${
                  r.prompt_metadata.difficulty === "basic" ? "bg-emerald-500/10 text-emerald-400/70" :
                  r.prompt_metadata.difficulty === "intermediate" ? "bg-amber-500/10 text-amber-400/70" :
                  "bg-red-500/10 text-red-400/70"
                }`} style={{ fontSize: 11 }}>{r.prompt_metadata.difficulty}</span>
              )}
            </div>
          </div>
          {/* Phase 6.5: Points-metaphor score card. Replaces the
              review-only 12/12 number which ignored file/program/behavior
              graders. The dominant number is `N / N points`, qualified by
              `across M graders`. The hover tooltip lists per-grader
              breakdown so users can drill down without scrolling. */}
          {hasGraders ? (
            <div
              className={`rounded-xl border px-6 py-3 text-center ${evalPassed ? "border-emerald-500/20 bg-emerald-500/10" : "border-red-500/20 bg-red-500/10"}`}
              title={
                "Per grader:\n" +
                perGraderBreakdown
                  .map(b => `${b.pass ? "✓" : "✗"} ${b.name} — ${b.passed}/${b.total}`)
                  .join("\n")
              }
            >
              <div className="mb-1 text-white/40" style={{ fontSize: 10, letterSpacing: 0.5, textTransform: "uppercase" }}>
                Points
              </div>
              <div className={evalPassed ? "text-emerald-400" : "text-red-400"} style={{ ...mono, fontSize: 32 }}>
                {pointTotals.passed} / {pointTotals.total}
              </div>
              <div className="text-white/30" style={{ fontSize: 11 }}>
                across {graderTotals.total} grader{graderTotals.total === 1 ? "" : "s"}
              </div>
            </div>
          ) : (
            // No graders ran (e.g. error rows, v2 reports without grader
            // results). Fall back to the legacy review score so the card
            // still has a number.
            <div className={`rounded-xl border px-6 py-3 text-center ${evalPassed ? "border-emerald-500/20 bg-emerald-500/10" : "border-red-500/20 bg-red-500/10"}`}>
              <div className="mb-1 text-white/40" style={{ fontSize: 10, letterSpacing: 0.5, textTransform: "uppercase" }}>
                Review Score
              </div>
              <div className={evalPassed ? "text-emerald-400" : "text-red-400"} style={{ ...mono, fontSize: 32 }}>{overallScore}</div>
              <div className="text-white/30" style={{ fontSize: 11 }}>/ {maxScore}</div>
            </div>
          )}
        </div>

        {/* Error banner */}
        {!evalPassed && errorMsg && (
          <div className="mb-6 rounded-xl border border-red-500/20 bg-red-500/[0.05] p-4">
            <div className="mb-1 flex items-center gap-2">
              <AlertTriangle className="h-4 w-4 text-red-400" />
              <span className="text-red-400" style={{ fontSize: 13 }}>{errorMsg}</span>
            </div>
            {rerunCommand && (
              <div className="mt-2 ml-6 flex items-center gap-2">
                <code className="rounded bg-white/5 px-2 py-0.5 text-white/40" style={{ ...mono, fontSize: 11 }}>{rerunCommand}</code>
                <CopyButton text={rerunCommand} />
              </div>
            )}
          </div>
        )}

        {/* Stat cards */}
        <div className="mb-6 grid grid-cols-2 gap-3 md:grid-cols-6">
          {[
            { label: "Total", value: `${durationTotal.toFixed(1)}s`, icon: Clock, color: "text-white/60" },
            { label: "Generation", value: `${durationGen.toFixed(1)}s`, icon: Cpu, color: "text-blue-400" },
            { label: "Review", value: `${durationReview.toFixed(1)}s`, icon: MessageSquare, color: "text-purple-400" },
            { label: "Input Tokens", value: envInputTokens.toLocaleString(), icon: Zap, color: "text-emerald-400" },
            { label: "Output Tokens", value: envOutputTokens.toLocaleString(), icon: Zap, color: "text-pink-400" },
            { label: "Files", value: generatedFiles.length, icon: FileCode2, color: "text-amber-400" },
          ].map(s => (
            <div key={s.label} className="rounded-lg border border-white/8 bg-white/[0.03] p-3">
              <div className="mb-1 flex items-center gap-1.5">
                <s.icon className={`h-3 w-3 ${s.color}`} />
                <span className="text-white/30" style={{ fontSize: 10 }}>{s.label}</span>
              </div>
              <span className="text-white" style={{ ...mono, fontSize: 16 }}>{s.value}</span>
            </div>
          ))}
        </div>

        {/* Environment info */}
        {environment && (
          <div className="mb-6 rounded-xl border border-white/8 bg-white/[0.03] p-5">
            <h3 className="mb-3 text-white" style={{ fontSize: 14 }}>Environment</h3>
            <div className="grid gap-x-8 gap-y-3 sm:grid-cols-2 md:grid-cols-3" style={{ fontSize: 12 }}>
              {envModel && <div><span className="text-white/30">Model:</span> <span className="text-white/60" style={mono}>{envModel}</span></div>}
              {r.prompt_metadata?.sdk_package && (
                <div><span className="text-white/30">SDK Package:</span> <span className="text-blue-400/70" style={mono}>{r.prompt_metadata.sdk_package}</span></div>
              )}
              <div><span className="text-white/30">Turns:</span> <span className="text-white/60" style={mono}>{envTurnCount}</span></div>
            </div>

            {/* Tools Available — Phase 6.3: group skills by their plugin /
                skill_dir parent. environment.skill_groups (v3) carries the
                parent linkage; we join by name. Builtins and MCP servers
                stay flat below the grouped section. v2 reports without
                skill_groups fall through to the flat layout. */}
            {((r.tool_availability?.length ?? 0) > 0 ||
              ((environment.skills_loaded ?? environment.skillsLoaded)?.length ?? 0) > 0 ||
              (environment.available_tools?.length ?? 0) > 0) && (() => {
              // Normalize Env field names — Go emits a mix of camelCase (most
              // fields) and snake_case (skill_groups). Read both casings.
              const envSkillsLoaded = environment.skills_loaded ?? environment.skillsLoaded ?? [];
              const envSkillsInvoked = environment.skills_invoked ?? environment.skillsInvoked ?? [];
              const envMcpServers = environment.mcp_servers ?? environment.mcpServers ?? [];
              const envSkillGroups = environment.skill_groups ?? [];

              // Prefer top-level r.tool_availability (v3, has type/used). Fall
              // back to environment.available_tools (legacy flat list) or
              // synthesize from envSkillsLoaded/envMcpServers (v2).
              type ToolRow = { name: string; type?: string; used?: boolean };
              let toolRows: ToolRow[] = [];
              if ((r.tool_availability?.length ?? 0) > 0) {
                toolRows = (r.tool_availability ?? []).map(t => ({
                  name: t.name, type: t.type, used: t.used,
                }));
              } else if ((environment.available_tools?.length ?? 0) > 0) {
                toolRows = (environment.available_tools ?? []).map(n => ({ name: n }));
              } else {
                toolRows = [
                  ...envSkillsLoaded.map(n => ({ name: n, type: "skill" })),
                  ...envMcpServers.map(n => ({ name: n, type: "mcp" })),
                ];
              }

              const linkage = new Map(envSkillGroups.map(g => [g.name, g]));
              const invokedSet = new Set(envSkillsInvoked);
              const skillNamesSet = new Set(envSkillsLoaded);

              const classify = (row: ToolRow): "skill" | "mcp" | "builtin" => {
                if (row.type === "skill") return "skill";
                if (row.type === "mcp") return "mcp";
                if (row.type === "builtin") return "builtin";
                if (skillNamesSet.has(row.name)) return "skill";
                if (envMcpServers.some(s =>
                  row.name === s || row.name.startsWith(s + ".") || row.name.startsWith(s + "_") || row.name.startsWith(s + "/")
                )) return "mcp";
                return "builtin";
              };

              type Group = { parent: string; parentKind: string; children: string[] };
              const groups = new Map<string, Group>();
              const builtins: string[] = [];
              const mcps: string[] = [];

              for (const row of toolRows) {
                const cat = classify(row);
                if (cat === "skill") {
                  const link = linkage.get(row.name);
                  if (link?.parent) {
                    const key = `${link.parent_kind ?? "skill_dir"}::${link.parent}`;
                    if (!groups.has(key)) {
                      groups.set(key, { parent: link.parent, parentKind: link.parent_kind ?? "skill_dir", children: [] });
                    }
                    groups.get(key)!.children.push(row.name);
                  } else {
                    const key = "__ungrouped__";
                    if (!groups.has(key)) {
                      groups.set(key, { parent: "Ungrouped skills", parentKind: "", children: [] });
                    }
                    groups.get(key)!.children.push(row.name);
                  }
                } else if (cat === "mcp") {
                  mcps.push(row.name);
                } else {
                  builtins.push(row.name);
                }
              }

              const skillBadge = (name: string) =>
                invokedSet.has(name) ? "bg-blue-500/20 text-blue-400" : "bg-blue-500/10 text-blue-400/40";

              return (
                <div className="mt-4 pt-4 border-t border-white/5">
                  <div className="mb-2 text-white/40" style={{ fontSize: 11 }}>Available Tools</div>

                  {/* Grouped skills: one parent header line per plugin /
                      skill_dir, children indented and bordered. */}
                  {[...groups.values()].map((g, gi) => (
                    <div key={`g-${gi}`} className="mb-3">
                      <div className="mb-1 flex items-center gap-2 text-white/50" style={{ fontSize: 11 }}>
                        <span style={mono}>{g.parent}</span>
                        {g.parentKind && (
                          <span className="rounded bg-white/5 px-1.5 py-0.5 text-white/30" style={{ fontSize: 9 }}>
                            {g.parentKind}
                          </span>
                        )}
                        <span className="text-white/25" style={{ fontSize: 10 }}>
                          ({g.children.length} skill{g.children.length === 1 ? "" : "s"})
                        </span>
                      </div>
                      <div className="ml-3 flex flex-wrap gap-1.5 border-l border-white/5 pl-3">
                        {g.children.map(name => (
                          <span
                            key={name}
                            className={`inline-flex items-center gap-1.5 rounded-md px-2 py-1 ${skillBadge(name)}`}
                            style={{ fontSize: 10 }}
                          >
                            <span style={mono}>{name}</span>
                            <span className="opacity-60">(skill)</span>
                          </span>
                        ))}
                      </div>
                    </div>
                  ))}

                  {/* MCP servers: flat (no parent linkage on tool_availability today). */}
                  {mcps.length > 0 && (
                    <div className="mb-3">
                      <div className="mb-1 text-white/50" style={{ fontSize: 11 }}>MCP servers</div>
                      <div className="ml-3 flex flex-wrap gap-1.5 border-l border-white/5 pl-3">
                        {mcps.map(name => (
                          <span
                            key={name}
                            className="inline-flex items-center gap-1.5 rounded-md bg-purple-500/20 px-2 py-1 text-purple-400"
                            style={{ fontSize: 10 }}
                          >
                            <span style={mono}>{name}</span>
                            <span className="opacity-60">(MCP)</span>
                          </span>
                        ))}
                      </div>
                    </div>
                  )}

                  {/* Builtins: flat. */}
                  {builtins.length > 0 && (
                    <div>
                      <div className="mb-1 text-white/50" style={{ fontSize: 11 }}>Built-in tools</div>
                      <div className="ml-3 flex flex-wrap gap-1.5 border-l border-white/5 pl-3">
                        {builtins.map(name => (
                          <span
                            key={name}
                            className="inline-flex items-center gap-1.5 rounded-md bg-white/5 px-2 py-1 text-white/30"
                            style={{ fontSize: 10 }}
                          >
                            <span style={mono}>{name}</span>
                            <span className="opacity-60">(builtin)</span>
                          </span>
                        ))}
                      </div>
                    </div>
                  )}
                </div>
              );
            })()}
          </div>
        )}

        {/* Workspace Delta (#566) */}
        {workspaceDelta && (
          <div className="mb-6 rounded-xl border border-white/8 bg-white/[0.03] p-5">
            <h3 className="mb-3 text-white" style={{ fontSize: 14 }}>Workspace Changes</h3>
            <div className="grid gap-x-8 gap-y-2 sm:grid-cols-2 md:grid-cols-4" style={{ fontSize: 12 }}>
              <div><span className="text-white/30">Files created:</span> <span className="text-emerald-400/70" style={mono}>{workspaceDelta.new_file_count ?? 0}</span></div>
              <div><span className="text-white/30">Files modified:</span> <span className="text-amber-400/70" style={mono}>{workspaceDelta.modified_file_count ?? 0}</span></div>
              <div><span className="text-white/30">Files deleted:</span> <span className="text-red-400/70" style={mono}>{workspaceDelta.deleted_file_count ?? 0}</span></div>
              <div><span className="text-white/30">Net bytes:</span> <span className="text-white/60" style={mono}>{(workspaceDelta.bytes_net ?? 0).toLocaleString()} bytes</span></div>
            </div>
          </div>
        )}

        {/* Grader Results (Phase 3 unified grader pipeline) */}
        {hasGraders && (
          <div className="mb-6 rounded-xl border border-white/8 bg-white/[0.03] p-5">
            <h3 className="mb-4 flex items-center gap-2 text-white" style={{ fontSize: 14 }}>
              <Wrench className="h-4 w-4 text-white/40" /> Grader Results ({graderResults.length})
            </h3>
            <div className="space-y-2">
              {graderResults.map((gr, i) => (
                <GraderResultRow key={i} result={gr} />
              ))}
            </div>
          </div>
        )}

        {/* Legacy Review Panel (backward compat for pre-Phase-3 reports) */}
        {!hasGraders && (review.summary || criteria.length > 0 || reviewPanel.length > 0) && (
          <div className="mb-6 rounded-xl border border-amber-500/20 bg-amber-500/[0.05] p-5">
            <div className="mb-3 flex items-center gap-2">
              <AlertTriangle className="h-4 w-4 text-amber-400/80" />
              <h3 className="text-amber-400/80" style={{ fontSize: 14 }}>
                Legacy Review (Pre-Phase-3)
              </h3>
            </div>
            <p className="mb-4 text-amber-400/60" style={{ fontSize: 12 }}>
              This report uses the old review format. Re-run the evaluation to see grader results.
            </p>
            {/* Consolidated Review */}
            {review.summary && (
              <div className="mb-4 rounded-lg border border-white/8 bg-white/[0.03] p-4">
                <div className="mb-2 flex items-center justify-between">
                  <span className="text-white/50" style={{ fontSize: 13 }}>Consolidated Review</span>
                  <span className={scoreColor} style={{ ...mono, fontSize: 16 }}>{overallScore}/{maxScore}</span>
                </div>
                <p className="mb-3 text-white/45" style={{ fontSize: 13, lineHeight: 1.6 }}>{review.summary}</p>
                <div className="grid gap-4 sm:grid-cols-2">
                  {review.strengths && review.strengths.length > 0 && (
                    <div>
                      <p className="mb-1.5 text-emerald-400/60" style={{ fontSize: 11 }}>Strengths</p>
                      {review.strengths.map((s, i) => (
                        <div key={i} className="mb-1 flex gap-1.5">
                          <CheckCircle2 className="mt-0.5 h-3 w-3 shrink-0 text-emerald-400/50" />
                          <span className="text-white/50" style={{ fontSize: 12 }}>{s}</span>
                        </div>
                      ))}
                    </div>
                  )}
                  {review.issues && review.issues.length > 0 && (
                    <div>
                      <p className="mb-1.5 text-red-400/60" style={{ fontSize: 11 }}>Issues</p>
                      {review.issues.map((s, i) => (
                        <div key={i} className="mb-1 flex gap-1.5">
                          <XCircle className="mt-0.5 h-3 w-3 shrink-0 text-red-400/50" />
                          <span className="text-white/50" style={{ fontSize: 12 }}>{s}</span>
                        </div>
                      ))}
                    </div>
                  )}
                </div>
              </div>
            )}

            {/* Criteria badges */}
            {criteria.length > 0 && (
              <div className="mb-4 flex flex-wrap gap-1.5">
                {criteria.map(c => (
                  <span key={c.name} className={`flex items-center gap-1 rounded-md px-2 py-0.5 ${c.passed ? "bg-emerald-500/10 text-emerald-400/70" : "bg-red-500/10 text-red-400/70"}`} style={{ fontSize: 10 }}>
                    {c.passed ? <CheckCircle2 className="h-2.5 w-2.5" /> : <XCircle className="h-2.5 w-2.5" />}
                    {c.name}
                  </span>
                ))}
              </div>
            )}

            {/* Individual reviewers */}
            {reviewPanel.length > 0 && (
              <div className="space-y-2">
                {reviewPanel.map((rev, i) => {
                  const revScorePct = rev.max_score > 0 ? (rev.overall_score / rev.max_score) * 100 : 0;
                  const revCriteria = rev.scores?.criteria || [];
                  return (
                    <div key={i} className="rounded-lg border border-white/5 bg-white/[0.02]">
                      <button
                        onClick={() => setExpandedReviewer(expandedReviewer === i ? -1 : i)}
                        className="flex w-full items-center justify-between p-3 text-left"
                      >
                        <div className="flex items-center gap-3">
                          <span className="text-white/40" style={{ fontSize: 12 }}>Reviewer {i + 1}</span>
                          <span className="text-white/25" style={{ ...mono, fontSize: 11 }}>{rev.model}</span>
                        </div>
                        <div className="flex items-center gap-3">
                          <span className={revScorePct >= 80 ? "text-emerald-400" : revScorePct >= 60 ? "text-amber-400" : "text-red-400"} style={{ ...mono, fontSize: 13 }}>
                            {rev.overall_score}/{rev.max_score}
                          </span>
                          {expandedReviewer === i ? <ChevronDown className="h-3.5 w-3.5 text-white/20" /> : <ChevronRight className="h-3.5 w-3.5 text-white/20" />}
                        </div>
                      </button>
                      {expandedReviewer === i && (
                        <div className="border-t border-white/5 p-4">
                          {rev.summary && <p className="mb-3 text-white/40" style={{ fontSize: 12, lineHeight: 1.6 }}>{rev.summary}</p>}
                          {revCriteria.length > 0 && (
                            <div className="flex flex-wrap gap-1.5">
                              {revCriteria.map(c => (
                                <span key={c.name} className={`flex items-center gap-1 rounded-md px-2 py-0.5 ${c.passed ? "bg-emerald-500/10 text-emerald-400/70" : "bg-red-500/10 text-red-400/70"}`} style={{ fontSize: 10 }}>
                                  {c.passed ? <CheckCircle2 className="h-2.5 w-2.5" /> : <XCircle className="h-2.5 w-2.5" />}
                                  {c.name}
                                </span>
                              ))}
                            </div>
                          )}
                        </div>
                      )}
                    </div>
                  );
                })}
              </div>
            )}
          </div>
        )}

        {/* Session Timeline */}
        {sessionEvents.length > 0 && (
          <div className="mb-6 rounded-xl border border-white/8 bg-white/[0.03]">
            <button
              onClick={() => setShowTimeline(!showTimeline)}
              className="flex w-full items-center justify-between p-5"
            >
              <h3 className="flex items-center gap-2 text-white" style={{ fontSize: 14 }}>
                <Zap className="h-4 w-4 text-amber-400" />
                Session Timeline ({sessionEvents.length} events)
              </h3>
              {showTimeline ? <ChevronDown className="h-4 w-4 text-white/20" /> : <ChevronRight className="h-4 w-4 text-white/20" />}
            </button>
            {showTimeline && (
              <div className="border-t border-white/8 p-5">
                <Timeline events={sessionEvents} />
              </div>
            )}
          </div>
        )}

        {/* Reviewer Timelines */}
        {reviewPanel.some(rev => ((rev.events || []) as SessionEvent[]).length > 0) && (
          <div className="mb-6 rounded-xl border border-white/8 bg-white/[0.03] p-5">
            <h3 className="mb-4 flex items-center gap-2 text-white" style={{ fontSize: 14 }}>
              <MessageSquare className="h-4 w-4 text-purple-400" />
              Reviewer Timelines ({reviewPanel.length} reviewers)
            </h3>
            <div className="space-y-2">
              {reviewPanel.map((rev, i) => (
                <ReviewerTimeline key={i} reviewer={rev} />
              ))}
            </div>
          </div>
        )}

        {/* Generated Files */}
        {generatedFiles.length > 0 && (
          <div className="mb-6 rounded-xl border border-white/8 bg-white/[0.03]">
            <button onClick={() => setShowFiles(!showFiles)} className="flex w-full items-center justify-between p-5">
              <h3 className="flex items-center gap-2 text-white" style={{ fontSize: 14 }}>
                <FileCode2 className="h-4 w-4 text-white/40" />
                Generated Files ({generatedFiles.length})
              </h3>
              {showFiles ? <ChevronDown className="h-4 w-4 text-white/20" /> : <ChevronRight className="h-4 w-4 text-white/20" />}
            </button>
            {showFiles && (
              <div className="border-t border-white/8 p-5 space-y-2">
                {generatedFiles.map((filePath, i) => {
                  const reviewed = reviewedFiles.find(rf => rf.path === filePath);
                  const isExpanded = expandedFile === filePath;
                  
                  return (
                    <div key={i} className="rounded-lg border border-white/5 bg-black/20">
                      <button
                        onClick={() => setExpandedFile(isExpanded ? null : filePath)}
                        className="flex w-full items-center justify-between px-4 py-2 text-left"
                      >
                        <div className="flex items-center gap-2">
                          <FileCode2 className="h-3.5 w-3.5 text-emerald-400/50" />
                          <span className="text-emerald-400/70" style={{ ...mono, fontSize: 12 }}>{filePath}</span>
                        </div>
                        {reviewed && (
                          isExpanded
                            ? <ChevronDown className="h-3.5 w-3.5 text-white/30" />
                            : <ChevronRight className="h-3.5 w-3.5 text-white/30" />
                        )}
                      </button>
                      {reviewed && isExpanded && (
                        <div className="border-t border-white/5 p-4">
                          {(() => {
                            const MAX_CONTENT_BYTES = 256 * 1024; // 256 KiB
                            const tooBig = reviewed.content.length > MAX_CONTENT_BYTES;
                            const display = tooBig
                              ? reviewed.content.slice(0, MAX_CONTENT_BYTES)
                              : reviewed.content;
                            return (
                              <>
                                <pre className="overflow-x-auto rounded-lg bg-black/40 p-4 text-white/70" style={{ ...mono, fontSize: 11, lineHeight: 1.5 }}>
                                  <code>{display}</code>
                                </pre>
                                {tooBig && (
                                  <div className="mt-2 text-amber-400/70" style={{ fontSize: 11 }}>
                                    File truncated for display ({reviewed.content.length.toLocaleString()} bytes; showing first {MAX_CONTENT_BYTES.toLocaleString()}).
                                  </div>
                                )}
                              </>
                            );
                          })()}
                        </div>
                      )}
                    </div>
                  );
                })}
              </div>
            )}
          </div>
        )}
      </div>
    </div>
  );
}
