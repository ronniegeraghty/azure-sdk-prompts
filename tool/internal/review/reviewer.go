package review

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	copilot "github.com/github/copilot-sdk/go"
)

// Reviewer runs LLM-as-judge code reviews via a separate Copilot session.
type Reviewer interface {
	Review(ctx context.Context, originalPrompt string, workDir string, referenceDir string) (*ReviewResult, error)
}

// CopilotReviewer uses a Copilot session to perform code reviews.
type CopilotReviewer struct {
	client *copilot.Client
	model  string
}

// NewCopilotReviewer creates a reviewer backed by the given Copilot client.
func NewCopilotReviewer(client *copilot.Client, model string) *CopilotReviewer {
	if model == "" {
		model = "claude-sonnet-4.5"
	}
	return &CopilotReviewer{client: client, model: model}
}

// Review creates a separate Copilot session, sends the review prompt, and parses results.
func (r *CopilotReviewer) Review(ctx context.Context, originalPrompt string, workDir string, referenceDir string) (*ReviewResult, error) {
	generatedFiles, err := readDirFiles(workDir)
	if err != nil {
		return nil, fmt.Errorf("reading generated files: %w", err)
	}
	if len(generatedFiles) == 0 {
		return nil, fmt.Errorf("no generated files found in %s", workDir)
	}

	var referenceFiles map[string]string
	if referenceDir != "" {
		referenceFiles, err = readDirFiles(referenceDir)
		if err != nil {
			// Non-fatal: proceed without reference
			referenceFiles = nil
		}
	}

	reviewPrompt := BuildReviewPrompt(originalPrompt, generatedFiles, referenceFiles)

	session, err := r.client.CreateSession(ctx, &copilot.SessionConfig{
		Model: r.model,
		SystemMessage: &copilot.SystemMessageConfig{
			Mode:    "append",
			Content: "You are a code review judge. Respond with ONLY valid JSON. No markdown, no explanation.",
		},
		WorkingDirectory:    workDir,
		OnPermissionRequest: copilot.PermissionHandler.ApproveAll,
	})
	if err != nil {
		return nil, fmt.Errorf("creating review session: %w", err)
	}
	defer session.Disconnect()

	// Capture the assistant's response
	var assistantContent strings.Builder
	var mu sync.Mutex
	unsub := session.On(func(event copilot.SessionEvent) {
		if event.Type == copilot.SessionEventTypeAssistantMessage && event.Data.Content != nil {
			mu.Lock()
			assistantContent.WriteString(*event.Data.Content)
			mu.Unlock()
		}
	})
	defer unsub()

	_, err = session.SendAndWait(ctx, copilot.MessageOptions{
		Prompt: reviewPrompt,
	})
	if err != nil {
		return nil, fmt.Errorf("review session send: %w", err)
	}

	mu.Lock()
	responseText := assistantContent.String()
	mu.Unlock()

	return parseReviewResponse(responseText)
}

// StubReviewer returns placeholder review results for testing.
type StubReviewer struct{}

// Review returns a stub review result.
func (s *StubReviewer) Review(_ context.Context, _ string, _ string, _ string) (*ReviewResult, error) {
	return &ReviewResult{
		Scores: ReviewScores{
			Correctness:   0,
			Completeness:  0,
			BestPractices: 0,
			ErrorHandling: 0,
			PackageUsage:  0,
			CodeQuality:   0,
		},
		OverallScore: 0,
		Summary:      "Review skipped (stub evaluator)",
		Issues:       []string{},
		Strengths:    []string{},
	}, nil
}

// parseReviewResponse extracts the JSON ReviewResult from the LLM response.
func parseReviewResponse(text string) (*ReviewResult, error) {
	// Try to find JSON in the response (LLM may wrap it in markdown fences)
	jsonStr := extractJSON(text)
	if jsonStr == "" {
		return nil, fmt.Errorf("no JSON found in review response: %.200s", text)
	}

	var result ReviewResult
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		return nil, fmt.Errorf("parsing review JSON: %w (response: %.200s)", err, jsonStr)
	}
	return &result, nil
}

// extractJSON finds the first JSON object in the text.
func extractJSON(text string) string {
	// Strip markdown code fences if present
	text = strings.TrimSpace(text)
	if strings.HasPrefix(text, "```json") {
		text = strings.TrimPrefix(text, "```json")
		if idx := strings.LastIndex(text, "```"); idx >= 0 {
			text = text[:idx]
		}
	} else if strings.HasPrefix(text, "```") {
		text = strings.TrimPrefix(text, "```")
		if idx := strings.LastIndex(text, "```"); idx >= 0 {
			text = text[:idx]
		}
	}
	text = strings.TrimSpace(text)

	// Find the first { and last }
	start := strings.Index(text, "{")
	end := strings.LastIndex(text, "}")
	if start >= 0 && end > start {
		return text[start : end+1]
	}
	return ""
}

// readDirFiles reads all files in a directory (non-recursive, skipping hidden/binary).
func readDirFiles(dir string) (map[string]string, error) {
	files := make(map[string]string)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // skip unreadable
		}
		if info.IsDir() {
			name := info.Name()
			if strings.HasPrefix(name, ".") && path != dir {
				return filepath.SkipDir
			}
			return nil
		}
		// Skip binary/large files
		if info.Size() > 1<<20 { // 1MB
			return nil
		}
		rel, err := filepath.Rel(dir, path)
		if err != nil {
			return nil
		}
		if strings.HasPrefix(filepath.Base(rel), ".") {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		files[rel] = string(data)
		return nil
	})
	return files, err
}
