package com.example.encryptedblob;

public final class EncryptedBlobException extends RuntimeException {
    EncryptedBlobException(String message) {
        super(message);
    }

    EncryptedBlobException(String message, Throwable cause) {
        super(message, cause);
    }
}
