package com.example.azureauth;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.identity.CredentialUnavailableException;

import java.util.Locale;

public final class AuthenticationFailureAnalyzer {
    private AuthenticationFailureAnalyzer() {
    }

    public static String explain(Throwable failure) {
        String details = collectMessages(failure);
        String normalized = details.toLowerCase(Locale.ROOT);

        if (containsAny(normalized, "aadsts7000222", "expired certificate", "certificate has expired")) {
            return "The client certificate or secret has expired. Rotate it and update the credential source.";
        }
        if (containsAny(normalized, "aadsts7000215", "invalid client secret")) {
            return "The client secret is invalid. Check that the secret value, not its identifier, was supplied.";
        }
        if (containsAny(normalized, "aadsts700027", "certificate was not found", "invalid certificate")) {
            return "The client certificate is invalid or is not registered on the application.";
        }
        if (containsAny(normalized, "aadsts90002", "tenant not found")) {
            return "The configured Microsoft Entra tenant does not exist or cannot be reached.";
        }
        if (containsAny(normalized, "aadsts700016", "application with identifier")) {
            return "The client application was not found in the configured tenant; verify client and tenant IDs.";
        }
        if (containsAny(normalized, "managed identity", "identity not found", "no identity has been assigned")) {
            return "No usable managed identity was found; assign one or verify its client ID and endpoint.";
        }
        if (containsAny(normalized, "federated", "subject claim", "token file")) {
            return "Workload identity federation is misconfigured; verify tenant, client, subject, and token file.";
        }
        if (failure instanceof CredentialUnavailableException) {
            return "No credential in the selected chain is configured or available.";
        }
        if (failure instanceof ClientAuthenticationException) {
            return "Microsoft Entra ID rejected the credential. SDK details: " + firstLine(details);
        }
        return "Token acquisition failed. SDK details: " + firstLine(details);
    }

    private static String collectMessages(Throwable failure) {
        StringBuilder messages = new StringBuilder();
        Throwable current = failure;
        while (current != null) {
            if (current.getMessage() != null && !current.getMessage().isBlank()) {
                if (!messages.isEmpty()) {
                    messages.append(" | ");
                }
                messages.append(current.getMessage());
            }
            current = current.getCause();
        }
        return messages.isEmpty() ? failure.getClass().getSimpleName() : messages.toString();
    }

    private static boolean containsAny(String value, String... fragments) {
        for (String fragment : fragments) {
            if (value.contains(fragment)) {
                return true;
            }
        }
        return false;
    }

    private static String firstLine(String value) {
        int lineBreak = value.indexOf('\n');
        return lineBreak >= 0 ? value.substring(0, lineBreak) : value;
    }
}
