package com.example.encryptedblob;

import com.azure.core.credential.TokenCredential;
import com.azure.core.exception.HttpResponseException;
import com.azure.security.keyvault.keys.KeyAsyncClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyAsyncClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;
import com.azure.security.keyvault.keys.cryptography.models.KeyWrapAlgorithm;
import reactor.core.publisher.Mono;

import java.security.SecureRandom;
import java.util.Arrays;

public final class AsyncKeyManagement {
    private static final int DATA_KEY_BYTES = 32;
    private static final KeyWrapAlgorithm WRAP_ALGORITHM = KeyWrapAlgorithm.RSA_OAEP_256;

    private final KeyAsyncClient keyClient;
    private final TokenCredential credential;
    private final String keyName;
    private final SecureRandom secureRandom;

    public AsyncKeyManagement(KeyAsyncClient keyClient, TokenCredential credential, String keyName) {
        this.keyClient = keyClient;
        this.credential = credential;
        this.keyName = keyName;
        this.secureRandom = new SecureRandom();
    }

    Mono<DataKeyEnvelope> generateAndWrapKey() {
        return Mono.defer(() -> {
            byte[] dataKey = new byte[DATA_KEY_BYTES];
            secureRandom.nextBytes(dataKey);

            return keyClient.getKey(keyName)
                .flatMap(key -> {
                    String versionedKeyId = key.getId();
                    return cryptoClient(versionedKeyId).wrapKey(WRAP_ALGORITHM, dataKey)
                        .map(result -> new DataKeyEnvelope(
                            dataKey,
                            new ProtectedDataKey(versionedKeyId, result.getEncryptedKey())));
                })
                .doOnError(ignored -> Arrays.fill(dataKey, (byte) 0))
                .doOnCancel(() -> Arrays.fill(dataKey, (byte) 0))
                .onErrorMap(HttpResponseException.class, exception ->
                    keyVaultFailure("wrap a data encryption key", exception));
        });
    }

    Mono<byte[]> unwrapKey(ProtectedDataKey protectedKey) {
        return cryptoClient(protectedKey.keyId())
            .unwrapKey(WRAP_ALGORITHM, protectedKey.wrappedKey())
            .map(result -> result.getKey())
            .onErrorMap(HttpResponseException.class, exception ->
                keyVaultFailure("unwrap the data encryption key", exception));
    }

    private CryptographyAsyncClient cryptoClient(String keyId) {
        return new CryptographyClientBuilder()
            .keyIdentifier(keyId)
            .credential(credential)
            .buildAsyncClient();
    }

    private static EnvelopeEncryptionException keyVaultFailure(String operation, HttpResponseException exception) {
        int status = exception.getResponse() == null ? -1 : exception.getResponse().getStatusCode();
        return new EnvelopeEncryptionException(
            "Key Vault could not " + operation + " (HTTP " + status
                + "). Verify that the stored key version is enabled and permits wrapKey/unwrapKey.",
            exception);
    }
}
