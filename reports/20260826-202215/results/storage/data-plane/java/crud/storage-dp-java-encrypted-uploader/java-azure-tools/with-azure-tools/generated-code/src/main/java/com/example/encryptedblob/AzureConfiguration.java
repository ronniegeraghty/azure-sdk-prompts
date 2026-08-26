package com.example.encryptedblob;

import com.azure.core.credential.TokenCredential;
import com.azure.identity.ManagedIdentityCredential;
import com.azure.identity.ManagedIdentityCredentialBuilder;
import com.azure.security.keyvault.keys.KeyAsyncClient;
import com.azure.security.keyvault.keys.KeyClient;
import com.azure.security.keyvault.keys.KeyClientBuilder;
import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.BlobServiceClientBuilder;

import java.net.URI;
import java.net.URISyntaxException;
import java.util.Map;

public final class AzureConfiguration {
    private final ManagedIdentityCredential credential;
    private final BlobContainerClient blobContainerClient;
    private final BlobContainerAsyncClient blobContainerAsyncClient;
    private final KeyClient keyClient;
    private final KeyAsyncClient keyAsyncClient;
    private final String keyName;

    private AzureConfiguration(
        ManagedIdentityCredential credential,
        BlobContainerClient blobContainerClient,
        BlobContainerAsyncClient blobContainerAsyncClient,
        KeyClient keyClient,
        KeyAsyncClient keyAsyncClient,
        String keyName
    ) {
        this.credential = credential;
        this.blobContainerClient = blobContainerClient;
        this.blobContainerAsyncClient = blobContainerAsyncClient;
        this.keyClient = keyClient;
        this.keyAsyncClient = keyAsyncClient;
        this.keyName = keyName;
    }

    public static AzureConfiguration fromEnvironment() {
        Map<String, String> environment = System.getenv();
        String storageEndpoint = required(environment, "AZURE_STORAGE_BLOB_ENDPOINT");
        String containerName = required(environment, "AZURE_STORAGE_CONTAINER");
        String vaultUrl = required(environment, "AZURE_KEY_VAULT_URL");
        String keyName = required(environment, "AZURE_KEY_NAME");
        validateHttpsEndpoint("AZURE_STORAGE_BLOB_ENDPOINT", storageEndpoint);
        validateHttpsEndpoint("AZURE_KEY_VAULT_URL", vaultUrl);

        ManagedIdentityCredentialBuilder credentialBuilder =
            new ManagedIdentityCredentialBuilder();
        String managedIdentityClientId = environment.get("AZURE_CLIENT_ID");
        if (managedIdentityClientId != null && !managedIdentityClientId.isBlank()) {
            credentialBuilder.clientId(managedIdentityClientId);
        }
        ManagedIdentityCredential credential = credentialBuilder.build();

        BlobServiceClientBuilder blobBuilder = new BlobServiceClientBuilder()
            .endpoint(storageEndpoint)
            .credential(credential);
        BlobContainerClient syncContainer = blobBuilder
            .buildClient()
            .getBlobContainerClient(containerName);
        BlobContainerAsyncClient asyncContainer = blobBuilder
            .buildAsyncClient()
            .getBlobContainerAsyncClient(containerName);

        KeyClientBuilder keyBuilder = new KeyClientBuilder()
            .vaultUrl(vaultUrl)
            .credential(credential);
        return new AzureConfiguration(
            credential,
            syncContainer,
            asyncContainer,
            keyBuilder.buildClient(),
            keyBuilder.buildAsyncClient(),
            keyName
        );
    }

    public TokenCredential credential() {
        return credential;
    }

    public BlobContainerClient blobContainerClient() {
        return blobContainerClient;
    }

    public BlobContainerAsyncClient blobContainerAsyncClient() {
        return blobContainerAsyncClient;
    }

    public KeyClient keyClient() {
        return keyClient;
    }

    public KeyAsyncClient keyAsyncClient() {
        return keyAsyncClient;
    }

    public String keyName() {
        return keyName;
    }

    private static String required(Map<String, String> environment, String name) {
        String value = environment.get(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is missing: " + name);
        }
        return value;
    }

    private static void validateHttpsEndpoint(String name, String value) {
        try {
            URI uri = new URI(value);
            if (!"https".equalsIgnoreCase(uri.getScheme()) || uri.getHost() == null) {
                throw new IllegalStateException(name + " must be an absolute HTTPS URL");
            }
        } catch (URISyntaxException e) {
            throw new IllegalStateException(name + " is not a valid URL", e);
        }
    }
}
