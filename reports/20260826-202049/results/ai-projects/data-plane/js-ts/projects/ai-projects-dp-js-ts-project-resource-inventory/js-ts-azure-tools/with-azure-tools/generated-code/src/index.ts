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
  console.log(`  Name: ${connection.name}`);
  console.log(`  Type: ${connection.type}`);
  console.log(`  Target: ${connection.target}`);
  console.log(`  Default: ${connection.isDefault}`);
}

function isModelDeployment(
  deployment: DeploymentUnion,
): deployment is ModelDeployment {
  return deployment.type === "ModelDeployment";
}

function printModelDeployment(deployment: ModelDeployment): void {
  console.log(`  Name: ${deployment.name}`);
  console.log(`  Model publisher: ${deployment.modelPublisher}`);
  console.log(`  Model name: ${deployment.modelName}`);
  console.log(`  Model version: ${deployment.modelVersion}`);
}

async function main(): Promise<void> {
  const configuration = loadConfiguration();
  const client = new AIProjectClient(
    configuration.projectEndpoint,
    new DefaultAzureCredential(),
  );

  console.log("Project connections:");
  for await (const connection of client.connections.list()) {
    printConnection(connection);
    console.log();
  }

  console.log(`Connection "${configuration.connectionName}":`);
  const connection = await client.connections.get(configuration.connectionName);
  printConnection(connection);

  console.log("\nProject model deployments:");
  for await (const deployment of client.deployments.list()) {
    if (isModelDeployment(deployment)) {
      printModelDeployment(deployment);
      console.log();
    }
  }

  console.log(`Model deployment "${configuration.deploymentName}":`);
  const deployment = await client.deployments.get(
    configuration.deploymentName,
  );

  if (!isModelDeployment(deployment)) {
    throw new Error(
      `Deployment "${configuration.deploymentName}" has type "${deployment.type}", not "ModelDeployment".`,
    );
  }

  printModelDeployment(deployment);
}

await main().catch((error: unknown) => {
  const message = error instanceof Error ? error.message : String(error);
  console.error(`Failed to inspect Foundry project resources: ${message}`);
  process.exitCode = 1;
});
