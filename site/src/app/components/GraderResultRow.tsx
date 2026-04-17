import { useState } from "react";
import { CheckCircle2, XCircle, ChevronDown, ChevronRight, AlertCircle } from "lucide-react";
import type { GraderResult } from "../data/types";

const mono = { fontFamily: "'JetBrains Mono', monospace" };

export interface GraderResultRowProps {
  result: GraderResult;
  defaultExpanded?: boolean;
}

export function GraderResultRow({ result, defaultExpanded = false }: GraderResultRowProps) {
  const [expanded, setExpanded] = useState(defaultExpanded);

  // Determine pass/fail status
  const passed = result.pass !== null && result.pass !== undefined ? result.pass : null;
  const hasScore = result.score !== undefined && result.score !== null;
  const hasOverallScore = result.overall_score !== undefined && result.max_score !== undefined;

  // Formatted score display
  let scoreDisplay = "—";
  if (hasScore && result.score !== undefined) {
    scoreDisplay = `${(result.score * 100).toFixed(0)}%`;
  } else if (hasOverallScore && result.overall_score !== undefined && result.max_score !== undefined) {
    scoreDisplay = `${result.overall_score}/${result.max_score}`;
  }

  // Badge color logic
  const badgeColor = passed === true
    ? "bg-emerald-500/10 text-emerald-400 border-emerald-500/20"
    : passed === false
    ? "bg-red-500/10 text-red-400 border-red-500/20"
    : "bg-white/5 text-white/30 border-white/10";

  const badgeIcon = passed === true
    ? <CheckCircle2 className="h-3 w-3" />
    : passed === false
    ? <XCircle className="h-3 w-3" />
    : <AlertCircle className="h-3 w-3" />;

  // Check if there's expandable content
  const hasDetails = !!(
    result.summary ||
    result.issues?.length ||
    result.strengths?.length ||
    result.file_details ||
    result.program_details ||
    result.prompt_details ||
    result.behavior_details ||
    result.review_details
  );

  const graderTypeLabel = result.grader_type
    .split("_")
    .map((w) => w.charAt(0).toUpperCase() + w.slice(1))
    .join(" ");

  return (
    <div className="rounded-lg border border-white/5 bg-white/[0.02]">
      {/* Header Row */}
      <div
        onClick={() => hasDetails && setExpanded(!expanded)}
        className={`flex items-center gap-3 p-3 ${hasDetails ? "cursor-pointer select-none" : ""}`}
      >
        {/* Pass/Fail Badge */}
        <div className={`flex items-center gap-1.5 rounded-md border px-2 py-1 ${badgeColor}`} style={{ fontSize: 10 }}>
          {badgeIcon}
          <span>{passed === true ? "PASS" : passed === false ? "FAIL" : "N/A"}</span>
        </div>

        {/* Grader Name & Type */}
        <div className="flex-1 min-w-0">
          <div className="text-white/80 truncate" style={{ fontSize: 13 }}>
            {result.grader_name}
          </div>
          <div className="text-white/30 truncate" style={{ fontSize: 10 }}>
            {graderTypeLabel}
            {result.model && ` • ${result.model}`}
          </div>
        </div>

        {/* Score */}
        {(hasScore || hasOverallScore) && (
          <div className="text-white/50" style={{ ...mono, fontSize: 12 }}>
            {scoreDisplay}
          </div>
        )}

        {/* Gate indicator */}
        {result.gate && (
          <div className="rounded bg-amber-500/10 px-2 py-0.5 text-amber-400/80" style={{ fontSize: 9 }}>
            GATE
          </div>
        )}

        {/* Expand arrow */}
        {hasDetails && (
          expanded ? (
            <ChevronDown className="h-3.5 w-3.5 shrink-0 text-white/20" />
          ) : (
            <ChevronRight className="h-3.5 w-3.5 shrink-0 text-white/20" />
          )
        )}
      </div>

      {/* Expanded Details */}
      {expanded && hasDetails && (
        <div className="border-t border-white/5 p-3 space-y-3">
          {/* Summary */}
          {result.summary && (
            <div>
              <div className="mb-1 text-white/25" style={{ fontSize: 10 }}>Summary</div>
              <p className="text-white/60" style={{ fontSize: 12, lineHeight: 1.5 }}>
                {result.summary}
              </p>
            </div>
          )}

          {/* Issues & Strengths */}
          {(result.issues?.length || result.strengths?.length) && (
            <div className="grid gap-3 sm:grid-cols-2">
              {result.issues && result.issues.length > 0 && (
                <div>
                  <div className="mb-1.5 text-red-400/60" style={{ fontSize: 11 }}>Issues</div>
                  {result.issues.map((issue, i) => (
                    <div key={i} className="mb-1 flex gap-1.5">
                      <XCircle className="mt-0.5 h-3 w-3 shrink-0 text-red-400/50" />
                      <span className="text-white/50" style={{ fontSize: 11 }}>{issue}</span>
                    </div>
                  ))}
                </div>
              )}
              {result.strengths && result.strengths.length > 0 && (
                <div>
                  <div className="mb-1.5 text-emerald-400/60" style={{ fontSize: 11 }}>Strengths</div>
                  {result.strengths.map((strength, i) => (
                    <div key={i} className="mb-1 flex gap-1.5">
                      <CheckCircle2 className="mt-0.5 h-3 w-3 shrink-0 text-emerald-400/50" />
                      <span className="text-white/50" style={{ fontSize: 11 }}>{strength}</span>
                    </div>
                  ))}
                </div>
              )}
            </div>
          )}

          {/* File Details */}
          {result.file_details && (
            <div>
              <div className="mb-1.5 text-white/25" style={{ fontSize: 10 }}>File Checks</div>
              <div className="space-y-1">
                {result.file_details.checked_files.map((f, i) => (
                  <div key={i} className="flex items-center gap-2 text-white/40" style={{ fontSize: 11 }}>
                    {f.exists ? (
                      <CheckCircle2 className="h-3 w-3 text-emerald-400/50" />
                    ) : (
                      <XCircle className="h-3 w-3 text-red-400/50" />
                    )}
                    <span style={mono}>{f.path}</span>
                    {f.pattern && (
                      <span className="text-white/25">
                        • pattern {f.pattern_matched ? "matched" : "not matched"}
                      </span>
                    )}
                  </div>
                ))}
              </div>
            </div>
          )}

          {/* Program Details */}
          {result.program_details && (
            <div>
              <div className="mb-1.5 text-white/25" style={{ fontSize: 10 }}>Program Execution</div>
              <div className="rounded-md bg-black/30 p-3 space-y-2">
                <div>
                  <span className="text-white/25" style={{ fontSize: 10 }}>Command: </span>
                  <code className="text-white/60" style={{ ...mono, fontSize: 11 }}>
                    {result.program_details.command}
                  </code>
                </div>
                <div>
                  <span className="text-white/25" style={{ fontSize: 10 }}>Exit code: </span>
                  <span
                    className={result.program_details.exit_code === 0 ? "text-emerald-400" : "text-red-400"}
                    style={{ ...mono, fontSize: 11 }}
                  >
                    {result.program_details.exit_code}
                  </span>
                </div>
                {result.program_details.stdout && (
                  <div>
                    <div className="text-white/25 mb-1" style={{ fontSize: 10 }}>Stdout:</div>
                    <pre className="text-white/50 overflow-auto max-h-40" style={{ ...mono, fontSize: 10 }}>
                      {result.program_details.stdout}
                    </pre>
                  </div>
                )}
                {result.program_details.stderr && (
                  <div>
                    <div className="text-red-400/60 mb-1" style={{ fontSize: 10 }}>Stderr:</div>
                    <pre className="text-red-400/50 overflow-auto max-h-40" style={{ ...mono, fontSize: 10 }}>
                      {result.program_details.stderr}
                    </pre>
                  </div>
                )}
              </div>
            </div>
          )}

          {/* Prompt Details (LLM-as-judge) */}
          {result.prompt_details && (
            <div>
              <div className="mb-1.5 text-white/25" style={{ fontSize: 10 }}>LLM Review</div>
              <div className="space-y-2">
                {result.prompt_details.model && (
                  <div className="text-white/40" style={{ fontSize: 11 }}>
                    Model: <span style={mono}>{result.prompt_details.model}</span>
                  </div>
                )}
                {result.prompt_details.reasoning && (
                  <div>
                    <div className="text-white/25 mb-1" style={{ fontSize: 10 }}>Reasoning:</div>
                    <p className="text-white/50" style={{ fontSize: 11, lineHeight: 1.5 }}>
                      {result.prompt_details.reasoning}
                    </p>
                  </div>
                )}
              </div>
            </div>
          )}

          {/* Behavior Details */}
          {result.behavior_details && (
            <div>
              <div className="mb-1.5 text-white/25" style={{ fontSize: 10 }}>Behavior Analysis</div>
              <div className="grid gap-2 sm:grid-cols-2 text-white/40" style={{ fontSize: 11 }}>
                {result.behavior_details.tools_used && result.behavior_details.tools_used.length > 0 && (
                  <div>
                    <span className="text-white/25">Tools used: </span>
                    {result.behavior_details.tools_used.join(", ")}
                  </div>
                )}
                {result.behavior_details.turn_count !== undefined && (
                  <div>
                    <span className="text-white/25">Turns: </span>
                    {result.behavior_details.turn_count}
                  </div>
                )}
                {result.behavior_details.total_actions !== undefined && (
                  <div>
                    <span className="text-white/25">Total actions: </span>
                    {result.behavior_details.total_actions}
                  </div>
                )}
                {result.behavior_details.violations && result.behavior_details.violations.length > 0 && (
                  <div className="col-span-2">
                    <span className="text-red-400/60">Violations: </span>
                    {result.behavior_details.violations.join(", ")}
                  </div>
                )}
              </div>
            </div>
          )}

          {/* Review Details (multi-model review panel) */}
          {result.review_details && (
            <div>
              <div className="mb-1.5 text-white/25" style={{ fontSize: 10 }}>Review Panel</div>
              <div className="space-y-2">
                {result.review_details.criteria && result.review_details.criteria.length > 0 && (
                  <div className="flex flex-wrap gap-1.5">
                    {result.review_details.criteria.map((c, i) => (
                      <span
                        key={i}
                        className={`flex items-center gap-1 rounded-md px-2 py-0.5 ${
                          c.passed
                            ? "bg-emerald-500/10 text-emerald-400/70"
                            : "bg-red-500/10 text-red-400/70"
                        }`}
                        style={{ fontSize: 10 }}
                      >
                        {c.passed ? (
                          <CheckCircle2 className="h-2.5 w-2.5" />
                        ) : (
                          <XCircle className="h-2.5 w-2.5" />
                        )}
                        {c.name}
                      </span>
                    ))}
                  </div>
                )}
              </div>
            </div>
          )}
        </div>
      )}
    </div>
  );
}
