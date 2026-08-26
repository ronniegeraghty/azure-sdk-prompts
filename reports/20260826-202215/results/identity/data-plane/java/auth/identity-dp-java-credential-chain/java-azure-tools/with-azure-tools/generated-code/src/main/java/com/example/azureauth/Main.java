package com.example.azureauth;

public final class Main {
    private static final String AZURE_RESOURCE_MANAGER_SCOPE =
        "https://management.azure.com/.default";

    private Main() {
    }

    public static void main(String[] args) {
        DeploymentEnvironment environment = new EnvironmentDetector().detect();
        boolean enableCae = readBooleanEnvironmentVariable("AZURE_ENABLE_CAE", true);
        CredentialSelection selection = new AzureCredentialFactory().create(environment, enableCae);

        System.out.println("Detected environment: " + environment);
        System.out.println("Credential strategy: " + selection.strategy());
        System.out.println("Target scope: " + AZURE_RESOURCE_MANAGER_SCOPE);

        System.out.println("\nSynchronous connectivity test:");
        new CredentialConnectivityTester().test(selection, AZURE_RESOURCE_MANAGER_SCOPE);

        System.out.println("\nAsynchronous connectivity test:");
        new AsyncCredentialConnectivityTester()
            .test(selection, AZURE_RESOURCE_MANAGER_SCOPE)
            .block();
    }

    private static boolean readBooleanEnvironmentVariable(String name, boolean defaultValue) {
        String value = System.getenv(name);
        return value == null || value.isBlank() ? defaultValue : Boolean.parseBoolean(value);
    }
}
