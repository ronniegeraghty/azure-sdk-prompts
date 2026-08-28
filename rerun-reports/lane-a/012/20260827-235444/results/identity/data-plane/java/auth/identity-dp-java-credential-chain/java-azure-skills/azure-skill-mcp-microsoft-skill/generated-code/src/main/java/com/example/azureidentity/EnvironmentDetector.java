package com.example.azureidentity;

import java.util.Map;
import java.util.Objects;
import java.util.function.Predicate;

public final class EnvironmentDetector {
    private static final String[] CI_MARKERS = {
        "TF_BUILD",
        "PIPELINE_WORKSPACE",
        "BUILD_BUILDID",
        "GITHUB_ACTIONS",
        "GITLAB_CI",
        "JENKINS_URL",
        "CI"
    };

    private static final String[] PRODUCTION_MARKERS = {
        "IDENTITY_ENDPOINT",
        "MSI_ENDPOINT",
        "IMDS_ENDPOINT",
        "AZURE_FEDERATED_TOKEN_FILE",
        "WEBSITE_INSTANCE_ID",
        "CONTAINER_APP_NAME"
    };

    private final Map<String, String> environment;

    public EnvironmentDetector() {
        this(System.getenv());
    }

    EnvironmentDetector(Map<String, String> environment) {
        this.environment = Map.copyOf(Objects.requireNonNull(environment, "environment"));
    }

    public DeploymentEnvironment detect() {
        if (containsAny(CI_MARKERS, this::isTruthyMarker)) {
            return DeploymentEnvironment.CI;
        }
        if (containsAny(PRODUCTION_MARKERS, this::hasValue)) {
            return DeploymentEnvironment.PRODUCTION;
        }
        return DeploymentEnvironment.DEVELOPMENT;
    }

    private boolean containsAny(String[] names, Predicate<String> predicate) {
        for (String name : names) {
            if (predicate.test(name)) {
                return true;
            }
        }
        return false;
    }

    private boolean hasValue(String name) {
        String value = environment.get(name);
        return value != null && !value.isBlank();
    }

    private boolean isTruthyMarker(String name) {
        if (!hasValue(name)) {
            return false;
        }
        String value = environment.get(name);
        return !value.equalsIgnoreCase("false") && !value.equals("0");
    }
}
