import type { Connection, ModelDeployment } from "@azure/ai-projects";
import { AIProjectClient } from "@azure/ai-projects";
import { DefaultAzureCredential } from "@azure/identity";

const endpoint = requiredEnvironmentVariable("FOUNDRY_PROJECT_ENDPOINT");
const connectionName = requiredEnvironmentVariable("CONNECTION_NAME");
const deploymentName = requiredEnvironmentVariable("DEPLOYMENT_NAME");
const client = new AIProjectClient(endpoint, new DefaultAzureCredential());

console.log("Connections:");
for await (const connection of client.connections.list()) {
  printConnection(connection);
}

console.log("Selected connection:");
printConnection(await client.connections.get(connectionName));

console.log("Model deployments:");
for await (const deployment of client.deployments.list()) {
  if (
    deployment.type === "ModelDeployment"
    && "modelPublisher" in deployment
    && "modelName" in deployment
    && "modelVersion" in deployment
  ) {
    printDeployment(deployment);
  }
}

console.log("Selected model deployment:");
const selectedDeployment = await client.deployments.get(deploymentName);
if (
  selectedDeployment.type !== "ModelDeployment"
  || !("modelPublisher" in selectedDeployment)
  || !("modelName" in selectedDeployment)
  || !("modelVersion" in selectedDeployment)
) {
  throw new Error(`${deploymentName} is not a model deployment.`);
}
printDeployment(selectedDeployment);

function printConnection(connection: Connection): void {
  console.log(
    `${connection.name} | ${connection.type} | ${connection.target} | default=${connection.isDefault}`,
  );
}

function printDeployment(deployment: ModelDeployment): void {
  console.log(
    `${deployment.name} | ${deployment.modelPublisher} | ${deployment.modelName} | ${deployment.modelVersion}`,
  );
}

function requiredEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`${name} is required.`);
  }
  return value;
}
