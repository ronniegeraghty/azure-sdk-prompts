export type OperationCategory = "storage" | "key-vault" | "cryptography";

export class EncryptedBlobError extends Error {
  public constructor(
    public readonly category: OperationCategory,
    public readonly operation: string,
    message: string,
    options?: ErrorOptions,
  ) {
    super(message, options);
    this.name = "EncryptedBlobError";
  }
}

interface AzureErrorShape {
  readonly code?: string;
  readonly statusCode?: number;
}

function readAzureErrorShape(error: unknown): AzureErrorShape {
  if (typeof error !== "object" || error === null) {
    return {};
  }

  const candidate = error as Record<string, unknown>;
  return {
    ...(typeof candidate.code === "string" ? { code: candidate.code } : {}),
    ...(typeof candidate.statusCode === "number"
      ? { statusCode: candidate.statusCode }
      : {}),
  };
}

export function describeAzureFailure(error: unknown): string {
  const { code, statusCode } = readAzureErrorShape(error);
  const details = [
    code ? `code ${code}` : undefined,
    statusCode ? `HTTP ${statusCode}` : undefined,
  ].filter((item): item is string => item !== undefined);

  if (error instanceof Error) {
    return details.length > 0
      ? `${error.message} (${details.join(", ")})`
      : error.message;
  }

  return details.length > 0 ? details.join(", ") : "Unknown Azure SDK error";
}
