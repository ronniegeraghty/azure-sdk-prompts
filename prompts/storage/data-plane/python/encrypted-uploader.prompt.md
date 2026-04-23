---
id: storage-dp-python-encrypted-uploader
properties:
  service: storage
  plane: data-plane
  language: python
  category: crud
  difficulty: advanced
  description: 'Can an agent implement client-side envelope encryption for Azure Blob Storage using Key Vault Keys for key
    wrapping, with AES-GCM local encryption, wrapped DEK stored as blob metadata, and proper key material lifecycle?

    '
  sdk_package: azure-storage-blob
  doc_url: https://learn.microsoft.com/en-us/python/api/overview/azure/storage-blob-readme
  created: '2026-04-10'
  author: copilot
tags:
- blob-storage
- key-vault
- encryption
- envelope-encryption
- aes-gcm
- key-wrap
- cryptography-client
- multi-service
- async
---

# Encrypted Uploader: Azure Blob Storage + Key Vault Keys (Python)

## Prompt

Create a Python project that uploads files to Azure Blob Storage with client-side encryption, where the encryption key material is managed in Azure Key Vault.

**Write the code to files (use file-write tools, do not reply with code blocks).**

The project needs:

- A **key management module** (both sync and async versions) that interacts with Azure Key Vault's Keys service (not Secrets) to perform cryptographic operations. It should implement envelope encryption: generate a data encryption key locally, use Key Vault to protect (wrap) it, and store the protected key alongside the encrypted blob. For decryption, have Key Vault recover (unwrap) the data key, then decrypt locally. The raw data key should never be persisted anywhere, and the vault's key material should never leave Key Vault.

- A **blob uploader/downloader module** (both sync and async versions) that handles the actual encryption and storage. For upload: generate a data key, encrypt the data locally using **AES-GCM** (authenticated encryption), protect the data key via Key Vault, then upload the ciphertext to Blob Storage with the protected key and any necessary cryptographic parameters (nonce/IV) stored as blob metadata. For download: read the blob and its metadata, recover the data key via Key Vault, and decrypt. Should handle errors from both services (e.g., the vault key may have been disabled, or the blob may not exist).

- A **configuration module** that builds the necessary Azure connections for both Blob Storage and Key Vault. It should read endpoints from environment variables and authenticate with `DefaultAzureCredential`. All connections should share a single credential instance.

- A **main script** that demos both implementations: runs the full encrypt-upload-download-decrypt round-trip using the sync implementation first, then repeats with the async implementation. Print the vault key ID used, the wrapped DEK (base64), and the decrypted output.

Include a `requirements.txt` with the necessary Azure SDK dependencies.

## Evaluation Criteria

### Dependencies
- Uses the Key Vault Keys SDK (NOT Secrets) — critical distinction
- Uses a cryptography library for local AES-GCM encryption

### Key Vault Integration
- Uses Key Vault's cryptographic client for key wrap and unwrap operations (NOT a secrets client)
- Specifies an RSA key wrap algorithm for protecting the DEK
- Key material never leaves Key Vault — wrap/unwrap happens server-side

### Envelope Encryption
- Generates a random 256-bit data encryption key (DEK) locally
- Encrypts data locally with AES-GCM using the DEK
- Wraps the DEK via Key Vault before storing
- Stores the wrapped DEK, nonce/IV, and vault key identifier as blob metadata
- Decryption retrieves metadata, unwraps the DEK via Key Vault, and decrypts locally
- Uses AES-GCM specifically (not CBC, ECB, or other modes)
- Generates a fresh random nonce for each encryption

### Error Handling
- Handles Key Vault errors (key disabled, key not found)
- Handles blob not found
- Code must build and run without import errors or runtime crashes

### Async Support
- Async versions use the async variants of Blob Storage and Key Vault clients

### Anti-Patterns
- Does NOT use a secrets client instead of a keys/crypto client
- Does NOT encrypt data directly with the vault key (must be envelope encryption)
- Does NOT store the raw DEK in plaintext

## Context

This is the most advanced scenario, combining two Azure services (Blob Storage + Key Vault)
with client-side cryptography. It tests a critical security pattern: envelope encryption where
a locally-generated data encryption key (DEK) is used to encrypt data with AES-GCM, then the
DEK itself is protected by wrapping it with an RSA key managed in Key Vault. The vault's key
material never leaves the HSM boundary. LLMs frequently make two mistakes here: using Key Vault
Secrets instead of Keys (wrong service), and encrypting data directly with the vault key instead
of implementing the DEK/KEK envelope pattern.
