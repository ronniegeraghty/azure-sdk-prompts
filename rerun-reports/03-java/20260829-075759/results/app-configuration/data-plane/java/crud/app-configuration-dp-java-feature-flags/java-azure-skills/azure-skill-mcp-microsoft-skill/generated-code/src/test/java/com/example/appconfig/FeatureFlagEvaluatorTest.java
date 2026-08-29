package com.example.appconfig;

import org.junit.jupiter.api.Test;

import java.util.Map;
import java.util.Optional;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

class FeatureFlagEvaluatorTest {
    private static final String FLAG_KEY =
        FeatureFlagEvaluator.FEATURE_FLAG_PREFIX + "beta-dashboard";

    @Test
    void evaluatesSimpleEnabledAndDisabledFlags() {
        FeatureFlagEvaluator enabled = evaluator("""
            {"id":"beta-dashboard","enabled":true,"conditions":{"client_filters":[]}}
            """);
        FeatureFlagEvaluator disabled = evaluator("""
            {"id":"beta-dashboard","enabled":false,"conditions":{"client_filters":[]}}
            """);

        assertTrue(enabled.isEnabled("beta-dashboard", "production"));
        assertFalse(disabled.isEnabled("beta-dashboard", "production"));
    }

    @Test
    void percentageRolloutIsStableForEachUser() {
        FeatureFlagEvaluator evaluator = evaluator(percentageFlag(30));

        boolean first = evaluator.isEnabledForUser("beta-dashboard", "production", "alice");

        assertEquals(first, evaluator.isEnabledForUser("beta-dashboard", "production", "alice"));
    }

    @Test
    void percentageBoundariesIncludeNobodyOrEverybody() {
        FeatureFlagEvaluator zero = evaluator(percentageFlag(0));
        FeatureFlagEvaluator hundred = evaluator(percentageFlag(100));

        assertFalse(zero.isEnabledForUser("beta-dashboard", "production", "alice"));
        assertTrue(hundred.isEnabledForUser("beta-dashboard", "production", "alice"));
    }

    @Test
    void missingFlagDefaultsToDisabled() {
        FeatureFlagEvaluator evaluator = new FeatureFlagEvaluator((key, label) -> Optional.empty());

        assertFalse(evaluator.isEnabledForUser("missing", "production", "alice"));
    }

    @Test
    void invalidPercentageIsRejected() {
        FeatureFlagEvaluator evaluator = evaluator(percentageFlag(101));

        assertThrows(
            IllegalArgumentException.class,
            () -> evaluator.isEnabledForUser("beta-dashboard", "production", "alice")
        );
    }

    private static FeatureFlagEvaluator evaluator(String payload) {
        Map<String, String> settings = Map.of(FLAG_KEY + "@production", payload);
        return new FeatureFlagEvaluator(
            (key, label) -> Optional.ofNullable(settings.get(key + "@" + label))
        );
    }

    private static String percentageFlag(int percentage) {
        return """
            {
              "id": "beta-dashboard",
              "enabled": true,
              "conditions": {
                "client_filters": [
                  {
                    "name": "Microsoft.Percentage",
                    "parameters": {"Value": %d}
                  }
                ]
              }
            }
            """.formatted(percentage);
    }
}
