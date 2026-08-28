import {
  type SecretLookup,
  type SecretValue,
} from "./secret-provider.js";

export interface SecretProvider {
  getSecret(
    name: string,
    defaultValue: string,
    version?: string,
  ): Promise<SecretValue>;
  isExpiringWithin(
    secret: Pick<SecretValue, "expiresOn">,
    warningWindowMs: number,
    now?: Date,
  ): boolean;
}

export interface CachedSecretProviderOptions {
  warningWindowMs?: number;
  now?: () => Date;
}

export class CachedSecretProvider {
  private readonly cache = new Map<string, SecretValue>();
  private readonly lookups = new Map<string, SecretLookup>();
  private readonly pendingRefreshes = new Map<string, Promise<SecretValue>>();
  private readonly warningWindowMs: number;
  private readonly now: () => Date;

  public constructor(
    private readonly provider: SecretProvider,
    options: CachedSecretProviderOptions = {},
  ) {
    this.warningWindowMs =
      options.warningWindowMs ?? 7 * 24 * 60 * 60 * 1_000;
    this.now = options.now ?? (() => new Date());
  }

  public async loadRequired(lookups: readonly SecretLookup[]): Promise<void> {
    for (const lookup of lookups) {
      this.lookups.set(lookup.name, lookup);
    }

    await Promise.all(lookups.map(({ name }) => this.refresh(name)));
  }

  public async get(name: string): Promise<string> {
    const cached = this.cache.get(name);

    if (cached === undefined) {
      return (await this.refresh(name)).value;
    }

    if (
      this.provider.isExpiringWithin(
        cached,
        this.warningWindowMs,
        this.now(),
      )
    ) {
      return (await this.refresh(name)).value;
    }

    return cached.value;
  }

  public async refresh(name: string): Promise<SecretValue> {
    const existingRefresh = this.pendingRefreshes.get(name);
    if (existingRefresh !== undefined) {
      return existingRefresh;
    }

    const lookup = this.lookups.get(name);
    if (lookup === undefined) {
      throw new Error(
        `No secret lookup is registered for "${name}". Load it at startup before refreshing it.`,
      );
    }

    const refresh = this.provider
      .getSecret(lookup.name, lookup.defaultValue, lookup.version)
      .then((secret) => {
        this.cache.set(name, secret);
        return secret;
      })
      .finally(() => {
        this.pendingRefreshes.delete(name);
      });

    this.pendingRefreshes.set(name, refresh);
    return refresh;
  }

  public getExpiringSecrets(): SecretValue[] {
    const now = this.now();

    return [...this.cache.values()].filter((secret) =>
      this.provider.isExpiringWithin(secret, this.warningWindowMs, now),
    );
  }

  public inspect(name: string): SecretValue | undefined {
    return this.cache.get(name);
  }
}
