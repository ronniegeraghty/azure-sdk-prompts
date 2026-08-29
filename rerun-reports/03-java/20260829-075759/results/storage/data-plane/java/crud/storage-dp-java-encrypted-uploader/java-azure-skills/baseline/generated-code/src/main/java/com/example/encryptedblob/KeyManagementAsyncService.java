package com.example.encryptedblob;

import com.azure.core.credential.TokenCredential;
import com.azure.core.exception.AzureException;
import com.azure.core.exception.HttpResponseException;
import com.azure.security.keyvault.keys.KeyAsyncClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyAsyncClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;
import com.azure.security.keyvault.keys.cryptography.models.KeyWrapAlgorithm;
import reactor.core.publisher.Mono;

import java.util.Arrays;
import java.util.Objects;

public final class KeyManagementAsyncService {
    private static final KeyWrapAlgorithm WRAP_ALGORITHM = KeyWrapAlgorithm.RSA_OAEP_256;
    private final TokenCredential credential;
    private final String currentKeyId;
    private final CryptographyAsyncClient currentKeyClient;

    private KeyManagementAsyncService(TokenCredential credential, String keyId) {
        this.credential = credential;
        this.currentKeyId = keyId;
        this.currentKeyClient = cryptographyClient(keyId);
    }

    public static Mono<KeyManagementAsyncService> create(
            KeyAsyncClient keyClient, TokenCredential credential, String keyName) {
        Objects.requireNonNull(credential, "credential");
        return keyClient.getKey(keyName)
                .map(key -> new KeyManagementAsyncService(credential, key.getId()))
                .onErrorMap(AzureException.class,
                        e -> keyVaultFailure("resolve wrapping key '" + keyName + "'", e));
    }

    public Mono<WrappedDataKey> generateAndWrapDataKey() {
        return Mono.defer(() -> {
            byte[] plaintextKey = CipherSupport.generateDataKey();
            return currentKeyClient.wrapKey(WRAP_ALGORITHM, plaintextKey)
                    .map(result -> new WrappedDataKey(
                            plaintextKey, result.getEncryptedKey(), currentKeyId))
                    .doFinally(ignored -> Arrays.fill(plaintextKey, (byte) 0));
        }).onErrorMap(AzureException.class,
                e -> keyVaultFailure("wrap the data encryption key", e));
    }

    public Mono<byte[]> unwrapDataKey(String keyId, byte[] wrappedKey) {
        return clientFor(keyId).unwrapKey(WRAP_ALGORITHM, wrappedKey)
                .map(result -> result.getKey())
                .onErrorMap(AzureException.class,
                        e -> keyVaultFailure(
                                "unwrap the data encryption key with Key Vault key " + keyId, e));
    }

    private CryptographyAsyncClient clientFor(String keyId) {
        return currentKeyId.equals(keyId) ? currentKeyClient : cryptographyClient(keyId);
    }

    private CryptographyAsyncClient cryptographyClient(String keyId) {
        return new CryptographyClientBuilder()
                .keyIdentifier(keyId)
                .credential(credential)
                .buildAsyncClient();
    }

    private static EnvelopeEncryptionException keyVaultFailure(
            String operation, AzureException cause) {
        String status = cause instanceof HttpResponseException httpError
                ? " (HTTP " + httpError.getResponse().getStatusCode() + ")"
                : "";
        return new EnvelopeEncryptionException(
                "Azure Key Vault could not " + operation + status,
                cause);
    }
}
