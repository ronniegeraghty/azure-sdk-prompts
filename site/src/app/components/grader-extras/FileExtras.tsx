import { CheckCircle2, XCircle } from "lucide-react";
import type { FileExtras as FileExtrasType } from "../../data/types";

const mono = { fontFamily: "'JetBrains Mono', monospace" };

export function FileExtras({ extras }: { extras: FileExtrasType }) {
  return (
    <div>
      <div className="mb-1.5 text-white/25" style={{ fontSize: 10 }}>
        File Checks
      </div>
      <div className="space-y-1">
        {extras.files.map((f, i) => (
          <div
            key={i}
            className="flex items-center gap-2 text-white/40"
            style={{ fontSize: 11 }}
          >
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
            {f.size !== undefined && (
              <span className="text-white/20">({f.size} bytes)</span>
            )}
          </div>
        ))}
      </div>
    </div>
  );
}
