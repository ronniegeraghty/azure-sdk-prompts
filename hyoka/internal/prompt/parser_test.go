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
if len(result) != 3 {
t.Fatalf("expected 3 entries, got %d", len(result))
}
if result[0].Prompt != "Maven dependency for azure-storage-blob" {
t.Errorf("entry[0].Prompt = %q, want %q", result[0].Prompt, "Maven dependency for azure-storage-blob")
}
if result[1].Prompt != "BlobServiceClientBuilder with DefaultAzureCredential" {
t.Errorf("entry[1].Prompt = %q", result[1].Prompt)
}
if result[2].Prompt != "BlobClient.uploadFromFile() and downloadToFile()" {
t.Errorf("entry[2].Prompt = %q", result[2].Prompt)
}
for i, e := range result {
if len(e.SubPoints) != 0 {
t.Errorf("entry[%d] should have no sub-points, got %v", i, e.SubPoints)
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
if len(result) != 2 {
t.Fatalf("expected 2 entries, got %d", len(result))
}
if result[0].Prompt != "CRUD operations" {
t.Errorf("entry[0].Prompt = %q", result[0].Prompt)
}
if len(result[0].SubPoints) != 3 {
t.Fatalf("expected 3 sub-points for entry[0], got %d", len(result[0].SubPoints))
}
expected := []string{"createItem()", "readItem()", "deleteItem()"}
for i, sp := range result[0].SubPoints {
if sp != expected[i] {
t.Errorf("entry[0].SubPoints[%d] = %q, want %q", i, sp, expected[i])
}
}
if result[1].Prompt != "Error handling" {
t.Errorf("entry[1].Prompt = %q", result[1].Prompt)
}
if len(result[1].SubPoints) != 1 {
t.Fatalf("expected 1 sub-point for entry[1], got %d", len(result[1].SubPoints))
}
if result[1].SubPoints[0] != "CosmosException handling with status codes" {
t.Errorf("entry[1].SubPoints[0] = %q", result[1].SubPoints[0])
}
}

func TestParseEvaluationCriteria_WithPreamble(t *testing.T) {
input := `The generated code should include:
- azure-identity dependency
- DefaultAzureCredential usage
- Exception handling for AuthenticationException`

result := ParseEvaluationCriteria(input)
if len(result) != 3 {
t.Fatalf("expected 3 entries (preamble skipped), got %d", len(result))
}
if result[0].Prompt != "azure-identity dependency" {
t.Errorf("entry[0].Prompt = %q", result[0].Prompt)
}
}

func TestParseEvaluationCriteria_TabSubPoints(t *testing.T) {
input := "- Parent item\n\t- Sub item one\n\t- Sub item two"

result := ParseEvaluationCriteria(input)
if len(result) != 1 {
t.Fatalf("expected 1 entry, got %d", len(result))
}
if len(result[0].SubPoints) != 2 {
t.Fatalf("expected 2 tab-indented sub-points, got %d", len(result[0].SubPoints))
}
if result[0].SubPoints[0] != "Sub item one" {
t.Errorf("sub[0] = %q", result[0].SubPoints[0])
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
if len(p.ParsedCriteria) != 3 {
t.Fatalf("expected 3 parsed criteria, got %d", len(p.ParsedCriteria))
}
if p.ParsedCriteria[0].Prompt != "Use of DefaultAzureCredential" {
t.Errorf("criteria[0].Prompt = %q", p.ParsedCriteria[0].Prompt)
}
if len(p.ParsedCriteria[0].SubPoints) != 2 {
t.Errorf("criteria[0] should have 2 sub-points, got %d", len(p.ParsedCriteria[0].SubPoints))
}
if p.ParsedCriteria[1].Prompt != "CRUD operations on secrets" {
t.Errorf("criteria[1].Prompt = %q", p.ParsedCriteria[1].Prompt)
}
if len(p.ParsedCriteria[2].SubPoints) != 0 {
t.Errorf("criteria[2] should have no sub-points")
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
{Prompt: "Use DefaultAzureCredential"},
{Prompt: "Handle authentication errors"},
{Prompt: "Use async methods"},
}
result := FormatParsedCriteria(entries)
expected := "1. Use DefaultAzureCredential\n2. Handle authentication errors\n3. Use async methods"
if result != expected {
t.Errorf("FormatParsedCriteria() =\n%q\n\nwant:\n%q", result, expected)
}
}

func TestFormatParsedCriteria_WithSubPoints(t *testing.T) {
entries := []CriterionEntry{
{
Prompt: "CRUD operations",
SubPoints: []string{
"createItem()",
"readItem()",
"deleteItem()",
},
},
{
Prompt:    "Error handling",
SubPoints: []string{"CosmosException handling"},
},
}
result := FormatParsedCriteria(entries)
expected := `1. CRUD operations
   - createItem()
   - readItem()
   - deleteItem()
2. Error handling
   - CosmosException handling`
if result != expected {
t.Errorf("FormatParsedCriteria() =\n%q\n\nwant:\n%q", result, expected)
}
}

func TestFormatParsedCriteria_SkipsEmptyEntries(t *testing.T) {
entries := []CriterionEntry{
{Prompt: "Use DefaultAzureCredential"},
{Prompt: ""},
{Prompt: "  \t  "},
{Prompt: "Handle authentication errors"},
}
result := FormatParsedCriteria(entries)
// Empty entries should be skipped but numbering should be preserved
expected := "1. Use DefaultAzureCredential\n4. Handle authentication errors"
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

if len(parsed) != 3 {
t.Fatalf("expected 3 parsed entries, got %d", len(parsed))
}

expected := `1. Use of DefaultAzureCredential
   - Must include azure-identity dependency
   - Must use DefaultAzureCredentialBuilder
2. CRUD operations on secrets
   - setSecret()
   - getSecret()
3. Error handling for authentication failures`

if result != expected {
t.Errorf("FormatParsedCriteria() =\n%q\n\nwant:\n%q", result, expected)
}
}

func TestFormatParsedCriteria_NumberedListsInInput(t *testing.T) {
// Test that numbered lists in markdown input are converted to bullet format
input := `- Authentication
  1. Use DefaultAzureCredential
  2. Handle AuthenticationException
- Storage operations
  1. Upload blob
  2. Download blob`

parsed := ParseEvaluationCriteria(input)
// The parser currently treats "1. Use..." as sub-points since they don't start with "- "
// This tests current behavior - if parser changes, update this test
if len(parsed) != 2 {
t.Fatalf("expected 2 parsed entries, got %d", len(parsed))
}
}

func TestFormatParsedCriteria_FourSpaceIndent(t *testing.T) {
// Test 4-space indented sub-points (currently dropped by parser)
input := `- Parent item
    - Four space sub-point
  - Two space sub-point`

parsed := ParseEvaluationCriteria(input)
_ = FormatParsedCriteria(parsed) // Just verify it does not panic

// Currently parser only recognizes 2-space indent as sub-points
// This documents current behavior - 4-space indent is silently dropped
if len(parsed) != 1 {
t.Fatalf("expected 1 parsed entry, got %d", len(parsed))
}
if len(parsed[0].SubPoints) != 1 {
t.Errorf("expected 1 sub-point (4-space indent dropped), got %d: %v",
len(parsed[0].SubPoints), parsed[0].SubPoints)
}
}
