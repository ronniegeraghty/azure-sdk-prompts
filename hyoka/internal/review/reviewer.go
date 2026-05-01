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

	reviewPrompt := BuildReviewPrompt(originalPrompt, generatedFiles, referenceFiles, evaluationCriteria, artifact)

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
		ConfigDir:           configDir,
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

	result, err := parseReviewResponse(responseText)
	if err != nil {
		slog.Error("Failed to parse review response", "model", r.model, "error", err)
		return nil, err
	}

	// Validate response; if invalid, retry up to 2 times
	if errs := validateReviewerResponse(result); len(errs) > 0 {
		slog.Warn("Review response validation failed", "model", r.model, "errors", errs)
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

// parseReviewResponse extracts the JSON ReviewResult from the LLM response.
// It supports both the new ReviewerResponse schema (flat criteria array) and
// the legacy nested scores.criteria format for backward compatibility.
func parseReviewResponse(text string) (*ReviewResult, error) {
	// Try to find JSON in the response (LLM may wrap it in markdown fences)
	jsonStr := utils.ExtractJSON(text)
	if jsonStr == "" {
		return nil, fmt.Errorf("no JSON found in review response: %.200s", text)
	}

	// First, try the new ReviewerResponse schema (flat criteria array)
	var newResp struct {
		Criteria  []CriterionJudgment `json:"criteria"`
		Summary   string              `json:"summary"`
		Issues    []string            `json:"issues"`
		Strengths []string            `json:"strengths"`
	}
	if err := json.Unmarshal([]byte(jsonStr), &newResp); err == nil && len(newResp.Criteria) > 0 {
		criteria := make([]CriterionResult, len(newResp.Criteria))
		for i, c := range newResp.Criteria {
			criteria[i] = CriterionResult{
				Name:   c.Criterion,
				Passed: c.Passed,
				Reason: c.Reasoning,
			}
		}
		scores := ReviewScores{Criteria: criteria}
		return &ReviewResult{
			Scores:       scores,
			OverallScore: scores.PassedCount(),
			MaxScore:     scores.TotalCount(),
			Summary:      newResp.Summary,
			Issues:       newResp.Issues,
			Strengths:    newResp.Strengths,
		}, nil
	}

	// Fall back to legacy format
	var result ReviewResult
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return nil, fmt.Errorf("parsing review JSON: %w (response: %.200s)", err, jsonStr)
	}
	// Ensure MaxScore and OverallScore are consistent with criteria
	if result.MaxScore == 0 && len(result.Scores.Criteria) > 0 {
		result.MaxScore = result.Scores.TotalCount()
	}
	if result.OverallScore == 0 && len(result.Scores.Criteria) > 0 {
		result.OverallScore = result.Scores.PassedCount()
	}
	return &result, nil
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

	reviewPrompt := BuildReviewPrompt(originalPrompt, generatedFiles, referenceFiles, evaluationCriteria, artifact)

	// Run reviewers sequentially — one Copilot session at a time
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
		result, reviewErr := p.runSingleReview(ctx, model, reviewPrompt, modelWorkDir)
		if result != nil {
			result.Model = model
		}
		if reviewErr != nil {
			slog.Warn("Panel reviewer failed", "model", model, "error", reviewErr)
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
	consolidated = deterministicVote(panel)
	consolidated.Model = "consensus"
	slog.Info("Panel review complete", "panel_size", len(panel), "consensus_score", consolidated.OverallScore, "max_score", consolidated.MaxScore)

	return panel, consolidated, nil
}

// Review implements the Reviewer interface using the panel (for backward compat).
func (p *PanelReviewer) Review(ctx context.Context, originalPrompt string, workDir string, referenceDir string, evaluationCriteria string, artifact *GeneratorArtifact) (*ReviewResult, error) {
	_, consolidated, err := p.ReviewPanel(ctx, originalPrompt, workDir, referenceDir, evaluationCriteria, artifact)
	return consolidated, err
}

// runSingleReview creates a Copilot client, runs a review session, and returns the result.
func (p *PanelReviewer) runSingleReview(ctx context.Context, model string, reviewPrompt string, workDir string) (*ReviewResult, error) {
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
		ConfigDir:           configDir,
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

	// Send initial review prompt, then validate and retry up to 2 times
	const maxRetries = 2
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
		result, err = parseReviewResponse(responseText)
		if err != nil {
			if attempt < maxRetries {
				currentPrompt = fmt.Sprintf("Your previous response was not valid JSON. Error: %v\n\nPlease respond again with ONLY a valid JSON object matching the required schema.", err)
				continue
			}
			return nil, err
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

// consolidate uses the first model to synthesize all individual reviews into a consensus.
func (p *PanelReviewer) consolidate(ctx context.Context, originalPrompt string, generatedFiles map[string]string, panel []ReviewResult) (*ReviewResult, error) {
	consolidatorModel := p.models[0]
	slog.Debug("Starting consolidation", "consolidator_model", consolidatorModel, "panel_size", len(panel))

	prompt := buildConsolidationPrompt(originalPrompt, panel)

	slog.Debug("Sending consolidation prompt", "consolidator_model", consolidatorModel)
	result, err := p.runSingleReview(ctx, consolidatorModel, prompt, "")
	if err != nil {
		return nil, fmt.Errorf("consolidation failed: %w", err)
	}
	slog.Debug("Consolidation complete", "overall_score", result.OverallScore, "max_score", result.MaxScore)
	return result, nil
}

// averageReview computes deterministic voting across a panel.
// For each criterion, it FAILS if ANY reviewer marked it failed (strict voting).
// This ensures no false passes when reviewers disagree.
func averageReview(panel []ReviewResult) *ReviewResult {
	if len(panel) == 0 {
		return &ReviewResult{Summary: "No reviews to consolidate"}
	}

	// Collect all criteria by name, track fail counts
	type criterionAgg struct {
		failCount int
		total     int
		reasons   []string
	}
	criteriaMap := make(map[string]*criterionAgg)
	var criteriaOrder []string

	for _, r := range panel {
		for _, c := range r.Scores.Criteria {
			agg, exists := criteriaMap[c.Name]
			if !exists {
				agg = &criterionAgg{}
				criteriaMap[c.Name] = agg
				criteriaOrder = append(criteriaOrder, c.Name)
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

	// Build consensus criteria — fails if ANY reviewer failed it
	var criteria []CriterionResult
	passedCount := 0
	for _, name := range criteriaOrder {
		agg := criteriaMap[name]
		passed := agg.failCount == 0 // strict: any fail = fail
		reason := fmt.Sprintf("%d/%d reviewers passed", agg.total-agg.failCount, agg.total)
		criteria = append(criteria, CriterionResult{
			Name:   name,
			Passed: passed,
			Reason: reason,
		})
		if passed {
			passedCount++
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
func deterministicVote(panel []ReviewResult) *ReviewResult {
	return averageReview(panel)
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
