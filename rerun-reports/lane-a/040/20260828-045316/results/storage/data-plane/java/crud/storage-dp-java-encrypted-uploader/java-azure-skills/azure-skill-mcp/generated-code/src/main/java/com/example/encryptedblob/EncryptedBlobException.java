package com.example.encryptedblob;

public final class EncryptedBlobException extends RuntimeException {
    public EncryptedBlobException(String message, Throwable cause) {
        super(message, cause);
    }

    public EncryptedBlobException(String message) {
        super(message);
    }
}
