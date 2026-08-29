# Azure Service Bus Java messaging demo

This Maven sample demonstrates:

- sending one queue message with `ServiceBusMessage`;
- sending five queue messages in a `ServiceBusMessageBatch`;
- receiving queue messages with `receiveMessages()` and completing them explicitly;
- continuous queue processing with `ServiceBusProcessorClient` handlers;
- sending to a topic and receiving from a subscription; and
- deterministic cleanup of senders, receivers, and processors with `close()`.

## Prerequisites

- Java 17 or later
- Maven 3.9 or later
- An existing Azure Service Bus namespace, queue, topic, and subscription
- A Microsoft Entra identity with **Azure Service Bus Data Sender** and
  **Azure Service Bus Data Receiver** access to the relevant entities

The demo uses `DefaultAzureCredential`; it does not store credentials or connection
strings in source code. For local development, authenticate with a supported developer
credential. In Azure, use a managed identity.

## Required Maven dependency

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-messaging-servicebus</artifactId>
    <version>7.17.20</version>
</dependency>
```

The project also includes `com.azure:azure-identity` for passwordless authentication.

## Configuration

Set these environment variables before running:

```powershell
$env:SERVICE_BUS_FQDN = "your-namespace.servicebus.windows.net"
$env:SERVICE_BUS_QUEUE_NAME = "your-queue"
$env:SERVICE_BUS_TOPIC_NAME = "your-topic"
$env:SERVICE_BUS_SUBSCRIPTION_NAME = "your-subscription"
```

The queue, topic, and subscription must already exist. This sample does not provision or
modify Azure resources.

## Run

```powershell
mvn compile exec:java
```

## References

- [Azure Service Bus client library for Java](https://learn.microsoft.com/java/api/overview/azure/messaging-servicebus-readme?view=azure-java-stable)
- [Send to and receive from Service Bus queues using Java](https://learn.microsoft.com/azure/service-bus-messaging/service-bus-java-how-to-use-queues)
- [Send to a topic and receive from a subscription using Java](https://learn.microsoft.com/azure/service-bus-messaging/service-bus-java-how-to-use-topics-subscriptions)
