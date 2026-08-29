package com.example.azure.identity;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.identity.CredentialUnavailableException;

import java.util.ArrayList;
import java.util.List;
import java.util.Locale;

final class AuthenticationFailureAnalyzer {
    private AuthenticationFailureAnalyzer() {
    }

    static String describe(Throwable failure) {
        Throwable root = rootCause(failure);
        String details = collectMessages(failure);
        String normalized = details.toLowerCase(Locale.ROOT);

        String category;
        if (containsAny(normalized, "aadsts7000222", "expired client secret", "certificate has expired",
            "expired certificate", "certificate is not within its validity period")) {
            category = "The client secret or certificate has expired";
        } else if (containsAny(normalized, "aadsts500011", "aadsts90002", "tenant not found",
            "invalid tenant", "wrong tenant")) {
            category = "The configured tenant is invalid or does not contain the requested application/resource";
        } else if (containsAny(normalized, "aadsts7000215", "invalid client secret", "invalid_client")) {
            category = "The client credential is invalid";
        } else if (containsAny(normalized, "aadsts700016", "application with identifier")) {
            category = "The configured client ID was not found in the tenant";
        } else if (containsAny(normalized, "aadsts700211", "federated identity record")) {
            category = "No matching workload identity federation record was found";
        } else if (containsAny(normalized, "aadsts700024", "client assertion is not within")) {
            category = "The workload identity assertion is expired or not yet valid";
        } else if (containsAny(normalized, "unauthorized_client", "consent", "permission", "forbidden")) {
            category = "The identity lacks consent or permission for this token request";
        } else if (failure instanceof CredentialUnavailableException
            || containsAny(normalized, "credentialunavailableexception", "authentication unavailable",
                "no accounts were found", "cannot be established", "not installed", "not logged in")) {
            category = "No configured credential source or Azure identity is available";
        } else if (failure instanceof ClientAuthenticationException) {
            category = "Microsoft Entra ID rejected the authentication request";
        } else if (containsAny(normalized, "timeout", "timed out", "connection refused", "unknownhost")) {
            category = "The identity endpoint could not be reached";
        } else {
            category = root.getClass().getSimpleName();
        }

        return category + ": " + sanitize(root.getMessage());
    }

    private static Throwable rootCause(Throwable failure) {
        Throwable current = failure;
        while (current.getCause() != null && current.getCause() != current) {
            current = current.getCause();
        }
        return current;
    }

    private static String collectMessages(Throwable failure) {
        List<String> messages = new ArrayList<>();
        Throwable current = failure;
        while (current != null) {
            if (current.getMessage() != null) {
                messages.add(current.getMessage());
            }
            current = current.getCause();
        }
        return String.join(" | ", messages);
    }

    private static boolean containsAny(String value, String... candidates) {
        for (String candidate : candidates) {
            if (value.contains(candidate)) {
                return true;
            }
        }
        return false;
    }

    private static String sanitize(String message) {
        if (message == null || message.isBlank()) {
            return "No additional details were supplied.";
        }
        return message.replaceAll("[\\r\\n]+", " ").trim();
    }
}
