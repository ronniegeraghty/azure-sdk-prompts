package com.example.encryptedblob;

import com.azure.core.util.BinaryData;
import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.models.BlobStorageException;
import com.azure.storage.blob.options.BlobParallelUploadOptions;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.security.SecureRandom;
import java.util.Base64;

public final class EncryptedBlobClient {
    private final BlobContainerClient containerClient;
    private final KeyManagementService keyManagement;
    private final SecureRandom secureRandom;

    public EncryptedBlobClient(
            BlobContainerClient containerClient,
            KeyManagementService keyManagement) {
        this(containerClient, keyManagement, new SecureRandom());
    }

    EncryptedBlobClient(
            BlobContainerClient containerClient,
            KeyManagementService keyManagement,
            SecureRandom secureRandom) {
        this.containerClient = containerClient;
        this.keyManagement = keyManagement;
        this.secureRandom = secureRandom;
    }

    public EncryptedBlobInfo upload(String blobName, byte[] plaintext) {
        try (ProtectedDataKey protectedKey = keyManagement.generateAndWrapDataKey()) {
            byte[] iv = EnvelopeCrypto.generateIv(secureRandom);
            BlobEncryptionMetadata metadata = BlobEncryptionMetadata.create(
                    protectedKey.keyId(), protectedKey.wrappedKey(), iv);
            byte[] ciphertext = EnvelopeCrypto.encrypt(
                    plaintext,
                    protectedKey.dataKey().bytes(),
                    iv,
                    metadata.authenticatedMetadata());

            try {
                BlobClient blobClient = containerClient.getBlobClient(blobName);
                blobClient.uploadWithResponse(
                        new BlobParallelUploadOptions(BinaryData.fromBytes(ciphertext))
                                .setMetadata(metadata.toMap()),
                        null,
                        null);
            } catch (BlobStorageException exception) {
                throw blobException("upload", blobName, exception);
            }

            return new EncryptedBlobInfo(
                    protectedKey.keyId(),
                    Base64.getEncoder().encodeToString(protectedKey.wrappedKey()));
        }
    }

    public EncryptedBlobInfo uploadFile(String blobName, Path source) {
        try {
            return upload(blobName, Files.readAllBytes(source));
        } catch (IOException exception) {
            throw new EncryptionStorageException("Could not read source file: " + source, exception);
        }
    }

    public byte[] download(String blobName) {
        BlobClient blobClient = containerClient.getBlobClient(blobName);
        BlobEncryptionMetadata metadata;
        byte[] ciphertext;
        try {
            metadata = BlobEncryptionMetadata.parse(blobClient.getProperties().getMetadata());
            ciphertext = blobClient.downloadContent().toBytes();
        } catch (BlobStorageException exception) {
            throw blobException("download", blobName, exception);
        }

        try (DataKey dataKey = keyManagement.unwrapDataKey(metadata.keyId(), metadata.wrappedKey())) {
            return EnvelopeCrypto.decrypt(
                    ciphertext,
                    dataKey.bytes(),
                    metadata.iv(),
                    metadata.authenticatedMetadata());
        }
    }

    public void downloadFile(String blobName, Path destination) {
        try {
            Files.write(destination, download(blobName));
        } catch (IOException exception) {
            throw new EncryptionStorageException("Could not write destination file: " + destination, exception);
        }
    }

    private static EncryptionStorageException blobException(
            String operation,
            String blobName,
            BlobStorageException exception) {
        String errorCode = exception.getErrorCode() == null
                ? "unknown"
                : exception.getErrorCode().toString();
        return new EncryptionStorageException(
                "Blob Storage could not " + operation + " blob '" + blobName
                        + "' (HTTP " + exception.getStatusCode() + ", " + errorCode + ")",
                exception);
    }
}
