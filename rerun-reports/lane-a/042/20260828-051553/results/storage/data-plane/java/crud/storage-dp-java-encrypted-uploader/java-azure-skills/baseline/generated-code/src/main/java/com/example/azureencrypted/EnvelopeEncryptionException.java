package com.example.azureencrypted;

public final class EnvelopeEncryptionException extends RuntimeException {
    public EnvelopeEncryptionException(String message, Throwable cause) {
        super(message, cause);
    }

    public EnvelopeEncryptionException(String message) {
        super(message);
    }
}
