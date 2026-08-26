package com.example.appconfig;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

class FeatureFlagEvaluatorTest {
    @Test
    void evaluatesSimpleFlags() {
        assertTrue(FeatureFlagEvaluator.evaluatePayload("flag", null, "{\"enabled\":true}"));
        assertFalse(FeatureFlagEvaluator.evaluatePayload("flag", "alice", "{\"enabled\":false}"));
    }

    @Test
    void percentageRolloutIsDeterministic() {
        String payload = """
            {
              "enabled": true,
              "conditions": {
                "client_filters": [
                  {"name": "Microsoft.Percentage", "parameters": {"Value": 30}}
                ]
              }
            }
            """;

        boolean first = FeatureFlagEvaluator.evaluatePayload("flag", "alice", payload);
        for (int attempt = 0; attempt < 20; attempt++) {
            assertEquals(first, FeatureFlagEvaluator.evaluatePayload("flag", "alice", payload));
        }
    }

    @Test
    void percentageBoundariesAreRespected() {
        String zeroPercent = percentagePayload(0);
        String fullRollout = percentagePayload(100);

        assertFalse(FeatureFlagEvaluator.evaluatePayload("flag", "alice", zeroPercent));
        assertTrue(FeatureFlagEvaluator.evaluatePayload("flag", "alice", fullRollout));
        assertFalse(FeatureFlagEvaluator.evaluatePayload("flag", null, fullRollout));
    }

    @Test
    void rejectsInvalidPercentage() {
        assertThrows(
            IllegalArgumentException.class,
            () -> FeatureFlagEvaluator.evaluatePayload("flag", "alice", percentagePayload(101))
        );
    }

    private static String percentagePayload(int percentage) {
        return """
            {
              "enabled": true,
              "conditions": {
                "client_filters": [
                  {"name": "Microsoft.Percentage", "parameters": {"Value": %d}}
                ]
              }
            }
            """.formatted(percentage);
    }
}
