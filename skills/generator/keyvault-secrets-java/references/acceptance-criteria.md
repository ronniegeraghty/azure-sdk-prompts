# Azure Key Vault Secrets SDK for Java Acceptance Criteria

**SDK**: `com.azure:azure-security-keyvault-secrets`
**Repository**: https://github.com/Azure/azure-sdk-for-java/tree/main/sdk/keyvault-v2/azure-security-keyvault-secrets

---

## 1. Correct Import Patterns

### 1.1 Client Imports

#### ✅ CORRECT: Secret Clients
```java
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;
import com.azure.security.keyvault.secrets.SecretAsyncClient;
```

#### ✅ CORRECT: Authentication
```java
import com.azure.identity.DefaultAzureCredentialBuilder;
```

### 1.2 Model Imports

#### ✅ CORRECT: Secret Models
```java
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;
import com.azure.security.keyvault.secrets.models.SecretProperties;
import com.azure.security.keyvault.secrets.models.DeletedSecret;
```

---

## 2. Client Creation Patterns

### 2.1 ✅ CORRECT: Builder with DefaultAzureCredential
```java
String vaultUrl = System.getenv("AZURE_KEYVAULT_URL");

SecretClient secretClient = new SecretClientBuilder()
    .vaultUrl(vaultUrl)
    .credential(new DefaultAzureCredentialBuilder().build())
    .buildClient();
```

### 2.2 ✅ CORRECT: Async Client
```java
SecretAsyncClient secretAsyncClient = new SecretClientBuilder()
    .vaultUrl(vaultUrl)
    .credential(new DefaultAzureCredentialBuilder().build())
    .buildAsyncClient();
```

---

## 3. Secret Operations

### 3.1 ✅ CORRECT: Set Secret
```java
KeyVaultSecret secret = secretClient.setSecret("database-password", "P@ssw0rd123!");
```

### 3.2 ✅ CORRECT: Get Secret
```java
KeyVaultSecret secret = secretClient.getSecret("database-password");
String value = secret.getValue();
```

### 3.3 ✅ CORRECT: List Secrets
```java
for (SecretProperties props : secretClient.listPropertiesOfSecrets()) {
    System.out.println("Secret: " + props.getName());
}
```

### 3.4 ✅ CORRECT: Delete Secret
```java
SyncPoller<DeletedSecret, Void> deletePoller = secretClient.beginDeleteSecret("old-secret");
deletePoller.waitForCompletion();
```

---

## 4. Error Handling

### 4.1 ✅ CORRECT: Exception Handling
```java
import com.azure.core.exception.HttpResponseException;
import com.azure.core.exception.ResourceNotFoundException;

try {
    KeyVaultSecret secret = secretClient.getSecret("my-secret");
} catch (ResourceNotFoundException e) {
    System.err.println("Secret not found: " + e.getMessage());
} catch (HttpResponseException e) {
    System.err.println("HTTP error: " + e.getResponse().getStatusCode());
}
```

---

## 5. Best Practices Checklist

- [ ] Use DefaultAzureCredential for authentication
- [ ] Use environment variables for vault URL
- [ ] Enable soft delete on vault
- [ ] Use tags to organize secrets
- [ ] Set expiration dates for credentials
- [ ] Set content type to indicate format
