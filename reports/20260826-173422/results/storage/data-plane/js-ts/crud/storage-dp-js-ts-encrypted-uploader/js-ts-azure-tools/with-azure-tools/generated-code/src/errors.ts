export type AzureService = "Azure Blob Storage" | "Azure Key Vault";

export class AzureOperationError extends Error {
  public readonly service: AzureService;
  public readonly operation: string;
  public readonly statusCode: number | undefined;
  public readonly code: string | undefined;

  public constructor(
    service: AzureService,
    operation: string,
    cause: unknown,
    detail?: string,
  ) {
    const azureError = getAzureErrorDetails(cause);
    const suffix = detail ?? azureError.message ?? "Unknown service error";
    super(`${service} ${operation} failed: ${suffix}`, { cause });
    this.name = "AzureOperationError";
    this.service = service;
    this.operation = operation;
    this.statusCode = azureError.statusCode;
    this.code = azureError.code;
  }
}

interface AzureErrorDetails {
  statusCode?: number;
  code?: string;
  message?: string;
}

function getAzureErrorDetails(error: unknown): AzureErrorDetails {
  if (typeof error !== "object" || error === null) {
    return {};
  }

  const candidate = error as Record<string, unknown>;
  return {
    ...(typeof candidate["statusCode"] === "number"
      ? { statusCode: candidate["statusCode"] }
      : {}),
    ...(typeof candidate["code"] === "string"
      ? { code: candidate["code"] }
      : {}),
    ...(typeof candidate["message"] === "string"
      ? { message: candidate["message"] }
      : {}),
  };
}
