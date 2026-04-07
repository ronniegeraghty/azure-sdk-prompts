---
name: "property-migration"
description: "How properties map works, snake_case convention"
domain: "architecture"
confidence: "high"
source: "hyoka/internal/config/tools.go, hyoka/internal/prompt/prompt.go"
---

## Context

Hyoka uses property-based tool filtering to dynamically match tools to prompts based on metadata (language, service, plane). This allows a single config to work across multiple prompts without hardcoding tool lists. Property migration is the process of updating the property schema when business logic or tool definitions change.

## Property System Overview

### How It Works

1. **Prompt declares properties:**
   ```yaml
   id: key-vault-dp-python-crud
   language: python
   service: key-vault
   plane: data-plane
   ```

2. **Tool defines property filters:**
   ```yaml
   # config.yaml
   tools:
     - name: python_key_vault_sdk
       include: true
       properties:
         language: python
         service: key-vault
   ```

3. **Engine matches at runtime:**
   ```
   Does prompt.language (python) match tool.properties.language (python)? ✓
   Does prompt.service (key-vault) match tool.properties.service (key-vault)? ✓
   → Tool included in session
   ```

4. **No match excludes tool:**
   ```
   Does prompt.language (dotnet) match tool.properties.language (python)? ✗
   → Tool hidden from session
   ```

## Property Convention: snake_case

All property names and values use **snake_case** (not camelCase or kebab-case):

**✓ Correct:**
```yaml
properties:
  language: python            # not "Language" or "PYTHON"
  service: key_vault          # not "keyVault" or "Key-Vault"
  deployment_plane: data_plane  # not "deploymentPlane"
```

**✗ Incorrect:**
```yaml
properties:
  Language: python            # CamelCase in key
  language: Python            # CamelCase in value
  service: Key-Vault          # Kebab-case
```

## Valid Property Names

- **language:** Python, dotnet, java, js_ts, go, rust, cpp
- **service:** identity, key_vault, storage, cosmos_db, app_config, etc.
- **plane:** data_plane, management_plane
- **category:** auth, crud, pagination, streaming, etc.
- **difficulty:** beginner, intermediate, advanced (optional)

## Migration Scenarios

### Scenario 1: Adding a New Property

When you want to filter tools by a new dimension:

1. **Define in prompt frontmatter:**
   ```yaml
   # prompts/storage/python/blob-upload.md
   tier: premium  # New property
   ```

2. **Add to tool filter (optional):**
   ```yaml
   # config.yaml
   tools:
     - name: storage_premium_tools
       properties:
         language: python
         service: storage
         tier: premium  # Now filtered
   ```

3. **Verify tool visibility:**
   ```bash
   go run ./hyoka run --prompt-id storage-dp-python-blob-upload --dry-run
   # Shows: storage_premium_tools included
   
   go run ./hyoka run --prompt-id storage-dp-python-blob-list --dry-run
   # Shows: storage_premium_tools NOT included (no tier=premium)
   ```

### Scenario 2: Renaming a Property

If you rename `service` to `azure_service`:

1. **Update all prompts** (batch script):
   ```bash
   for f in prompts/**/*.md; do
     sed -i 's/^service:/azure_service:/' "$f"
   done
   ```

2. **Update all configs:**
   ```yaml
   # Before
   properties:
     service: key_vault
   
   # After
   properties:
     azure_service: key_vault
   ```

3. **Validate no orphaned properties:**
   ```bash
   grep -r "service:" prompts/  # Should be empty (all renamed)
   grep -r "service:" configs/  # Should be empty
   ```

### Scenario 3: Changing Property Values

If `plane` values change from `data-plane` to `data_plane`:

1. **Update all prompts:**
   ```bash
   for f in prompts/**/*.md; do
     sed -i 's/plane: data-plane/plane: data_plane/' "$f"
     sed -i 's/plane: management-plane/plane: management_plane/' "$f"
   done
   ```

2. **Update all configs:**
   ```yaml
   # Before
   properties:
     plane: data-plane
   
   # After
   properties:
     plane: data_plane
   ```

3. **Test property matching:**
   ```bash
   go run ./hyoka list --plane data_plane
   # Verify correct prompts listed
   ```

## Property Matching Logic

```go
// hyoka/internal/config/tools.go
func matchesProperties(prompt *Prompt, toolProps map[string]string) bool {
    for key, expectedValue := range toolProps {
        actualValue := prompt.GetProperty(key)
        if actualValue != expectedValue {
            return false  // Property mismatch → tool not included
        }
    }
    return true  // All properties match → tool included
}
```

## Best Practices

1. **Use properties consistently:** If a property exists, use it in ALL relevant places (prompts + tools)
2. **Document property semantics:** List valid values in architecture docs or code comments
3. **Test after migration:** Run `go run ./hyoka list` to verify filtering still works
4. **Batch migrations:** Update prompts and configs in the same commit
5. **Avoid redundant properties:** Don't encode info that's already in prompt ID

## Testing Property Matches

Create a test to verify property matching:

```go
func TestPropertyMatching(t *testing.T) {
    tests := []struct {
        prompt   *Prompt
        toolProp map[string]string
        expect   bool
    }{
        {
            prompt: &Prompt{Language: "python", Service: "key_vault"},
            toolProp: map[string]string{
                "language": "python",
                "service": "key_vault",
            },
            expect: true,
        },
        {
            prompt: &Prompt{Language: "dotnet", Service: "key_vault"},
            toolProp: map[string]string{
                "language": "python",  // Mismatch
                "service": "key_vault",
            },
            expect: false,
        },
    }
    
    for _, tt := range tests {
        if matchesProperties(tt.prompt, tt.toolProp) != tt.expect {
            t.Fail()
        }
    }
}
```

## Code Locations

- **Property matching logic:** `hyoka/internal/config/tools.go`
- **Prompt property parsing:** `hyoka/internal/prompt/prompt.go`
- **Tool configuration:** `hyoka/internal/config/config.go`
- **Example prompts:** `hyoka/prompts/**/*.md`

## Anti-Patterns

- Using camelCase or kebab-case for property names (use snake_case)
- Hardcoding tool names instead of using property filters
- Property values that don't match prompt metadata
- Leaving orphaned properties in configs after migration
- Not testing property matching after schema changes
