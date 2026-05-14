import { CheckCircle2, XCircle } from "lucide-react";
import type { ReviewExtras as ReviewExtrasType } from "../../data/types";

const mono = { fontFamily: "'JetBrains Mono', monospace" };

export function ReviewExtras({ extras }: { extras: ReviewExtrasType }) {
  return (
    <div>
      <div className="mb-1.5 text-white/25" style={{ fontSize: 10 }}>
        Review Panel
      </div>
      <div className="space-y-3">
        {extras.model && (
          <div className="text-white/40" style={{ fontSize: 11 }}>
            <span className="text-white/25">Model: </span>
            <span style={mono}>{extras.model}</span>
          </div>
        )}
        {extras.is_consensus && (
          <div className="inline-flex items-center rounded-md bg-blue-500/10 px-2 py-0.5 text-blue-400/80" style={{ fontSize: 9 }}>
            CONSENSUS
          </div>
        )}
        {extras.summary && (
          <div>
            <div className="text-white/25 mb-1" style={{ fontSize: 10 }}>
              Summary:
            </div>
            <p className="text-white/50" style={{ fontSize: 11, lineHeight: 1.5 }}>
              {extras.summary}
            </p>
          </div>
        )}
        {(extras.issues?.length || extras.strengths?.length) && (
          <div className="grid gap-3 sm:grid-cols-2">
            {extras.issues && extras.issues.length > 0 && (
              <div>
                <div className="mb-1.5 text-red-400/60" style={{ fontSize: 11 }}>
                  Issues
                </div>
                {extras.issues.map((issue, i) => (
                  <div key={i} className="mb-1 flex gap-1.5">
                    <XCircle className="mt-0.5 h-3 w-3 shrink-0 text-red-400/50" />
                    <span className="text-white/50" style={{ fontSize: 11 }}>
                      {issue}
                    </span>
                  </div>
                ))}
              </div>
            )}
            {extras.strengths && extras.strengths.length > 0 && (
              <div>
                <div className="mb-1.5 text-emerald-400/60" style={{ fontSize: 11 }}>
                  Strengths
                </div>
                {extras.strengths.map((strength, i) => (
                  <div key={i} className="mb-1 flex gap-1.5">
                    <CheckCircle2 className="mt-0.5 h-3 w-3 shrink-0 text-emerald-400/50" />
                    <span className="text-white/50" style={{ fontSize: 11 }}>
                      {strength}
                    </span>
                  </div>
                ))}
              </div>
            )}
          </div>
        )}
        {extras.panel_results && extras.panel_results.length > 0 && (
          <div>
            <div className="text-white/25 mb-2" style={{ fontSize: 10 }}>
              Panel Members ({extras.panel_results.length}):
            </div>
            <div className="space-y-2">
              {extras.panel_results.map((panel, i) => (
                <div
                  key={i}
                  className="rounded-md border border-white/5 bg-white/[0.01] p-2"
                >
                  <div className="mb-1 flex items-center justify-between gap-2">
                    <span className="text-white/50" style={{ ...mono, fontSize: 10 }}>
                      {panel.model}
                    </span>
                    <div className="flex items-center gap-2">
                      <span className={panel.pass ? "text-emerald-400/80" : "text-red-400/80"} style={{ fontSize: 10 }}>
                        {panel.pass ? "PASS" : "FAIL"}
                      </span>
                      <span className="text-white/30" style={{ fontSize: 10 }}>
                        {panel.score}
                      </span>
                    </div>
                  </div>
                  {(panel.issues?.length || panel.strengths?.length) && (
                    <div className="grid gap-2 sm:grid-cols-2">
                      {panel.issues && panel.issues.length > 0 && (
                        <div>
                          <div className="mb-1 text-red-400/60" style={{ fontSize: 10 }}>Issues</div>
                          {panel.issues.map((issue, idx) => (
                            <div key={idx} className="mb-1 flex gap-1.5">
                              <XCircle className="mt-0.5 h-3 w-3 shrink-0 text-red-400/50" />
                              <span className="text-white/40" style={{ fontSize: 10 }}>{issue}</span>
                            </div>
                          ))}
                        </div>
                      )}
                      {panel.strengths && panel.strengths.length > 0 && (
                        <div>
                          <div className="mb-1 text-emerald-400/60" style={{ fontSize: 10 }}>Strengths</div>
                          {panel.strengths.map((strength, idx) => (
                            <div key={idx} className="mb-1 flex gap-1.5">
                              <CheckCircle2 className="mt-0.5 h-3 w-3 shrink-0 text-emerald-400/50" />
                              <span className="text-white/40" style={{ fontSize: 10 }}>{strength}</span>
                            </div>
                          ))}
                        </div>
                      )}
                    </div>
                  )}
                </div>
              ))}
            </div>
          </div>
        )}
        {extras.duration_seconds !== undefined && (
          <div className="text-white/30" style={{ fontSize: 10 }}>
            Duration: {extras.duration_seconds.toFixed(2)}s
          </div>
        )}
      </div>
    </div>
  );
}
