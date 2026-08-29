# Azure Key Vault secret CRUD with Java

This Maven application creates `my-secret`, reads and prints its value, creates
a new version with the value `updated-value`, then soft-deletes and purges the
secret.

## Prerequisites

- JDK 17 and Maven
- An existing Azure Key Vault with soft delete enabled
- An identity available to `DefaultAzureCredential`
- Secret data-plane permissions to set, get, delete, and purge secrets

Set the vault URL without storing credentials in source code:

```powershell
$env:AZURE_KEY_VAULT_URL = "https://<vault-name>.vault.azure.net"
```

For local development, `DefaultAzureCredential` can use an existing Azure CLI
login. In Azure, prefer a managed identity with least-privilege Key Vault RBAC.

## Build and run

```powershell
mvn compile
mvn exec:java
```

Purging is irreversible. If purge protection is enabled on the vault, Azure
Key Vault rejects manual purge until the retention period expires.

Reference: [Azure Key Vault Secret client library for Java](https://learn.microsoft.com/java/api/overview/azure/security-keyvault-secrets-readme?view=azure-java-stable)
