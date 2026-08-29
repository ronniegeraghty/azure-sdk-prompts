export class ConfigurationError extends Error {
  public constructor(message: string) {
    super(message);
    this.name = "ConfigurationError";
  }
}

export class KeyVaultOperationError extends Error {
  public constructor(
    operation: string,
    message: string,
    options?: ErrorOptions,
  ) {
    super(`Key Vault ${operation} failed: ${message}`, options);
    this.name = "KeyVaultOperationError";
  }
}

export class BlobOperationError extends Error {
  public constructor(
    operation: string,
    message: string,
    options?: ErrorOptions,
  ) {
    super(`Blob ${operation} failed: ${message}`, options);
    this.name = "BlobOperationError";
  }
}

export class EncryptionMetadataError extends Error {
  public constructor(message: string) {
    super(`Invalid encryption metadata: ${message}`);
    this.name = "EncryptionMetadataError";
  }
}

export function getErrorMessage(error: unknown): string {
  return error instanceof Error ? error.message : String(error);
}

export function getStatusCode(error: unknown): number | undefined {
  if (
    typeof error === "object" &&
    error !== null &&
    "statusCode" in error &&
    typeof error.statusCode === "number"
  ) {
    return error.statusCode;
  }

  return undefined;
}
