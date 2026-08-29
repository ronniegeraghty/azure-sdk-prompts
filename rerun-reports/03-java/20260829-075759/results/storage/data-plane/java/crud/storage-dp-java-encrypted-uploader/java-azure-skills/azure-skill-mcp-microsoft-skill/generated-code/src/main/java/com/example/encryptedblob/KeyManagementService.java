package com.example.encryptedblob;

import com.azure.core.credential.TokenCredential;
import com.azure.core.exception.HttpResponseException;
import com.azure.security.keyvault.keys.KeyClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;
import com.azure.security.keyvault.keys.cryptography.models.KeyWrapAlgorithm;
import com.azure.security.keyvault.keys.models.KeyVaultKey;

import java.security.SecureRandom;

public final class KeyManagementService {
    static final KeyWrapAlgorithm WRAP_ALGORITHM = KeyWrapAlgorithm.RSA_OAEP_256;
    static final String WRAP_ALGORITHM_NAME = "RSA-OAEP-256";
    private static final int DATA_KEY_BYTES = 32;

    private final KeyClient keyClient;
    private final TokenCredential credential;
    private final String keyName;
    private final SecureRandom secureRandom;

    public KeyManagementService(KeyClient keyClient, TokenCredential credential, String keyName) {
        this(keyClient, credential, keyName, new SecureRandom());
    }

    KeyManagementService(
            KeyClient keyClient,
            TokenCredential credential,
            String keyName,
            SecureRandom secureRandom) {
        this.keyClient = keyClient;
        this.credential = credential;
        this.keyName = keyName;
        this.secureRandom = secureRandom;
    }

    ProtectedDataKey generateAndWrapDataKey() {
        try {
            KeyVaultKey key = keyClient.getKey(keyName);
            DataKey dataKey = generateDataKey();
            try {
                byte[] wrappedKey = cryptographyClient(key.getId())
                        .wrapKey(WRAP_ALGORITHM, dataKey.bytes())
                        .getEncryptedKey();
                return new ProtectedDataKey(dataKey, wrappedKey, key.getId());
            } catch (RuntimeException exception) {
                dataKey.close();
                throw exception;
            }
        } catch (HttpResponseException exception) {
            throw new KeyManagementException(
                    "Key Vault could not wrap the data key with key '" + keyName
                            + "' (HTTP " + exception.getResponse().getStatusCode() + ")",
                    exception);
        }
    }

    DataKey unwrapDataKey(String keyId, byte[] wrappedKey) {
        try {
            byte[] rawKey = cryptographyClient(keyId)
                    .unwrapKey(WRAP_ALGORITHM, wrappedKey)
                    .getKey();
            return new DataKey(rawKey);
        } catch (HttpResponseException exception) {
            throw new KeyManagementException(
                    "Key Vault could not unwrap the data key with key '" + keyId
                            + "' (HTTP " + exception.getResponse().getStatusCode() + ")",
                    exception);
        }
    }

    private DataKey generateDataKey() {
        byte[] key = new byte[DATA_KEY_BYTES];
        secureRandom.nextBytes(key);
        return new DataKey(key);
    }

    private CryptographyClient cryptographyClient(String keyId) {
        return new CryptographyClientBuilder()
                .keyIdentifier(keyId)
                .credential(credential)
                .buildClient();
    }
}
