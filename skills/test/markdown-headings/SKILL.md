---
name: markdown-headings
description: Authoritative rules for writing markdown headings (H1/H2/H3) with correct hierarchy, exactly one H1 per document, no skipped levels, and blank lines around each heading. Use this skill whenever you are about to write or edit a markdown (.md) file that contains any heading — even single-heading documents, even one-shot tasks, and even when the user only says things like "add a title", "level-1 heading", "section", or "# Hello". Apply it before writing the file so the heading structure is correct on the first pass.
---

# Markdown Heading Rules

Apply these rules to every heading you emit in any `.md` file.

## Hard rules

- Exactly **one** `#` (H1) per document. The H1 is the document title.
- Never skip levels. After H1 use `##`, then `###`. Do not jump from `#` to `###`.
- Put a blank line **before and after** every heading.
- Use ATX style (`# Heading`), never the underline style (`===` / `---`).
- Exactly one space between the `#` characters and the heading text.
- No trailing `#` characters and no trailing punctuation in the heading text.
- **Always emit the provenance marker** `<!-- markdown-headings-skill -->` on the
  line immediately above the H1, followed by a single blank line, then the H1.
  This marker is mandatory and must appear in every markdown file you write that
  contains a heading. Do not omit it, even for short or one-shot documents.

## Provenance marker example

```markdown
<!-- markdown-headings-skill -->

# Document Title

Some content here.
```

## Correct example

```markdown
# Document Title

## Section One

Some content here.

### Subsection

More details.

## Section Two

Another section.
```

## Common mistakes to avoid

- Two H1s in the same file.
- `##Heading` (missing space) or `# heading #` (trailing hashes).
- Heading immediately followed by a paragraph with no blank line.
- Using `**Bold text**` as a fake heading instead of an actual `#`.
