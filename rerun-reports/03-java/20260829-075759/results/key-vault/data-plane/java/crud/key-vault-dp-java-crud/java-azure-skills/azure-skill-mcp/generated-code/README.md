# Azure Key Vault secret CRUD (Java)

This console application creates `my-secret`, reads and prints its value, creates
a new version with the value `updated-value`, then deletes and purges the secret.
It uses `DefaultAzureCredential`, so no credentials are stored in source code.

Set the URL of an existing vault:

```powershell
$env:AZURE_KEY_VAULT_URL = "https://<vault-name>.vault.azure.net"
```

The authenticated identity needs secret get, set, delete, and purge permissions.
Purge also requires purge protection to be disabled; Azure intentionally blocks
purging while purge protection is enabled.

Build the application without contacting Azure:

```powershell
mvn compile
```

Run it only when you intend to modify the configured vault:

```powershell
mvn exec:java
```

References:

- [Azure Key Vault Secret client library for Java quickstart](https://learn.microsoft.com/azure/key-vault/secrets/quick-create-java)
- [Azure SDK for Java Key Vault Secrets samples](https://github.com/Azure/azure-sdk-for-java/tree/main/sdk/keyvault/azure-security-keyvault-secrets/src/samples/java/com/azure/security/keyvault/secrets)
