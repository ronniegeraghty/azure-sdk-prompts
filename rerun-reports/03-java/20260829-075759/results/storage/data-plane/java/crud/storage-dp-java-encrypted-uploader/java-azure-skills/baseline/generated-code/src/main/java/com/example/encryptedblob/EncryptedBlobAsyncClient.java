package com.example.encryptedblob;

import com.azure.core.util.BinaryData;
import com.azure.core.exception.AzureException;
import com.azure.storage.blob.BlobAsyncClient;
import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.models.BlobStorageException;
import com.azure.storage.blob.options.BlobParallelUploadOptions;
import reactor.core.publisher.Mono;
import reactor.core.scheduler.Schedulers;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Arrays;
import java.util.Map;

public final class EncryptedBlobAsyncClient {
    private final BlobContainerAsyncClient containerClient;
    private final KeyManagementAsyncService keyManagement;

    public EncryptedBlobAsyncClient(
            BlobContainerAsyncClient containerClient,
            KeyManagementAsyncService keyManagement) {
        this.containerClient = containerClient;
        this.keyManagement = keyManagement;
    }

    public Mono<UploadReceipt> upload(String blobName, byte[] plaintext) {
        return keyManagement.generateAndWrapDataKey()
                .flatMap(dataKey -> uploadEncrypted(blobName, plaintext, dataKey)
                        .doFinally(ignored -> dataKey.destroy()));
    }

    public Mono<UploadReceipt> uploadFile(String blobName, Path source) {
        return Mono.fromCallable(() -> Files.readAllBytes(source))
                .subscribeOn(Schedulers.boundedElastic())
                .onErrorMap(IOException.class,
                        e -> new EnvelopeEncryptionException(
                                "Could not read input file " + source, e))
                .flatMap(bytes -> upload(blobName, bytes));
    }

    public Mono<byte[]> download(String blobName) {
        BlobAsyncClient blob = blob(blobName);
        return blob.getProperties()
                .map(properties -> EnvelopeMetadata.parse(properties.getMetadata()))
                .flatMap(envelope -> keyManagement
                        .unwrapDataKey(envelope.keyId(), envelope.wrappedKey())
                        .flatMap(key -> blob.downloadContent()
                                .map(content -> decrypt(blobName, content, envelope, key))
                                .doFinally(ignored -> Arrays.fill(key, (byte) 0))))
                .onErrorMap(AzureException.class,
                        e -> blobFailure("download", blobName, e));
    }

    public Mono<Void> downloadFile(String blobName, Path destination) {
        return download(blobName)
                .flatMap(bytes -> Mono.fromRunnable(() -> write(destination, bytes))
                        .subscribeOn(Schedulers.boundedElastic())
                        .then());
    }

    private Mono<UploadReceipt> uploadEncrypted(
            String blobName, byte[] plaintext, WrappedDataKey dataKey) {
        byte[] plaintextKey = dataKey.plaintextKey();
        try {
            byte[] aad = EnvelopeMetadata.authenticatedData(blobName, dataKey.keyId());
            CipherSupport.EncryptedData encrypted =
                    CipherSupport.encrypt(plaintext, plaintextKey, aad);
            Map<String, String> metadata = EnvelopeMetadata.create(
                    dataKey.keyId(), dataKey.wrappedKey(), encrypted.iv());
            return blob(blobName)
                    .uploadWithResponse(
                            new BlobParallelUploadOptions(
                                    BinaryData.fromBytes(encrypted.ciphertext()))
                                    .setMetadata(metadata))
                    .map(ignored -> new UploadReceipt(dataKey.keyId(), dataKey.wrappedKey()))
                    .onErrorMap(AzureException.class,
                            e -> blobFailure("upload", blobName, e))
                    .doFinally(ignored -> Arrays.fill(plaintextKey, (byte) 0));
        } catch (RuntimeException e) {
            Arrays.fill(plaintextKey, (byte) 0);
            return Mono.error(e);
        }
    }

    private byte[] decrypt(
            String blobName,
            BinaryData content,
            EnvelopeMetadata.Parsed envelope,
            byte[] plaintextKey) {
        byte[] aad = EnvelopeMetadata.authenticatedData(blobName, envelope.keyId());
        return CipherSupport.decrypt(content.toBytes(), plaintextKey, envelope.iv(), aad);
    }

    private BlobAsyncClient blob(String blobName) {
        return containerClient.getBlobAsyncClient(blobName);
    }

    private static void write(Path destination, byte[] bytes) {
        try {
            Files.write(destination, bytes);
        } catch (IOException e) {
            throw new EnvelopeEncryptionException(
                    "Could not write decrypted file " + destination, e);
        }
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
