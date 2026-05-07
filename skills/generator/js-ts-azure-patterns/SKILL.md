# Azure SDK for JavaScript/TypeScript — Generator Patterns

You are generating Azure SDK code for JavaScript or TypeScript. Follow these patterns to produce production-quality code that aligns with Azure SDK best practices.

## 1. SDK Logging with @azure/logger (CRITICAL — most commonly missed)

**Always** set up SDK logging using `@azure/logger`. This is the standard way to enable diagnostics in Azure SDK for JS/TS.

```typescript
import { setLogLevel } from "@azure/logger";

// Set log level based on environment or explicitly
setLogLevel("info"); // Options: "verbose", "info", "warning", "error"
```

- Import `setLogLevel` from `@azure/logger`
- Call it early in the program before creating any clients
- Use `"info"` for production, `"verbose"` for debugging
- Add `@azure/logger` to `package.json` dependencies

## 2. RestError Exception Handling (CRITICAL — 80% failure rate)

**Always** catch `RestError` with `statusCode` checks for Azure service calls. Do not use bare `try/catch` with generic `Error`.

```typescript
import { RestError } from "@azure/core-rest-pipeline";

try {
  const secret = await client.getSecret("my-secret");
} catch (e) {
  if (e instanceof RestError) {
    if (e.statusCode === 404) {
      console.log("Secret not found");
    } else if (e.statusCode === 409) {
      console.log("Conflict — resource already exists or is locked");
    } else {
      console.error(`Azure service error: ${e.statusCode} — ${e.message}`);
    }
  } else {
    throw e; // Re-throw unexpected errors
  }
}
```

Key patterns per service:
- **Key Vault**: 404 for not-found secrets/keys, check `e.statusCode`
- **Storage**: 404 for missing blobs, 409 for lease conflicts, 412 for ETag mismatches
- **Cosmos DB**: Check `e.statusCode` on operations — 404 for missing items, 409 for conflicts
- **Service Bus**: `ServiceBusError` has `code` property (e.g., `"MessageLockLost"`, `"SessionCannotBeLocked"`)

## 3. Authentication — Production-Aware Patterns

Use `DefaultAzureCredential` for development, but prefer specific credentials for production:

```typescript
import { DefaultAzureCredential, ManagedIdentityCredential } from "@azure/identity";

// Development — uses environment variables, CLI, VS Code, etc.
const credential = new DefaultAzureCredential();

// Production — prefer explicit credential for faster auth and clearer errors
const prodCredential = new ManagedIdentityCredential();
```

Rules:
- **Always** use `@azure/identity` — never hardcode connection strings, keys, or secrets
- Read endpoint URLs from environment variables, not hardcoded strings
- `DefaultAzureCredential` is fine for samples but note when `ManagedIdentityCredential` is better for production
- Add `@azure/identity` to `package.json` dependencies

## 4. Client Construction with Endpoint and Credential

Always construct clients with an endpoint URL and credential:

```typescript
import { SecretClient } from "@azure/keyvault-secrets";
import { DefaultAzureCredential } from "@azure/identity";

const vaultUrl = process.env.AZURE_KEYVAULT_URL || "https://my-vault.vault.azure.net";
const client = new SecretClient(vaultUrl, new DefaultAzureCredential());
```

- Read endpoints from environment variables
- Pass credential as second argument (not connection strings)
- Never hardcode vault URLs, storage account URLs, or other endpoints

## 5. Async/Await with for-await-of Pagination

Use `for await...of` for paginated operations — never collect all items into an array:

```typescript
// ✅ Correct — streams pages, constant memory
for await (const secret of client.listPropertiesOfSecrets()) {
  console.log(secret.name);
}

// ❌ Wrong — loads everything into memory
const allSecrets = [];
for await (const secret of client.listPropertiesOfSecrets()) {
  allSecrets.push(secret); // Defeats pagination purpose
}
```

## 6. Long-Running Operations (LROs)

Use the `begin*` + `pollUntilDone()` pattern for operations that take time:

```typescript
// Key Vault — delete and purge requires waiting
const poller = await client.beginDeleteSecret("my-secret");
const deletedSecret = await poller.pollUntilDone();
// Only AFTER poller completes:
await client.purgeDeletedSecret("my-secret");
```

- Methods starting with `begin` return a poller
- Always `await poller.pollUntilDone()` before proceeding
- Never assume deletion is instantaneous — always use the poller
- Never fire-and-forget a `beginDelete*` call

## 7. Service Bus Message Settlement (commonly missed)

When receiving messages, always settle them properly:

```typescript
const receiver = sbClient.createReceiver("my-queue");
const messages = await receiver.receiveMessages(10);

for (const msg of messages) {
  try {
    await processMessage(msg);
    await receiver.completeMessage(msg);      // Success — remove from queue
  } catch (err) {
    if (isTransient(err)) {
      await receiver.abandonMessage(msg);     // Retry later
    } else {
      await receiver.deadLetterMessage(msg, { // Permanent failure
        deadLetterReason: "ProcessingError",
        deadLetterErrorDescription: err.message,
      });
    }
  }
}
```

- Always call `completeMessage()`, `abandonMessage()`, or `deadLetterMessage()`
- For streaming: use `subscribe()` with `processMessage` and `processError` handlers
- Always call `close()` on sender, receiver, and client when done

## 8. package.json Dependencies

Always include a `package.json` with correct `@azure/` scoped packages:

```json
{
  "dependencies": {
    "@azure/identity": "^4.0.0",
    "@azure/keyvault-secrets": "^4.8.0",
    "@azure/logger": "^1.1.0",
    "@azure/core-rest-pipeline": "^1.16.0"
  }
}
```

Rules:
- All Azure SDK packages use the `@azure/` scope
- Never use deprecated packages (e.g., `azure-storage`, `azure-keyvault`)
- Always include `@azure/identity` and `@azure/logger`
- Include `@azure/core-rest-pipeline` if catching `RestError`

## 9. Cosmos DB Patterns

```typescript
import { CosmosClient } from "@azure/cosmos";

const client = new CosmosClient({ endpoint, key });
const { database } = await client.databases.createIfNotExists({ id: "mydb" });
const { container } = await database.containers.createIfNotExists({
  id: "mycontainer",
  partitionKey: { paths: ["/partitionKey"] },
});

// CRUD with partition key
const { resource } = await container.items.create({ id: "1", partitionKey: "pk1", data: "value" });
const { resource: item } = await container.item("1", "pk1").read();
await container.item("1", "pk1").replace({ ...item, data: "updated" });
await container.item("1", "pk1").delete();

// Query with SQL
const { resources } = await container.items
  .query({ query: "SELECT * FROM c WHERE c.partitionKey = @pk", parameters: [{ name: "@pk", value: "pk1" }] })
  .fetchAll();
```

- Always specify partition key in operations
- Use `createIfNotExists()` for database and container setup
- Use `SqlQuerySpec` with parameters for queries (not string interpolation)

## 10. Event Hubs Patterns

```typescript
import { EventHubProducerClient, EventHubConsumerClient } from "@azure/event-hubs";
import { ContainerClient } from "@azure/storage-blob";
import { BlobCheckpointStore } from "@azure/eventhubs-checkpointstore-blob";
import { DefaultAzureCredential } from "@azure/identity";

const credential = new DefaultAzureCredential();
const fullyQualifiedNamespace = process.env.AZURE_EVENTHUB_FULLY_QUALIFIED_NAMESPACE || "my-namespace.servicebus.windows.net";
const eventHubName = process.env.AZURE_EVENTHUB_NAME || "my-event-hub";

// Producer — use fully qualified namespace + credential (not connection strings)
const producer = new EventHubProducerClient(fullyQualifiedNamespace, eventHubName, credential);
const batch = await producer.createBatch();
batch.tryAdd({ body: "event data" });
await producer.sendBatch(batch);
await producer.close();

// Consumer with checkpoint store
const storageAccountUrl = process.env.AZURE_STORAGE_ACCOUNT_URL || "https://mystorage.blob.core.windows.net";
const containerClient = new ContainerClient(`${storageAccountUrl}/checkpoints`, credential);
const checkpointStore = new BlobCheckpointStore(containerClient);
const consumer = new EventHubConsumerClient("$Default", fullyQualifiedNamespace, eventHubName, credential, checkpointStore);
const subscription = consumer.subscribe({
  processEvents: async (events, context) => {
    for (const event of events) {
      console.log(event.body);
    }
    await context.updateCheckpoint(events[events.length - 1]);
  },
  processError: async (err, context) => {
    console.error(err);
  },
});
```

- Use fully qualified namespace + `DefaultAzureCredential` (never connection strings)
- Use `createBatch()` + `tryAdd()` for producing (not `send()` with raw arrays)
- Use `BlobCheckpointStore` for consumer checkpointing
- Always call `updateCheckpoint()` in `processEvents`
- Always call `close()` for cleanup
