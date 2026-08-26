package com.example.foundry;

import com.azure.ai.projects.AIProjectClientBuilder;
import com.azure.ai.projects.ConnectionsClient;
import com.azure.ai.projects.DeploymentsClient;
import com.azure.ai.projects.models.Connection;
import com.azure.ai.projects.models.Deployment;
import com.azure.ai.projects.models.ModelDeployment;
import com.azure.core.http.rest.PagedIterable;
import com.azure.identity.DefaultAzureCredentialBuilder;

public final class ProjectInventory {
    private ProjectInventory() {
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

        System.out.println("Project connections");
        PagedIterable<Connection> connections = connectionsClient.listConnections();
        for (Connection connection : connections) {
            printConnection(connection);
        }

        System.out.printf("%nConnection '%s' (credentials excluded)%n", connectionName);
        Connection selectedConnection = connectionsClient.getConnection(connectionName, false);
        printConnection(selectedConnection);

        System.out.println("\nModel deployments");
        PagedIterable<Deployment> deployments = deploymentsClient.listDeployments();
        for (Deployment deployment : deployments) {
            if (deployment instanceof ModelDeployment modelDeployment) {
                printModelDeployment(modelDeployment);
            }
        }

        System.out.printf("%nModel deployment '%s'%n", deploymentName);
        Deployment selectedDeployment = deploymentsClient.getDeployment(deploymentName);
        if (!(selectedDeployment instanceof ModelDeployment modelDeployment)) {
            throw new IllegalStateException(
                "Deployment '%s' is not a model deployment (type: %s)."
                    .formatted(deploymentName, selectedDeployment.getType()));
        }
        printModelDeployment(modelDeployment);
    }

    private static void printConnection(Connection connection) {
        System.out.printf(
            "name=%s, type=%s, target=%s, default=%s%n",
            connection.getName(),
            connection.getType(),
            connection.getTarget(),
            connection.isDefault());
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
            throw new IllegalStateException("Required environment variable is not set: " + name);
        }
        return value;
    }
}
