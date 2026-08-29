package com.example.encryptedblob;

import com.azure.core.util.BinaryData;
import com.azure.core.util.Context;
import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.models.BlobProperties;
import com.azure.storage.blob.models.BlobRequestConditions;
import com.azure.storage.blob.options.BlobParallelUploadOptions;

import javax.crypto.AEADBadTagException;
import javax.crypto.Cipher;
import javax.crypto.spec.GCMParameterSpec;
import javax.crypto.spec.SecretKeySpec;
import java.security.GeneralSecurityException;
import java.security.SecureRandom;
import java.util.Objects;

public final class EncryptedBlobStore {
    private static final String CIPHER = "AES/GCM/NoPadding";
    private static final int GCM_TAG_BITS = 128;
    private static final int IV_BYTES = 12;

    private final BlobClient blobClient;
    private final KeyManagementService keyManagement;
    private final SecureRandom secureRandom;

    public EncryptedBlobStore(BlobClient blobClient, KeyManagementService keyManagement) {
        this(blobClient, keyManagement, new SecureRandom());
    }

    EncryptedBlobStore(
        BlobClient blobClient,
        KeyManagementService keyManagement,
        SecureRandom secureRandom
    ) {
        this.blobClient = Objects.requireNonNull(blobClient, "blobClient");
        this.keyManagement = Objects.requireNonNull(keyManagement, "keyManagement");
        this.secureRandom = Objects.requireNonNull(secureRandom, "secureRandom");
    }

    public BlobEncryptionMetadata upload(byte[] plaintext, boolean overwrite) {
        Objects.requireNonNull(plaintext, "plaintext");
        try (GeneratedDataKey generatedKey = keyManagement.generateAndWrapDataKey()) {
            byte[] iv = new byte[IV_BYTES];
            secureRandom.nextBytes(iv);
            byte[] ciphertext = encrypt(plaintext, generatedKey.plaintextKey().bytes(), iv);
            WrappedDataKey wrappedKey = generatedKey.wrappedKey();
            BlobEncryptionMetadata metadata = new BlobEncryptionMetadata(
                wrappedKey.keyId(), wrappedKey.algorithm(), wrappedKey.bytes(), iv);

            BlobParallelUploadOptions options = new BlobParallelUploadOptions(BinaryData.fromBytes(ciphertext))
                .setMetadata(metadata.toMap());
            if (!overwrite) {
                options.setRequestConditions(new BlobRequestConditions().setIfNoneMatch("*"));
            }
            blobClient.uploadWithResponse(options, null, Context.NONE);
            return metadata;
        } catch (GeneralSecurityException exception) {
            throw new EncryptionStorageException("Local encryption failed", exception);
        } catch (RuntimeException exception) {
            throw serviceFailure("Encrypted blob upload failed", exception);
        }
    }

    public byte[] download() {
        try {
            BlobProperties properties = blobClient.getProperties();
            BlobEncryptionMetadata metadata = BlobEncryptionMetadata.fromMap(properties.getMetadata());
            BlobRequestConditions conditions =
                new BlobRequestConditions().setIfMatch(properties.getETag());
            byte[] ciphertext = blobClient
                .downloadContentWithResponse(null, conditions, null, Context.NONE)
                .getValue()
                .toBytes();

            WrappedDataKey wrappedKey = new WrappedDataKey(
                metadata.keyId(),
                metadata.wrapAlgorithm(),
                metadata.wrappedDataKey());
            try (DataKey dataKey = keyManagement.unwrapDataKey(wrappedKey)) {
                return decrypt(ciphertext, dataKey.bytes(), metadata.initializationVector());
            }
        } catch (AEADBadTagException exception) {
            throw new EncryptionStorageException(
                "Ciphertext authentication failed; the blob or metadata may have been modified",
                exception);
        } catch (GeneralSecurityException exception) {
            throw new EncryptionStorageException("Local decryption failed", exception);
        } catch (RuntimeException exception) {
            throw serviceFailure("Encrypted blob download failed", exception);
        }
    }

    private static byte[] encrypt(byte[] plaintext, byte[] key, byte[] iv)
        throws GeneralSecurityException {
        Cipher cipher = Cipher.getInstance(CIPHER);
        cipher.init(Cipher.ENCRYPT_MODE, new SecretKeySpec(key, "AES"), new GCMParameterSpec(GCM_TAG_BITS, iv));
        return cipher.doFinal(plaintext);
    }

    private static byte[] decrypt(byte[] ciphertext, byte[] key, byte[] iv)
        throws GeneralSecurityException {
        Cipher cipher = Cipher.getInstance(CIPHER);
        cipher.init(Cipher.DECRYPT_MODE, new SecretKeySpec(key, "AES"), new GCMParameterSpec(GCM_TAG_BITS, iv));
        return cipher.doFinal(ciphertext);
    }

    private static RuntimeException serviceFailure(String message, RuntimeException exception) {
        if (exception instanceof EncryptionStorageException) {
            return exception;
        }
        return new EncryptionStorageException(message, exception);
    }
}
