package com.example.azure.identity;

public final class Main {
    private static final String AZURE_RESOURCE_MANAGER_SCOPE
        = "https://management.azure.com/.default";

    private Main() {
    }

    public static void main(String[] args) {
        DeploymentEnvironment environment = new EnvironmentDetector().detect();
        boolean caeEnabled = Boolean.parseBoolean(
            System.getenv().getOrDefault("AZURE_ENABLE_CAE", "false"));
        CredentialFactory.BuiltCredential credential
            = new CredentialFactory().create(environment, caeEnabled);

        System.out.println("Detected environment: " + environment);
        System.out.println("Credential strategy: " + credential.strategy());
        System.out.println("Azure scope: " + AZURE_RESOURCE_MANAGER_SCOPE);
        System.out.println();

        new CredentialConnectivityTester().test(credential, AZURE_RESOURCE_MANAGER_SCOPE);
        System.out.println();
        new AsyncCredentialConnectivityTester()
            .test(credential, AZURE_RESOURCE_MANAGER_SCOPE)
            .block();
    }
}
