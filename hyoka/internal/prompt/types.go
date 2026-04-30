package prompt

// Prompt represents a parsed prompt file (.prompt.md, .prompt.yaml, or .prompt.yml)
// with metadata and prompt text.
//
// Metadata string fields (service, language, plane, etc.) are stored in the
// Properties map rather than as dedicated struct fields. Use the getter methods
// (Service(), Language(), etc.) or the generic Property() method to access them.
// Non-string typed fields (Tags, Timeout, etc.) remain as struct fields.
type Prompt struct {
ID              string            `yaml:"id" json:"id"`
Tags            []string          `yaml:"tags" json:"tags"`
ProjectContext  map[string]string `yaml:"project_context" json:"project_context,omitempty"`
StarterProject  string            `yaml:"starter_project" json:"starter_project,omitempty"`
ReferenceAnswer string            `yaml:"reference_answer" json:"reference_answer,omitempty"`
Timeout           int               `yaml:"timeout" json:"timeout,omitempty"`
MaxSessionActions int               `yaml:"max_session_actions" json:"max_session_actions,omitempty"`
MaxTurns          int               `yaml:"max_turns" json:"max_turns,omitempty"`
ExpectedPkgs      []string          `yaml:"expected_packages" json:"expected_packages,omitempty"`
ExpectedTools     []string          `yaml:"expected_tools" json:"expected_tools,omitempty"`

// Group is an optional logical grouping label used by reports/site to
// cluster related prompt evaluations (e.g., "crud-operations",
// "auth-flows"). Empty string means ungrouped. Validation rules:
// kebab-case (lowercase ASCII letters/digits/hyphens), 1-64 chars,
// must start with a letter, no leading/trailing/consecutive hyphens.
Group string `yaml:"group" json:"group,omitempty"`

// Properties holds all metadata string fields (service, language, plane, etc.).
// Keys must be snake_case.
Properties map[string]string `yaml:"properties" json:"properties"`

// The prompt text: extracted from ## Prompt section (.prompt.md)
// or from the prompt_text field (.prompt.yaml/.prompt.yml)
PromptText string `yaml:"-" json:"prompt_text"`

// The evaluation criteria: extracted from ## Evaluation Criteria section (.prompt.md)
// or from the evaluation_criteria field (.prompt.yaml/.prompt.yml)
EvaluationCriteria string `yaml:"-" json:"evaluation_criteria,omitempty"`

// ParsedCriteria is the structured representation of EvaluationCriteria,
// extracted by ParseEvaluationCriteria during prompt parsing.
ParsedCriteria []CriterionEntry `yaml:"-" json:"parsed_criteria,omitempty"`

// Source file path
FilePath string `yaml:"-" json:"file_path"`
}

// CriterionEntry represents the parsed evaluation criteria from a prompt
// file. After the grader-redesign rewrite of ParseEvaluationCriteria, exactly
// one CriterionEntry is produced per prompt: the leading non-bullet text is
// captured as Prompt (preamble) and every "- " bullet becomes one entry in
// Checks. SubPoints is preserved for back-compat with on-disk reports/tests
// but is no longer populated by the parser.
type CriterionEntry struct {
Prompt    string   `json:"prompt"`
Checks    []string `json:"checks,omitempty"`
SubPoints []string `json:"sub_points,omitempty"`
}

// Property returns the value of a metadata property by key.
func (p *Prompt) Property(key string) string {
return p.Properties[key]
}

// Convenience getters for well-known metadata properties.

func (p *Prompt) Service() string     { return p.Properties["service"] }
func (p *Prompt) Plane() string       { return p.Properties["plane"] }
func (p *Prompt) Language() string    { return p.Properties["language"] }
func (p *Prompt) Category() string    { return p.Properties["category"] }
func (p *Prompt) Difficulty() string  { return p.Properties["difficulty"] }
func (p *Prompt) Description() string { return p.Properties["description"] }
func (p *Prompt) SDKPackage() string  { return p.Properties["sdk_package"] }
func (p *Prompt) DocURL() string      { return p.Properties["doc_url"] }
func (p *Prompt) Created() string     { return p.Properties["created"] }
func (p *Prompt) Author() string      { return p.Properties["author"] }

// Filter defines criteria for selecting prompts.
// Filters is a map of property key to value pairs; all must match.
// Tags and PromptID are handled separately.
type Filter struct {
Filters  map[string]string
Tags     []string
PromptID string
}

// Matches returns true if the prompt matches all non-empty filter criteria.
func (p *Prompt) Matches(f Filter) bool {
if f.PromptID != "" && p.ID != f.PromptID {
return false
}
for key, value := range f.Filters {
if p.Property(key) != value {
return false
}
}
if len(f.Tags) > 0 {
tagSet := make(map[string]bool, len(p.Tags))
for _, t := range p.Tags {
tagSet[t] = true
}
for _, required := range f.Tags {
if !tagSet[required] {
return false
}
}
}
return true
}
