package com.example.encryptedblob;

public final class EnvelopeEncryptionException extends RuntimeException {
    public EnvelopeEncryptionException(String message) {
        super(message);
    }

    public EnvelopeEncryptionException(String message, Throwable cause) {
        super(message, cause);
    }
}
