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

// GraderHooks are optional per-grader lifecycle callbacks. Either field may
// be nil, in which case the corresponding event is simply not invoked.
//
// OnStart fires immediately before a grader's Grade() method is called.
// OnComplete fires after Grade() returns, with the final (post-error-wrap,
// post-weight/gate) result value. Neither hook returns an error — they are
// UX signals only and must not influence grading outcomes.
type GraderHooks struct {
	OnStart    func(g graders.Grader)
	OnComplete func(g graders.Grader, result graders.GraderResult)
}

// RunGraders executes all graders sequentially and returns their results.
func RunGraders(ctx context.Context, graderInstances []graders.Grader, configs []graders.GraderConfig, input graders.GraderInput) []graders.GraderResult {
	return RunGradersWithHooks(ctx, graderInstances, configs, input, GraderHooks{})
}

// RunGradersWithHooks executes graders sequentially, invoking the supplied
// hooks around each grader so callers (e.g. the progress display) can render
// per-grader lifecycle events. Hook fields may be nil.
func RunGradersWithHooks(ctx context.Context, graderInstances []graders.Grader, configs []graders.GraderConfig, input graders.GraderInput, hooks GraderHooks) []graders.GraderResult {
	results := make([]graders.GraderResult, 0, len(graderInstances))

	configMap := make(map[string]graders.GraderConfig, len(configs))
	for _, c := range configs {
		configMap[c.Name] = c
	}

	for _, g := range graderInstances {
		if hooks.OnStart != nil {
			hooks.OnStart(g)
		}

		// Track completion state to ensure OnComplete fires even on panic (#819).
		completeFired := false
		gc, _ := configMap[g.Name()]
		defer func(grader graders.Grader, gcfg graders.GraderConfig) {
			if !completeFired && hooks.OnComplete != nil {
				// Grader interrupted or panicked — synthesize a failing
				// result that still satisfies the "every grader must emit
				// ≥ 1 Point" invariant so the renderer doesn't fall back
				// to "PASS"/"100%".
				hooks.OnComplete(grader, graders.NewErrorResult(
					grader.Kind(), grader.Name(), gcfg, "grader interrupted or panicked",
				))
			}
		}(g, gc)

		// Set the grader's own config in the input.
		ginput := input
		if ok := gc.Name != ""; ok {
			ginput.Config = gc
		}

		result, err := g.Grade(ctx, ginput)
		if err != nil {
			// Synthesize a failing result with a single failure Point so the
			// site never sees a Points-less GraderResult (Phase 3 invariant).
			result = graders.NewErrorResult(
				g.Kind(), g.Name(), gc,
				fmt.Sprintf("grader execution error: %v", err),
			)
		}

		// Apply weight and gate from config.
		if gc, ok := configMap[g.Name()]; ok {
			result.Weight = gc.EffectiveWeight()
			result.Gate = gc.Gate
		}

		results = append(results, result)

		if hooks.OnComplete != nil {
			hooks.OnComplete(g, result)
		}
		completeFired = true
	}

	return results
}
