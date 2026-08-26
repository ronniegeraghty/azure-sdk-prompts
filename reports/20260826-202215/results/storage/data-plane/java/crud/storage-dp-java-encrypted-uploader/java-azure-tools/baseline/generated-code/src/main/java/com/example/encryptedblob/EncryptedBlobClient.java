package com.example.encryptedblob;

import com.azure.core.util.BinaryData;
import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.models.BlobProperties;
import com.azure.storage.blob.models.BlobStorageException;
import com.azure.storage.blob.options.BlobParallelUploadOptions;

import java.security.SecureRandom;
import java.util.Base64;
import java.util.HashMap;
import java.util.Map;

public final class EncryptedBlobClient {
    private static final String META_VERSION = "encversion";
    private static final String META_ALGORITHM = "encalgorithm";
    private static final String META_IV = "enciv";
    private static final String META_WRAPPED_KEY = "wrappeddek";
    private static final String META_KEY_ID = "keyid";
    private static final String FORMAT_VERSION = "1";

    private final BlobContainerClient containerClient;
    private final KeyManagement keyManagement;
    private final SecureRandom secureRandom;

    public EncryptedBlobClient(BlobContainerClient containerClient, KeyManagement keyManagement) {
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
        byte[] iv = BlobEncryption.newIv(secureRandom);

        try (DataEncryptionKey dataKey = keyManagement.generateDataKey()) {
            byte[] ciphertext = BlobEncryption.encrypt(plaintext, dataKey, iv);
            byte[] wrappedKey = keyManagement.wrap(dataKey);
            Map<String, String> metadata = metadata(iv, wrappedKey, keyManagement.keyId());

            try {
                blobClient.uploadWithResponse(
                        new BlobParallelUploadOptions(BinaryData.fromBytes(ciphertext))
                                .setMetadata(metadata),
                        null,
                        null);
            } catch (BlobStorageException e) {
                throw blobFailure("upload encrypted blob", blobName, e);
            }
            return new UploadResult(keyManagement.keyId(), Base64.getEncoder().encodeToString(wrappedKey));
        }
    }

    public byte[] download(String blobName) {
        BlobClient blobClient = containerClient.getBlobClient(blobName);
        BlobProperties properties;
        byte[] ciphertext;
        try {
            properties = blobClient.getProperties();
            ciphertext = blobClient.downloadContent().toBytes();
        } catch (BlobStorageException e) {
            throw blobFailure("download encrypted blob", blobName, e);
        }

        EncryptionMetadata metadata = parseMetadata(properties.getMetadata(), blobName);
        try (DataEncryptionKey dataKey =
                     keyManagement.unwrap(metadata.wrappedKey(), metadata.keyId())) {
            return BlobEncryption.decrypt(ciphertext, dataKey, metadata.iv());
        }
    }

    static Map<String, String> metadata(byte[] iv, byte[] wrappedKey, String keyId) {
        Map<String, String> metadata = new HashMap<>();
        metadata.put(META_VERSION, FORMAT_VERSION);
        metadata.put(META_ALGORITHM, BlobEncryption.ALGORITHM);
        metadata.put(META_IV, Base64.getEncoder().encodeToString(iv));
        metadata.put(META_WRAPPED_KEY, Base64.getEncoder().encodeToString(wrappedKey));
        metadata.put(META_KEY_ID, keyId);
        return metadata;
    }

    static EncryptionMetadata parseMetadata(Map<String, String> metadata, String blobName) {
        String version = required(metadata, META_VERSION, blobName);
        String algorithm = required(metadata, META_ALGORITHM, blobName);
        if (!FORMAT_VERSION.equals(version) || !BlobEncryption.ALGORITHM.equals(algorithm)) {
            throw new EncryptedBlobException(
                    "Blob '" + blobName + "' uses unsupported encryption format "
                            + version + "/" + algorithm);
        }
        try {
            byte[] iv = Base64.getDecoder().decode(required(metadata, META_IV, blobName));
            if (iv.length != BlobEncryption.IV_SIZE_BYTES) {
                throw new EncryptedBlobException(
                        "Blob '" + blobName + "' has an invalid AES-GCM IV length");
            }
            byte[] wrappedKey = Base64.getDecoder().decode(required(metadata, META_WRAPPED_KEY, blobName));
            String keyId = required(metadata, META_KEY_ID, blobName);
            return new EncryptionMetadata(iv, wrappedKey, keyId);
        } catch (IllegalArgumentException e) {
            throw new EncryptedBlobException(
                    "Blob '" + blobName + "' contains invalid base64 encryption metadata", e);
        }
    }

    private static String required(Map<String, String> metadata, String name, String blobName) {
        String value = metadata.get(name);
        if (value == null || value.isBlank()) {
            throw new EncryptedBlobException(
                    "Blob '" + blobName + "' is missing encryption metadata '" + name + "'");
        }
        return value;
    }

    private static EncryptedBlobException blobFailure(
            String operation, String blobName, BlobStorageException cause) {
        return new EncryptedBlobException(
                "Blob Storage could not " + operation + " '" + blobName
                        + "' (status " + cause.getStatusCode() + ", code "
                        + cause.getErrorCode() + ")",
                cause);
    }

    public record UploadResult(String keyId, String wrappedDataKeyBase64) {
    }

    record EncryptionMetadata(byte[] iv, byte[] wrappedKey, String keyId) {
    }
}
