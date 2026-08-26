# Foundry prompt agreement

Foundry prompt sets evaluate the same application scenario across .NET, Java,
JavaScript/TypeScript, and Python.

## Parity

- Keep the application behavior and user-visible inputs identical across all four
  prompts.
- Adapt only package names, client APIs, project structure, and language idioms.
- Apply equivalent scenario-specific checks in every language. A check can use a
  different mechanism when SDK designs differ, but it must verify the same behavior.
- Treat one scenario across four languages as a single prompt set. Review parity
  whenever one member changes.

## Evaluation scope

Prompt-specific criteria verify the selected SDK workflow: client construction,
operation sequence, request values, polling or paging behavior, result traversal,
and resource lifecycle.

Shared criteria under `criteria/` verify cross-cutting concerns such as
authentication, secrets, general error handling, and language conventions. Do not
duplicate those checks in a prompt unless they are essential to the scenario.

## Evidence

Use the current SDK source, README, and samples as the source of truth. Before
accepting a prompt set:

1. Persist a reference implementation for every language under `reference-apps/` so
   developers can inspect and update the exact SDK usage.
2. Restore and build each reference application without requiring live Azure
   resources.
3. When practical, run equivalent evaluator-owned behavioral tests against controlled
   SDK responses. Treat live Foundry execution as optional manual verification.
4. Confirm that deliberately broken implementations fail, including invented APIs,
   omitted workflow steps, incorrect identifiers, and hardcoded output.
5. Use model review only for criteria that executable checks cannot establish.
