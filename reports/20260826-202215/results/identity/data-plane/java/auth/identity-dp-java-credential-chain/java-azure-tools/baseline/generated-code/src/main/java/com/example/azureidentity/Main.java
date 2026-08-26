package com.example.azureidentity;

public final class Main {
    private static final String AZURE_RESOURCE_MANAGER_SCOPE =
        "https://management.azure.com/.default";

    private Main() {
    }

    public static void main(String[] args) {
        boolean caeEnabled = readCaeSetting();
        DeploymentEnvironment environment = new EnvironmentDetector().detect();
        AzureCredentialFactory.CredentialSelection selection =
            new AzureCredentialFactory().create(environment, caeEnabled);

        System.out.println("Detected environment: " + environment);
        System.out.println("Credential strategy: " + selection.strategy());
        System.out.println("Azure scope: " + AZURE_RESOURCE_MANAGER_SCOPE);
        System.out.println();

        new CredentialConnectivityTester().test(
            selection.credential(), AZURE_RESOURCE_MANAGER_SCOPE, selection.caeEnabled());
        System.out.println();
        new AsyncCredentialConnectivityTester().test(
            selection.credential(), AZURE_RESOURCE_MANAGER_SCOPE, selection.caeEnabled()).block();
    }

    private static boolean readCaeSetting() {
        String value = System.getenv("AZURE_ENABLE_CAE");
        return value == null || value.isBlank() || Boolean.parseBoolean(value);
    }
}
