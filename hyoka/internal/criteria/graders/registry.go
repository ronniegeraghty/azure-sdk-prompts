package graders

import (
	"fmt"
	"log/slog"
)

// NewGrader creates a Grader instance from a GraderConfig by dispatching
// on the Kind field. Returns an error for unknown kinds or invalid configs.
func NewGrader(gc GraderConfig) (Grader, error) {
	decoded, err := gc.DecodeConfig()
	if err != nil {
		return nil, fmt.Errorf("decoding config for grader %q: %w", gc.Name, err)
	}

	// Emit deprecation warning if gate is set.
	if gc.Gate {
		slog.Warn("grader field 'gate' is deprecated; gate semantics are no longer enforced — use 'tool' grader check kinds or separate explicit graders instead", "name", gc.Name)
	}

	switch gc.Kind {
	case KindProgram:
		cfg, ok := decoded.(*ProgramConfig)
		if !ok {
			return nil, fmt.Errorf("grader %q: expected *ProgramConfig, got %T", gc.Name, decoded)
		}
		g, err := NewProgramGrader(gc.Name, cfg)
		if err != nil {
			return nil, err
		}
		return g, nil

	case KindPrompt:
		cfg, ok := decoded.(*PromptConfig)
		if !ok {
			return nil, fmt.Errorf("grader %q: expected *PromptConfig, got %T", gc.Name, decoded)
		}
		g, err := NewPromptGrader(gc.Name, promptConfigToMap(cfg))
		if err != nil {
			return nil, err
		}
		return &PromptGraderAdapter{inner: g}, nil

	case KindTool:
		cfg, ok := decoded.(*ToolConfig)
		if !ok {
			return nil, fmt.Errorf("grader %q: expected *ToolConfig, got %T", gc.Name, decoded)
		}
		return NewToolGrader(gc.Name, cfg)

	case "workspace":
		cfg, ok := decoded.(*WorkspaceConfig)
		if !ok {
			return nil, fmt.Errorf("grader %q: expected *WorkspaceConfig, got %T", gc.Name, decoded)
		}
		return NewWorkspaceGrader(gc.Name, cfg)

	case "activity":
		cfg, ok := decoded.(*ActivityConfig)
		if !ok {
			return nil, fmt.Errorf("grader %q: expected *ActivityConfig, got %T", gc.Name, decoded)
		}
		return NewActivityGrader(gc.Name, cfg)

	default:
		return nil, fmt.Errorf("unknown grader kind %q for %q", gc.Kind, gc.Name)
	}
}

// Config-to-map adapter for graders that still accept map[string]any.
func promptConfigToMap(cfg *PromptConfig) map[string]any {
	m := map[string]any{
		"model":  cfg.Model,
		"rubric": cfg.Rubric,
	}
	return m
}
