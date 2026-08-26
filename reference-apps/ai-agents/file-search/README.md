# Agent file-search reference applications

These applications implement the four `file-search.prompt.md` files. Restoring or
building them doesn't call Azure. Running them requires a Foundry project, model
deployment, and Azure credentials.

Each application uploads the same trail-guide fact, waits for vector indexing,
grounds an agent response with file search, and cleans up every created resource.
