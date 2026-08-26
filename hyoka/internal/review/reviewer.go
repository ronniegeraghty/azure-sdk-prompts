// Package review provides code review functionality using Copilot.
package review

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"

	copilot "github.com/github/copilot-sdk/go"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/artifact"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/copilotperm"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/utils"
)

// GeneratorArtifact is a type alias for the generator artifact type.
// This allows review to access artifact.GeneratorArtifact without creating an import cycle.
type GeneratorArtifact = artifact.GeneratorArtifact

// Reviewer runs LLM-as-judge code reviews via a separate Copilot session.
type Reviewer interface {
	Review(ctx context.Context, originalPrompt string, workDir string, referenceDir string, evaluationCriteria string, artifact *GeneratorArtifact) (*ReviewResult, error)
}

// CopilotReviewer uses a Copilot session to perform code reviews.
type CopilotReviewer struct {
	client            *copilot.Client
	model             string
	maxSessionActions int
	skillDirectories  []string
	sessionTimeout    time.Duration
	systemPrompt      string
}

// NewCopilotReviewer creates a reviewer backed by the given Copilot client.
func NewCopilotReviewer(client *copilot.Client, model string, maxSessionActions int) *CopilotReviewer {
	if model == "" {
		model = "claude-sonnet-4.5"
	}
	return &CopilotReviewer{client: client, model: model, maxSessionActions: maxSessionActions}
}

// SetSkillDirectories configures skill directories for the review session.
func (r *CopilotReviewer) SetSkillDirectories(dirs []string) {
	r.skillDirectories = dirs
}

// SetSessionTimeout configures the maximum duration for a single review
// SendAndWait call. Zero means use the default (10 minutes).
func (r *CopilotReviewer) SetSessionTimeout(d time.Duration) {
	r.sessionTimeout = d
}

// SetSystemPrompt configures a custom system prompt for the review session.
// An empty string means no system prompt is sent.
func (r *CopilotReviewer) SetSystemPrompt(prompt string) {
	r.systemPrompt = prompt
}

// Review creates a separate Copilot session, sends the review prompt, and parses results.
func (r *CopilotReviewer) Review(ctx context.Context, originalPrompt string, workDir string, referenceDir string, evaluationCriteria string, artifact *GeneratorArtifact) (*ReviewResult, error) {
	slog.Debug("Reading generated files for review", "workDir", workDir)
	generatedFiles, err := utils.ReadDirFiles(workDir)
	if err != nil {
		return nil, fmt.Errorf("reading generated files: %w", err)
	}

	// Empty workspace is acceptable if we have an artifact with a response
	if len(generatedFiles) == 0 {
		if artifact == nil || artifact.FinalResponse == "" {
			return nil, fmt.Errorf("no generated files found in %s and no agent response to review", workDir)
		}
		slog.Debug("No generated files, reviewing agent's final response only")
	} else {
		slog.Debug("Generated files loaded", "file_count", len(generatedFiles))
	}

	var referenceFiles map[string]string
	if referenceDir != "" {
		referenceFiles, err = utils.ReadDirFiles(referenceDir)
		if err != nil {
			// Non-fatal: proceed without reference
			slog.Warn("Could not read reference files, proceeding without", "referenceDir", referenceDir, "error", err)
			referenceFiles = nil
		}
	}

	checks := criteriaStringToChecks(evaluationCriteria)
	reviewPrompt := BuildReviewPrompt(originalPrompt, generatedFiles, referenceFiles, checks, artifact)

	// Create isolated config directory to prevent user-level skills from
	// leaking into the review session (#21).
	configDir, err := os.MkdirTemp("", "hyoka-config-*")
	if err != nil {
		return nil, fmt.Errorf("creating isolated config dir: %w", err)
	}
	defer os.RemoveAll(configDir)

	reviewCtx, reviewCancel := context.WithCancel(ctx)
	defer reviewCancel()

	// Capture the assistant's response and all session events
	collector := newEventCollector(r.model, r.maxSessionActions, reviewCancel)

	slog.Info("Starting review session", "model", r.model, "skills", len(r.skillDirectories), "work_dir", workDir)
	slog.Debug("Creating review session", "model", r.model)
	sessionCfg := &copilot.SessionConfig{
		Model:               r.model,
		ConfigDirectory:     configDir,
		WorkingDirectory:    workDir,
		OnPermissionRequest: copilotperm.ApproveAll,
		SkillDirectories:    r.skillDirectories,
		OnEvent:             collector.handleEvent,
	}
	if r.systemPrompt != "" {
		sessionCfg.SystemMessage = &copilot.SystemMessageConfig{
			Mode:    "append",
			Content: r.systemPrompt,
		}
	}
	session, err := r.client.CreateSession(reviewCtx, sessionCfg)
	if err != nil {
		slog.Error("Failed to create review session", "model", r.model, "error", err)
		return nil, fmt.Errorf("creating review session: %w", err)
	}
	// Clean up session state (#62). DeleteSession removes session-state dir
	// and SQLite entry while client is still connected. Then Disconnect
	// releases in-memory resources.
	defer func() {
		deleteCtx, deleteCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer deleteCancel()
		if err := r.client.DeleteSession(deleteCtx, session.SessionID); err != nil {
			slog.Debug("review session delete failed", "sessionID", session.SessionID, "error", err)
		}
		done := make(chan struct{})
		go func() { session.Disconnect(); close(done) }()
		select {
		case <-done:
		case <-time.After(15 * time.Second):
		}
	}()

	// Apply an explicit deadline so the SDK does not fall back to its
	// hard-coded 60-second default (see copilot-sdk session.go).
	reviewTimeout := 10 * time.Minute
	if r.sessionTimeout > 0 {
		reviewTimeout = r.sessionTimeout
	}
	sendCtx, sendCancel := context.WithTimeout(reviewCtx, reviewTimeout)
	defer sendCancel()

	slog.Debug("Sending review prompt", "model", r.model, "timeout", reviewTimeout, "length", len(reviewPrompt))
	_, err = session.SendAndWait(sendCtx, copilot.MessageOptions{
		Prompt: reviewPrompt,
	})
	if err != nil {
		slog.Error("Review session send failed", "model", r.model, "error", err)
		return nil, fmt.Errorf("review session send: %w", err)
	}

	responseText, capturedEvents := collector.response()

	result, validationErrors := parseReviewResponseV2(responseText, checks)
	if len(validationErrors) > 0 {
		slog.Error("Failed to parse review response", "model", r.model, "errors", validationErrors)
		return nil, fmt.Errorf("parsing review response: %s", strings.Join(validationErrors, "; "))
	}

	result.Events = capturedEvents
	slog.Info("Review complete", "model", r.model, "overall_score", result.OverallScore, "max_score", result.MaxScore)
	return result, nil
}

// StubReviewer returns placeholder review results for testing.
type StubReviewer struct{}

// Review returns a stub review result.
func (s *StubReviewer) Review(_ context.Context, _ string, _ string, _ string, _ string, _ *GeneratorArtifact) (*ReviewResult, error) {
	return &ReviewResult{
		Scores: ReviewScores{
			Criteria: []CriterionResult{
				{Name: "stub_criterion", Passed: true, Reason: "stub mode"},
			},
		},
		OverallScore: 1,
		MaxScore:     1,
		Summary:      "Review skipped (stub evaluator)",
		Issues:       []string{},
		Strengths:    []string{},
	}, nil
}

// ReviewBuckets returns a stub review result with one criterion per bucket so
// StubReviewer satisfies MultiBucketReviewer for tests.
func (s *StubReviewer) ReviewBuckets(_ context.Context, _ string, _ string, _ string, buckets []Bucket, _ *GeneratorArtifact) (*ReviewResult, error) {
	criteria := make([]CriterionResult, 0, len(buckets))
	for _, b := range buckets {
		criteria = append(criteria, CriterionResult{
			Name: "stub_criterion_" + b.Name, Passed: true, Reason: "stub mode",
		})
	}
	if len(criteria) == 0 {
		criteria = append(criteria, CriterionResult{Name: "stub_criterion", Passed: true, Reason: "stub mode"})
	}
	return &ReviewResult{
		Scores:       ReviewScores{Criteria: criteria},
		OverallScore: len(criteria),
		MaxScore:     len(criteria),
		Summary:      "Review skipped (stub evaluator, bucketed)",
		Issues:       []string{},
		Strengths:    []string{},
	}, nil
}

// parseReviewResponseV2 parses an id-aware review response and validates against expected checks.
// Returns the populated ReviewResult with canonical labels from expected, plus any validation errors.
// The Name field of each CriterionResult is set to the canonical text from expected (indexed by ID).
func parseReviewResponseV2(text string, expected []ReviewCheck) (*ReviewResult, []string) {
	jsonStr := utils.ExtractJSON(text)
	if jsonStr == "" {
		return nil, []string{fmt.Sprintf("no JSON found in review response: %.200s", text)}
	}

	var resp struct {
		Criteria  []CriterionJudgment `json:"criteria"`
		Summary   string              `json:"summary"`
		Issues    []string            `json:"issues"`
		Strengths []string            `json:"strengths"`
	}
	if err := json.Unmarshal([]byte(jsonStr), &resp); err != nil {
		return nil, []string{fmt.Sprintf("parsing review JSON: %v (response: %.200s)", err, jsonStr)}
	}

	// Build expected id → text lookup
	expectedIDs := make(map[string]string)
	for _, c := range expected {
		expectedIDs[c.ID] = c.Text
	}

	// Validate: returned IDs must match expected IDs exactly
	returnedIDs := make(map[string]bool)
	var validationErrors []string
	for _, c := range resp.Criteria {
		if c.ID == "" {
			validationErrors = append(validationErrors, "found criterion with empty id")
			continue
		}
		returnedIDs[c.ID] = true
		if _, ok := expectedIDs[c.ID]; !ok {
			validationErrors = append(validationErrors, fmt.Sprintf("unexpected id: %s", c.ID))
		}
		if c.Reasoning == "" {
			validationErrors = append(validationErrors, fmt.Sprintf("id %s has empty reasoning", c.ID))
		}
	}

	// Check for missing IDs
	var missing []string
	for id := range expectedIDs {
		if !returnedIDs[id] {
			missing = append(missing, id)
		}
	}
	if len(missing) > 0 {
		validationErrors = append(validationErrors, fmt.Sprintf("missing ids: %v", missing))
	}

	// If validation failed, return errors
	if len(validationErrors) > 0 {
		return nil, validationErrors
	}

	// Populate criteria with canonical labels and stable IDs
	criteria := make([]CriterionResult, len(resp.Criteria))
	for i, c := range resp.Criteria {
		criteria[i] = CriterionResult{
			ID:     c.ID,              // stable id for vote keying
			Name:   expectedIDs[c.ID], // canonical label from YAML
			Passed: c.Passed,
			Reason: c.Reasoning,
		}
	}

	scores := ReviewScores{Criteria: criteria}
	return &ReviewResult{
		Scores:       scores,
		OverallScore: scores.PassedCount(),
		MaxScore:     scores.TotalCount(),
		Summary:      resp.Summary,
		Issues:       resp.Issues,
		Strengths:    resp.Strengths,
	}, nil
}

// validateReviewerResponse checks that a parsed response contains valid criteria.
// Returns a list of validation errors; nil means valid.
func validateReviewerResponse(result *ReviewResult) []string {
	var errs []string
	if result == nil {
		return []string{"nil review result"}
	}
	if len(result.Scores.Criteria) == 0 {
		errs = append(errs, "no criteria in response")
	}
	for i, c := range result.Scores.Criteria {
		if c.Name == "" {
			errs = append(errs, fmt.Sprintf("criterion %d has empty name", i))
		}
	}
	return errs
}

// PanelReviewer runs multiple reviewers in parallel and consolidates results.
type PanelReviewer struct {
	clientOpts        *copilot.ClientOptions
	models            []string // first model is the consolidator
	maxSessionActions int
	skillDirectories  []string
	sessionTimeout    time.Duration
	systemPrompt      string
}

// NewPanelReviewer creates a panel reviewer that runs multiple models concurrently.
// The first model in the list is used as the consolidator.
func NewPanelReviewer(clientOpts *copilot.ClientOptions, models []string, maxSessionActions int) *PanelReviewer {
	return &PanelReviewer{
		clientOpts:        clientOpts,
		models:            models,
		maxSessionActions: maxSessionActions,
	}
}

// SetSkillDirectories configures skill directories for all review sessions.
func (p *PanelReviewer) SetSkillDirectories(dirs []string) {
	p.skillDirectories = dirs
}

// SetSessionTimeout configures the maximum duration for a single review
// SendAndWait call. Zero means use the default (10 minutes).
func (p *PanelReviewer) SetSessionTimeout(d time.Duration) {
	p.sessionTimeout = d
}

// SetSystemPrompt configures a custom system prompt for all review sessions.
// An empty string means no system prompt is sent.
func (p *PanelReviewer) SetSystemPrompt(prompt string) {
	p.systemPrompt = prompt
}

// Models returns the list of reviewer models.
func (p *PanelReviewer) Models() []string {
	return p.models
}

// ReviewPanel runs all reviewer models sequentially and returns individual results
// plus a consolidated result. The consolidated result is produced by the first model
// in the list, which receives all other reviewers' outputs.
// Reviews run one at a time so each Copilot session starts, completes, and stops
// before the next begins, reducing peak memory usage.
func (p *PanelReviewer) ReviewPanel(ctx context.Context, originalPrompt string, workDir string, referenceDir string, evaluationCriteria string, artifact *GeneratorArtifact) (panel []ReviewResult, consolidated *ReviewResult, err error) {
	slog.Info("Starting sequential panel review", "model_count", len(p.models), "models", p.models)
	if len(p.models) == 0 {
		return nil, nil, fmt.Errorf("no reviewer models configured")
	}

	generatedFiles, err := utils.ReadDirFiles(workDir)
	if err != nil || len(generatedFiles) == 0 {
		// Empty workspace is acceptable if we have an artifact with a response
		if artifact == nil || artifact.FinalResponse == "" {
			return nil, nil, fmt.Errorf("no generated files to review in %s and no agent response to review", workDir)
		}
		slog.Debug("No generated files, reviewing agent's final response only")
	}

	var referenceFiles map[string]string
	if referenceDir != "" {
		var readErr error
		referenceFiles, readErr = utils.ReadDirFiles(referenceDir)
		if readErr != nil {
			slog.Warn("Failed to read reference files", "dir", referenceDir, "error", readErr)
		}
	}

	checks := criteriaStringToChecks(evaluationCriteria)
	reviewPrompt := BuildReviewPrompt(originalPrompt, generatedFiles, referenceFiles, checks, artifact)

	// Run reviewers sequentially — one Copilot session at a time
	var skipped []SkippedReviewer
	for i, model := range p.models {
		// Bail early if the parent context was cancelled (#129).
		if ctx.Err() != nil {
			break
		}
		slog.Debug("Panel reviewer starting", "model", model, "index", i)
		modelWorkDir, copyErr := copyDirToTemp(workDir, fmt.Sprintf("hyoka-review-%s-*", model))
		if copyErr != nil {
			slog.Warn("Failed to create workspace copy for reviewer", "model", model, "error", copyErr)
			modelWorkDir = workDir
		} else {
			defer os.RemoveAll(modelWorkDir)
		}
		result, reviewErr := p.runSingleReview(ctx, model, reviewPrompt, modelWorkDir, checks)
		if result != nil {
			result.Model = model
		}
		if reviewErr != nil {
			slog.Warn("Panel reviewer failed", "model", model, "error", reviewErr)
			skipped = append(skipped, SkippedReviewer{Model: model, Error: reviewErr.Error()})
			continue
		}
		slog.Debug("Panel reviewer complete", "model", model, "overall_score", result.OverallScore, "max_score", result.MaxScore)
		panel = append(panel, *result)
	}

	if len(panel) == 0 {
		return nil, nil, fmt.Errorf("all reviewers failed")
	}

	// Deterministic multi-model voting: for each criterion, if ANY reviewer
	// says it failed, mark it as failed. No AI consolidation needed.
	slog.Info("Computing deterministic consensus (any-fail voting)", "panel_size", len(panel))
	consolidated = deterministicVote(panel, checks)
	consolidated.Model = "consensus"
	consolidated.SkippedReviewers = skipped
	slog.Info("Panel review complete", "panel_size", len(panel), "consensus_score", consolidated.OverallScore, "max_score", consolidated.MaxScore)

	return panel, consolidated, nil
}

// Review implements the Reviewer interface using the panel (for backward compat).
func (p *PanelReviewer) Review(ctx context.Context, originalPrompt string, workDir string, referenceDir string, evaluationCriteria string, artifact *GeneratorArtifact) (*ReviewResult, error) {
	_, consolidated, err := p.ReviewPanel(ctx, originalPrompt, workDir, referenceDir, evaluationCriteria, artifact)
	return consolidated, err
}

// runSingleReview creates a Copilot client, runs a review session, and returns the result.
// If checks are provided, uses id-aware parser with retry-on-validation-error.
func (p *PanelReviewer) runSingleReview(ctx context.Context, model string, reviewPrompt string, workDir string, checks []ReviewCheck) (*ReviewResult, error) {
	slog.Debug("Starting single review", "model", model)
	opts := *p.clientOpts
	client := copilot.NewClient(&opts)

	if err := client.Start(ctx); err != nil {
		return nil, fmt.Errorf("starting reviewer client for %s: %w", model, err)
	}
	var panelSessionID string
	defer func() {
		// Delete session state before stopping client (#62)
		if panelSessionID != "" {
			deleteCtx, deleteCancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer deleteCancel()
			if err := client.DeleteSession(deleteCtx, panelSessionID); err != nil {
				slog.Debug("panel review session delete failed",
					"sessionID", panelSessionID, "model", model, "error", err)
			}
		}
		done := make(chan struct{})
		go func() { client.Stop(); close(done) }()
		select {
		case <-done:
		case <-time.After(10 * time.Second):
			client.ForceStop()
		}
	}()

	// Create isolated config directory to prevent user-level skills from
	// leaking into the review session (#21).
	configDir, err := os.MkdirTemp("", "hyoka-config-*")
	if err != nil {
		return nil, fmt.Errorf("creating isolated config dir for %s: %w", model, err)
	}
	defer os.RemoveAll(configDir)

	reviewCtx, reviewCancel := context.WithCancel(ctx)
	defer reviewCancel()

	collector := newEventCollector(model, p.maxSessionActions, reviewCancel)

	slog.Info("Starting review session", "model", model, "skills", len(p.skillDirectories), "work_dir", workDir)
	slog.Debug("Creating review session", "model", model)
	sessionCfg := &copilot.SessionConfig{
		Model:               model,
		ConfigDirectory:     configDir,
		WorkingDirectory:    workDir,
		OnPermissionRequest: copilotperm.ApproveAll,
		SkillDirectories:    p.skillDirectories,
		OnEvent:             collector.handleEvent,
	}
	if p.systemPrompt != "" {
		sessionCfg.SystemMessage = &copilot.SystemMessageConfig{
			Mode:    "append",
			Content: p.systemPrompt,
		}
	}
	session, err := client.CreateSession(reviewCtx, sessionCfg)
	if err != nil {
		return nil, fmt.Errorf("creating review session for %s: %w", model, err)
	}
	panelSessionID = session.SessionID

	// Apply an explicit deadline so the SDK does not fall back to its
	// hard-coded 60-second default (see copilot-sdk session.go).
	panelTimeout := 10 * time.Minute
	if p.sessionTimeout > 0 {
		panelTimeout = p.sessionTimeout
	}
	sendCtx, sendCancel := context.WithTimeout(reviewCtx, panelTimeout)
	defer sendCancel()

	slog.Debug("Sending review prompt", "model", model, "timeout", panelTimeout, "length", len(reviewPrompt))

	// Send initial review prompt, then validate and retry up to 3 times
	const maxRetries = 3
	var result *ReviewResult
	currentPrompt := reviewPrompt

	for attempt := 0; attempt <= maxRetries; attempt++ {
		if attempt > 0 {
			slog.Info("Retrying review with validation feedback", "model", model, "attempt", attempt)
		}

		_, err = session.SendAndWait(sendCtx, copilot.MessageOptions{
			Prompt: currentPrompt,
		})
		if err != nil {
			return nil, fmt.Errorf("review session send for %s: %w", model, err)
		}

		responseText, _ := collector.response()

		// Use id-aware parser if checks provided, otherwise fall back to legacy
		if len(checks) > 0 {
			var validationErrors []string
			result, validationErrors = parseReviewResponseV2(responseText, checks)
			if len(validationErrors) > 0 {
				if attempt < maxRetries {
					// Build precise retry prompt with missing/extra id details
					missing := []string{}
					extra := []string{}
					expectedIDs := make(map[string]bool)
					for _, c := range checks {
						expectedIDs[c.ID] = true
					}

					// Parse to extract returned IDs for better error message
					jsonStr := utils.ExtractJSON(responseText)
					if jsonStr != "" {
						var resp struct {
							Criteria []struct {
								ID string `json:"id"`
							} `json:"criteria"`
						}
						if json.Unmarshal([]byte(jsonStr), &resp) == nil {
							returnedIDs := make(map[string]bool)
							for _, c := range resp.Criteria {
								returnedIDs[c.ID] = true
								if !expectedIDs[c.ID] {
									extra = append(extra, c.ID)
								}
							}
							for id := range expectedIDs {
								if !returnedIDs[id] {
									missing = append(missing, id)
								}
							}
						}
					}

					var retryMsg strings.Builder
					retryMsg.WriteString("Your response has validation errors:\n")
					for _, e := range validationErrors {
						fmt.Fprintf(&retryMsg, "- %s\n", e)
					}
					if len(missing) > 0 {
						fmt.Fprintf(&retryMsg, "\nYou MUST include these missing check IDs in your response: %v\n", missing)
					}
					if len(extra) > 0 {
						fmt.Fprintf(&retryMsg, "You included unexpected IDs: %v\n", extra)
					}
					fmt.Fprintf(&retryMsg, "\nPlease return a COMPLETE response with exactly these IDs: %s\n", formatIDList(checks))
					currentPrompt = retryMsg.String()
					slog.Debug("Retry prompt prepared", "model", model, "attempt", attempt, "missing", missing, "extra", extra)
					continue
				}
				// Max retries exceeded: synthesize failing checks for missing IDs instead of dropping reviewer
				slog.Warn("reviewer failed to return valid criteria after 3 retries, synthesizing failures for missing checks",
					"model", model,
					"validation_errors", validationErrors)

				// Determine which IDs are missing
				expectedIDs := make(map[string]bool)
				for _, c := range checks {
					expectedIDs[c.ID] = true
				}
				returnedIDs := make(map[string]bool)
				if result != nil {
					for _, c := range result.Scores.Criteria {
						returnedIDs[c.ID] = true
					}
				}

				// Synthesize failures for missing IDs
				if result == nil {
					result = &ReviewResult{
						Model: model,
						Scores: ReviewScores{
							Criteria: []CriterionResult{},
						},
					}
				}
				for _, check := range checks {
					if !returnedIDs[check.ID] {
						result.Scores.Criteria = append(result.Scores.Criteria, CriterionResult{
							ID:     check.ID,
							Name:   check.Text,
							Passed: false,
							Reason: "reviewer failed to return a vote after 3 attempts",
						})
					}
				}
			}
		}

		if errs := validateReviewerResponse(result); len(errs) > 0 {
			if attempt < maxRetries {
				currentPrompt = fmt.Sprintf("Your response had validation errors: %s\n\nPlease respond again with ONLY a valid JSON object. Every criterion must have a non-empty name.", strings.Join(errs, "; "))
				continue
			}
			slog.Warn("Review response validation failed after retries", "model", model, "errors", errs)
		}
		break
	}

	_, capturedEvents := collector.response()
	result.Events = capturedEvents
	return result, nil
}

// averageReview computes deterministic voting across a panel.
// For each criterion, it FAILS if ANY reviewer marked it failed (strict voting).
// This ensures no false passes when reviewers disagree.
// Criteria are keyed by ID (bucket::check_id for bucketed, check_id for combined).
func averageReview(panel []ReviewResult, expected []ReviewCheck) *ReviewResult {
	if len(panel) == 0 {
		return &ReviewResult{Summary: "No reviews to consolidate"}
	}

	// Build id → canonical label lookup from expected
	expectedLabels := make(map[string]string)
	for _, c := range expected {
		expectedLabels[c.ID] = c.Text
	}

	// Collect all criteria by stable ID, track fail counts
	type criterionAgg struct {
		id         string
		bucketName string // extracted from prefixed Name
		label      string // canonical label from expected
		failCount  int
		total      int
		reasons    []string
	}
	criteriaMap := make(map[string]*criterionAgg)
	var observedOrder []string // for legacy path when expected is empty

	// First pass: collect votes from all reviewers
	for _, r := range panel {
		for _, c := range r.Scores.Criteria {
			// Extract bucket name if present in Name
			bucketName := ""
			if strings.HasPrefix(c.Name, "[") {
				closeIdx := strings.Index(c.Name, "]")
				if closeIdx > 0 {
					bucketName = c.Name[1:closeIdx]
				}
			}

			// Build vote key: bucket::id or just id
			var voteKey string
			if bucketName != "" && c.ID != "" {
				voteKey = bucketName + "::" + c.ID
			} else if c.ID != "" {
				voteKey = c.ID
			} else {
				// Fallback for legacy path: key by name
				voteKey = c.Name
			}

			agg, exists := criteriaMap[voteKey]
			if !exists {
				label := c.Name // default to display name
				if c.ID != "" && expectedLabels[c.ID] != "" {
					label = expectedLabels[c.ID]
					// Prefix with bucket if applicable
					if bucketName != "" {
						label = fmt.Sprintf("[%s] %s", bucketName, label)
					}
				}
				agg = &criterionAgg{
					id:         c.ID,
					bucketName: bucketName,
					label:      label,
				}
				criteriaMap[voteKey] = agg
				observedOrder = append(observedOrder, voteKey)
			}
			agg.total++
			if !c.Passed {
				agg.failCount++
			}
			if c.Reason != "" {
				agg.reasons = append(agg.reasons, c.Reason)
			}
		}
	}

	// Build consensus criteria
	var criteria []CriterionResult
	passedCount := 0

	if len(expected) > 0 {
		// New path: anchor to expected checks (in expected order)
		// This ensures MaxScore is deterministic even if reviewers skip checks
		for _, exp := range expected {
			// Check for votes on this expected ID (may be bucketed or not)
			foundVotes := false

			// First, check for non-bucketed vote (direct ID match)
			if agg, hasVotes := criteriaMap[exp.ID]; hasVotes {
				foundVotes = true
				passed := agg.failCount == 0 // strict: any fail = fail
				reason := fmt.Sprintf("%d/%d reviewers passed", agg.total-agg.failCount, agg.total)
				criteria = append(criteria, CriterionResult{
					ID:     agg.id,
					Name:   agg.label,
					Passed: passed,
					Reason: reason,
				})
				if passed {
					passedCount++
				}
			}

			// Then, check for bucketed votes (bucket::ID pattern)
			// This handles cases where the same check appears in multiple buckets
			for voteKey, agg := range criteriaMap {
				// Extract ID from bucket::id pattern
				var checkID string
				if strings.Contains(voteKey, "::") {
					parts := strings.SplitN(voteKey, "::", 2)
					if len(parts) == 2 {
						checkID = parts[1]
					}
				}

				// If this vote matches the expected ID and wasn't already counted
				if checkID == exp.ID && voteKey != exp.ID {
					foundVotes = true
					passed := agg.failCount == 0 // strict: any fail = fail
					reason := fmt.Sprintf("%d/%d reviewers passed", agg.total-agg.failCount, agg.total)
					criteria = append(criteria, CriterionResult{
						ID:     agg.id,
						Name:   agg.label,
						Passed: passed,
						Reason: reason,
					})
					if passed {
						passedCount++
					}
				}
			}

			if !foundVotes {
				// No reviewer voted on this check → mark as failed
				criteria = append(criteria, CriterionResult{
					ID:     exp.ID,
					Name:   exp.Text,
					Passed: false,
					Reason: "no reviewer returned a vote for this check",
				})
			}
		}
	} else {
		// Legacy path: walk observed votes (when expected is empty)
		for _, voteKey := range observedOrder {
			agg := criteriaMap[voteKey]
			passed := agg.failCount == 0 // strict: any fail = fail
			reason := fmt.Sprintf("%d/%d reviewers passed", agg.total-agg.failCount, agg.total)
			criteria = append(criteria, CriterionResult{
				ID:     agg.id,
				Name:   agg.label, // canonical label (bucket-prefixed if applicable)
				Passed: passed,
				Reason: reason,
			})
			if passed {
				passedCount++
			}
		}
	}

	// Merge issues and strengths
	issueSet := make(map[string]bool)
	var issues []string
	strengthSet := make(map[string]bool)
	var strengths []string
	for _, r := range panel {
		for _, iss := range r.Issues {
			if !issueSet[iss] {
				issueSet[iss] = true
				issues = append(issues, iss)
			}
		}
		for _, s := range r.Strengths {
			if !strengthSet[s] {
				strengthSet[s] = true
				strengths = append(strengths, s)
			}
		}
	}

	return &ReviewResult{
		Model: "consensus (strict-vote)",
		Scores: ReviewScores{
			Criteria: criteria,
		},
		OverallScore: passedCount,
		MaxScore:     len(criteria),
		Summary:      fmt.Sprintf("Strict consensus from %d reviewers: %d/%d reviewer checks passed (any-fail voting)", len(panel), passedCount, len(criteria)),
		Issues:       issues,
		Strengths:    strengths,
	}
}

// deterministicVote computes a consensus result using strict any-fail voting.
// For each criterion, if ANY reviewer says it failed, the criterion fails.
// This replaces AI consolidation with deterministic, reproducible logic.
func deterministicVote(panel []ReviewResult, expected []ReviewCheck) *ReviewResult {
	return averageReview(panel, expected)
}

func copyDirToTemp(src string, pattern string) (string, error) {
	dst, err := os.MkdirTemp("", pattern)
	if err != nil {
		return "", fmt.Errorf("creating temp dir: %w", err)
	}
	err = filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			name := d.Name()
			if strings.HasPrefix(name, ".") && path != src {
				return filepath.SkipDir
			}
			if utils.IsDefaultExcludedDir(name) {
				return filepath.SkipDir
			}
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(target, 0755)
		}
		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()
		dstFile, err := os.Create(target)
		if err != nil {
			return err
		}
		defer dstFile.Close()
		_, err = io.Copy(dstFile, srcFile)
		return err
	})
	if err != nil {
		os.RemoveAll(dst)
		return "", err
	}
	return dst, nil
}
