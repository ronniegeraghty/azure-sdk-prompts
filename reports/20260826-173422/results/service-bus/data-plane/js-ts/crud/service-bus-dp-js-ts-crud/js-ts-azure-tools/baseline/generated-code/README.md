# Azure Service Bus TypeScript demo

This sample uses `@azure/service-bus` to send and receive queue messages, send
a five-message batch, process messages with both `receiveMessages()` and
`subscribe()`, and send to a topic for receipt through a subscription.

The queue, topic, and topic subscription must already exist. The sample does
not provision Azure resources.

```powershell
npm install
$env:AZURE_SERVICE_BUS_CONNECTION_STRING = "<service-bus-connection-string>"
$env:AZURE_SERVICE_BUS_QUEUE_NAME = "<queue-name>"
$env:AZURE_SERVICE_BUS_TOPIC_NAME = "<topic-name>"
$env:AZURE_SERVICE_BUS_SUBSCRIPTION_NAME = "<subscription-name>"
npm start
```

Compile without running the sample:

```powershell
npm run build
```
