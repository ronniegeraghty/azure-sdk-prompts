package com.example.azureidentity;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.identity.CredentialUnavailableException;

import java.util.Locale;

final class AuthenticationFailureAnalyzer {
    private AuthenticationFailureAnalyzer() {
    }

    static String explain(Throwable error) {
        Throwable root = rootCause(error);
        String detail = firstMessage(error, root);
        String normalized = detail.toLowerCase(Locale.ROOT);

        if (containsAny(normalized, "expired certificate", "certificate has expired",
            "aadsts7000222", "client secret is expired", "credential has expired")) {
            return "The client certificate or secret has expired. " + detail;
        }
        if (containsAny(normalized, "aadsts90002", "tenant not found", "invalid tenant",
            "tenant does not exist")) {
            return "The configured Microsoft Entra tenant is invalid or unavailable. " + detail;
        }
        if (containsAny(normalized, "aadsts700016", "application with identifier",
            "was not found in the directory")) {
            return "The client ID is not registered in the configured tenant; check AZURE_CLIENT_ID "
                + "and AZURE_TENANT_ID. " + detail;
        }
        if (containsAny(normalized, "aadsts7000215", "invalid client secret",
            "invalid_client")) {
            return "The client secret or certificate is invalid. " + detail;
        }
        if (containsAny(normalized, "aadsts500011", "invalid scope", "resource principal")) {
            return "The requested Azure scope/resource is invalid for this tenant. " + detail;
        }
        if (containsAny(normalized, "managed identity", "identity not found", "no identity",
            "credential unavailable", "unavailable", "imds endpoint")) {
            return "No configured managed, workload, pipeline, or developer identity was available. "
                + detail;
        }
        if (containsAny(normalized, "federated token", "oidc", "token file")) {
            return "The workload or pipeline federated-identity configuration is invalid. " + detail;
        }
        if (containsAny(normalized, "timeout", "timed out", "connection", "dns",
            "unknown host", "network")) {
            return "The Microsoft Entra authentication endpoint could not be reached. " + detail;
        }
        if (error instanceof CredentialUnavailableException) {
            return "No credential in the selected chain could authenticate. " + detail;
        }
        if (error instanceof ClientAuthenticationException) {
            return "Microsoft Entra ID rejected the credential. " + detail;
        }
        return error.getClass().getSimpleName() + ": " + detail;
    }

    private static Throwable rootCause(Throwable error) {
        Throwable current = error;
        while (current.getCause() != null && current.getCause() != current) {
            current = current.getCause();
        }
        return current;
    }

    private static String firstMessage(Throwable error, Throwable root) {
        if (error.getMessage() != null && !error.getMessage().isBlank()) {
            return error.getMessage();
        }
        if (root.getMessage() != null && !root.getMessage().isBlank()) {
            return root.getMessage();
        }
        return root.getClass().getSimpleName();
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
