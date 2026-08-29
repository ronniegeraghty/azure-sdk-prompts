package com.example.encryptedblob;

import com.azure.security.keyvault.keys.cryptography.CryptographyClient;
import com.azure.security.keyvault.keys.cryptography.models.KeyWrapAlgorithm;

import java.security.SecureRandom;
import java.util.Arrays;
import java.util.Objects;

public final class KeyManagementService {
    public static final String WRAP_ALGORITHM = "RSA-OAEP-256";
    private static final KeyWrapAlgorithm SDK_WRAP_ALGORITHM = KeyWrapAlgorithm.RSA_OAEP_256;
    private static final int DATA_KEY_BYTES = 32;

    private final CryptographyClient cryptographyClient;
    private final String keyId;
    private final SecureRandom secureRandom;

    public KeyManagementService(CryptographyClient cryptographyClient, String keyId) {
        this(cryptographyClient, keyId, new SecureRandom());
    }

    KeyManagementService(CryptographyClient cryptographyClient, String keyId, SecureRandom secureRandom) {
        this.cryptographyClient = Objects.requireNonNull(cryptographyClient, "cryptographyClient");
        this.keyId = Objects.requireNonNull(keyId, "keyId");
        this.secureRandom = Objects.requireNonNull(secureRandom, "secureRandom");
    }

    GeneratedDataKey generateAndWrapDataKey() {
        byte[] plaintextKey = new byte[DATA_KEY_BYTES];
        secureRandom.nextBytes(plaintextKey);
        try {
            byte[] wrapped = cryptographyClient.wrapKey(SDK_WRAP_ALGORITHM, plaintextKey).getEncryptedKey();
            return new GeneratedDataKey(
                new DataKey(plaintextKey),
                new WrappedDataKey(keyId, WRAP_ALGORITHM, wrapped));
        } catch (RuntimeException exception) {
            Arrays.fill(plaintextKey, (byte) 0);
            throw exception;
        }
    }

    DataKey unwrapDataKey(WrappedDataKey wrappedDataKey) {
        validateWrappedKey(wrappedDataKey);
        byte[] plaintextKey = cryptographyClient
            .unwrapKey(SDK_WRAP_ALGORITHM, wrappedDataKey.bytes())
            .getKey();
        return new DataKey(plaintextKey);
    }

    private void validateWrappedKey(WrappedDataKey wrappedDataKey) {
        if (!keyId.equals(wrappedDataKey.keyId())) {
            throw new IllegalArgumentException("The wrapped DEK references a different Key Vault key version");
        }
        if (!WRAP_ALGORITHM.equals(wrappedDataKey.algorithm())) {
            throw new IllegalArgumentException("Unsupported key-wrap algorithm: " + wrappedDataKey.algorithm());
        }
    }
}
