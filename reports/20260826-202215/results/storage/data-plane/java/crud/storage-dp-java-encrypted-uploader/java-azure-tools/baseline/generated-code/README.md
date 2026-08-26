# Encrypted Azure Blob uploader

Java 17 example using envelope encryption:

- AES-256-GCM encrypts blob data locally with a fresh data encryption key (DEK).
- Azure Key Vault **Keys** wraps and unwraps the DEK with RSA-OAEP-256.
- The wrapped DEK, key version ID, IV, algorithm, and format version are stored as blob metadata.
- One managed-identity credential instance is shared by all synchronous and asynchronous clients.

The container and RSA key must already exist. The managed identity needs Blob Data Contributor
access on the container and `get`, `wrapKey`, and `unwrapKey` permissions on the Key Vault key.

Set:

```text
AZURE_STORAGE_BLOB_ENDPOINT=https://ACCOUNT.blob.core.windows.net
AZURE_STORAGE_CONTAINER_NAME=CONTAINER
AZURE_KEY_VAULT_ENDPOINT=https://VAULT.vault.azure.net
AZURE_KEY_VAULT_KEY_NAME=KEY_NAME
```

Then run in an Azure environment with managed identity:

```text
mvn compile exec:java
```

The demo overwrites `sync-encrypted-demo.bin` and `async-encrypted-demo.bin`.
