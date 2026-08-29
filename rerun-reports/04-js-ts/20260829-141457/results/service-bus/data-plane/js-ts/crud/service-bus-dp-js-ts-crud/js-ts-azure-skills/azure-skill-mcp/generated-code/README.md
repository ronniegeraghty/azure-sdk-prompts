# Azure Service Bus TypeScript demo

This example uses `@azure/service-bus` to send and receive queue messages,
settle messages explicitly, process messages with `subscribe()`, and exchange a
message through a topic and subscription.

## Install and build

```powershell
npm install
npm run build
```

Create an Azure Service Bus queue, topic, and subscription separately. Copy
`.env.example` values into your shell environment; do not commit a real
connection string.

```powershell
$env:SERVICE_BUS_CONNECTION_STRING = "<connection-string>"
$env:SERVICE_BUS_QUEUE_NAME = "<queue-name>"
$env:SERVICE_BUS_TOPIC_NAME = "<topic-name>"
$env:SERVICE_BUS_SUBSCRIPTION_NAME = "<subscription-name>"
npm start
```

The connection string must grant send and listen permissions to the configured
entities. For production applications, prefer Microsoft Entra ID and managed
identity over connection-string credentials.

Reference: [Azure Service Bus client library for JavaScript](https://learn.microsoft.com/javascript/api/overview/azure/service-bus-readme?view=azure-node-latest)
