package com.example.azurecredentials;

import com.azure.core.credential.TokenCredential;

import java.util.Objects;

public record CredentialSelection(
    TokenCredential credential,
    String strategy,
    boolean caeEnabled
) {
    public CredentialSelection {
        Objects.requireNonNull(credential, "credential");
        Objects.requireNonNull(strategy, "strategy");
    }
}
