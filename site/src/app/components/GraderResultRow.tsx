import { useState } from "react";
import { CheckCircle2, XCircle, ChevronDown, ChevronRight } from "lucide-react";
import type { GraderResult } from "../data/types";
import { formatGraderScore } from "../lib/graderScore";
import {
  FileExtras,
  ProgramExtras,
  PromptExtras,
  BehaviorExtras,
  ActionSequenceExtras,
  ToolConstraintExtras,
  OutputCheckExtras,
  ReviewExtras,
} from "./grader-extras";

const mono = { fontFamily: "'JetBrains Mono', monospace" };

export interface GraderResultRowProps {
  result: GraderResult;
  defaultExpanded?: boolean;
}

export function GraderResultRow({ result, defaultExpanded = false }: GraderResultRowProps) {
  // v4 schema: Points is required, pass is always boolean
  const hasMultiplePoints = result.points.length > 1;
  const hasFailed = !result.pass;
  
  // Auto-expand multi-point graders OR failed graders by default
  const [expanded, setExpanded] = useState(
    defaultExpanded || hasMultiplePoints || hasFailed
  );

  // Badge color based on result.pass (v4: always boolean)
  const badgeColor = result.pass
    ? "bg-emerald-500/10 text-emerald-400 border-emerald-500/20"
    : "bg-red-500/10 text-red-400 border-red-500/20";

  const badgeIcon = result.pass ? (
    <CheckCircle2 className="h-3 w-3" />
  ) : (
    <XCircle className="h-3 w-3" />
  );

  // Check if there's expandable content (Points list + Extras)
  const hasDetails = result.points.length > 0 || !!result.extras;

  const graderTypeLabel = result.grader_type
    .split("_")
    .map((w) => w.charAt(0).toUpperCase() + w.slice(1))
    .join(" ");

  // Get model from extras if present (for review graders)
  const modelName = result.extras?.review?.model || result.extras?.prompt?.model;

  return (
    <div className="rounded-lg border border-white/5 bg-white/[0.02]">
      {/* Header Row — single source of truth for score, no duplication */}
      <div
        onClick={() => hasDetails && setExpanded(!expanded)}
        className={`flex items-center gap-3 p-3 ${
          hasDetails ? "cursor-pointer select-none" : ""
        }`}
      >
        {/* Grader Name & Type */}
        <div className="flex-1 min-w-0">
          <div className="text-white/80 truncate" style={{ fontSize: 13 }}>
            {result.grader_name}
          </div>
          <div className="text-white/30 truncate" style={{ fontSize: 10 }}>
            {graderTypeLabel}
            {modelName && ` • ${modelName}`}
          </div>
        </div>

        {/* Score (canonical format: "N/M points") */}
        <div className="text-white/50" style={{ ...mono, fontSize: 12 }}>
          {formatGraderScore(result)}
        </div>

        {/* Pass/Fail Icon Badge (icon-only, score string already shows count) */}
        <div
          className={`flex items-center justify-center rounded-md border w-6 h-6 ${badgeColor}`}
        >
          {badgeIcon}
        </div>

        {/* Gate indicator */}
        {result.gate && (
          <div
            className="rounded bg-amber-500/10 px-2 py-0.5 text-amber-400/80"
            style={{ fontSize: 9 }}
          >
            GATE
          </div>
        )}

        {/* Expand arrow */}
        {hasDetails &&
          (expanded ? (
            <ChevronDown className="h-3.5 w-3.5 shrink-0 text-white/20" />
          ) : (
            <ChevronRight className="h-3.5 w-3.5 shrink-0 text-white/20" />
          ))}
      </div>

      {/* Expanded Details — v4: Points list (always) + KindExtras (when present) */}
      {expanded && hasDetails && (
        <div className="border-t border-white/5 p-3 space-y-3">
          {/* Points List — ALWAYS rendered when there are points */}
          {result.points.length > 0 && (
            <div>
              <div className="mb-1.5 text-white/25" style={{ fontSize: 10 }}>
                Points
              </div>
              <div className="ml-6 space-y-1">
                {result.points.map((p, i) => (
                  <div
                    key={i}
                    className="flex items-start gap-2 py-1"
                  >
                    {p.pass ? (
                      <CheckCircle2 className="mt-0.5 h-3 w-3 shrink-0 text-emerald-400/80" />
                    ) : (
                      <XCircle className="mt-0.5 h-3 w-3 shrink-0 text-red-400/80" />
                    )}
                    <div className="min-w-0 flex-1">
                      <div
                        className={p.pass ? "text-white/70" : "text-red-400/80"}
                        style={{ fontSize: 12 }}
                      >
                        {p.label}
                      </div>
                      {p.message && (
                        <div
                          className="text-white/40"
                          style={{ fontSize: 11, lineHeight: 1.5 }}
                        >
                          {p.message}
                        </div>
                      )}
                      {p.evidence && Object.keys(p.evidence).length > 0 && (
                        <div className="mt-1 flex flex-wrap gap-1">
                          {Object.entries(p.evidence).map(([k, v]) => (
                            <span
                              key={k}
                              className="rounded bg-white/5 px-1.5 py-0.5 text-white/30"
                              style={{ fontSize: 9 }}
                            >
                              {k}: {v}
                            </span>
                          ))}
                        </div>
                      )}
                    </div>
                  </div>
                ))}
              </div>
            </div>
          )}

          {/* Message (headline summary) */}
          {result.message && (
            <div>
              <div className="mb-1 text-white/25" style={{ fontSize: 10 }}>
                Summary
              </div>
              <p className="text-white/60" style={{ fontSize: 12, lineHeight: 1.5 }}>
                {result.message}
              </p>
            </div>
          )}

          {/* Kind-specific Extras — single dispatcher */}
          {result.extras && (
            <div>
              {result.extras.file && <FileExtras extras={result.extras.file} />}
              {result.extras.program && <ProgramExtras extras={result.extras.program} />}
              {result.extras.prompt && <PromptExtras extras={result.extras.prompt} />}
              {result.extras.behavior && <BehaviorExtras extras={result.extras.behavior} />}
              {result.extras.action_sequence && (
                <ActionSequenceExtras extras={result.extras.action_sequence} />
              )}
              {result.extras.tool_constraint && (
                <ToolConstraintExtras extras={result.extras.tool_constraint} />
              )}
              {result.extras.output_check && (
                <OutputCheckExtras extras={result.extras.output_check} />
              )}
              {result.extras.review && <ReviewExtras extras={result.extras.review} />}
            </div>
          )}
        </div>
      )}
    </div>
  );
}
