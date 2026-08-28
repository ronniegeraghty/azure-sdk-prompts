# Client-side encrypted Azure Blob uploader

This project uses envelope encryption for Blob Storage:

1. Generate a fresh 256-bit data encryption key (DEK) in process memory.
2. Encrypt the file locally with AES-256-GCM.
3. Ask Azure Key Vault Keys to wrap the DEK with RSA-OAEP-256.
4. Store only ciphertext, the wrapped DEK, the versioned Key Vault key ID,
   and encryption parameters in blob metadata.
5. On download, ask Key Vault to unwrap the DEK and decrypt locally.

The plaintext DEK is never persisted. The key-encryption key remains in Key
Vault because wrap and unwrap operations use the versioned remote key ID.

## Setup

Use Python 3.9 or later, install the dependencies, and set:

```text
AZURE_STORAGE_ACCOUNT_URL=https://<account>.blob.core.windows.net
AZURE_STORAGE_CONTAINER_NAME=<existing-container>
AZURE_KEY_VAULT_URL=https://<vault>.vault.azure.net
AZURE_KEY_NAME=<existing-rsa-key>
```

Optionally set `DEMO_FILE_PATH`, `DEMO_SYNC_BLOB_NAME`, or
`DEMO_ASYNC_BLOB_NAME`. The default input is `demo-input.txt`.

Authenticate `DefaultAzureCredential` through your normal development
identity, then run:

```text
python main.py
```

The identity needs Blob data read/write access to the container and Key Vault
key get, wrap, and unwrap permissions. With Azure RBAC, assign least-privilege
roles such as **Storage Blob Data Contributor** at container scope and **Key
Vault Crypto User** at key or vault scope.

The container and RSA key must already exist. This project does not provision
or modify Azure resources.

## Local tests

```text
python -m unittest discover -s tests -v
```

Tests use in-memory fakes and do not contact Azure.

## References

- [Azure Key Vault Keys Python client](https://learn.microsoft.com/python/api/overview/azure/keyvault-keys-readme)
- [Manage blob properties and metadata with Python](https://learn.microsoft.com/azure/storage/blobs/storage-blob-properties-metadata-python)
- [Passwordless Azure Storage connections](https://learn.microsoft.com/azure/storage/common/migrate-azure-credentials)
