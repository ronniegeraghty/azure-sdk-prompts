package criteria

import (
"path/filepath"
"runtime"
"testing"
)

// TestLoadUnifiedDir_RealCriteriaFixtures exercises the back-compat loader
// against the repo's on-disk criteria/ directory. Every file there uses the
// legacy shape (no `type:` on entries); all must translate cleanly.
func TestLoadUnifiedDir_RealCriteriaFixtures(t *testing.T) {
_, here, _, _ := runtime.Caller(0)
// here = .../hyoka/internal/graders/unified_realfixtures_test.go
repoRoot := filepath.Join(filepath.Dir(here), "..", "..", "..")
dir := filepath.Join(repoRoot, "criteria")

bundle, err := LoadUnifiedDir(dir)
if err != nil {
t.Fatalf("walk: %v", err)
}
if len(bundle.FileErrors) != 0 {
for _, fe := range bundle.FileErrors {
t.Errorf("unexpected deferred error on real fixture: %v", fe)
}
}
if len(bundle.Configs) == 0 {
t.Fatal("expected at least one real criteria file to load")
}
for _, cfg := range bundle.Configs {
for _, g := range cfg.Graders {
if g.Type == "" {
t.Errorf("%s: entry %q has empty type after translation", cfg.Source, g.Name)
}
}
for _, grp := range cfg.Groups {
for _, g := range grp.Graders {
if g.Type == "" {
t.Errorf("%s group %q: entry %q has empty type after translation",
cfg.Source, grp.Name, g.Name)
}
}
}
}
}
