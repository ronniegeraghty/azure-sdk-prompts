import { DefaultAzureCredential } from "@azure/identity";
import { AzureLogLevel, setLogLevel } from "@azure/logger";
import {
  BlobServiceClient,
  StorageRetryPolicyType,
} from "@azure/storage-blob";

const LOG_LEVELS: readonly AzureLogLevel[] = [
  "verbose",
  "info",
  "warning",
  "error",
];

function requiredEnvironmentVariable(name: string): string {
  const value = process.env[name]?.trim();
  if (!value) {
    throw new Error(`The ${name} environment variable is required.`);
  }
  return value;
}

function nonNegativeIntegerEnvironmentVariable(
  name: string,
  defaultValue: number,
): number {
  const rawValue = process.env[name]?.trim();
  if (!rawValue) {
    return defaultValue;
  }

  const value = Number(rawValue);
  if (!Number.isSafeInteger(value) || value < 0) {
    throw new Error(`${name} must be a non-negative integer.`);
  }
  return value;
}

function configureSdkLogging(): void {
  const rawLevel = process.env.AZURE_SDK_LOG_LEVEL?.trim().toLowerCase();
  if (!rawLevel) {
    return;
  }

  if (!LOG_LEVELS.includes(rawLevel as AzureLogLevel)) {
    throw new Error(
      `AZURE_SDK_LOG_LEVEL must be one of: ${LOG_LEVELS.join(", ")}.`,
    );
  }
  setLogLevel(rawLevel as AzureLogLevel);
}

export interface BlobStorageConfiguration {
  blobServiceClient: BlobServiceClient;
  containerName: string;
}

export function createBlobStorageConfiguration(): BlobStorageConfiguration {
  configureSdkLogging();

  const accountEndpoint = requiredEnvironmentVariable(
    "AZURE_STORAGE_ACCOUNT_ENDPOINT",
  );
  const containerName = requiredEnvironmentVariable(
    "AZURE_STORAGE_CONTAINER_NAME",
  );

  const blobServiceClient = new BlobServiceClient(
    accountEndpoint,
    new DefaultAzureCredential(),
    {
      retryOptions: {
        retryPolicyType: StorageRetryPolicyType.EXPONENTIAL,
        // Azure's maxTries includes the initial request.
        maxTries:
          nonNegativeIntegerEnvironmentVariable(
            "AZURE_STORAGE_MAX_RETRIES",
            5,
          ) + 1,
        retryDelayInMs: nonNegativeIntegerEnvironmentVariable(
          "AZURE_STORAGE_RETRY_DELAY_MS",
          1_000,
        ),
        maxRetryDelayInMs: nonNegativeIntegerEnvironmentVariable(
          "AZURE_STORAGE_MAX_RETRY_DELAY_MS",
          30_000,
        ),
      },
    },
  );

  return { blobServiceClient, containerName };
}
