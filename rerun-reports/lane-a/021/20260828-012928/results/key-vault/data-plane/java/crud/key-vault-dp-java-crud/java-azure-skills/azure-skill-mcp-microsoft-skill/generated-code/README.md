# Azure Key Vault secret CRUD

This Maven application uses `DefaultAzureCredential` to create, read, update,
soft-delete, and purge the `my-secret` secret. Updating a secret value creates a
new version because existing secret values are immutable.

Set the vault URL and provide any credential supported by
`DefaultAzureCredential`:

```powershell
$env:AZURE_KEYVAULT_URL = "https://<vault-name>.vault.azure.net"
mvn compile exec:java
```

The authenticated identity needs permissions to get, set, delete, and purge
secrets. Purging is not possible when purge protection is enabled.
