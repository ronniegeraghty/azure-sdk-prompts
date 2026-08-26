package com.example.encryptedblob;

import com.azure.core.util.BinaryData;
import com.azure.storage.blob.BlobAsyncClient;
import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.models.BlobProperties;
import com.azure.storage.blob.models.BlobStorageException;
import com.azure.storage.blob.options.BlobParallelUploadOptions;
import reactor.core.publisher.Mono;

import java.security.SecureRandom;
import java.util.Base64;

public final class AsyncEncryptedBlobClient {
    private final BlobContainerAsyncClient containerClient;
    private final AsyncKeyManagement keyManagement;
    private final SecureRandom secureRandom;

    public AsyncEncryptedBlobClient(
            BlobContainerAsyncClient containerClient,
            AsyncKeyManagement keyManagement) {
        this(containerClient, keyManagement, new SecureRandom());
    }

    AsyncEncryptedBlobClient(
            BlobContainerAsyncClient containerClient,
            AsyncKeyManagement keyManagement,
            SecureRandom secureRandom) {
        this.containerClient = containerClient;
        this.keyManagement = keyManagement;
        this.secureRandom = secureRandom;
    }

    public Mono<EncryptedBlobClient.UploadResult> upload(String blobName, byte[] plaintext) {
        BlobAsyncClient blobClient = containerClient.getBlobAsyncClient(blobName);

        return Mono.using(
                keyManagement::generateDataKey,
                dataKey -> {
                    byte[] iv = BlobEncryption.newIv(secureRandom);
                    byte[] ciphertext = BlobEncryption.encrypt(plaintext, dataKey, iv);
                    return keyManagement.wrap(dataKey)
                            .flatMap(wrappedKey -> blobClient
                                    .uploadWithResponse(new BlobParallelUploadOptions(
                                                    BinaryData.fromBytes(ciphertext))
                                            .setMetadata(EncryptedBlobClient.metadata(
                                                    iv, wrappedKey, keyManagement.keyId())))
                                    .thenReturn(new EncryptedBlobClient.UploadResult(
                                            keyManagement.keyId(),
                                            Base64.getEncoder().encodeToString(wrappedKey))));
                },
                DataEncryptionKey::close)
                .onErrorMap(
                        BlobStorageException.class,
                        e -> blobFailure("upload encrypted blob", blobName, e));
    }

    public Mono<byte[]> download(String blobName) {
        BlobAsyncClient blobClient = containerClient.getBlobAsyncClient(blobName);
        Mono<BlobProperties> properties = blobClient.getProperties();
        Mono<byte[]> ciphertext = blobClient.downloadContent().map(BinaryData::toBytes);

        return Mono.zip(properties, ciphertext)
                .flatMap(tuple -> {
                    EncryptedBlobClient.EncryptionMetadata metadata =
                            EncryptedBlobClient.parseMetadata(
                                    tuple.getT1().getMetadata(), blobName);
                    return keyManagement.unwrap(metadata.wrappedKey(), metadata.keyId())
                            .map(dataKey -> {
                                try (dataKey) {
                                    return BlobEncryption.decrypt(
                                            tuple.getT2(), dataKey, metadata.iv());
                                }
                            });
                })
                .onErrorMap(
                        BlobStorageException.class,
                        e -> blobFailure("download encrypted blob", blobName, e));
    }

    private static EncryptedBlobException blobFailure(
            String operation, String blobName, BlobStorageException cause) {
        return new EncryptedBlobException(
                "Blob Storage could not " + operation + " '" + blobName
                        + "' (status " + cause.getStatusCode() + ", code "
                        + cause.getErrorCode() + ")",
                cause);
    }
}
