import {
  KeyVaultSecretProvider,
  type SecretResult,
} from "./secret-provider.js";

export type RequiredSecrets = Readonly<Record<string, string>>;

export class SecretCache {
  private readonly entries = new Map<string, SecretResult>();

  public constructor(
    private readonly provider: KeyVaultSecretProvider,
    private readonly defaults: RequiredSecrets,
    private readonly expiryWarningWindowMs: number,
  ) {
    if (expiryWarningWindowMs < 0) {
      throw new RangeError("expiryWarningWindowMs must not be negative.");
    }
  }

  public async loadRequired(): Promise<void> {
    await Promise.all(
      Object.keys(this.defaults).map(async (name) => {
        await this.refresh(name);
      }),
    );
  }

  public async get(name: string): Promise<string> {
    const cached = this.entries.get(name);
    if (cached === undefined) {
      return (await this.refresh(name)).value;
    }

    if (
      this.provider.expiresWithin(cached, this.expiryWarningWindowMs)
    ) {
      return (await this.refresh(name)).value;
    }

    return cached.value;
  }

  public inspect(name: string): SecretResult | undefined {
    return this.entries.get(name);
  }

  public inspectAll(): readonly SecretResult[] {
    return [...this.entries.values()];
  }

  public async refresh(name: string): Promise<SecretResult> {
    const defaultValue = this.defaults[name];
    if (defaultValue === undefined) {
      throw new Error(`No default value is configured for secret "${name}".`);
    }

    const result = await this.provider.getSecret(name, { defaultValue });
    this.entries.set(name, result);
    return result;
  }

  public async refreshExpiring(now = new Date()): Promise<readonly SecretResult[]> {
    const expiring = [...this.entries.values()].filter((entry) =>
      this.provider.expiresWithin(
        entry,
        this.expiryWarningWindowMs,
        now,
      ),
    );

    return Promise.all(expiring.map(({ name }) => this.refresh(name)));
  }

  public expiringSecrets(now = new Date()): readonly SecretResult[] {
    return [...this.entries.values()].filter((entry) =>
      this.provider.expiresWithin(
        entry,
        this.expiryWarningWindowMs,
        now,
      ),
    );
  }
}
