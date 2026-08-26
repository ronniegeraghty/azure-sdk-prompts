package com.example.encryptedblob;

import com.azure.storage.blob.models.BlobStorageException;

public final class BlobEncryptionException extends RuntimeException {
    public BlobEncryptionException(String message) {
        super(message);
    }

    public BlobEncryptionException(String message, Throwable cause) {
        super(message, cause);
    }

    static BlobEncryptionException storageFailure(
        String operation,
        String blobName,
        BlobStorageException cause
    ) {
        return new BlobEncryptionException(
            "Blob Storage could not " + operation + " '" + blobName + "' (HTTP "
                + cause.getStatusCode() + ", error " + cause.getErrorCode() + ")",
            cause
        );
    }
}
