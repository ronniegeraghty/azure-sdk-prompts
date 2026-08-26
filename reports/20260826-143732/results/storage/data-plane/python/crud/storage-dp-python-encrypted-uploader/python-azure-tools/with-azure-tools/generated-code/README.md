# Client-side encrypted Azure Blob uploader

This project encrypts blob data locally with a fresh 256-bit AES-GCM data
encryption key (DEK) for every upload. An RSA key in Azure Key Vault wraps the
DEK with RSA-OAEP-256. Only the wrapped DEK, nonce, algorithms, format version,
and versioned Key Vault key ID are persisted as blob metadata.

The sync and async demos each reuse one `DefaultAzureCredential` instance
across their Blob Storage, Key Vault Keys, and cryptography clients. Sync and
async credentials are separate because the Azure SDK exposes different
credential protocols for those execution models.

## Configuration

Create the target container and an RSA key separately, then grant the caller
least-privilege data-plane roles that allow blob read/write and Key Vault
`get`, `wrapKey`, and `unwrapKey` operations. Set:

```text
AZURE_STORAGE_ACCOUNT_URL=https://<account>.blob.core.windows.net
AZURE_STORAGE_CONTAINER=<container>
AZURE_KEY_VAULT_URL=https://<vault>.vault.azure.net
AZURE_KEY_VAULT_KEY_NAME=<rsa-key-name>
AZURE_BLOB_NAME=encrypted-demo.bin
DEMO_PLAINTEXT=optional demo text
```

Install and run:

```text
python -m pip install -r requirements.txt
python main.py
```

For Azure-hosted deployments, use managed identity through
`DefaultAzureCredential` and set `AZURE_TOKEN_CREDENTIALS=prod` to restrict
the credential chain to production-safe credentials.

The one-shot AES-GCM API buffers each file in memory. For very large files,
use a reviewed chunked authenticated-encryption format rather than splitting
AES-GCM ciphertext without a protocol for authenticating chunk order.

## SDK references

- https://learn.microsoft.com/python/api/overview/azure/identity-readme
- https://learn.microsoft.com/python/api/overview/azure/keyvault-keys-readme
- https://learn.microsoft.com/python/api/overview/azure/storage-blob-readme
