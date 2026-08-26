package com.example.azureauth;

import java.time.OffsetDateTime;

public record ConnectivityTestResult(
    boolean successful,
    OffsetDateTime expiresAt,
    boolean caeRequested,
    String failureReason
) {
    static ConnectivityTestResult success(OffsetDateTime expiresAt, boolean caeRequested) {
        return new ConnectivityTestResult(true, expiresAt, caeRequested, null);
    }

    static ConnectivityTestResult failure(boolean caeRequested, String failureReason) {
        return new ConnectivityTestResult(false, null, caeRequested, failureReason);
    }
}
