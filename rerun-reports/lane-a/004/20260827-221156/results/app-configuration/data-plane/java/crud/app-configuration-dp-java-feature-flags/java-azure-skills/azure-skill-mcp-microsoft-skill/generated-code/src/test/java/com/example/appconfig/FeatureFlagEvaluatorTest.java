package com.example.appconfig;

import org.junit.jupiter.api.Test;

import java.util.Map;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

class FeatureFlagEvaluatorTest {
    @Test
    void evaluatesSimpleEnabledAndDisabledFlags() {
        assertTrue(FeatureFlagEvaluator.evaluate("flag", "user", """
            {"id":"flag","enabled":true,"conditions":{"client_filters":[]}}
            """));
        assertFalse(FeatureFlagEvaluator.evaluate("flag", "user", """
            {"id":"flag","enabled":false,"conditions":{"client_filters":[]}}
            """));
    }

    @Test
    void percentageRolloutIsDeterministic() {
        String payload = """
            {
              "id": "flag",
              "enabled": true,
              "conditions": {
                "client_filters": [
                  {"name":"Microsoft.Percentage","parameters":{"Value":30}}
                ]
              }
            }
            """;

        boolean first = FeatureFlagEvaluator.evaluate("flag", "alice", payload);
        assertEquals(first, FeatureFlagEvaluator.evaluate("flag", "alice", payload));
        assertEquals(
            FeatureFlagEvaluator.rolloutBucket("flag", "alice") < 30,
            first);
    }

    @Test
    void rejectsMissingUserForPercentageRollout() {
        String payload = """
            {
              "enabled": true,
              "conditions": {
                "client_filters": [
                  {"name":"Microsoft.Percentage","parameters":{"Value":50}}
                ]
              }
            }
            """;
        assertThrows(
            IllegalArgumentException.class,
            () -> FeatureFlagEvaluator.evaluate("flag", null, payload));
    }

    @Test
    void evaluatorUsesAzureFeatureFlagKeyAndLabel() {
        ConfigurationReader reader = new ConfigurationReader() {
            @Override
            public String getSetting(String key, String label) {
                assertEquals(".appconfig.featureflag/beta", key);
                assertEquals("staging", label);
                return """
                    {"id":"beta","enabled":true,"conditions":{"client_filters":[]}}
                    """;
            }

            @Override
            public Map<String, String> listSettings(String keyPrefix, String label) {
                return Map.of();
            }
        };

        assertTrue(new FeatureFlagEvaluator(reader).isEnabled("beta", "user-1", "staging"));
    }
}
