package com.example.encryptedblob;

import com.azure.core.credential.TokenCredential;
import com.azure.core.exception.AzureException;
import com.azure.core.exception.HttpResponseException;
import com.azure.security.keyvault.keys.KeyClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;
import com.azure.security.keyvault.keys.cryptography.models.KeyWrapAlgorithm;

import java.util.Arrays;
import java.util.Objects;

public final class KeyManagementService {
    private static final KeyWrapAlgorithm WRAP_ALGORITHM = KeyWrapAlgorithm.RSA_OAEP_256;
    private final TokenCredential credential;
    private final String currentKeyId;
    private final CryptographyClient currentKeyClient;

    public KeyManagementService(
            KeyClient keyClient, TokenCredential credential, String keyName) {
        this.credential = Objects.requireNonNull(credential, "credential");
        try {
            this.currentKeyId = keyClient.getKey(keyName).getId();
            this.currentKeyClient = cryptographyClient(currentKeyId);
        } catch (AzureException e) {
            throw keyVaultFailure("resolve wrapping key '" + keyName + "'", e);
        }
    }

    public WrappedDataKey generateAndWrapDataKey() {
        byte[] plaintextKey = CipherSupport.generateDataKey();
        try {
            byte[] wrappedKey = currentKeyClient.wrapKey(WRAP_ALGORITHM, plaintextKey)
                    .getEncryptedKey();
            return new WrappedDataKey(plaintextKey, wrappedKey, currentKeyId);
        } catch (AzureException e) {
            throw keyVaultFailure("wrap the data encryption key", e);
        } finally {
            Arrays.fill(plaintextKey, (byte) 0);
        }
    }

    public byte[] unwrapDataKey(String keyId, byte[] wrappedKey) {
        try {
            return clientFor(keyId).unwrapKey(WRAP_ALGORITHM, wrappedKey).getKey();
        } catch (AzureException e) {
            throw keyVaultFailure(
                    "unwrap the data encryption key with Key Vault key " + keyId, e);
        }
    }

    private CryptographyClient clientFor(String keyId) {
        return currentKeyId.equals(keyId) ? currentKeyClient : cryptographyClient(keyId);
    }

    private CryptographyClient cryptographyClient(String keyId) {
        return new CryptographyClientBuilder()
                .keyIdentifier(keyId)
                .credential(credential)
                .buildClient();
    }

    private static EnvelopeEncryptionException keyVaultFailure(
            String operation, AzureException cause) {
        String status = cause instanceof HttpResponseException httpError
                ? " (HTTP " + httpError.getResponse().getStatusCode() + ")"
                : "";
        return new EnvelopeEncryptionException(
                "Azure Key Vault could not " + operation + status,
                cause);
    }
}
