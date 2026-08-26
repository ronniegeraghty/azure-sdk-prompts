# Encrypted Azure Blob Uploader

Java 17 sample for client-side envelope encryption with Azure Blob Storage and
Azure Key Vault Keys. Each upload uses a new local AES-256-GCM data encryption
key (DEK). An RSA key in Key Vault wraps the DEK with RSA-OAEP-256, and only the
wrapped DEK, versioned Key Vault key ID, IV, and algorithm identifiers are
stored as blob metadata.

The sample expects an existing blob container and an existing RSA or RSA-HSM
Key Vault key that permits `wrapKey` and `unwrapKey`. It does not create or
modify Azure resources.

## Configuration

Set these environment variables on an Azure host with managed identity:

```text
AZURE_STORAGE_BLOB_ENDPOINT=https://<account>.blob.core.windows.net
AZURE_STORAGE_CONTAINER=<existing-container>
AZURE_KEY_VAULT_URL=https://<vault>.vault.azure.net
AZURE_KEY_NAME=<existing-rsa-key>
AZURE_CLIENT_ID=<optional-user-assigned-managed-identity-client-id>
```

Grant the identity only the required data-plane permissions, such as **Storage
Blob Data Contributor** on the target container and **Key Vault Crypto User**
on the target key.

## Build and run

```text
mvn clean verify
mvn exec:java
```

`Main` performs one synchronous and one asynchronous round trip. The classes
also expose `Path`-based upload and download methods for small files. The sample
buffers each file in memory; use chunked authenticated encryption for large
production files.

References:

- https://learn.microsoft.com/java/api/overview/azure/security-keyvault-keys-readme
- https://learn.microsoft.com/java/api/overview/azure/storage-blob-readme
- https://learn.microsoft.com/java/api/overview/azure/identity-readme
- https://learn.microsoft.com/azure/storage/blobs/client-side-encryption
