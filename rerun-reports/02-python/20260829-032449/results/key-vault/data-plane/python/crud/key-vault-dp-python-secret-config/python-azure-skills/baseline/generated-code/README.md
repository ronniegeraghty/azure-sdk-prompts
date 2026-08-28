# Azure Key Vault configuration provider

This project provides synchronous and asynchronous secret providers, in-memory
caches, secure `DefaultAzureCredential` client factories, and safe
delete/purge/recreate rotation helpers.

Install dependencies and set the vault URL:

```powershell
python -m pip install -r requirements.txt
$env:AZURE_KEY_VAULT_URL = "https://your-vault.vault.azure.net"
python main.py
```

`main.py` reads `database-url`, `api-key`, and `feature-flags`. Missing secrets
return configured defaults (`feature-flags` defaults to `{}`). To enable the
rotation portion of the demo, set `ROTATION_SECRET_NAME` and
`ROTATION_SECRET_VALUE`. Rotation permanently purges the deleted secret so its
name can be reused; the identity therefore needs secret delete, purge, and set
permissions, and rotation cannot work while purge protection is enabled.

In Azure, assign the workload a managed identity with only the required Key
Vault data-plane permissions. `DefaultAzureCredential` automatically uses that
identity without storing credentials in application code.
