---
id: storage-dp-js-ts-encrypted-uploader
properties:
  service: storage
  plane: data-plane
  language: js-ts
  category: crud
  difficulty: advanced
  description: >
    Can an agent implement client-side envelope encryption for Azure Blob
    Storage using Key Vault Keys for key wrapping, with AES-GCM local
    encryption via Node.js crypto, wrapped DEK and auth tag stored as blob
    metadata, and proper key material lifecycle?
  sdk_package: '@azure/storage-blob'
  doc_url: https://learn.microsoft.com/en-us/javascript/api/overview/azure/storage-blob-readme
  created: '2026-04-30'
  author: kaghiya
tags:
  - blob-storage
  - key-vault
  - encryption
  - envelope-encryption
  - aes-gcm
  - key-wrap
  - cryptography-client
  - multi-service
---

# Encrypted Uploader: Azure Blob Storage + Key Vault Keys (TypeScript)

## Prompt

Create a TypeScript Node.js project that uploads files to Azure Blob Storage with client-side encryption, where the encryption key material is managed in Azure Key Vault.

The project needs:

- A **key management class** that interacts with Azure Key Vault's Keys service (not Secrets) to perform cryptographic operations. It should implement envelope encryption: generate a data encryption key locally, use Key Vault to protect (wrap) it, and store the protected key alongside the encrypted blob. For decryption, have Key Vault recover (unwrap) the data key, then decrypt locally. The raw data key should never be persisted anywhere, and the vault's key material should never leave Key Vault.

- A **blob uploader/downloader class** that handles the actual encryption and storage. For upload: generate a data key, encrypt the data locally using AES-GCM, protect the data key via Key Vault, then upload the ciphertext to Blob Storage with the protected key and any necessary cryptographic parameters stored as blob metadata (including the initialization vector and the authentication tag, which in Node.js is separate from the ciphertext). For download: read the blob and its metadata, recover the data key via Key Vault, and decrypt. Handle errors from both services using `RestError` from `@azure/core-rest-pipeline` with `statusCode` checks (e.g., the vault key may have been disabled, or the blob may not exist — check for 404, 403, etc.).

- A **configuration module** that builds the necessary Azure connections for both Blob Storage and Key Vault. It should read endpoints from environment variables and authenticate with managed identity. All connections should share a single credential instance.

- A **main script** that demos the full encrypt-upload-download-decrypt round-trip: encrypts and uploads a sample string, then downloads and decrypts it back. Print the vault key ID used, the wrapped DEK (base64), and the decrypted output to verify the round-trip.

Enable SDK diagnostic logging using `@azure/logger` with a configurable log level for debugging.

Include a complete `package.json` with the necessary Azure SDK dependencies and a `tsconfig.json`.

## Evaluation Criteria

### Dependencies (scenario-specific)
- Uses `@azure/keyvault-keys` (Keys, NOT Secrets) — critical distinction
- Uses Node.js built-in `crypto` module for local AES-GCM encryption

### Client Construction (scenario-specific)
- Uses `KeyClient` for key management and `CryptographyClient` for wrap/unwrap operations (NOT `SecretClient`)
- Constructs `CryptographyClient` with the key ID or key name

### Key Vault Keys Patterns (critical)
- Uses `CryptographyClient` for `wrapKey()` and `unwrapKey()` operations
- Specifies RSA key wrap algorithm (e.g., `"RSA-OAEP"` or `"RSA-OAEP-256"`)
- Key material never leaves Key Vault (wrap/unwrap is server-side)

### Envelope Encryption Patterns (critical)
- Generates a random AES-256 DEK locally (32 bytes via `crypto.randomBytes`)
- Encrypts data with AES-GCM locally using the DEK (`crypto.createCipheriv("aes-256-gcm", ...)`)
- Wraps the DEK via Key Vault `wrapKey()`
- Stores wrapped DEK as blob metadata (base64-encoded)
- Stores IV (initialization vector) in blob metadata (base64-encoded)
- Stores GCM auth tag in blob metadata (base64-encoded) — in Node.js the auth tag is separate from ciphertext
- Stores vault key identifier in blob metadata
- For decryption: retrieves wrapped DEK from metadata, unwraps via Key Vault, sets auth tag via `decipher.setAuthTag()`, decrypts locally

### AES-GCM (Node.js specific)
- Uses AES-GCM (not AES-CBC, AES-ECB, or other modes)
- Generates random IV for each encryption (typically 12 bytes for GCM)
- Retrieves auth tag via `cipher.getAuthTag()` after encryption finalize
- Sets auth tag via `decipher.setAuthTag()` before decryption finalize

### Scenario-Specific Error Handling
- Handles Key Vault errors (key disabled, key not found) via RestError
- Handles Storage errors (blob not found) via RestError with statusCode

### Anti-Patterns (scenario-specific)
- NOT using `SecretClient` instead of `KeyClient`/`CryptographyClient`
- NOT encrypting data directly with the vault key (should be envelope encryption)
- NOT storing raw DEK in plaintext
- NOT omitting the GCM auth tag from blob metadata (decryption will fail)

## Context

This is the most advanced scenario, combining two Azure services (Blob Storage + Key Vault)
with client-side cryptography. It tests a critical security pattern: envelope encryption where
a locally-generated data encryption key (DEK) is used to encrypt data with AES-GCM, then the
DEK itself is protected by wrapping it with an RSA key managed in Key Vault. The vault's key
material never leaves the HSM boundary. LLMs frequently make two mistakes here: using Key Vault
Secrets instead of Keys (wrong service), and encrypting data directly with the vault key instead
of implementing the DEK/KEK envelope pattern. In Node.js specifically, the GCM authentication
tag is a separate value from the ciphertext (unlike Java where it's appended), so it must be
explicitly stored and restored — omitting it is a common Node.js-specific failure mode.
