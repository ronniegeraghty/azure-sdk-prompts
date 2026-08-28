# Azure client-side encrypted blob uploader

This project performs envelope encryption before data reaches Azure Blob
Storage. Each upload gets a random 256-bit AES-GCM data encryption key (DEK).
Azure Key Vault Keys wraps that DEK with RSA-OAEP-256, and the versioned Key
Vault key ID, wrapped DEK, nonce, and algorithms are stored as blob metadata.
Only ciphertext is uploaded. The raw DEK exists only in process memory, and the
Key Vault private key never leaves the vault.

## Configuration

Use an existing storage container and an RSA key in Azure Key Vault. The
authenticated identity needs Blob Data Contributor access to the container and
Key Vault permissions to read the key and perform wrap/unwrap operations.

Set these environment variables:

```text
AZURE_STORAGE_ACCOUNT_URL=https://<account>.blob.core.windows.net
AZURE_STORAGE_CONTAINER_NAME=<container>
AZURE_KEY_VAULT_URL=https://<vault>.vault.azure.net
AZURE_KEY_NAME=<rsa-key-name>
AZURE_KEY_VERSION=<optional-version>
AZURE_BLOB_NAME=encrypted-demo.bin
```

`AZURE_KEY_VERSION` and `AZURE_BLOB_NAME` are optional. Install dependencies,
authenticate in a way supported by `DefaultAzureCredential`, then run:

```text
python -m pip install -r requirements.txt
python -m encrypted_blob.main
python -m encrypted_blob.main path\to\local-file.txt
```

The sync and async demos use separate credential types because Azure async
clients require an async credential. Within each implementation, Blob Storage
and Key Vault share one credential instance. The container and vault key must
already exist; the demo does not provision Azure resources.
