package com.example;

import com.azure.core.exception.AzureException;
import com.azure.core.exception.HttpResponseException;
import com.azure.core.util.polling.SyncPoller;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;
import com.azure.security.keyvault.secrets.models.DeletedSecret;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;

public final class KeyVaultSecretCrud {
    private static final String SECRET_NAME = "my-secret";

    private KeyVaultSecretCrud() {
    }

    public static void main(String[] args) {
        String vaultUrl = System.getenv("KEY_VAULT_URL");
        if (vaultUrl == null || vaultUrl.isBlank()) {
            System.err.println(
                    "KEY_VAULT_URL must be set (for example, https://my-vault.vault.azure.net).");
            System.exit(2);
        }

        try {
            SecretClient secretClient = new SecretClientBuilder()
                    .vaultUrl(vaultUrl)
                    .credential(new DefaultAzureCredentialBuilder().build())
                    .buildClient();

            KeyVaultSecret createdSecret = secretClient.setSecret(SECRET_NAME, "my-secret-value");
            System.out.printf("Created secret version: %s%n", createdSecret.getProperties().getVersion());

            KeyVaultSecret readSecret = secretClient.getSecret(SECRET_NAME);
            System.out.printf("Secret value: %s%n", readSecret.getValue());

            KeyVaultSecret updatedSecret = secretClient.setSecret(SECRET_NAME, "updated-value");
            System.out.printf("Updated secret version: %s%n", updatedSecret.getProperties().getVersion());

            SyncPoller<DeletedSecret, Void> deletionPoller =
                    secretClient.beginDeleteSecret(SECRET_NAME);
            deletionPoller.waitForCompletion();
            System.out.printf("Deleted secret: %s%n", SECRET_NAME);

            secretClient.purgeDeletedSecret(SECRET_NAME);
            System.out.printf("Purged secret: %s%n", SECRET_NAME);
        } catch (HttpResponseException exception) {
            System.err.printf(
                    "Azure Key Vault request failed (HTTP %d): %s%n",
                    exception.getResponse().getStatusCode(),
                    exception.getMessage());
            System.exit(1);
        } catch (AzureException exception) {
            System.err.printf("Azure authentication or SDK operation failed: %s%n",
                    exception.getMessage());
            System.exit(1);
        } catch (IllegalArgumentException exception) {
            System.err.printf("Invalid Key Vault configuration: %s%n", exception.getMessage());
            System.exit(2);
        }
    }
}
