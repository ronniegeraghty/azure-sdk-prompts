package com.example;

import com.azure.messaging.eventhubs.EventData;
import com.azure.messaging.eventhubs.EventDataBatch;
import com.azure.messaging.eventhubs.EventHubClientBuilder;
import com.azure.messaging.eventhubs.EventHubProducerClient;
import com.azure.messaging.eventhubs.checkpointstore.blob.BlobCheckpointStore;
import com.azure.messaging.eventhubs.models.EventContext;
import com.azure.messaging.eventhubs.models.EventPosition;
import com.azure.messaging.eventhubs.models.ErrorContext;
import com.azure.messaging.eventhubs.EventProcessorClient;
import com.azure.messaging.eventhubs.EventProcessorClientBuilder;
import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.BlobContainerClientBuilder;

import java.time.Duration;
import java.time.Instant;
import java.util.UUID;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;
import java.util.function.Consumer;

public final class EventHubsBatchProcessorDemo {
    private static final int EVENT_COUNT = 10;
    private static final Duration RECEIVE_TIMEOUT = Duration.ofSeconds(60);

    private EventHubsBatchProcessorDemo() {
    }

    public static void main(String[] args) throws InterruptedException {
        String eventHubConnectionString = requiredEnvironmentVariable("EVENT_HUB_CONNECTION_STRING");
        String storageConnectionString = requiredEnvironmentVariable("BLOB_STORAGE_CONNECTION_STRING");
        String blobContainerName = requiredEnvironmentVariable("BLOB_CONTAINER_NAME");
        String consumerGroup = System.getenv().getOrDefault(
            "EVENT_HUB_CONSUMER_GROUP",
            EventHubClientBuilder.DEFAULT_CONSUMER_GROUP_NAME);

        String runId = UUID.randomUUID().toString();
        Instant receiveFrom = Instant.now().minusSeconds(5);

        sendEvents(eventHubConnectionString, runId);

        BlobContainerAsyncClient blobContainerClient = new BlobContainerClientBuilder()
            .connectionString(storageConnectionString)
            .containerName(blobContainerName)
            .buildAsyncClient();

        CountDownLatch receivedEvents = new CountDownLatch(EVENT_COUNT);
        Consumer<EventContext> processEvent = eventContext -> {
            EventData event = eventContext.getEventData();
            String eventRunId = String.valueOf(event.getProperties().get("runId"));

            if (runId.equals(eventRunId)) {
                System.out.printf(
                    "Partition %s, sequence %d: %s%n",
                    eventContext.getPartitionContext().getPartitionId(),
                    event.getSequenceNumber(),
                    event.getBodyAsString());
                receivedEvents.countDown();
            }

            eventContext.updateCheckpoint();
        };

        Consumer<ErrorContext> processError = errorContext -> System.err.printf(
            "Error in partition %s: %s%n",
            errorContext.getPartitionContext() == null
                ? "unknown"
                : errorContext.getPartitionContext().getPartitionId(),
            errorContext.getThrowable());

        EventProcessorClient processor = new EventProcessorClientBuilder()
            .connectionString(eventHubConnectionString)
            .consumerGroup(consumerGroup)
            .checkpointStore(new BlobCheckpointStore(blobContainerClient))
            .initialPartitionEventPosition(partitionId -> EventPosition.fromEnqueuedTime(receiveFrom))
            .processEvent(processEvent)
            .processError(processError)
            .buildEventProcessorClient();

        try {
            processor.start();
            System.out.printf("Waiting up to %d seconds for this run's events...%n",
                RECEIVE_TIMEOUT.toSeconds());

            if (!receivedEvents.await(RECEIVE_TIMEOUT.toSeconds(), TimeUnit.SECONDS)) {
                throw new IllegalStateException(
                    "Timed out after receiving " + (EVENT_COUNT - receivedEvents.getCount())
                        + " of " + EVENT_COUNT + " events for run " + runId);
            }
        } finally {
            processor.stop();
        }
    }

    private static void sendEvents(String connectionString, String runId) {
        try (EventHubProducerClient producer = new EventHubClientBuilder()
            .connectionString(connectionString)
            .buildProducerClient()) {

            EventDataBatch batch = producer.createBatch();
            for (int i = 1; i <= EVENT_COUNT; i++) {
                EventData event = new EventData("Event " + i);
                event.getProperties().put("eventNumber", i);
                event.getProperties().put("runId", runId);
                event.getProperties().put("sentAt", Instant.now().toString());

                if (!batch.tryAdd(event)) {
                    throw new IllegalStateException(
                        "Event " + i + " does not fit in the batch; no partial batch was sent.");
                }
            }

            producer.send(batch);
            System.out.printf("Sent %d events for run %s.%n", EVENT_COUNT, runId);
        }
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is not set: " + name);
        }
        return value;
    }
}
