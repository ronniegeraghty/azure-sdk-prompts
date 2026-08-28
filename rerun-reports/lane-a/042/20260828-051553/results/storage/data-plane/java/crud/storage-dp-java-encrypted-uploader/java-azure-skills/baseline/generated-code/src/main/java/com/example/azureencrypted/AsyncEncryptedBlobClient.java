package com.example.azureencrypted;

import com.azure.core.util.BinaryData;
import com.azure.storage.blob.BlobAsyncClient;
import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.models.BlobStorageException;
import com.azure.storage.blob.options.BlobParallelUploadOptions;
import reactor.core.publisher.Mono;

import javax.crypto.Cipher;
import java.security.SecureRandom;
import java.util.Arrays;
import java.util.Map;

public final class AsyncEncryptedBlobClient {
    private static final int IV_BYTES = 12;

    private final BlobContainerAsyncClient containerClient;
    private final AsyncKeyManagement keyManagement;
    private final SecureRandom secureRandom;

    public AsyncEncryptedBlobClient(
            BlobContainerAsyncClient containerClient, AsyncKeyManagement keyManagement) {
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
        return Mono.defer(() -> {
            BlobAsyncClient blobClient = containerClient.getBlobAsyncClient(blobName);
            byte[] iv = new byte[IV_BYTES];
            secureRandom.nextBytes(iv);

            return Mono.usingWhen(
                    keyManagement.generateAndWrapDataKey(),
                    dataKey -> {
                        byte[] ciphertext = EncryptedBlobClient.crypt(
                                Cipher.ENCRYPT_MODE, plaintext, dataKey.plaintext(), iv);
                        Map<String, String> metadata = EncryptedBlobClient.metadata(
                                iv, dataKey.wrapped(), dataKey.keyId());
                        BlobParallelUploadOptions options =
                                new BlobParallelUploadOptions(BinaryData.fromBytes(ciphertext))
                                        .setMetadata(metadata);
                        return blobClient.uploadWithResponse(options)
                                .thenReturn(new EncryptedBlobClient.UploadResult(
                                        dataKey.keyId(), metadata.get("wrappeddek")));
                    },
                    dataKey -> Mono.fromRunnable(dataKey::close),
                    (dataKey, error) -> Mono.fromRunnable(dataKey::close),
                    dataKey -> Mono.fromRunnable(dataKey::close))
                    .onErrorMap(
                            BlobStorageException.class,
                            exception -> new EnvelopeEncryptionException(
                                    "Blob Storage could not upload blob " + blobName, exception));
        });
    }

    public Mono<byte[]> download(String blobName) {
        return Mono.defer(() -> {
            BlobAsyncClient blobClient = containerClient.getBlobAsyncClient(blobName);
            return blobClient.getProperties()
                    .flatMap(properties -> {
                        EncryptedBlobClient.EncryptionMetadata metadata =
                                EncryptedBlobClient.parseMetadata(properties.getMetadata(), blobName);
                        return blobClient.downloadContent()
                                .map(BinaryData::toBytes)
                                .flatMap(ciphertext -> keyManagement
                                        .unwrapDataKey(
                                                metadata.wrappedDataKey(), metadata.keyId())
                                        .map(dataKey -> {
                                            try {
                                                return EncryptedBlobClient.crypt(
                                                        Cipher.DECRYPT_MODE,
                                                        ciphertext,
                                                        dataKey,
                                                        metadata.iv());
                                            } finally {
                                                Arrays.fill(dataKey, (byte) 0);
                                            }
                                        }));
                    })
                    .onErrorMap(
                            error -> error instanceof BlobStorageException exception
                                    && exception.getStatusCode() == 404,
                            exception -> new EnvelopeEncryptionException(
                                    "Encrypted blob does not exist: " + blobName, exception))
                    .onErrorMap(
                            BlobStorageException.class,
                            exception -> new EnvelopeEncryptionException(
                                    "Blob Storage could not download blob " + blobName, exception));
        });
    }
}
