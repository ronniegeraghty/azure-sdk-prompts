package utils

import (
"os"
"path/filepath"
"strings"
"testing"
)

// ---------------------------------------------------------------------------
// ExtractJSON tests
// ---------------------------------------------------------------------------

func TestExtractJSON(t *testing.T) {
tests := []struct {
name  string
input string
want  string
}{
{"plain json", `{"a":1}`, `{"a":1}`},
{"with text", `Here is the result: {"a":1} done.`, `{"a":1}`},
{"markdown json fence", "```json\n{\"a\":1}\n```", `{"a":1}`},
{"markdown plain fence", "```\n{\"a\":1}\n```", `{"a":1}`},
{"no json", "hello world", ""},
{"empty string", "", ""},
{"only whitespace", "   \n\t  ", ""},
{"nested json", `{"a":{"b":2}}`, `{"a":{"b":2}}`},
{"only open brace", "{", ""},
{"only close brace", "}", ""},
{"braces wrong order", "}before{", ""},
{"json with leading text", `text before {"key":"val"} text after`, `{"key":"val"}`},
{"multiple json objects picks outermost", `{"a":1} ignore {"b":2}`, `{"a":1} ignore {"b":2}`},
{"json fence with extra whitespace", "```json\n  {\"a\":1}  \n```", `{"a":1}`},
{"json fence missing closing", "```json\n{\"a\":1}", `{"a":1}`},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
got := ExtractJSON(tt.input)
if got != tt.want {
t.Errorf("ExtractJSON(%q) = %q, want %q", tt.input, got, tt.want)
}
})
}
}

// ---------------------------------------------------------------------------
// IsBuildArtifactDir tests
// ---------------------------------------------------------------------------

func TestIsBuildArtifactDir(t *testing.T) {
tests := []struct {
name string
want bool
}{
{"target", true},
{"node_modules", true},
{"__pycache__", true},
{".venv", true},
{"venv", true},
{"bin", true},
{"obj", true},
{"build", true},
{"dist", true},
{"out", true},
{"vendor", true},
{"packages", true},
{".gradle", true},
{".cargo", true},
{"debug", true},
{"release", true},
{"src", false},
{"cmd", false},
{"internal", false},
{"lib", false},
{"docs", false},
{"", false},
{"Build", false},
{"NODE_MODULES", false},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
got := IsBuildArtifactDir(tt.name)
if got != tt.want {
t.Errorf("IsBuildArtifactDir(%q) = %v, want %v", tt.name, got, tt.want)
}
})
}
}

// ---------------------------------------------------------------------------
// ReadDirFiles tests
// ---------------------------------------------------------------------------

func TestReadDirFilesBasic(t *testing.T) {
dir := t.TempDir()
writeFile(t, dir, "main.go", "package main")
mkDir(t, dir, "pkg")
writeFile(t, dir, filepath.Join("pkg", "lib.go"), "package pkg")

files, err := ReadDirFiles(dir)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if len(files) != 2 {
t.Fatalf("file count = %d, want 2", len(files))
}
if files["main.go"] != "package main" {
t.Errorf("main.go content = %q", files["main.go"])
}
if files[filepath.Join("pkg", "lib.go")] != "package pkg" {
t.Errorf("pkg/lib.go content = %q", files[filepath.Join("pkg", "lib.go")])
}
}

func TestReadDirFilesEmptyDir(t *testing.T) {
dir := t.TempDir()
files, err := ReadDirFiles(dir)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if len(files) != 0 {
t.Errorf("expected 0 files, got %d", len(files))
}
}

func TestReadDirFilesSkipsHiddenFiles(t *testing.T) {
dir := t.TempDir()
writeFile(t, dir, "visible.go", "visible")
writeFile(t, dir, ".hidden", "hidden")

files, err := ReadDirFiles(dir)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if _, ok := files[".hidden"]; ok {
t.Error("should skip hidden files")
}
if _, ok := files["visible.go"]; !ok {
t.Error("should include visible files")
}
}

func TestReadDirFilesSkipsHiddenDirs(t *testing.T) {
dir := t.TempDir()
writeFile(t, dir, "root.go", "root")
mkDir(t, dir, ".git")
writeFile(t, dir, filepath.Join(".git", "config"), "gitconfig")

files, err := ReadDirFiles(dir)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if len(files) != 1 {
t.Errorf("file count = %d, want 1", len(files))
}
if _, ok := files[filepath.Join(".git", "config")]; ok {
t.Error("should skip files in hidden dirs")
}
}

func TestReadDirFilesSkipsBuildArtifactDirs(t *testing.T) {
dir := t.TempDir()
writeFile(t, dir, "main.go", "main")
for _, artifact := range []string{"node_modules", "vendor", "__pycache__"} {
mkDir(t, dir, artifact)
writeFile(t, dir, filepath.Join(artifact, "file.txt"), "content")
}

files, err := ReadDirFiles(dir)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if len(files) != 1 {
t.Errorf("file count = %d, want 1 (only main.go)", len(files))
}
}

func TestReadDirFilesSkipsLargeFiles(t *testing.T) {
dir := t.TempDir()
writeFile(t, dir, "small.txt", "small")
large := strings.Repeat("x", 1<<20+1)
writeFile(t, dir, "large.bin", large)

files, err := ReadDirFiles(dir)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if _, ok := files["large.bin"]; ok {
t.Error("should skip files over 1MB")
}
if _, ok := files["small.txt"]; !ok {
t.Error("should include small files")
}
}

func TestReadDirFilesTotalSizeCap(t *testing.T) {
dir := t.TempDir()
chunk := strings.Repeat("a", 1<<20-1) // just under 1MB
for i := 0; i < 11; i++ {
name := "file_" + string(rune('a'+i)) + ".txt"
writeFile(t, dir, name, chunk)
}

files, err := ReadDirFiles(dir)
if err != nil {
t.Fatalf("unexpected error: %v", err)
}
if len(files) > 10 {
t.Errorf("expected at most 10 files under 10MB cap, got %d", len(files))
}
if len(files) == 0 {
t.Error("expected at least some files to be read")
}
}

func TestReadDirFilesNonexistentDir(t *testing.T) {
files, err := ReadDirFiles("/nonexistent/path/that/does/not/exist")
_ = err
if len(files) != 0 {
t.Errorf("expected 0 files for nonexistent dir, got %d", len(files))
}
}

// ---------------------------------------------------------------------------
// Test helpers
// ---------------------------------------------------------------------------

func writeFile(t *testing.T, base, rel, content string) {
t.Helper()
path := filepath.Join(base, rel)
if err := os.WriteFile(path, []byte(content), 0644); err != nil {
t.Fatalf("writeFile %s: %v", rel, err)
}
}

func mkDir(t *testing.T, base, rel string) {
t.Helper()
if err := os.MkdirAll(filepath.Join(base, rel), 0755); err != nil {
t.Fatalf("mkDir %s: %v", rel, err)
}
}
