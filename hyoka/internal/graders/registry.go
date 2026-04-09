package graders

import (
"context"
"fmt"
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
cfg, ok := decoded.(*ToolConstraintConfig)
if !ok {
return nil, fmt.Errorf("grader %q: expected *ToolConstraintConfig, got %T", gc.Name, decoded)
}
g, err := NewToolConstraintGrader(gc.Name, toolConstraintConfigToMap(cfg))
if err != nil {
return nil, err
}
return g, nil

default:
return nil, fmt.Errorf("unknown grader kind %q for %q", gc.Kind, gc.Name)
}
}

// InstantiateGraders creates Grader instances from a list of GraderConfig.
func InstantiateGraders(configs []GraderConfig) ([]Grader, error) {
graderInstances := make([]Grader, 0, len(configs))
for _, gc := range configs {
g, err := NewGrader(gc)
if err != nil {
return nil, fmt.Errorf("instantiating grader %q: %w", gc.Name, err)
}
graderInstances = append(graderInstances, g)
}
return graderInstances, nil
}

// RunGraders executes all graders sequentially and returns their results.
func RunGraders(ctx context.Context, graderInstances []Grader, configs []GraderConfig, input GraderInput) []GraderResult {
results := make([]GraderResult, 0, len(graderInstances))

configMap := make(map[string]GraderConfig, len(configs))
for _, c := range configs {
configMap[c.Name] = c
}

for _, g := range graderInstances {
// Set the grader's own config in the input.
ginput := input
if gc, ok := configMap[g.Name()]; ok {
ginput.Config = gc
}

result, err := g.Grade(ctx, ginput)
if err != nil {
result = GraderResult{
Name:    g.Name(),
Kind:    g.Kind(),
Pass:    false,
Score:   0,
Message: fmt.Sprintf("grader execution error: %v", err),
}
}

// Apply weight and gate from config.
if gc, ok := configMap[g.Name()]; ok {
result.Weight = gc.EffectiveWeight()
result.Gate = gc.Gate
}

results = append(results, result)
}

return results
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
