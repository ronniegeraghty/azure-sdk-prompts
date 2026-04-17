package serve

import (
"encoding/json"
"net/http"
"net/http/httptest"
"os"
"path/filepath"
"testing"

"github.com/ronniegeraghty/hyoka/hyoka/internal/report"
)

// setupTestReportsWithEvals creates a reports directory with eval report.json files.
func setupTestReportsWithEvals(t *testing.T) string {
t.Helper()
dir := t.TempDir()

runID := "20260327-113302"
runDir := filepath.Join(dir, runID)
os.MkdirAll(runDir, 0755)

summary := map[string]any{
"run_id":            runID,
"timestamp":         "2026-03-27T18:33:02Z",
"total_evaluations": 1,
}
data, _ := json.Marshal(summary)
os.WriteFile(filepath.Join(runDir, "summary.json"), data, 0644)

evalDir := filepath.Join(runDir, "results", "identity", "data-plane", "python", "auth", "test-prompt-one", "baseline/claude-opus-4.6")
os.MkdirAll(evalDir, 0755)

pass := true
evalReport := report.EvalReport{
SchemaVersion: 2,
PromptID:      "test-prompt-one",
ConfigName:    "baseline/claude-opus-4.6",
Timestamp:     "2026-03-27T18:33:02Z",
Duration:      42.5,
Success:       true,
GraderResults: []report.GraderResult{
{
GraderName:   "llm-review",
GraderType:   "review",
Model:        "claude-opus-4.6",
OverallScore: 85,
MaxScore:     100,
Summary:      "Good implementation",
Score:        0.85,
Weight:       2.0,
Pass:         &pass,
},
{
GraderName: "file-check",
GraderType: "file",
Score:      1.0,
Weight:     1.0,
Pass:       &pass,
Gate:       true,
},
},
ActionTimeline: &report.ActionTimelineReport{
Events: []report.ActionEventReport{
{Sequence: 1, Type: "tool_call", Tool: "read_file", TurnNumber: 1},
{Sequence: 2, Type: "tool_call", Tool: "edit_file", TurnNumber: 2},
},
Summary: report.ActionSummaryReport{
TotalEvents:    10,
TotalTurns:     3,
TotalActions:   8,
TotalToolCalls: 5,
},
},
ScoreBreakdown: &report.ScoreBreakdown{
Formula:       "Final Score = sum(grader_score * weight) / sum(weights)",
FinalScore:    0.9,
FinalScorePct: 90.0,
TotalWeight:   3.0,
WeightedSum:   2.7,
Contributions: []report.ScoreContribution{
{Name: "llm-review", Kind: "review", Score: 0.85, Weight: 2.0, WeightedScore: 1.7, Pass: true},
{Name: "file-check", Kind: "file", Score: 1.0, Weight: 1.0, WeightedScore: 1.0, Pass: true, Gate: true},
},
},
}
evalData, _ := json.Marshal(evalReport)
os.WriteFile(filepath.Join(evalDir, "report.json"), evalData, 0644)

return dir
}

// setupComparisonReports creates reports for two configs with shared prompts.
func setupComparisonReports(t *testing.T) string {
t.Helper()
dir := t.TempDir()

pass := true
writeReport := func(runID, promptID, config, timestamp string, score float64) {
configDir := filepath.Join(dir, runID, "results", promptID, filepath.FromSlash(config))
os.MkdirAll(configDir, 0755)
r := report.EvalReport{
SchemaVersion: 2,
PromptID:      promptID,
ConfigName:    config,
Timestamp:     timestamp,
Success:       true,
GraderResults: []report.GraderResult{
{GraderName: "correctness", GraderType: "prompt", Score: score, Weight: 1.0, Pass: &pass},
},
}
r.ScoreBreakdown = report.BuildScoreBreakdown(r.GraderResults)
data, _ := json.Marshal(r)
os.WriteFile(filepath.Join(configDir, "report.json"), data, 0644)
}

writeReport("run-001", "prompt-alpha", "baseline/opus", "2025-01-01T10:00:00Z", 0.6)
writeReport("run-001", "prompt-beta", "baseline/opus", "2025-01-01T10:00:00Z", 0.8)
writeReport("run-001", "prompt-alpha", "azure-mcp/opus", "2025-01-01T10:00:00Z", 0.9)
writeReport("run-001", "prompt-beta", "azure-mcp/opus", "2025-01-01T10:00:00Z", 0.7)
writeReport("run-000", "prompt-alpha", "baseline/opus", "2024-06-01T10:00:00Z", 0.4)

return dir
}

func TestAPIGraders(t *testing.T) {
dir := setupTestReportsWithEvals(t)
mux := buildMux(Options{ReportsDir: dir})

req := httptest.NewRequest("GET", "/api/runs/20260327-113302/graders", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)

if rec.Code != http.StatusOK {
t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
}

var result []json.RawMessage
if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
t.Fatalf("failed to decode: %v", err)
}
if len(result) != 1 {
t.Fatalf("expected 1 eval, got %d", len(result))
}
}

func TestAPIGradersNotFound(t *testing.T) {
dir := setupTestReportsWithEvals(t)
mux := buildMux(Options{ReportsDir: dir})
req := httptest.NewRequest("GET", "/api/runs/nonexistent/graders", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)
if rec.Code != http.StatusNotFound {
t.Errorf("expected 404, got %d", rec.Code)
}
}

func TestAPITimeline(t *testing.T) {
dir := setupTestReportsWithEvals(t)
mux := buildMux(Options{ReportsDir: dir})
req := httptest.NewRequest("GET", "/api/runs/20260327-113302/timeline", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)
if rec.Code != http.StatusOK {
t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
}
var result []json.RawMessage
if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
t.Fatalf("failed to decode: %v", err)
}
if len(result) != 1 {
t.Fatalf("expected 1 eval, got %d", len(result))
}
}

func TestAPITimelineNotFound(t *testing.T) {
dir := setupTestReportsWithEvals(t)
mux := buildMux(Options{ReportsDir: dir})
req := httptest.NewRequest("GET", "/api/runs/nonexistent/timeline", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)
if rec.Code != http.StatusNotFound {
t.Errorf("expected 404, got %d", rec.Code)
}
}

func TestAPIScoreBreakdown(t *testing.T) {
dir := setupTestReportsWithEvals(t)
mux := buildMux(Options{ReportsDir: dir})
req := httptest.NewRequest("GET", "/api/runs/20260327-113302/score-breakdown", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)
if rec.Code != http.StatusOK {
t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
}
var result []json.RawMessage
if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
t.Fatalf("failed to decode: %v", err)
}
if len(result) != 1 {
t.Fatalf("expected 1 eval, got %d", len(result))
}
}

func TestAPIScoreBreakdownNotFound(t *testing.T) {
dir := setupTestReportsWithEvals(t)
mux := buildMux(Options{ReportsDir: dir})
req := httptest.NewRequest("GET", "/api/runs/nonexistent/score-breakdown", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)
if rec.Code != http.StatusNotFound {
t.Errorf("expected 404, got %d", rec.Code)
}
}

func TestAPICompareConfigs(t *testing.T) {
dir := setupComparisonReports(t)
mux := buildMux(Options{ReportsDir: dir})
req := httptest.NewRequest("GET", "/api/compare/configs?a=baseline/opus&b=azure-mcp/opus", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)
if rec.Code != http.StatusOK {
t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
}
var result map[string]any
if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
t.Fatalf("failed to decode: %v", err)
}
if result["kind"] != "configs" {
t.Errorf("expected kind configs, got %v", result["kind"])
}
if result["label_a"] != "baseline/opus" {
t.Errorf("expected label_a baseline/opus, got %v", result["label_a"])
}
if result["label_b"] != "azure-mcp/opus" {
t.Errorf("expected label_b azure-mcp/opus, got %v", result["label_b"])
}
perPrompt, ok := result["per_prompt"].([]any)
if !ok {
t.Fatal("expected per_prompt array")
}
if len(perPrompt) != 2 {
t.Errorf("expected 2 prompt diffs, got %d", len(perPrompt))
}
summary, ok := result["summary"].(map[string]any)
if !ok {
t.Fatal("expected summary object")
}
if summary["improved"].(float64)+summary["regressed"].(float64) == 0 {
t.Error("expected some improved or regressed prompts")
}
}

func TestAPICompareConfigsMissingParams(t *testing.T) {
dir := t.TempDir()
mux := buildMux(Options{ReportsDir: dir})
tests := []struct {
name string
url  string
}{
{"missing both", "/api/compare/configs"},
{"missing b", "/api/compare/configs?a=x"},
{"missing a", "/api/compare/configs?b=y"},
}
for _, tc := range tests {
t.Run(tc.name, func(t *testing.T) {
req := httptest.NewRequest("GET", tc.url, nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)
if rec.Code != http.StatusBadRequest {
t.Errorf("expected 400, got %d", rec.Code)
}
})
}
}

func TestAPICompareConfigsNoData(t *testing.T) {
dir := t.TempDir()
mux := buildMux(Options{ReportsDir: dir})
req := httptest.NewRequest("GET", "/api/compare/configs?a=nonexistent&b=also-missing", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)
if rec.Code != http.StatusInternalServerError {
t.Errorf("expected 500 for missing data, got %d", rec.Code)
}
}

func TestAPICompareRuns(t *testing.T) {
dir := setupComparisonReports(t)
mux := buildMux(Options{ReportsDir: dir})
req := httptest.NewRequest("GET", "/api/compare/runs?a=run-000&b=run-001", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)
if rec.Code != http.StatusOK {
t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
}
var result map[string]any
if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
t.Fatalf("failed to decode: %v", err)
}
if result["kind"] != "runs" {
t.Errorf("expected kind runs, got %v", result["kind"])
}
if result["label_a"] != "run-000" {
t.Errorf("expected label_a run-000, got %v", result["label_a"])
}
if result["label_b"] != "run-001" {
t.Errorf("expected label_b run-001, got %v", result["label_b"])
}
}

func TestAPICompareRunsMissingParams(t *testing.T) {
dir := t.TempDir()
mux := buildMux(Options{ReportsDir: dir})
req := httptest.NewRequest("GET", "/api/compare/runs?a=x", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)
if rec.Code != http.StatusBadRequest {
t.Errorf("expected 400, got %d", rec.Code)
}
}

func TestAPICompareRunsTraversalBlocked(t *testing.T) {
dir := t.TempDir()
mux := buildMux(Options{ReportsDir: dir})
req := httptest.NewRequest("GET", "/api/compare/runs?a=../etc&b=run-001", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)
if rec.Code != http.StatusBadRequest {
t.Errorf("expected 400 for traversal, got %d", rec.Code)
}
}

func TestAPICompareTemporal(t *testing.T) {
dir := setupComparisonReports(t)
mux := buildMux(Options{ReportsDir: dir})
req := httptest.NewRequest("GET", "/api/compare/temporal?config=baseline/opus&since=2024-12-01T00:00:00Z", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)
if rec.Code != http.StatusOK {
t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
}
var result map[string]any
if err := json.NewDecoder(rec.Body).Decode(&result); err != nil {
t.Fatalf("failed to decode: %v", err)
}
if result["config"] != "baseline/opus" {
t.Errorf("expected config baseline/opus, got %v", result["config"])
}
}

func TestAPICompareTemporalDateOnly(t *testing.T) {
dir := setupComparisonReports(t)
mux := buildMux(Options{ReportsDir: dir})
req := httptest.NewRequest("GET", "/api/compare/temporal?config=baseline/opus&since=2024-12-01", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)
if rec.Code != http.StatusOK {
t.Fatalf("expected 200 with date-only, got %d: %s", rec.Code, rec.Body.String())
}
}

func TestAPICompareTemporalMissingParams(t *testing.T) {
dir := t.TempDir()
mux := buildMux(Options{ReportsDir: dir})
tests := []struct {
name string
url  string
}{
{"missing both", "/api/compare/temporal"},
{"missing since", "/api/compare/temporal?config=x"},
{"missing config", "/api/compare/temporal?since=2025-01-01"},
}
for _, tc := range tests {
t.Run(tc.name, func(t *testing.T) {
req := httptest.NewRequest("GET", tc.url, nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)
if rec.Code != http.StatusBadRequest {
t.Errorf("expected 400, got %d", rec.Code)
}
})
}
}

func TestAPICompareTemporalBadDate(t *testing.T) {
dir := t.TempDir()
mux := buildMux(Options{ReportsDir: dir})
req := httptest.NewRequest("GET", "/api/compare/temporal?config=x&since=not-a-date", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)
if rec.Code != http.StatusBadRequest {
t.Errorf("expected 400 for bad date, got %d", rec.Code)
}
}

func TestAPITrends(t *testing.T) {
dir := setupTestReportsWithEvals(t)
mux := buildMux(Options{ReportsDir: dir})
req := httptest.NewRequest("GET", "/api/trends?promptId=test-prompt-one", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)
if rec.Code != http.StatusOK {
t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
}
}

func TestAPITrendsMissingFilter(t *testing.T) {
dir := t.TempDir()
mux := buildMux(Options{ReportsDir: dir})
req := httptest.NewRequest("GET", "/api/trends", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)
if rec.Code != http.StatusBadRequest {
t.Errorf("expected 400, got %d", rec.Code)
}
}

func TestAPIPromptHistory(t *testing.T) {
dir := setupTestReportsWithEvals(t)
mux := buildMux(Options{ReportsDir: dir})
req := httptest.NewRequest("GET", "/api/prompts/test-prompt-one/history", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)
if rec.Code != http.StatusOK {
t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
}
}

func TestAPIPromptHistoryEmpty(t *testing.T) {
dir := t.TempDir()
mux := buildMux(Options{ReportsDir: dir})
req := httptest.NewRequest("GET", "/api/prompts/nonexistent-prompt/history", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)
if rec.Code != http.StatusOK {
t.Fatalf("expected 200 (empty array), got %d: %s", rec.Code, rec.Body.String())
}
}

func TestAPIRunDetailUnknownSubRoute(t *testing.T) {
dir := setupTestReportsWithEvals(t)
mux := buildMux(Options{ReportsDir: dir})
req := httptest.NewRequest("GET", "/api/runs/20260327-113302/nonexistent", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)
if rec.Code != http.StatusNotFound {
t.Errorf("expected 404, got %d", rec.Code)
}
}

func TestAPIPairwise(t *testing.T) {
dir := setupTestReportsWithEvals(t)
runID := "20260327-113302"

// Seed pairwise.json so the endpoint has something to serve.
pairwisePayload := map[string]any{
"run_id":    runID,
"timestamp": "2026-03-27T18:33:02Z",
"reports": []map[string]any{
{
"prompt_id": "test-prompt-one",
"baseline": map[string]any{
"config_name": "baseline/claude-opus-4.6",
"score":       0.85,
"max_score":   1.0,
"success":     true,
},
"variants": []any{},
"impacts":  []any{},
},
},
"aggregate_impacts": []any{},
}
data, _ := json.Marshal(pairwisePayload)
if err := os.WriteFile(filepath.Join(dir, runID, "pairwise.json"), data, 0644); err != nil {
t.Fatalf("seed pairwise.json: %v", err)
}

mux := buildMux(Options{ReportsDir: dir})
req := httptest.NewRequest("GET", "/api/runs/"+runID+"/pairwise", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)

if rec.Code != http.StatusOK {
t.Fatalf("expected 200, got %d: %s", rec.Code, rec.Body.String())
}
if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
t.Errorf("expected JSON content type, got %q", ct)
}
var got map[string]any
if err := json.Unmarshal(rec.Body.Bytes(), &got); err != nil {
t.Fatalf("decode: %v", err)
}
if got["run_id"] != runID {
t.Errorf("expected run_id=%s, got %v", runID, got["run_id"])
}
}

func TestAPIPairwiseMissing(t *testing.T) {
dir := setupTestReportsWithEvals(t)
mux := buildMux(Options{ReportsDir: dir})

// No pairwise.json in this run → 404.
req := httptest.NewRequest("GET", "/api/runs/20260327-113302/pairwise", nil)
rec := httptest.NewRecorder()
mux.ServeHTTP(rec, req)
if rec.Code != http.StatusNotFound {
t.Errorf("expected 404 for run without pairwise.json, got %d", rec.Code)
}
}
