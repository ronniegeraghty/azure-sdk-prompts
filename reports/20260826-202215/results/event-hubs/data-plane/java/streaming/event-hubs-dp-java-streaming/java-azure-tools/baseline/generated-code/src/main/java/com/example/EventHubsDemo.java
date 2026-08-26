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

public final class EventHubsDemo {
    private static final int EVENT_COUNT = 10;
    private static final Duration RECEIVE_TIMEOUT = Duration.ofSeconds(60);

    private EventHubsDemo() {
    }

    public static void main(String[] args) throws InterruptedException {
        String eventHubsConnectionString = requireEnvironmentVariable("EVENT_HUB_CONNECTION_STRING");
        String storageConnectionString = requireEnvironmentVariable("AZURE_STORAGE_CONNECTION_STRING");
        String blobContainerName = requireEnvironmentVariable("BLOB_CONTAINER_NAME");
        String runId = UUID.randomUUID().toString();

        EventHubClientBuilder eventHubClientBuilder = new EventHubClientBuilder()
            .connectionString(eventHubsConnectionString);

        sendEvents(eventHubClientBuilder, runId);

        BlobContainerAsyncClient blobContainerClient = new BlobContainerClientBuilder()
            .connectionString(storageConnectionString)
            .containerName(blobContainerName)
            .buildAsyncClient();
        blobContainerClient.createIfNotExists().block();

        CountDownLatch receivedEvents = new CountDownLatch(EVENT_COUNT);
        EventProcessorClient processor = new EventProcessorClientBuilder()
            .connectionString(eventHubsConnectionString)
            .consumerGroup(EventHubClientBuilder.DEFAULT_CONSUMER_GROUP_NAME)
            .checkpointStore(new BlobCheckpointStore(blobContainerClient))
            .initialPartitionEventPosition(partitionId -> EventPosition.earliest())
            .processEvent(context -> processEvent(context, runId, receivedEvents))
            .processError(EventHubsDemo::processError)
            .buildEventProcessorClient();

        try {
            processor.start();
            boolean receivedAll = receivedEvents.await(RECEIVE_TIMEOUT.toSeconds(), TimeUnit.SECONDS);
            if (!receivedAll) {
                System.err.printf(
                    "Timed out after %d seconds; received %d of %d events from this run.%n",
                    RECEIVE_TIMEOUT.toSeconds(),
                    EVENT_COUNT - receivedEvents.getCount(),
                    EVENT_COUNT);
            }
        } finally {
            processor.stop();
        }
    }

    private static void sendEvents(EventHubClientBuilder clientBuilder, String runId) {
        try (EventHubProducerClient producer = clientBuilder.buildProducerClient()) {
            EventDataBatch batch = producer.createBatch();

            for (int i = 1; i <= EVENT_COUNT; i++) {
                EventData event = new EventData("Event " + i);
                event.getProperties().put("sequence-number", i);
                event.getProperties().put("sample-run-id", runId);
                event.getProperties().put("source", "java-event-hubs-demo");

                if (!batch.tryAdd(event)) {
                    throw new IllegalStateException("The 10-event sample batch exceeded the Event Hubs size limit.");
                }
            }

            producer.send(batch);
            System.out.printf("Sent %d events.%n", EVENT_COUNT);
        }
    }

    private static void processEvent(
        EventContext context,
        String runId,
        CountDownLatch receivedEvents) {

        EventData event = context.getEventData();
        System.out.printf(
            "Partition %s, sequence %d: %s%n",
            context.getPartitionContext().getPartitionId(),
            event.getSequenceNumber(),
            event.getBodyAsString());

        context.updateCheckpoint();

        if (runId.equals(event.getProperties().get("sample-run-id"))) {
            receivedEvents.countDown();
        }
    }

    private static void processError(ErrorContext context) {
        String partitionId = context.getPartitionContext() == null
            ? "unknown"
            : context.getPartitionContext().getPartitionId();
        System.err.printf(
            "Error while processing partition %s: %s%n",
            partitionId,
            context.getThrowable().getMessage());
        context.getThrowable().printStackTrace(System.err);
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is not set: " + name);
        }
        return value;
    }
}
