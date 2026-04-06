package report

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/ronniegeraghty/hyoka/internal/review"
)

const (
	// MaxReportBytes is the threshold (50 MB) above which verbose fields are
	// truncated before writing. This prevents unbounded memory growth when
	// evaluating large prompt suites.
	MaxReportBytes int64 = 50 * 1024 * 1024

	// truncatedMessage is appended to fields that were shortened.
	truncatedMessage = "... [truncated — exceeded memory bounds]"

	// maxFieldBytes caps individual verbose fields during truncation.
	maxFieldBytes = 64 * 1024 // 64 KB per field
)

// estimateSize returns a rough byte count for the JSON encoding of v.
// It uses json.Marshal which is wasteful for very large objects, but the
// caller only invokes this once per report so it is acceptable.
func estimateSize(v any) int64 {
	data, err := json.Marshal(v)
	if err != nil {
		return 0
	}
	return int64(len(data))
}

// TruncateReport checks the estimated serialized size of an EvalReport and,
// if it exceeds MaxReportBytes, truncates verbose fields (session event
// content, tool results, reviewed file contents, error details, and review
// event payloads) to bring it under budget. It returns true when truncation
// was applied.
func TruncateReport(r *EvalReport) bool {
	return truncateReportWithLimit(r, MaxReportBytes)
}

// truncateReportWithLimit is the internal implementation that accepts a
// configurable limit, used by tests.
func truncateReportWithLimit(r *EvalReport, limit int64) bool {
	size := estimateSize(r)
	if size <= limit {
		return false
	}

	slog.Warn("Report exceeds memory bound — truncating verbose fields",
		"prompt_id", r.PromptID,
		"config", r.ConfigName,
		"estimated_bytes", size,
		"limit_bytes", limit,
	)

	truncated := false

	// 1. Session events — Content, ToolResult, ToolArgs tend to be large.
	for i := range r.SessionEvents {
		ev := &r.SessionEvents[i]
		if truncateField(&ev.Content) {
			truncated = true
		}
		if truncateField(&ev.ToolResult) {
			truncated = true
		}
		if truncateField(&ev.ToolArgs) {
			truncated = true
		}
	}

	// 2. Reviewed file contents.
	for i := range r.ReviewedFiles {
		if truncateField(&r.ReviewedFiles[i].Content) {
			truncated = true
		}
	}

	// 3. Error details.
	if truncateField(&r.ErrorDetails) {
		truncated = true
	}

	// 4. Review events (both consolidated and panel).
	truncateReviewEvents := func(rr *review.ReviewResult) {
		if rr == nil {
			return
		}
		for i := range rr.Events {
			ev := &rr.Events[i]
			if truncateField(&ev.Content) {
				truncated = true
			}
			if truncateField(&ev.Result) {
				truncated = true
			}
			if truncateField(&ev.ToolArgs) {
				truncated = true
			}
		}
	}
	truncateReviewEvents(r.Review)
	for i := range r.ReviewPanel {
		truncateReviewEvents(&r.ReviewPanel[i])
	}

	if truncated {
		slog.Warn("Verbose fields truncated",
			"prompt_id", r.PromptID,
			"original_bytes", size,
			"max_field_bytes", maxFieldBytes,
		)
	}
	return truncated
}

// truncateField shortens *s in-place if it exceeds maxFieldBytes.
// Returns true when truncation was applied.
func truncateField(s *string) bool {
	if len(*s) <= maxFieldBytes {
		return false
	}
	// Truncate to maxFieldBytes and append the marker. We cut on a rune
	// boundary by converting to []rune, but for speed we just slice bytes
	// and accept a possible mid-rune cut — the marker makes it clear the
	// content is incomplete.
	*s = (*s)[:maxFieldBytes] + fmt.Sprintf(" %s", truncatedMessage)
	return true
}
