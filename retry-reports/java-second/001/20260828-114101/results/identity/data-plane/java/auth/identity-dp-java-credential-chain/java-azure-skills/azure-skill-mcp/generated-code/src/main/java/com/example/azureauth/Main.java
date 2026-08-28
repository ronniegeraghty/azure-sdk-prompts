package com.example.azureauth;

import java.time.Duration;

public final class Main {
    private static final String ARM_SCOPE = "https://management.azure.com/.default";

    private Main() {
    }

    public static void main(String[] args) {
        DeploymentEnvironment environment = new EnvironmentDetector().detect();
        boolean enableCae = !"false".equalsIgnoreCase(System.getenv("AZURE_ENABLE_CAE"));
        CredentialSelection selection = new CredentialFactory().create(environment, enableCae);

        System.out.println("Detected environment: " + environment);
        System.out.println("Credential strategy: " + selection.strategy());
        System.out.println("CAE requested: " + selection.caeEnabled());
        System.out.println("Scope: " + ARM_SCOPE);
        System.out.println();

        new CredentialConnectivityTester().test(selection, ARM_SCOPE);
        System.out.println();
        new AsyncCredentialConnectivityTester()
            .test(selection, ARM_SCOPE)
            .block(Duration.ofMinutes(2));
    }
}
