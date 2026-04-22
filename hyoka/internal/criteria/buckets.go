package criteria

import (
	"fmt"
	"strings"
)

// ReviewMode constants used by the engine to select bucket construction.
const (
	ReviewModeCombined = "combined"
	ReviewModeIsolated = "isolated"
)

// MatchedGrader pairs a matched grader entry with metadata about the group it
// belongs to (if any). This lets BuildReviewBuckets honor group-level isolation
// without losing the grader's own isolate flag.
type MatchedGrader struct {
	Entry        GraderEntry
	GroupName    string // empty for top-level (non-grouped) graders
	GroupIsolate bool
}

// ReviewBucket is a unit of work for the reviewer. Each bucket becomes one
// reviewer session per panel model. In combined mode there is exactly one
// bucket. In isolated mode there is one bucket per isolated grader/group plus
// (optionally) one shared bucket for the remaining criteria.
type ReviewBucket struct {
	// Name is a stable identifier used for logging and to disambiguate
	// criterion names when bucket results are merged. The default shared
	// bucket is named "combined".
	Name string

	// Graders are the attribute-matched graders included in this bucket.
	// The combined bucket may also carry PromptCriteria; isolated buckets
	// carry only their owning grader(s).
	Graders []GraderEntry

	// PromptCriteria is the prompt's own criteria text, attached only to
	// the combined bucket so it is reviewed exactly once.
	PromptCriteria string
}

// FormatCriteria renders a bucket's criteria as a single string suitable for
// passing to the reviewer. Mirrors MergeCriteria's behavior so combined-mode
// output is byte-identical to the legacy single-string path.
func (b ReviewBucket) FormatCriteria() string {
	return MergeCriteria(b.Graders, b.PromptCriteria)
}

// MatchingGradersWithIsolation returns the matched graders along with their
// group context so callers can construct isolation buckets. The semantics of
// when-condition resolution match MatchingGraders exactly.
func MatchingGradersWithIsolation(configs []GraderConfig, props map[string]string) []MatchedGrader {
	var out []MatchedGrader
	for _, gc := range configs {
		if !matchesWhen(gc.When, props) {
			continue
		}
		for _, g := range gc.Graders {
			when := mergeWhen(gc.When, g.When)
			if !matchesWhen(when, props) {
				continue
			}
			out = append(out, MatchedGrader{Entry: g})
		}
		for _, grp := range gc.Groups {
			grpWhen := mergeWhen(gc.When, grp.When)
			if !matchesWhen(grpWhen, props) {
				continue
			}
			for _, g := range grp.Graders {
				when := mergeWhen(grpWhen, g.When)
				if !matchesWhen(when, props) {
					continue
				}
				out = append(out, MatchedGrader{
					Entry:        g,
					GroupName:    grp.Name,
					GroupIsolate: grp.Isolate,
				})
			}
		}
	}
	return out
}

// HasIsolation reports whether any matched grader (or its group) requests
// isolation. Used by the engine to decide whether isolated mode actually has
// any work to do.
func HasIsolation(matched []MatchedGrader) bool {
	for _, m := range matched {
		if m.GroupIsolate || m.Entry.Isolate {
			return true
		}
	}
	return false
}

// groupIdentity returns a stable bucket key for an isolated group. Anonymous
// groups (no name) are bucketed by their order via the synthetic key suffix
// produced by BuildReviewBuckets.
func groupIdentity(name string) string {
	n := strings.TrimSpace(name)
	if n == "" {
		return ""
	}
	return n
}

// BuildReviewBuckets constructs the list of review buckets according to mode.
//
// In combined mode (default), it returns a single bucket containing every
// matched grader plus the prompt's own criteria — byte-identical to the
// legacy single-session behavior.
//
// In isolated mode:
//   - Each grader marked Isolate goes into its own bucket.
//   - Each group marked Isolate goes into a single bucket for the whole group.
//     A grader's own Isolate flag is ignored when its group is isolated.
//   - All remaining graders + the prompt criteria share a single "combined"
//     bucket. If there are no remaining items, no combined bucket is added.
//
// The returned slice is never empty: when isolated mode is requested but
// nothing is marked isolate, the function falls back to the combined-mode
// single bucket so the engine can warn and proceed.
func BuildReviewBuckets(matched []MatchedGrader, promptCriteria, mode string) []ReviewBucket {
	if mode != ReviewModeIsolated {
		return []ReviewBucket{combinedBucket(matched, promptCriteria)}
	}
	if !HasIsolation(matched) {
		return []ReviewBucket{combinedBucket(matched, promptCriteria)}
	}

	var buckets []ReviewBucket
	var leftover []GraderEntry

	// Group graders by their owning group so isolated groups go as a unit.
	type groupBag struct {
		name    string
		isolate bool
		graders []GraderEntry
	}
	var groupsOrder []string
	groupsByKey := map[string]*groupBag{}
	groupKeySeq := 0

	for _, m := range matched {
		if m.GroupName == "" && !m.GroupIsolate {
			// Top-level grader.
			if m.Entry.Isolate {
				buckets = append(buckets, ReviewBucket{
					Name:    bucketNameFor(m.Entry.Name, len(buckets)),
					Graders: []GraderEntry{m.Entry},
				})
			} else {
				leftover = append(leftover, m.Entry)
			}
			continue
		}
		// Grader belonging to a group (named or anonymous).
		key := groupIdentity(m.GroupName)
		if key == "" {
			key = fmt.Sprintf("__anon_group_%d", groupKeySeq)
			groupKeySeq++
		}
		bag, ok := groupsByKey[key]
		if !ok {
			bag = &groupBag{name: m.GroupName, isolate: m.GroupIsolate}
			groupsByKey[key] = bag
			groupsOrder = append(groupsOrder, key)
		}
		bag.graders = append(bag.graders, m.Entry)
	}

	for _, key := range groupsOrder {
		bag := groupsByKey[key]
		if bag.isolate {
			name := bag.name
			if name == "" {
				name = fmt.Sprintf("group-%d", len(buckets))
			}
			buckets = append(buckets, ReviewBucket{
				Name:    bucketNameFor(name, len(buckets)),
				Graders: append([]GraderEntry(nil), bag.graders...),
			})
			continue
		}
		// Non-isolated group: per-grader isolate still applies.
		for _, g := range bag.graders {
			if g.Isolate {
				buckets = append(buckets, ReviewBucket{
					Name:    bucketNameFor(g.Name, len(buckets)),
					Graders: []GraderEntry{g},
				})
			} else {
				leftover = append(leftover, g)
			}
		}
	}

	if len(leftover) > 0 || strings.TrimSpace(promptCriteria) != "" {
		buckets = append(buckets, ReviewBucket{
			Name:           "combined",
			Graders:        leftover,
			PromptCriteria: promptCriteria,
		})
	}
	if len(buckets) == 0 {
		return []ReviewBucket{combinedBucket(matched, promptCriteria)}
	}
	return buckets
}

func combinedBucket(matched []MatchedGrader, promptCriteria string) ReviewBucket {
	graders := make([]GraderEntry, 0, len(matched))
	for _, m := range matched {
		graders = append(graders, m.Entry)
	}
	return ReviewBucket{Name: "combined", Graders: graders, PromptCriteria: promptCriteria}
}

func bucketNameFor(raw string, index int) string {
	n := strings.TrimSpace(raw)
	if n == "" {
		return fmt.Sprintf("bucket-%d", index)
	}
	return n
}
