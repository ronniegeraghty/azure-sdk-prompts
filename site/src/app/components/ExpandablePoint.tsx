import { useState } from "react";
import { CheckCircle2, XCircle, ChevronRight } from "lucide-react";
import type { GraderPoint } from "../data/types";

const mono = { fontFamily: "'JetBrains Mono', monospace" };

interface ReviewerVote {
  model: string;
  passed: boolean;
  reason?: string;
}

interface ExpandablePointProps {
  point: GraderPoint;
  labelText: string;
  secondary: string | null;
  reviewerVotes: ReviewerVote[];
}

export function ExpandablePoint({
  point: p,
  labelText,
  secondary,
  reviewerVotes,
}: ExpandablePointProps) {
  // Calculate disagreement
  const passedCount = reviewerVotes.filter(v => v.passed).length;
  const totalVotes = reviewerVotes.length;
  const hasDisagreement = totalVotes > 0 && passedCount > 0 && passedCount < totalVotes;
  
  // Point expansion state (auto-expand if disagreement)
  const [pointExpanded, setPointExpanded] = useState(hasDisagreement);
  
  return (
    <div>
      <div
        className={`flex items-start gap-2 py-1 ${reviewerVotes.length > 0 ? "cursor-pointer" : ""}`}
        onClick={() => reviewerVotes.length > 0 && setPointExpanded(!pointExpanded)}
        role={reviewerVotes.length > 0 ? "button" : undefined}
        aria-expanded={reviewerVotes.length > 0 ? pointExpanded : undefined}
        tabIndex={reviewerVotes.length > 0 ? 0 : undefined}
        onKeyDown={(e) => {
          if (reviewerVotes.length > 0 && (e.key === "Enter" || e.key === " ")) {
            e.preventDefault();
            setPointExpanded(!pointExpanded);
          }
        }}
      >
        {p.pass ? (
          <CheckCircle2 className="mt-0.5 h-3 w-3 shrink-0 text-emerald-400/80" />
        ) : (
          <XCircle className="mt-0.5 h-3 w-3 shrink-0 text-red-400/80" />
        )}
        <div className="min-w-0 flex-1">
          <div className="flex items-center gap-2">
            <span
              className={p.pass ? "text-white/70" : "text-red-400/80"}
              style={{ fontSize: 12 }}
            >
              {labelText}
            </span>
            {hasDisagreement && (
              <span
                className="inline-flex items-center gap-1 rounded bg-amber-500/10 px-1.5 py-0.5 text-amber-400/80"
                style={{ fontSize: 9 }}
              >
                ⚠️ {passedCount}/{totalVotes}
              </span>
            )}
          </div>
          {secondary && (
            <div
              className="text-white/40"
              style={{ fontSize: 11, lineHeight: 1.5 }}
            >
              {secondary}
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
        {reviewerVotes.length > 0 && (
          <ChevronRight
            className={`h-3 w-3 shrink-0 text-white/20 transition-transform ${
              pointExpanded ? "rotate-90" : ""
            }`}
          />
        )}
      </div>
      
      {/* Per-reviewer votes (expandable) */}
      {pointExpanded && reviewerVotes.length > 0 && (
        <div className="ml-8 mt-1 space-y-1 border-l-2 border-white/5 pl-3">
          {reviewerVotes.map((vote, vi) => (
            <div key={vi} className="flex items-start gap-2 py-0.5">
              {vote.passed ? (
                <CheckCircle2 className="mt-0.5 h-3 w-3 shrink-0 text-emerald-400/80" />
              ) : (
                <XCircle className="mt-0.5 h-3 w-3 shrink-0 text-red-400/80" />
              )}
              <span
                className="text-white/50"
                style={{ ...mono, fontSize: 10 }}
              >
                {vote.model}:
              </span>
              <span
                className="text-white/40 flex-1"
                style={{ fontSize: 10, lineHeight: 1.4 }}
              >
                {vote.reason || (vote.passed ? "Pass" : "Fail")}
              </span>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
