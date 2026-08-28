import {
  expiresWithin,
  type KeyVaultSecretProvider,
  type SecretValue,
} from "./secret-provider.js";

export interface CachedSecret extends SecretValue {
  fetchedAt: Date;
}

export interface RequiredSecret {
  name: string;
  defaultValue: string;
}

export class CachingSecretProvider {
  private readonly cache = new Map<string, CachedSecret>();
  private readonly defaults = new Map<string, string>();

  public constructor(
    private readonly provider: KeyVaultSecretProvider,
    private readonly warningWindowMs: number,
  ) {
    if (!Number.isFinite(warningWindowMs) || warningWindowMs < 0) {
      throw new RangeError("warningWindowMs must be a non-negative finite number.");
    }
  }

  public async loadRequired(secrets: readonly RequiredSecret[]): Promise<ReadonlyMap<string, CachedSecret>> {
    await Promise.all(
      secrets.map(async ({ name, defaultValue }) => {
        this.defaults.set(name, defaultValue);
        await this.refresh(name);
      }),
    );

    return new Map(this.cache);
  }

  public async get(name: string, defaultValue = ""): Promise<CachedSecret> {
    if (!this.defaults.has(name)) {
      this.defaults.set(name, defaultValue);
    }

    const cached = this.cache.get(name);
    if (!cached || expiresWithin(cached.expiresOn, this.warningWindowMs)) {
      return this.refresh(name);
    }

    return cached;
  }

  public async refresh(name: string, defaultValue?: string): Promise<CachedSecret> {
    if (defaultValue !== undefined) {
      this.defaults.set(name, defaultValue);
    }

    const resolvedDefault = this.defaults.get(name) ?? "";
    const secret = await this.provider.getSecret(name, resolvedDefault);
    const cached = { ...secret, fetchedAt: new Date() };
    this.cache.set(name, cached);
    return cached;
  }

  public getNearExpiry(now = new Date()): readonly CachedSecret[] {
    return [...this.cache.values()].filter((secret) =>
      expiresWithin(secret.expiresOn, this.warningWindowMs, now),
    );
  }

  public async refreshNearExpiry(now = new Date()): Promise<readonly CachedSecret[]> {
    const names = this.getNearExpiry(now).map(({ name }) => name);
    return Promise.all(names.map((name) => this.refresh(name)));
  }

  public snapshot(): ReadonlyMap<string, CachedSecret> {
    return new Map(this.cache);
  }
}
