package prompt

import (
"bytes"
"fmt"
"regexp"
"strings"

"gopkg.in/yaml.v3"
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

sections := splitSections(body)
if s, ok := sections["Prompt"]; ok {
p.PromptText = s
}
if s, ok := sections["Evaluation Criteria"]; ok {
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
p.PromptText = fm.PromptTextField
p.EvaluationCriteria = fm.EvaluationCriteriaField
if p.EvaluationCriteria != "" {
p.ParsedCriteria = ParseEvaluationCriteria(p.EvaluationCriteria)
}
p.FilePath = filePath

if p.ID == "" {
return nil, fmt.Errorf("prompt missing required 'id' field: %s", filePath)
}

return p, nil
}

// ParseEvaluationCriteria parses a raw evaluation criteria text into
// structured CriterionEntry values. Top-level bullet points (lines starting
// with "- ") become individual entries. Sub-points (lines starting with
// "  - " under a parent) are grouped with the parent entry.
func ParseEvaluationCriteria(text string) []CriterionEntry {
if strings.TrimSpace(text) == "" {
return nil
}

var entries []CriterionEntry
var current *CriterionEntry

for _, line := range strings.Split(text, "\n") {
trimmed := strings.TrimRight(line, " \t\r")

// Sub-point: starts with two+ spaces then "- "
if (strings.HasPrefix(trimmed, "  - ") || strings.HasPrefix(trimmed, "\t- ")) && current != nil {
sub := strings.TrimSpace(trimmed)
sub = strings.TrimPrefix(sub, "- ")
if sub != "" {
current.SubPoints = append(current.SubPoints, sub)
}
continue
}

// Top-level bullet: starts with "- "
if strings.HasPrefix(trimmed, "- ") {
if current != nil {
entries = append(entries, *current)
}
prompt := strings.TrimPrefix(trimmed, "- ")
current = &CriterionEntry{Prompt: prompt}
continue
}
}

if current != nil {
entries = append(entries, *current)
}

return entries
}
