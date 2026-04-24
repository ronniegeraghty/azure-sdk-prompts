---
name: markdown-lists
description: "Guidance on bullet and numbered list formatting. Use when generating markdown files with lists."
applyTo: "**/*.md"
---

# Markdown List Best Practices

When writing bullet lists:

- Start each list item with a dash `-` or asterisk `*` followed by a space
- Use consistent indentation (2 or 4 spaces for nested items)
- Add a blank line before the list starts (after preceding text)
- Keep list items concise and focused

For numbered lists:

1. Start each item with a number followed by a period and space
2. Markdown will auto-renumber if needed, so you can use `1.` for all items
3. Use consistent indentation for nested numbered lists

Example bullet list:

```markdown
Some introductory text.

- First item
- Second item
  - Nested item with 2-space indent
  - Another nested item
- Third item
```

Example numbered list:

```markdown
Steps to follow:

1. First step
2. Second step
3. Third step
```
