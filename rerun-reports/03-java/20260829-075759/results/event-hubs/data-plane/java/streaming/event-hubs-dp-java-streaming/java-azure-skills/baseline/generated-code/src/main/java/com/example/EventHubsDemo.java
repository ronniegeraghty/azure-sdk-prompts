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
import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.BlobContainerClientBuilder;

import java.time.Duration;
import java.time.Instant;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;

public final class EventHubsDemo {
    private static final int EVENT_COUNT = 10;
    private static final CountDownLatch EVENTS_RECEIVED = new CountDownLatch(EVENT_COUNT);

    private EventHubsDemo() {
    }

    public static void main(String[] args) throws InterruptedException {
        String eventHubsConnectionString = requiredEnvironmentVariable(
            "EVENT_HUBS_CONNECTION_STRING");
        String eventHubName = requiredEnvironmentVariable("EVENT_HUB_NAME");
        String storageConnectionString = requiredEnvironmentVariable(
            "AZURE_STORAGE_CONNECTION_STRING");
        String blobContainerName = requiredEnvironmentVariable(
            "BLOB_CONTAINER_NAME");

        sendEvents(eventHubsConnectionString, eventHubName);

        BlobContainerAsyncClient blobContainerClient = new BlobContainerClientBuilder()
            .connectionString(storageConnectionString)
            .containerName(blobContainerName)
            .buildAsyncClient();

        EventProcessorClient processor = new EventProcessorClientBuilder()
            .connectionString(eventHubsConnectionString, eventHubName)
            .consumerGroup(EventHubClientBuilder.DEFAULT_CONSUMER_GROUP_NAME)
            .checkpointStore(new BlobCheckpointStore(blobContainerClient))
            .processEvent(EventHubsDemo::processEvent)
            .processError(EventHubsDemo::processError)
            .buildEventProcessorClient();

        try {
            processor.start();
            System.out.println("Waiting for events...");

            boolean receivedAllEvents = EVENTS_RECEIVED.await(
                Duration.ofSeconds(60).toSeconds(), TimeUnit.SECONDS);
            if (!receivedAllEvents) {
                System.err.println("Timed out before 10 events were received.");
            }
        } finally {
            processor.stop();
        }
    }

    private static void sendEvents(String connectionString, String eventHubName) {
        try (EventHubProducerClient producer = new EventHubClientBuilder()
            .connectionString(connectionString, eventHubName)
            .buildProducerClient()) {

            EventDataBatch batch = producer.createBatch();
            for (int i = 1; i <= EVENT_COUNT; i++) {
                EventData event = new EventData("Event body " + i);
                event.getProperties().put("eventNumber", i);
                event.getProperties().put("createdAt", Instant.now().toString());

                if (!batch.tryAdd(event)) {
                    throw new IllegalStateException(
                        "The 10 events do not fit in a single EventDataBatch.");
                }
            }

            producer.send(batch);
            System.out.println("Sent " + EVENT_COUNT + " events.");
        }
    }

    private static void processEvent(EventContext context) {
        EventData event = context.getEventData();
        System.out.printf(
            "Received partition=%s sequence=%d body=%s properties=%s%n",
            context.getPartitionContext().getPartitionId(),
            event.getSequenceNumber(),
            event.getBodyAsString(),
            event.getProperties());

        context.updateCheckpoint();
        EVENTS_RECEIVED.countDown();
    }

    private static void processError(ErrorContext context) {
        System.err.printf(
            "Error in partition %s: %s%n",
            context.getPartitionContext() == null
                ? "unknown"
                : context.getPartitionContext().getPartitionId(),
            context.getThrowable().getMessage());
        context.getThrowable().printStackTrace(System.err);
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
