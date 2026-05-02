---
name: markdown-lists
description: Authoritative rules for formatting markdown bullet lists and numbered lists, including marker choice, spacing, blank lines around lists, and nested-item indentation. Use this skill whenever you are about to write or edit a markdown (.md) file that contains any list — including short three-item lists, single-bullet lists, and tasks where the user says things like "a bullet list", "a list of items", "three items", or "numbered steps". Apply it before writing the file so the list formatting is correct on the first pass.
---

# Markdown List Rules

Apply these rules to every list you emit in any `.md` file.

## Bullet lists

- Start each item with `-` followed by a single space (prefer `-` over `*` for consistency).
- Put a blank line **before** the list (after any preceding paragraph) and **after** the list.
- Indent nested items by exactly **2 spaces** relative to the parent item.
- Keep each item concise and on a single line when possible.

## Numbered lists

- Start each item with a number, a period, and a single space: `1. Item`.
- Numbers may be written as `1.` for every item; renderers auto-renumber. Sequential numbering (`1.`, `2.`, `3.`) is also fine.
- Same blank-line and indentation rules as bullet lists.

## Correct examples

Bullet list:

```markdown
Some introductory text.

- First item
- Second item
  - Nested item with 2-space indent
  - Another nested item
- Third item
```

Numbered list:

```markdown
Steps to follow:

1. First step
2. Second step
3. Third step
```

## Common mistakes to avoid

- Missing blank line between the preceding paragraph and the first list item (some renderers won't recognize the list).
- Mixing `-`, `*`, and `+` markers within the same list.
- Inconsistent indentation for nested items (e.g., 3 spaces, or tabs mixed with spaces).
- Using `•` or other non-markdown bullet characters.
