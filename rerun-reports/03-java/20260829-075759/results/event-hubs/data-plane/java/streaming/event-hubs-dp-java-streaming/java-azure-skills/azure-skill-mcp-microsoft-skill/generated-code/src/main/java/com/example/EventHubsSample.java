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
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;
import java.util.function.Consumer;

public final class EventHubsSample {
    private static final int EVENT_COUNT = 10;
    private static final Duration RECEIVE_TIMEOUT = Duration.ofSeconds(60);

    private EventHubsSample() {
    }

    public static void main(String[] args) throws InterruptedException {
        String eventHubsConnectionString = requiredEnvironmentVariable(
            "EVENT_HUBS_CONNECTION_STRING");
        String eventHubName = requiredEnvironmentVariable("EVENT_HUB_NAME");
        String storageConnectionString = requiredEnvironmentVariable(
            "AZURE_STORAGE_CONNECTION_STRING");
        String checkpointContainer = requiredEnvironmentVariable(
            "BLOB_CHECKPOINT_CONTAINER");

        sendEvents(eventHubsConnectionString, eventHubName);

        CountDownLatch receivedEvents = new CountDownLatch(EVENT_COUNT);
        Consumer<EventContext> processEvent = eventContext -> {
            EventData event = eventContext.getEventData();
            System.out.printf(
                "Received partition=%s sequence=%d body=%s properties=%s%n",
                eventContext.getPartitionContext().getPartitionId(),
                event.getSequenceNumber(),
                event.getBodyAsString(),
                event.getProperties());

            // Checkpoint only after the event has been processed successfully.
            eventContext.updateCheckpoint();
            receivedEvents.countDown();
        };

        Consumer<ErrorContext> processError = errorContext -> {
            String partitionId = errorContext.getPartitionContext() == null
                ? "N/A"
                : errorContext.getPartitionContext().getPartitionId();
            System.err.printf("Error on partition %s: %s%n",
                partitionId, errorContext.getThrowable().getMessage());
        };

        BlobContainerAsyncClient blobContainerClient =
            new BlobContainerClientBuilder()
                .connectionString(storageConnectionString)
                .containerName(checkpointContainer)
                .buildAsyncClient();

        EventProcessorClient processor = new EventProcessorClientBuilder()
            .connectionString(eventHubsConnectionString, eventHubName)
            .consumerGroup(EventHubClientBuilder.DEFAULT_CONSUMER_GROUP_NAME)
            .checkpointStore(new BlobCheckpointStore(blobContainerClient))
            .initialPartitionEventPosition(
                partitionId -> EventPosition.earliest())
            .processEvent(processEvent)
            .processError(processError)
            .buildEventProcessorClient();

        try {
            processor.start();
            boolean receivedAll = receivedEvents.await(
                RECEIVE_TIMEOUT.toSeconds(), TimeUnit.SECONDS);
            if (!receivedAll) {
                System.err.printf(
                    "Timed out after %d seconds; received %d of %d expected events.%n",
                    RECEIVE_TIMEOUT.toSeconds(),
                    EVENT_COUNT - receivedEvents.getCount(),
                    EVENT_COUNT);
            }
        } finally {
            processor.stop();
        }
    }

    private static void sendEvents(
        String connectionString,
        String eventHubName) {

        try (EventHubProducerClient producer = new EventHubClientBuilder()
            .connectionString(connectionString, eventHubName)
            .buildProducerClient()) {

            EventDataBatch batch = producer.createBatch();
            for (int i = 1; i <= EVENT_COUNT; i++) {
                EventData event = new EventData("Sample event " + i);
                event.getProperties().put("eventNumber", i);
                event.getProperties().put("source", "java-sample");
                event.getProperties().put("category",
                    i % 2 == 0 ? "even" : "odd");

                if (!batch.tryAdd(event)) {
                    throw new IllegalStateException(
                        "The 10 sample events do not fit in one EventDataBatch.");
                }
            }

            producer.send(batch);
            System.out.printf("Sent %d events.%n", batch.getCount());
        }
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(
                "Required environment variable is not set: " + name);
        }
        return value;
    }
}
