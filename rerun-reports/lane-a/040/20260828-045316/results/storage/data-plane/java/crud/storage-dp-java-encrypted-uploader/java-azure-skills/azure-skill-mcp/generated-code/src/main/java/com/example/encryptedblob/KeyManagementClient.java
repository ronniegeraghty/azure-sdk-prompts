package com.example.encryptedblob;

import com.azure.core.credential.TokenCredential;
import com.azure.core.exception.HttpResponseException;
import com.azure.security.keyvault.keys.KeyClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyClient;
import com.azure.security.keyvault.keys.cryptography.CryptographyClientBuilder;
import com.azure.security.keyvault.keys.cryptography.models.KeyWrapAlgorithm;
import com.azure.security.keyvault.keys.models.KeyVaultKey;

import java.security.SecureRandom;

public final class KeyManagementClient {
    public static final String WRAP_ALGORITHM = "RSA-OAEP-256";
    private static final int DATA_KEY_BYTES = 32;

    private final KeyClient keyClient;
    private final TokenCredential credential;
    private final String keyName;
    private final SecureRandom secureRandom;

    public KeyManagementClient(KeyClient keyClient, TokenCredential credential, String keyName) {
        this(keyClient, credential, keyName, new SecureRandom());
    }

    KeyManagementClient(
            KeyClient keyClient,
            TokenCredential credential,
            String keyName,
            SecureRandom secureRandom) {
        this.keyClient = keyClient;
        this.credential = credential;
        this.keyName = keyName;
        this.secureRandom = secureRandom;
    }

    GeneratedDataKey generateAndProtectDataKey() {
        byte[] rawKey = new byte[DATA_KEY_BYTES];
        secureRandom.nextBytes(rawKey);
        DataKeyMaterial plaintextKey = new DataKeyMaterial(rawKey);

        try {
            KeyVaultKey vaultKey = keyClient.getKey(keyName);
            String keyId = vaultKey.getId();
            CryptographyClient cryptographyClient = cryptographyClient(keyId);
            byte[] wrappedKey = cryptographyClient
                    .wrapKey(KeyWrapAlgorithm.RSA_OAEP_256, plaintextKey.bytesForWrapping())
                    .getEncryptedKey();
            return new GeneratedDataKey(
                    plaintextKey,
                    new ProtectedDataKey(keyId, WRAP_ALGORITHM, wrappedKey));
        } catch (HttpResponseException exception) {
            plaintextKey.close();
            throw new KeyManagementException("Key Vault could not protect the data key", exception);
        } catch (RuntimeException exception) {
            plaintextKey.close();
            throw exception;
        }
    }

    DataKeyMaterial recoverDataKey(ProtectedDataKey protectedKey) {
        validateAlgorithm(protectedKey.wrapAlgorithm());
        try {
            byte[] rawKey = cryptographyClient(protectedKey.keyId())
                    .unwrapKey(KeyWrapAlgorithm.RSA_OAEP_256, protectedKey.wrappedKey())
                    .getKey();
            return new DataKeyMaterial(rawKey);
        } catch (HttpResponseException exception) {
            throw new KeyManagementException(
                    "Key Vault could not recover the data key; the key may be disabled or inaccessible",
                    exception);
        }
    }

    private CryptographyClient cryptographyClient(String keyId) {
        return new CryptographyClientBuilder()
                .keyIdentifier(keyId)
                .credential(credential)
                .buildClient();
    }

    static void validateAlgorithm(String algorithm) {
        if (!WRAP_ALGORITHM.equals(algorithm)) {
            throw new IllegalArgumentException("Unsupported key wrap algorithm: " + algorithm);
        }
    }
}
