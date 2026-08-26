import {
  SecretProvider,
  type ResolvedSecret,
} from "./secret-provider.js";

export interface RequiredSecret {
  name: string;
  defaultValue: string;
  version?: string;
}

export interface CacheEntry extends ResolvedSecret {
  loadedAt: Date;
}

export class SecretCache {
  private readonly entries = new Map<string, CacheEntry>();
  private readonly definitions = new Map<string, RequiredSecret>();

  public constructor(
    private readonly provider: SecretProvider,
    private readonly warningWindowMs: number,
  ) {
    if (warningWindowMs < 0) {
      throw new RangeError("warningWindowMs must be zero or greater");
    }
  }

  public async loadRequired(
    requiredSecrets: readonly RequiredSecret[],
  ): Promise<readonly CacheEntry[]> {
    const duplicate = findDuplicate(requiredSecrets.map(({ name }) => name));
    if (duplicate !== undefined) {
      throw new Error(`Duplicate required secret '${duplicate}'`);
    }

    const loaded = await Promise.all(
      requiredSecrets.map(async (definition) => ({
        definition,
        entry: await this.fetch(definition),
      })),
    );

    for (const { definition, entry } of loaded) {
      this.definitions.set(definition.name, definition);
      this.entries.set(definition.name, entry);
    }

    return loaded.map(({ entry }) => entry);
  }

  public getCached(name: string): CacheEntry {
    const entry = this.entries.get(name);
    if (entry === undefined) {
      throw new Error(`Secret '${name}' has not been loaded into the cache`);
    }
    return entry;
  }

  public async get(name: string, now = new Date()): Promise<CacheEntry> {
    const entry = this.getCached(name);
    if (this.provider.isNearExpiry(entry, this.warningWindowMs, now)) {
      return this.refresh(name);
    }
    return entry;
  }

  public async refresh(name: string): Promise<CacheEntry> {
    const definition = this.definitions.get(name);
    if (definition === undefined) {
      throw new Error(`No required-secret definition exists for '${name}'`);
    }

    const entry = await this.fetch(definition);
    this.entries.set(name, entry);
    return entry;
  }

  public findNearExpiry(now = new Date()): readonly CacheEntry[] {
    return [...this.entries.values()].filter((entry) =>
      this.provider.isNearExpiry(entry, this.warningWindowMs, now),
    );
  }

  public async refreshNearExpiry(
    now = new Date(),
  ): Promise<readonly CacheEntry[]> {
    return Promise.all(
      this.findNearExpiry(now).map(({ name }) => this.refresh(name)),
    );
  }

  private async fetch(definition: RequiredSecret): Promise<CacheEntry> {
    const secret = await this.provider.getSecret(
      definition.name,
      definition.defaultValue,
      definition.version,
    );
    return { ...secret, loadedAt: new Date() };
  }
}

function findDuplicate(values: readonly string[]): string | undefined {
  const seen = new Set<string>();
  for (const value of values) {
    if (seen.has(value)) {
      return value;
    }
    seen.add(value);
  }
  return undefined;
}
