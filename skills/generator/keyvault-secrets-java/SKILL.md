---
name: azure-security-keyvault-secrets-java
description: Azure Key Vault Secrets Java SDK for secret management. Use when storing, retrieving, or managing passwords, API keys, connection strings, or other sensitive configuration data.
package: com.azure:azure-security-keyvault-secrets
---

# Azure Key Vault Secrets (Java)

Securely store and manage secrets like passwords, API keys, and connection strings.

## Installation

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-security-keyvault-secrets</artifactId>
    <version>4.9.0</version>
</dependency>
```

## Client Creation

```java
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;
import com.azure.identity.DefaultAzureCredentialBuilder;

// Sync client
SecretClient secretClient = new SecretClientBuilder()
    .vaultUrl("https://<vault-name>.vault.azure.net")
    .credential(new DefaultAzureCredentialBuilder().build())
    .buildClient();

// Async client
SecretAsyncClient secretAsyncClient = new SecretClientBuilder()
    .vaultUrl("https://<vault-name>.vault.azure.net")
    .credential(new DefaultAzureCredentialBuilder().build())
    .buildAsyncClient();
```

## Create/Set Secret

```java
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;

// Simple secret
KeyVaultSecret secret = secretClient.setSecret("database-password", "P@ssw0rd123!");
System.out.println("Secret name: " + secret.getName());
System.out.println("Secret ID: " + secret.getId());

// Secret with options
KeyVaultSecret secretWithOptions = secretClient.setSecret(
    new KeyVaultSecret("api-key", "sk_live_abc123xyz")
        .setProperties(new SecretProperties()
            .setContentType("application/json")
            .setExpiresOn(OffsetDateTime.now().plusYears(1))
            .setNotBefore(OffsetDateTime.now())
            .setEnabled(true)
            .setTags(Map.of(
                "environment", "production",
                "service", "payment-api"
            ))
        )
);
```

## Get Secret

```java
// Get latest version
KeyVaultSecret secret = secretClient.getSecret("database-password");
String value = secret.getValue();

// Get specific version
KeyVaultSecret specificVersion = secretClient.getSecret("database-password", "<version-id>");
```

## List Secrets

```java
import com.azure.security.keyvault.secrets.models.SecretProperties;

for (SecretProperties secretProps : secretClient.listPropertiesOfSecrets()) {
    System.out.println("Secret: " + secretProps.getName());
    System.out.println("  Enabled: " + secretProps.isEnabled());
}

// List versions of a secret
for (SecretProperties version : secretClient.listPropertiesOfSecretVersions("database-password")) {
    System.out.println("Version: " + version.getVersion());
}
```

## Delete Secret

```java
import com.azure.core.util.polling.SyncPoller;
import com.azure.security.keyvault.secrets.models.DeletedSecret;

SyncPoller<DeletedSecret, Void> deletePoller = secretClient.beginDeleteSecret("old-secret");
deletePoller.waitForCompletion();
```

## Recover Deleted Secret

```java
SyncPoller<KeyVaultSecret, Void> recoverPoller = secretClient.beginRecoverDeletedSecret("old-secret");
recoverPoller.waitForCompletion();
```

## Backup and Restore

```java
byte[] backup = secretClient.backupSecret("important-secret");
KeyVaultSecret restored = secretClient.restoreSecretBackup(backup);
```

## Error Handling

```java
import com.azure.core.exception.HttpResponseException;
import com.azure.core.exception.ResourceNotFoundException;

try {
    KeyVaultSecret secret = secretClient.getSecret("my-secret");
} catch (ResourceNotFoundException e) {
    System.out.println("Secret not found");
} catch (HttpResponseException e) {
    int status = e.getResponse().getStatusCode();
    if (status == 403) {
        System.out.println("Access denied - check permissions");
    } else if (status == 429) {
        System.out.println("Rate limited - retry later");
    }
}
```

## Best Practices

1. **Enable Soft Delete** — Protects against accidental deletion
2. **Use Tags** — Tag secrets with environment, service, owner
3. **Set Expiration** — Use `setExpiresOn()` for credentials that should rotate
4. **Content Type** — Set `contentType` to indicate format
5. **Version Management** — Don't delete old versions immediately during rotation
6. **Least Privilege** — Use separate vaults for different environments

## Environment Variables

```bash
AZURE_KEYVAULT_URL=https://<vault-name>.vault.azure.net
```

## Trigger Phrases

- "Key Vault secrets Java", "secret management Java"
- "store password", "store API key", "connection string"
- "retrieve secret", "rotate secret"
- "Azure secrets", "vault secrets"
