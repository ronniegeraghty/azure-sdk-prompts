import { ManagedIdentityCredential } from "@azure/identity";
import { setLogLevel, type AzureLogLevel } from "@azure/logger";
import {
  BlobServiceClient,
  StorageRetryPolicyType,
} from "@azure/storage-blob";

export interface StorageConfiguration {
  blobServiceClient: BlobServiceClient;
  containerName: string;
}

const LOG_LEVELS: readonly AzureLogLevel[] = [
  "verbose",
  "info",
  "warning",
  "error",
];

function requireEnvironmentVariable(name: string): string {
  const value = process.env[name]?.trim();
  if (!value) {
    throw new Error(`Required environment variable ${name} is not set.`);
  }
  return value;
}

function readPositiveInteger(name: string, defaultValue: number): number {
  const rawValue = process.env[name]?.trim();
  if (!rawValue) {
    return defaultValue;
  }

  const value = Number(rawValue);
  if (!Number.isSafeInteger(value) || value <= 0) {
    throw new Error(`${name} must be a positive integer.`);
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
  const value = (process.env.AZURE_LOG_LEVEL ?? "warning").toLowerCase();
  if (!LOG_LEVELS.includes(value as AzureLogLevel)) {
    throw new Error(
      `AZURE_LOG_LEVEL must be one of: ${LOG_LEVELS.join(", ")}.`,
    );
  }
  return value as AzureLogLevel;
}

function readSecureEndpoint(): string {
  const rawEndpoint = requireEnvironmentVariable("AZURE_STORAGE_BLOB_ENDPOINT");
  let endpoint: URL;

  try {
    endpoint = new URL(rawEndpoint);
  } catch {
    throw new Error("AZURE_STORAGE_BLOB_ENDPOINT must be a valid URL.");
  }

  if (endpoint.protocol !== "https:") {
    throw new Error("AZURE_STORAGE_BLOB_ENDPOINT must use HTTPS.");
  }
  if (endpoint.username || endpoint.password || endpoint.search || endpoint.hash) {
    throw new Error(
      "AZURE_STORAGE_BLOB_ENDPOINT must not contain credentials, query parameters, or fragments.",
    );
  }

  return endpoint.toString().replace(/\/$/, "");
}

export function createStorageConfiguration(): StorageConfiguration {
  const endpoint = readSecureEndpoint();
  const containerName = requireEnvironmentVariable("AZURE_STORAGE_CONTAINER");
  const maxRetries = readNonNegativeInteger("AZURE_STORAGE_MAX_RETRIES", 5);
  const retryDelayInMs = readPositiveInteger(
    "AZURE_STORAGE_RETRY_DELAY_MS",
    800,
  );
  const maxRetryDelayInMs = readPositiveInteger(
    "AZURE_STORAGE_MAX_RETRY_DELAY_MS",
    10_000,
  );

  setLogLevel(readLogLevel());

  const managedIdentityClientId = process.env.AZURE_CLIENT_ID?.trim();
  const credential = managedIdentityClientId
    ? new ManagedIdentityCredential({ clientId: managedIdentityClientId })
    : new ManagedIdentityCredential();

  const blobServiceClient = new BlobServiceClient(endpoint, credential, {
    retryOptions: {
      maxTries: maxRetries + 1,
      retryDelayInMs,
      maxRetryDelayInMs,
      retryPolicyType: StorageRetryPolicyType.EXPONENTIAL,
    },
  });

  return { blobServiceClient, containerName };
}
