package com.example.encryptedblob;

import com.azure.core.util.BinaryData;
import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.models.BlobRequestConditions;
import com.azure.storage.blob.models.BlobStorageException;
import com.azure.storage.blob.models.DownloadRetryOptions;
import com.azure.storage.blob.options.BlockBlobSimpleUploadOptions;
import reactor.core.publisher.Mono;
import reactor.core.scheduler.Schedulers;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Base64;

public final class EncryptedBlobAsyncClient {
    private final BlobContainerAsyncClient containerClient;
    private final AsyncKeyManagementClient keyManagementClient;
    private final EnvelopeEncryption encryption;

    public EncryptedBlobAsyncClient(
            BlobContainerAsyncClient containerClient,
            AsyncKeyManagementClient keyManagementClient) {
        this.containerClient = containerClient;
        this.keyManagementClient = keyManagementClient;
        this.encryption = new EnvelopeEncryption();
    }

    public Mono<UploadResult> upload(Path source, String blobName) {
        return Mono.fromCallable(() -> Files.readAllBytes(source))
                .subscribeOn(Schedulers.boundedElastic())
                .onErrorMap(
                        IOException.class,
                        exception -> new EncryptedBlobException(
                                "Could not read source file: " + source,
                                exception))
                .flatMap(bytes -> upload(bytes, blobName));
    }

    public Mono<UploadResult> upload(byte[] plaintext, String blobName) {
        return keyManagementClient.generateAndProtectDataKey()
                .flatMap(generatedKey -> Mono.using(
                        () -> generatedKey,
                        key -> {
                            EnvelopeEncryption.EncryptedPayload payload =
                                    encryption.encrypt(plaintext, key.plaintextKey(), blobName);
                            ProtectedDataKey protectedKey = key.protectedKey();
                            BlockBlobSimpleUploadOptions options =
                                    new BlockBlobSimpleUploadOptions(
                                            BinaryData.fromBytes(payload.ciphertext()))
                                            .setMetadata(encryption.metadata(payload, protectedKey));

                            return containerClient.getBlobAsyncClient(blobName)
                                    .getBlockBlobAsyncClient()
                                    .uploadWithResponse(options)
                                    .map(ignored -> new UploadResult(
                                            protectedKey.keyId(),
                                            Base64.getEncoder().encodeToString(
                                                    protectedKey.wrappedKey())));
                        },
                        GeneratedDataKey::close))
                .onErrorMap(
                        BlobStorageException.class,
                        exception -> new EncryptedBlobException(
                                "Blob Storage upload failed for '" + blobName + "': "
                                        + exception.getErrorCode(),
                                exception));
    }

    public Mono<byte[]> download(String blobName) {
        var blobClient = containerClient.getBlobAsyncClient(blobName).getBlockBlobAsyncClient();
        return blobClient.getProperties()
                .flatMap(properties -> {
                    var metadata = encryption.parseMetadata(properties.getMetadata());
                    BlobRequestConditions conditions = new BlobRequestConditions()
                            .setIfMatch(properties.getETag());
                    return blobClient.downloadContentWithResponse(
                                    new DownloadRetryOptions(),
                                    conditions)
                            .map(response -> response.getValue().toBytes())
                            .flatMap(ciphertext -> keyManagementClient
                                    .recoverDataKey(metadata.protectedKey())
                                    .flatMap(dataKey -> Mono.using(
                                            () -> dataKey,
                                            key -> Mono.fromCallable(() -> encryption.decrypt(
                                                    ciphertext,
                                                    metadata.iv(),
                                                    key,
                                                    blobName)),
                                            DataKeyMaterial::close)));
                })
                .onErrorMap(
                        BlobStorageException.class,
                        exception -> {
                            String detail = exception.getStatusCode() == 404
                                    ? "blob does not exist"
                                    : String.valueOf(exception.getErrorCode());
                            return new EncryptedBlobException(
                                    "Blob Storage download failed for '" + blobName + "': " + detail,
                                    exception);
                        });
    }

    public Mono<Void> download(String blobName, Path destination) {
        return download(blobName)
                .flatMap(bytes -> Mono.fromCallable(() -> {
                            Files.write(destination, bytes);
                            return destination;
                        })
                        .subscribeOn(Schedulers.boundedElastic())
                        .onErrorMap(
                                IOException.class,
                                exception -> new EncryptedBlobException(
                                        "Could not write decrypted file: " + destination,
                                        exception)))
                .then();
    }
}
