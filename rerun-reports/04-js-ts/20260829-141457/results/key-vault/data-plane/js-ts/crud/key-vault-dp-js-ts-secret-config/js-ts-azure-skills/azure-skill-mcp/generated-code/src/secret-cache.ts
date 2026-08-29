import {
  KeyVaultSecretProvider,
  type ResolvedSecret
} from "./secret-provider.js";

export interface RequiredConfigKey {
  name: string;
  defaultValue: string;
}

export interface CachedSecret extends ResolvedSecret {
  cachedAt: Date;
  defaultValue: string;
}

export type AutoRefreshErrorHandler = (error: unknown) => void;

export class SecretCache {
  private readonly entries = new Map<string, CachedSecret>();
  private readonly inFlight = new Map<string, Promise<CachedSecret>>();

  public constructor(
    private readonly provider: KeyVaultSecretProvider,
    private readonly expiryWarningWindowMs: number
  ) {
    if (expiryWarningWindowMs < 0) {
      throw new Error("Expiry warning window must not be negative.");
    }
  }

  public async bulkLoad(keys: readonly RequiredConfigKey[]): Promise<ReadonlyMap<string, CachedSecret>> {
    const uniqueKeys = new Map(keys.map((key) => [key.name, key]));
    await Promise.all(
      [...uniqueKeys.values()].map((key) => this.fetchAndCache(key.name, key.defaultValue))
    );
    return this.snapshot();
  }

  public async get(name: string, defaultValue = ""): Promise<string> {
    const cached = this.entries.get(name);
    if (!cached) {
      return (await this.fetchAndCache(name, defaultValue)).value;
    }

    if (this.provider.isExpiringSoon(cached, this.expiryWarningWindowMs)) {
      return (await this.fetchAndCache(name, cached.defaultValue)).value;
    }

    return cached.value;
  }

  public async refresh(name: string): Promise<CachedSecret> {
    const cached = this.entries.get(name);
    if (!cached) {
      throw new Error(`Cannot refresh "${name}" before it has been loaded into the cache.`);
    }

    return this.fetchAndCache(name, cached.defaultValue);
  }

  public async refreshExpiring(): Promise<readonly CachedSecret[]> {
    const expiring = this.getExpiringSecrets();
    return Promise.all(
      expiring.map((entry) => this.fetchAndCache(entry.name, entry.defaultValue))
    );
  }

  public getExpiringSecrets(now = new Date()): readonly CachedSecret[] {
    return [...this.entries.values()].filter((entry) =>
      this.provider.isExpiringSoon(entry, this.expiryWarningWindowMs, now)
    );
  }

  public getEntry(name: string): CachedSecret | undefined {
    return this.entries.get(name);
  }

  public snapshot(): ReadonlyMap<string, CachedSecret> {
    return new Map(this.entries);
  }

  public startAutoRefresh(
    intervalMs: number,
    onError: AutoRefreshErrorHandler
  ): () => void {
    if (intervalMs <= 0) {
      throw new Error("Automatic refresh interval must be greater than zero.");
    }

    const timer = setInterval(() => {
      void this.refreshExpiring().catch(onError);
    }, intervalMs);
    timer.unref();

    return () => clearInterval(timer);
  }

  private fetchAndCache(name: string, defaultValue: string): Promise<CachedSecret> {
    const existingRequest = this.inFlight.get(name);
    if (existingRequest) {
      return existingRequest;
    }

    const request = this.provider
      .getSecret(name, { defaultValue })
      .then((secret): CachedSecret => {
        const entry = {
          ...secret,
          defaultValue,
          cachedAt: new Date()
        };
        this.entries.set(name, entry);
        return entry;
      })
      .finally(() => {
        this.inFlight.delete(name);
      });

    this.inFlight.set(name, request);
    return request;
  }
}
