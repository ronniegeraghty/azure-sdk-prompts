package prompt

import (
	"encoding/json"
	"strings"
	"testing"
)

// TestPromptGroupJSONOmitempty verifies that the `group` field on
// prompt.Prompt honors `json:"group,omitempty"`:
//   - When Group is empty, the JSON output must NOT contain a "group" key.
//   - When Group is non-empty, the JSON output must round-trip cleanly.
//
// This guards against accidental removal of the omitempty tag, which would
// pollute every ungrouped prompt's JSON payload with an empty string and
// break downstream "is this prompt grouped?" checks in the site/report
// layer (#606 polish, #608).
func TestPromptGroupJSONOmitempty(t *testing.T) {
	t.Run("empty group is omitted", func(t *testing.T) {
		p := &Prompt{
			ID:         "test-dp-python-sample",
			PromptText: "do the thing",
			Properties: map[string]string{"service": "storage"},
			// Group intentionally zero-valued.
		}
		raw, err := json.Marshal(p)
		if err != nil {
			t.Fatalf("Marshal: %v", err)
		}
		if strings.Contains(string(raw), `"group"`) {
			t.Errorf(`empty Group should be omitted via omitempty; got JSON containing "group": %s`, raw)
		}
	})

	t.Run("non-empty group round-trips", func(t *testing.T) {
		in := &Prompt{
			ID:         "test-dp-python-sample",
			PromptText: "do the thing",
			Group:      "crud-operations",
			Properties: map[string]string{"service": "storage"},
		}
		raw, err := json.Marshal(in)
		if err != nil {
			t.Fatalf("Marshal: %v", err)
		}
		if !strings.Contains(string(raw), `"group":"crud-operations"`) {
			t.Errorf(`expected "group":"crud-operations" in JSON; got: %s`, raw)
		}

		var out Prompt
		if err := json.Unmarshal(raw, &out); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if out.Group != in.Group {
			t.Errorf("Group round-trip mismatch: got %q, want %q", out.Group, in.Group)
		}
	})

	t.Run("absent group unmarshals as empty", func(t *testing.T) {
		raw := []byte(`{"id":"x","prompt_text":"y","properties":{"service":"s"}}`)
		var out Prompt
		if err := json.Unmarshal(raw, &out); err != nil {
			t.Fatalf("Unmarshal: %v", err)
		}
		if out.Group != "" {
			t.Errorf("absent group should unmarshal to empty string; got %q", out.Group)
		}
	})

	t.Run("explicit empty group also omitted on remarshal", func(t *testing.T) {
		// Simulates: load a prompt that had group:"x", clear it, re-marshal.
		p := &Prompt{
			ID:         "test-dp-python-sample",
			PromptText: "do the thing",
			Group:      "",
			Properties: map[string]string{"service": "storage"},
		}
		raw, err := json.Marshal(p)
		if err != nil {
			t.Fatalf("Marshal: %v", err)
		}
		if strings.Contains(string(raw), `"group"`) {
			t.Errorf(`cleared Group should still be omitted via omitempty; got: %s`, raw)
		}
	})
}
