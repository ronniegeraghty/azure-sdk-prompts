package com.example.encryptedblob;

import com.azure.core.util.BinaryData;
import com.azure.core.util.Context;
import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.models.BlobHttpHeaders;
import com.azure.storage.blob.models.BlobStorageException;
import com.azure.storage.blob.options.BlobParallelUploadOptions;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Arrays;
import java.util.Objects;

public final class SyncEncryptedBlobClient {
    private final BlobContainerClient containerClient;
    private final SyncKeyManagementClient keyManagementClient;

    public SyncEncryptedBlobClient(
        BlobContainerClient containerClient,
        SyncKeyManagementClient keyManagementClient
    ) {
        this.containerClient = Objects.requireNonNull(containerClient, "containerClient");
        this.keyManagementClient =
            Objects.requireNonNull(keyManagementClient, "keyManagementClient");
    }

    public UploadResult upload(Path source, String blobName) throws IOException {
        return upload(Files.readAllBytes(source), blobName);
    }

    public UploadResult upload(byte[] plaintext, String blobName) {
        Objects.requireNonNull(plaintext, "plaintext");
        Objects.requireNonNull(blobName, "blobName");

        try (EnvelopeKey envelopeKey = keyManagementClient.generateAndWrapDataKey()) {
            ProtectedDataKey protectedKey = envelopeKey.protectedDataKey();
            byte[] authenticatedData = EncryptionMetadata.authenticatedData(protectedKey);
            LocalAesGcm.EncryptedPayload encrypted = LocalAesGcm.encrypt(
                envelopeKey.dataKey(),
                plaintext,
                authenticatedData
            );
            EncryptionMetadata metadata =
                EncryptionMetadata.create(protectedKey, encrypted.iv());
            BlobClient blobClient = containerClient.getBlobClient(blobName);
            BlobParallelUploadOptions options = new BlobParallelUploadOptions(
                BinaryData.fromBytes(encrypted.ciphertext())
            )
                .setMetadata(metadata.toMap())
                .setHeaders(new BlobHttpHeaders().setContentType("application/octet-stream"));

            try {
                blobClient.uploadWithResponse(options, null, Context.NONE);
            } catch (BlobStorageException e) {
                throw BlobEncryptionException.storageFailure("upload", blobName, e);
            }
            return new UploadResult(protectedKey.keyId(), protectedKey.wrappedKey());
        }
    }

    public void download(String blobName, Path destination) throws IOException {
        Files.write(destination, download(blobName));
    }

    public byte[] download(String blobName) {
        Objects.requireNonNull(blobName, "blobName");
        BlobClient blobClient = containerClient.getBlobClient(blobName);

        EncryptionMetadata metadata;
        byte[] ciphertext;
        try {
            metadata = EncryptionMetadata.parse(blobClient.getProperties().getMetadata());
            ciphertext = blobClient.downloadContent().toBytes();
        } catch (BlobStorageException e) {
            throw BlobEncryptionException.storageFailure("download", blobName, e);
        }

        byte[] dataKey = keyManagementClient.unwrapDataKey(metadata.protectedDataKey());
        try {
            return LocalAesGcm.decrypt(
                dataKey,
                ciphertext,
                metadata.iv(),
                metadata.authenticatedData()
            );
        } finally {
            Arrays.fill(dataKey, (byte) 0);
        }
    }
}
