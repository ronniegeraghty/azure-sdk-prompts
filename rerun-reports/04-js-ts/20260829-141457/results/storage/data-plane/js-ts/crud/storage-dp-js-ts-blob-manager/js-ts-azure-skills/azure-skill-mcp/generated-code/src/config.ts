import {
  ManagedIdentityCredential,
  type TokenCredential,
} from "@azure/identity";
import {
  BlobServiceClient,
  StorageRetryPolicyType,
  type StoragePipelineOptions,
} from "@azure/storage-blob";
import {
  setLogLevel,
  type AzureLogLevel,
} from "@azure/logger";

export interface BlobStorageConfig {
  accountEndpoint: string;
  containerName: string;
  maxRetries: number;
  retryDelayInMs: number;
  maxRetryDelayInMs: number;
  logLevel: AzureLogLevel;
  managedIdentityClientId?: string;
}

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

function nonNegativeInteger(name: string, defaultValue: number): number {
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

function positiveInteger(name: string, defaultValue: number): number {
  const value = nonNegativeInteger(name, defaultValue);
  if (value === 0) {
    throw new Error(`${name} must be greater than zero.`);
  }
  return value;
}

function storageEndpoint(value: string): string {
  const endpoint = new URL(value);
  if (endpoint.protocol !== "https:") {
    throw new Error("AZURE_STORAGE_ACCOUNT_ENDPOINT must use HTTPS.");
  }
  if (endpoint.username || endpoint.password) {
    throw new Error("AZURE_STORAGE_ACCOUNT_ENDPOINT must not contain credentials.");
  }
  return endpoint.toString().replace(/\/$/, "");
}

function logLevel(): AzureLogLevel {
  const value = process.env.AZURE_SDK_LOG_LEVEL?.trim().toLowerCase() ?? "warning";
  if (!isAzureLogLevel(value)) {
    throw new Error(
      `AZURE_SDK_LOG_LEVEL must be one of: ${LOG_LEVELS.join(", ")}.`,
    );
  }
  return value;
}

function isAzureLogLevel(value: string): value is AzureLogLevel {
  return LOG_LEVELS.some((level) => level === value);
}

export function loadBlobStorageConfig(): BlobStorageConfig {
  const managedIdentityClientId = process.env.AZURE_CLIENT_ID?.trim();
  const retryDelayInMs = positiveInteger(
    "AZURE_STORAGE_RETRY_DELAY_MS",
    800,
  );
  const maxRetryDelayInMs = positiveInteger(
    "AZURE_STORAGE_MAX_RETRY_DELAY_MS",
    30_000,
  );
  if (maxRetryDelayInMs < retryDelayInMs) {
    throw new Error(
      "AZURE_STORAGE_MAX_RETRY_DELAY_MS must be greater than or equal to AZURE_STORAGE_RETRY_DELAY_MS.",
    );
  }

  return {
    accountEndpoint: storageEndpoint(
      requiredEnvironmentVariable("AZURE_STORAGE_ACCOUNT_ENDPOINT"),
    ),
    containerName:
      process.env.AZURE_STORAGE_CONTAINER_NAME?.trim() || "blob-manager-demo",
    maxRetries: nonNegativeInteger("AZURE_STORAGE_MAX_RETRIES", 5),
    retryDelayInMs,
    maxRetryDelayInMs,
    logLevel: logLevel(),
    ...(managedIdentityClientId ? { managedIdentityClientId } : {}),
  };
}

function createCredential(config: BlobStorageConfig): TokenCredential {
  return config.managedIdentityClientId
    ? new ManagedIdentityCredential({ clientId: config.managedIdentityClientId })
    : new ManagedIdentityCredential();
}

export function createBlobServiceClient(
  config: BlobStorageConfig,
): BlobServiceClient {
  setLogLevel(config.logLevel);

  const options: StoragePipelineOptions = {
    retryOptions: {
      retryPolicyType: StorageRetryPolicyType.EXPONENTIAL,
      // maxTries includes the initial request.
      maxTries: config.maxRetries + 1,
      retryDelayInMs: config.retryDelayInMs,
      maxRetryDelayInMs: config.maxRetryDelayInMs,
    },
  };

  return new BlobServiceClient(
    config.accountEndpoint,
    createCredential(config),
    options,
  );
}
