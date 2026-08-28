package com.example.encryptedblob;

import com.azure.core.exception.HttpResponseException;
import com.azure.core.util.BinaryData;
import com.azure.core.util.Context;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.models.BlobRequestConditions;
import com.azure.storage.blob.models.BlobStorageException;
import com.azure.storage.blob.models.DownloadRetryOptions;
import com.azure.storage.blob.options.BlockBlobSimpleUploadOptions;
import com.azure.storage.blob.specialized.BlockBlobClient;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.time.Duration;
import java.util.Base64;

public final class EncryptedBlobClient {
    private final BlobContainerClient containerClient;
    private final KeyManagementClient keyManagementClient;
    private final EnvelopeEncryption encryption;

    public EncryptedBlobClient(
            BlobContainerClient containerClient,
            KeyManagementClient keyManagementClient) {
        this.containerClient = containerClient;
        this.keyManagementClient = keyManagementClient;
        this.encryption = new EnvelopeEncryption();
    }

    public UploadResult upload(Path source, String blobName) {
        try {
            return upload(Files.readAllBytes(source), blobName);
        } catch (IOException exception) {
            throw new EncryptedBlobException("Could not read source file: " + source, exception);
        }
    }

    public UploadResult upload(byte[] plaintext, String blobName) {
        try (GeneratedDataKey generatedKey = keyManagementClient.generateAndProtectDataKey()) {
            EnvelopeEncryption.EncryptedPayload payload =
                    encryption.encrypt(plaintext, generatedKey.plaintextKey(), blobName);
            ProtectedDataKey protectedKey = generatedKey.protectedKey();
            BlockBlobSimpleUploadOptions options = new BlockBlobSimpleUploadOptions(
                    BinaryData.fromBytes(payload.ciphertext()))
                    .setMetadata(encryption.metadata(payload, protectedKey));

            containerClient.getBlobClient(blobName)
                    .getBlockBlobClient()
                    .uploadWithResponse(options, Duration.ofMinutes(2), Context.NONE);

            return new UploadResult(
                    protectedKey.keyId(),
                    Base64.getEncoder().encodeToString(protectedKey.wrappedKey()));
        } catch (BlobStorageException exception) {
            throw new EncryptedBlobException(
                    "Blob Storage upload failed for '" + blobName + "': "
                            + exception.getErrorCode(),
                    exception);
        }
    }

    public byte[] download(String blobName) {
        BlockBlobClient blobClient = containerClient.getBlobClient(blobName).getBlockBlobClient();
        try {
            var properties = blobClient.getProperties();
            var metadata = encryption.parseMetadata(properties.getMetadata());
            BlobRequestConditions conditions = new BlobRequestConditions()
                    .setIfMatch(properties.getETag());
            byte[] ciphertext = blobClient.downloadContentWithResponse(
                            new DownloadRetryOptions(),
                            conditions,
                            Duration.ofMinutes(2),
                            Context.NONE)
                    .getValue()
                    .toBytes();

            try (DataKeyMaterial dataKey =
                         keyManagementClient.recoverDataKey(metadata.protectedKey())) {
                return encryption.decrypt(
                        ciphertext,
                        metadata.iv(),
                        dataKey,
                        blobName);
            }
        } catch (BlobStorageException exception) {
            String detail = exception.getStatusCode() == 404
                    ? "blob does not exist"
                    : String.valueOf(exception.getErrorCode());
            throw new EncryptedBlobException(
                    "Blob Storage download failed for '" + blobName + "': " + detail,
                    exception);
        } catch (HttpResponseException exception) {
            throw new EncryptedBlobException(
                    "Blob Storage download failed for '" + blobName + "'",
                    exception);
        }
    }

    public void download(String blobName, Path destination) {
        try {
            Files.write(destination, download(blobName));
        } catch (IOException exception) {
            throw new EncryptedBlobException(
                    "Could not write decrypted file: " + destination,
                    exception);
        }
    }
}
