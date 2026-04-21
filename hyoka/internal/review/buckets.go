// Multi-bucket reviewer: runs one Copilot session per (panel-model, bucket)
// so isolation requested via --review-mode isolated translates into actual
// session-level separation rather than just text concatenation.
package review

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/utils"
)

// Bucket is a single unit of review work — one set of criteria text rendered
// for one Copilot session. Bucket names are used to disambiguate criterion
// results when multiple buckets are merged together.
type Bucket struct {
	Name     string
	Criteria string
}

// MultiBucketReviewer is implemented by reviewers that can split a review
// across multiple sessions, one per bucket. Implementations may merge
// per-bucket results into a single consolidated ReviewResult.
type MultiBucketReviewer interface {
	ReviewBuckets(ctx context.Context, originalPrompt, workDir, referenceDir string, buckets []Bucket) (*ReviewResult, error)
}

// MultiBucketPanelReviewer is implemented by panel reviewers that can split
// each panel-model's review across multiple buckets and return both the
// per-model panel and the consensus ReviewResult.
type MultiBucketPanelReviewer interface {
	ReviewPanelBuckets(ctx context.Context, originalPrompt, workDir, referenceDir string, buckets []Bucket) (panel []ReviewResult, consolidated *ReviewResult, err error)
}

// ReviewBuckets runs one Copilot session per bucket for this CopilotReviewer's
// single model and merges the per-bucket criterion results into one
// ReviewResult. With a single bucket, behavior is identical to Review.
func (r *CopilotReviewer) ReviewBuckets(ctx context.Context, originalPrompt, workDir, referenceDir string, buckets []Bucket) (*ReviewResult, error) {
	if len(buckets) == 0 {
		return nil, fmt.Errorf("ReviewBuckets called with no buckets")
	}
	if len(buckets) == 1 {
		return r.Review(ctx, originalPrompt, workDir, referenceDir, buckets[0].Criteria)
	}
	results := make([]bucketResult, 0, len(buckets))
	for _, b := range buckets {
		if ctx.Err() != nil {
			break
		}
		slog.Info("CopilotReviewer bucket starting", "model", r.model, "bucket", b.Name)
		res, err := r.Review(ctx, originalPrompt, workDir, referenceDir, b.Criteria)
		if err != nil {
			slog.Warn("CopilotReviewer bucket failed", "model", r.model, "bucket", b.Name, "error", err)
			continue
		}
		results = append(results, bucketResult{name: b.Name, result: res})
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("all buckets failed for model %s", r.model)
	}
	merged := mergeBucketResults(results)
	merged.Model = r.model
	return merged, nil
}

// ReviewPanelBuckets runs each panel model against every bucket and merges
// per-model bucket results into a single ReviewResult per model. The panel is
// then consolidated via deterministic any-fail voting (matching ReviewPanel).
func (p *PanelReviewer) ReviewPanelBuckets(ctx context.Context, originalPrompt, workDir, referenceDir string, buckets []Bucket) ([]ReviewResult, *ReviewResult, error) {
	if len(p.models) == 0 {
		return nil, nil, fmt.Errorf("no reviewer models configured")
	}
	if len(buckets) == 0 {
		return nil, nil, fmt.Errorf("ReviewPanelBuckets called with no buckets")
	}
	if len(buckets) == 1 {
		return p.ReviewPanel(ctx, originalPrompt, workDir, referenceDir, buckets[0].Criteria)
	}

	generatedFiles, err := utils.ReadDirFiles(workDir)
	if err != nil || len(generatedFiles) == 0 {
		return nil, nil, fmt.Errorf("no generated files to review in %s", workDir)
	}
	var referenceFiles map[string]string
	if referenceDir != "" {
		var rerr error
		referenceFiles, rerr = utils.ReadDirFiles(referenceDir)
		if rerr != nil {
			slog.Warn("Failed to read reference files", "dir", referenceDir, "error", rerr)
		}
	}

	slog.Info("Starting bucketed panel review",
		"models", p.models, "bucket_count", len(buckets))

	var panel []ReviewResult
	for _, model := range p.models {
		if ctx.Err() != nil {
			break
		}
		modelWorkDir, copyErr := copyDirToTemp(workDir, fmt.Sprintf("hyoka-review-%s-*", sanitize(model)))
		if copyErr != nil {
			slog.Warn("Failed to create workspace copy for reviewer", "model", model, "error", copyErr)
			modelWorkDir = workDir
		} else {
			defer os.RemoveAll(modelWorkDir)
		}

		results := make([]bucketResult, 0, len(buckets))
		for _, b := range buckets {
			if ctx.Err() != nil {
				break
			}
			slog.Debug("Bucket review starting", "model", model, "bucket", b.Name)
			reviewPrompt := BuildReviewPrompt(originalPrompt, generatedFiles, referenceFiles, b.Criteria)
			res, rerr := p.runSingleReview(ctx, model, reviewPrompt, modelWorkDir)
			if rerr != nil {
				slog.Warn("Bucket review failed", "model", model, "bucket", b.Name, "error", rerr)
				continue
			}
			results = append(results, bucketResult{name: b.Name, result: res})
		}
		if len(results) == 0 {
			slog.Warn("All buckets failed for model", "model", model)
			continue
		}
		merged := mergeBucketResults(results)
		merged.Model = model
		panel = append(panel, *merged)
	}
	if len(panel) == 0 {
		return nil, nil, fmt.Errorf("all reviewers failed across all buckets")
	}
	consolidated := deterministicVote(panel)
	consolidated.Model = "consensus"
	slog.Info("Bucketed panel review complete",
		"panel_size", len(panel),
		"buckets", len(buckets),
		"consensus_score", consolidated.OverallScore,
		"max_score", consolidated.MaxScore)
	return panel, consolidated, nil
}

type bucketResult struct {
	name   string
	result *ReviewResult
}

// mergeBucketResults combines multiple per-bucket review results into one
// ReviewResult. Criterion names from non-"combined" buckets are prefixed with
// "[bucket-name] " to keep them distinguishable when the deterministic vote
// runs across panel models. Issues, strengths, and summaries are concatenated.
func mergeBucketResults(parts []bucketResult) *ReviewResult {
	merged := &ReviewResult{}
	var summaries []string
	var allEvents []ReviewEvent

	for _, p := range parts {
		r := p.result
		if r == nil {
			continue
		}
		for _, c := range r.Scores.Criteria {
			cc := c
			if p.name != "" && p.name != "combined" {
				cc.Name = fmt.Sprintf("[%s] %s", p.name, c.Name)
			}
			merged.Scores.Criteria = append(merged.Scores.Criteria, cc)
		}
		merged.OverallScore += r.OverallScore
		merged.MaxScore += r.MaxScore
		if s := strings.TrimSpace(r.Summary); s != "" {
			if p.name != "" && p.name != "combined" {
				summaries = append(summaries, fmt.Sprintf("[%s] %s", p.name, s))
			} else {
				summaries = append(summaries, s)
			}
		}
		merged.Issues = append(merged.Issues, r.Issues...)
		merged.Strengths = append(merged.Strengths, r.Strengths...)
		allEvents = append(allEvents, r.Events...)
	}
	merged.Summary = strings.Join(summaries, "\n\n")
	merged.Events = allEvents
	return merged
}

// sanitize makes a string safe for use in temp-dir patterns.
func sanitize(s string) string {
	return strings.NewReplacer("/", "_", " ", "_", ":", "_").Replace(s)
}
