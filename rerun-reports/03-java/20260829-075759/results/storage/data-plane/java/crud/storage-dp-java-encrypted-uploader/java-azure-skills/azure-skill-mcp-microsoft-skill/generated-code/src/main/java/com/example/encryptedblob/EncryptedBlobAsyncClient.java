package com.example.encryptedblob;

import com.azure.core.util.BinaryData;
import com.azure.storage.blob.BlobAsyncClient;
import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.models.BlobStorageException;
import com.azure.storage.blob.options.BlobParallelUploadOptions;
import reactor.core.publisher.Mono;
import reactor.core.scheduler.Schedulers;

import java.nio.file.Files;
import java.nio.file.Path;
import java.security.SecureRandom;
import java.util.Base64;

public final class EncryptedBlobAsyncClient {
    private final BlobContainerAsyncClient containerClient;
    private final AsyncKeyManagementService keyManagement;
    private final SecureRandom secureRandom;

    public EncryptedBlobAsyncClient(
            BlobContainerAsyncClient containerClient,
            AsyncKeyManagementService keyManagement) {
        this(containerClient, keyManagement, new SecureRandom());
    }

    EncryptedBlobAsyncClient(
            BlobContainerAsyncClient containerClient,
            AsyncKeyManagementService keyManagement,
            SecureRandom secureRandom) {
        this.containerClient = containerClient;
        this.keyManagement = keyManagement;
        this.secureRandom = secureRandom;
    }

    public Mono<EncryptedBlobInfo> upload(String blobName, byte[] plaintext) {
        return keyManagement.generateAndWrapDataKey()
                .flatMap(protectedKey -> Mono.using(
                        () -> protectedKey,
                        key -> {
                            byte[] iv = EnvelopeCrypto.generateIv(secureRandom);
                            BlobEncryptionMetadata metadata = BlobEncryptionMetadata.create(
                                    key.keyId(), key.wrappedKey(), iv);
                            byte[] ciphertext = EnvelopeCrypto.encrypt(
                                    plaintext,
                                    key.dataKey().bytes(),
                                    iv,
                                    metadata.authenticatedMetadata());
                            BlobAsyncClient blobClient = containerClient.getBlobAsyncClient(blobName);
                            BlobParallelUploadOptions options =
                                    new BlobParallelUploadOptions(BinaryData.fromBytes(ciphertext))
                                            .setMetadata(metadata.toMap());
                            return blobClient.uploadWithResponse(options)
                                    .thenReturn(new EncryptedBlobInfo(
                                            key.keyId(),
                                            Base64.getEncoder().encodeToString(key.wrappedKey())));
                        },
                        ProtectedDataKey::close))
                .onErrorMap(
                        BlobStorageException.class,
                        exception -> blobException("upload", blobName, exception));
    }

    public Mono<EncryptedBlobInfo> uploadFile(String blobName, Path source) {
        return Mono.fromCallable(() -> Files.readAllBytes(source))
                .subscribeOn(Schedulers.boundedElastic())
                .onErrorMap(
                        exception -> new EncryptionStorageException(
                                "Could not read source file: " + source, exception))
                .flatMap(bytes -> upload(blobName, bytes));
    }

    public Mono<byte[]> download(String blobName) {
        BlobAsyncClient blobClient = containerClient.getBlobAsyncClient(blobName);
        return blobClient.getProperties()
                .map(properties -> BlobEncryptionMetadata.parse(properties.getMetadata()))
                .zipWith(blobClient.downloadContent())
                .flatMap(tuple -> {
                    BlobEncryptionMetadata metadata = tuple.getT1();
                    byte[] ciphertext = tuple.getT2().toBytes();
                    return keyManagement.unwrapDataKey(metadata.keyId(), metadata.wrappedKey())
                            .map(dataKey -> {
                                try (dataKey) {
                                    return EnvelopeCrypto.decrypt(
                                            ciphertext,
                                            dataKey.bytes(),
                                            metadata.iv(),
                                            metadata.authenticatedMetadata());
                                }
                            });
                })
                .onErrorMap(
                        BlobStorageException.class,
                        exception -> blobException("download", blobName, exception));
    }

    public Mono<Void> downloadFile(String blobName, Path destination) {
        return download(blobName)
                .flatMap(bytes -> Mono.fromCallable(() -> Files.write(destination, bytes))
                        .subscribeOn(Schedulers.boundedElastic())
                        .onErrorMap(
                                exception -> new EncryptionStorageException(
                                        "Could not write destination file: " + destination, exception)))
                .then();
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
