import { KeyVaultSecretProvider } from "./secret-provider.js";
import type { ExpiryInspection, SecretSnapshot } from "./types.js";

export class CachedSecretProvider {
  private readonly cache = new Map<string, SecretSnapshot>();
  private readonly defaults = new Map<string, string>();
  private readonly refreshes = new Map<string, Promise<SecretSnapshot>>();

  public constructor(
    private readonly provider: KeyVaultSecretProvider,
    private readonly warningWindowDays = 7,
  ) {
    if (warningWindowDays < 0) {
      throw new RangeError("warningWindowDays must be zero or greater");
    }
  }

  public async loadRequired(
    required: Readonly<Record<string, string>>,
  ): Promise<ReadonlyMap<string, SecretSnapshot>> {
    await Promise.all(
      Object.entries(required).map(([name, defaultValue]) =>
        this.refresh(name, defaultValue),
      ),
    );
    return new Map(this.cache);
  }

  public async get(
    name: string,
    defaultValue = "",
  ): Promise<SecretSnapshot> {
    const cached = this.cache.get(name);
    if (cached === undefined) {
      return this.refresh(name, defaultValue);
    }

    const expiry = this.provider.inspectSnapshotExpiry(
      cached,
      this.warningWindowDays,
    );
    if (expiry.isNearExpiry) {
      return this.refresh(name, this.defaults.get(name) ?? defaultValue);
    }

    return cached;
  }

  public async refresh(
    name: string,
    defaultValue = this.defaults.get(name) ?? "",
  ): Promise<SecretSnapshot> {
    const existingRefresh = this.refreshes.get(name);
    if (existingRefresh !== undefined) {
      return existingRefresh;
    }

    this.defaults.set(name, defaultValue);
    const refresh = this.provider
      .getSecret(name, defaultValue)
      .then((secret) => {
        this.cache.set(name, secret);
        return secret;
      })
      .finally(() => {
        this.refreshes.delete(name);
      });

    this.refreshes.set(name, refresh);
    return refresh;
  }

  public getExpiryWarnings(): ExpiryInspection[] {
    return [...this.cache.values()]
      .map((secret) =>
        this.provider.inspectSnapshotExpiry(secret, this.warningWindowDays),
      )
      .filter((inspection) => inspection.isNearExpiry);
  }

  public async refreshExpiringSecrets(): Promise<SecretSnapshot[]> {
    return Promise.all(
      this.getExpiryWarnings().map(({ name }) => this.refresh(name)),
    );
  }
}
