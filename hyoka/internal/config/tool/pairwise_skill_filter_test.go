package tool

import (
	"context"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// repoRootForTest returns the repo root by walking up from this test file's
// location. Avoids hardcoded absolute paths that break in CI.
func repoRootForTest(t *testing.T) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller failed")
	}
	// thisFile is <root>/hyoka/internal/config/tool/pairwise_skill_filter_test.go
	// -> repo root is 4 levels up
	return filepath.Clean(filepath.Join(filepath.Dir(thisFile), "..", "..", "..", ".."))
}

// TestPairwiseDeepVariantSkillsLoadedFilter verifies that when a pairwise
// deep variant excludes a skill via ExcludedSkills, that skill's path does
// NOT appear in the ToolLoadReport.GeneratorSkillDirs() result (the
// session-truth that drives SessionConfig.SkillDirectories).
//
// This test locks the contract that pairwise: deep on a skill_dir actually
// narrows the live tool set, not just the report's display field. It prevents
// regression of the split-brain between ResolveSkills (filter-aware,
// report-only) and validateSkillDirEntry (currently unfiltered, session-truth).
//
// Expected behavior:
// - Baseline variant with NO exclusions: GeneratorSkillDirs returns paths for
//   both markdown-headings AND markdown-lists
// - Variant excluding markdown-headings: GeneratorSkillDirs returns ONLY
//   markdown-lists path (excluding builtins like customize-cloud-agent)
// - Variant excluding markdown-lists: GeneratorSkillDirs returns ONLY
//   markdown-headings path
//
// This test SHOULD initially FAIL on main (demonstrating the bug exists).
// After the fix to validateSkillDirEntry (honoring entry.ExcludedSkills),
// this test will PASS.
func TestPairwiseDeepVariantSkillsLoadedFilter(t *testing.T) {
	// Use absolute path from repo root for test skills
	repoRoot := repoRootForTest(t)

	tests := []struct {
		name            string
		excludedSkills  []string
		wantSkills      []string // skill names that should be loaded
		excludeBuiltins bool     // if true, filter out builtins from assertion
	}{
		{
			name:            "baseline - no exclusions",
			excludedSkills:  nil,
			wantSkills:      []string{"markdown-headings", "markdown-lists"},
			excludeBuiltins: true,
		},
		{
			name:            "exclude markdown-headings",
			excludedSkills:  []string{"markdown-headings"},
			wantSkills:      []string{"markdown-lists"},
			excludeBuiltins: true,
		},
		{
			name:            "exclude markdown-lists",
			excludedSkills:  []string{"markdown-lists"},
			wantSkills:      []string{"markdown-headings"},
			excludeBuiltins: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a tool entry for test-skills with pairwise: deep
			entry := Entry{
				Name:           "test-skills",
				Type:           "skill",
				Source:         "local",
				Path:           "skills/test",
				SkillDir:       true,
				Pairwise:       "deep",
				ExcludedSkills: tt.excludedSkills,
			}

			// Validate the entry
			ctx := context.Background()
			report, err := ValidateAndExpand(ctx, ValidationInput{
				GeneratorTools: []Entry{entry},
				ConfigDir:      repoRoot,
				PluginsDir:     "",
				Emit:           nil,
			})

			if err != nil {
				t.Fatalf("ValidateAndExpand failed: %v", err)
			}

			// Get the loaded skill directories (session-truth)
			skillDirs := report.GeneratorSkillDirs()

			// Extract skill names from paths
			var loadedSkills []string
			for _, dir := range skillDirs {
				skillName := filepath.Base(dir)
				// Filter out builtins if requested
				if tt.excludeBuiltins && (skillName == "customize-cloud-agent" || strings.HasPrefix(skillName, "builtin-")) {
					continue
				}
				loadedSkills = append(loadedSkills, skillName)
			}

			// Check we have the expected number of skills
			if len(loadedSkills) != len(tt.wantSkills) {
				t.Errorf("Expected %d skills loaded, got %d: %v", len(tt.wantSkills), len(loadedSkills), loadedSkills)
				t.Logf("Full skillDirs: %v", skillDirs)
				t.Logf("ToolLoadReport items:")
				for i, item := range report.Items {
					t.Logf("  [%d] Kind=%s Name=%s Parent=%s Status=%s Path=%s", i, item.Kind, item.Name, item.Parent, item.Status, item.Path)
				}
			}

			// Check each expected skill is present
			for _, wantSkill := range tt.wantSkills {
				found := false
				for _, loaded := range loadedSkills {
					if loaded == wantSkill {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected skill %q to be loaded, but it was not in %v", wantSkill, loadedSkills)
				}
			}

			// Check no excluded skills are present
			for _, excludedSkill := range tt.excludedSkills {
				for _, loaded := range loadedSkills {
					if loaded == excludedSkill {
						t.Errorf("Expected skill %q to be excluded, but it was loaded", excludedSkill)
					}
				}
			}
		})
	}
}
