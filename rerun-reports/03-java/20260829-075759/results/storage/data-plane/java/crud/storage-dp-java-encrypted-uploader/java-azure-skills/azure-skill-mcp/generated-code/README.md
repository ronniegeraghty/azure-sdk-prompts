# Client-side encrypted Azure Blob uploader

A Java 17 Maven sample that encrypts blob contents locally with a fresh AES-256-GCM data
encryption key (DEK) per upload. Azure Key Vault Keys wraps the DEK with RSA-OAEP-256. Only
the ciphertext, wrapped DEK, versioned Key Vault key ID, IV, and algorithm identifiers are
stored in Blob Storage metadata.

The sample expects the container and an RSA Key Vault key to already exist. It does not create
or modify Azure resources.

## Configuration

Set these environment variables:

| Variable | Value |
|---|---|
| `AZURE_STORAGE_BLOB_ENDPOINT` | `https://<account>.blob.core.windows.net` |
| `AZURE_STORAGE_CONTAINER` | Existing blob container name |
| `AZURE_KEY_VAULT_ENDPOINT` | `https://<vault>.vault.azure.net` |
| `AZURE_KEY_VAULT_KEY_NAME` | Existing RSA key name |
| `AZURE_CLIENT_ID` | Optional client ID for a user-assigned managed identity |

The managed identity needs Blob data read/write access to the container and Key Vault key
`get`, `wrapKey`, and `unwrapKey` data-plane permissions. With Azure RBAC, assign narrowly
scoped roles that provide those operations, such as **Storage Blob Data Contributor** and
**Key Vault Crypto Service Encryption User**.

## Run

Run this from an Azure-hosted environment with managed identity available:

```shell
mvn compile exec:java
```

`Main` performs sync and async upload/download round trips using separate blobs. The Azure
SDK clients share one `ManagedIdentityCredential` instance.

## Cryptographic behavior

- A 256-bit DEK and 96-bit IV are generated with `SecureRandom` for every upload.
- AES-GCM authenticates the ciphertext and detects changes during decryption.
- Key Vault performs RSA-OAEP-256 wrapping and unwrapping; the vault key material never leaves
  Key Vault.
- The plaintext DEK is held only in process memory and its backing byte array is overwritten
  after use. Java and JVM copies cannot be guaranteed to be erased by application code.
- The versioned Key Vault key ID is stored with each blob, so decryption remains tied to the
  exact key version used for wrapping.
