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

# Inline graders — specific to THIS prompt's contract (rust code block,
# no extra source files). Distinct from the generic test-language graders
# in criteria/language/test.yaml.
graders:
  - type: workspace
    name: Code-block prompt contract
    weight: 1.0
    checks:
      - kind: require_to_create
        files: [hello.md]
      # hello.md must include the fenced rust block — checked by substring,
      # not by a generic "any code block" rule.
      - kind: file
        name: hello.md
        state: present
        min_bytes: 100
        contains: "```rust"
      - kind: file
        name: hello.md
        state: present
        contains: "fn main"
      # The prompt explicitly says "Do not write any other files." Specifically,
      # the rust snippet must live INSIDE hello.md as a fenced block, not as a
      # separate Rust source file or build manifest.
      - kind: forbidden_to_create
        files: [main.rs, hello.rs, Cargo.toml, Cargo.lock]
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
