# Azure Service Bus Java demo

This example uses existing Azure Service Bus entities; it does not create or
deploy any Azure resources.

## Configuration

Set these environment variables:

```text
SERVICE_BUS_CONNECTION_STRING=<service-bus-connection-string>
SERVICE_BUS_QUEUE_NAME=<existing-queue-name>
SERVICE_BUS_TOPIC_NAME=<existing-topic-name>
SERVICE_BUS_SUBSCRIPTION_NAME=<existing-subscription-name>
```

The connection string must permit sending and receiving for the named entities.
The subscription must already belong to the configured topic.

## Run

```text
mvn compile exec:java
```

The program sends one queue message, sends a five-message batch, receives and
completes queue messages, runs a continuous processor, and demonstrates topic
and subscription messaging. Every sender, receiver, and processor client is
closed in a `finally` block.
