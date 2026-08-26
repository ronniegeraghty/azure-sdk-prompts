package com.example.encryptedblob;

import com.azure.core.credential.TokenCredential;
import com.azure.core.exception.HttpResponseException;
import com.azure.security.keyvault.keys.KeyClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;
import com.azure.security.keyvault.keys.cryptography.models.KeyWrapAlgorithm;
import com.azure.security.keyvault.keys.models.KeyVaultKey;

import java.security.SecureRandom;
import java.util.Arrays;
import java.util.Objects;

public final class SyncKeyManagementClient {
    static final KeyWrapAlgorithm WRAP_ALGORITHM = KeyWrapAlgorithm.RSA_OAEP_256;
    private static final int DATA_KEY_BYTES = 32;

    private final KeyClient keyClient;
    private final TokenCredential credential;
    private final String keyName;
    private final SecureRandom secureRandom;

    public SyncKeyManagementClient(KeyClient keyClient, TokenCredential credential, String keyName) {
        this(keyClient, credential, keyName, new SecureRandom());
    }

    SyncKeyManagementClient(
        KeyClient keyClient,
        TokenCredential credential,
        String keyName,
        SecureRandom secureRandom
    ) {
        this.keyClient = Objects.requireNonNull(keyClient, "keyClient");
        this.credential = Objects.requireNonNull(credential, "credential");
        this.keyName = Objects.requireNonNull(keyName, "keyName");
        this.secureRandom = Objects.requireNonNull(secureRandom, "secureRandom");
    }

    EnvelopeKey generateAndWrapDataKey() {
        byte[] dataKey = new byte[DATA_KEY_BYTES];
        secureRandom.nextBytes(dataKey);

        try {
            KeyVaultKey key = keyClient.getKey(keyName);
            String keyId = key.getId();
            byte[] wrappedKey = cryptographyClient(keyId)
                .wrapKey(WRAP_ALGORITHM, dataKey)
                .getEncryptedKey();
            ProtectedDataKey protectedKey =
                new ProtectedDataKey(keyId, WRAP_ALGORITHM.toString(), wrappedKey);
            return new EnvelopeKey(protectedKey, dataKey);
        } catch (HttpResponseException e) {
            Arrays.fill(dataKey, (byte) 0);
            throw keyVaultFailure("wrap a new data encryption key", e);
        } catch (RuntimeException e) {
            Arrays.fill(dataKey, (byte) 0);
            throw e;
        }
    }

    byte[] unwrapDataKey(ProtectedDataKey protectedKey) {
        Objects.requireNonNull(protectedKey, "protectedKey");
        validateAlgorithm(protectedKey.algorithm());

        try {
            return cryptographyClient(protectedKey.keyId())
                .unwrapKey(WRAP_ALGORITHM, protectedKey.wrappedKey())
                .getKey();
        } catch (HttpResponseException e) {
            throw keyVaultFailure("unwrap the data encryption key with " + protectedKey.keyId(), e);
        }
    }

    private CryptographyClient cryptographyClient(String keyId) {
        return new CryptographyClientBuilder()
            .keyIdentifier(keyId)
            .credential(credential)
            .buildClient();
    }

    static void validateAlgorithm(String algorithm) {
        if (!WRAP_ALGORITHM.toString().equals(algorithm)) {
            throw new IllegalArgumentException("Unsupported key wrap algorithm: " + algorithm);
        }
    }

    private static KeyManagementException keyVaultFailure(String operation, HttpResponseException e) {
        int status = e.getResponse() == null ? -1 : e.getResponse().getStatusCode();
        return new KeyManagementException(
            "Azure Key Vault could not " + operation + " (HTTP " + status
                + "). Verify that the key exists, is enabled, and permits wrapKey/unwrapKey.",
            e
        );
    }
}
