import { DefaultAzureCredential } from "@azure/identity";
import { setLogLevel, type AzureLogLevel } from "@azure/logger";
import {
  BlobServiceClient,
  StorageRetryPolicyType,
} from "@azure/storage-blob";

export interface BlobStorageConfig {
  accountEndpoint: string;
  containerName: string;
  maxRetries: number;
  retryDelayInMs: number;
  maxRetryDelayInMs: number;
  logLevel: AzureLogLevel;
}

const LOG_LEVELS: readonly AzureLogLevel[] = [
  "verbose",
  "info",
  "warning",
  "error",
];

function readRequiredEnvironmentVariable(name: string): string {
  const value = process.env[name]?.trim();
  if (!value) {
    throw new Error(`Required environment variable ${name} is not set.`);
  }
  return value;
}

function readNonNegativeInteger(name: string, defaultValue: number): number {
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

function readLogLevel(): AzureLogLevel {
  const value = process.env.AZURE_STORAGE_LOG_LEVEL?.trim().toLowerCase();
  if (!value) {
    return "warning";
  }
  if (LOG_LEVELS.includes(value as AzureLogLevel)) {
    return value as AzureLogLevel;
  }
  throw new Error(
    `AZURE_STORAGE_LOG_LEVEL must be one of: ${LOG_LEVELS.join(", ")}.`,
  );
}

export function loadBlobStorageConfig(): BlobStorageConfig {
  const accountEndpoint = readRequiredEnvironmentVariable(
    "AZURE_STORAGE_ACCOUNT_ENDPOINT",
  );
  const endpoint = new URL(accountEndpoint);
  if (endpoint.protocol !== "https:") {
    throw new Error("AZURE_STORAGE_ACCOUNT_ENDPOINT must use HTTPS.");
  }

  return {
    accountEndpoint: endpoint.toString(),
    containerName: readRequiredEnvironmentVariable(
      "AZURE_STORAGE_CONTAINER_NAME",
    ),
    maxRetries: readNonNegativeInteger("AZURE_STORAGE_MAX_RETRIES", 5),
    retryDelayInMs: readNonNegativeInteger(
      "AZURE_STORAGE_RETRY_DELAY_MS",
      1_000,
    ),
    maxRetryDelayInMs: readNonNegativeInteger(
      "AZURE_STORAGE_MAX_RETRY_DELAY_MS",
      30_000,
    ),
    logLevel: readLogLevel(),
  };
}

export function createBlobServiceClient(
  config: BlobStorageConfig,
): BlobServiceClient {
  setLogLevel(config.logLevel);

  const credential = new DefaultAzureCredential();
  return new BlobServiceClient(config.accountEndpoint, credential, {
    retryOptions: {
      maxTries: config.maxRetries + 1,
      retryPolicyType: StorageRetryPolicyType.EXPONENTIAL,
      retryDelayInMs: config.retryDelayInMs,
      maxRetryDelayInMs: config.maxRetryDelayInMs,
    },
  });
}
