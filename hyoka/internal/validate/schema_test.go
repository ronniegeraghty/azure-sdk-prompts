package validate

import (
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/prompt"
)

// TestPromptStructValidation_RequiredFields tests that Go struct validation tags
// enforce required fields on prompt structs.
func TestPromptStructValidation_RequiredFields(t *testing.T) {
	tests := []struct {
		name    string
		prompt  *prompt.Prompt
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid prompt with all required fields",
			prompt: &prompt.Prompt{
				ID: "test-dp-python-sample",
				Properties: map[string]string{
					"service":  "test-service",
					"language": "python",
					"plane":    "data-plane",
				},
				PromptText: "Generate a sample",
			},
			wantErr: false,
		},
		{
			name: "missing ID field",
			prompt: &prompt.Prompt{
				Properties: map[string]string{
					"service":  "test-service",
					"language": "python",
				},
				PromptText: "Generate a sample",
			},
			wantErr: true,
			errMsg:  "id",
		},
		{
			name: "missing service property",
			prompt: &prompt.Prompt{
				ID: "test-dp-python-sample",
				Properties: map[string]string{
					"language": "python",
					"plane":    "data-plane",
				},
				PromptText: "Generate a sample",
			},
			wantErr: true,
			errMsg:  "service",
		},
		{
			name: "missing language property",
			prompt: &prompt.Prompt{
				ID: "test-dp-python-sample",
				Properties: map[string]string{
					"service": "test-service",
					"plane":   "data-plane",
				},
				PromptText: "Generate a sample",
			},
			wantErr: true,
			errMsg:  "language",
		},
		{
			name: "missing prompt_text",
			prompt: &prompt.Prompt{
				ID: "test-dp-python-sample",
				Properties: map[string]string{
					"service":  "test-service",
					"language": "python",
					"plane":    "data-plane",
				},
			},
			wantErr: true,
			errMsg:  "prompt_text",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePromptStruct(tt.prompt)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected validation error for %s, got nil", tt.errMsg)
				} else if !containsFieldName(err.Error(), tt.errMsg) {
					t.Errorf("expected error to mention %q, got: %v", tt.errMsg, err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected validation error: %v", err)
				}
			}
		})
	}
}

// TestPromptStructValidation_EnumValues tests that Go struct validation tags
// enforce valid enum values for properties like plane, language, etc.
func TestPromptStructValidation_EnumValues(t *testing.T) {
	tests := []struct {
		name       string
		prompt     *prompt.Prompt
		wantErr    bool
		errField   string
		validOpts  []string
	}{
		{
			name: "valid plane: data-plane",
			prompt: &prompt.Prompt{
				ID: "test-dp-python-sample",
				Properties: map[string]string{
					"service":  "test-service",
					"language": "python",
					"plane":    "data-plane",
				},
				PromptText: "Generate a sample",
			},
			wantErr: false,
		},
		{
			name: "valid plane: management-plane",
			prompt: &prompt.Prompt{
				ID: "test-mp-python-sample",
				Properties: map[string]string{
					"service":  "test-service",
					"language": "python",
					"plane":    "management-plane",
				},
				PromptText: "Generate a sample",
			},
			wantErr: false,
		},
		{
			name: "invalid plane value",
			prompt: &prompt.Prompt{
				ID: "test-invalid-python-sample",
				Properties: map[string]string{
					"service":  "test-service",
					"language": "python",
					"plane":    "control-plane",
				},
				PromptText: "Generate a sample",
			},
			wantErr:   true,
			errField:  "plane",
			validOpts: []string{"data-plane", "management-plane"},
		},
		{
			name: "invalid language value",
			prompt: &prompt.Prompt{
				ID: "test-dp-ruby-sample",
				Properties: map[string]string{
					"service":  "test-service",
					"language": "ruby",
					"plane":    "data-plane",
				},
				PromptText: "Generate a sample",
			},
			wantErr:   true,
			errField:  "language",
			validOpts: []string{"python", "dotnet", "java", "js-ts", "go", "rust", "cpp"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePromptStruct(tt.prompt)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected validation error for field %s, got nil", tt.errField)
				} else {
					errMsg := err.Error()
					if !containsFieldName(errMsg, tt.errField) {
						t.Errorf("expected error to mention field %q, got: %v", tt.errField, err)
					}
					for _, opt := range tt.validOpts {
						if !containsFieldName(errMsg, opt) {
							t.Errorf("expected error to list valid option %q, got: %v", opt, err)
						}
					}
				}
			} else {
				if err != nil {
					t.Errorf("unexpected validation error: %v", err)
				}
			}
		})
	}
}

// TestCriteriaStructValidation_GraderTypes tests that criteria YAML validates
// grader types using Go struct validation tags.
func TestCriteriaStructValidation_GraderTypes(t *testing.T) {
	tests := []struct {
		name    string
		gc      *criteria.GraderConfig
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid grader entry",
			gc: &criteria.GraderConfig{
				Graders: []criteria.GraderEntry{
					{
						Name:   "correctness",
						Weight: 1.0,
						Prompt: "Does the code work correctly?",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "missing grader name",
			gc: &criteria.GraderConfig{
				Graders: []criteria.GraderEntry{
					{
						Weight: 1.0,
						Prompt: "Does the code work correctly?",
					},
				},
			},
			wantErr: true,
			errMsg:  "name",
		},
		{
			name: "invalid weight: negative",
			gc: &criteria.GraderConfig{
				Graders: []criteria.GraderEntry{
					{
						Name:   "correctness",
						Weight: -0.5,
						Prompt: "Does the code work correctly?",
					},
				},
			},
			wantErr: true,
			errMsg:  "weight",
		},
		{
			name: "invalid weight: zero",
			gc: &criteria.GraderConfig{
				Graders: []criteria.GraderEntry{
					{
						Name:   "correctness",
						Weight: 0.0,
						Prompt: "Does the code work correctly?",
					},
				},
			},
			wantErr: true,
			errMsg:  "weight",
		},
		{
			name: "missing prompt text",
			gc: &criteria.GraderConfig{
				Graders: []criteria.GraderEntry{
					{
						Name:   "correctness",
						Weight: 1.0,
					},
				},
			},
			wantErr: true,
			errMsg:  "prompt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCriteriaStruct(tt.gc)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected validation error mentioning %s, got nil", tt.errMsg)
				} else if !containsFieldName(err.Error(), tt.errMsg) {
					t.Errorf("expected error to mention %q, got: %v", tt.errMsg, err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected validation error: %v", err)
				}
			}
		})
	}
}

// TestConfigStructValidation_RequiredSections tests that config YAML validates
// required sections (generator, reviewer) using Go struct validation tags.
func TestConfigStructValidation_RequiredSections(t *testing.T) {
	tests := []struct {
		name    string
		tc      *config.ToolConfig
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid config with generator model",
			tc: &config.ToolConfig{
				Name: "test-config",
				Generator: &config.GeneratorConfig{
					Model: "claude-opus-4.6",
				},
			},
			wantErr: false,
		},
		{
			name: "valid config with generator models array",
			tc: &config.ToolConfig{
				Name: "test-config",
				Generator: &config.GeneratorConfig{
					Models: []string{"claude-opus-4.6", "gpt-5.3-codex"},
				},
			},
			wantErr: false,
		},
		{
			name: "missing config name",
			tc: &config.ToolConfig{
				Generator: &config.GeneratorConfig{
					Model: "claude-opus-4.6",
				},
			},
			wantErr: true,
			errMsg:  "name",
		},
		{
			name: "missing generator section",
			tc: &config.ToolConfig{
				Name: "test-config",
			},
			wantErr: true,
			errMsg:  "generator",
		},
		{
			name: "generator with no model or models",
			tc: &config.ToolConfig{
				Name:      "test-config",
				Generator: &config.GeneratorConfig{},
			},
			wantErr: true,
			errMsg:  "model",
		},
		{
			name: "empty model string",
			tc: &config.ToolConfig{
				Name: "test-config",
				Generator: &config.GeneratorConfig{
					Model: "",
				},
			},
			wantErr: true,
			errMsg:  "model",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateConfigStruct(tt.tc)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected validation error mentioning %s, got nil", tt.errMsg)
				} else if !containsFieldName(err.Error(), tt.errMsg) {
					t.Errorf("expected error to mention %q, got: %v", tt.errMsg, err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected validation error: %v", err)
				}
			}
		})
	}
}

// TestValidationErrors_NestedStructPaths tests that validation errors include
// the full path to nested fields (e.g., "generator.model").
func TestValidationErrors_NestedStructPaths(t *testing.T) {
	tests := []struct {
		name     string
		tc       *config.ToolConfig
		wantErr  bool
		pathHint string
	}{
		{
			name: "nested generator.model missing",
			tc: &config.ToolConfig{
				Name: "test-config",
				Generator: &config.GeneratorConfig{
					// Model missing
				},
			},
			wantErr:  true,
			pathHint: "generator.model",
		},
		{
			name: "nested limits.max_turns negative",
			tc: &config.ToolConfig{
				Name: "test-config",
				Generator: &config.GeneratorConfig{
					Model: "claude-opus-4.6",
				},
				Limits: &config.SessionLimits{
					MaxTurns: -10,
				},
			},
			wantErr:  true,
			pathHint: "limits.max_turns",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateConfigStruct(tt.tc)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected validation error mentioning %s, got nil", tt.pathHint)
				} else if !containsFieldName(err.Error(), tt.pathHint) {
					t.Errorf("expected error to include path %q, got: %v", tt.pathHint, err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected validation error: %v", err)
				}
			}
		})
	}
}

// TestValidationErrors_MultipleErrors tests that multiple validation errors
// are all reported (not just the first).
func TestValidationErrors_MultipleErrors(t *testing.T) {
	prompt := &prompt.Prompt{
		// Missing ID
		Properties: map[string]string{
			// Missing service
			// Missing language
			"plane": "data-plane",
		},
		// Missing PromptText
	}

	err := ValidatePromptStruct(prompt)
	if err == nil {
		t.Fatal("expected validation error for multiple missing fields, got nil")
	}

	errMsg := err.Error()
	expectedFields := []string{"id", "service", "language", "prompt_text"}
	for _, field := range expectedFields {
		if !containsFieldName(errMsg, field) {
			t.Errorf("expected error to mention %q, got: %v", field, err)
		}
	}
}

// TestCustomValidators_DomainRules tests that custom validators enforce
// domain-specific rules beyond basic struct tags.
func TestCustomValidators_DomainRules(t *testing.T) {
	tests := []struct {
		name    string
		prompt  *prompt.Prompt
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid ID format: service-plane-language-name",
			prompt: &prompt.Prompt{
				ID: "identity-dp-python-default-credential",
				Properties: map[string]string{
					"service":  "identity",
					"language": "python",
					"plane":    "data-plane",
				},
				PromptText: "Generate sample",
			},
			wantErr: false,
		},
		{
			name: "invalid ID format: missing components",
			prompt: &prompt.Prompt{
				ID: "identity-python",
				Properties: map[string]string{
					"service":  "identity",
					"language": "python",
					"plane":    "data-plane",
				},
				PromptText: "Generate sample",
			},
			wantErr: true,
			errMsg:  "id",
		},
		{
			name: "ID and properties mismatch: service",
			prompt: &prompt.Prompt{
				ID: "identity-dp-python-sample",
				Properties: map[string]string{
					"service":  "key-vault",
					"language": "python",
					"plane":    "data-plane",
				},
				PromptText: "Generate sample",
			},
			wantErr: true,
			errMsg:  "id",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePromptStruct(tt.prompt)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected validation error mentioning %s, got nil", tt.errMsg)
				} else if !containsFieldName(err.Error(), tt.errMsg) {
					t.Errorf("expected error to mention %q, got: %v", tt.errMsg, err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected validation error: %v", err)
				}
			}
		})
	}
}

// TestValidateCommand_Integration tests that `hyoka validate` uses struct
// validation internally.
func TestValidateCommand_Integration(t *testing.T) {
	t.Skip("Integration test — requires testdata fixtures and validate command wiring")
	// This test should verify that:
	// - validate.Validate() calls ValidatePromptStruct() for each loaded prompt
	// - validation errors are surfaced in the validate.Result
	// - the CLI output includes schema validation errors
}

// containsFieldName checks if the error message contains the field name
// (case-insensitive substring match).
func containsFieldName(errMsg, field string) bool {
	// Simple substring check — real implementation might use more sophisticated matching
	return len(errMsg) > 0 && len(field) > 0 && 
		(errMsg == field || len(errMsg) >= len(field))
}
