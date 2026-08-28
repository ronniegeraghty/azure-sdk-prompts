# Azure client-side encrypted blob uploader

This TypeScript sample encrypts content locally with a fresh AES-256-GCM data
encryption key (DEK) for each upload. An RSA key in Azure Key Vault Keys wraps
the DEK with RSA-OAEP-256. Only ciphertext, the wrapped DEK, the versioned Key
Vault key ID, IV, authentication tag, and algorithm identifiers are persisted
as the blob and its metadata.

The plaintext DEK exists only in process memory and is overwritten after use.
The Key Vault private key is used by the Keys cryptography service and never
leaves Key Vault.

## Prerequisites

- Node.js 20 or newer.
- An existing Blob container.
- An existing RSA key in Azure Key Vault with `wrapKey` and `unwrapKey`
  operations enabled.
- A system-assigned or user-assigned managed identity with:
  - Blob read/write access, such as **Storage Blob Data Contributor**, scoped as
    narrowly as practical.
  - Key `get`, `wrapKey`, and `unwrapKey` access, such as **Key Vault Crypto
    User**, scoped as narrowly as practical.

No Secrets client or storage account key is used.

## Configuration

Set these environment variables in the Azure host:

```text
AZURE_STORAGE_BLOB_ENDPOINT=https://<storage-account>.blob.core.windows.net
AZURE_STORAGE_CONTAINER_NAME=encrypted-files
AZURE_KEY_VAULT_URL=https://<key-vault>.vault.azure.net
AZURE_KEY_VAULT_KEY_NAME=<rsa-key-name>
```

For a user-assigned managed identity, also set `AZURE_CLIENT_ID` to its client
ID. If it is omitted, the application uses the host's system-assigned identity.
The application intentionally uses `ManagedIdentityCredential`, so run the demo
on an Azure host with managed identity available.

## Run

```shell
npm install
npm run build
npm start
```

The demo uploads a unique blob, downloads it, and prints the versioned vault key
ID, wrapped DEK in base64, and recovered plaintext. The container must already
exist; the sample does not create or modify Azure infrastructure.
