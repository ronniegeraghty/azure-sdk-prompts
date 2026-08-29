# Azure Key Vault secret CRUD example

This Java 17 application creates `my-secret`, reads and prints its value,
updates it, waits for soft deletion to finish, and then permanently purges it.
Authentication uses `DefaultAzureCredential`.

Set the vault URL and provide credentials supported by
`DefaultAzureCredential`, then run:

```powershell
$env:KEY_VAULT_URL = "https://<vault-name>.vault.azure.net"
mvn compile exec:java
```

For local development, `DefaultAzureCredential` can use environment-based
credentials such as `AZURE_CLIENT_ID`, `AZURE_TENANT_ID`, and
`AZURE_CLIENT_SECRET`. The identity needs secret get, set, delete, and purge
permissions in the target vault.

The required Maven dependencies are:

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-security-keyvault-secrets</artifactId>
    <version>4.9.4</version>
</dependency>
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-identity</artifactId>
    <version>1.16.1</version>
</dependency>
```
