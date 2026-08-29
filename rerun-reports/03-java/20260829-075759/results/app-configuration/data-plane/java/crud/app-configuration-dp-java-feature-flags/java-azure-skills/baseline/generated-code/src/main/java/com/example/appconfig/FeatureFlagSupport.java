package com.example.appconfig;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.math.BigDecimal;
import java.math.BigInteger;
import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.Iterator;

final class FeatureFlagSupport {
    static final String KEY_PREFIX = ".appconfig.featureflag/";
    private static final String PERCENTAGE_FILTER = "Microsoft.Percentage";
    private static final ObjectMapper OBJECT_MAPPER = new ObjectMapper();
    private static final BigDecimal ONE_HUNDRED = BigDecimal.valueOf(100);
    private static final BigInteger BUCKET_COUNT = BigInteger.valueOf(10_000);

    private FeatureFlagSupport() {
    }

    static boolean evaluate(String flagName, String payload, String userId) {
        final JsonNode flag;
        try {
            flag = OBJECT_MAPPER.readTree(payload);
        } catch (JsonProcessingException exception) {
            throw new IllegalArgumentException("Invalid JSON for feature flag " + flagName, exception);
        }

        if (!flag.path("enabled").asBoolean(false)) {
            return false;
        }

        JsonNode filters = flag.path("conditions").path("client_filters");
        if (!filters.isArray() || filters.isEmpty()) {
            return true;
        }

        Iterator<JsonNode> iterator = filters.elements();
        while (iterator.hasNext()) {
            JsonNode filter = iterator.next();
            String filterName = filter.path("name").asText();
            if (!PERCENTAGE_FILTER.equals(filterName)) {
                throw new UnsupportedOperationException("Unsupported feature filter: " + filterName);
            }
            if (userId == null || userId.isBlank()) {
                throw new IllegalArgumentException("A user ID is required for percentage rollout");
            }

            JsonNode valueNode = filter.path("parameters").path("Value");
            if (valueNode.isMissingNode()) {
                throw new IllegalArgumentException("Percentage filter is missing parameters.Value");
            }
            BigDecimal percentage;
            try {
                percentage = new BigDecimal(valueNode.asText());
            } catch (NumberFormatException exception) {
                throw new IllegalArgumentException("Invalid rollout percentage: " + valueNode.asText(), exception);
            }
            if (percentage.signum() < 0 || percentage.compareTo(ONE_HUNDRED) > 0) {
                throw new IllegalArgumentException("Rollout percentage must be between 0 and 100");
            }
            return bucket(flagName, userId) < percentage.multiply(BigDecimal.valueOf(100)).intValue();
        }
        return true;
    }

    static int bucket(String flagName, String userId) {
        try {
            MessageDigest digest = MessageDigest.getInstance("SHA-256");
            byte[] hash = digest.digest((flagName + ":" + userId).getBytes(StandardCharsets.UTF_8));
            return new BigInteger(1, hash).mod(BUCKET_COUNT).intValue();
        } catch (NoSuchAlgorithmException exception) {
            throw new IllegalStateException("SHA-256 is not available", exception);
        }
    }
}
