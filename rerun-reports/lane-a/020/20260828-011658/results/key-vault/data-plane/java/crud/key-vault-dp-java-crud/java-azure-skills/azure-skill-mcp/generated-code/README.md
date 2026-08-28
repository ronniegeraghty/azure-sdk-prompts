# Azure Key Vault Secrets CRUD (Java)

This Maven application creates, reads, updates, deletes, and purges the
`my-secret` secret using `DefaultAzureCredential`.

## Prerequisites

- JDK 17 or later
- Maven 3.9 or later
- An existing soft-delete-enabled Azure Key Vault
- Key Vault data-plane permissions to get, set, delete, and purge secrets

Set the vault URL:

```powershell
$env:AZURE_KEY_VAULT_URL = "https://<vault-name>.vault.azure.net"
```

Configure any credential source supported by `DefaultAzureCredential`. For
example, local development can use environment-based service principal
variables:

```powershell
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:AZURE_CLIENT_ID = "<client-id>"
$env:AZURE_CLIENT_SECRET = "<client-secret>"
```

Do not store credentials in source control.

## Build and run

```powershell
mvn compile
mvn exec:java
```

The Azure SDK BOM in `pom.xml` keeps the `azure-security-keyvault-secrets` and
`azure-identity` dependency versions compatible:

```xml
<dependencies>
    <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-security-keyvault-secrets</artifactId>
    </dependency>
    <dependency>
        <groupId>com.azure</groupId>
        <artifactId>azure-identity</artifactId>
    </dependency>
</dependencies>
```

If purge protection is enabled on the vault, Azure rejects immediate purge
until the configured retention period expires.
