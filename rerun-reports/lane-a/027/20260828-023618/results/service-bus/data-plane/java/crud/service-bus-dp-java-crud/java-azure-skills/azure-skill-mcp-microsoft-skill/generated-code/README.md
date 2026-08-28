# Azure Service Bus Java demo

This Maven project demonstrates queue sends, a five-message batch, synchronous
peek-lock receiving and completion, continuous processor handlers, and
topic/subscription messaging.

Set these environment variables to existing Azure Service Bus entities:

```powershell
$env:AZURE_SERVICE_BUS_CONNECTION_STRING = "<namespace-connection-string>"
$env:AZURE_SERVICE_BUS_QUEUE_NAME = "<queue-name>"
$env:AZURE_SERVICE_BUS_TOPIC_NAME = "<topic-name>"
$env:AZURE_SERVICE_BUS_SUBSCRIPTION_NAME = "<subscription-name>"
```

Run the sample:

```powershell
mvn compile exec:java
```

The required SDK dependency is:

```xml
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-messaging-servicebus</artifactId>
    <version>7.17.17</version>
</dependency>
```
