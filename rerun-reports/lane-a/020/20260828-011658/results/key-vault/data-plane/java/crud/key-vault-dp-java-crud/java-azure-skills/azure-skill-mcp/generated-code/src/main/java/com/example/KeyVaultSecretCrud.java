package com.example;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.exception.HttpResponseException;
import com.azure.core.exception.ResourceNotFoundException;
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
    private static final String INITIAL_VALUE = "my-secret-value";
    private static final String UPDATED_VALUE = "updated-value";

    private KeyVaultSecretCrud() {
    }

    public static void main(String[] args) {
        try {
            String vaultUrl = requireEnvironmentVariable("AZURE_KEY_VAULT_URL");
            SecretClient secretClient = createSecretClient(vaultUrl);

            KeyVaultSecret createdSecret = secretClient.setSecret(SECRET_NAME, INITIAL_VALUE);
            System.out.printf("Created secret '%s'.%n", createdSecret.getName());

            KeyVaultSecret retrievedSecret = secretClient.getSecret(SECRET_NAME);
            System.out.printf("Read secret value: %s%n", retrievedSecret.getValue());

            KeyVaultSecret updatedSecret = secretClient.setSecret(SECRET_NAME, UPDATED_VALUE);
            System.out.printf("Updated secret '%s' to value: %s%n",
                updatedSecret.getName(), updatedSecret.getValue());

            SyncPoller<DeletedSecret, Void> deletePoller =
                secretClient.beginDeleteSecret(SECRET_NAME);
            deletePoller.waitForCompletion();
            System.out.printf("Deleted secret '%s'.%n", SECRET_NAME);

            secretClient.purgeDeletedSecret(SECRET_NAME);
            System.out.printf("Purged secret '%s'.%n", SECRET_NAME);
        } catch (IllegalArgumentException exception) {
            System.err.println("Configuration error: " + exception.getMessage());
            System.exit(2);
        } catch (CredentialUnavailableException exception) {
            System.err.println("No DefaultAzureCredential source is available: "
                + exception.getMessage());
            System.exit(3);
        } catch (ClientAuthenticationException exception) {
            System.err.println("Authentication failed. Configure a supported "
                + "DefaultAzureCredential source: " + exception.getMessage());
            System.exit(3);
        } catch (ResourceNotFoundException exception) {
            System.err.println("The vault or secret was not found: " + exception.getMessage());
            System.exit(4);
        } catch (HttpResponseException exception) {
            System.err.printf("Key Vault request failed (HTTP %d): %s%n",
                exception.getResponse().getStatusCode(), exception.getMessage());
            System.exit(5);
        } catch (RuntimeException exception) {
            System.err.println("Unexpected failure while managing the secret: "
                + exception.getMessage());
            System.exit(1);
        }
    }

    private static SecretClient createSecretClient(String vaultUrl) {
        DefaultAzureCredential credential = new DefaultAzureCredentialBuilder().build();

        return new SecretClientBuilder()
            .vaultUrl(vaultUrl)
            .credential(credential)
            .buildClient();
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(name + " must be set.");
        }
        return value;
    }
}
