package trends

import (
"testing"
)

// --- SliceTrends tests ---

func TestSliceTrendsFiltersByLanguage(t *testing.T) {
entries := []TrendEntry{
{RunID: "run1", PromptID: "identity-dp-python-auth", ConfigName: "baseline", Success: true, Properties: map[string]string{"language": "python", "service": "identity"}},
{RunID: "run1", PromptID: "kv-dp-dotnet-crud", ConfigName: "baseline", Success: true, Properties: map[string]string{"language": "dotnet", "service": "key-vault"}},
{RunID: "run1", PromptID: "storage-dp-python-blob", ConfigName: "baseline", Success: false, Properties: map[string]string{"language": "python", "service": "storage"}},
}

slice := SliceTrends(entries, map[string]string{"language": "python"})

if slice.TotalRuns != 2 {
t.Fatalf("expected 2 entries, got %d", slice.TotalRuns)
}
if len(slice.PromptTrends) != 2 {
t.Fatalf("expected 2 prompt trends, got %d", len(slice.PromptTrends))
}
for _, e := range slice.Entries {
if e.Properties["language"] != "python" {
t.Errorf("unexpected language in slice: %s", e.Properties["language"])
}
}
if slice.Properties["language"] != "python" {
t.Error("slice Properties should reflect the filter")
}
}

func TestSliceTrendsMultipleProperties(t *testing.T) {
entries := []TrendEntry{
{RunID: "run1", PromptID: "p1", ConfigName: "baseline", Properties: map[string]string{"language": "python", "service": "identity"}},
{RunID: "run1", PromptID: "p2", ConfigName: "baseline", Properties: map[string]string{"language": "python", "service": "key-vault"}},
{RunID: "run1", PromptID: "p3", ConfigName: "baseline", Properties: map[string]string{"language": "dotnet", "service": "identity"}},
}

slice := SliceTrends(entries, map[string]string{"language": "python", "service": "identity"})

if slice.TotalRuns != 1 {
t.Fatalf("expected 1 entry for python+identity, got %d", slice.TotalRuns)
}
if slice.Entries[0].PromptID != "p1" {
t.Errorf("expected p1, got %s", slice.Entries[0].PromptID)
}
}

func TestSliceTrendsEmptyFilter(t *testing.T) {
entries := []TrendEntry{
{RunID: "run1", PromptID: "p1", ConfigName: "baseline"},
{RunID: "run1", PromptID: "p2", ConfigName: "baseline"},
}

slice := SliceTrends(entries, map[string]string{})

if slice.TotalRuns != 2 {
t.Fatalf("empty filter should return all entries, got %d", slice.TotalRuns)
}
}

func TestSliceTrendsNilFilter(t *testing.T) {
entries := []TrendEntry{
{RunID: "run1", PromptID: "p1", ConfigName: "baseline"},
}

slice := SliceTrends(entries, nil)

if slice.TotalRuns != 1 {
t.Fatalf("nil filter should return all entries, got %d", slice.TotalRuns)
}
}

func TestSliceTrendsNoMatch(t *testing.T) {
entries := []TrendEntry{
{RunID: "run1", PromptID: "p1", ConfigName: "baseline", Properties: map[string]string{"language": "python"}},
}

slice := SliceTrends(entries, map[string]string{"language": "rust"})

if slice.TotalRuns != 0 {
t.Fatalf("expected 0 entries for non-matching filter, got %d", slice.TotalRuns)
}
if len(slice.PromptTrends) != 0 {
t.Fatalf("expected 0 prompt trends, got %d", len(slice.PromptTrends))
}
}

func TestSliceTrendsPreservesRunOrder(t *testing.T) {
entries := []TrendEntry{
{RunID: "run2", PromptID: "p1", ConfigName: "baseline", Timestamp: "2025-01-02T12:00:00Z", Properties: map[string]string{"language": "python"}},
{RunID: "run1", PromptID: "p1", ConfigName: "baseline", Timestamp: "2025-01-01T12:00:00Z", Properties: map[string]string{"language": "python"}},
}

slice := SliceTrends(entries, map[string]string{"language": "python"})

if len(slice.RunIDs) != 2 {
t.Fatalf("expected 2 run IDs, got %d", len(slice.RunIDs))
}
// Entries should be sorted by timestamp after slicing.
if slice.Entries[0].RunID != "run1" {
t.Errorf("expected run1 first (earlier timestamp), got %s", slice.Entries[0].RunID)
}
}

func TestSliceTrendsNilProperties(t *testing.T) {
// Entry with nil Properties should not match any property filter.
entries := []TrendEntry{
{RunID: "run1", PromptID: "p1", ConfigName: "baseline", Properties: nil},
{RunID: "run1", PromptID: "p2", ConfigName: "baseline", Properties: map[string]string{"language": "python"}},
}

slice := SliceTrends(entries, map[string]string{"language": "python"})

if slice.TotalRuns != 1 {
t.Fatalf("expected 1 entry (nil props shouldn't match), got %d", slice.TotalRuns)
}
}

func TestSliceTrendsBuildsTrends(t *testing.T) {
entries := []TrendEntry{
{RunID: "run1", PromptID: "p1", ConfigName: "baseline", Success: true, Timestamp: "2025-01-01T12:00:00Z", Properties: map[string]string{"language": "python"}},
{RunID: "run2", PromptID: "p1", ConfigName: "baseline", Success: true, Timestamp: "2025-01-02T12:00:00Z", Properties: map[string]string{"language": "python"}},
{RunID: "run1", PromptID: "p1", ConfigName: "azure-mcp", Success: false, Timestamp: "2025-01-01T12:00:00Z", Properties: map[string]string{"language": "python"}},
}

slice := SliceTrends(entries, map[string]string{"language": "python"})

if len(slice.PromptTrends) != 1 {
t.Fatalf("expected 1 prompt trend, got %d", len(slice.PromptTrends))
}
pt := slice.PromptTrends[0]
if len(pt.Configs) != 2 {
t.Errorf("expected 2 configs, got %d", len(pt.Configs))
}
if pt.Trend["baseline"] != TrendStable {
t.Errorf("expected baseline trend stable, got %s", pt.Trend["baseline"])
}
}

// --- DetectRegressions tests ---

func TestDetectRegressionsOverallScore(t *testing.T) {
previous := []TrendEntry{
{PromptID: "p1", ConfigName: "baseline", Score: 4, MaxScore: 5}, // 0.8
}
current := []TrendEntry{
{PromptID: "p1", ConfigName: "baseline", Score: 3, MaxScore: 5}, // 0.6
}

regs := DetectRegressions(current, previous, 0.1)

if len(regs) != 1 {
t.Fatalf("expected 1 regression, got %d", len(regs))
}
r := regs[0]
if r.PromptID != "p1" {
t.Errorf("expected p1, got %s", r.PromptID)
}
if r.Config != "baseline" {
t.Errorf("expected baseline, got %s", r.Config)
}
if r.GraderName != "" {
t.Errorf("expected empty grader name for overall score, got %s", r.GraderName)
}
if r.Delta < 0.19 || r.Delta > 0.21 {
t.Errorf("expected delta ~0.2, got %f", r.Delta)
}
}

func TestDetectRegressionsPerGrader(t *testing.T) {
previous := []TrendEntry{
{
PromptID: "p1", ConfigName: "baseline", Score: 4, MaxScore: 5,
GraderScores: []GraderScore{
{GraderName: "correctness", Score: 0.9, MaxScore: 5},
{GraderName: "style", Score: 0.8, MaxScore: 5},
},
},
}
current := []TrendEntry{
{
PromptID: "p1", ConfigName: "baseline", Score: 4, MaxScore: 5,
GraderScores: []GraderScore{
{GraderName: "correctness", Score: 0.5, MaxScore: 5}, // dropped 0.4
{GraderName: "style", Score: 0.75, MaxScore: 5},       // dropped only 0.05
},
},
}

regs := DetectRegressions(current, previous, 0.1)

// Should find correctness regression but not style (delta too small).
correctnessFound := false
for _, r := range regs {
if r.GraderName == "correctness" {
correctnessFound = true
if r.Delta < 0.39 || r.Delta > 0.41 {
t.Errorf("expected correctness delta ~0.4, got %f", r.Delta)
}
}
if r.GraderName == "style" {
t.Error("style should not be flagged as regression (delta < threshold)")
}
}
if !correctnessFound {
t.Error("expected correctness grader regression to be detected")
}
}

func TestDetectRegressionsNoRegression(t *testing.T) {
previous := []TrendEntry{
{PromptID: "p1", ConfigName: "baseline", Score: 3, MaxScore: 5},
}
current := []TrendEntry{
{PromptID: "p1", ConfigName: "baseline", Score: 4, MaxScore: 5}, // improved
}

regs := DetectRegressions(current, previous, 0.1)

if len(regs) != 0 {
t.Fatalf("expected 0 regressions for improved score, got %d", len(regs))
}
}

func TestDetectRegressionsDefaultThreshold(t *testing.T) {
previous := []TrendEntry{
{PromptID: "p1", ConfigName: "baseline", Score: 5, MaxScore: 5},
}
current := []TrendEntry{
{PromptID: "p1", ConfigName: "baseline", Score: 4, MaxScore: 5}, // 0.2 drop
}

// threshold <= 0 should use default (0.1).
regs := DetectRegressions(current, previous, 0)

if len(regs) != 1 {
t.Fatalf("expected 1 regression with default threshold, got %d", len(regs))
}
}

func TestDetectRegressionsNoPreviousData(t *testing.T) {
current := []TrendEntry{
{PromptID: "p1", ConfigName: "baseline", Score: 3, MaxScore: 5},
}

regs := DetectRegressions(current, nil, 0.1)

if len(regs) != 0 {
t.Fatalf("expected 0 regressions without previous data, got %d", len(regs))
}
}

func TestDetectRegressionsMultipleConfigs(t *testing.T) {
previous := []TrendEntry{
{PromptID: "p1", ConfigName: "baseline", Score: 5, MaxScore: 5},
{PromptID: "p1", ConfigName: "azure-mcp", Score: 5, MaxScore: 5},
}
current := []TrendEntry{
{PromptID: "p1", ConfigName: "baseline", Score: 2, MaxScore: 5},   // regressed
{PromptID: "p1", ConfigName: "azure-mcp", Score: 5, MaxScore: 5}, // stable
}

regs := DetectRegressions(current, previous, 0.1)

if len(regs) != 1 {
t.Fatalf("expected 1 regression (baseline only), got %d", len(regs))
}
if regs[0].Config != "baseline" {
t.Errorf("expected baseline regression, got %s", regs[0].Config)
}
}

func TestDetectRegressionsAveragesMultipleEntries(t *testing.T) {
previous := []TrendEntry{
{PromptID: "p1", ConfigName: "baseline", Score: 5, MaxScore: 5},
{PromptID: "p1", ConfigName: "baseline", Score: 5, MaxScore: 5},
}
current := []TrendEntry{
{PromptID: "p1", ConfigName: "baseline", Score: 3, MaxScore: 5},
{PromptID: "p1", ConfigName: "baseline", Score: 3, MaxScore: 5},
}

regs := DetectRegressions(current, previous, 0.1)

if len(regs) != 1 {
t.Fatalf("expected 1 regression, got %d", len(regs))
}
// Previous avg: 1.0, Current avg: 0.6, Delta: 0.4
if regs[0].Delta < 0.39 || regs[0].Delta > 0.41 {
t.Errorf("expected delta ~0.4, got %f", regs[0].Delta)
}
}

func TestDetectRegressionsZeroMaxScore(t *testing.T) {
previous := []TrendEntry{
{PromptID: "p1", ConfigName: "baseline", Score: 0, MaxScore: 0},
}
current := []TrendEntry{
{PromptID: "p1", ConfigName: "baseline", Score: 0, MaxScore: 0},
}

regs := DetectRegressions(current, previous, 0.1)

if len(regs) != 0 {
t.Fatalf("expected 0 regressions for zero max score, got %d", len(regs))
}
}

func TestDetectRegressionsExactThreshold(t *testing.T) {
previous := []TrendEntry{
{PromptID: "p1", ConfigName: "baseline", Score: 10, MaxScore: 10}, // 1.0
}
current := []TrendEntry{
{PromptID: "p1", ConfigName: "baseline", Score: 9, MaxScore: 10}, // 0.9, delta = exactly 0.1
}

// Exact threshold should NOT trigger (> threshold, not >=).
regs := DetectRegressions(current, previous, 0.1)

if len(regs) != 0 {
t.Fatalf("expected 0 regressions for delta exactly at threshold, got %d", len(regs))
}
}

func TestDetectRegressionsSorted(t *testing.T) {
previous := []TrendEntry{
{PromptID: "p2", ConfigName: "baseline", Score: 5, MaxScore: 5},
{PromptID: "p1", ConfigName: "azure-mcp", Score: 5, MaxScore: 5},
{PromptID: "p1", ConfigName: "baseline", Score: 5, MaxScore: 5},
}
current := []TrendEntry{
{PromptID: "p2", ConfigName: "baseline", Score: 1, MaxScore: 5},
{PromptID: "p1", ConfigName: "azure-mcp", Score: 1, MaxScore: 5},
{PromptID: "p1", ConfigName: "baseline", Score: 1, MaxScore: 5},
}

regs := DetectRegressions(current, previous, 0.1)

if len(regs) != 3 {
t.Fatalf("expected 3 regressions, got %d", len(regs))
}
// Should be sorted: p1/azure-mcp, p1/baseline, p2/baseline.
if regs[0].PromptID != "p1" || regs[0].Config != "azure-mcp" {
t.Errorf("first regression should be p1/azure-mcp, got %s/%s", regs[0].PromptID, regs[0].Config)
}
if regs[1].PromptID != "p1" || regs[1].Config != "baseline" {
t.Errorf("second regression should be p1/baseline, got %s/%s", regs[1].PromptID, regs[1].Config)
}
if regs[2].PromptID != "p2" || regs[2].Config != "baseline" {
t.Errorf("third regression should be p2/baseline, got %s/%s", regs[2].PromptID, regs[2].Config)
}
}

// --- matchesProperties tests ---

func TestMatchesPropertiesEmptyProps(t *testing.T) {
entry := TrendEntry{Properties: map[string]string{"language": "python"}}
if !matchesProperties(entry, map[string]string{}) {
t.Error("empty filter should match any entry")
}
}

func TestMatchesPropertiesMissingKey(t *testing.T) {
entry := TrendEntry{Properties: map[string]string{"language": "python"}}
if matchesProperties(entry, map[string]string{"service": "identity"}) {
t.Error("entry without service property should not match service filter")
}
}
