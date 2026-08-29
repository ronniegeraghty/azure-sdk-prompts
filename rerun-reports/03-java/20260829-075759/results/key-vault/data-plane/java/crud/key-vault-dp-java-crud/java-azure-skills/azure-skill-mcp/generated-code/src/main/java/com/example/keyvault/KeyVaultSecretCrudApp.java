package com.example.keyvault;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.exception.HttpResponseException;
import com.azure.core.util.polling.SyncPoller;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.security.keyvault.secrets.SecretClient;
import com.azure.security.keyvault.secrets.SecretClientBuilder;
import com.azure.security.keyvault.secrets.models.DeletedSecret;
import com.azure.security.keyvault.secrets.models.KeyVaultSecret;

public final class KeyVaultSecretCrudApp {
    private static final String VAULT_URL_ENVIRONMENT_VARIABLE = "AZURE_KEY_VAULT_URL";
    private static final String SECRET_NAME = "my-secret";
    private static final String INITIAL_VALUE = "my-secret-value";
    private static final String UPDATED_VALUE = "updated-value";

    private KeyVaultSecretCrudApp() {
    }

    public static void main(String[] args) {
        String operation = "configuration";

        try {
            String vaultUrl = requireEnvironmentVariable(VAULT_URL_ENVIRONMENT_VARIABLE);

            operation = "client creation";
            SecretClient secretClient = new SecretClientBuilder()
                .vaultUrl(vaultUrl)
                .credential(new DefaultAzureCredentialBuilder().build())
                .buildClient();

            operation = "create";
            secretClient.setSecret(SECRET_NAME, INITIAL_VALUE);
            System.out.printf("Created secret \"%s\".%n", SECRET_NAME);

            operation = "read";
            KeyVaultSecret secret = secretClient.getSecret(SECRET_NAME);
            System.out.printf("Read secret value: %s%n", secret.getValue());

            operation = "update";
            secretClient.setSecret(SECRET_NAME, UPDATED_VALUE);
            System.out.printf("Updated secret \"%s\".%n", SECRET_NAME);

            operation = "delete";
            SyncPoller<DeletedSecret, Void> deletionPoller =
                secretClient.beginDeleteSecret(SECRET_NAME);
            deletionPoller.waitForCompletion();
            System.out.printf("Deleted secret \"%s\".%n", SECRET_NAME);

            operation = "purge";
            secretClient.purgeDeletedSecret(SECRET_NAME);
            System.out.printf("Purged secret \"%s\".%n", SECRET_NAME);
        } catch (ClientAuthenticationException exception) {
            System.err.printf(
                "Authentication failed during %s: %s%n",
                operation,
                exception.getMessage());
            System.exit(1);
        } catch (HttpResponseException exception) {
            int statusCode = exception.getResponse() == null
                ? -1
                : exception.getResponse().getStatusCode();
            System.err.printf(
                "Azure Key Vault operation \"%s\" failed (HTTP %d): %s%n",
                operation,
                statusCode,
                exception.getMessage());
            System.exit(1);
        } catch (IllegalArgumentException exception) {
            System.err.printf("Invalid configuration: %s%n", exception.getMessage());
            System.exit(1);
        }
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(
                "Set " + name + " to a vault URL such as https://<vault-name>.vault.azure.net");
        }
        return value;
    }
}
