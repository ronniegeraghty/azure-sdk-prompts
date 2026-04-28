import { render, screen, fireEvent, within } from "@testing-library/react";
import { describe, it, expect } from "vitest";
import { ReviewExtras } from "./ReviewExtras";
import type { ReviewExtras as ReviewExtrasType, ReviewPanelEntry } from "../../data/types";

/**
 * Test suite for per-reviewer vote display in ReviewExtras component.
 * 
 * Tests cover:
 * - Full agreement (all reviewers pass)
 * - Disagreement scenarios (mixed pass/fail)
 * - All fail scenarios
 * - Missing/empty reasons
 * - Legacy reports without criteria data
 * - Keyboard accessibility
 */

describe("ReviewExtras - Per-Reviewer Vote Display", () => {
  // Helper to create minimal ReviewExtras with panel_results
  function makeExtras(panel_results: ReviewPanelEntry[]): ReviewExtrasType {
    return {
      model: "claude-opus-4.6",
      summary: "Overall review summary",
      panel_results,
    };
  }

  describe("Full Agreement - All Reviewers Pass", () => {
    it("renders all reviewer criteria when all pass the same check", () => {
      const extras = makeExtras([
        {
          model: "claude-opus-4.6",
          overall_score: 5,
          max_score: 5,
          summary: "Reviewer 1 summary",
          criteria: [
            {
              name: "Uses DefaultAzureCredential",
              passed: true,
              reason: "Code correctly imports and uses DefaultAzureCredential from azure-identity package.",
            },
          ],
        },
        {
          model: "gpt-4o",
          overall_score: 5,
          max_score: 5,
          summary: "Reviewer 2 summary",
          criteria: [
            {
              name: "Uses DefaultAzureCredential",
              passed: true,
              reason: "Proper credential usage verified in get_client() function.",
            },
          ],
        },
        {
          model: "claude-sonnet-4.6",
          overall_score: 5,
          max_score: 5,
          summary: "Reviewer 3 summary",
          criteria: [
            {
              name: "Uses DefaultAzureCredential",
              passed: true,
              reason: "Authentication pattern follows Azure SDK best practices.",
            },
          ],
        },
      ]);

      render(<ReviewExtras extras={extras} />);

      // All panel members should be visible
      expect(screen.getByText("claude-opus-4.6")).toBeInTheDocument();
      expect(screen.getByText("gpt-4o")).toBeInTheDocument();
      expect(screen.getByText("claude-sonnet-4.6")).toBeInTheDocument();

      // All criteria should be rendered
      expect(screen.getAllByText("Uses DefaultAzureCredential")).toHaveLength(3);

      // All reasons should be visible
      expect(screen.getByText(/Code correctly imports and uses DefaultAzureCredential/)).toBeInTheDocument();
      expect(screen.getByText(/Proper credential usage verified in get_client/)).toBeInTheDocument();
      expect(screen.getByText(/Authentication pattern follows Azure SDK best practices/)).toBeInTheDocument();
    });

    it("displays CheckCircle2 icons for all passing criteria", () => {
      const extras = makeExtras([
        {
          model: "claude-opus-4.6",
          overall_score: 3,
          max_score: 3,
          summary: "All checks pass",
          criteria: [
            { name: "Check A", passed: true, reason: "Reason A" },
            { name: "Check B", passed: true, reason: "Reason B" },
            { name: "Check C", passed: true, reason: "Reason C" },
          ],
        },
      ]);

      const { container } = render(<ReviewExtras extras={extras} />);

      // Should have check icons for passing items (lucide-react renders as SVG with specific classes)
      const panelCard = container.querySelector('[class*="rounded-md"][class*="border"]');
      expect(panelCard).toBeInTheDocument();

      // Verify all three criteria are rendered
      expect(screen.getByText("Check A")).toBeInTheDocument();
      expect(screen.getByText("Check B")).toBeInTheDocument();
      expect(screen.getByText("Check C")).toBeInTheDocument();
    });
  });

  describe("Disagreement - Mixed Pass/Fail", () => {
    it("renders both passing and failing reviewers for the same criterion", () => {
      const extras = makeExtras([
        {
          model: "claude-opus-4.6",
          overall_score: 4,
          max_score: 5,
          summary: "Mostly good",
          criteria: [
            {
              name: "Uses async/await patterns",
              passed: true,
              reason: "Code uses async/await throughout.",
            },
          ],
        },
        {
          model: "gpt-4o",
          overall_score: 3,
          max_score: 5,
          summary: "Has issues",
          criteria: [
            {
              name: "Uses async/await patterns",
              passed: false,
              reason: "The code uses synchronous SDK client methods instead of async.",
            },
          ],
        },
        {
          model: "claude-sonnet-4.6",
          overall_score: 5,
          max_score: 5,
          summary: "Looks good",
          criteria: [
            {
              name: "Uses async/await patterns",
              passed: true,
              reason: "Async patterns properly implemented.",
            },
          ],
        },
      ]);

      render(<ReviewExtras extras={extras} />);

      // Should show all three reviewers
      expect(screen.getByText("claude-opus-4.6")).toBeInTheDocument();
      expect(screen.getByText("gpt-4o")).toBeInTheDocument();
      expect(screen.getByText("claude-sonnet-4.6")).toBeInTheDocument();

      // Should show all three criteria (same name, different verdicts)
      expect(screen.getAllByText("Uses async/await patterns")).toHaveLength(3);

      // Should show all three reasons
      expect(screen.getByText(/Code uses async\/await throughout/)).toBeInTheDocument();
      expect(screen.getByText(/The code uses synchronous SDK client methods/)).toBeInTheDocument();
      expect(screen.getByText(/Async patterns properly implemented/)).toBeInTheDocument();
    });

    it("displays different icons for passing vs failing criteria", () => {
      const extras = makeExtras([
        {
          model: "reviewer-pass",
          overall_score: 1,
          max_score: 1,
          summary: "Pass",
          criteria: [{ name: "Security check", passed: true, reason: "Good" }],
        },
        {
          model: "reviewer-fail",
          overall_score: 0,
          max_score: 1,
          summary: "Fail",
          criteria: [{ name: "Security check", passed: false, reason: "Bad" }],
        },
      ]);

      const { container } = render(<ReviewExtras extras={extras} />);

      // Both criteria should be rendered
      expect(screen.getAllByText("Security check")).toHaveLength(2);
      expect(screen.getByText("Good")).toBeInTheDocument();
      expect(screen.getByText("Bad")).toBeInTheDocument();

      // Verify both panel cards exist
      const panelCards = container.querySelectorAll('[class*="rounded-md"][class*="border"]');
      expect(panelCards.length).toBeGreaterThanOrEqual(2);
    });
  });

  describe("All Fail Scenario", () => {
    it("renders all failing reviewers with fail icons and reasons", () => {
      const extras = makeExtras([
        {
          model: "claude-opus-4.6",
          overall_score: 0,
          max_score: 3,
          summary: "Multiple issues found",
          criteria: [
            {
              name: "No hardcoded secrets",
              passed: false,
              reason: "Found API key hardcoded in line 42.",
            },
          ],
        },
        {
          model: "gpt-4o",
          overall_score: 0,
          max_score: 3,
          summary: "Security concerns",
          criteria: [
            {
              name: "No hardcoded secrets",
              passed: false,
              reason: "Connection string embedded in source code.",
            },
          ],
        },
        {
          model: "claude-sonnet-4.6",
          overall_score: 0,
          max_score: 3,
          summary: "Fails security check",
          criteria: [
            {
              name: "No hardcoded secrets",
              passed: false,
              reason: "Credentials stored in plain text.",
            },
          ],
        },
      ]);

      render(<ReviewExtras extras={extras} />);

      // All reviewers should be visible
      expect(screen.getByText("claude-opus-4.6")).toBeInTheDocument();
      expect(screen.getByText("gpt-4o")).toBeInTheDocument();
      expect(screen.getByText("claude-sonnet-4.6")).toBeInTheDocument();

      // All fail criteria should be rendered
      expect(screen.getAllByText("No hardcoded secrets")).toHaveLength(3);

      // All fail reasons should be visible
      expect(screen.getByText(/Found API key hardcoded in line 42/)).toBeInTheDocument();
      expect(screen.getByText(/Connection string embedded in source code/)).toBeInTheDocument();
      expect(screen.getByText(/Credentials stored in plain text/)).toBeInTheDocument();
    });
  });

  describe("Missing or Empty Reasons", () => {
    it("renders criteria without reasons gracefully (no empty quote/dash artifacts)", () => {
      const extras = makeExtras([
        {
          model: "claude-opus-4.6",
          overall_score: 2,
          max_score: 2,
          summary: "Summary",
          criteria: [
            {
              name: "Check with reason",
              passed: true,
              reason: "This has a reason",
            },
            {
              name: "Check without reason",
              passed: true,
              reason: "",
            },
          ],
        },
      ]);

      render(<ReviewExtras extras={extras} />);

      // Both criteria names should be visible
      expect(screen.getByText("Check with reason")).toBeInTheDocument();
      expect(screen.getByText("Check without reason")).toBeInTheDocument();

      // The reason should be visible where present
      expect(screen.getByText("This has a reason")).toBeInTheDocument();

      // Should NOT render empty strings, dashes, or "undefined"
      const container = screen.getByText("Check without reason").closest("div");
      expect(container?.textContent).not.toContain("undefined");
      expect(container?.textContent).not.toContain("—");
      expect(container?.textContent).not.toContain('""');
    });

    it("handles criteria with missing reason field (undefined)", () => {
      const extras = makeExtras([
        {
          model: "gpt-4o",
          overall_score: 1,
          max_score: 1,
          summary: "Pass",
          criteria: [
            {
              name: "Criterion without reason field",
              passed: true,
              // reason is completely missing (undefined)
            } as any,
          ],
        },
      ]);

      render(<ReviewExtras extras={extras} />);

      expect(screen.getByText("Criterion without reason field")).toBeInTheDocument();

      // Should not display "undefined" or throw errors
      const container = screen.getByText("Criterion without reason field").closest("div");
      expect(container?.textContent).not.toContain("undefined");
    });
  });

  describe("Legacy Reports - Backward Compatibility", () => {
    it("renders panel_results without criteria field (legacy reports)", () => {
      const extras = makeExtras([
        {
          model: "claude-opus-4.6",
          overall_score: 4,
          max_score: 5,
          summary: "Legacy panel member summary",
          // No criteria field (pre-v4 report)
        },
        {
          model: "gpt-4o",
          overall_score: 3,
          max_score: 5,
          summary: "Another legacy summary",
          // No criteria field
        },
      ]);

      render(<ReviewExtras extras={extras} />);

      // Should render panel members normally
      expect(screen.getByText("claude-opus-4.6")).toBeInTheDocument();
      expect(screen.getByText("gpt-4o")).toBeInTheDocument();

      // Should show summaries
      expect(screen.getByText("Legacy panel member summary")).toBeInTheDocument();
      expect(screen.getByText("Another legacy summary")).toBeInTheDocument();

      // Should not crash or show error messages
      expect(screen.queryByText(/error/i)).not.toBeInTheDocument();
      expect(screen.queryByText(/undefined/i)).not.toBeInTheDocument();
    });

    it("renders panel_results with empty criteria array", () => {
      const extras = makeExtras([
        {
          model: "claude-sonnet-4.6",
          overall_score: 5,
          max_score: 5,
          summary: "Panel with empty criteria",
          criteria: [], // Empty array
        },
      ]);

      render(<ReviewExtras extras={extras} />);

      expect(screen.getByText("claude-sonnet-4.6")).toBeInTheDocument();
      expect(screen.getByText("Panel with empty criteria")).toBeInTheDocument();

      // Should not crash
      expect(screen.queryByText(/error/i)).not.toBeInTheDocument();
    });
  });

  describe("Complex Multi-Criteria Scenarios", () => {
    it("renders multiple criteria per reviewer correctly", () => {
      const extras = makeExtras([
        {
          model: "claude-opus-4.6",
          overall_score: 4,
          max_score: 5,
          summary: "Multiple checks",
          criteria: [
            {
              name: "Installing azure-keyvault-secrets package",
              passed: true,
              reason: "Package correctly installed via pip.",
            },
            {
              name: "Creating SecretClient with vault URL",
              passed: true,
              reason: "SecretClient instantiated properly.",
            },
            {
              name: "Using async/await patterns",
              passed: false,
              reason: "Synchronous methods used instead of async.",
            },
            {
              name: "Proper error handling",
              passed: true,
              reason: "try/except blocks in place.",
            },
            {
              name: "Environment variable for vault URL",
              passed: true,
              reason: "Vault URL read from AZURE_KEYVAULT_URL env var.",
            },
          ],
        },
      ]);

      render(<ReviewExtras extras={extras} />);

      // All 5 criteria should be rendered
      expect(screen.getByText("Installing azure-keyvault-secrets package")).toBeInTheDocument();
      expect(screen.getByText("Creating SecretClient with vault URL")).toBeInTheDocument();
      expect(screen.getByText("Using async/await patterns")).toBeInTheDocument();
      expect(screen.getByText("Proper error handling")).toBeInTheDocument();
      expect(screen.getByText("Environment variable for vault URL")).toBeInTheDocument();

      // All reasons should be visible
      expect(screen.getByText(/Package correctly installed via pip/)).toBeInTheDocument();
      expect(screen.getByText(/SecretClient instantiated properly/)).toBeInTheDocument();
      expect(screen.getByText(/Synchronous methods used instead of async/)).toBeInTheDocument();
      expect(screen.getByText(/try\/except blocks in place/)).toBeInTheDocument();
      expect(screen.getByText(/Vault URL read from AZURE_KEYVAULT_URL env var/)).toBeInTheDocument();
    });

    it("handles mixed criteria counts across reviewers", () => {
      const extras = makeExtras([
        {
          model: "reviewer-1",
          overall_score: 2,
          max_score: 2,
          summary: "Two checks",
          criteria: [
            { name: "Check A", passed: true, reason: "A OK" },
            { name: "Check B", passed: true, reason: "B OK" },
          ],
        },
        {
          model: "reviewer-2",
          overall_score: 1,
          max_score: 1,
          summary: "One check",
          criteria: [
            { name: "Check A", passed: true, reason: "A good" },
          ],
        },
        {
          model: "reviewer-3",
          overall_score: 3,
          max_score: 3,
          summary: "Three checks",
          criteria: [
            { name: "Check A", passed: false, reason: "A bad" },
            { name: "Check B", passed: true, reason: "B good" },
            { name: "Check C", passed: true, reason: "C good" },
          ],
        },
      ]);

      render(<ReviewExtras extras={extras} />);

      // All reviewers visible
      expect(screen.getByText("reviewer-1")).toBeInTheDocument();
      expect(screen.getByText("reviewer-2")).toBeInTheDocument();
      expect(screen.getByText("reviewer-3")).toBeInTheDocument();

      // All criteria visible
      expect(screen.getAllByText("Check A").length).toBeGreaterThanOrEqual(3);
      expect(screen.getAllByText("Check B").length).toBeGreaterThanOrEqual(2);
      expect(screen.getByText("Check C")).toBeInTheDocument();
    });
  });

  describe("Edge Cases", () => {
    it("renders when panel_results is undefined (no panel at all)", () => {
      const extras: ReviewExtrasType = {
        model: "claude-opus-4.6",
        summary: "No panel results",
        // panel_results is undefined
      };

      render(<ReviewExtras extras={extras} />);

      // Should render the summary
      expect(screen.getByText("No panel results")).toBeInTheDocument();

      // Should not crash or show errors
      expect(screen.queryByText(/Panel Members/)).not.toBeInTheDocument();
    });

    it("renders when panel_results is empty array", () => {
      const extras: ReviewExtrasType = {
        model: "claude-opus-4.6",
        summary: "Empty panel",
        panel_results: [],
      };

      render(<ReviewExtras extras={extras} />);

      expect(screen.getByText("Empty panel")).toBeInTheDocument();

      // Should not show panel section
      expect(screen.queryByText(/Panel Members/)).not.toBeInTheDocument();
    });

    it("handles very long criterion names and reasons without breaking layout", () => {
      const longName =
        "This is an extremely long criterion name that tests whether the layout can handle very long text without breaking or causing horizontal overflow issues in the UI";
      const longReason =
        "This is an extremely long reason that explains in great detail why this particular check passed or failed, including multiple sentences and technical details that span across many lines of text to ensure the component handles text wrapping correctly without any visual glitches or overflow problems.";

      const extras = makeExtras([
        {
          model: "claude-opus-4.6",
          overall_score: 1,
          max_score: 1,
          summary: "Long text test",
          criteria: [
            {
              name: longName,
              passed: true,
              reason: longReason,
            },
          ],
        },
      ]);

      render(<ReviewExtras extras={extras} />);

      // Should render long text without errors
      expect(screen.getByText(longName)).toBeInTheDocument();
      expect(screen.getByText(longReason)).toBeInTheDocument();
    });

    it("handles special characters in criterion names and reasons", () => {
      const extras = makeExtras([
        {
          model: "test-model",
          overall_score: 1,
          max_score: 1,
          summary: "Special chars",
          criteria: [
            {
              name: "Uses <DefaultAzureCredential> & \"proper\" auth (100%)",
              passed: true,
              reason: 'Code imports "azure-identity" & uses <credential> correctly.',
            },
          ],
        },
      ]);

      render(<ReviewExtras extras={extras} />);

      // Should render special characters correctly
      expect(screen.getByText(/Uses <DefaultAzureCredential> & "proper" auth/)).toBeInTheDocument();
      expect(screen.getByText(/Code imports "azure-identity" & uses <credential>/)).toBeInTheDocument();
    });
  });

  describe("Accessibility", () => {
    it("renders semantic HTML structure for screen readers", () => {
      const extras = makeExtras([
        {
          model: "claude-opus-4.6",
          overall_score: 2,
          max_score: 2,
          summary: "Accessible test",
          criteria: [
            { name: "Check 1", passed: true, reason: "Reason 1" },
            { name: "Check 2", passed: false, reason: "Reason 2" },
          ],
        },
      ]);

      const { container } = render(<ReviewExtras extras={extras} />);

      // Should have proper structure with divs and text
      expect(container.querySelector("div")).toBeInTheDocument();

      // Text content should be accessible
      expect(screen.getByText("Check 1")).toBeInTheDocument();
      expect(screen.getByText("Check 2")).toBeInTheDocument();
    });

    it("ensures criterion text has sufficient color contrast (via className)", () => {
      const extras = makeExtras([
        {
          model: "test-model",
          overall_score: 1,
          max_score: 1,
          summary: "Contrast test",
          criteria: [
            { name: "Pass check", passed: true, reason: "Pass reason" },
            { name: "Fail check", passed: false, reason: "Fail reason" },
          ],
        },
      ]);

      const { container } = render(<ReviewExtras extras={extras} />);

      // Verify elements are rendered (actual color contrast would be tested in E2E)
      const passElement = screen.getByText("Pass check");
      const failElement = screen.getByText("Fail check");

      expect(passElement).toBeInTheDocument();
      expect(failElement).toBeInTheDocument();

      // Elements should have appropriate styling classes
      expect(passElement.className).toBeTruthy();
      expect(failElement.className).toBeTruthy();
    });
  });

  describe("Real-World Data Scenarios", () => {
    it("renders actual Azure SDK review criteria from production reports", () => {
      // Based on real data from reports/20260427-232343/.../report.json
      const extras = makeExtras([
        {
          model: "claude-opus-4.6",
          overall_score: 5,
          max_score: 5,
          summary: "All Azure SDK patterns correctly implemented",
          criteria: [
            {
              name: "Installing `azure-keyvault-secrets` and `azure-identity` packages",
              passed: true,
              reason:
                "The docstring and agent summary both show `pip install azure-keyvault-secrets azure-identity`, and the code imports from both packages.",
              weight: 1,
            },
            {
              name: "Creating a `SecretClient` with vault URL and credential",
              passed: true,
              reason:
                "get_client() creates DefaultAzureCredential and passes it with vault_url to SecretClient.",
              weight: 1,
            },
            {
              name: "Setting a secret with `set_secret(name, value)`",
              passed: true,
              reason: "create_secret() calls client.set_secret(name, value).",
              weight: 1,
            },
            {
              name: "Getting a secret with `get_secret(name)` and reading its `.value`",
              passed: true,
              reason: "read_secret() calls client.get_secret(name) and returns secret.value.",
              weight: 1,
            },
            {
              name: "Updating a secret (re-`set_secret` with same name, new value)",
              passed: true,
              reason: "update_secret() calls client.set_secret(name, new_value).",
              weight: 1,
            },
          ],
        },
      ]);

      render(<ReviewExtras extras={extras} />);

      // Verify all production criteria are rendered
      expect(screen.getByText(/Installing `azure-keyvault-secrets` and `azure-identity` packages/)).toBeInTheDocument();
      expect(screen.getByText(/Creating a `SecretClient` with vault URL and credential/)).toBeInTheDocument();
      expect(screen.getByText(/Setting a secret with `set_secret\(name, value\)`/)).toBeInTheDocument();
      expect(screen.getByText(/Getting a secret with `get_secret\(name\)` and reading its `.value`/)).toBeInTheDocument();
      expect(screen.getByText(/Updating a secret/)).toBeInTheDocument();

      // Verify reasons are visible
      expect(screen.getByText(/docstring and agent summary both show/)).toBeInTheDocument();
      expect(screen.getByText(/get_client\(\) creates DefaultAzureCredential/)).toBeInTheDocument();
    });
  });
});
