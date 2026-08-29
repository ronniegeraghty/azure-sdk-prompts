import { ResourceManagementClient } from "@azure/arm-resources";
import {
  AggregateAuthenticationError,
  AzureCliCredential,
  ChainedTokenCredential,
  CredentialUnavailableError,
  ManagedIdentityCredential,
} from "@azure/identity";

const ARM_SCOPE = "https://management.azure.com/.default";

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name];
  if (!value) {
    throw new Error(`Required environment variable ${name} is not set.`);
  }

  return value;
}

async function main(): Promise<void> {
  const subscriptionId = requireEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
  const userAssignedClientId = requireEnvironmentVariable(
    "AZURE_CLIENT_ID",
  );

  // With no client ID, ManagedIdentityCredential uses the system-assigned identity.
  const systemAssignedCredential = new ManagedIdentityCredential();

  // Supplying a client ID selects a user-assigned managed identity.
  const userAssignedCredential = new ManagedIdentityCredential({
    clientId: userAssignedClientId,
  });

  // Azure CLI is tried only after both managed identities are unavailable.
  const credential = new ChainedTokenCredential(
    systemAssignedCredential,
    userAssignedCredential,
    new AzureCliCredential(),
  );

  try {
    await systemAssignedCredential.getToken(ARM_SCOPE);
    console.log("System-assigned managed identity is available.");
  } catch (error: unknown) {
    if (error instanceof CredentialUnavailableError) {
      console.log(
        "System-assigned managed identity is unavailable; trying the configured user-assigned identity, then Azure CLI.",
      );
    } else {
      throw error;
    }
  }

  const client = new ResourceManagementClient(
    credential,
    subscriptionId,
  );

  console.log("Resource groups:");
  for await (const resourceGroup of client.resourceGroups.list()) {
    console.log(`- ${resourceGroup.name ?? "(unnamed)"}`);
  }
}

main().catch((error: unknown) => {
  if (error instanceof CredentialUnavailableError) {
    console.error(
      "Managed Identity is unavailable. Run this program in Azure or sign in locally with `az login`.",
    );
  } else if (error instanceof AggregateAuthenticationError) {
    console.error(
      "No credential in the chain could authenticate. Configure Managed Identity in Azure or sign in locally with `az login`.",
    );
    console.error(error.message);
  } else {
    console.error("The Azure SDK operation failed:", error);
  }

  process.exitCode = 1;
});
