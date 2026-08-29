# Azure client-side encrypted blob demo

This Node.js TypeScript project uses envelope encryption. It generates an
AES-256 data encryption key (DEK) locally, encrypts content with AES-GCM, and
wraps the DEK with an RSA key in Azure Key Vault's Keys service. Only the
ciphertext, wrapped DEK, key ID, IV, authentication tag, and algorithm names are
stored in Blob Storage metadata.

## Configuration

The managed identity requires Blob Data Contributor access to the target
container and permission to `get`, `wrapKey`, and `unwrapKey` on the Key Vault
key. Configure these environment variables:

```text
AZURE_STORAGE_BLOB_ENDPOINT=https://<account>.blob.core.windows.net
AZURE_STORAGE_CONTAINER_NAME=<container>
AZURE_KEY_VAULT_URL=https://<vault>.vault.azure.net
AZURE_KEY_NAME=<rsa-key-name>
AZURE_CLIENT_ID=<optional-user-assigned-managed-identity-client-id>
```

The container and RSA key must already exist. The demo does not create or
modify Azure resources.

## Run

```text
npm install
npm run demo
```
