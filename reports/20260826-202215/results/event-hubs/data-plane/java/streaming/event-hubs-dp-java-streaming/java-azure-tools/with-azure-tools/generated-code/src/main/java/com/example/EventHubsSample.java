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
import java.time.Instant;
import java.util.UUID;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;
import java.util.function.Consumer;

public final class EventHubsSample {
    private static final int EVENT_COUNT = 10;
    private static final Duration RECEIVE_TIMEOUT = Duration.ofMinutes(2);

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
        String sampleRunId = UUID.randomUUID().toString();

        sendEvents(eventHubsConnectionString, eventHubName, sampleRunId);

        CountDownLatch receivedSampleEvents = new CountDownLatch(EVENT_COUNT);
        BlobContainerAsyncClient blobContainer = new BlobContainerClientBuilder()
            .connectionString(storageConnectionString)
            .containerName(checkpointContainer)
            .buildAsyncClient();

        Consumer<EventContext> processEvent = eventContext -> {
            EventData event = eventContext.getEventData();
            String partitionId = eventContext.getPartitionContext().getPartitionId();

            System.out.printf(
                "Received partition=%s, sequence=%d, body=%s, properties=%s%n",
                partitionId,
                event.getSequenceNumber(),
                event.getBodyAsString(),
                event.getProperties());

            // Checkpoint only after the event has been processed successfully.
            eventContext.updateCheckpoint();

            if (sampleRunId.equals(event.getProperties().get("sampleRunId"))) {
                receivedSampleEvents.countDown();
            }
        };

        Consumer<ErrorContext> processError = errorContext -> {
            String partitionId = errorContext.getPartitionContext() == null
                ? "N/A"
                : errorContext.getPartitionContext().getPartitionId();
            System.err.printf(
                "Processor error on partition %s: %s%n",
                partitionId,
                errorContext.getThrowable());
        };

        EventProcessorClient processor = new EventProcessorClientBuilder()
            .connectionString(eventHubsConnectionString, eventHubName)
            .consumerGroup(EventHubClientBuilder.DEFAULT_CONSUMER_GROUP_NAME)
            .checkpointStore(new BlobCheckpointStore(blobContainer))
            .initialPartitionEventPosition(partitionId -> EventPosition.earliest())
            .processEvent(processEvent)
            .processError(processError)
            .buildEventProcessorClient();

        try {
            processor.start();
            System.out.printf(
                "Processor started; waiting up to %d seconds for this run's events.%n",
                RECEIVE_TIMEOUT.toSeconds());

            boolean receivedAll = receivedSampleEvents.await(
                RECEIVE_TIMEOUT.toMillis(), TimeUnit.MILLISECONDS);
            if (!receivedAll) {
                throw new IllegalStateException(
                    "Timed out waiting for all events; remaining="
                        + receivedSampleEvents.getCount());
            }
        } finally {
            processor.stop();
            System.out.println("Processor stopped.");
        }
    }

    private static void sendEvents(
        String connectionString,
        String eventHubName,
        String sampleRunId) {

        try (EventHubProducerClient producer = new EventHubClientBuilder()
            .connectionString(connectionString, eventHubName)
            .buildProducerClient()) {

            EventDataBatch batch = producer.createBatch();
            for (int i = 1; i <= EVENT_COUNT; i++) {
                EventData event = new EventData("Sample event " + i);
                event.getProperties().put("eventNumber", i);
                event.getProperties().put("source", "java-event-hubs-sample");
                event.getProperties().put("sampleRunId", sampleRunId);
                event.getProperties().put("createdAt", Instant.now().toString());

                if (!batch.tryAdd(event)) {
                    throw new IllegalStateException(
                        "Event " + i + " does not fit in the EventDataBatch.");
                }
            }

            producer.send(batch);
            System.out.printf("Sent %d events in one batch.%n", batch.getCount());
        }
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(
                "Required environment variable is not set: " + name);
        }
        return value;
    }
}
