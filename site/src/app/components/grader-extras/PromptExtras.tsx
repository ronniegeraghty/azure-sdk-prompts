import type { PromptExtras as PromptExtrasType } from "../../data/types";

const mono = { fontFamily: "'JetBrains Mono', monospace" };

export function PromptExtras({ extras }: { extras: PromptExtrasType }) {
  return (
    <div>
      <div className="mb-1.5 text-white/25" style={{ fontSize: 10 }}>
        LLM Review
      </div>
      <div className="space-y-2">
        <div className="text-white/40" style={{ fontSize: 11 }}>
          Model: <span style={mono}>{extras.model}</span>
        </div>
        {extras.rubric && (
          <div className="text-white/40" style={{ fontSize: 11 }}>
            Rubric: <span style={mono}>{extras.rubric}</span>
          </div>
        )}
        <div className="text-white/40" style={{ fontSize: 11 }}>
          Score:{" "}
          <span style={mono}>
            {extras.raw_score}/{extras.max_score}
          </span>
        </div>
        {extras.reasoning && (
          <div>
            <div className="text-white/25 mb-1" style={{ fontSize: 10 }}>
              Reasoning:
            </div>
            <p className="text-white/50" style={{ fontSize: 11, lineHeight: 1.5 }}>
              {extras.reasoning}
            </p>
          </div>
        )}
      </div>
    </div>
  );
}
