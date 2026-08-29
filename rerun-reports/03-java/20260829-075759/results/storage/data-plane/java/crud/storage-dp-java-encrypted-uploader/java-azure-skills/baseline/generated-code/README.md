# Azure client-side encrypted blob uploader

Java 17 example using envelope encryption:

- A fresh 256-bit AES data encryption key (DEK) is generated for each upload.
- Blob content is encrypted locally with AES-GCM.
- Azure Key Vault Keys wraps the DEK with `RSA-OAEP-256`; Key Vault key material never
  leaves the vault.
- The wrapped DEK, IV, algorithms, format version, and versioned Key Vault key ID are stored
  as blob metadata. The plaintext DEK is kept only in memory and cleared after use.

The sample reads these environment variables:

| Variable | Description |
|---|---|
| `AZURE_STORAGE_BLOB_ENDPOINT` | Blob service endpoint, such as `https://account.blob.core.windows.net` |
| `AZURE_STORAGE_CONTAINER` | Existing container name |
| `AZURE_KEY_VAULT_ENDPOINT` | Vault endpoint, such as `https://vault.vault.azure.net` |
| `AZURE_KEY_VAULT_KEY_NAME` | Existing RSA key name |
| `AZURE_CLIENT_ID` | Optional client ID for a user-assigned managed identity |

The managed identity needs blob data read/write permissions and Key Vault `keys/get`,
`keys/wrapKey`, and `keys/unwrapKey` permissions. The container and RSA key must already exist;
the sample does not create Azure resources.

Run the sync round trip followed by the async round trip:

```text
mvn compile exec:java
```

This compact example buffers each file in memory and is intended for small files. For large
files, use a framed/chunked authenticated-encryption format rather than loading the entire
plaintext or ciphertext into one byte array.
