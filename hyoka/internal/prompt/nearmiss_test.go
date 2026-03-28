package prompt

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScanNearMisses_HyphenatedPrompt(t *testing.T) {
	dir := t.TempDir()
	// Create a file with hyphen instead of dot: storage-prompt.md
	os.WriteFile(filepath.Join(dir, "storage-prompt.md"), []byte("# test"), 0644)

	misses := ScanNearMisses(dir)
	if len(misses) != 1 {
		t.Fatalf("expected 1 near miss, got %d: %v", len(misses), misses)
	}
	if misses[0] != "storage-prompt.md" {
		t.Errorf("expected 'storage-prompt.md', got %q", misses[0])
	}
}

func TestScanNearMisses_WrongExtension(t *testing.T) {
	dir := t.TempDir()
	// Create a file with wrong extension: auth.prompt.txt
	os.WriteFile(filepath.Join(dir, "auth.prompt.txt"), []byte("# test"), 0644)

	misses := ScanNearMisses(dir)
	if len(misses) != 1 {
		t.Fatalf("expected 1 near miss, got %d: %v", len(misses), misses)
	}
	if misses[0] != "auth.prompt.txt" {
		t.Errorf("expected 'auth.prompt.txt', got %q", misses[0])
	}
}

func TestScanNearMisses_FrontmatterMd(t *testing.T) {
	dir := t.TempDir()
	// Create an .md file with YAML frontmatter (looks like a prompt but wrong name)
	content := "---\nid: test-prompt\nservice: storage\n---\n\n## Prompt\n\nDo something.\n"
	os.WriteFile(filepath.Join(dir, "storage-auth.md"), []byte(content), 0644)

	misses := ScanNearMisses(dir)
	if len(misses) != 1 {
		t.Fatalf("expected 1 near miss, got %d: %v", len(misses), misses)
	}
	if misses[0] != "storage-auth.md" {
		t.Errorf("expected 'storage-auth.md', got %q", misses[0])
	}
}

func TestScanNearMisses_CorrectFilesIgnored(t *testing.T) {
	dir := t.TempDir()
	// A correctly named prompt file should not appear as a near miss
	os.WriteFile(filepath.Join(dir, "storage-auth.prompt.md"), []byte(testPromptContent), 0644)
	// A plain .md without frontmatter should also be ignored
	os.WriteFile(filepath.Join(dir, "readme.md"), []byte("# Just a readme"), 0644)
	// A non-.md file should be ignored
	os.WriteFile(filepath.Join(dir, "data.json"), []byte("{}"), 0644)

	misses := ScanNearMisses(dir)
	if len(misses) != 0 {
		t.Errorf("expected 0 near misses, got %d: %v", len(misses), misses)
	}
}

func TestScanNearMisses_MultiplePatterns(t *testing.T) {
	dir := t.TempDir()
	// Create multiple near-miss patterns
	os.WriteFile(filepath.Join(dir, "auth-prompt.md"), []byte("# test"), 0644)
	os.WriteFile(filepath.Join(dir, "crud.prompt.txt"), []byte("# test"), 0644)
	content := "---\nid: test\n---\n## Prompt\n\nHello.\n"
	os.WriteFile(filepath.Join(dir, "keyvault.md"), []byte(content), 0644)
	// Correct file — should not appear
	os.WriteFile(filepath.Join(dir, "correct.prompt.md"), []byte(testPromptContent), 0644)

	misses := ScanNearMisses(dir)
	if len(misses) != 3 {
		t.Errorf("expected 3 near misses, got %d: %v", len(misses), misses)
	}
}

func TestScanNearMisses_NestedDirectories(t *testing.T) {
	dir := t.TempDir()
	sub := filepath.Join(dir, "storage", "data-plane")
	os.MkdirAll(sub, 0755)
	os.WriteFile(filepath.Join(sub, "auth-prompt.md"), []byte("# test"), 0644)

	misses := ScanNearMisses(dir)
	if len(misses) != 1 {
		t.Fatalf("expected 1 near miss, got %d: %v", len(misses), misses)
	}
	// Should be a relative path
	expected := filepath.Join("storage", "data-plane", "auth-prompt.md")
	if misses[0] != expected {
		t.Errorf("expected %q, got %q", expected, misses[0])
	}
}

func TestScanNearMisses_EmptyDirectory(t *testing.T) {
	dir := t.TempDir()
	misses := ScanNearMisses(dir)
	if len(misses) != 0 {
		t.Errorf("expected 0 near misses for empty dir, got %d", len(misses))
	}
}

func TestScanNearMisses_Deduplication(t *testing.T) {
	dir := t.TempDir()
	// A file that matches both hyphenated AND has frontmatter should only appear once
	content := "---\nid: test\n---\n"
	os.WriteFile(filepath.Join(dir, "auth-prompt.md"), []byte(content), 0644)

	misses := ScanNearMisses(dir)
	if len(misses) != 1 {
		t.Errorf("expected 1 near miss (deduplicated), got %d: %v", len(misses), misses)
	}
}

func TestSuggestFix(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"hyphen to dot", "auth-prompt.md", "auth.prompt.md"},
		{"prompt.txt to prompt.md", "auth.prompt.txt", "auth.prompt.md"},
		{"prompt.yaml to prompt.md", "auth.prompt.yaml", "auth.prompt.md"},
		{"nested hyphen", "sub/dir/auth-prompt.md", "sub/dir/auth.prompt.md"},
		{"nested wrong ext", "sub/crud.prompt.txt", "sub/crud.prompt.md"},
		{"no obvious fix", "readme.md", ""},
		{"plain file", "data.json", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := suggestFix(tt.input)
			if got != tt.want {
				t.Errorf("suggestFix(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestLoadPrompts_ZeroPromptsWithNearMissSuggestions(t *testing.T) {
	dir := t.TempDir()
	// Create a near-miss file
	os.WriteFile(filepath.Join(dir, "auth-prompt.md"), []byte("# test"), 0644)

	_, err := LoadPrompts(dir)
	if err == nil {
		t.Fatal("expected error for zero prompts")
	}
	errMsg := err.Error()
	if !contains(errMsg, "no prompts found") {
		t.Errorf("expected 'no prompts found' in error, got %q", errMsg)
	}
	if !contains(errMsg, "Did you mean") {
		t.Errorf("expected 'Did you mean' suggestion in error, got %q", errMsg)
	}
	if !contains(errMsg, "auth-prompt.md") {
		t.Errorf("expected near-miss file in error, got %q", errMsg)
	}
}

func TestLoadPrompts_ZeroPromptsNoNearMisses(t *testing.T) {
	dir := t.TempDir()
	// Only non-prompt files
	os.WriteFile(filepath.Join(dir, "readme.txt"), []byte("# readme"), 0644)

	_, err := LoadPrompts(dir)
	if err == nil {
		t.Fatal("expected error for zero prompts")
	}
	errMsg := err.Error()
	if !contains(errMsg, "no prompts found") {
		t.Errorf("expected 'no prompts found' in error, got %q", errMsg)
	}
	if contains(errMsg, "Did you mean") {
		t.Errorf("did not expect suggestions when no near misses, got %q", errMsg)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && containsSubstr(s, substr))
}

func containsSubstr(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
