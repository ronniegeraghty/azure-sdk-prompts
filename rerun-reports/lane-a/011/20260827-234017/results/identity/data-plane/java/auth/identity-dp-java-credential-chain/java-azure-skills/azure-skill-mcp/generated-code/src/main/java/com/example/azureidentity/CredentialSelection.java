package com.example.azureidentity;

import com.azure.core.credential.TokenCredential;

import java.util.Objects;

public record CredentialSelection(
    TokenCredential credential,
    DeploymentEnvironment environment,
    String strategy,
    boolean caeEnabled
) {
    public CredentialSelection {
        Objects.requireNonNull(credential, "credential");
        Objects.requireNonNull(environment, "environment");
        Objects.requireNonNull(strategy, "strategy");
    }
}
