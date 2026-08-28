package com.example.encryptedblob;

import com.azure.core.util.BinaryData;
import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.models.BlobStorageException;
import com.azure.storage.blob.options.BlobParallelUploadOptions;
import reactor.core.publisher.Mono;
import reactor.core.scheduler.Schedulers;

import javax.crypto.Cipher;
import javax.crypto.spec.GCMParameterSpec;
import javax.crypto.spec.SecretKeySpec;
import java.nio.file.Files;
import java.nio.file.Path;
import java.security.GeneralSecurityException;
import java.security.SecureRandom;
import java.util.Arrays;
import java.util.Map;

public final class AsyncEncryptedBlobClient {
    private static final String CONTENT_ALGORITHM = "AES/GCM/NoPadding";
    private static final int GCM_TAG_BITS = 128;
    private static final int IV_BYTES = 12;

    private final BlobContainerAsyncClient containerClient;
    private final AsyncKeyManagement keyManagement;
    private final SecureRandom secureRandom;

    public AsyncEncryptedBlobClient(
        BlobContainerAsyncClient containerClient,
        AsyncKeyManagement keyManagement
    ) {
        this.containerClient = containerClient;
        this.keyManagement = keyManagement;
        this.secureRandom = new SecureRandom();
    }

    public Mono<EncryptedBlobClient.UploadResult> upload(Path source, String blobName) {
        return Mono.fromCallable(() -> Files.readAllBytes(source))
            .subscribeOn(Schedulers.boundedElastic())
            .onErrorMap(exception ->
                new EnvelopeEncryptionException("Could not read source file: " + source, exception))
            .flatMap(plaintext -> upload(plaintext, blobName));
    }

    public Mono<EncryptedBlobClient.UploadResult> upload(byte[] plaintext, String blobName) {
        return Mono.usingWhen(
                keyManagement.generateAndWrapKey(),
                envelope -> Mono.defer(() -> {
                    byte[] iv = new byte[IV_BYTES];
                    secureRandom.nextBytes(iv);
                    byte[] ciphertext = encrypt(plaintext, envelope.plaintextKey(), iv);
                    ProtectedDataKey protectedKey = envelope.protectedKey();
                    Map<String, String> metadata = BlobEncryptionMetadata.create(protectedKey, iv);

                    return containerClient.getBlobAsyncClient(blobName)
                        .uploadWithResponse(
                            new BlobParallelUploadOptions(BinaryData.fromBytes(ciphertext))
                                .setMetadata(metadata))
                        .map(ignored -> new EncryptedBlobClient.UploadResult(
                            protectedKey.keyId(),
                            protectedKey.wrappedKey()));
                }),
                envelope -> Mono.fromRunnable(envelope::close),
                (envelope, ignored) -> Mono.fromRunnable(envelope::close),
                envelope -> Mono.fromRunnable(envelope::close))
            .onErrorMap(BlobStorageException.class, exception ->
                blobFailure("upload", blobName, exception));
    }

    public Mono<byte[]> download(String blobName) {
        var blobClient = containerClient.getBlobAsyncClient(blobName);
        return blobClient.getProperties()
            .flatMap(properties -> {
                BlobEncryptionMetadata metadata = BlobEncryptionMetadata.parse(properties.getMetadata());
                return blobClient.downloadContent()
                    .flatMap(content -> keyManagement.unwrapKey(metadata.protectedKey())
                        .flatMap(dataKey -> Mono.using(
                            () -> dataKey,
                            key -> Mono.fromCallable(() ->
                                EncryptedBlobClient.decrypt(content.toBytes(), key, metadata.iv())),
                            key -> Arrays.fill(key, (byte) 0))));
            })
            .onErrorMap(BlobStorageException.class, exception ->
                blobFailure("download", blobName, exception));
    }

    public Mono<Void> download(String blobName, Path destination) {
        return download(blobName)
            .flatMap(plaintext -> Mono.<Void>fromRunnable(() -> {
                try {
                    Files.write(destination, plaintext);
                } catch (Exception exception) {
                    throw new EnvelopeEncryptionException(
                        "Could not write decrypted file: " + destination, exception);
                } finally {
                    Arrays.fill(plaintext, (byte) 0);
                }
            }).subscribeOn(Schedulers.boundedElastic()));
    }

    private static byte[] encrypt(byte[] plaintext, byte[] dataKey, byte[] iv) {
        try {
            Cipher cipher = Cipher.getInstance(CONTENT_ALGORITHM);
            cipher.init(Cipher.ENCRYPT_MODE, new SecretKeySpec(dataKey, "AES"), new GCMParameterSpec(GCM_TAG_BITS, iv));
            return cipher.doFinal(plaintext);
        } catch (GeneralSecurityException exception) {
            throw new EnvelopeEncryptionException("Local AES-GCM encryption failed.", exception);
        }
    }

    private static EnvelopeEncryptionException blobFailure(
        String operation,
        String blobName,
        BlobStorageException exception
    ) {
        String detail = exception.getStatusCode() == 404
            ? "The blob or container does not exist."
            : "Storage returned " + exception.getErrorCode() + ".";
        return new EnvelopeEncryptionException(
            "Could not " + operation + " blob '" + blobName + "' (HTTP "
                + exception.getStatusCode() + "). " + detail,
            exception);
    }
}
