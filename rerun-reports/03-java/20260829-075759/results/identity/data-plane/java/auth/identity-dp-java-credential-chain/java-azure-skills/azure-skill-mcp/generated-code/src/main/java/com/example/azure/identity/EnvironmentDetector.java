package com.example.azure.identity;

import java.util.Locale;
import java.util.Map;
import java.util.Set;

public final class EnvironmentDetector {
    private static final Set<String> CI_MARKERS = Set.of(
        "CI",
        "TF_BUILD",
        "BUILD_BUILDID",
        "PIPELINE_WORKSPACE",
        "GITHUB_ACTIONS",
        "GITLAB_CI",
        "JENKINS_URL"
    );
    private static final Set<String> MANAGED_IDENTITY_MARKERS = Set.of(
        "IDENTITY_ENDPOINT",
        "MSI_ENDPOINT",
        "IMDS_ENDPOINT"
    );

    private final Map<String, String> environment;

    public EnvironmentDetector() {
        this(System.getenv());
    }

    EnvironmentDetector(Map<String, String> environment) {
        this.environment = Map.copyOf(environment);
    }

    public DeploymentEnvironment detect() {
        String override = environment.get("APP_ENVIRONMENT");
        if (isPresent(override)) {
            return switch (override.trim().toLowerCase(Locale.ROOT)) {
                case "dev", "development", "local" -> DeploymentEnvironment.DEV;
                case "ci", "pipeline" -> DeploymentEnvironment.CI;
                case "prod", "production" -> DeploymentEnvironment.PRODUCTION;
                default -> throw new IllegalArgumentException(
                    "APP_ENVIRONMENT must be one of: dev, ci, production");
            };
        }

        if (CI_MARKERS.stream().anyMatch(this::hasVariable)) {
            return DeploymentEnvironment.CI;
        }
        if (MANAGED_IDENTITY_MARKERS.stream().anyMatch(this::hasVariable) || hasWorkloadIdentityConfiguration()) {
            return DeploymentEnvironment.PRODUCTION;
        }
        return DeploymentEnvironment.DEV;
    }

    private boolean hasWorkloadIdentityConfiguration() {
        return hasVariable("AZURE_FEDERATED_TOKEN_FILE")
            && hasVariable("AZURE_TENANT_ID")
            && hasVariable("AZURE_CLIENT_ID");
    }

    private boolean hasVariable(String name) {
        return isPresent(environment.get(name));
    }

    private static boolean isPresent(String value) {
        return value != null && !value.isBlank();
    }
}
