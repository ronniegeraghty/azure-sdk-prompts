interface AzureErrorShape {
  code?: unknown;
  statusCode?: unknown;
}

function isAzureErrorShape(value: unknown): value is AzureErrorShape {
  return typeof value === "object" && value !== null;
}

export function describeServiceError(error: unknown): string {
  const message = error instanceof Error ? error.message : String(error);

  if (!isAzureErrorShape(error)) {
    return message;
  }

  const details: string[] = [];
  if (typeof error.code === "string") {
    details.push(`code ${error.code}`);
  }
  if (typeof error.statusCode === "number") {
    details.push(`HTTP ${error.statusCode}`);
  }

  return details.length > 0 ? `${message} (${details.join(", ")})` : message;
}

export class KeyManagementError extends Error {
  constructor(operation: string, cause: unknown) {
    super(
      `Azure Key Vault ${operation} failed: ${describeServiceError(cause)}`,
      { cause },
    );
    this.name = "KeyManagementError";
  }
}

export class EncryptedBlobError extends Error {
  constructor(operation: string, cause: unknown) {
    super(
      `Azure Blob Storage ${operation} failed: ${describeServiceError(cause)}`,
      { cause },
    );
    this.name = "EncryptedBlobError";
  }
}
