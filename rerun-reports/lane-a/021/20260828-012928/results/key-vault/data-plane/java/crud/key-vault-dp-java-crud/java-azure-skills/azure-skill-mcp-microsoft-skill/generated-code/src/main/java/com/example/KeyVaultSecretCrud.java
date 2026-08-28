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
    private static final String SECRET_NAME = "my-secret";
    private static final String INITIAL_VALUE = "my-secret-value";
    private static final String UPDATED_VALUE = "updated-value";

    private KeyVaultSecretCrud() {
    }

    public static void main(String[] args) {
        String vaultUrl = System.getenv("AZURE_KEYVAULT_URL");
        if (vaultUrl == null || vaultUrl.isBlank()) {
            System.err.println(
                "AZURE_KEYVAULT_URL must be set, for example: https://<vault-name>.vault.azure.net");
            System.exit(2);
        }

        try {
            SecretClient secretClient = new SecretClientBuilder()
                .vaultUrl(vaultUrl)
                .credential(new DefaultAzureCredentialBuilder().build())
                .buildClient();

            KeyVaultSecret created = secretClient.setSecret(SECRET_NAME, INITIAL_VALUE);
            System.out.printf("Created secret \"%s\" (version %s).%n",
                created.getName(), created.getProperties().getVersion());

            KeyVaultSecret read = secretClient.getSecret(SECRET_NAME);
            System.out.printf("Read secret value: %s%n", read.getValue());

            // Secret values are immutable; setting the same name creates a new version.
            KeyVaultSecret updated = secretClient.setSecret(SECRET_NAME, UPDATED_VALUE);
            System.out.printf("Updated secret \"%s\" to version %s.%n",
                updated.getName(), updated.getProperties().getVersion());

            SyncPoller<DeletedSecret, Void> deletePoller =
                secretClient.beginDeleteSecret(SECRET_NAME);
            deletePoller.waitForCompletion();
            System.out.printf("Deleted secret \"%s\".%n", SECRET_NAME);

            secretClient.purgeDeletedSecret(SECRET_NAME);
            System.out.printf("Purged secret \"%s\" permanently.%n", SECRET_NAME);
        } catch (ClientAuthenticationException e) {
            System.err.println("Authentication failed. Check the credentials available to "
                + "DefaultAzureCredential: " + e.getMessage());
            System.exit(1);
        } catch (ResourceNotFoundException e) {
            System.err.println("The requested vault or secret was not found: " + e.getMessage());
            System.exit(1);
        } catch (HttpResponseException e) {
            int statusCode = e.getResponse() == null ? -1 : e.getResponse().getStatusCode();
            System.err.printf("Azure Key Vault request failed (HTTP %d): %s%n",
                statusCode, e.getMessage());
            if (statusCode == 403) {
                System.err.println("Ensure the identity has secret get, set, delete, and purge permissions. "
                    + "Purge also fails when purge protection is enabled.");
            } else if (statusCode == 429) {
                System.err.println("The vault is throttling requests; retry after the server's delay.");
            }
            System.exit(1);
        } catch (IllegalArgumentException e) {
            System.err.println("Invalid Key Vault configuration: " + e.getMessage());
            System.exit(2);
        }
    }
}
