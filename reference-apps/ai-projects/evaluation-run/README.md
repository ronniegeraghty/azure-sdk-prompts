# Evaluation-run reference applications

These applications implement the four `evaluation-run.prompt.md` files. Restoring
or building them doesn't call Azure. Running them requires
`FOUNDRY_PROJECT_ENDPOINT` and Azure credentials.

Each application obtains the project OpenAI client, creates the same F1 evaluation,
runs it against one inline JSONL item, polls to terminal state, traverses every
output item and metric result, and deletes the evaluation.
