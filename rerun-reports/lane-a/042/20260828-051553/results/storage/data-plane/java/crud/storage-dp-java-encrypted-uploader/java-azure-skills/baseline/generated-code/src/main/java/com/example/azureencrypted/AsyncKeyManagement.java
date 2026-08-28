package com.example.azureencrypted;

import com.azure.core.exception.HttpResponseException;
import com.azure.core.credential.TokenCredential;
import com.azure.security.keyvault.keys.cryptography.CryptographyAsyncClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;
import reactor.core.publisher.Mono;

import java.security.SecureRandom;
import java.util.Arrays;
import java.util.concurrent.atomic.AtomicBoolean;

public final class AsyncKeyManagement {
    private final CryptographyAsyncClient cryptographyClient;
    private final TokenCredential credential;
    private final String keyId;
    private final SecureRandom secureRandom;

    public AsyncKeyManagement(
            CryptographyAsyncClient cryptographyClient,
            TokenCredential credential,
            String keyId) {
        this(cryptographyClient, credential, keyId, new SecureRandom());
    }

    AsyncKeyManagement(
            CryptographyAsyncClient cryptographyClient,
            TokenCredential credential,
            String keyId,
            SecureRandom secureRandom) {
        this.cryptographyClient = cryptographyClient;
        this.credential = credential;
        this.keyId = keyId;
        this.secureRandom = secureRandom;
    }

    Mono<GeneratedDataKey> generateAndWrapDataKey() {
        return Mono.defer(() -> {
            byte[] dataKey = new byte[KeyManagement.DATA_KEY_BYTES];
            secureRandom.nextBytes(dataKey);
            AtomicBoolean ownershipTransferred = new AtomicBoolean();

            return cryptographyClient.wrapKey(KeyManagement.KEY_WRAP_ALGORITHM, dataKey)
                    .map(result -> {
                        ownershipTransferred.set(true);
                        return new GeneratedDataKey(dataKey, result.getEncryptedKey(), keyId);
                    })
                    .onErrorMap(
                            HttpResponseException.class,
                            exception -> new EnvelopeEncryptionException(
                                    "Key Vault could not wrap the data key with key " + keyId,
                                    exception))
                    .doFinally(signal -> {
                        if (!ownershipTransferred.get()) {
                            Arrays.fill(dataKey, (byte) 0);
                        }
                    });
        });
    }

    Mono<byte[]> unwrapDataKey(byte[] wrappedDataKey, String wrappingKeyId) {
        CryptographyAsyncClient client = keyId.equals(wrappingKeyId)
                ? cryptographyClient
                : new CryptographyClientBuilder()
                        .keyIdentifier(wrappingKeyId)
                        .credential(credential)
                        .buildAsyncClient();
        return client.unwrapKey(KeyManagement.KEY_WRAP_ALGORITHM, wrappedDataKey)
                .map(result -> result.getKey())
                .onErrorMap(
                        HttpResponseException.class,
                        exception -> new EnvelopeEncryptionException(
                                "Key Vault could not unwrap the data key with key " + wrappingKeyId,
                                exception));
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
