# Azure Service Bus TypeScript demo

This sample uses `@azure/service-bus` to send and receive queue and topic
messages. The queue, topic, and subscription must already exist.

## Install

```powershell
npm install
```

The required runtime package is:

```powershell
npm install @azure/service-bus
```

## Configure and run

Set the values from `.env.example` in your shell. Do not commit a real Service
Bus connection string.

```powershell
$env:SERVICEBUS_CONNECTION_STRING = "Endpoint=sb://..."
$env:SERVICEBUS_QUEUE_NAME = "my-queue"
$env:SERVICEBUS_TOPIC_NAME = "my-topic"
$env:SERVICEBUS_SUBSCRIPTION_NAME = "my-subscription"

npm run build
npm start
```

The connection string needs send and receive permissions for all configured
entities. For production applications, prefer Microsoft Entra authentication
with a managed identity instead of a connection string.

SDK references:

- https://learn.microsoft.com/javascript/api/overview/azure/service-bus-readme
- https://learn.microsoft.com/azure/service-bus-messaging/service-bus-nodejs-how-to-use-queues
