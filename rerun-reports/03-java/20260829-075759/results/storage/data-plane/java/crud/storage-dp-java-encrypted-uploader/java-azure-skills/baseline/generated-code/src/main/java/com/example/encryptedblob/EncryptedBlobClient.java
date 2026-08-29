package com.example.encryptedblob;

import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.core.exception.AzureException;
import com.azure.storage.blob.models.BlobStorageException;
import com.azure.storage.blob.options.BlobParallelUploadOptions;
import com.azure.core.util.BinaryData;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Arrays;
import java.util.Map;

public final class EncryptedBlobClient {
    private final BlobContainerClient containerClient;
    private final KeyManagementService keyManagement;

    public EncryptedBlobClient(
            BlobContainerClient containerClient, KeyManagementService keyManagement) {
        this.containerClient = containerClient;
        this.keyManagement = keyManagement;
    }

    public UploadReceipt upload(String blobName, byte[] plaintext) {
        WrappedDataKey dataKey = keyManagement.generateAndWrapDataKey();
        byte[] plaintextKey = dataKey.plaintextKey();
        try {
            byte[] aad = EnvelopeMetadata.authenticatedData(blobName, dataKey.keyId());
            CipherSupport.EncryptedData encrypted =
                    CipherSupport.encrypt(plaintext, plaintextKey, aad);
            Map<String, String> metadata = EnvelopeMetadata.create(
                    dataKey.keyId(), dataKey.wrappedKey(), encrypted.iv());
            blob(blobName).uploadWithResponse(
                    new BlobParallelUploadOptions(BinaryData.fromBytes(encrypted.ciphertext()))
                            .setMetadata(metadata),
                    null,
                    null);
            return new UploadReceipt(dataKey.keyId(), dataKey.wrappedKey());
        } catch (AzureException e) {
            throw blobFailure("upload", blobName, e);
        } finally {
            Arrays.fill(plaintextKey, (byte) 0);
            dataKey.destroy();
        }
    }

    public UploadReceipt uploadFile(String blobName, Path source) {
        try {
            return upload(blobName, Files.readAllBytes(source));
        } catch (IOException e) {
            throw new EnvelopeEncryptionException("Could not read input file " + source, e);
        }
    }

    public byte[] download(String blobName) {
        BlobClient blob = blob(blobName);
        try {
            Map<String, String> metadata = blob.getProperties().getMetadata();
            EnvelopeMetadata.Parsed envelope = EnvelopeMetadata.parse(metadata);
            byte[] plaintextKey =
                    keyManagement.unwrapDataKey(envelope.keyId(), envelope.wrappedKey());
            try {
                byte[] aad = EnvelopeMetadata.authenticatedData(blobName, envelope.keyId());
                return CipherSupport.decrypt(
                        blob.downloadContent().toBytes(), plaintextKey, envelope.iv(), aad);
            } finally {
                Arrays.fill(plaintextKey, (byte) 0);
            }
        } catch (AzureException e) {
            throw blobFailure("download", blobName, e);
        }
    }

    public void downloadFile(String blobName, Path destination) {
        try {
            Files.write(destination, download(blobName));
        } catch (IOException e) {
            throw new EnvelopeEncryptionException(
                    "Could not write decrypted file " + destination, e);
        }
    }

    private BlobClient blob(String blobName) {
        return containerClient.getBlobClient(blobName);
    }

    private static EnvelopeEncryptionException blobFailure(
            String operation, String blobName, AzureException cause) {
        String details = cause instanceof BlobStorageException storageError
                ? " (HTTP " + storageError.getStatusCode()
                        + ", error " + storageError.getErrorCode() + ")"
                : "";
        return new EnvelopeEncryptionException(
                "Azure Blob Storage could not " + operation + " blob '" + blobName + "'"
                        + details,
                cause);
    }
}
