package com.example.azureencrypted;

import com.azure.core.exception.HttpResponseException;
import com.azure.core.credential.TokenCredential;
import com.azure.security.keyvault.keys.cryptography.CryptographyClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;
import com.azure.security.keyvault.keys.cryptography.models.KeyWrapAlgorithm;
import com.azure.security.keyvault.keys.cryptography.models.UnwrapResult;
import com.azure.security.keyvault.keys.cryptography.models.WrapResult;

import java.security.SecureRandom;
import java.util.Arrays;

public final class KeyManagement {
    static final int DATA_KEY_BYTES = 32;
    static final KeyWrapAlgorithm KEY_WRAP_ALGORITHM = KeyWrapAlgorithm.RSA_OAEP_256;

    private final CryptographyClient cryptographyClient;
    private final TokenCredential credential;
    private final String keyId;
    private final SecureRandom secureRandom;

    public KeyManagement(
            CryptographyClient cryptographyClient, TokenCredential credential, String keyId) {
        this(cryptographyClient, credential, keyId, new SecureRandom());
    }

    KeyManagement(
            CryptographyClient cryptographyClient,
            TokenCredential credential,
            String keyId,
            SecureRandom secureRandom) {
        this.cryptographyClient = cryptographyClient;
        this.credential = credential;
        this.keyId = keyId;
        this.secureRandom = secureRandom;
    }

    GeneratedDataKey generateAndWrapDataKey() {
        byte[] dataKey = new byte[DATA_KEY_BYTES];
        secureRandom.nextBytes(dataKey);
        try {
            WrapResult result = cryptographyClient.wrapKey(KEY_WRAP_ALGORITHM, dataKey);
            return new GeneratedDataKey(dataKey, result.getEncryptedKey(), keyId);
        } catch (HttpResponseException exception) {
            Arrays.fill(dataKey, (byte) 0);
            throw new EnvelopeEncryptionException(
                    "Key Vault could not wrap the data key with key " + keyId, exception);
        }
    }

    byte[] unwrapDataKey(byte[] wrappedDataKey, String wrappingKeyId) {
        try {
            CryptographyClient client = keyId.equals(wrappingKeyId)
                    ? cryptographyClient
                    : new CryptographyClientBuilder()
                            .keyIdentifier(wrappingKeyId)
                            .credential(credential)
                            .buildClient();
            UnwrapResult result = client.unwrapKey(KEY_WRAP_ALGORITHM, wrappedDataKey);
            return result.getKey();
        } catch (HttpResponseException exception) {
            throw new EnvelopeEncryptionException(
                    "Key Vault could not unwrap the data key with key " + wrappingKeyId, exception);
        }
    }

    static final class GeneratedDataKey implements AutoCloseable {
        private final byte[] plaintext;
        private final byte[] wrapped;
        private final String keyId;

        private GeneratedDataKey(byte[] plaintext, byte[] wrapped, String keyId) {
            this.plaintext = plaintext;
            this.wrapped = wrapped;
            this.keyId = keyId;
        }

        byte[] plaintext() {
            return plaintext;
        }

        byte[] wrapped() {
            return wrapped;
        }

        String keyId() {
            return keyId;
        }

        @Override
        public void close() {
            Arrays.fill(plaintext, (byte) 0);
        }
    }
}
