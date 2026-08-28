import {
  expiryStatusFromSecret,
  type SecretExpiryStatus,
  type SecretReader,
  type SecretResult,
} from "./secret-provider.js";

export interface RequiredConfigKey {
  name: string;
  defaultValue?: string;
}

export class CachedSecretProvider {
  private readonly cache = new Map<string, SecretResult>();
  private readonly defaultValues = new Map<string, string>();

  public constructor(
    private readonly provider: SecretReader,
    private readonly warningWindowDays = 7,
    private readonly now: () => Date = () => new Date(),
  ) {
    if (!Number.isFinite(warningWindowDays) || warningWindowDays < 0) {
      throw new Error("Expiry warning window must be a non-negative number of days.");
    }
  }

  public async loadRequired(keys: readonly RequiredConfigKey[]): Promise<Map<string, SecretResult>> {
    const loaded = await Promise.all(
      keys.map(async ({ name, defaultValue = "" }) => {
        this.defaultValues.set(name, defaultValue);
        return [name, await this.refresh(name, defaultValue)] as const;
      }),
    );

    return new Map(loaded);
  }

  public async get(name: string, defaultValue?: string): Promise<SecretResult> {
    if (defaultValue !== undefined) {
      this.defaultValues.set(name, defaultValue);
    }

    const cached = this.cache.get(name);
    const effectiveDefault = this.defaultValues.get(name) ?? defaultValue ?? "";

    if (cached === undefined || this.isNearExpiry(cached)) {
      return this.refresh(name, effectiveDefault);
    }

    return cached;
  }

  public peek(name: string): SecretResult | undefined {
    return this.cache.get(name);
  }

  public async refresh(name: string, defaultValue?: string): Promise<SecretResult> {
    const effectiveDefault = defaultValue ?? this.defaultValues.get(name) ?? "";
    this.defaultValues.set(name, effectiveDefault);

    const fresh = await this.provider.getSecret(name, effectiveDefault);
    this.cache.set(name, fresh);
    return fresh;
  }

  public async refreshExpiring(): Promise<readonly string[]> {
    const expiringNames = this.getExpiryStatuses()
      .filter(({ isNearExpiry }) => isNearExpiry)
      .map(({ name }) => name);

    await Promise.all(expiringNames.map((name) => this.refresh(name)));
    return expiringNames;
  }

  public getExpiryStatuses(): readonly SecretExpiryStatus[] {
    return [...this.cache.values()].map((secret) =>
      expiryStatusFromSecret(secret, this.warningWindowDays, this.now()),
    );
  }

  private isNearExpiry(secret: SecretResult): boolean {
    return expiryStatusFromSecret(secret, this.warningWindowDays, this.now()).isNearExpiry;
  }
}
