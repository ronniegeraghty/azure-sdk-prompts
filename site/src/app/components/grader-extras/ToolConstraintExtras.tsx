import type { ToolConstraintExtras as ToolConstraintExtrasType } from "../../data/types";

const mono = { fontFamily: "'JetBrains Mono', monospace" };

export function ToolConstraintExtras({ extras }: { extras: ToolConstraintExtrasType }) {
  return (
    <div>
      <div className="mb-1.5 text-white/25" style={{ fontSize: 10 }}>
        Tool Constraints
      </div>
      <div className="space-y-2">
        <div
          className={`text-${
            extras.constraints_met ? "emerald" : "red"
          }-400/80`}
          style={{ fontSize: 11 }}
        >
          <span className="text-white/25">Status: </span>
          {extras.constraints_met ? "All constraints met" : "Constraints violated"}
        </div>
        {extras.tools_used && extras.tools_used.length > 0 && (
          <div className="text-white/40" style={{ fontSize: 11 }}>
            <span className="text-white/25">Tools used: </span>
            {extras.tools_used.join(", ")}
          </div>
        )}
        {Object.keys(extras.tool_counts).length > 0 && (
          <div>
            <div className="text-white/25 mb-1" style={{ fontSize: 10 }}>
              Tool call counts:
            </div>
            <div className="grid grid-cols-2 gap-x-4 gap-y-1">
              {Object.entries(extras.tool_counts).map(([tool, count]) => (
                <div key={tool} className="text-white/40" style={{ ...mono, fontSize: 10 }}>
                  <span className="text-white/50">{tool}:</span> {count}
                </div>
              ))}
            </div>
          </div>
        )}
        {extras.missing_tools && extras.missing_tools.length > 0 && (
          <div className="text-red-400/60" style={{ fontSize: 11 }}>
            <span className="text-white/25">Missing required tools: </span>
            {extras.missing_tools.join(", ")}
          </div>
        )}
        {extras.forbidden_used && extras.forbidden_used.length > 0 && (
          <div className="text-red-400/60" style={{ fontSize: 11 }}>
            <span className="text-white/25">Forbidden tools used: </span>
            {extras.forbidden_used.join(", ")}
          </div>
        )}
        {extras.violations && extras.violations.length > 0 && (
          <div>
            <div className="text-red-400/60 mb-1" style={{ fontSize: 10 }}>
              Violations:
            </div>
            <ul className="list-disc list-inside text-red-400/50" style={{ fontSize: 11 }}>
              {extras.violations.map((v, i) => (
                <li key={i}>{v}</li>
              ))}
            </ul>
          </div>
        )}
      </div>
    </div>
  );
}
