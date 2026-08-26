# Azure Service Bus TypeScript demo

Install the required package and development tooling:

```powershell
npm install
```

The runtime dependency is `@azure/service-bus`.

Set configuration for existing Service Bus entities. Keep the connection string
in an environment variable rather than source code:

```powershell
$env:AZURE_SERVICE_BUS_CONNECTION_STRING = "<service-bus-connection-string>"
$env:AZURE_SERVICE_BUS_QUEUE_NAME = "<queue-name>"
$env:AZURE_SERVICE_BUS_TOPIC_NAME = "<topic-name>"
$env:AZURE_SERVICE_BUS_SUBSCRIPTION_NAME = "<subscription-name>"
```

The queue, topic, and topic subscription must already exist. Build and run:

```powershell
npm run build
npm start
```

The program demonstrates single and batched queue sends, pull-based receiving
with explicit completion, event-driven subscription handlers, and topic
publishing with subscription receiving. It closes the subscription, receivers,
senders, and client in dependency order.

For production, prefer Microsoft Entra authentication with managed identity over
a connection string.
