package com.example.appconfig;

import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.nio.ByteBuffer;
import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.Iterator;
import java.util.Locale;
import java.util.Map;
import java.util.OptionalDouble;

final class FeatureFlag {
    static final String KEY_PREFIX = ".appconfig.featureflag/";
    private static final ObjectMapper OBJECT_MAPPER = new ObjectMapper();

    private FeatureFlag() {
    }

    static boolean evaluate(String flagId, String json, String userId) {
        JsonNode root;
        try {
            root = OBJECT_MAPPER.readTree(json);
        } catch (Exception exception) {
            throw new IllegalArgumentException("Invalid JSON for feature flag '" + flagId + "'", exception);
        }

        if (!root.path("enabled").asBoolean(false)) {
            return false;
        }

        OptionalDouble percentage = findPercentage(root);
        if (percentage.isEmpty()) {
            return true;
        }

        double value = percentage.getAsDouble();
        if (value < 0 || value > 100) {
            throw new IllegalArgumentException(
                "Percentage for feature flag '" + flagId + "' must be between 0 and 100");
        }
        if (value == 100) {
            return true;
        }
        if (value == 0 || userId == null || userId.isBlank()) {
            return false;
        }
        return rolloutBucket(flagId, userId) < value;
    }

    static int rolloutBucket(String flagId, String userId) {
        try {
            MessageDigest digest = MessageDigest.getInstance("SHA-256");
            byte[] hash = digest.digest((flagId + ":" + userId).getBytes(StandardCharsets.UTF_8));
            long firstEightBytes = ByteBuffer.wrap(hash).getLong();
            return (int) Long.remainderUnsigned(firstEightBytes, 100);
        } catch (NoSuchAlgorithmException exception) {
            throw new IllegalStateException("SHA-256 is required by the Java runtime", exception);
        }
    }

    private static OptionalDouble findPercentage(JsonNode root) {
        JsonNode allocationPercentage = root.path("allocation").path("percentage");
        if (!allocationPercentage.isMissingNode()) {
            return OptionalDouble.of(asPercentage(allocationPercentage));
        }

        JsonNode filters = root.path("conditions").path("client_filters");
        if (!filters.isArray()) {
            return OptionalDouble.empty();
        }

        for (JsonNode filter : filters) {
            String name = filter.path("name").asText("").toLowerCase(Locale.ROOT);
            if (name.equals("percentage") || name.endsWith(".percentage")) {
                JsonNode parameters = filter.path("parameters");
                Iterator<Map.Entry<String, JsonNode>> fields = parameters.fields();
                while (fields.hasNext()) {
                    Map.Entry<String, JsonNode> field = fields.next();
                    if (field.getKey().equalsIgnoreCase("value")) {
                        return OptionalDouble.of(asPercentage(field.getValue()));
                    }
                }
                throw new IllegalArgumentException("Percentage filter is missing its Value parameter");
            }
        }
        return OptionalDouble.empty();
    }

    private static double asPercentage(JsonNode node) {
        if (node.isNumber()) {
            return node.asDouble();
        }
        if (node.isTextual()) {
            try {
                return Double.parseDouble(node.asText());
            } catch (NumberFormatException exception) {
                throw new IllegalArgumentException("Percentage must be numeric", exception);
            }
        }
        throw new IllegalArgumentException("Percentage must be numeric");
    }
}
