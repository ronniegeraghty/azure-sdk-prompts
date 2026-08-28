# Encrypted Azure Blob Uploader

Java 17 sample using envelope encryption:

- Each upload generates a random AES-256 data encryption key (DEK).
- File contents are encrypted locally with AES-GCM.
- Azure Key Vault Keys wraps the DEK with `RSA-OAEP-256`.
- Blob metadata stores the wrapped DEK, Key Vault key ID, IV, and algorithm identifiers.
- Downloads unwrap the DEK in Key Vault and decrypt locally. Plaintext DEKs only exist briefly in process memory; the application's raw-key buffers are zeroed after use.

The managed identity needs Blob data access to the target container and Key Vault key permissions for `get`, `wrapKey`, and `unwrapKey`. The configured Key Vault key must be an RSA key that supports those operations.

## Configuration

```text
AZURE_STORAGE_BLOB_ENDPOINT=https://<account>.blob.core.windows.net
AZURE_STORAGE_CONTAINER_NAME=<container>
AZURE_KEY_VAULT_ENDPOINT=https://<vault>.vault.azure.net
AZURE_KEY_VAULT_KEY_NAME=<rsa-key-name>
AZURE_CLIENT_ID=<optional-user-assigned-managed-identity-client-id>
```

The container and RSA key must already exist. This project does not create or modify Azure resources.

## Build and run

```text
mvn clean package
mvn exec:java -Dexec.args="C:\path\to\input.txt"
```

Without an argument, `Main` creates a temporary UTF-8 demo file. It performs a sync round trip followed by an async round trip.
