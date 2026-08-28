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

public final class KeyManagement {
    private static final int DATA_KEY_BYTES = 32;
    private static final KeyWrapAlgorithm WRAP_ALGORITHM = KeyWrapAlgorithm.RSA_OAEP_256;

    private final KeyClient keyClient;
    private final TokenCredential credential;
    private final String keyName;
    private final SecureRandom secureRandom;

    public KeyManagement(KeyClient keyClient, TokenCredential credential, String keyName) {
        this.keyClient = keyClient;
        this.credential = credential;
        this.keyName = keyName;
        this.secureRandom = new SecureRandom();
    }

    DataKeyEnvelope generateAndWrapKey() {
        byte[] dataKey = new byte[DATA_KEY_BYTES];
        secureRandom.nextBytes(dataKey);

        try {
            KeyVaultKey vaultKey = keyClient.getKey(keyName);
            String versionedKeyId = vaultKey.getId();
            CryptographyClient cryptographyClient = cryptoClient(versionedKeyId);
            byte[] wrappedKey = cryptographyClient.wrapKey(WRAP_ALGORITHM, dataKey).getEncryptedKey();
            return new DataKeyEnvelope(dataKey, new ProtectedDataKey(versionedKeyId, wrappedKey));
        } catch (HttpResponseException exception) {
            Arrays.fill(dataKey, (byte) 0);
            throw keyVaultFailure("wrap a data encryption key", exception);
        } catch (RuntimeException exception) {
            Arrays.fill(dataKey, (byte) 0);
            throw exception;
        }
    }

    byte[] unwrapKey(ProtectedDataKey protectedKey) {
        try {
            return cryptoClient(protectedKey.keyId())
                .unwrapKey(WRAP_ALGORITHM, protectedKey.wrappedKey())
                .getKey();
        } catch (HttpResponseException exception) {
            throw keyVaultFailure("unwrap the data encryption key", exception);
        }
    }

    private CryptographyClient cryptoClient(String keyId) {
        return new CryptographyClientBuilder()
            .keyIdentifier(keyId)
            .credential(credential)
            .buildClient();
    }

    private static EnvelopeEncryptionException keyVaultFailure(String operation, HttpResponseException exception) {
        int status = exception.getResponse() == null ? -1 : exception.getResponse().getStatusCode();
        return new EnvelopeEncryptionException(
            "Key Vault could not " + operation + " (HTTP " + status
                + "). Verify that the stored key version is enabled and permits wrapKey/unwrapKey.",
            exception);
    }
}
