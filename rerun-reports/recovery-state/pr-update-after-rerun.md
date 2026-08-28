# PR update after the three-way rerun

Use this document to prepare the four language comments and the final description for
[`ronniegeraghty/hyoka#656`](https://github.com/ronniegeraghty/hyoka/pull/656).

## Result selection

1. Start with the 65 valid results retained in the original language runs.
2. Use `rerun-manifest.csv` to locate each replacement result.
3. For every manifest row, prefer the valid rerun over the original invalid or missing result.
4. Require one valid result for every expected prompt/config combination.
5. If any combination remains invalid after retry, exclude that entire prompt triplet from comparative totals and disclose it in the relevant language comment.

## Timeout handling

Treat an SDK `session.idle` timeout as an execution failure, not a failed evaluation:

1. After both lanes finish, rerun each timed-out manifest row in a controlled retry pass.
2. If a retry succeeds, use the successful report and include the prompt triplet normally.
3. Record the number of recovered timeouts as execution metadata, but do not reduce any score.
4. If a combination still times out after the retry pass, exclude its entire three-arm prompt
   triplet so every comparison keeps equal prompts and denominators.
5. List persistent exclusions by prompt ID, config, attempts, and timeout reason.

Do not convert a timeout into zero passed checks, an evaluation failure, or evidence that one
configuration produces worse code.

## Check classification

Report these categories separately:

| Category | Included in scored aggregates | Identification |
|---|---|---|
| Prompt checks | Yes; primary result | Graders from the prompt file |
| Language checks | Yes; supplemental result | `criteria_file` graders with grader type `prompt` |
| Workspace checks | No | Grader type `workspace` |
| Tool/MCP checks | No | Grader type `tool` |

Python's `Output Files Exist` workspace check and `Tool Usage Verification` Azure MCP check
must not contribute to prompt-check totals, language-check totals, pass rates, or headline
comparisons. Show them only as diagnostics. Apply the same rule if workspace or tool graders
are added for another language.

Do not use Hyoka's all-or-nothing evaluation pass count as the experiment headline. A failed
workspace or tool check does not establish that the generated solution is incorrect.

## Comparison rules

The three arms are:

| Short label | Config suffix |
|---|---|
| Baseline | `baseline` |
| Azure Skill + MCP | `azure-skill-mcp` |
| Azure Skill + MCP + Microsoft skill | `azure-skill-mcp-microsoft-skill` |

For each language, calculate:

- Passed/total prompt checks and percentage for each arm.
- Prompt-check difference from baseline for both enhanced arms.
- Prompt-check difference between the Microsoft-skill arm and Azure Skill + MCP.
- Improved, regressed, and tied prompt counts for each enhanced arm versus baseline.
- Improved, regressed, and tied prompt counts for the Microsoft-skill arm versus Azure Skill + MCP.
- The same aggregates separately for language checks when language criteria exist.
- Workspace and MCP tool results as unscored diagnostics.

Compare per-prompt outcomes only when all three arms have valid reports for that prompt.
Cross-language rollups must use prompt checks only. Do not rank languages against each other
because their prompts and criteria differ.

## Per-language PR comment

Post one comment for Python, JavaScript/TypeScript, Java, and .NET.

```markdown
## <Language> three-way comparison

The report includes **<N> complete prompt triplets** and **<3N> valid evaluations**.

Report: [`<report path>`](<committed report URL>)

### Prompt checks

| Arm | Passed | Rate | Difference from baseline |
|---|---:|---:|---:|
| Baseline | <passed>/<total> | <rate>% | — |
| Azure Skill + MCP | <passed>/<total> | <rate>% | <signed difference> |
| Azure Skill + MCP + Microsoft skill | <passed>/<total> | <rate>% | <signed difference> |

| Pairwise prompt outcome | Improved | Regressed | Tied |
|---|---:|---:|---:|
| Azure Skill + MCP vs baseline | <N> | <N> | <N> |
| Microsoft skill vs baseline | <N> | <N> | <N> |
| Microsoft skill vs Azure Skill + MCP | <N> | <N> | <N> |

<Concise prompt-check interpretation. Identify where the additional Microsoft skill helped,
regressed, or made no measurable difference.>

### Language checks

| Arm | Passed | Rate | Difference from baseline |
|---|---:|---:|---:|
| Baseline | <passed>/<total> | <rate>% | — |
| Azure Skill + MCP | <passed>/<total> | <rate>% | <signed difference> |
| Azure Skill + MCP + Microsoft skill | <passed>/<total> | <rate>% | <signed difference> |

<Concise language-check interpretation, or state that no language criteria are configured.>

### Excluded diagnostics

| Diagnostic | Baseline | Azure Skill + MCP | Microsoft skill |
|---|---:|---:|---:|
| Workspace checks | <passed>/<total or N/A> | <passed>/<total or N/A> | <passed>/<total or N/A> |
| Azure MCP usage checks | <passed>/<total or N/A> | <passed>/<total or N/A> | <passed>/<total or N/A> |

Workspace and tool checks are excluded from scored aggregates because they measure the
evaluation environment or observed tool invocation rather than generated-code correctness.

<List any excluded prompt triplets, retries, persistent SDK failures, or other language-specific
limitations. Omit this paragraph when none apply.>
```

When timeouts recover, add a short execution note:

```markdown
### Execution notes

<N> evaluations hit a Copilot SDK `session.idle` timeout during the main rerun. All recovered
in the controlled retry pass and contribute their successful reports to the totals.
```

When timeouts persist, use:

```markdown
### Execution exclusions

The following prompt triplets are excluded from every arm's totals because at least one
configuration did not produce a valid report after retry:

| Prompt ID | Timed-out config | Attempts | Reason |
|---|---|---:|---|
| `<prompt ID>` | `<config>` | <N> | Copilot SDK `session.idle` timeout |
```

For .NET, replace the language-check table with:

```markdown
### Language checks

No generic .NET language criteria are configured, so this section has no scored results.
```

## Final PR description

Preserve the existing config and planned-prompt documentation, then replace the result section
with the following structure.

```markdown
## Results

Prompt checks are the primary task-correctness measure. Language checks are supplemental and
are reported separately. Workspace and tool/MCP checks are excluded from scored aggregates.

### Prompt checks

| Language | Complete triplets | Baseline | Azure Skill + MCP | Difference | Microsoft skill | Difference vs baseline | Difference vs Azure Skill + MCP |
|---|---:|---:|---:|---:|---:|---:|---:|
| [Python](<comment-or-report-link>) | <N> | <P/T (%)> | <P/T (%)> | <delta> | <P/T (%)> | <delta> | <delta> |
| [JavaScript/TypeScript](<link>) | <N> | <P/T (%)> | <P/T (%)> | <delta> | <P/T (%)> | <delta> | <delta> |
| [Java](<link>) | <N> | <P/T (%)> | <P/T (%)> | <delta> | <P/T (%)> | <delta> | <delta> |
| [.NET](<link>) | <N> | <P/T (%)> | <P/T (%)> | <delta> | <P/T (%)> | <delta> | <delta> |
| **Informational rollup** | **<N>** | **<P/T (%)>** | **<P/T (%)>** | **<delta>** | **<P/T (%)>** | **<delta>** | **<delta>** |

The informational rollup combines equivalent prompt checks only. It is not a language ranking.

### Language checks

| Language | Baseline | Azure Skill + MCP | Microsoft skill |
|---|---:|---:|---:|
| Python | <P/T (%)> | <P/T (%)> | <P/T (%)> |
| JavaScript/TypeScript | <P/T (%)> | <P/T (%)> | <P/T (%)> |
| Java | <P/T (%)> | <P/T (%)> | <P/T (%)> |
| .NET | Not configured | Not configured | Not configured |

### Findings

- <Overall effect of Azure Skill + MCP compared with baseline.>
- <Incremental effect of adding the Microsoft language skill.>
- <Languages or prompt groups with meaningful gains or regressions.>
- <Whether Azure MCP was actually invoked, reported only as an unscored diagnostic.>
- <Any excluded triplets or persistent execution failures.>

## Interpretation limits

- This is a single trial per prompt/config combination and is subject to model variance.
- Cross-language scores are not directly comparable because prompt inventories and criteria differ.
- Loaded skills might not be invoked for every prompt.
- MCP invocation is diagnostic evidence, not a generated-code correctness check.
- Workspace checks measure report collection behavior and are excluded from correctness totals.
```

## Publication sequence

1. Confirm all manifest rows resolve to one valid report.
2. Generate and commit the consolidated reports and JSON artifacts.
3. Post the four language comments.
4. Link each language row in the PR description to its comment or committed analysis.
5. Update the PR description with the all-language summary.
6. Replace the progress comment with `216 / 216 complete`.
