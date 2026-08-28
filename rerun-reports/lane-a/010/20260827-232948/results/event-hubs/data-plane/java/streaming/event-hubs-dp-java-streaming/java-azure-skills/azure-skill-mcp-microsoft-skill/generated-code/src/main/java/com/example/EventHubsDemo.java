package com.example;

import com.azure.messaging.eventhubs.EventData;
import com.azure.messaging.eventhubs.EventDataBatch;
import com.azure.messaging.eventhubs.EventHubClientBuilder;
import com.azure.messaging.eventhubs.EventHubProducerClient;
import com.azure.messaging.eventhubs.EventProcessorClient;
import com.azure.messaging.eventhubs.EventProcessorClientBuilder;
import com.azure.messaging.eventhubs.checkpointstore.blob.BlobCheckpointStore;
import com.azure.messaging.eventhubs.models.ErrorContext;
import com.azure.messaging.eventhubs.models.EventContext;
import com.azure.messaging.eventhubs.models.EventPosition;
import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.BlobContainerClientBuilder;

import java.time.Duration;
import java.util.UUID;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;
import java.util.function.Consumer;

public final class EventHubsDemo {
    private static final int EVENT_COUNT = 10;
    private static final Duration RECEIVE_TIMEOUT = Duration.ofSeconds(60);

    private EventHubsDemo() {
    }

    public static void main(String[] args) throws InterruptedException {
        String eventHubsConnectionString = requireEnvironmentVariable(
            "EVENT_HUBS_CONNECTION_STRING");
        String eventHubName = requireEnvironmentVariable("EVENT_HUB_NAME");
        String storageConnectionString = requireEnvironmentVariable(
            "AZURE_STORAGE_CONNECTION_STRING");
        String checkpointContainerName = requireEnvironmentVariable(
            "BLOB_CHECKPOINT_CONTAINER");
        String runId = UUID.randomUUID().toString();

        sendEvents(eventHubsConnectionString, eventHubName, runId);

        BlobContainerAsyncClient blobContainerClient =
            new BlobContainerClientBuilder()
                .connectionString(storageConnectionString)
                .containerName(checkpointContainerName)
                .buildAsyncClient();

        CountDownLatch receivedEvents = new CountDownLatch(EVENT_COUNT);

        Consumer<EventContext> processEvent = eventContext -> {
            EventData event = eventContext.getEventData();
            System.out.printf(
                "Received partition=%s, sequence=%d, body=%s, properties=%s%n",
                eventContext.getPartitionContext().getPartitionId(),
                event.getSequenceNumber(),
                event.getBodyAsString(),
                event.getProperties());

            if (runId.equals(event.getProperties().get("runId"))) {
                receivedEvents.countDown();
            }

            eventContext.updateCheckpoint();
        };

        Consumer<ErrorContext> processError = errorContext -> {
            String partitionId = errorContext.getPartitionContext() == null
                ? "N/A"
                : errorContext.getPartitionContext().getPartitionId();
            System.err.printf(
                "Error while processing partition %s: %s%n",
                partitionId,
                errorContext.getThrowable());
        };

        EventProcessorClient processor = new EventProcessorClientBuilder()
            .connectionString(eventHubsConnectionString, eventHubName)
            .consumerGroup(EventHubClientBuilder.DEFAULT_CONSUMER_GROUP_NAME)
            .checkpointStore(new BlobCheckpointStore(blobContainerClient))
            .initialPartitionEventPosition(partitionId -> EventPosition.earliest())
            .processEvent(processEvent)
            .processError(processError)
            .buildEventProcessorClient();

        try {
            processor.start();
            System.out.printf(
                "Processor started; waiting up to %d seconds for this run's events.%n",
                RECEIVE_TIMEOUT.toSeconds());

            if (!receivedEvents.await(RECEIVE_TIMEOUT.toSeconds(), TimeUnit.SECONDS)) {
                throw new IllegalStateException(
                    "Timed out waiting for all events; remaining=" + receivedEvents.getCount());
            }
        } finally {
            processor.stop();
            System.out.println("Processor stopped.");
        }
    }

    private static void sendEvents(
        String connectionString,
        String eventHubName,
        String runId) {

        try (EventHubProducerClient producer = new EventHubClientBuilder()
            .connectionString(connectionString, eventHubName)
            .buildProducerClient()) {

            EventDataBatch batch = producer.createBatch();

            for (int i = 1; i <= EVENT_COUNT; i++) {
                EventData event = new EventData("Demo event " + i);
                event.getProperties().put("eventNumber", i);
                event.getProperties().put("source", "java-event-hubs-demo");
                event.getProperties().put("runId", runId);

                if (!batch.tryAdd(event)) {
                    throw new IllegalStateException(
                        "The 10 demo events do not fit in one EventDataBatch.");
                }
            }

            producer.send(batch);
            System.out.printf("Sent %d events with runId=%s%n", batch.getCount(), runId);
        }
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(
                "Required environment variable is not set: " + name);
        }
        return value;
    }
}
