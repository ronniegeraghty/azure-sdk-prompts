# Encrypted Azure Blob Uploader

Java 17 sample that encrypts files locally with a fresh AES-256-GCM data encryption key (DEK), wraps that DEK with a versioned RSA key in Azure Key Vault, and stores only ciphertext and the wrapped DEK in Azure Blob Storage.

## Prerequisites

- Java 17 and Maven 3.9+
- An existing blob container
- An existing RSA or RSA-HSM Key Vault key enabled for `wrapKey` and `unwrapKey`
- A managed identity with `Storage Blob Data Contributor` on the container and permission to get, wrap, and unwrap the Key Vault key

Set these environment variables:

```text
AZURE_STORAGE_BLOB_ENDPOINT=https://<account>.blob.core.windows.net
AZURE_STORAGE_CONTAINER=<existing-container>
AZURE_KEY_VAULT_ENDPOINT=https://<vault>.vault.azure.net
AZURE_KEY_VAULT_KEY_NAME=<existing-rsa-key>
AZURE_CLIENT_ID=<optional-user-assigned-managed-identity-client-id>
```

Run the demo from an Azure-hosted environment with that managed identity:

```text
mvn compile exec:java
```

The demo performs synchronous and asynchronous round trips. Blob metadata contains the versioned Key Vault key ID, RSA-OAEP-256-wrapped DEK, AES-GCM IV, and algorithm identifiers. The plaintext DEK exists only in process memory and is overwritten after each operation.
