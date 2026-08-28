package com.example.azureauth;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.identity.CredentialUnavailableException;

import java.util.ArrayList;
import java.util.List;
import java.util.Locale;

final class AuthenticationFailureReporter {
    private AuthenticationFailureReporter() {
    }

    static String describe(Throwable failure) {
        List<String> messages = new ArrayList<>();
        Throwable current = failure;
        while (current != null) {
            if (current.getMessage() != null && !current.getMessage().isBlank()) {
                messages.add(current.getMessage());
            }
            current = current.getCause();
        }

        String combined = String.join(" | ", messages);
        String normalized = combined.toLowerCase(Locale.ROOT);
        String reason;

        if (normalized.contains("aadsts7000222")
            || (normalized.contains("certificate") && normalized.contains("expired"))) {
            reason = "The client secret or certificate has expired.";
        } else if (normalized.contains("aadsts90002")
            || normalized.contains("tenant not found")
            || normalized.contains("wrong tenant")) {
            reason = "The tenant ID is wrong or the tenant cannot be reached.";
        } else if (normalized.contains("aadsts7000215")
            || normalized.contains("invalid client secret")) {
            reason = "The client secret is invalid.";
        } else if (normalized.contains("aadsts700016")
            || normalized.contains("application with identifier")) {
            reason = "The client ID is wrong or the application is not registered in this tenant.";
        } else if (normalized.contains("federated")
            || normalized.contains("subject claim")
            || normalized.contains("token file")) {
            reason = "The workload identity token or federated identity configuration is invalid.";
        } else if (failure instanceof CredentialUnavailableException
            || normalized.contains("credentialunavailable")
            || normalized.contains("no managed identity")
            || normalized.contains("identity not found")
            || normalized.contains("imds endpoint")) {
            reason = "No usable identity is available in this environment.";
        } else if (normalized.contains("unauthorized_client")
            || normalized.contains("access_denied")
            || normalized.contains("forbidden")) {
            reason = "The identity is not authorized for this token request.";
        } else if (failure instanceof ClientAuthenticationException) {
            reason = "Microsoft Entra ID rejected the authentication request.";
        } else {
            reason = "Token acquisition failed before authentication completed.";
        }

        return reason + System.lineSeparator()
            + "  SDK exception: " + failure.getClass().getSimpleName() + System.lineSeparator()
            + "  Detail: " + sanitize(combined);
    }

    private static String sanitize(String message) {
        if (message == null || message.isBlank()) {
            return "(no detail supplied by the credential)";
        }
        String oneLine = message.replaceAll("\\s+", " ").trim();
        return oneLine.length() <= 600 ? oneLine : oneLine.substring(0, 600) + "...";
    }
}
