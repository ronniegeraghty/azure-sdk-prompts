package com.example.azureidentity;

import java.net.HttpURLConnection;
import java.net.URI;
import java.time.Duration;
import java.util.Map;
import java.util.function.BooleanSupplier;

public final class EnvironmentDetector {
    private static final String IMDS_URL =
        "http://169.254.169.254/metadata/instance?api-version=2021-02-01";
    private static final Duration PROBE_TIMEOUT = Duration.ofMillis(250);

    private final Map<String, String> environment;
    private final BooleanSupplier managedIdentityEndpointProbe;

    public EnvironmentDetector() {
        this(System.getenv(), EnvironmentDetector::isImdsAvailable);
    }

    EnvironmentDetector(Map<String, String> environment, BooleanSupplier managedIdentityEndpointProbe) {
        this.environment = Map.copyOf(environment);
        this.managedIdentityEndpointProbe = managedIdentityEndpointProbe;
    }

    public DeploymentEnvironment detect() {
        if (hasAny("CI", "TF_BUILD", "BUILD_BUILDID", "PIPELINE_WORKSPACE",
            "GITHUB_ACTIONS", "GITHUB_WORKSPACE", "GITLAB_CI", "CI_PROJECT_DIR", "JENKINS_URL")) {
            return DeploymentEnvironment.CI;
        }

        if (hasAny("IDENTITY_ENDPOINT", "MSI_ENDPOINT", "IMDS_ENDPOINT")
            || hasWorkloadIdentityConfiguration()
            || managedIdentityEndpointProbe.getAsBoolean()) {
            return DeploymentEnvironment.PRODUCTION;
        }

        return DeploymentEnvironment.DEVELOPMENT;
    }

    private boolean hasWorkloadIdentityConfiguration() {
        return hasValue("AZURE_CLIENT_ID")
            && hasValue("AZURE_TENANT_ID")
            && hasValue("AZURE_FEDERATED_TOKEN_FILE");
    }

    private boolean hasAny(String... names) {
        for (String name : names) {
            if (hasValue(name)) {
                return true;
            }
        }
        return false;
    }

    private boolean hasValue(String name) {
        String value = environment.get(name);
        return value != null && !value.isBlank() && !"false".equalsIgnoreCase(value);
    }

    private static boolean isImdsAvailable() {
        HttpURLConnection connection = null;
        try {
            connection = (HttpURLConnection) URI.create(IMDS_URL).toURL().openConnection();
            connection.setRequestMethod("GET");
            connection.setRequestProperty("Metadata", "true");
            connection.setConnectTimeout((int) PROBE_TIMEOUT.toMillis());
            connection.setReadTimeout((int) PROBE_TIMEOUT.toMillis());
            int status = connection.getResponseCode();
            return status >= 200 && status < 500;
        } catch (Exception ignored) {
            return false;
        } finally {
            if (connection != null) {
                connection.disconnect();
            }
        }
    }
}
