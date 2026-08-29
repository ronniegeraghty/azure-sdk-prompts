package com.example.azureauth;

import java.util.Locale;
import java.util.Map;
import java.util.Set;

public final class EnvironmentDetector {
    private static final Set<String> CI_MARKERS = Set.of(
        "CI",
        "TF_BUILD",
        "GITHUB_ACTIONS",
        "GITLAB_CI",
        "JENKINS_URL",
        "BUILD_SOURCESDIRECTORY",
        "PIPELINE_WORKSPACE"
    );

    private static final Set<String> PRODUCTION_MARKERS = Set.of(
        "IDENTITY_ENDPOINT",
        "MSI_ENDPOINT",
        "IMDS_ENDPOINT",
        "WEBSITE_INSTANCE_ID",
        "CONTAINER_APP_NAME",
        "AZURE_FEDERATED_TOKEN_FILE"
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
        if (hasText(override)) {
            return parseOverride(override);
        }
        if (hasAnyMarker(CI_MARKERS)) {
            return DeploymentEnvironment.CI;
        }
        if (hasAnyMarker(PRODUCTION_MARKERS) || hasWorkloadIdentityConfiguration()) {
            return DeploymentEnvironment.PRODUCTION;
        }
        return DeploymentEnvironment.DEVELOPMENT;
    }

    private boolean hasAnyMarker(Set<String> names) {
        return names.stream().map(environment::get).anyMatch(EnvironmentDetector::hasText);
    }

    private boolean hasWorkloadIdentityConfiguration() {
        return hasText(environment.get("AZURE_TENANT_ID"))
            && hasText(environment.get("AZURE_CLIENT_ID"))
            && hasText(environment.get("AZURE_FEDERATED_TOKEN_FILE"));
    }

    private static DeploymentEnvironment parseOverride(String value) {
        return switch (value.trim().toLowerCase(Locale.ROOT)) {
            case "dev", "development", "local" -> DeploymentEnvironment.DEVELOPMENT;
            case "ci", "pipeline" -> DeploymentEnvironment.CI;
            case "prod", "production" -> DeploymentEnvironment.PRODUCTION;
            default -> throw new IllegalArgumentException(
                "APP_ENVIRONMENT must be one of: dev, ci, or production"
            );
        };
    }

    private static boolean hasText(String value) {
        return value != null && !value.isBlank();
    }
}
