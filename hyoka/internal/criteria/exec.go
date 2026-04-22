// Grader instantiation and execution driver used by the eval engine.
//
// These helpers operate on a set of runtime GraderConfig values (typically
// produced by UnifiedGraderEntry.ToRuntimeConfig) and bridge from the
// file-level matching layer (criteria) down to the grader-type layer
// (criteria/graders) for per-grader construction and scoring.
package criteria

import (
	"context"
	"fmt"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria/graders"
)

// InstantiateGraders creates Grader instances from a list of GraderConfig.
func InstantiateGraders(configs []graders.GraderConfig) ([]graders.Grader, error) {
	graderInstances := make([]graders.Grader, 0, len(configs))
	for _, gc := range configs {
		g, err := graders.NewGrader(gc)
		if err != nil {
			return nil, fmt.Errorf("instantiating grader %q: %w", gc.Name, err)
		}
		graderInstances = append(graderInstances, g)
	}
	return graderInstances, nil
}

// RunGraders executes all graders sequentially and returns their results.
func RunGraders(ctx context.Context, graderInstances []graders.Grader, configs []graders.GraderConfig, input graders.GraderInput) []graders.GraderResult {
	results := make([]graders.GraderResult, 0, len(graderInstances))

	configMap := make(map[string]graders.GraderConfig, len(configs))
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
			result = graders.GraderResult{
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
