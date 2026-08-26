package com.example.appconfig;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.math.BigInteger;
import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.Locale;
import java.util.OptionalDouble;

final class FeatureFlagLogic {
    private static final ObjectMapper JSON = new ObjectMapper();
    private static final BigInteger TEN_THOUSAND = BigInteger.valueOf(10_000);

    private FeatureFlagLogic() {
    }

    static boolean evaluate(String flagName, String payload, String userId) {
        JsonNode flag;
        try {
            flag = JSON.readTree(payload);
        } catch (JsonProcessingException exception) {
            throw new IllegalArgumentException("Invalid JSON for feature flag " + flagName, exception);
        }

        if (!flag.path("enabled").asBoolean(false)) {
            return false;
        }

        OptionalDouble percentage = findPercentage(flag);
        if (percentage.isEmpty()) {
            return true;
        }
        if (userId == null || userId.isBlank()) {
            return false;
        }

        double value = percentage.getAsDouble();
        if (value < 0 || value > 100) {
            throw new IllegalArgumentException(
                "Percentage for feature flag " + flagName + " must be between 0 and 100");
        }
        return rolloutBucket(flagName, userId) < value * 100;
    }

    private static OptionalDouble findPercentage(JsonNode flag) {
        JsonNode filters = flag.path("conditions").path("client_filters");
        if (!filters.isArray()) {
            return OptionalDouble.empty();
        }

        for (JsonNode filter : filters) {
            String name = filter.path("name").asText("").toLowerCase(Locale.ROOT);
            if (name.endsWith("percentage")) {
                JsonNode parameters = filter.path("parameters");
                JsonNode value = parameters.has("Value")
                    ? parameters.get("Value")
                    : parameters.get("value");
                if (value == null || (!value.isNumber() && !value.isTextual())) {
                    throw new IllegalArgumentException("Percentage filter is missing its Value parameter");
                }
                try {
                    return OptionalDouble.of(value.isNumber()
                        ? value.asDouble()
                        : Double.parseDouble(value.asText()));
                } catch (NumberFormatException exception) {
                    throw new IllegalArgumentException("Percentage filter Value must be numeric", exception);
                }
            }
        }
        return OptionalDouble.empty();
    }

    private static int rolloutBucket(String flagName, String userId) {
        try {
            MessageDigest digest = MessageDigest.getInstance("SHA-256");
            byte[] hash = digest.digest((flagName + ":" + userId).getBytes(StandardCharsets.UTF_8));
            byte[] firstEightBytes = new byte[8];
            System.arraycopy(hash, 0, firstEightBytes, 0, firstEightBytes.length);
            return new BigInteger(1, firstEightBytes).mod(TEN_THOUSAND).intValue();
        } catch (NoSuchAlgorithmException exception) {
            throw new IllegalStateException("SHA-256 is unavailable", exception);
        }
    }
}
