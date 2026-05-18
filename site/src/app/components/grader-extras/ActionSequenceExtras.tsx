import { CheckCircle2, XCircle, ArrowRight } from "lucide-react";
import type { ActionSequenceExtras as ActionSequenceExtrasType } from "../../data/types";

const mono = { fontFamily: "'JetBrains Mono', monospace" };

export function ActionSequenceExtras({ extras }: { extras: ActionSequenceExtrasType }) {
  // Build a unified sequence view showing expected vs actual at each step
  const maxLength = Math.max(
    extras.expected_sequence.length,
    extras.actual_sequence.length
  );

  return (
    <div>
      <div className="mb-1.5 text-white/25" style={{ fontSize: 10 }}>
        Action Sequence (Expected vs Actual)
      </div>
      <div className="space-y-2">
        <div className="text-white/40" style={{ fontSize: 11 }}>
          <span className="text-white/25">Matched: </span>
          {extras.matched_actions} / {extras.expected_sequence.length}
        </div>
        {/* Per-step diff table */}
        <div className="space-y-1">
          {Array.from({ length: maxLength }).map((_, i) => {
            const expected = extras.expected_sequence[i];
            const actual = extras.actual_sequence[i];
            const matches = expected === actual;
            return (
              <div
                key={i}
                className={`flex items-center gap-2 rounded-md border px-2 py-1.5 ${
                  matches
                    ? "border-emerald-500/15 bg-emerald-500/5"
                    : "border-red-500/20 bg-red-500/5"
                }`}
              >
                {matches ? (
                  <CheckCircle2 className="h-3 w-3 shrink-0 text-emerald-400/80" />
                ) : (
                  <XCircle className="h-3 w-3 shrink-0 text-red-400/80" />
                )}
                <div className="flex items-center gap-2 flex-1 min-w-0">
                  <span className="text-white/30" style={{ fontSize: 10 }}>
                    {i + 1}.
                  </span>
                  <span
                    className={matches ? "text-emerald-400/80" : "text-red-400/60"}
                    style={{ ...mono, fontSize: 11 }}
                  >
                    {expected || <span className="text-white/20">(none)</span>}
                  </span>
                  {!matches && actual && (
                    <>
                      <ArrowRight className="h-3 w-3 text-white/20" />
                      <span
                        className="text-amber-400/70"
                        style={{ ...mono, fontSize: 11 }}
                      >
                        {actual}
                      </span>
                    </>
                  )}
                  {!matches && !actual && expected && (
                    <span className="text-white/30" style={{ fontSize: 10 }}>
                      (missing)
                    </span>
                  )}
                </div>
              </div>
            );
          })}
        </div>
        {extras.tools_used && extras.tools_used.length > 0 && (
          <div className="text-white/40" style={{ fontSize: 11 }}>
            <span className="text-white/25">Tools used: </span>
            {extras.tools_used.join(", ")}
          </div>
        )}
      </div>
    </div>
  );
}
