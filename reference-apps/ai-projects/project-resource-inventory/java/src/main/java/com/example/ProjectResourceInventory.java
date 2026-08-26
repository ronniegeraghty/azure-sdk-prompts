package com.example;

import com.azure.ai.projects.AIProjectClientBuilder;
import com.azure.ai.projects.ConnectionsClient;
import com.azure.ai.projects.DeploymentsClient;
import com.azure.ai.projects.models.Connection;
import com.azure.ai.projects.models.Deployment;
import com.azure.ai.projects.models.ModelDeployment;
import com.azure.identity.DefaultAzureCredentialBuilder;

public final class ProjectResourceInventory {
    private ProjectResourceInventory() {
    }

    public static void main(String[] args) {
        String endpoint = requireEnvironmentVariable("FOUNDRY_PROJECT_ENDPOINT");
        String connectionName = requireEnvironmentVariable("CONNECTION_NAME");
        String deploymentName = requireEnvironmentVariable("DEPLOYMENT_NAME");
        AIProjectClientBuilder builder = new AIProjectClientBuilder()
            .endpoint(endpoint)
            .credential(new DefaultAzureCredentialBuilder().build());
        ConnectionsClient connections = builder.buildConnectionsClient();
        DeploymentsClient deployments = builder.buildDeploymentsClient();

        System.out.println("Connections:");
        for (Connection connection : connections.listConnections()) {
            printConnection(connection);
        }

        System.out.println("Selected connection:");
        printConnection(connections.getConnection(connectionName, false));

        System.out.println("Model deployments:");
        for (Deployment deployment : deployments.listDeployments()) {
            if (deployment instanceof ModelDeployment modelDeployment) {
                printDeployment(modelDeployment);
            }
        }

        System.out.println("Selected model deployment:");
        Deployment selectedDeployment = deployments.getDeployment(deploymentName);
        if (!(selectedDeployment instanceof ModelDeployment modelDeployment)) {
            throw new IllegalStateException(
                deploymentName + " is not a model deployment.");
        }
        printDeployment(modelDeployment);
    }

    private static void printConnection(Connection connection) {
        System.out.printf(
            "%s | %s | %s | default=%s%n",
            connection.getName(),
            connection.getType(),
            connection.getTarget(),
            connection.isDefault());
    }

    private static void printDeployment(ModelDeployment deployment) {
        System.out.printf(
            "%s | %s | %s | %s%n",
            deployment.getName(),
            deployment.getModelPublisher(),
            deployment.getModelName(),
            deployment.getModelVersion());
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " is required.");
        }
        return value;
    }
}
