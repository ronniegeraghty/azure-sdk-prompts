package com.example.encryptedblob;

import com.azure.core.util.BinaryData;
import com.azure.storage.blob.BlobAsyncClient;
import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.models.BlobHttpHeaders;
import com.azure.storage.blob.models.BlobStorageException;
import com.azure.storage.blob.options.BlobParallelUploadOptions;
import reactor.core.publisher.Mono;
import reactor.core.scheduler.Schedulers;

import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Arrays;
import java.util.Objects;

public final class AsyncEncryptedBlobClient {
    private final BlobContainerAsyncClient containerClient;
    private final AsyncKeyManagementClient keyManagementClient;

    public AsyncEncryptedBlobClient(
        BlobContainerAsyncClient containerClient,
        AsyncKeyManagementClient keyManagementClient
    ) {
        this.containerClient = Objects.requireNonNull(containerClient, "containerClient");
        this.keyManagementClient =
            Objects.requireNonNull(keyManagementClient, "keyManagementClient");
    }

    public Mono<UploadResult> upload(Path source, String blobName) {
        return Mono.fromCallable(() -> Files.readAllBytes(source))
            .subscribeOn(Schedulers.boundedElastic())
            .flatMap(bytes -> upload(bytes, blobName));
    }

    public Mono<UploadResult> upload(byte[] plaintext, String blobName) {
        return Mono.defer(() -> {
            Objects.requireNonNull(plaintext, "plaintext");
            Objects.requireNonNull(blobName, "blobName");

            return keyManagementClient.generateAndWrapDataKey()
                .flatMap(envelopeKey -> Mono.using(
                    () -> envelopeKey,
                    key -> {
                        ProtectedDataKey protectedKey = key.protectedDataKey();
                        LocalAesGcm.EncryptedPayload encrypted = LocalAesGcm.encrypt(
                            key.dataKey(),
                            plaintext,
                            EncryptionMetadata.authenticatedData(protectedKey)
                        );
                        EncryptionMetadata metadata =
                            EncryptionMetadata.create(protectedKey, encrypted.iv());
                        BlobParallelUploadOptions options = new BlobParallelUploadOptions(
                            BinaryData.fromBytes(encrypted.ciphertext())
                        )
                            .setMetadata(metadata.toMap())
                            .setHeaders(
                                new BlobHttpHeaders().setContentType("application/octet-stream")
                            );

                        return containerClient.getBlobAsyncClient(blobName)
                            .uploadWithResponse(options)
                            .thenReturn(
                                new UploadResult(
                                    protectedKey.keyId(),
                                    protectedKey.wrappedKey()
                                )
                            );
                    },
                    EnvelopeKey::close
                ));
        }).onErrorMap(
            BlobStorageException.class,
            e -> BlobEncryptionException.storageFailure("upload", blobName, e)
        );
    }

    public Mono<Void> download(String blobName, Path destination) {
        return download(blobName)
            .flatMap(bytes -> Mono.fromCallable(() -> Files.write(destination, bytes)))
            .subscribeOn(Schedulers.boundedElastic())
            .then();
    }

    public Mono<byte[]> download(String blobName) {
        return Mono.defer(() -> {
            Objects.requireNonNull(blobName, "blobName");
            BlobAsyncClient blobClient = containerClient.getBlobAsyncClient(blobName);

            return blobClient.getProperties()
                .flatMap(properties -> {
                    EncryptionMetadata metadata =
                        EncryptionMetadata.parse(properties.getMetadata());
                    return blobClient.downloadContent()
                        .flatMap(content -> decrypt(content.toBytes(), metadata));
                });
        }).onErrorMap(
            BlobStorageException.class,
            e -> BlobEncryptionException.storageFailure("download", blobName, e)
        );
    }

    private Mono<byte[]> decrypt(byte[] ciphertext, EncryptionMetadata metadata) {
        return keyManagementClient.unwrapDataKey(metadata.protectedDataKey())
            .map(dataKey -> {
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
            });
    }
}
