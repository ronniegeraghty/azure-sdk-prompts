package com.example.encryptedblob;

import com.azure.core.util.BinaryData;
import com.azure.storage.blob.BlobAsyncClient;
import com.azure.storage.blob.models.BlobRequestConditions;
import com.azure.storage.blob.options.BlobParallelUploadOptions;
import reactor.core.publisher.Mono;

import javax.crypto.AEADBadTagException;
import javax.crypto.Cipher;
import javax.crypto.spec.GCMParameterSpec;
import javax.crypto.spec.SecretKeySpec;
import java.security.GeneralSecurityException;
import java.security.SecureRandom;
import java.util.Objects;

public final class AsyncEncryptedBlobStore {
    private static final String CIPHER = "AES/GCM/NoPadding";
    private static final int GCM_TAG_BITS = 128;
    private static final int IV_BYTES = 12;

    private final BlobAsyncClient blobClient;
    private final AsyncKeyManagementService keyManagement;
    private final SecureRandom secureRandom;

    public AsyncEncryptedBlobStore(
        BlobAsyncClient blobClient,
        AsyncKeyManagementService keyManagement
    ) {
        this(blobClient, keyManagement, new SecureRandom());
    }

    AsyncEncryptedBlobStore(
        BlobAsyncClient blobClient,
        AsyncKeyManagementService keyManagement,
        SecureRandom secureRandom
    ) {
        this.blobClient = Objects.requireNonNull(blobClient, "blobClient");
        this.keyManagement = Objects.requireNonNull(keyManagement, "keyManagement");
        this.secureRandom = Objects.requireNonNull(secureRandom, "secureRandom");
    }

    public Mono<BlobEncryptionMetadata> upload(byte[] plaintext, boolean overwrite) {
        Objects.requireNonNull(plaintext, "plaintext");
        return keyManagement.generateAndWrapDataKey()
            .flatMap(generatedKey -> Mono.using(
                () -> generatedKey,
                key -> encryptAndUpload(plaintext, key, overwrite),
                GeneratedDataKey::close))
            .onErrorMap(
                exception -> !(exception instanceof EncryptionStorageException),
                exception -> new EncryptionStorageException("Encrypted blob upload failed", exception));
    }

    public Mono<byte[]> download() {
        return blobClient.getProperties()
            .flatMap(properties -> {
                BlobEncryptionMetadata metadata =
                    BlobEncryptionMetadata.fromMap(properties.getMetadata());
                BlobRequestConditions conditions =
                    new BlobRequestConditions().setIfMatch(properties.getETag());
                return blobClient.downloadContentWithResponse(null, conditions)
                    .map(response -> new DownloadedCiphertext(
                        response.getValue().toBytes(),
                        metadata));
            })
            .flatMap(downloaded -> {
                BlobEncryptionMetadata metadata = downloaded.metadata();
                WrappedDataKey wrappedKey = new WrappedDataKey(
                    metadata.keyId(),
                    metadata.wrapAlgorithm(),
                    metadata.wrappedDataKey());
                return keyManagement.unwrapDataKey(wrappedKey)
                    .flatMap(dataKey -> Mono.using(
                        () -> dataKey,
                        key -> decrypt(downloaded.ciphertext(), key, metadata.initializationVector()),
                        DataKey::close));
            })
            .onErrorMap(
                exception -> !(exception instanceof EncryptionStorageException),
                exception -> new EncryptionStorageException("Encrypted blob download failed", exception));
    }

    private Mono<BlobEncryptionMetadata> encryptAndUpload(
        byte[] plaintext,
        GeneratedDataKey generatedKey,
        boolean overwrite
    ) {
        return Mono.fromCallable(() -> {
            byte[] iv = new byte[IV_BYTES];
            secureRandom.nextBytes(iv);
            byte[] ciphertext = encrypt(plaintext, generatedKey.plaintextKey().bytes(), iv);
            WrappedDataKey wrappedKey = generatedKey.wrappedKey();
            BlobEncryptionMetadata metadata = new BlobEncryptionMetadata(
                wrappedKey.keyId(), wrappedKey.algorithm(), wrappedKey.bytes(), iv);
            return new PendingUpload(ciphertext, metadata);
        }).flatMap(pending -> {
            BlobParallelUploadOptions options =
                new BlobParallelUploadOptions(BinaryData.fromBytes(pending.ciphertext()))
                    .setMetadata(pending.metadata().toMap());
            if (!overwrite) {
                options.setRequestConditions(new BlobRequestConditions().setIfNoneMatch("*"));
            }
            return blobClient.uploadWithResponse(options).thenReturn(pending.metadata());
        });
    }

    private static Mono<byte[]> decrypt(
        byte[] ciphertext,
        DataKey key,
        byte[] initializationVector
    ) {
        return Mono.fromCallable(() -> {
            try {
                Cipher cipher = Cipher.getInstance(CIPHER);
                cipher.init(
                    Cipher.DECRYPT_MODE,
                    new SecretKeySpec(key.bytes(), "AES"),
                    new GCMParameterSpec(GCM_TAG_BITS, initializationVector));
                return cipher.doFinal(ciphertext);
            } catch (AEADBadTagException exception) {
                throw new EncryptionStorageException(
                    "Ciphertext authentication failed; the blob or metadata may have been modified",
                    exception);
            } catch (GeneralSecurityException exception) {
                throw new EncryptionStorageException("Local decryption failed", exception);
            }
        });
    }

    private static byte[] encrypt(byte[] plaintext, byte[] key, byte[] iv)
        throws GeneralSecurityException {
        Cipher cipher = Cipher.getInstance(CIPHER);
        cipher.init(Cipher.ENCRYPT_MODE, new SecretKeySpec(key, "AES"), new GCMParameterSpec(GCM_TAG_BITS, iv));
        return cipher.doFinal(plaintext);
    }

    private record PendingUpload(byte[] ciphertext, BlobEncryptionMetadata metadata) {
    }

    private record DownloadedCiphertext(byte[] ciphertext, BlobEncryptionMetadata metadata) {
    }
}
