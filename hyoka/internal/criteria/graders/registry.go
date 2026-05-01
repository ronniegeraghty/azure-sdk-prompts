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

	switch gc.Kind {
	case KindFile:
		slog.Warn("grader kind 'file' is deprecated; use 'output_check' with require_files instead", "name", gc.Name, "kind", gc.Kind)
		cfg, ok := decoded.(*FileConfig)
		if !ok {
			return nil, fmt.Errorf("grader %q: expected *FileConfig, got %T", gc.Name, decoded)
		}
		g, err := NewFileGrader(gc.Name, cfg)
		if err != nil {
			return nil, err
		}
		return g, nil

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

	case KindBehavior:
		slog.Warn("grader kind 'behavior' is deprecated; use 'tool' kind instead", "name", gc.Name, "kind", gc.Kind)
		cfg, ok := decoded.(*BehaviorConfig)
		if !ok {
			return nil, fmt.Errorf("grader %q: expected *BehaviorConfig, got %T", gc.Name, decoded)
		}
		g, err := NewBehaviorGrader(gc.Name, behaviorConfigToMap(cfg))
		if err != nil {
			return nil, err
		}
		return g, nil

	case KindActionSequence:
		cfg, ok := decoded.(*ActionSequenceConfig)
		if !ok {
			return nil, fmt.Errorf("grader %q: expected *ActionSequenceConfig, got %T", gc.Name, decoded)
		}
		g, err := NewActionSequenceGrader(gc.Name, actionSequenceConfigToMap(cfg))
		if err != nil {
			return nil, err
		}
		return g, nil

	case KindToolConstraint:
		slog.Warn("grader kind 'tool_constraint' is deprecated; use 'tool' kind instead", "name", gc.Name, "kind", gc.Kind)
		cfg, ok := decoded.(*ToolConstraintConfig)
		if !ok {
			return nil, fmt.Errorf("grader %q: expected *ToolConstraintConfig, got %T", gc.Name, decoded)
		}
		g, err := NewToolConstraintGrader(gc.Name, toolConstraintConfigToMap(cfg))
		if err != nil {
			return nil, err
		}
		return g, nil

	case KindOutputCheck:
		cfg, ok := decoded.(*OutputCheckConfig)
		if !ok {
			return nil, fmt.Errorf("grader %q: expected *OutputCheckConfig, got %T", gc.Name, decoded)
		}
		g, err := NewOutputCheckGrader(gc.Name, cfg)
		if err != nil {
			return nil, err
		}
		return g, nil

	case KindToolUsage:
		slog.Warn("grader kind 'tool_usage' is deprecated; use 'tool' kind instead", "name", gc.Name, "kind", gc.Kind)
		cfg, ok := decoded.(*ToolUsageConfig)
		if !ok {
			return nil, fmt.Errorf("grader %q: expected *ToolUsageConfig, got %T", gc.Name, decoded)
		}
		return NewToolUsageGrader(gc.Name, cfg)

	default:
		return nil, fmt.Errorf("unknown grader kind %q for %q", gc.Kind, gc.Name)
	}
}

// Config-to-map adapters for graders that still accept map[string]any.

func promptConfigToMap(cfg *PromptConfig) map[string]any {
	m := map[string]any{
		"model":  cfg.Model,
		"rubric": cfg.Rubric,
	}
	return m
}

func behaviorConfigToMap(cfg *BehaviorConfig) map[string]any {
	m := map[string]any{}
	if len(cfg.RequiredTools) > 0 {
		m["required_tools"] = cfg.RequiredTools
	}
	if len(cfg.ForbiddenTools) > 0 {
		m["forbidden_tools"] = cfg.ForbiddenTools
	}
	if cfg.MaxTurns > 0 {
		m["max_turns"] = cfg.MaxTurns
	}
	return m
}

func actionSequenceConfigToMap(cfg *ActionSequenceConfig) map[string]any {
	return map[string]any{
		"expected_actions": cfg.ExpectedActions,
	}
}

func toolConstraintConfigToMap(cfg *ToolConstraintConfig) map[string]any {
	m := map[string]any{}
	if len(cfg.Required) > 0 {
		m["required"] = cfg.Required
	}
	if len(cfg.Forbidden) > 0 {
		m["forbidden"] = cfg.Forbidden
	}
	if cfg.MinCalls != nil {
		m["min_calls"] = cfg.MinCalls
	}
	if cfg.MaxCalls != nil {
		m["max_calls"] = cfg.MaxCalls
	}
	return m
}
