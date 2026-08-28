# Azure Key Vault configuration provider

This project provides synchronous and asynchronous Key Vault secret providers,
in-memory caches, expiry-aware refresh, and safe delete-and-recreate rotation.
It authenticates with `DefaultAzureCredential`; no application credentials are
stored in source code.

## Run

1. Install dependencies with `python -m pip install -r requirements.txt`.
2. Set `AZURE_KEY_VAULT_URL` to an HTTPS Key Vault URL.
3. Ensure the current managed identity or developer identity has permissions to
   get, set, delete, purge, and inspect deleted secrets.
4. Run `python main.py`.

The demo never prints secret values. It runs the synchronous flow first and the
asynchronous flow second.

## Rotation behavior

`rotate_secret` and `rotate_secret_async` wait for the
`begin_delete_secret()` long-running operation, purge the soft-deleted secret,
wait until the deleted record disappears, and retry creation if name reuse is
temporarily unavailable. A vault with purge protection enabled intentionally
prevents this delete-and-recreate workflow and the purge error is propagated.
