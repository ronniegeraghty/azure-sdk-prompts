---
id: test-dp-test-hello-markdown
properties:
  service: test
  plane: data-plane
  language: test
  category: test
  difficulty: basic
  description: 'Minimal markdown output for fast grader iteration'
  created: '2025-04-24'
  author: switch
tags:
- test-fixture
- markdown
---

# Hello Markdown Test

## Prompt

Write a file called `hello.md` to the workspace containing:
- A level-1 heading with the text "Hello"
- A single bullet list with exactly three items: "First", "Second", "Third"

Do not write any other files. Complete this task in one turn.

## Evaluation Criteria

The generated output should include:
- A file named `hello.md`
- Proper markdown heading syntax (single `#`)
- Exactly three bullet list items

## Context

This is a minimal test fixture designed for fast grader iteration. The AI should
complete this in seconds with minimal token usage.
