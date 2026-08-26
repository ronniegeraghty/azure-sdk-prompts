package com.example.azureidentity;

import java.time.OffsetDateTime;

public record ConnectivityTestResult(
    boolean successful,
    OffsetDateTime expiresAt,
    boolean caeEnabled,
    String failureReason) {

    static ConnectivityTestResult success(OffsetDateTime expiresAt, boolean caeEnabled) {
        return new ConnectivityTestResult(true, expiresAt, caeEnabled, null);
    }

    static ConnectivityTestResult failure(boolean caeEnabled, Throwable error) {
        return new ConnectivityTestResult(
            false, null, caeEnabled, AuthenticationFailureAnalyzer.explain(error));
    }
}
