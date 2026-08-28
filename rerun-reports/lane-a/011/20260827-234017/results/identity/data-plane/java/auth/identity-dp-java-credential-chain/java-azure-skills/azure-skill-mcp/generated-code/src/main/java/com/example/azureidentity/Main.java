package com.example.azureidentity;

public final class Main {
    private static final String ARM_SCOPE = "https://management.azure.com/.default";

    private Main() {
    }

    public static void main(String[] args) {
        boolean enableCae = isCaeEnabled(args);
        DeploymentEnvironment environment = new EnvironmentDetector().detect();
        CredentialSelection selection = new AzureCredentialFactory().create(environment, enableCae);

        System.out.println("Detected environment: " + environment);
        System.out.println("Credential strategy: " + selection.strategy());
        System.out.println("CAE requested: " + selection.caeEnabled());
        System.out.println();

        new SyncCredentialConnectivityTester()
            .test(selection.credential(), ARM_SCOPE, selection.caeEnabled());
        System.out.println();
        new AsyncCredentialConnectivityTester()
            .test(selection.credential(), ARM_SCOPE, selection.caeEnabled())
            .block();
    }

    private static boolean isCaeEnabled(String[] args) {
        for (String arg : args) {
            if ("--cae".equalsIgnoreCase(arg)) {
                return true;
            }
        }
        return Boolean.parseBoolean(System.getenv().getOrDefault("AZURE_ENABLE_CAE", "false"));
    }
}
