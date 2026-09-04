import { Component, type ErrorInfo, type ReactNode } from "react";
import { AlertCircle } from "lucide-react";

const mono = { fontFamily: "'JetBrains Mono', monospace" };

export interface ErrorBoundaryProps {
  /** Friendly label for the surface that failed (e.g. "Dashboard"). */
  surface?: string;
  children: ReactNode;
}

interface ErrorBoundaryState {
  error: Error | null;
}

/**
 * Generic React error boundary with a friendlier fallback than the bare
 * runtime stack trace. Use sparingly — wrap whole pages or expensive
 * sections, not individual cards. The reload button forces a remount.
 */
export class ErrorBoundary extends Component<ErrorBoundaryProps, ErrorBoundaryState> {
  state: ErrorBoundaryState = { error: null };

  static getDerivedStateFromError(error: Error): ErrorBoundaryState {
    return { error };
  }

  componentDidCatch(error: Error, info: ErrorInfo) {
    // Surface to the console so devtools still has the stack;
    // production telemetry can hook in here later.
    // eslint-disable-next-line no-console
    console.error(`[ErrorBoundary${this.props.surface ? `: ${this.props.surface}` : ""}]`, error, info);
  }

  handleReset = () => {
    this.setState({ error: null });
  };

  render() {
    if (!this.state.error) return this.props.children;

    const { surface } = this.props;
    const message = this.state.error.message || "Unknown error";

    return (
      <div className="flex min-h-[60vh] items-center justify-center bg-[#0a0a0f] px-6 py-10">
        <div className="max-w-xl rounded-xl border border-amber-500/30 bg-amber-500/[0.04] p-6">
          <div className="flex items-center gap-3">
            <AlertCircle className="h-5 w-5 text-amber-400" />
            <h2 className="text-lg text-amber-300" style={{ fontWeight: 500 }}>
              {surface ? `${surface} hit a snag` : "Something went sideways"}
            </h2>
          </div>
          <p className="mt-3 text-sm text-white/70">
            The page couldn't render. The underlying data is probably missing a
            field the UI assumed was always present.
          </p>
          <pre
            className="mt-4 overflow-x-auto rounded-md border border-white/8 bg-black/30 px-3 py-2 text-amber-200/80"
            style={{ ...mono, fontSize: 11 }}
          >
            {message}
          </pre>
          <div className="mt-5 flex gap-3">
            <button
              type="button"
              onClick={this.handleReset}
              className="rounded-md border border-white/10 bg-white/[0.04] px-3 py-1.5 text-xs text-white/80 transition hover:bg-white/[0.08]"
            >
              Try again
            </button>
            <button
              type="button"
              onClick={() => window.location.reload()}
              className="rounded-md border border-white/10 bg-white/[0.04] px-3 py-1.5 text-xs text-white/80 transition hover:bg-white/[0.08]"
            >
              Reload page
            </button>
          </div>
          <p className="mt-4 text-[11px] text-white/30">
            If you're a developer, the full stack is in the browser console.
          </p>
        </div>
      </div>
    );
  }
}
