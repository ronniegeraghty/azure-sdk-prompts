package com.example.azureauth;

import java.time.OffsetDateTime;

public record ConnectivityTestResult(
    boolean successful,
    OffsetDateTime expiresAt,
    boolean caeRequested,
    String failureReason
) {
    public static ConnectivityTestResult success(OffsetDateTime expiresAt, boolean caeRequested) {
        return new ConnectivityTestResult(true, expiresAt, caeRequested, null);
    }

    public static ConnectivityTestResult failure(boolean caeRequested, String reason) {
        return new ConnectivityTestResult(false, null, caeRequested, reason);
    }
}
