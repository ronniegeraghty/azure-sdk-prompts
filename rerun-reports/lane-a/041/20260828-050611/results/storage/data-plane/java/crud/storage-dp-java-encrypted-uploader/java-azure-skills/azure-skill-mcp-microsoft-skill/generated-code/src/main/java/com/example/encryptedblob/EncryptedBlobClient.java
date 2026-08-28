package com.example.encryptedblob;

import com.azure.core.util.BinaryData;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.models.BlobStorageException;
import com.azure.storage.blob.options.BlobParallelUploadOptions;

import javax.crypto.AEADBadTagException;
import javax.crypto.Cipher;
import javax.crypto.spec.GCMParameterSpec;
import javax.crypto.spec.SecretKeySpec;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.security.GeneralSecurityException;
import java.security.SecureRandom;
import java.util.Arrays;
import java.util.Base64;
import java.util.Map;

public final class EncryptedBlobClient {
    private static final String CONTENT_ALGORITHM = "AES/GCM/NoPadding";
    private static final int GCM_TAG_BITS = 128;
    private static final int IV_BYTES = 12;

    private final BlobContainerClient containerClient;
    private final KeyManagement keyManagement;
    private final SecureRandom secureRandom;

    public EncryptedBlobClient(BlobContainerClient containerClient, KeyManagement keyManagement) {
        this.containerClient = containerClient;
        this.keyManagement = keyManagement;
        this.secureRandom = new SecureRandom();
    }

    public UploadResult upload(Path source, String blobName) {
        try {
            return upload(Files.readAllBytes(source), blobName);
        } catch (IOException exception) {
            throw new EnvelopeEncryptionException("Could not read source file: " + source, exception);
        }
    }

    public UploadResult upload(byte[] plaintext, String blobName) {
        try (DataKeyEnvelope envelope = keyManagement.generateAndWrapKey()) {
            byte[] iv = new byte[IV_BYTES];
            secureRandom.nextBytes(iv);
            byte[] ciphertext = encrypt(plaintext, envelope.plaintextKey(), iv);
            ProtectedDataKey protectedKey = envelope.protectedKey();

            Map<String, String> metadata = BlobEncryptionMetadata.create(protectedKey, iv);
            containerClient.getBlobClient(blobName).uploadWithResponse(
                new BlobParallelUploadOptions(BinaryData.fromBytes(ciphertext))
                    .setMetadata(metadata),
                null,
                null);

            return new UploadResult(protectedKey.keyId(), protectedKey.wrappedKey());
        } catch (BlobStorageException exception) {
            throw blobFailure("upload", blobName, exception);
        }
    }

    public byte[] download(String blobName) {
        try {
            var blobClient = containerClient.getBlobClient(blobName);
            Map<String, String> metadata = blobClient.getProperties().getMetadata();
            BlobEncryptionMetadata encryptionMetadata = BlobEncryptionMetadata.parse(metadata);
            byte[] ciphertext = blobClient.downloadContent().toBytes();
            byte[] dataKey = keyManagement.unwrapKey(encryptionMetadata.protectedKey());

            try {
                return decrypt(ciphertext, dataKey, encryptionMetadata.iv());
            } finally {
                Arrays.fill(dataKey, (byte) 0);
            }
        } catch (BlobStorageException exception) {
            throw blobFailure("download", blobName, exception);
        }
    }

    public void download(String blobName, Path destination) {
        byte[] plaintext = download(blobName);
        try {
            Files.write(destination, plaintext);
        } catch (IOException exception) {
            throw new EnvelopeEncryptionException("Could not write decrypted file: " + destination, exception);
        } finally {
            Arrays.fill(plaintext, (byte) 0);
        }
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

    static byte[] decrypt(byte[] ciphertext, byte[] dataKey, byte[] iv) {
        try {
            Cipher cipher = Cipher.getInstance(CONTENT_ALGORITHM);
            cipher.init(Cipher.DECRYPT_MODE, new SecretKeySpec(dataKey, "AES"), new GCMParameterSpec(GCM_TAG_BITS, iv));
            return cipher.doFinal(ciphertext);
        } catch (AEADBadTagException exception) {
            throw new EnvelopeEncryptionException(
                "Ciphertext authentication failed; the blob data or encryption metadata was modified.", exception);
        } catch (GeneralSecurityException exception) {
            throw new EnvelopeEncryptionException("Local AES-GCM decryption failed.", exception);
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

    public record UploadResult(String keyId, byte[] wrappedKey) {
        public UploadResult {
            wrappedKey = Arrays.copyOf(wrappedKey, wrappedKey.length);
        }

        @Override
        public byte[] wrappedKey() {
            return Arrays.copyOf(wrappedKey, wrappedKey.length);
        }

        public String wrappedKeyBase64() {
            return Base64.getEncoder().encodeToString(wrappedKey);
        }
    }
}
