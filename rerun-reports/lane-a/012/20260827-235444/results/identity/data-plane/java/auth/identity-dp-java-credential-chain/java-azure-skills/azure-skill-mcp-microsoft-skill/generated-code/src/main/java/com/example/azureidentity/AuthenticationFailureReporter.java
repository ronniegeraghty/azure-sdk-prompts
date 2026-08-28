package com.example.azureidentity;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.identity.CredentialUnavailableException;

import java.util.Locale;

final class AuthenticationFailureReporter {
    private AuthenticationFailureReporter() {
    }

    static String describe(Throwable failure) {
        String details = collectMessages(failure);
        String normalized = details.toLowerCase(Locale.ROOT);

        if (failure instanceof CredentialUnavailableException
            || normalized.contains("credentialunavailableexception")
            || normalized.contains("no credential")) {
            return "No identity is available for this credential strategy. " + details;
        }
        if (normalized.contains("aadsts7000222")
            || normalized.contains("client secret") && normalized.contains("expired")) {
            return "The service principal client secret has expired. " + details;
        }
        if (normalized.contains("certificate") && normalized.contains("expired")) {
            return "The client certificate has expired. " + details;
        }
        if (normalized.contains("aadsts90002")
            || normalized.contains("tenant") && normalized.contains("not found")) {
            return "The configured tenant does not exist or is not accessible. " + details;
        }
        if (normalized.contains("aadsts700016")
            || normalized.contains("application") && normalized.contains("not found")) {
            return "The client ID is not registered in the configured tenant (possibly the wrong tenant). " + details;
        }
        if (normalized.contains("aadsts7000215") || normalized.contains("invalid client secret")) {
            return "The configured client secret is invalid. " + details;
        }
        if (normalized.contains("aadsts70011") || normalized.contains("invalid scope")) {
            return "The requested scope is invalid for this identity provider. " + details;
        }
        if (normalized.contains("timeout")
            || normalized.contains("connection")
            || normalized.contains("unknownhost")) {
            return "The identity endpoint could not be reached. " + details;
        }
        if (failure instanceof ClientAuthenticationException) {
            return "Microsoft Entra ID rejected the authentication request. " + details;
        }
        return "Credential configuration or token acquisition failed. " + details;
    }

    private static String collectMessages(Throwable failure) {
        StringBuilder messages = new StringBuilder();
        Throwable current = failure;
        int depth = 0;
        while (current != null && depth++ < 8) {
            String message = current.getMessage();
            if (message != null && !message.isBlank()) {
                if (!messages.isEmpty()) {
                    messages.append(" Caused by: ");
                }
                messages.append(message.replaceAll("\\s+", " ").trim());
            }
            current = current.getCause();
        }
        return messages.isEmpty() ? failure.getClass().getSimpleName() : messages.toString();
    }
}
