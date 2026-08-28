package com.example;

import com.azure.messaging.eventhubs.EventData;
import com.azure.messaging.eventhubs.EventDataBatch;
import com.azure.messaging.eventhubs.EventHubClientBuilder;
import com.azure.messaging.eventhubs.EventHubProducerClient;
import com.azure.messaging.eventhubs.EventProcessorClient;
import com.azure.messaging.eventhubs.EventProcessorClientBuilder;
import com.azure.messaging.eventhubs.checkpointstore.blob.BlobCheckpointStore;
import com.azure.messaging.eventhubs.models.EventPosition;
import com.azure.storage.blob.BlobContainerAsyncClient;
import com.azure.storage.blob.BlobContainerClientBuilder;

import java.time.Duration;
import java.util.UUID;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;

public final class EventHubsSendReceive {
    private static final int EVENT_COUNT = 10;
    private static final Duration RECEIVE_TIMEOUT = Duration.ofSeconds(60);

    private EventHubsSendReceive() {
    }

    public static void main(String[] args) throws InterruptedException {
        String eventHubsConnectionString = requiredEnvironmentVariable("EVENT_HUBS_CONNECTION_STRING");
        String eventHubName = requiredEnvironmentVariable("EVENT_HUB_NAME");
        String storageConnectionString = requiredEnvironmentVariable("AZURE_STORAGE_CONNECTION_STRING");
        String blobContainerName = requiredEnvironmentVariable("BLOB_CONTAINER_NAME");
        String runId = UUID.randomUUID().toString();

        EventHubProducerClient producer = new EventHubClientBuilder()
            .connectionString(eventHubsConnectionString, eventHubName)
            .buildProducerClient();

        EventProcessorClient processor = null;
        try {
            sendEvents(producer, runId);

            BlobContainerAsyncClient blobContainerClient = new BlobContainerClientBuilder()
                .connectionString(storageConnectionString)
                .containerName(blobContainerName)
                .buildAsyncClient();

            CountDownLatch receivedEvents = new CountDownLatch(EVENT_COUNT);
            processor = new EventProcessorClientBuilder()
                .connectionString(eventHubsConnectionString, eventHubName)
                .consumerGroup(EventHubClientBuilder.DEFAULT_CONSUMER_GROUP_NAME)
                .checkpointStore(new BlobCheckpointStore(blobContainerClient))
                .initialPartitionEventPosition(partitionId -> EventPosition.earliest())
                .processEvent(eventContext -> {
                    EventData event = eventContext.getEventData();
                    System.out.printf(
                        "Partition %s, sequence %d: %s%n",
                        eventContext.getPartitionContext().getPartitionId(),
                        event.getSequenceNumber(),
                        event.getBodyAsString());

                    // Checkpoint only after the event has been processed successfully.
                    eventContext.updateCheckpoint();

                    if (runId.equals(event.getProperties().get("runId"))) {
                        receivedEvents.countDown();
                    }
                })
                .processError(errorContext -> System.err.printf(
                    "Error in partition %s: %s%n",
                    errorContext.getPartitionContext() == null
                        ? "<not associated with a partition>"
                        : errorContext.getPartitionContext().getPartitionId(),
                    errorContext.getThrowable()))
                .buildEventProcessorClient();

            processor.start();
            System.out.println("Processor started; waiting for the sent events...");

            if (!receivedEvents.await(RECEIVE_TIMEOUT.toSeconds(), TimeUnit.SECONDS)) {
                throw new IllegalStateException(
                    "Timed out waiting for all events; remaining: " + receivedEvents.getCount());
            }
        } finally {
            if (processor != null) {
                processor.stop();
            }
            producer.close();
        }
    }

    private static void sendEvents(EventHubProducerClient producer, String runId) {
        EventDataBatch batch = producer.createBatch();

        for (int index = 1; index <= EVENT_COUNT; index++) {
            EventData event = new EventData("Event " + index);
            event.getProperties().put("runId", runId);
            event.getProperties().put("eventNumber", index);
            event.getProperties().put("source", "java-sample");

            if (!batch.tryAdd(event)) {
                throw new IllegalStateException("The 10 sample events do not fit in one EventDataBatch.");
            }
        }

        producer.send(batch);
        System.out.printf("Sent %d events with runId %s.%n", batch.getCount(), runId);
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Set the " + name + " environment variable.");
        }
        return value;
    }
}
