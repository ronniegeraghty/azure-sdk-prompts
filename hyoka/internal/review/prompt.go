package review

import (
	"fmt"
	"strings"
)

// BuildReviewPrompt constructs a structured review prompt for the LLM-as-judge.
// It includes the original prompt, generated code files, optional reference answer,
// and evaluation criteria. The reviewer evaluates ONLY against the provided criteria
// — no general rubric is injected.
func BuildReviewPrompt(originalPrompt string, generatedFiles map[string]string, referenceFiles map[string]string, evaluationCriteria string) string {
	var b strings.Builder

	b.WriteString("You are evaluating another AI agent's work. The agent was given the prompt below ")
	b.WriteString("and asked to produce code. Review the generated code against the original prompt ")
	b.WriteString("and the evaluation criteria provided.\n\n")

	b.WriteString("## Original Prompt\n\n")
	b.WriteString(originalPrompt)
	b.WriteString("\n\n")

	if evaluationCriteria != "" {
		b.WriteString("## Evaluation Criteria\n\n")
		b.WriteString("Evaluate EACH criterion individually as pass/fail:\n\n")
		b.WriteString(evaluationCriteria)
		b.WriteString("\n\n")
	}

	b.WriteString("## Generated Code\n\n")
	for name, content := range generatedFiles {
		fmt.Fprintf(&b, "### %s\n```\n%s\n```\n\n", name, content)
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
	b.WriteString("The overall score = number of passed criteria out of total criteria.\n\n")
	b.WriteString("Respond with ONLY a JSON object, no markdown fencing, no explanation.\n")
	b.WriteString("Use this EXACT schema:\n\n")
	b.WriteString(`{"criteria":[{"criterion":"original criteria text","passed":true,"reasoning":"brief explanation"}],"summary":"...","issues":["..."],"strengths":["..."]}`)
	b.WriteString("\n\nEach criterion from the Evaluation Criteria section above MUST appear exactly once in the criteria array.\n")
	b.WriteString("Use the original criterion text as the \"criterion\" field value.\n")

	return b.String()
}
