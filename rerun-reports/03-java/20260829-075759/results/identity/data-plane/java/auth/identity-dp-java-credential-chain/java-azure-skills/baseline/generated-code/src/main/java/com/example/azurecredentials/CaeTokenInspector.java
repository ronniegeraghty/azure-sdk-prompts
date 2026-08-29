package com.example.azurecredentials;

import com.azure.core.credential.AccessToken;

import java.nio.charset.StandardCharsets;
import java.util.Base64;
import java.util.regex.Pattern;

final class CaeTokenInspector {
    private static final Pattern CAE_CLAIM = Pattern.compile(
        "\"xms_cc\"\\s*:\\s*\\[[^]]*\"cp1\"",
        Pattern.CASE_INSENSITIVE);

    private CaeTokenInspector() {
    }

    static String status(AccessToken token, boolean requested) {
        if (!requested) {
            return "disabled (not requested)";
        }

        String tokenValue = token.getToken();
        String[] segments = tokenValue.split("\\.");
        if (segments.length < 2) {
            return "requested; not inspectable (token is not a JWT)";
        }

        try {
            String payload = new String(
                Base64.getUrlDecoder().decode(segments[1]),
                StandardCharsets.UTF_8);
            return CAE_CLAIM.matcher(payload).find()
                ? "enabled (xms_cc/cp1 claim present)"
                : "requested; token does not advertise xms_cc/cp1";
        } catch (IllegalArgumentException exception) {
            return "requested; not inspectable (invalid JWT payload)";
        }
    }
}
