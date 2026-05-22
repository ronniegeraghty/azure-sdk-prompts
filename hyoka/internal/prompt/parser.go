package prompt

import (
"bytes"
"fmt"
"log/slog"
"regexp"
"strings"

"gopkg.in/yaml.v3"

"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria"
)

// sectionHeadingRe matches level-2 markdown headings at the start of a line.
var sectionHeadingRe = regexp.MustCompile(`(?m)^## (.+)$`)

// frontmatter is the YAML structure decoded from prompt file frontmatter.
// All metadata string fields live under the properties: map.
type frontmatter struct {
ID              string            `yaml:"id"`
Tags            []string          `yaml:"tags"`
ProjectContext  map[string]string `yaml:"project_context"`
StarterProject  string            `yaml:"starter_project"`
ReferenceAnswer string            `yaml:"reference_answer"`
Timeout           int               `yaml:"timeout"`
MaxSessionActions int               `yaml:"max_session_actions"`
MaxTurns          int               `yaml:"max_turns"`
ExpectedPkgs      []string          `yaml:"expected_packages"`
ExpectedTools     []string          `yaml:"expected_tools"`
Group             string            `yaml:"group"`

Properties map[string]string `yaml:"properties"`

// Inline graders (both .prompt.yaml and .prompt.md frontmatter)
Graders []interface{} `yaml:"graders,omitempty"`

// YAML-only prompt fields (used by .prompt.yaml/.prompt.yml)
PromptTextField         string `yaml:"prompt_text"`
EvaluationCriteriaField string `yaml:"evaluation_criteria"`
}

// frontmatterToPrompt converts a decoded frontmatter into a Prompt.
func frontmatterToPrompt(fm *frontmatter) *Prompt {
props := fm.Properties
if props == nil {
props = make(map[string]string)
}
return &Prompt{
ID:                fm.ID,
Tags:              fm.Tags,
ProjectContext:    fm.ProjectContext,
StarterProject:    fm.StarterProject,
ReferenceAnswer:   fm.ReferenceAnswer,
Timeout:           fm.Timeout,
MaxSessionActions: fm.MaxSessionActions,
MaxTurns:          fm.MaxTurns,
ExpectedPkgs:      fm.ExpectedPkgs,
ExpectedTools:     fm.ExpectedTools,
Group:             strings.TrimSpace(fm.Group),
Properties:        props,
}
}

// splitSections splits a markdown body at all ## headings into a map of
// heading name → section content. Content before any ## heading is stored
// under the empty-string key.
func splitSections(body string) map[string]string {
sections := make(map[string]string)
locs := sectionHeadingRe.FindAllStringSubmatchIndex(body, -1)
if len(locs) == 0 {
if t := strings.TrimSpace(body); t != "" {
sections[""] = t
}
return sections
}
// Content before the first heading
if preamble := strings.TrimSpace(body[:locs[0][0]]); preamble != "" {
sections[""] = preamble
}
for i, loc := range locs {
heading := strings.TrimSpace(body[loc[2]:loc[3]])
contentStart := loc[1]
var contentEnd int
if i+1 < len(locs) {
contentEnd = locs[i+1][0]
} else {
contentEnd = len(body)
}
sections[heading] = strings.TrimSpace(body[contentStart:contentEnd])
}
return sections
}

// ParsePromptFile parses a .prompt.md file's content into a Prompt struct.
// Metadata must use the nested properties: format in frontmatter.
// For .prompt.yaml/.prompt.yml files, use ParsePromptYAML instead.
func ParsePromptFile(content []byte, filePath string) (*Prompt, error) {
text := string(content)

if !strings.HasPrefix(text, "---") {
return nil, fmt.Errorf("file does not start with frontmatter delimiter: %s", filePath)
}
parts := strings.SplitN(text[3:], "---", 2)
if len(parts) < 2 {
return nil, fmt.Errorf("missing closing frontmatter delimiter: %s", filePath)
}
fmText := strings.TrimSpace(parts[0])
body := parts[1]

var fm frontmatter
dec := yaml.NewDecoder(bytes.NewReader([]byte(fmText)))
if err := dec.Decode(&fm); err != nil {
return nil, fmt.Errorf("parsing frontmatter in %s: %w", filePath, err)
}

p := frontmatterToPrompt(&fm)

// Parse inline graders if present
if len(fm.Graders) > 0 {
	graders, err := criteria.ParseInlineGraders(fm.Graders, p.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing inline graders in %s: %w", filePath, err)
	}
	p.Graders = graders
}

sections := splitSections(body)
if s, ok := sections["Prompt"]; ok {
p.PromptText = s
}
if s, ok := sections["Evaluation Criteria"]; ok {
	slog.Warn("DEPRECATED: '## Evaluation Criteria' markdown section is deprecated and will be removed in a future release. Migrate to 'graders:' frontmatter with 'type: prompt'.",
		"file", filePath,
		"migration_guide", "Replace '## Evaluation Criteria' section with frontmatter 'graders:' field using 'type: prompt'")
	p.EvaluationCriteria = s
	p.ParsedCriteria = ParseEvaluationCriteria(s)
}

p.FilePath = filePath

if p.ID == "" {
return nil, fmt.Errorf("prompt missing required 'id' field: %s", filePath)
}

return p, nil
}

// ParsePromptYAML parses a pure YAML prompt file (.prompt.yaml or .prompt.yml)
// into a Prompt struct. All fields including prompt_text and evaluation_criteria
// are expressed as top-level YAML keys.
func ParsePromptYAML(content []byte, filePath string) (*Prompt, error) {
var fm frontmatter
dec := yaml.NewDecoder(bytes.NewReader(content))
if err := dec.Decode(&fm); err != nil {
return nil, fmt.Errorf("parsing YAML prompt %s: %w", filePath, err)
}

p := frontmatterToPrompt(&fm)

// Parse inline graders if present
if len(fm.Graders) > 0 {
	graders, err := criteria.ParseInlineGraders(fm.Graders, p.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing inline graders in %s: %w", filePath, err)
	}
	p.Graders = graders
}

p.PromptText = fm.PromptTextField
p.EvaluationCriteria = fm.EvaluationCriteriaField
if p.EvaluationCriteria != "" {
	slog.Warn("DEPRECATED: 'evaluation_criteria:' frontmatter field is deprecated and will be removed in a future release. Migrate to 'graders:' frontmatter with 'type: prompt'.",
		"file", filePath,
		"migration_guide", "Replace 'evaluation_criteria:' field with 'graders:' using 'type: prompt'")
	p.ParsedCriteria = ParseEvaluationCriteria(p.EvaluationCriteria)
}
p.FilePath = filePath

if p.ID == "" {
return nil, fmt.Errorf("prompt missing required 'id' field: %s", filePath)
}

return p, nil
}

// ParseEvaluationCriteria parses an evaluation-criteria text block into a
// single CriterionEntry. All non-bullet leading lines become the entry's
// Prompt field (preamble shown to the LLM judge). All "- " bullet lines
// (top-level or nested) become individual entries in the Checks slice —
// each check is one pass/fail criterion the judge must score.
//
// When the input contains no bullets at all, the entire trimmed text is
// used as a single Check so the judge still has something to score against.
// Empty input returns nil.
//
// Returns []CriterionEntry (length 0 or 1) for API symmetry with the
// previous one-entry-per-bullet shape.
func ParseEvaluationCriteria(text string) []CriterionEntry {
if strings.TrimSpace(text) == "" {
return nil
}

var preambleLines []string
var checks []string
sawBullet := false

for _, line := range strings.Split(text, "\n") {
trimmedRight := strings.TrimRight(line, " \t\r")
stripped := strings.TrimLeft(trimmedRight, " \t")

if strings.HasPrefix(stripped, "- ") {
sawBullet = true
c := strings.TrimSpace(strings.TrimPrefix(stripped, "- "))
if c != "" {
checks = append(checks, c)
}
continue
}

// Non-bullet line. If we haven't seen a bullet yet, treat as preamble.
// Lines after the first bullet are ignored under the new flat-checks model.
if !sawBullet {
preambleLines = append(preambleLines, trimmedRight)
}
}

if !sawBullet {
whole := strings.TrimSpace(text)
if whole == "" {
return nil
}
return []CriterionEntry{{Checks: []string{whole}}}
}

preamble := strings.TrimSpace(strings.Join(preambleLines, "\n"))
if preamble == "" && len(checks) == 0 {
return nil
}
return []CriterionEntry{{Prompt: preamble, Checks: checks}}
}

// FormatParsedCriteria renders parsed criterion entries into a deterministic
// text block for injection into the LLM-judge prompt. Output shape:
//
//	<preamble>
//	1. check1
//	2. check2
//
// The preamble (if non-empty) is rendered unnumbered, followed by each check
// as a numbered top-level item. When the input contains multiple entries
// (legacy callers), they are concatenated with blank-line separators, with
// numbering reset within each entry.
//
// For back-compat: if an entry has no Checks but legacy SubPoints, the
// SubPoints are rendered as the numbered checks under the entry's Prompt.
func FormatParsedCriteria(entries []CriterionEntry) string {
if len(entries) == 0 {
return ""
}
var blocks []string
for _, e := range entries {
preamble := strings.TrimSpace(e.Prompt)
var checks []string
for _, c := range e.Checks {
c = strings.TrimSpace(c)
if c != "" {
checks = append(checks, c)
}
}
if len(checks) == 0 && len(e.SubPoints) > 0 {
for _, s := range e.SubPoints {
s = strings.TrimSpace(s)
if s != "" {
checks = append(checks, s)
}
}
}
if preamble == "" && len(checks) == 0 {
continue
}
var b strings.Builder
if preamble != "" {
b.WriteString(preamble)
if len(checks) > 0 {
b.WriteString("\n")
}
}
for i, c := range checks {
fmt.Fprintf(&b, "%d. %s", i+1, c)
if i < len(checks)-1 {
b.WriteString("\n")
}
}
blocks = append(blocks, b.String())
}
return strings.Join(blocks, "\n\n")
}
