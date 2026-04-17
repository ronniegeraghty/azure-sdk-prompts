package cmd

import (
"bytes"
"encoding/json"
"io"
"os"
"path/filepath"
"strings"
"testing"

"github.com/ronniegeraghty/hyoka/hyoka/internal/report"
)

func TestCompareCmdHelp(t *testing.T) {
cmd := compareCmd()
cmd.SetOut(io.Discard)
cmd.SetErr(io.Discard)
cmd.SetArgs([]string{"--help"})
if err := cmd.Execute(); err != nil {
t.Fatalf("compare --help failed: %v", err)
}
}

func TestCompareCmdFlagDefaults(t *testing.T) {
cmd := compareCmd()
cmd.SetOut(io.Discard)
cmd.SetErr(io.Discard)
cmd.SetArgs([]string{"--help"})
_ = cmd.Execute()

tests := []struct {
flag     string
expected string
}{
{"config-a", ""},
{"config-b", ""},
{"run-a", ""},
{"run-b", ""},
{"config", ""},
{"since", ""},
{"reports-dir", "./reports"},
{"format", "table"},
{"top", "0"},
}

for _, tt := range tests {
f := cmd.Flags().Lookup(tt.flag)
if f == nil {
t.Errorf("expected flag %q to be registered", tt.flag)
continue
}
if f.DefValue != tt.expected {
t.Errorf("flag %q: expected default %q, got %q", tt.flag, tt.expected, f.DefValue)
}
}
}

func TestDetectMode_ConfigPair(t *testing.T) {
mode, err := detectMode("a", "b", "", "", "", "")
if err != nil {
t.Fatal(err)
}
if mode != "configs" {
t.Errorf("expected configs, got %s", mode)
}
}

func TestDetectMode_RunPair(t *testing.T) {
mode, err := detectMode("", "", "r1", "r2", "", "")
if err != nil {
t.Fatal(err)
}
if mode != "runs" {
t.Errorf("expected runs, got %s", mode)
}
}

func TestDetectMode_Temporal(t *testing.T) {
mode, err := detectMode("", "", "", "", "cfg", "2025-01-01")
if err != nil {
t.Fatal(err)
}
if mode != "temporal" {
t.Errorf("expected temporal, got %s", mode)
}
}

func TestDetectMode_NoFlags(t *testing.T) {
_, err := detectMode("", "", "", "", "", "")
if err == nil {
t.Error("expected error when no flags set")
}
}

func TestDetectMode_ConflictingModes(t *testing.T) {
_, err := detectMode("a", "b", "r1", "r2", "", "")
if err == nil {
t.Error("expected error for conflicting modes")
}
}

func TestDetectMode_IncompleteConfigPair(t *testing.T) {
_, err := detectMode("a", "", "", "", "", "")
if err == nil {
t.Error("expected error for incomplete config pair")
}
}

func TestDetectMode_IncompleteRunPair(t *testing.T) {
_, err := detectMode("", "", "r1", "", "", "")
if err == nil {
t.Error("expected error for incomplete run pair")
}
}

func TestDetectMode_IncompleteTemporal(t *testing.T) {
_, err := detectMode("", "", "", "", "cfg", "")
if err == nil {
t.Error("expected error for incomplete temporal")
}
}

// writeTestReport creates a minimal report.json for compare command testing.
func writeTestReport(t *testing.T, dir, runID, promptID, config, timestamp string, score float64) {
t.Helper()
pass := score >= 0.5
graders := []report.GraderResult{
{GraderName: "correctness", GraderType: "prompt", Score: score, Weight: 1.0, Pass: &pass},
}
r := report.EvalReport{
SchemaVersion:  2,
PromptID:       promptID,
ConfigName:     config,
Timestamp:      timestamp,
Success:        true,
GraderResults:  graders,
ScoreBreakdown: report.BuildScoreBreakdown(graders),
}

configDir := filepath.Join(dir, runID, "results", promptID, filepath.FromSlash(config))
if err := os.MkdirAll(configDir, 0755); err != nil {
t.Fatal(err)
}
data, err := json.Marshal(r)
if err != nil {
t.Fatal(err)
}
if err := os.WriteFile(filepath.Join(configDir, "report.json"), data, 0644); err != nil {
t.Fatal(err)
}
}

func TestCompareCmd_ConfigComparison_Table(t *testing.T) {
dir := t.TempDir()
writeTestReport(t, dir, "run-001", "prompt-alpha", "config-a", "2025-01-01T10:00:00Z", 0.6)
writeTestReport(t, dir, "run-001", "prompt-alpha", "config-b", "2025-01-01T10:00:00Z", 0.9)

var buf bytes.Buffer
cmd := compareCmd()
cmd.SetOut(&buf)
cmd.SetErr(io.Discard)
cmd.SetArgs([]string{"--config-a", "config-a", "--config-b", "config-b", "--reports-dir", dir})

if err := cmd.Execute(); err != nil {
t.Fatalf("compare failed: %v", err)
}

out := buf.String()
if !strings.Contains(out, "Config comparison") {
t.Error("expected 'Config comparison' header in output")
}
if !strings.Contains(out, "prompt-alpha") {
t.Error("expected prompt-alpha in output")
}
if !strings.Contains(out, "improved") {
t.Error("expected 'improved' in summary")
}
}

func TestCompareCmd_ConfigComparison_JSON(t *testing.T) {
dir := t.TempDir()
writeTestReport(t, dir, "run-001", "prompt-alpha", "config-a", "2025-01-01T10:00:00Z", 0.6)
writeTestReport(t, dir, "run-001", "prompt-alpha", "config-b", "2025-01-01T10:00:00Z", 0.9)

var buf bytes.Buffer
cmd := compareCmd()
cmd.SetOut(&buf)
cmd.SetErr(io.Discard)
cmd.SetArgs([]string{"--config-a", "config-a", "--config-b", "config-b", "--reports-dir", dir, "--format", "json"})

if err := cmd.Execute(); err != nil {
t.Fatalf("compare --format json failed: %v", err)
}

var result map[string]any
if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
t.Fatalf("output is not valid JSON: %v\n%s", err, buf.String())
}
if result["kind"] != "configs" {
t.Errorf("expected kind = configs, got %v", result["kind"])
}
if result["label_a"] != "config-a" {
t.Errorf("expected label_a = config-a, got %v", result["label_a"])
}
if result["label_b"] != "config-b" {
t.Errorf("expected label_b = config-b, got %v", result["label_b"])
}
}

func TestCompareCmd_RunComparison(t *testing.T) {
dir := t.TempDir()
writeTestReport(t, dir, "20250101-100000", "prompt-alpha", "config-a", "2025-01-01T10:00:00Z", 0.5)
writeTestReport(t, dir, "20250102-100000", "prompt-alpha", "config-a", "2025-01-02T10:00:00Z", 0.8)

var buf bytes.Buffer
cmd := compareCmd()
cmd.SetOut(&buf)
cmd.SetErr(io.Discard)
cmd.SetArgs([]string{"--run-a", "20250101-100000", "--run-b", "20250102-100000", "--reports-dir", dir})

if err := cmd.Execute(); err != nil {
t.Fatalf("compare runs failed: %v", err)
}

out := buf.String()
if !strings.Contains(out, "Run comparison") {
t.Error("expected 'Run comparison' header")
}
}

func TestCompareCmd_TemporalComparison(t *testing.T) {
dir := t.TempDir()
writeTestReport(t, dir, "run-old", "prompt-alpha", "config-a", "2025-01-01T10:00:00Z", 0.4)
writeTestReport(t, dir, "run-new", "prompt-alpha", "config-a", "2025-02-01T10:00:00Z", 0.9)

var buf bytes.Buffer
cmd := compareCmd()
cmd.SetOut(&buf)
cmd.SetErr(io.Discard)
cmd.SetArgs([]string{"--config", "config-a", "--since", "2025-01-15", "--reports-dir", dir})

if err := cmd.Execute(); err != nil {
t.Fatalf("compare temporal failed: %v", err)
}

out := buf.String()
if !strings.Contains(out, "Temporal comparison") {
t.Error("expected 'Temporal comparison' header")
}
}

func TestCompareCmd_TopN(t *testing.T) {
dir := t.TempDir()
for i := 0; i < 5; i++ {
pid := "prompt-" + string(rune('a'+i))
writeTestReport(t, dir, "run-001", pid, "config-a", "2025-01-01T10:00:00Z", 0.3+float64(i)*0.1)
writeTestReport(t, dir, "run-001", pid, "config-b", "2025-01-01T10:00:00Z", 0.8)
}

var buf bytes.Buffer
cmd := compareCmd()
cmd.SetOut(&buf)
cmd.SetErr(io.Discard)
cmd.SetArgs([]string{"--config-a", "config-a", "--config-b", "config-b", "--reports-dir", dir, "--top", "2"})

if err := cmd.Execute(); err != nil {
t.Fatalf("compare --top 2 failed: %v", err)
}

out := buf.String()
if !strings.Contains(out, "showing top 2 of 5") {
t.Errorf("expected top N indicator in output:\n%s", out)
}
}

func TestCompareCmd_NoFlagsError(t *testing.T) {
cmd := compareCmd()
cmd.SetOut(io.Discard)
cmd.SetErr(io.Discard)
cmd.SilenceErrors = true
cmd.SilenceUsage = true
cmd.SetArgs([]string{})

err := cmd.Execute()
if err == nil {
t.Error("expected error when no comparison flags provided")
}
}
