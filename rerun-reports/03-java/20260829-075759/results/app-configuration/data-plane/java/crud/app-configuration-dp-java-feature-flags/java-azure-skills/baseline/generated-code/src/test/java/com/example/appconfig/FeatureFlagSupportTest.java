package com.example.appconfig;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

class FeatureFlagSupportTest {
    @Test
    void disabledFlagIsAlwaysDisabled() {
        assertFalse(FeatureFlagSupport.evaluate("test", """
            {"enabled": false}
            """, "alice"));
    }

    @Test
    void enabledFlagWithoutFiltersIsEnabled() {
        assertTrue(FeatureFlagSupport.evaluate("test", """
            {"enabled": true}
            """, null));
    }

    @Test
    void percentageRolloutIsDeterministic() {
        String payload = """
            {
              "enabled": true,
              "conditions": {
                "client_filters": [{
                  "name": "Microsoft.Percentage",
                  "parameters": {"Value": 30}
                }]
              }
            }
            """;

        boolean first = FeatureFlagSupport.evaluate("test", payload, "alice");
        assertEquals(first, FeatureFlagSupport.evaluate("test", payload, "alice"));
    }

    @Test
    void percentageBoundariesAreHonored() {
        String zeroPercent = percentagePayload("0");
        String oneHundredPercent = percentagePayload("100");

        assertFalse(FeatureFlagSupport.evaluate("test", zeroPercent, "alice"));
        assertTrue(FeatureFlagSupport.evaluate("test", oneHundredPercent, "alice"));
    }

    @Test
    void percentageRequiresUserId() {
        assertThrows(IllegalArgumentException.class,
            () -> FeatureFlagSupport.evaluate("test", percentagePayload("30"), null));
    }

    @Test
    void unsupportedFiltersFailExplicitly() {
        String payload = """
            {
              "enabled": true,
              "conditions": {
                "client_filters": [{"name": "Microsoft.Targeting", "parameters": {}}]
              }
            }
            """;

        assertThrows(UnsupportedOperationException.class,
            () -> FeatureFlagSupport.evaluate("test", payload, "alice"));
    }

    private static String percentagePayload(String percentage) {
        return """
            {
              "enabled": true,
              "conditions": {
                "client_filters": [{
                  "name": "Microsoft.Percentage",
                  "parameters": {"Value": "%s"}
                }]
              }
            }
            """.formatted(percentage);
    }
}
