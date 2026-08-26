package com.example.encryptedblob;

import com.azure.core.credential.TokenCredential;
import com.azure.core.exception.HttpResponseException;
import com.azure.security.keyvault.keys.KeyAsyncClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyAsyncClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;
import reactor.core.publisher.Mono;

import java.security.SecureRandom;
import java.util.Arrays;
import java.util.Objects;

public final class AsyncKeyManagementClient {
    private static final int DATA_KEY_BYTES = 32;

    private final KeyAsyncClient keyClient;
    private final TokenCredential credential;
    private final String keyName;
    private final SecureRandom secureRandom;

    public AsyncKeyManagementClient(
        KeyAsyncClient keyClient,
        TokenCredential credential,
        String keyName
    ) {
        this(keyClient, credential, keyName, new SecureRandom());
    }

    AsyncKeyManagementClient(
        KeyAsyncClient keyClient,
        TokenCredential credential,
        String keyName,
        SecureRandom secureRandom
    ) {
        this.keyClient = Objects.requireNonNull(keyClient, "keyClient");
        this.credential = Objects.requireNonNull(credential, "credential");
        this.keyName = Objects.requireNonNull(keyName, "keyName");
        this.secureRandom = Objects.requireNonNull(secureRandom, "secureRandom");
    }

    Mono<EnvelopeKey> generateAndWrapDataKey() {
        return Mono.defer(() -> {
            byte[] dataKey = new byte[DATA_KEY_BYTES];
            secureRandom.nextBytes(dataKey);

            return keyClient.getKey(keyName)
                .flatMap(key -> cryptographyClient(key.getId())
                    .wrapKey(SyncKeyManagementClient.WRAP_ALGORITHM, dataKey)
                    .map(result -> new EnvelopeKey(
                        new ProtectedDataKey(
                            key.getId(),
                            SyncKeyManagementClient.WRAP_ALGORITHM.toString(),
                            result.getEncryptedKey()
                        ),
                        dataKey
                    )))
                .doOnError(ignored -> Arrays.fill(dataKey, (byte) 0));
        }).onErrorMap(
            HttpResponseException.class,
            e -> keyVaultFailure("wrap a new data encryption key", e)
        );
    }

    Mono<byte[]> unwrapDataKey(ProtectedDataKey protectedKey) {
        return Mono.defer(() -> {
            Objects.requireNonNull(protectedKey, "protectedKey");
            SyncKeyManagementClient.validateAlgorithm(protectedKey.algorithm());
            return cryptographyClient(protectedKey.keyId())
                .unwrapKey(SyncKeyManagementClient.WRAP_ALGORITHM, protectedKey.wrappedKey())
                .map(result -> result.getKey());
        }).onErrorMap(
            HttpResponseException.class,
            e -> keyVaultFailure(
                "unwrap the data encryption key with " + protectedKey.keyId(),
                e
            )
        );
    }

    private CryptographyAsyncClient cryptographyClient(String keyId) {
        return new CryptographyClientBuilder()
            .keyIdentifier(keyId)
            .credential(credential)
            .buildAsyncClient();
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
