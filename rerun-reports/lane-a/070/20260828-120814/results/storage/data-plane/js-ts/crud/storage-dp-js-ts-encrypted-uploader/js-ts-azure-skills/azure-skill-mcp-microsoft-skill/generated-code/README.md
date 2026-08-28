# Azure encrypted blob uploader

This TypeScript sample encrypts data locally with a fresh AES-256-GCM data
encryption key (DEK) for every upload. Azure Key Vault Keys wraps the DEK with
an RSA key, and Blob Storage stores only the ciphertext plus the versioned key
ID, wrapped DEK, IV, authentication tag, and algorithm identifiers in metadata.

The plaintext DEK exists only in process memory and is overwritten after use.
The Key Vault private key is used remotely by the Keys cryptography API and
never leaves the vault.

## Prerequisites

- Node.js 20 or later.
- An existing blob container.
- An existing RSA Key Vault key enabled for `wrapKey` and `unwrapKey`.
- A managed identity with Blob Data Contributor access to the container and
  Key Vault Crypto User access to the key.

Copy `.env.example` values into the process environment. The application uses
a system-assigned managed identity unless
`AZURE_MANAGED_IDENTITY_CLIENT_ID` selects a user-assigned identity.

## Run

```powershell
npm install
npm run build
npm start
```

The demo uploads `encrypted-demo.txt`, downloads it, authenticates and decrypts
it locally, then prints the versioned vault key ID, wrapped DEK, and plaintext.

`EncryptedBlobClient.uploadFile` and `downloadToFile` provide file-oriented
helpers. They buffer the complete file because AES-GCM authentication is
performed over the complete payload; use this sample for files that fit safely
in application memory.
