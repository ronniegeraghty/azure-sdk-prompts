import { ManagedIdentityCredential } from "@azure/identity";
import {
  type AzureLogLevel,
  setLogLevel,
} from "@azure/logger";
import {
  BlobServiceClient,
  StorageRetryPolicyType,
} from "@azure/storage-blob";

export interface StorageConfig {
  accountEndpoint: string;
  containerName: string;
  maxRetries: number;
  retryDelayMs: number;
  maxRetryDelayMs: number;
  leaseWaitMs: number;
  leasePollMs: number;
  uploadBufferSize: number;
  uploadConcurrency: number;
  sdkLogLevel?: AzureLogLevel;
  managedIdentityClientId?: string;
}

const LOG_LEVELS = new Set<AzureLogLevel>([
  "error",
  "warning",
  "info",
  "verbose",
]);

function requiredEnvironmentVariable(name: string): string {
  const value = process.env[name]?.trim();
  if (!value) {
    throw new Error(`Environment variable ${name} is required.`);
  }
  return value;
}

function positiveInteger(name: string, fallback: number): number {
  const rawValue = process.env[name]?.trim();
  if (!rawValue) {
    return fallback;
  }

  const value = Number(rawValue);
  if (!Number.isSafeInteger(value) || value <= 0) {
    throw new Error(`${name} must be a positive integer.`);
  }
  return value;
}

function sdkLogLevel(): AzureLogLevel | undefined {
  const value = (process.env.AZURE_SDK_LOG_LEVEL ?? "warning").toLowerCase();
  if (value === "off") {
    return undefined;
  }
  if (!LOG_LEVELS.has(value as AzureLogLevel)) {
    throw new Error(
      "AZURE_SDK_LOG_LEVEL must be off, error, warning, info, or verbose.",
    );
  }
  return value as AzureLogLevel;
}

function storageEndpoint(): string {
  const value = requiredEnvironmentVariable("AZURE_STORAGE_ACCOUNT_ENDPOINT");
  let endpoint: URL;
  try {
    endpoint = new URL(value);
  } catch {
    throw new Error(
      "AZURE_STORAGE_ACCOUNT_ENDPOINT must be a valid HTTPS URL.",
    );
  }

  if (endpoint.protocol !== "https:") {
    throw new Error("AZURE_STORAGE_ACCOUNT_ENDPOINT must use HTTPS.");
  }
  return endpoint.toString().replace(/\/$/, "");
}

export function loadStorageConfig(): StorageConfig {
  return {
    accountEndpoint: storageEndpoint(),
    containerName:
      process.env.AZURE_STORAGE_CONTAINER_NAME?.trim() || "blob-manager-demo",
    maxRetries: positiveInteger("AZURE_STORAGE_MAX_RETRIES", 5),
    retryDelayMs: positiveInteger("AZURE_STORAGE_RETRY_DELAY_MS", 1_000),
    maxRetryDelayMs: positiveInteger(
      "AZURE_STORAGE_MAX_RETRY_DELAY_MS",
      30_000,
    ),
    leaseWaitMs: positiveInteger("AZURE_STORAGE_LEASE_WAIT_MS", 30_000),
    leasePollMs: positiveInteger("AZURE_STORAGE_LEASE_POLL_MS", 1_000),
    uploadBufferSize: positiveInteger(
      "AZURE_STORAGE_UPLOAD_BUFFER_SIZE",
      8 * 1024 * 1024,
    ),
    uploadConcurrency: positiveInteger(
      "AZURE_STORAGE_UPLOAD_CONCURRENCY",
      5,
    ),
    sdkLogLevel: sdkLogLevel(),
    managedIdentityClientId: process.env.AZURE_CLIENT_ID?.trim() || undefined,
  };
}

export function createBlobServiceClient(
  config: StorageConfig,
): BlobServiceClient {
  setLogLevel(config.sdkLogLevel);

  const credential = config.managedIdentityClientId
    ? new ManagedIdentityCredential({
        clientId: config.managedIdentityClientId,
      })
    : new ManagedIdentityCredential();

  return new BlobServiceClient(config.accountEndpoint, credential, {
    retryOptions: {
      maxTries: config.maxRetries,
      retryDelayInMs: config.retryDelayMs,
      maxRetryDelayInMs: config.maxRetryDelayMs,
      retryPolicyType: StorageRetryPolicyType.EXPONENTIAL,
    },
  });
}
