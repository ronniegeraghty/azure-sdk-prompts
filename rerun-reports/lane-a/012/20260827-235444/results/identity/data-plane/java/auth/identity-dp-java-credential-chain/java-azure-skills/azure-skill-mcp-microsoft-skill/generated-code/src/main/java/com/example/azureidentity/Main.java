package com.example.azureidentity;

public final class Main {
    private static final String AZURE_RESOURCE_MANAGER_SCOPE = "https://management.azure.com/.default";

    private Main() {
    }

    public static void main(String[] args) {
        boolean enableCae = readCaeSetting();
        DeploymentEnvironment environment = new EnvironmentDetector().detect();
        CredentialSelection selection = new AzureCredentialFactory().create(environment, enableCae);

        System.out.println("Detected environment: " + environment);
        System.out.println("Credential strategy: " + selection.strategy());
        System.out.println("Scope: " + AZURE_RESOURCE_MANAGER_SCOPE);

        System.out.println("\nSynchronous connectivity test");
        new CredentialConnectivityTester().test(
            selection.credential(),
            AZURE_RESOURCE_MANAGER_SCOPE,
            selection.caeEnabled()
        );

        System.out.println("\nAsynchronous connectivity test");
        new AsyncCredentialConnectivityTester().test(
            selection.credential(),
            AZURE_RESOURCE_MANAGER_SCOPE,
            selection.caeEnabled()
        ).block();
    }

    private static boolean readCaeSetting() {
        String value = System.getenv().getOrDefault("AZURE_ENABLE_CAE", "true");
        return Boolean.parseBoolean(value);
    }
}
