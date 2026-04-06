// Package graders provides pluggable grader implementations.
package graders

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"

	"github.com/ronniegeraghty/hyoka/internal/utils"
)

// DefaultMaxScore is the default maximum score a prompt grader assigns.
const DefaultMaxScore = 10

// LLMCaller abstracts the LLM invocation so the grader is testable
// without a live Copilot session.
type LLMCaller interface {
	Call(ctx context.Context, model string, prompt string) (string, error)
}


// PromptGradeResult holds the normalized output of a prompt grader.
type PromptGradeResult struct {
	Score   float64             `json:"score"`   // 0.0–1.0
	Passed  bool                `json:"passed"`  // score > 0
	Details PromptGraderDetails `json:"details"`
}

// LLMResponse is the JSON schema the LLM is expected to return.
type LLMResponse struct {
	Score     int    `json:"score"`
	Reasoning string `json:"reasoning"`
}

// PromptGrader wraps a single LLM model + rubric as one grader instance (DM19).
type PromptGrader struct {
	Name     string
	Model    string
	Rubric   string
	MaxScore int
	Caller   LLMCaller
}

// NewPromptGrader creates a PromptGrader from a name and raw config map.
// Required keys: "model" (string), "rubric" (string).
// Optional: "max_score" (int, default 10).
func NewPromptGrader(name string, cfg map[string]any) (*PromptGrader, error) {
	model, ok := cfg["model"].(string)
	if !ok || model == "" {
		return nil, fmt.Errorf("prompt grader %q: model is required", name)
	}

	rubric, ok := cfg["rubric"].(string)
	if !ok || rubric == "" {
		return nil, fmt.Errorf("prompt grader %q: rubric is required", name)
	}

	maxScore := DefaultMaxScore
	if v, exists := cfg["max_score"]; exists {
		switch ms := v.(type) {
		case int:
			maxScore = ms
		case float64:
			maxScore = int(ms)
		default:
			return nil, fmt.Errorf("prompt grader %q: max_score must be an integer", name)
		}
	}
	if maxScore <= 0 {
		return nil, fmt.Errorf("prompt grader %q: max_score must be positive, got %d", name, maxScore)
	}

	return &PromptGrader{
		Name:     name,
		Model:    model,
		Rubric:   rubric,
		MaxScore: maxScore,
	}, nil
}

// BuildPrompt constructs the LLM prompt from the rubric and workspace files.
func (pg *PromptGrader) BuildPrompt(workspaceFiles map[string]string) string {
	var b strings.Builder

	b.WriteString("You are a code quality evaluator. Score the following code on a scale of 0 to ")
	fmt.Fprintf(&b, "%d based on the rubric below.\n\n", pg.MaxScore)

	b.WriteString("## Rubric\n\n")
	b.WriteString(pg.Rubric)
	b.WriteString("\n\n")

	b.WriteString("## Code Under Review\n\n")
	if len(workspaceFiles) == 0 {
		b.WriteString("(no files found)\n\n")
	} else {
		for path, content := range workspaceFiles {
			fmt.Fprintf(&b, "### %s\n```\n%s\n```\n\n", path, content)
		}
	}

	b.WriteString("## Response Format\n\n")
	b.WriteString("Respond with ONLY a JSON object (no markdown fences, no extra text):\n")
	fmt.Fprintf(&b, `{"score": <0-%d>, "reasoning": "<brief explanation>"}`, pg.MaxScore)
	b.WriteString("\n")

	return b.String()
}

// ParseResponse extracts a score and reasoning from the LLM's text response.
func (pg *PromptGrader) ParseResponse(text string) (*LLMResponse, error) {
	jsonStr := utils.ExtractJSON(text)
	if jsonStr == "" {
		return nil, fmt.Errorf("no JSON found in LLM response: %.200s", text)
	}

	var resp LLMResponse
	if err := json.Unmarshal([]byte(jsonStr), &resp); err != nil {
		return nil, fmt.Errorf("parsing LLM response JSON: %w (response: %.200s)", err, jsonStr)
	}

	// Clamp score to valid range.
	if resp.Score < 0 {
		resp.Score = 0
	}
	if resp.Score > pg.MaxScore {
		resp.Score = pg.MaxScore
	}

	return &resp, nil
}

// NormalizeScore converts a raw integer score to the 0.0–1.0 range.
func NormalizeScore(rawScore, maxScore int) float64 {
	if maxScore <= 0 {
		return 0
	}
	score := float64(rawScore) / float64(maxScore)
	if score < 0 {
		return 0
	}
	if score > 1.0 {
		return 1.0
	}
	return score
}

// Grade runs the LLM review and returns a normalized result.
// workDir is the directory of generated code to evaluate.
func (pg *PromptGrader) Grade(ctx context.Context, workDir string) (*PromptGradeResult, error) {
	if pg.Caller == nil {
		return nil, fmt.Errorf("prompt grader %q: no LLM caller configured", pg.Name)
	}

	slog.Debug("Reading workspace files for prompt grader", "grader", pg.Name, "workDir", workDir)
	files, err := utils.ReadDirFiles(workDir)
	if err != nil {
		return nil, fmt.Errorf("prompt grader %q: reading workspace: %w", pg.Name, err)
	}

	prompt := pg.BuildPrompt(files)
	slog.Debug("Sending prompt to LLM", "grader", pg.Name, "model", pg.Model, "prompt_len", len(prompt))

	responseText, err := pg.Caller.Call(ctx, pg.Model, prompt)
	if err != nil {
		return nil, fmt.Errorf("prompt grader %q: LLM call failed: %w", pg.Name, err)
	}

	resp, err := pg.ParseResponse(responseText)
	if err != nil {
		return nil, fmt.Errorf("prompt grader %q: %w", pg.Name, err)
	}

	normalized := NormalizeScore(resp.Score, pg.MaxScore)
	slog.Info("Prompt grader complete",
		"grader", pg.Name,
		"model", pg.Model,
		"raw_score", resp.Score,
		"max_score", pg.MaxScore,
		"normalized", normalized,
	)

	return &PromptGradeResult{
		Score:  normalized,
		Passed: resp.Score > 0,
		Details: PromptGraderDetails{
			Model:     pg.Model,
			Rubric:    pg.Rubric,
			Reasoning: resp.Reasoning,
			RawScore:  resp.Score,
			MaxScore:  pg.MaxScore,
		},
	}, nil
}
