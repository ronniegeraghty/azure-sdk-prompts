import { render, screen, fireEvent } from "@testing-library/react";
import { describe, it, expect } from "vitest";
import { ExpandablePoint } from "./ExpandablePoint";
import type { GraderPoint } from "../data/types";

/**
 * Test suite for ExpandablePoint component — per-reviewer vote display.
 * 
 * Tests cover:
 * - Full agreement (all reviewers pass)
 * - Disagreement scenarios (mixed pass/fail) with auto-expand
 * - All fail scenarios
 * - Missing/empty reasons
 * - Amber badge display for split votes
 * - Keyboard accessibility
 */

describe("ExpandablePoint - Per-Reviewer Vote Display", () => {
  const basePoint: GraderPoint = {
    label: "Uses DefaultAzureCredential",
    pass: true,
    message: "Authentication implemented correctly",
  };

  describe("Full Agreement - All Reviewers Pass", () => {
    it("renders all reviewer votes when all pass the same check", () => {
      const reviewerVotes = [
        {
          model: "claude-opus-4.6",
          passed: true,
          reason: "Code correctly imports and uses DefaultAzureCredential from azure-identity package.",
        },
        {
          model: "gpt-4o",
          passed: true,
          reason: "Proper credential usage verified in get_client() function.",
        },
        {
          model: "claude-sonnet-4.6",
          passed: true,
          reason: "Authentication pattern follows Azure SDK best practices.",
        },
      ];

      render(
        <ExpandablePoint
          point={basePoint}
          labelText="Uses DefaultAzureCredential"
          secondary={null}
          reviewerVotes={reviewerVotes}
        />
      );

      // Point should be clickable (has reviewer votes)
      const pointDiv = screen.getByRole("button");
      expect(pointDiv).toBeInTheDocument();

      // No disagreement badge (all pass)
      expect(screen.queryByText(/⚠️/)).not.toBeInTheDocument();

      // Expand to see reviewer votes
      fireEvent.click(pointDiv);

      // All reviewer models should be visible
      expect(screen.getByText(/claude-opus-4\.6:/)).toBeInTheDocument();
      expect(screen.getByText(/gpt-4o:/)).toBeInTheDocument();
      expect(screen.getByText(/claude-sonnet-4\.6:/)).toBeInTheDocument();

      // All reasons should be visible
      expect(screen.getByText(/Code correctly imports and uses DefaultAzureCredential/)).toBeInTheDocument();
      expect(screen.getByText(/Proper credential usage verified in get_client/)).toBeInTheDocument();
      expect(screen.getByText(/Authentication pattern follows Azure SDK best practices/)).toBeInTheDocument();
    });

    it("does not auto-expand when all reviewers agree (pass)", () => {
      const reviewerVotes = [
        { model: "opus", passed: true, reason: "Good" },
        { model: "sonnet", passed: true, reason: "Good" },
      ];

      render(
        <ExpandablePoint
          point={basePoint}
          labelText="Check A"
          secondary={null}
          reviewerVotes={reviewerVotes}
        />
      );

      // Should be collapsed (reviewer reasons not visible)
      expect(screen.queryByText(/opus:/)).not.toBeInTheDocument();
      expect(screen.queryByText(/sonnet:/)).not.toBeInTheDocument();
    });
  });

  describe("Disagreement - Split Votes", () => {
    it("auto-expands when reviewers disagree on a criterion", () => {
      const reviewerVotes = [
        { model: "opus", passed: true, reason: "Looks good" },
        { model: "sonnet", passed: false, reason: "Missing error handling" },
        { model: "gpt", passed: true, reason: "Correct pattern" },
      ];

      render(
        <ExpandablePoint
          point={basePoint}
          labelText="Error handling"
          secondary={null}
          reviewerVotes={reviewerVotes}
        />
      );

      // Should auto-expand on disagreement
      expect(screen.getByText(/opus:/)).toBeInTheDocument();
      expect(screen.getByText(/sonnet:/)).toBeInTheDocument();
      expect(screen.getByText(/gpt:/)).toBeInTheDocument();
    });

    it("displays amber badge with split vote count (⚠️ N/M)", () => {
      const reviewerVotes = [
        { model: "opus", passed: true, reason: "Pass" },
        { model: "sonnet", passed: false, reason: "Fail" },
        { model: "gpt", passed: true, reason: "Pass" },
      ];

      render(
        <ExpandablePoint
          point={basePoint}
          labelText="Check"
          secondary={null}
          reviewerVotes={reviewerVotes}
        />
      );

      // Amber badge should show 2/3 (2 passed out of 3)
      expect(screen.getByText(/⚠️\s*2\/3/)).toBeInTheDocument();
    });

    it("renders different icons for passing vs failing criteria", () => {
      const reviewerVotes = [
        { model: "opus", passed: true, reason: "Good" },
        { model: "sonnet", passed: false, reason: "Bad" },
      ];

      const { container } = render(
        <ExpandablePoint
          point={basePoint}
          labelText="Check"
          secondary={null}
          reviewerVotes={reviewerVotes}
        />
      );

      // Should auto-expand due to disagreement
      // Should have both CheckCircle2 and XCircle icons
      // Testing library doesn't give us direct icon access, but we can check
      // the structure by looking for the reviewer rows
      expect(screen.getByText(/opus:/)).toBeInTheDocument();
      expect(screen.getByText(/sonnet:/)).toBeInTheDocument();
      expect(screen.getByText(/Good/)).toBeInTheDocument();
      expect(screen.getByText(/Bad/)).toBeInTheDocument();
    });
  });

  describe("All Fail Scenarios", () => {
    it("renders all failing reviewers with fail icons and reasons", () => {
      const reviewerVotes = [
        { model: "opus", passed: false, reason: "Missing imports" },
        { model: "sonnet", passed: false, reason: "Wrong package" },
        { model: "gpt", passed: false, reason: "No credential instantiation" },
      ];

      render(
        <ExpandablePoint
          point={{ ...basePoint, pass: false }}
          labelText="Uses DefaultAzureCredential"
          secondary={null}
          reviewerVotes={reviewerVotes}
        />
      );

      // No disagreement (all fail)
      expect(screen.queryByText(/⚠️/)).not.toBeInTheDocument();

      // Expand to see reasons
      fireEvent.click(screen.getByRole("button"));

      expect(screen.getByText(/Missing imports/)).toBeInTheDocument();
      expect(screen.getByText(/Wrong package/)).toBeInTheDocument();
      expect(screen.getByText(/No credential instantiation/)).toBeInTheDocument();
    });
  });

  describe("Missing/Empty Reasons", () => {
    it("renders criteria without reasons gracefully (shows Pass/Fail fallback)", () => {
      const reviewerVotes = [
        { model: "opus", passed: true },
        { model: "sonnet", passed: true },
      ];

      render(
        <ExpandablePoint
          point={basePoint}
          labelText="Check"
          secondary={null}
          reviewerVotes={reviewerVotes}
        />
      );

      // No disagreement, so not auto-expanded. Click to expand.
      fireEvent.click(screen.getByRole("button"));

      // Should show "Pass" as fallback (no disagreement, so need to expand first)
      const passText = screen.getAllByText(/Pass/);
      expect(passText.length).toBeGreaterThan(0);
    });

    it("handles criteria with undefined reason field", () => {
      const reviewerVotes = [
        { model: "opus", passed: true, reason: undefined },
        { model: "sonnet", passed: true, reason: undefined },
      ];

      render(
        <ExpandablePoint
          point={basePoint}
          labelText="Check"
          secondary={null}
          reviewerVotes={reviewerVotes}
        />
      );

      // No disagreement, expand manually
      fireEvent.click(screen.getByRole("button"));

      // Should not crash, should show fallback text
      const passText = screen.getAllByText(/Pass/);
      expect(passText.length).toBeGreaterThan(0);
    });

    it("handles empty string reason", () => {
      const reviewerVotes = [
        { model: "opus", passed: true, reason: "" },
      ];

      render(
        <ExpandablePoint
          point={basePoint}
          labelText="Check"
          secondary={null}
          reviewerVotes={reviewerVotes}
        />
      );

      fireEvent.click(screen.getByRole("button"));

      // Empty reason should show fallback
      expect(screen.getByText(/Pass/)).toBeInTheDocument();
    });
  });

  describe("No Reviewers (Legacy Fallback)", () => {
    it("renders point without reviewer votes (no expansion)", () => {
      render(
        <ExpandablePoint
          point={basePoint}
          labelText="Uses DefaultAzureCredential"
          secondary="Authentication implemented correctly"
          reviewerVotes={[]}
        />
      );

      // Should render the label
      expect(screen.getByText("Uses DefaultAzureCredential")).toBeInTheDocument();
      expect(screen.getByText("Authentication implemented correctly")).toBeInTheDocument();

      // Should NOT be clickable (no role=button)
      expect(screen.queryByRole("button")).not.toBeInTheDocument();

      // Should NOT have expand chevron
      const chevrons = document.querySelectorAll('svg');
      // Only the CheckCircle2 or XCircle icon, no ChevronRight
      expect(chevrons.length).toBe(1);
    });
  });

  describe("Multiple Criteria per Point", () => {
    it("renders multiple reviewer votes for a single criterion", () => {
      const reviewerVotes = [
        { model: "opus", passed: true, reason: "Reason 1" },
        { model: "sonnet", passed: true, reason: "Reason 2" },
        { model: "gpt", passed: true, reason: "Reason 3" },
        { model: "claude", passed: true, reason: "Reason 4" },
      ];

      render(
        <ExpandablePoint
          point={basePoint}
          labelText="Check"
          secondary={null}
          reviewerVotes={reviewerVotes}
        />
      );

      fireEvent.click(screen.getByRole("button"));

      // All 4 reviewers should be visible
      expect(screen.getByText(/opus:/)).toBeInTheDocument();
      expect(screen.getByText(/sonnet:/)).toBeInTheDocument();
      expect(screen.getByText(/gpt:/)).toBeInTheDocument();
      expect(screen.getByText(/claude:/)).toBeInTheDocument();
    });
  });

  describe("Edge Cases", () => {
    it("handles very long criterion names without breaking layout", () => {
      const longName = "This is a very long criterion name that should wrap gracefully without breaking the UI layout or causing horizontal scroll bars to appear unexpectedly";

      render(
        <ExpandablePoint
          point={basePoint}
          labelText={longName}
          secondary={null}
          reviewerVotes={[]}
        />
      );

      expect(screen.getByText(longName)).toBeInTheDocument();
    });

    it("handles very long reviewer reasons without breaking layout", () => {
      const longReason = "This is an extremely long reason that might be provided by a verbose LLM reviewer and should wrap properly without breaking the UI layout or causing any visual artifacts or horizontal scrolling issues in the interface";

      const reviewerVotes = [
        { model: "verbose-model", passed: true, reason: longReason },
      ];

      render(
        <ExpandablePoint
          point={basePoint}
          labelText="Check"
          secondary={null}
          reviewerVotes={reviewerVotes}
        />
      );

      fireEvent.click(screen.getByRole("button"));

      expect(screen.getByText(longReason)).toBeInTheDocument();
    });

    it("handles special characters in criterion names", () => {
      const specialName = 'Uses "DefaultAzureCredential" & <escaping> test';

      render(
        <ExpandablePoint
          point={basePoint}
          labelText={specialName}
          secondary={null}
          reviewerVotes={[]}
        />
      );

      expect(screen.getByText(specialName)).toBeInTheDocument();
    });

    it("handles special characters in reviewer reasons", () => {
      const specialReason = 'Code uses "correct" pattern & <proper> escaping';

      const reviewerVotes = [
        { model: "opus", passed: true, reason: specialReason },
      ];

      render(
        <ExpandablePoint
          point={basePoint}
          labelText="Check"
          secondary={null}
          reviewerVotes={reviewerVotes}
        />
      );

      fireEvent.click(screen.getByRole("button"));

      expect(screen.getByText(specialReason)).toBeInTheDocument();
    });
  });

  describe("Accessibility", () => {
    it("renders semantic HTML with proper ARIA attributes", () => {
      const reviewerVotes = [
        { model: "opus", passed: true, reason: "Good" },
      ];

      render(
        <ExpandablePoint
          point={basePoint}
          labelText="Check"
          secondary={null}
          reviewerVotes={reviewerVotes}
        />
      );

      const button = screen.getByRole("button");
      expect(button).toHaveAttribute("aria-expanded", "false");
      expect(button).toHaveAttribute("tabIndex", "0");

      // Expand it
      fireEvent.click(button);
      expect(button).toHaveAttribute("aria-expanded", "true");
    });

    it("supports keyboard navigation (Enter and Space keys)", () => {
      const reviewerVotes = [
        { model: "opus", passed: true, reason: "Good" },
      ];

      render(
        <ExpandablePoint
          point={basePoint}
          labelText="Check"
          secondary={null}
          reviewerVotes={reviewerVotes}
        />
      );

      const button = screen.getByRole("button");

      // Initially collapsed
      expect(screen.queryByText(/opus:/)).not.toBeInTheDocument();

      // Press Enter to expand
      fireEvent.keyDown(button, { key: "Enter" });
      expect(screen.getByText(/opus:/)).toBeInTheDocument();

      // Press Space to collapse
      fireEvent.keyDown(button, { key: " " });
      expect(screen.queryByText(/opus:/)).not.toBeInTheDocument();
    });

    it("has sufficient color contrast for criterion text", () => {
      const reviewerVotes = [
        { model: "opus", passed: true, reason: "Good" },
      ];

      const { container } = render(
        <ExpandablePoint
          point={basePoint}
          labelText="Check"
          secondary={null}
          reviewerVotes={reviewerVotes}
        />
      );

      // Check that the label has proper contrast class
      const labelSpan = container.querySelector('[class*="text-white"]');
      expect(labelSpan).toBeInTheDocument();
    });
  });
});
