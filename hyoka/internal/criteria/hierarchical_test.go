package criteria

import (
"os"
"testing"
)

// TestHierarchicalWhenFileLevelOnly tests file-level when conditions.
func TestHierarchicalWhenFileLevelOnly(t *testing.T) {
configs := []GraderConfig{
{
When: map[string]string{"language": "python"},
Graders: []GraderEntry{
{Name: "PEP8", Weight: 1.0},
{Name: "Type Hints", Weight: 1.0},
},
},
}

// Matches file-level when
got := MatchingGraders(configs, map[string]string{"language": "python"})
if len(got) != 2 {
t.Errorf("expected 2 graders, got %d", len(got))
}

// Does not match file-level when
got = MatchingGraders(configs, map[string]string{"language": "java"})
if len(got) != 0 {
t.Errorf("expected 0 graders for non-matching language, got %d", len(got))
}
}

// TestHierarchicalWhenGraderLevelOverride tests grader-level when overriding file-level.
func TestHierarchicalWhenGraderLevelOverride(t *testing.T) {
configs := []GraderConfig{
{
When: map[string]string{"language": "python"},
Graders: []GraderEntry{
{Name: "General Check", Weight: 1.0}, // No grader-level when
{Name: "KeyVault Specific", Weight: 1.0, When: map[string]string{"service": "keyvault"}},
},
},
}

// Python + no service: only general check applies
got := MatchingGraders(configs, map[string]string{"language": "python"})
if len(got) != 1 || got[0].Name != "General Check" {
t.Errorf("expected 1 grader (General Check), got %d", len(got))
}

// Python + keyvault: both apply
got = MatchingGraders(configs, map[string]string{"language": "python", "service": "keyvault"})
if len(got) != 2 {
t.Errorf("expected 2 graders, got %d", len(got))
}

// Python + storage: only general check applies
got = MatchingGraders(configs, map[string]string{"language": "python", "service": "storage"})
if len(got) != 1 || got[0].Name != "General Check" {
t.Errorf("expected 1 grader (General Check), got %d", len(got))
}
}

// TestHierarchicalWhenGroupLevel tests group-level when conditions.
func TestHierarchicalWhenGroupLevel(t *testing.T) {
configs := []GraderConfig{
{
When: map[string]string{"language": "java"},
Graders: []GraderEntry{
{Name: "Top-Level Java Check", Weight: 1.0},
},
Groups: []GraderGroup{
{
Name: "KeyVault Checks",
When: map[string]string{"service": "keyvault"},
Graders: []GraderEntry{
{Name: "Vault URI Format", Weight: 1.0},
{Name: "Secret Client", Weight: 1.0},
},
},
},
},
}

// Java + no service: only top-level grader
got := MatchingGraders(configs, map[string]string{"language": "java"})
if len(got) != 1 || got[0].Name != "Top-Level Java Check" {
t.Errorf("expected 1 grader (Top-Level Java Check), got %d: %v", len(got), got)
}

// Java + keyvault: all 3 graders
got = MatchingGraders(configs, map[string]string{"language": "java", "service": "keyvault"})
if len(got) != 3 {
t.Errorf("expected 3 graders, got %d", len(got))
}

// Java + storage: only top-level grader
got = MatchingGraders(configs, map[string]string{"language": "java", "service": "storage"})
if len(got) != 1 {
t.Errorf("expected 1 grader, got %d", len(got))
}
}

// TestHierarchicalWhenThreeLevels tests file → group → grader resolution.
func TestHierarchicalWhenThreeLevels(t *testing.T) {
configs := []GraderConfig{
{
When: map[string]string{"language": "python"},
Groups: []GraderGroup{
{
Name: "Auth Group",
When: map[string]string{"category": "auth"},
Graders: []GraderEntry{
{Name: "General Auth", Weight: 1.0},
{Name: "KeyVault Auth", Weight: 1.0, When: map[string]string{"service": "keyvault"}},
},
},
},
},
}

// Python + auth + no service: only general auth
got := MatchingGraders(configs, map[string]string{
"language": "python",
"category": "auth",
})
if len(got) != 1 || got[0].Name != "General Auth" {
t.Errorf("expected 1 grader (General Auth), got %d", len(got))
}

// Python + auth + keyvault: both
got = MatchingGraders(configs, map[string]string{
"language": "python",
"category": "auth",
"service":  "keyvault",
})
if len(got) != 2 {
t.Errorf("expected 2 graders, got %d", len(got))
}

// Python + crud + keyvault: nothing (category mismatch at group level)
got = MatchingGraders(configs, map[string]string{
"language": "python",
"category": "crud",
"service":  "keyvault",
})
if len(got) != 0 {
t.Errorf("expected 0 graders (group when doesn't match), got %d", len(got))
}
}

// TestMergeWhen tests the mergeWhen helper function.
func TestMergeWhen(t *testing.T) {
tests := []struct {
name     string
parent   map[string]string
child    map[string]string
expected map[string]string
}{
{
name:     "both empty",
parent:   nil,
child:    nil,
expected: nil,
},
{
name:     "parent only",
parent:   map[string]string{"language": "go"},
child:    nil,
expected: map[string]string{"language": "go"},
},
{
name:     "child only",
parent:   nil,
child:    map[string]string{"service": "storage"},
expected: map[string]string{"service": "storage"},
},
{
name:     "no overlap",
parent:   map[string]string{"language": "go"},
child:    map[string]string{"service": "storage"},
expected: map[string]string{"language": "go", "service": "storage"},
},
{
name:     "child overrides parent",
parent:   map[string]string{"language": "go", "service": "keyvault"},
child:    map[string]string{"service": "storage"},
expected: map[string]string{"language": "go", "service": "storage"},
},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
got := mergeWhen(tt.parent, tt.child)
if len(got) != len(tt.expected) {
t.Fatalf("expected %d keys, got %d", len(tt.expected), len(got))
}
for k, v := range tt.expected {
if got[k] != v {
t.Errorf("expected %s=%s, got %s=%s", k, v, k, got[k])
}
}
})
}
}

// TestLoadFileWithGroups tests loading a criteria file with groups.
func TestLoadFileWithGroups(t *testing.T) {
dir := t.TempDir()
content := `
when:
  language: python
graders:
  - name: Top-Level Check
    weight: 1.0
groups:
  - name: Auth Checks
    when:
      category: auth
    graders:
      - name: DefaultAzureCredential
        weight: 1.0
      - name: No Hardcoded Keys
        weight: 1.0
`
path := dir + "/test.yaml"
if err := os.WriteFile(path, []byte(content), 0644); err != nil {
t.Fatal(err)
}

gc, err := loadFile(path)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if len(gc.Graders) != 1 {
t.Errorf("expected 1 top-level grader, got %d", len(gc.Graders))
}
if len(gc.Groups) != 1 {
t.Errorf("expected 1 group, got %d", len(gc.Groups))
}
if gc.Groups[0].Name != "Auth Checks" {
t.Errorf("expected group name 'Auth Checks', got %q", gc.Groups[0].Name)
}
if len(gc.Groups[0].Graders) != 2 {
t.Errorf("expected 2 graders in group, got %d", len(gc.Groups[0].Graders))
}
}

// TestLoadFileEmptyRejection tests that files with no graders OR groups are rejected.
func TestLoadFileEmptyRejection(t *testing.T) {
dir := t.TempDir()

// File with when but no graders or groups
content := `when:
  language: python
`
path := dir + "/empty.yaml"
if err := os.WriteFile(path, []byte(content), 0644); err != nil {
t.Fatal(err)
}

_, err := loadFile(path)
if err == nil {
t.Error("expected error for file with no graders or groups")
}
}

// TestIsolatePropertyPreserved tests that isolate: true is parsed and preserved.
func TestIsolatePropertyPreserved(t *testing.T) {
dir := t.TempDir()
content := `
when:
  language: go
graders:
  - name: Regular Check
    weight: 1.0
  - name: Isolated Check
    weight: 1.0
    isolate: true
`
path := dir + "/test.yaml"
if err := os.WriteFile(path, []byte(content), 0644); err != nil {
t.Fatal(err)
}

gc, err := loadFile(path)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if len(gc.Graders) != 2 {
t.Fatalf("expected 2 graders, got %d", len(gc.Graders))
}
if gc.Graders[0].Isolate {
t.Error("expected first grader to NOT be isolated")
}
if !gc.Graders[1].Isolate {
t.Error("expected second grader to be isolated")
}
}
