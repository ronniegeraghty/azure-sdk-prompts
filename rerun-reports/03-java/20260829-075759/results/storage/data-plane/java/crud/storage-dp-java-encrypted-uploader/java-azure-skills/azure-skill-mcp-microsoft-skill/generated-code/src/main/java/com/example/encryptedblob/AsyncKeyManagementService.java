package com.example.encryptedblob;

import com.azure.core.credential.TokenCredential;
import com.azure.core.exception.HttpResponseException;
import com.azure.security.keyvault.keys.KeyAsyncClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyAsyncClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;
import reactor.core.publisher.Mono;

import java.security.SecureRandom;

public final class AsyncKeyManagementService {
    private static final int DATA_KEY_BYTES = 32;

    private final KeyAsyncClient keyClient;
    private final TokenCredential credential;
    private final String keyName;
    private final SecureRandom secureRandom;

    public AsyncKeyManagementService(KeyAsyncClient keyClient, TokenCredential credential, String keyName) {
        this(keyClient, credential, keyName, new SecureRandom());
    }

    AsyncKeyManagementService(
            KeyAsyncClient keyClient,
            TokenCredential credential,
            String keyName,
            SecureRandom secureRandom) {
        this.keyClient = keyClient;
        this.credential = credential;
        this.keyName = keyName;
        this.secureRandom = secureRandom;
    }

    Mono<ProtectedDataKey> generateAndWrapDataKey() {
        return keyClient.getKey(keyName)
                .flatMap(key -> Mono.defer(() -> {
                    DataKey dataKey = generateDataKey();
                    return cryptographyClient(key.getId())
                            .wrapKey(KeyManagementService.WRAP_ALGORITHM, dataKey.bytes())
                            .map(result -> new ProtectedDataKey(
                                    dataKey, result.getEncryptedKey(), key.getId()))
                            .doOnError(ignored -> dataKey.close());
                }))
                .onErrorMap(
                        HttpResponseException.class,
                        exception -> keyVaultException("wrap", keyName, exception));
    }

    Mono<DataKey> unwrapDataKey(String keyId, byte[] wrappedKey) {
        return cryptographyClient(keyId)
                .unwrapKey(KeyManagementService.WRAP_ALGORITHM, wrappedKey)
                .map(result -> new DataKey(result.getKey()))
                .onErrorMap(
                        HttpResponseException.class,
                        exception -> keyVaultException("unwrap", keyId, exception));
    }

    private DataKey generateDataKey() {
        byte[] key = new byte[DATA_KEY_BYTES];
        secureRandom.nextBytes(key);
        return new DataKey(key);
    }

    private CryptographyAsyncClient cryptographyClient(String keyId) {
        return new CryptographyClientBuilder()
                .keyIdentifier(keyId)
                .credential(credential)
                .buildAsyncClient();
    }

    private static KeyManagementException keyVaultException(
            String operation,
            String key,
            HttpResponseException exception) {
        return new KeyManagementException(
                "Key Vault could not " + operation + " the data key with key '" + key
                        + "' (HTTP " + exception.getResponse().getStatusCode() + ")",
                exception);
    }
}
