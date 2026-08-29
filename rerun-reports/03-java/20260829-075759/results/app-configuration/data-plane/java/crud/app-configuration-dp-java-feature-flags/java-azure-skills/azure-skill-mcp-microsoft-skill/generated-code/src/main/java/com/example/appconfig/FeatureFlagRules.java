package com.example.appconfig;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.nio.ByteBuffer;
import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.ArrayList;
import java.util.List;
import java.util.Locale;

final class FeatureFlagRules {
    private static final ObjectMapper JSON = new ObjectMapper();
    private static final String PERCENTAGE_FILTER = "Microsoft.Percentage";

    private FeatureFlagRules() {
    }

    static boolean evaluate(String payload, String featureId, String userId) {
        JsonNode flag = parse(payload);
        if (!flag.path("enabled").asBoolean(false)) {
            return false;
        }

        JsonNode filters = flag.path("conditions").path("client_filters");
        if (!filters.isArray() || filters.isEmpty()) {
            return true;
        }

        String requirement = flag.path("conditions").path("requirement_type").asText("Any");
        List<Boolean> results = new ArrayList<>();
        filters.forEach(filter -> results.add(evaluateFilter(filter, featureId, userId)));
        return "all".equals(requirement.toLowerCase(Locale.ROOT))
            ? results.stream().allMatch(Boolean::booleanValue)
            : results.stream().anyMatch(Boolean::booleanValue);
    }

    static double bucket(String featureId, String userId) {
        try {
            byte[] digest = MessageDigest.getInstance("SHA-256")
                .digest((featureId + ":" + userId).getBytes(StandardCharsets.UTF_8));
            long unsignedPrefix = ByteBuffer.wrap(digest).getLong() & Long.MAX_VALUE;
            return (unsignedPrefix % 10_000) / 100.0;
        } catch (NoSuchAlgorithmException exception) {
            throw new IllegalStateException("SHA-256 is required by the Java runtime", exception);
        }
    }

    private static boolean evaluateFilter(JsonNode filter, String featureId, String userId) {
        if (!PERCENTAGE_FILTER.equals(filter.path("name").asText())) {
            return false;
        }
        if (userId == null || userId.isBlank()) {
            return false;
        }

        JsonNode value = filter.path("parameters").path("Value");
        double percentage;
        try {
            percentage = value.isNumber() ? value.doubleValue() : Double.parseDouble(value.asText());
        } catch (NumberFormatException exception) {
            throw new IllegalArgumentException("Percentage filter Value must be numeric", exception);
        }
        if (percentage < 0 || percentage > 100) {
            throw new IllegalArgumentException("Percentage filter Value must be between 0 and 100");
        }
        return bucket(featureId, userId) < percentage;
    }

    private static JsonNode parse(String payload) {
        try {
            return JSON.readTree(payload);
        } catch (JsonProcessingException exception) {
            throw new IllegalArgumentException("Feature flag contains invalid JSON", exception);
        }
    }
}
