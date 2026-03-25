# Azure Key Vault Secrets SDK for Java — Examples

Comprehensive code examples for the Azure Key Vault Secrets SDK for Java.

## Maven Dependency

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-security-keyvault-secrets</artifactId>
    <version>4.9.0</version>
</dependency>

<!-- Required for authentication -->
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-identity</artifactId>
    <version>1.18.2</version>
</dependency>
```

## Client Creation

### Sync SecretClient

```java
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;

SecretClient secretClient = new SecretClientBuilder()
    .credential(new DefaultAzureCredentialBuilder().build())
    .vaultUrl("<your-key-vault-url>")
    .buildClient();
```

### Async SecretClient

```java
import com.azure.security.keyvault.secrets.SecretAsyncClient;

SecretAsyncClient secretAsyncClient = new SecretClientBuilder()
    .credential(new DefaultAzureCredentialBuilder().build())
    .vaultUrl("<your-key-vault-url>")
    .buildAsyncClient();
```

## Setting Secrets

### Simple Secret

```java
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;

KeyVaultSecret secret = secretClient.setSecret("<secret-name>", "<secret-value>");
System.out.printf("Secret created with name \"%s\" and value \"%s\"%n",
    secret.getName(), secret.getValue());
```

### Secret with Properties (Expiration)

```java
import com.azure.security.keyvault.secrets.models.SecretProperties;
import java.time.OffsetDateTime;

KeyVaultSecret newSecret = new KeyVaultSecret("secretName", "secretValue")
    .setProperties(new SecretProperties().setExpiresOn(OffsetDateTime.now().plusDays(60)));

KeyVaultSecret returnedSecret = secretClient.setSecret(newSecret);
```

## Getting Secrets

### Get Current Version

```java
KeyVaultSecret secret = secretClient.getSecret("secretName");
System.out.printf("Secret returned with name %s and value %s%n",
    secret.getName(), secret.getValue());
```

### Get Specific Version

```java
String secretVersion = "6A385B124DEF4096AF1361A85B16C204";
KeyVaultSecret secretWithVersion = secretClient.getSecret("secretName", secretVersion);
```

## Listing Secrets

```java
import com.azure.security.keyvault.secrets.models.SecretProperties;

for (SecretProperties secretProps : secretClient.listPropertiesOfSecrets()) {
    KeyVaultSecret secretWithValue = secretClient.getSecret(
        secretProps.getName(),
        secretProps.getVersion()
    );
    System.out.printf("Secret: %s = %s%n",
        secretWithValue.getName(),
        secretWithValue.getValue());
}
```

## Deleting and Recovering Secrets

### Delete Secret

```java
import com.azure.core.util.polling.SyncPoller;
import com.azure.security.keyvault.secrets.models.DeletedSecret;

SyncPoller<DeletedSecret, Void> deleteSecretPoller = secretClient.beginDeleteSecret("secretName");
deleteSecretPoller.waitForCompletion();
```

### Recover Deleted Secret

```java
SyncPoller<KeyVaultSecret, Void> recoverPoller =
    secretClient.beginRecoverDeletedSecret("deletedSecretName");
recoverPoller.waitForCompletion();
```

## Async Client Patterns

```java
// Set secret async
secretAsyncClient.setSecret("asyncSecretName", "asyncSecretValue")
    .subscribe(
        secret -> System.out.printf("Created secret: %s%n", secret.getName()),
        error -> System.err.println("Error: " + error.getMessage())
    );

// List secrets async
secretAsyncClient.listPropertiesOfSecrets()
    .flatMap(secretProps -> secretAsyncClient.getSecret(
        secretProps.getName(),
        secretProps.getVersion()
    ))
    .subscribe(
        secret -> System.out.printf("Secret: %s%n", secret.getName()),
        error -> System.err.println("Error: " + error.getMessage())
    );
```

## Error Handling

```java
import com.azure.core.exception.HttpResponseException;
import com.azure.core.exception.ResourceNotFoundException;

try {
    KeyVaultSecret secret = secretClient.getSecret("nonexistent-secret");
} catch (ResourceNotFoundException e) {
    System.err.println("Secret not found: " + e.getMessage());
} catch (HttpResponseException e) {
    System.err.println("HTTP error: " + e.getResponse().getStatusCode());
}
```
