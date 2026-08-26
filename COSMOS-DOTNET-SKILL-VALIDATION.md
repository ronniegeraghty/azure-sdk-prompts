# Cosmos DB .NET skill validation

## Summary

Hyoka can compare generated code with and without skills through pairwise tool
ablation. The Cosmos DB .NET data-plane scenario is not ready for a meaningful
skill comparison on the current `main` branch:

- The prompt prescribes SDK APIs instead of describing only the customer task.
- The repository has no .NET pairwise config.
- The repository has no .NET language criteria file.
- The current `microsoft/skills` .NET plugin has no Cosmos DB data-plane skill
  for `Microsoft.Azure.Cosmos`.

Do not run the .NET prompt with `python-pairwise`. Hyoka accepts that combination,
but it loads Python skills for C# generation and produces invalid evidence.

## Teams conclusions

The Teams discussion established these principles:

- Prompts describe what the customer wants, not which SDK classes or methods to
  call.
- A prompt may identify the Azure service, but skills should supply the detailed
  SDK guidance.
- Use the same realistic prompt with and without the relevant skill or MCP
  server. This makes the grounding source the experimental variable.
- Treat the prompt like an interview question: measure how well the skill helps
  the agent solve an open-ended but well-defined task.
- Keep each skill and its `references/examples.md` content consistent. The
  missing generation workflow is tracked by
  [microsoft/skills#311](https://github.com/microsoft/skills/issues/311).

The discussion did not settle how to aggregate repeated runs, such as whether to
use a mean or median score.

## Target prompt

The requested file is
`prompts/cosmos-db/data-plane/dotnet/crud-items.prompt.md`, with prompt ID:

```text
cosmos-db-dp-dotnet-crud
```

The version on `main` names specific methods such as
`CreateDatabaseIfNotExistsAsync`, `CreateItemAsync`, and
`GetItemQueryIterator`. This conflicts with the Teams guidance because the
prompt itself supplies much of the knowledge that the skill is meant to add.

Pan's simplified version exists on the unmerged branch
`origin/JonathanCrd/net-skills-validations`, in commit `cda5e2d3`. It asks the
agent to connect, insert, read, query, and handle failures without prescribing
the exact SDK methods. Use that version as the basis for the experiment.

## Python precedent

The equivalent Python prompt is
`prompts/cosmos-db/data-plane/python/crud-items.prompt.md`, with prompt ID:

```text
cosmos-db-dp-python-crud
```

The general `configs/python-pairwise.yaml` combines:

- Azure MCP
- Local generator skills
- The remote `azure-sdk-python` plugin from `microsoft/skills`
- A general Python skill
- Three generator models

With `--pairwise`, Hyoka creates a baseline and four single-tool ablations. The
three models therefore produce 15 evaluations. The current Python plugin
contains `azure-cosmos-py` and `azure-cosmos-db-py`.

For this investigation, `configs/python-cosmos-skill.yaml` narrows the
experiment to two cases. Azure MCP remains enabled in both cases, and pairwise
mode toggles only the remote `azure-sdk-python` plugin.

The exact command run on August 26, 2026 was:

```powershell
$env:npm_config_registry = "https://pkgs.dev.azure.com/azure-sdk/public/_packaging/azure-sdk-for-js/npm/registry/"
go run .\hyoka run `
  --prompt-id cosmos-db-dp-python-crud `
  --config python-cosmos-skill `
  --pairwise `
  --criteria-dir .\criteria `
  --progress off `
  --yes `
  --log-level debug `
  --log-file .\cosmos-python-gpt56-evaluation.log
```

The config uses `gpt-5.6-sol` for generation and review.

## Python comparison result

The completed report is in `reports/20260826-102427`.

| Case | Skills plugin | Score | Result |
|------|---------------|-------|--------|
| `python-cosmos-skill/baseline` | Included | 10/13 | Failed |
| `python-cosmos-skill/without-azure-sdk-python` | Removed | 9/13 | Failed |

The skill-enabled result passed one additional prompt-specific check:
`database_client.create_database_if_not_exists()`. Both variants missed the
required `enable_cross_partition_query` argument. Both generated
`cosmos_crud.py` and `requirements.txt`.

Both variants also failed two experiment-level checks:

- The workspace grader reported no `*.py` file even though each report contains
  `generated-code/cosmos_crud.py`.
- The Azure MCP usage grader reported no `azure` tool use. The model generated
  code without calling that MCP server.

The summary matrix therefore reports 10/13 with skills and 9/13 without skills.
However, `pairwise.json` reports both variants as 100 with impact 0 because its
top-level pairwise score does not incorporate the per-check totals. Use the
10/13 versus 9/13 result for this run, and treat the pairwise impact field as a
reporting defect.

## Prior Python report status

No prior Cosmos DB Python report was found in:

- The local `reports` directory, which does not exist in this checkout
- Committed repository history
- Pull request #623 discussion
- Xiang's `xiangyan99/hyoka` fork
- The referenced Teams group chat

Teams confirms that Xiang completed several Python prompt runs. On April 22,
2026, Xiang said the evaluation required changing
`baseline-sonnet-skills.yaml` from `azure-sdk-java@skills` to
`azure-sdk-python@skills`. On April 23, Xiang wrote:

> I have finished several runs with Python prompts. I can share my observations
> in the meeting.

That message has no attachment. Searches for reports, results, scores, and
observations found no artifact or score summary in the chat. The report was
therefore likely retained only in the runner's ignored local `reports`
directory or presented verbally in the meeting.

## Current .NET skill gap

The current
[`microsoft/skills` .NET plugin](https://github.com/microsoft/skills/tree/main/.github/plugins/azure-sdk-dotnet)
contains `azure-resource-manager-cosmosdb-dotnet`. That skill covers management
operations through `Azure.ResourceManager.CosmosDB` and explicitly excludes
document CRUD through `Microsoft.Azure.Cosmos`.

The Teams-referenced
[microsoft/skills#301](https://github.com/microsoft/skills/pull/301) also
updated only the management-plane Cosmos DB skill. A data-plane .NET skill must
be added or restored before this prompt can measure Cosmos DB skill impact.

## Validation workflow

Validate prompt, config, and criteria syntax with:

```powershell
go run .\hyoka validate
```

The current tree passes validation with 92 prompts, 15 configs, and 5 criteria
files.

The run initially failed before evaluation because this repository pins
`github.com/github/copilot-sdk/go` v0.2.0. That SDK decodes the CLI ping
timestamp as an integer, while Copilot CLI 1.0.81-9 returns an ISO timestamp.
The report above was produced with a temporary local compatibility shim. Hyoka
needs a full upgrade to Copilot SDK v1 because v1 also changes permission and
typed session-event APIs; changing only the dependency version does not compile.

The .NET comparison completed on August 26, 2026:

- Run ID: `20260826-111051`
- With `azure-sdk-dotnet`: 7/7 in 148.3 seconds, with three files generated
- Without `azure-sdk-dotnet`: 7/7 in 293.4 seconds, with two files generated
- Measured impact: 0
- Report: [`reports/20260826-111051/summary.md`](reports/20260826-111051/summary.md)

Both variants called Azure MCP and generated buildable implementations. The
skill-enabled variant loaded the full .NET plugin, including the management-plane
Cosmos DB skill, but no data-plane skill. The equal scores therefore don't show
whether a relevant `Microsoft.Azure.Cosmos` skill would help.

Reproduce the comparison with:

```powershell
$env:npm_config_registry = "https://pkgs.dev.azure.com/azure-sdk/public/_packaging/azure-sdk-for-js/npm/registry/"
go run .\hyoka run `
  --prompt-id cosmos-db-dp-dotnet-crud `
  --config dotnet-cosmos-skill `
  --pairwise `
  --criteria-dir .\criteria `
  --progress off `
  --yes
```

`configs/dotnet-cosmos-skill.yaml` uses `gpt-5.6-sol` for generation and review.
It holds Azure MCP constant and produces two variants:

1. `dotnet-cosmos-skill/baseline`
2. `dotnet-cosmos-skill/without-azure-sdk-dotnet`

The current remote plugin has no `Microsoft.Azure.Cosmos` data-plane skill. Do
not interpret this result as a data-plane skill comparison.

## JavaScript/TypeScript and Java configs

The report branch also includes focused configs for languages that have
relevant Cosmos DB data-plane skills:

| Language | Config | Skill |
|----------|--------|-------|
| JavaScript/TypeScript | `js-ts-cosmos-skill` | `azure-cosmos-ts` |
| Java | `java-cosmos-skill` | `azure-cosmos-java` |

Run them with:

```powershell
go run .\hyoka run `
  --prompt-id cosmos-db-dp-js-ts-crud `
  --config js-ts-cosmos-skill `
  --pairwise `
  --criteria-dir .\criteria `
  --progress off `
  --yes

go run .\hyoka run `
  --prompt-id cosmos-db-dp-java-crud `
  --config java-cosmos-skill `
  --pairwise `
  --criteria-dir .\criteria `
  --progress off `
  --yes
```

Each config holds Azure MCP constant and toggles only its language SDK plugin.

## Required preparation

Before running the .NET evaluation:

1. Replace the current prompt with the task-level wording from
   `origin/JonathanCrd/net-skills-validations`, or an equivalent rewrite.
2. Add or identify a `Microsoft.Azure.Cosmos` data-plane .NET skill in
   `microsoft/skills`. Pin the config's `version:` when evaluating a branch,
   tag, or commit instead of the default branch.
3. Add `criteria/language/dotnet.yaml` with build, package, partition-key,
   parameterized-query, and Cosmos-specific error-handling checks.
4. Run one prompt and one config first, then repeat enough times to account for
   model variance.
