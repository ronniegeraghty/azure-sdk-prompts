package com.example.azureauth;

import com.azure.core.credential.AccessToken;

import java.nio.charset.StandardCharsets;
import java.util.Base64;

final class TokenDetails {
    private TokenDetails() {
    }

    static String caeStatus(AccessToken token, boolean caeRequested) {
        if (!caeRequested) {
            return "not requested";
        }

        String[] segments = token.getToken().split("\\.");
        if (segments.length < 2) {
            return "requested; token is opaque, so the cp1 claim cannot be inspected";
        }

        try {
            String claims = new String(
                Base64.getUrlDecoder().decode(segments[1]),
                StandardCharsets.UTF_8
            );
            boolean hasCaeCapability = claims.matches(
                "(?s).*\"xms_cc\"\\s*:\\s*\\{.*?\"values\"\\s*:\\s*\\[.*?\"cp1\".*?].*?}.*"
            );
            return hasCaeCapability
                ? "enabled (cp1 capability present)"
                : "requested; cp1 capability is not present in the returned token";
        } catch (IllegalArgumentException exception) {
            return "requested; token claims could not be decoded";
        }
    }
}
