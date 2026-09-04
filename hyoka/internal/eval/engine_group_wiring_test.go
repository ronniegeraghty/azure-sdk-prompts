package eval

import (
	"context"
	"testing"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/prompt"
)

// TestPromptGroupPropagatesToEvalReport is an observable-wiring test
// (#587 trap) for PR #606. It proves that a prompt's `group` frontmatter
// field actually reaches the EvalReport.PromptMeta["group"] runtime payload
// that reports and the site consume.
//
// The wire-up point is hyoka/internal/eval/engine_eval.go:78-80 — if a
// future refactor drops that block, unit tests on the parser/validator
// would still pass while the site-visible grouping silently breaks.
//
// This runs the full engine.Run pipeline with a StubRunner so we hit the
// real PromptMeta construction, then asserts on the exact runtime payload.
func TestPromptGroupPropagatesToEvalReport(t *testing.T) {
	tests := []struct {
		name      string
		group     string
		wantKey   bool   // should "group" key exist in PromptMeta?
		wantValue string // expected value if key exists
	}{
		{
			name:      "group set propagates to PromptMeta",
			group:     "crud-operations",
			wantKey:   true,
			wantValue: "crud-operations",
		},
		{
			name:    "empty group omitted from PromptMeta",
			group:   "",
			wantKey: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outputDir := t.TempDir()
			engine := NewEngine(&StubRunner{}, quietOpts(EngineOptions{
				Workers:   1,
				OutputDir: outputDir,
			}))

			prompts := []*prompt.Prompt{{
				ID:    "group-wiring-check",
				Group: tt.group,
				Properties: map[string]string{
					"service":  "storage",
					"language": "python",
					"plane":    "data-plane",
					"category": "crud",
				},
			}}
			configs := []config.ToolConfig{{
				Name:      "test-config",
				Generator: &config.GeneratorConfig{Model: "gpt-4"},
			}}

			summary, err := engine.Run(context.Background(), prompts, configs)
			if err != nil {
				t.Fatalf("engine.Run: %v", err)
			}
			if len(summary.Results) != 1 {
				t.Fatalf("expected 1 result, got %d", len(summary.Results))
			}
			r := summary.Results[0]
			if r.PromptMeta == nil {
				t.Fatalf("PromptMeta is nil — wiring is dead")
			}

			got, ok := r.PromptMeta["group"]
			if tt.wantKey {
				if !ok {
					t.Fatalf(`PromptMeta["group"] missing — frontmatter group %q never reached the engine wiring point (engine_eval.go:78-80)`, tt.group)
				}
				gotStr, _ := got.(string)
				if gotStr != tt.wantValue {
					t.Errorf(`PromptMeta["group"] = %q, want %q`, gotStr, tt.wantValue)
				}
			} else {
				if ok {
					t.Errorf(`PromptMeta["group"] = %v, want key absent for empty Group`, got)
				}
			}
		})
	}
}
