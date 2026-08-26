package com.example.encryptedblob;

import com.azure.core.exception.HttpResponseException;
import com.azure.security.keyvault.keys.cryptography.CryptographyClient;
import com.azure.security.keyvault.keys.cryptography.models.KeyWrapAlgorithm;
import com.azure.security.keyvault.keys.cryptography.models.UnwrapResult;
import com.azure.security.keyvault.keys.cryptography.models.WrapResult;

import java.security.SecureRandom;
import java.util.function.Function;

public final class KeyManagement {
    private static final int DEK_SIZE_BYTES = 32;

    private final CryptographyClient cryptographyClient;
    private final String keyId;
    private final SecureRandom secureRandom;
    private final Function<String, CryptographyClient> cryptographyClientFactory;

    public KeyManagement(
            CryptographyClient cryptographyClient,
            String keyId,
            Function<String, CryptographyClient> cryptographyClientFactory) {
        this(cryptographyClient, keyId, cryptographyClientFactory, new SecureRandom());
    }

    KeyManagement(
            CryptographyClient cryptographyClient,
            String keyId,
            Function<String, CryptographyClient> cryptographyClientFactory,
            SecureRandom secureRandom) {
        this.cryptographyClient = cryptographyClient;
        this.keyId = keyId;
        this.cryptographyClientFactory = cryptographyClientFactory;
        this.secureRandom = secureRandom;
    }

    DataEncryptionKey generateDataKey() {
        byte[] key = new byte[DEK_SIZE_BYTES];
        secureRandom.nextBytes(key);
        return new DataEncryptionKey(key);
    }

    byte[] wrap(DataEncryptionKey dataKey) {
        try {
            WrapResult result = cryptographyClient.wrapKey(KeyWrapAlgorithm.RSA_OAEP_256, dataKey.bytes());
            return result.getEncryptedKey();
        } catch (HttpResponseException e) {
            throw new KeyManagementException(
                    "Key Vault could not wrap the data key with " + keyId
                            + "; verify that the key is enabled and permits wrapKey",
                    e);
        }
    }

    DataEncryptionKey unwrap(byte[] wrappedKey, String wrappingKeyId) {
        try {
            UnwrapResult result = cryptographyClientFactory.apply(wrappingKeyId)
                    .unwrapKey(KeyWrapAlgorithm.RSA_OAEP_256, wrappedKey);
            return new DataEncryptionKey(result.getKey());
        } catch (HttpResponseException e) {
            throw new KeyManagementException(
                    "Key Vault could not unwrap the data key with " + wrappingKeyId
                            + "; verify that the key version is enabled and permits unwrapKey",
                    e);
        }
    }

    public String keyId() {
        return keyId;
    }
}
