import type { SecretClientLike, SecretSnapshot, ExpiryInspection } from "./types.js";

const DAY_IN_MILLISECONDS = 24 * 60 * 60 * 1000;

function isSecretNotFound(error: unknown): boolean {
  if (typeof error !== "object" || error === null) {
    return false;
  }

  const candidate = error as { code?: unknown; statusCode?: unknown };
  return candidate.statusCode === 404 || candidate.code === "SecretNotFound";
}

export class KeyVaultSecretProvider {
  public constructor(
    private readonly client: SecretClientLike,
    private readonly now: () => Date = () => new Date(),
  ) {}

  public async getSecret(
    name: string,
    defaultValue: string,
    version?: string,
  ): Promise<SecretSnapshot> {
    try {
      const secret = await this.client.getSecret(
        name,
        version === undefined ? undefined : { version },
      );

      return {
        name,
        value: secret.value ?? defaultValue,
        found: secret.value !== undefined,
        ...(secret.properties.version === undefined
          ? {}
          : { version: secret.properties.version }),
        ...(secret.properties.expiresOn === undefined
          ? {}
          : { expiresOn: secret.properties.expiresOn }),
      };
    } catch (error: unknown) {
      if (isSecretNotFound(error)) {
        return { name, value: defaultValue, found: false };
      }

      throw error;
    }
  }

  public async inspectExpiry(
    name: string,
    warningWindowDays: number,
    version?: string,
  ): Promise<ExpiryInspection> {
    if (warningWindowDays < 0) {
      throw new RangeError("warningWindowDays must be zero or greater");
    }

    const secret = await this.getSecret(name, "", version);
    return this.inspectSnapshotExpiry(secret, warningWindowDays);
  }

  public inspectSnapshotExpiry(
    secret: SecretSnapshot,
    warningWindowDays: number,
  ): ExpiryInspection {
    if (warningWindowDays < 0) {
      throw new RangeError("warningWindowDays must be zero or greater");
    }

    if (secret.expiresOn === undefined) {
      return {
        name: secret.name,
        isExpired: false,
        isNearExpiry: false,
      };
    }

    const millisecondsRemaining =
      secret.expiresOn.getTime() - this.now().getTime();

    return {
      name: secret.name,
      expiresOn: secret.expiresOn,
      isExpired: millisecondsRemaining <= 0,
      isNearExpiry:
        millisecondsRemaining <= warningWindowDays * DAY_IN_MILLISECONDS,
      millisecondsRemaining,
    };
  }
}
