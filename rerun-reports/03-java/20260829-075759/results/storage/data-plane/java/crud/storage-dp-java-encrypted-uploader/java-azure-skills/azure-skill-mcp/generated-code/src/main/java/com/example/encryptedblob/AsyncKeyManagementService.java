package com.example.encryptedblob;

import com.azure.security.keyvault.keys.cryptography.CryptographyAsyncClient;
import com.azure.security.keyvault.keys.cryptography.models.KeyWrapAlgorithm;
import reactor.core.publisher.Mono;

import java.security.SecureRandom;
import java.util.Arrays;
import java.util.Objects;

public final class AsyncKeyManagementService {
    private static final KeyWrapAlgorithm SDK_WRAP_ALGORITHM = KeyWrapAlgorithm.RSA_OAEP_256;
    private static final int DATA_KEY_BYTES = 32;

    private final CryptographyAsyncClient cryptographyClient;
    private final String keyId;
    private final SecureRandom secureRandom;

    public AsyncKeyManagementService(CryptographyAsyncClient cryptographyClient, String keyId) {
        this(cryptographyClient, keyId, new SecureRandom());
    }

    AsyncKeyManagementService(
        CryptographyAsyncClient cryptographyClient,
        String keyId,
        SecureRandom secureRandom
    ) {
        this.cryptographyClient = Objects.requireNonNull(cryptographyClient, "cryptographyClient");
        this.keyId = Objects.requireNonNull(keyId, "keyId");
        this.secureRandom = Objects.requireNonNull(secureRandom, "secureRandom");
    }

    Mono<GeneratedDataKey> generateAndWrapDataKey() {
        return Mono.defer(() -> {
            byte[] plaintextKey = new byte[DATA_KEY_BYTES];
            secureRandom.nextBytes(plaintextKey);
            return cryptographyClient.wrapKey(SDK_WRAP_ALGORITHM, plaintextKey)
                .map(result -> new GeneratedDataKey(
                    new DataKey(plaintextKey),
                    new WrappedDataKey(
                        keyId,
                        KeyManagementService.WRAP_ALGORITHM,
                        result.getEncryptedKey())))
                .doOnError(ignored -> Arrays.fill(plaintextKey, (byte) 0));
        });
    }

    Mono<DataKey> unwrapDataKey(WrappedDataKey wrappedDataKey) {
        return Mono.defer(() -> {
            validateWrappedKey(wrappedDataKey);
            return cryptographyClient.unwrapKey(SDK_WRAP_ALGORITHM, wrappedDataKey.bytes())
                .map(result -> new DataKey(result.getKey()));
        });
    }

    private void validateWrappedKey(WrappedDataKey wrappedDataKey) {
        if (!keyId.equals(wrappedDataKey.keyId())) {
            throw new IllegalArgumentException("The wrapped DEK references a different Key Vault key version");
        }
        if (!KeyManagementService.WRAP_ALGORITHM.equals(wrappedDataKey.algorithm())) {
            throw new IllegalArgumentException("Unsupported key-wrap algorithm: " + wrappedDataKey.algorithm());
        }
    }
}
