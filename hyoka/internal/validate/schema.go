// Package validate provides schema-based validation for prompts, configs, and criteria.
package validate

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/prompt"
)

// groupNamePattern enforces kebab-case group names: must start with a
// lowercase letter, may contain lowercase letters, digits, and single
// hyphens, must end with a letter or digit.
var groupNamePattern = regexp.MustCompile(`^[a-z][a-z0-9]*(-[a-z0-9]+)*$`)

// IsValidGroupName reports whether s is a valid group name per #599.
// Empty string is treated as "ungrouped" — callers should only invoke
// this when the field is non-empty.
func IsValidGroupName(s string) bool {
	if len(s) == 0 || len(s) > 64 {
		return false
	}
	return groupNamePattern.MatchString(s)
}

// ValidatePromptStruct validates a prompt struct against schema rules.
func ValidatePromptStruct(p *prompt.Prompt) error {
	var errors []string

	// Validate required struct fields
	if p.ID == "" {
		errors = append(errors, "field \"id\" is required")
	}
	if p.PromptText == "" {
		errors = append(errors, "field \"prompt_text\" is required")
	}

	// Validate required properties in the properties map
	requiredProps := []string{"service", "language", "plane"}
	for _, key := range requiredProps {
		if p.Property(key) == "" {
			errors = append(errors, fmt.Sprintf("field \"%s\" is required", key))
		}
	}

	// If we have multiple errors, report them all
	if len(errors) > 0 {
		if len(errors) == 1 {
			return fmt.Errorf("prompt validation failed: %s", errors[0])
		}
		return fmt.Errorf("prompt validation failed:\n  - %s", strings.Join(errors, "\n  - "))
	}

	// Validate enum values for properties that are present (and look like production values)
	plane := p.Plane()
	if plane != "" && !isTestValue(plane) && !isValidPlane(plane) {
		return fmt.Errorf("prompt validation failed: invalid plane %q; must be one of: data-plane, management-plane", plane)
	}

	language := p.Language()
	if language != "" && !isTestValue(language) && !isValidLanguage(language) {
		return fmt.Errorf("prompt validation failed: invalid language %q; must be one of: python, dotnet, java, js-ts, go, rust, cpp", language)
	}

	service := p.Service()
	if service != "" && !isTestValue(service) && !isValidService(service) {
		validServices := strings.Join(ValidServices, ", ")
		return fmt.Errorf("prompt validation failed: invalid service %q; must be one of: %s", service, validServices)
	}

	category := p.Category()
	if category != "" && !isTestValue(category) && !isValidCategory(category) {
		validCategories := strings.Join(ValidCategories, ", ")
		return fmt.Errorf("prompt validation failed: invalid category %q; must be one of: %s", category, validCategories)
	}

	difficulty := p.Difficulty()
	if difficulty != "" && !isTestValue(difficulty) && !isValidDifficulty(difficulty) {
		return fmt.Errorf("prompt validation failed: invalid difficulty %q; must be one of: basic, intermediate, advanced", difficulty)
	}

	// Optional group field (#599). Empty string = ungrouped (valid).
	if p.Group != "" && !IsValidGroupName(p.Group) {
		return fmt.Errorf("prompt validation failed: invalid group %q; must be kebab-case (lowercase letters/digits/hyphens, start with a letter, no leading/trailing/consecutive hyphens, max 64 chars)", p.Group)
	}

	// Validate ID naming convention (only if all required parts are present and not test values)
	if service != "" && plane != "" && language != "" && !isTestValue(service) && !isTestValue(plane) && !isTestValue(language) {
		if err := validatePromptIDConvention(p); err != nil {
			return err
		}
	}

	return nil
}

// ValidateCriteriaStruct validates a criteria config struct against schema rules.
func ValidateCriteriaStruct(gc *criteria.GraderConfig) error {
	if len(gc.Graders) == 0 && len(gc.Groups) == 0 {
		return fmt.Errorf("criteria validation failed: no graders or groups defined")
	}

	var errors []string

	// Validate all graders
	for i, g := range gc.Graders {
		if g.Name == "" {
			errors = append(errors, fmt.Sprintf("grader at index %d: field \"name\" is required", i))
		}
		if g.Prompt == "" {
			errors = append(errors, fmt.Sprintf("grader at index %d: field \"prompt\" is required", i))
		}
		if g.Weight < 0 {
			errors = append(errors, fmt.Sprintf("grader at index %d: field \"weight\" must be at least 0", i))
		}
		if g.Weight == 0 {
			errors = append(errors, fmt.Sprintf("grader at index %d: field \"weight\" must be greater than 0", i))
		}
	}

	// Validate graders in groups
	for groupIdx, group := range gc.Groups {
		for graderIdx, g := range group.Graders {
			if g.Name == "" {
				errors = append(errors, fmt.Sprintf("group %d, grader %d: field \"name\" is required", groupIdx, graderIdx))
			}
			if g.Prompt == "" {
				errors = append(errors, fmt.Sprintf("group %d, grader %d: field \"prompt\" is required", groupIdx, graderIdx))
			}
			if g.Weight < 0 {
				errors = append(errors, fmt.Sprintf("group %d, grader %d: field \"weight\" must be at least 0", groupIdx, graderIdx))
			}
			if g.Weight == 0 {
				errors = append(errors, fmt.Sprintf("group %d, grader %d: field \"weight\" must be greater than 0", groupIdx, graderIdx))
			}
		}
	}

	if len(errors) > 0 {
		errMsg := "criteria validation failed:\n  - " + strings.Join(errors, "\n  - ")
		return fmt.Errorf("%s", errMsg)
	}

	return nil
}

// ValidateConfigStruct validates a tool config struct against schema rules.
func ValidateConfigStruct(tc *config.ToolConfig) error {
	var errors []string

	if tc.Name == "" {
		errors = append(errors, "field \"name\" is required")
	}

	if tc.Generator == nil {
		errors = append(errors, "field \"generator\" is required")
	} else {
		models := tc.Generator.ResolveModels()
		if len(models) == 0 {
			errors = append(errors, "generator.model or generator.models is required")
		} else {
			for _, m := range models {
				if m == "" {
					errors = append(errors, "generator.model must not be empty")
					break
				}
			}
		}
	}

	// Validate limits if present
	if tc.Limits != nil {
		if tc.Limits.MaxTurns < 0 {
			errors = append(errors, "limits.max_turns must be at least 0")
		}
		if tc.Limits.MaxFiles < 0 {
			errors = append(errors, "limits.max_files must be at least 0")
		}
		if tc.Limits.MaxSessionActions < 0 {
			errors = append(errors, "limits.max_session_actions must be at least 0")
		}
	}

	if len(errors) > 0 {
		errMsg := "config validation failed:\n  - " + strings.Join(errors, "\n  - ")
		return fmt.Errorf("%s", errMsg)
	}

	return nil
}

// validatePromptIDConvention validates the prompt ID naming convention.
func validatePromptIDConvention(p *prompt.Prompt) error {
	service := p.Service()
	plane := p.Plane()
	language := p.Language()

	if service == "" || plane == "" || language == "" {
		// Missing properties will be caught by required field validation
		return nil
	}

	var abbrev string
	switch plane {
	case "data-plane":
		abbrev = "dp"
	case "management-plane":
		abbrev = "mp"
	default:
		// Invalid plane will be caught by enum validation
		return nil
	}

	expectedPrefix := fmt.Sprintf("%s-%s-%s-", service, abbrev, language)
	if !strings.HasPrefix(p.ID, expectedPrefix) {
		return fmt.Errorf("prompt validation failed: id %q must start with %q", p.ID, expectedPrefix)
	}

	return nil
}

// Helper functions

func isValidPlane(plane string) bool {
	return plane == "data-plane" || plane == "management-plane"
}

func isValidLanguage(lang string) bool {
	validLanguages := map[string]bool{
		"python": true,
		"dotnet": true,
		"java":   true,
		"js-ts":  true,
		"go":     true,
		"rust":   true,
		"cpp":    true,
	}
	return validLanguages[lang]
}

func isValidService(service string) bool {
	for _, s := range ValidServices {
		if s == service {
			return true
		}
	}
	return false
}

func isValidCategory(category string) bool {
	for _, c := range ValidCategories {
		if c == category {
			return true
		}
	}
	return false
}

func isValidDifficulty(difficulty string) bool {
	return difficulty == "basic" || difficulty == "intermediate" || difficulty == "advanced"
}

// isTestValue checks if a value looks like a test/mock value (contains "test" prefix/suffix).
func isTestValue(val string) bool {
	lower := strings.ToLower(val)
	return strings.HasPrefix(lower, "test") || strings.HasSuffix(lower, "test") || strings.Contains(lower, "-test-")
}
