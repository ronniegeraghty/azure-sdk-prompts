package com.example.appconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.nio.ByteBuffer;
import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.Iterator;
import java.util.Locale;

public final class FeatureFlagEvaluator {
    static final String FEATURE_FLAG_PREFIX = ".appconfig.featureflag/";
    private static final ObjectMapper OBJECT_MAPPER = new ObjectMapper();

    private final ConfigurationService configurationService;

    public FeatureFlagEvaluator(ConfigurationService configurationService) {
        this.configurationService = configurationService;
    }

    public boolean isEnabled(String flagId) {
        return isEnabled(flagId, null, null);
    }

    public boolean isEnabled(String flagId, String userId) {
        return isEnabled(flagId, userId, null);
    }

    public boolean isEnabled(String flagId, String userId, String label) {
        try {
            String payload = configurationService.getSetting(FEATURE_FLAG_PREFIX + flagId, label);
            return evaluatePayload(flagId, userId, payload);
        } catch (ResourceNotFoundException exception) {
            return false;
        }
    }

    static boolean evaluatePayload(String flagId, String userId, String payload) {
        try {
            JsonNode flag = OBJECT_MAPPER.readTree(payload);
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
                if (filterName.toLowerCase(Locale.ROOT).endsWith("percentage")) {
                    if (userId == null || userId.isBlank()) {
                        return false;
                    }
                    double percentage = readPercentage(filter.path("parameters"));
                    return bucket(flagId, userId) < percentage * 100.0;
                }
            }
            return false;
        } catch (JsonProcessingException exception) {
            throw new IllegalArgumentException("Invalid feature flag JSON for '" + flagId + "'", exception);
        }
    }

    private static double readPercentage(JsonNode parameters) {
        JsonNode value = parameters.has("Value") ? parameters.get("Value") : parameters.get("value");
        if (value == null || !value.isNumber()) {
            throw new IllegalArgumentException("Percentage filter requires a numeric Value parameter");
        }
        double percentage = value.asDouble();
        if (percentage < 0.0 || percentage > 100.0) {
            throw new IllegalArgumentException("Percentage filter Value must be between 0 and 100");
        }
        return percentage;
    }

    private static int bucket(String flagId, String userId) {
        try {
            MessageDigest digest = MessageDigest.getInstance("SHA-256");
            byte[] hash = digest.digest((flagId + ":" + userId).getBytes(StandardCharsets.UTF_8));
            long unsignedPrefix = Integer.toUnsignedLong(ByteBuffer.wrap(hash).getInt());
            return (int) (unsignedPrefix % 10_000);
        } catch (NoSuchAlgorithmException exception) {
            throw new IllegalStateException("SHA-256 is required by Java 17", exception);
        }
    }
}
