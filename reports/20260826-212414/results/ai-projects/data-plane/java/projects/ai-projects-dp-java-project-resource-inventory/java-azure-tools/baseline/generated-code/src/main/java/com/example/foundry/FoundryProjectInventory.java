package com.example.foundry;

import com.azure.ai.projects.AIProjectClientBuilder;
import com.azure.ai.projects.ConnectionsClient;
import com.azure.ai.projects.DeploymentsClient;
import com.azure.ai.projects.models.Connection;
import com.azure.ai.projects.models.Deployment;
import com.azure.ai.projects.models.ModelDeployment;
import com.azure.core.http.rest.PagedIterable;
import com.azure.identity.DefaultAzureCredentialBuilder;

public final class FoundryProjectInventory {
    private FoundryProjectInventory() {
    }

    public static void main(String[] args) {
        String endpoint = requireEnvironmentVariable("FOUNDRY_PROJECT_ENDPOINT");
        String connectionName = requireEnvironmentVariable("CONNECTION_NAME");
        String deploymentName = requireEnvironmentVariable("DEPLOYMENT_NAME");

        AIProjectClientBuilder clientBuilder = new AIProjectClientBuilder()
            .endpoint(endpoint)
            .credential(new DefaultAzureCredentialBuilder().build());

        ConnectionsClient connectionsClient = clientBuilder.buildConnectionsClient();
        DeploymentsClient deploymentsClient = clientBuilder.buildDeploymentsClient();

        printAllConnections(connectionsClient);
        printNamedConnection(connectionsClient, connectionName);
        printAllModelDeployments(deploymentsClient);
        printNamedModelDeployment(deploymentsClient, deploymentName);
    }

    private static void printAllConnections(ConnectionsClient client) {
        System.out.println("Project connections");
        PagedIterable<Connection> connections = client.listConnections();
        for (Connection connection : connections) {
            printConnection(connection);
        }
    }

    private static void printNamedConnection(ConnectionsClient client, String connectionName) {
        System.out.printf("%nConnection '%s' (credentials excluded)%n", connectionName);
        Connection connection = client.getConnection(connectionName, false);
        printConnection(connection);
    }

    private static void printConnection(Connection connection) {
        System.out.printf(
            "name=%s, type=%s, target=%s, default=%s%n",
            connection.getName(),
            connection.getType(),
            connection.getTarget(),
            connection.isDefault());
    }

    private static void printAllModelDeployments(DeploymentsClient client) {
        System.out.println("\nProject model deployments");
        PagedIterable<Deployment> deployments = client.listDeployments();
        for (Deployment deployment : deployments) {
            if (deployment instanceof ModelDeployment modelDeployment) {
                printModelDeployment(modelDeployment);
            }
        }
    }

    private static void printNamedModelDeployment(DeploymentsClient client, String deploymentName) {
        System.out.printf("%nModel deployment '%s'%n", deploymentName);
        Deployment deployment = client.getDeployment(deploymentName);
        if (!(deployment instanceof ModelDeployment modelDeployment)) {
            throw new IllegalStateException(
                "Deployment '%s' has type '%s'; expected ModelDeployment."
                    .formatted(deploymentName, deployment.getType()));
        }
        printModelDeployment(modelDeployment);
    }

    private static void printModelDeployment(ModelDeployment deployment) {
        System.out.printf(
            "name=%s, publisher=%s, model=%s, version=%s%n",
            deployment.getName(),
            deployment.getModelPublisher(),
            deployment.getModelName(),
            deployment.getModelVersion());
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException("Required environment variable is missing or blank: " + name);
        }
        return value;
    }
}
