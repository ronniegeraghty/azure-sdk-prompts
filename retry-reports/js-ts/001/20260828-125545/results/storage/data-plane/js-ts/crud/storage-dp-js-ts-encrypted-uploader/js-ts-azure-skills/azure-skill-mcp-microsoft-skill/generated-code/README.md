# Azure client-side encrypted blob uploader

This TypeScript example encrypts data locally with a fresh AES-256-GCM data
encryption key (DEK) for every upload. An RSA key in Azure Key Vault wraps the
DEK with RSA-OAEP-256. Only the ciphertext, wrapped DEK, versioned Key Vault key
ID, IV, and authentication tag are persisted as blob data and metadata.

The plaintext DEK exists only in process memory while encryption or decryption
is running and is zeroed afterward. The Key Vault key material never leaves Key
Vault.

## Prerequisites

- Node.js 20 or later
- A pre-existing Blob Storage container
- A pre-existing RSA or RSA-HSM Key Vault key enabled for `wrapKey` and
  `unwrapKey`
- A managed identity with:
  - `Storage Blob Data Contributor` on the target container or storage account
  - `Key Vault Crypto User` on the target key or vault

No connection strings, account keys, secrets, or local developer credential
fallbacks are used.

## Configure and run

Set the variables shown in `.env.example` in the process environment. Use
`AZURE_MANAGED_IDENTITY_CLIENT_ID` only for a user-assigned managed identity;
omit it for a system-assigned identity.

```powershell
npm install
npm run demo
```

The demo uploads `sample.txt.encrypted` by default, downloads it, unwraps the
DEK through Key Vault, decrypts locally, and prints the versioned vault key ID,
the wrapped DEK in base64, and the recovered plaintext.

The container and Key Vault key must already exist. This project never creates,
updates, or deletes Azure resources.
