// Unified matching, partitioning, review-bucket construction, and legacy
// graders.GraderConfig translation for the Phase 2 execution path (issue #625).
//
// The engine calls these helpers to:
//   1. Match grader entries from a Bundle against a prompt's properties,
//      honoring hierarchical file/group/grader when resolution.
//   2. Partition the matched entries into prompt (LLM-review) and typed
//      (output_check, file, program, ...) entries.
//   3. Build review buckets from the prompt entries — one bucket per
//      isolated grader/group plus a combined bucket for the rest, or a
//      single combined bucket in combined mode.
//   4. Bridge a UnifiedGraderEntry to the runtime graders.GraderConfig used by
//      NewGrader.
//
// This file replaces the matching/bucket logic in internal/criteria (ported
// here so internal/graders is self-contained). The criteria package remains
// on disk during Phase 2 and is deleted in Phase 3.
package criteria

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/criteria/graders"
)

// ReviewMode constants used by the engine to select bucket construction.
const (
	ReviewModeCombined = "combined"
	ReviewModeIsolated = "isolated"
)

// MatchedUnifiedEntry pairs a matched UnifiedGraderEntry with metadata about
// the group it belongs to (if any). This lets BuildUnifiedReviewBuckets
// honor group-level isolation without losing the grader's own isolate flag.
type MatchedUnifiedEntry struct {
	Entry        UnifiedGraderEntry
	GroupName    string // empty for top-level (non-grouped) entries
	GroupIsolate bool
	Source       string // originating YAML file path
}

// MatchingUnifiedEntries returns every grader entry from bundle.Configs
// whose file/group/grader when conditions match props. File-level and
// group-level when conditions are merged (child wins) before matching.
//
// Returned slice order is stable: file-walk order → top-level graders →
// groups → entries within each group.
func MatchingUnifiedEntries(bundle *Bundle, props map[string]string) []MatchedUnifiedEntry {
	if bundle == nil {
		return nil
	}
	var out []MatchedUnifiedEntry
	for _, gc := range bundle.Configs {
		if !matchesUnifiedWhen(gc.When, props) {
			continue
		}
		for _, e := range gc.Graders {
			when := mergeUnifiedWhen(gc.When, e.When)
			if !matchesUnifiedWhen(when, props) {
				continue
			}
			out = append(out, MatchedUnifiedEntry{Entry: e, Source: gc.Source})
		}
		for _, grp := range gc.Groups {
			grpWhen := mergeUnifiedWhen(gc.When, grp.When)
			if !matchesUnifiedWhen(grpWhen, props) {
				continue
			}
			for _, e := range grp.Graders {
				when := mergeUnifiedWhen(grpWhen, e.When)
				if !matchesUnifiedWhen(when, props) {
					continue
				}
				out = append(out, MatchedUnifiedEntry{
					Entry:        e,
					GroupName:    grp.Name,
					GroupIsolate: grp.Isolate,
					Source:       gc.Source,
				})
			}
		}
	}
	return out
}

// PartitionMatched splits matched entries into prompt (type=prompt, review
// panel input) and typed (everything else, NewGrader input) slices while
// preserving their relative order.
func PartitionMatched(matched []MatchedUnifiedEntry) (promptEntries, typedEntries []MatchedUnifiedEntry) {
	for _, m := range matched {
		if m.Entry.Type == graders.KindPrompt {
			promptEntries = append(promptEntries, m)
		} else {
			typedEntries = append(typedEntries, m)
		}
	}
	return promptEntries, typedEntries
}

// HasUnifiedIsolation reports whether any matched prompt entry (or its
// owning group) requests isolation. Used by the engine to decide whether
// isolated mode has any work to do.
func HasUnifiedIsolation(matched []MatchedUnifiedEntry) bool {
	for _, m := range matched {
		if m.GroupIsolate || m.Entry.Isolate {
			return true
		}
	}
	return false
}

// FormatUnifiedPromptEntries renders prompt-type entries as a text block
// suitable for injection into a review prompt.
//
// The output uses the entry name as a bolded section heading (NOT a numbered
// outer item). When Checks is non-empty, the optional Prompt acts as
// preamble and each check is rendered as a numbered top-level item:
//
//	**Name**
//	preamble text
//	1. check1
//	2. check2
//
// When Checks is empty, the entry is rendered as context only:
//
//	**Name**: Prompt text
//
// (or just `**Name**` when no Prompt). Multiple entries are separated by a
// blank line.
//
// Why no outer numbering: the LLM judge counts numbered items as scorable
// criteria. A leading `1. **Name**` line caused the judge to treat the
// section heading itself as a check, producing the intermittent
// "3 points returned instead of 2" off-by-one bug.
func FormatUnifiedPromptEntries(entries []UnifiedGraderEntry) string {
	if len(entries) == 0 {
		return ""
	}
	var blocks []string
	for _, e := range entries {
		var b strings.Builder
		if len(e.Checks) > 0 {
			fmt.Fprintf(&b, "**%s**\n", e.Name)
			if preamble := strings.TrimSpace(e.Prompt); preamble != "" {
				fmt.Fprintf(&b, "%s\n", preamble)
			}
			for j, c := range e.Checks {
				fmt.Fprintf(&b, "%d. %s", j+1, strings.TrimSpace(c))
				if j < len(e.Checks)-1 {
					b.WriteString("\n")
				}
			}
			blocks = append(blocks, b.String())
			continue
		}
		// No checks — render as context only.
		fmt.Fprintf(&b, "**%s**", e.Name)
		if p := strings.TrimSpace(e.Prompt); p != "" {
			fmt.Fprintf(&b, ": %s", p)
		}
		blocks = append(blocks, b.String())
	}
	return strings.Join(blocks, "\n\n")
}

// MergeUnifiedCriteria combines rendered prompt entries with a prompt's own
// evaluation criteria text. Mirrors criteria.MergeCriteria's output.
func MergeUnifiedCriteria(entries []UnifiedGraderEntry, promptCriteria string) string {
	parts := make([]string, 0, 2)
	if formatted := FormatUnifiedPromptEntries(entries); formatted != "" {
		parts = append(parts, "### Attribute-Matched Criteria\n\n"+formatted)
	}
	promptCriteria = strings.TrimSpace(promptCriteria)
	if promptCriteria != "" {
		parts = append(parts, "### Prompt-Specific Criteria\n\n"+promptCriteria)
	}
	return strings.Join(parts, "\n")
}

// BuildUnifiedReviewBuckets constructs review buckets from the prompt-type
// entries matched against a prompt. Each bucket's Criteria field is the
// rendered text that will be handed to the review panel.
//
// Prompt-frontmatter criteria and criteria-file entries are ALWAYS separated
// into distinct buckets, regardless of mode:
//   - promptCriteria (if non-empty) becomes its own bucket named
//     "Criteria from prompt file"
//   - Matched criteria-file entries each get their own bucket, named after
//     the entry's name field
//
// In combined mode (default), each matched criteria-file entry becomes its
// own top-level bucket (one bucket per entry).
//
// In isolated mode:
//   - Each prompt entry marked Isolate goes into its own bucket (named
//     after the entry).
//   - Each group marked Isolate goes into one bucket for the whole group
//     (named after the group).
//   - All remaining entries share one "combined" bucket. If nothing remains,
//     the combined bucket is omitted.
//
// When isolated mode is requested but nothing is marked isolate, matched
// criteria-file entries fall back to per-entry buckets (same as combined mode).
func BuildUnifiedReviewBuckets(matched []MatchedUnifiedEntry, promptCriteria, mode string) []graders.ReviewBucket {
	var buckets []graders.ReviewBucket

	// Prompt-frontmatter criteria always get their own bucket, regardless of mode.
	promptCriteria = strings.TrimSpace(promptCriteria)
	if promptCriteria != "" {
		buckets = append(buckets, graders.ReviewBucket{
			Name:     "Criteria from prompt file",
			Criteria: MergeUnifiedCriteria(nil, promptCriteria),
		})
	}

	// Build buckets for matched criteria-file entries.
	if mode != ReviewModeIsolated || !HasUnifiedIsolation(matched) {
		// Combined mode (default): each matched entry gets its own bucket.
		for _, m := range matched {
			buckets = append(buckets, graders.ReviewBucket{
				Name:     bucketName(m.Entry.Name, len(buckets)),
				Criteria: MergeUnifiedCriteria([]UnifiedGraderEntry{m.Entry}, ""),
			})
		}
		return buckets
	}

	// Isolated mode: honor isolate flags.
	var leftover []UnifiedGraderEntry

	type groupBag struct {
		name    string
		isolate bool
		entries []UnifiedGraderEntry
	}
	var groupsOrder []string
	groupsByKey := map[string]*groupBag{}
	anonSeq := 0

	for _, m := range matched {
		if m.GroupName == "" && !m.GroupIsolate {
			if m.Entry.Isolate {
				buckets = append(buckets, graders.ReviewBucket{
					Name:     bucketName(m.Entry.Name, len(buckets)),
					Criteria: MergeUnifiedCriteria([]UnifiedGraderEntry{m.Entry}, ""),
				})
			} else {
				leftover = append(leftover, m.Entry)
			}
			continue
		}
		key := strings.TrimSpace(m.GroupName)
		if key == "" {
			key = fmt.Sprintf("__anon_group_%d", anonSeq)
			anonSeq++
		}
		bag, ok := groupsByKey[key]
		if !ok {
			bag = &groupBag{name: m.GroupName, isolate: m.GroupIsolate}
			groupsByKey[key] = bag
			groupsOrder = append(groupsOrder, key)
		}
		bag.entries = append(bag.entries, m.Entry)
	}

	for _, key := range groupsOrder {
		bag := groupsByKey[key]
		if bag.isolate {
			name := bag.name
			if name == "" {
				name = fmt.Sprintf("group-%d", len(buckets))
			}
			buckets = append(buckets, graders.ReviewBucket{
				Name:     bucketName(name, len(buckets)),
				Criteria: MergeUnifiedCriteria(append([]UnifiedGraderEntry(nil), bag.entries...), ""),
			})
			continue
		}
		for _, e := range bag.entries {
			if e.Isolate {
				buckets = append(buckets, graders.ReviewBucket{
					Name:     bucketName(e.Name, len(buckets)),
					Criteria: MergeUnifiedCriteria([]UnifiedGraderEntry{e}, ""),
				})
			} else {
				leftover = append(leftover, e)
			}
		}
	}

	// Leftover criteria-file entries go into a "combined" bucket.
	if len(leftover) > 0 {
		buckets = append(buckets, graders.ReviewBucket{
			Name:     "combined",
			Criteria: MergeUnifiedCriteria(leftover, ""),
		})
	}

	return buckets
}

// combinedCriteriaFileBucket constructs a "combined" bucket containing only
// matched criteria-file entries (no prompt-frontmatter criteria).
func combinedCriteriaFileBucket(matched []MatchedUnifiedEntry) graders.ReviewBucket {
	entries := make([]UnifiedGraderEntry, 0, len(matched))
	for _, m := range matched {
		entries = append(entries, m.Entry)
	}
	return graders.ReviewBucket{
		Name:     "combined",
		Criteria: MergeUnifiedCriteria(entries, ""),
	}
}

func bucketName(raw string, index int) string {
	n := strings.TrimSpace(raw)
	if n == "" {
		return fmt.Sprintf("bucket-%d", index)
	}
	return n
}

// ToRuntimeConfig converts a UnifiedGraderEntry into the graders.GraderConfig shape
// consumed by NewGrader. Typed entries (type != prompt) carry their payload
// in Details; prompt entries are not expected to flow through NewGrader
// under the Phase 2 design (they feed the review panel instead).
//
// The returned graders.GraderConfig has Kind=Type, Config=Details, Weight and Name
// copied, Gate=false (Phase 2 locked decision: no gating), and an empty
// WhenMap — matching has already been resolved by MatchingUnifiedEntries.
func (e UnifiedGraderEntry) ToRuntimeConfig() graders.GraderConfig {
	return graders.GraderConfig{
		Kind:   e.Type,
		Name:   e.Name,
		Config: cloneYAMLNode(e.Details),
		Weight: e.EffectiveWeight(),
	}
}

// cloneYAMLNode returns a deep copy of n so the runtime config can be
// decoded without mutating the source Details node held by the Bundle.
// yaml.Node contains slices of child pointers; a shallow copy would share
// state with every eval that reuses the same Bundle entry.
func cloneYAMLNode(n yaml.Node) yaml.Node {
	out := yaml.Node{
		Kind:        n.Kind,
		Style:       n.Style,
		Tag:         n.Tag,
		Value:       n.Value,
		Anchor:      n.Anchor,
		HeadComment: n.HeadComment,
		LineComment: n.LineComment,
		FootComment: n.FootComment,
		Line:        n.Line,
		Column:      n.Column,
	}
	if len(n.Content) > 0 {
		out.Content = make([]*yaml.Node, len(n.Content))
		for i, c := range n.Content {
			if c == nil {
				continue
			}
			cloned := cloneYAMLNode(*c)
			out.Content[i] = &cloned
		}
	}
	if n.Alias != nil {
		cloned := cloneYAMLNode(*n.Alias)
		out.Alias = &cloned
	}
	return out
}
