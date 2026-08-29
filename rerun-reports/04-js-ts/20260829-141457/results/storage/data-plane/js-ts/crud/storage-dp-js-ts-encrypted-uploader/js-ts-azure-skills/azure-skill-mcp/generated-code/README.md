# Azure client-side encrypted blob uploader

This TypeScript sample uses envelope encryption:

1. Generate a random 256-bit data encryption key (DEK) in local memory.
2. Encrypt the payload locally with AES-256-GCM.
3. Wrap the DEK with a versioned RSA key in Azure Key Vault by using
   `RSA-OAEP-256`.
4. Store only ciphertext, the wrapped DEK, IV, authentication tag, algorithms,
   version, and Key Vault key ID in Blob metadata.
5. On download, ask Key Vault to unwrap the DEK and decrypt locally.

The plaintext DEK is overwritten in memory after use and is never persisted.
The Key Vault private key never leaves Key Vault. The existing Key Vault key
must be an RSA key enabled for `wrapKey` and `unwrapKey` operations.

## Configure and run

Use an existing storage container, Key Vault, RSA key, and managed identity.
Do not put credentials in environment variables. Copy the endpoint and resource
settings from `.env.example` into the host's environment, then run:

```powershell
npm install
npm run build
npm start
```

The managed identity needs data-plane access equivalent to **Storage Blob Data
Contributor** on the container and **Key Vault Crypto User** on the key, scoped
as narrowly as possible. A user-assigned identity can be selected with
`AZURE_CLIENT_ID`; otherwise, the system-assigned identity is used.

The demo intentionally does not create the container or key. It uploads a sample
string, downloads it, and prints the versioned key ID, wrapped DEK, and recovered
plaintext.

## References

- [Azure Key Vault `CryptographyClient`](https://learn.microsoft.com/javascript/api/@azure/keyvault-keys/cryptographyclient)
- [Azure Blob Storage JavaScript client library](https://learn.microsoft.com/javascript/api/overview/azure/storage-blob-readme)
- [Managed identity authentication](https://learn.microsoft.com/entra/identity/managed-identities-azure-resources/overview)
