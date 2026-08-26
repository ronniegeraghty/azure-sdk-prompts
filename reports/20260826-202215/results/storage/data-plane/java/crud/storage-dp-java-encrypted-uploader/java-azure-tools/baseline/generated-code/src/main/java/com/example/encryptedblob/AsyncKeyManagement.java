package com.example.encryptedblob;

import com.azure.core.exception.HttpResponseException;
import com.azure.security.keyvault.keys.cryptography.CryptographyAsyncClient;
import com.azure.security.keyvault.keys.cryptography.models.KeyWrapAlgorithm;
import reactor.core.publisher.Mono;

import java.security.SecureRandom;
import java.util.function.Function;

public final class AsyncKeyManagement {
    private static final int DEK_SIZE_BYTES = 32;

    private final CryptographyAsyncClient cryptographyClient;
    private final String keyId;
    private final SecureRandom secureRandom;
    private final Function<String, CryptographyAsyncClient> cryptographyClientFactory;

    public AsyncKeyManagement(
            CryptographyAsyncClient cryptographyClient,
            String keyId,
            Function<String, CryptographyAsyncClient> cryptographyClientFactory) {
        this(cryptographyClient, keyId, cryptographyClientFactory, new SecureRandom());
    }

    AsyncKeyManagement(
            CryptographyAsyncClient cryptographyClient,
            String keyId,
            Function<String, CryptographyAsyncClient> cryptographyClientFactory,
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

    Mono<byte[]> wrap(DataEncryptionKey dataKey) {
        return cryptographyClient
                .wrapKey(KeyWrapAlgorithm.RSA_OAEP_256, dataKey.bytes())
                .map(result -> result.getEncryptedKey())
                .onErrorMap(
                        HttpResponseException.class,
                        e -> new KeyManagementException(
                                "Key Vault could not wrap the data key with " + keyId
                                        + "; verify that the key is enabled and permits wrapKey",
                                e));
    }

    Mono<DataEncryptionKey> unwrap(byte[] wrappedKey, String wrappingKeyId) {
        return cryptographyClientFactory.apply(wrappingKeyId)
                .unwrapKey(KeyWrapAlgorithm.RSA_OAEP_256, wrappedKey)
                .map(result -> new DataEncryptionKey(result.getKey()))
                .onErrorMap(
                        HttpResponseException.class,
                        e -> new KeyManagementException(
                                "Key Vault could not unwrap the data key with " + wrappingKeyId
                                        + "; verify that the key version is enabled and permits unwrapKey",
                                e));
    }

    public String keyId() {
        return keyId;
    }
}
