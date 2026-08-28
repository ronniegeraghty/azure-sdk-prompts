package com.example.azureidentity;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.identity.CredentialUnavailableException;

import java.util.Locale;

final class AuthenticationFailureAnalyzer {
    private AuthenticationFailureAnalyzer() {
    }

    static String describe(Throwable failure) {
        String details = collectMessages(failure);
        String normalized = details.toLowerCase(Locale.ROOT);

        if (contains(normalized, "aadsts7000222", "expired client secret", "certificate has expired",
            "expired certificate")) {
            return "The client secret or certificate has expired. " + details;
        }
        if (contains(normalized, "aadsts90002", "tenant not found", "invalid tenant",
            "aadsts700016", "application with identifier")) {
            return "The tenant ID is wrong, unavailable, or does not contain the configured application. " + details;
        }
        if (contains(normalized, "aadsts7000215", "invalid client secret")) {
            return "The client secret is invalid. " + details;
        }
        if (contains(normalized, "aadsts700027", "certificate", "client assertion is not within")) {
            return "The client certificate or assertion is invalid. " + details;
        }
        if (contains(normalized, "aadsts700024", "federated identity credential", "subject claim",
            "issuer claim")) {
            return "Workload identity federation is expired or does not match its federated credential. " + details;
        }
        if (contains(normalized, "aadsts500011", "invalid_resource", "resource principal")) {
            return "The requested Azure scope/resource is invalid for this tenant. " + details;
        }
        if (failure instanceof CredentialUnavailableException
            || contains(normalized, "credentialunavailableexception", "no managed identity",
                "managed identity endpoint", "identity not found", "no accounts were found")) {
            return "No configured identity is available in this environment. " + details;
        }
        if (failure instanceof ClientAuthenticationException) {
            return "Microsoft Entra ID rejected the credential. " + details;
        }
        return failure.getClass().getSimpleName() + ": " + details;
    }

    private static boolean contains(String value, String... candidates) {
        for (String candidate : candidates) {
            if (value.contains(candidate)) {
                return true;
            }
        }
        return false;
    }

    private static String collectMessages(Throwable failure) {
        StringBuilder messages = new StringBuilder();
        Throwable current = failure;
        while (current != null) {
            String message = current.getMessage();
            if (message != null && !message.isBlank() && messages.indexOf(message) < 0) {
                if (!messages.isEmpty()) {
                    messages.append(" Caused by: ");
                }
                messages.append(message.replaceAll("\\s+", " ").trim());
            }
            current = current.getCause();
        }
        return messages.isEmpty() ? "No diagnostic message was provided." : messages.toString();
    }
}
