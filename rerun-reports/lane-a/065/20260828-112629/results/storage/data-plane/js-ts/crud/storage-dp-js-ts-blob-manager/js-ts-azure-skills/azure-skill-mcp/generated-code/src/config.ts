import {
  ManagedIdentityCredential,
  type TokenCredential,
} from "@azure/identity";
import { setLogLevel, type AzureLogLevel } from "@azure/logger";
import {
  BlobServiceClient,
  StorageRetryPolicyType,
  type StoragePipelineOptions,
} from "@azure/storage-blob";

export interface BlobStorageConfig {
  endpoint: string;
  maxRetries: number;
  retryDelayInMs: number;
  maxRetryDelayInMs: number;
  logLevel: AzureLogLevel | undefined;
  managedIdentityClientId?: string;
}

const LOG_LEVELS = new Set<AzureLogLevel>([
  "verbose",
  "info",
  "warning",
  "error",
]);

function readNonNegativeInteger(
  name: string,
  defaultValue: number,
): number {
  const rawValue = process.env[name];
  if (rawValue === undefined) {
    return defaultValue;
  }

  const value = Number(rawValue);
  if (!Number.isSafeInteger(value) || value < 0) {
    throw new Error(`${name} must be a non-negative integer.`);
  }

  return value;
}

function readLogLevel(): AzureLogLevel | undefined {
  const rawValue = process.env.AZURE_SDK_LOG_LEVEL?.toLowerCase();
  if (rawValue === undefined || rawValue === "off") {
    return undefined;
  }

  if (!LOG_LEVELS.has(rawValue as AzureLogLevel)) {
    throw new Error(
      "AZURE_SDK_LOG_LEVEL must be verbose, info, warning, error, or off.",
    );
  }

  return rawValue as AzureLogLevel;
}

export function loadBlobStorageConfig(): BlobStorageConfig {
  const endpoint = process.env.AZURE_STORAGE_BLOB_ENDPOINT;
  if (!endpoint) {
    throw new Error(
      "AZURE_STORAGE_BLOB_ENDPOINT is required (for example, https://<account>.blob.core.windows.net).",
    );
  }

  const parsedEndpoint = new URL(endpoint);
  if (parsedEndpoint.protocol !== "https:") {
    throw new Error("AZURE_STORAGE_BLOB_ENDPOINT must use HTTPS.");
  }

  const retryDelayInMs = readNonNegativeInteger(
    "AZURE_STORAGE_RETRY_DELAY_MS",
    800,
  );
  const maxRetryDelayInMs = readNonNegativeInteger(
    "AZURE_STORAGE_MAX_RETRY_DELAY_MS",
    30_000,
  );
  if (maxRetryDelayInMs < retryDelayInMs) {
    throw new Error(
      "AZURE_STORAGE_MAX_RETRY_DELAY_MS must be greater than or equal to AZURE_STORAGE_RETRY_DELAY_MS.",
    );
  }

  const managedIdentityClientId = process.env.AZURE_CLIENT_ID;

  return {
    endpoint: parsedEndpoint.toString(),
    maxRetries: readNonNegativeInteger("AZURE_STORAGE_MAX_RETRIES", 4),
    retryDelayInMs,
    maxRetryDelayInMs,
    logLevel: readLogLevel(),
    ...(managedIdentityClientId ? { managedIdentityClientId } : {}),
  };
}

export function createBlobServiceClient(
  config: BlobStorageConfig = loadBlobStorageConfig(),
): BlobServiceClient {
  setLogLevel(config.logLevel);

  const credential: TokenCredential = config.managedIdentityClientId
    ? new ManagedIdentityCredential(config.managedIdentityClientId)
    : new ManagedIdentityCredential();

  const options: StoragePipelineOptions = {
    retryOptions: {
      maxTries: config.maxRetries + 1,
      retryDelayInMs: config.retryDelayInMs,
      maxRetryDelayInMs: config.maxRetryDelayInMs,
      retryPolicyType: StorageRetryPolicyType.EXPONENTIAL,
    },
  };

  return new BlobServiceClient(config.endpoint, credential, options);
}
