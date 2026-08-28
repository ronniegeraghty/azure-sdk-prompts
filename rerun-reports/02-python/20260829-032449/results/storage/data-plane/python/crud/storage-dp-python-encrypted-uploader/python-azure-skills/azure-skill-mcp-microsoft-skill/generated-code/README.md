# Encrypted Azure Blob uploader

This project encrypts files locally with a fresh 256-bit AES-GCM data
encryption key (DEK) for every upload. Azure Key Vault Keys wraps the DEK with
RSA-OAEP-256. The blob stores only ciphertext and metadata containing the
versioned Key Vault key ID, wrapped DEK, nonce, algorithms, and format version.
The plaintext DEK exists only in process memory and project-owned mutable
buffers are overwritten after use; the Key Vault private key is never exported.

## Configuration

Install the dependencies:

```text
python -m pip install -r requirements.txt
```

Set these required environment variables:

| Variable | Example |
|---|---|
| `AZURE_STORAGE_ACCOUNT_URL` | `https://myaccount.blob.core.windows.net` |
| `AZURE_STORAGE_CONTAINER_NAME` | `encrypted-files` |
| `AZURE_KEYVAULT_URL` | `https://myvault.vault.azure.net` |
| `AZURE_KEY_NAME` | `blob-wrapping-key` |

Optional variables are `DEMO_INPUT_FILE`, `DEMO_SYNC_BLOB_NAME`,
`DEMO_ASYNC_BLOB_NAME`, `DEMO_SYNC_OUTPUT_FILE`, and
`DEMO_ASYNC_OUTPUT_FILE`.

Authentication uses `DefaultAzureCredential`. The identity needs Blob data
read/write access (for example, **Storage Blob Data Contributor**) and Key
Vault `keys/get`, `keys/wrapKey`, and `keys/unwrapKey` data-plane permissions
(for example, **Key Vault Crypto User**). In production, set
`AZURE_TOKEN_CREDENTIALS=prod` to restrict the credential chain to
production-safe credentials.

Run both round trips:

```text
python main.py
```

The implementation reads each file into memory because AES-GCM is applied as a
single authenticated message. Use it for files that comfortably fit in process
memory.

## References

- [Azure Key Vault Keys Python client](https://learn.microsoft.com/python/api/overview/azure/keyvault-keys-readme)
- [CryptographyClient API](https://learn.microsoft.com/python/api/azure-keyvault-keys/azure.keyvault.keys.crypto.cryptographyclient)
- [Azure Blob Storage Python client](https://learn.microsoft.com/azure/storage/blobs/storage-quickstart-blobs-python)
- [DefaultAzureCredential](https://learn.microsoft.com/python/api/azure-identity/azure.identity.defaultazurecredential)
