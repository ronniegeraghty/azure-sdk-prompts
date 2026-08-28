# Azure client-side encrypted blob uploader

This project encrypts each payload locally with a fresh AES-256-GCM data
encryption key (DEK). An RSA key in Azure Key Vault's Keys service wraps the
DEK with RSA-OAEP-256. Only the wrapped DEK, nonce, algorithms, version, and
Key Vault key ID are persisted as blob metadata.

## Prerequisites

- An existing Blob Storage account and container.
- An existing RSA key in Azure Key Vault.
- An identity supported by `DefaultAzureCredential` with blob data access and
  Key Vault `keys/get`, `keys/wrapKey`, and `keys/unwrapKey` permissions.

The project does not create or modify Azure resources.

## Run

Create and activate a virtual environment, install `requirements.txt`, and set
the variables shown in `.env.example` in your shell. Then run:

    python main.py

The demo uploads two blobs: one through the synchronous clients and one through
the asynchronous clients. It then downloads and decrypts both.

The sync Azure clients share one synchronous `DefaultAzureCredential` instance.
The async clients share one asynchronous `DefaultAzureCredential` instance;
the Azure SDK requires separate sync and async credential protocols.
