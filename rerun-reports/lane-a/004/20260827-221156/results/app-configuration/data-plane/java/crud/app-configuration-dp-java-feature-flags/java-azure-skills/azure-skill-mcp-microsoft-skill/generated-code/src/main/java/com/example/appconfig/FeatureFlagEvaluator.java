package com.example.appconfig;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.nio.ByteBuffer;
import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;

public final class FeatureFlagEvaluator {
    static final String FEATURE_FLAG_PREFIX = ".appconfig.featureflag/";
    private static final String PERCENTAGE_FILTER = "Microsoft.Percentage";
    private static final ObjectMapper OBJECT_MAPPER = new ObjectMapper();

    private final ConfigurationReader configuration;

    public FeatureFlagEvaluator(ConfigurationReader configuration) {
        this.configuration = configuration;
    }

    public boolean isEnabled(String flagId, String userId) {
        return isEnabled(flagId, userId, null);
    }

    public boolean isEnabled(String flagId, String userId, String label) {
        String payload = configuration.getSetting(FEATURE_FLAG_PREFIX + flagId, label);
        return evaluate(flagId, userId, payload);
    }

    static boolean evaluate(String flagId, String userId, String payload) {
        JsonNode flag;
        try {
            flag = OBJECT_MAPPER.readTree(payload);
        } catch (JsonProcessingException e) {
            throw new IllegalArgumentException("Invalid JSON for feature flag '" + flagId + "'", e);
        }

        if (!flag.path("enabled").asBoolean(false)) {
            return false;
        }

        for (JsonNode filter : flag.path("conditions").path("client_filters")) {
            if (PERCENTAGE_FILTER.equals(filter.path("name").asText())) {
                double percentage = filter.path("parameters").path("Value").asDouble(0);
                return percentage >= 100 || percentage > 0 && rolloutBucket(flagId, userId) < percentage;
            }
        }
        return true;
    }

    static double rolloutBucket(String flagId, String userId) {
        if (userId == null || userId.isBlank()) {
            throw new IllegalArgumentException("userId is required for percentage rollout");
        }
        try {
            byte[] digest = MessageDigest.getInstance("SHA-256")
                .digest((flagId + ":" + userId).getBytes(StandardCharsets.UTF_8));
            long unsignedPrefix = Integer.toUnsignedLong(ByteBuffer.wrap(digest).getInt());
            return (unsignedPrefix % 10_000) / 100.0;
        } catch (NoSuchAlgorithmException e) {
            throw new IllegalStateException("SHA-256 is unavailable", e);
        }
    }
}
