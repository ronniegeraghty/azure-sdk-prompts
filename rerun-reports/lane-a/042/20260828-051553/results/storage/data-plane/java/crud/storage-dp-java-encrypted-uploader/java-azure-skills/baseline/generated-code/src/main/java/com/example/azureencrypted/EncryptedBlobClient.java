package com.example.azureencrypted;

import com.azure.core.util.BinaryData;
import com.azure.core.util.Context;
import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.models.BlobProperties;
import com.azure.storage.blob.models.BlobStorageException;
import com.azure.storage.blob.options.BlobParallelUploadOptions;

import javax.crypto.AEADBadTagException;
import javax.crypto.Cipher;
import javax.crypto.spec.GCMParameterSpec;
import javax.crypto.spec.SecretKeySpec;
import java.security.GeneralSecurityException;
import java.security.SecureRandom;
import java.util.Arrays;
import java.util.Base64;
import java.util.HashMap;
import java.util.Map;

public final class EncryptedBlobClient {
    private static final String CONTENT_ALGORITHM = "AES/GCM/NoPadding";
    private static final String KEY_ALGORITHM = "AES";
    private static final int IV_BYTES = 12;
    private static final int TAG_BITS = 128;
    private static final String META_VERSION = "encryptionversion";
    private static final String META_ALGORITHM = "contentalgorithm";
    private static final String META_IV = "iv";
    private static final String META_WRAPPED_KEY = "wrappeddek";
    private static final String META_KEY_ID = "keyid";
    private static final String META_WRAP_ALGORITHM = "keywrapalgorithm";

    private final BlobContainerClient containerClient;
    private final KeyManagement keyManagement;
    private final SecureRandom secureRandom;

    public EncryptedBlobClient(
            BlobContainerClient containerClient, KeyManagement keyManagement) {
        this(containerClient, keyManagement, new SecureRandom());
    }

    EncryptedBlobClient(
            BlobContainerClient containerClient,
            KeyManagement keyManagement,
            SecureRandom secureRandom) {
        this.containerClient = containerClient;
        this.keyManagement = keyManagement;
        this.secureRandom = secureRandom;
    }

    public UploadResult upload(String blobName, byte[] plaintext) {
        BlobClient blobClient = containerClient.getBlobClient(blobName);
        byte[] iv = new byte[IV_BYTES];
        secureRandom.nextBytes(iv);

        try (KeyManagement.GeneratedDataKey dataKey = keyManagement.generateAndWrapDataKey()) {
            byte[] ciphertext = crypt(Cipher.ENCRYPT_MODE, plaintext, dataKey.plaintext(), iv);
            Map<String, String> metadata =
                    metadata(iv, dataKey.wrapped(), dataKey.keyId());
            try {
                BlobParallelUploadOptions options =
                        new BlobParallelUploadOptions(BinaryData.fromBytes(ciphertext))
                                .setMetadata(metadata);
                blobClient.uploadWithResponse(options, null, Context.NONE);
            } catch (BlobStorageException exception) {
                throw new EnvelopeEncryptionException(
                        "Blob Storage could not upload blob " + blobName, exception);
            }
            return new UploadResult(dataKey.keyId(), metadata.get(META_WRAPPED_KEY));
        }
    }

    public byte[] download(String blobName) {
        BlobClient blobClient = containerClient.getBlobClient(blobName);
        BlobProperties properties;
        byte[] ciphertext;
        try {
            properties = blobClient.getProperties();
            ciphertext = blobClient.downloadContent().toBytes();
        } catch (BlobStorageException exception) {
            if (exception.getStatusCode() == 404) {
                throw new EnvelopeEncryptionException(
                        "Encrypted blob does not exist: " + blobName, exception);
            }
            throw new EnvelopeEncryptionException(
                    "Blob Storage could not download blob " + blobName, exception);
        }

        EncryptionMetadata metadata = parseMetadata(properties.getMetadata(), blobName);
        byte[] dataKey =
                keyManagement.unwrapDataKey(metadata.wrappedDataKey(), metadata.keyId());
        try {
            return crypt(Cipher.DECRYPT_MODE, ciphertext, dataKey, metadata.iv());
        } finally {
            Arrays.fill(dataKey, (byte) 0);
        }
    }

    static byte[] crypt(int mode, byte[] input, byte[] dataKey, byte[] iv) {
        try {
            Cipher cipher = Cipher.getInstance(CONTENT_ALGORITHM);
            cipher.init(mode, new SecretKeySpec(dataKey, KEY_ALGORITHM), new GCMParameterSpec(TAG_BITS, iv));
            return cipher.doFinal(input);
        } catch (AEADBadTagException exception) {
            throw new EnvelopeEncryptionException(
                    "Ciphertext authentication failed; the data or metadata may be corrupted", exception);
        } catch (GeneralSecurityException exception) {
            throw new EnvelopeEncryptionException("Local AES-GCM operation failed", exception);
        }
    }

    static Map<String, String> metadata(byte[] iv, byte[] wrappedKey, String keyId) {
        Base64.Encoder encoder = Base64.getEncoder();
        Map<String, String> metadata = new HashMap<>();
        metadata.put(META_VERSION, "1");
        metadata.put(META_ALGORITHM, CONTENT_ALGORITHM);
        metadata.put(META_IV, encoder.encodeToString(iv));
        metadata.put(META_WRAPPED_KEY, encoder.encodeToString(wrappedKey));
        metadata.put(META_KEY_ID, keyId);
        metadata.put(META_WRAP_ALGORITHM, KeyManagement.KEY_WRAP_ALGORITHM.toString());
        return metadata;
    }

    static EncryptionMetadata parseMetadata(Map<String, String> metadata, String blobName) {
        try {
            if (!"1".equals(requiredMetadata(metadata, META_VERSION))
                    || !CONTENT_ALGORITHM.equals(requiredMetadata(metadata, META_ALGORITHM))
                    || !KeyManagement.KEY_WRAP_ALGORITHM.toString()
                            .equals(requiredMetadata(metadata, META_WRAP_ALGORITHM))) {
                throw new EnvelopeEncryptionException(
                        "Blob " + blobName + " uses an unsupported encryption format");
            }
            return new EncryptionMetadata(
                    Base64.getDecoder().decode(requiredMetadata(metadata, META_IV)),
                    Base64.getDecoder().decode(requiredMetadata(metadata, META_WRAPPED_KEY)),
                    requiredMetadata(metadata, META_KEY_ID));
        } catch (IllegalArgumentException exception) {
            throw new EnvelopeEncryptionException(
                    "Blob " + blobName + " has invalid encryption metadata", exception);
        }
    }

    private static String requiredMetadata(Map<String, String> metadata, String name) {
        String value = metadata.get(name);
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException("Missing metadata field " + name);
        }
        return value;
    }

    public record UploadResult(String keyId, String wrappedDataKeyBase64) {
    }

    record EncryptionMetadata(byte[] iv, byte[] wrappedDataKey, String keyId) {
    }
}
