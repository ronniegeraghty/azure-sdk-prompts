package com.example.azureidentity;

import java.time.OffsetDateTime;

public record ConnectivityResult(
    boolean successful,
    String scope,
    OffsetDateTime expiresAt,
    boolean caeEnabled,
    String failureReason
) {
    public static ConnectivityResult success(String scope, OffsetDateTime expiresAt, boolean caeEnabled) {
        return new ConnectivityResult(true, scope, expiresAt, caeEnabled, null);
    }

    public static ConnectivityResult failure(String scope, boolean caeEnabled, String reason) {
        return new ConnectivityResult(false, scope, null, caeEnabled, reason);
    }

    public void print(String testName) {
        System.out.println(testName + " connectivity test:");
        System.out.println("  Scope: " + scope);
        System.out.println("  CAE-enabled request: " + (caeEnabled ? "yes" : "no"));
        if (successful) {
            System.out.println("  Result: SUCCESS");
            System.out.println("  Token expires at: " + expiresAt);
        } else {
            System.out.println("  Result: FAILURE");
            System.out.println("  Reason: " + failureReason);
        }
    }
}
