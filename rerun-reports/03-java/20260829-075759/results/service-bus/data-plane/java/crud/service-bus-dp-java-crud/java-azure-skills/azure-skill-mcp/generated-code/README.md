# Azure Service Bus Java messaging demo

This Maven project demonstrates:

- queue sender creation with `ServiceBusClientBuilder`
- one `ServiceBusMessage` and a five-message `ServiceBusMessageBatch`
- synchronous `receiveMessages(...)` with explicit `complete(...)`
- continuous processing with `ServiceBusProcessorClient` handlers
- topic sending and subscription receiving
- deterministic cleanup with `close()`

The required Service Bus SDK dependency is:

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-messaging-servicebus</artifactId>
    <version>7.17.20</version>
</dependency>
```

The sample also uses `com.azure:azure-identity` so `DefaultAzureCredential`
can authenticate without embedding a connection string or key.

## Run

Create the queue, topic, and subscription beforehand, then grant the signed-in
identity the **Azure Service Bus Data Sender** and **Azure Service Bus Data
Receiver** roles for the required entities.

In PowerShell:

```powershell
$env:SERVICE_BUS_FQ_NAMESPACE = "<namespace>.servicebus.windows.net"
$env:SERVICE_BUS_QUEUE_NAME = "<queue>"
$env:SERVICE_BUS_TOPIC_NAME = "<topic>"
$env:SERVICE_BUS_SUBSCRIPTION_NAME = "<subscription>"
mvn compile exec:java
```

When the environment variables are absent, the program exits without making a
network connection.

## References

- [Send messages to an Azure Service Bus topic and receive messages from subscriptions (Java)](https://learn.microsoft.com/azure/service-bus-messaging/service-bus-java-how-to-use-topics-subscriptions)
- [Azure Service Bus client library for Java](https://learn.microsoft.com/java/api/overview/azure/messaging-servicebus-readme)
