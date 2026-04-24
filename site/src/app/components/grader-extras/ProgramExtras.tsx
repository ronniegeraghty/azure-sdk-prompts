import type { ProgramExtras as ProgramExtrasType } from "../../data/types";

const mono = { fontFamily: "'JetBrains Mono', monospace" };

export function ProgramExtras({ extras }: { extras: ProgramExtrasType }) {
  return (
    <div>
      <div className="mb-1.5 text-white/25" style={{ fontSize: 10 }}>
        Program Execution
      </div>
      <div className="rounded-md bg-black/30 p-3 space-y-2">
        <div>
          <span className="text-white/25" style={{ fontSize: 10 }}>
            Command:{" "}
          </span>
          <code className="text-white/60" style={{ ...mono, fontSize: 11 }}>
            {extras.command}
            {extras.args && extras.args.length > 0 && ` ${extras.args.join(" ")}`}
          </code>
        </div>
        <div>
          <span className="text-white/25" style={{ fontSize: 10 }}>
            Exit code:{" "}
          </span>
          <span
            className={extras.exit_code === 0 ? "text-emerald-400" : "text-red-400"}
            style={{ ...mono, fontSize: 11 }}
          >
            {extras.exit_code}
          </span>
        </div>
        {extras.duration_ms !== undefined && (
          <div>
            <span className="text-white/25" style={{ fontSize: 10 }}>
              Duration:{" "}
            </span>
            <span className="text-white/40" style={{ fontSize: 11 }}>
              {extras.duration_ms}ms
            </span>
          </div>
        )}
        {extras.stdout && (
          <div>
            <div className="text-white/25 mb-1" style={{ fontSize: 10 }}>
              Stdout:
            </div>
            <pre
              className="text-white/50 overflow-auto max-h-40"
              style={{ ...mono, fontSize: 10 }}
            >
              {extras.stdout}
            </pre>
          </div>
        )}
        {extras.stderr && (
          <div>
            <div className="text-red-400/60 mb-1" style={{ fontSize: 10 }}>
              Stderr:
            </div>
            <pre
              className="text-red-400/50 overflow-auto max-h-40"
              style={{ ...mono, fontSize: 10 }}
            >
              {extras.stderr}
            </pre>
          </div>
        )}
      </div>
    </div>
  );
}
