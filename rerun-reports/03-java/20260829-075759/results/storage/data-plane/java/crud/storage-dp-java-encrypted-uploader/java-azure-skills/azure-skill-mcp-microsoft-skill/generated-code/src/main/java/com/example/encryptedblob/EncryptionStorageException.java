package com.example.encryptedblob;

public final class EncryptionStorageException extends RuntimeException {
    public EncryptionStorageException(String message) {
        super(message);
    }

    public EncryptionStorageException(String message, Throwable cause) {
        super(message, cause);
    }
}
