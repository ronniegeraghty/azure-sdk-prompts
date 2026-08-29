package com.example;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.exception.HttpResponseException;
import com.azure.core.exception.ResourceNotFoundException;
import com.azure.core.util.polling.SyncPoller;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;
import com.azure.security.keyvault.secrets.models.DeletedSecret;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;

public final class KeyVaultSecretCrud {
    private static final String VAULT_URL_ENV = "AZURE_KEY_VAULT_URL";
    private static final String SECRET_NAME = "my-secret";

    private KeyVaultSecretCrud() {
    }

    public static void main(String[] args) {
        try {
            SecretClient secretClient = createSecretClient();
            runCrudOperations(secretClient);
        } catch (ResourceNotFoundException exception) {
            System.err.printf("The secret or deleted secret was not found: %s%n", exception.getMessage());
            System.exit(1);
        } catch (ClientAuthenticationException exception) {
            System.err.printf("Azure authentication failed: %s%n", exception.getMessage());
            System.err.println("Sign in with Azure CLI or configure a supported DefaultAzureCredential identity.");
            System.exit(1);
        } catch (HttpResponseException exception) {
            int statusCode = exception.getResponse() == null
                ? -1
                : exception.getResponse().getStatusCode();
            System.err.printf("Key Vault request failed (HTTP %d): %s%n", statusCode, exception.getMessage());
            if (statusCode == 403) {
                System.err.println(
                    "Verify secret data-plane permissions and that purge protection does not block purging.");
            }
            System.exit(1);
        } catch (IllegalStateException exception) {
            System.err.println(exception.getMessage());
            System.exit(1);
        }
    }

    private static SecretClient createSecretClient() {
        String vaultUrl = System.getenv(VAULT_URL_ENV);
        if (vaultUrl == null || vaultUrl.isBlank()) {
            throw new IllegalStateException(
                "Set " + VAULT_URL_ENV + " to a vault URL such as https://my-vault.vault.azure.net.");
        }

        return new SecretClientBuilder()
            .vaultUrl(vaultUrl)
            .credential(new DefaultAzureCredentialBuilder().build())
            .buildClient();
    }

    private static void runCrudOperations(SecretClient secretClient) {
        KeyVaultSecret createdSecret = secretClient.setSecret(SECRET_NAME, "my-secret-value");
        System.out.printf("Created secret \"%s\".%n", createdSecret.getName());

        KeyVaultSecret readSecret = secretClient.getSecret(SECRET_NAME);
        System.out.printf("Read secret value: %s%n", readSecret.getValue());

        // A new value creates a new secret version; properties are updated through a separate API.
        KeyVaultSecret updatedSecret = secretClient.setSecret(SECRET_NAME, "updated-value");
        System.out.printf(
            "Updated secret \"%s\" by creating version \"%s\".%n",
            updatedSecret.getName(),
            updatedSecret.getProperties().getVersion());

        SyncPoller<DeletedSecret, Void> deletePoller = secretClient.beginDeleteSecret(SECRET_NAME);
        deletePoller.waitForCompletion();
        System.out.printf("Deleted secret \"%s\".%n", SECRET_NAME);

        secretClient.purgeDeletedSecret(SECRET_NAME);
        System.out.printf("Purged secret \"%s\" permanently.%n", SECRET_NAME);
    }
}
