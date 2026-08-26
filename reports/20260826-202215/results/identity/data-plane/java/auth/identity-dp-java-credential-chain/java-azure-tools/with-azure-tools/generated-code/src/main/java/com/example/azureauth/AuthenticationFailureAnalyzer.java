package com.example.azureauth;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.identity.CredentialUnavailableException;

import java.util.Locale;
import java.util.concurrent.CompletionException;
import java.util.concurrent.ExecutionException;

final class AuthenticationFailureAnalyzer {
    private AuthenticationFailureAnalyzer() {
    }

    static String describe(Throwable failure) {
        Throwable cause = unwrap(failure);
        String messages = collectMessages(cause);
        String normalized = messages.toLowerCase(Locale.ROOT);

        String reason;
        if (containsAny(normalized, "aadsts7000222", "client secret keys for app") && normalized.contains("expired")) {
            reason = "The client secret has expired.";
        } else if (containsAny(normalized, "aadsts7000215", "invalid client secret")) {
            reason = "The client secret is invalid.";
        } else if (containsAny(normalized, "certificate") && containsAny(normalized, "expired", "not yet valid", "validity period")) {
            reason = "The client certificate is expired or not yet valid.";
        } else if (containsAny(normalized, "aadsts90002", "tenant not found", "invalid tenant", "wrong tenant")) {
            reason = "The tenant ID is wrong or the tenant cannot be found.";
        } else if (containsAny(normalized, "aadsts700016", "application with identifier") && normalized.contains("not found")) {
            reason = "The client ID is wrong or the application is not registered in this tenant.";
        } else if (containsAny(normalized, "federated", "subject mismatch", "issuer mismatch")) {
            reason = "Workload identity federation is misconfigured or its projected token is invalid.";
        } else if (containsAny(normalized, "no identity", "identity not found", "no managed identity",
            "managed identity is not available", "imds endpoint cannot be established")) {
            reason = "No usable managed identity is available on this host.";
        } else if (cause instanceof CredentialUnavailableException) {
            reason = "No credential in the selected chain is configured and available.";
        } else if (containsAny(normalized, "aadsts70011", "invalid scope")) {
            reason = "The requested Azure scope is invalid.";
        } else if (containsAny(normalized, "unknownhost", "connection refused", "connect timed out", "connection timeout")) {
            reason = "The identity endpoint could not be reached; check DNS, proxy, and network access.";
        } else if (cause instanceof ClientAuthenticationException) {
            reason = "Microsoft Entra ID rejected the credential.";
        } else {
            reason = "Token acquisition failed unexpectedly.";
        }

        String detail = firstNonBlankMessage(cause);
        return detail.isEmpty() ? reason : reason + " SDK detail: " + abbreviate(detail, 500);
    }

    private static Throwable unwrap(Throwable failure) {
        Throwable current = failure;
        while ((current instanceof CompletionException || current instanceof ExecutionException)
            && current.getCause() != null) {
            current = current.getCause();
        }
        return current;
    }

    private static String collectMessages(Throwable failure) {
        StringBuilder result = new StringBuilder();
        Throwable current = failure;
        while (current != null) {
            if (current.getMessage() != null) {
                result.append(' ').append(current.getMessage());
            }
            current = current.getCause();
        }
        return result.toString();
    }

    private static String firstNonBlankMessage(Throwable failure) {
        Throwable current = failure;
        while (current != null) {
            if (current.getMessage() != null && !current.getMessage().isBlank()) {
                return current.getMessage().replaceAll("\\s+", " ").trim();
            }
            current = current.getCause();
        }
        return "";
    }

    private static boolean containsAny(String value, String... candidates) {
        for (String candidate : candidates) {
            if (value.contains(candidate)) {
                return true;
            }
        }
        return false;
    }

    private static String abbreviate(String value, int maxLength) {
        return value.length() <= maxLength ? value : value.substring(0, maxLength - 3) + "...";
    }
}
