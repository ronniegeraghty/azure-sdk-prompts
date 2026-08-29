# Encrypted Azure Blob Uploader

A Java 17 sample that performs client-side envelope encryption for Azure Blob Storage.
File contents are encrypted locally with a fresh AES-256-GCM data encryption key (DEK).
An RSA key in Azure Key Vault wraps the DEK with RSA-OAEP-256; only the wrapped DEK,
versioned Key Vault key ID, IV, and algorithm identifiers are stored as blob metadata.

## Configuration

The application uses one `ManagedIdentityCredential` instance for all synchronous and
asynchronous clients. Set these environment variables:

```text
AZURE_STORAGE_BLOB_ENDPOINT=https://<account>.blob.core.windows.net
AZURE_STORAGE_CONTAINER=<container>
AZURE_KEY_VAULT_ENDPOINT=https://<vault>.vault.azure.net
AZURE_KEY_VAULT_KEY_NAME=<rsa-key-name>
AZURE_CLIENT_ID=<optional-user-assigned-managed-identity-client-id>
```

The RSA key must already exist and permit `wrapKey` and `unwrapKey`. Grant the managed
identity least-privilege data-plane access to blobs and Key Vault cryptographic operations.
The application does not create or modify Azure resources.

## Run

```text
mvn compile exec:java
mvn compile exec:java -Dexec.args="Text to encrypt"
```

The demo creates the configured blob container if it does not exist, then performs sync and
async upload/download round trips. `uploadFile` and `downloadFile` methods are also available
for `Path` inputs. This compact sample buffers each file in memory and is intended for small
files; use streaming and chunked authenticated encryption for large files.

References:

- https://learn.microsoft.com/azure/storage/blobs/storage-quickstart-blobs-java
- https://learn.microsoft.com/azure/key-vault/keys/about-keys
- https://learn.microsoft.com/java/api/overview/azure/identity-readme
