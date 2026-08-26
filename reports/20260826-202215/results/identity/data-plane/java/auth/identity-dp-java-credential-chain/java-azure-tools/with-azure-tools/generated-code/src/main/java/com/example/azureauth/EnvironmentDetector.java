package com.example.azureauth;

import java.util.Map;
import java.util.Objects;
import java.util.Set;

public final class EnvironmentDetector {
    private static final Set<String> CI_SIGNALS = Set.of(
        "CI",
        "TF_BUILD",
        "GITHUB_ACTIONS",
        "GITHUB_WORKSPACE",
        "BUILD_BUILDID",
        "BUILD_SOURCESDIRECTORY",
        "JENKINS_URL",
        "GITLAB_CI"
    );

    private static final Set<String> PRODUCTION_SIGNALS = Set.of(
        "IDENTITY_ENDPOINT",
        "MSI_ENDPOINT",
        "IMDS_ENDPOINT",
        "WEBSITE_SITE_NAME",
        "FUNCTIONS_WORKER_RUNTIME",
        "AZURE_FEDERATED_TOKEN_FILE",
        "KUBERNETES_SERVICE_HOST"
    );

    private final Map<String, String> environment;

    public EnvironmentDetector() {
        this(System.getenv());
    }

    EnvironmentDetector(Map<String, String> environment) {
        this.environment = Map.copyOf(Objects.requireNonNull(environment, "environment"));
    }

    public DeploymentEnvironment detect() {
        if (hasAny(CI_SIGNALS)) {
            return DeploymentEnvironment.CI;
        }
        if (hasAny(PRODUCTION_SIGNALS)) {
            return DeploymentEnvironment.PRODUCTION;
        }
        return DeploymentEnvironment.DEVELOPMENT;
    }

    private boolean hasAny(Set<String> names) {
        return names.stream().map(environment::get).anyMatch(EnvironmentDetector::isPresent);
    }

    private static boolean isPresent(String value) {
        return value != null && !value.isBlank() && !"false".equalsIgnoreCase(value);
    }
}
