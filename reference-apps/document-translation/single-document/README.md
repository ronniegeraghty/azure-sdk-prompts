# Single-document reference applications

These applications implement the four `single-document.prompt.md` files. Restoring
or building them doesn't call Azure. Running them requires
`AZURE_DOCUMENT_TRANSLATION_ENDPOINT`, Azure credentials, an input path, target
language, and output path.

Each application supplies binary document content, filename, and inferred media
type to the single-document client and writes the response bytes unchanged.
