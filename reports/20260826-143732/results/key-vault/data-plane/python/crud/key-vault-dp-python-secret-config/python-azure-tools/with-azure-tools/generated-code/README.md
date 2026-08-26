# Azure Key Vault configuration provider

This project provides synchronous and asynchronous Key Vault secret providers,
in-memory caches with expiry-aware refresh, and soft-delete-aware rotation.

## Setup

1. Create a Python 3.9+ virtual environment and install `requirements.txt`.
2. Grant the application's managed identity the minimum Key Vault data-plane
   permissions it needs. Rotation requires get, set, delete, and purge.
3. Set `AZURE_KEYVAULT_URL` to the vault URL. In production, set
   `AZURE_TOKEN_CREDENTIALS=prod` to constrain `DefaultAzureCredential` to
   production credentials.
4. Set `DEMO_ROTATED_SECRET_VALUE` to the new demo value.
5. Run `python main.py`.

The demo reports whether cached values are available without printing secret
values. It runs the synchronous flow first and then the asynchronous flow.

Rotation permanently purges the soft-deleted secret before recreating it.
Consequently, it requires purge permission and cannot be used when purge
protection is enabled. For most production rotation workflows, creating a new
secret version with `set_secret` is preferable because it preserves rollback
history; this project performs delete/purge/recreate to demonstrate the
explicit lifecycle requested here.

The synchronous SDK exposes deletion as a long-running operation, so the
rotator waits on the poller from `begin_delete_secret()` before purging. The
current asynchronous SDK exposes the equivalent operation as the awaited
`delete_secret()` method rather than a `begin_delete_secret()` poller.

## SDK references

- https://learn.microsoft.com/python/api/overview/azure/keyvault-secrets-readme
- https://learn.microsoft.com/python/api/overview/azure/identity-readme
