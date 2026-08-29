import {
  KeyVaultSecretProvider,
  type SecretResult,
} from "./secret-provider.js";

export interface RequiredSecret {
  name: string;
  defaultValue?: string;
}

export interface CachedSecret extends SecretResult {
  fetchedAt: Date;
}

export class SecretCache {
  private readonly entries = new Map<string, CachedSecret>();

  constructor(
    private readonly provider: KeyVaultSecretProvider,
    private readonly warningWindowMs = 7 * 24 * 60 * 60 * 1_000,
  ) {
    if (warningWindowMs < 0) {
      throw new RangeError("warningWindowMs must be non-negative");
    }
  }

  async loadRequired(secrets: readonly RequiredSecret[]): Promise<void> {
    const loaded = await Promise.all(
      secrets.map(async ({ name, defaultValue = "" }) => {
        const secret = await this.provider.getSecret(name, defaultValue);
        return [name, this.toCached(secret)] as const;
      }),
    );

    for (const [name, secret] of loaded) {
      this.entries.set(name, secret);
    }
  }

  async get(name: string, defaultValue = ""): Promise<string> {
    const cached = this.entries.get(name);
    if (
      cached === undefined ||
      this.provider.isNearExpiry(cached, this.warningWindowMs)
    ) {
      return (await this.refresh(name, defaultValue)).value;
    }

    return cached.value;
  }

  async refresh(name: string, defaultValue = ""): Promise<CachedSecret> {
    const secret = this.toCached(
      await this.provider.getSecret(name, defaultValue),
    );
    this.entries.set(name, secret);
    return secret;
  }

  expiringSoon(now = new Date()): CachedSecret[] {
    return [...this.entries.values()].filter((secret) =>
      this.provider.isNearExpiry(secret, this.warningWindowMs, now),
    );
  }

  snapshot(): ReadonlyMap<string, CachedSecret> {
    return new Map(this.entries);
  }

  private toCached(secret: SecretResult): CachedSecret {
    return { ...secret, fetchedAt: new Date() };
  }
}
