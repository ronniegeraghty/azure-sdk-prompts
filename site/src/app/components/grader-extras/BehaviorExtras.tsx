import type { BehaviorExtras as BehaviorExtrasType } from "../../data/types";

export function BehaviorExtras({ extras }: { extras: BehaviorExtrasType }) {
  return (
    <div>
      <div className="mb-1.5 text-white/25" style={{ fontSize: 10 }}>
        Behavior Analysis
      </div>
      <div className="grid gap-2 sm:grid-cols-2 text-white/40" style={{ fontSize: 11 }}>
        {extras.tools_used && extras.tools_used.length > 0 && (
          <div>
            <span className="text-white/25">Tools used: </span>
            {extras.tools_used.join(", ")}
          </div>
        )}
        {extras.turn_count !== undefined && (
          <div>
            <span className="text-white/25">Turns: </span>
            {extras.turn_count}
            {extras.max_turns !== undefined && ` / ${extras.max_turns}`}
          </div>
        )}
        {extras.total_actions !== undefined && (
          <div>
            <span className="text-white/25">Total actions: </span>
            {extras.total_actions}
          </div>
        )}
        {extras.turn_limit_hit && (
          <div className="text-amber-400/70">
            <span className="text-white/25">Turn limit: </span>
            Hit
          </div>
        )}
        {extras.missing_tools && extras.missing_tools.length > 0 && (
          <div className="col-span-2 text-red-400/60">
            <span className="text-white/25">Missing tools: </span>
            {extras.missing_tools.join(", ")}
          </div>
        )}
        {extras.forbidden_used && extras.forbidden_used.length > 0 && (
          <div className="col-span-2 text-red-400/60">
            <span className="text-white/25">Forbidden tools used: </span>
            {extras.forbidden_used.join(", ")}
          </div>
        )}
        {extras.violations && extras.violations.length > 0 && (
          <div className="col-span-2 text-red-400/60">
            <span className="text-white/25">Violations: </span>
            {extras.violations.join(", ")}
          </div>
        )}
      </div>
    </div>
  );
}
