package eval

import (
"context"
"os"
"path/filepath"
"sort"
"strings"
"testing"

"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
"github.com/ronniegeraghty/hyoka/hyoka/internal/prompt"
)

func TestSnapshotDir_CapturesFilesAndDirs(t *testing.T) {
dir := t.TempDir()
os.WriteFile(filepath.Join(dir, "hello.py"), []byte("print('hi')"), 0644)
os.Mkdir(filepath.Join(dir, "subdir"), 0755)
os.Mkdir(filepath.Join(dir, ".hidden"), 0755)
snap := snapshotDir(dir)
if snap == nil { t.Fatal("snapshotDir returned nil") }
if !snap["hello.py"] { t.Error("snapshotDir should capture files") }
if !snap["subdir"] { t.Error("snapshotDir should capture directories") }
if snap[".hidden"] { t.Error("snapshotDir should skip hidden entries") }
}

func TestRecoverMisplacedFiles_RecoversNewFiles(t *testing.T) {
home := t.TempDir(); workspace := t.TempDir()
os.WriteFile(filepath.Join(home, "existing.txt"), []byte("old"), 0644)
snap := snapshotDir(home)
os.WriteFile(filepath.Join(home, "main.py"), []byte("print('hello')"), 0644)
recovered := recoverMisplacedFiles(home, snap, workspace, "test")
if recovered != 1 { t.Fatalf("expected 1 recovered, got %d", recovered) }
if _, err := os.Stat(filepath.Join(workspace, "main.py")); err != nil { t.Error("main.py should be in workspace") }
if _, err := os.Stat(filepath.Join(home, "main.py")); err == nil { t.Error("main.py should be removed from home") }
}

func TestRecoverMisplacedFiles_DeletesJunkDirs(t *testing.T) {
home := t.TempDir(); workspace := t.TempDir()
snap := snapshotDir(home)
pycache := filepath.Join(home, "__pycache__")
os.Mkdir(pycache, 0755)
os.WriteFile(filepath.Join(pycache, "mod.cpython-311.pyc"), []byte{0}, 0644)
recovered := recoverMisplacedFiles(home, snap, workspace, "test")
if recovered != 1 { t.Fatalf("expected 1 recovered, got %d", recovered) }
if _, err := os.Stat(pycache); err == nil { t.Error("__pycache__ should be deleted") }
}

func TestRecoverMisplacedFiles_MovesNewDirToWorkspace(t *testing.T) {
home := t.TempDir(); workspace := t.TempDir()
snap := snapshotDir(home)
projDir := filepath.Join(home, "myproject")
os.Mkdir(projDir, 0755)
os.WriteFile(filepath.Join(projDir, "app.py"), []byte("app"), 0644)
recovered := recoverMisplacedFiles(home, snap, workspace, "test")
if recovered != 1 { t.Fatalf("expected 1 recovered, got %d", recovered) }
if _, err := os.Stat(filepath.Join(workspace, "myproject", "app.py")); err != nil { t.Error("myproject/app.py should be in workspace") }
}

func TestRecoverMisplacedFiles_SkipsPreExistingDirs(t *testing.T) {
home := t.TempDir(); workspace := t.TempDir()
os.Mkdir(filepath.Join(home, "Documents"), 0755)
snap := snapshotDir(home)
recovered := recoverMisplacedFiles(home, snap, workspace, "test")
if recovered != 0 { t.Fatalf("expected 0 recovered, got %d", recovered) }
}

func TestFilterExcludedDirs_Empty(t *testing.T) {
files := []string{"main.go", "lib/utils.go"}
if got := filterExcludedDirs(files, nil); len(got) != 2 { t.Errorf("expected 2, got %d", len(got)) }
if got := filterExcludedDirs(files, []string{}); len(got) != 2 { t.Errorf("expected 2, got %d", len(got)) }
}

func TestFilterExcludedDirs_ExcludesTopLevel(t *testing.T) {
files := []string{"main.go", "node_modules/express/index.js", "src/app.js", "dist/bundle.js"}
got := filterExcludedDirs(files, []string{"node_modules", "dist"})
if len(got) != 2 { t.Errorf("expected 2, got %d: %v", len(got), got) }
}

func TestFilterExcludedDirs_KeepsNonMatching(t *testing.T) {
files := []string{"src/main.go", "src/utils.go", "README.md"}
if got := filterExcludedDirs(files, []string{"node_modules"}); len(got) != 3 { t.Errorf("expected 3, got %d", len(got)) }
}

func TestFilterExcludedDirs_NestedMatch(t *testing.T) {
files := []string{"project/target/classes/App.class", "project/src/App.java", "target/output.jar"}
if got := filterExcludedDirs(files, []string{"target"}); len(got) != 1 { t.Errorf("expected 1, got %d: %v", len(got), got) }
}

func TestFilterExcludedDirs_TrailingSlash(t *testing.T) {
files := []string{"dist/app.js", "src/main.ts"}
if got := filterExcludedDirs(files, []string{"dist/"}); len(got) != 1 { t.Errorf("expected 1, got %d: %v", len(got), got) }
}

func TestDefaultIgnoreDirs_CoversAllLanguages(t *testing.T) {
for _, dir := range []string{"node_modules", "__pycache__", "venv", "target", "vendor", ".gradle", "bin", "obj", "dist", "build"} {
if !DefaultIgnoreDirs[dir] { t.Errorf("missing %q", dir) }
}
}

func TestJunkDirsIsDefaultIgnoreDirs(t *testing.T) {
for k := range DefaultIgnoreDirs { if !junkDirs[k] { t.Errorf("junkDirs missing %q", k) } }
for k := range junkDirs { if !DefaultIgnoreDirs[k] { t.Errorf("DefaultIgnoreDirs missing %q", k) } }
}

func TestCopyDir_CopiesFilesAndSubdirs(t *testing.T) {
src := t.TempDir(); dst := t.TempDir()
os.WriteFile(filepath.Join(src, "main.py"), []byte("print('hello')"), 0644)
os.MkdirAll(filepath.Join(src, "sub"), 0755)
os.WriteFile(filepath.Join(src, "sub", "util.py"), []byte("# util"), 0644)
if err := copyDir(src, dst); err != nil { t.Fatalf("copyDir failed: %v", err) }
data, _ := os.ReadFile(filepath.Join(dst, "main.py"))
if string(data) != "print('hello')" { t.Errorf("unexpected content: %q", data) }
}

func TestCopyDir_SkipsHiddenDirs(t *testing.T) {
src := t.TempDir(); dst := t.TempDir()
os.MkdirAll(filepath.Join(src, ".git"), 0755)
os.WriteFile(filepath.Join(src, ".git", "config"), []byte("x"), 0644)
os.WriteFile(filepath.Join(src, "main.py"), []byte("code"), 0644)
copyDir(src, dst)
if _, err := os.Stat(filepath.Join(dst, ".git")); err == nil { t.Error(".git should not be copied") }
if _, err := os.Stat(filepath.Join(dst, "main.py")); err != nil { t.Error("main.py should be copied") }
}

func TestCopyDir_MissingSrcReturnsError(t *testing.T) {
if err := copyDir("/nonexistent/path", t.TempDir()); err == nil { t.Fatal("expected error") }
}

func TestCopyDir_EmptyDirectory(t *testing.T) {
src := t.TempDir(); dst := t.TempDir()
copyDir(src, dst)
files, _ := listFiles(dst)
if len(files) != 0 { t.Errorf("expected 0 files, got %d", len(files)) }
}

func TestResolveStarterDir_Relative(t *testing.T) {
p := &prompt.Prompt{StarterProject: "./starter/", FilePath: "/home/user/prompts/storage/blob.prompt.md"}
if got := resolveStarterDir(p); got != "/home/user/prompts/storage/starter" { t.Errorf("got %q", got) }
}

func TestResolveStarterDir_Absolute(t *testing.T) {
p := &prompt.Prompt{StarterProject: "/absolute/starter", FilePath: "/home/user/prompts/blob.prompt.md"}
if got := resolveStarterDir(p); got != "/absolute/starter" { t.Errorf("got %q", got) }
}

func TestResolveStarterDir_NoFilePath(t *testing.T) {
p := &prompt.Prompt{StarterProject: "starter/"}
if got := resolveStarterDir(p); got != "starter/" { t.Errorf("got %q", got) }
}

// --- CopyStarterFiles tests (#127) ---

func TestCopyStarterFiles_CopiesOnlyDeclaredFiles(t *testing.T) {
starterDir := t.TempDir()
os.WriteFile(filepath.Join(starterDir, "main.py"), []byte("print('hello')"), 0644)
os.MkdirAll(filepath.Join(starterDir, "src"), 0755)
os.WriteFile(filepath.Join(starterDir, "src", "utils.py"), []byte("# utils"), 0644)
os.WriteFile(filepath.Join(starterDir, "requirements.txt"), []byte("azure-identity"), 0644)

ws, _ := NewWorkspace("test-prompt", "test-config")
defer ws.Cleanup()

files, err := ws.CopyStarterFiles(&prompt.Prompt{StarterProject: starterDir})
if err != nil { t.Fatalf("CopyStarterFiles failed: %v", err) }
sort.Strings(files)
expected := []string{"main.py", "requirements.txt", "src/utils.py"}
if len(files) != len(expected) { t.Fatalf("expected %d files, got %d: %v", len(expected), len(files), files) }
for i, f := range expected { if files[i] != f { t.Errorf("file[%d] = %q, want %q", i, files[i], f) } }
data, _ := os.ReadFile(filepath.Join(ws.Dir, "main.py"))
if string(data) != "print('hello')" { t.Error("main.py content mismatch") }
}

func TestCopyStarterFiles_NoLeakOfHiddenFiles(t *testing.T) {
starterDir := t.TempDir()
os.WriteFile(filepath.Join(starterDir, "app.py"), []byte("app"), 0644)
os.MkdirAll(filepath.Join(starterDir, ".git"), 0755)
os.WriteFile(filepath.Join(starterDir, ".git", "config"), []byte("x"), 0644)

ws, _ := NewWorkspace("test-prompt", "test-config")
defer ws.Cleanup()
files, _ := ws.CopyStarterFiles(&prompt.Prompt{StarterProject: starterDir})
if len(files) != 1 || files[0] != "app.py" { t.Errorf("expected [app.py], got %v", files) }
if _, err := os.Stat(filepath.Join(ws.Dir, ".git")); err == nil { t.Error(".git should not be copied") }
}

func TestCopyStarterFiles_ExcludesBuildArtifacts(t *testing.T) {
starterDir := t.TempDir()
os.WriteFile(filepath.Join(starterDir, "main.py"), []byte("code"), 0644)
os.MkdirAll(filepath.Join(starterDir, "node_modules", "express"), 0755)
os.WriteFile(filepath.Join(starterDir, "node_modules", "express", "index.js"), []byte("//"), 0644)

ws, _ := NewWorkspace("test-prompt", "test-config")
defer ws.Cleanup()
files, _ := ws.CopyStarterFiles(&prompt.Prompt{StarterProject: starterDir})
if len(files) != 1 || files[0] != "main.py" { t.Errorf("expected [main.py], got %v", files) }
if _, err := os.Stat(filepath.Join(ws.Dir, "node_modules")); err == nil { t.Error("node_modules should not be copied") }
}

func TestCopyStarterFiles_NoStarterProject(t *testing.T) {
ws, _ := NewWorkspace("test", "test"); defer ws.Cleanup()
files, err := ws.CopyStarterFiles(&prompt.Prompt{})
if err != nil { t.Fatalf("unexpected error: %v", err) }
if files != nil { t.Errorf("expected nil, got %v", files) }
}

func TestCopyStarterFiles_InvalidPath(t *testing.T) {
ws, _ := NewWorkspace("test", "test"); defer ws.Cleanup()
_, err := ws.CopyStarterFiles(&prompt.Prompt{StarterProject: "/nonexistent/path"})
if err == nil { t.Fatal("expected error") }
}

func TestCopyStarterFiles_RelativeToPromptFile(t *testing.T) {
promptDir := t.TempDir()
os.MkdirAll(filepath.Join(promptDir, "starter"), 0755)
os.WriteFile(filepath.Join(promptDir, "starter", "setup.py"), []byte("setup"), 0644)
ws, _ := NewWorkspace("test", "test"); defer ws.Cleanup()
files, _ := ws.CopyStarterFiles(&prompt.Prompt{StarterProject: "./starter", FilePath: filepath.Join(promptDir, "test.prompt.md")})
if len(files) != 1 || files[0] != "setup.py" { t.Errorf("expected [setup.py], got %v", files) }
}

// --- Workspace Cleanup tests (#128) ---

func TestWorkspaceCleanup_EphemeralRemovedOnSuccess(t *testing.T) {
ws, _ := NewWorkspace("test", "test"); dir := ws.Dir
os.WriteFile(filepath.Join(dir, "output.py"), []byte("result"), 0644)
ws.Cleanup()
if _, err := os.Stat(dir); !os.IsNotExist(err) { t.Error("ephemeral workspace should be removed") }
}

func TestWorkspaceCleanup_PersistentSurvives(t *testing.T) {
dir := t.TempDir(); ws, _ := NewWorkspaceAt(dir)
os.WriteFile(filepath.Join(dir, "output.py"), []byte("result"), 0644)
ws.Cleanup()
if _, err := os.Stat(filepath.Join(dir, "output.py")); err != nil { t.Error("persistent workspace should survive") }
}

func TestWorkspaceCleanup_OnError(t *testing.T) {
ws, _ := NewWorkspace("test", "test"); dir := ws.Dir
func() { defer ws.Cleanup(); os.WriteFile(filepath.Join(dir, "partial.py"), []byte("x"), 0644) }()
if _, err := os.Stat(dir); !os.IsNotExist(err) { t.Error("workspace should be cleaned up on error path") }
}

func TestWorkspaceCleanup_OnContextCancellation(t *testing.T) {
ws, _ := NewWorkspace("test", "test"); dir := ws.Dir
ctx, cancel := context.WithCancel(context.Background())
func() { defer ws.Cleanup(); cancel(); _ = ctx.Err() }()
if _, err := os.Stat(dir); !os.IsNotExist(err) { t.Error("workspace should be cleaned up after cancellation") }
}

func TestWorkspaceCleanup_DoubleCleanup(t *testing.T) {
ws, _ := NewWorkspace("test", "test")
if err := ws.Cleanup(); err != nil { t.Fatal(err) }
if err := ws.Cleanup(); err != nil { t.Fatal("second cleanup should not fail:", err) }
}

// --- EvalWorkspace tests (#126) ---

func TestNewEvalWorkspace_CreatesIsolatedDir(t *testing.T) {
	ws, err := NewEvalWorkspace()
	if err != nil {
		t.Fatalf("NewEvalWorkspace() error: %v", err)
	}
	defer ws.Cleanup()

	info, err := os.Stat(ws.Dir)
	if err != nil {
		t.Fatalf("workspace dir does not exist: %v", err)
	}
	if !info.IsDir() {
		t.Fatal("workspace path is not a directory")
	}
	if !strings.HasPrefix(filepath.Base(ws.Dir), EvalWorkspacePrefix) {
		t.Errorf("workspace dir %q does not have prefix %q", ws.Dir, EvalWorkspacePrefix)
	}
	entries, err := os.ReadDir(ws.Dir)
	if err != nil {
		t.Fatalf("reading workspace dir: %v", err)
	}
	if len(entries) != 0 {
		t.Errorf("workspace should be empty, got %d entries", len(entries))
	}
	if ws.CreatedAt.IsZero() {
		t.Error("CreatedAt should be set")
	}
}

func TestEvalWorkspace_Cleanup(t *testing.T) {
	ws, err := NewEvalWorkspace()
	if err != nil {
		t.Fatalf("NewEvalWorkspace() error: %v", err)
	}
	dir := ws.Dir
	os.WriteFile(filepath.Join(dir, "test.py"), []byte("code"), 0644)
	if err := ws.Cleanup(); err != nil {
		t.Fatalf("Cleanup() error: %v", err)
	}
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		t.Errorf("workspace dir should be removed after Cleanup")
	}
	if err := ws.Cleanup(); err != nil {
		t.Errorf("second Cleanup() should not error, got: %v", err)
	}
}

func TestEvalWorkspace_CleanupNil(t *testing.T) {
	var ws *EvalWorkspace
	if err := ws.Cleanup(); err != nil {
		t.Errorf("Cleanup on nil workspace should not error, got: %v", err)
	}
}

func TestEvalWorkspace_CopyStarterFiles(t *testing.T) {
	ws, err := NewEvalWorkspace()
	if err != nil {
		t.Fatalf("NewEvalWorkspace() error: %v", err)
	}
	defer ws.Cleanup()
	starterDir := t.TempDir()
	os.WriteFile(filepath.Join(starterDir, "main.py"), []byte("starter"), 0644)
	os.MkdirAll(filepath.Join(starterDir, "lib"), 0755)
	os.WriteFile(filepath.Join(starterDir, "lib", "utils.py"), []byte("utils"), 0644)
	if err := ws.CopyStarterFiles(starterDir); err != nil {
		t.Fatalf("CopyStarterFiles() error: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(ws.Dir, "main.py"))
	if err != nil {
		t.Fatal("main.py not copied to workspace")
	}
	if string(data) != "starter" {
		t.Errorf("unexpected content: %q", data)
	}
	data, err = os.ReadFile(filepath.Join(ws.Dir, "lib", "utils.py"))
	if err != nil {
		t.Fatal("lib/utils.py not copied to workspace")
	}
	if string(data) != "utils" {
		t.Errorf("unexpected content: %q", data)
	}
}

func TestEvalWorkspace_ListFiles(t *testing.T) {
	ws, err := NewEvalWorkspace()
	if err != nil {
		t.Fatalf("NewEvalWorkspace() error: %v", err)
	}
	defer ws.Cleanup()
	files, err := ws.ListFiles()
	if err != nil {
		t.Fatalf("ListFiles() error: %v", err)
	}
	if len(files) != 0 {
		t.Errorf("expected 0 files, got %d", len(files))
	}
	os.WriteFile(filepath.Join(ws.Dir, "app.py"), []byte("app"), 0644)
	os.MkdirAll(filepath.Join(ws.Dir, "pkg"), 0755)
	os.WriteFile(filepath.Join(ws.Dir, "pkg", "mod.py"), []byte("mod"), 0644)
	files, err = ws.ListFiles()
	if err != nil {
		t.Fatalf("ListFiles() error: %v", err)
	}
	if len(files) != 2 {
		t.Errorf("expected 2 files, got %d: %v", len(files), files)
	}
}

func TestEvalWorkspace_IsIsolated(t *testing.T) {
	ws, err := NewEvalWorkspace()
	if err != nil {
		t.Fatalf("NewEvalWorkspace() error: %v", err)
	}
	defer ws.Cleanup()
	cwdFiles, _ := os.ReadDir(".")
	wsFiles, _ := os.ReadDir(ws.Dir)
	if len(cwdFiles) == 0 {
		t.Skip("CWD has no files")
	}
	if len(wsFiles) != 0 {
		t.Errorf("workspace should be empty (isolated from CWD), got %d entries", len(wsFiles))
	}
}

// --- Workspace lifecycle tests (#261) ---

// TestWorkspaceLifecycle_FilesAvailableBeforeCleanup verifies that workspace
// files persist through the evaluation lifecycle and can be copied to the
// report directory before session cleanup runs.
func TestWorkspaceLifecycle_FilesAvailableBeforeCleanup(t *testing.T) {
	genWs, err := NewWorkspace("test-prompt", "test-config")
	if err != nil {
		t.Fatalf("NewWorkspace() error: %v", err)
	}
	defer genWs.Cleanup()

	// Simulate the Copilot agent writing files to the workspace
	os.WriteFile(filepath.Join(genWs.Dir, "main.py"), []byte("print('hello')"), 0644)
	os.MkdirAll(filepath.Join(genWs.Dir, "pkg"), 0755)
	os.WriteFile(filepath.Join(genWs.Dir, "pkg", "utils.py"), []byte("# utils"), 0644)

	// Create persistent report directory (mirrors engine.go ws setup)
	reportDir := filepath.Join(t.TempDir(), "generated-code")
	ws, err := NewWorkspaceAt(reportDir)
	if err != nil {
		t.Fatalf("NewWorkspaceAt() error: %v", err)
	}

	// Copy files BEFORE cleanup — this is the critical ordering (#261)
	if err := copyDir(genWs.Dir, ws.Dir); err != nil {
		t.Fatalf("copyDir() error: %v", err)
	}

	// Verify files are in the report directory
	files, err := ws.ListFiles()
	if err != nil {
		t.Fatalf("ListFiles() error: %v", err)
	}
	if len(files) != 2 {
		t.Errorf("expected 2 files in report dir, got %d: %v", len(files), files)
	}

	// Now simulate cleanup (as if CleanupFn ran)
	genWs.Cleanup()

	// Report dir files should survive cleanup of the temp workspace
	files, err = ws.ListFiles()
	if err != nil {
		t.Fatalf("ListFiles() after cleanup error: %v", err)
	}
	if len(files) != 2 {
		t.Errorf("report dir should retain files after cleanup, got %d", len(files))
	}

	// Verify file contents are correct
	data, err := os.ReadFile(filepath.Join(ws.Dir, "main.py"))
	if err != nil {
		t.Fatalf("reading main.py: %v", err)
	}
	if string(data) != "print('hello')" {
		t.Errorf("unexpected content: %q", data)
	}
}

// TestCleanupFn_CalledAfterFileCopy verifies the CleanupFn pattern: session
// state is only deleted after workspace files have been safely copied.
func TestCleanupFn_CalledAfterFileCopy(t *testing.T) {
	genDir := t.TempDir()
	os.WriteFile(filepath.Join(genDir, "output.py"), []byte("result"), 0644)

	var cleanupCalled bool
	result := &EvalResult{
		GeneratedFiles: []string{"output.py"},
		Success:        true,
		CleanupFn: func() {
			// Simulate DeleteSession removing workspace files
			os.RemoveAll(genDir)
			cleanupCalled = true
		},
	}

	// Simulate engine.go copying files before calling CleanupFn
	reportDir := t.TempDir()
	if err := copyDir(genDir, reportDir); err != nil {
		t.Fatalf("copyDir() error: %v", err)
	}

	// Verify files exist in report dir BEFORE cleanup
	data, err := os.ReadFile(filepath.Join(reportDir, "output.py"))
	if err != nil {
		t.Fatalf("file not copied to report dir: %v", err)
	}
	if string(data) != "result" {
		t.Errorf("unexpected content: %q", data)
	}

	// Call CleanupFn — simulates session deletion
	if result.CleanupFn != nil {
		result.CleanupFn()
	}

	if !cleanupCalled {
		t.Error("CleanupFn should have been called")
	}

	// genDir is now removed (simulating DeleteSession behavior)
	if _, err := os.Stat(genDir); !os.IsNotExist(err) {
		t.Error("genDir should be removed by CleanupFn")
	}

	// But report dir files should survive
	data, err = os.ReadFile(filepath.Join(reportDir, "output.py"))
	if err != nil {
		t.Fatalf("report dir file should survive CleanupFn: %v", err)
	}
	if string(data) != "result" {
		t.Errorf("unexpected content after cleanup: %q", data)
	}
}

// TestCleanupFn_NilSafe verifies that nil CleanupFn doesn't panic.
func TestCleanupFn_NilSafe(t *testing.T) {
	result := &EvalResult{Success: true}
	// Should not panic
	if result.CleanupFn != nil {
		result.CleanupFn()
	}
}

// TestCleanupFn_StubEvaluatorHasNoCleanup verifies StubEvaluator returns nil CleanupFn.
func TestCleanupFn_StubEvaluatorHasNoCleanup(t *testing.T) {
	stub := &StubEvaluator{}
	result, err := stub.Evaluate(context.Background(),
		&prompt.Prompt{ID: "test"},
		&config.ToolConfig{Name: "test", Generator: &config.GeneratorConfig{Model: "test"}},
		t.TempDir())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.CleanupFn != nil {
		t.Error("StubEvaluator should have nil CleanupFn")
	}
}
