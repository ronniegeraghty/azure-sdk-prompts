package com.example;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.exception.HttpResponseException;
import com.azure.core.exception.ResourceNotFoundException;
import com.azure.core.util.polling.SyncPoller;
import com.azure.identity.CredentialUnavailableException;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;
import com.azure.security.keyvault.secrets.models.DeletedSecret;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;

public final class KeyVaultSecretCrudApp {
    private static final String VAULT_URL_ENVIRONMENT_VARIABLE = "AZURE_KEY_VAULT_URL";
    private static final String SECRET_NAME = "my-secret";

    private KeyVaultSecretCrudApp() {
    }

    public static void main(String[] args) {
        String vaultUrl = System.getenv(VAULT_URL_ENVIRONMENT_VARIABLE);
        if (vaultUrl == null || vaultUrl.isBlank()) {
            System.err.printf(
                "Set %s to a vault URL such as https://<vault-name>.vault.azure.net/.%n",
                VAULT_URL_ENVIRONMENT_VARIABLE);
            System.exit(2);
        }

        SecretClient secretClient = new SecretClientBuilder()
            .vaultUrl(vaultUrl)
            .credential(new DefaultAzureCredentialBuilder().build())
            .buildClient();

        try {
            runCrudOperations(secretClient);
        } catch (CredentialUnavailableException exception) {
            System.err.println("No credential was available to DefaultAzureCredential: "
                + exception.getMessage());
            System.exit(1);
        } catch (ClientAuthenticationException exception) {
            System.err.println("Azure authentication failed: " + exception.getMessage());
            System.exit(1);
        } catch (ResourceNotFoundException exception) {
            System.err.printf("Secret \"%s\" was not found: %s%n",
                SECRET_NAME, exception.getMessage());
            System.exit(1);
        } catch (HttpResponseException exception) {
            handleHttpError(exception);
            System.exit(1);
        }
    }

    private static void runCrudOperations(SecretClient secretClient) {
        KeyVaultSecret createdSecret = secretClient.setSecret(SECRET_NAME, "my-secret-value");
        System.out.printf("Created secret \"%s\" (version %s).%n",
            createdSecret.getName(), createdSecret.getProperties().getVersion());

        KeyVaultSecret readSecret = secretClient.getSecret(SECRET_NAME);
        System.out.printf("Read secret value: %s%n", readSecret.getValue());

        KeyVaultSecret updatedSecret = secretClient.setSecret(SECRET_NAME, "updated-value");
        System.out.printf("Updated secret \"%s\" to a new version (%s).%n",
            updatedSecret.getName(), updatedSecret.getProperties().getVersion());

        SyncPoller<DeletedSecret, Void> deletePoller =
            secretClient.beginDeleteSecret(SECRET_NAME);
        deletePoller.waitForCompletion();
        System.out.printf("Deleted secret \"%s\".%n", SECRET_NAME);

        secretClient.purgeDeletedSecret(SECRET_NAME);
        System.out.printf("Purged secret \"%s\" permanently.%n", SECRET_NAME);
    }

    private static void handleHttpError(HttpResponseException exception) {
        int statusCode = exception.getResponse() == null
            ? -1
            : exception.getResponse().getStatusCode();

        switch (statusCode) {
            case 403 ->
                System.err.println("Access denied. Grant the identity permission to set, get, "
                    + "delete, and purge secrets.");
            case 409 ->
                System.err.println("The operation conflicted with the vault state. A previously "
                    + "deleted secret may still exist, or purge protection may be enabled.");
            case 429 ->
                System.err.println("Key Vault throttled the request after SDK retries. Try again later.");
            default ->
                System.err.printf("Key Vault request failed with HTTP status %d: %s%n",
                    statusCode, exception.getMessage());
        }
    }
}
