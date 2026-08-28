package com.example.encryptedblob;

import com.azure.core.credential.TokenCredential;
import com.azure.core.exception.HttpResponseException;
import com.azure.security.keyvault.keys.KeyAsyncClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyAsyncClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;
import com.azure.security.keyvault.keys.cryptography.models.KeyWrapAlgorithm;
import reactor.core.publisher.Mono;

import java.security.SecureRandom;

public final class AsyncKeyManagementClient {
    private static final int DATA_KEY_BYTES = 32;

    private final KeyAsyncClient keyClient;
    private final TokenCredential credential;
    private final String keyName;
    private final SecureRandom secureRandom;

    public AsyncKeyManagementClient(
            KeyAsyncClient keyClient,
            TokenCredential credential,
            String keyName) {
        this(keyClient, credential, keyName, new SecureRandom());
    }

    AsyncKeyManagementClient(
            KeyAsyncClient keyClient,
            TokenCredential credential,
            String keyName,
            SecureRandom secureRandom) {
        this.keyClient = keyClient;
        this.credential = credential;
        this.keyName = keyName;
        this.secureRandom = secureRandom;
    }

    Mono<GeneratedDataKey> generateAndProtectDataKey() {
        return Mono.defer(() -> {
            byte[] rawKey = new byte[DATA_KEY_BYTES];
            secureRandom.nextBytes(rawKey);
            DataKeyMaterial plaintextKey = new DataKeyMaterial(rawKey);

            return keyClient.getKey(keyName)
                    .flatMap(vaultKey -> {
                        String keyId = vaultKey.getId();
                        return cryptographyClient(keyId)
                                .wrapKey(KeyWrapAlgorithm.RSA_OAEP_256, plaintextKey.bytesForWrapping())
                                .map(result -> new GeneratedDataKey(
                                        plaintextKey,
                                        new ProtectedDataKey(
                                                keyId,
                                                KeyManagementClient.WRAP_ALGORITHM,
                                                result.getEncryptedKey())));
                    })
                    .doOnError(ignored -> plaintextKey.close())
                    .onErrorMap(
                            HttpResponseException.class,
                            exception -> new KeyManagementException(
                                    "Key Vault could not protect the data key",
                                    exception));
        });
    }

    Mono<DataKeyMaterial> recoverDataKey(ProtectedDataKey protectedKey) {
        return Mono.defer(() -> {
            KeyManagementClient.validateAlgorithm(protectedKey.wrapAlgorithm());
            return cryptographyClient(protectedKey.keyId())
                    .unwrapKey(KeyWrapAlgorithm.RSA_OAEP_256, protectedKey.wrappedKey())
                    .map(result -> new DataKeyMaterial(result.getKey()))
                    .onErrorMap(
                            HttpResponseException.class,
                            exception -> new KeyManagementException(
                                    "Key Vault could not recover the data key; "
                                            + "the key may be disabled or inaccessible",
                                    exception));
        });
    }

    private CryptographyAsyncClient cryptographyClient(String keyId) {
        return new CryptographyClientBuilder()
                .keyIdentifier(keyId)
                .credential(credential)
                .buildAsyncClient();
    }
}
