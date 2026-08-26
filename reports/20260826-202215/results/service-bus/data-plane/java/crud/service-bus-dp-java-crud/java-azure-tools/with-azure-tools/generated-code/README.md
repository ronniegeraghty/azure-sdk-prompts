# Azure Service Bus Java messaging demo

This Maven sample demonstrates:

- creating queue and topic senders with `ServiceBusClientBuilder`;
- sending one message and a five-message `ServiceBusMessageBatch`;
- receiving with `receiveMessages()` and explicitly completing messages;
- continuous queue processing with `ServiceBusProcessorClient` handlers; and
- sending to a topic and receiving from a subscription.

The queue, topic, and subscription must already exist. The sample does not provision Azure resources.

## Configuration

The program uses `DefaultAzureCredential`, which supports local developer credentials and managed
identity without storing Service Bus keys in source code. Set these environment variables:

```powershell
$env:SERVICE_BUS_NAMESPACE = "<namespace>.servicebus.windows.net"
$env:SERVICE_BUS_QUEUE_NAME = "<queue-name>"
$env:SERVICE_BUS_TOPIC_NAME = "<topic-name>"
$env:SERVICE_BUS_SUBSCRIPTION_NAME = "<subscription-name>"
```

The signed-in identity needs the **Azure Service Bus Data Sender** and **Azure Service Bus Data
Receiver** roles for the demonstrated entities.

## Build and run

```powershell
mvn compile
mvn exec:java "-Dexec.mainClass=com.example.servicebus.ServiceBusMessagingDemo"
```

The required Service Bus Maven dependency is:

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-messaging-servicebus</artifactId>
    <version>7.17.17</version>
</dependency>
```

`azure-identity` is also included because the sample uses passwordless `DefaultAzureCredential`.

Reference: [Azure Service Bus client library for Java](https://learn.microsoft.com/java/api/overview/azure/messaging-servicebus-readme?view=azure-java-stable)
