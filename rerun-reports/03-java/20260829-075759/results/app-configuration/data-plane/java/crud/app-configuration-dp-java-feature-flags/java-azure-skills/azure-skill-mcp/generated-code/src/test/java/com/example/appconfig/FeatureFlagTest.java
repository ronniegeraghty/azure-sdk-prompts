package com.example.appconfig;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

class FeatureFlagTest {
    @Test
    void disabledFlagIsAlwaysDisabled() {
        assertFalse(FeatureFlag.evaluate("checkout", "{\"enabled\":false}", "alice"));
    }

    @Test
    void enabledFlagWithoutFiltersIsEnabled() {
        assertTrue(FeatureFlag.evaluate("checkout", "{\"enabled\":true}", "alice"));
    }

    @Test
    void percentageRolloutIsDeterministic() {
        String json = """
            {
              "enabled": true,
              "conditions": {
                "client_filters": [{
                  "name": "Microsoft.Percentage",
                  "parameters": { "Value": "30" }
                }]
              }
            }
            """;

        boolean first = FeatureFlag.evaluate("checkout", json, "alice");
        assertEquals(first, FeatureFlag.evaluate("checkout", json, "alice"));
        assertEquals(FeatureFlag.rolloutBucket("checkout", "alice") < 30, first);
    }

    @Test
    void percentageBoundariesAreHandled() {
        String zero = """
            {"enabled":true,"conditions":{"client_filters":[
              {"name":"Percentage","parameters":{"Value":0}}
            ]}}
            """;
        String hundred = """
            {"enabled":true,"conditions":{"client_filters":[
              {"name":"Percentage","parameters":{"Value":100}}
            ]}}
            """;

        assertFalse(FeatureFlag.evaluate("checkout", zero, "alice"));
        assertTrue(FeatureFlag.evaluate("checkout", hundred, null));
    }

    @Test
    void invalidPercentageIsRejected() {
        String json = """
            {"enabled":true,"conditions":{"client_filters":[
              {"name":"Percentage","parameters":{"Value":101}}
            ]}}
            """;

        assertThrows(IllegalArgumentException.class,
            () -> FeatureFlag.evaluate("checkout", json, "alice"));
    }
}
