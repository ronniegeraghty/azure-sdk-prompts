import type { SecretLike, SecretStore } from "./secret-store.js";

export interface RotationResult {
  secret: SecretLike;
  previousVersion?: string;
}

export class SecretRotationHelper {
  public constructor(private readonly client: SecretStore) {}

  public async rotate(
    name: string,
    newValue: string,
    expiresOn: Date,
  ): Promise<RotationResult> {
    if (expiresOn.getTime() <= Date.now()) {
      throw new RangeError("The rotated secret expiry must be in the future");
    }

    const previous = await this.tryGetCurrent(name);
    const secret = await this.client.setSecret(name, newValue, {
      enabled: true,
      expiresOn,
      tags: { rotatedOn: new Date().toISOString() },
    });

    return {
      secret,
      ...(previous?.properties.version === undefined
        ? {}
        : { previousVersion: previous.properties.version }),
    };
  }

  public async deleteAndPurgeForNameReuse(
    name: string,
    confirmation: string,
  ): Promise<void> {
    if (confirmation !== name) {
      throw new Error(
        "Cleanup confirmation must exactly match the secret name",
      );
    }

    // Key Vault deletion is name-scoped and removes every version.
    const deletePoller = await this.client.beginDeleteSecret(name);
    await deletePoller.pollUntilDone();
    await this.client.purgeDeletedSecret(name);
  }

  private async tryGetCurrent(name: string): Promise<SecretLike | undefined> {
    try {
      return await this.client.getSecret(name);
    } catch (error: unknown) {
      if (isNotFound(error)) {
        return undefined;
      }
      throw error;
    }
  }
}

function isNotFound(error: unknown): boolean {
  if (typeof error !== "object" || error === null) {
    return false;
  }
  const candidate = error as { code?: unknown; statusCode?: unknown };
  return candidate.statusCode === 404 || candidate.code === "SecretNotFound";
}
