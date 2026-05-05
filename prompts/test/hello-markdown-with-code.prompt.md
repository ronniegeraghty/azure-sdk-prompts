---
id: test-dp-test-hello-markdown-with-code
properties:
  service: test
  plane: data-plane
  language: test
  category: test
  difficulty: basic
  description: 'Companion to hello-markdown that satisfies the rust-code-block check for pass-side grader rendering'
  created: '2025-04-24'
  author: neo
tags:
- test-fixture
- markdown
- code-block
---

# Hello Markdown with Code Test

## Prompt

Write a file called `hello.md` to the workspace containing:
- A level-1 heading with the text "Hello"
- A single bullet list with exactly three items: "First", "Second", "Third"
- A fenced code block tagged with the language `rust` containing a trivial Rust snippet (e.g., `fn main() { println!("Hello"); }`)

Do not write any other files. Complete this task in one turn.

## Evaluation Criteria

The generated output should include:
- A file named `hello.md`
- Proper markdown heading syntax (single `#`)
- Exactly three bullet list items
- A fenced code block with the `rust` language tag

## Context

This is a companion test fixture to `hello-markdown.prompt.md` that satisfies the
rust-code-block check in `criteria/language/test.yaml`. It demonstrates the pass-side
rendering of the same grader criteria, providing clean A/B comparison for grader iteration.
