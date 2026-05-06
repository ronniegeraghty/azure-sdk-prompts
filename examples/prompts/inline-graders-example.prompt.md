---
id: identity-dp-python-inline-graders
tags:
  - auth
  - msal
  - inline-graders-demo
properties:
  service: identity
  plane: data-plane
  language: python
  category: auth
  difficulty: intermediate
  description: >
    Example demonstrating inline graders defined in markdown frontmatter.
    This prompt shows how inline `graders:` coexist with the traditional
    `## Evaluation Criteria` markdown section.
  sdk_package: azure-identity
  doc_url: https://learn.microsoft.com/en-us/python/api/overview/azure/identity-readme
  created: '2026-05-06'
  author: Oracle

# Inline graders using the unified grader schema
# These are evaluated AFTER implicit "Criteria from prompt file" (from ## Evaluation Criteria)
# and BEFORE matched criteria-file graders.
#
# IMPORTANT: Inline graders forbid `when:` clauses (hard error). They apply only to this
# prompt file and always execute.
graders:
  - type: prompt
    name: MSAL Usage
    weight: 1.5
    prompt: "Review MSAL 1.0+ best practices:"
    checks:
      - Uses PublicClientApplication for user authentication flows
      - Cache token results using FilesystemTokenCache or similar
      - Handles MsalError exceptions explicitly
      - Uses scopes parameter (not resource) for API access

  - type: workspace
    name: Files Created
    weight: 0.5
    checks:
      - kind: require_to_create
        files: [main.py]
      - kind: file
        name: main.py
        state: present
        min_bytes: 100

  - type: tool
    name: Standard Tools Used
    weight: 0.3
    checks:
      - kind: tool_used
        tool: create_file
      - kind: tool_used
        tool: run_terminal_command
---

# Azure Identity: MSAL Authentication (Python)

## Prompt

Write a Python script that authenticates to Azure using the MSAL (Microsoft Authentication Library) library with DefaultAzureCredential as a fallback.

Your script should:

1. Use `PublicClientApplication` from `azure-identity` for interactive user authentication
2. Request a token for the Microsoft Graph API scope: `https://graph.microsoft.com/.default`
3. Handle token caching to avoid prompting for credentials repeatedly
4. Include error handling for `MsalError` exceptions
5. Fall back to `DefaultAzureCredential` if interactive authentication fails

## Evaluation Criteria

- Correctly imports `PublicClientApplication` from `msal` (not from `azure.identity`)
- Instantiates `PublicClientApplication` with the correct `client_id`
- Uses the `acquire_token_interactive()` method with appropriate scopes
- Implements token cache using `SerializableTokenCache` or filesystem-based cache
- Handles `MsalError` and other authentication exceptions
- Uses `DefaultAzureCredential` as fallback in production environments
- Code is executable and demonstrates the complete auth flow

## Notes

This is an example prompt demonstrating inline graders in `.prompt.md` format. The inline graders
(defined in frontmatter above) coexist with this `## Evaluation Criteria` markdown section. Both
are sent to the reviewer as separate criterion groups in the final evaluation.

The inline graders demonstrate:
- **Prompt grader** with custom preamble and explicit checks
- **Workspace grader** validating file creation
- **Tool grader** validating tool usage patterns

Inline graders cannot use `when:` clauses—they always apply to this prompt file.
