import { render, screen, fireEvent, waitFor } from "@testing-library/react";
import { describe, it, expect, vi } from "vitest";

// Inline minimal component for testing — matches the structure in eval-detail-page.tsx
function RerunCommand({ command, mono }: { command: string; mono: object }) {
  const [copied, setCopied] = React.useState(false);
  const handleCopy = () => {
    navigator.clipboard.writeText(command);
    setCopied(true);
    setTimeout(() => setCopied(false), 2000);
  };
  return (
    <div className="mb-6 rounded-lg border border-white/10 bg-white/[0.02] p-4" data-testid="rerun-panel">
      <div className="mb-2 flex items-center gap-2 text-white/50" style={{ fontSize: 12 }}>
        <span>Reproduce this eval</span>
      </div>
      <div className="flex items-center gap-2">
        <code className="flex-1 overflow-auto rounded bg-black/40 px-3 py-2 text-white/70" style={{ ...mono, fontSize: 11 }}>
          {command}
        </code>
        <button onClick={handleCopy} className="text-white/30 transition hover:text-white/60">
          {copied ? "✓" : "Copy"}
        </button>
      </div>
    </div>
  );
}

// eslint-disable-next-line @typescript-eslint/no-namespace
declare global {
  interface Window {
    clipboardData: any;
  }
}

const React = await import("react");

describe("RerunCommand", () => {
  const mono = { fontFamily: "'JetBrains Mono', monospace" };
  
  // Mock clipboard API
  const mockClipboard = {
    writeText: vi.fn().mockResolvedValue(undefined),
  };
  Object.assign(navigator, { clipboard: mockClipboard });

  it("renders the command text correctly", () => {
    render(<RerunCommand command="hyoka run --prompt-id test --config baseline/opus" mono={mono} />);
    expect(screen.getByText("hyoka run --prompt-id test --config baseline/opus")).toBeInTheDocument();
  });

  it("shows 'Reproduce this eval' label", () => {
    render(<RerunCommand command="hyoka run --prompt-id test --config baseline/opus" mono={mono} />);
    expect(screen.getByText("Reproduce this eval")).toBeInTheDocument();
  });

  it("copies command to clipboard when copy button clicked", async () => {
    render(<RerunCommand command="hyoka run --prompt-id test --config baseline/opus" mono={mono} />);
    const copyButton = screen.getByRole("button", { name: /copy/i });
    fireEvent.click(copyButton);
    await waitFor(() => {
      expect(mockClipboard.writeText).toHaveBeenCalledWith("hyoka run --prompt-id test --config baseline/opus");
    });
  });

  it("shows checkmark temporarily after successful copy", async () => {
    render(<RerunCommand command="hyoka run --prompt-id test --config baseline/opus" mono={mono} />);
    const copyButton = screen.getByRole("button", { name: /copy/i });
    fireEvent.click(copyButton);
    await waitFor(() => {
      expect(screen.getByText("✓")).toBeInTheDocument();
    });
  });

  it("handles commands with special characters", () => {
    render(<RerunCommand command='hyoka run --prompt-id "key-vault-dp-python-crud" --config "baseline/claude-opus-4.6,azure-mcp/opus"' mono={mono} />);
    expect(screen.getByText(/hyoka run --prompt-id "key-vault-dp-python-crud"/)).toBeInTheDocument();
  });

  it("renders in a bordered panel with proper styling", () => {
    render(<RerunCommand command="hyoka run --prompt-id test --config baseline/opus" mono={mono} />);
    const panel = screen.getByTestId("rerun-panel");
    expect(panel).toHaveClass("rounded-lg", "border", "border-white/10");
  });
});
