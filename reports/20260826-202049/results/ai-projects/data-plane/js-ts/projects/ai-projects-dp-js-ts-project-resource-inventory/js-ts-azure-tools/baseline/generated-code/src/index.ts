import {
  AIProjectClient,
  type Connection,
  type DeploymentUnion,
  type ModelDeployment,
} from "@azure/ai-projects";
import { DefaultAzureCredential } from "@azure/identity";

interface Configuration {
  projectEndpoint: string;
  connectionName: string;
  deploymentName: string;
}

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name]?.trim();
  if (!value) {
    throw new Error(`Required environment variable ${name} is not set.`);
  }
  return value;
}

function loadConfiguration(): Configuration {
  return {
    projectEndpoint: requireEnvironmentVariable("FOUNDRY_PROJECT_ENDPOINT"),
    connectionName: requireEnvironmentVariable("CONNECTION_NAME"),
    deploymentName: requireEnvironmentVariable("DEPLOYMENT_NAME"),
  };
}

function printConnection(connection: Connection): void {
  console.log({
    name: connection.name,
    type: connection.type,
    target: connection.target,
    isDefault: connection.isDefault,
  });
}

function isModelDeployment(deployment: DeploymentUnion): deployment is ModelDeployment {
  return deployment.type === "ModelDeployment";
}

function printModelDeployment(deployment: ModelDeployment): void {
  console.log({
    name: deployment.name,
    modelPublisher: deployment.modelPublisher,
    modelName: deployment.modelName,
    modelVersion: deployment.modelVersion,
  });
}

async function main(): Promise<void> {
  const config = loadConfiguration();
  const project = new AIProjectClient(
    config.projectEndpoint,
    new DefaultAzureCredential(),
  );

  console.log("Project connections:");
  for await (const connection of project.connections.list()) {
    printConnection(connection);
  }

  console.log(`Connection "${config.connectionName}":`);
  const connection = await project.connections.get(config.connectionName);
  printConnection(connection);

  console.log("Project model deployments:");
  for await (const deployment of project.deployments.list()) {
    if (isModelDeployment(deployment)) {
      printModelDeployment(deployment);
    }
  }

  console.log(`Deployment "${config.deploymentName}":`);
  const deployment = await project.deployments.get(config.deploymentName);
  if (!isModelDeployment(deployment)) {
    throw new Error(
      `Deployment "${config.deploymentName}" has type "${deployment.type}", not "ModelDeployment".`,
    );
  }
  printModelDeployment(deployment);
}

await main().catch((error: unknown) => {
  console.error("Failed to inspect the Foundry project:", error);
  process.exitCode = 1;
});
