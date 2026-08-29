import type { SecretClientLike } from "./types.js";

export interface RotationResult {
  name: string;
  previousVersion?: string;
  newVersion?: string;
  expiresOn: Date;
}

function isSecretNotFound(error: unknown): boolean {
  if (typeof error !== "object" || error === null) {
    return false;
  }

  const candidate = error as { code?: unknown; statusCode?: unknown };
  return candidate.statusCode === 404 || candidate.code === "SecretNotFound";
}

export class SecretRotationHelper {
  public constructor(
    private readonly client: SecretClientLike,
    private readonly now: () => Date = () => new Date(),
  ) {}

  public async rotate(
    name: string,
    newValue: string,
    expiresOn: Date,
  ): Promise<RotationResult> {
    if (Number.isNaN(expiresOn.getTime()) || expiresOn <= this.now()) {
      throw new RangeError("expiresOn must be a valid future date");
    }

    let previousVersion: string | undefined;
    try {
      const current = await this.client.getSecret(name);
      previousVersion = current.properties.version;
    } catch (error: unknown) {
      if (!isSecretNotFound(error)) {
        throw error;
      }
    }

    // setSecret creates a new version while retaining prior versions.
    const rotated = await this.client.setSecret(name, newValue, {
      enabled: true,
      expiresOn,
      tags: {
        rotationStatus: "active",
        rotatedOn: this.now().toISOString(),
      },
    });

    return {
      name,
      expiresOn,
      ...(previousVersion === undefined ? {} : { previousVersion }),
      ...(rotated.properties.version === undefined
        ? {}
        : { newVersion: rotated.properties.version }),
    };
  }

  public async deleteAndPurgeForNameReuse(
    name: string,
    confirmSecretName: string,
  ): Promise<void> {
    if (confirmSecretName !== name) {
      throw new Error(
        `Refusing to purge '${name}': confirmation must exactly match the secret name`,
      );
    }

    // Key Vault deletes the secret name and every version, not one old version.
    const deletePoller = await this.client.beginDeleteSecret(name);
    await deletePoller.pollUntilDone();
    await this.client.purgeDeletedSecret(name);
  }
}
