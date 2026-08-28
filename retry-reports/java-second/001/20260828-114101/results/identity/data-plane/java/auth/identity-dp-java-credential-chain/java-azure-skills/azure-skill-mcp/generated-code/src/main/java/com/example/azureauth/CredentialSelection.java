package com.example.azureauth;

import com.azure.core.credential.TokenCredential;

public record CredentialSelection(
    TokenCredential credential,
    String strategy,
    boolean caeEnabled
) {
}
