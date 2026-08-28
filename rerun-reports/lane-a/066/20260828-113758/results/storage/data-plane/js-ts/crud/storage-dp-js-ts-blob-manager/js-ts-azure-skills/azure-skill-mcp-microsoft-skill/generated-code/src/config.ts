import { ManagedIdentityCredential } from "@azure/identity";
import { setLogLevel, type AzureLogLevel } from "@azure/logger";
import {
  BlobServiceClient,
  StorageRetryPolicyType,
} from "@azure/storage-blob";

const LOG_LEVELS = new Set<AzureLogLevel>([
  "verbose",
  "info",
  "warning",
  "error",
]);

function requiredEnvironmentVariable(name: string): string {
  const value = process.env[name]?.trim();
  if (!value) {
    throw new Error(`Required environment variable ${name} is not set.`);
  }
  return value;
}

function integerEnvironmentVariable(
  name: string,
  defaultValue: number,
  minimum: number,
): number {
  const rawValue = process.env[name]?.trim();
  if (!rawValue) {
    return defaultValue;
  }

  const value = Number(rawValue);
  if (!Number.isSafeInteger(value) || value < minimum) {
    throw new Error(`${name} must be an integer greater than or equal to ${minimum}.`);
  }
  return value;
}

function configureAzureLogging(): void {
  const rawLevel = process.env.AZURE_LOG_LEVEL?.trim().toLowerCase() ?? "warning";
  if (!LOG_LEVELS.has(rawLevel as AzureLogLevel)) {
    throw new Error(
      `AZURE_LOG_LEVEL must be one of: ${[...LOG_LEVELS].join(", ")}.`,
    );
  }
  setLogLevel(rawLevel as AzureLogLevel);
}

export interface StorageConfiguration {
  blobServiceClient: BlobServiceClient;
  containerName: string;
}

export function createStorageConfiguration(): StorageConfiguration {
  configureAzureLogging();

  const accountUrl = requiredEnvironmentVariable("AZURE_STORAGE_ACCOUNT_URL");
  const parsedUrl = new URL(accountUrl);
  if (parsedUrl.protocol !== "https:") {
    throw new Error("AZURE_STORAGE_ACCOUNT_URL must use HTTPS.");
  }

  const maxRetries = integerEnvironmentVariable(
    "AZURE_STORAGE_MAX_RETRIES",
    5,
    0,
  );
  const retryDelayInMs = integerEnvironmentVariable(
    "AZURE_STORAGE_RETRY_DELAY_MS",
    1_000,
    1,
  );
  const maxRetryDelayInMs = integerEnvironmentVariable(
    "AZURE_STORAGE_MAX_RETRY_DELAY_MS",
    30_000,
    retryDelayInMs,
  );

  const credential = new ManagedIdentityCredential();
  const blobServiceClient = new BlobServiceClient(accountUrl, credential, {
    retryOptions: {
      maxTries: maxRetries + 1,
      retryDelayInMs,
      maxRetryDelayInMs,
      retryPolicyType: StorageRetryPolicyType.EXPONENTIAL,
    },
  });

  return {
    blobServiceClient,
    containerName:
      process.env.AZURE_STORAGE_CONTAINER_NAME?.trim() || "blob-manager-demo",
  };
}
