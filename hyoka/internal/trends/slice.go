package trends

import (
"math"
"sort"
)

// DefaultRegressionThreshold is the default score drop (10%) that triggers a regression.
const DefaultRegressionThreshold = 0.1

// TrendSlice contains trend data for prompts matching a property filter.
type TrendSlice struct {
Properties   map[string]string `json:"properties"`
TotalRuns    int               `json:"total_runs"`
Entries      []TrendEntry      `json:"entries"`
PromptTrends []PromptTrend     `json:"prompt_trends"`
RunIDs       []string          `json:"run_ids"`
}

// Regression represents a detected score drop between runs.
type Regression struct {
PromptID      string  `json:"prompt_id"`
Config        string  `json:"config"`
PreviousScore float64 `json:"previous_score"`
CurrentScore  float64 `json:"current_score"`
Delta         float64 `json:"delta"`
GraderName    string  `json:"grader_name,omitempty"` // empty for overall score
}

// SliceTrends filters entries by prompt properties (AND logic) and returns
// trend data for the matching subset.
//
// Example: SliceTrends(entries, map[string]string{"language": "python"})
// returns trends only for Python prompts.
func SliceTrends(entries []TrendEntry, properties map[string]string) *TrendSlice {
filtered := filterByProperties(entries, properties)

sort.Slice(filtered, func(i, j int) bool {
return filtered[i].Timestamp < filtered[j].Timestamp
})

promptTrends, runIDs := buildPromptTrends(filtered)

return &TrendSlice{
Properties:   properties,
TotalRuns:    len(filtered),
Entries:      filtered,
PromptTrends: promptTrends,
RunIDs:       runIDs,
}
}

// filterByProperties returns entries whose Properties map contains all the
// specified key-value pairs (AND logic).
func filterByProperties(entries []TrendEntry, properties map[string]string) []TrendEntry {
if len(properties) == 0 {
// No filter — return a copy of all entries.
out := make([]TrendEntry, len(entries))
copy(out, entries)
return out
}

var out []TrendEntry
for _, e := range entries {
if matchesProperties(e, properties) {
out = append(out, e)
}
}
return out
}

// matchesProperties checks whether an entry's Properties contain all the
// specified key-value pairs.
func matchesProperties(entry TrendEntry, properties map[string]string) bool {
for k, v := range properties {
if entry.Properties[k] != v {
return false
}
}
return true
}

// DetectRegressions compares current and previous entries to find score drops
// exceeding the threshold. The threshold is a fraction (0.0–1.0) of the
// normalized score (score / maxScore). A threshold of 0.1 detects drops of
// 10% or more.
//
// Both overall scores and per-grader scores are checked.
func DetectRegressions(current, previous []TrendEntry, threshold float64) []Regression {
if threshold <= 0 {
threshold = DefaultRegressionThreshold
}

prevScores := buildScoreMap(previous)
currScores := buildScoreMap(current)

var regressions []Regression

for key, curr := range currScores {
prev, ok := prevScores[key]
if !ok {
continue // no previous data to compare
}

// Check overall score regression.
if prev.overall > 0 && curr.overall > 0 {
delta := prev.overall - curr.overall
if delta > threshold {
regressions = append(regressions, Regression{
PromptID:      key.promptID,
Config:        key.config,
PreviousScore: roundTo4(prev.overall),
CurrentScore:  roundTo4(curr.overall),
Delta:         roundTo4(delta),
})
}
}

// Check per-grader regressions.
for grader, prevGraderScore := range prev.graders {
currGraderScore, ok := curr.graders[grader]
if !ok {
continue
}
if prevGraderScore > 0 && currGraderScore > 0 {
delta := prevGraderScore - currGraderScore
if delta > threshold {
regressions = append(regressions, Regression{
PromptID:      key.promptID,
Config:        key.config,
PreviousScore: roundTo4(prevGraderScore),
CurrentScore:  roundTo4(currGraderScore),
Delta:         roundTo4(delta),
GraderName:    grader,
})
}
}
}
}

// Sort for deterministic output.
sort.Slice(regressions, func(i, j int) bool {
if regressions[i].PromptID != regressions[j].PromptID {
return regressions[i].PromptID < regressions[j].PromptID
}
if regressions[i].Config != regressions[j].Config {
return regressions[i].Config < regressions[j].Config
}
return regressions[i].GraderName < regressions[j].GraderName
})

return regressions
}

type scoreKey struct {
promptID string
config   string
}

type scoreData struct {
overall float64
graders map[string]float64
}

// buildScoreMap aggregates entries into average normalized scores per
// prompt+config combination. When multiple entries exist for the same
// prompt+config, scores are averaged.
func buildScoreMap(entries []TrendEntry) map[scoreKey]*scoreData {
type acc struct {
overallSum   float64
overallCount int
graderSum    map[string]float64
graderCount  map[string]int
}

accMap := map[scoreKey]*acc{}
for _, e := range entries {
k := scoreKey{e.PromptID, e.ConfigName}
a, ok := accMap[k]
if !ok {
a = &acc{
graderSum:   map[string]float64{},
graderCount: map[string]int{},
}
accMap[k] = a
}

if e.MaxScore > 0 {
a.overallSum += float64(e.Score) / float64(e.MaxScore)
a.overallCount++
}

for _, gs := range e.GraderScores {
a.graderSum[gs.GraderName] += gs.Score
a.graderCount[gs.GraderName]++
}
}

result := make(map[scoreKey]*scoreData, len(accMap))
for k, a := range accMap {
sd := &scoreData{graders: map[string]float64{}}
if a.overallCount > 0 {
sd.overall = a.overallSum / float64(a.overallCount)
}
for g, sum := range a.graderSum {
if a.graderCount[g] > 0 {
sd.graders[g] = sum / float64(a.graderCount[g])
}
}
result[k] = sd
}

return result
}

func roundTo4(v float64) float64 {
return math.Round(v*10000) / 10000
}
