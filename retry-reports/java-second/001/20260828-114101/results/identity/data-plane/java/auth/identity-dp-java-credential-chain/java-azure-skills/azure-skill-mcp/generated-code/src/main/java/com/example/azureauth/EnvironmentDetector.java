package com.example.azureauth;

import java.util.Locale;
import java.util.Map;
import java.util.Set;

public final class EnvironmentDetector {
    private static final Set<String> CI_MARKERS = Set.of(
        "BUILD_BUILDID",
        "BUILD_SOURCESDIRECTORY",
        "SYSTEM_TEAMPROJECT",
        "SYSTEM_OIDCREQUESTURI",
        "GITHUB_ACTIONS",
        "GITLAB_CI",
        "JENKINS_URL",
        "TF_BUILD"
    );

    private static final Set<String> PRODUCTION_MARKERS = Set.of(
        "IDENTITY_ENDPOINT",
        "MSI_ENDPOINT",
        "IMDS_ENDPOINT",
        "WEBSITE_INSTANCE_ID",
        "CONTAINER_APP_NAME",
        "KUBERNETES_SERVICE_HOST",
        "AZURE_FEDERATED_TOKEN_FILE"
    );

    public DeploymentEnvironment detect() {
        return detect(System.getenv());
    }

    DeploymentEnvironment detect(Map<String, String> environment) {
        String override = environment.get("APP_DEPLOYMENT_ENVIRONMENT");
        if (hasText(override)) {
            return parseOverride(override);
        }

        if (isTrue(environment.get("CI")) || containsAny(environment, CI_MARKERS)) {
            return DeploymentEnvironment.CI;
        }

        if (containsAny(environment, PRODUCTION_MARKERS)) {
            return DeploymentEnvironment.PRODUCTION;
        }

        return DeploymentEnvironment.DEVELOPMENT;
    }

    private DeploymentEnvironment parseOverride(String value) {
        return switch (value.trim().toLowerCase(Locale.ROOT)) {
            case "dev", "development", "local" -> DeploymentEnvironment.DEVELOPMENT;
            case "ci", "pipeline" -> DeploymentEnvironment.CI;
            case "prod", "production" -> DeploymentEnvironment.PRODUCTION;
            default -> throw new IllegalArgumentException(
                "APP_DEPLOYMENT_ENVIRONMENT must be dev, ci, or production; got: " + value
            );
        };
    }

    private boolean containsAny(Map<String, String> environment, Set<String> names) {
        return names.stream().anyMatch(name -> hasText(environment.get(name)));
    }

    private boolean isTrue(String value) {
        return hasText(value) && !"false".equalsIgnoreCase(value) && !"0".equals(value);
    }

    private boolean hasText(String value) {
        return value != null && !value.isBlank();
    }
}
