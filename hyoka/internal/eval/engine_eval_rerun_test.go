package eval

import (
	"testing"
)

func TestBuildRerunCommand(t *testing.T) {
	tests := []struct {
		name            string
		promptID        string
		configName      string
		baseConfigName  string
		generatorModel  string
		pairwiseVariant string
		opts            EngineOptions
		want            string
	}{
		{
			name:            "single-model config",
			promptID:        "test-prompt-id",
			configName:      "baseline/claude-opus-4.6",
			baseConfigName:  "",
			generatorModel:  "claude-opus-4.6",
			pairwiseVariant: "",
			opts:            EngineOptions{MaxSessionActions: 50},
			want:            "hyoka run --prompt-id test-prompt-id --config baseline/claude-opus-4.6",
		},
		{
			name:            "multi-model config with --model",
			promptID:        "test-prompt-id",
			configName:      "python-pairwise/claude-opus-4.6",
			baseConfigName:  "python-pairwise",
			generatorModel:  "claude-opus-4.6",
			pairwiseVariant: "",
			opts:            EngineOptions{MaxSessionActions: 50},
			want:            "hyoka run --prompt-id test-prompt-id --config python-pairwise --model claude-opus-4.6",
		},
		{
			name:            "multi-model config with different model",
			promptID:        "key-vault-dp-python-crud",
			configName:      "python-pairwise/gpt-5.3-codex",
			baseConfigName:  "python-pairwise",
			generatorModel:  "gpt-5.3-codex",
			pairwiseVariant: "",
			opts:            EngineOptions{MaxSessionActions: 50},
			want:            "hyoka run --prompt-id key-vault-dp-python-crud --config python-pairwise --model gpt-5.3-codex",
		},
		{
			name:            "with skip-review flag",
			promptID:        "test-prompt",
			configName:      "baseline/opus",
			baseConfigName:  "",
			generatorModel:  "claude-opus-4.6",
			pairwiseVariant: "",
			opts:            EngineOptions{SkipReview: true, MaxSessionActions: 50},
			want:            "hyoka run --prompt-id test-prompt --config baseline/opus --skip-review",
		},
		{
			name:            "with monitor-resources flag",
			promptID:        "test-prompt",
			configName:      "baseline/opus",
			baseConfigName:  "",
			generatorModel:  "claude-opus-4.6",
			pairwiseVariant: "",
			opts:            EngineOptions{MonitorResources: true, MaxSessionActions: 50},
			want:            "hyoka run --prompt-id test-prompt --config baseline/opus --monitor-resources",
		},
		{
			name:            "with custom max-session-actions",
			promptID:        "test-prompt",
			configName:      "baseline/opus",
			baseConfigName:  "",
			generatorModel:  "claude-opus-4.6",
			pairwiseVariant: "",
			opts:            EngineOptions{MaxSessionActions: 100},
			want:            "hyoka run --prompt-id test-prompt --config baseline/opus --max-session-actions=100",
		},
		{
			name:            "multi-model with all flags",
			promptID:        "complex-test",
			configName:      "multi-model/claude-sonnet-4.6",
			baseConfigName:  "multi-model",
			generatorModel:  "claude-sonnet-4.6",
			pairwiseVariant: "",
			opts:            EngineOptions{SkipReview: true, MonitorResources: true, MaxSessionActions: 75},
			want:            "hyoka run --prompt-id complex-test --config multi-model --model claude-sonnet-4.6 --skip-review --monitor-resources --max-session-actions=75",
		},
		{
			name:            "prompt ID with hyphens",
			promptID:        "identity-dp-python-default-credential",
			configName:      "baseline/claude-opus-4.6",
			baseConfigName:  "",
			generatorModel:  "claude-opus-4.6",
			pairwiseVariant: "",
			opts:            EngineOptions{MaxSessionActions: 50},
			want:            "hyoka run --prompt-id identity-dp-python-default-credential --config baseline/claude-opus-4.6",
		},
		// ── Option F: Pairwise variant tests ──────────────────────────
		{
			name:            "pairwise baseline variant",
			promptID:        "test-prompt",
			configName:      "python-pairwise/baseline/claude-opus-4.6",
			baseConfigName:  "python-pairwise",
			generatorModel:  "claude-opus-4.6",
			pairwiseVariant: "baseline",
			opts:            EngineOptions{MaxSessionActions: 50},
			want:            "hyoka run --prompt-id test-prompt --config python-pairwise --model claude-opus-4.6 --pairwise-variant baseline",
		},
		{
			name:            "pairwise simple ablation (without-azure)",
			promptID:        "test-prompt",
			configName:      "python-pairwise/without-azure/claude-opus-4.6",
			baseConfigName:  "python-pairwise",
			generatorModel:  "claude-opus-4.6",
			pairwiseVariant: "without-azure",
			opts:            EngineOptions{MaxSessionActions: 50},
			want:            "hyoka run --prompt-id test-prompt --config python-pairwise --model claude-opus-4.6 --pairwise-variant without-azure",
		},
		{
			name:            "pairwise deep MCP variant (quoted slash)",
			promptID:        "test-prompt",
			configName:      "python-pairwise/without-azure/storage_blob_list/claude-opus-4.6",
			baseConfigName:  "python-pairwise",
			generatorModel:  "claude-opus-4.6",
			pairwiseVariant: "without-azure/storage_blob_list",
			opts:            EngineOptions{MaxSessionActions: 50},
			want:            `hyoka run --prompt-id test-prompt --config python-pairwise --model claude-opus-4.6 --pairwise-variant "without-azure/storage_blob_list"`,
		},
		{
			name:            "pairwise variant + skip-review",
			promptID:        "test-prompt",
			configName:      "python-pairwise/without-azure/claude-opus-4.6",
			baseConfigName:  "python-pairwise",
			generatorModel:  "claude-opus-4.6",
			pairwiseVariant: "without-azure",
			opts:            EngineOptions{SkipReview: true, MaxSessionActions: 50},
			want:            "hyoka run --prompt-id test-prompt --config python-pairwise --model claude-opus-4.6 --pairwise-variant without-azure --skip-review",
		},
		{
			name:            "pairwise variant + all flags",
			promptID:        "complex-test",
			configName:      "multi-model/baseline/claude-sonnet-4.6",
			baseConfigName:  "multi-model",
			generatorModel:  "claude-sonnet-4.6",
			pairwiseVariant: "baseline",
			opts:            EngineOptions{SkipReview: true, MonitorResources: true, MaxSessionActions: 75},
			want:            "hyoka run --prompt-id complex-test --config multi-model --model claude-sonnet-4.6 --pairwise-variant baseline --skip-review --monitor-resources --max-session-actions=75",
		},
		{
			name:            "config+model+variant composition",
			promptID:        "identity-dp-python-default-credential",
			configName:      "python-pairwise/without-azure/gpt-5.3-codex",
			baseConfigName:  "python-pairwise",
			generatorModel:  "gpt-5.3-codex",
			pairwiseVariant: "without-azure",
			opts:            EngineOptions{MaxSessionActions: 50},
			want:            "hyoka run --prompt-id identity-dp-python-default-credential --config python-pairwise --model gpt-5.3-codex --pairwise-variant without-azure",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildRerunCommand(tt.promptID, tt.configName, tt.baseConfigName, tt.generatorModel, tt.pairwiseVariant, tt.opts)
			if got != tt.want {
				t.Errorf("buildRerunCommand() = %q, want %q", got, tt.want)
			}
		})
	}
}
