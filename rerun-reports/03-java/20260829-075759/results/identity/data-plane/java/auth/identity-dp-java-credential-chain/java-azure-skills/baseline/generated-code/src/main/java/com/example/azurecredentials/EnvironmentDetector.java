package com.example.azurecredentials;

import java.util.Map;
import java.util.Objects;

public final class EnvironmentDetector {
    private static final String[] CI_MARKERS = {
        "CI",
        "TF_BUILD",
        "GITHUB_ACTIONS",
        "GITHUB_WORKSPACE",
        "PIPELINE_WORKSPACE",
        "BUILD_BUILDID",
        "JENKINS_URL"
    };

    private static final String[] MANAGED_IDENTITY_MARKERS = {
        "IDENTITY_ENDPOINT",
        "MSI_ENDPOINT",
        "IMDS_ENDPOINT"
    };

    private final Map<String, String> environment;

    public EnvironmentDetector() {
        this(System.getenv());
    }

    EnvironmentDetector(Map<String, String> environment) {
        this.environment = Map.copyOf(Objects.requireNonNull(environment, "environment"));
    }

    public DeploymentEnvironment detect() {
        if (hasAny(CI_MARKERS)) {
            return DeploymentEnvironment.CI;
        }
        if (hasAny(MANAGED_IDENTITY_MARKERS) || hasWorkloadIdentityConfiguration()) {
            return DeploymentEnvironment.PRODUCTION;
        }
        return DeploymentEnvironment.DEVELOPMENT;
    }

    private boolean hasWorkloadIdentityConfiguration() {
        return hasValue("AZURE_FEDERATED_TOKEN_FILE")
            && hasValue("AZURE_TENANT_ID")
            && hasValue("AZURE_CLIENT_ID");
    }

    private boolean hasAny(String[] names) {
        for (String name : names) {
            if (hasValue(name)) {
                return true;
            }
        }
        return false;
    }

    private boolean hasValue(String name) {
        String value = environment.get(name);
        return value != null && !value.isBlank();
    }
}
