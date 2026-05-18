package review

import (
	"fmt"
	"strings"
)

// criteriaStringToChecks converts a legacy string-based criteria to []ReviewCheck.
// This is a temporary shim during migration. Each numbered line becomes one check.
// Non-numbered lines are treated as preamble for the next check.
func criteriaStringToChecks(criteria string) []ReviewCheck {
	if strings.TrimSpace(criteria) == "" {
		return nil
	}
	
	var checks []ReviewCheck
	var preambleLines []string
	checkNum := 1
	
	for _, line := range strings.Split(criteria, "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		
		// Check if line starts with number (e.g., "1. ", "2. ")
		numbered := false
		for i := 1; i <= 100; i++ {
			prefix := fmt.Sprintf("%d. ", i)
			if strings.HasPrefix(trimmed, prefix) {
				text := strings.TrimSpace(strings.TrimPrefix(trimmed, prefix))
				if text != "" {
					preamble := strings.TrimSpace(strings.Join(preambleLines, "\n"))
					checks = append(checks, ReviewCheck{
						ID:       fmt.Sprintf("check_%d", checkNum),
						Text:     text,
						Preamble: preamble,
					})
					checkNum++
					preambleLines = nil
				}
				numbered = true
				break
			}
		}
		
		// Bullet list (- item)
		if !numbered && strings.HasPrefix(trimmed, "- ") {
			text := strings.TrimSpace(strings.TrimPrefix(trimmed, "- "))
			if text != "" {
				preamble := strings.TrimSpace(strings.Join(preambleLines, "\n"))
				checks = append(checks, ReviewCheck{
					ID:       fmt.Sprintf("check_%d", checkNum),
					Text:     text,
					Preamble: preamble,
				})
				checkNum++
				preambleLines = nil
			}
			numbered = true
		}
		
		// Non-numbered line: accumulate as preamble
		if !numbered {
			preambleLines = append(preambleLines, line)
		}
	}
	
	// If no checks were found, treat entire text as one check
	if len(checks) == 0 {
		checks = append(checks, ReviewCheck{
			ID:   "check_1",
			Text: strings.TrimSpace(criteria),
		})
	}
	
	return checks
}

// BuildReviewPrompt constructs a structured review prompt for the LLM-as-judge.
// It includes the original prompt, agent output files, optional reference answer,
// and evaluation criteria. The reviewer evaluates ONLY against the provided criteria
// — no general rubric is injected.
//
// This id-aware version accepts []ReviewCheck with stable check_N IDs and renders
// them in the format "check_1: <text>". The reviewer must return matching IDs.
func BuildReviewPrompt(originalPrompt string, generatedFiles map[string]string, referenceFiles map[string]string, checks []ReviewCheck, artifact *GeneratorArtifact) string {
	var b strings.Builder

	b.WriteString("You are evaluating another AI agent's work. The agent was given the prompt below ")
	b.WriteString("and asked to produce output. Review the agent output against the original prompt ")
	b.WriteString("and the evaluation criteria provided.\n\n")

	b.WriteString("## Original Prompt\n\n")
	b.WriteString(originalPrompt)
	b.WriteString("\n\n")

	if len(checks) > 0 {
		b.WriteString("## Evaluation Criteria\n\n")
		b.WriteString("You MUST return one judgment per check id below. No extras, no missing entries.\n\n")
		
		// Group checks by preamble to avoid repeating context
		lastPreamble := ""
		for _, c := range checks {
			if c.Preamble != "" && c.Preamble != lastPreamble {
				fmt.Fprintf(&b, "%s\n\n", c.Preamble)
				lastPreamble = c.Preamble
			}
			fmt.Fprintf(&b, "- %s: %s\n", c.ID, c.Text)
		}
		b.WriteString("\n")
	}

	// If files were generated, show them; otherwise indicate no files
	if len(generatedFiles) > 0 {
		b.WriteString("## Generated Code\n\n")
		for name, content := range generatedFiles {
			fmt.Fprintf(&b, "### %s\n```\n%s\n```\n\n", name, content)
		}
	} else {
		b.WriteString("## Generated Code\n\n")
		b.WriteString("(No files were created by the agent.)\n\n")
	}

	// Agent's final response — shown when available
	if artifact != nil && artifact.FinalResponse != "" {
		b.WriteString("## Agent's Final Response\n\n")
		if len(generatedFiles) == 0 {
			b.WriteString("The agent did not create files, but provided this response:\n\n")
		}
		b.WriteString(artifact.FinalResponse)
		b.WriteString("\n\n")
	}

	if len(referenceFiles) > 0 {
		b.WriteString("## Reference Answer\n\n")
		for name, content := range referenceFiles {
			fmt.Fprintf(&b, "### %s\n```\n%s\n```\n\n", name, content)
		}
	} else {
		b.WriteString("## Reference Answer\n\nNo reference answer provided.\n\n")
	}

	b.WriteString("## Scoring Instructions\n\n")
	b.WriteString("For EACH criterion listed above, determine:\n")
	b.WriteString("- **passed**: true if the criterion is fully met, false otherwise\n")
	b.WriteString("- **reasoning**: brief explanation of why it passed or failed\n\n")
	if len(generatedFiles) == 0 && artifact != nil && artifact.FinalResponse != "" {
		b.WriteString("Note: Since no files were created, evaluate the agent's response text against the criteria.\n\n")
	}
	b.WriteString("The overall score = number of passed criteria out of total criteria.\n\n")
	
	b.WriteString("Respond with ONLY a JSON object, no markdown fencing, no explanation.\n")
	b.WriteString("Use this EXACT schema:\n\n")
	
	// Build example with actual IDs
	var exampleIDs []string
	for i, c := range checks {
		exampleIDs = append(exampleIDs, fmt.Sprintf(`{"id":"%s","passed":true,"reasoning":"..."}`, c.ID))
		if i >= 2 {
			break // Show up to 3 examples
		}
	}
	if len(exampleIDs) == 0 {
		exampleIDs = []string{`{"id":"check_1","passed":true,"reasoning":"..."}`}
	}
	
	fmt.Fprintf(&b, `{"criteria":[%s],"summary":"...","issues":["..."],"strengths":["..."]}`, strings.Join(exampleIDs, ","))
	b.WriteString("\n\n")
	
	b.WriteString("Rules:\n")
	if len(checks) > 0 {
		fmt.Fprintf(&b, "- Return exactly one entry per id listed above: %s.\n", formatIDList(checks))
		fmt.Fprintf(&b, "- The \"id\" field MUST be one of: %s.\n", formatIDList(checks))
	}
	b.WriteString("- Do NOT invent ids. Do NOT omit ids. Do NOT echo the criterion text.\n")

	return b.String()
}

// formatIDList formats check IDs as a comma-separated list for the prompt.
func formatIDList(checks []ReviewCheck) string {
	if len(checks) == 0 {
		return "(none)"
	}
	ids := make([]string, len(checks))
	for i, c := range checks {
		ids[i] = c.ID
	}
	return strings.Join(ids, ", ")
}
