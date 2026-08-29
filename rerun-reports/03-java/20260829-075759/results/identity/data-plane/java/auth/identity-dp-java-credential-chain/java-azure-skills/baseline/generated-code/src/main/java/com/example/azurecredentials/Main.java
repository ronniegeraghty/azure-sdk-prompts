package com.example.azurecredentials;

public final class Main {
    private static final String AZURE_RESOURCE_MANAGER_SCOPE =
        "https://management.azure.com/.default";

    private Main() {
    }

    public static void main(String[] args) {
        boolean enableCae = readCaeSetting();
        DeploymentEnvironment environment = new EnvironmentDetector().detect();
        CredentialSelection selection =
            new AzureCredentialFactory().create(environment, enableCae);

        System.out.println("Detected environment: " + environment);
        System.out.println("Credential strategy: " + selection.strategy());
        System.out.println("CAE requested: " + selection.caeEnabled());
        System.out.println();

        new CredentialConnectivityTester().test(selection, AZURE_RESOURCE_MANAGER_SCOPE);
        System.out.println();

        Boolean asyncResult = new AsyncCredentialConnectivityTester()
            .test(selection, AZURE_RESOURCE_MANAGER_SCOPE)
            .block();
        if (asyncResult == null) {
            System.out.println("[async] Authentication test completed without a result");
        }
    }

    private static boolean readCaeSetting() {
        // Developer-tool credentials do not all support CAE, so it is explicitly opt-in.
        String value = System.getenv().getOrDefault("AZURE_ENABLE_CAE", "false");
        return Boolean.parseBoolean(value);
    }
}
