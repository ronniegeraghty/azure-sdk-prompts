# Active Decisions

## Inline Graders Must Encode Prompt-Specific Contract (2026-05-06)

**Decision:** Inline `graders:` on prompt files should encode the **prompt's unique evaluation contract**, not duplicate generic criteria-file graders.

**Rationale:**
- Inline graders that duplicate checks already in `criteria/language/*.yaml` or `criteria/service/*.yaml` don't earn their keep
- They create maintenance drag and obscure what makes each prompt special
- Specialized behavior belongs inline; reusable checks belong in criteria files

**Resolution:**
- If an inline grader is identical to a criteria-file grader → move it to criteria files or specialize it
- If inline grader encodes prompt-specific behavior (e.g., "this prompt must not generate .rs files") → keep it inline and document the unique contract

**Example:**
- `prompts/test/hello-markdown-with-code.prompt.md` inline graders now check Rust code-block + no extra source files (prompt-specific)
- `prompts/test/hello-yaml.prompt.yaml` inline graders check exact bullet labels + no code fence + no extra files (prompt-specific)
- Both removed duplicates that were already in `criteria/language/test.yaml`

**Follow-up:** During docs reviews or prompt creation, enforce this pattern: inline graders are **specializations**, not duplicates.
