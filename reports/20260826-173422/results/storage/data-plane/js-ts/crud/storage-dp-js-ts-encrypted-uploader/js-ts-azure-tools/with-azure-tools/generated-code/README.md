# Azure client-side encrypted blob uploader

This TypeScript Node.js sample uses envelope encryption:

1. It generates a random 256-bit data encryption key (DEK) locally.
2. It encrypts content locally with AES-256-GCM.
3. Azure Key Vault Keys wraps the DEK with a versioned RSA key using
   RSA-OAEP-256.
4. It stores only ciphertext, the wrapped DEK, the key ID, IV, authentication
   tag, and algorithm identifiers in Azure Blob Storage.
5. On download, Key Vault unwraps the DEK and the application decrypts the
   blob locally.

The raw DEK is never persisted and is overwritten in application buffers after
use. The Key Vault key material never leaves Key Vault.

## Prerequisites

- Node.js 20 or later
- An existing Azure Storage account and blob container
- An existing Azure Key Vault RSA key enabled for `wrapKey` and `unwrapKey`
- A system-assigned or user-assigned managed identity

Grant the managed identity only the required data-plane permissions:

- **Storage Blob Data Contributor** scoped to the target container
- **Key Vault Crypto Service Encryption User** scoped to the wrapping key

The application intentionally does not create Azure resources or containers.

## Configure

Copy `.env.example` values into the environment used by the Azure-hosted
workload:

```text
AZURE_STORAGE_BLOB_ENDPOINT=https://<storage-account>.blob.core.windows.net
AZURE_STORAGE_CONTAINER_NAME=encrypted-files
AZURE_KEY_VAULT_URL=https://<vault-name>.vault.azure.net
AZURE_KEY_VAULT_KEY_NAME=blob-encryption-kek
```

For a user-assigned managed identity, also set
`AZURE_MANAGED_IDENTITY_CLIENT_ID`. Otherwise, the system-assigned identity is
used.

## Build and run

```shell
npm install
npm run demo
```

The demo uploads a timestamped blob, downloads and decrypts it, then prints the
versioned Key Vault key ID, wrapped DEK in base64, and decrypted text.

`EncryptedBlobStorage` also provides `uploadFile` and `downloadToFile`. These
methods buffer the complete file in memory and are intended for files below
100 MB. A production implementation for larger files should use chunked
authenticated encryption and staged block uploads.

## References

- [Azure Key Vault Keys JavaScript SDK](https://learn.microsoft.com/javascript/api/overview/azure/keyvault-keys-readme)
- [Azure Blob Storage JavaScript SDK](https://learn.microsoft.com/javascript/api/overview/azure/storage-blob-readme)
- [Managed identities with Azure Identity](https://learn.microsoft.com/azure/developer/javascript/sdk/authentication/azure-hosted-apps)
