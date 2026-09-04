import type { OutputCheckExtras as OutputCheckExtrasType } from "../../data/types";

const mono = { fontFamily: "'JetBrains Mono', monospace" };

export function OutputCheckExtras({ extras }: { extras: OutputCheckExtrasType }) {
  const totalSize = extras.produced_files.reduce((sum, f) => sum + f.size, 0);

  return (
    <div>
      <div className="mb-1.5 text-white/25" style={{ fontSize: 10 }}>
        Produced Files
      </div>
      <div className="space-y-2">
        <div className="text-white/40" style={{ fontSize: 11 }}>
          <span className="text-white/25">File count: </span>
          {extras.produced_files.length}
        </div>
        <div className="text-white/40" style={{ fontSize: 11 }}>
          <span className="text-white/25">Total size: </span>
          {formatBytes(totalSize)}
        </div>
        {extras.produced_files.length > 0 && (
          <div>
            <div className="text-white/25 mb-1" style={{ fontSize: 10 }}>
              Files:
            </div>
            <div className="space-y-1">
              {extras.produced_files.map((f, i) => (
                <div
                  key={i}
                  className="flex items-center justify-between text-white/40"
                  style={{ ...mono, fontSize: 10 }}
                >
                  <span className="truncate">{f.path}</span>
                  <span className="text-white/25 ml-2 shrink-0">
                    {formatBytes(f.size)}
                  </span>
                </div>
              ))}
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

function formatBytes(bytes: number): string {
  if (bytes === 0) return "0 B";
  const k = 1024;
  const sizes = ["B", "KB", "MB", "GB"];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return `${parseFloat((bytes / Math.pow(k, i)).toFixed(1))} ${sizes[i]}`;
}
