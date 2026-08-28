package com.example.azureidentity;

import java.util.Map;
import java.util.stream.Stream;

public final class EnvironmentDetector {
    private static final String[] CI_MARKERS = {
        "CI",
        "TF_BUILD",
        "BUILD_BUILDID",
        "BUILD_SOURCESDIRECTORY",
        "SYSTEM_TEAMFOUNDATIONCOLLECTIONURI",
        "GITHUB_ACTIONS",
        "GITLAB_CI",
        "JENKINS_URL"
    };

    private static final String[] PRODUCTION_MARKERS = {
        "IDENTITY_ENDPOINT",
        "MSI_ENDPOINT",
        "IMDS_ENDPOINT",
        "AZURE_FEDERATED_TOKEN_FILE",
        "WEBSITE_INSTANCE_ID",
        "CONTAINER_APP_NAME",
        "KUBERNETES_SERVICE_HOST"
    };

    public DeploymentEnvironment detect() {
        return detect(System.getenv());
    }

    public DeploymentEnvironment detect(Map<String, String> environment) {
        if (hasAnyNonBlank(environment, CI_MARKERS)) {
            return DeploymentEnvironment.CI;
        }
        if (hasAnyNonBlank(environment, PRODUCTION_MARKERS)) {
            return DeploymentEnvironment.PRODUCTION;
        }
        return DeploymentEnvironment.DEV;
    }

    private boolean hasAnyNonBlank(Map<String, String> environment, String[] names) {
        return Stream.of(names).map(environment::get).anyMatch(EnvironmentDetector::isNonBlank);
    }

    private static boolean isNonBlank(String value) {
        return value != null && !value.isBlank();
    }
}
