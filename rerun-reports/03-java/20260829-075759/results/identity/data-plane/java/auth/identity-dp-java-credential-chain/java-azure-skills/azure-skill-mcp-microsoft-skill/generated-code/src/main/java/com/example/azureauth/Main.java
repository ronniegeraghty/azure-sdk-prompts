package com.example.azureauth;

public final class Main {
    private static final String ARM_SCOPE = "https://management.azure.com/.default";

    private Main() {
    }

    public static void main(String[] args) {
        boolean enableCae = readCaeSetting();
        DeploymentEnvironment environment = new EnvironmentDetector().detect();
        CredentialSelection selection = new AzureCredentialFactory().create(environment, enableCae);

        System.out.println("Detected environment: " + environment);
        System.out.println("Credential strategy: " + selection.strategy());
        System.out.println("Azure Resource Manager scope: " + ARM_SCOPE);

        System.out.println("\nSynchronous connectivity test:");
        new SyncCredentialConnectivityTester().test(selection, ARM_SCOPE);

        System.out.println("\nAsynchronous connectivity test:");
        new AsyncCredentialConnectivityTester().test(selection, ARM_SCOPE).block();
    }

    private static boolean readCaeSetting() {
        String configured = System.getenv().getOrDefault("AZURE_CAE_ENABLED", "true");
        return Boolean.parseBoolean(configured);
    }
}
