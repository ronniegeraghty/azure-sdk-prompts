# Dataset-lifecycle reference applications

These applications implement the four `dataset-lifecycle.prompt.md` files.
Restoring or building them doesn't call Azure. Running them requires a Foundry
project with a default Azure Storage connection and these environment variables:

- `FOUNDRY_PROJECT_ENDPOINT`
- `DATASET_NAME`
- `DATASET_VERSION`
- `DATA_FILE_PATH`
- `DOWNLOAD_FILE_PATH`

Each application uploads one exact dataset version, retrieves typed metadata,
downloads the referenced blob through the language's Azure Storage Blob SDK,
verifies the persisted bytes, and deletes that exact version.
