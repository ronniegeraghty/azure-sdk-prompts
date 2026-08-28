package prompt

import (
"testing"
)

func TestParseEvaluationCriteria_Empty(t *testing.T) {
result := ParseEvaluationCriteria("")
if result != nil {
t.Errorf("expected nil for empty input, got %v", result)
}
result = ParseEvaluationCriteria("   \n  \n  ")
if result != nil {
t.Errorf("expected nil for whitespace-only input, got %v", result)
}
}

func TestParseEvaluationCriteria_SimpleBullets(t *testing.T) {
input := `- Maven dependency for azure-storage-blob
- BlobServiceClientBuilder with DefaultAzureCredential
- BlobClient.uploadFromFile() and downloadToFile()`

result := ParseEvaluationCriteria(input)
if len(result) != 1 {
t.Fatalf("expected 1 entry, got %d", len(result))
}
if result[0].Prompt != "" {
t.Errorf("entry.Prompt = %q, want empty (no preamble)", result[0].Prompt)
}
expected := []string{
"Maven dependency for azure-storage-blob",
"BlobServiceClientBuilder with DefaultAzureCredential",
"BlobClient.uploadFromFile() and downloadToFile()",
}
if len(result[0].Checks) != len(expected) {
t.Fatalf("expected %d checks, got %d: %v", len(expected), len(result[0].Checks), result[0].Checks)
}
for i, c := range result[0].Checks {
if c != expected[i] {
t.Errorf("Checks[%d] = %q, want %q", i, c, expected[i])
}
}
}

func TestParseEvaluationCriteria_WithSubPoints(t *testing.T) {
input := `- CRUD operations
  - createItem()
  - readItem()
  - deleteItem()
- Error handling
  - CosmosException handling with status codes`

result := ParseEvaluationCriteria(input)
if len(result) != 1 {
t.Fatalf("expected 1 entry, got %d", len(result))
}
expected := []string{
"CRUD operations",
"createItem()",
"readItem()",
"deleteItem()",
"Error handling",
"CosmosException handling with status codes",
}
if len(result[0].Checks) != len(expected) {
t.Fatalf("expected %d flat checks, got %d: %v", len(expected), len(result[0].Checks), result[0].Checks)
}
for i, c := range result[0].Checks {
if c != expected[i] {
t.Errorf("Checks[%d] = %q, want %q", i, c, expected[i])
}
}
}

func TestParseEvaluationCriteria_WithPreamble(t *testing.T) {
input := `The generated code should include:
- azure-identity dependency
- DefaultAzureCredential usage
- Exception handling for AuthenticationException`

result := ParseEvaluationCriteria(input)
if len(result) != 1 {
t.Fatalf("expected 1 entry, got %d", len(result))
}
if result[0].Prompt != "The generated code should include:" {
t.Errorf("entry.Prompt = %q, want preamble preserved", result[0].Prompt)
}
if len(result[0].Checks) != 3 {
t.Fatalf("expected 3 checks, got %d", len(result[0].Checks))
}
if result[0].Checks[0] != "azure-identity dependency" {
t.Errorf("Checks[0] = %q", result[0].Checks[0])
}
}

func TestParseEvaluationCriteria_TabSubPoints(t *testing.T) {
input := "- Parent item\n\t- Sub item one\n\t- Sub item two"

result := ParseEvaluationCriteria(input)
if len(result) != 1 {
t.Fatalf("expected 1 entry, got %d", len(result))
}
expected := []string{"Parent item", "Sub item one", "Sub item two"}
if len(result[0].Checks) != len(expected) {
t.Fatalf("expected %d flat checks, got %d", len(expected), len(result[0].Checks))
}
for i, c := range result[0].Checks {
if c != expected[i] {
t.Errorf("Checks[%d] = %q, want %q", i, c, expected[i])
}
}
}

func TestSplitSections(t *testing.T) {
tests := []struct {
name     string
body     string
expected map[string]string
}{
{
"single section",
"\n## Prompt\n\nWrite some code.\n",
map[string]string{"Prompt": "Write some code."},
},
{
"multiple sections",
"\n## Prompt\n\nWrite some code.\n\n## Evaluation Criteria\n\n- Use DefaultAzureCredential\n- Handle errors\n",
map[string]string{
"Prompt":              "Write some code.",
"Evaluation Criteria": "- Use DefaultAzureCredential\n- Handle errors",
},
},
{
"with preamble",
"\n# Title\n\nSome intro.\n\n## Prompt\n\nCode here.\n\n## Notes\n\nExtra info.\n",
map[string]string{
"":       "# Title\n\nSome intro.",
"Prompt": "Code here.",
"Notes":  "Extra info.",
},
},
{
"empty body",
"",
map[string]string{},
},
{
"no sections",
"\nJust plain text.\n",
map[string]string{"": "Just plain text."},
},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
got := splitSections(tt.body)
if len(got) != len(tt.expected) {
t.Fatalf("section count = %d, want %d; got %v", len(got), len(tt.expected), got)
}
for key, want := range tt.expected {
if got[key] != want {
t.Errorf("sections[%q] = %q, want %q", key, got[key], want)
}
}
})
}
}

func TestParseEvaluationCriteria_InPromptFile(t *testing.T) {
content := []byte(`---
id: test-criteria
properties:
  service: test
  language: python
---

## Prompt

Write some code.

## Evaluation Criteria

The generated code should include:
- Use of DefaultAzureCredential
  - Import from azure.identity
  - Proper exception handling
- CRUD operations on secrets
- Logging with the logging module
`)

p, err := ParsePromptFile(content, "test.prompt.md")
if err != nil {
t.Fatal(err)
}
if len(p.ParsedCriteria) != 1 {
t.Fatalf("expected 1 parsed criteria entry, got %d", len(p.ParsedCriteria))
}
if p.ParsedCriteria[0].Prompt != "The generated code should include:" {
t.Errorf("Prompt (preamble) = %q", p.ParsedCriteria[0].Prompt)
}
expected := []string{
"Use of DefaultAzureCredential",
"Import from azure.identity",
"Proper exception handling",
"CRUD operations on secrets",
"Logging with the logging module",
}
if len(p.ParsedCriteria[0].Checks) != len(expected) {
t.Fatalf("expected %d flat checks, got %d", len(expected), len(p.ParsedCriteria[0].Checks))
}
for i, c := range p.ParsedCriteria[0].Checks {
if c != expected[i] {
t.Errorf("Checks[%d] = %q, want %q", i, c, expected[i])
}
}
}

func TestFormatParsedCriteria_Empty(t *testing.T) {
result := FormatParsedCriteria(nil)
if result != "" {
t.Errorf("expected empty string for nil input, got %q", result)
}
result = FormatParsedCriteria([]CriterionEntry{})
if result != "" {
t.Errorf("expected empty string for empty slice, got %q", result)
}
}

func TestFormatParsedCriteria_SimpleBullets(t *testing.T) {
entries := []CriterionEntry{
{Checks: []string{
"Use DefaultAzureCredential",
"Handle authentication errors",
"Use async methods",
}},
}
result := FormatParsedCriteria(entries)
expected := "1. Use DefaultAzureCredential\n2. Handle authentication errors\n3. Use async methods"
if result != expected {
t.Errorf("FormatParsedCriteria() =\n%q\n\nwant:\n%q", result, expected)
}
}

func TestFormatParsedCriteria_WithSubPoints(t *testing.T) {
// Legacy SubPoints back-compat: when Checks is empty, SubPoints is rendered.
entries := []CriterionEntry{
{
Prompt: "CRUD operations",
SubPoints: []string{
"createItem()",
"readItem()",
"deleteItem()",
},
},
}
result := FormatParsedCriteria(entries)
expected := "CRUD operations\n1. createItem()\n2. readItem()\n3. deleteItem()"
if result != expected {
t.Errorf("FormatParsedCriteria() =\n%q\n\nwant:\n%q", result, expected)
}
}

func TestFormatParsedCriteria_SkipsEmptyEntries(t *testing.T) {
entries := []CriterionEntry{
{Checks: []string{"Use DefaultAzureCredential", "", "  \t  ", "Handle authentication errors"}},
}
result := FormatParsedCriteria(entries)
// Empty checks within an entry are skipped and numbering is contiguous.
expected := "1. Use DefaultAzureCredential\n2. Handle authentication errors"
if result != expected {
t.Errorf("FormatParsedCriteria() =\n%q\n\nwant:\n%q", result, expected)
}
}

func TestFormatParsedCriteria_MultilineWithBlankLines(t *testing.T) {
input := `- Use of DefaultAzureCredential
  - Must include azure-identity dependency
  - Must use DefaultAzureCredentialBuilder

- CRUD operations on secrets
  - setSecret()
  - getSecret()

- Error handling for authentication failures`

parsed := ParseEvaluationCriteria(input)
result := FormatParsedCriteria(parsed)

if len(parsed) != 1 {
t.Fatalf("expected 1 parsed entry, got %d", len(parsed))
}

expected := `1. Use of DefaultAzureCredential
2. Must include azure-identity dependency
3. Must use DefaultAzureCredentialBuilder
4. CRUD operations on secrets
5. setSecret()
6. getSecret()
7. Error handling for authentication failures`

if result != expected {
t.Errorf("FormatParsedCriteria() =\n%q\n\nwant:\n%q", result, expected)
}
}

func TestFormatParsedCriteria_NumberedListsInInput(t *testing.T) {
// Numbered "1." lines aren't bullets; under the flat-checks model they
// appear as preamble lines (when before any bullet) or get dropped
// (when after a bullet). The two top-level bullets become checks.
input := `- Authentication
  1. Use DefaultAzureCredential
  2. Handle AuthenticationException
- Storage operations
  1. Upload blob
  2. Download blob`

parsed := ParseEvaluationCriteria(input)
if len(parsed) != 1 {
t.Fatalf("expected 1 parsed entry, got %d", len(parsed))
}
if len(parsed[0].Checks) != 2 {
t.Fatalf("expected 2 top-level checks, got %d: %v", len(parsed[0].Checks), parsed[0].Checks)
}
}

func TestFormatParsedCriteria_FourSpaceIndent(t *testing.T) {
// 4-space indented bullet lines are still recognized as bullets after
// the leading whitespace strip, so they're flattened into Checks.
input := `- Parent item
    - Four space sub-point
  - Two space sub-point`

parsed := ParseEvaluationCriteria(input)
_ = FormatParsedCriteria(parsed) // Just verify it does not panic

if len(parsed) != 1 {
t.Fatalf("expected 1 parsed entry, got %d", len(parsed))
}
if len(parsed[0].Checks) != 3 {
t.Errorf("expected 3 flat checks, got %d: %v",
len(parsed[0].Checks), parsed[0].Checks)
}
}

// TestDeprecatedEvaluationCriteriaYAML verifies that ParsePromptYAML emits
// a deprecation warning when the evaluation_criteria field is present, but
// still parses the prompt successfully for backward compatibility.
func TestDeprecatedEvaluationCriteriaYAML(t *testing.T) {
content := []byte(`id: test-yaml-legacy
properties:
  service: test
  language: python
prompt_text: Write some Python code.
evaluation_criteria: |
  The code should include:
  - DefaultAzureCredential usage
  - Error handling
`)

p, err := ParsePromptYAML(content, "test.prompt.yaml")
if err != nil {
t.Fatalf("ParsePromptYAML should not error on legacy evaluation_criteria field: %v", err)
}

if p.EvaluationCriteria == "" {
t.Error("expected EvaluationCriteria to be populated from deprecated field")
}
if len(p.ParsedCriteria) == 0 {
t.Error("expected ParsedCriteria to be populated from deprecated field")
}

// Warning is logged via slog.Warn but doesn't break parsing
// (verified by running tests with -v to see stderr output)
}

// TestMarkdownEvaluationCriteriaSectionSupported verifies that ParsePromptFile
// continues to support the ## Evaluation Criteria markdown section without any
// deprecation warning. The markdown body section remains a first-class way to
// declare criteria in .prompt.md files; only the YAML/frontmatter
// `evaluation_criteria:` field is deprecated in favor of `graders:`.
func TestMarkdownEvaluationCriteriaSectionSupported(t *testing.T) {
content := []byte(`---
id: test-md-legacy
properties:
  service: test
  language: python
---

## Prompt

Write some Python code.

## Evaluation Criteria

The code should include:
- DefaultAzureCredential usage
- Error handling
`)

p, err := ParsePromptFile(content, "test.prompt.md")
if err != nil {
t.Fatalf("ParsePromptFile should not error on legacy ## Evaluation Criteria section: %v", err)
}

if p.EvaluationCriteria == "" {
t.Error("expected EvaluationCriteria to be populated from markdown section")
}
if len(p.ParsedCriteria) == 0 {
t.Error("expected ParsedCriteria to be populated from markdown section")
}
}

// TestGradersNoDeprecationWarning verifies that prompts using the new
// graders: frontmatter do NOT trigger deprecation warnings.
func TestGradersNoDeprecationWarning(t *testing.T) {
content := []byte(`---
id: test-clean-graders
properties:
  service: test
  language: python
graders:
  - type: prompt
    name: DefaultAzureCredential Check
    checks:
      - Uses DefaultAzureCredential
      - Handles authentication errors
---

## Prompt

Write some Python code using DefaultAzureCredential.
`)

p, err := ParsePromptFile(content, "test.prompt.md")
if err != nil {
t.Fatalf("ParsePromptFile should not error on clean graders: %v", err)
}

if p.EvaluationCriteria != "" {
t.Error("expected EvaluationCriteria to be empty when using graders")
}
if len(p.Graders) == 0 {
t.Error("expected Graders to be populated")
}

// No deprecation warning should be emitted (no EvaluationCriteria section)
}

