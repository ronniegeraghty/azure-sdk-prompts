package com.example.appconfig;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

class FeatureFlagLogicTest {
    @Test
    void evaluatesSimpleEnabledAndDisabledFlags() {
        assertTrue(FeatureFlagLogic.evaluate("simple", "{\"enabled\":true}", null));
        assertFalse(FeatureFlagLogic.evaluate("simple", "{\"enabled\":false}", "alice"));
    }

    @Test
    void percentageRolloutIsDeterministic() {
        String payload = percentageFlag(30);
        boolean first = FeatureFlagLogic.evaluate("beta", payload, "alice");

        for (int attempt = 0; attempt < 20; attempt++) {
            assertEquals(first, FeatureFlagLogic.evaluate("beta", payload, "alice"));
        }
    }

    @Test
    void percentageBoundaryValuesAreRespected() {
        assertFalse(FeatureFlagLogic.evaluate("beta", percentageFlag(0), "alice"));
        assertTrue(FeatureFlagLogic.evaluate("beta", percentageFlag(100), "alice"));
        assertFalse(FeatureFlagLogic.evaluate("beta", percentageFlag(30), null));
    }

    @Test
    void percentageDistributionTracksConfiguredValue() {
        int enabled = 0;
        for (int user = 0; user < 10_000; user++) {
            if (FeatureFlagLogic.evaluate("beta", percentageFlag(30), "user-" + user)) {
                enabled++;
            }
        }

        assertTrue(enabled >= 2_800 && enabled <= 3_200, "enabled users: " + enabled);
    }

    @Test
    void rejectsMalformedPayloadsAndPercentages() {
        assertThrows(IllegalArgumentException.class,
            () -> FeatureFlagLogic.evaluate("beta", "not-json", "alice"));
        assertThrows(IllegalArgumentException.class,
            () -> FeatureFlagLogic.evaluate("beta", percentageFlag(101), "alice"));
    }

    private static String percentageFlag(int percentage) {
        return """
            {
              "enabled": true,
              "conditions": {
                "client_filters": [
                  {
                    "name": "Microsoft.Percentage",
                    "parameters": {"Value": "%d"}
                  }
                ]
              }
            }
            """.formatted(percentage);
    }
}
