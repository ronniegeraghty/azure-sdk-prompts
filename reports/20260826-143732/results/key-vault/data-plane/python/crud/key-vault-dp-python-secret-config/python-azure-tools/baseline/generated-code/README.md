# Azure Key Vault configuration provider

This project provides synchronous and asynchronous Key Vault secret providers,
expiry-aware in-memory caches, and safe secret rotation.

## Setup

Use Python 3.10 or later, then install the dependencies:

```powershell
python -m pip install -r requirements.txt
```

Authentication uses `DefaultAzureCredential`. No application credential is
stored in this project. Set these environment variables before running the
demo:

```powershell
$env:AZURE_KEY_VAULT_URL = "https://your-vault.vault.azure.net"
$env:REQUIRED_CONFIG_KEYS = "database-url,api-key,feature-flags"
$env:ROTATION_SECRET_NAME = "api-key"
$env:ROTATED_SECRET_VALUE = "the-new-secret-value"
python main.py
```

The first two variables after the vault URL are optional and use the demo
defaults when omitted. `ROTATED_SECRET_VALUE` is required.

## Rotation warning

The demo is intentionally destructive: it runs the rotation workflow once
with the synchronous client and again with the asynchronous client. Rotation
waits for deletion, purges the soft-deleted secret, waits for purge propagation,
and then recreates the secret with a 90-day expiry.

The authenticated Azure identity needs secret get, set, delete, and purge
permissions. Rotation cannot recreate the same secret name when purge
protection is enabled; in that case it raises `SecretRotationError` rather than
pretending rotation succeeded.

Run the local tests without contacting Azure:

```powershell
python -m unittest discover -v
```
