package com.example.azurecredentials;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.identity.CredentialUnavailableException;

import java.util.Locale;

final class AuthenticationFailureReporter {
    private AuthenticationFailureReporter() {
    }

    static String describe(Throwable failure) {
        Throwable root = unwrap(failure);
        String details = combinedMessages(failure);
        String normalized = details.toLowerCase(Locale.ROOT);

        if (containsAny(normalized, "expired", "aadsts7000222", "certificate has expired")) {
            return "The client secret or certificate has expired. " + details;
        }
        if (containsAny(normalized, "aadsts700016", "application with identifier", "wrong tenant",
            "tenant is not found", "aadsts90002")) {
            return "The application or tenant configuration is incorrect. " + details;
        }
        if (containsAny(normalized, "invalid_client", "aadsts7000215", "invalid client secret",
            "certificate validation")) {
            return "The client secret or certificate is invalid. " + details;
        }
        if (containsAny(normalized, "federated identity credential", "aadsts700211",
            "aadsts700212", "subject mismatch")) {
            return "The workload identity federation configuration does not match. " + details;
        }
        if (containsAny(normalized, "unauthorized", "forbidden", "permission", "consent")) {
            return "The identity lacks permission or consent for the requested scope. " + details;
        }
        if (root instanceof CredentialUnavailableException
            || containsAny(normalized, "credentialunavailableexception", "authentication unavailable",
                "no managed identity endpoint", "managed identity is not available",
                "no account", "not logged in")) {
            return "No usable identity was available. " + details;
        }
        if (root instanceof ClientAuthenticationException) {
            return "Azure rejected the authentication request. " + details;
        }
        return root.getClass().getSimpleName() + ": " + details;
    }

    private static Throwable unwrap(Throwable failure) {
        Throwable current = failure;
        while (current.getCause() != null && current.getCause() != current) {
            current = current.getCause();
        }
        return current;
    }

    private static String combinedMessages(Throwable failure) {
        StringBuilder messages = new StringBuilder();
        Throwable current = failure;
        while (current != null) {
            String message = current.getMessage();
            if (message != null && !message.isBlank()
                && messages.indexOf(message) < 0) {
                if (!messages.isEmpty()) {
                    messages.append(" | ");
                }
                messages.append(message.replaceAll("\\s+", " ").trim());
            }
            current = current.getCause();
        }
        return messages.isEmpty() ? "No additional details were supplied." : messages.toString();
    }

    private static boolean containsAny(String value, String... candidates) {
        for (String candidate : candidates) {
            if (value.contains(candidate)) {
                return true;
            }
        }
        return false;
    }
}
