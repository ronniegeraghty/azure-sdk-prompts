# Batch-container reference applications

These applications implement the four `batch-container.prompt.md` files. Restoring
or building them doesn't call Azure. Running them requires
`AZURE_DOCUMENT_TRANSLATION_ENDPOINT`, Azure credentials, source and target
container URIs, and a target language.

Each application waits for the batch operation, enumerates every document status,
and prints returned failure code and message details.
