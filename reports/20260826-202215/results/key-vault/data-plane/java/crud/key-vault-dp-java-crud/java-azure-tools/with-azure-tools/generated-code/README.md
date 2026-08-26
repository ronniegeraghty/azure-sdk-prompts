# Azure Key Vault secret CRUD with Java

This Maven console application creates `my-secret`, reads and prints its value,
creates a new version with the value `updated-value`, then soft-deletes and
purges the secret.

## Prerequisites

- Java 17 and Maven
- An existing Azure Key Vault with soft delete enabled and purge protection
  disabled
- An identity authorized to set, get, delete, and purge secrets
- Local authentication available to `DefaultAzureCredential`, or managed
  identity when hosted in Azure

Set the vault URL and run the application from PowerShell:

```powershell
$env:AZURE_KEY_VAULT_URL = "https://<vault-name>.vault.azure.net/"
mvn compile exec:java
```

`DefaultAzureCredential` supports development credentials and managed identity
without storing credentials in source code. Purging is permanent and requires
the Key Vault purge permission.

References:

- [Azure Key Vault Secrets client library for Java](https://learn.microsoft.com/java/api/overview/azure/security-keyvault-secrets-readme)
- [Azure Identity client library for Java](https://learn.microsoft.com/java/api/overview/azure/identity-readme)
