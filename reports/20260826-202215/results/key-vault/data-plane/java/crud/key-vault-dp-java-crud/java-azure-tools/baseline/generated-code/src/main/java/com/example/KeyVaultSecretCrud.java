package com.example;

import com.azure.core.exception.HttpResponseException;
import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.util.polling.SyncPoller;
import com.azure.identity.CredentialUnavailableException;
import com.azure.identity.DefaultAzureCredential;
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
        try {
            String vaultUrl = requireEnvironmentVariable("AZURE_KEY_VAULT_URL");
            DefaultAzureCredential credential = new DefaultAzureCredentialBuilder().build();
            SecretClient secretClient = new SecretClientBuilder()
                    .vaultUrl(vaultUrl)
                    .credential(credential)
                    .buildClient();

            runCrudOperations(secretClient);
        } catch (CredentialUnavailableException exception) {
            System.err.println("No supported Azure credential was available: " + exception.getMessage());
            System.exit(1);
        } catch (ClientAuthenticationException exception) {
            System.err.println("Azure authentication failed: " + exception.getMessage());
            System.exit(1);
        } catch (HttpResponseException exception) {
            System.err.printf(
                    "Key Vault request failed with status %d: %s%n",
                    exception.getResponse().getStatusCode(),
                    exception.getMessage());
            System.exit(1);
        } catch (IllegalArgumentException exception) {
            System.err.println("Configuration error: " + exception.getMessage());
            System.exit(1);
        }
    }

    private static void runCrudOperations(SecretClient secretClient) {
        KeyVaultSecret createdSecret = secretClient.setSecret(SECRET_NAME, "my-secret-value");
        System.out.println("Created secret version: " + createdSecret.getProperties().getVersion());

        KeyVaultSecret readSecret = secretClient.getSecret(SECRET_NAME);
        System.out.println("Secret value: " + readSecret.getValue());

        KeyVaultSecret updatedSecret = secretClient.setSecret(SECRET_NAME, "updated-value");
        System.out.println("Updated secret value: " + updatedSecret.getValue());

        SyncPoller<DeletedSecret, Void> deletionPoller = secretClient.beginDeleteSecret(SECRET_NAME);
        deletionPoller.waitForCompletion();
        System.out.println("Deleted secret: " + SECRET_NAME);

        secretClient.purgeDeletedSecret(SECRET_NAME);
        System.out.println("Purged secret: " + SECRET_NAME);
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(
                    name + " must be set to a vault URL such as https://your-vault.vault.azure.net");
        }
        return value;
    }
}
