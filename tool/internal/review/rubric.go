package review

import (
	"fmt"
	"strings"
)

// BuildReviewPrompt constructs a structured review prompt for the LLM-as-judge.
// It includes the original prompt, generated code files, and optionally a reference answer.
func BuildReviewPrompt(originalPrompt string, generatedFiles map[string]string, referenceFiles map[string]string) string {
	var b strings.Builder

	b.WriteString("You are a senior Azure SDK code reviewer. Review the generated code below.\n\n")

	b.WriteString("## Original Prompt\n\n")
	b.WriteString(originalPrompt)
	b.WriteString("\n\n")

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

	b.WriteString(scoringRubric(len(referenceFiles) > 0))
	return b.String()
}

func scoringRubric(hasReference bool) string {
	refLine := ""
	if hasReference {
		refLine = `7. **Reference Similarity** — How similar is it to the reference? (1-10)`
	} else {
		refLine = `7. **Reference Similarity** — Skip (no reference provided), output 0`
	}

	return fmt.Sprintf(`## Scoring Rubric

Score each dimension from 1-10:

1. **Correctness** — Does the code correctly implement what was asked?
2. **Completeness** — Are all requirements addressed? Missing features?
3. **Best Practices** — Does it follow Azure SDK best practices? (DefaultAzureCredential, proper disposal, async patterns, etc.)
4. **Error Handling** — Are errors handled properly? Retries? Timeouts?
5. **Package Usage** — Are the correct and latest SDK packages used?
6. **Code Quality** — Clean, readable, well-structured code?
%s

## Output Format

Respond with ONLY a JSON object, no markdown fencing, no explanation:
{"scores":{"correctness":N,"completeness":N,"best_practices":N,"error_handling":N,"package_usage":N,"code_quality":N,"reference_similarity":N},"overall_score":N,"summary":"...","issues":["..."],"strengths":["..."]}
`, refLine)
}
